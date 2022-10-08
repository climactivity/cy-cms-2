<script lang="ts">
    import { Answer, Question } from '$lib/stores/stores';

    export let question: Question;

    const addAnswer = () => {
        question.answers = [...question.answers, new Answer()];
    };

    const removeAnswer = (index) => {
        question.answers = question.answers.filter((u, i) => index !== i);
    };
</script>

<div>
    <label for="question">Frage</label>
    <input
        required
        type="textarea"
        name="question"
        placeholder="question"
        bind:value={question.question}
    />
</div>

<table >
    <thead>
        <tr>
            <th width={"2%"}>#</th>
            <th>Antwort</th>
            <th>Richtig</th>
            <th>Entfernen</th>
        </tr>
    </thead>
    <tbody>
        {#each question.answers as answer, index}
            <tr>
                <th scope="row">{index + 1}</th>
                <td>
                    <input
                        required
                        name={`question.answers[${index}].value`}
                        placeholder="Antwort"
                        bind:value={answer.value}
                    />
                </td>
                <td>
                    <input
                        class="checkbox"
                        type="checkbox"
                        name={`question.answers[${index}].correct`}
                        bind:checked={answer.correct}
                    />
                </td>
                <td>
                    <button
                       
                        type="button"
                        name={`question.answers[${index}].delete`}
                        on:click={() => removeAnswer(index)}
                    >
                        Entfernen
                    </button>
                </td>
            </tr>
        {/each}
    </tbody>
</table>
<button
    color="success"
    type="button"
    name="Neue Antwort hinzufügen"
    on:click={addAnswer}
>
    Neue Antwort Hinzufügen
</button>
