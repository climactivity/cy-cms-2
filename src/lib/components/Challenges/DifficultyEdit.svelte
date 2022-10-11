<script lang="ts">
	import TodoEdit from './TodosEdit.svelte';

	import StringEdit from '../Editor/string-edit.svelte';
	import ReorderableList from '../ReorderableList.svelte';
	import type { Challenge, Difficulty } from './challenges';
	import TodoListEntry from './TodoListEntry.svelte';

	export let data: Difficulty;
</script>

<div class="space-y-4">
	<StringEdit id="name" label="Name" type="text" placeholder="Name" bind:value={data.name} />

	<StringEdit
		id="taskDescription"
		label="Beschreibung"
		type="richtext"
		placeholder="taskDescription"
		bind:value={data.taskDescription}
	/>

	<StringEdit
		id="upgradeText"
		label="Level geschafft Text"
		type="richtext"
		placeholder="upgradeText"
		bind:value={data.upgradeText}
	/>

	<TodoEdit bind:data={data.todos} />
</div>
