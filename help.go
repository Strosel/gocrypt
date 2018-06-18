package gocrypt

import "fmt"

func Help(){
  fmt.Println("Caesar cipher encrypt - Caesar_enc(s int, m string) string")
  fmt.Println("Caesar cipher decrypt - Caesar_dec(s int, m string) string")
  fmt.Println("Vigenère cipher encrypt - Vigen_enc(k, m string) string")
  fmt.Println("Vigenère cipher decrypt - Vigen_dec(k, m string) string")
}
