<script lang="ts">
  import { baseUrl } from "$lib/stores/stores";

  export const refetch = () => {
    console.log("refreshing quest list");
    fetchQuests();
  };

  const fetchQuests = async () => {
    const response = await fetch(
      `${baseUrl}quest-management?${new URLSearchParams({
        r: "DE",
        l: "DE",
      })}`,
      {
        credentials: "include",
      }
    );

    return await response.json();
  };

  let questPromise: Promise<any[]> = fetchQuests();

  let term: string = "";
</script>

<div class="">
  <a
    href="/quests/neueaufgabe"
    class="block py-2 px-4  bg-gradient-to-l from-green-400 to-blue-500 text-white w-full transition ease-in duration-400 text-center text-base font-semibold shadow-xl rounded   hover:to-blue-700"
    style="width: 100%"
  >
    Aufgabe erstellen
  </a>

  <input
    class="rounded-lg border-transparent flex-1 border border-gray-300 w-full py-2 px-4 bg-white text-gray-700 placeholder-gray-400 shadow-md text-base focus:outline-none focus:ring-2 focus:ring-green-600 focus:border-transparent my-2 "
    placeholder="Aufgaben suchen.."
    bind:value={term}
  />
</div>

{#await questPromise}
  Lade ...
{:then quests}
  <div class=" text-gray-500 p-2">
    Wir haben {quests ? quests.length : ""} Aufgaben fertig!
  </div>
  {#each quests
    .sort((a, b) => {
      if (a.title < b.title) {
        return -1;
      }
      if (a.title > b.title) {
        return 1;
      }
      return 0;
    })
    .filter((aufgabe) => aufgabe.title
        .toLowerCase()
        .match(term.toLowerCase())) as quest}
    <a
      href="quests/{quest._id} "
      class="block py-1 px-2 mb-2 border-b border-gray-300"
    >
      <h3 class="inline-block font-medium text-gray-700  hover:underline">
        {quest.title}
      </h3>

      <p class=" text-sm text-gray-500  ">
        Created Date: {quest.createdAt.slice(0, 10)}, Last Modified: {quest.updatedAt.slice(
          0,
          10
        )}
      </p>
    </a>
  {/each}
{/await}

<style>
</style>
