package flexmem

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/eaciit/toolkit"
)

type RpcProxy struct {
	fns map[string]reflect.Value
	log *toolkit.LogEngine
}

func NewRpcProxy() *RpcProxy {
	r := new(RpcProxy)
	return r
}

func (r *RpcProxy) setLogger(l *toolkit.LogEngine) *RpcProxy {
	r.log = l
	return r
}

func proxyLog(p *RpcProxy) *toolkit.LogEngine {
	if p.log == nil {
		p.log = toolkit.NewLogEngine(true, false, "", "", "")
	}
	return p.log
}

func (r *RpcProxy) Call(request Request, response *Response) error {
	var err error

	name := strings.ToLower(request.Name)
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				err = fmt.Errorf("%s panic error. %v", name, rec)
			}
		}()

		if r.fns == nil {
			err = fmt.Errorf("%s error. invalid initialization", name)
			return
		}

		fn, ok := r.fns[name]
		if !ok {
			err = fmt.Errorf("service %s is not exist", name)
			return
		}

		var ins []reflect.Value
		var outs []reflect.Value
		if len(request.Parm) > 0 {
			ins = make([]reflect.Value, len(request.Parm))
			for idx, parm := range request.Parm {
				ins[idx] = reflect.ValueOf(parm)
			}
		}
		//fmt.Println("call", name, "number of parms", len(request.Parm), "raw:", ins, "json:", toolkit.JsonString(ins))
		outs = fn.Call(ins)
		response.Data = toolkit.ToBytes(outs[0].Interface(), "")
	}()

	return err
}

func RegisterRpcObjectToProxy(proxy *RpcProxy, objs ...interface{}) *RpcProxy {
	if proxy == nil {
		proxy = NewRpcProxy()
	}

	for _, obj := range objs {
		registerObjToProxy(proxy, obj)
	}

	return proxy
}

func registerObjToProxy(proxy *RpcProxy, obj interface{}) error {
	rv := reflect.ValueOf(obj)
	rt := rv.Type()

	fnCount := rv.NumMethod()
	proxyLog(proxy).Infof("registering object to proxy: %s. %d method(s) found",
		rt.Elem().Name(), fnCount)
	for i := 0; i < fnCount; i++ {
		fn := rv.Method(i)
		ft := rt.Method(i)
		fnName := strings.ToLower(fmt.Sprintf("%s.%s",
			rt.Elem().Name(), ft.Name))

		if proxy.fns == nil {
			proxy.fns = map[string]reflect.Value{}
		}

		proxy.fns[fnName] = fn
		proxyLog(proxy).Infof("adding to proxy: %s", fnName)
	}

	return nil
}
