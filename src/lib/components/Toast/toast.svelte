<script lang="ts">
  import { ToastStyle } from "./toast-types";
  import { fly } from "svelte/transition";
  export let triggerToast = (
    style: ToastStyle,
    message: string,
    duration: number = 3000
  ) => {
    hidden = false;
    _message = message;
    _style = style;
    setTimeout(() => (hidden = true), duration);
    console.log(getColor());
  };

  let hidden = true;
  let _message = "";
  let _style = ToastStyle.info;
  const getColor = () => {
    switch (_style) {
      case ToastStyle.info:
        return ToastStyle.info.toString();
      case ToastStyle.failed:
        return ToastStyle.failed.toString();
      case ToastStyle.success:
        return ToastStyle.success.toString();
    }
  };
</script>

<div class="absolute right-1/2 top-[5.5rem]">
  {#if !hidden}
    <div class="toast" transition:fly={{ y: -50, duration: 1000 }}>
      <div
        class="shadow-xl rounded-md p-2 text-zinc-800 bg-white border-2"
        style="border-color: {getColor()};"
      >
        <span>{_message}</span>
      </div>
    </div>
  {/if}
</div>
<slot {triggerToast} {...$$props} />

<style lang="scss">
  .toast {
    @apply absolute top-1/2 shadow-2xl z-[10000];
  }
</style>
