package main

import (
	"cmp"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const modelName = "gemini-1.5-flash"
const defaultPort = "9000"

// Server state holding the context of the gemini client and the generative model.
type geminiServer struct {
	ctx   context.Context
	model *genai.GenerativeModel
}

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable to create a client.
	apiKey := os.Getenv("GOOGLE_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Could not create Gemini client %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)

	server := &geminiServer{
		ctx:   ctx,
		model: model,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /chat", server.chatHandler)
	mux.HandleFunc("POST /stream", server.streamingChatHandler)

	// Access preferred port the server must listen to as an environment variable if provided.
	port := cmp.Or(os.Getenv("PORT"), defaultPort)
	addr := "localhost:" + port
	log.Println("Listening on ", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

type part struct {
	Text string
}
type content struct {
	Role  string
	Parts []part
}
type chatRequest struct {
	Chat    string
	History []content
}

// chatHandler returns the complete response of the model to the client. Expects a JSON payload in
// the request with the following format:
// Request:
//   - chat: string
//   - history: Array
//
// Returns a JSON payload containing the model response with the following format.
// Response:
//   - text: string
func (gs *geminiServer) chatHandler(w http.ResponseWriter, r *http.Request) {
	cr := &chatRequest{}
	if err := parseRequestJSON(r, cr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cs := gs.startChat(cr.History)
	res, err := cs.SendMessage(gs.ctx, genai.Text(cr.Chat))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resTxt, err := responseString(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderResponseJSON(w, part{Text: resTxt})
}

// streamingChatHandler continuously streams the response of the model to the client. Expects a
// JSON payload in the request with the following format:
// Request:
//   - chat: string,
//   - history: Array,
//
// Partial responses from the model is text.
func (gs *geminiServer) streamingChatHandler(w http.ResponseWriter, r *http.Request) {
	cr := &chatRequest{}
	if err := parseRequestJSON(r, cr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cs := gs.startChat(cr.History)
	iter := cs.SendMessageStream(gs.ctx, genai.Text(cr.Chat))

	w.Header().Set("Content-Type", "text/event-stream")

	for {
		res, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			break
		}

		resTxt, err := responseString(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			break
		}

		fmt.Fprint(w, resTxt)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

// startChat starts a chat session with the model using the given history.
func (gs *geminiServer) startChat(hist []content) *genai.ChatSession {
	cs := gs.model.StartChat()
	cs.History = encodeHistory(hist)
	return cs
}

// encodeHistory converts []content to a []*genai.Content which is accepted by the model's chat session.
func encodeHistory(cs []content) []*genai.Content {
	gcs := make([]*genai.Content, len(cs))
	for i, c := range cs {
		gcs[i] = c.geminiCompatible()
	}

	return gcs
}

// geminiCompatible converts content to genai.Content accepted by the chat session.
func (c *content) geminiCompatible() *genai.Content {
	gc := &genai.Content{}
	gc.Role = c.Role
	ps := make([]genai.Part, len(c.Parts))
	for i, p := range c.Parts {
		ps[i] = genai.Text(p.Text)
	}
	gc.Parts = ps
	return gc
}

// responseString parses the model response of type genai.GenerateContentResponse to a string.
func responseString(resp *genai.GenerateContentResponse) (string, error) {
	// Only taking the first candidate since GenerationConfig.CandidateCount defaults to 1.
	if len(resp.Candidates) > 0 {
		if cs := contentString(resp.Candidates[0].Content); cs != nil {
			return *cs, nil
		}
	}

	return "", fmt.Errorf("invalid response from Gemini model")
}

// contentString converts genai.Content to a string.
func contentString(c *genai.Content) *string {
	if c == nil || c.Parts == nil {
		return nil
	}

	contStrs := make([]string, len(c.Parts))
	for i, part := range c.Parts {
		if pt, ok := part.(genai.Text); ok {
			contStrs[i] = string(pt)
		} else {
			return nil
		}
	}

	contStr := strings.Join(contStrs, "\n")
	return &contStr
}
