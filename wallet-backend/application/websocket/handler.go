package websocket

import (
	"fmt"
	"net/http"

	"encoding/json"

	"context"

	"bytes"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/websocket"
	"github.com/kowala-tech/wallet-backend/application/command"
)

//Handler is an http.Handler used to bind a websocket connection to all use cases that are sent
//through websocket.
type Handler struct {
	Logger      log.Logger
	GetBlockCmd command.GetBlockHeightHandler
}

//Request encapsulates a request sent through a websocket to the handler.
type Request struct {
	Action string `json:"action"`
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		h.Logger.Log(
			"type",
			"alert",
			"msg",
			fmt.Sprintf("Error with creating socket: %s", err),
		)
		return
	}
	defer conn.Close()

	h.readAndHandle(conn)
}

func (h *Handler) readAndHandle(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			h.Logger.Log(
				"type",
				"alert",
				"msg",
				fmt.Sprintf("Error reading from socket: %s", err),
			)
			break
		}

		wsReq, err := parseMsg(msg)
		if err != nil {
			h.Logger.Log(
				"type",
				"alert",
				"msg",
				fmt.Sprintf("Invalid request: %s", err),
			)
			break
		}

		if wsReq.Action == "" {
			h.Logger.Log(
				"type",
				"alert",
				"msg",
				"Invalid request action",
			)
			break
		}

		resp, err := h.executeAction(wsReq)
		if err != nil {
			h.Logger.Log(
				"type",
				"alert",
				"msg",
				fmt.Sprintf("Error getting block: %s", err),
			)
			break
		}

		conn.WriteMessage(websocket.TextMessage, resp)
	}
}

func (h *Handler) executeAction(request *Request) ([]byte, error) {
	ctx := context.Background()
	var response []byte
	b := bytes.NewBuffer(response)

	if request.Action == "blockheight" {
		blockHeight, err := h.GetBlockCmd.Handle(ctx)
		if err != nil {
			return nil, err
		}

		json.NewEncoder(b).Encode(blockHeight)
	}

	return b.Bytes(), nil
}

func parseMsg(msg []byte) (*Request, error) {
	var r Request

	err := json.Unmarshal(msg, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
