package util

import (
	"math/rand"
	"time"
)

type Result struct {
	Code int         `json:"code"` // return code
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}


// NewResult creates a result with Code=0, Msg="", Data=nil.
func  NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}


func RandString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Nanosecond)

	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}