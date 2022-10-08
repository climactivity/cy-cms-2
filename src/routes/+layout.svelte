<script>
	import "./styles.css";
	import Header from "$lib/components/header/Header.svelte";
	import Searchbar from "$lib/components/header/Searchbar.svelte";
	import Footer from "$lib/components/Footer.svelte";
  </script>
  
  <!--navigation here-->
  <!-- <main class="layout bg-zinc-200"> -->
  
	<!-- <div class="header important shadow-md align-middle flex flex-row place-content-between">
		<h1 class="text-2xl my-auto ml-2">Cy CMS v2.0</h1>
	</div>
  
	<nav class="nav bg-zinc-200 pt-2">
	  <ul class="grid grid-flow-row fixed ml-2 space-y-2">
		<a href="/quests" class="navlink">Aufgaben</a>
		<a href="/infobytes" class="navlink">Infobytes</a>
		<a href="/media" class="navlink">Medien</a>
		<a href="/aspects" class="navlink">Aspekte</a>
		<a href="/treetemplates" class="navlink">B&#xE4;ume</a>
		<a href="/analytics" class="navlink">Analytics</a>
		<a href="#" class="navlink">Abmelden</a>
		<hr class="border-2 border-white rounded-full">
		<a href="https://nk.climactivity.de" class="navlink">Server Status</a>
	  </ul>
	</nav> -->
  
  <main class=" flex h-screen font-sans">
	<nav class="fixed ">
	  <Header />
	</nav>
	<div class=" flex-grow pl-20 content-start md:pl-32 pr-3 mt-0">
	  <content class="">
		<Searchbar />
		<slot />
	  </content>
  
	  <footer class="absolute bottom-0 pl-40  flex-grow-1 ">
		<!-- TODO footer content-->
		<!-- <Footer/> -->
	  </footer>
	</div>
  </main>
  
  <style>
	.footer {
	  grid-area: footer;
	}
  </style>
  