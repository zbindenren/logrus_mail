package logrus_mail

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewMailAuthHook(t *testing.T) {

	// invalid port
	_, err := NewMailAuthHook("testapp", "smtp.gmail.com", 10, "user.name@gmail.com", "user.name@gmail.com", "user.name", "password", []logrus.Level{})
	if err == nil {
		t.Errorf("no error on invalid port")
	}

	// invalid mail host
	_, err = NewMailAuthHook("testapp", "www.gmail.com", 587, "user.name@gmail.com", "user.name@gmail.com", "user.name", "password", []logrus.Level{logrus.InfoLevel})
	if err == nil {
		t.Errorf("no error on invalid hostname")
	}

	// invalid email address
	_, err = NewMailAuthHook("testapp", "smtp.gmail.com", 587, "user.name", "user.name@gmail.com", "user.name", "password", []logrus.Level{logrus.ErrorLevel})
	if err == nil {
		t.Errorf("no error on invalid email address")
	}

}
