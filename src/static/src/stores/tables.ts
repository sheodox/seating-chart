import { writable } from 'svelte/store';
import { send, on } from '../socket';

export interface Table {
	id: string;
	name: string;
	capacity: number;
	posX: number;
	posY: number;
}

function initFromLocalStorage(key: string, fallback: number) {
	const val = localStorage.getItem(key);
	if (!val) {
		return fallback;
	}

	const int = parseInt(val, 10);
	if (isNaN(int)) {
		return fallback;
	}
	return int;
}

export const tables = writable<Table[]>([]);
export const highlightTable = writable('');
const tableZoomKey = 'tableZoom';
export const tableZoom = writable(initFromLocalStorage(tableZoomKey, 10));

tableZoom.subscribe((val) => {
	if (typeof val === 'number' && !isNaN(val)) {
		localStorage.setItem(tableZoomKey, '' + val);
	}
});

interface TableAdd {
	name: string;
	posX: number;
	posY: number;
}

export const tableOps = {
	add(args: TableAdd) {
		send('tables/add', args);
	},
	edit(table: Table) {
		send('tables/edit', table);
	},
	delete(id: string) {
		send('tables/delete', { id });
	},
};

on('tables/add', (table: Table) => {
	tables.update((tables) => [...tables, table]);
});

on('tables/delete', (id: string) => {
	tables.update((tables) => tables.filter((table) => table.id !== id));
});

on('tables/edit', (t: Table) => {
	tables.update((tables) => tables.map((table) => (table.id === t.id ? t : table)));
});

on('tables/list', (t: Table[]) => {
	tables.set(t);
});

send('tables/list');
