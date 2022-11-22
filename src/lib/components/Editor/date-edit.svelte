<script lang="ts">
	export let value: string = '';
	export let label: string = '';
	export let id: string;
	export let onChange = (e) => {};
	export let onInput = (e) => {};
</script>

<div>
	<label class="font-semibold" for={id}>{label}</label>
	<input type="date" />
</div>

<style lang="scss">
	.form-field {
		@apply w-full border-zinc-300 px-3 py-2 rounded-md shadow-sm;
	}
</style>
