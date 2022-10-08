<script lang="ts">
    export let required: boolean = false;
    export let inputId: string;
    export let name: string;
    export let value: string;
</script>

<label for={inputId}>{name}</label>
<input {required} id={inputId} name={inputId} bind:value />
