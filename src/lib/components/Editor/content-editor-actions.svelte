<script lang="ts">
  import { goto } from "$app/navigation";
  import { baseUrl } from "$lib/stores/stores";
  import { ToastStyle } from "$lib/components/Toast/toast-types";
  import Toast from "$lib/components/Toast/toast.svelte";

  import type { Content } from "./editor-types";
  export let saveTarget;
  const save = async () => {
    saving = true;
    const res = await fetch(
      `${baseUrl}${saveTarget}${data._id ? "/" + data._id : ""}`,
      {
        credentials: "include",
        method: data._id ? "PUT" : "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      }
    );
    setTimeout(() => (saving = false), 100); // make the button animate even if saving only takes > 100ms
    return res.ok;
  };
  export let data: Content;
  export let title = "";
  const goBack = (event) => {
    history.back();
  };
  let saving = false;
  let unsavedChanges = false;
</script>

<div {...$$props}>
  <div class="pb-8 mx-auto w-full flex flex-row justify-center">
    <div class="bg-white rounded-md shadow-none bg-white w-full ">
      <nav class="grid grid-cols-3  ">
        <div class="flex flex-row col-span-1 items-stretch">
          <button class="text-lg my-auto  p-2" on:click|preventDefault={goBack}
            >🡄</button
          >
          <span class="text-lg my-auto">{title}</span>
          {#if unsavedChanges}
            * (ungespeicherte änderungen)
          {/if}
        </div>
        <dev
          class="flex flex-row col-span-1 col-start-3 place-content-end mr-4"
        >
          <label class="text-lg my-auto"
            >Veröffentlichen <input
              type="checkbox"
              bind:checked={data.published}
            /></label
          >

          <Toast let:triggerToast>
            <button
              class=" button cta ml-2 my-auto "
              class:saving
              on:click={async () => {
                let ok = await save();
                let style = ok ? ToastStyle.success : ToastStyle.failed;
                let message = ok
                  ? "Gespeichert"
                  : "Es ist etwas schief gelaufen 😣";
                triggerToast(style, message, 3000);
              }}>Speichern</button
            >
          </Toast>
        </dev>
      </nav>
    </div>
  </div>
</div>

<style lang="scss">
  .saving {
    @apply bg-slate-300 transition-colors;
  }
</style>
