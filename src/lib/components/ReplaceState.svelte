<script lang="ts">
	import { client } from '$lib/stores/stores';
	import ChallengeCard from './Challenges/ChallengeCard.svelte';

	const replaceState = async () => {
		console.log('boom');

		const data = JSON.parse(newState);
		console.log(data);
		const { tags, topics, challenges } = data;

		console.log(tags.length, topics.length, challenges.length);
		const replaceCollection = async (collection, records) => {
			const oldData = await client.records.getFullList(collection);

			// use old school for loops to not fire ~100 xhrs at once, browsers
			// don't seem to like that
			if (oldData && oldData.length > 0) {
				for (const k in oldData) {
					const record = oldData[k];

					console.log('DROP', record);
					await client.records.delete(collection, record.id);
				}
			}

			for (const k in records) {
				const record = records[k];
				console.log('ADD', record);
				await client.records.create(collection, record);
			}

			return;
		};

		await replaceCollection('tags', tags);
		await replaceCollection('topics', topics);

		const findOneByField = async (collection, field, query) => {
			const res = await client.records.getList(collection, 0, 1, { [field]: query });
			return res.items[0]?.id ?? null;
		};

		const replaceChallenges = async (records) => {
			const oldData = await client.records.getFullList('challenges');

			// use old school for loops to not fire ~100 xhrs at once, browsers
			// don't seem to like that
			if (oldData && oldData.length > 0) {
				for (const k in oldData) {
					const record = oldData[k];

					console.log('DROP', record);
					await client.records.delete('challenges', record.id);
				}
			}

			for (const k in records) {
				let record = records[k];

				// replace topic slug with id

				record.topic = await findOneByField('topics', 'label', record.topic);

				// replace tag slugs with id

				// record.tags = await Promise.all(
				// 	record.tags.map((label) => findOneByField('tags', 'label', label))
				// );

				let tagIds = [];
				for (const k in record.tags) {
					const tag = record.tags[k];

					const tagId = await findOneByField('tags', 'label', tag.label);
					tagIds = [...tagIds, tagId];
				}

				record.tags = tagIds;
				// remove images from import, i guess ill have to do this manually

				console.log('ADD', record);
				await client.records.create('challenges', record);
			}

			return;
		};

		await replaceChallenges(challenges);
	};

	let newState;
</script>

<div class="border-4 border-red-300 bg-red-100 rounded">
	<label for="replacer" class="text-center font-bold text-red-500 flex flex-col "
		>Zustand ersetzen! (Gefährlich)

		<textarea
			name="replacer"
			id="replacer"
			autocorrect="off"
			bind:value={newState}
			class="m-4 rounded text-black font-normal font-mono p-4 min-h-[30rem]"
		/>
	</label>

	<div class="float-right">
		<button
			class="bg-red-700 text-white font-bol px-2 py-1 rounded"
			on:click={(e) => replaceState()}>Do it (alle Daten werden zerstört!!)</button
		>
	</div>
</div>
