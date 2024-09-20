# GPT Translate API

A documentação deste projeto também está disponível em **[Inglês](README.md)**.

GPT Translate API é uma aplicação simples que utiliza a API do OpenAI para traduzir textos entre diferentes idiomas. A API recebe um texto, o idioma de origem e o idioma de destino, e retorna a tradução gerada pelo modelo GPT-4o Mini.

## Estrutura do Projeto

```bash
gpt-translate-api/
│
├── .env                # Variáveis de ambiente (chaves e configurações sensíveis)
├── Makefile            # Atalhos para executar comandos (build, run, etc.)
├── go.mod              # Módulo Go e dependências
├── go.sum              # Checksums de dependências
├── index.go            # Arquivo principal de inicialização
└── src/
    ├── config/         # Configurações do aplicativo
    ├── handlers/       # Lógica de processamento de tradução
    ├── models/         # Estruturas de dados e modelos
    ├── routes/         # Definições e configuração de rotas de API usando Gin
    └── services/       # Funções para interagir com a API OpenAI
```

## Pré-requisitos

- **Go** (versão 1.20+)
- **Postman** ou outra ferramenta para testar APIs
- Uma conta OpenAI e uma chave de API válida

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/gpt-translate-api.git
   cd gpt-translate-api
   ```

2. Instale as dependências:

   ```bash
   go mod tidy
   ```

3. Crie o arquivo `.env` com as seguintes variáveis:

   ```bash
   OPENAI_MODEL=gpt-3.5-turbo
   OPENAI_AUTHORIZATION=your_openai_api_key
   OPENAI_ORGANIZATION=your_organization_id
   OPENAI_PROJECT=your_project_id
   OPENAI_URL=https://api.openai.com/v1/chat/completions
   PORT=8080
   ```

4. Execute o projeto:

   Com `go run`:
   
   ```bash
   go run ./index.go
   ```

   Ou usando o `Makefile`:

   ```bash
   make run
   ```

## Uso

### Endpoint: `/translate`

- **Método:** `POST`
- **URL:** `http://localhost:8080/translate`

### Corpo da Requisição

```json
{
  "q": "Hello, how are you?",
  "source": "en",
  "target": "es"
}
```

| Campo     | Tipo   | Descrição                          |
|-----------|--------|------------------------------------|
| `q`       | string | Texto que será traduzido           |
| `source`  | string | Código do idioma de origem (ISO 639-1) |
| `target`  | string | Código do idioma de destino (ISO 639-1) |

### Exemplo de Requisição:

```bash
curl -X POST http://localhost:8080/translate \
  -H "Content-Type: application/json" \
  -d '{
    "q": "Hello, how are you?",
    "source": "en",
    "target": "es"
  }'
```

### Exemplo de Resposta:

```json
{
  "text": "Hello, how are you?",
  "source": "en",
  "target": "es",
  "result": "Hola, ¿cómo estás?"
}
```

## Testes

Para testar a API, você pode usar o [Postman](https://www.postman.com/) ou o comando `curl` como mostrado acima. 

## Tecnologias Utilizadas

- [Go](https://golang.org/) - Linguagem de programação
- [Gin](https://github.com/gin-gonic/gin) - Framework web para Go
- [Resty](https://github.com/go-resty/resty) - Biblioteca HTTP para Go
- [OpenAI API](https://beta.openai.com/docs/) - API para tradução usando GPT-3.5 Turbo

## Como Contribuir

1. Faça um fork do projeto
2. Crie uma nova branch: `git checkout -b feature/sua-feature`
3. Faça suas alterações e commit: `git commit -m 'Adicionei minha feature'`
4. Faça push para a branch: `git push origin feature/sua-feature`
5. Abra um Pull Request
