package DB

import (
	"fmt"
	"SalesTracker/common"
	"strconv"
)

const (
	MariaDataBase     = "MariaDataBase"
)

// Initializing DB Details
func (d *AllUsedDatabases) Init() {
	dbconfig := common.ReadTomlConfig("../dbconfig.toml")

	//setting Maria db connection details
	d.MariaDB.Server = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBServer"])
	d.MariaDB.Port, _ = strconv.Atoi(fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBPort"]))
	d.MariaDB.User = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBUser"])
	d.MariaDB.Password = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBPassword"])
	d.MariaDB.Database = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBDatabase"])
	d.MariaDB.DBType = fmt.Sprintf("%v", dbconfig.(map[string]interface{})["MariaDBDBType"])
	d.MariaDB.DB = MariaDataBase
	
}
