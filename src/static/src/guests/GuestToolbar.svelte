<style>
	textarea {
		height: 10rem;
		width: 40rem;
	}
</style>

<div>
	<button on:click={() => (showImport = true)}>
		<Icon icon="file-import" />
		Import Guests
	</button>
	<button on:click={exportAssignments}>
		<Icon icon="file-export" />
		Export CSV
	</button>
	<button on:click={() => (showExportText = true)}>
		<Icon icon="copy" />
		Export Text
	</button>
</div>

{#if showImport}
	<Modal bind:visible={showImport} title="Import Guests">
		<ImportGuests on:imported={() => (showImport = false)} />
	</Modal>
{/if}

{#if showExportText}
	<Modal title="Export Text" bind:visible={showExportText}>
		<div class="modal-body f-column gap-3">
			<TextInput id="export-text-separator" bind:value={exportTextSeparator}>Name/table separator</TextInput>

			{#if numUnassignedGuests > 0}
				<p class="sx-badge-red">
					{#if numUnassignedGuests === 1}
						There is a guest without a table assignment!
					{:else}
						There are {numUnassignedGuests} guests without table assignments!
					{/if}
				</p>
			{/if}

			<fieldset>
				<legend>Tables to include</legend>
				{#each $tables as table}
					{@const included = tablesToInclude.includes(table.id)}
					<button
						on:click={() => toggleTableIncluded(table.id, included)}
						aria-pressed={included}
						class:primary={included}><span>{table.name}<span /></span></button
					>
				{/each}
			</fieldset>

			<label>
				Result
				<br />
				<textarea value={exportedText} />
			</label>

			<button on:click={downloadTextExport}><Icon icon="download" />Download</button>
		</div>
	</Modal>
{/if}

<script lang="ts">
	import { Icon, Modal, TextInput } from 'sheodox-ui';
	import { guestsAlphabetized } from '../stores/guests';
	import { tables } from '../stores/tables';
	import type { Guest } from '../stores/guests';
	import type { Table } from '../stores/tables';
	import ImportGuests from './ImportGuests.svelte';

	let showImport = false,
		showExportText = false,
		exportTextSeparator = ' - ',
		numUnassignedGuests = 0,
		tablesToInclude = $tables.map((t) => t.id);

	$: exportedText = genExportText(exportTextSeparator, $guestsAlphabetized, $tables, tablesToInclude);

	function toggleTableIncluded(id: string, included: boolean) {
		included ? tablesToInclude.splice(tablesToInclude.indexOf(id), 1) : tablesToInclude.push(id);

		tablesToInclude = tablesToInclude;
	}

	function genExportText(sep: string, guests: Guest[], tables: Table[], tablesToInclude: string[]) {
		const rows = [];
		numUnassignedGuests = 0;

		let lastFirstLetter = '';

		for (const guest of guests) {
			if (!guest.tableId && guest.going) {
				numUnassignedGuests += 1;
			}

			if (!tablesToInclude.includes(guest.tableId)) {
				continue;
			}

			const tableName = tables.find((table) => table.id === guest.tableId)?.name,
				lastNameFirstLetter = guest.lastName.at(0);

			if (lastFirstLetter && lastNameFirstLetter !== lastFirstLetter) {
				// add an empty blank line between blocks of names that start with the first letter, to make it easier to paste
				// all names that start with the same letter into groups alphabetically when printing
				rows.push('');
			}

			lastFirstLetter = lastNameFirstLetter;

			rows.push(`${guest.lastName}, ${guest.firstName}${sep}${tableName}`);
		}

		return rows.join('\n');
	}

	function download(fileName: string, mimeType: string, text: string) {
		const a = document.createElement('a');
		a.download = fileName;
		a.href = URL.createObjectURL(new Blob([text], { type: mimeType }));
		document.body.appendChild(a);
		a.click();
		a.remove();
	}

	function exportAssignments() {
		const rows = [['Name', 'Table']];
		for (const guest of $guestsAlphabetized) {
			const tableName = $tables.find((table) => table.id === guest.tableId)?.name;

			rows.push([`"${guest.lastName}, ${guest.firstName}"`, tableName]);
		}

		const csv = rows
			.map((row) => {
				return row.join(',');
			})
			.join('\n');

		download('guest-list.csv', 'text/csv', csv);
	}

	function downloadTextExport() {
		download('guest-list.txt', 'text/plain', exportedText);
	}
</script>
