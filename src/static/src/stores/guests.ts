import { writable, derived } from 'svelte/store';
import { send, on } from '../socket';

export interface Guest {
	id: string;
	tableId: string;
	firstName: string;
	lastName: string;
	people: number;
	going: boolean;
	//notes: string;
}
function sortableName(guest: Guest) {
	return `${guest.lastName}, ${guest.firstName}`;
}

export const guests = writable<Guest[]>([]);
export const guestsAlphabetized = derived(guests, (guests) => {
	const sorted = [...guests];

	sorted.sort((g1, g2) => {
		return sortableName(g1).localeCompare(sortableName(g2));
	});

	return sorted;
});

export const draggingGuest = writable<Guest>(null);

export function guestDragStart(e: DragEvent, guest: Guest) {
	e.dataTransfer.setData('guestId', guest.id);
	e.stopPropagation();
	draggingGuest.set(guest);

	window.addEventListener(
		'dragend',
		() => {
			draggingGuest.set(null);
		},
		{ once: true }
	);
}

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
	add(args: { firstName: string; lastName: string; people: number; going: boolean }) {
		send('guests/add', args);
	},
};
