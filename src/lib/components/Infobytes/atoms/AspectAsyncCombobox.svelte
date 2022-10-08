<script lang="ts">

    export let inputId: string;
    export let name: string;
    export let value: string;
    export let optionsPromise: Promise<any[]>;
</script>

<label for={inputId}>{name}</label>
{#await optionsPromise}

{:then data}
    <select id={inputId} name={inputId} type="select" bind:value>
        <option disabled selected value={null}> -- {name} -- </option>
        {#each data as aspect}
            <option value={aspect._id}>{aspect.name ?? aspect.title ?? aspect._id}</option>
        {/each}
    </select>
    {:catch error}
    <span>An error occurred: {error}!</span>
{/await}
