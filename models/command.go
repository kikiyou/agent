package models

import (
	"errors"

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

func (m COMMANDS) AddCommand(COMMAND string, LABEL string) (err error) {
	_, err = db.GetDB().Exec("INSERT INTO COMMANDS(COMMAND, LABEL, ISDYNAMIC) VALUES($1, $2, 0)", COMMAND, LABEL)
	if err != nil {
		return err
	}
	return nil
}

func (m COMMANDS) DeleteCommand(ID int64) (err error) {
	_, err = m.GetCommandAndIsDynamicByID(ID)

	if err != nil {
		return errors.New("ID not found")
	}

	_, err = db.GetDB().Exec("DELETE FROM COMMANDS WHERE id=$1", ID)

	return err
}
