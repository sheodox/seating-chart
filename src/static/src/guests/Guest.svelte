<style>
	input {
		width: 100%;
	}
	tr {
		border-bottom: 1px solid var(--shdx-gray-300);
	}
</style>

<tr>
	{#if editMode}
		<td>
			<input aria-label="first name" bind:value={firstName} />
		</td>
		<td>
			<input aria-label="last name" bind:value={lastName} />
		</td>
		<td>
			<input aria-label="people" bind:value={people} type="number" class="people" />
		</td>
		<td>
			<button on:click={save} disabled={!dirty} class="primary">Save</button>
			<button on:click={toggleMode}>Cancel</button>
			<button on:click={del} class="danger">Delete</button>
		</td>
	{:else}
		<td>{guest.firstName}</td>
		<td>{guest.lastName}</td>
		<td>{guest.people}</td>
		<td>
			<button on:click={toggleMode}>Edit</button>
		</td>
	{/if}
</tr>

<script lang="ts">
	import { send } from '../socket';
	import type { Guest } from '../stores/guests';
	export let guest: Guest;

	let editMode = false;

	$: dirty = editMode && (firstName !== guest.firstName || lastName !== guest.lastName || people !== guest.people);

	let firstName: string, lastName: string, people: number;

	function toggleMode() {
		editMode = !editMode;
		if (editMode) {
			firstName = guest.firstName;
			lastName = guest.lastName;
			people = guest.people;
		}
	}

	function save() {
		send('guests/edit', {
			id: guest.id,
			firstName,
			lastName,
			people,
		});
		toggleMode();
	}

	function del() {
		if (confirm(`Are you sure you want to delete "${guest.firstName} ${guest.lastName}"?`))
			send('guests/delete', { id: guest.id });
	}
</script>
