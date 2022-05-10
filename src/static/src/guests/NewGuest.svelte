<style>
	form {
		width: 20rem;
	}
</style>

<form on:submit|preventDefault={add}>
	<fieldset class="f-column gap-2">
		<legend>New Guest</legend>
		<TextInput id="new-guest-fn" bind:value={firstName}>First Name</TextInput>
		<TextInput id="new-guest-ln" bind:value={lastName}>Last Name</TextInput>
		<TextInput id="new-guest-people" bind:value={people} type="number">People</TextInput>
		<button class="primary">Add</button>
	</fieldset>
</form>

<script lang="ts">
	import { TextInput } from 'sheodox-ui';
	import { send } from '../socket';

	let people = 1,
		firstName = '',
		lastName = '';

	function add() {
		send('guests/add', {
			firstName,
			lastName,
			people,
		});

		people = 1;
		lastName = '';
		firstName = '';
	}
</script>
