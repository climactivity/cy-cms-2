<script>
	let on = true;
	const check = () => {
		on = !on;
	};
</script>

<div
	class="{on ? 'sidebar' : 'side'} h-screen menu bg-white p-4 flex items-start nunito  shadow z-40"
>
	<ul class="list-reset ">
		<li class="py-1.5 md:my-0">
			<a
				href="/"
				class="block py-1 md:py-3 pl-1 align-middle text-black no-underline hover:text-indigo-400"
			>
				<i class="fa-solid fa-solar-panel " /><span class="w-full inline-block pb-1 md:pb-0 text-sm"
					>Dashboard</span
				>
			</a>
		</li>

		<div class="my-2 border border-black w-8" />

		<li class="my-2 md:my-0">
			<a
				href="/challenges"
				class="block py-1 md:py-3 pl-1 align-middle text-black no-underline hover:text-indigo-400"
			>
				<i class="fa-solid fa-medal " /><span class="w-full inline-block pb-1 md:pb-0 text-sm"
					>Challegnes</span
				>
			</a>
		</li>

		<div class="my-2 border border-black w-8" />

		<li class="my-2 md:my-0">
			<a
				href="https://nk.climactivity.de"
				target="_blank"
				class="block py-1 md:py-3 pl-1 align-middle text-black no-underline hover:text-indigo-400"
			>
				<i class="fa-solid fa-globe  " /><span class="w-full inline-block pb-1 md:pb-0 text-sm"
					>Gameserver</span
				>
			</a>
		</li>

		<li class="my-2 md:my-0">
			<a
				href="/_/"
				target="_blank"
				class="block py-1 md:py-3 pl-1 align-middle text-black no-underline hover:text-indigo-400"
			>
				<i class="fa-solid fa-database  " /><span class="w-full inline-block pb-1 md:pb-0 text-sm"
					>Database</span
				>
			</a>
		</li>

		<div class="my-2 border border-black w-8" />

		<li class="my-2 md:my-0">
			<button
				on:click={check}
				class="block py-1 md:py-3 pl-1 align-middle text-black no-underline hover:text-indigo-400"
			>
				<i class="fa-solid fa-arrows-left-right " />
			</button>
		</li>
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
