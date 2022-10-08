<script lang="ts">
    export let showDebug = false;
    export let data;
</script>

{#if data}
    <div>
        <input type="checkbox" bind:checked={showDebug} />
        <span>DEBUG</span>
        <pre
            class="{showDebug
                ? ''
                : 'hidden'} overflow-x-scroll max-w-3xl debug">{JSON.stringify(
                data,
                null,
                2
            )}</pre>
    </div>
{/if}

<style lang="scss">
    .debug {
        white-space: pre-wrap;
    }
</style>
