package filepath

import (
	mod "path/filepath"
	"reflect"

	pygolin_runtime "github.com/pygolin/runtime"
)

func fun(f *pygolin_runtime.Frame, _ []*pygolin_runtime.Object) (*pygolin_runtime.Object, *pygolin_runtime.BaseException) {
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Abs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Abs", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Base)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Base", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Clean)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Clean", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Dir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Dir", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ErrBadPattern)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ErrBadPattern", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.EvalSymlinks)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "EvalSymlinks", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Ext)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Ext", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.FromSlash)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "FromSlash", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Glob)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Glob", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.HasPrefix)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "HasPrefix", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.IsAbs)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "IsAbs", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Join)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Join", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.ListSeparator))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ListSeparator", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Match)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Match", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Rel)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Rel", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(uint(mod.Separator))); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Separator", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SkipDir)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SkipDir", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Split)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Split", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.SplitList)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "SplitList", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.ToSlash)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "ToSlash", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.VolumeName)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "VolumeName", o); raised != nil {
		return nil, raised
	}
	if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(mod.Walk)); raised != nil {
		return nil, raised
	} else if raised = f.Globals().SetItemString(f, "Walk", o); raised != nil {
		return nil, raised
	}
	if true {
		var x mod.WalkFunc
		if o, raised := pygolin_runtime.WrapNative(f, reflect.ValueOf(x)); raised != nil {
			return nil, raised
		} else if raised = f.Globals().SetItemString(f, "WalkFunc", o.Type().ToObject()); raised != nil {
			return nil, raised
		}
	}

	return nil, nil
}

var Code = pygolin_runtime.NewCode("<module>", "path/filepath", nil, 0, fun)

func init() {
	pygolin_runtime.RegisterModule("__go__/path/filepath", Code)
}
