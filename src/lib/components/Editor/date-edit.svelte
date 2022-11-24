<script lang="ts">
	import { onMount } from 'svelte';

	export let value: string;
	export let label: string = '';
	export let id: string;

	let date: string;

	const setValue = (_e) => {
		console.log(date);
		value = date;
	};

	onMount(() => {
		if (value) {
			const _date = new Date(value);
			console.log(_date);
			if (_date) {
				date = _date.toISOString().slice(0, 10);
			}
		}
	});
</script>

<div>
	<label class="font-semibold" for={id}>{label}</label>
	<input type="date" bind:value={date} on:change={setValue} />
</div>

<style lang="scss">
	.form-field {
		@apply w-full border-zinc-300 px-3 py-2 rounded-md shadow-sm;
	}
</style>
