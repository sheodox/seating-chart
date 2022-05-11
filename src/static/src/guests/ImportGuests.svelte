<style>
	input {
		display: none;
	}
</style>

<div class="modal-body">
	<p class="text-align-center">
		<input type="file" id="upload-guest-list" on:change={uploadGuests} accept=".csv" />
		<label class="button" class:primary={!newGuests.length} for="upload-guest-list"
			><Icon icon="file-upload" />Upload</label
		>
	</p>

	{#if newGuests.length}
		<p>
			<Checkbox id="skip-first" bind:checked={skipFirst}>Skip First Row (use if the first row is a heading)</Checkbox>
		</p>

		<table>
			<thead>
				<tr>
					<th>Last Name</th>
					<th>First Name</th>
					<th>People</th>
				</tr>
			</thead>
			<tbody>
				{#each relevantGuests as guest}
					<tr>
						<td>{guest.lastName}</td>
						<td>{guest.firstName}</td>
						<td>{guest.people}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	{/if}
</div>

<div class="modal-footer">
	<button on:click={addGuests} class="primary" disabled={!relevantGuests.length}>Add</button>
</div>

<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { Icon, Checkbox } from 'sheodox-ui';
	import { guestOps } from '../stores/guests';

	const dispatch = createEventDispatcher<{ imported: void }>();

	let newGuests: { firstName: string; lastName: string; people: number }[] = [],
		skipFirst = false;

	$: relevantGuests = newGuests.slice(skipFirst ? 1 : 0);

	function addGuests() {
		for (const guest of relevantGuests) {
			guestOps.add(guest);
		}

		dispatch('imported');
	}

	function uploadGuests(e: Event) {
		const file = (e.target as HTMLInputElement).files[0],
			reader = new FileReader();

		reader.onload = (readEvent) => {
			parseGuests(readEvent.target.result.toString());
		};
		reader.readAsText(file);
	}

	function parseGuests(csv: string) {
		newGuests = csv.split(/\n/g).map((row) => {
			const [lastName, firstName, people] = row
				.trim()
				.split(',')
				.map((str) => str.trim());
			return { lastName, firstName, people: parseInt(people) };
		});
	}
</script>
