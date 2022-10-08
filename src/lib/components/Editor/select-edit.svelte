<script lang="ts" >
    export let id: string;
    export let label: string;
    export let value: any; 
    export let options: {label: string, value: any}[];
    export let placeholder: string;
</script>


<div>
    <label for={id}>{label}</label>
    <select id={id} bind:value={value} class="w-full rounded-md shadow-sm px-4 py-2">
    {#if placeholder}
        <option value={null}>{placeholder}</option>
    {/if}
    {#each options as option}
        <option value={option.value}>{option.label}</option>
    {/each}
    </select>
</div>