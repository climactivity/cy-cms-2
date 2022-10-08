<script lang="ts">
  import { currentInfobyte, Infobyte } from '$lib/stores/stores'

  const newInfobyte = () => {
    currentInfobyte.set(new Infobyte());
  };

  export let small = false;

  let title = small ? "+" : "Infobyte erstellen";
</script>

<button style={"min-width: 20px"} color="success" on:click={newInfobyte}
  >{title}
</button>
