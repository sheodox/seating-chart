<div>
	<button on:click={() => (showImport = true)}>
		<Icon icon="file-import" />
		Import Guests
	</button>
	<button on:click={exportAssignments}>
		<Icon icon="file-export" />
		Export Assignments
	</button>
</div>

{#if showImport}
	<Modal bind:visible={showImport} title="Import Guests">
		<ImportGuests on:imported={() => (showImport = false)} />
	</Modal>
{/if}

<script lang="ts">
	import { Icon, Modal } from 'sheodox-ui';
	import { guests } from '../stores/guests';
	import { tables } from '../stores/tables';
	import ImportGuests from './ImportGuests.svelte';

	let showImport = false;

	function exportAssignments() {
		const rows = [['Name', 'Table']];
		for (const guest of $guests) {
			const tableName = $tables.find((table) => table.id === guest.tableId)?.name;

			rows.push([`"${guest.lastName}, ${guest.firstName}"`, tableName]);
		}

		const csv = rows
			.map((row) => {
				return row.join(',');
			})
			.join('\n');

		const a = document.createElement('a');
		a.download = 'guest-list.csv';
		a.href = URL.createObjectURL(new Blob([csv], { type: 'text/csv' }));
		document.body.appendChild(a);
		a.click();
		a.remove();
	}
</script>
