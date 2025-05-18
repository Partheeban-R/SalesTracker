package DB

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


type DatabaseType struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
	DBType   string
	DB       string
}


type AllUsedDatabases struct {
	MariaDB       DatabaseType
}

func MakeDB_Conn(pDBtype string) (*sql.DB, error) {
	lDbDetails := new(AllUsedDatabases)
	lDbDetails.Init()

	//log.Println(DbDetails)

	lConnString := ""
	lLocalDBtype := ""

	var lDb *sql.DB
	var lErr error
	var lDataBaseConnection DatabaseType
	// get connection details
	if pDBtype == lDbDetails.MariaDB.DB {
		lDataBaseConnection = lDbDetails.MariaDB
		lLocalDBtype = lDbDetails.MariaDB.DBType
	}
	// Prepare connection string
	if lLocalDBtype == "mysql" {
		lConnString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", lDataBaseConnection.User, lDataBaseConnection.Password, lDataBaseConnection.Server, lDataBaseConnection.Port, lDataBaseConnection.Database)
	}

	//make a connection to db
	if lLocalDBtype != "" {
		lDb, lErr = sql.Open(lLocalDBtype, lConnString)
		if lErr != nil {
			log.Println("Open connection failed:", lErr.Error())
		} 
	} else {
		return lDb, fmt.Errorf(" Invalid DB Details")
	}
	return lDb, lErr
}


func InsertBulkData(pDB *sql.DB, pSqlStringValues string, pSqlString string) error {
	log.Println("InsertBulkData (+)")
	pSqlStringValues = pSqlStringValues[0 : len(pSqlStringValues)-1]
	_, lErr := pDB.Exec(pSqlString + pSqlStringValues)
	log.Println("pSqlString + pSqlStringValues :\n",pSqlString + pSqlStringValues)
	if lErr != nil {
		log.Println(lErr)
		return lErr
	}
	log.Println("InsertBulkData (-)")
	return nil
}
