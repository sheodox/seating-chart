import { writable } from 'svelte/store';
import { send, on } from '../socket';

export interface Guest {
	id: string;
	tableId: string;
	firstName: string;
	lastName: string;
	people: number;
	//notes: string;
}

export const guests = writable<Guest[]>([]);

export const draggingGuest = writable<Guest>(null);

on('guests/list', (g: Guest[]) => {
	sortGuests(g);
	guests.set(g);
});

on('guests/add', (g: Guest) => {
	guests.update((oldGuests) => {
		const guests = [...oldGuests, g];
		sortGuests(guests);
		return guests;
	});
});

export function sortGuests(guests: Guest[]) {
	guests.sort((aGuest, bGuest) => {
		const a = aGuest.lastName.toLowerCase(),
			b = bGuest.lastName.toLowerCase();

		if (a === b) {
			return 0;
		}
		return a < b ? -1 : 1;
	});
}

on('guests/edit', (g: Guest) => {
	guests.update((oldGuests) => {
		const guests = oldGuests.map((guest) => (guest.id === g.id ? g : guest));
		sortGuests(guests);
		return guests;
	});
});

on('guests/delete', (id: string) => {
	guests.update((guests) => {
		return guests.filter((guest) => guest.id !== id);
	});
});

send('guests/list');

export const guestOps = {
	assign(guestId: string, tableId: string) {
		send('guests/assign', { tableId, guestId });
	},
};
