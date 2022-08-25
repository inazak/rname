package rname

import (
  "testing"
)


func TestPrependZeros(t *testing.T) {

  ps := []struct{
    Sample   string
    Width    int
    Expect string
  }{
    {
      Sample:  "1",
      Width:   4,
      Expect:"0001",
    },
    {
      Sample:  "abc99",
      Width:   4,
      Expect:"abc0099",
    },
    {
      Sample:  "99xyz",
      Width:   4,
      Expect:"0099xyz",
    },
    {
      Sample:  "abc55xyz",
      Width:   3,
      Expect:"abc055xyz",
    },
    {
      Sample:  "abcxyz",
      Width:   3,
      Expect:"abcxyz",
    },
    {
      Sample:  "12abc45xyz",
      Width:   3,
      Expect:"12abc045xyz",
    },
  }

  for _, p := range ps {
    out := prependZeros(p.Sample, p.Width)
    if out != p.Expect {
      t.Errorf("prependZeros('%s', %d) expected=%s got=%s",
        p.Sample, p.Width, p.Expect, out)
    }
  }
}

