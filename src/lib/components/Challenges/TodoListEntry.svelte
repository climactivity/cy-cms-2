<script lang="ts">
	import { fade, slide, blur } from 'svelte/transition';
	import type { Todo } from './challenges';

	export let value: Todo;
	export let values;
	export let onDelete = (e) => {
		values = [...values.filter((v) => v.name !== value.name)];
	};
</script>

<div>
	{@html value.name}
</div>
<div>
	<button
		class="w-6 h-full bg-black text-white text-center font-bold rounded hover:bg-red-500"
		on:click={() => onDelete(value)}
	>
		-
	</button>
</div>
