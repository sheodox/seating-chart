{#if $activeRoute === 'guests'}
	<GuestList />
{:else}
	<Editor />
{/if}

<script lang="ts">
	import { activeRoute } from './stores/routing';
	import GuestList from './guests/GuestList.svelte';
	import Editor from './chart/Editor.svelte';

	$: console.log({ route: $activeRoute });
</script>
