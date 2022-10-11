<script lang="ts">
	import { onMount } from 'svelte';
	// import Quill from "quill";
	// import Delta from "quill-delta"
	let editorElement;

	export let toolbarOptions = [
		[{ header: 1 }, { header: 2 }, 'blockquote', 'link'],
		['bold', 'italic', 'underline', 'strike'],
		[{ list: 'ordered' }, { list: 'ordered' }],
		[{ align: [] }],
		['clean']
	];

	let quill;

	let editor;

	export let value = '';
	export let textLenght;
	let initialValue;
	onMount(async () => {
		const { default: Quill } = await import('quill');
		const { default: Delta } = await import('quill-delta');

		setTimeout(() => {
			if (editorElement) editorElement.innerHTML = value;
			quill = new Quill(editorElement, {
				modules: {
					toolbar: toolbarOptions
				},
				theme: 'snow',
				placeholder: ''
			});

			editor = editorElement.querySelector('.ql-editor');
			textLenght = editor.innerText?.length ?? 0;

			quill.on('text-change', function (delta) {
				value = editor.innerHTML;
				textLenght = editor.innerText?.length ?? 0;
			});
		}, 1);
	});
</script>

<div class="editor-wrapper bg-white flex flex-col">
	<div id="editorElement" bind:this={editorElement} />
</div>

<style>
	@import 'https://cdn.quilljs.com/1.3.6/quill.snow.css';

	.editor-wrapper {
		max-width: inherit;
		height: 100%;
	}

	#editor-element {
	}
	.ql-editor {
		min-height: 100%;
		overflow-y: scroll;
	}
</style>
