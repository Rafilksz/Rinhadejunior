package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/ask", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		var requestBody map[string]string
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
			return
		}

		respostas := []string{
			"Não sei, foi um junior que me programou",
			"Desculpe, meu código está agindo estranho hoje. Acho que o desenvolvedor júnior fez algumas alterações.",
			"Parece que o desenvolvedor júnior mexeu na lógica do sistema novamente. Vou tentar corrigir isso.",
			"Talvez um dia eu descubra a resposta.",
			"Essa pergunta me deixou perplexo!",
			"Oh, não, mais um bug? Deixe-me adivinhar, foi o desenvolvedor júnior, certo?",
			"O junior que me programou era burro demais para saber esta resposta",
			"Parece que o desenvolvedor júnior estava \"pensando fora da caixa\" novamente. Vou verificar as consequências disso.",
			"Desculpe por qualquer inconveniente. O desenvolvedor júnior estava tentando inovar",
			"Meu desenvolvedor não me programou direito, vou pedir para o @Criascript corrigir isso",
			"Meu desenvolvedor faltou as aulas online no youtube que ele comprou do AD do youtube, me desculpe",
			"Pergunta no Posto ipiranga !",
			"Esta esta no nivel GPT, pergunta uma que seja facil por favor",
			"não importa a pergunta, o resultado sempre sera 42 !",
			"Sua pergunta eu nao sei, mas com certeza o google sabe !",
			"ja tentou ver no google antes de me perguntar algo ?",
			"verifica se o @Criascript ta aovivo na twitch, talvez ele saiba",
			"Vou ficar te devendo campeão.",
			"Antes da resposta me faça um pix de 5 reais, chave=rafael1989longas@gmail.com.",
			"esta ai é facil demais, manda outra mais dificil",
		}

		// Escolha uma resposta aleatória.
		rand.Seed(time.Now().UnixNano())
		response := map[string]string{"answer": respostas[rand.Intn(len(respostas))]}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:5363"}) // Substitua pelo seu domínio Svelte.

	fmt.Println("Servidor rodando na porta 8081...")
	http.ListenAndServe(":8081", handlers.CORS(headers, methods, origins)(r))
}
