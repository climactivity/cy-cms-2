<script lang="ts">
  import type { Question } from '$lib/stores/stores';
  import SidebarTitleHeader from '../molcules/SidebarTitleHeader.svelte';
  import QuestionSelectionList from "../molecules/QuestionSelectionList.svelte";

  export let questions: Question[];
</script>

<SidebarTitleHeader title={"Antworten:"} />

<QuestionSelectionList {questions} />
