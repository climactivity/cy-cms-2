export interface Challenge {
	id?: string;
	slug: string;
	title?: string;
	topic?: string;
	tags?: string[];

	lead?: boolean;
	impact?: ChallengeImpact;
	type?: ChallengeType;

	frontMatter?: string;
	summary?: string;
	tips?: string;
	todos?: string;
	content?: string;
	difficulties?: Difficulties;
	completedText?: string;

	image?: File | string;

	reminderText?: string;
	notificationDays?: number[];

	score?: number;

	published?: boolean;
}

export interface Difficulties {
	[difficulties: string]: Difficulty;
}

export interface Difficulty {
	name: string;
	index: number;
	taskDescription: string;
	todos: Todo[];
	upgradeText: string;
}

export interface Todo {
	name: string;
}

export type ChallengeImpact = 'bigpoint' | 'peanut';
export type ChallengeType = 'one-time' | 'recurring' | 'repeatable';
