<script lang="ts">
	import { client } from '$lib/stores/stores';
	import RichTextEditor from '../RichTextEditor.svelte';

	export let value: string[];
	export let placeholder: string = '';
	export let label: string = '';
	export let id: string;
	let tags = client.records.getFullList('tags');
</script>

{#await tags then options}
	<div>
		{#if label}
			<label class="font-semibold" for={id}>{label}</label>
		{/if}
		<div
			class="grid grid-cols-6 place-items-stretch"
			style="grid-template-columns: repeat(20px,4);"
		>
			{#each options as tag}
				<label class="bg-zinc-200 border border-zinc-400 px-1  mx-2 my-1 whitespace-nowrap">
					{tag.label}
					<input
						type="checkbox"
						checked={value ? value.some((v) => v === tag.id) : false}
						on:change={(e) => {
							let checked = e.target.checked;
							if (!value) {
								value = [];
							}
							if (checked) {
								value = [...value, tag.id];
							} else {
								value = value.filter((_v) => _v !== tag.id);
							}
						}}
					/>
				</label>
			{/each}
		</div>
	</div>
{/await}

<style lang="scss">
	.form-field {
		@apply w-full border-zinc-300 px-3 py-2 rounded-md shadow-sm;
	}
</style>
