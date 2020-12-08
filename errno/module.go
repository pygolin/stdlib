package errno
import (
	πg "github.com/pygolin/runtime"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/errno.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßEINVAL := πg.InternStr("EINVAL")
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: EINVAL = 22
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ßEINVAL.ToObject(), πg.NewInt(22).ToObject()); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("errno", Code)
}

