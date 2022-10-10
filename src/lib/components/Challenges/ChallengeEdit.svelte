<script lang="ts">
	import { makeSlug } from '$lib/stores/stores';
	import ContentEditorSection from '../Editor/content-editor-section.svelte';
	import ContentEditor from '../Editor/content-editor.svelte';
	import SelectEdit from '../Editor/select-edit.svelte';
	import StringEdit from '../Editor/string-edit.svelte';
	import type { Challenge, Difficulty } from './challenges';
	import DifficultyCard from './DifficultyCard.svelte';
	import DifficultyEdit from './DifficultyEdit.svelte';

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
</script>

<ContentEditor
	{dataclass}
	bind:data
	bind:title={data.title}
	titleplaceholder="Neue Challenge"
	saveTarget="challenges"
	{editing}
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

		<StringEdit
			id="title"
			label="Title"
			type="string"
			placeholder="Title"
			bind:value={data.title}
			onChange={() => {
				console.log('input has happened');
				if (!editing) {
					data.slug = makeSlug(data.title);
				}
			}}
		/>

		<StringEdit
			id="topic"
			label="Bereich"
			type="string"
			placeholder="Bereich"
			bind:value={data.topic}
		/>

		<StringEdit
			id="tags"
			label="Tags"
			type="string"
			placeholder="tag1, tag2, ..."
			value={data.tags?.join(', ') ?? ''}
			onChange={(e) => {
				console.log(e.target.value.split(','));
				data.tags = e.target.value.split(',').map((v) => v.trim());
			}}
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
				<input id="lead" type="checkbox" />
			</label>
		</div>

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
	</ContentEditorSection>

	<ContentEditorSection label="Inhalt">
		<StringEdit
			id="frontMatter"
			label="Einleitung"
			type="richtext"
			placeholder="frontMatter"
			bind:value={data.frontMatter}
		/>

		<StringEdit
			id="summary"
			label="Warum"
			type="richtext"
			placeholder="summary"
			bind:value={data.summary}
		/>

		<StringEdit id="tips" label="Tips" type="richtext" placeholder="tips" bind:value={data.tips} />

		<StringEdit
			id="todos"
			label="Todos"
			type="richtext"
			placeholder="todos"
			bind:value={data.todos}
		/>

		<StringEdit
			id="content"
			label="Tiefer eingehen"
			type="richtext"
			placeholder="content"
			bind:value={data.content}
		/>

		<StringEdit
			id="upgradeText"
			label="Alle Level geschafft Text"
			type="richtext"
			placeholder="upgradeText"
			bind:value={data.completedText}
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
			<span class="font-semibold">Schwierigkeitsgrade</span>
			{#each data.difficulties ? Object.values(data.difficulties) : [] as diff}
				<svelte:component this={DifficultyCard} bind:value={diff} />
			{/each}
		</div>

		<DifficultyEdit bind:data={selectedDifficulty} />
		<div class=" w-full flex place-content-center">
			<button
				class="text-black text-lg px-4 py-2 border-black border-2 rounded-lg"
				on:click={(e) => {
					addDifficulty();
				}}>Schwierigkeitsgrad hinzufügen <b>+</b></button
			>
		</div>
	</ContentEditorSection>

	<ContentEditorSection label="Benachrichtigungen">
		<StringEdit
			id="reminderText"
			label="reminderText"
			type="string"
			placeholder="reminderText"
			bind:value={data.reminderText}
		/>
	</ContentEditorSection>
</ContentEditor>

<pre>
    {JSON.stringify(data, null, 2)}
</pre>
