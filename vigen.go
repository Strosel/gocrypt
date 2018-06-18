package gocrypt

import (
  "math"
  "strings"
)

//Vigenère cipher
func Vigen_enc(k, m string) string{ //encode with k(ey) and m(essage)
  k = aao(strings.ToLower(k))
  m = aao(strings.ToLower(m))

  r := rep(k, m)
  enc := ""
  for i := 0; i<len(m); i++{
    enc += t[string(r[i])][string(m[i])]
  }
  return enc
}

func Vigen_dec(k, m string) string{ //decode with k(ey) and m(essage)
  k = aao(strings.ToLower(k))
  m = aao(strings.ToLower(m))

  r := rep(k, m)
  dec := ""
  for i := 0; i<len(m); i++{
    for k, v := range t[string(r[i])]{
      if v == string(m[i]){
        dec += k
      }
    }
  }
  return dec
}

func rep(k, m string) string{ // repeat the k(ey) to fit the m(essage)
  r := mult( k, int( math.Floor( float64( len(m)/len(k) ))))
  r += k[:len(m)-len(r)]
  return r
}

func mult(s string, i int) string{ // string multiplication
  out := ""
  for j := 0; j<i; j++{
    out += s
  }
  return out
}

var t map[string]map[string]string

func init(){ //generate translation table
  t = make(map[string]map[string]string)
  alf := []byte{
    'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
    'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
  for i, v := range alf{
    t[string(v)] = make(map[string]string)
    t[string(v)][" "] = " "
    for _, w := range alf{
      if byte(i) + w <= 122{
        t[string(v)][string(w)] = string(w + byte(i))
      }else{
        t[string(v)][string(w)] = string(w + byte(i) - 26)
      }
    }
  }
}
