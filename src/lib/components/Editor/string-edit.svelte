<script lang="ts">
	import RichTextEditor from '../RichTextEditor.svelte';

	export let value: string = '';
	export let placeholder: string = '';
	export let label: string = '';
	export let id: string;
	export let type = 'text';
	export let onChange = (e) => {};
	export let onInput = (e) => {};
	export let readonly = false;
</script>

<div class:readonly>
	{#if label}
		<label class="font-semibold" for={id}>{label} {readonly ? '(readonly)' : ''}</label>
	{/if}
	{#if type === 'textarea'}
		<textarea {id} type="text" bind:value {placeholder} {...$$props} class="form-field" />
	{:else if type === 'richtext'}
		<RichTextEditor bind:value />
	{:else}
		<input
			{id}
			type="text"
			bind:value
			{placeholder}
			{...$$props}
			on:change={onChange}
			on:input={onInput}
			class="form-field"
			{readonly}
		/>
	{/if}
</div>

<style lang="scss">
	.form-field {
		@apply w-full border-zinc-300 px-3 py-2 rounded-md shadow-sm;
	}

	.readonly {
		@apply text-gray-400 grid grid-flow-col;
		grid-template-columns: 1fr 4fr;
	}
</style>
