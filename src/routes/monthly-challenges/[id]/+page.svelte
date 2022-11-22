<script>
	import MonthlyChallengeEdit from '$lib/components/Challenges/MonthlyChallengeEdit.svelte';
	import { client } from '$lib/stores/stores';
	/** @type {import('./$types').PageData} */
	export let data;
	const challengeRequest = client.records.getOne('monthly_challenges', data.id);
</script>

<div>
	{#await challengeRequest then record}
		<MonthlyChallengeEdit data={record} editing />
	{/await}
</div>
