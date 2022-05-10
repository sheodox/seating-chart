package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sheodox/seating-chart/controllers"
	"github.com/sheodox/seating-chart/interactors"
	"github.com/sheodox/seating-chart/repositories"
)

var (
	upgrader = getUpgrader()
	isDev    = os.Getenv("APP_ENV") == "development"
)

func getUpgrader() websocket.Upgrader {
	fmt.Printf("env %q\n", os.Getenv("APP_ENV"))

	if isDev {
		return websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
			return true
		}}
	}
	return websocket.Upgrader{}
}

type WSMessageReq struct {
	// a string mapping to some handler on a controller
	Route string `json:"route"`
	// the data relevant to the operation
	Data map[string]any `json:"data"`
}

type WSMessageRes struct {
	// a string mapping to some handler on a controller
	Route string `json:"route"`
	// the data relevant to the operation
	Data any `json:"data"`
}

type Server struct {
	Guests  controllers.GuestController
	Tables  controllers.TableController
	clients []*websocket.Conn
}

func NewServer() *Server {
	reps := repositories.NewRepositories()

	return &Server{
		Guests: controllers.GuestController{
			Interactor: interactors.GuestInteractor{
				Repo: reps.Guest,
			},
		},
		Tables: controllers.TableController{
			Interactor: interactors.TableInteractor{
				Repo: reps.Table,
			},
		},
		clients: make([]*websocket.Conn, 0),
	}
}

func (s *Server) Start() {
	e := echo.New()

	if isDev {
		// allow stuff like logout redirects to route back to the dev server
		e.GET("/", func(c echo.Context) error {
			return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
		})

		// CORS for the dev server
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowCredentials: true,
		}))
	}

	e.GET("/ws", s.connectClient)

	e.Logger.Fatal(e.Start(":5007"))
}

func (s *Server) connectClient(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return err
	}

	defer ws.Close()

	s.clients = append(s.clients, ws)
	defer s.removeClient(ws)

	for {
		_, msg, err := ws.ReadMessage()

		if err != nil {
			return err
		}

		err = s.handleWSMessage(ws, msg)

		if err != nil {
			log.Printf("Error handling message: %v\n", err)
			return err
		}
	}
}

func (s *Server) handleWSMessage(client *websocket.Conn, msg []byte) error {
	var m WSMessageReq
	err := json.Unmarshal(msg, &m)

	if err != nil {
		return err
	}

	switch m.Route {
	case "guests/add":
		guest, err := s.Guests.Add(m.Data)
		if err != nil {
			s.sendError(client, err)
			return nil
		}
		s.broadcast("guests/add", guest)
	case "guests/edit":
		guest, err := s.Guests.Edit(m.Data)
		if err != nil {
			s.sendError(client, err)
			return nil
		}
		s.broadcast("guests/edit", guest)
	case "guests/delete":
		id, err := s.Guests.Delete(m.Data)
		if err != nil {
			s.sendError(client, err)
			return nil
		}
		s.broadcast("guests/delete", id)
	case "guests/list":
		guests, err := s.Guests.List()

		if err != nil {
			s.sendError(client, err)
			return nil
		}

		client.WriteJSON(WSMessageRes{
			Route: "guests/list",
			Data:  guests,
		})
	case "guests/assign":
		guest, err := s.Guests.Assign(m.Data)

		if err != nil {
			s.sendError(client, err)
			return nil
		}

		s.broadcast("guests/edit", guest)
	case "tables/add":
		table, err := s.Tables.Add(m.Data)

		if err != nil {
			s.sendError(client, err)
			return nil
		}

		s.broadcast("tables/add", table)
	case "tables/list":
		tables, err := s.Tables.List()

		if err != nil {
			s.sendError(client, err)
			return nil
		}

		client.WriteJSON(WSMessageRes{
			Route: "tables/list",
			Data:  tables,
		})
	case "tables/edit":
		table, err := s.Tables.Edit(m.Data)
		if err != nil {
			s.sendError(client, err)
			return nil
		}
		s.broadcast("tables/edit", table)
	case "tables/delete":
		id, err := s.Tables.Delete(m.Data)
		if err != nil {
			s.sendError(client, err)
			return nil
		}
		s.broadcast("tables/delete", id)

		// guests might have been unassigned from a table, send the list again
		guests, err := s.Guests.List()

		if err != nil {
			s.sendError(client, err)
			return nil
		}

		s.broadcast("guests/list", guests)
	default:
		log.Printf("Missing socket handler for %q\n", m.Route)
		s.sendError(client, errors.New(fmt.Sprintf("missing handler for route %q", m.Route)))
	}

	return nil
}

func (s *Server) sendError(client *websocket.Conn, err error) {
	client.WriteJSON(WSMessageRes{
		Route: "error",
		Data:  err.Error(),
	})
}

func (s *Server) broadcast(route string, data any) {
	msg := WSMessageRes{
		Route: route,
		Data:  data,
	}

	for _, client := range s.clients {
		client.WriteJSON(msg)
	}
}

func (s *Server) removeClient(ws *websocket.Conn) {
	for i, conn := range s.clients {
		if conn == ws {
			// overwrite the client with the last one in the slice, shorten the slice to exclude the duplicate
			// https://stackoverflow.com/a/37335777/2675087
			s.clients[i] = s.clients[len(s.clients)-1]
			s.clients = s.clients[:len(s.clients)-1]
		}
	}
}
