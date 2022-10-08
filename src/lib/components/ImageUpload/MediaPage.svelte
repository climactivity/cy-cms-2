<script lang="ts">
  import ImageUpload from "$lib/components/Media/ImageUpload.svelte";
  import ImageGallery from "$lib/components/ImageUpload/ImageGallery.svelte";

  let modalOpen = false;
  const toggleModal = () => {
    modalOpen = !modalOpen;
  };

  const imageUploadedCallback = () => {
    console.log("imageUploadedCallback");
    toggleModal();
    flashToast(2500);
    if (gallery) {
      gallery.refetch();
    }
  };

  const imageModifiedCallback = () => {
    if (gallery) {
      gallery.refetch();
    }
  };

  let showToast = false;

  const flashToast = (duration) => {
    showToast = true;
    setTimeout(() => {
      showToast = false;
    }, duration);
  };

  let gallery;
</script>

<div class={showToast ? "p-3 mb-3 absolute" : "hidden absolute"}>
  <div class="mr-1" success isOpen={showToast}>
    <div>📸 Bild hochgeladen!</div>
    <div>
      Dein Bild ist jetzt hochgeladen! Kopier den Permalink mit dem Button in
      der Gallery!
    </div>
  </div>
</div>

<button success on:click={toggleModal} style="width: 100%">
  Bild hochladen
</button>
<ImageGallery bind:this={gallery} {imageModifiedCallback} />

<div isOpen={modalOpen} {toggleModal} transitionOptions>
  <h1 toggle={toggleModal}>Bild hochladen</h1>
  <h3>
    <ImageUpload {imageUploadedCallback} />
  </h3>
</div>

<style>
  .hidden {
    display: none;
  }
  .absolute {
    position: absolute;
    right: 1rem;
    top: 1rem;
  }
</style>
