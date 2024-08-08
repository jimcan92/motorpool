<script lang="ts">
	import { makers } from '$lib/states/makers.svelte';
	import { main } from '$lib/wailsjs/go/models';
	import { Plus } from 'lucide-svelte';
	import MakerDialog from './MakerDialog.svelte';

	let open = $state(false);
	let query = $state('');
	let makerToUpdate = $state<main.Maker>();

	const filteredMakers = $derived.by(() => {
		let all = makers.all;
		if (query !== '')
			all = makers.all.filter((m) => m.name.toLowerCase().includes(query.toLowerCase()));
		return all;
	});

	$effect(() => {
		if (makerToUpdate && !open) makerToUpdate = undefined;
	});
</script>

<div class="h-full w-full">
	<!-- <div class="bg-base-2000 border-base-content/30 h-full rounded-lg border p-4"> -->
	<div class="flex h-20 items-end justify-between">
		<!-- <div class="flex items-center gap-4"> -->
		<div class="form-control">
			<label for="search" class="label label-text">Search Manufacturer's Name</label>
			<input type="text" id="search" class="input input-bordered" bind:value={query} />
		</div>
		<!-- </div> -->
		<button class="btn btn-outline" onclick={() => (open = true)}>
			<Plus class="h-5 w-5" />
			New Manufacturer
		</button>
	</div>
	<div class="bg-base-300 mt-4 h-[calc(100%-6rem)] w-full overflow-auto rounded-lg p-4">
		<table class="table">
			<thead>
				<tr>
					<th>Manufaturer</th>
				</tr>
			</thead>
			<tbody>
				{#each filteredMakers as item}
					<tr
						class="hover cursor-pointer"
						onclick={() => {
							makerToUpdate = item;
							open = true;
						}}
					>
						<td>{item.name}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
	<!-- </div> -->
</div>

{#if makerToUpdate}
	<MakerDialog bind:open manufacturer={makerToUpdate} />
{:else}
	<MakerDialog bind:open />
{/if}
