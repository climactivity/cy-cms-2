<script lang="ts">
	export let data;
</script>

{#if data}
	<div>
		{data.slug ?? 'No slug'}
	</div>
{/if}
