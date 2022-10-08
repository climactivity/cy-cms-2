<script lang="ts">
  import { fetchAspects, getInfobytes } from "../InfobyteService";
  import InfobyteSelectionList from "../molecules/InfobyteSelectionList.svelte";
  import MasterInfobyteHeader from "../molecules/MasterInfobyteHeader.svelte";
  import type {
    aspectIdentifier,
    infobyteIdentifier,
  } from "../particles/Infobyte";

  let infobytes: infobyteIdentifier[];
  let aspects: aspectIdentifier[];

  async function fetchInfobytesAndAspects() {
    infobytes = await getInfobytes();
    aspects = await fetchAspects("DE", "DE");
  }
</script>

<MasterInfobyteHeader title={"Infobytes:"} />

{#await fetchInfobytesAndAspects()}
  <p>...lade</p>
{:then}
  <InfobyteSelectionList {infobytes} {aspects} />
{:catch error}
  <p>Hmm da ist was schief gelaufen:</p>
  <pre>{error}</pre>
{/await}
