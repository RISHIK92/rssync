package mailer

import (
	"fmt"
	"os"

	"github.com/nawafswe/gomailer"
	"github.com/nawafswe/gomailer/message"
)

func SendEmail(sender string, recipients []string, title string, link string, pubDate string, description string) (string, error) {
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	if username == "" || password == "" {
		return "", fmt.Errorf("SMTP credentials are not set in environment variables")
	}
	mail := gomailer.NewMailer("smtp.gmail.com", 587, username, password)
    msg := message.NewMessage()
    msg.From = sender
    msg.Recipients = recipients
    msg.Subject = title
    htmlTemplate := `
	<div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px; color: #333333; line-height: 1.6; border: 1px solid #e0e0e0; border-radius: 8px; background-color: #ffffff;">
		
		<h2 style="margin-top: 0; margin-bottom: 8px; font-size: 22px; line-height: 1.3;">
			<a href="%s" style="color: #2c3e50; text-decoration: none;">%s</a>
		</h2>
		
		<p style="font-size: 13px; color: #888888; margin-top: 0; margin-bottom: 16px; border-bottom: 1px solid #eeeeee; padding-bottom: 12px;">
			Published: %s
		</p>
		
		<p style="font-size: 15px; color: #444444; margin-bottom: 24px;">
			%s
		</p>
		
		<a href="%s" style="display: inline-block; padding: 10px 18px; background-color: #0056b3; color: #ffffff; text-decoration: none; border-radius: 4px; font-weight: bold; font-size: 14px;">
			Read full article &rarr;
		</a>
		
	</div>
	`
	msg.HTMLBody = fmt.Sprintf(htmlTemplate, link, title, pubDate, description, link)
	err := mail.Send(msg)
	if err != nil {
		return "", err
	}
	return "success", nil
}