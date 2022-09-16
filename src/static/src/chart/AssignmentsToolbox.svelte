<h1 class="my-0">Guests</h1>
<div class="py-3">
	<Checkbox id="hide-guests" bind:checked={$showGuests}>Show guests on tables</Checkbox>
</div>

<div class="f-column gap-6">
	<TabList {tabs} bind:selectedTab />
	<Tab {selectedTab} tabId="unassigned">
		<AssignmentGuestList guests={unassigned} listTitle="Unassigned" showTable={false} />
	</Tab>
	<Tab {selectedTab} tabId="assigned">
		<AssignmentGuestList guests={assigned} listTitle="Assigned" showTable />
	</Tab>
</div>

<script lang="ts">
	import { TabList, Tab } from 'sheodox-ui';
	import AssignmentGuestList from './AssignmentGuestList.svelte';
	import { guests } from '../stores/guests';
	import { Checkbox } from 'sheodox-ui';
	import { showGuests } from '../stores/tables';

	let selectedTab = 'unassigned';

	const tabs = [
		{
			id: 'unassigned',
			title: 'Unassigned',
		},
		{ id: 'assigned', title: 'Assigned' },
	];

	$: unassigned = $guests.filter((g) => g.going && !g.tableId);
	$: assigned = $guests.filter((g) => g.going && !!g.tableId);
</script>
