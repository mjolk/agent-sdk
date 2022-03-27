/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package rest represents REST API.
package rest

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"

	"github.com/google/uuid"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"github.com/trustbloc/agent-sdk/cmd/agent-mobile/pkg/api"
	"github.com/trustbloc/agent-sdk/cmd/agent-mobile/pkg/wrappers/config"
)

// Aries is an Aries implementation with endpoints to execute operations.
type Aries struct {
	endpoints map[string]map[string]endpoint

	URL          string
	WebsocketURL string
	Token        string
	mutex        sync.RWMutex
	subscribers  map[string]map[string][]api.Handler
}

// NewAries returns a new Aries instance.
// Use this if you want your requests to be handled by a remote agent.
func NewAries(opts *config.Options) (*Aries, error) {
	if opts == nil || opts.AgentURL == "" {
		return nil, errors.New("no agent url provided")
	}

	endpoints := getControllerEndpoints()

	a := &Aries{
		endpoints:    endpoints,
		URL:          opts.AgentURL,
		Token:        opts.APIToken,
		WebsocketURL: opts.WebsocketURL,
		subscribers:  make(map[string]map[string][]api.Handler),
	}

	go a.startNotificationListener()

	return a, nil
}

// Incoming represents WebSocket event message.
type Incoming struct {
	ID      string      `json:"id"`
	Topic   string      `json:"topic"`
	Message interface{} `json:"message"`
}

func (ar *Aries) startNotificationListener() {
	if ar.WebsocketURL == "" {
		return
	}

	conn, _, err := websocket.Dial(context.Background(), ar.WebsocketURL, nil) // nolint: bodyclose
	if err != nil {
		logger.Errorf("notification listener: websocket dial: %v", err)

		return
	}

	defer func() {
		err = conn.Close(websocket.StatusNormalClosure, "closing the connection")
		if err != nil && websocket.CloseStatus(err) != websocket.StatusNormalClosure {
			logger.Errorf("notification listener: close connection: %v", err)
		}
	}()

	for {
		var incoming *Incoming
		// listens for notifications
		if err := wsjson.Read(context.Background(), conn, &incoming); err != nil {
			if websocket.CloseStatus(err) != websocket.StatusNormalClosure {
				logger.Errorf("notification listener: read: close connection: %v", err)
			}

			// exit if websocket is closed
			if websocket.CloseStatus(err) != -1 {
				return
			}

			continue
		}

		ar.notifySubscribers(incoming)
	}
}

func (ar *Aries) notifySubscribers(incoming *Incoming) {
	ar.mutex.RLock()
	defer ar.mutex.RUnlock()

	// gets all handlers that were subscribed for the topic
	for _, handlers := range ar.subscribers[incoming.Topic] {
		raw, err := json.Marshal(incoming)
		if err != nil {
			logger.Errorf("notification listener: marshal: %v", err)

			break
		}

		// send the payload to the subscribers
		for _, handler := range handlers {
			if err := handler.Handle(incoming.Topic, raw); err != nil {
				logger.Errorf("notification listener: handle: %v", err)
			}
		}
	}
}

// RegisterHandler registers a handler to process incoming notifications from the framework.
// Handler is implemented by mobile apps.
func (ar *Aries) RegisterHandler(h api.Handler, topics string) string {
	ar.mutex.Lock()
	defer ar.mutex.Unlock()

	id := uuid.New().String()

	for _, topic := range strings.Split(topics, ",") {
		if ar.subscribers[topic] == nil {
			ar.subscribers[topic] = map[string][]api.Handler{}
		}

		ar.subscribers[topic][id] = append(ar.subscribers[topic][id], h)
	}

	return id
}

// UnregisterHandler unregisters a handler by given id.
func (ar *Aries) UnregisterHandler(id string) {
	ar.mutex.Lock()
	defer ar.mutex.Unlock()

	for topic := range ar.subscribers {
		for key := range ar.subscribers[topic] {
			if key == id {
				delete(ar.subscribers[topic], id)
			}
		}
	}
}
