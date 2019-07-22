package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/s27y/learn-go-with-tests/command-line"
)

func TestBlindAlerter(t *testing.T) {

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}

				got := blindAlerter.Alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}

	})

}

func assertScheduledAlert(t *testing.T, got poker.ScheduledAlert, want poker.ScheduledAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf("got amount %d want %d", got.Amount, want.Amount)
	}
	if got.At != want.At {
		t.Errorf("got scheduled time of %v want %v", got.At, want.At)
	}

}
