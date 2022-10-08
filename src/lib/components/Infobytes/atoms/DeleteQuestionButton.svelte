<script lang="ts">
  import { Confirm } from "svelte-confirm";
  import { currentQuestionName, Infobyte } from '$lib/stores/stores'
  import { createOrUpdateInfobyte } from "../InfobyteService";

  export let infobyte: Infobyte;
  export let selectedQuestionIndex: number;

  let open = false;
  const toggle = () => {
    open = !open;
    currentQuestionName.set(null);
  };

  const handleDelete = () => {
    infobyte.questions = infobyte.questions.filter(
      (q, index) => index !== selectedQuestionIndex
    );
    createOrUpdateInfobyte(infobyte).then(() => (open = true));
  };
</script>

<Confirm
  let:confirm={confirmThis}
  confirmTitle="Frage Löschen"
  cancelTitle="Abbrechen"
>
  <button color="danger" on:click={confirmThis(handleDelete)}>Löschen</button>
  <span slot="title">
    Bist du dir sicher dass du diese Frage löschen möchtest?
  </span>
  <span slot="description">
    Das Löschen eines Frage kann aktuell nicht rückgängig gemacht werden!
  </span>
</Confirm>

<button on:click={toggle}  >
  <h1>Frage wurde erfolgreich gelöscht!</h1>
  <h3>
    Das Frage konnte erfolgreich gelöscht werden. Aktuell ist es hilfreich die
    Anwendung neu zu laden um sicher zu gehen dass auch wirklich alles richtig
    gespeichert wurde.
  </h3>
  <div>
    <button color="danger" on:click={toggle}>Schließen</button>
  </div>
</button>
