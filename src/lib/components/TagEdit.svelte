<script lang="ts">
	import { client, makeSlug } from '$lib/stores/stores';
	import ContentEditorSection from './Editor/content-editor-section.svelte';
	import StringEdit from './Editor/string-edit.svelte';

	let data: { label: string; topic: string; id?: string; [k: string]: string } = {
		label: '',
		topic: ''
	};

	export let afterSubmit = () => {};
	const addTag = () => {
		console.log(data);
		client.records.create('tags', data);
	};

	const updateTag = () => {
		client.records.update('tags', data.id, data);
	};
</script>

<ContentEditorSection label="Neuen Tag erstellen">
	{#if data?.id}
		<StringEdit id="id" label="id" type="string" placeholder="id" readonly bind:value={data.id} />
	{/if}
	<StringEdit
		id="slug"
		label="slug"
		type="string"
		placeholder="slug"
		readonly
		bind:value={data.tag}
	/>
	<StringEdit
		id="slug"
		label="Tag"
		type="string"
		placeholder="tag"
		bind:value={data.label}
		onInput={() => {
			console.log('input has happened');
			if (!data.id) {
				data.tag = makeSlug(data.label);
			}
		}}
	/>

	<div class=" w-full flex place-content-center">
		<button
			class="text-black text-lg px-4 py-2 border-black border-2 rounded-lg"
			on:click={async (e) => {
				if (!data.id) {
					await addTag();
				} else {
					await updateTag();
				}
				afterSubmit();
			}}>Tag hinzufügen <b>+</b></button
		>
	</div>
</ContentEditorSection>
