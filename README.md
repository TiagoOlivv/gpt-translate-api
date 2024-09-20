# GPT Translate API

This project documentation is also available in **[Portuguese](README.pt-BR.md)**.

GPT Translate API is a simple application that utilizes the OpenAI API to translate texts between different languages. The API receives a text, the source language, and the target language, and returns the translation generated by the GPT-3.5 Turbo model.

## Project Structure

```bash
gpt-translate-api/
│
├── .env                # Environment variables (keys and sensitive configurations)
├── Makefile            # Shortcuts for running commands (build, run, etc.)
├── go.mod              # Go module and dependencies
├── go.sum              # Dependency checksums
├── index.go            # Main initialization file
└── src/
    ├── config/         # Application configurations
    ├── handlers/       # Translation processing logic
    ├── models/         # Data structures and models
    ├── routes/         # API route definitions and setup using Gin
    └── services/       # Functions to interact with OpenAI API
```

## Prerequisites

- **Go** (version 1.20+)
- **Postman** or another tool to test APIs
- An OpenAI account and a valid API key

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/tiagoolivv/gpt-translate-api.git
   cd gpt-translate-api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Create the `.env` file with the following variables:

   ```bash
   OPENAI_MODEL=gpt-3.5-turbo
   OPENAI_AUTHORIZATION=your_openai_api_key
   OPENAI_ORGANIZATION=your_organization_id
   OPENAI_PROJECT=your_project_id
   OPENAI_URL=https://api.openai.com/v1/chat/completions
   PORT=8080
   ```

4. Run the project:

   Using `go run`:
   
   ```bash
   go run ./index.go
   ```

   Or with the `Makefile`:

   ```bash
   make run
   ```

## Usage

### Endpoint: `/translate`

- **Method:** `POST`
- **URL:** `http://localhost:8080/translate`

### Request Body

```json
{
  "q": "Hello, how are you?",
  "source": "en",
  "target": "es"
}
```

| Field    | Type   | Description                         |
|----------|--------|-------------------------------------|
| `q`      | string | The text to be translated           |
| `source` | string | The source language code (ISO 639-1) |
| `target` | string | The target language code (ISO 639-1) |

### Request Example:

```bash
curl -X POST http://localhost:8080/translate \
  -H "Content-Type: application/json" \
  -d '{
    "q": "Hello, how are you?",
    "source": "en",
    "target": "es"
  }'
```

### Response Example:

```json
{
  "text": "Hello, how are you?",
  "source": "en",
  "target": "es",
  "result": "Hola, ¿cómo estás?"
}
```

## Testing

You can test the API using [Postman](https://www.postman.com/) or the `curl` command as shown above.

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gin](https://github.com/gin-gonic/gin) - Web framework for Go
- [Resty](https://github.com/go-resty/resty) - HTTP client for Go
- [OpenAI API](https://beta.openai.com/docs/) - API for translation using GPT-3.5 Turbo

## How to Contribute

1. Fork the project
2. Create a new branch: `git checkout -b feature/your-feature`
3. Make your changes and commit: `git commit -m 'Added my feature'`
4. Push to the branch: `git push origin feature/your-feature`
5. Open a Pull Request