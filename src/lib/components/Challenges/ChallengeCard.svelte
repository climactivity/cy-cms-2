<script lang="ts">
	import { client } from '$lib/stores/stores';
	import type { Challenge } from './challenges';

	export let data: Challenge;
</script>

{#if data}
	<div
		class="p-2 bg-zinc-50 border border-zinc-300 hover:bg-white transition flex flex-1 place-content-between gap-1"
	>
		<div>
			<span class="text-sm text-zinc-400">{data.id ?? 'No Title'}</span>
			<span class="">{data.title ?? 'No Title'}</span>
			<span>({data.slug ?? 'No slug'})</span>
		</div>

		<div>
			{#if data.image}
				<!-- svelte-ignore a11y-missing-attribute -->
				<img
					src={client.records.getFileUrl(data, data.image, { thumb: '100x100' })}
					class="w-6 h-6 inline"
				/>
			{/if}
			<span>{data.updated ?? ''}</span>
			<span class="fa-solid fa-pen" />
		</div>
	</div>
{/if}
