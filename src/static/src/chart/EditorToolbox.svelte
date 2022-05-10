<style>
	aside {
		background-color: var(--shdx-gray-400);
	}
</style>

<aside class="p-2">
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
