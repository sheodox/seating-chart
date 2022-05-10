<h1 class="my-0">Guests</h1>
<AssignmentGuestList guests={unassigned} listTitle="Unassigned" showTable={false} />
<AssignmentGuestList guests={assigned} listTitle="Assigned" showTable />

<script lang="ts">
	import AssignmentGuestList from './AssignmentGuestList.svelte';
	import { guests } from '../stores/guests';

	$: unassigned = $guests.filter((g) => !g.tableId);
	$: assigned = $guests.filter((g) => !!g.tableId);
</script>
