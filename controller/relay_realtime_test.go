package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

func TestRealtimeUpgraderNegotiatesOpenAIBetaSubprotocol(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("upgrade failed: %v", err)
			return
		}
		defer conn.Close()
	}))
	defer server.Close()

	url := "ws" + server.URL[len("http"):]
	dialer := websocket.Dialer{
		Subprotocols: []string{"openai-beta.realtime-v1"},
	}
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial realtime websocket: %v", err)
	}
	defer conn.Close()

	if got := conn.Subprotocol(); got != "openai-beta.realtime-v1" {
		t.Fatalf("negotiated subprotocol = %q, want %q", got, "openai-beta.realtime-v1")
	}
}
