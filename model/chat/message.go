package chat

type Message struct {
	MessageID      int32  `gorm:"primary_key;not null" json:"message_id"`
	MessageType    int    `gorm:"int" json:"message_type"`
	MessageContent string `json:"message_content"`
}
