import { writable } from 'svelte/store';

export type EditorMode = 'lines' | 'tables' | 'assignments';

export const editorMode = writable<EditorMode>('assignments');
