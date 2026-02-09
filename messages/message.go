// Package messages provides translation message bundles with pluralization support.
package messages

import (
	"encoding/json"
	"fmt"
)

// MessageSet holds messages for a single locale.
type MessageSet struct {
	tag      string
	messages map[string]*Message
}

// NewMessageSet creates an empty MessageSet for the given locale tag.
func NewMessageSet(tag string) *MessageSet {
	return &MessageSet{
		tag:      tag,
		messages: make(map[string]*Message),
	}
}

// Get retrieves a message by ID, returning nil if not found.
func (ms *MessageSet) Get(id string) *Message {
	return ms.messages[id]
}

// Set adds or updates a message.
func (ms *MessageSet) Set(m *Message) {
	ms.messages[m.ID] = m
}

// Message represents a single translatable message.
type Message struct {
	ID          string `json:"id"`
	Translation any    `json:"translation"` // string or PluralTranslations map
}

// PluralTranslations holds CLDR plural category translations.
type PluralTranslations struct {
	Zero  string `json:"zero,omitempty"`
	One   string `json:"one,omitempty"`
	Two   string `json:"two,omitempty"`
	Few   string `json:"few,omitempty"`
	Many  string `json:"many,omitempty"`
	Other string `json:"other"`
}

// IsPlural returns true if this message has plural translations.
func (m *Message) IsPlural() bool {
	_, ok := m.Translation.(map[string]any)
	return ok
}

// GetSingular returns the translation as a simple string.
// For plural messages, returns the "other" form.
func (m *Message) GetSingular() string {
	switch v := m.Translation.(type) {
	case string:
		return v
	case map[string]any:
		if other, ok := v["other"].(string); ok {
			return other
		}
	}
	return ""
}

// GetPlural returns the plural translations, or nil if not plural.
func (m *Message) GetPlural() *PluralTranslations {
	v, ok := m.Translation.(map[string]any)
	if !ok {
		return nil
	}

	pt := &PluralTranslations{}
	if s, ok := v["zero"].(string); ok {
		pt.Zero = s
	}
	if s, ok := v["one"].(string); ok {
		pt.One = s
	}
	if s, ok := v["two"].(string); ok {
		pt.Two = s
	}
	if s, ok := v["few"].(string); ok {
		pt.Few = s
	}
	if s, ok := v["many"].(string); ok {
		pt.Many = s
	}
	if s, ok := v["other"].(string); ok {
		pt.Other = s
	}
	return pt
}

// MessagesFile represents the JSON structure for a messages file.
// This format is compatible with go-i18n.
type MessagesFile struct {
	Messages []Message `json:"messages"`
}

// ParseMessagesJSON parses a go-i18n compatible JSON messages file.
func ParseMessagesJSON(data []byte) (*MessagesFile, error) {
	var mf MessagesFile
	if err := json.Unmarshal(data, &mf); err != nil {
		return nil, fmt.Errorf("parsing messages JSON: %w", err)
	}
	return &mf, nil
}
