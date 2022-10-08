<script lang="ts">
  import { Confirm } from "svelte-confirm";

  import { currentInfobitIndex, Infobyte } from '$lib/stores/stores'
  import { createOrUpdateInfobyte } from '../InfobyteEditor.svelte';

  export let infobitIndex: number;
  export let infobyte: Infobyte;

  let open = false;
  const toggle = () => {
    open = !open;
    currentInfobitIndex.set(null);
  };

  const handleDelete = () => {
    infobyte.infobits = infobyte.infobits.filter(
      (i, index) => index !== infobitIndex
    );
    createOrUpdateInfobyte(infobyte).then(() => (open = true));
  };
</script>

<Confirm
  let:confirm={confirmThis}
  confirmTitle="Infobit Löschen"
  cancelTitle="Abbrechen"
>
  <Button color="danger" on:click={confirmThis(handleDelete)}>Löschen</Button>
  <span slot="title">
    Bist du dir sicher dass du dieses Infobit löschen möchtest?
  </span>
  <span slot="description">
    Das löschen eines Infobits kann aktuell nicht rückgängig gemacht werden!
  </span>
</Confirm>

<Modal isOpen={open} {toggle} transitionOptions>
  <ModalHeader>Infobit wurde erfolgreich gelöscht!</ModalHeader>
  <ModalBody>
    Das Infobit konnte erfolgreich gelöscht werden. Aktuell ist es hilfreich die
    Anwendung neu zu laden um sicher zu gehen dass auch wirklich alles richtig
    gespeichert wurde.
  </ModalBody>
  <ModalFooter>
    <Button color="danger" on:click={toggle}>Schließen</Button>
  </ModalFooter>
</Modal>
