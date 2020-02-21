package go_aspace

import (
	"github.com/spf13/viper"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUsername(t *testing.T) {
	Init()
	want := "donbot"
	got := viper.GetString("username")
	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
