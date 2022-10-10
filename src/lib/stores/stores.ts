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

export const client = new PocketBase('http://127.0.0.1:8090');

if (browser) {
	client.authStore.loadFromCookie(document.cookie, 'PB_AUTH');

	if (!client.authStore.isValid) {
		goto(`/login?next=${window.location}`);
	}
}

export const login = async (email, password) => {
	await client.admins
		.authViaEmail(email, password)
		.then((res) => {
			console.log(email);
			document.cookie = `PB_AUTH=${client.authStore.exportToCookie()}`;
			const queryString = window.location.search;
			const urlParams = new URLSearchParams(queryString);
			const next = urlParams.get('next');
			if (next) {
				goto(next);
			} else {
				goto('/');
			}
		})
		.catch((err) => {
			console.warn(err);
		});
};
