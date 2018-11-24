package CMD

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Message struct {
	Pcmd Cmd
	Scmd Cmd
	Size uint16
	Data []byte
}

type Writer interface {
	Write(message Message)
}
type Reader interface {
	Read(message *Message) error
}

func newMessage(pcmd, scmd Cmd, size uint16, data []byte) Message {
	return Message{Pcmd:pcmd, Scmd:scmd, Size:size, Data:data}
}

func NewUnRegisterMessage(message Message) Message {
	data := []byte(fmt.Sprintf("unknow Message: %s", message))
	return newMessage(C_Err, C_None, uint16(len(data)), data)
}

func (m *Message) WriteTo(w io.Writer) (err error) {
	_, err = w.Write(m.Header())
	if err != nil {
		return err
	}

	_, err = w.Write(m.Data)
	return err
}

func ReadFrom(r io.Reader) (Message, error) {
	header := make([]byte, 4)
	_, err := r.Read(header)
	if err != nil {
		return Message{}, err
	}

	msg := Message{}
	msg.UnmarshalBinary(header)

	buf := make([]byte, msg.Size)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return Message{}, err
	}

	return msg, nil
}

func (m *Message) Header() []byte {
	data := make([]byte, 4)
	data[0] = byte(m.Pcmd)
	data[1] = byte(m.Scmd)
	binary.BigEndian.PutUint16(data, m.Size)
	return data
}

func (m *Message) UnmarshalBinary(data []byte) error {
	m.Pcmd = Cmd(data[0])
	m.Scmd = Cmd(data[1])
	m.Size = binary.BigEndian.Uint16(data[2:])
	return nil
}