<script lang="ts">
	import { client } from '$lib/stores/stores';
	import { onMount } from 'svelte';

	export let label: string = '';
	export let id: string;
	let files;
	export let value;

	export let record;
</script>

<div>
	{#if label}
		<label class="font-semibold" for={id}>{label}</label>
	{/if}

	{#if value}
		<div>
			{#if !(value.item instanceof Function)}
				Aktuelles Bild: {value}

				<img src={client.records.getFileUrl(record, value)} class="w-20 h-20" />
			{:else}
				neues Bild:
				<img src={URL.createObjectURL(value.item(0))} class="w-20 h-20" />
			{/if}
		</div>
	{/if}
	<input
		{id}
		type="file"
		bind:files
		class="form-field"
		on:change={() => {
			value = files;
		}}
	/>
</div>

<!-- <pre>
	{JSON.stringify(value)}
</pre> -->
<style lang="scss">
	.form-field {
		@apply w-full border-zinc-300 px-3 py-2 rounded-md shadow-sm;
	}
</style>
