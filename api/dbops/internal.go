package dbops

import (
	"strconv"
	"sync"
	"video_server/defs"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr:=strconv.FormatInt(ttl, 10)
	stmtIns, err:=dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?,?,?)")
	if err!=nil{
		return err
	}
	err = stmtIns.Exec(sid, ttlstr, uname)
	defer stmtIns.Close()
	return nil
}