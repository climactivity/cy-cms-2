<script>
	import SidebarButton from './SidebarButton.svelte';

	let on = true;
</script>

<div
	class="{on ? 'sidebar' : 'side'} h-screen menu bg-white p-4 flex items-start nunito  shadow z-40"
>
	<ul class="list-reset ">
		<SidebarButton iconClass="fa-solar-panel" target="/" label="Home" />

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
	</ul>
</div>

<style>
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

	.side {
		transition: ease-in-out all 0.3s;
		z-index: 9999;
		width: 9rem;
	}

	.side span {
		transition: ease-in-out all 0.1s;
		opacity: 1;
	}
	/* .side:hover {
    width: 130px;
    box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
    opacity: 1;
  }

  .side:hover span {
    opacity: 1;
  } */
</style>
