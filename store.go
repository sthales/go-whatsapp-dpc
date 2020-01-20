package whatsapp

type Store struct {
	Contacts map[string]Contact
	Chats    map[string]Chat
}

type Contact struct {
	// Jid    string
	// Notify string
	// Name   string
	// Short  string

	Jid           string
	Name          string
	Type          string
	ShortName     string
	Pushname      string
	StatusMute    string
	SectionHeader string
	VerifiedName  string
	// Labels string
	IsEnterprise string
}

type Chat struct {
	// Jid             string
	// Name            string
	// Unread          string
	// LastMessageTime string
	// IsMuted         string
	// IsMarkedSpam    string
	// isReadOnly string

	Jid                string
	T                  string
	Type               string
	Kind               string
	Keys               string
	Before             string
	Archive            bool
	IsReadOnly         bool
	UnreadCount        string
	MuteExpiration     string
	ModifyTag          string
	Name               string
	PendingMsgs        bool
	Star               bool
	NotSpam            bool
	Pin                string
	ChangeNumberOldJid string
	ChangeNumberNewJid string
	EphemeralDuration  string
	// Labels string
}

func newStore() *Store {
	return &Store{
		make(map[string]Contact),
		make(map[string]Chat),
	}
}

// func (wac *Conn) updateContacts(contacts interface{}, jid string) {

// 	wac.Store.Contacts[jid] = contacts

// 	// c, ok := contacts.([]interface{})
// 	// if !ok {
// 	// 	return
// 	// }

// 	// for _, contact := range c {
// 	// 	contactNode, ok := contact.(binary.Node)
// 	// 	if !ok {
// 	// 		continue
// 	// 	}

// 	// 	jid := strings.Replace(contactNode.Attributes["jid"], "@c.us", "@s.whatsapp.net", 1)
// 	// 	wac.Store.Contacts[jid] = Contact{
// 	// 		contactNode.Attributes["jid"],
// 	// 		contactNode.Attributes["name"],
// 	// 		contactNode.Attributes["type"],
// 	// 		contactNode.Attributes["short"],
// 	// 		contactNode.Attributes["notify"],
// 	// 		"",
// 	// 		"",
// 	// 		contactNode.Attributes["status_mute"],
// 	// 		contactNode.Attributes["index"],
// 	// 		contactNode.Attributes["vname"],
// 	// 		contactNode.Attributes["enterprise"],
// 	// 	}
// 	// }
// }

// func (wac *Conn) updateChats(chats interface{}) {
// 	c, ok := chats.([]interface{})
// 	if !ok {
// 		return
// 	}

// 	for _, chat := range c {
// 		chatNode, ok := chat.(binary.Node)
// 		if !ok {
// 			continue
// 		}

// 		jid := strings.Replace(chatNode.Attributes["jid"], "@c.us", "@s.whatsapp.net", 1)
// 		wac.Store.Chats[jid] = Chat{
// 			chatNode.Attributes["jid"],
// 			chatNode.Attributes["t"],
// 			chatNode.Attributes["type"],
// 			chatNode.Attributes["kind"],
// 			"",
// 			chatNode.Attributes["before"],
// 			chatNode.Attributes["archive"],
// 			chatNode.Attributes["read_only"],
// 			chatNode.Attributes["count"],
// 			chatNode.Attributes["mute"],
// 			chatNode.Attributes["modify_tag"],
// 			chatNode.Attributes["name"],
// 			chatNode.Attributes["message"],
// 			chatNode.Attributes["star"],
// 			chatNode.Attributes["spam"],
// 			chatNode.Attributes["pin"],
// 			chatNode.Attributes["old_jid"],
// 			chatNode.Attributes["new_jid"],
// 			chatNode.Attributes["ephemeral"],
// 		}
// 	}
// }
