package kvdb

import (
	"git.eaciitapp.com/sebar/dbflex"
)

type QueryResult struct {
	Error error
}

func (k *KvDB) Query(name string, parm *dbflex.QueryParam) *QueryResult {
	res := new(QueryResult)
	return res
}
