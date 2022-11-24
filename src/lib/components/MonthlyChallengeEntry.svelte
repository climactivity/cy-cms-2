<script lang="ts">
	import { goto } from '$app/navigation';
	import type { MonthlyChallenge } from './Challenges/challenges';

	export let value: MonthlyChallenge;
	export let onDelete = (value) => {};
</script>

<div class="text-sm text-gray-500 font-mono">
	{value.id}
</div>
<div>
	{value.from.slice(0, 10)}
</div>
<div>
	{value.to.slice(0, 10)}
</div>
<div>
	{value.title}
</div>
<div>
	<button
		class="w-full h-full bg-black text-white text-center font-bold rounded hover:bg-yellow-500"
		on:click={() => goto(`/monthly-challenges/${value.id}`)}
	>
		Edit
	</button>
</div>

<style lang="scss">
	div {
		@apply px-2 py-1;
	}
</style>
