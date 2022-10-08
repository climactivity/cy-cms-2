<script lang="ts">
  import {
    currentInfobit,
    currentInfobitIndex,
    currentInfobyte,
    currentQuestion,
    currentQuestionName,
    Infobit,
    Infobyte,
    Question,
  } from '$lib/stores/stores';

  import InfobitSidebar from "./organisms/InfobitSidebar.svelte";
  import InfobyteSidebar from "./organisms/InfobyteSidebar.svelte";
  import QuestionSidebar from "./organisms/QuestionSidebar.svelte";
  // import InfobitEditorForm from "./templates/InfobitEditorForm.svelte";
  // import InfobyteEditor from "./templates/InfobyteEditor.svelte";
  // import QuestionEditorForm from "./templates/QuestionEditorForm.svelte";

  let selectedInfobyte: Infobyte = null;
  let selectedInfobitIndex: number = null;
  let selectedQuestionName: string = null;
  let selectedQuestionIndex: number = null;

  let selectedQuestions: Question[] = [];

  currentInfobyte.subscribe((value) => {
    selectedInfobyte = value;
  });

  currentInfobitIndex.subscribe((value) => {
    selectedInfobitIndex = value;
    selectedQuestions = selectedInfobyte?.questions?.filter(
      (q) => q.infobit === selectedInfobitIndex + 1,
    );
  });

  currentInfobit.subscribe((infobit: Infobit) => {
    if (!infobit) return;
    if (!selectedInfobyte?.infobits?.includes(infobit)) {
      selectedInfobyte?.infobits?.push(infobit);
      currentInfobitIndex.set(selectedInfobyte?.infobits?.length - 1);
    }
  });

  currentQuestion.subscribe((question: Question) => {
    if (!question) return;
    question.question = "Frage Nr: " + selectedInfobyte.questions.length;

    if (!selectedInfobyte?.questions?.includes(question)) {
      question.infobit = selectedInfobitIndex + 1;
      selectedInfobyte?.questions?.push(question);

      selectedQuestions = selectedInfobyte.questions.filter(
        (q) => q.infobit === selectedInfobitIndex + 1,
      );

      currentQuestionName.set(question.question);
    }
  });

  currentQuestionName.subscribe((value) => {
    selectedQuestionName = value;
    selectedQuestionIndex = selectedInfobyte?.questions?.findIndex(
      (q) => q.question === selectedQuestionName,
    );
    selectedQuestions = selectedInfobyte?.questions?.filter(
      (q) => q.infobit === selectedInfobitIndex + 1,
    );
  });

  function handleBack() {
    if (selectedQuestionName) {
      currentQuestionName.set(null);
    } else if (
      selectedInfobitIndex !== null &&
      selectedInfobitIndex !== undefined
    ) {
      currentInfobitIndex.set(null);
    } else {
      currentInfobyte.set(null);
    }
  }

  $: if (
    selectedInfobyte?.questions?.filter(
      (q) => q.infobit === selectedInfobitIndex + 1,
    )
  ) {
    selectedQuestions = selectedInfobyte?.questions?.filter(
      (q) => q.infobit === selectedInfobitIndex + 1,
    );
  }
</script>

<div>
  <div >
    <button
      
      color={"link"}
      disabled={!selectedInfobyte}
      on:click={handleBack}>Zurück</button
    >
  </div>
</div>
<div>
  <div >
    <container>
      {#if !selectedInfobyte}
        <InfobyteSidebar />
      {:else if selectedInfobitIndex === null || selectedInfobitIndex === undefined}
        <InfobitSidebar infobits={selectedInfobyte.infobits} />
      {:else}
        <QuestionSidebar questions={selectedQuestions} />
      {/if}
    </container>
  </div>
  <!-- <Col sm="9">
    {#if selectedQuestionName}
      <QuestionEditorForm
        bind:selectedInfobyte
        bind:selectedInfobitIndex
        bind:selectedQuestionIndex
      />
    {:else if selectedInfobitIndex !== null && selectedInfobitIndex !== undefined}
      <!-- <InfobitEditorForm bind:selectedInfobyte bind:selectedInfobitIndex /> -->
    <!-- {:else if selectedInfobyte}
      <InfobyteEditor bind:selectedInfobyte />
    {/if}
  </Col> --> -->
</div>
