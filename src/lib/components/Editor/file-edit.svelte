<script lang="ts">
	export let value;
	export let label: string = '';
	export let id: string;
</script>

<div>
	{#if label}
		<label class="font-semibold" for={id}>{label}</label>
	{/if}
	<input {id} type="file" bind:value {...$$props} class="form-field" />
</div>

<style lang="scss">
	.form-field {
		@apply w-full border-zinc-300 px-3 py-2 rounded-md shadow-sm;
	}
</style>
