<script lang="ts">
	export let value;
	export let onDelete = (value) => {};
</script>

<div>
	{value.id}
</div>
<div>
	{value.topic}
</div>
<div>
	{value.label}
</div>
<div>
	<button
		class="w-6 h-full bg-black text-white text-center font-bold rounded hover:bg-red-500"
		on:click={() => onDelete(value)}
	>
		-
	</button>
</div>

<style lang="scss">
	div {
		@apply px-2 py-1;
	}
</style>
