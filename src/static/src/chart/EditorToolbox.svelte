<style>
	aside {
		background-color: var(--sx-gray-400);
		width: 25rem;
		height: 90vh;
		overflow: auto;
		border-radius: 10px 0 0 10px;
	}
	aside :global(th) {
		text-align: left;
	}
</style>

<aside class="p-2 f-column">
	{#if $editorMode === 'lines'}
		<LinesToolbox />
	{:else if $editorMode === 'tables'}
		<TablesToolbox />
	{:else if $editorMode === 'assignments'}
		<AssignmentsToolbox />
	{/if}
</aside>

<script lang="ts">
	import { editorMode } from '../stores/editor';
	import LinesToolbox from './LinesToolbox.svelte';
	import TablesToolbox from './TablesToolbox.svelte';
	import AssignmentsToolbox from './AssignmentsToolbox.svelte';
</script>
