<script>
	import { goto } from '$app/navigation';
	import { client } from '$lib/stores/stores';
	import SidebarButton from './SidebarButton.svelte';

	let on = true;

	const signOut = async () => {
		client.authStore.clear();
		await goto('/login');
	};
</script>

<div
	class="{on
		? 'sidebar'
		: 'side'} h-screen menu bg-white p-4 flex items-start nunito flex-1 shadow z-40"
>
	<ul class="list-reset ">
		<SidebarButton iconClass="fa-home" target="/" label="Home" />

		<div class="my-2 border border-black w-8" />

		<SidebarButton iconClass="fa-swatchbook" target="/topics" label="Bereiche" />
		<SidebarButton iconClass="fa-tags" target="/tags" label="Tags" />
		<SidebarButton iconClass="fa-medal" target="/challenges" label="Challenges" />

		<div class="my-2 border border-black w-8" />

		<SidebarButton iconClass="fa-globe" target="https://nk.climactivity.de" label="Gameserver" />
		<SidebarButton
			iconClass="fa-database"
			target="https://cms2.climactivity.de/_/"
			label="Database"
		/>
		<div class="my-2 border border-black w-8" />

		<SidebarButton iconClass="fa-warning" target="/replace" label="Excel import" />
		<li title="Abmelden" class=" md:my-0 sb-elem group show-label relative ">
			<button
				on:click={(e) => signOut()}
				class="fa-solid fa-sign-out py-3 pl-1"
				target="/logout"
				label="Abmelden"
			/>
		</li>
	</ul>
</div>

<style lang="scss">
	.nunito {
		font-family: 'nunito', font-sans;
	}
	.list-reset > li > a {
		display: grid;
		grid-auto-flow: column;
		grid-template-columns: 1.5rem 1fr;
		gap: 1rem;
		place-content: center;
		place-items: center;
	}

	.list-reset > li > button {
		display: grid;
		grid-auto-flow: column;
		grid-template-columns: 1.5rem 1fr;
		gap: 1rem;
		place-content: center;
		place-items: center;
	}

	.sidebar {
		transition: ease-in-out all 0.3s;
		z-index: 9999;
		width: 64px;
		opacity: 1;
	}

	.sidebar span {
		transition: ease-in-out all 0.1s;
		opacity: 0;
	}

	.side span {
		transition: ease-in-out all 0.1s;
		opacity: 1;
	}

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
