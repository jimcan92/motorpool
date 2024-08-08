import { CreateMaker, DeleteMaker, GetAllMakers, UpdateMaker } from '$lib/wailsjs/go/main/App';
import type { main } from '$lib/wailsjs/go/models';

let _makers = $state<main.Maker[]>([]);

export const makers = {
	get all() {
		return _makers;
	},
	async load() {
		_makers = await GetAllMakers();
	},
	async add(maker: main.Maker) {
		await CreateMaker(maker);
		await this.load();
	},
	async delete(id: string) {
		await DeleteMaker(id);
		await this.load();
	},
	async update(maker: main.Maker) {
		await UpdateMaker(maker);
		await this.load();
	}
};
