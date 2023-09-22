// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
	"tlan-trust-spaces/pkg/utils"
)

type Conversation struct {
	CreatedDate     time.Time `json:"createdDate"`
	ID              string    `json:"id"`
	LastMessageDate time.Time `json:"lastMessageDate"`
	Messages        []Message `json:"messages"`
	Read            bool      `json:"read"`
	Subject         string    `json:"subject"`
}

func (c Conversation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Conversation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Conversation) GetCreatedDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedDate
}

func (o *Conversation) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Conversation) GetLastMessageDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.LastMessageDate
}

func (o *Conversation) GetMessages() []Message {
	if o == nil {
		return []Message{}
	}
	return o.Messages
}

func (o *Conversation) GetRead() bool {
	if o == nil {
		return false
	}
	return o.Read
}

func (o *Conversation) GetSubject() string {
	if o == nil {
		return ""
	}
	return o.Subject
}
