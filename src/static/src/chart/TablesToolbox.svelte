<table>
	<thead>
		<tr>
			<th>Name</th>
			<th>Capacity</th>
		</tr>
	</thead>
	<tbody>
		{#each $tables as table}
			<tr>
				<td>{table.name}</td>
				<td>
					<button on:click={() => promptCapacity(table)}>
						{table.capacity}
					</button>
				</td>
			</tr>
		{/each}
	</tbody>
</table>

<script lang="ts">
	import { Table, tableOps, tables } from '../stores/tables';

	function promptCapacity(table: Table) {
		const capacity = parseInt(prompt('Enter a new capacity', '' + table.capacity), 10);

		if (capacity && capacity > 0) {
			tableOps.edit({
				...table,
				capacity,
			});
		}
	}
</script>
