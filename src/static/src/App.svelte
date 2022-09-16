<Header appName="Seating Chart">
	<nav slot="headerCenter">
		{#if $userLoggedIn}
			<ul>
				{#each nav as navItem}
					<li>
						<Link href={navItem.href} classes={navItem.route === $activeRoute ? 'active' : ''}>
							<Icon icon={navItem.icon} />
							{navItem.text}
						</Link>
					</li>
				{/each}
			</ul>
		{/if}
	</nav>
</Header>

<Toasts />

<main class="f-column f-1">
	{#if $userLoggedIn}
		<Routing />
	{:else if $userLoggedIn === false}
		<Login />
	{:else}
		<Loading />
	{/if}
</main>

<script lang="ts">
	import { Header, Icon, Loading, Toasts } from 'sheodox-ui';
	import Routing from './Routing.svelte';
	import Link from './lib/Link.svelte';
	import Login from './Login.svelte';
	import { activeRoute } from './stores/routing';
	import { userLoggedIn } from './stores/app';

	const nav = [
		{
			route: 'chart',
			href: '/',
			icon: 'map',
			text: 'Chart',
		},
		{
			route: 'guests',
			href: '/guests',
			icon: 'users',
			text: 'Guests',
		},
	];
</script>
