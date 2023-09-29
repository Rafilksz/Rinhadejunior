<script>
  let question = '';
  let answer = '';

  const handleSubmit = async () => {
    const response = await fetch('http://localhost:8081/api/ask', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ question }),
    });

    if (response.ok) {
      const data = await response.json();
      answer = data.answer;
      // Limpar o campo de pergunta
      question = '';
    } else {
      answer = 'Erro ao enviar a pergunta.';
    }
  };
</script>

<div class="chat-window">
  <div class="chat-messages">
    {#if answer}
      <div class="answer-box">
        <div class="answer">{answer}</div>
      </div>
    {/if}
  </div>
  <div class="chat-input">
    <input bind:value={question} placeholder="Pergunte qualquer coisa 🙊🙊" />
    <button on:click={handleSubmit}>Enviar</button>
  </div>
</div>

<style>
  .chat-window {
    border: 3px solid #ccc;
    border-radius: 20px;
    padding: 10px;
    max-width: 400px;
    margin: 0 auto;
    margin-bottom: 100px;
  }

  .chat-messages {
    margin-bottom: 2px;
  }

  .answer-box {
    margin-bottom: 10px;
  }

  .answer {
    background-color: #f5f5f5;
    padding: 20px;
    border-radius: 3px;
  }

  .chat-input {
    display: flex;
    gap: 10px;
  }

  input {
    flex-grow: 1;
    padding: 10px;
    border: 3px solid #ccc;
    border-radius: 5px;
  }

  button {
    background-color: #007bff;
    color: #fff;
    border: none;
    border-radius: 5px;
    padding: 5px 10px;
    cursor: pointer;
  }
</style>
