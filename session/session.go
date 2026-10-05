package session

import (
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Key struct {
	GuildID string
	UserID  string
}

type Handler interface {
	Handle(msg *discordgo.MessageCreate) (done bool, err error)
}

type Entry struct {
	Handler   Handler
	ExpiresAt time.Time
}

type Manager struct {
	mu       sync.RWMutex
	sessions map[Key]Entry
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[Key]Entry),
	}
}

func (m *Manager) Set(guildID, userID string, handler Handler, ttl time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.sessions[Key{
		GuildID: guildID,
		UserID:  userID,
	}] = Entry{
		Handler:   handler,
		ExpiresAt: time.Now().Add(ttl),
	}
}

func (m *Manager) Get(guildID, userID string) (Handler, bool) {
	key := Key{
		GuildID: guildID,
		UserID:  userID,
	}

	m.mu.RLock()
	entry, ok := m.sessions[key]
	m.mu.RUnlock()

	if !ok {
		return nil, false
	}

	if time.Now().After(entry.ExpiresAt) {
		m.Delete(guildID, userID)
		return nil, false
	}

	return entry.Handler, true
}

func (m *Manager) Delete(guildID, userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.sessions, Key{
		GuildID: guildID,
		UserID:  userID,
	})
}
