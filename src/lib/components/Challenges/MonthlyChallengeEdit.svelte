<script lang="ts">
	import { client, makeSlug } from '$lib/stores/stores';
	import { filter } from 'd3';
	import AnalyticsPage from '../Analytics/AnalyticsPage.svelte';
	import ContentEditorSection from '../Editor/content-editor-section.svelte';
	import ContentEditor from '../Editor/content-editor.svelte';
	import DateEdit from '../Editor/date-edit.svelte';
	import FileEdit from '../Editor/file-edit.svelte';
	import RtfEdit from '../Editor/rtf-edit.svelte';
	import SelectEdit from '../Editor/select-edit.svelte';
	import StringEdit from '../Editor/string-edit.svelte';
	import ChallengeCard from './ChallengeCard.svelte';
	import type { Challenge, Difficulty, MonthlyChallenge } from './challenges';
	import DifficultyCard from './DifficultyCard.svelte';
	import DifficultyEdit from './DifficultyEdit.svelte';
	import TagsEdit from './TagsEdit.svelte';
	import TopicsEdit from './TopicsEdit.svelte';

	export let data: MonthlyChallenge = {
		from: '',
		to: '',
		challenges: [],
		title: '',
		body: '',
		sources: ''
	};
	let newChallenge;
	const dataclass = 'monthly_challenge';
	export let editing = false;

	const jsonFields = [];
	const listFields = ['challenges'];

	let permaLink;

	const externalBaseUrl =
		import.meta.env.VITE_APP_BASE_URL || 'https://cy-challenge-app-beta.netlify.app';

	$: if (editing && data.id) {
		permaLink = `${externalBaseUrl}/monthly-challenges/${data.id}`;
	}

	const fetchChallenge = (challengeId) => {
		return client.records.getOne('challenges', challengeId);
	};

	const fetchChallenges = () => {
		return client.records.getFullList('challenges', 100, {});
	};
</script>

<ContentEditor
	{dataclass}
	bind:data
	bind:title={data.title}
	titleplaceholder="Neue MonatsChallenge"
	saveTarget="monthly_challenges"
	{editing}
	{jsonFields}
	{listFields}
>
	<ContentEditorSection label="Metadata">
		<StringEdit id="id" label="id" type="string" placeholder="id" readonly bind:value={data.id} />

		<div class="flex flex-col text-sm text-zinc-700">
			<span class="font-semibold text-black text-base">Link in der App</span>
			Speichern überträgt die Änderungen sofort - funktioniert auch ohne "Veröffentlichen", die Challenge
			kann dann nur nicht in der Challenge-Liste gefunden werden.
			<a target="_blank" href={permaLink} class="text-blue-800 text-base underline">
				{permaLink}
			</a>
		</div>
	</ContentEditorSection>

	<ContentEditorSection label="Seitensetup">
		<StringEdit
			id="title"
			label="Überschrift"
			type="string"
			placeholder="Title"
			bind:value={data.title}
			recommendedLength={50}
		/>

		<RtfEdit
			id="frontMatter"
			label="Einleitung"
			description="Wird auf der Seite angezeigt?"
			bind:value={data.body}
			recommendedLength={150}
		/>

		<FileEdit id="image" label="Image" bind:value={data.headerImage} record={data} />
	</ContentEditorSection>

	<ContentEditorSection label="Zeitraum">
		<DateEdit id="from" label="von" bind:value={data.from} />

		<DateEdit id="to" label="bis" bind:value={data.to} />
	</ContentEditorSection>

	<ContentEditorSection label="Challenges">
		{#if data.challenges}
			{#each data.challenges as challengeId}
				{#await fetchChallenge(challengeId) then challenge}
					<div
						on:click={(e) => {
							data.challenges = [...data.challenges.filter((value) => value !== challenge.id)];
						}}
					>
						<ChallengeCard data={challenge} />
					</div>
				{/await}
			{/each}
		{/if}

		{#await fetchChallenges() then challenges}
			<div>
				<label for="addChallenge">Challenge hinzufügen</label>
				<select
					id="addChallenge"
					bind:value={newChallenge}
					class="w-full rounded-md shadow-sm px-4 py-2"
				>
					<option value={null}>keine</option>
					{#each challenges as challenge}
						<option value={challenge.id}>{challenge.title}</option>
					{/each}
				</select>
			</div>

			<button
				class="px-2 py-1  h-full bg-opacity-0 border-2 border-black text-black text-center font-bold rounded hover:bg-green-500"
				on:click={(e) => {
					if (!data.challenges) {
						data.challenges = [];
					}
					data.challenges = [...data.challenges, newChallenge];
					newChallenge = undefined;
				}}>Challenge hinzufügen +</button
			>
		{/await}
	</ContentEditorSection>

	<ContentEditorSection label="Quellen">
		<RtfEdit
			id="sources"
			label="Quellen"
			description=""
			bind:value={data.sources}
			recommendedLength={100}
		/>
	</ContentEditorSection>
</ContentEditor>

<!-- <pre>
    {JSON.stringify(data, null, 2)}
</pre> -->
