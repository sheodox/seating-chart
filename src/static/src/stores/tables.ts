import { writable } from 'svelte/store';
import { send, on } from '../socket';

export interface Table {
	id: string;
	name: string;
	capacity: number;
	posX: number;
	posY: number;
}

export const tables = writable<Table[]>([]);
export const draggingOverTable = writable('');

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
