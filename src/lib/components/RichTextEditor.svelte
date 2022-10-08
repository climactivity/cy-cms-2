<script lang="ts">
	import { onMount } from "svelte";

  let editorElement; 


  export let toolbarOptions = [
		[{ header: 1 }, { header: 2 }, "blockquote", "link", "image", "video"],
		["bold", "italic", "underline", "strike"],
		[{ list: "ordered" }, { list: "ordered" }],
		[{ align: [] }],
		["clean"]
	];


  let quill;
  onMount(async () => {
    const { default: Quill } = await import("quill");

    quill = new Quill("#editorElement", {
      modules: {
        toolbar: toolbarOptions
      },
      theme: "snow",
      placeholder: "Write your story..."
    });


  })

  const save = () => {
    console.log("Content")
    console.log(editorElement)
    const editor = editorElement.querySelector('.ql-editor')
    console.log(editor)
    console.log(editor.innerHTML)
    console.log(editorElement.querySelector('.ql-editor'))

    console.log(quill.getContents())
    console.log(quill.getText())
    console.log(quill.getText())


  }
</script>

<div class="editor-wrapper">
  <button on:click={save}>SAVE</button>
<div id="editorElement" bind:this={editorElement}/>
</div>

<style>
  @import 'https://cdn.quilljs.com/1.3.6/quill.snow.css';
</style>