package CMD

func NewPingMessage() Message {
	return newMessage(C_Ping, C_None, 0, nil)
}
