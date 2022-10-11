<script lang="ts">
	import { activeTabStore } from '$lib/stores/stores';
	import { getContext, onMount } from 'svelte';
	import { slide } from 'svelte/transition';

	export let label = 'Section Heading';
</script>

{#if $activeTabStore === '' || $activeTabStore === label}
	<div id={label} title={label} class="bg-zinc-50 p-2 rounded-md shadow-sm my-2 mr-2">
		<div class="w-full text-lg font-bold ">
			<span>{label}</span>
		</div>
		<div class="p-2 space-y-2 ">
			<slot />
		</div>
	</div>
{/if}
