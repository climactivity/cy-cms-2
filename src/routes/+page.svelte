<script lang="ts">
	import welcome from '$lib/images/svelte-welcome.webp';
	import welcome_fallback from '$lib/images/svelte-welcome.png';

	import PocketBase from 'pocketbase';
	import { findRecord } from '$lib/find-record';
	import { client } from '$lib/stores/stores';
	import ReplaceState from '$lib/components/ReplaceState.svelte';

	// const records =  client.records.getFullList('test', 200 /* batch size */, {
	// 	sort: '-created',
	// });
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<h1 class="font-bold text-2xl py-4">Home</h1>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	h1 {
		width: 100%;
	}

	.welcome {
		display: block;
		position: relative;
		width: 100%;
		height: 0;
		padding: 0 0 calc(100% * 495 / 2048) 0;
	}

	.welcome img {
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		display: block;
	}
</style>
