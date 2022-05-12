<style>
	.toolbar {
		align-self: center;
		background: var(--shdx-gray-500);
		border-radius: 0 0 6px 6px;
	}
	button {
		margin: 0;
	}
	button[aria-pressed='true'] {
		background-color: var(--shdx-gray-300);
	}
	#editors {
		flex: 1;
		position: relative;
		aspect-ratio: 1;
	}
</style>

<div class="toolbar f-row gap-1 p-1">
	{#each modes as m}
		<button aria-pressed={$editorMode === m.id} on:click={() => ($editorMode = m.id)}><span>{m.name}</span></button>
	{/each}
</div>

<div class="f-1 f-row">
	<div id="editors">
		<Lines />
		<Tables />
	</div>
	<EditorToolbox />
</div>

<script lang="ts">
	import Lines from './Lines.svelte';
	import Tables from './Tables.svelte';
	import EditorToolbox from './EditorToolbox.svelte';
	import { EditorMode, editorMode } from '../stores/editor';

	const modes: { id: EditorMode; name: string }[] = [
		{
			name: 'Table Assignments',
			id: 'assignments',
		},
		{
			name: 'Tables',
			id: 'tables',
		},
		{
			name: 'Lines',
			id: 'lines',
		},
	];
</script>
