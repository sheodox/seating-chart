import { writable } from 'svelte/store';
import page from 'page';
import type { Context } from 'page';

//the name of the header nav item this route lives within
export const activeApp = writable<string>('');
//a rough pseudocode-y ID that corresponds to a component that's mounted in Routing.svelte
export const activeRoute = writable<string>('');
export const activeRouteParams = writable<Record<string, string>>({});
export const activeQueryParams = writable<Record<string, string>>({});

function setRoute(app: string, routeId: string) {
	return (context: Context) => {
		activeApp.set(app);
		activeRoute.set(routeId);
		//activeRouteParams.set(context.params);

		//const queryParams: Record<string, string> = {};
		//for (const [param, value] of new URLSearchParams(context.querystring)) {
		//queryParams[param] = decodeURIComponent(value);
		//}
		//activeQueryParams.set(queryParams);
	};
}

page(`/guests`, setRoute('guests', 'guests'));
page(`/`, setRoute('chart', 'chart'));
page();
