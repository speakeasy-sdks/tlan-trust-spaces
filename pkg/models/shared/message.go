// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
	"tlan-trust-spaces/pkg/utils"
)

type Message struct {
	ConversationID string            `json:"conversationId"`
	CreatedDate    time.Time         `json:"createdDate"`
	Documents      []MessageDocument `json:"documents"`
	ID             string            `json:"id"`
	Message        string            `json:"message"`
	Sender         User              `json:"sender"`
}

func (m Message) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *Message) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Message) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

func (o *Message) GetCreatedDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedDate
}

func (o *Message) GetDocuments() []MessageDocument {
	if o == nil {
		return []MessageDocument{}
	}
	return o.Documents
}

func (o *Message) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Message) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *Message) GetSender() User {
	if o == nil {
		return User{}
	}
	return o.Sender
}
