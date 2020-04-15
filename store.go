package whatsapp

import (
	"strings"

	"github.com/Rhymen/go-whatsapp/binary"
)

type Store struct {
	Contacts map[string]Contact
	Chats    map[string]Chat
}

type Contact struct {
	Jid    string
	Notify string
	Name   string
	Short  string

	// Jid           string
	// Name          string
	// Type          string
	// ShortName     string
	// Pushname      string
	// StatusMute    string
	// SectionHeader string
	// VerifiedName  string
	// // Labels string
	// IsEnterprise string
}

type Chat struct {
	Jid             string
	Name            string
	Unread          string
	LastMessageTime string
	Mute            string
	IsMarkedSpam    string
	Pin             string
	Archive         bool
	// isReadOnly      string

	// Jid                string
	// T                  string
	// Type               string
	// Kind               string
	// Keys               string
	// Before             string
	// Archive            bool
	// IsReadOnly         bool
	// UnreadCount        string
	// MuteExpiration     string
	// ModifyTag          string
	// Name               string
	// PendingMsgs        bool
	// Star               bool
	// NotSpam            bool
	// Pin                string
	// ChangeNumberOldJid string
	// ChangeNumberNewJid string
	// EphemeralDuration  string
	// // Labels string
}

func newStore() *Store {
	return &Store{
		make(map[string]Contact),
		make(map[string]Chat),
	}
}

func (wac *Conn) updateContacts(contacts interface{}) {
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
		wac.Store.Contacts[jid] = Contact{
			jid,
			contactNode.Attributes["notify"],
			contactNode.Attributes["name"],
			contactNode.Attributes["short"],
		}
	}
}

func (wac *Conn) updateChats(chats interface{}) {
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
		wac.Store.Chats[jid] = Chat{
			jid,
			chatNode.Attributes["name"],
			chatNode.Attributes["count"],
			chatNode.Attributes["t"],
			chatNode.Attributes["mute"],
			chatNode.Attributes["spam"],
			chatNode.Attributes["pin"],
			chatNode.Attributes["archive"] == "true",
		}
	}
}
