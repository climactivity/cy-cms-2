<script lang="ts">
  import { onMount } from "svelte";
  import { baseUrl } from "$lib/stores/stores";
  import {
    FilterType,
    QuestAspectFilter,
    QuestDto,
    QuestQuestFilter,
  } from "./QuestTypes";
  // import RichTextEditor from "$lib/components/RichTextEditor.svelte";
  import ImagePicker from "$lib/components/ImageUpload/ImagePicker.svelte";
  import RewardForm from '$lib/components/Rewards/Rewardform.svelte';

  export let questModifiedCallback;
  export let quest: QuestDto;
  let iconImg;
  let newAspectFilter: QuestAspectFilter = new QuestAspectFilter();
  let newQuestFilter: QuestQuestFilter = new QuestQuestFilter();
  const fetchAspects = async () => {
    const response = await fetch(
      `${baseUrl}localized-aspect?${new URLSearchParams({
        r: quest.region,
        l: quest.language,
      })}`,
      {
        credentials: "include",
      },
    );

    return await response.json();
  };
  let aspectsPromise: Promise<any[]> = fetchAspects();

  const fetchQuests = async () => {
    const response = await fetch(
      `${baseUrl}quest-management?${new URLSearchParams({
        r: quest.region,
        l: quest.language,
      })}`,
      {
        credentials: "include",
      },
    );

    return await response.json();
  };

  let questPromise: Promise<any[]> = fetchQuests();
  const refetchAspects = async (l, r) => {
    aspectsPromise = fetchAspects();
  };

  $: if (iconImg) {
    quest.altIcon = iconImg.relativeUrl;
  }

  var language,
    region = "DE";
  $: {
    refetchAspects(language, region);
  }
  const save = async () => {
    console.log(quest);
    const response = await fetch(
      `${baseUrl}quest-management${quest._id ? "/" + quest._id : ""}`,
      {
        credentials: "include",
        method: quest._id ? "PUT" : "POST",
        headers: {
          "Content-Type": "application/json",
          // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: JSON.stringify(quest),
      },
    );

    if (!response.ok) {
      console.log(response);
      alert("It broke");
      return;
    }
    quest = { ...quest, ...(await response.json()) };
    questModifiedCallback();
  };

  const deleteQuest = async () => {
    if (quest._id) {
      const response = await fetch(
        `${baseUrl}quest-management${quest._id ? "/" + quest._id : ""}`,
        {
          credentials: "include",
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
            // 'Content-Type': 'application/x-www-form-urlencoded',
          },
        },
      );
      if (!response.ok) {
        console.log(response);
        alert("It broke");
        return;
      }
    }
    quest = new QuestDto();
    questModifiedCallback();
  };

  const getTrackingLevel = async (
    aspectId,trackingLevel,inclueName = false,) => {
    let aspect = (await aspectsPromise).find((a) => a._id == aspectId);

    if (trackingLevel < 0) {
      if (inclueName) return { aspect: aspect.name, option: "Nicht getrackt" };
      return "Nicht getrackt";
    }
    let option = aspect?.localizedTrackingData?.options?.find(
      (opt) => opt.level >= trackingLevel,
    )?.option;
    if (!option) {
      return aspect?.localizedTrackingData?.options?.slice(-1)[0]?.option;
    }
    if (inclueName) return { aspect: aspect.name, option };
    return option;
  };

  const getNumTrackingLevels = async (aspectId) => {
    let aspect = (await aspectsPromise).find((a) => a._id == aspectId);
    return aspect?.localizedTrackingData?.options?.length();
  };
</script>

<div>
  <h1>Aufgabe bearbeiten</h1>
  <div class="wrapper">
    <div id="localization-input">
      <label for="region">Region</label>
      <select
        id="region"
        name="region"
        type="select"
        bind:value={quest.region}
        on:change={() => (region = quest.region)}
      >
        <option value={null}> -- Region -- </option>
        <option>DE</option>
      </select>
      <div class="localization">
        <label for="language">Sprache</label>
        <select
          id="language"
          name="language"
          type="select"
          bind:value={quest.language}
          on:change={() => (language = quest.language)}
        >
          <option value={null}> -- Sprache -- </option>
          <option>DE</option>
        </select>

        <label for="quest-title"> Titel </label>
        <input
          required
          id="quest-title"
          name="-title"
          type="text"
          bind:value={quest.title}
        />

        <!-- <RichTextEditor bind:value={quest.text} bind:key={quest._id} /> -->

        <div id="meta-input">
          <label for="linkToAfterInput"> Link nachdem fertig </label>
          <input
            required
            id="linkToAfterInput"
            name="linkToAfterInput"
            type="string"
            bind:value={quest.linkToAfter}
          />


          <label for={"published"} style="padding: 5px;">
            <input
              class="checkbox"
              type="checkbox"
              name={"published"}
              bind:checked={quest.triggerTrackingUpdate}
            />
            Tracking abfragen nachdem Aufgabe bearbeit?
          </label>
        </div>
      </div>
    </div>
    <div id="time-input">
      <label for="startDateInput"> Start Date </label>
      <input
        required
        id="startDateInput"
        name="startDateInput"
        type="datetime-local"
        bind:value={quest.startDate}
      />

      <label for="deadlineInput"> Deadline </label>
      <input
        required
        id="deadlineInput"
        name="deadlineInput"
        type="datetime-local"
        bind:value={quest.deadline}
      />

      <label for="maxDurationInput"> maximale Dauer (in Stunden) </label>
      <input
        required
        id="maxDurationInput"
        name="maxDurationInput"
        type="number"
        bind:value={quest.maxDuration}
      />

      <label for="repeatAfterInput"> Wiederholen nach (in Stunden) </label>
      <input
        required
        id="repeatAfterInput"
        name="repeatAfterInput"
        type="number"
        bind:value={quest.repeatAfter}
      />

      <label for="numRepeatsInput"> Wiederholen </label>
      <input
        required
        id="numRepeatsInput"
        name="numRepeatsInput"
        type="number"
        bind:value={quest.numRepeat}
      />
    </div>
    Alternatives Icon
    <ImagePicker bind:img={iconImg} />

    <div id="aspect-filters">
      <h1>Aspekt Filter</h1>
      <div>
        {#if !quest.questAspectFilter}
          Noch keine Aspektfilter
        {/if}
        {#each quest.questAspectFilter as aspectFilter, i}
          <div>
            {#await getTrackingLevel(aspectFilter.aspectId, aspectFilter.trackingLevel, true)}
              ...
            {:then option}
              {option.aspect} - {option.option}
            {/await}
            - {FilterType[aspectFilter.filterType]}
            <!--die warnung ist einfach mal falsch, cool https://www.typescriptlang.org/docs/handbook/enums.html -->
          </div>
          <button
            on:click={() => {
              quest.questAspectFilter = quest.questAspectFilter.filter(
                (_, j) => j !== i,
              );
            }}
          >
            entfernen
          </button>
        {/each}
      </div>
      <hr />
      <div>

        <select
          id="region"
          name="region"
          type="select"
          bind:value={newAspectFilter.filterType}
        >
          <option value={FilterType.REQURIED}> REQURIED </option>
          <option value={FilterType.EXCLUDED}> EXCLUDED </option>
        </select>
        <div>
          <select
            id="trackingLevel"
            name="trackingLevel"
            type="range"
            min="-1"
            max="4"
            step="1"
            bind:value={newAspectFilter.trackingLevel}
          />
          {#await getTrackingLevel(newAspectFilter.aspectId, newAspectFilter.trackingLevel)}
            ...
          {:then option}
            {option}
          {/await}
        </div>
      </div>
      <div>
        <button
          on:click={() => {
            if (
              newAspectFilter.aspectId !== undefined &&
              newAspectFilter.trackingLevel !== undefined
            ) {
              quest.questAspectFilter = [
                ...quest.questAspectFilter,
                newAspectFilter,
              ];
              newAspectFilter = new QuestAspectFilter();
            }
          }}
        >
          Aspektfilter speichern
        </button>
      </div>
    </div>

    <div id="quest-filters">
      <h1>Aufgabenfilter</h1>
      <div>
        {#if !quest.questQuestFilter}
          Noch keine Aufgabenfilter
        {/if}
        {#each quest.questQuestFilter as questFilter, i}
          <div>
            {questFilter.questId} - - {FilterType[questFilter.filterType]}
            <!--die warnung ist einfach mal falsch, cool https://www.typescriptlang.org/docs/handbook/enums.html -->
          </div>
          <button
            on:click={() => {
              quest.questQuestFilter = quest.questQuestFilter.filter(
                (_, j) => j !== i,
              );
            }}
          >
            entfernen
          </button>
        {/each}
      </div>
      <hr />
      <div>
        <select
          id="region"
          name="region"
          type="select"
          bind:value={newQuestFilter.filterType}
        >
          <option value={FilterType.REQURIED}> REQURIED </option>
          <option value={FilterType.EXCLUDED}> EXCLUDED </option>
          <option value={FilterType.PREFERED}> PREFERED </option>
        </select>
      </div>
      <div>
        <button
          on:click={() => {
            if (newQuestFilter.questId !== undefined) {
              quest.questQuestFilter = [
                ...quest.questQuestFilter,
                newQuestFilter,
              ];
              newQuestFilter = new QuestQuestFilter();
            }
          }}
        >
          Aufgabenfilter speichern
        </button>
      </div>
    </div>

    <RewardForm reward={quest.reward} />

    <label for={"userSelectable"} style="padding: 5px;">
      <input
        class="checkbox"
        type="checkbox"
        name={"userSelectable"}
        bind:checked={quest.userSelectable}
      />
      Kann ausgewählt werden
    </label>

    <label for={"published"} style="padding: 5px;">
      <input
        class="checkbox"
        type="checkbox"
        name={"published"}
        bind:checked={quest.published}
      />
      Veröffentlichen
    </label>
  </div>
  <div>
    <div >
      <button
        
        on:click={(e) => {
          save();
        }}>Speichern</button
      >
    </div>

    <div >
      <button
        on:click={(e) => {
          deleteQuest();
        }}>Löschen</button
      >
    </div>
  </div>
</div>
