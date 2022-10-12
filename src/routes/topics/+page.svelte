<script lang="ts">
	import { browser } from '$app/environment';
	import { client } from '$lib/stores/stores';

	import TopicEdit from '$lib/components/TopicEdit.svelte';
	import TopicListEntry from '$lib/components/TopicListEntry.svelte';

	let topics = client.records.getFullList('topics');

	const update = () => {
		topics = client.records.getFullList('topics');
	};
	const onDelete = async (topic) => {
		console.log('delete topic', topic.id);
		await client.records.delete('topics', topic.id);
		update();
	};
</script>

<h1 class="font-bold text-2xl py-4">Bereiche</h1>
{#await topics}
	Lade Bereiche...
{:then data}
	{#if browser}
		<div
			class="list grid  items-start grid-cols-4 list gap-x-2 gap-y1 border"
			style="grid-template-columns: 1fr 4fr 4fr 1fr;"
		>
			<div class="font-bold border-b px-2 py-1">ID</div>
			<div class="font-bold border-b px-2 py-1">slug</div>
			<div class="font-bold border-b px-2 py-1">Bereich</div>
			<div class="font-bold border-b px-2 py-1">Löschen</div>
			{#each data as topic}
				<TopicListEntry value={topic} {onDelete} />
			{/each}
		</div>
	{/if}
{:catch e}
	<pre class="bg-red-100">{JSON.stringify(e, null, 2)}</pre>
{/await}

<TopicEdit afterSubmit={update} />

<style>
	.list > div:nth-child(8n + 5),
	.list > div:nth-child(8n + 6),
	.list > div:nth-child(8n + 7),
	.list > div:nth-child(8n + 8) {
		background: #ddd;
	}
</style>
