package gocrypt

import "strings"

func aao(in string) string{ //åäö
  in = strings.Replace(in, "å", "a", -1)
  in = strings.Replace(in, "ä", "a", -1)
  in = strings.Replace(in, "ö", "o", -1)
  return in
}
