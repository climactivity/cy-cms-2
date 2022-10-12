<script lang="ts">
	import { browser } from '$app/environment';
	import { client } from '$lib/stores/stores';

	import TopicListEntry from '$lib/components/TopicListEntry.svelte';
	import TagEdit from '$lib/components/TagEdit.svelte';
	import TagListEntry from '$lib/components/TagListEntry.svelte';

	let topics = client.records.getFullList('tags');

	const update = () => {
		topics = client.records.getFullList('tags');
	};
	const onDelete = async (topic) => {
		console.log('delete tag', topic.id);
		await client.records.delete('tags', topic.id);
		update();
	};
</script>

<h1 class="font-bold text-2xl py-4">Tags</h1>
{#await topics}
	Lade Tags...
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
			{#each data as tag}
				<TagListEntry value={tag} {onDelete} />
			{/each}
		</div>
	{/if}
{:catch e}
	<pre class="bg-red-100">{JSON.stringify(e, null, 2)}</pre>
{/await}

<TagEdit afterSubmit={update} />

<style>
	.list > div:nth-child(8n + 5),
	.list > div:nth-child(8n + 6),
	.list > div:nth-child(8n + 7),
	.list > div:nth-child(8n + 8) {
		background: #ddd;
	}
</style>
