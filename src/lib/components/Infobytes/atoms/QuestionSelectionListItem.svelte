<script lang="ts">

    import { currentQuestionName, Question } from '$lib/stores/stores';

    export let question: Question;
    export let active: boolean;

    const onInfobitSelected = () => {
        if (active) {
            currentQuestionName.set(null);
        } else {
            currentQuestionName.set(question.question);
        }
    };
</script>

<button  on:click={onInfobitSelected} {active}>
    {question.question}
</button>
