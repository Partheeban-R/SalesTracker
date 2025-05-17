package DB

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Structure to hold database connection details
type DatabaseType struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
	DBType   string
	DB       string
}

// structure to hold all db connection details used in this program
type AllUsedDatabases struct {
	MariaDB       DatabaseType
}

// ---------------------------------------------------------------------------------
// function opens the db connection and return connection variable
// ---------------------------------------------------------------------------------
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
log.Println("conn :",lConnString)
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

// --------------------------------------------------------------------
//
//	execute bulk inserts
//
// --------------------------------------------------------------------
func ExecuteBulkStatement(db *sql.DB, sqlStringValues string, sqlString string) error {
	log.Println("ExecuteBulkStatement+")
	//trim the last ,
	sqlStringValues = sqlStringValues[0 : len(sqlStringValues)-1]
	_, err := db.Exec(sqlString + sqlStringValues)
	if err != nil {
		log.Println(err)
		log.Println("ExecuteBulkStatement-")
		return err
	} else {
		log.Println("inserted Sucessfully")
	}
	log.Println("ExecuteBulkStatement-")
	return nil
}
