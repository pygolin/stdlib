package seq_tests
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/unittest"
	// _ "github.com/pygolin/stdlib/test/test_support"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß123 := πg.InternStr("123")
		ßCommonTest := πg.InternStr("CommonTest")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßIndexError := πg.InternStr("IndexError")
		ßIterFunc := πg.InternStr("IterFunc")
		ßIterFuncStop := πg.InternStr("IterFuncStop")
		ßIterGen := πg.InternStr("IterGen")
		ßIterGenExc := πg.InternStr("IterGenExc")
		ßIterNextOnly := πg.InternStr("IterNextOnly")
		ßIterNoNext := πg.InternStr("IterNoNext")
		ßLyingList := πg.InternStr("LyingList")
		ßLyingTuple := πg.InternStr("LyingTuple")
		ßMemoryError := πg.InternStr("MemoryError")
		ßNone := πg.InternStr("None")
		ßSequence := πg.InternStr("Sequence")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTestCase := πg.InternStr("TestCase")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ßZeroDivisionError := πg.InternStr("ZeroDivisionError")
		ß__contains__ := πg.InternStr("__contains__")
		ß__data := πg.InternStr("__data")
		ß__doc__ := πg.InternStr("__doc__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__getslice__ := πg.InternStr("__getslice__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__imul__ := πg.InternStr("__imul__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__len__ := πg.InternStr("__len__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__mul__ := πg.InternStr("__mul__")
		ß__name__ := πg.InternStr("__name__")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertIn := πg.InternStr("assertIn")
		ßassertIsNot := πg.InternStr("assertIsNot")
		ßassertNotIn := πg.InternStr("assertNotIn")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßcheck_free_after_iterating := πg.InternStr("check_free_after_iterating")
		ßcount := πg.InternStr("count")
		ßdo := πg.InternStr("do")
		ßeggs := πg.InternStr("eggs")
		ßgrumpy := πg.InternStr("grumpy")
		ßhasattr := πg.InternStr("hasattr")
		ßi := πg.InternStr("i")
		ßid := πg.InternStr("id")
		ßindex := πg.InternStr("index")
		ßiter := πg.InternStr("iter")
		ßiterfunc := πg.InternStr("iterfunc")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßlong := πg.InternStr("long")
		ßmax := πg.InternStr("max")
		ßmaxint := πg.InternStr("maxint")
		ßmin := πg.InternStr("min")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßpow := πg.InternStr("pow")
		ßrange := πg.InternStr("range")
		ßreversed := πg.InternStr("reversed")
		ßseqn := πg.InternStr("seqn")
		ßskip := πg.InternStr("skip")
		ßslice := πg.InternStr("slice")
		ßspam := πg.InternStr("spam")
		ßspameggs := πg.InternStr("spameggs")
		ßstr := πg.InternStr("str")
		ßsupport := πg.InternStr("support")
		ßsys := πg.InternStr("sys")
		ßtest_addmul := πg.InternStr("test_addmul")
		ßtest_bigrepeat := πg.InternStr("test_bigrepeat")
		ßtest_constructors := πg.InternStr("test_constructors")
		ßtest_contains := πg.InternStr("test_contains")
		ßtest_contains_fake := πg.InternStr("test_contains_fake")
		ßtest_contains_order := πg.InternStr("test_contains_order")
		ßtest_count := πg.InternStr("test_count")
		ßtest_free_after_iterating := πg.InternStr("test_free_after_iterating")
		ßtest_getitem := πg.InternStr("test_getitem")
		ßtest_getitemoverwriteiter := πg.InternStr("test_getitemoverwriteiter")
		ßtest_getslice := πg.InternStr("test_getslice")
		ßtest_iadd := πg.InternStr("test_iadd")
		ßtest_imul := πg.InternStr("test_imul")
		ßtest_index := πg.InternStr("test_index")
		ßtest_len := πg.InternStr("test_len")
		ßtest_minmax := πg.InternStr("test_minmax")
		ßtest_repeat := πg.InternStr("test_repeat")
		ßtest_subscript := πg.InternStr("test_subscript")
		ßtest_truth := πg.InternStr("test_truth")
		ßtuple := πg.InternStr("tuple")
		ßtype2test := πg.InternStr("type2test")
		ßunittest := πg.InternStr("unittest")
		ßx := πg.InternStr("x")
		ßxrange := πg.InternStr("xrange")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Dict
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """
			πF.SetLineno(1)
			// line 1: """
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nTests common to tuple, list and UserList.UserList\n").ToObject()); πE != nil {
				continue
			}
			// line 5: import unittest
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: import sys
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: from test import test_support as support
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßsupport.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: def iterfunc(seqn):
			πF.SetLineno(11)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "seqn", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("iterfunc", "/usr/lib/python2.7/test/seq_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µseqn *πg.Object = πArgs[0]; _ = µseqn
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 bool
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
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
						// line 12: 'Regular generator'
						πF.SetLineno(12)
						if πE = πg.CheckLocal(πF, µseqn, "seqn"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, µseqn); πE != nil {
							continue
						}
						πF.PushCheckpoint(2)
						πTemp002 = false
					Label1:
						if πE != nil || πR != nil {
							continue
						}
						if πTemp002 {
							πF.PopCheckpoint()
							goto Label3
						}
						if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
							} else if isStop {
								πE = nil
								πF.RestoreExc(nil, nil)
							}
							πTemp003 = !isStop
						} else {
							πTemp003 = true
							µi = πTemp004
						}
						if πE != nil || !πTemp003 {
							continue
						}
						πF.PushCheckpoint(1)            
						// line 14: yield i
						πF.SetLineno(14)
						if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
							continue
						}
						πF.PushCheckpoint(4)
						return µi, nil
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
			if πE = πF.Globals().SetItem(πF, ßiterfunc.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: 'Regular generator'
			πF.SetLineno(12)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Regular generator").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßiterfunc); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 16: class Sequence(object):
			πF.SetLineno(16)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Sequence", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 17: 'Sequence using __getitem__'
					πF.SetLineno(17)
					// line 17: 'Sequence using __getitem__'
					πF.SetLineno(17)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Sequence using __getitem__").ToObject()); πE != nil {
						continue
					}
					// line 18: def __init__(self, seqn):
					πF.SetLineno(18)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "seqn", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseqn *πg.Object = πArgs[1]; _ = µseqn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 19: self.seqn = seqn
							πF.SetLineno(19)
							if πE = πg.CheckLocal(πF, µseqn, "seqn"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µseqn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseqn, πTemp001); πE != nil {
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
					// line 20: def __getitem__(self, i):
					πF.SetLineno(20)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µi *πg.Object = πArgs[1]; _ = µi
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
							// line 21: return self.seqn[i]
							πF.SetLineno(21)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßseqn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πR = πTemp002
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
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("Sequence").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSequence.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 23: class IterFunc(object):
			πF.SetLineno(23)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("IterFunc", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 24: 'Sequence using iterator protocol'
					πF.SetLineno(24)
					// line 24: 'Sequence using iterator protocol'
					πF.SetLineno(24)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Sequence using iterator protocol").ToObject()); πE != nil {
						continue
					}
					// line 25: def __init__(self, seqn):
					πF.SetLineno(25)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "seqn", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseqn *πg.Object = πArgs[1]; _ = µseqn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 26: self.seqn = seqn
							πF.SetLineno(26)
							if πE = πg.CheckLocal(πF, µseqn, "seqn"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µseqn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseqn, πTemp001); πE != nil {
								continue
							}
							// line 27: self.i = 0
							πF.SetLineno(27)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßi, πTemp001); πE != nil {
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
					// line 28: def __iter__(self):
					πF.SetLineno(28)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 29: return self
							πF.SetLineno(29)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 30: def next(self):
					πF.SetLineno(30)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µv *πg.Object = πg.UnboundLocal; _ = µv
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßi, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßseqn, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GE(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 31: if self.i >= len(self.seqn): raise StopIteration
							πF.SetLineno(31)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							// line 31: if self.i >= len(self.seqn): raise StopIteration
							πF.SetLineno(31)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							goto Label2
						Label2:
							// line 32: v = self.seqn[self.i]
							πF.SetLineno(32)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßi, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßseqn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µv = πTemp002
							// line 33: self.i += 1
							πF.SetLineno(33)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßi, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßi, πTemp002); πE != nil {
								continue
							}
							// line 34: return v
							πF.SetLineno(34)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πR = µv
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("IterFunc").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIterFunc.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 36: class IterGen(object):
			πF.SetLineno(36)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("IterGen", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 37: 'Sequence using iterator protocol defined with a generator'
					πF.SetLineno(37)
					// line 37: 'Sequence using iterator protocol defined with a generator'
					πF.SetLineno(37)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Sequence using iterator protocol defined with a generator").ToObject()); πE != nil {
						continue
					}
					// line 38: def __init__(self, seqn):
					πF.SetLineno(38)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "seqn", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseqn *πg.Object = πArgs[1]; _ = µseqn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 39: self.seqn = seqn
							πF.SetLineno(39)
							if πE = πg.CheckLocal(πF, µseqn, "seqn"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µseqn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseqn, πTemp001); πE != nil {
								continue
							}
							// line 40: self.i = 0
							πF.SetLineno(40)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßi, πTemp001); πE != nil {
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
					// line 41: def __iter__(self):
					πF.SetLineno(41)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µval *πg.Object = πg.UnboundLocal; _ = µval
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
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
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßseqn, nil); πE != nil {
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
									µval = πTemp002
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 43: yield val
								πF.SetLineno(43)
								if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return µval, nil
							Label4:
								πTemp002 = πSent
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("IterGen").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIterGen.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 45: class IterNextOnly(object):
			πF.SetLineno(45)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("IterNextOnly", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 46: 'Missing __getitem__ and __iter__'
					πF.SetLineno(46)
					// line 46: 'Missing __getitem__ and __iter__'
					πF.SetLineno(46)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Missing __getitem__ and __iter__").ToObject()); πE != nil {
						continue
					}
					// line 47: def __init__(self, seqn):
					πF.SetLineno(47)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "seqn", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseqn *πg.Object = πArgs[1]; _ = µseqn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 48: self.seqn = seqn
							πF.SetLineno(48)
							if πE = πg.CheckLocal(πF, µseqn, "seqn"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µseqn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseqn, πTemp001); πE != nil {
								continue
							}
							// line 49: self.i = 0
							πF.SetLineno(49)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßi, πTemp001); πE != nil {
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
					// line 50: def next(self):
					πF.SetLineno(50)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µv *πg.Object = πg.UnboundLocal; _ = µv
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßi, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßseqn, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GE(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 51: if self.i >= len(self.seqn): raise StopIteration
							πF.SetLineno(51)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							// line 51: if self.i >= len(self.seqn): raise StopIteration
							πF.SetLineno(51)
							πE = πF.Raise(πTemp001, nil, nil)
							continue
							goto Label2
						Label2:
							// line 52: v = self.seqn[self.i]
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßi, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßseqn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µv = πTemp002
							// line 53: self.i += 1
							πF.SetLineno(53)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßi, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßi, πTemp002); πE != nil {
								continue
							}
							// line 54: return v
							πF.SetLineno(54)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πR = µv
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("IterNextOnly").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIterNextOnly.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 56: class IterNoNext(object):
			πF.SetLineno(56)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("IterNoNext", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 57: 'Iterator missing next()'
					πF.SetLineno(57)
					// line 57: 'Iterator missing next()'
					πF.SetLineno(57)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Iterator missing next()").ToObject()); πE != nil {
						continue
					}
					// line 58: def __init__(self, seqn):
					πF.SetLineno(58)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "seqn", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseqn *πg.Object = πArgs[1]; _ = µseqn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 59: self.seqn = seqn
							πF.SetLineno(59)
							if πE = πg.CheckLocal(πF, µseqn, "seqn"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µseqn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseqn, πTemp001); πE != nil {
								continue
							}
							// line 60: self.i = 0
							πF.SetLineno(60)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßi, πTemp001); πE != nil {
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
					// line 61: def __iter__(self):
					πF.SetLineno(61)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 62: return self
							πF.SetLineno(62)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("IterNoNext").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIterNoNext.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 64: class IterGenExc(object):
			πF.SetLineno(64)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("IterGenExc", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 65: 'Test propagation of exceptions'
					πF.SetLineno(65)
					// line 65: 'Test propagation of exceptions'
					πF.SetLineno(65)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Test propagation of exceptions").ToObject()); πE != nil {
						continue
					}
					// line 66: def __init__(self, seqn):
					πF.SetLineno(66)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "seqn", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseqn *πg.Object = πArgs[1]; _ = µseqn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 67: self.seqn = seqn
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µseqn, "seqn"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µseqn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseqn, πTemp001); πE != nil {
								continue
							}
							// line 68: self.i = 0
							πF.SetLineno(68)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßi, πTemp001); πE != nil {
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
					// line 69: def __iter__(self):
					πF.SetLineno(69)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 70: return self
							πF.SetLineno(70)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 71: def next(self):
					πF.SetLineno(71)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 72: 3 // 0
							πF.SetLineno(72)
							if πTemp001, πE = πg.FloorDiv(πF, πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("IterGenExc").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIterGenExc.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 74: class IterFuncStop(object):
			πF.SetLineno(74)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("IterFuncStop", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 75: 'Test immediate stop'
					πF.SetLineno(75)
					// line 75: 'Test immediate stop'
					πF.SetLineno(75)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Test immediate stop").ToObject()); πE != nil {
						continue
					}
					// line 76: def __init__(self, seqn):
					πF.SetLineno(76)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "seqn", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseqn *πg.Object = πArgs[1]; _ = µseqn
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 77: pass
							πF.SetLineno(77)
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
					// line 78: def __iter__(self):
					πF.SetLineno(78)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 79: return self
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πR = µself
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 80: def next(self):
					πF.SetLineno(80)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							// line 81: raise StopIteration
							πF.SetLineno(81)
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
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("IterFuncStop").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIterFuncStop.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 88: class LyingTuple(tuple):
			πF.SetLineno(88)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("LyingTuple", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 89: def __iter__(self):
					πF.SetLineno(89)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								default: panic("unexpected function state")
								}
								// line 90: yield 1
								πF.SetLineno(90)
								πF.PushCheckpoint(1)
								return πg.NewInt(1).ToObject(), nil
							Label1:
								πTemp001 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("LyingTuple").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLyingTuple.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 92: class LyingList(list):
			πF.SetLineno(92)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("LyingList", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 93: def __iter__(self):
					πF.SetLineno(93)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								default: panic("unexpected function state")
								}
								// line 94: yield 1
								πF.SetLineno(94)
								πF.PushCheckpoint(1)
								return πg.NewInt(1).ToObject(), nil
							Label1:
								πTemp001 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("LyingList").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLyingList.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 96: class CommonTest(unittest.TestCase):
			πF.SetLineno(96)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp006 = πg.NewDict()
			if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
				continue
			}
			_, πE = πg.NewCode("CommonTest", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 *πg.Object
				_ = πTemp017
				var πTemp018 *πg.Object
				_ = πTemp018
				var πTemp019 *πg.Object
				_ = πTemp019
				var πTemp020 *πg.Object
				_ = πTemp020
				var πTemp021 *πg.Object
				_ = πTemp021
				var πTemp022 *πg.Object
				_ = πTemp022
				var πTemp023 *πg.Object
				_ = πTemp023
				var πTemp024 *πg.Object
				_ = πTemp024
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 98: type2test = None
					πF.SetLineno(98)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 100: def test_constructors(self):
					πF.SetLineno(100)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_constructors", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl0 *πg.Object = πg.UnboundLocal; _ = µl0
						var µl1 *πg.Object = πg.UnboundLocal; _ = µl1
						var µl2 *πg.Object = πg.UnboundLocal; _ = µl2
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µu0 *πg.Object = πg.UnboundLocal; _ = µu0
						var µu1 *πg.Object = πg.UnboundLocal; _ = µu1
						var µu2 *πg.Object = πg.UnboundLocal; _ = µu2
						var µuu *πg.Object = πg.UnboundLocal; _ = µuu
						var µuu0 *πg.Object = πg.UnboundLocal; _ = µuu0
						var µuu1 *πg.Object = πg.UnboundLocal; _ = µuu1
						var µuu2 *πg.Object = πg.UnboundLocal; _ = µuu2
						var µv *πg.Object = πg.UnboundLocal; _ = µv
						var µOtherSeq *πg.Object = πg.UnboundLocal; _ = µOtherSeq
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µv0 *πg.Object = πg.UnboundLocal; _ = µv0
						var µvv *πg.Object = πg.UnboundLocal; _ = µvv
						var µg *πg.Object = πg.UnboundLocal; _ = µg
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 bool
						_ = πTemp013
						var πTemp014 []*πg.Object
						_ = πTemp014
						var πTemp015 []πg.Param
						_ = πTemp015
						var πTemp016 []*πg.Object
						_ = πTemp016
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
							// line 101: l0 = []
							πF.SetLineno(101)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µl0 = πTemp002
							// line 102: l1 = [0]
							πF.SetLineno(102)
							πTemp001 = make([]*πg.Object, 1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µl1 = πTemp002
							// line 103: l2 = [0, 1]
							πF.SetLineno(103)
							πTemp001 = make([]*πg.Object, 2)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp001[1] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µl2 = πTemp002
							// line 105: u = self.type2test()
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µu = πTemp003
							// line 106: u0 = self.type2test(l0)
							πF.SetLineno(106)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl0, "l0"); πE != nil {
								continue
							}
							πTemp001[0] = µl0
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu0 = πTemp003
							// line 107: u1 = self.type2test(l1)
							πF.SetLineno(107)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl1, "l1"); πE != nil {
								continue
							}
							πTemp001[0] = µl1
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu1 = πTemp003
							// line 108: u2 = self.type2test(l2)
							πF.SetLineno(108)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl2, "l2"); πE != nil {
								continue
							}
							πTemp001[0] = µl2
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu2 = πTemp003
							// line 110: uu = self.type2test(u)
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µuu = πTemp003
							// line 111: uu0 = self.type2test(u0)
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu0, "u0"); πE != nil {
								continue
							}
							πTemp001[0] = µu0
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µuu0 = πTemp003
							// line 112: uu1 = self.type2test(u1)
							πF.SetLineno(112)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu1, "u1"); πE != nil {
								continue
							}
							πTemp001[0] = µu1
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µuu1 = πTemp003
							// line 113: uu2 = self.type2test(u2)
							πF.SetLineno(113)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp001[0] = µu2
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µuu2 = πTemp003
							// line 115: v = self.type2test(tuple(u))
							πF.SetLineno(115)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp004[0] = µu
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µv = πTemp003
							// line 116: class OtherSeq(object):
							πF.SetLineno(116)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("OtherSeq", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
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
									// line 117: def __init__(self, initseq):
									πF.SetLineno(117)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "initseq", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µinitseq *πg.Object = πArgs[1]; _ = µinitseq
										var πTemp001 *πg.Object
										_ = πTemp001
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 118: self.__data = initseq
											πF.SetLineno(118)
											if πE = πg.CheckLocal(πF, µinitseq, "initseq"); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µinitseq); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß__data, πTemp001); πE != nil {
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
									// line 119: def __len__(self):
									πF.SetLineno(119)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 120: return len(self.__data)
											πF.SetLineno(120)
											πTemp001 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µself, ß__data, nil); πE != nil {
												continue
											}
											πTemp001[0] = πTemp002
											if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp001)
											πR = πTemp003
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
									// line 121: def __getitem__(self, i):
									πF.SetLineno(121)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "i", Def: nil}
									πTemp004 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µi *πg.Object = πArgs[1]; _ = µi
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
											// line 122: return self.__data[i]
											πF.SetLineno(122)
											if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
												continue
											}
											πTemp001 = µi
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetAttr(πF, µself, ß__data, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
												continue
											}
											πR = πTemp002
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
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp003 == nil {
								πTemp003 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("OtherSeq").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µOtherSeq = πTemp006
							// line 123: s = OtherSeq(u0)
							πF.SetLineno(123)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu0, "u0"); πE != nil {
								continue
							}
							πTemp001[0] = µu0
							if πE = πg.CheckLocal(πF, µOtherSeq, "OtherSeq"); πE != nil {
								continue
							}
							if πTemp002, πE = µOtherSeq.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp002
							// line 124: v0 = self.type2test(s)
							πF.SetLineno(124)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µv0 = πTemp003
							// line 125: self.assertEqual(len(v0), len(s))
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µv0, "v0"); πE != nil {
								continue
							}
							πTemp004[0] = µv0
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp004[0] = µs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 127: s = "this is also a sequence"
							πF.SetLineno(127)
							µs = πg.NewStr("this is also a sequence").ToObject()
							// line 128: vv = self.type2test(s)
							πF.SetLineno(128)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvv = πTemp003
							// line 129: self.assertEqual(len(vv), len(s))
							πF.SetLineno(129)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvv, "vv"); πE != nil {
								continue
							}
							πTemp004[0] = µvv
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp004[0] = µs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1000).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp006 = πg.NewTuple2(ßdo.ToObject(), πg.NewFloat(1.2).ToObject()).ToObject()
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewInt(2000).ToObject()
							πTemp001[1] = πg.NewInt(2200).ToObject()
							πTemp001[2] = πg.NewInt(5).ToObject()
							if πTemp008, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πg.NewTuple5(ß123.ToObject(), ß.ToObject(), πTemp007, πTemp006, πTemp009).ToObject()
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp010 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp010 {
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
								πTemp011 = !isStop
							} else {
								πTemp011 = true
								µs = πTemp003
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πTemp007, πE = πg.ResolveGlobal(πF, ßSequence); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ßIterFunc); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßIterGen); πE != nil {
								continue
							}
							if πTemp012, πE = πg.ResolveGlobal(πF, ßiterfunc); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple4(πTemp007, πTemp008, πTemp009, πTemp012).ToObject()
							if πTemp003, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp011 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp011 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp013 = !isStop
							} else {
								πTemp013 = true
								µg = πTemp006
							}
							if πE != nil || !πTemp013 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 136: self.assertEqual(self.type2test(g(s)), self.type2test(s))
							πF.SetLineno(136)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp014 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp014[0] = µs
							if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
								continue
							}
							if πTemp006, πE = µg.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp004[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp007
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp004[0] = µs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 137: self.assertEqual(self.type2test(IterFuncStop(s)), self.type2test())
							πF.SetLineno(137)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp014 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp014[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIterFuncStop); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp004[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 138: self.assertEqual(self.type2test(c for c in "123"), self.type2test("123"))
							πF.SetLineno(138)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp015 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/seq_tests.py", πTemp015, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µc *πg.Object = πg.UnboundLocal; _ = µc
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
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
										if πTemp001, πE = πg.Iter(πF, ß123.ToObject()); πE != nil {
											continue
										}
										πF.PushCheckpoint(2)
										πTemp002 = false
									Label1:
										if πE != nil || πR != nil {
											continue
										}
										if πTemp002 {
											πF.PopCheckpoint()
											goto Label3
										}
										if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
											isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
											if exc != nil {
												πE = exc
											} else if isStop {
												πE = nil
												πF.RestoreExc(nil, nil)
											}
											πTemp003 = !isStop
										} else {
											πTemp003 = true
											µc = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 138: self.assertEqual(self.type2test(c for c in "123"), self.type2test("123"))
										πF.SetLineno(138)
										if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return µc, nil
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
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp007
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ß123.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 139: self.assertRaises(TypeError, self.type2test, IterNextOnly(s))
							πF.SetLineno(139)
							πTemp001 = πF.MakeArgs(3)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp004[0] = µs
							if πTemp006, πE = πg.ResolveGlobal(πF, ßIterNextOnly); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[2] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 140: self.assertRaises(TypeError, self.type2test, IterNoNext(s))
							πF.SetLineno(140)
							πTemp001 = πF.MakeArgs(3)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp004[0] = µs
							if πTemp006, πE = πg.ResolveGlobal(πF, ßIterNoNext); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[2] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 141: self.assertRaises(ZeroDivisionError, self.type2test, IterGenExc(s))
							πF.SetLineno(141)
							πTemp001 = πF.MakeArgs(3)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp004[0] = µs
							if πTemp006, πE = πg.ResolveGlobal(πF, ßIterGenExc); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[2] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 144: self.assertEqual(self.type2test(LyingTuple((2,))), self.type2test((1,)))
							πF.SetLineno(144)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp014 = πF.MakeArgs(1)
							πTemp002 = πg.NewTuple1(πg.NewInt(2).ToObject()).ToObject()
							πTemp014[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßLyingTuple); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp004[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp006
							πTemp004 = πF.MakeArgs(1)
							πTemp002 = πg.NewTuple1(πg.NewInt(1).ToObject()).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 145: self.assertEqual(self.type2test(LyingList([2])), self.type2test([1]))
							πF.SetLineno(145)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp014 = πF.MakeArgs(1)
							πTemp016 = make([]*πg.Object, 1)
							πTemp016[0] = πg.NewInt(2).ToObject()
							πTemp002 = πg.NewList(πTemp016...).ToObject()
							πTemp014[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßLyingList); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp004[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp006
							πTemp004 = πF.MakeArgs(1)
							πTemp014 = make([]*πg.Object, 1)
							πTemp014[0] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp014...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_constructors.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 147: def test_truth(self):
					πF.SetLineno(147)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_truth", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 148: self.assertFalse(self.type2test())
							πF.SetLineno(148)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 149: self.assertTrue(self.type2test([42]))
							πF.SetLineno(149)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(42).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_truth.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 151: def test_getitem(self):
					πF.SetLineno(151)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_getitem", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
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
							case 1: goto Label1
							case 2: goto Label2
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 152: u = self.type2test([0, 1, 2, 3, 4])
							πF.SetLineno(152)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 5)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							πTemp002[3] = πg.NewInt(3).ToObject()
							πTemp002[4] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp002[0] = µu
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp006 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µi = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 154: self.assertEqual(u[i], i)
							πF.SetLineno(154)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µu, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[1] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 155: self.assertEqual(u[long(i)], i)
							πF.SetLineno(155)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002[0] = µi
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp004 = πTemp008
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µu, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[1] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp002[0] = µu
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.Neg(πF, πTemp008); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp006 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µi = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 157: self.assertEqual(u[i], len(u)+i)
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µu, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp002[0] = µu
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp008, µi); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 158: self.assertEqual(u[long(i)], len(u)+i)
							πF.SetLineno(158)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002[0] = µi
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp004 = πTemp008
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µu, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp002[0] = µu
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp008, µi); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 159: self.assertRaises(IndexError, u.__getitem__, -len(u)-1)
							πF.SetLineno(159)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp002[0] = µu
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.Neg(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
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
							// line 160: self.assertRaises(IndexError, u.__getitem__, len(u))
							πF.SetLineno(160)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp002[0] = µu
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[2] = πTemp004
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
							// line 161: self.assertRaises(ValueError, u.__getitem__, slice(0,10,0))
							πF.SetLineno(161)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(10).ToObject()
							πTemp002[2] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[2] = πTemp004
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
							// line 163: u = self.type2test()
							πF.SetLineno(163)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µu = πTemp004
							// line 164: self.assertRaises(IndexError, u.__getitem__, 0)
							πF.SetLineno(164)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(0).ToObject()
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
							// line 165: self.assertRaises(IndexError, u.__getitem__, -1)
							πF.SetLineno(165)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
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
							// line 167: self.assertRaises(TypeError, u.__getitem__)
							πF.SetLineno(167)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
							// line 169: a = self.type2test([10, 11])
							πF.SetLineno(169)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp002[1] = πg.NewInt(11).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 170: self.assertEqual(a[0], 10)
							πF.SetLineno(170)
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 171: self.assertEqual(a[1], 11)
							πF.SetLineno(171)
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(11).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 172: self.assertEqual(a[-2], 10)
							πF.SetLineno(172)
							πTemp001 = πF.MakeArgs(2)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 173: self.assertEqual(a[-1], 11)
							πF.SetLineno(173)
							πTemp001 = πF.MakeArgs(2)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(11).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 174: self.assertRaises(IndexError, a.__getitem__, -3)
							πF.SetLineno(174)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
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
							// line 175: self.assertRaises(IndexError, a.__getitem__, 3)
							πF.SetLineno(175)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(3).ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_getitem.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 178: def test_getslice(self):
					πF.SetLineno(178)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_getslice", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 179: l = [0, 1, 2, 3, 4]
							πF.SetLineno(179)
							πTemp001 = make([]*πg.Object, 5)
							πTemp001[0] = πg.NewInt(0).ToObject()
							πTemp001[1] = πg.NewInt(1).ToObject()
							πTemp001[2] = πg.NewInt(2).ToObject()
							πTemp001[3] = πg.NewInt(3).ToObject()
							πTemp001[4] = πg.NewInt(4).ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µl = πTemp002
							// line 180: u = self.type2test(l)
							πF.SetLineno(180)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp001[0] = µl
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp003
							// line 182: self.assertEqual(u[0:0], self.type2test())
							πF.SetLineno(182)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
							// line 183: self.assertEqual(u[1:2], self.type2test([1]))
							πF.SetLineno(183)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 184: self.assertEqual(u[-2:-1], self.type2test([3]))
							πF.SetLineno(184)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πTemp006, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(3).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 185: self.assertEqual(u[-1000:1000], u)
							πF.SetLineno(185)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1000).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.NewInt(1000).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[1] = µu
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
							// line 186: self.assertEqual(u[1000:-1000], self.type2test([]))
							πF.SetLineno(186)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1000).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1000).ToObject(), πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 187: self.assertEqual(u[:], u)
							πF.SetLineno(187)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[1] = µu
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
							// line 188: self.assertEqual(u[1:None], self.type2test([1, 2, 3, 4]))
							πF.SetLineno(188)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(3).ToObject()
							πTemp005[3] = πg.NewInt(4).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 189: self.assertEqual(u[None:3], self.type2test([0, 1, 2]))
							πF.SetLineno(189)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 192: self.assertEqual(u[::], u)
							πF.SetLineno(192)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[1] = µu
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
							// line 193: self.assertEqual(u[::2], self.type2test([0, 2, 4]))
							πF.SetLineno(193)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.NewInt(2).ToObject()}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(4).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 194: self.assertEqual(u[1::2], self.type2test([1, 3]))
							πF.SetLineno(194)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.NewInt(2).ToObject()}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(3).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 195: self.assertEqual(u[::-1], self.type2test([4, 3, 2, 1, 0]))
							πF.SetLineno(195)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 5)
							πTemp005[0] = πg.NewInt(4).ToObject()
							πTemp005[1] = πg.NewInt(3).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp005[4] = πg.NewInt(0).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 196: self.assertEqual(u[::-2], self.type2test([4, 2, 0]))
							πF.SetLineno(196)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(4).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(0).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 197: self.assertEqual(u[3::-2], self.type2test([3, 1]))
							πF.SetLineno(197)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.None, πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(3).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 198: self.assertEqual(u[3:3:-2], self.type2test([]))
							πF.SetLineno(198)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject(), πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 199: self.assertEqual(u[3:2:-2], self.type2test([3]))
							πF.SetLineno(199)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.NewInt(2).ToObject(), πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(3).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 200: self.assertEqual(u[3:1:-2], self.type2test([3]))
							πF.SetLineno(200)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.NewInt(1).ToObject(), πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(3).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 201: self.assertEqual(u[3:0:-2], self.type2test([3, 1]))
							πF.SetLineno(201)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.NewInt(0).ToObject(), πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(3).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 202: self.assertEqual(u[::-100], self.type2test([4]))
							πF.SetLineno(202)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(4).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 203: self.assertEqual(u[100:-100:], self.type2test([]))
							πF.SetLineno(203)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(100).ToObject(), πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 204: self.assertEqual(u[-100:100:], u)
							πF.SetLineno(204)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.NewInt(100).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[1] = µu
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
							// line 205: self.assertEqual(u[100:-100:-1], u[::-1])
							πF.SetLineno(205)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(100).ToObject(), πTemp003, πTemp006}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πTemp003}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
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
							// line 206: self.assertEqual(u[-100:100:-1], self.type2test([]))
							πF.SetLineno(206)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.NewInt(100).ToObject(), πTemp006}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 207: self.assertEqual(u[-100L:100L:2L], self.type2test([0, 2, 4]))
							πF.SetLineno(207)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewLongFromBytes([]byte{0x64,}).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.NewLongFromBytes([]byte{0x64,}).ToObject(), πg.NewLongFromBytes([]byte{0x2,}).ToObject()}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µu, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(4).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 210: a = self.type2test([0,1,2,3,4])
							πF.SetLineno(210)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = make([]*πg.Object, 5)
							πTemp004[0] = πg.NewInt(0).ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							πTemp004[2] = πg.NewInt(2).ToObject()
							πTemp004[3] = πg.NewInt(3).ToObject()
							πTemp004[4] = πg.NewInt(4).ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp003
							// line 211: self.assertEqual(a[ -pow(2,128L): 3 ], self.type2test([0,1,2]))
							πF.SetLineno(211)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(2).ToObject()
							πTemp004[1] = πg.NewLongFromBytes([]byte{0x80,}).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßpow); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Neg(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µa, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 212: self.assertEqual(a[ 3: pow(2,145L) ], self.type2test([3,4]))
							πF.SetLineno(212)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(2).ToObject()
							πTemp004[1] = πg.NewLongFromBytes([]byte{0x91,}).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßpow); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πTemp006, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µa, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(3).ToObject()
							πTemp005[1] = πg.NewInt(4).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 214: self.assertRaises(TypeError, u.__getslice__)
							πF.SetLineno(214)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µu, ß__getslice__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_getslice.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 177: @unittest.skip('grumpy')
					πF.SetLineno(177)
					πTemp006 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_getslice); πE != nil {
						continue
					}
					πTemp006[0] = πTemp007
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = ßgrumpy.ToObject()
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp007, ßskip, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp009, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_getslice.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 216: def test_contains(self):
					πF.SetLineno(216)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_contains", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 []*πg.Object
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
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
							// line 217: u = self.type2test([0, 1, 2])
							πF.SetLineno(217)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µu); πE != nil {
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
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µi = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 219: self.assertIn(i, u)
							πF.SetLineno(219)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[1] = µu
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							if πTemp008, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.Sub(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							if πTemp009, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp008, πE = πg.Add(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp005 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µi = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 221: self.assertNotIn(i, u)
							πF.SetLineno(221)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[1] = µu
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 223: self.assertRaises(TypeError, u.__contains__)
							πF.SetLineno(223)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ß__contains__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_contains.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 225: def test_contains_fake(self):
					πF.SetLineno(225)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_contains_fake", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µAllEq *πg.Object = πg.UnboundLocal; _ = µAllEq
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 226: class AllEq(object):
							πF.SetLineno(226)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
							_, πE = πg.NewCode("AllEq", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
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
									// line 230: def __eq__(self, other):
									πF.SetLineno(230)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 231: return True
											πF.SetLineno(231)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 232: __hash__ = None # Can't meet hash invariant requirements
									πF.SetLineno(232)
									if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp003); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("AllEq").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µAllEq = πTemp005
							// line 233: self.assertNotIn(AllEq(), self.type2test([]))
							πF.SetLineno(233)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µAllEq, "AllEq"); πE != nil {
								continue
							}
							if πTemp002, πE = µAllEq.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 234: self.assertIn(AllEq(), self.type2test([1]))
							πF.SetLineno(234)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µAllEq, "AllEq"); πE != nil {
								continue
							}
							if πTemp002, πE = µAllEq.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 1)
							πTemp007[0] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_contains_fake.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 236: def test_contains_order(self):
					πF.SetLineno(236)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_contains_order", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µDoNotTestEq *πg.Object = πg.UnboundLocal; _ = µDoNotTestEq
						var µStopCompares *πg.Object = πg.UnboundLocal; _ = µStopCompares
						var µcheckfirst *πg.Object = πg.UnboundLocal; _ = µcheckfirst
						var µchecklast *πg.Object = πg.UnboundLocal; _ = µchecklast
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
							// line 240: class DoNotTestEq(Exception):
							πF.SetLineno(240)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
							_, πE = πg.NewCode("DoNotTestEq", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 241: pass
									πF.SetLineno(241)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("DoNotTestEq").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µDoNotTestEq = πTemp005
							// line 242: class StopCompares(object):
							πF.SetLineno(242)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
							_, πE = πg.NewCode("StopCompares", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
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
									// line 243: def __eq__(self, other):
									πF.SetLineno(243)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µDoNotTestEq, "DoNotTestEq"); πE != nil {
												continue
											}
											// line 244: raise DoNotTestEq
											πF.SetLineno(244)
											πE = πF.Raise(µDoNotTestEq, nil, nil)
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("StopCompares").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µStopCompares = πTemp005
							// line 246: checkfirst = self.type2test([1, StopCompares()])
							πF.SetLineno(246)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µStopCompares, "StopCompares"); πE != nil {
								continue
							}
							if πTemp002, πE = µStopCompares.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp002
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µcheckfirst = πTemp004
							// line 247: self.assertIn(1, checkfirst)
							πF.SetLineno(247)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µcheckfirst, "checkfirst"); πE != nil {
								continue
							}
							πTemp003[1] = µcheckfirst
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 248: checklast = self.type2test([StopCompares(), 1])
							πF.SetLineno(248)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µStopCompares, "StopCompares"); πE != nil {
								continue
							}
							if πTemp002, πE = µStopCompares.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							πTemp006[1] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µchecklast = πTemp004
							// line 249: self.assertRaises(DoNotTestEq, checklast.__contains__, 1)
							πF.SetLineno(249)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µDoNotTestEq, "DoNotTestEq"); πE != nil {
								continue
							}
							πTemp003[0] = µDoNotTestEq
							if πE = πg.CheckLocal(πF, µchecklast, "checklast"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µchecklast, ß__contains__, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							πTemp003[2] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_contains_order.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 251: def test_len(self):
					πF.SetLineno(251)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_len", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
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
							// line 252: self.assertEqual(len(self.type2test()), 0)
							πF.SetLineno(252)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 253: self.assertEqual(len(self.type2test([])), 0)
							πF.SetLineno(253)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp006...).ToObject()
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 254: self.assertEqual(len(self.type2test([0])), 1)
							πF.SetLineno(254)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							πTemp006[0] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp006...).ToObject()
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 255: self.assertEqual(len(self.type2test([0, 1, 2])), 3)
							πF.SetLineno(255)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 3)
							πTemp006[0] = πg.NewInt(0).ToObject()
							πTemp006[1] = πg.NewInt(1).ToObject()
							πTemp006[2] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp006...).ToObject()
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_len.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 257: def test_minmax(self):
					πF.SetLineno(257)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_minmax", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 258: u = self.type2test([0, 1, 2])
							πF.SetLineno(258)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 259: self.assertEqual(min(u), 0)
							πF.SetLineno(259)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp002[0] = µu
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 260: self.assertEqual(max(u), 2)
							πF.SetLineno(260)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp002[0] = µu
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_minmax.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 263: def test_addmul(self):
					πF.SetLineno(263)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_addmul", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu1 *πg.Object = πg.UnboundLocal; _ = µu1
						var µu2 *πg.Object = πg.UnboundLocal; _ = µu2
						var µsubclass *πg.Object = πg.UnboundLocal; _ = µsubclass
						var µu3 *πg.Object = πg.UnboundLocal; _ = µu3
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Dict
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 264: u1 = self.type2test([0])
							πF.SetLineno(264)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu1 = πTemp004
							// line 265: u2 = self.type2test([0, 1])
							πF.SetLineno(265)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu2 = πTemp004
							// line 266: self.assertEqual(u1, u1 + self.type2test())
							πF.SetLineno(266)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu1, "u1"); πE != nil {
								continue
							}
							πTemp001[0] = µu1
							if πE = πg.CheckLocal(πF, µu1, "u1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µu1, πTemp005); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 267: self.assertEqual(u1, self.type2test() + u1)
							πF.SetLineno(267)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu1, "u1"); πE != nil {
								continue
							}
							πTemp001[0] = µu1
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu1, "u1"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, µu1); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 268: self.assertEqual(u1 + self.type2test([1]), u2)
							πF.SetLineno(268)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu1, "u1"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp004 = πg.NewList(πTemp006...).ToObject()
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Add(πF, µu1, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp001[1] = µu2
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 269: self.assertEqual(self.type2test([-1]) + u1, self.type2test([-1, 0]))
							πF.SetLineno(269)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							πTemp004 = πg.NewList(πTemp006...).ToObject()
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µu1, "u1"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, µu1); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							πTemp006[1] = πg.NewInt(0).ToObject()
							πTemp003 = πg.NewList(πTemp006...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 270: self.assertEqual(self.type2test(), u2*0)
							πF.SetLineno(270)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µu2, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 271: self.assertEqual(self.type2test(), 0*u2)
							πF.SetLineno(271)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewInt(0).ToObject(), µu2); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 272: self.assertEqual(self.type2test(), u2*0L)
							πF.SetLineno(272)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µu2, πg.NewLongFromBytes([]byte{}).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 273: self.assertEqual(self.type2test(), 0L*u2)
							πF.SetLineno(273)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewLongFromBytes([]byte{}).ToObject(), µu2); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 274: self.assertEqual(u2, u2*1)
							πF.SetLineno(274)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp001[0] = µu2
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µu2, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 275: self.assertEqual(u2, 1*u2)
							πF.SetLineno(275)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp001[0] = µu2
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewInt(1).ToObject(), µu2); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 276: self.assertEqual(u2, u2*1L)
							πF.SetLineno(276)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp001[0] = µu2
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µu2, πg.NewLongFromBytes([]byte{0x1,}).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 277: self.assertEqual(u2, 1L*u2)
							πF.SetLineno(277)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							πTemp001[0] = µu2
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewLongFromBytes([]byte{0x1,}).ToObject(), µu2); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 278: self.assertEqual(u2+u2, u2*2)
							πF.SetLineno(278)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µu2, µu2); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µu2, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 279: self.assertEqual(u2+u2, 2*u2)
							πF.SetLineno(279)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µu2, µu2); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), µu2); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 280: self.assertEqual(u2+u2, u2*2L)
							πF.SetLineno(280)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µu2, µu2); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µu2, πg.NewLongFromBytes([]byte{0x2,}).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 281: self.assertEqual(u2+u2, 2L*u2)
							πF.SetLineno(281)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µu2, µu2); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewLongFromBytes([]byte{0x2,}).ToObject(), µu2); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 282: self.assertEqual(u2+u2+u2, u2*3)
							πF.SetLineno(282)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µu2, µu2); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µu2); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µu2, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 283: self.assertEqual(u2+u2+u2, 3*u2)
							πF.SetLineno(283)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µu2, µu2); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µu2); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu2, "u2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewInt(3).ToObject(), µu2); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 285: class subclass(self.type2test):
							πF.SetLineno(285)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp007 = πg.NewDict()
							if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
								continue
							}
							_, πE = πg.NewCode("subclass", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp007
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 286: pass
									πF.SetLineno(286)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("subclass").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
								continue
							}
							µsubclass = πTemp005
							// line 287: u3 = subclass([0, 1])
							πF.SetLineno(287)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							if πTemp003, πE = µsubclass.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu3 = πTemp003
							// line 288: self.assertEqual(u3, u3*1)
							πF.SetLineno(288)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu3, "u3"); πE != nil {
								continue
							}
							πTemp001[0] = µu3
							if πE = πg.CheckLocal(πF, µu3, "u3"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µu3, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 289: self.assertIsNot(u3, u3*1)
							πF.SetLineno(289)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu3, "u3"); πE != nil {
								continue
							}
							πTemp001[0] = µu3
							if πE = πg.CheckLocal(πF, µu3, "u3"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µu3, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIsNot, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_addmul.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 262: @unittest.skip('grumpy')
					πF.SetLineno(262)
					πTemp006 = πF.MakeArgs(1)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßtest_addmul); πE != nil {
						continue
					}
					πTemp006[0] = πTemp014
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = ßgrumpy.ToObject()
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßskip, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp015.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp015, πE = πTemp014.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_addmul.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 291: def test_iadd(self):
					πF.SetLineno(291)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_iadd", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 292: u = self.type2test([0, 1])
							πF.SetLineno(292)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 293: u += self.type2test()
							πF.SetLineno(293)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µu, πTemp004); πE != nil {
								continue
							}
							µu = πTemp003
							// line 294: self.assertEqual(u, self.type2test([0, 1]))
							πF.SetLineno(294)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 295: u += self.type2test([2, 3])
							πF.SetLineno(295)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(2).ToObject()
							πTemp002[1] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, µu, πTemp004); πE != nil {
								continue
							}
							µu = πTemp003
							// line 296: self.assertEqual(u, self.type2test([0, 1, 2, 3]))
							πF.SetLineno(296)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 297: u += self.type2test([4, 5])
							πF.SetLineno(297)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(4).ToObject()
							πTemp002[1] = πg.NewInt(5).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, µu, πTemp004); πE != nil {
								continue
							}
							µu = πTemp003
							// line 298: self.assertEqual(u, self.type2test([0, 1, 2, 3, 4, 5]))
							πF.SetLineno(298)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 6)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(4).ToObject()
							πTemp005[5] = πg.NewInt(5).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 300: u = self.type2test("spam")
							πF.SetLineno(300)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßspam.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 301: u += self.type2test("eggs")
							πF.SetLineno(301)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßeggs.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, µu, πTemp004); πE != nil {
								continue
							}
							µu = πTemp003
							// line 302: self.assertEqual(u, self.type2test("spameggs"))
							πF.SetLineno(302)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßspameggs.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_iadd.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 304: def test_imul(self):
					πF.SetLineno(304)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_imul", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 305: u = self.type2test([0, 1])
							πF.SetLineno(305)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 306: u *= 3
							πF.SetLineno(306)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IMul(πF, µu, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							µu = πTemp003
							// line 307: self.assertEqual(u, self.type2test([0, 1, 0, 1, 0, 1]))
							πF.SetLineno(307)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							πTemp001[0] = µu
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 6)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							πTemp005[2] = πg.NewInt(0).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp005[4] = πg.NewInt(0).ToObject()
							πTemp005[5] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_imul.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 309: def test_getitemoverwriteiter(self):
					πF.SetLineno(309)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_getitemoverwriteiter", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µT *πg.Object = πg.UnboundLocal; _ = µT
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 311: class T(self.type2test):
							πF.SetLineno(311)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
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
							_, πE = πg.NewCode("T", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
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
									// line 312: def __getitem__(self, key):
									πF.SetLineno(312)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µkey *πg.Object = πArgs[1]; _ = µkey
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 []*πg.Object
										_ = πTemp002
										var πTemp003 *πg.Object
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
											// line 313: return str(key) + '!!!'
											πF.SetLineno(313)
											πTemp002 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
												continue
											}
											πTemp002[0] = µkey
											if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
												continue
											}
											if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp002)
											if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr("!!!").ToObject()); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("T").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µT = πTemp005
							// line 314: self.assertEqual(iter(T((1,2))).next(), 1)
							πF.SetLineno(314)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							πTemp002 = πg.NewTuple2(πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µT, "T"); πE != nil {
								continue
							}
							if πTemp002, πE = µT.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßnext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(1).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_getitemoverwriteiter.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 317: def test_repeat(self):
					πF.SetLineno(317)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("test_repeat", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µn *πg.Object = πg.UnboundLocal; _ = µn
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
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
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(4).ToObject()
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
								µm = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 319: s = tuple(range(m))
							πF.SetLineno(319)
							πTemp002 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp007[0] = µm
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp004
							πTemp002 = πF.MakeArgs(2)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewInt(5).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp006 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
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
								µn = πTemp004
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 321: self.assertEqual(self.type2test(s*n), self.type2test(s)*n)
							πF.SetLineno(321)
							πTemp002 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, µs, µn); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002[0] = πTemp008
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp007[0] = µs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp010, µn); πE != nil {
								continue
							}
							πTemp002[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 322: self.assertEqual(self.type2test(s)*(-4), self.type2test([]))
							πF.SetLineno(322)
							πTemp002 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp007[0] = µs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp008, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp007 = πF.MakeArgs(1)
							πTemp011 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp011...).ToObject()
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002[1] = πTemp004
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
							// line 323: self.assertEqual(id(s), id(s*1))
							πF.SetLineno(323)
							πTemp002 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp007[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002[0] = πTemp004
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µs, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002[1] = πTemp004
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
					if πE = πClass.SetItem(πF, ßtest_repeat.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 316: @unittest.skip('grumpy')
					πF.SetLineno(316)
					πTemp006 = πF.MakeArgs(1)
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßtest_repeat); πE != nil {
						continue
					}
					πTemp006[0] = πTemp018
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = ßgrumpy.ToObject()
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp019, πE = πg.GetAttr(πF, πTemp018, ßskip, nil); πE != nil {
						continue
					}
					if πTemp018, πE = πTemp019.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp019, πE = πTemp018.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_repeat.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 325: def test_bigrepeat(self):
					πF.SetLineno(325)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("test_bigrepeat", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsys *πg.Object = πg.UnboundLocal; _ = µsys
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 326: import sys
							πF.SetLineno(326)
							if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
								continue
							}
							πTemp001 = πTemp002[0]
							µsys = πTemp001
							if πE = πg.CheckLocal(πF, µsys, "sys"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsys, ßmaxint, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, πTemp003, πg.NewInt(2147483647).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 327: if sys.maxint <= 2147483647:
							πF.SetLineno(327)
						Label1:
							// line 328: x = self.type2test([0])
							πF.SetLineno(328)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µx = πTemp003
							// line 333: x *= 1<<16
							πF.SetLineno(333)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(16).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IMul(πF, µx, πTemp001); πE != nil {
								continue
							}
							µx = πTemp003
							// line 334: self.assertRaises(MemoryError, x.__mul__, 1<<16)
							πF.SetLineno(334)
							πTemp002 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMemoryError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µx, ß__mul__, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(16).ToObject()); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp002[0] = µx
							πTemp002[1] = ß__imul__.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 335: if hasattr(x, '__imul__'):
							πF.SetLineno(335)
						Label3:
							// line 336: self.assertRaises(MemoryError, x.__imul__, 1<<16)
							πF.SetLineno(336)
							πTemp002 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMemoryError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µx, ß__imul__, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(16).ToObject()); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label4
						Label4:
							goto Label2
						Label2:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_bigrepeat.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 338: def test_subscript(self):
					πF.SetLineno(338)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("test_subscript", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 339: a = self.type2test([10, 11])
							πF.SetLineno(339)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(10).ToObject()
							πTemp002[1] = πg.NewInt(11).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 340: self.assertEqual(a.__getitem__(0L), 10)
							πF.SetLineno(340)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewLongFromBytes([]byte{}).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 341: self.assertEqual(a.__getitem__(1L), 11)
							πF.SetLineno(341)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewLongFromBytes([]byte{0x1,}).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(11).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 342: self.assertEqual(a.__getitem__(-2L), 10)
							πF.SetLineno(342)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.Neg(πF, πg.NewLongFromBytes([]byte{0x2,}).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 343: self.assertEqual(a.__getitem__(-1L), 11)
							πF.SetLineno(343)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.Neg(πF, πg.NewLongFromBytes([]byte{0x1,}).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(11).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 344: self.assertRaises(IndexError, a.__getitem__, -3)
							πF.SetLineno(344)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
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
							// line 345: self.assertRaises(IndexError, a.__getitem__, 3)
							πF.SetLineno(345)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(3).ToObject()
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
							// line 346: self.assertEqual(a.__getitem__(slice(0,1)), self.type2test([10]))
							πF.SetLineno(346)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(10).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 347: self.assertEqual(a.__getitem__(slice(1,2)), self.type2test([11]))
							πF.SetLineno(347)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(11).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 348: self.assertEqual(a.__getitem__(slice(0,2)), self.type2test([10, 11]))
							πF.SetLineno(348)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(10).ToObject()
							πTemp005[1] = πg.NewInt(11).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 349: self.assertEqual(a.__getitem__(slice(0,3)), self.type2test([10, 11]))
							πF.SetLineno(349)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp005[1] = πg.NewInt(3).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 2)
							πTemp005[0] = πg.NewInt(10).ToObject()
							πTemp005[1] = πg.NewInt(11).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 350: self.assertEqual(a.__getitem__(slice(3,5)), self.type2test([]))
							πF.SetLineno(350)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewInt(3).ToObject()
							πTemp005[1] = πg.NewInt(5).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 351: self.assertRaises(ValueError, a.__getitem__, slice(0, 10, 0))
							πF.SetLineno(351)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(10).ToObject()
							πTemp002[2] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[2] = πTemp004
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
							// line 352: self.assertRaises(TypeError, a.__getitem__, 'x')
							πF.SetLineno(352)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = ßx.ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_subscript.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 354: def test_count(self):
					πF.SetLineno(354)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("test_count", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µBadExc *πg.Object = πg.UnboundLocal; _ = µBadExc
						var µBadCmp *πg.Object = πg.UnboundLocal; _ = µBadCmp
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Dict
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 355: a = self.type2test([0, 1, 2])*3
							πF.SetLineno(355)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = make([]*πg.Object, 3)
							πTemp003[0] = πg.NewInt(0).ToObject()
							πTemp003[1] = πg.NewInt(1).ToObject()
							πTemp003[2] = πg.NewInt(2).ToObject()
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Mul(πF, πTemp005, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							µa = πTemp001
							// line 356: self.assertEqual(a.count(0), 3)
							πF.SetLineno(356)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 357: self.assertEqual(a.count(1), 3)
							πF.SetLineno(357)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 358: self.assertEqual(a.count(3), 0)
							πF.SetLineno(358)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 360: self.assertRaises(TypeError, a.count)
							πF.SetLineno(360)
							πTemp002 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
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
							// line 362: class BadExc(Exception):
							πF.SetLineno(362)
							πTemp002 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							πTemp006 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadExc", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp006
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 363: pass
									πF.SetLineno(363)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadExc").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
								continue
							}
							µBadExc = πTemp005
							// line 365: class BadCmp(object):
							πF.SetLineno(365)
							πTemp002 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							πTemp006 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadCmp", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 366: def __eq__(self, other):
									πF.SetLineno(366)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 bool
										_ = πTemp002
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, µother, πg.NewInt(2).ToObject()); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label1
											}
											goto Label2
											// line 367: if other == 2:
											πF.SetLineno(367)
										Label1:
											if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
												continue
											}
											if πTemp001, πE = µBadExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 368: raise BadExc()
											πF.SetLineno(368)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label2
										Label2:
											// line 369: return False
											πF.SetLineno(369)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadCmp").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
								continue
							}
							µBadCmp = πTemp005
							// line 371: self.assertRaises(BadExc, a.count, BadCmp())
							πF.SetLineno(371)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
								continue
							}
							πTemp002[0] = µBadExc
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µa, ßcount, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp001, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
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
					if πE = πClass.SetItem(πF, ßtest_count.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 374: def test_index(self):
					πF.SetLineno(374)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("test_index", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µu *πg.Object = πg.UnboundLocal; _ = µu
						var µBadExc *πg.Object = πg.UnboundLocal; _ = µBadExc
						var µBadCmp *πg.Object = πg.UnboundLocal; _ = µBadCmp
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 375: u = self.type2test([0, 1])
							πF.SetLineno(375)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 376: self.assertEqual(u.index(0), 0)
							πF.SetLineno(376)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 377: self.assertEqual(u.index(1), 1)
							πF.SetLineno(377)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 378: self.assertRaises(ValueError, u.index, 2)
							πF.SetLineno(378)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(2).ToObject()
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
							// line 380: u = self.type2test([-2, -1, 0, 0, 1, 2])
							πF.SetLineno(380)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 6)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							πTemp002[2] = πg.NewInt(0).ToObject()
							πTemp002[3] = πg.NewInt(0).ToObject()
							πTemp002[4] = πg.NewInt(1).ToObject()
							πTemp002[5] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µu = πTemp004
							// line 381: self.assertEqual(u.count(0), 2)
							πF.SetLineno(381)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 382: self.assertEqual(u.index(0), 2)
							πF.SetLineno(382)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 383: self.assertEqual(u.index(0, 2), 2)
							πF.SetLineno(383)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 384: self.assertEqual(u.index(-2, -10), 0)
							πF.SetLineno(384)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 385: self.assertEqual(u.index(0, 3), 3)
							πF.SetLineno(385)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 386: self.assertEqual(u.index(0, 3, 4), 3)
							πF.SetLineno(386)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(3).ToObject()
							πTemp002[2] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 387: self.assertRaises(ValueError, u.index, 2, 0, -10)
							πF.SetLineno(387)
							πTemp001 = πF.MakeArgs(5)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(2).ToObject()
							πTemp001[3] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp003
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
							// line 389: self.assertRaises(TypeError, u.index)
							πF.SetLineno(389)
							πTemp001 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µu, "u"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µu, ßindex, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
							// line 391: class BadExc(Exception):
							πF.SetLineno(391)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadExc", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 392: pass
									πF.SetLineno(392)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadExc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µBadExc = πTemp006
							// line 394: class BadCmp(object):
							πF.SetLineno(394)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadCmp", "/usr/lib/python2.7/test/seq_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
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
									// line 395: def __eq__(self, other):
									πF.SetLineno(395)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 bool
										_ = πTemp002
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, µother, πg.NewInt(2).ToObject()); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label1
											}
											goto Label2
											// line 396: if other == 2:
											πF.SetLineno(396)
										Label1:
											if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
												continue
											}
											if πTemp001, πE = µBadExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 397: raise BadExc()
											πF.SetLineno(397)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label2
										Label2:
											// line 398: return False
											πF.SetLineno(398)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadCmp").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µBadCmp = πTemp006
							// line 400: a = self.type2test([0, 1, 2, 3])
							πF.SetLineno(400)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							πTemp002[3] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 401: self.assertRaises(BadExc, a.index, BadCmp())
							πF.SetLineno(401)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µBadExc, "BadExc"); πE != nil {
								continue
							}
							πTemp001[0] = µBadExc
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp003, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
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
							// line 403: a = self.type2test([-2, -1, 0, 0, 1, 2])
							πF.SetLineno(403)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 6)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							πTemp002[2] = πg.NewInt(0).ToObject()
							πTemp002[3] = πg.NewInt(0).ToObject()
							πTemp002[4] = πg.NewInt(1).ToObject()
							πTemp002[5] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 404: self.assertEqual(a.index(0), 2)
							πF.SetLineno(404)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 405: self.assertEqual(a.index(0, 2), 2)
							πF.SetLineno(405)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 406: self.assertEqual(a.index(0, -4), 2)
							πF.SetLineno(406)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 407: self.assertEqual(a.index(-2, -10), 0)
							πF.SetLineno(407)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 408: self.assertEqual(a.index(0, 3), 3)
							πF.SetLineno(408)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 409: self.assertEqual(a.index(0, -3), 3)
							πF.SetLineno(409)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 410: self.assertEqual(a.index(0, 3, 4), 3)
							πF.SetLineno(410)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							πTemp002[1] = πg.NewInt(3).ToObject()
							πTemp002[2] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 411: self.assertEqual(a.index(0, -3, -2), 3)
							πF.SetLineno(411)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 412: self.assertEqual(a.index(0, -4*sys.maxint, 4*sys.maxint), 2)
							πF.SetLineno(412)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = πg.NewInt(0).ToObject()
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßmaxint, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßmaxint, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewInt(4).ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 413: self.assertRaises(ValueError, a.index, 0, 4*sys.maxint,-4*sys.maxint)
							πF.SetLineno(413)
							πTemp001 = πF.MakeArgs(5)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(0).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßmaxint, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewInt(4).ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[3] = πTemp003
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßmaxint, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πTemp001[4] = πTemp003
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
							// line 414: self.assertRaises(ValueError, a.index, 2, 0, -10)
							πF.SetLineno(414)
							πTemp001 = πF.MakeArgs(5)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßindex, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(2).ToObject()
							πTemp001[3] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp003
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_index.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 373: @unittest.skip('grumpy')
					πF.SetLineno(373)
					πTemp006 = πF.MakeArgs(1)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßtest_index); πE != nil {
						continue
					}
					πTemp006[0] = πTemp022
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = ßgrumpy.ToObject()
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp023, πE = πg.GetAttr(πF, πTemp022, ßskip, nil); πE != nil {
						continue
					}
					if πTemp022, πE = πTemp023.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp023, πE = πTemp022.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_index.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 417: def test_free_after_iterating(self):
					πF.SetLineno(417)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("test_free_after_iterating", "/usr/lib/python2.7/test/seq_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 418: support.check_free_after_iterating(self, iter, self.type2test)
							πF.SetLineno(418)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsupport); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcheck_free_after_iterating, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 419: support.check_free_after_iterating(self, reversed, self.type2test)
							πF.SetLineno(419)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßreversed); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsupport); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcheck_free_after_iterating, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_free_after_iterating.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 416: @unittest.skip('grumpy')
					πF.SetLineno(416)
					πTemp006 = πF.MakeArgs(1)
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßtest_free_after_iterating); πE != nil {
						continue
					}
					πTemp006[0] = πTemp023
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = ßgrumpy.ToObject()
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp024, πE = πg.GetAttr(πF, πTemp023, ßskip, nil); πE != nil {
						continue
					}
					if πTemp023, πE = πTemp024.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp024, πE = πTemp023.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_free_after_iterating.ToObject(), πTemp024); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("CommonTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCommonTest.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("test.seq_tests", Code)
}

