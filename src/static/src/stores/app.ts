import { writable } from 'svelte/store';

export const userLoggedIn = writable<null | boolean>(null);
