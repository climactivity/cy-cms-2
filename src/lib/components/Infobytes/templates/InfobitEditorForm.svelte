<script lang="ts">
  import { createForm } from "svelte-forms-lib";

  import { Infobyte, isProd } from "$lib/stores/stores";
  import DebugInfobyteOutput from "../atoms/DebugInfobyteOutput.svelte";
  import DeleteInfobitButton from "../atoms/DeleteInfobitButton.svelte";
  import { createOrUpdateInfobyte } from "../InfobyteService";
  import InfoBitEditor from "./InfoBitEditor.svelte";
  import { currentInfobit, Infobit } from "$lib/stores/stores";

  export let selectedInfobyte: Infobyte;
  export let selectedInfobitIndex: number;

  let open = false;
  const toggle = () => (open = !open);

  const { form, errors, handleSubmit } = createForm({
    initialValues: selectedInfobyte,
    onSubmit: async (values: Infobyte) => {
      let infobyte = values;
      if (infobyte._id === "") delete infobyte._id;

      createOrUpdateInfobyte(infobyte).then(() => (open = true));
      selectedInfobyte.infobits[selectedInfobitIndex] =
        $form.infobits[selectedInfobitIndex];
    },
  });

  const handleSaveAndCreate = () => {
    currentInfobit.set(new Infobit());
  };

  $: if (selectedInfobyte && selectedInfobitIndex) {
    $form.infobits = selectedInfobyte.infobits || [];
    $errors.infobits = selectedInfobyte.infobits || [];
  }
</script>

<section>
  <form on:submit={handleSubmit}>
    <div>
      <Col sm="6">
        <h4 class="mt-2">Infobyte: {$form.name}</h4>
        <h2 class="mt-2">Infobit Nr: {selectedInfobitIndex + 1}</h2>
      </Col>
      <Col sm="3">
        <Button color={"success"} type="submit">Speichern</Button>
      </Col>
      <Col sm="3">
        <Button color={"success"} on:click={handleSaveAndCreate}
          >Speichern und neues Infobit</Button
        >
      </Col>
    </div>
    <InfoBitEditor bind:value={$form.infobits[selectedInfobitIndex]} />
  </form>
  <div>
    <Col sm={{ size: 3, offset: 9 }}>
      <DeleteInfobitButton
        bind:infobitIndex={selectedInfobitIndex}
        bind:infobyte={selectedInfobyte}
      />
    </Col>
  </div>

  <DebugInfobyteOutput visible={!isProd} {form} />
</section>

<Modal isOpen={open} {toggle} transitionOptions>
  <ModalHeader>Infobit wurde gespeichert!</ModalHeader>
  <ModalBody>
    Das Infobit konnte erfolgreich gespeichert werden. Aktuell ist es hilfreich
    die Anwendung neu zu laden um sicher zu gehen dass auch wirklich alles
    richtig gespeichert wurde.
  </ModalBody>
  <ModalFooter>
    <Button color="danger" on:click={toggle}>Schließen</Button>
  </ModalFooter>
</Modal>
