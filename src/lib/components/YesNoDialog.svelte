<script lang="ts">
    export let caption = "Sicher?",
        text = "",
        confirmAction = "Ja",
        cancelAction = "Nein";
    export let onConfirm: (e: Event) => void;
    export let confirm = () => {
        visible = !visible;
    };

    let visible = false;
</script>

{#if visible}
    <main
        class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 shadow-2xl z-[10000] bg-white"
    >
        <div class="bg-red-300 p-2">
            <span>{caption}</span>
        </div>
        {text}
        <div class="actions p-2">
            <button on:click={confirm} class="button mr-8"
                >{cancelAction}</button
            >
            <button on:click={onConfirm} class="danger button"
                >{confirmAction}</button
            >
        </div>
    </main>
{/if}
<slot {confirm} />
