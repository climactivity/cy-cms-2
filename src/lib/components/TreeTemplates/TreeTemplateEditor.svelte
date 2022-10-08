<script lang="ts">
  import { createForm } from "svelte-forms-lib";
  import { baseUrl } from "$lib/stores/stores";

  export let currentTemplate;
  $: {
    console.log(currentTemplate);
    if (!currentTemplate) {
      currentTemplate = {
        template_name: "tree0",
        texture_name: "tree0",
        ui_name: "Test Baum",
        preview_name: "2",
        coin_value: 0,
        unlock_level: 0,
        bigpoint_available: ["ernährung"],
        archetype: "tree",
        initial_state: {
          stage: 0,
          water: 0.0,
          water_needed: 40.0,
          bigpoint: "ernährung",
          aspect: "pflanzliche_ernährung",
        },
      };
    } else {
      $form = currentTemplate;
    }
  }

  const { form, errors, state, handleChange, handleSubmit, handleReset } =
    createForm({
      initialValues: {
        currentTemplate,
      },
      onSubmit: async (values) => {
        let treeTemplate = values;
        console.log($form, values, treeTemplate);
        if (treeTemplate._id === "") delete treeTemplate._id;
        let result = await fetch(
          `${baseUrl}tree-template${
            treeTemplate._id ? "/" + treeTemplate._id : ""
          }`,
          {
            method: treeTemplate._id ? "PUT" : "POST",
            headers: {
              "Content-Type": "application/json",
              // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            credentials: "include",

            body: JSON.stringify(treeTemplate),
          },
        );
        //currentInfobyte.set(new Infobyte());
      },
    });

  async function deleteTreeTemplate() {
    let id = $form._id;
    if (!id) {
      handleReset();
      return;
    }

    let result = await fetch(`${baseUrl}tree-template/${id}`, {
      credentials: "include",

      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
    });

    handleReset();
  }
</script>

<section>
  <h1>Baum Template Editor</h1>

  <form on:submit={handleSubmit}>
    <ul>
      <label for="template_name">Template Name</label>
      <input
        on:change={handleChange}
        on:blur={handleChange}
        bind:value={$form.template_name}
      />
    </ul>
    <ul>
      <label for="texture_name">texture_name</label>
      <input
        on:change={handleChange}
        on:blur={handleChange}
        bind:value={$form.texture_name}
      />
    </ul>
    <ul>
      <label for="ui_name">ui_name</label>
      <input
        on:change={handleChange}
        on:blur={handleChange}
        bind:value={$form.ui_name}
      />
    </ul>
    <ul>
      <label for="preview_name">preview_name</label>
      <input
        on:change={handleChange}
        on:blur={handleChange}
        bind:value={$form.preview_name}
      />
    </ul>
    <ul>
      <label for="coin_value">coin_value</label>
      <input
        on:change={handleChange}
        on:blur={handleChange}
        bind:value={$form.coin_value}
      />
    </ul>
    <ul>
      <label for="unlock_level">unlock_level</label>
      <input
        on:change={handleChange}
        on:blur={handleChange}
        bind:value={$form.unlock_level}
      />
    </ul>
    <ul>
      <label for="bigpoint_available">bigpoint_available</label>
      <input
        on:change={handleChange}
        on:blur={handleChange}
        bind:value={$form.bigpoint_available}
      />
    </ul>
    <ul>
      <label for="archetype">archetype</label>
      <input
        on:change={handleChange}
        on:blur={handleChange}
        bind:value={$form.archetype}
      />
  </ul>
  </form>
  <div>
    <div>
      <button  on:click={handleSubmit}>Speichern</button>
    </div>
    <div>
      <button  on:click={deleteTreeTemplate}>Löschen</button>
    </div>
  </div>
</section>

<style>
  .bordered {
    font-family: inherit;
    font-size: inherit;
    max-width: 400px;
    width: 100%;
    padding: 12px;
    box-sizing: border-box;
    border: 1px solid var(--grey);
    border-radius: 4px;
    transition: all 150ms ease;
    background: var(--white);
  }

  input:focus,
  select:focus {
    outline: none;
    box-shadow: 0 0 0 4px rgb(227, 227, 245);
    border-color: var(--grey);
  }
</style>
