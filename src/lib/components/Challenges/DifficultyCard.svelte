<script lang="ts">
	import { thresholdScott } from 'd3';
	import { each } from 'svelte/internal';
	import YesNoDialog from '../YesNoDialog.svelte';
	import type { Difficulty } from './challenges';

	export let value: Difficulty;
	export let index;

	if (index) value.index = index;

	export let onClick = (e) => {};
	export let onDelete = () => {};
	let confirmDelete;
</script>

<div class="grid grid-flow-col" style="grid-template-columns: 1fr 2fr 4fr 4fr 1fr 1fr;">
	<div>
		{value.index}
	</div>
	<div on:click={onClick}>
		{value.name}
	</div>
	<div>
		{@html value.taskDescription}
	</div>
	<div>
		{#if value.todos && value.todos.length > 0}
			{#each value.todos as todo}
				{todo.name}
			{/each}
		{/if}
	</div>
	<div>
		{value.upgradeText}
	</div>
	<div>
		<button on:click={(e) => onDelete()} class="w-8 h-8"> - </button>
		<!-- <YesNoDialog let:confirm onConfirm={onDelete}>
		</YesNoDialog> -->
	</div>
</div>
