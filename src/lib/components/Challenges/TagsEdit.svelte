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
		<div>
			{#each options as tag}
				<label>
					{tag.label}
					<input
						type="checkbox"
						checked={value.some((v) => v === tag.id)}
						on:change={(e) => {
							let checked = e.target.checked;
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
