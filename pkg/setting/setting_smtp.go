package setting

import (
	"fmt"

	"github.com/grafana/grafana/pkg/util"
	"github.com/prometheus/common/model"
)

type SmtpSettings struct {
	Enabled        bool
	Host           string
	User           string
	Password       string
	CertFile       string
	KeyFile        string
	FromAddress    string
	FromName       string
	EhloIdentity   string
	StartTLSPolicy string
	SkipVerify     bool
	StaticHeaders  map[string]string

	SendWelcomeEmailOnSignUp bool
	TemplatesPatterns        []string
	ContentTypes             []string
}

func (cfg *Cfg) readSmtpSettings() error {
	sec := cfg.Raw.Section("smtp")
	cfg.Smtp.Enabled = sec.Key("enabled").MustBool(false)
	cfg.Smtp.Host = sec.Key("host").String()
	cfg.Smtp.User = sec.Key("user").String()
	cfg.Smtp.Password = sec.Key("password").String()
	cfg.Smtp.CertFile = sec.Key("cert_file").String()
	cfg.Smtp.KeyFile = sec.Key("key_file").String()
	cfg.Smtp.FromAddress = sec.Key("from_address").String()
	cfg.Smtp.FromName = sec.Key("from_name").String()
	cfg.Smtp.EhloIdentity = sec.Key("ehlo_identity").String()
	cfg.Smtp.StartTLSPolicy = sec.Key("startTLS_policy").String()
	cfg.Smtp.SkipVerify = sec.Key("skip_verify").MustBool(false)

	emails := cfg.Raw.Section("emails")
	cfg.Smtp.SendWelcomeEmailOnSignUp = emails.Key("welcome_email_on_sign_up").MustBool(false)
	cfg.Smtp.TemplatesPatterns = util.SplitString(emails.Key("templates_pattern").MustString("emails/*.html, emails/*.txt"))
	cfg.Smtp.ContentTypes = util.SplitString(emails.Key("content_types").MustString("text/html"))

	// populate static headers
	if err := cfg.readGrafanaSmtpStaticHeaders(); err != nil {
		return err
	}

	return nil
}

func (cfg *Cfg) readGrafanaSmtpStaticHeaders() error {
	staticHeadersSection := cfg.Raw.Section("smtp.static_headers")
	keys := staticHeadersSection.Keys()
	cfg.Smtp.StaticHeaders = make(map[string]string, len(keys))

	for _, key := range keys {
		labelName := model.LabelName(key.Name())
		labelValue := model.LabelValue(key.Value())

		if !labelName.IsValid() {
			return fmt.Errorf("invalid label name in [smtp.static_headers] configuration. name %q", labelName)
		}

		if !labelValue.IsValid() {
			return fmt.Errorf("invalid label value in [smtp.static_headers] configuration. name %q value %q", labelName, labelValue)
		}

		cfg.Smtp.StaticHeaders[string(labelName)] = string(labelValue)
	}

	return nil
}
