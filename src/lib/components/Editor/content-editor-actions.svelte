<script lang="ts">
	import { goto } from '$app/navigation';

	import { ToastStyle } from '$lib/components/Toast/toast-types';
	import Toast from '$lib/components/Toast/toast.svelte';
	import { client } from '$lib/stores/stores';

	import type { Content } from './editor-types';

	export let jsonFields = ['difficulties', 'notificationDays'];
	export let listFields = ['tags'];

	export let editing;
	export let saveTarget;
	const save = async () => {
		saving = true;
		const formData = new FormData();

		Object.keys(data).forEach((k, i) => {
			const v = data[k];
			console.log('appending', k);
			if (v instanceof FileList) {
				console.log("We're uploading a File!");
				formData.append(k, v[0]);
				return;
			}

			if (jsonFields.some((f) => f === k)) {
				formData.append(k, JSON.stringify(v));
				return;
			}

			if (listFields.some((f) => f === k)) {
				console.log("We're setting a list!");

				const arr: any[] = v;
				console.log(arr);
				arr.forEach((element) => formData.append(k, element));
				// formData.append(k, JSON.stringify(v));
				return;
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
		if (unsavedChanges > 1) {
		}
	}

	// const preventAccidentalBack = (e) => {
	// 	if (unsavedChanges > 1) {
	// 		console.log('Argh');
	// 		e.preventDefault();
	// 		const message = 'Ungespeicherte Änderungen gehen verloren!';
	// 		e.returnValue = message;
	// 		return message;
	// 	}
	// };
</script>

<div {...$$props} class="sticky top-0 z-50">
	<div class="pb-8 mx-auto w-full flex flex-row justify-center ">
		<div class="bg-white  shadow-none bg-white w-full ">
			<nav class="flex justify-between ">
				<div class="flex flex-row col-span-1 items-center">
					<button class="text-lg my-auto  p-2" on:click|preventDefault={goBack}
						><div class="bg-zinc-100 hover:text-white hover:bg-black rounded-full w-6">
							🡄
						</div></button
					>
					<span class="text-lg my-auto">{title}</span>
					{#if unsavedChanges > 2}
						<div
							class=" fa-solid fa-pencil rounded-full bg-yellow-50 border border-yellow-500 px-2 py-1 mx-1 text-xs"
						/>
					{/if}
					{#if editing}
						<span
							class="fa-solid fa-database rounded-full bg-gray-50 border border-gray-500 px-2 py-1 mx-1 text-xs"
						/>
					{/if}
				</div>
				<dev class="flex flex-row col-span-1 col-start-3 place-content-end mr-4">
					<label
						class="text-lg my-auto select-none  px-2 py-0.5 cursor-pointer bg-white hover:text-green-500 rounded "
						>Veröffentlichen <input type="checkbox" bind:checked={data.published} /></label
					>

					<Toast let:triggerToast>
						<button
							class=" button cta ml-2 my-auto text-lg px-2 py-0.5 cursor-pointer bg-white hover:bg-green-200 border-2 border-black rounded"
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
