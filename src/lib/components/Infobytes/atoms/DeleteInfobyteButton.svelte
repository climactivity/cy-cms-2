<script lang="ts">
  import { Confirm } from "svelte-confirm";

  import { currentInfobyte } from '$lib/stores/stores'
  import { deleteInfobyte } from "../InfobyteService";

  export let infobyteId: string;

  let open = false;
  const toggle = () => {
    open = !open;
    currentInfobyte.set(null);
  };

  const handleDelete = () => {
    if (!infobyteId) return;

    deleteInfobyte(infobyteId).then(() => (open = true));
  };
</script>

<Confirm
  let:confirm={confirmThis}
  confirmTitle="Infobyte Löschen"
  cancelTitle="Abbrechen"
>
  <button color="danger" on:click={confirmThis(handleDelete)}>Löschen</button>
  <span slot="title">
    Bist du dir sicher dass du dieses Infobyte löschen möchtest?
  </span>
  <span slot="description"
    >Das löschen eines Infobytes kann aktuell nicht rückgängig gemacht werden!
  </span>
</Confirm>

<button on:click={toggle} >
  <h1>Infobyte wurde erfolgreich gelöscht!</h1>
  <h3>
    Das Infobyte konnte erfolgreich gelöscht werden. Aktuell ist es hilfreich
    die Anwendung neu zu laden um sicher zu gehen dass auch wirklich alles
    richtig gespeichert wurde.
  </h3>
  <div>
    <button color="danger" on:click={toggle}>Schließen</button>
  </div>
</button>
