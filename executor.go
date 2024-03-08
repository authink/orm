package orm

import (
	"database/sql"
	"errors"

	"github.com/authink/orm/model"
	"github.com/jmoiron/sqlx"
)

type Executor interface {
	NamedExec(string, any) (sql.Result, error)
	PrepareNamed(string) (*sqlx.NamedStmt, error)
}

func NamedInsert(executor Executor, statement string, m model.Identifier) error {
	return namedExec(executor, statement, m, afterInsert)
}

func NamedUpdate(executor Executor, statement string, m model.Identifier) error {
	return namedExec(executor, statement, m, afterUpdate)
}

func NamedSave(executor Executor, statement string, m model.Identifier) error {
	return namedExec(executor, statement, m, afterSave)
}

func Get(executor Executor, statement string, m model.Identifier) error {
	return prepareNamed(executor, statement, m, m, false)
}

func Delete(executor Executor, statement string, m model.Identifier) (err error) {
	return namedExec(executor, statement, m, afterDelete)
}

func Count(executor Executor, statement string, c *int, arg model.Arg) error {
	return prepareNamed(executor, statement, c, arg, false)
}

func Select(executor Executor, statement string, list any, arg model.Arg) error {
	return prepareNamed(executor, statement, list, arg, true)
}

func namedExec(executor Executor, statement string, m model.Identifier, afterExec func(sql.Result, model.Identifier) error) (err error) {
	result, err := executor.NamedExec(
		statement,
		m,
	)
	if err != nil {
		return
	}

	err = afterExec(result, m)
	return
}

func prepareNamed(executor Executor, statement string, dest, arg any, isSelect bool) (err error) {
	stmt, err := executor.PrepareNamed(statement)
	if err != nil {
		return
	}
	if isSelect {
		err = stmt.Select(dest, arg)
	} else {
		err = stmt.Get(dest, arg)
	}
	return
}

func afterInsert(result sql.Result, m model.Identifier) (err error) {
	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}

	m.SetId(uint32(lastId))
	return
}

func afterSave(result sql.Result, m model.Identifier) (err error) {
	if err = afterInsert(result, m); err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	} else if rowsAffected == 0 {
		err = errors.New("duplicate key")
	}
	return
}

func afterUpdate(sql.Result, model.Identifier) error {
	return nil
}

func afterDelete(sql.Result, model.Identifier) error {
	return nil
}
