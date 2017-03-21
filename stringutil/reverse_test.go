package stringutil

import "testing"

func TestNothing(t *testing.T) {
    //t.Errorf("Failure!")
}

func TestReverse(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := Reverse(c.in)
        if got != c.want {
            t.Errorf("Reverse (%q) expected=%q actual=%q", c.in, c.want, got)
        }
    }
    //t.Errorf("Failure 2!")
}
