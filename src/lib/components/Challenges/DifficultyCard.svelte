<script lang="ts">
	import { thresholdScott } from 'd3';
	import { each } from 'svelte/internal';
	import type { Difficulty } from './challenges';

	export let value: Difficulty;
</script>

<div class="grid grid-flow-col" style="grid-template-columns: 1fr 2fr 4fr 4fr 1fr 1fr;">
	<div>
		{value.index}
	</div>
	<div>
		{value.name}
	</div>
	<div>
		{value.taskDescription}
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
</div>
