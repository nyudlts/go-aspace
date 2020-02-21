package go_aspace

import (
	"regexp"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestConfValues(t *testing.T) {
	Init()
	got := len(conf.username)
	if got <= 0 {
		t.Errorf("username string length is %d", got)
	}

	got = len(conf.password)
	if got <= 0 {
		t.Errorf("password string length is 0")
	}

	got = len(conf.url)
	if got <= 0 {
		t.Errorf("url string length is 0")
	}

	var urlPat = regexp.MustCompile(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)
	if !urlPat.MatchString(conf.url) {
		t.Errorf("URL is malformed")
	}
}

func TestGetSessionKey(t *testing.T) {
	Init()
	getSessionKey()
	want := 64
	got := len(SessionKey)
	if want != got {
		t.Errorf("wanted key length of %d, got %d", want, got)
	}

}
