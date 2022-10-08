<script lang="ts">
	import { baseUrl } from "$lib/stores/stores";
	import type {
		FilterType,
		QuestAspectFilter,
		QuestQuestFilter,
		QuestDto,
	} from "./QuestTypes";
	import { goto } from "$app/navigation";
	import { fade } from "svelte/transition";
	import { page } from "$app/stores";

	// export let quest : QuestDto;
	export let quest = {
		deadline: "",
		maxDuration: "",
		startDate: "",
		region: "DE",
		language: "DE",
		title: "",
		text: "",
		published: false,
		questFollowup: "", // immediately following quest
		alertTrackedAspect: "", // affected aspect _id field
		linkToAfter: "",
		altIcon: "",
		userSelectable: true,
		triggerTrackingUpdate: false,
		numRepeat: "1",
		repeatAfter: "0",
	};

	let showToast = false;
	const flashToast = (duration) => {
		showToast = true;
		setTimeout(() => {
			showToast = false;
		}, duration);
	};

	const handleSubmit = async (e) => {
		e.preventDefault();

		const response = await fetch(`${baseUrl}quest-management`, {
			credentials: "include",
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(quest),
		});

		if (!response.ok) {
			console.log(response);
			alert("It broke");
			return;
		} else {
			console.log("yeeah", quest);
		}

		quest = { ...quest, ...(await response.json()) };

		flashToast(500);
		setTimeout(() => {
			const relativeBase = $page.routeId.split("/")[0];
			goto(`/${relativeBase}/`, {
				replaceState: false,
			});
		}, 600);
	};

	export let title = "";
</script>

<!-- Toaster -->
<div
	transition:fade
	class="{showToast
		? 'p-3 mb-3 absolute top-2 right-2'
		: 'hidden'} flex w-90 max-w-sm mx-auto overflow-hidden  bg-slate-100 rounded-lg shadow-md dark:bg-gray-800 opacity-90"
>
	<div class="flex items-center justify-center w-8 bg-gradient-to-l from-green-400 to-blue-500 ">
		<svg
			class="w-6 h-6 text-white fill-current"
			viewBox="0 0 40 40"
			xmlns="http://www.w3.org/2000/svg"
		>
			<path
				d="M20 3.33331C10.8 3.33331 3.33337 10.8 3.33337 20C3.33337 29.2 10.8 36.6666 20 36.6666C29.2 36.6666 36.6667 29.2 36.6667 20C36.6667 10.8 29.2 3.33331 20 3.33331ZM16.6667 28.3333L8.33337 20L10.6834 17.65L16.6667 23.6166L29.3167 10.9666L31.6667 13.3333L16.6667 28.3333Z"
			/>
		</svg>
	</div>

	<div class="px-4 py-2 -mx-3">
		<div class="mx-3">
			<span class="font-semibold text-green-500 dark:text-emerald-400"
				>Speichert</span
			>
			<p class="text-sm text-gray-600 dark:text-gray-200">Aufgabe Speichert!</p>
		</div>
	</div>
</div>

<h1 class="text-2xl text-center">{title ? title : "Neue Aufgabe erstellen"}</h1>
<form on:submit={handleSubmit} class="form">
	<div class="grid">
		<div id="">
			<label for="region">Region</label>
			<select
				id="region"
				name="region"
				type="select"
				bind:value={quest.region}
				on:change={() => (quest.region = quest.region)}
			>
				<option value={null}> -- Region -- </option>
				<option>DE</option>
			</select>
		</div>
		<div class="localization">
			<label for="language">Sprache</label>
			<select
				id="language"
				name="language"
				type="select"
				bind:value={quest.language}
				on:change={() => (quest.language = quest.language)}
			>
				<option value={null}> -- Sprache -- </option>
				<option>DE</option>
			</select>
		</div>
		<div>
			<label for="quest-title"> Titel </label>
			<input
				required
				id="quest-title"
				name="-title"
				type="text"
				bind:value={quest.title}
			/>
		</div>
		<div id="meta-input">
			<label for="linkToAfterInput"> Link nachdem fertig </label>
			<input
				id="linkToAfterInput"
				name="linkToAfterInput"
				type="string"
				bind:value={quest.linkToAfter}
			/>
		</div>
		<div>
			<label for="description">Aufgabe Description</label>
			<textarea
				type="text"
				name="description"
				id="description"
				value={quest.text}
			/>
		</div>
		<div>
			<label for={"published"} style="padding: 5px;">
				<input
					class="checkbox"
					type="checkbox"
					name={"published"}
					bind:checked={quest.triggerTrackingUpdate}
				/>
				Tracking abfragen nachdem Aufgabe bearbeit?
			</label>
		</div>

		<div id="time-input">
			<label for="startDateInput"> Start Date </label>
			<input
				required
				id="startDateInput"
				name="startDateInput"
				type="datetime-local"
				bind:value={quest.startDate}
			/>
		</div>
		<div>
			<label for="deadlineInput"> Deadline </label>
			<input
				required
				id="deadlineInput"
				name="deadlineInput"
				type="datetime-local"
				bind:value={quest.deadline}
			/>
		</div>
		<div>
			<label for="maxDurationInput"> maximale Dauer (in Stunden) </label>
			<input
				id="maxDurationInput"
				name="maxDurationInput"
				type="number"
				bind:value={quest.maxDuration}
			/>
		</div>
		<div>
			<label for="repeatAfterInput"> Wiederholen nach (in Stunden) </label>
			<input
				id="repeatAfterInput"
				name="repeatAfterInput"
				type="number"
				bind:value={quest.repeatAfter}
			/>
		</div>
		<div>
			<label for="numRepeatsInput"> Wiederholen </label>
			<input
				id="numRepeatsInput"
				name="numRepeatsInput"
				type="number"
				bind:value={quest.numRepeat}
			/>
		</div>

		<input
			type="submit"
			value="Speichern"
			class="mt-6 py-2 px-4  bg-gradient-to-l from-green-400 to-blue-500 text-white w-full transition ease-in duration-400 text-center text-base font-semibold shadow-xl rounded   hover:to-blue-700 "
		/>
	</div>
</form>

<style>
	.form label {
		display: block;
	}

	.form input,
	select {
		display: block;
		width: 100%;
		height: 40px;
		padding: 5px;
		border: 1px solid #ccc;
		border-radius: 5px;
		margin-bottom: 10px;
	}

	.form textarea {
		width: 100%;
		height: 150px;
		border: 1px solid #ccc;
		border-radius: 5px;
	}

	.grid {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 30px;
		margin-bottom: 20px;
	}

	@media (max-width: 600px) {
		.grid {
			grid-template-columns: 1fr;
		}
	}
</style>
