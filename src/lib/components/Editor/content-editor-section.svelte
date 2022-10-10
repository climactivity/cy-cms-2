<script lang="ts">
	import { getContext, onMount } from 'svelte';

	export let label = 'Section Heading';
</script>

<div id={label} title={label} class="bg-zinc-50 p-2 rounded-md shadow-sm my-2 mr-2">
	<div class="w-full text-lg font-bold ">
		<span>{label}</span>
	</div>
	<div class="p-2 space-y-2 ">
		<slot />
	</div>
</div>
