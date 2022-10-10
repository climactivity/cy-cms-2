<script>
  let on = true;
  const check = () => {
    on = !on;
  };


</script>

<div
  class="{on
    ? 'sidebar'
    : 'side'} h-screen menu bg-white px-4 flex items-center nunito  shadow z-40"
>
  <ul class="list-reset ">
    <li class="my-2 md:my-0">
      <a
        href="/"
        class="block py-1 md:py-3 pl-1 align-middle text-gray-600 no-underline hover:text-indigo-400"
      >
        <i class="fa-solid fa-solar-panel mr-3 " /><span
          class="w-full inline-block pb-1 md:pb-0 text-sm">Dashboard</span
        >
      </a>
    </li>
    <li class="my-2 md:my-0">
      <a
        href="/"
        class="block py-1 md:py-3 pl-1 align-middle text-gray-600 no-underline hover:text-indigo-400"
      >
        <i class="fa-solid fa-person-circle-question mr-3 " /><span
          class="w-full inline-block pb-1 md:pb-0 text-sm">Aufgaben</span
        >
      </a>
    </li>
    <li class="my-2 md:my-0">
      <a
        href="/"
        class="block py-1 md:py-3 pl-1 align-middle text-gray-600 no-underline hover:text-indigo-400"
      >
        <i class="fa-solid fa-glasses mr-3" /><span
          class="w-full inline-block pb-1 md:pb-0 text-sm">Aspekte</span
        >
      </a>
    </li>
    <li class="my-2 md:my-0 ">
      <a
        href="/"
        class="block py-1 md:py-3 pl-1 align-middle text-gray-600 no-underline hover:text-indigo-400"
      >
        <i class="fa-solid fa-info mr-3 " /><span
          class="w-full inline-block pb-1 md:pb-0 text-sm">Infobytes</span
        >
      </a>
    </li>
    <li class="my-2 md:my-0">
      <a
        href="/media"
        class="block py-1 md:py-3 pl-1 align-middle text-gray-600 no-underline hover:text-indigo-400"
      >
        <i class="fa-solid fa-photo-film mr-3" /><span
          class="w-full inline-block pb-1 md:pb-0 text-sm">Median</span
        >
      </a>
    </li>
    <li class="my-2 md:my-0">
      <a
        href="/editor"
        class="block py-1 md:py-3 pl-1 align-middle text-gray-600 no-underline hover:text-indigo-400"
      >
        <i class="fa-solid fa-tree mr-3" /><span
          class="w-full inline-block pb-1 md:pb-0 text-sm">B&#xE4;ume</span
        >
      </a>
    </li>
    <li class="my-2 md:my-0">
      <a
        href="/"
        class="block py-1 md:py-3 pl-1 align-middle text-gray-600 no-underline hover:text-indigo-400"
      >
        <i class="fas fa-chart-area fa-fw mr-3 " /><span
          class="w-full inline-block pb-1 md:pb-0 text-sm">Analytics</span
        >
      </a>
    </li>
    <li class="my-2 md:my-0">
      <dr
        on:click={check}
        class="block py-1 md:py-3 pl-1 align-middle text-gray-600 no-underline hover:text-indigo-400"
      >
        <i class="fa-solid fa-arrows-left-right mr-3" /><span
          class="w-full inline-block pb-1 md:pb-0 text-sm"
          >{on ? "Hov on" : "Hov Off"}</span
        >
      </dr>
    </li>
  </ul>
</div>

<style>
  .nunito {
    font-family: "nunito", font-sans;
  }

  .sidebar {
    transition: ease-in-out all 0.3s;
    z-index: 9999;
    width: 64px;
    opacity:1;
    box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  }

  .sidebar span {
    position: absolute;
    transition: ease-in-out all 0.1s;
    opacity: 0;
  }


  .side {
    transition: ease-in-out all 0.3s;
    z-index: 9999;
    width: 130px;
    box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);

  }

  .side span {
    position: absolute;
    transition: ease-in-out all 0.1s;
    opacity: 1;
  }
  /* .side:hover {
    width: 130px;
    box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
    opacity: 1;
  }

  .side:hover span {
    opacity: 1;
  } */
</style>
