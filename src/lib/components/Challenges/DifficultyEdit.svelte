<script>
	import StringEdit from '../Editor/string-edit.svelte';

	export let data;
</script>

<div>
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
</div>
