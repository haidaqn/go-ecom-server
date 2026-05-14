package sendmail

import (
	"errors"
	"fmt"
	"strings"

	"github.com/resend/resend-go/v2"
)

var (
	ErrMissingAPIKey  = errors.New("sendmail: RESEND_API_KEY is empty")
	ErrMissingFrom    = errors.New("sendmail: RESEND_FROM is empty")
	ErrNoRecipients   = errors.New("sendmail: no recipients")
	ErrEmptyOTP       = errors.New("sendmail: otp is empty")
	ErrNilMailService = errors.New("sendmail: service is nil")
)

type IMailService interface {
	SendMail(to []string, otp string) error
}

type mailService struct {
	client *resend.Client
	from   string
}

func NewMailService(apiKey, from string) IMailService {
	return &mailService{
		client: resend.NewClient(strings.TrimSpace(apiKey)),
		from:   strings.TrimSpace(from),
	}
}

func normalizeRecipients(recipients []string) []string {
	out := make([]string, 0, len(recipients))
	for _, r := range recipients {
		r = strings.TrimSpace(r)
		if r != "" {
			out = append(out, r)
		}
	}
	return out
}

func (m *mailService) SendMail(to []string, otp string) error {
	if m == nil || m.client == nil {
		return ErrNilMailService
	}
	if strings.TrimSpace(m.client.ApiKey) == "" {
		return ErrMissingAPIKey
	}
	if m.from == "" {
		return ErrMissingFrom
	}
	recipients := normalizeRecipients(to)
	if len(recipients) == 0 {
		return ErrNoRecipients
	}
	otp = strings.TrimSpace(otp)
	if otp == "" {
		return ErrEmptyOTP
	}

	params := &resend.SendEmailRequest{
		From:    m.from,
		To:      recipients,
		Subject: "Mã xác thực tài khoản",
		Html: fmt.Sprintf(`
			<div style="font-family:sans-serif">
				<h2>Xác thực OTP</h2>
				<p>Mã OTP của bạn:</p>
				<h1>%s</h1>
				<p>Mã có hiệu lực trong 5 phút. Không chia sẻ mã này với ai.</p>
			</div>
		`, otp),
		Text: fmt.Sprintf("Mã OTP của bạn là: %s. Hiệu lực 5 phút.", otp),
	}

	_, err := m.client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("sendmail: resend: %w", err)
	}
	return nil
}
