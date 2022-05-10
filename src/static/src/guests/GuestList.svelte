<style>
	th:is(.first-name, .last-name) {
		width: 10rem;
	}
	th.people {
		width: 10rem;
	}
	th.actions {
		width: 15rem;
	}
	table {
		border: 1px solid var(--shdx-gray-300);
	}
	th {
		padding: var(--shdx-spacing-3);
		text-align: left;
	}
</style>

<section class="f-column align-items-center gap-5 my-5">
	<NewGuest />

	<table>
		<caption>Total {$guests.length} guests with {totalPeople} people.</caption>
		<thead>
			<tr>
				<th class="last-name">Last Name</th>
				<th class="first-name">First Name</th>
				<th class="people">People</th>
				<th class="actions">Actions</th>
			</tr>
		</thead>
		<tbody>
			{#each $guests as guest}
				<Guest {guest} />
			{/each}
		</tbody>
	</table>
</section>

<script lang="ts">
	import { guests } from '../stores/guests';
	import NewGuest from './NewGuest.svelte';
	import Guest from './Guest.svelte';

	$: totalPeople = $guests.reduce((people, guest) => people + guest.people, 0);
</script>
