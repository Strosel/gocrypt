package gocrypt

import "strings"

//ceasar cipher
func Caesar_enc(s int, m string) string{ // encode m(essage) with s(tep)
  m = aao(strings.ToLower(m))
  enc := ""
  for i := 0; i<len(m); i++{
    if string(m[i]) == " "{
      enc += " "
    }else if m[i] - byte(s) < byte('a'){
      enc += string(byte(26)+m[i]-byte(s)) // move to the end of the alfabet (+26) and go
    }else{
      enc += string(m[i]-byte(s))
    }
  }
  return enc
}

func Caesar_dec(s int, m string) string{ // decode m(essage) with s(tep)
  m = aao(strings.ToLower(m))
  dec := ""
  for i := 0; i<len(m); i++{
    if string(m[i]) == " "{
      dec += " "
    }else if m[i] + byte(s) > byte('z'){
      dec += string(m[i]-byte(26)+byte(s)) // move to the start of the alfabet (-26) and go
    }else{
      dec += string(m[i]+byte(s))
    }
  }
  return dec
}
