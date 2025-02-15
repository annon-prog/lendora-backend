package models

// request body received for email messaging
type EmailResponse struct {
	Message    string         `json:"message"`
	Recipient  string         `json:"recipient"`
	Subject    string         `json:"subject"`
	Type       string         `json:"type"`
	Attachment AttachmentData `json:"attachment,omitempty"`
}

// request sub body received for attachments under email request body
type AttachmentData struct {
	Filename    string `json:"filename"`
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
}
