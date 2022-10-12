<script lang="ts">
	import RichTextEditor from '../RichTextEditor.svelte';

	export let value: string = '';
	export let label: string = '';
	export let id: string;
	export let readonly = false;
	export let description: string = '';
	import {
		Dialog,
		DialogOverlay,
		DialogTitle,
		DialogDescription
	} from '@rgossiaux/svelte-headlessui';
	import { fade, slide } from 'svelte/transition';
	let isOpen = false;

	let y;

	let returnToThisY;

	const open = (e) => {
		returnToThisY = y;
		isOpen = true;
	};

	const close = (e) => {
		y = returnToThisY;
		isOpen = false;
	};

	let textLenght = 0;
</script>

<svelte:window bind:scrollY={y} />

<Dialog open={isOpen} on:close={close}>
	<DialogOverlay
		style={'position: absolute; inset: 0; background-color: rgb(0 0 0); opacity: 0.3; z-index: 51'}
	/>

	<div class="absolute inset-56 bg-white rounded-lg p-4 space-y-4 shadow-xl z-[52]">
		<DialogTitle class="text-xl font-bold">
			<div class="flex flex-row place-content-between">
				<div>
					{label}
				</div>
				<div>
					<button on:click={close} class="fa-solid fa-close px-2" />
				</div>
			</div>
		</DialogTitle>
		<DialogDescription>{description}</DialogDescription>

		<div class="h-[75%]">
			<RichTextEditor bind:value bind:textLenght />
		</div>

		{#if textLenght > 1}
			<div transition:fade class="text-right">{textLenght} Zeichen</div>
		{/if}
	</div>
</Dialog>

<div class:readonly>
	{#if label}
		<label class="font-semibold" for={id}>{label} {readonly ? '(readonly)' : ''}</label>
	{/if}
	<button on:click={open} class="fa-solid fa-pen hover:text-zinc-400" />

	<div class="prose bg-white min-w-full rounded p-2 mt-2">
		{@html value}
	</div>
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
