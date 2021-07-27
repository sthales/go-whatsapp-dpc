package whatsapp

import (
	"fmt"
	"os"
	"strings"

	"github.com/Rhymen/go-whatsapp/binary"
	"github.com/Rhymen/go-whatsapp/binary/proto"
)

/*
The Handler interface is the minimal interface that needs to be implemented
to be accepted as a valid handler for our dispatching system.
The minimal handler is used to dispatch error messages. These errors occur on unexpected behavior by the websocket
connection or if we are unable to handle or interpret an incoming message. Error produced by user actions are not
dispatched through this handler. They are returned as an error on the specific function call.
*/
type Handler interface {
	HandleError(err error)
}

type SyncHandler interface {
	Handler
	ShouldCallSynchronously() bool
}

/*
The TextMessageHandler interface needs to be implemented to receive text messages dispatched by the dispatcher.
*/
type TextMessageHandler interface {
	Handler
	HandleTextMessage(message TextMessage)
}

/*
The ImageMessageHandler interface needs to be implemented to receive image messages dispatched by the dispatcher.
*/
type ImageMessageHandler interface {
	Handler
	HandleImageMessage(message ImageMessage)
}

/*
The VideoMessageHandler interface needs to be implemented to receive video messages dispatched by the dispatcher.
*/
type VideoMessageHandler interface {
	Handler
	HandleVideoMessage(message VideoMessage)
}

/*
The AudioMessageHandler interface needs to be implemented to receive audio messages dispatched by the dispatcher.
*/
type AudioMessageHandler interface {
	Handler
	HandleAudioMessage(message AudioMessage)
}

/*
The DocumentMessageHandler interface needs to be implemented to receive document messages dispatched by the dispatcher.
*/
type DocumentMessageHandler interface {
	Handler
	HandleDocumentMessage(message DocumentMessage)
}

/*
The LiveLocationMessageHandler interface needs to be implemented to receive live location messages dispatched by the dispatcher.
*/
type LiveLocationMessageHandler interface {
	Handler
	HandleLiveLocationMessage(message LiveLocationMessage)
}

/*
The LocationMessageHandler interface needs to be implemented to receive location messages dispatched by the dispatcher.
*/
type LocationMessageHandler interface {
	Handler
	HandleLocationMessage(message LocationMessage)
}

/*
The StickerMessageHandler interface needs to be implemented to receive sticker messages dispatched by the dispatcher.
*/
type StickerMessageHandler interface {
	Handler
	HandleStickerMessage(message StickerMessage)
}

/*
The ContactMessageHandler interface needs to be implemented to receive contact messages dispatched by the dispatcher.
*/
type ContactMessageHandler interface {
	Handler
	HandleContactMessage(message ContactMessage)
}

/*
The ContactsArrayMessageHandler interface needs to be implemented to receive contacts dispatched by dispatcher.
*/
type ContactsArrayMessageHandler interface {
	Handler
	HandleContactsArrayMessage(message ContactsArrayMessage)
}

/*
The JsonMessageHandler interface needs to be implemented to receive json messages dispatched by the dispatcher.
These json messages contain status updates of every kind sent by WhatsAppWeb servers. WhatsAppWeb uses these messages
to built a Store, which is used to save these "secondary" information. These messages may contain
presence (available, last see) information, or just the battery status of your phone.
*/
type JsonMessageHandler interface {
	Handler
	HandleJsonMessage(message string)
}

/**
The RawMessageHandler interface needs to be implemented to receive raw messages dispatched by the dispatcher.
Raw messages are the raw protobuf structs instead of the easy-to-use structs in TextMessageHandler, ImageMessageHandler, etc..
*/
type RawMessageHandler interface {
	Handler
	HandleRawMessage(message *proto.WebMessageInfo)
}

/**
The ContactListHandler interface needs to be implemented to applky custom actions to contact lists dispatched by the dispatcher.
*/
type ContactListHandler interface {
	Handler
	HandleContactList(contacts []Contact)
}

/**
The ChatListHandler interface needs to be implemented to apply custom actions to chat lists dispatched by the dispatcher.
*/
type ChatListHandler interface {
	Handler
	HandleChatList(contacts []Chat)
}

type ActionHandler interface {
	Handler
	HandleAction(action Action)
}

/*
AddHandler adds an handler to the list of handler that receive dispatched messages.
The provided handler must at least implement the Handler interface. Additionally implemented
handlers(TextMessageHandler, ImageMessageHandler) are optional. At runtime it is checked if they are implemented
and they are called if so and needed.
*/
func (wac *Conn) AddHandler(handler Handler) {
	wac.handler = append(wac.handler, handler)
}

// RemoveHandler removes a handler from the list of handlers that receive dispatched messages.
func (wac *Conn) RemoveHandler(handler Handler) bool {
	i := -1
	for k, v := range wac.handler {
		if v == handler {
			i = k
			break
		}
	}
	if i > -1 {
		wac.handler = append(wac.handler[:i], wac.handler[i+1:]...)
		return true
	}
	return false
}

// RemoveHandlers empties the list of handlers that receive dispatched messages.
func (wac *Conn) RemoveHandlers() {
	wac.handler = make([]Handler, 0)
}

func (wac *Conn) shouldCallSynchronously(handler Handler) bool {
	sh, ok := handler.(SyncHandler)
	return ok && sh.ShouldCallSynchronously()
}

func (wac *Conn) handle(message interface{}) {
	wac.handleWithCustomHandlers(message, wac.handler)
}

func (wac *Conn) handleWithCustomHandlers(message interface{}, handlers []Handler) {
	switch m := message.(type) {
	case error:
		for _, h := range handlers {
			if wac.shouldCallSynchronously(h) {
				h.HandleError(m)
			} else {
				go h.HandleError(m)
			}
		}
	case string:
		for _, h := range handlers {
			if x, ok := h.(JsonMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleJsonMessage(m)
				} else {
					go x.HandleJsonMessage(m)
				}
			}
		}
	case TextMessage:
		for _, h := range handlers {
			if x, ok := h.(TextMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleTextMessage(m)
				} else {
					go x.HandleTextMessage(m)
				}
			}
		}
	case ImageMessage:
		for _, h := range handlers {
			if x, ok := h.(ImageMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleImageMessage(m)
				} else {
					go x.HandleImageMessage(m)
				}
			}
		}
	case VideoMessage:
		for _, h := range handlers {
			if x, ok := h.(VideoMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleVideoMessage(m)
				} else {
					go x.HandleVideoMessage(m)
				}
			}
		}
	case AudioMessage:
		for _, h := range handlers {
			if x, ok := h.(AudioMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleAudioMessage(m)
				} else {
					go x.HandleAudioMessage(m)
				}
			}
		}
	case DocumentMessage:
		for _, h := range handlers {
			if x, ok := h.(DocumentMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleDocumentMessage(m)
				} else {
					go x.HandleDocumentMessage(m)
				}
			}
		}
	case LocationMessage:
		for _, h := range handlers {
			if x, ok := h.(LocationMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleLocationMessage(m)
				} else {
					go x.HandleLocationMessage(m)
				}
			}
		}
	case LiveLocationMessage:
		for _, h := range handlers {
			if x, ok := h.(LiveLocationMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleLiveLocationMessage(m)
				} else {
					go x.HandleLiveLocationMessage(m)
				}
			}
		}

	case StickerMessage:
		for _, h := range handlers {
			if x, ok := h.(StickerMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleStickerMessage(m)
				} else {
					go x.HandleStickerMessage(m)
				}
			}
		}

	case ContactMessage:
		for _, h := range handlers {
			if x, ok := h.(ContactMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleContactMessage(m)
				} else {
					go x.HandleContactMessage(m)
				}
			}
		}
	case ContactsArrayMessage:
		for _, h := range handlers {
			if x, ok := h.(ContactsArrayMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleContactsArrayMessage(m)
				} else {
					go x.HandleContactsArrayMessage(m)
				}
			}
		}

	case *proto.WebMessageInfo:
		for _, h := range handlers {
			if x, ok := h.(RawMessageHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleRawMessage(m)
				} else {
					go x.HandleRawMessage(m)
				}
			}
		}
	}

}

func (wac *Conn) handleContacts(contacts interface{}) {
	var contactList []Contact
	c, ok := contacts.([]interface{})
	if !ok {
		return
	}
	for _, contact := range c {
		contactNode, ok := contact.(binary.Node)
		if !ok {
			continue
		}

		jid := strings.Replace(contactNode.Attributes["jid"], "@c.us", "@s.whatsapp.net", 1)
		contactList = append(contactList, Contact{
			jid,
			contactNode.Attributes["notify"],
			contactNode.Attributes["name"],
			contactNode.Attributes["short"],
		})
	}
	for _, h := range wac.handler {
		if x, ok := h.(ContactListHandler); ok {
			if wac.shouldCallSynchronously(h) {
				x.HandleContactList(contactList)
			} else {
				go x.HandleContactList(contactList)
			}
		}
	}
}

func (wac *Conn) handleChats(chats interface{}) {
	var chatList []Chat
	c, ok := chats.([]interface{})
	if !ok {
		return
	}
	for _, chat := range c {
		chatNode, ok := chat.(binary.Node)
		if !ok {
			continue
		}

		jid := strings.Replace(chatNode.Attributes["jid"], "@c.us", "@s.whatsapp.net", 1)
		chatList = append(chatList, Chat{
			jid,
			chatNode.Attributes["name"],
			chatNode.Attributes["count"],
			chatNode.Attributes["t"],
			chatNode.Attributes["mute"],
			chatNode.Attributes["spam"],
			chatNode.Attributes["pin"],
			chatNode.Attributes["archive"] == "true",
		})
	}
	for _, h := range wac.handler {
		if x, ok := h.(ChatListHandler); ok {
			if wac.shouldCallSynchronously(h) {
				x.HandleChatList(chatList)
			} else {
				go x.HandleChatList(chatList)
			}
		}
	}
}

// Marshal actions to json and handle with HandleJsonMessages. NEEDS A LOT OF IMRPOVEMENTS. DO IT LATER.
func (wac *Conn) handleActionNode(actionContent interface{}) {

	ac, ok := actionContent.([]binary.Node)
	if !ok {
		return
	}

	for _, action := range ac {

		groupedAction := groupActionsByType(action)
		switch groupedAction.Description {
		case "chat":

			jid := strings.Replace(action.Attributes["jid"], "@c.us", "@s.whatsapp.net", 1)
			chat := wac.Store.Chats[jid]

			switch action.Attributes["type"] {
			case "pin":
				chat.Pin = action.Attributes["pin"]
			case "mute":
				chat.Mute = action.Attributes["mute"]
			case "archive":
				chat.Archive = true
			case "unarchive":
				chat.Archive = false
			case "read":
				chat.Unread = "0"
			case "unread":
				chat.Unread = "-1"
			}

		default:
			continue
		}

		for _, h := range wac.handler {
			if x, ok := h.(ActionHandler); ok {
				if wac.shouldCallSynchronously(h) {
					x.HandleAction(groupedAction)
				} else {
					go x.HandleAction(groupedAction)
				}
			}
		}

	}
}

// Action defines actions struct
type Action struct {
	Description string
	Attributes  map[string]string
	Content     interface{}
}

// WhatsApp does this way, maybe theres a better way to do it.
func groupActionsByType(action binary.Node) Action {

	var groupedAction Action

	switch action.Description {
	case "read":
		groupedAction.Description = "chat"
		groupedAction.Attributes = action.Attributes
		if action.Attributes["type"] == "false" {
			groupedAction.Attributes["type"] = "unread"
		} else {
			groupedAction.Attributes["type"] = "read"
		}
	default:
		groupedAction.Description = action.Description
		groupedAction.Attributes = action.Attributes
		groupedAction.Content = action.Content
	}

	return groupedAction
}

func (wac *Conn) dispatch(msg interface{}) {
	if msg == nil {
		return
	}

	switch message := msg.(type) {

	case *binary.Node:

		if message.Description == "action" {
			fmt.Printf("The type is: %T\n", message.Content)
			if con, ok := message.Content.([]interface{}); ok {

				for a := range con {
					if v, ok := con[a].(*proto.WebMessageInfo); ok {
						wac.handle(v)
						wac.handle(ParseProtoMessage(v))
					}
				}
			} else {
				wac.handleActionNode(message.Content)
			}

		} else if message.Description == "response" && message.Attributes["type"] == "contacts" {
			wac.updateContacts(message.Content)
			wac.handleContacts(message.Content)
			// fmt.Printf("contacts received")
		} else if message.Description == "response" && message.Attributes["type"] == "chat" {
			wac.updateChats(message.Content)
			wac.handleChats(message.Content)
		}
	case error:
		wac.handle(message)
	case string:
		wac.handle(message)
	default:
		fmt.Fprintf(os.Stderr, "unknown type in dipatcher chan: %T", msg)
	}
}
