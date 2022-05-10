{#if guests.length}
	<h2>{listTitle}</h2>
	<table>
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
						<button draggable="true" on:dragstart={(e) => dragStart(e, guest)}>
							{guest.firstName}
							{guest.lastName}
							<span class="shdx-badge-pink fw-bold px-1">{guest.people} <Icon icon="users" /></span>
						</button>
					</td>
					{#if showTable}
						<td>
							{tableName(guest.tableId)}
							<button on:click={() => clearAssignment(guest)} title="Unassign from table">
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
	import { draggingGuest, Guest, guestOps } from '../stores/guests';

	export let guests: Guest[];
	export let listTitle: string;
	export let showTable: boolean;

	function dragStart(e: DragEvent, guest: Guest) {
		e.dataTransfer.setData('guestId', guest.id);
		$draggingGuest = guest;

		window.addEventListener(
			'dragend',
			() => {
				$draggingGuest = null;
			},
			{ once: true }
		);
	}

	function tableName(tableId: string) {
		return $tables.find((table) => table.id === tableId)?.name;
	}

	function clearAssignment(guest: Guest) {
		guestOps.assign(guest.id, null);
	}
</script>
