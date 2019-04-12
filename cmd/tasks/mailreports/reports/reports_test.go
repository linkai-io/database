package reports_test

import (
	"os"
	"testing"

	"github.com/linkai-io/am/mock"
	"github.com/linkai-io/am/pkg/initializers"
	"github.com/linkai-io/database/cmd/tasks/mailreports/reports"
	"github.com/rs/zerolog/log"
)

var (
	appConfig initializers.AppConfig
	runType   string
	weekly    bool
)

func TestGenerateReports(t *testing.T) {
	//t.Skip("you must populate the database before running this test")
	appConfig.Env = os.Getenv("APP_ENV")
	appConfig.Region = os.Getenv("APP_REGION")
	appConfig.SelfRegister = os.Getenv("APP_SELF_REGISTER")
	appConfig.ServiceKey = "linkai_admin"
	runType = os.Getenv("APP_RUN_TYPE")
	if runType == "weekly" {
		weekly = true
	}

	m := &mock.Mailer{}
	m.InitFn = func(config []byte) error {
		return nil
	}

	m.SendMailFn = func(subject, to, html, text string) error {
		t.Logf("Sending %s to %s with body\n%s\n", subject, to, html)
		if subject != "Your daily hakken service email report" {
			t.Fatalf("wrong subject headline")
		}
		return nil
	}

	log.Logger = log.With().Str("service", "Email Reporting Task").Logger()

	reports.GenerateReports(appConfig, false, m)
	if m.SendMailInvoked == false {
		t.Fatalf("error invoking email")
	}

	m.SendMailInvoked = false
	m.SendMailFn = func(subject, to, html, text string) error {
		t.Logf("Sending %s to %s with body\n%s\n", subject, to, html)
		if subject != "Your weekly hakken service email report" {
			t.Fatalf("wrong subject headline")
		}
		return nil
	}

	log.Logger = log.With().Str("service", "Email Reporting Task").Logger()

	reports.GenerateReports(appConfig, true, m)
	if m.SendMailInvoked == false {
		t.Fatalf("error invoking email")
	}
}
