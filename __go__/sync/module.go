package sync

import (
	"reflect"
	mod "sync"

	pygolin_runtime "github.com/pygolin/runtime"
)

func fun(f *pygolin_runtime.Frame, _ []*pygolin_runtime.Object) (*pygolin_runtime.Object, *pygolin_runtime.BaseException) {
	if true {
		var x mod.Cond
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Cond", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Map
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Map", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Mutex
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Mutex", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.NewCond)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NewCond", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.Once
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Once", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.Pool
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "Pool", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.RWMutex
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "RWMutex", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}
	if true {
		var x mod.WaitGroup
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "WaitGroup", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}

	return nil, nil
}

var Code = pygolin_runtime.NewCode("<module>", "sync", nil, 0, fun)

func init() {
	pygolin_runtime.RegisterModule("__go__/sync", Code)
}
