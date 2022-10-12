import type { Writable } from 'svelte/store';
import { writable } from 'svelte/store';

import PocketBase from 'pocketbase';
import { goto } from '$app/navigation';
import { browser } from '$app/environment';

export class Answer {
	value: string = '';
	correct: boolean = false;
}

export class Question {
	question: string = '';
	answers: Answer[] = [new Answer()];
	infobit?: number = 0;
}

export class Infobit {
	doc = {
		type: 'doc',
		content: [
			{
				type: 'paragraph'
			}
		]
	};
	selection = {
		type: 'text',
		anchor: 1,
		head: 1
	};
}

export class Infobyte {
	_id: string = '';
	name: string = '';
	published: boolean = false;
	region: string = 'DE';
	language: string = 'DE';
	frontmatter?: string = '';
	questions?: Question[] = [new Question()];
	infobits: any[] = [];
	aspect: string;
	factor: string;
	difficulty: number;
	position: number;
}

export class Aspect {
	_id: string = '';
	name: string = '';
}

export const currentInfobyte: Writable<Infobyte> = writable(null);
export const currentInfobitIndex: Writable<number> = writable(null);
export const currentInfobit: Writable<Infobit> = writable(null);
export const currentQuestionName: Writable<string> = writable(null);
export const currentQuestionIndex: Writable<number> = writable(null);
export const currentQuestion: Writable<Question> = writable(null);
export const isProd = false;

export const client = new PocketBase(import.meta.env.VITE_PB_BASE_URL ?? '/');

if (browser) {
	console.log(document.cookie);
	// await client.authStore.loadFromCookie(document.cookie, 'pocketbase_auth');

	if (!client.authStore.isValid) {
		const next = window.location.pathname;
		if (next === '/login') {
			goto(`/`);
		}
		goto(`/login?next=${next}`);
	} else {
		const path = window.location.pathname;
		if (path.includes('login')) {
			goto('/');
		}
	}
}

export const login = async (email, password) => {
	await client.admins
		.authViaEmail(email, password)
		.then((res) => {
			console.log(email);
			// document.cookie = `PB_AUTH=${client.authStore.exportToCookie()}`;
			const queryString = window.location.search;
			const urlParams = new URLSearchParams(queryString);
			const next = urlParams.get('next');
			if (next) {
				if (next.includes('login')) {
					goto('/');
				} else {
					goto(next);
				}
			} else {
				goto('/');
			}
		})
		.catch((err) => {
			console.warn(err);
		});
};

export const makeSlug = (title) => {
	const regex = /[^a-z_]/g;
	return title.split('(')[0].trim().toLowerCase().replace(regex, '_');
};

export const activeTabStore = writable('');
