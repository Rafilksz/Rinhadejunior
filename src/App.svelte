<script>
	import ChatWindow from './ChatWindow.svelte';
  </script>
  
  <style>
	body {
	  background-color: rgb(85, 83, 83); 
	}
  
	main {
	  text-align: center;
	  padding: 1em;
	}
  
	button {
	  margin: 5px; 
	  padding: 10px 20px; 
	  border: none;
	  border-radius: 5px;
	  cursor: pointer;
	}
  </style>
  
  <main>
	<div style="display: flex; flex-direction: column; align-items: center;">
	  <img
		src="https://yt3.googleusercontent.com/HHjS1nkjdwqeYo7yfr4nnyVkWVLgjNdxinQKV9rvokMo1yfyP1cDXj05-3LPe7u9rMcO7eiV=s176-c-k-c0x00ffffff-no-rj"
		alt="Imagem"
		style="width: 176px; height: auto; border-radius: 50%;"
	  />
	  <h1 style="color: black;">Chat Gepeto dos cria.</h1>
	</div>
	<div style="margin-top: 50px;">
	  <ChatWindow />
	</div>
	<div style="margin-top: 20px;">
	  <p class="text-above-buttons">Redes sociais do Macaco calvo porem cheiroso que promoveu esta rinha de junior!</p>
	  <div style="display: flex; justify-content: center;">
		<a href="https://www.twitch.tv/criascript" style="text-decoration: none;" target="_blank">
		  <button style="background-color: #9147FF; color: #fff;">Twitch</button>
		</a>
		<a href="https://www.youtube.com/@criascript" style="text-decoration: none;" target="_blank">
		  <button style="background-color: #FF0000; color: #fff;">Youtube</button>
		</a>
		<a href="https://www.instagram.com/criascript/" style="text-decoration: none;" target="_blank">
		  <button style="background-color: #1877F2; color: #fff;">Instagram</button>
		</a>
		<a href="https://www.tiktok.com/@criascript" style="text-decoration: none;" target="_blank">
		  <button style="background-color: black; color: #fff;">TikTok</button>
		</a>
		<a href="https://github.com/criascript" style="text-decoration: none;" target="_blank">
		  <button style="background-color: black; color: #fff;">Github</button>
		</a>
		<a href="https://sincerocidio.files.wordpress.com/2010/07/macaco.png" style="text-decoration: none;" target="_blank">
		  <button style="background-color: #00AFF0; color: #fff;">Only Fans</button>
		</a>
		
	  </div>
	  <p class="text-above-buttons">Pagina desenvolvida e criada por Rafael Adolfo Longas, todos os direitos preservados por este junior que vos fala <a href="https://github.com/Rafilksz" style="color: black; text-decoration: underline;" target="_blank">https://github.com/Rafilksz</a> </p>
	</div>
  </main>
  