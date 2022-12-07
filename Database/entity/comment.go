package entity

import "database/sql"

type Comment struct {
	Id      int32
	Email   string
	Comment sql.NullString
}
