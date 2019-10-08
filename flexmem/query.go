package flexmem

import (
	"fmt"

	"git.eaciitapp.com/sebar/dbflex"
	df "git.eaciitapp.com/sebar/dbflex"
	"github.com/eaciit/toolkit"
	. "github.com/eaciit/toolkit"
)

type Query struct {
	dbflex.QueryBase
}

func (q *Query) BuildCommand() (interface{}, error) {
	return nil, nil
}

func (q *Query) BuildFilter(f *df.Filter) (interface{}, error) {
	return f, nil
}

func (q *Query) Cursor(m M) df.ICursor {
	cursor := new(Cursor)
	cursor.SetThis(cursor)
	conn := q.Connection().(*Connection)

	tablename := q.Config(df.ConfigKeyTableName, "").(string)

	parts := q.Config(df.ConfigKeyGroupedQueryItems, df.GroupedQueryItems{}).(df.GroupedQueryItems)
	where := q.Config(df.ConfigKeyWhere, M{}).(M)

	aggrs, hasAggr := parts[df.QueryAggr]
	groupby, hasGroup := parts[df.QueryGroup]

	queryParm := toolkit.M{}
	queryParm.Set(df.ConfigKeyWhere, where)
	if hasAggr {
		queryParm.Set(df.QueryAggr, aggrs)
	}
	if hasGroup {
		queryParm.Set(df.QueryGroup, groupby)
	}

	if res := conn.client.Call("KvDB.Query", tablename, queryParm); res.Err() != nil {
		cursor.SetError(fmt.Errorf("cursor error: %s", res.Err()))
		return cursor
	} else {
		ms := []toolkit.M{}
		toolkit.FromBytes(res.Data, "json", ms)
		cursor.data = ms
	}

	return cursor
}

func (q *Query) Execute(m M) (interface{}, error) {
	tablename := q.Config(df.ConfigKeyTableName, "").(string)
	conn := q.Connection().(*Connection)
	data := m.Get("data")

	//parts := q.Config(df.ConfigKeyGroupedQueryItems, df.GroupedQueryItems{}).(df.GroupedQueryItems)
	where := q.Config(df.ConfigKeyWhere, new(dbflex.Filter)).(*dbflex.Filter)
	//hasWhere := where != nil

	ct := q.Config(df.ConfigKeyCommandType, "N/A")
	switch ct {
	case df.QueryInsert:
		res := conn.client.Call("KvDB.insert", tablename, data)
		return res.Data, res.Err()

	case df.QueryUpdate:
		res := conn.client.Call("KvDB.update", tablename, where, data)
		return res.Data, res.Err()

	case df.QueryDelete:
		res := conn.client.Call("KvDB.delete", tablename, where, data)
		return res.Data, res.Err()

	case df.QuerySave:
		res := conn.client.Call("KvDB.save", tablename, where, data)
		return res.Data, res.Err()
	}

	return nil, nil
}

/*
func (q *Query) SetThis(q dbflex.IQuery) {
	panic("not implemented")
}

func (q *Query) This() dbflex.IQuery {
	panic("not implemented")
}

func (q *Query) BuildFilter(*dbflex.Filter) (interface{}, error) {
	panic("not implemented")
}

func (q *Query) BuildCommand() (interface{}, error) {
	panic("not implemented")
}

func (q *Query) Cursor(toolkit.M) dbflex.ICursor {
	panic("not implemented")
}

func (q *Query) Execute(toolkit.M) (interface{}, error) {
	panic("not implemented")
}

func (q *Query) SetConfig(string, interface{}) {
	panic("not implemented")
}

func (q *Query) SetConfigM(toolkit.M) {
	panic("not implemented")
}

func (q *Query) Config(string, interface{}) interface{} {
	panic("not implemented")
}

func (q *Query) ConfigRef(string, interface{}, interface{}) {
	panic("not implemented")
}

func (q *Query) DeleteConfig(...string) {
	panic("not implemented")
}

func (q *Query) Connection() dbflex.IConnection {
	panic("not implemented")
}

func (q *Query) SetConnection(dbflex.IConnection) {
	panic("not implemented")
}
*/
