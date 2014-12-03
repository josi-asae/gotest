package flow

import (
	"code.google.com/p/go.net/websocket"
	"github.com/nu7hatch/gouuid"
	"testing"
)

var (
	r       *Runtime
	started bool
)

func ensureRuntimeStarted() {
	if !started {
		r = new(Runtime)
		r.Init()
		go r.Listen("localhost:13014")
		started = true
		<-r.Ready()
	}
}

// Tests runtime information support
func TestRuntimeGetRuntime(t *testing.T) {
	ensureRuntimeStarted()
	// Create a WebSocket client
	ws, err := websocket.Dial("ws://localhost:13014/", "", "http://localhost/")
	if err != nil {
		t.Error(err.Error())
	}
	// Send a runtime request and check the response
	if err = websocket.JSON.Send(ws, &Message{"runtime", "getruntime", nil}); err != nil {
		t.Error(err.Error())
	}
	var res runtimeInfo
	if err = websocket.JSON.Receive(ws, &res); err != nil {
		t.Error(err.Error())
	}
	if res.Type != "goflow" {
		t.Errorf("Invalid protocol type: %s\n", res.Type)
	}
	if res.Version != "0.4" {
		t.Errorf("Invalid protocol version: %s\n", res.Version)
	}
	if len(res.Capabilities) != 5 {
		t.Errorf("Invalid number of supported capabilities: %v\n", res.Capabilities)
	}
	if res.Id == "" {
		t.Error("Runtime Id is empty")
	}
}

//Tests Init
func TestInit(t *testing.T) {
	uv4, err := uuid.NewV4()
	if err != nil {
		t.Error(err.Error())
	}
	if uv4 == nil {
		t.Error(err.Error())
	}
}

//Test Id
func TestId(t *testing.T) {

}

//Test Ready
func TestReady(t *testing.T) {

}

//Test Stop
func TestStop(t *testing.T) {

}

//Test Handle
func TestHandle(t *testing.T) {

}

//Test Listen
func TestListen(t *testing.T) {

}

