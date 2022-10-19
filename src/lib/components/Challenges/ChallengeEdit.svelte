<script lang="ts">
	import { client, makeSlug } from '$lib/stores/stores';
	import { filter } from 'd3';
	import AnalyticsPage from '../Analytics/AnalyticsPage.svelte';
	import ContentEditorSection from '../Editor/content-editor-section.svelte';
	import ContentEditor from '../Editor/content-editor.svelte';
	import FileEdit from '../Editor/file-edit.svelte';
	import RtfEdit from '../Editor/rtf-edit.svelte';
	import SelectEdit from '../Editor/select-edit.svelte';
	import StringEdit from '../Editor/string-edit.svelte';
	import type { Challenge, Difficulty } from './challenges';
	import DifficultyCard from './DifficultyCard.svelte';
	import DifficultyEdit from './DifficultyEdit.svelte';
	import TagsEdit from './TagsEdit.svelte';
	import TopicsEdit from './TopicsEdit.svelte';

	export let data: Challenge = {
		slug: '',
		published: false
	};

	const dataclass = 'challenge';
	export let editing = false;

	let selectedDifficulty: Partial<Difficulty> = {};
	const addDifficulty = () => {
		const name = selectedDifficulty.name;
		const difficulties = data.difficulties ?? {};
		difficulties[name] = selectedDifficulty as Difficulty;
		data.difficulties = difficulties;
		console.log(difficulties);
		selectedDifficulty = {};
	};

	const jsonFields = ['difficulties', 'notificationDays'];
	const listFields = ['tags'];

	let permaLink;

	const externalBaseUrl =
		import.meta.env.VITE_APP_BASE_URL || 'https://cy-challenge-app-beta.netlify.app';

	$: if (editing && data.id) {
		permaLink = `${externalBaseUrl}/challenge/${data.slug}`;
	}
</script>

<ContentEditor
	{dataclass}
	bind:data
	bind:title={data.title}
	titleplaceholder="Neue Challenge"
	saveTarget="challenges"
	{editing}
	{jsonFields}
	{listFields}
>
	<ContentEditorSection label="Metadata">
		<StringEdit
			id="slug"
			label="slug"
			type="string"
			placeholder="slug"
			readonly
			bind:value={data.slug}
		/>

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

	<ContentEditorSection label="Titel">
		<StringEdit
			id="title"
			label="Überschrift"
			type="string"
			placeholder="Title"
			bind:value={data.title}
			onChange={() => {
				console.log('input has happened');
				if (!editing) {
					data.slug = makeSlug(data.title);
				}
			}}
			recommendedLength={50}
		/>

		<TopicsEdit id="topic" label="Bereich" placeholder="-" bind:value={data.topic} />

		<TagsEdit id="tags" label="Tags" placeholder="tag1, tag2, ..." bind:value={data.tags} />

		<FileEdit id="image" label="Image" bind:value={data.image} record={data} />

		<RtfEdit
			id="frontMatter"
			label="Einleitung"
			description="Wird über den Aktionen angezeigt"
			bind:value={data.frontMatter}
			recommendedLength={150}
		/>
	</ContentEditorSection>

	<ContentEditorSection label="Impact">
		<SelectEdit
			id="impact"
			label="Impact"
			placeholder="-"
			bind:value={data.impact}
			options={[
				{ value: 'bigpoint', label: 'Bigpoint' },
				{ value: 'peanut', label: 'Peanut' }
			]}
		/>

		<div class="py-4">
			<label for="lead">
				Ist Super-Challenge?
				<input id="lead" type="checkbox" bind:checked={data.lead} />
			</label>
		</div>
	</ContentEditorSection>

	<ContentEditorSection label="Inhalt">
		<RtfEdit
			id="summary"
			label="Warum"
			description="Wird unter den Aktionen angezeigt, fasst das Warum zusammen "
			bind:value={data.summary}
			recommendedLength={100}
		/>

		<RtfEdit
			id="tips"
			label="Tips"
			description="Hilfreiche (optionale) Todos"
			bind:value={data.tips}
			recommendedLength={200}
		/>

		<RtfEdit
			id="content"
			label="Tiefer eingehen"
			description="Vertiefende Informationen, da unter dem Fold weiter scrollt kann hier mehr stehen (aber bitte nicht mehr als 1000 Zeichen"
			bind:value={data.content}
			recommendedLength={800}
		/>

		<RtfEdit
			id="upgradeText"
			label="Alle Level geschafft Text"
			description="Wird angeizeigt wenn alle Schwierigkeitsgrade der Challenge geschafft wurden, klopf der Nutzer:in nochmal auf die Schulter"
			bind:value={data.completedText}
			recommendedLength={100}
		/>
	</ContentEditorSection>

	<ContentEditorSection label="Fortschritt">
		<!-- <ReorderableList
			indexField="name"
			listItem={DifficultyCard}
			value={data.difficulties ? Object.values(data.difficulties) : []}
			listCols="1fr 2fr 4fr 4fr 1fr 1fr"
			headers={['Index', 'Name', 'Beschreibung', 'Todos', 'Geschafft']}
		/> -->

		<div>
			<span class="font-semibold"
				>Schwierigkeitsgrade <span class="text-sm text-zinc-600"
					>(zum bearbeiten anklicken, Inhalt wird in den Eingabefeldern unten angeizeigt)</span
				></span
			>
			{#each data.difficulties ? Object.values(data.difficulties) : [] as diff}
				<svelte:component
					this={DifficultyCard}
					bind:value={diff}
					index={Object.values(data.difficulties).indexOf(diff) ?? 0}
					onClick={(e) => {
						console.log(diff);
						selectedDifficulty = diff;
					}}
					onDelete={() => {
						console.log('delete diff');
						delete data.difficulties[diff.name];
						data.difficulties = { ...data.difficulties };
					}}
				/>
			{/each}
		</div>

		<DifficultyEdit bind:data={selectedDifficulty} />
		<div class=" w-full flex place-content-center">
			<button
				class="text-black text-lg px-4 py-2 border-black border-2 rounded-lg"
				on:click={(e) => {
					addDifficulty();
				}}
			>
				{#if data && selectedDifficulty && data.difficulties && data.difficulties[selectedDifficulty.name]}
					Schwierigkeitsgrad aktualisieren
				{:else}
					Schwierigkeitsgrad hinzufügen <b>+</b>
				{/if}
			</button>
		</div>
	</ContentEditorSection>

	<ContentEditorSection label="Benachrichtigungen">
		<SelectEdit
			id="type"
			label="Challenge Art"
			placeholder="-"
			bind:value={data.type}
			options={[
				{ value: 'recurring', label: 'wöchentlich' },
				{ value: 'one-time', label: 'einmal' },
				{ value: 'repeatable', label: 'wiederholbar' }
			]}
		/>

		<StringEdit
			id="reminderText"
			label="Benachrichtigungstext"
			type="string"
			placeholder="reminderText"
			bind:value={data.reminderText}
			recommendedLength={50}
		/>

		<StringEdit
			id="tags"
			label="Erninnern nach ... Tagen"
			type="string"
			placeholder="7, 14, ..."
			value={data.notificationDays && data.notificationDays instanceof Array
				? data.notificationDays.join(', ')
				: data.notificationDays}
			onChange={(e) => {
				console.log(e.target.value.split(','));
				data.notificationDays = e.target.value
					.split(',')
					.map((v) => {
						let val = parseInt(v.trim());
						if (isNaN(val)) return;
						return val;
					})
					.filter((v) => v != null);
			}}
			recommendedLength={-1}
		/>
	</ContentEditorSection>
</ContentEditor>

<!-- <pre>
    {JSON.stringify(data, null, 2)}
</pre> -->
