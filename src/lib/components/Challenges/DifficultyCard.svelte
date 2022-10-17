<script lang="ts">
	import { thresholdScott } from 'd3';
	import { each } from 'svelte/internal';
	import YesNoDialog from '../YesNoDialog.svelte';
	import type { Difficulty } from './challenges';

	export let value: Difficulty;
	export let index;

	value.index = index ?? 0;

	export let onClick = (e) => {};
	export let onDelete = () => {};
	let confirmDelete;
</script>

<div class="grid grid-flow-row bg-white p-2 my-2 cursor-pointer" style="" on:click={onClick}>
	<div>
		Stufenname: {value.name}
	</div>
	<div>
		{@html value.taskDescription}
	</div>
	<div>
		<span class="font-semibold">Todos</span>
		{#if value.todos && value.todos.length > 0}
			<ul>
				{#each value.todos as todo}
					<li class="list-disc ml-5">
						{todo.name}
					</li>
				{/each}
			</ul>
		{/if}
	</div>
	<div>
		{value.upgradeText ?? 'Stufenaufstieg?'}
	</div>
	<div>
		<button
			on:click={(e) => onDelete()}
			class="h-8 hover:bg-red-100 border hover:border-red-500 hover:text-red-500 rounded px-2 py-1"
		>
			Löschen
		</button>
		<!-- <YesNoDialog let:confirm onConfirm={onDelete}>
		</YesNoDialog> -->
	</div>
</div>
