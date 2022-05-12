<h1 class="my-0">Guests</h1>
<div class="f-column gap-6">
	<AssignmentGuestList guests={unassigned} listTitle="Unassigned" showTable={false} />
	<AssignmentGuestList guests={assigned} listTitle="Assigned" showTable />
</div>

<script lang="ts">
	import AssignmentGuestList from './AssignmentGuestList.svelte';
	import { guests } from '../stores/guests';

	$: unassigned = $guests.filter((g) => !g.tableId);
	$: assigned = $guests.filter((g) => !!g.tableId);
</script>
