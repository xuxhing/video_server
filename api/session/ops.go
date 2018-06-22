package session

import (
	"time"
	"sync"
)

type SimpleSession struct {
	Username string
	TTL int64
}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB() {

}

func GenerateNewSessionId(un string) string {

}

func IsSessionExpired(sid string) (string, bool) {
	
}