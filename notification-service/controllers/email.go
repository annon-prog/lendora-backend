package controllers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"notification-service/models"

	"gopkg.in/gomail.v2"

	"github.com/gin-gonic/gin"
)

/**
 * @Function Name: SendEmails
 *
 * @Description:
 * Handler function used to send out email notifications
 *
 * @Params:
 * c *gin.Context
 *
 * @Returns:
 */
func SendEmails(c *gin.Context) {

	var data models.EmailResponse

	//define our smtp credentials
	from := os.Getenv("FROM_ADDRESS")
	pass := os.Getenv("EMAIL_PASS")
	host := os.Getenv("EMAIL_HOST")
	portNumber := os.Getenv("EMAIL_PORT")

	port, err := strconv.Atoi(portNumber)
	if err != nil {
		fmt.Printf("Failed to convert port to int: %v", port)
	}

	//fetch message and recipient from request body
	//parse request body
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Printf("Failed to parse request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to parse request body:%v", err.Error()),
		})
	}

	//create a new message
	message := gomail.NewMessage()

	//set email headers
	message.SetHeader("From", from)
	message.SetHeader("To", data.Recipient)
	message.SetHeader("Subject", data.Subject)

	//set email body
	switch data.Type {
	case "plain":
		message.SetBody("text/plain", data.Message)
	case "html":
		message.SetBody("text/html", data.Message)
	default:
		message.SetBody("text/plain", data.Message)

	}

	// Handle base64 attachment if present
	if data.Attachment.Content != "" {
		// Decode base64 content
		fileData, err := base64.StdEncoding.DecodeString(data.Attachment.Content)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Invalid base64 attachment: %v", err),
			})
			return
		}

		// Attach the decoded file
		message.Attach(data.Attachment.Filename,
			gomail.SetCopyFunc(func(w io.Writer) error {
				_, err := w.Write(fileData)
				return err
			}),
			gomail.SetHeader(map[string][]string{
				"Content-Type": {data.Attachment.ContentType},
			}),
		)
	}

	//set up the smtp dialer
	dialer := gomail.NewDialer(host, port, from, pass)

	//send out the email
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Printf("Failed to send out email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to send out email:%v", err),
		})
	} else {

		//if successful, log and return success message
		fmt.Printf("Successfully sent out email to %s", data.Recipient)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "successfully sent out message",
		})

	}

}
