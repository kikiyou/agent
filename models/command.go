package models

import (
	"github.com/kikiyou/agent/db"
)

type COMMANDS struct {
	// db tag lets you specify the column name if it differs from the struct field
	ID        int64  `db:"ID"`
	COMMAND   string `db:"COMMAND,size:255"`
	LABEL     string `db:"LABEL,size:35"`
	ISDYNAMIC int64  `db:"ISDYNAMIC"`
}

// dbmap.AddTableWithName(models.COMMANDS{}, "COMMANDS").SetKeys(true, "ID")
// err = dbmap.CreateTablesIfNotExists()
//One ...
func (m COMMANDS) GetCommandAndIsDynamicByID(ID int64) (c COMMANDS, err error) {

	err = db.GetDB().SelectOne(&c, "SELECT COMMAND,ISDYNAMIC FROM COMMANDS where ID=$1", ID)
	return c, err
}

func (m COMMANDS) GetCommandList() (c []COMMANDS, err error) {
	_, err = db.GetDB().Select(&c, "SELECT * from COMMANDS")
	return c, err
}
