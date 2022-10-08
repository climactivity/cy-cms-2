<script lang="ts">
    import StringEdit from "../Editor/string-edit.svelte";
    import YesNoDialog from "../YesNoDialog.svelte";

    export let values, value;
    export let index;
</script>

<div>
    {value.id}
</div>
{#if value.editing}
    <div>
        <StringEdit
            bind:value={value.name}
            id="name"
            placeholder="Name"
            type="textarea"
            rows="5"
            class="w-full"
        />
    </div>
    <div>
        <StringEdit
            bind:value={value.intro}
            id="intro"
            placeholder="intro"
            type="textarea"
            rows="5"
            class="w-full"
        />
    </div>

    <div class="grid place-content-center">
        <YesNoDialog
            let:confirm
            onConfirm={() => {
                values = values.filter((e) => e.id !== value.id);
            }}
        >
            <button
                class=""
                on:click|preventDefault={(event) => {
                    confirm();
                }}>💣</button
            >
        </YesNoDialog>
    </div>

    <div class="grid place-content-center">
        <button
            class=""
            on:click|preventDefault={(event) => {
                values.forEach((e, j) => {
                    delete e.editing;
                });
                value.editing = false;
            }}>✔</button
        >
    </div>
{:else}
    <div>
        {value.name}
    </div>
    <div>
        {value.intro}
    </div>

    <div class="grid place-content-center">
        <button
            class=""
            on:click|preventDefault={(event) => {
                values.forEach((e) => {
                    delete e.editing;
                });
                value.editing = true;
            }}>✏</button
        >
    </div>
{/if}
