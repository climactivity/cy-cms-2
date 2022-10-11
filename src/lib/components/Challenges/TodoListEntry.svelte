<script lang="ts">
	import { fade, slide, blur } from 'svelte/transition';
	import type { Todo } from './challenges';

	export let value: Todo;
	export let values;
	export let onDelete = (e) => {};
</script>

<div class="grid grid-flow-col w-full" style="grid-template-cols: 5fr 1fr">
	<div>
		{value.name}
	</div>
	<div>
		<button on:click={onDelete}> - </button>
	</div>
</div>
