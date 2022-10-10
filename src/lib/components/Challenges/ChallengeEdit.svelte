<script lang="ts">
	import ContentEditorSection from '../Editor/content-editor-section.svelte';
	import ContentEditor from '../Editor/content-editor.svelte';
	import SelectEdit from '../Editor/select-edit.svelte';
	import StringEdit from '../Editor/string-edit.svelte';

	export let data = {
		slug: '',
		content: '',
		title: '',
		notificationText: '',
		type: null,
		image: undefined
	};

	const makeSlug = (title) => {
		const regex = /[^a-z_]/g;
		return title.split('(')[0].trim().toLowerCase().replace(regex, '_');
	};
</script>

<ContentEditor
	bind:data
	bind:title={data.title}
	titleplaceholder="Neue Challenge"
	saveTarget="challenges"
>
	<ContentEditorSection label="Metadata">
		<StringEdit
			id="title"
			label="Title"
			type="string"
			placeholder="Title"
			bind:value={data.title}
			onChange={() => {
				console.log('input has happened');
				data.slug = makeSlug(data.title);
			}}
		/>

		<StringEdit
			id="content"
			label="content"
			type="richtext"
			placeholder="content"
			bind:value={data.content}
		/>

		<SelectEdit
			id="type"
			label="type"
			placeholder="type"
			bind:value={data.type}
			options={[
				{ value: 'recurring', label: 'wöchentlich' },
				{ value: 'one-time', label: 'einmal' },
				{ value: 'repeatable', label: 'wiederholbar' }
			]}
		/>

		<StringEdit
			id="notificationText"
			label="notificationText"
			type="string"
			placeholder="notificationText"
			bind:value={data.notificationText}
		/>
	</ContentEditorSection>
</ContentEditor>

<pre>
    {JSON.stringify(data, null, 2)}
</pre>
