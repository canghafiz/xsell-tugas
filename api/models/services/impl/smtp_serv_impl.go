package impl

import (
	"be/models/domains"
	"crypto/tls"
	"fmt"
	"log"
	"mime"
	"net/smtp"
	"strings"
)

type SmtpServImpl struct {
	Smtp    domains.Smtp
	AppName string
}

func NewSmtpServImpl(smtp domains.Smtp, appName string) *SmtpServImpl {
	return &SmtpServImpl{Smtp: smtp, AppName: appName}
}

func (serv *SmtpServImpl) SendEmailOtp(email string, code string) error {
	subject := fmt.Sprintf("Verification Code - %s", serv.AppName)

	// Encode subject agar aman
	encodedSubject := mime.QEncoding.Encode("utf-8", subject)

	htmlBody := fmt.Sprintf(`
	<html>
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>Your OTP Code</title>
	</head>
	<body style="margin: 0; padding: 0; background-color: #f9f9f9; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;">
	  <table width="100%%" cellpadding="0" cellspacing="0" style="background-color: #f9f9f9; padding: 20px;">
	    <tr>
	      <td align="center">
	        <table width="100%%" max-width="600" cellpadding="0" cellspacing="0" style="background-color: #ffffff; border-radius: 12px; overflow: hidden; box-shadow: 0 4px 12px rgba(0,0,0,0.08);">
	          <tr>
	            <td style="padding: 32px 24px 24px; text-align: center; background: linear-gradient(135deg, #c00 0%%, #e53935 100%%); color: white;">
	              <h1 style="margin: 0; font-size: 24px; font-weight: 600;">One-Time Password</h1>
	            </td>
	          </tr>
	          <tr>
	            <td style="padding: 24px;">
	              <p style="margin: 0 0 16px; color: #333; font-size: 16px;">Hello,</p>
	              <p style="margin: 0 0 24px; color: #555; font-size: 16px;">Use this verification code to complete your request:</p>
	              <div style="background-color: #fff5f5; border: 2px dashed #ef9a9a; border-radius: 10px; padding: 16px; margin: 24px 0; display: inline-block;">
	                <strong style="font-size: 28px; letter-spacing: 8px; color: #c62828; font-family: monospace;">%s</strong>
	              </div>
	              <p style="margin: 0 0 16px; color: #555; font-size: 15px;">This code is valid for <strong>5 minutes</strong>.</p>
	              <p style="margin: 0; color: #888; font-size: 14px;"><small>Do not share this code with anyone.</small></p>
	            </td>
	          </tr>
	          <tr>
	            <td style="padding: 16px 24px; background-color: #f5f5f5; color: #777; font-size: 13px; text-align: center; border-top: 1px solid #eee;">
	              © 2025 Your App. All rights reserved.<br>This is an automated message. Please do not reply.
	            </td>
	          </tr>
	        </table>
	      </td>
	    </tr>
	  </table>
	</body>
	</html>`, code)

	// ⚠️ Escape % di HTML dengan %%
	// Tapi lebih baik gunakan raw string + fmt.Sprintf terpisah

	// Buat header dan body dengan pemisah yang benar
	msg := "From: " + serv.Smtp.From + "\r\n" +
		"To: " + email + "\r\n" +
		"Subject: " + encodedSubject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" + // <-- PENTING: dua CRLF sebelum body
		htmlBody

	addr := fmt.Sprintf("%s:%s", serv.Smtp.Host, serv.Smtp.Port)
	client, err := smtp.Dial(addr)
	if err != nil {
		log.Printf("[SMTP] Dial error: %v", err)
		return serv.handleSmtpError(err, "connection")
	}
	defer client.Close()

	if err = client.Hello("localhost"); err != nil {
		return serv.handleSmtpError(err, "hello")
	}

	// STARTTLS untuk port 587
	if serv.Smtp.Port == "587" {
		if ok, _ := client.Extension("STARTTLS"); ok {
			tlsConfig := &tls.Config{
				ServerName: serv.Smtp.Host,
				// JANGAN gunakan InsecureSkipVerify di production
			}
			if err = client.StartTLS(tlsConfig); err != nil {
				return serv.handleSmtpError(err, "tls")
			}
		}
	}

	if serv.Smtp.Username != "" {
		auth := smtp.PlainAuth("", serv.Smtp.Username, serv.Smtp.Password, serv.Smtp.Host)
		if err = client.Auth(auth); err != nil {
			return serv.handleSmtpError(err, "auth")
		}
	}

	if err = client.Mail(serv.Smtp.From); err != nil {
		return serv.handleSmtpError(err, "sender")
	}
	if err = client.Rcpt(email); err != nil {
		return serv.handleSmtpError(err, "recipient")
	}

	writer, err := client.Data()
	if err != nil {
		return serv.handleSmtpError(err, "data")
	}
	defer writer.Close()

	if _, err = writer.Write([]byte(msg)); err != nil {
		return serv.handleSmtpError(err, "write")
	}

	log.Printf("[SMTP] Email sent to %s", email)
	return nil
}

func (serv *SmtpServImpl) handleSmtpError(err error, stage string) error {
	errStr := err.Error()

	// Network/Connection errors
	if stage == "connection" {
		if strings.Contains(errStr, "connection refused") {
			return fmt.Errorf("mail server unavailable, please check SMTP_HOST and SMTP_PORT configuration")
		}
		if strings.Contains(errStr, "timeout") || strings.Contains(errStr, "deadline exceeded") {
			return fmt.Errorf("connection timeout, mail server not responding")
		}
		if strings.Contains(errStr, "no such host") {
			return fmt.Errorf("invalid SMTP host: %s", serv.Smtp.Host)
		}
		return fmt.Errorf("failed to connect to mail server: %v", err)
	}

	// TLS errors
	if stage == "tls" {
		if strings.Contains(errStr, "certificate") {
			return fmt.Errorf("TLS certificate error, please verify SMTP server configuration")
		}
		if strings.Contains(errStr, "handshake") {
			return fmt.Errorf("TLS handshake failed, SMTP server may not support STARTTLS on port %s", serv.Smtp.Port)
		}
		return fmt.Errorf("TLS connection failed: %v", err)
	}

	// Authentication errors
	if stage == "auth" {
		if strings.Contains(errStr, "535") {
			return fmt.Errorf("authentication failed: invalid SMTP username or password")
		}
		if strings.Contains(errStr, "530") {
			return fmt.Errorf("authentication required: SMTP server requires valid credentials")
		}
		if strings.Contains(errStr, "534") {
			return fmt.Errorf("authentication mechanism not supported by SMTP server")
		}
		return fmt.Errorf("authentication failed: %v", err)
	}

	// Sender/Recipient errors
	if stage == "sender" {
		if strings.Contains(errStr, "550") || strings.Contains(errStr, "553") {
			return fmt.Errorf("sender address rejected: %s", serv.Smtp.From)
		}
		return fmt.Errorf("failed to set sender: %v", err)
	}

	if stage == "recipient" {
		if strings.Contains(errStr, "550") {
			return fmt.Errorf("recipient address rejected or does not exist")
		}
		if strings.Contains(errStr, "551") {
			return fmt.Errorf("recipient not local, relay denied")
		}
		if strings.Contains(errStr, "552") {
			return fmt.Errorf("mailbox full or quota exceeded")
		}
		if strings.Contains(errStr, "554") {
			return fmt.Errorf("transaction failed, recipient address may be invalid")
		}
		return fmt.Errorf("failed to set recipient: %v", err)
	}

	// Data transfer errors
	if stage == "data" || stage == "write" {
		if strings.Contains(errStr, "552") {
			return fmt.Errorf("message size exceeds limit")
		}
		if strings.Contains(errStr, "554") {
			return fmt.Errorf("message rejected by server")
		}
		return fmt.Errorf("failed to send email data: %v", err)
	}

	// Generic SMTP errors
	if stage == "hello" {
		return fmt.Errorf("SMTP handshake failed: %v", err)
	}

	// Default
	return fmt.Errorf("smtp error at %s stage: %v", stage, err)
}
