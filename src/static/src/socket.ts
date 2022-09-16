import { createAutoExpireToast } from 'sheodox-ui';
import { userLoggedIn } from './stores/app';

const RECONNECT_DELAY_MS = 1000;

let socket: WebSocket | null = null;
let queued: WSMessage[] = [];

type Handler = (data: any) => any;
interface WSMessage {
	route: string;
	data: any;
}
const handlers: Map<string, Handler> = new Map();

export function on(route: string, handler: Handler) {
	if (handlers.has(route)) {
		console.warn(`There's already a handler for "${route}", overwriting.`);
	}
	handlers.set(route, handler);
}

export function send(route: string, data?: any) {
	const msg = { route, data };
	console.log('sent message: ', { route, data });
	if (!socket) {
		queued.push(msg);
	} else {
		sendMsg(msg);
	}
}

function sendMsg(msg: WSMessage) {
	socket.send(JSON.stringify(msg));
}

async function setupSocket() {
	const res = await fetch(`/api/auth/status`, {
		credentials: 'include',
	});

	if (res.status === 401) {
		userLoggedIn.set(false);
		return;
	}

	userLoggedIn.set(true);

	const socketProtocol = location.protocol === 'http:' ? 'ws' : 'wss';
	const ws = new WebSocket(`${socketProtocol}://${location.hostname}:${location.port}/api/ws`);

	ws.addEventListener('error', (e) => {
		console.log(e);
	});

	ws.addEventListener('message', (msg) => {
		if (msg.data === 'pong') {
			return;
		}

		const { route, data } = JSON.parse(msg.data);

		console.log('received message: ', { route, data });

		if (handlers.has(route)) {
			handlers.get(route)(data);
		} else {
			console.warn(`Received message for "${route}", but no handler exists`);
		}
	});

	ws.addEventListener('open', () => {
		// need to ping occasionally to keep the connection alive
		//function ping() {
		//if (ws.readyState === WebSocket.OPEN) {
		//ws.send('ping');
		//setTimeout(ping, PING_INTERVAL_MS);
		//}
		//}
		//ping();
		socket = ws;

		for (const msg of queued) {
			sendMsg(msg);
		}
		queued = [];
	});

	//if the connection is closed we should try to reconnect it
	ws.addEventListener('close', () => {
		socket = null;
		setTimeout(setupSocket, RECONNECT_DELAY_MS);
	});
}

setupSocket();

on('error', (msg) => {
	console.error(`SocketErr: ${msg}`);

	createAutoExpireToast({
		variant: 'error',
		title: 'Error',
		message: msg,
	});
});
