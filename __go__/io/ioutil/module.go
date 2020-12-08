package ioutil

import (
	mod "io/ioutil"
	"reflect"

	pygolin_runtime "github.com/pygolin/runtime"
)

func fun(f *pygolin_runtime.Frame, _ []*pygolin_runtime.Object) (*pygolin_runtime.Object, *pygolin_runtime.BaseException) {
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Discard)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Discard", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.NopCloser)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "NopCloser", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ReadAll)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadAll", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ReadDir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ReadFile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ReadFile", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.TempDir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TempDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.TempFile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "TempFile", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.WriteFile)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "WriteFile", o); raised != nil {
		return nil, raised
	}

	return nil, nil
}

var Code = pygolin_runtime.NewCode("<module>", "io/ioutil", nil, 0, fun)

func init() {
	pygolin_runtime.RegisterModule("__go__/io/ioutil", Code)
}
