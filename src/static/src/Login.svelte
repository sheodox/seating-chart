<style>
	form {
		max-width: 90vw;
		width: 20rem;
	}
</style>

<section class="f-column f-1 justify-content-center align-items-center">
	<form class="f-column gap-3 m-3" on:submit|preventDefault={submit}>
		<TextInput bind:value={username} id="login-email" type="email">Email</TextInput>
		<TextInput bind:value={password} id="login-password" type="password">Password</TextInput>
		<button class="primary">Login</button>
	</form>
</section>

<script lang="ts">
	import { createAutoExpireToast, TextInput } from 'sheodox-ui';

	let username: string, password: string;

	async function submit() {
		const res = await fetch('/api/auth/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ username, password }),
		});

		if (res.status === 200) {
			location.reload();
		} else {
			createAutoExpireToast({
				title: 'Email or password incorrect.',
				message: '',
				variant: 'error',
			});
		}
	}
</script>
