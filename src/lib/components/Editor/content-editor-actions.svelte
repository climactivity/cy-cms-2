<script lang="ts">
	import { goto } from '$app/navigation';

	import { ToastStyle } from '$lib/components/Toast/toast-types';
	import Toast from '$lib/components/Toast/toast.svelte';
	import { client } from '$lib/stores/stores';

	import type { Content } from './editor-types';

	export let editing;
	export let saveTarget;
	const save = async () => {
		saving = true;

		const formData = new FormData();
		Object.keys(data).map((k, i) => {
			const v = data[k];
			if (v instanceof Blob) {
				console.log("We're uploading a blob!");
				// formData.append(k, v)
			}
			formData.append(k, v);
		});

		console.log('save data!', data);

		let res;
		if (editing && formData.has('id')) {
			res = await client.records.update(saveTarget, data.id, formData);
		} else {
			res = await client.records.create(saveTarget, formData);
		}

		unsavedChanges = 1;
		return res;
	};
	export let data: Content;
	export let title = '';
	const goBack = (event) => {
		history.back();
	};
	let saving = false;
	let unsavedChanges = 0;

	$: if (data) {
		unsavedChanges = unsavedChanges + 1;
	}
</script>

<div {...$$props} class="sticky top-0 z-50">
	<div class="pb-8 mx-auto w-full flex flex-row justify-center ">
		<div class="bg-white  shadow-none bg-white w-full ">
			<nav class="flex justify-between ">
				<div class="flex flex-row col-span-1 items-center">
					<button class="text-lg my-auto  p-2" on:click|preventDefault={goBack}>🡄</button>
					<span class="text-lg my-auto">{title}</span>
					{#if unsavedChanges > 2}
						<span class="rounded-full bg-yellow-50 border border-yellow-500 px-2 py-1 mx-2 text-xs">
							ungespeicherte Änderungen
						</span>
					{/if}
					{#if editing}
						<span class="rounded-full bg-gray-50 border border-gray-500 px-2 py-1 mx-2 text-xs">
							bearbeite bestehendes Document</span
						>
					{/if}
				</div>
				<dev class="flex flex-row col-span-1 col-start-3 place-content-end mr-4">
					<label class="text-lg my-auto"
						>Veröffentlichen <input type="checkbox" bind:checked={data.published} /></label
					>

					<Toast let:triggerToast>
						<button
							class=" button cta ml-2 my-auto "
							class:saving
							on:click={async () => {
								let ok = await save();
								let style = ok ? ToastStyle.success : ToastStyle.failed;
								let message = ok ? 'Gespeichert' : 'Es ist etwas schief gelaufen 😣';
								triggerToast(style, message, 3000);
							}}>Speichern</button
						>
					</Toast>
				</dev>
			</nav>
		</div>
	</div>
</div>

<style lang="scss">
	.saving {
		@apply bg-slate-300 transition-colors;
	}
</style>
