package main

import (
	"io"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"video_server/api/defs"
	"video_server/api/session"
	"video_server/api/dbops"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 首先从request里面读出create user相关的body
	res, _:=ioutil.ReadAll(r.Body)
	ubody:=&defs.UserCredential{}

	if err:=json.Unmarshal(res, ubody);err!=nil{
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err:=dbops.AddUserCredential(ubody.Username, ubody.Pwd); err!=nil {
		sendErrorResponse(w, defs.ErrorDBError)
	}

	id:=session.GenerateNewSessionId(ubody.Username)
	su:=&defs.SignedUp{Success: true, SessionId: id}

	if resp, err:=json.Marshal(su);err!=nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}