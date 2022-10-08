<script lang="ts">
    import { flip } from "svelte/animate";

    export let value: any[];
    export let listItem;
    export let listCols = "1fr 3fr 7fr 1fr"
    export let headers : string[] = ["ID", "Name", "Einleitung"]
    export let indexField = "id"
    let hovering: boolean | number = false;

    //stolen from https://svelte.dev/repl/cd4d1bc127834d11812b1d156a60cdd7?version=3.20.1
    const drop = (event: DragEvent, target) => {
        event.dataTransfer.dropEffect = "move";
        const start = parseInt(event.dataTransfer.getData("text/plain"));
        const newTracklist = value;

        if (start < target) {
            newTracklist.splice(target + 1, 0, newTracklist[start]);
            newTracklist.splice(start, 1);
        } else {
            newTracklist.splice(target, 0, newTracklist[start]);
            newTracklist.splice(start + 1, 1);
        }
        newTracklist.forEach((f, i) => (f[indexField] = i));
        value = newTracklist;
        hovering = null;
    };

    const dragstart = (event, i) => {
        event.dataTransfer.effectAllowed = "move";
        event.dataTransfer.dropEffect = "move";
        const start = i;
        event.dataTransfer.setData("text/plain", start);
    };

</script>


<div class="shadow-md rounded-sm bg-white" style="--list-cols: {listCols}">
    <div class="list-item" style="--list-cols: {listCols}">
        {#each headers as colHeader}
            <div>{colHeader}</div>
        {/each}
    </div>
    {#if value}
    {#each value as f, index (f[indexField])}
        <div
            class="list-item" style="--list-cols: {listCols}"
            animate:flip
            draggable={true}
            on:dragstart={(event) => dragstart(event, index)}
            on:drop|preventDefault={(event) => drop(event, index)}
            ondragover="return false"
            on:dragenter={() => (hovering = index)}
            class:important={hovering === index}
        >
            <svelte:component this={listItem} bind:value={f} bind:values={value}/>
        </div>

    {/each}
    {/if}
</div>

<style lang="scss">
    .list-item {
        grid-template-columns: var(--list-cols);
        @apply py-2 px-4 border-b-[1px] border-solid border-zinc-200 grid grid-flow-col;
    }

    .list-item,
    ::not(:last-child) {
        @apply border-0;
    }
    .list-item > * {
        @apply mx-2 px-2;
    }
</style>