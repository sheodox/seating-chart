<Header appName="Seating Chart">
	<nav slot="headerCenter">
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
	</nav>
</Header>

<Toasts />

<main class="f-column f-1">
	<Routing />
</main>

<script lang="ts">
	import { Header, Icon, Toasts } from 'sheodox-ui';
	import Routing from './Routing.svelte';
	import Link from './lib/Link.svelte';
	import { activeRoute } from './stores/routing';

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
