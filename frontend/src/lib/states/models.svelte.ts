import { CreateModel, DeleteModel, GetAllModels, UpdateModel } from '$lib/wailsjs/go/main/App';
import { main } from '$lib/wailsjs/go/models';

let _models = $state<main.Model[]>([]);

export const models = {
	get all() {
		return _models;
	},
	async load() {
		_models = await GetAllModels();
	},
	async add(model: main.Model) {
		await CreateModel(model);
		await this.load();
	},
	async update(model: main.Model) {
		await UpdateModel(model);
		await this.load();
	},
	async delete(id: string) {
		await DeleteModel(id);
		await this.load();
	}
};
