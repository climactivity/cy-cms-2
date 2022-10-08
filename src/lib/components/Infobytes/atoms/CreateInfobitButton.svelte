<script lang="ts">

  import { currentInfobit, Infobit } from '$lib/stores/stores'
  const newInfobit = () => {
    currentInfobit.set(new Infobit());
  };
</script>

<button color="success" on:click={() => newInfobit()}>Infobit erstellen</button>
