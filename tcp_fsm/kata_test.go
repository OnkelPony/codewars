package kata

import (
	"testing"
)

func TestTraverseTCPStates(t *testing.T) {
	cases := []struct {
		events  []string
		want    string
	}{
		{[]string{"APP_ACTIVE_OPEN","RCV_SYN_ACK","RCV_FIN"}, "CLOSE_WAIT"},
	}

	for _, c := range cases {
		got := TraverseTCPStates(c.events)
		if got != c.want {
			t.Errorf("TraverseTCPStates(%v) => %s, want %s", c.events, got, c.want)
		}
	}
}
