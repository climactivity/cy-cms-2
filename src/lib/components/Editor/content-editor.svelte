<script lang="ts">
	import { afterUpdate, getContext, setContext, onMount } from 'svelte';
	import ContentEditorActions from './content-editor-actions.svelte';

	export let data, title, titleplaceholder, editing, dataclass;
	export let saveTarget;
	let content;

	let tabs = [];
	let activeTab = { title: 'Metadata' };
	const makeTabs = () => {
		if (content.children.length > 0) {
			tabs = [];
			//ew
			for (let i = 0; i < content.children.length; i++) {
				const child = content.children[i];
				const title = child.getAttribute('title');
				tabs = [
					...tabs,
					{
						child: child,
						title: title
					}
				];
			}
		}
	};

	const setActiveTab = (tab) => {
		activeTab = tab;
	};

	// get size of static descendants here
	onMount(makeTabs);

	// get slot size (and static descendants) here
	afterUpdate(makeTabs);
</script>

<div class="layout ">
	<ContentEditorActions
		bind:data
		title={!title || title.length == 0 ? titleplaceholder : title}
		class="header"
		bind:saveTarget
		{editing}
	/>
	<div
		id="breadcrumbs"
		class="sticky top-10 h-7 z-50 bg-white py-1 text-xs border-b  border-zinc-200 mr-3"
	>
		springe zu:
		{#each tabs as tab}
			<a
				href={`#${tab.title}`}
				on:click={() => setActiveTab(tab)}
				class="text-zinc-400 bg-white border-zinc-400 border px-1 py-0.5 mx-1 rounded-md"
				class:activeTab={tab.title === activeTab.title}>{tab.title}</a
			>
		{/each}
	</div>
	<div class="content w-full " bind:this={content}>
		<slot />
	</div>
</div>

<style lang="scss">
	.layout {
		display: grid;
		width: 100%;
		height: 100%;
		grid-template-rows: 50px 30px 1fr;
		grid-template-columns: 1fr;

		grid-template-areas:
			'header'
			'breadcrumbs'
			'content';
	}

	.header {
		grid-area: header;
	}

	.breadcrumbs {
		grid-area: breadcrumbs;
	}

	.content {
		grid-area: content;
	}

	.activeTab {
		@apply text-white bg-zinc-400 border-0;
	}
</style>
