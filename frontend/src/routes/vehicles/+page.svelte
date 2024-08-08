<script lang="ts">
	import { fuelTypes, models, vehicles, vehicleTypes } from '$lib/states';
	import { main } from '$lib/wailsjs/go/models';
	import { Plus } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import VehicleDialog from './VehicleDialog.svelte';

	let open = $state(false);
	let query = $state('');
	let vehicleToUpdate = $state<main.Vehicle>();
	let vehicleTypeFilter = $state('All');
	let fuelTypeFilter = $state('All');

	const filteredVehicles = $derived.by(() => {
		let all = vehicles.all;
		if (vehicleTypeFilter !== 'All') {
			all = vehicles.all.filter((m) => m.vehicleType === vehicleTypeFilter);
		}
		if (fuelTypeFilter !== 'All') {
			all = vehicles.all.filter((m) => m.fuelType === fuelTypeFilter);
		}
		if (query !== '')
			all = all.filter((m) => {
				const model = models.all.find((model) => model.id === m.modelId);

				return (
					model?.name.toLowerCase().includes(query.toLowerCase()) ||
					m.plateNo.toLowerCase().includes(query.toLowerCase())
				);
			});
		return all;
	});

	$effect(() => {
		if (vehicleToUpdate && !open) vehicleToUpdate = undefined;
	});

	onMount(async () => {
		vehicles.load();
	});
</script>

<div class="h-full w-full">
	<!-- <div class="bg-base-2000 border-base-content/30 h-full rounded-lg border p-4"> -->
	<div class="flex h-20 items-end justify-between">
		<div class="flex items-center gap-4">
			<div class="form-control">
				<label for="search" class="label label-text"> Model Name/Plate Number </label>
				<input
					type="text"
					id="search"
					class="input input-bordered"
					placeholder="Model S or PJ672KL"
					bind:value={query}
				/>
			</div>
			<div class="form-control">
				<label for="make" class="label label-text">Vehicle Type</label>
				<select id="make" class="select select-bordered" bind:value={vehicleTypeFilter}>
					<option value="All">All</option>
					{#each vehicleTypes as item}
						<option value={item}>{item}</option>
					{/each}
				</select>
			</div>
			<div class="form-control">
				<label for="make" class="label label-text">Fuel Type</label>
				<select id="make" class="select select-bordered" bind:value={fuelTypeFilter}>
					<option value="All">All</option>
					{#each fuelTypes as item}
						<option value={item}>{item}</option>
					{/each}
				</select>
			</div>
		</div>
		<button class="btn btn-outline" onclick={() => (open = true)}>
			<Plus class="h-5 w-5" />
			New Vehicle
		</button>
	</div>
	<div class="bg-base-300 mt-4 h-[calc(100%-6rem)] w-full overflow-auto rounded-lg p-4">
		<table class="table">
			<thead>
				<tr>
					<th>Model</th>
					<th>Vehicle Type</th>
					<th>Fuel Type</th>
					<th>Plate Number</th>
					<th>Capacity</th>
					<th>Department Assigned</th>
				</tr>
			</thead>
			<tbody>
				{#each filteredVehicles as item}
					<tr
						class="hover cursor-pointer"
						onclick={() => {
							vehicleToUpdate = item;
							open = true;
						}}
					>
						<th>{models.all.find((m) => m.id === item.modelId)?.name}</th>
						<td>{item.vehicleType}</td>
						<td>{item.fuelType}</td>
						<td>{item.plateNo}</td>
						<td>{item.capacity}</td>
						<td>{item.deptAssigned}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
	<!-- </div> -->
</div>
{#if vehicleToUpdate}
	<VehicleDialog bind:open vehicle={vehicleToUpdate} />
{:else}
	<VehicleDialog bind:open />
{/if}
