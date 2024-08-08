<script lang="ts">
	import { makers } from '$lib/states/makers.svelte';
	import { models } from '$lib/states/models.svelte';
	import { main } from '$lib/wailsjs/go/models';
	import dayjs from 'dayjs';
	import { Delete, RefreshCcw, Save, X } from 'lucide-svelte';
	import { onMount } from 'svelte';

	type Props = {
		open?: boolean;
		model?: main.Model;
	};

	let { open = $bindable(), model }: Props = $props();

	let el: HTMLDialogElement;

	let make = $state<main.Maker>(makers.all[0]);
	let name = $state('');
	let year = $state(dayjs().year());

	$effect(() => {
		if (open) el.show();
		else el.close();
	});

	onMount(() => {
		if (model) {
			make = makers.all.find((m) => m.id === model.makerId) ?? makers.all[0];
			name = model.name;
			year = model.year;
		}
	});

	async function onSave(e: SubmitEvent) {
		e.preventDefault();

		if (model) {
			await models.update({
				id: model.id,
				makerId: make.id,
				name,
				year
			});
		} else {
			const model: main.Model = {
				id: `${make.id}_${name.replaceAll(' ', '_')}`,
				makerId: make.id,
				name,
				year
			};

			await models.add(model);
		}
		close();
	}

	function close() {
		clearFields();
		open = false;
	}

	async function onDelete() {
		if (model) {
			await models.delete(model.id);
			close();
		}
	}

	function clearFields() {
		make = makers.all[0];
		name = '';
		year = dayjs().year();
	}
</script>

<dialog class="modal bg-base-300/70" bind:this={el}>
	<form class="modal-box border-base-content/30 rounded-lg border shadow-none" onsubmit={onSave}>
		<h3 class="mb-4 text-lg font-semibold">
			{model ? `Update '${model.name}' model` : 'Add New Model'}
		</h3>
		<div class="form-control">
			<label for="make" class="label label-text">Manufacturer</label>
			<!-- <input type="text" name="make" id="make" class="input input-bordered" bind:value={make} /> -->
			<select name="make" id="make" class="select select-bordered" bind:value={make}>
				{#each makers.all as item}
					<option value={item}>{item.name}</option>
				{/each}
			</select>
		</div>
		<div class="form-control">
			<label for="name" class="label label-text">Model Name</label>
			<input type="text" name="name" id="name" class="input input-bordered" bind:value={name} />
		</div>
		<div class="form-control">
			<label for="year" class="label label-text">Year</label>
			<input
				type="number"
				name="year"
				id="year"
				class="input input-bordered"
				min={dayjs().add(-100, 'year').year()}
				max={dayjs().year()}
				bind:value={year}
			/>
		</div>
		<div class="modal-action">
			<button type="button" class="btn btn-ghost" onclick={close}>
				<X class="h-5 w-5" />
				Cancel
			</button>
			{#if model}
				<button type="button" class="btn btn-warning" onclick={onDelete}>
					<Delete class="h-5 w-5" />
					Delete
				</button>
			{/if}
			<button class="btn btn-accent">
				{#if model}
					<RefreshCcw class="h-5 w-5" />
					Update
				{:else}
					<Save class="h-5 w-5" />
					Save
				{/if}
			</button>
		</div>
	</form>
</dialog>
