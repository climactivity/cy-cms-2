<script lang="ts">
	import { onMount } from 'svelte';
	import type { Content, ContentMetadata } from './editor-types';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import Record from 'pocketbase';
	import Searchbar from '../header/Searchbar.svelte';

	export let ListElementComponent;

	export let fetchContentMetaData: (offset: number, limit: number, query?: string) => Promise<any>;

	export let offset: number, limit: number;
	export let searchTarget = 'title';
	export let newPath = 'new';
	export let listCols = 'auto';

	let selectedId;
	let data: ContentMetadata;
	let selectedData: Content[] = [];

	export let searchQuery = '';

	const refetch = async () => {
		console.log('fetching list content!');
		data = await fetchContentMetaData(offset, limit, searchQuery);
	};

	const selectElement = (event: Event, elem) => {
		console.log(`selected ${elem.name}`);
		selectedId = elem.id;
		const relativeBase = $page.routeId.split('/')[0];
		goto(`/${relativeBase}/${selectedId}`, {
			replaceState: false
		})
			.then(() => console.log('navigated'))
			.catch(() => console.log('failed'));
	};

	const newElement = (event: Event) => {
		const relativeBase = $page.routeId.split('/')[0];
		goto(`/${relativeBase}/${newPath}`, {
			replaceState: false
		})
			.then(() => console.log('navigated'))
			.catch(() => console.log('failed'));
	};

	const search = (element) =>
		(element[searchTarget] as String) &&
		element[searchTarget].toLowerCase().includes(searchQuery.toLowerCase());

	$: selectedData = data ? (searchQuery.length > 3 ? data.data.filter(search) : data.data) : [];

	onMount(async () => {
		refetch();
	});
</script>

<div>
	<div
		class="grid grid-flow-col border-b-2 border-gray-300"
		style="grid-template-columns: 5fr 2fr;"
	>
		<Searchbar bind:searchQuery />
		<div
			class="cursor-pointer w-full bg-white m-4 mb-6 text-black border-2 border-black font-bold grid items-center rounded-lg"
			on:click={(e) => newElement(e)}
		>
			<button>Hinzufügen +</button>
		</div>
	</div>

	{#if data}
		<div class="grid space-y-2 space-x-2 list" style="--list-cols: {listCols}">
			{#each selectedData as elem}
				<div
					on:click|preventDefault={(e) => selectElement(e, elem)}
					class="clickable my-2 p-2 bg-slate-50"
					class:selected={selectedId === elem['id']}
				>
					<svelte:component this={ListElementComponent} data={elem} />
				</div>
			{/each}
		</div>
	{:else}
		{JSON.stringify(data)}
		no Items found
	{/if}
</div>

<style lang="scss">
	.selected {
		@apply bg-white shadow-md transition-all;
	}

	.clickable {
		@apply cursor-pointer select-none rounded-sm;
	}

	.list {
		grid-template-columns: var(--list-cols);
	}
</style>
