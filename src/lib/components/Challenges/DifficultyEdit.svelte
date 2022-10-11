<script lang="ts">
	import TodoEdit from './TodosEdit.svelte';

	import StringEdit from '../Editor/string-edit.svelte';
	import ReorderableList from '../ReorderableList.svelte';
	import type { Challenge, Difficulties, Difficulty } from './challenges';
	import TodoListEntry from './TodoListEntry.svelte';
	import RtfEdit from '../Editor/rtf-edit.svelte';

	export let data: Difficulty;
</script>

<div class="space-y-4">
	<StringEdit id="name" label="Name" type="text" placeholder="Name" bind:value={data.name} />

	<RtfEdit id="taskDescription" label="Beschreibung" bind:value={data.taskDescription} />

	<RtfEdit
		id="upgradeText"
		label="Level geschafft Text"
		description="Wird angeizeigt wenn dieser Schwierigkeitsgrad der Challenge geschafft wurden, klopf der Nutzer:in nochmal auf die Schulter"
		bind:value={data.upgradeText}
	/>

	<TodoEdit bind:data={data.todos} />
</div>
