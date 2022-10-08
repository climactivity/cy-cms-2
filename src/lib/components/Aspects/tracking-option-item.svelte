<script lang="ts">
import StringEdit from "../Editor/string-edit.svelte";
import YesNoDialog from "../YesNoDialog.svelte";
import type { LocalizedTrackingOption, TrackingOption } from "./AspectTypes";


    export let values: LocalizedTrackingOption[], value: LocalizedTrackingOption;
    export let index
</script>

<div>
    {value.level}
</div>
{#if value.editing}
    <div>
        <StringEdit
            bind:value={value.option}
            id="name"
            placeholder="Name"
            type="textarea"
            rows="5"
            class="w-full"
        />
    </div>
    <div>
        <StringEdit
            bind:value={value.co2value}
            id="intro"
            placeholder="intro"
            type="textarea"
            rows="5"
            class="w-full"
        />
    </div>

    <div>
        <StringEdit
            bind:value={value.waterFactor}
            id="intro"
            placeholder="intro"
            type="textarea"
            rows="5"
            class="w-full"
        />
    </div>

    <div>
        <StringEdit
            bind:value={value.reward.xp}
            id="intro"
            placeholder="intro"
            type="textarea"
            rows="5"
            class="w-full"
        />
    </div>

    <div>
        <StringEdit
            bind:value={value.reward.coins}
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
                values = values.filter((e) => e.level !== value.level)
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
        {value.option}
    </div>
    <div>
        {value.co2value}
    </div>
    <div>
        {value.waterFactor}
    </div>
    <div>
        {value.reward.xp}
    </div>
    <div>
        {value.reward.coins}
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