package controllers

import (
	"crypto/subtle"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	session_cookie_name = "seating_chart_session"
)

var (
	valid_sessions = make(map[string]bool)
)

type AuthController struct {
	allowed_username string
	allowed_password string
}

func NewAuthController(username, password string) AuthController {
	return AuthController{username, password}
}

type statusResponse struct{}

func (a *AuthController) Status(c echo.Context) error {
	sess, err := session.Get("session", c)

	if err != nil {
		return err
	}

	val := sess.Values[session_cookie_name]

	if allowed, ok := val.(bool); ok && allowed {
		c.NoContent(http.StatusOK)
	} else {
		c.NoContent(http.StatusUnauthorized)
	}
	return nil
}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthController) Login(c echo.Context) error {
	authBody := new(authRequest)

	if err := c.Bind(authBody); err != nil {
		return err
	}

	// authenticate the user
	sess, err := session.Get("session", c)

	if subtle.ConstantTimeCompare([]byte(authBody.Username), []byte(a.allowed_username)) == 1 && subtle.ConstantTimeCompare([]byte(authBody.Password), []byte(a.allowed_password)) == 1 {
		if err != nil {
			return err
		}

		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			HttpOnly: true,
		}

		sess.Values[session_cookie_name] = true
		err = sess.Save(c.Request(), c.Response())

		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	} else {
		sess.Values[session_cookie_name] = false
		err = sess.Save(c.Request(), c.Response())

		if err != nil {
			return err
		}

		return c.NoContent(http.StatusUnauthorized)
	}
}

func (a *AuthController) RequireUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		not_allowed := func() error {
			return c.NoContent(http.StatusUnauthorized)
		}

		sess, err := session.Get("session", c)

		if err != nil {
			return err
		}

		sess_val := sess.Values[session_cookie_name]
		allowed, ok := sess_val.(bool)

		if !ok || !allowed {
			return not_allowed()
		}

		return next(c)
	}
}

func (a *AuthController) Logout(c echo.Context) error {
	sess, err := session.Get("session", c)

	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30,
		HttpOnly: true,
	}

	sess.Values[session_cookie_name] = nil
	err = sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
