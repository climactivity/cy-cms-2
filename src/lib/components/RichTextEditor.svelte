<script lang="ts">
	import { onMount } from 'svelte';
	// import Quill from "quill";
	// import Delta from "quill-delta"
	let editorElement;

	export let toolbarOptions = [
		[{ header: 1 }, { header: 2 }, 'blockquote', 'link', 'image', 'video'],
		['bold', 'italic', 'underline', 'strike'],
		[{ list: 'ordered' }, { list: 'ordered' }],
		[{ align: [] }],
		['clean']
	];

	let quill;

	let editor;

	export let value = '<p>placeholder</p>';

	let initialValue;
	onMount(async () => {
		const { default: Quill } = await import('quill');
		const { default: Delta } = await import('quill-delta');

		editorElement.innerHTML = value;

		setTimeout(() => {
			quill = new Quill(editorElement, {
				modules: {
					toolbar: toolbarOptions
				},
				theme: 'snow',
				placeholder: ''
			});

			editor = editorElement.querySelector('.ql-editor');

			quill.on('text-change', function (delta) {
				value = editor.innerHTML;
			});
		}, 1);
	});
</script>

<div class="editor-wrapper bg-white">
	<div id="editorElement" bind:this={editorElement} />
</div>

<style>
	@import 'https://cdn.quilljs.com/1.3.6/quill.snow.css';

	.editor-wrapper {
		max-width: inherit;
	}
</style>
