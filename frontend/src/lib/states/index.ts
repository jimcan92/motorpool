import { FuelType, VehicleType } from '$lib/types/vehicle';
import { makers } from './makers.svelte';
import { models } from './models.svelte';
import { vehicles } from './vehicles.svelte';

async function loadAll() {
	await makers.load();
	await models.load();
	await vehicles.load();
}

const vehicleTypes: VehicleType[] = [
	VehicleType.Car,
	VehicleType.Motorcycle,
	VehicleType.Truck,
	VehicleType.Van
];

const fuelTypes: FuelType[] = [FuelType.Diesel, FuelType.Gasoline];

export { fuelTypes, loadAll, makers, models, vehicles, vehicleTypes };
