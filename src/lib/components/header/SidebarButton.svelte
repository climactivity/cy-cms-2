<script lang="ts">
	export let target, label, iconClass;
</script>

<li title={label} class="my-2 md:my-0 sb-elem group show-label relative">
	<a
		href={target}
		class="block py-1 md:py-3 pl-2 align-middle text-black no-underline hover:text-zinc-400"
	>
		<i class="fa-solid {iconClass}" />
	</a>
</li>

<style lang="scss">
	.show-label {
		&::after {
			content: attr(title);
			position: absolute;
			top: 0.5rem;
			left: 1.5rem;
			background: black;
			color: white;
			padding: 0.25rem 0.5rem;
			border-radius: 0.25rem;
			font-size: smaller;
			transition: all 75ms;
			opacity: 0;
		}
	}

	.show-label:hover {
		&::after {
			opacity: 1;
			left: 2.5rem;
			transition-duration: 150ms;
		}
	}
</style>
