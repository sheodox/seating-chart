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
				<td>{table.capacity}</td>
			</tr>
		{/each}
	</tbody>
</table>

<script lang="ts">
	import { tables } from '../stores/tables';
</script>
