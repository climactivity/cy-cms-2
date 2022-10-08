<script lang="ts">
  import { createForm } from "svelte-forms-lib";

  import { Infobyte, isProd, Question } from "$lib/stores/stores";
  import DebugInfobyteOutput from "../atoms/DebugInfobyteOutput.svelte";
  import DeleteQuestionButton from "../atoms/DeleteQuestionButton.svelte";
  import { createOrUpdateInfobyte } from "../InfobyteService";
  import QuestionInput from "../organisms/QuestionInput.svelte";

  export let selectedInfobyte: Infobyte;
  export let selectedInfobitIndex: number;
  export let selectedQuestionIndex: number;

  let open = false;
  const toggle = () => (open = !open);

  const { form, errors, handleSubmit } = createForm({
    initialValues: selectedInfobyte,
    onSubmit: async (values: Infobyte) => {
      let infobyte = values;
      if (infobyte._id === "") delete infobyte._id;

      createOrUpdateInfobyte(infobyte).then(() => (open = true));
    },
  });

  $: if (selectedInfobyte) {
    $form.questions = selectedInfobyte.questions || [new Question()];
    $errors.questions = selectedInfobyte.questions || [new Question()];
  }

  $: if (selectedInfobyte.questions[selectedQuestionIndex]) {
    $form.questions[selectedQuestionIndex] =
      selectedInfobyte.questions[selectedQuestionIndex];
  }
</script>

<section>
  <form on:submit={handleSubmit}>
    <div>
      <Col sm="9">
        <h6 class="mt-2">Infobyte: {$form.name}</h6>
        <h4 class="mt-2">Infobit Nr: {selectedInfobitIndex}</h4>
        <h2 class="mt-2">
          Frage: {selectedInfobyte.questions[selectedQuestionIndex].question}
        </h2>
      </Col>
      <Col sm="3">
        <Button color={"success"} type="submit">Speichern</Button>
      </Col>
    </div>

    <QuestionInput
      bind:question={selectedInfobyte.questions[selectedQuestionIndex]}
    />
  </form>

  <div>
    <Col sm={{ size: 3, offset: 9 }}>
      <DeleteQuestionButton
        bind:infobyte={selectedInfobyte}
        bind:selectedQuestionIndex
      />
    </Col>
  </div>

  <DebugInfobyteOutput visible={!isProd} {form} />
</section>

<Modal isOpen={open} {toggle} transitionOptions>
  <ModalHeader>Frage wurde gespeichert!</ModalHeader>
  <ModalBody>
    Das Frage konnte erfolgreich gespeichert werden. Aktuell ist es hilfreich
    die Anwendung neu zu laden um sicher zu gehen dass auch wirklich alles
    richtig gespeichert wurde.
  </ModalBody>
  <ModalFooter>
    <Button color="danger" on:click={toggle}>Schließen</Button>
  </ModalFooter>
</Modal>
