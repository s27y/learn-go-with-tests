package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/s27y/learn-go-with-tests/command-line"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func assertGameStartedWith(t *testing.T, game *poker.GameSpy, numerOfPlayers int) {
	t.Helper()
	if game.StartedWith != numerOfPlayers {
		t.Errorf("wanted Start called with %d but got %d", numerOfPlayers, game.StartedWith)
	}
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func assertFinishCalledWith(t *testing.T, game *poker.GameSpy, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf("wanted Finished with %q but got %q", winner, game.FinishedWith)
	}
}
func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}
		in := userSends("3", "Chris wins")

		poker.NewCLI(in, stdout, game).PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()
		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &poker.GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 0)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)

	})

}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
