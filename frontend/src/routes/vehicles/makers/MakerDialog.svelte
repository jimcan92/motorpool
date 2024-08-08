<script lang="ts">
	import { makers } from '$lib/states/makers.svelte';
	import { main } from '$lib/wailsjs/go/models';
	import { Delete, RefreshCcw, Save, X } from 'lucide-svelte';
	import { onMount } from 'svelte';

	type Props = {
		open?: boolean;
		manufacturer?: main.Maker;
	};

	let { open = $bindable(), manufacturer }: Props = $props();

	let el: HTMLDialogElement;

	let make = $state('');

	$effect(() => {
		if (open) el.show();
		else el.close();
	});

	onMount(() => {
		if (manufacturer) make = manufacturer.name;
	});

	async function onSave(e: SubmitEvent) {
		e.preventDefault();

		if (manufacturer) await makers.update({ id: manufacturer.id, name: make });
		else await makers.add({ id: make.replaceAll(' ', '_'), name: make });
		close();
	}

	function close() {
		clearFields();
		open = false;
	}

	async function onDelete() {
		if (manufacturer) {
			await makers.delete(manufacturer.id);
			close();
		}
	}

	function clearFields() {
		make = '';
	}
</script>

<dialog class="modal bg-base-300/70" bind:this={el}>
	<form class="modal-box border-base-content/30 rounded-lg border shadow-none" onsubmit={onSave}>
		<h3 class="mb-4 text-lg font-semibold">
			{manufacturer ? `Update '${manufacturer.name}' manufacturer` : 'Add New Manufacturer'}
		</h3>
		<div class="form-control">
			<label for="make" class="label label-text">Manufacturer</label>
			<input type="text" name="make" id="make" class="input input-bordered" bind:value={make} />
		</div>
		<div class="modal-action">
			<button type="button" class="btn btn-ghost" onclick={close}>
				<X class="h-5 w-5" />
				Cancel
			</button>
			{#if manufacturer}
				<button type="button" class="btn btn-warning" onclick={onDelete}>
					<Delete class="h-5 w-5" />
					Delete
				</button>
			{/if}
			<button class="btn btn-accent">
				{#if manufacturer}
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
