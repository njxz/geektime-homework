package geektime

import (
	"database/sql"
	"errors"
	"fmt"
)

func SQLinsert(sql_action string) (string, error) {
	return "", sql.ErrNoRows
}

func CallSQLinsert() error {
	command := ""
	_, err := SQLinsert(command)
	return fmt.Errorf("%w,%s", err, "insert '"+command+"' error")
}

func handle() {
	err := CallSQLinsert()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	handle()
}
