<script lang="ts">
	import { makers } from '$lib/states/makers.svelte';
	import { models } from '$lib/states/models.svelte';
	import { main } from '$lib/wailsjs/go/models';
	import { Plus } from 'lucide-svelte';
	import ModelDialog from './ModelDialog.svelte';

	let open = $state(false);
	let query = $state('');
	let modelToUpdate = $state<main.Model>();
	let makeFilter = $state('All');

	const filteredModels = $derived.by(() => {
		let all = models.all;
		if (makeFilter !== 'All') {
			all = models.all.filter((m) => m.makerId === makeFilter);
		}
		if (query !== '') all = all.filter((m) => m.name.toLowerCase().includes(query.toLowerCase()));
		return all;
	});

	$effect(() => {
		if (modelToUpdate && !open) modelToUpdate = undefined;
	});
</script>

<div class="h-full w-full">
	<!-- <div class="bg-base-2000 border-base-content/30 h-full rounded-lg border p-4"> -->
	<div class="flex h-20 items-end justify-between">
		<div class="flex items-center gap-4">
			<div class="form-control">
				<label for="search" class="label label-text">Search Model Name</label>
				<input type="text" id="search" class="input input-bordered" bind:value={query} />
			</div>
			<div class="form-control">
				<label for="make" class="label label-text">Filter by Manufaturer</label>
				<select id="make" class="select select-bordered" bind:value={makeFilter}>
					<option value="All">All</option>
					{#each makers.all as item}
						<option value={item.id}>{item.name}</option>
					{/each}
				</select>
			</div>
		</div>
		<button class="btn btn-outline" onclick={() => (open = true)}>
			<Plus class="h-5 w-5" />
			New Model
		</button>
	</div>
	<div class="bg-base-300 mt-4 h-[calc(100%-6rem)] w-full overflow-auto rounded-lg p-4">
		<table class="table">
			<thead>
				<tr>
					<th>Manufaturer</th>
					<th>Model Name</th>
					<th>Year</th>
				</tr>
			</thead>
			<tbody>
				{#each filteredModels as item}
					<tr
						class="hover cursor-pointer"
						onclick={() => {
							modelToUpdate = item;
							open = true;
						}}
					>
						<td>{makers.all.find((m) => m.id === item.makerId)?.name}</td>
						<th>{item.name}</th>
						<td>{item.year}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
	<!-- </div> -->
</div>
{#if modelToUpdate}
	<ModelDialog bind:open model={modelToUpdate} />
{:else}
	<ModelDialog bind:open />
{/if}
