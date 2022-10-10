<script lang="ts">
	import RichTextEditor from '../RichTextEditor.svelte';

	export let value: string = '';
	export let placeholder: string = '';
	export let label: string = '';
	export let id: string;
	export let type = 'text';
	export let onChange = (e) => {};
</script>

<div>
	{#if label}
		<label for={id}>{label}</label>
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
			class="form-field"
		/>
	{/if}
</div>

<style lang="scss">
	.form-field {
		@apply w-full border-zinc-300 px-3 py-2 rounded-md shadow-sm;
	}
</style>
