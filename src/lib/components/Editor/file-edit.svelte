<script lang="ts">
	export let label: string = '';
	export let id: string;
	let files;
	export let value;
</script>

<div>
	{#if label}
		<label class="font-semibold" for={id}>{label}</label>
	{/if}

	{#if value}
		Aktuelles Bild: {value}
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

<style lang="scss">
	.form-field {
		@apply w-full border-zinc-300 px-3 py-2 rounded-md shadow-sm;
	}
</style>
