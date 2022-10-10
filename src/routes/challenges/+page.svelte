<script lang="ts">
	import { browser } from '$app/environment';
	import ContentTypeList from '$lib/components/Editor/content-type-list.svelte';
	import { client } from '$lib/stores/stores';

	import ChallengeCard from '$lib/components/Challenges/ChallengeCard.svelte';
	export const getChallengesMetaData = async (offset, limit) => {
		const resultset = await client.records.getList('challenges', offset, limit);

		const items = resultset.items;
		return { data: items, count: resultset.totalItems };
	};
</script>

<h1>Hi</h1>

{#if browser}
	<ContentTypeList
		ListElementComponent={ChallengeCard}
		fetchContentMetaData={getChallengesMetaData}
		newPath="new"
		offset={0}
		limit={100}
	/>
{/if}
