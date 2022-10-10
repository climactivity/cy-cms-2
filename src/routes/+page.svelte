<script lang="ts">
	import welcome from '$lib/images/svelte-welcome.webp';
	import welcome_fallback from '$lib/images/svelte-welcome.png';

	import PocketBase from 'pocketbase';
	import { findRecord } from '$lib/find-record';

	const client = new PocketBase('http://127.0.0.1:8090');

	// const records =  client.records.getFullList('test', 200 /* batch size */, {
	// 	sort: '-created',
	// });

	const authData = client.users.authViaEmail('test@climactivity.de', '12345678');

	const findThing = findRecord(client, 'test_challenge');
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<h1>
		<span class="welcome">
			<picture>
				<source srcset={welcome} type="image/webp" />
				<img src={welcome_fallback} alt="Welcome" />
			</picture>
		</span>

		to your new<br />SvelteKit app
	</h1>

	<div class="text-green-600">Hello</div>
	{#await authData then userData}
		you are logged in as {JSON.stringify(userData.user.email)}
	{/await}

	<div>
		{#await findThing then thing}
			{JSON.stringify(thing)}
		{/await}
	</div>
</section>

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
