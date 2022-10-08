<script lang="ts">

  import CreateInfobyteButton from "../atoms/CreateInfobyteButton.svelte";

  export let title: string;
</script>

<div>
  <div >
    <h4 class="mt-3">{title}</h4>
  </div>
  <div>
    <CreateInfobyteButton small />
  </div>
</div>
