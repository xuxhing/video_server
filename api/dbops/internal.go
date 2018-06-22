package dbops

import (
	"strconv"
	"database/sql"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr:=strconv.FormatInt(ttl, 10)
	stmtIns, err:=dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?,?,?)")
	if err!=nil{
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss:=&defs.SimpleSession{}
	stmtOut, err:=dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id=?")
	if err!=nil {
		return nil, err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err!=nil && err!=sql.ErrNoRows {
		return nil, err
	}
	if res, err:=strconv.ParseInt(ttl, 10, 64);err==nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	return nil, nil
}