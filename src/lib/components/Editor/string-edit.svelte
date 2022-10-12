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
	export let recommendedLength = 999;
	export let textLenght = 0;
</script>

<div class:readonly>
	{#if label}
		<label
			class="font-semibold {recommendedLength -
				(type === 'richtext' ? textLenght : value?.length ?? 0) <
			0
				? 'text-red-500'
				: ''}"
			for={id}
			>{label}
			{readonly ? '(readonly)' : ''}{#if !readonly && recommendedLength > 0}
				- {`${type === 'richtext' ? textLenght : value?.length ?? 0} / ${recommendedLength}`} Zeichen
			{/if}</label
		>
	{/if}
	{#if type === 'textarea'}
		<textarea {id} type="text" bind:value {placeholder} {...$$props} class="form-field" />
	{:else if type === 'richtext'}
		<RichTextEditor bind:textLenght bind:value {...$$props} />
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
