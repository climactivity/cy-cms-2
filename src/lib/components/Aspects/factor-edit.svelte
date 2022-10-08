<script lang="ts">
    import type { LocalizedFactor } from "./AspectTypes";
    import FactorListItem from "./factor-list-item.svelte";
    import ReorderableList from "../ReorderableList.svelte";

    export let value: LocalizedFactor[] = [];

    const addFactor = (event) => {
        value = [
            ...value,
            { id: value.length, intro: "", name: "", editing: true },
        ];
    };
</script>

<ReorderableList
    listItem={FactorListItem}
    bind:value={value}
    headers={["ID", "Name", "Einleitung"]}
    listCols="1fr 3fr 7fr 1fr"
/>
<div class="grid col-span-4 place-content-end">
    <button class="button cta w-20 m-2 " on:click={addFactor}>+</button>
</div>
