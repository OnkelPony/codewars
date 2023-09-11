package kata

import (
    "testing"

)

func TestRGB(t *testing.T) {
    cases := []struct {
        r, g, b int
        want    string
    }{
        {0, 0, 0, "000000"},
        {255, 255, 255, "FFFFFF"},
        {254, 253, 252, "FEFDFC"},
        {128, 128, 128, "808080"},
    }

    for _, c := range cases {
        got := RGB(c.r, c.g, c.b)
        if got != c.want {
            t.Errorf("RGB(%d, %d, %d) => %q, want %q", c.r, c.g, c.b, got, c.want)
        }
    }
}