package c3

import (
	"github.com/lovelaced/glitzz/config"
	"github.com/lovelaced/glitzz/modules"
	"testing"
	"time"
)

func TestGetDaysToNextCongressAndCongressNumber(t *testing.T) {
	now := createDate(2018, time.December, 20)
	days, number := getDaysToNextCongressAndCongressNumber(now)
	if days != 7 {
		t.Errorf("Days was %d", days)
	}

	if number != 35 {
		t.Errorf("Number was %d", number)
	}
}

func TestGetDaysToNextCongressAndCongressNumberAfter(t *testing.T) {
	now := createDate(2018, time.December, 29)
	days, number := getDaysToNextCongressAndCongressNumber(now)
	if days != 363 {
		t.Errorf("Days was %d", days)
	}

	if number != 36 {
		t.Errorf("Number was %d", number)
	}
}

func TestC3(t *testing.T) {
	p := New(nil, config.Default())
	output, err := p.RunCommand(modules.Command{Text: ".c3", Nick: "nick"})
	if err != nil {
		t.Errorf("error was not nil %s", err)
	}
	if len(output) != 1 {
		t.Errorf("invalid output length %d", len(output))
	}
}
