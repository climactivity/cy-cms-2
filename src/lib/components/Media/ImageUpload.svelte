<script lang="ts">
  import { baseUrl } from "$lib/stores/stores";
  export let images;
  export let imageUploadedCallback;
  let modalOpen = false;

  const defaultImageMetadata = {
    title: "",
    license: "",
    credit: ""
  }

  let imageMetadata = {...defaultImageMetadata};

  const toggleModal = () => {
    modalOpen = !modalOpen;
  };
  function uploadImage() {
    const formData = new FormData();
    formData.append("file", images[0]);
    formData.append("meta", JSON.stringify(imageMetadata));
    fetch(`${baseUrl}image-upload/upload`, {
      method: "POST",
      body: formData,
    }).then((res) => {
      console.log(res);
      if (+status <= 400) {
        //toggleModal()
        if (imageUploadedCallback !== null) {
          imageUploadedCallback();
        }
      }
    });
  }

  function reset() {
    images = null;
    imageMetadata =  {...defaultImageMetadata};
  }
</script>

<label for="fileUpload">Bild</label>
<input required id="fileUpload" type="file" bind:files={images} />

<label for="titleInput">Title</label>
<input
  required
  id="titleInput"
  name="titleInput"
  bind:value={imageMetadata.title}
/>

<label for="licenseInput">Lizenz</label>
<input
  required
  id="licenseInput"
  name="licenseInput"
  bind:value={imageMetadata.license}
/>

<label for="creditInput">Urheber</label>
<input
  required
  id="creditInput"
  name="creditInput"
  bind:value={imageMetadata.credit}
/>
<div>
  <button on:click={uploadImage}>Hochladen</button>
</div>
<!-- <div xs="6">
<Button on:click={reset}>Zurücksetzen</Button>
</div> -->
