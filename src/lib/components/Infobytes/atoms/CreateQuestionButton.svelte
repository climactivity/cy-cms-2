<script lang="ts">
  import { currentQuestion, Question } from '$lib/stores/stores'

  const newQuestion = () => {
    currentQuestion.set(new Question());
  };
</script>

<button color="success" on:click={() => newQuestion()}>Frage erstellen</button>
