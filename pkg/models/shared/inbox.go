// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Inbox - ok
type Inbox struct {
	Conversations []ConversationListItem `json:"conversations"`
}

func (o *Inbox) GetConversations() []ConversationListItem {
	if o == nil {
		return []ConversationListItem{}
	}
	return o.Conversations
}
