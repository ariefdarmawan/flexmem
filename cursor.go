package flexmem

import (
	"fmt"

	"git.eaciitapp.com/sebar/dbflex"
	"github.com/eaciit/toolkit"
)

type Cursor struct {
	dbflex.CursorBase

	tablename  string
	data       []toolkit.M
	currentIdx int
}

func (cr *Cursor) Close() {
	cr.currentIdx = 0
	cr.data = make([]toolkit.M, 0)
}

func (cr *Cursor) Count() int {
	return len(cr.data)
}

func (cr *Cursor) Fetch(out interface{}) error {
	if cr.Error() != nil {
		return toolkit.Errorf("unable to fetch data. %s", cr.Error())
	}

	if cr.currentIdx >= cr.Count() {

	}

	m := cr.data[cr.currentIdx]
	if err := toolkit.Serde(m, out, ""); err != nil {
		return fmt.Errorf("cursor unable to cast. %s", err.Error())
	}

	cr.currentIdx++
	return nil
}

func (cr *Cursor) Fetchs(result interface{}, n int) error {
	if cr.Error() != nil {
		return toolkit.Errorf("unable to fetch data: %s", cr.Error())
	}

	read := 0
	ms := make([]toolkit.M, n)
	count := cr.Count()
	for {
		if cr.currentIdx >= count {
			break
		}

		ms[read] = cr.data[cr.currentIdx]

		read++
		cr.currentIdx++
		if n != 0 && read == n {
			break
		}
	}

	if read < n {
		ms = ms[0:read]
	}

	if err := toolkit.Serde(ms, result, ""); err != nil {
		return fmt.Errorf("cursor unable to cast: %s", err.Error())
	}

	return nil
}

/*
func (cr *Cursor) Reset() error {
	panic("not implemented")
}

func (cr *Cursor) Fetch(interface{}) error {
	panic("not implemented")
}

func (cr *Cursor) Fetchs(interface{}, int) error {
	panic("not implemented")
}


func (cr *Cursor) CountAsync() <-chan int {
	panic("not implemented")
}

func (cr *Cursor) Error() error {
	panic("not implemented")
}

func (cr *Cursor) CloseAfterFetch() bool {
	panic("not implemented")
}

func (cr *Cursor) SetCountCommand(dbflex.ICommand) {
	panic("not implemented")
}

func (cr *Cursor) CountCommand() dbflex.ICommand {
	panic("not implemented")
}

func (cr *Cursor) Connection() dbflex.IConnection {
	panic("not implemented")
}

func (cr *Cursor) SetConnection(dbflex.IConnection) {
	panic("not implemented")
}

func (cr *Cursor) ConfigRef(key string, def interface{}, out interface{}) {
	panic("not implemented")
}

func (cr *Cursor) Set(key string, value interface{}) {
	panic("not implemented")
}

func (cr *Cursor) SetCloseAfterFetch() dbflex.ICursor {
	panic("not implemented")
}

func (cr *Cursor) AutoClose(time.Duration) dbflex.ICursor {
	panic("not implemented")
}
*/
