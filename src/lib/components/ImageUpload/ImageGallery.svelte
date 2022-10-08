<script lang="ts">
  import { baseUrl } from "$lib/stores/stores";
  import ImageThumbnail from "./ImageThumbnail.svelte";
  let query = "";

  let offset = 0;
  export let pageSize = 20;
  let hasNext = true;
  let hasPrev = false;
  const fetchPage = async () => {
    const response = await fetch(
      `${baseUrl}image-upload?${new URLSearchParams({
        offset: offset + "",
        pageSize: pageSize + "",
        query: query,
      })}`,
      {
        credentials: "include",
      },
    );
    return await response.json();
  };
  export let select;
  export let imageModifiedCallback: Function;
  export const refetch = () => {
    page = fetchPage();
  };

  let page: Promise<any> = fetchPage();
  const nextPage = async () => {
    const { count } = await page;
    offset = Math.min(count, offset + pageSize);
    refetch();
  };
  const prevPage = () => {
    offset = Math.max(0, offset - pageSize);
    refetch();
  };
  const firstPage = () => {
    offset = 0;
    refetch();
  };
  const lastPage = () => {
    offset = 0;
    refetch();
  };

  const setPage = (pageNum) => {
    offset = pageNum * pageSize;
    refetch();
  };

  var timeout;
  $: if (query !== null) {
    clearInterval(timeout);
    timeout = setTimeout(function () {
      offset = 0;
      refetch();
    }, 300);
  }
</script>

<div>
  <div>
    {#await page then data}
      {#if Array.isArray(data.data)}
        <ul>
          <div class="grid grid-cols">
            {#each data.data as image}
              <ImageThumbnail {image} {imageModifiedCallback} {select} />
            {/each}
          </div>
        </ul>
        <div>
          <div style="width: 100%;">
            <button on:click={firstPage} disabled={offset == 0}>&lt;&lt;</button
            >
            <button on:click={prevPage} disabled={offset == 0}>&lt;</button>
            {#each Array(Math.ceil(data.count / pageSize)) as _, pageNum}
              <button  on:click={() => setPage(pageNum)}>
                {pageNum + 1}
              </button>
            {/each}
            <button
              on:click={nextPage}
              disabled={offset + pageSize >= data.count}>&gt;</button
            >
            <button
              on:click={lastPage}
              disabled={offset + pageSize >= data.count}>&gt;&gt;</button
            >
          </div>
          {offset + 1} bis {Math.min(offset + pageSize, data.count)} von {data.count}
          Bilder(n) insgesammt
        </div>
      {:else}
        <ul>Ein Fehler ist Aufgetreten: {data}!</ul>
      {/if}
    {/await}
  </div>
</div>

<style>
  .grid {
    display: grid;
    padding: 10px;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    max-width: 100%;
  }
</style>
