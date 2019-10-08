package flexmem

import (
	"context"
	"fmt"

	"github.com/ariefdarmawan/flexmem/fmclient"

	"git.eaciitapp.com/sebar/dbflex"
)

type Connection struct {
	dbflex.ConnectionBase `bson:"-" json:"-"`
	ctx                   context.Context
	client                *fmclient.Client
}

func (c *Connection) Connect() error {
	cl, err := fmclient.NewClient(c.Host)
	if err != nil {
		return fmt.Errorf("sebarmem connect fail. %s", err.Error())
	}
	c.client = cl
	return nil
}

func (c *Connection) State() string {
	if c.client == nil {
		return dbflex.StateUnknown
	} else {
		return dbflex.StateConnected
	}
}

func (c *Connection) Close() {
	if c.client != nil {
		c.client.Close()
		c.client = nil
	}
}

func (c *Connection) NewQuery() dbflex.IQuery {
	q := new(Query)
	q.SetThis(q)
	q.SetConnection(c)

	return q
}

/*
func (c *Connection) DropTable(name string) error {
	return c.db.Collection(name).Drop(c.ctx)
}

func (c *Connection) Prepare(dbflex.ICommand) (dbflex.IQuery, error) {
	panic("not implemented")
}

func (c *Connection) Execute(dbflex.ICommand, toolkit.M) (interface{}, error) {
	panic("not implemented")
}

func (c *Connection) Cursor(dbflex.ICommand, toolkit.M) dbflex.ICursor {
	panic("not implemented")
}

func (c *Connection) NewQuery() dbflex.IQuery {
	panic("not implemented")
}

func (c *Connection) ObjectNames(dbflex.ObjTypeEnum) []string {
	panic("not implemented")
}

func (c *Connection) ValidateTable(interface{}, bool) error {
	panic("not implemented")
}

func (c *Connection) DropTable(string) error {
	panic("not implemented")
}

func (c *Connection) SetThis(dbflex.IConnection) dbflex.IConnection {
	panic("not implemented")
}

func (c *Connection) This() dbflex.IConnection {
	panic("not implemented")
}

func (c *Connection) SetFieldNameTag(string) {
	panic("not implemented")
}

func (c *Connection) FieldNameTag() string {
	panic("not implemented")
}
*/
