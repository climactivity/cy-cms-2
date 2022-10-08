<script lang="ts">
export let value: string = "";
export let placeholder: string = "";
export let label: string = "";
export let id: string; 
export let type = "text";
</script>

<div>
    {#if label}
        <label for="{id}">{label}</label>
    {/if}
    {#if type === "textarea"}
        <textarea id="{id}" type="text" bind:value={value} placeholder={placeholder} {...$$props} class="form-field"/>
    {:else}
        <input id="{id}" type="text" bind:value={value} placeholder={placeholder} {...$$props} class="form-field"/>
    {/if}
</div>


<style lang="scss">
.form-field {
    @apply w-full border-zinc-300 px-3 py-2 rounded-md shadow-sm 
}
</style>