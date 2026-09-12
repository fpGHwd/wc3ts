package peer_test

import (
	"testing"

	"github.com/kradalby/wc3ts/lan"
	"github.com/kradalby/wc3ts/peer"
)

func TestResponderPortDoesNotConflictWithGamePort(t *testing.T) {
	t.Parallel()

	if peer.DefaultResponderPort == lan.DefaultPort {
		t.Fatalf(
			"wc3ts responder port %d conflicts with Warcraft III port %d",
			peer.DefaultResponderPort,
			lan.DefaultPort,
		)
	}
}
