<script lang="ts">
	import { client } from '$lib/stores/stores';

	export let id: string;
	export let label: string;
	export let value: any;
	export let placeholder: string;
	let topics = client.records.getFullList('topics');
</script>

<div>
	<label for={id}>{label}</label>
	{#await topics then options}
		<select {id} bind:value class="w-full rounded-md shadow-sm px-4 py-2">
			{#if placeholder}
				<option value={null}>{placeholder}</option>
			{/if}
			{#each options as option}
				<option value={option.id}>{option.label}</option>
			{/each}
		</select>
	{/await}
</div>
