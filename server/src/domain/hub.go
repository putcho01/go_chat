package domain

import (
	"context"

	"github.com/putcho01/go_chat/src/services"
)

type Hub struct {
	Clients      map[*Client]bool //現在の参加者
	RegisterCh   chan *Client     //入室時
	UnRegisterCh chan *Client     //退出時
	BroadcastCh  chan []byte      //送信時
	pubsub       *services.PubSubService
}

const broadCastChan = "broadcast"

func NewHub(pubsub *services.PubSubService) *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroadcastCh:  make(chan []byte),
		pubsub:       pubsub,
	}
}

func (h *Hub) RunLoop() {
	for {
		select {
		case client := <-h.RegisterCh:
			h.register(client)

		case client := <-h.UnRegisterCh:
			h.unregister(client)

		case msg := <-h.BroadcastCh:
			h.publishMessage(msg)
		}
	}
}

func (h *Hub) SubscribeMessages() {
	ch := h.pubsub.Subscribe(context.TODO(), broadCastChan)

	for msg := range ch {
		h.broadCastToAllClient([]byte(msg.Payload))
	}
}

func (h *Hub) publishMessage(msg []byte) {
	h.pubsub.Publish(context.TODO(), broadCastChan, msg)
}

func (h *Hub) register(c *Client) {
	h.Clients[c] = true
}

func (h *Hub) unregister(c *Client) {
	delete(h.Clients, c)
}

func (h *Hub) broadCastToAllClient(msg []byte) {
	for c := range h.Clients {
		c.sendCh <- msg
	}
}
