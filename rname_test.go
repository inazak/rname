package rname

import (
  "testing"
)


func TestPrependZeros(t *testing.T) {

  ps := []struct{
    Sample   string
    Width    int
    Expected string
  }{
    {
      Sample:  "1",
      Width:   4,
      Expected:"0001",
    },
    {
      Sample:  "abc99",
      Width:   4,
      Expected:"abc0099",
    },
    {
      Sample:  "99xyz",
      Width:   4,
      Expected:"0099xyz",
    },
    {
      Sample:  "abc55xyz",
      Width:   3,
      Expected:"abc055xyz",
    },
    {
      Sample:  "abcxyz",
      Width:   3,
      Expected:"abcxyz",
    },
    {
      Sample:  "abc1234xyz",
      Width:   3,
      Expected:"abc1234xyz",
    },
  }

  for _, p := range ps {
    out := prependZeros(p.Sample, p.Width)
    if out != p.Expected {
      t.Errorf("prependZeros('%s', %d) expected=%s got=%s",
        p.Sample, p.Width, p.Expected, out)
    }
  }
}

