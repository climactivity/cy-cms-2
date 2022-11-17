<script lang="ts">
	import { fade, slide, blur } from 'svelte/transition';
	import RtfEdit from '../Editor/rtf-edit.svelte';
	import type { Todo } from './challenges';

	export let value: Todo;
	export let values;
	export let onDelete = (e) => {
		values = [...values.filter((v) => v.name !== value.name)];
	};

	let editing = false;
</script>

{#if editing}
	<RtfEdit
		id="upgradeText"
		label="Alle Level geschafft Text"
		description="Wird angeizeigt wenn alle Schwierigkeitsgrade der Challenge geschafft wurden, klopf der Nutzer:in nochmal auf die Schulter"
		bind:value={value.name}
		recommendedLength={100}
	/>
{:else}
	<div>
		{@html value.name}
	</div>
{/if}

<div>
	<button
		class="w-6 h-full bg-black text-white text-center font-bold rounded hover:bg-yellow-500"
		on:click={() => (editing = true)}
	>
		E
	</button>
</div>

<div>
	<button
		class="w-6 h-full bg-black text-white text-center font-bold rounded hover:bg-red-500"
		on:click={() => onDelete(value)}
	>
		-
	</button>
</div>
