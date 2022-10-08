<script lang="ts">
    import ImageGallery from "./ImageGallery.svelte";
    
    let modalOpen = false;
    export let img
    let input
    const showModal = () => {
        modalOpen = true;
    };
    const closeModal = () => {
        modalOpen = false;
    };
    const toggleModal = () => {
        modalOpen = !modalOpen;
    };

    const select = (i) => {
        console.log(i);
        closeModal();
        img = i
    };
</script>

<div class="thumb-container">

{#if img}
    <a href={img.relativeUrl} target="_blank">
        <img class="thumb" src={img.relativeUrl} alt={img.title} />
    </a>

{/if}

<button on:click={showModal}>Bild auswählen</button>
</div>

<div isOpen={modalOpen} {toggleModal} transitionOptions>
    <h1 toggle={toggleModal}>Bild auswählen</h1>
    <h3>
        <ImageGallery select={select} />
    </h3>
</div>


<style>
    .thumb {
        width: 200px;
        height: 200px;
        object-fit: contain;
    }
    .thumb-container {
        width: auto;
        height: 210px;

        display: flex;
        flex-direction: row;
        justify-content: flex-start;
        align-items: center;
        padding: 10px;
        margin-bottom: 10px;
        border: 1px;
        border-radius: 5px;
        border-color: lightgray;
        border-width: 1px;
        border-style: solid;
    }
</style>
