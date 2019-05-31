package main

import (
	"fmt"
	"os"
	"io"
	"net/http" 
)

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}
type logWriter struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

resp, err := http.Get("http://google.com")
if err != nil{
	fmt.Println("Error:", err)
	os.Exit(1)
}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	os.Exit(0)





} // main func END -----------

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Just show length of the bytes", len(bs))
	return len(bs), nil
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi, I'm English bot"
}

func (spanishBot) getGreeting() string {
	return "Hi, I'm Spanish bot"
}
