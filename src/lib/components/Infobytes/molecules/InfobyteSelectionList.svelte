<script lang="ts">

    import CreateInfobyteButton from "../atoms/CreateInfobyteButton.svelte";
    import type {
        aspectIdentifier,
        infobyteIdentifier,
    } from "../particles/Infobyte";
    import AspectGroupListItems from "$lib/components/Infobytes/molecules/AspectGroupListItems.svelte";

    export let infobytes: infobyteIdentifier[];
    export let aspects: aspectIdentifier[];
</script>

<ul >
    {#each aspects as aspect}
        <AspectGroupListItems
            {aspect}
            infobytes={infobytes.filter((i) => i.aspect == aspect._id)}
        />
    {/each}
    <CreateInfobyteButton />
</ul>
