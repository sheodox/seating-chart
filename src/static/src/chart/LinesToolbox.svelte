<style>
	button[aria-pressed='true'] {
		background-color: var(--shdx-gray-300);
	}
</style>

<div class="f-row justify-content-between">
	<h1 class="my-0">Lines</h1>
	<button on:click={lineOps.add} class="primary"><Icon icon="plus" />New Line</button>
</div>

<table>
	<thead>
		<tr>
			<th>Action</th>
			<th>Closed</th>
		</tr>
	</thead>
	<tbody>
		{#each $lines as line, lineIndex}
			<tr
				on:mouseenter={() => ($highlightingLineIndex = lineIndex)}
				on:mouseleave={() => ($highlightingLineIndex = -1)}
			>
				<td>
					<button
						on:click={() => ($editingLineIndex = lineIndex)}
						class="primary"
						aria-pressed={$editingLineIndex === lineIndex}
					>
						<span>Draw</span>
					</button>
					<button on:click={() => deleteLine(line.id)}> Delete </button>
				</td>
				<td>
					<Checkbox id="line-closed-{line.id}" on:change={(e) => changeClosed(line, e)} checked={line.closed}
						>Closed</Checkbox
					>
				</td>
			</tr>
		{/each}
	</tbody>
</table>

<script lang="ts">
	import { Checkbox, Icon } from 'sheodox-ui';
	import { editingLineIndex, highlightingLineIndex, lineOps, lines, Line } from '../stores/lines';

	function deleteLine(id: string) {
		lineOps.delete(id);
	}

	function changeClosed(line: Line, e: Event) {
		lineOps.edit({
			...line,
			closed: (e.target as HTMLInputElement).checked,
		});
	}
</script>
