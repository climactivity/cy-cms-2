<script lang="ts">
	import StringEdit from '../Editor/string-edit.svelte';
	import ReorderableList from '../ReorderableList.svelte';
	import type { Challenge, Difficulty, Todo } from './challenges';
	import TodoListEntry from './TodoListEntry.svelte';

	export let data: Todo[];

	let newTodoDiff = { name: '', index: data?.length ?? 0 };
</script>

<div class="space-y-2">
	<ReorderableList
		listItem={TodoListEntry}
		bind:value={data}
		headers={['Todos', '-']}
		indexField="index"
		listCols="5fr"
	/>
	<StringEdit id="diffTodo" type="richtext" label="Neues todo" bind:value={newTodoDiff.name} />
	<button
		class="text-black text-sm px-2 py-1 border-black border-2 rounded-lg"
		on:click={(e) => {
			if (!data || !(data instanceof Array)) {
				data = [];
			}
			data = [...data, newTodoDiff];
			newTodoDiff = { name: '', index: data?.length ?? 0 };
			console.log(data);
		}}>Todo hinzufügen <b>+</b></button
	>
</div>
