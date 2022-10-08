<script lang="ts">
  import RewardDisplayer from "../Rewards/RewardDisplayer.svelte";
  import { baseUrl } from "$lib/stores/stores";
  export let image: any;
  export let select;
  export let imageModifiedCallback: Function;
  let modalOpen = false;

  const toggleModal = () => {
    modalOpen = !modalOpen;
  };

  const updateImage = () => {};

  const deleteImage = () => {
    console.log(image._id);
    fetch(`${baseUrl}image-upload/delete/${image._id}`, {
      method: "DELETE",
    })
      .then(async (res) => {
        console.log(res);
        if (+status <= 400) {
          if (imageModifiedCallback !== null) {
            toggleModal();
            imageModifiedCallback();
            console.log(await res.json());
          }
        }
      })
      .then((data) => {
        console.log(data);
      });
  };
</script>

<div class="container">
  <div class="row">
    <div style="width: 100%; display: flex; flex-direction: row-reverse;">
      <button on:click={toggleModal}>X</button>
    </div>

    <div class="row-item">
      <a href={image.relativeUrl} target="_blank">
        <img class="thumb" src={image.relativeUrl} alt={image.title} />
      </a>
    </div>
    <div class="row-item">
      {image.title} <br />
      {`©${image.credit}, Lizenz: ${image.license}`}
    </div>
    <div class="row-item">
      <button
        on:click={(e) => {
          if (select) {
            console.log(select);
            select(image);
            return;
          }
          navigator.clipboard
            .writeText(
              `${window.location.protocol}//${window.location.host}${image.relativeUrl}`,
            )
            .then(
              function () {
                console.log(`${baseUrl.slice(0, -1)}${image.relativeUrl}`);
              },
              function (err) {
                console.error("Async: Could not copy text: ", err);
              },
            );
        }}>Link Kopieren</button
      >
    </div>
  </div>
</div>

<div isOpen={modalOpen} {toggleModal} transitionOptions>
  <h1 toggle={toggleModal}>Bild löschen</h1>
  <h3>Kann nicht rückgängig gemacht werden!</h3>
  <div>
    <button  on:click={deleteImage}>Löschen</button>
  </div>
</div>

<style>
  .thumb {
    width: 200px;
    height: 200px;
    object-fit: contain;
  }
  .row {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    margin-bottom: 10px;
    border: 1px;
    border-radius: 5px;
    border-color: lightgray;
    border-width: 1px;
    border-style: solid;
  }
  .row-item {
    width: auto;
  }
</style>
