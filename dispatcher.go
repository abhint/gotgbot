package gotgbot

import (
	"gotgbot/ext"
)

type Dispatcher struct {
	Bot      ext.Bot
	updates  chan Update
	handlers *[]Handler
}

func NewDispatcher(bot ext.Bot, updates chan Update) Dispatcher {
	d := Dispatcher{}
	d.Bot = bot
	d.updates = updates
	d.handlers = new([]Handler)
	return d
}

func (d Dispatcher) Start() {
	for upd := range d.updates {
		d.processUpdate(upd)
	}
}

func (d Dispatcher) processUpdate(update Update) {
	for _, handler := range *d.handlers {
		if handler.CheckUpdate(update) {
			handler.HandleUpdate(update, d)
			break
		}
	}
}

func (d Dispatcher) AddHandler(handler Handler) {
	*d.handlers = append(*d.handlers, handler)

}
