<script lang="ts">
import ContentEditorActions from "./content-editor-actions.svelte";


    export let data, title;
    export let saveTarget
</script>

<div class="layout ">
    <ContentEditorActions bind:data title={title} class="header" bind:saveTarget/>
    <div class="content w-full ">
        <slot />
    </div>
    
</div>

<style lang="scss">
.layout {
    display: grid;
    width: 100%;
    height: 100%;
    grid-template-rows: 50px 1fr;
    grid-template-columns:1fr;

    grid-template-areas:
      "header"
      "content";

  }

  .header {
    grid-area: header;
  }

  .content {
    grid-area: content;
  }

</style>