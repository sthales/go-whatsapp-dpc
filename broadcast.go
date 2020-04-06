package whatsapp

// GetBroadcastMetadata returns broadcastlist recliplents, status and broadcast name
func (wac *Conn) GetBroadcastMetadata(jid string) (<-chan string, error) {
	data := []interface{}{"query", "contact", jid}
	return wac.writeJson(data)
}
