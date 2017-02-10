package main

import (
    "fmt"
    "net/http"
	"crypto/sha1"
	"Imprinted/authService/register"
	)

func handler(w http.ResponseWriter, r *http.Request){
		if(r.URL.Query().Get("Action")=="Register"){
			if(userManagement.Load(encrypt(r.URL.Query().Get("UserName")))== nil){
			userManagement.Save(&userManagement.User{UserName: encrypt(r.URL.Query().Get("UserName")), Password: encrypt(r.URL.Query().Get("UserName"))})
			}
		}else{
			if(userManagement.Load(encrypt(r.URL.Query().Get("UserName")))!= nil){
				fmt.Fprintf(w, "true")
				}else{
				fmt.Fprintf(w, "false")
				}
		}
}
func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func encrypt(s string) []byte {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)
	return bs
}
