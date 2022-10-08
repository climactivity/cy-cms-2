<script lang="ts">
    import { dataset_dev } from "svelte/internal";


    export let inputId: string;
    export let name: string;
    export let value: string;
    export let optionsPromise: Promise<any[]>;
</script>

<label for={inputId}>{name}</label>
{#await optionsPromise}
    <select id={inputId} name={inputId} type="select" >
        <option disabled selected > -- {name} -- </option>
        <option disabled >Lade...</option>
    </select>
{:then data}
    {#if Array.isArray(data)}
        {#if data.length > 0}
            <select id={inputId} name={inputId} type="select" bind:value>
                <option disabled selected value={null}> -- {name} -- </option>
                {#each data as factor}
                    <option value={factor.id}>{factor.name}</option>
                {/each}
            </select>
        {:else}
            <select id={inputId} name={inputId} type="select" >
                <option disabled> -- {name} -- </option>
                <option disabled 
                    >Wähle zu erst einen Aspekt!</option
                >
            </select>
        {/if}
    {:else}
        <select id={inputId} name={inputId} type="select">
            <option disabled > -- {name} -- </option>
            <option disabled
                >Es gibt noch keine Gesichtspunkte!</option
            >
        </select>
    {/if}
{:catch error}
    <select id={inputId} name={inputId} type="select" bind:value>
        <option disabled selected value={null}> -- {name} -- </option>
        <option disabled value={null}>An error occurred!</option>
    </select>
{/await}
