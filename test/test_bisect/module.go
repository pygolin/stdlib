package test_bisect
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/UserList"
	// _ "github.com/pygolin/stdlib/test/test_support"
	// _ "github.com/pygolin/stdlib/bisect"
	// _ "github.com/pygolin/stdlib/unittest"
	// _ "github.com/pygolin/stdlib/random"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/test/test_bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß0123456789 := πg.InternStr("0123456789")
		ß02468 := πg.InternStr("02468")
		ßAttributeError := πg.InternStr("AttributeError")
		ßCmpErr := πg.InternStr("CmpErr")
		ßGetOnly := πg.InternStr("GetOnly")
		ßIndexError := πg.InternStr("IndexError")
		ßLenOnly := πg.InternStr("LenOnly")
		ßNone := πg.InternStr("None")
		ßOverflowError := πg.InternStr("OverflowError")
		ßRange := πg.InternStr("Range")
		ßTestBisect := πg.InternStr("TestBisect")
		ßTestCase := πg.InternStr("TestCase")
		ßTestErrorHandling := πg.InternStr("TestErrorHandling")
		ßTestInsort := πg.InternStr("TestInsort")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUserList := πg.InternStr("UserList")
		ßValueError := πg.InternStr("ValueError")
		ßZeroDivisionError := πg.InternStr("ZeroDivisionError")
		ß_UserList := πg.InternStr("_UserList")
		ß__cmp__ := πg.InternStr("__cmp__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__init__ := πg.InternStr("__init__")
		ß__len__ := πg.InternStr("__len__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__test__ := πg.InternStr("__test__")
		ß_bisect := πg.InternStr("_bisect")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßbisect := πg.InternStr("bisect")
		ßbisect_left := πg.InternStr("bisect_left")
		ßbisect_right := πg.InternStr("bisect_right")
		ßchoice := πg.InternStr("choice")
		ßdata := πg.InternStr("data")
		ßgettotalrefcount := πg.InternStr("gettotalrefcount")
		ßhasattr := πg.InternStr("hasattr")
		ßinsert := πg.InternStr("insert")
		ßinsort := πg.InternStr("insort")
		ßinsort_left := πg.InternStr("insort_left")
		ßinsort_right := πg.InternStr("insort_right")
		ßlast_insert := πg.InternStr("last_insert")
		ßlen := πg.InternStr("len")
		ßlibreftest := πg.InternStr("libreftest")
		ßlist := πg.InternStr("list")
		ßmax := πg.InternStr("max")
		ßmaxsize := πg.InternStr("maxsize")
		ßmin := πg.InternStr("min")
		ßmodule := πg.InternStr("module")
		ßmodules := πg.InternStr("modules")
		ßobject := πg.InternStr("object")
		ßprecomputedCases := πg.InternStr("precomputedCases")
		ßpy_bisect := πg.InternStr("py_bisect")
		ßrandrange := πg.InternStr("randrange")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsetUp := πg.InternStr("setUp")
		ßskipTest := πg.InternStr("skipTest")
		ßsort := πg.InternStr("sort")
		ßsorted := πg.InternStr("sorted")
		ßstart := πg.InternStr("start")
		ßstop := πg.InternStr("stop")
		ßsys := πg.InternStr("sys")
		ßtest_arg_parsing := πg.InternStr("test_arg_parsing")
		ßtest_backcompatibility := πg.InternStr("test_backcompatibility")
		ßtest_cmp_err := πg.InternStr("test_cmp_err")
		ßtest_get_only := πg.InternStr("test_get_only")
		ßtest_keyword_args := πg.InternStr("test_keyword_args")
		ßtest_large_pyrange := πg.InternStr("test_large_pyrange")
		ßtest_large_range := πg.InternStr("test_large_range")
		ßtest_len_only := πg.InternStr("test_len_only")
		ßtest_listDerived := πg.InternStr("test_listDerived")
		ßtest_main := πg.InternStr("test_main")
		ßtest_negative_lo := πg.InternStr("test_negative_lo")
		ßtest_non_sequence := πg.InternStr("test_non_sequence")
		ßtest_optionalSlicing := πg.InternStr("test_optionalSlicing")
		ßtest_precomputed := πg.InternStr("test_precomputed")
		ßtest_random := πg.InternStr("test_random")
		ßtest_support := πg.InternStr("test_support")
		ßtest_vsBuiltinSort := πg.InternStr("test_vsBuiltinSort")
		ßunittest := πg.InternStr("unittest")
		ßxrange := πg.InternStr("xrange")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 *πg.Dict
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 []πg.Param
		_ = πTemp008
		var πTemp009 πg.KWArgs
		_ = πTemp009
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: import sys
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 2: import unittest
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: from test import test_support
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import UserList as _UserList
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "UserList"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_UserList.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: UserList = _UserList.UserList
			πF.SetLineno(6)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_UserList); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßUserList, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUserList.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 12: sys.modules['_bisect'] = 0
			πF.SetLineno(12)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmodules, nil); πE != nil {
				continue
			}
			πTemp003 = ß_bisect.ToObject()
			if πE = πg.SetItem(πF, πTemp004, πTemp003, πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmodules, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πg.Contains(πF, πTemp004, ßbisect.ToObject()); πE != nil {
				continue
			}
			πTemp001 = πg.GetBool(πTemp005).ToObject()
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label1
			}
			goto Label2
			// line 14: if 'bisect' in sys.modules:
			πF.SetLineno(14)
		Label1:
			// line 15: del sys.modules['bisect']
			πF.SetLineno(15)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmodules, nil); πE != nil {
				continue
			}
			πTemp001 = ßbisect.ToObject()
			if πE = πg.DelItem(πF, πTemp003, πTemp001); πE != nil {
				continue
			}
			goto Label2
		Label2:
			// line 18: import bisect as py_bisect
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "bisect"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßpy_bisect.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 21: del sys.modules['_bisect']
			πF.SetLineno(21)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmodules, nil); πE != nil {
				continue
			}
			πTemp001 = ß_bisect.ToObject()
			if πE = πg.DelItem(πF, πTemp003, πTemp001); πE != nil {
				continue
			}
			// line 22: del sys.modules['bisect']
			πF.SetLineno(22)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmodules, nil); πE != nil {
				continue
			}
			πTemp001 = ßbisect.ToObject()
			if πE = πg.DelItem(πF, πTemp003, πTemp001); πE != nil {
				continue
			}
			// line 28: class Range(object):
			πF.SetLineno(28)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Range", "/usr/lib/python2.7/test/test_bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 29: """A trivial xrange()-like object without any integer width limitations."""
					πF.SetLineno(29)
					// line 29: """A trivial xrange()-like object without any integer width limitations."""
					πF.SetLineno(29)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A trivial xrange()-like object without any integer width limitations.").ToObject()); πE != nil {
						continue
					}
					// line 30: def __init__(self, start, stop):
					πF.SetLineno(30)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "start", Def: nil}
					πTemp002[2] = πg.Param{Name: "stop", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstart *πg.Object = πArgs[1]; _ = µstart
						var µstop *πg.Object = πArgs[2]; _ = µstop
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 31: self.start = start
							πF.SetLineno(31)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstart); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstart, πTemp001); πE != nil {
								continue
							}
							// line 32: self.stop = stop
							πF.SetLineno(32)
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstop); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstop, πTemp001); πE != nil {
								continue
							}
							// line 33: self.last_insert = None
							πF.SetLineno(33)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlast_insert, πTemp002); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 35: def __len__(self):
					πF.SetLineno(35)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 36: return self.stop - self.start
							πF.SetLineno(36)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstop, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstart, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 38: def __getitem__(self, idx):
					πF.SetLineno(38)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "idx", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µidx *πg.Object = πArgs[1]; _ = µidx
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 39: n = self.stop - self.start
							πF.SetLineno(39)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstop, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstart, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							µn = πTemp001
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µidx, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 40: if idx < 0:
							πF.SetLineno(40)
						Label1:
							// line 41: idx += n
							πF.SetLineno(41)
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µidx, µn); πE != nil {
								continue
							}
							µidx = πTemp001
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µidx, µn); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 42: if idx >= n:
							πF.SetLineno(42)
						Label3:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							πTemp005[0] = µidx
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 43: raise IndexError(idx)
							πF.SetLineno(43)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label4
						Label4:
							// line 44: return self.start + idx
							πF.SetLineno(44)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µidx); πE != nil {
								continue
							}
							πR = πTemp001
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 46: def insert(self, idx, item):
					πF.SetLineno(46)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "idx", Def: nil}
					πTemp002[2] = πg.Param{Name: "item", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("insert", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µidx *πg.Object = πArgs[1]; _ = µidx
						var µitem *πg.Object = πArgs[2]; _ = µitem
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 47: self.last_insert = idx, item
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µidx, µitem).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlast_insert, πTemp002); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßinsert.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Range").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRange.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 50: class TestBisect(unittest.TestCase):
			πF.SetLineno(50)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestBisect", "/usr/lib/python2.7/test/test_bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 52: module = py_bisect
					πF.SetLineno(52)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpy_bisect); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßmodule.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 54: def setUp(self):
					πF.SetLineno(54)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("setUp", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 55: self.precomputedCases = [
							πF.SetLineno(55)
							πTemp001 = make([]*πg.Object, 78)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[6] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[7] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[8] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[9] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[10] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(4).ToObject()).ToObject()
							πTemp001[11] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(4).ToObject()).ToObject()
							πTemp001[12] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[13] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[14] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(1.5).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[15] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[16] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(3).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[17] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[18] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[19] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(1.5).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[20] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(4).ToObject()).ToObject()
							πTemp001[21] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(3).ToObject(), πg.NewInt(4).ToObject()).ToObject()
							πTemp001[22] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[23] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[24] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(1.5).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[25] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[26] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(2.5).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[27] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[28] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(4).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[29] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[30] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[31] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(1.5).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[32] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[33] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(2.5).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[34] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(3).ToObject(), πg.NewInt(6).ToObject()).ToObject()
							πTemp001[35] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(3.5).ToObject(), πg.NewInt(6).ToObject()).ToObject()
							πTemp001[36] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(4).ToObject(), πg.NewInt(10).ToObject()).ToObject()
							πTemp001[37] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(5).ToObject(), πg.NewInt(10).ToObject()).ToObject()
							πTemp001[38] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[39] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[40] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[41] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[42] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[43] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[44] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[45] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[46] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[47] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[48] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[49] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[50] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(4).ToObject()).ToObject()
							πTemp001[51] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[52] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[53] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(1.5).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[54] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[55] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(3).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[56] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[57] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[58] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(1.5).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[59] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[60] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(3).ToObject(), πg.NewInt(4).ToObject()).ToObject()
							πTemp001[61] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[62] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[63] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(1.5).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[64] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[65] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(2.5).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[66] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(3).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[67] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(4).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[68] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[69] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πTemp001[70] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(1.5).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[71] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(2).ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[72] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(2.5).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[73] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[74] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewFloat(3.5).ToObject(), πg.NewInt(6).ToObject()).ToObject()
							πTemp001[75] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(4).ToObject(), πg.NewInt(6).ToObject()).ToObject()
							πTemp001[76] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 10)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(3).ToObject()
							πTemp005[5] = πg.NewInt(3).ToObject()
							πTemp005[6] = πg.NewInt(4).ToObject()
							πTemp005[7] = πg.NewInt(4).ToObject()
							πTemp005[8] = πg.NewInt(4).ToObject()
							πTemp005[9] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002 = πg.NewTuple4(πTemp004, πTemp003, πg.NewInt(5).ToObject(), πg.NewInt(10).ToObject()).ToObject()
							πTemp001[77] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßprecomputedCases, πTemp003); πE != nil {
								continue
							}
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsetUp.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 137: def test_precomputed(self):
					πF.SetLineno(137)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_precomputed", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var µexpected *πg.Object = πg.UnboundLocal; _ = µexpected
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßprecomputedCases, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp003 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
									continue
								}
								µfunc = πTemp005
								µdata = πTemp006
								µelem = πTemp007
								µexpected = πTemp008
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 139: self.assertEqual(func(data, elem), expected)
							πF.SetLineno(139)
							πTemp009 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp010[0] = µdata
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp010[1] = µelem
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp002, πE = µfunc.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							πTemp009[1] = µexpected
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							// line 140: self.assertEqual(func(UserList(data), elem), expected)
							πF.SetLineno(140)
							πTemp009 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(2)
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp011[0] = µdata
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUserList); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp010[0] = πTemp005
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp010[1] = µelem
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp002, πE = µfunc.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							πTemp009[1] = µexpected
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_precomputed.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 142: def test_negative_lo(self):
					πF.SetLineno(142)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_negative_lo", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmod *πg.Object = πg.UnboundLocal; _ = µmod
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 144: mod = self.module
							πF.SetLineno(144)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							µmod = πTemp001
							// line 145: self.assertRaises(ValueError, mod.bisect_left, [1, 2, 3], 5, -1, 3),
							πF.SetLineno(145)
							πTemp002 = πF.MakeArgs(6)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							πTemp003 = make([]*πg.Object, 3)
							πTemp003[0] = πg.NewInt(1).ToObject()
							πTemp003[1] = πg.NewInt(2).ToObject()
							πTemp003[2] = πg.NewInt(3).ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							πTemp002[2] = πTemp001
							πTemp002[3] = πg.NewInt(5).ToObject()
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[4] = πTemp001
							πTemp002[5] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 146: self.assertRaises(ValueError, mod.bisect_right, [1, 2, 3], 5, -1, 3),
							πF.SetLineno(146)
							πTemp002 = πF.MakeArgs(6)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							πTemp003 = make([]*πg.Object, 3)
							πTemp003[0] = πg.NewInt(1).ToObject()
							πTemp003[1] = πg.NewInt(2).ToObject()
							πTemp003[2] = πg.NewInt(3).ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							πTemp002[2] = πTemp001
							πTemp002[3] = πg.NewInt(5).ToObject()
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[4] = πTemp001
							πTemp002[5] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 147: self.assertRaises(ValueError, mod.insort_left, [1, 2, 3], 5, -1, 3),
							πF.SetLineno(147)
							πTemp002 = πF.MakeArgs(6)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßinsort_left, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							πTemp003 = make([]*πg.Object, 3)
							πTemp003[0] = πg.NewInt(1).ToObject()
							πTemp003[1] = πg.NewInt(2).ToObject()
							πTemp003[2] = πg.NewInt(3).ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							πTemp002[2] = πTemp001
							πTemp002[3] = πg.NewInt(5).ToObject()
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[4] = πTemp001
							πTemp002[5] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 148: self.assertRaises(ValueError, mod.insort_right, [1, 2, 3], 5, -1, 3),
							πF.SetLineno(148)
							πTemp002 = πF.MakeArgs(6)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßinsort_right, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							πTemp003 = make([]*πg.Object, 3)
							πTemp003[0] = πg.NewInt(1).ToObject()
							πTemp003[1] = πg.NewInt(2).ToObject()
							πTemp003[2] = πg.NewInt(3).ToObject()
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							πTemp002[2] = πTemp001
							πTemp002[3] = πg.NewInt(5).ToObject()
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[4] = πTemp001
							πTemp002[5] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_negative_lo.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 150: def test_large_range(self):
					πF.SetLineno(150)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_large_range", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmod *πg.Object = πg.UnboundLocal; _ = µmod
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 152: mod = self.module
							πF.SetLineno(152)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							µmod = πTemp001
							// line 153: n = sys.maxsize
							πF.SetLineno(153)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmaxsize, nil); πE != nil {
								continue
							}
							µn = πTemp002
							// line 154: try:
							πF.SetLineno(154)
							πF.PushCheckpoint(2)
							// line 155: data = xrange(n-1)
							πF.SetLineno(155)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µdata = πTemp002
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 156: except OverflowError:
							πF.SetLineno(156)
						Label3:
							// line 157: self.skipTest("can't create a xrange() object of size `sys.maxsize`")
							πF.SetLineno(157)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("can't create a xrange() object of size `sys.maxsize`").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßskipTest, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 158: self.assertEqual(mod.bisect_left(data, n-3), n-3)
							πF.SetLineno(158)
							πTemp003 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp007[0] = µdata
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_left, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 159: self.assertEqual(mod.bisect_right(data, n-3), n-2)
							πF.SetLineno(159)
							πTemp003 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp007[0] = µdata
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_right, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 160: self.assertEqual(mod.bisect_left(data, n-3, n-10, n), n-3)
							πF.SetLineno(160)
							πTemp003 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp007[0] = µdata
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp007[2] = πTemp001
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp007[3] = µn
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_left, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 161: self.assertEqual(mod.bisect_right(data, n-3, n-10, n), n-2)
							πF.SetLineno(161)
							πTemp003 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp007[0] = µdata
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp007[2] = πTemp001
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp007[3] = µn
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_right, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_large_range.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 163: def test_large_pyrange(self):
					πF.SetLineno(163)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_large_pyrange", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmod *πg.Object = πg.UnboundLocal; _ = µmod
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 165: mod = self.module
							πF.SetLineno(165)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							µmod = πTemp001
							// line 166: n = sys.maxsize
							πF.SetLineno(166)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmaxsize, nil); πE != nil {
								continue
							}
							µn = πTemp002
							// line 167: data = Range(0, n-1)
							πF.SetLineno(167)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRange); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µdata = πTemp002
							// line 168: self.assertEqual(mod.bisect_left(data, n-3), n-3)
							πF.SetLineno(168)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp004[0] = µdata
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_left, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 169: self.assertEqual(mod.bisect_right(data, n-3), n-2)
							πF.SetLineno(169)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp004[0] = µdata
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_right, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 170: self.assertEqual(mod.bisect_left(data, n-3, n-10, n), n-3)
							πF.SetLineno(170)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp004[0] = µdata
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp004[2] = πTemp001
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp004[3] = µn
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_left, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 171: self.assertEqual(mod.bisect_right(data, n-3, n-10, n), n-2)
							πF.SetLineno(171)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp004[0] = µdata
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp004[2] = πTemp001
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp004[3] = µn
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßbisect_right, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 172: x = n - 100
							πF.SetLineno(172)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							µx = πTemp001
							// line 173: mod.insort_left(data, x, x - 50, x + 50)
							πF.SetLineno(173)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003[0] = µdata
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003[1] = µx
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µx, πg.NewInt(50).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µx, πg.NewInt(50).ToObject()); πE != nil {
								continue
							}
							πTemp003[3] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßinsort_left, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 174: self.assertEqual(data.last_insert, (x, x))
							πF.SetLineno(174)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata, ßlast_insert, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µx, µx).ToObject()
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 175: x = n - 200
							πF.SetLineno(175)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µn, πg.NewInt(200).ToObject()); πE != nil {
								continue
							}
							µx = πTemp001
							// line 176: mod.insort_right(data, x, x - 50, x + 50)
							πF.SetLineno(176)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003[0] = µdata
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003[1] = µx
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µx, πg.NewInt(50).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µx, πg.NewInt(50).ToObject()); πE != nil {
								continue
							}
							πTemp003[3] = πTemp001
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µmod, ßinsort_right, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 177: self.assertEqual(data.last_insert, (x + 1, x))
							πF.SetLineno(177)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdata, ßlast_insert, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µx, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, µx).ToObject()
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_large_pyrange.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 179: def test_random(self, n=25):
					πF.SetLineno(179)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: πg.NewInt(25).ToObject()}
					πTemp007 = πg.NewFunction(πg.NewCode("test_random", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
						var µ_random *πg.Object = πg.UnboundLocal; _ = µ_random
						var µrandrange *πg.Object = πg.UnboundLocal; _ = µrandrange
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var µip *πg.Object = πg.UnboundLocal; _ = µip
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []πg.Param
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 181: import random as _random
							πF.SetLineno(181)
							if πTemp002, πE = πg.ImportModule(πF, "random"); πE != nil {
								continue
							}
							πTemp001 = πTemp002[0]
							µ_random = πTemp001
							// line 182: randrange = _random.randrange
							πF.SetLineno(182)
							if πE = πg.CheckLocal(πF, µ_random, "_random"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µ_random, ßrandrange, nil); πE != nil {
								continue
							}
							µrandrange = πTemp001
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp002[0] = µn
							if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp005 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µi = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 185: data = [randrange(0, n, 2) for j in xrange(i)]
							πF.SetLineno(185)
							πTemp007 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/test_bisect.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µj *πg.Object = πg.UnboundLocal; _ = µj
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 []*πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 bool
								_ = πTemp005
								var πTemp006 bool
								_ = πTemp006
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										default: panic("unexpected function state")
										}
										πTemp002 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										πTemp002[0] = µi
										if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp005 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp005 {
											πF.PopCheckpoint()
											goto Label3
										}
										if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp006 = !isStop
										} else {
											πTemp006 = true
											µj = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 185: data = [randrange(0, n, 2) for j in xrange(i)]
										πF.SetLineno(185)
										πTemp002 = πF.MakeArgs(3)
										πTemp002[0] = πg.NewInt(0).ToObject()
										if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
											continue
										}
										πTemp002[1] = µn
										πTemp002[2] = πg.NewInt(2).ToObject()
										if πE = πg.CheckLocal(πF, µrandrange, "randrange"); πE != nil {
											continue
										}
										if πTemp003, πE = µrandrange.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										πF.PushCheckpoint(4)
										return πTemp003, nil
									Label4:
										πTemp004 = πSent
										continue
									Label2:
										if πE != nil || πR != nil {
											continue
										}
									Label3:
									}
									return nil, πE
								}).ToObject(), nil
							}), πF.Globals()).ToObject()
							if πTemp008, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
								continue
							}
							µdata = πTemp003
							// line 186: data.sort()
							πF.SetLineno(186)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdata, ßsort, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 187: elem = randrange(-1, n+1)
							πF.SetLineno(187)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µrandrange, "randrange"); πE != nil {
								continue
							}
							if πTemp003, πE = µrandrange.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µelem = πTemp003
							// line 188: ip = self.module.bisect_left(data, elem)
							πF.SetLineno(188)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp002[1] = µelem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µip = πTemp003
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.LT(πF, µip, πTemp009); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 189: if ip < len(data):
							πF.SetLineno(189)
						Label4:
							// line 190: self.assertTrue(elem <= data[ip])
							πF.SetLineno(190)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							πTemp008 = µip
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µdata, πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LE(πF, µelem, πTemp009); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µip, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 191: if ip > 0:
							πF.SetLineno(191)
						Label6:
							// line 192: self.assertTrue(data[ip-1] < elem)
							πF.SetLineno(192)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Sub(πF, µip, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp009
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µdata, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, πTemp009, µelem); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label7
						Label7:
							// line 193: ip = self.module.bisect_right(data, elem)
							πF.SetLineno(193)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp002[1] = µelem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µip = πTemp003
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.LT(πF, µip, πTemp009); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 194: if ip < len(data):
							πF.SetLineno(194)
						Label8:
							// line 195: self.assertTrue(elem < data[ip])
							πF.SetLineno(195)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							πTemp008 = µip
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µdata, πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µelem, πTemp009); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µip, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label10
							}
							goto Label11
							// line 196: if ip > 0:
							πF.SetLineno(196)
						Label10:
							// line 197: self.assertTrue(data[ip-1] <= elem)
							πF.SetLineno(197)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Sub(πF, µip, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008 = πTemp009
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µdata, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LE(πF, πTemp009, µelem); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label11
						Label11:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_random.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 199: def test_optionalSlicing(self):
					πF.SetLineno(199)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_optionalSlicing", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfunc *πg.Object = πg.UnboundLocal; _ = µfunc
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var µexpected *πg.Object = πg.UnboundLocal; _ = µexpected
						var µlo *πg.Object = πg.UnboundLocal; _ = µlo
						var µhi *πg.Object = πg.UnboundLocal; _ = µhi
						var µip *πg.Object = πg.UnboundLocal; _ = µip
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 []*πg.Object
						_ = πTemp014
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 4: goto Label4
							case 5: goto Label5
							case 7: goto Label7
							case 8: goto Label8
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßprecomputedCases, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp003 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
									continue
								}
								µfunc = πTemp005
								µdata = πTemp006
								µelem = πTemp007
								µexpected = πTemp008
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp009 = πF.MakeArgs(1)
							πTemp009[0] = πg.NewInt(4).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp004 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp010 = !isStop
							} else {
								πTemp010 = true
								µlo = πTemp005
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 202: lo = min(len(data), lo)
							πF.SetLineno(202)
							πTemp009 = πF.MakeArgs(2)
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp011[0] = µdata
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							πTemp009[1] = µlo
							if πTemp005, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							µlo = πTemp006
							πTemp009 = πF.MakeArgs(2)
							πTemp009[0] = πg.NewInt(3).ToObject()
							πTemp009[1] = πg.NewInt(8).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp005, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp010 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp010 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp006, πE = πg.Next(πF, πTemp005); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp012 = !isStop
							} else {
								πTemp012 = true
								µhi = πTemp006
							}
							if πE != nil || !πTemp012 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 204: hi = min(len(data), hi)
							πF.SetLineno(204)
							πTemp009 = πF.MakeArgs(2)
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp011[0] = µdata
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp009[0] = πTemp007
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							πTemp009[1] = µhi
							if πTemp006, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							µhi = πTemp007
							// line 205: ip = func(data, elem, lo, hi)
							πF.SetLineno(205)
							πTemp009 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp009[0] = µdata
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp009[1] = µelem
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							πTemp009[2] = µlo
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							πTemp009[3] = µhi
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp006, πE = µfunc.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							µip = πTemp006
							// line 206: self.assertTrue(lo <= ip <= hi)
							πF.SetLineno(206)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LE(πF, µlo, µip); πE != nil {
								continue
							}
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LE(πF, µip, µhi); πE != nil {
								continue
							}
						Label10:
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp008, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(µfunc == πTemp013).ToObject()
							πTemp006 = πTemp007
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.LT(πF, µip, µhi); πE != nil {
								continue
							}
							πTemp006 = πTemp007
						Label11:
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label12
							}
							goto Label13
							// line 207: if func is self.module.bisect_left and ip < hi:
							πF.SetLineno(207)
						Label12:
							// line 208: self.assertTrue(elem <= data[ip])
							πF.SetLineno(208)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							πTemp007 = µip
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µdata, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LE(πF, µelem, πTemp008); πE != nil {
								continue
							}
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp008, ßbisect_left, nil); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(µfunc == πTemp013).ToObject()
							πTemp006 = πTemp007
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GT(πF, µip, µlo); πE != nil {
								continue
							}
							πTemp006 = πTemp007
						Label14:
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label15
							}
							goto Label16
							// line 209: if func is self.module.bisect_left and ip > lo:
							πF.SetLineno(209)
						Label15:
							// line 210: self.assertTrue(data[ip-1] < elem)
							πF.SetLineno(210)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Sub(πF, µip, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µdata, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LT(πF, πTemp008, µelem); πE != nil {
								continue
							}
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label16
						Label16:
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp008, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(µfunc == πTemp013).ToObject()
							πTemp006 = πTemp007
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label17
							}
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.LT(πF, µip, µhi); πE != nil {
								continue
							}
							πTemp006 = πTemp007
						Label17:
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label18
							}
							goto Label19
							// line 211: if func is self.module.bisect_right and ip < hi:
							πF.SetLineno(211)
						Label18:
							// line 212: self.assertTrue(elem < data[ip])
							πF.SetLineno(212)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							πTemp007 = µip
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µdata, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LT(πF, µelem, πTemp008); πE != nil {
								continue
							}
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label19
						Label19:
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp008, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(µfunc == πTemp013).ToObject()
							πTemp006 = πTemp007
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label20
							}
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GT(πF, µip, µlo); πE != nil {
								continue
							}
							πTemp006 = πTemp007
						Label20:
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label21
							}
							goto Label22
							// line 213: if func is self.module.bisect_right and ip > lo:
							πF.SetLineno(213)
						Label21:
							// line 214: self.assertTrue(data[ip-1] <= elem)
							πF.SetLineno(214)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Sub(πF, µip, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µdata, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LE(πF, πTemp008, µelem); πE != nil {
								continue
							}
							πTemp009[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							goto Label22
						Label22:
							// line 215: self.assertEqual(ip, max(lo, min(hi, expected)))
							πF.SetLineno(215)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µip, "ip"); πE != nil {
								continue
							}
							πTemp009[0] = µip
							πTemp011 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							πTemp011[0] = µlo
							πTemp014 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							πTemp014[0] = µhi
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							πTemp014[1] = µexpected
							if πTemp006, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp011[1] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp009[1] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_optionalSlicing.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 217: def test_backcompatibility(self):
					πF.SetLineno(217)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_backcompatibility", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 218: self.assertEqual(self.module.bisect, self.module.bisect_right)
							πF.SetLineno(218)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßbisect, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßbisect_right, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_backcompatibility.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 220: def test_keyword_args(self):
					πF.SetLineno(220)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_keyword_args", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 221: data = [10, 20, 30, 40, 50]
							πF.SetLineno(221)
							πTemp001 = make([]*πg.Object, 5)
							πTemp001[0] = πg.NewInt(10).ToObject()
							πTemp001[1] = πg.NewInt(20).ToObject()
							πTemp001[2] = πg.NewInt(30).ToObject()
							πTemp001[3] = πg.NewInt(40).ToObject()
							πTemp001[4] = πg.NewInt(50).ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µdata = πTemp002
							// line 222: self.assertEqual(self.module.bisect_left(a=data, x=25, lo=1, hi=3), 2)
							πF.SetLineno(222)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"a", µdata},
								{"x", πg.NewInt(25).ToObject()},
								{"lo", πg.NewInt(1).ToObject()},
								{"hi", πg.NewInt(3).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßbisect_left, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 223: self.assertEqual(self.module.bisect_right(a=data, x=25, lo=1, hi=3), 2)
							πF.SetLineno(223)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"a", µdata},
								{"x", πg.NewInt(25).ToObject()},
								{"lo", πg.NewInt(1).ToObject()},
								{"hi", πg.NewInt(3).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßbisect_right, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 224: self.assertEqual(self.module.bisect(a=data, x=25, lo=1, hi=3), 2)
							πF.SetLineno(224)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"a", µdata},
								{"x", πg.NewInt(25).ToObject()},
								{"lo", πg.NewInt(1).ToObject()},
								{"hi", πg.NewInt(3).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßbisect, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 225: self.module.insort_left(a=data, x=25, lo=1, hi=3)
							πF.SetLineno(225)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"a", µdata},
								{"x", πg.NewInt(25).ToObject()},
								{"lo", πg.NewInt(1).ToObject()},
								{"hi", πg.NewInt(3).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßinsort_left, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, πTemp003); πE != nil {
								continue
							}
							// line 226: self.module.insort_right(a=data, x=25, lo=1, hi=3)
							πF.SetLineno(226)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"a", µdata},
								{"x", πg.NewInt(25).ToObject()},
								{"lo", πg.NewInt(1).ToObject()},
								{"hi", πg.NewInt(3).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßinsort_right, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, πTemp003); πE != nil {
								continue
							}
							// line 227: self.module.insort(a=data, x=25, lo=1, hi=3)
							πF.SetLineno(227)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"a", µdata},
								{"x", πg.NewInt(25).ToObject()},
								{"lo", πg.NewInt(1).ToObject()},
								{"hi", πg.NewInt(3).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßinsort, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, πTemp003); πE != nil {
								continue
							}
							// line 228: self.assertEqual(data, [10, 20, 25, 25, 25, 30, 40, 50])
							πF.SetLineno(228)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							πTemp005 = make([]*πg.Object, 8)
							πTemp005[0] = πg.NewInt(10).ToObject()
							πTemp005[1] = πg.NewInt(20).ToObject()
							πTemp005[2] = πg.NewInt(25).ToObject()
							πTemp005[3] = πg.NewInt(25).ToObject()
							πTemp005[4] = πg.NewInt(25).ToObject()
							πTemp005[5] = πg.NewInt(30).ToObject()
							πTemp005[6] = πg.NewInt(40).ToObject()
							πTemp005[7] = πg.NewInt(50).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_keyword_args.ToObject(), πTemp010); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TestBisect").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestBisect.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 238: class TestInsort(unittest.TestCase):
			πF.SetLineno(238)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestInsort", "/usr/lib/python2.7/test/test_bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 240: module = py_bisect
					πF.SetLineno(240)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpy_bisect); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßmodule.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 242: def test_vsBuiltinSort(self, n=500):
					πF.SetLineno(242)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: πg.NewInt(500).ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("test_vsBuiltinSort", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
						var µ_random *πg.Object = πg.UnboundLocal; _ = µ_random
						var µchoice *πg.Object = πg.UnboundLocal; _ = µchoice
						var µinsorted *πg.Object = πg.UnboundLocal; _ = µinsorted
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µdigit *πg.Object = πg.UnboundLocal; _ = µdigit
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 244: import random as _random
							πF.SetLineno(244)
							if πTemp002, πE = πg.ImportModule(πF, "random"); πE != nil {
								continue
							}
							πTemp001 = πTemp002[0]
							µ_random = πTemp001
							// line 245: choice = _random.choice
							πF.SetLineno(245)
							if πE = πg.CheckLocal(πF, µ_random, "_random"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µ_random, ßchoice, nil); πE != nil {
								continue
							}
							µchoice = πTemp001
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßUserList); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp007 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µinsorted = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp002[0] = µn
							if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp008 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µi = πTemp004
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 249: digit = choice("0123456789")
							πF.SetLineno(249)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ß0123456789.ToObject()
							if πE = πg.CheckLocal(πF, µchoice, "choice"); πE != nil {
								continue
							}
							if πTemp004, πE = µchoice.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdigit = πTemp004
							if πE = πg.CheckLocal(πF, µdigit, "digit"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, ß02468.ToObject(), µdigit); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label7
							}
							goto Label8
							// line 250: if digit in "02468":
							πF.SetLineno(250)
						Label7:
							// line 251: f = self.module.insort_left
							πF.SetLineno(251)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßinsort_left, nil); πE != nil {
								continue
							}
							µf = πTemp005
							goto Label9
						Label8:
							// line 253: f = self.module.insort_right
							πF.SetLineno(253)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßinsort_right, nil); πE != nil {
								continue
							}
							µf = πTemp005
							goto Label9
						Label9:
							// line 254: f(insorted, digit)
							πF.SetLineno(254)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µinsorted, "insorted"); πE != nil {
								continue
							}
							πTemp002[0] = µinsorted
							if πE = πg.CheckLocal(πF, µdigit, "digit"); πE != nil {
								continue
							}
							πTemp002[1] = µdigit
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp004, πE = µf.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 255: self.assertEqual(sorted(insorted), insorted)
							πF.SetLineno(255)
							πTemp002 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinsorted, "insorted"); πE != nil {
								continue
							}
							πTemp010[0] = µinsorted
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µinsorted, "insorted"); πE != nil {
								continue
							}
							πTemp002[1] = µinsorted
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_vsBuiltinSort.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 257: def test_backcompatibility(self):
					πF.SetLineno(257)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_backcompatibility", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 258: self.assertEqual(self.module.insort, self.module.insort_right)
							πF.SetLineno(258)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsort, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsort_right, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_backcompatibility.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 260: def test_listDerived(self):
					πF.SetLineno(260)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_listDerived", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µList *πg.Object = πg.UnboundLocal; _ = µList
						var µlst *πg.Object = πg.UnboundLocal; _ = µlst
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 261: class List(list):
							πF.SetLineno(261)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							πTemp001 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("List", "/usr/lib/python2.7/test/test_bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								var πTemp001 []*πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 []πg.Param
								_ = πTemp003
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 262: data = []
									πF.SetLineno(262)
									πTemp001 = make([]*πg.Object, 0)
									πTemp002 = πg.NewList(πTemp001...).ToObject()
									if πE = πClass.SetItem(πF, ßdata.ToObject(), πTemp002); πE != nil {
										continue
									}
									// line 263: def insert(self, index, item):
									πF.SetLineno(263)
									πTemp003 = make([]πg.Param, 3)
									πTemp003[0] = πg.Param{Name: "self", Def: nil}
									πTemp003[1] = πg.Param{Name: "index", Def: nil}
									πTemp003[2] = πg.Param{Name: "item", Def: nil}
									πTemp002 = πg.NewFunction(πg.NewCode("insert", "/usr/lib/python2.7/test/test_bisect.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µindex *πg.Object = πArgs[1]; _ = µindex
										var µitem *πg.Object = πArgs[2]; _ = µitem
										var πTemp001 []*πg.Object
										_ = πTemp001
										var πTemp002 *πg.Object
										_ = πTemp002
										var πTemp003 *πg.Object
										_ = πTemp003
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 264: self.data.insert(index, item)
											πF.SetLineno(264)
											πTemp001 = πF.MakeArgs(2)
											if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
												continue
											}
											πTemp001[0] = µindex
											if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
												continue
											}
											πTemp001[1] = µitem
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp001)
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ßinsert.ToObject(), πTemp002); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("List").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µList = πTemp005
							// line 266: lst = List()
							πF.SetLineno(266)
							if πE = πg.CheckLocal(πF, µList, "List"); πE != nil {
								continue
							}
							if πTemp002, πE = µList.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlst = πTemp002
							// line 267: self.module.insort_left(lst, 10)
							πF.SetLineno(267)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp003[0] = µlst
							πTemp003[1] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßinsort_left, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 268: self.module.insort_right(lst, 5)
							πF.SetLineno(268)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							πTemp003[0] = µlst
							πTemp003[1] = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßinsort_right, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 269: self.assertEqual([5, 10], lst.data)
							πF.SetLineno(269)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = make([]*πg.Object, 2)
							πTemp006[0] = πg.NewInt(5).ToObject()
							πTemp006[1] = πg.NewInt(10).ToObject()
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlst, ßdata, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_listDerived.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TestInsort").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestInsort.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 280: class LenOnly(object):
			πF.SetLineno(280)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("LenOnly", "/usr/lib/python2.7/test/test_bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 281: "Dummy sequence class defining __len__ but not __getitem__."
					πF.SetLineno(281)
					// line 281: "Dummy sequence class defining __len__ but not __getitem__."
					πF.SetLineno(281)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Dummy sequence class defining __len__ but not __getitem__.").ToObject()); πE != nil {
						continue
					}
					// line 282: def __len__(self):
					πF.SetLineno(282)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 283: return 10
							πF.SetLineno(283)
							πR = πg.NewInt(10).ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 288: def __getitem__(self, ndx):
					πF.SetLineno(288)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ndx", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µndx *πg.Object = πArgs[1]; _ = µndx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							// line 289: raise AttributeError
							πF.SetLineno(289)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("LenOnly").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLenOnly.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 291: class GetOnly(object):
			πF.SetLineno(291)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("GetOnly", "/usr/lib/python2.7/test/test_bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 292: "Dummy sequence class defining __getitem__ but not __len__."
					πF.SetLineno(292)
					// line 292: "Dummy sequence class defining __getitem__ but not __len__."
					πF.SetLineno(292)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Dummy sequence class defining __getitem__ but not __len__.").ToObject()); πE != nil {
						continue
					}
					// line 293: def __getitem__(self, ndx):
					πF.SetLineno(293)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ndx", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µndx *πg.Object = πArgs[1]; _ = µndx
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 294: return 10
							πF.SetLineno(294)
							πR = πg.NewInt(10).ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 296: def __len__(self):
					πF.SetLineno(296)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							// line 297: raise AttributeError
							πF.SetLineno(297)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("GetOnly").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGetOnly.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 299: class CmpErr(object):
			πF.SetLineno(299)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("CmpErr", "/usr/lib/python2.7/test/test_bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 300: "Dummy element that always raises an error during comparison"
					πF.SetLineno(300)
					// line 300: "Dummy element that always raises an error during comparison"
					πF.SetLineno(300)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Dummy element that always raises an error during comparison").ToObject()); πE != nil {
						continue
					}
					// line 301: def __cmp__(self, other):
					πF.SetLineno(301)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__cmp__", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
								continue
							}
							// line 302: raise ZeroDivisionError
							πF.SetLineno(302)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__cmp__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("CmpErr").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCmpErr.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 304: class TestErrorHandling(unittest.TestCase):
			πF.SetLineno(304)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestErrorHandling", "/usr/lib/python2.7/test/test_bisect.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp006
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 306: module = py_bisect
					πF.SetLineno(306)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpy_bisect); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßmodule.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 308: def test_non_sequence(self):
					πF.SetLineno(308)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_non_sequence", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßinsort_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßinsort_right, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp004, πTemp005, πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp008 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µf = πTemp002
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 311: self.assertRaises(TypeError, f, 10, 10)
							πF.SetLineno(311)
							πTemp010 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp010[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp010[1] = µf
							πTemp010[2] = πg.NewInt(10).ToObject()
							πTemp010[3] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_non_sequence.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 313: def test_len_only(self):
					πF.SetLineno(313)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_len_only", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßinsort_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßinsort_right, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp004, πTemp005, πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp008 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µf = πTemp002
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 316: self.assertRaises(AttributeError, f, LenOnly(), 10)
							πF.SetLineno(316)
							πTemp010 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp010[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp010[1] = µf
							if πTemp002, πE = πg.ResolveGlobal(πF, ßLenOnly); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp010[2] = πTemp003
							πTemp010[3] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_len_only.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 318: def test_get_only(self):
					πF.SetLineno(318)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_get_only", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßinsort_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßinsort_right, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp004, πTemp005, πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp008 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µf = πTemp002
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 321: self.assertRaises(AttributeError, f, GetOnly(), 10)
							πF.SetLineno(321)
							πTemp010 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp010[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp010[1] = µf
							if πTemp002, πE = πg.ResolveGlobal(πF, ßGetOnly); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp010[2] = πTemp003
							πTemp010[3] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_get_only.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 323: def test_cmp_err(self):
					πF.SetLineno(323)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_cmp_err", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseq *πg.Object = πg.UnboundLocal; _ = µseq
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 324: seq = [CmpErr(), CmpErr(), CmpErr()]
							πF.SetLineno(324)
							πTemp001 = make([]*πg.Object, 3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßCmpErr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßCmpErr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßCmpErr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µseq = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßbisect_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßbisect_right, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßinsort_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßinsort_right, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple4(πTemp005, πTemp006, πTemp007, πTemp008).ToObject()
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp009 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp010 = !isStop
							} else {
								πTemp010 = true
								µf = πTemp003
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 327: self.assertRaises(ZeroDivisionError, f, seq, 10)
							πF.SetLineno(327)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[1] = µf
							if πE = πg.CheckLocal(πF, µseq, "seq"); πE != nil {
								continue
							}
							πTemp001[2] = µseq
							πTemp001[3] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_cmp_err.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 329: def test_arg_parsing(self):
					πF.SetLineno(329)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_arg_parsing", "/usr/lib/python2.7/test/test_bisect.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßbisect_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßbisect_right, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßinsort_left, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmodule, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßinsort_right, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp004, πTemp005, πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp008 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µf = πTemp002
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 332: self.assertRaises(TypeError, f, 10)
							πF.SetLineno(332)
							πTemp010 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp010[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp010[1] = µf
							πTemp010[2] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_arg_parsing.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TestErrorHandling").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestErrorHandling.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 342: libreftest = """
			πF.SetLineno(342)
			// line 342: libreftest = """
			πF.SetLineno(342)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nExample from the Library Reference:  Doc/library/bisect.rst\nThe bisect() function is generally useful for categorizing numeric data.\nThis example uses bisect() to look up a letter grade for an exam total\n(say) based on a set of ordered numeric breakpoints: 85 and up is an `A',\n75..84 is a `B', etc.\n    >>> grades = \"FEDCBA\"\n    >>> breakpoints = [30, 44, 66, 75, 85]\n    >>> from bisect import bisect\n    >>> def grade(total):\n    ...           return grades[bisect(breakpoints, total)]\n    ...\n    >>> grade(66)\n    'C'\n    >>> map(grade, [33, 99, 77, 44, 12, 88])\n    ['E', 'A', 'B', 'D', 'F', 'A']\n").ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlibreftest.ToObject(), πg.NewStr("\nExample from the Library Reference:  Doc/library/bisect.rst\nThe bisect() function is generally useful for categorizing numeric data.\nThis example uses bisect() to look up a letter grade for an exam total\n(say) based on a set of ordered numeric breakpoints: 85 and up is an `A',\n75..84 is a `B', etc.\n    >>> grades = \"FEDCBA\"\n    >>> breakpoints = [30, 44, 66, 75, 85]\n    >>> from bisect import bisect\n    >>> def grade(total):\n    ...           return grades[bisect(breakpoints, total)]\n    ...\n    >>> grade(66)\n    'C'\n    >>> map(grade, [33, 99, 77, 44, 12, 88])\n    ['E', 'A', 'B', 'D', 'F', 'A']\n").ToObject()); πE != nil {
				continue
			}
			// line 362: __test__ = {'libreftest' : libreftest}
			πF.SetLineno(362)
			πTemp006 = πg.NewDict()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßlibreftest); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ßlibreftest.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πTemp006.ToObject()
			if πE = πF.Globals().SetItem(πF, ß__test__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 364: def test_main(verbose=None):
			πF.SetLineno(364)
			πTemp008 = make([]πg.Param, 1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp008[0] = πg.Param{Name: "verbose", Def: πTemp003}
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "/usr/lib/python2.7/test/test_bisect.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µverbose *πg.Object = πArgs[0]; _ = µverbose
				var µtest_classes *πg.Object = πg.UnboundLocal; _ = µtest_classes
				var µcounts *πg.Object = πg.UnboundLocal; _ = µcounts
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 370: test_classes = [TestBisect, TestInsort, TestErrorHandling]
					πF.SetLineno(370)
					πTemp001 = make([]*πg.Object, 3)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTestBisect); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTestInsort); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTestErrorHandling); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µtest_classes = πTemp002
					// line 372: test_support.run_unittest(*test_classes)
					πF.SetLineno(372)
					if πE = πg.CheckLocal(πF, µtest_classes, "test_classes"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrun_unittest, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp003, nil, µtest_classes, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
						continue
					}
					πTemp002 = µverbose
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					πTemp001 = πF.MakeArgs(2)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp001[1] = ßgettotalrefcount.ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp005
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 376: if verbose and hasattr(sys, "gettotalrefcount"):
					πF.SetLineno(376)
				Label2:
					// line 378: counts = [None] * 5
					πF.SetLineno(378)
					πTemp001 = make([]*πg.Object, 1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					if πTemp002, πE = πg.Mul(πF, πTemp003, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					µcounts = πTemp002
					πTemp001 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcounts, "counts"); πE != nil {
						continue
					}
					πTemp006[0] = µcounts
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp001[0] = πTemp005
					if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(5)
					πTemp004 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µi = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(4)            
					// line 380: test_support.run_unittest(*test_classes)
					πF.SetLineno(380)
					if πE = πg.CheckLocal(πF, µtest_classes, "test_classes"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßrun_unittest, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp005, nil, µtest_classes, nil, nil); πE != nil {
						continue
					}
					// line 382: counts[i] = sys.gettotalrefcount()
					πF.SetLineno(382)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßgettotalrefcount, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcounts, "counts"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp008 = µi
					if πE = πg.SetItem(πF, µcounts, πTemp008, πTemp005); πE != nil {
						continue
					}
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 383: print counts
					πF.SetLineno(383)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µcounts, "counts"); πE != nil {
						continue
					}
					πTemp001[0] = µcounts
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
					goto Label3
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp004, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label3
			}
			goto Label4
			// line 385: if __name__ == "__main__":
			πF.SetLineno(385)
		Label3:
			// line 386: test_main(verbose=True)
			πF.SetLineno(386)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp009 = πg.KWArgs{
				{"verbose", πTemp003},
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.Call(πF, nil, πTemp009); πE != nil {
				continue
			}
			goto Label4
		Label4:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_bisect", Code)
}

