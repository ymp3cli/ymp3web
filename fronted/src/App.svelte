<script>
  import { onMount } from "svelte";
  const endpoint = `http://${location.hostname}:1234`
  let song = [];
  let display = true;
  setInterval(() => (display = !display), 1000);

  onMount(async () => {
    try {
      const response = await fetch(endpoint);
      song = await response.json();
    } catch (error) {
      song.ip = "I couldn't find a ymp3cli online instance on your local network, check https://github.com/paij0se/ymp3cli for more info ";
      song.songs = error
    }
  });
</script>

<main>
  <h1>ymp3web</h1>
  <h3>🟢:{song.ip}</h3>
  <h2>Listening {song.title} By {song.by}</h2>
  <div class="disk">
    <img src={song.img} alt="" />
  </div>
  <form class="number_song">
    <input type="text" />
  </form>

  <h2>Available songs:</h2>
  <!--
    Lol
  -->
  <p>{song.songs}</p>
</main>

<style>
  .disk {
    animation-name: spin;
    animation-duration: 50000ms;
    animation-iteration-count: infinite;
    animation-timing-function: linear;
  }

  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }
  img {
    width: 200px;
    height: 200px;
    margin: 5px;
    border-radius: 50%;
    border: 2px solid black;
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
