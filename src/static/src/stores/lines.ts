import { on, send } from '../socket';
import { writable } from 'svelte/store';

export interface LineCoordinate {
	x: number;
	y: number;
}
export interface Line {
	id: string;
	coords: LineCoordinate[];
	closed: boolean;
}

export const lines = writable<Line[]>([]);
export const editingLineIndex = writable<number>(0);
export const highlightingLineIndex = writable<number>(-1);

export const lineOps = {
	edit(line: Line) {
		send('lines/edit', line);
	},
	add() {
		send('lines/add');
	},
	delete(id: string) {
		send('lines/delete', { id });
	},
};

on('lines/edit', (line: Line) => {
	lines.update((lines) => lines.map((l) => (l.id === line.id ? line : l)));
});

on('lines/delete', (id: string) => {
	lines.update((lines) => lines.filter((l) => l.id !== id));
});

on('lines/add', (line: Line) => {
	lines.update((lines) => [...lines, line]);
});

on('lines/list', (l: Line[]) => {
	lines.set(l);
});

send('lines/list');
