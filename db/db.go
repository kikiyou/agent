package db

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/go-gorp/gorp"
	"github.com/kikiyou/agent/g"
	_ "github.com/mattn/go-sqlite3"
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorp.DbMap

// var commandModel = new(models.COMMANDS)

//Init ...
func Init() {
	dbinfo := filepath.Join(g.Root, "db/command_set.sqlite3")
	fmt.Println("######dbinfo##############")

	fmt.Println(dbinfo)
	fmt.Println("######dbinfo##############")
	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests

	if err != nil {
		return nil, err
	}
	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
