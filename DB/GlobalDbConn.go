package DB

import (
	"database/sql"
	"log"
)

type AllDbConn struct {
	MariaDB *sql.DB
}


var GDbConn AllDbConn


func OpenDB_Conn() error {
	log.Println("OpenDB_Conn (+)")

	var lErr error
	GDbConn.MariaDB, lErr = MakeDB_Conn(MariaDataBase)
	if lErr != nil {
		log.Println("Error in DB connect")
		return lErr
	}
	
	log.Println("OpenDB_Conn (-)")

	return nil
}

func CloseDB_Conn(){
	log.Println("CloseDB_Conn (+)")

	if GDbConn.MariaDB!=nil{
		GDbConn.MariaDB.Close()
	}
	
	log.Println("CloseDB_Conn (-)")

}