<script>
	import { client, login } from '$lib/stores/stores';

	let email, password;
</script>

<div class="mx-auto grid place-content-center h-[100vh]">
	<div class="grid grid-flow-row gap-8">
		<h1 class="text-center text-4xl font-bold">Login</h1>

		<form
			on:submit={(e) => {
				e.preventDefault();
				login(email, password);
			}}
			class="grid grid-flow-row gap-8 bg-white rounded shadow p-8"
		>
			<label for="username" class="grid grid-flow-row gap-2"
				>Benutzer
				<input
					id="username"
					type="email"
					bind:value={email}
					class="rounded bg-gray-50 border px-2 py-1"
				/>
			</label>

			<label for="password" class="grid grid-flow-row gap-2"
				>Passwort
				<input
					id="password"
					type="password"
					bind:value={password}
					class="rounded bg-gray-50 border px-2 py-1"
				/>
			</label>

			<button action="s" class="bg-black text-white font-bold rounded  px-4 py-2">Login</button>
		</form>
	</div>
</div>
