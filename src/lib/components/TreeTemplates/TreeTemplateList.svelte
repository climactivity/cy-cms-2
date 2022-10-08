<script lang="ts">
	import { baseUrl } from "$lib/stores/stores";
	import TreeTemplateEditor from "./TreeTemplateEditor.svelte";

	const refetch = async () => {
		const response = await fetch(`${baseUrl}tree-template`, {
			credentials: "include",
		});
		return await response.json();
	};

	class TreeTemplate {
		template_name = "tree0";
		texture_name = "tree0";
		ui_name = "Test Baum";
		preview_name = "2";
		coin_value = 0;
		bigpoint_available = ["ernährung"];
		initial_state: {
			stage: 0;
			water: 0.0;
			water_needed: 40.0;
			bigpoint: "ernährung";
			aspect: "pflanzliche_ernährung";
		};
	}

	let treeTemplates: Promise<[{ name: string; _id: string }]> = refetch();

	let selectedTreeTemplate = new TreeTemplate();

	const ontreeTemplateSelected = async (treeTemplate: {
		name: string;
		_id: string;
	}) => {
		let treeTemplateData = await fetchTreeTemplate(treeTemplate._id);
		selectedTreeTemplate = treeTemplateData;
	};

	const fetchTreeTemplate = async (id) => {
		if (!id) return;
		const response = await fetch(`${baseUrl}tree-template/${id}`);
		return await response.json();
	};

	const newtreeTemplate = () => {
		selectedTreeTemplate = new TreeTemplate();
	};
</script>

<div>
	<div>
		<div>
			{#await treeTemplates}
				<p>...waiting</p>
			{:then data}
				<ul>
					{#each data as treeTemplate}
						<li
							class="flex items-center justify-between px-4 py-2 border-b border-gray-300 dark:border-gray-700 hover:bg-gray-100 dark:bg-gray-700"
							on:click={() => ontreeTemplateSelected(treeTemplate)}
						>
							{treeTemplate.ui_name}
						</li>
					{/each}
				</ul>
			{:catch error}
				<p>An error occurred!</p>
			{/await}
		</div>
	</div>
	<button on:click={newtreeTemplate}> Add new </button>
</div>
<div>
	<TreeTemplateEditor bind:currentTemplate={selectedTreeTemplate} />
</div>
