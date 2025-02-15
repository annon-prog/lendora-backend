package models

type EmailResponse struct {
	Message    string         `json:"message"`
	Recipient  string         `json:"recipient"`
	Subject    string         `json:"subject"`
	Type       string         `json:"type"`
	Attachment AttachmentData `json:"attachment,omitempty"`
}

type AttachmentData struct {
	Filename    string `json:"filename"`
	Content     string `json:"content"` // base64 encoded content
	ContentType string `json:"contentType"`
}
