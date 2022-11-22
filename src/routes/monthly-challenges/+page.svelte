<script lang="ts">
	import { browser } from '$app/environment';
	import { client } from '$lib/stores/stores';

	import TopicEdit from '$lib/components/TopicEdit.svelte';
	import TopicListEntry from '$lib/components/TopicListEntry.svelte';
	import MonthlyChallengeEntry from '$lib/components/MonthlyChallengeEntry.svelte';
	import { goto } from '$app/navigation';

	let monthlyChallenges = client.records.getFullList('monthly_challenges');

	const update = () => {
		monthlyChallenges = client.records.getFullList('monthly_challenges');
	};
</script>

<h1 class="font-bold text-2xl py-4">Monatschallenges</h1>
{#await monthlyChallenges}
	Lade Monatschallenges...
{:then data}
	{#if browser}
		<div
			class="list grid  items-start grid-cols-4 list gap-x-2 gap-y1 border"
			style="grid-template-columns: 1fr 2fr 2fr 4fr 1fr;"
		>
			<div class="font-bold border-b px-2 py-1">ID</div>
			<div class="font-bold border-b px-2 py-1">from</div>
			<div class="font-bold border-b px-2 py-1">to</div>
			<div class="font-bold border-b px-2 py-1">title</div>
			<div class="font-bold border-b px-2 py-1">Edit</div>

			{#each data as challenge}
				<MonthlyChallengeEntry value={challenge} />
			{/each}
		</div>
	{/if}
{:catch e}
	<pre class="bg-red-100">{JSON.stringify(e, null, 2)}</pre>
{/await}

<div class=" w-full flex place-content-center py-4">
	<button
		class="text-black text-lg px-4 py-2 border-black border-2 rounded-lg"
		on:click={() => goto(`/monthly-challenges/new`)}>Monatschallenge hinzufügen <b>+</b></button
	>
</div>

<style>
	.list > div:nth-child(8n + 5),
	.list > div:nth-child(8n + 6),
	.list > div:nth-child(8n + 7),
	.list > div:nth-child(8n + 8) {
		background: #ddd;
	}
</style>
