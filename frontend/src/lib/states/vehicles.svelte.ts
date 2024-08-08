import { GetAllVehicles } from '$lib/wailsjs/go/main/App';
import { main } from '$lib/wailsjs/go/models';

const _vehicles = $state<main.Vehicle[]>([]);

export const vehicles = {
	get all() {
		return _vehicles;
	},
	async load() {
		await GetAllVehicles();
	}
};
