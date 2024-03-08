package orm

import (
	"github.com/authink/orm/model"
	"github.com/authink/orm/sql"
	"github.com/jmoiron/sqlx"
)

type Inserter[M model.Identifier] interface {
	Insert(M) error
	InsertTx(*sqlx.Tx, M) error
}

type Saver[M model.Identifier] interface {
	Save(M) error
	SaveTx(*sqlx.Tx, M) error
}

type Updater[M model.Identifier] interface {
	Update(M) error
	UpdateTx(*sqlx.Tx, M) error
}

type Geter[M model.Identifier] interface {
	Get(M) error
	GetTx(*sqlx.Tx, M) error
}

type Deleter[M model.Identifier] interface {
	Delete(M) error
	DeleteTx(*sqlx.Tx, M) error
}

type Finder[M model.Identifier] interface {
	Find(...model.Arg) ([]M, error)
}

type Counter interface {
	Count(...model.Arg) (int, error)
	CountTx(*sqlx.Tx, ...model.Arg) (int, error)
}

type Pager[M model.Identifier] interface {
	PaginationTx(*sqlx.Tx, model.Pager) ([]M, error)
}

type ORM[M model.Identifier] interface {
	Inserter[M]
	Saver[M]
	Updater[M]
	Geter[M]
	Deleter[M]
	Finder[M]
	Counter
	Pager[M]
}

type ORMBase[M model.Identifier, S sql.SQL] struct {
	Executor
	Stmt S
}

func (o *ORMBase[M, S]) Insert(m M) error {
	return NamedInsert(o.Executor, o.Stmt.Insert(), m)
}

func (o *ORMBase[M, S]) InsertTx(tx *sqlx.Tx, m M) error {
	return NamedInsert(tx, o.Stmt.Insert(), m)
}

func (o *ORMBase[M, S]) Save(m M) error {
	return NamedSave(o.Executor, o.Stmt.Save(), m)
}

func (o *ORMBase[M, S]) SaveTx(tx *sqlx.Tx, m M) error {
	return NamedSave(tx, o.Stmt.Save(), m)
}

func (o *ORMBase[M, S]) Update(m M) error {
	return NamedUpdate(o.Executor, o.Stmt.Update(), m)
}

func (o *ORMBase[M, S]) UpdateTx(tx *sqlx.Tx, m M) error {
	return NamedUpdate(tx, o.Stmt.Update(), m)
}

func (o *ORMBase[M, S]) Get(m M) error {
	return Get(o.Executor, o.Stmt.Get(), m)
}

func (o *ORMBase[M, S]) GetTx(tx *sqlx.Tx, m M) error {
	return Get(tx, o.Stmt.GetForUpdate(), m)
}

func (o *ORMBase[M, S]) Delete(m M) error {
	return Delete(o.Executor, o.Stmt.Delete(), m)
}

func (o *ORMBase[M, S]) DeleteTx(tx *sqlx.Tx, m M) error {
	return Delete(tx, o.Stmt.Delete(), m)
}

func (o *ORMBase[M, S]) Find(args ...model.Arg) (list []M, err error) {
	err = Select(o.Executor, o.Stmt.Find(), &list, o.parseArg(args))
	return
}

func (o *ORMBase[M, S]) Count(args ...model.Arg) (c int, err error) {
	err = Count(o.Executor, o.Stmt.Count(), &c, o.parseArg(args))
	return
}

func (o *ORMBase[M, S]) CountTx(tx *sqlx.Tx, args ...model.Arg) (c int, err error) {
	err = Count(tx, o.Stmt.Count(), &c, o.parseArg(args))
	return
}

func (o *ORMBase[M, S]) PaginationTx(tx *sqlx.Tx, pager model.Pager) (list []M, err error) {
	err = Select(tx, o.Stmt.Pagination(), &list, pager)
	return
}

func (o *ORMBase[M, S]) parseArg(args []model.Arg) model.Arg {
	var arg model.Arg
	if len(args) > 0 {
		arg = args[0]
	} else {
		arg = &model.Argument{}
	}
	return arg
}
