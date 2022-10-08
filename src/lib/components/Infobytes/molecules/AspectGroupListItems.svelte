<script lang="ts">
  import InfobyteSelectionListItem from "../atoms/InfobyteSelectionListItem.svelte";
  import type {
    aspectIdentifier,
    infobyteIdentifier,
  } from "../particles/Infobyte";

  export let aspect: aspectIdentifier;
  export let infobytes: infobyteIdentifier[];

  let isOpen = false;
</script>

<button on:click={() => (isOpen = !isOpen)}>
  {aspect.name}
</button>

<div>
  {#each infobytes.filter((i) => i.aspect == aspect._id) as infobyte}
    <InfobyteSelectionListItem {infobyte} />
  {/each}
</div>
