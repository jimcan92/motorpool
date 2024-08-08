<script lang="ts">
	import { page } from '$app/stores';
	import { loadAll } from '$lib/states';
	import { WindowSetDarkTheme, WindowSetLightTheme } from '$lib/wailsjs/runtime/runtime';
	import { Car, LayoutDashboard, ParkingMeter, Settings, Wrench } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import '../app.css';

	const { children } = $props();

	let light = $state(false);
	let open = $state(false);
	let pathnames = $derived($page.url.pathname.split('/'));
	const paths = $derived.by(() => {
		const h = $page.url.pathname.split('/').slice(1);
		h.pop();
		return h.map((p) => p[0].toUpperCase() + p.substring(1).toLowerCase());
	});
	const last = $derived.by(() => {
		const l = pathnames.at(-1);

		if (l && l !== '') return l[0].toUpperCase() + l.substring(1).toLowerCase();
		else return 'Dashboard';
	});
	// let paths = $derived(
	// 	pathnames.map((p) => {
	// 		return p[0].toUpperCase() + p.substring(1).toLowerCase();
	// 	})
	// );

	$effect(() => {
		if (light) WindowSetLightTheme();
		else WindowSetDarkTheme();
	});

	onMount(async () => await loadAll());

	function getPath(p: string) {
		const i = pathnames.indexOf(p.toLowerCase());
		const pn = pathnames.slice(0, i + 1);
		console.log(p, pn.join('/'));
		return pn.join('/');
	}
</script>

<div class="drawer lg:drawer-open h-svh">
	<input id="my-drawer-3" type="checkbox" class="drawer-toggle" bind:checked={open} />
	<div class="drawer-content flex h-svh flex-col">
		<!-- Navbar -->
		<div class="navbar bg-base-300 h-16 w-full">
			<div class="flex-none lg:hidden">
				<label for="my-drawer-3" aria-label="open sidebar" class="btn btn-square btn-ghost">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						class="inline-block h-6 w-6 stroke-current"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M4 6h16M4 12h16M4 18h16"
						></path>
					</svg>
				</label>
			</div>
			<div class="mx-2 flex-1 px-2">
				<div class="breadcrumbs">
					<ul>
						<li><a href="/">MOTOR POOL</a></li>
						{#each paths as path}
							<li><a href={getPath(path)}>{path}</a></li>
						{/each}
						<li><span>{last}</span></li>
					</ul>
				</div>
			</div>
		</div>
		<div class="h-[calc(100%-4rem)]">
			{@render children()}
		</div>
	</div>
	<div class="drawer-side">
		<label for="my-drawer-3" aria-label="close sidebar" class="drawer-overlay"></label>
		<div class="bg-base-300 z-50 flex h-svh w-64 flex-col justify-between overflow-y-auto p-2">
			<ul class="menu gap-2 p-2">
				<!-- Sidebar content here -->
				<li>
					<a href="/" class:active={$page.url.pathname === '/'} onclick={() => (open = false)}>
						<LayoutDashboard class="h-4 w-4" />
						Dashboard
					</a>
				</li>
				<li>
					<a
						href="/vehicles"
						class:active={$page.url.pathname.split('/').includes('vehicles')}
						onclick={() => (open = false)}
					>
						<Car class="h-4 w-4" />
						Vehicles
					</a>
				</li>
				<li>
					<a
						href="/usage"
						class:active={$page.url.pathname.split('/').includes('usage')}
						onclick={() => (open = false)}
					>
						<ParkingMeter class="h-4 w-4" />
						Dispatch
					</a>
				</li>
				<li>
					<a
						href="/maintenance"
						class:active={$page.url.pathname.split('/').includes('maintenance')}
						onclick={() => (open = false)}
					>
						<Wrench class="h-4 w-4" />
						Maintenance
					</a>
				</li>
			</ul>
			<ul class="menu border-base-content/30 border-t p-2">
				<li>
					<a
						href="/settings"
						class:active={$page.url.pathname.split('/').includes('settings')}
						onclick={() => (open = false)}
					>
						<Settings class="h-4 w-4" />
						Settings
					</a>
				</li>
			</ul>
		</div>
	</div>
</div>
