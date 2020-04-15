package dbops

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func AddUserCredential(loginName string, pwd string) error {

}

func GetUserCredential(loginName string) (string, error) {

}

/*
1. api的操作主要是对DB的CRUD
 */