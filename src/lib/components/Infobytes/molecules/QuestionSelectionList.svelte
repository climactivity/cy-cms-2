<script lang="ts">

    import { currentQuestionName, Question } from '$lib/stores/stores';
    import CreateQuestionButton from "../atoms/CreateQuestionButton.svelte";
    import QuestionSelectionListItem from "../atoms/QuestionSelectionListItem.svelte";

    export let questions: Question[];

    let selectedQuestionName: string;

    currentQuestionName.subscribe((name) => (selectedQuestionName = name));
</script>

<ul >
    {#each questions as question}
        <QuestionSelectionListItem
            {question}
            active={question.question === selectedQuestionName}
        />
    {/each}
    <CreateQuestionButton />
</ul>
