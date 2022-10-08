<script lang="ts">
    export let visible: boolean;
    export let form: any;
</script>

{#if visible}
    <div>
        <span>Debug</span>
        <pre>{JSON.stringify($form, null, 2)}</pre>
    </div>
{/if}
