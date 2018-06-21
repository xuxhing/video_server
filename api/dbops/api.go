package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name, pwd) values(?,?)")
	if err!=nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}	

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err!=nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err:=dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err!=nil {
		log.Printf("Delete user error: %s", err)
	}
	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}