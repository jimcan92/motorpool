<script lang="ts">
	import { page } from '$app/stores';
	import { Boxes, CarFront, SwatchBook } from 'lucide-svelte';

	const { children } = $props();
</script>

<div class="h-full w-full">
	<div class="flex h-16 items-center gap-4 px-4">
		<a
			href="/vehicles"
			class={`flex items-center p-2 font-semibold gap-2${$page.url.pathname !== '/vehicles' ? ' text-base-content/50' : ' border-base-content border-b'}`}
		>
			<CarFront class="h-5 w-5" />
			VEHICLES
		</a>
		<a
			href="/vehicles/models"
			class={`flex items-center p-2 font-semibold gap-2${$page.url.pathname !== '/vehicles/models' ? ' text-base-content/50' : ' border-base-content border-b'}`}
		>
			<SwatchBook class="h-5 w-5" />
			MODELS
		</a>
		<a
			href="/vehicles/makers"
			class={`flex items-center p-2 font-semibold gap-2${$page.url.pathname !== '/vehicles/makers' ? ' text-base-content/50' : ' border-base-content border-b'}`}
		>
			<Boxes class="h-5 w-5" />
			MANUFACTURERS
		</a>
	</div>
	<div class="h-[calc(100%-4rem)] p-4">
		{@render children()}
	</div>
</div>
