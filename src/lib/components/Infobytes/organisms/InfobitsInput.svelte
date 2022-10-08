<script lang="ts">
    import InfoBitEditor from "../templates/InfoBitEditor.svelte";
    import { Infobit } from '$lib/stores/stores'

    export let infobits: any[];
    export let errorInfobits: any;

    const addInfobit = () => {
        infobits = infobits.concat(new Infobit());
        errorInfobits = errorInfobits.concat(new Infobit());
    };

    const removeInfobit = (i) => () => {
        infobits = infobits.filter((_, j) => j !== i);
        errorInfobits = errorInfobits.filter((_, j) => j !== i);
    };
</script>

{#if infobits.length !== 0}
    {#each infobits as infobit, i}
        <InfoBitEditor bind:value={infobit} />
        <div class="form-control-row">
            {#if i === infobits.length - 1}
                <button
                    type="button"
                    class="form-control-button"
                    on:click={addInfobit}>+</button
                >
            {/if}
            {#if infobits.length !== 1}
                <button
                    type="button"
                    class="form-control-button"
                    on:click={removeInfobit(i)}>-</button
                >
            {/if}
        </div>
    {/each}
{:else}
    <div 
        >Klicke auf 'Infobit hinzufügen +' um ein Infobit hinzu zufügen!</div
    >
{/if}
{#if infobits.length === 0}
    <button type="button" class="form-control-button" on:click={addInfobit}
        >Infobit hinzufügen +</button
    >
{/if}

<style>
    .form-control-button {
        min-width: auto;
        margin-right: 1em;
    }
</style>
