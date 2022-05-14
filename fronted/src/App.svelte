<script>
  import { onMount } from "svelte";
  const endpoint = "http://localhost:1234";
  let song = [];

  onMount(async () => {
    const response = await fetch(endpoint);
    song = await response.json();
    console.log(song);
  });
</script>

<main>
  <h1>ymp3web</h1>
  <h3>🟢: {song.ip}</h3>
  <h2>Listening {song.title} By {song.by}</h2>
  <img src={song.img} alt="" />
  <h1>Available songs:</h1>
  <!--
    Lol
  -->
 <p>{song.songs}</p>
</main>

<style>
  img {
    width: 200px;
    height: 200px;
    margin: 5px;
    border-radius: 50%;
    border: 1px solid black;
    display: inline-block;
    box-shadow: 0px 0px 5px black;
  }
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ffb5da;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 900;
  }
  h3 {
    color: #000;
    font-weight: 900;
  }
  h2 {
    color: #ffb5da;
    font-weight: 900;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
