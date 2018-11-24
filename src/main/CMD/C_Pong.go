package CMD

func NewPongMessage() Message {
	return newMessage(C_Pong, C_None, 0, nil)
}
