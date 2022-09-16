{#if guests.length}
	<table>
		<caption class="shdx-font-size-5 fw-bold">{listTitle}</caption>
		<thead>
			<tr>
				<th>Guest</th>
				{#if showTable}
					<th>Table</th>
				{/if}
			</tr>
		</thead>
		<tbody>
			{#each guests as guest}
				<tr on:mouseenter={() => ($highlightTable = guest.tableId)} on:mouseleave={() => ($highlightTable = '')}>
					<td>
						<button draggable="true" on:dragstart={(e) => guestDragStart(e, guest)}>
							{guest.firstName}
							{guest.lastName}
							<span class="shdx-badge-pink fw-bold px-1">{guest.people} <Icon icon="users" /></span>
						</button>
					</td>
					{#if showTable}
						<td>
							{tableName(guest.tableId, $tables)}
							<button on:click={() => clearAssignment(guest)} title="Unassign from {tableName(guest.tableId, $tables)}">
								<Icon icon="unlink" variant="icon-only" />
							</button>
						</td>
					{/if}
				</tr>
			{/each}
		</tbody>
	</table>
{/if}

<script lang="ts">
	import { Icon } from 'sheodox-ui';
	import { highlightTable, tables } from '../stores/tables';
	import { guestDragStart, guestOps, sortGuests } from '../stores/guests';
	import type { Table } from '../stores/tables';
	import type { Guest } from '../stores/guests';

	export let guests: Guest[];
	export let listTitle: string;
	export let showTable: boolean;

	let sortedGuests = [];
	$: {
		sortedGuests = [...guests];
		sortGuests(sortedGuests);
	}

	function tableName(tableId: string, tables: Table[]) {
		return tables.find((table) => table.id === tableId)?.name || '';
	}

	function clearAssignment(guest: Guest) {
		guestOps.assign(guest.id, null);
	}
</script>
