package CMD

type Cmd int8

const (
	C_None Cmd = iota
	C_Err
	C_Ping
	C_Pong
)
