package gosqlcsv

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/thomas-bamilo/dbconf"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	// SQL Server driver
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/sqltocsv"
)

type dbStruct struct {
	ConnStr string
	DbType  string
}

func (f *dbStruct) SetConnStr(name string) { f.ConnStr = name }
func (f *dbStruct) SetDbType(name string)  { f.DbType = name }

// GoSQLCsv function fetches information from csv_name.txt, db_config.yaml and sql_query.txt to query database and print result to csv
func GoSQLCsv() {

	var dbConf dbconf.DbConf
	dbConf.ReadYamlDbConf()

	// create connection string depending on dbChoice
	dbStruct1 := dbStruct{}

	if dbConf.DbChoice == "sc" {
		dbStruct1.SetDbType("mysql")
		dbStruct1.SetConnStr(dbConf.ScUser + ":" + dbConf.ScPw + "@tcp(" + dbConf.ScHost + ")/" + dbConf.ScDb)

	} else if dbConf.DbChoice == "oms" {
		dbStruct1.SetDbType("mysql")
		dbStruct1.SetConnStr(dbConf.OmsUser + ":" + dbConf.OmsPw + "@tcp(" + dbConf.OmsHost + ")/" + dbConf.OmsDb)

	} else if dbConf.DbChoice == "bi" {
		dbStruct1.SetDbType("sqlserver")
		dbStruct1.SetConnStr(`sqlserver://` + dbConf.BiUser + ":" + dbConf.BiPw + "@" + dbConf.BiHost + "/" + dbConf.BiDb)

	}

	// open database connection
	db, err := sql.Open(dbStruct1.DbType, dbStruct1.ConnStr)
	if err != nil {

		log.Fatal(err)
	}

	// test database connection
	err = db.Ping()
	if err != nil {
		log.Println("Connection failed")
		log.Fatal(err)
	} else {
		log.Println("Connection successful!")
	}

	// read query from sql_query.txt in the same folder as the executable file
	query, err := ioutil.ReadFile("sql_query.txt")
	if err != nil {
		log.Fatal(err)
	}
	queryStr := string(query) // convert content to a 'string'

	// read csv name from csv_name.txt
	csvName, err := ioutil.ReadFile("csv_name.txt")
	if err != nil {
		log.Fatal(err)
	}
	csvNameStr := string(csvName) // convert content to a 'string'

	rows, _ := db.Query(queryStr)

	err = sqltocsv.WriteFile(csvNameStr, rows)
	if err != nil {
		log.Fatal(err)
	}

}
