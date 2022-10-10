<script>
	import ChallengeEdit from '$lib/components/Challenges/ChallengeEdit.svelte';
	import { client } from '$lib/stores/stores';
	/** @type {import('./$types').PageData} */
	export let data;
	const challengeRequest = client.records.getOne('challenges', data.id);
</script>

<div>
	{#await challengeRequest then record}
		<ChallengeEdit data={record} editing />
	{/await}
</div>
