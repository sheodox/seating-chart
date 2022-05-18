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
		<div class="py-2">
			<Checkbox bind:checked={going} id="new-guest-going">Going</Checkbox>
		</div>
		<button class="primary">Add</button>
	</fieldset>
</form>

<script lang="ts">
	import { TextInput, Checkbox } from 'sheodox-ui';
	import { guestOps } from '../stores/guests';

	let people = 1,
		firstName = '',
		lastName = '',
		going = true;

	function add() {
		guestOps.add({ firstName, lastName, people, going });

		people = 1;
		lastName = '';
		firstName = '';
		going = true;
	}
</script>
