package CMD

import (
	"fmt"
	"sync"
)

type ProcessF func(message Message, writeCh chan Message)
type Matcher interface {
	Register(cmd Cmd, f ProcessF) error
	Process(message Message, writeCh chan Message)
}

func NewMatch() Matcher {
	return &Match{}
}

type Match struct {
	data map[Cmd]ProcessF
	sync.RWMutex
}

func (m *Match) Register(cmd Cmd, f ProcessF) error {
	m.Lock()
	defer m.Unlock()

	if m.data == nil {
		m.data = make(map[Cmd]ProcessF)
	}

	_, exist := m.data[cmd]
	if exist {
		return fmt.Errorf("cmd: %s process func exist", cmd)
	}

	m.data[cmd] = f
	return nil
}

func (m *Match) Process(message Message, writeCh chan Message) {
	m.RLock()
	defer m.RUnlock()

	if f, exist := m.data[message.Pcmd]; exist {
		f(message, writeCh)
		return
	}

	if f, exist := m.data[C_None]; exist {
		f(message, writeCh)
		return
	}
}

