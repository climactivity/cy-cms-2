<script lang="ts">
  import { onMount } from "svelte";
  import type { Content, ContentMetadata } from "./editor-types";
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  export let ListElementComponent;
  export let fetchContentMetaData: (
    offset: number,
    limit: number,
    query?: string
  ) => Promise<ContentMetadata>;
  export let offset: number, limit: number;
  export let searchTarget = "title";
  export let newPath = "new";
  export let listCols = "auto";
  let selectedId;
  let data: ContentMetadata;
  let selectedData: Content[] = [];
  export let searchQuery = "";
  const refetch = async () => {
    console.log("fetching list content!");
    data = await fetchContentMetaData(offset, limit, searchQuery);
  };

  const selectElement = (event: Event, elem) => {
    console.log(`selected ${elem.name}`);
    selectedId = elem._id;
    const relativeBase = $page.routeId.split("/")[0];
    goto(`/${relativeBase}/${selectedId}`, {
      replaceState: false,
    })
      .then(() => console.log("navigated"))
      .catch(() => console.log("failed"));
  };

  const newElement = (event: Event) => {
    const relativeBase = $page.routeId.split("/")[0];
    goto(`/${relativeBase}/${newPath}`, {
      replaceState: false,
    })
      .then(() => console.log("navigated"))
      .catch(() => console.log("failed"));
  };

  const search = (element) =>
    (element[searchTarget] as String) &&
    element[searchTarget].toLowerCase().includes(searchQuery.toLowerCase());

  $: selectedData = data
    ? searchQuery.length > 3
      ? data.data.filter(search)
      : data.data
    : [];

  onMount(async () => {
    refetch();
  });
</script>

<div>
  <div class="clickable cta my-2 p-2" on:click={(e) => newElement(e)}>
    <button> + hinzufügen </button>
  </div>

  <div class="clickable my-2 p-2 flex flex-row">
    <span>🔎</span>
    <input id="search" placeholder="Suche" bind:value={searchQuery} />
  </div>
  {#if data}
    <div class="grid space-y-2 space-x-2 list" style="--list-cols: {listCols}">
      {#each selectedData as elem}
        <div
          on:click|preventDefault={(e) => selectElement(e, elem)}
          class="clickable my-2 p-2 bg-slate-50"
          class:selected={selectedId === elem._id}
        >
          <svelte:component this={ListElementComponent} data={elem} />
        </div>
      {/each}
    </div>
  {:else}
    Loading...
  {/if}
</div>

<style lang="scss">
  .selected {
    @apply bg-white shadow-md transition-all;
  }

  .clickable {
    @apply cursor-pointer select-none rounded-sm;
  }

  .list {
    grid-template-columns: var(--list-cols);
  }
</style>
