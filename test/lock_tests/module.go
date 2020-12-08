package lock_tests
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/threading"
	// _ "github.com/pygolin/stdlib/test/test_support"
	// _ "github.com/pygolin/stdlib/thread"
	// _ "github.com/pygolin/stdlib/time"
	// _ "github.com/pygolin/stdlib/unittest"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBaseLockTests := πg.InternStr("BaseLockTests")
		ßBaseSemaphoreTests := πg.InternStr("BaseSemaphoreTests")
		ßBaseTestCase := πg.InternStr("BaseTestCase")
		ßBoundedSemaphoreTests := πg.InternStr("BoundedSemaphoreTests")
		ßBunch := πg.InternStr("Bunch")
		ßConditionTests := πg.InternStr("ConditionTests")
		ßEventTests := πg.InternStr("EventTests")
		ßFalse := πg.InternStr("False")
		ßLock := πg.InternStr("Lock")
		ßLockTests := πg.InternStr("LockTests")
		ßNone := πg.InternStr("None")
		ßRLockTests := πg.InternStr("RLockTests")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßSemaphoreTests := πg.InternStr("SemaphoreTests")
		ßTestCase := πg.InternStr("TestCase")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß_Condition__lock := πg.InternStr("_Condition__lock")
		ß_Event__cond := πg.InternStr("_Event__cond")
		ß__doc__ := πg.InternStr("__doc__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_can_exit := πg.InternStr("_can_exit")
		ß_check_notify := πg.InternStr("_check_notify")
		ß_is_owned := πg.InternStr("_is_owned")
		ß_reset_internal_locks := πg.InternStr("_reset_internal_locks")
		ß_threads := πg.InternStr("_threads")
		ß_wait := πg.InternStr("_wait")
		ßacquire := πg.InternStr("acquire")
		ßappend := πg.InternStr("append")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertIs := πg.InternStr("assertIs")
		ßassertIsNot := πg.InternStr("assertIsNot")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßclear := πg.InternStr("clear")
		ßcondtype := πg.InternStr("condtype")
		ßdo_finish := πg.InternStr("do_finish")
		ßenumerate := πg.InternStr("enumerate")
		ßeventtype := πg.InternStr("eventtype")
		ßf := πg.InternStr("f")
		ßfinished := πg.InternStr("finished")
		ßget_ident := πg.InternStr("get_ident")
		ßgrumpy := πg.InternStr("grumpy")
		ßis_set := πg.InternStr("is_set")
		ßlen := πg.InternStr("len")
		ßlocktype := πg.InternStr("locktype")
		ßmaxint := πg.InternStr("maxint")
		ßn := πg.InternStr("n")
		ßnotify := πg.InternStr("notify")
		ßnotify_all := πg.InternStr("notify_all")
		ßobject := πg.InternStr("object")
		ßrange := πg.InternStr("range")
		ßreap_children := πg.InternStr("reap_children")
		ßrelease := πg.InternStr("release")
		ßsemtype := πg.InternStr("semtype")
		ßset := πg.InternStr("set")
		ßsetUp := πg.InternStr("setUp")
		ßskip := πg.InternStr("skip")
		ßsleep := πg.InternStr("sleep")
		ßsorted := πg.InternStr("sorted")
		ßstart_new_thread := πg.InternStr("start_new_thread")
		ßstarted := πg.InternStr("started")
		ßsupport := πg.InternStr("support")
		ßsys := πg.InternStr("sys")
		ßtearDown := πg.InternStr("tearDown")
		ßtest__is_owned := πg.InternStr("test__is_owned")
		ßtest_acquire := πg.InternStr("test_acquire")
		ßtest_acquire_contended := πg.InternStr("test_acquire_contended")
		ßtest_acquire_destroy := πg.InternStr("test_acquire_destroy")
		ßtest_acquire_release := πg.InternStr("test_acquire_release")
		ßtest_constructor := πg.InternStr("test_constructor")
		ßtest_default_value := πg.InternStr("test_default_value")
		ßtest_different_thread := πg.InternStr("test_different_thread")
		ßtest_is_set := πg.InternStr("test_is_set")
		ßtest_notify := πg.InternStr("test_notify")
		ßtest_reacquire := πg.InternStr("test_reacquire")
		ßtest_release_unacquired := πg.InternStr("test_release_unacquired")
		ßtest_reset_internal_locks := πg.InternStr("test_reset_internal_locks")
		ßtest_thread_leak := πg.InternStr("test_thread_leak")
		ßtest_timeout := πg.InternStr("test_timeout")
		ßtest_try_acquire := πg.InternStr("test_try_acquire")
		ßtest_try_acquire_contended := πg.InternStr("test_try_acquire_contended")
		ßtest_unacquired_notify := πg.InternStr("test_unacquired_notify")
		ßtest_unacquired_wait := πg.InternStr("test_unacquired_wait")
		ßtest_with := πg.InternStr("test_with")
		ßthreading := πg.InternStr("threading")
		ßthreading_cleanup := πg.InternStr("threading_cleanup")
		ßthreading_setup := πg.InternStr("threading_setup")
		ßtime := πg.InternStr("time")
		ßtype := πg.InternStr("type")
		ßunittest := πg.InternStr("unittest")
		ßwait := πg.InternStr("wait")
		ßwait_for_finished := πg.InternStr("wait_for_finished")
		ßwait_for_started := πg.InternStr("wait_for_started")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 []πg.Param
		_ = πTemp004
		var πTemp005 *πg.Dict
		_ = πTemp005
		var πTemp006 *πg.Object
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nVarious tests for synchronization primitives.\n").ToObject()); πE != nil {
				continue
			}
			// line 5: import sys
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: import time
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: from thread import start_new_thread, get_ident
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "thread"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßstart_new_thread); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstart_new_thread.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßget_ident); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßget_ident.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 8: import threading
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "threading"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: import unittest
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: from test import test_support as support
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßsupport.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: def _wait():
			πF.SetLineno(14)
			πTemp004 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("_wait", "/usr/lib/python2.7/test/lock_tests.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 16: time.sleep(0.01)
					πF.SetLineno(16)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewFloat(0.01).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsleep, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_wait.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: class Bunch(object):
			πF.SetLineno(18)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Bunch", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp005 *πg.Object
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 19: """
					πF.SetLineno(19)
					// line 19: """
					πF.SetLineno(19)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    A bunch of threads.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 22: def __init__(self, f, n, wait_before_exit=False):
					πF.SetLineno(22)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "f", Def: nil}
					πTemp002[2] = πg.Param{Name: "n", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "wait_before_exit", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πArgs[1]; _ = µf
						var µn *πg.Object = πArgs[2]; _ = µn
						var µwait_before_exit *πg.Object = πArgs[3]; _ = µwait_before_exit
						var µtask *πg.Object = πg.UnboundLocal; _ = µtask
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []πg.Param
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 23: """
							πF.SetLineno(23)
							// line 28: self.f = f
							πF.SetLineno(28)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µf); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßf, πTemp001); πE != nil {
								continue
							}
							// line 29: self.n = n
							πF.SetLineno(29)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßn, πTemp001); πE != nil {
								continue
							}
							// line 30: self.started = []
							πF.SetLineno(30)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstarted, πTemp003); πE != nil {
								continue
							}
							// line 31: self.finished = []
							πF.SetLineno(31)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfinished, πTemp003); πE != nil {
								continue
							}
							// line 32: self._can_exit = not wait_before_exit
							πF.SetLineno(32)
							if πE = πg.CheckLocal(πF, µwait_before_exit, "wait_before_exit"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µwait_before_exit); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_can_exit, πTemp003); πE != nil {
								continue
							}
							// line 33: def task():
							πF.SetLineno(33)
							πTemp005 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("task", "/usr/lib/python2.7/test/lock_tests.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µtid *πg.Object = πg.UnboundLocal; _ = µtid
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
								var πTemp007 bool
								_ = πTemp007
								var πTemp008 bool
								_ = πTemp008
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									case 1: goto Label1
									case 2: goto Label2
									case 3: goto Label3
									default: panic("unexpected function state")
									}
									// line 34: tid = get_ident()
									πF.SetLineno(34)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßget_ident); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									µtid = πTemp002
									// line 35: self.started.append(tid)
									πF.SetLineno(35)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µtid, "tid"); πE != nil {
										continue
									}
									πTemp003[0] = µtid
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßstarted, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 36: try:
									πF.SetLineno(36)
									πF.PushCheckpoint(1)
									// line 37: f()
									πF.SetLineno(37)
									if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
										continue
									}
									if πTemp001, πE = µf.Call(πF, nil, nil); πE != nil {
										continue
									}
									πF.PopCheckpoint()
									goto Label1
								Label1:
									πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
									// line 39: self.finished.append(tid)
									πF.SetLineno(39)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µtid, "tid"); πE != nil {
										continue
									}
									πTemp003[0] = µtid
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßfinished, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 40: while not self._can_exit:
									πF.SetLineno(40)
									πF.PushCheckpoint(3)
									πTemp006 = false
								Label2:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp006 {
										πF.PopCheckpoint()
										goto Label4
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ß_can_exit, nil); πE != nil {
										continue
									}
									if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(!πTemp008).ToObject()
									if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πE != nil || !πTemp007 {
										continue
									}
									πF.PushCheckpoint(2)            
									// line 41: _wait()
									πF.SetLineno(41)
									if πTemp001, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									continue
								Label3:
									if πE != nil || πR != nil {
										continue
									}
								Label4:
									if πTemp004 != nil {
										πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
										continue
									}
									if πR != nil {
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
							µtask = πTemp001
							// line 42: try:
							πF.SetLineno(42)
							πF.PushCheckpoint(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp002[0] = µn
							if πTemp006, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp004 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µi = πTemp006
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 44: start_new_thread(task, ())
							πF.SetLineno(44)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtask, "task"); πE != nil {
								continue
							}
							πTemp002[0] = µtask
							πTemp006 = πg.NewTuple0().ToObject()
							πTemp002[1] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßstart_new_thread); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							goto Label6
							// line 45: except:
							πF.SetLineno(45)
						Label6:
							// line 46: self._can_exit = True
							πF.SetLineno(46)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_can_exit, πTemp006); πE != nil {
								continue
							}
							// line 47: raise
							πF.SetLineno(47)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
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
					// line 23: """
					πF.SetLineno(23)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("\n        Construct a bunch of `n` threads running the same function `f`.\n        If `wait_before_exit` is True, the threads won't terminate until\n        do_finish() is called.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 49: def wait_for_started(self):
					πF.SetLineno(49)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("wait_for_started", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 50: while len(self.started) < self.n:
							πF.SetLineno(50)
							πF.PushCheckpoint(2)
							πTemp001 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp001 {
								πF.PopCheckpoint()
								goto Label3
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstarted, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 51: _wait()
							πF.SetLineno(51)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßwait_for_started.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 53: def wait_for_finished(self):
					πF.SetLineno(53)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("wait_for_finished", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 54: while len(self.finished) < self.n:
							πF.SetLineno(54)
							πF.PushCheckpoint(2)
							πTemp001 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp001 {
								πF.PopCheckpoint()
								goto Label3
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßfinished, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 55: _wait()
							πF.SetLineno(55)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßwait_for_finished.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 57: def do_finish(self):
					πF.SetLineno(57)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("do_finish", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 58: self._can_exit = True
							πF.SetLineno(58)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_can_exit, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdo_finish.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Bunch").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBunch.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 61: class BaseTestCase(unittest.TestCase):
			πF.SetLineno(61)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp008
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BaseTestCase", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
					// line 62: def setUp(self):
					πF.SetLineno(62)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("setUp", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 63: self._threads = support.threading_setup()
							πF.SetLineno(63)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsupport); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßthreading_setup, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_threads, πTemp002); πE != nil {
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
					// line 65: def tearDown(self):
					πF.SetLineno(65)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("tearDown", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 66: support.threading_cleanup(*self._threads)
							πF.SetLineno(66)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_threads, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsupport); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßthreading_cleanup, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp003, nil, πTemp001, nil, nil); πE != nil {
								continue
							}
							// line 67: support.reap_children()
							πF.SetLineno(67)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsupport); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßreap_children, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtearDown.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("BaseTestCase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBaseTestCase.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 70: class BaseLockTests(BaseTestCase):
			πF.SetLineno(70)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßBaseTestCase); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BaseLockTests", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 71: """
					πF.SetLineno(71)
					// line 71: """
					πF.SetLineno(71)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Tests for both recursive and non-recursive locks.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 75: def test_constructor(self):
					πF.SetLineno(75)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_constructor", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
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
							// line 76: lock = self.locktype()
							πF.SetLineno(76)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 77: del lock
							πF.SetLineno(77)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							µlock = πg.UnboundLocal
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_constructor.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 79: def test_acquire_destroy(self):
					πF.SetLineno(79)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_acquire_destroy", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
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
							// line 80: lock = self.locktype()
							πF.SetLineno(80)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 81: lock.acquire()
							πF.SetLineno(81)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 82: del lock
							πF.SetLineno(82)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							µlock = πg.UnboundLocal
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_acquire_destroy.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 84: def test_acquire_release(self):
					πF.SetLineno(84)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_acquire_release", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
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
							// line 85: lock = self.locktype()
							πF.SetLineno(85)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 86: lock.acquire()
							πF.SetLineno(86)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 87: lock.release()
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 88: del lock
							πF.SetLineno(88)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							µlock = πg.UnboundLocal
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_acquire_release.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 90: def test_try_acquire(self):
					πF.SetLineno(90)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_try_acquire", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
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
							// line 91: lock = self.locktype()
							πF.SetLineno(91)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 92: self.assertTrue(lock.acquire(False))
							πF.SetLineno(92)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 93: lock.release()
							πF.SetLineno(93)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_try_acquire.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 95: def test_try_acquire_contended(self):
					πF.SetLineno(95)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_try_acquire_contended", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 96: lock = self.locktype()
							πF.SetLineno(96)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 97: lock.acquire()
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 98: result = []
							πF.SetLineno(98)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µresult = πTemp001
							// line 99: def f():
							πF.SetLineno(99)
							πTemp004 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 100: result.append(lock.acquire(False))
									πF.SetLineno(100)
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πF.MakeArgs(1)
									if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									πTemp001[0] = πTemp004
									if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
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
							µf = πTemp001
							// line 101: Bunch(f, 1).wait_for_finished()
							πF.SetLineno(101)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[0] = µf
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 102: self.assertFalse(result[0])
							πF.SetLineno(102)
							πTemp003 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µresult, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 103: lock.release()
							πF.SetLineno(103)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_try_acquire_contended.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 105: def test_acquire_contended(self):
					πF.SetLineno(105)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_acquire_contended", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var µN *πg.Object = πg.UnboundLocal; _ = µN
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 106: lock = self.locktype()
							πF.SetLineno(106)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 107: lock.acquire()
							πF.SetLineno(107)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 108: N = 5
							πF.SetLineno(108)
							µN = πg.NewInt(5).ToObject()
							// line 109: def f():
							πF.SetLineno(109)
							πTemp003 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 110: lock.acquire()
									πF.SetLineno(110)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 111: lock.release()
									πF.SetLineno(111)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
							µf = πTemp001
							// line 113: b = Bunch(f, N)
							πF.SetLineno(113)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp004[0] = µf
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							πTemp004[1] = µN
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µb = πTemp005
							// line 114: b.wait_for_started()
							πF.SetLineno(114)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßwait_for_started, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 115: _wait()
							πF.SetLineno(115)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 116: self.assertEqual(len(b.finished), 0)
							πF.SetLineno(116)
							πTemp004 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßfinished, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp005
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 117: lock.release()
							πF.SetLineno(117)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 118: b.wait_for_finished()
							πF.SetLineno(118)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 119: self.assertEqual(len(b.finished), N)
							πF.SetLineno(119)
							πTemp004 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßfinished, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp005
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							πTemp004[1] = µN
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_acquire_contended.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 121: def test_with(self):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_with", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µ_with *πg.Object = πg.UnboundLocal; _ = µ_with
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 122: lock = self.locktype()
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 123: def f():
							πF.SetLineno(123)
							πTemp003 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 124: lock.acquire()
									πF.SetLineno(124)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 125: lock.release()
									πF.SetLineno(125)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
							µf = πTemp001
							// line 126: def _with(err=None):
							πF.SetLineno(126)
							πTemp003 = make([]πg.Param, 1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πg.Param{Name: "err", Def: πTemp004}
							πTemp002 = πg.NewFunction(πg.NewCode("_with", "/usr/lib/python2.7/test/lock_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µerr *πg.Object = πArgs[0]; _ = µerr
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								var πTemp005 bool
								_ = πTemp005
								var πTemp006 *πg.BaseException
								_ = πTemp006
								var πTemp007 *πg.Traceback
								_ = πTemp007
								var πTemp008 *πg.Type
								_ = πTemp008
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									case 1: goto Label1
									default: panic("unexpected function state")
									}
									// line 127: with lock:
									πF.SetLineno(127)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock.Type().ToObject(), ß__exit__, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µlock.Type().ToObject(), ß__enter__, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp002.Call(πF, πg.Args{µlock}, nil); πE != nil {
										continue
									}
									πF.PushCheckpoint(1)
									if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp003 = πg.GetBool(µerr != πTemp004).ToObject()
									if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									if πTemp005 {
										goto Label2
									}
									goto Label3
									// line 128: if err is not None:
									πF.SetLineno(128)
								Label2:
									if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
										continue
									}
									// line 129: raise err
									πF.SetLineno(129)
									πE = πF.Raise(µerr, nil, nil)
									continue
									goto Label3
								Label3:
									πF.PopCheckpoint()
								Label1:
									πTemp006, πTemp007 = nil, nil
									if πE != nil {
										πTemp006, πTemp007 = πF.ExcInfo()
									}
									if πTemp006 != nil {
										πTemp008 = πTemp006.Type()
										if πTemp003, πE = πTemp001.Call(πF, πg.Args{µlock, πTemp008.ToObject(), πTemp006.ToObject(), πTemp007.ToObject()}, nil); πE != nil {
											continue
										}
									} else {
										if πTemp003, πE = πTemp001.Call(πF, πg.Args{µlock, πg.None, πg.None, πg.None}, nil); πE != nil {
											continue
										}
									}
									if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									if πTemp006 != nil && πTemp005 != true {
										πE = πF.Raise(nil, nil, nil)
										continue
									}
									if πR != nil {
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
							µ_with = πTemp002
							// line 130: _with()
							πF.SetLineno(130)
							if πE = πg.CheckLocal(πF, µ_with, "_with"); πE != nil {
								continue
							}
							if πTemp004, πE = µ_with.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 132: Bunch(f, 1).wait_for_finished()
							πF.SetLineno(132)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp005[0] = µf
							πTemp005[1] = πg.NewInt(1).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 133: self.assertRaises(TypeError, _with, TypeError)
							πF.SetLineno(133)
							πTemp005 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µ_with, "_with"); πE != nil {
								continue
							}
							πTemp005[1] = µ_with
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp005[2] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 135: Bunch(f, 1).wait_for_finished()
							πF.SetLineno(135)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp005[0] = µf
							πTemp005[1] = πg.NewInt(1).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.GetAttr(πF, πTemp006, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_with.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 137: def test_thread_leak(self):
					πF.SetLineno(137)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_thread_leak", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 140: lock = self.locktype()
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 141: def f():
							πF.SetLineno(141)
							πTemp003 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 142: lock.acquire()
									πF.SetLineno(142)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 143: lock.release()
									πF.SetLineno(143)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
							µf = πTemp001
							// line 144: n = len(threading.enumerate())
							πF.SetLineno(144)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßenumerate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µn = πTemp005
							// line 147: Bunch(f, 15).wait_for_finished()
							πF.SetLineno(147)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp004[0] = µf
							πTemp004[1] = πg.NewInt(15).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 148: self.assertEqual(n, len(threading.enumerate()))
							πF.SetLineno(148)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp004[0] = µn
							πTemp006 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßenumerate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[1] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_thread_leak.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("BaseLockTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBaseLockTests.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 151: class LockTests(BaseLockTests):
			πF.SetLineno(151)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßBaseLockTests); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("LockTests", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
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
					// line 152: """
					πF.SetLineno(152)
					// line 152: """
					πF.SetLineno(152)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Tests for non-recursive, weak locks\n    (which can be acquired and released from different threads).\n    ").ToObject()); πE != nil {
						continue
					}
					// line 156: def test_reacquire(self):
					πF.SetLineno(156)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_reacquire", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var µphase *πg.Object = πg.UnboundLocal; _ = µphase
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
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
							// line 158: lock = self.locktype()
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 159: phase = []
							πF.SetLineno(159)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µphase = πTemp001
							// line 160: def f():
							πF.SetLineno(160)
							πTemp004 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 []*πg.Object
								_ = πTemp003
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 161: lock.acquire()
									πF.SetLineno(161)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 162: phase.append(None)
									πF.SetLineno(162)
									πTemp003 = πF.MakeArgs(1)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp003[0] = πTemp001
									if πE = πg.CheckLocal(πF, µphase, "phase"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µphase, ßappend, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 163: lock.acquire()
									πF.SetLineno(163)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 164: phase.append(None)
									πF.SetLineno(164)
									πTemp003 = πF.MakeArgs(1)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp003[0] = πTemp001
									if πE = πg.CheckLocal(πF, µphase, "phase"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µphase, ßappend, nil); πE != nil {
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
							µf = πTemp001
							// line 165: start_new_thread(f, ())
							πF.SetLineno(165)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[0] = µf
							πTemp002 = πg.NewTuple0().ToObject()
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstart_new_thread); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 166: while len(phase) == 0:
							πF.SetLineno(166)
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
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µphase, "phase"); πE != nil {
								continue
							}
							πTemp003[0] = µphase
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Eq(πF, πTemp008, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 167: _wait()
							πF.SetLineno(167)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 168: _wait()
							πF.SetLineno(168)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 169: self.assertEqual(len(phase), 1)
							πF.SetLineno(169)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µphase, "phase"); πE != nil {
								continue
							}
							πTemp009[0] = µphase
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 170: lock.release()
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 171: while len(phase) == 1:
							πF.SetLineno(171)
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
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µphase, "phase"); πE != nil {
								continue
							}
							πTemp003[0] = µphase
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Eq(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 172: _wait()
							πF.SetLineno(172)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 173: self.assertEqual(len(phase), 2)
							πF.SetLineno(173)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µphase, "phase"); πE != nil {
								continue
							}
							πTemp009[0] = µphase
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_reacquire.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 175: def test_different_thread(self):
					πF.SetLineno(175)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_different_thread", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 177: lock = self.locktype()
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 178: lock.acquire()
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 179: def f():
							πF.SetLineno(179)
							πTemp003 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 180: lock.release()
									πF.SetLineno(180)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
							µf = πTemp001
							// line 181: b = Bunch(f, 1)
							πF.SetLineno(181)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp004[0] = µf
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µb = πTemp005
							// line 182: b.wait_for_finished()
							πF.SetLineno(182)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 183: lock.acquire()
							πF.SetLineno(183)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 184: lock.release()
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_different_thread.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("LockTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLockTests.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 187: class RLockTests(BaseLockTests):
			πF.SetLineno(187)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßBaseLockTests); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("RLockTests", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp005 *πg.Object
				_ = πTemp005
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 188: """
					πF.SetLineno(188)
					// line 188: """
					πF.SetLineno(188)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Tests for recursive locks.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 191: def test_reacquire(self):
					πF.SetLineno(191)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_reacquire", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
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
							// line 192: lock = self.locktype()
							πF.SetLineno(192)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 193: lock.acquire()
							πF.SetLineno(193)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 194: lock.acquire()
							πF.SetLineno(194)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 195: lock.release()
							πF.SetLineno(195)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 196: lock.acquire()
							πF.SetLineno(196)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 197: lock.release()
							πF.SetLineno(197)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 198: lock.release()
							πF.SetLineno(198)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_reacquire.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 200: def test_release_unacquired(self):
					πF.SetLineno(200)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_release_unacquired", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 202: lock = self.locktype()
							πF.SetLineno(202)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 203: self.assertRaises(RuntimeError, lock.release)
							πF.SetLineno(203)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 204: lock.acquire()
							πF.SetLineno(204)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 205: lock.acquire()
							πF.SetLineno(205)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 206: lock.release()
							πF.SetLineno(206)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 207: lock.acquire()
							πF.SetLineno(207)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 208: lock.release()
							πF.SetLineno(208)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 209: lock.release()
							πF.SetLineno(209)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 210: self.assertRaises(RuntimeError, lock.release)
							πF.SetLineno(210)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_release_unacquired.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 212: def test_different_thread(self):
					πF.SetLineno(212)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_different_thread", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 214: lock = self.locktype()
							πF.SetLineno(214)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 215: def f():
							πF.SetLineno(215)
							πTemp003 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 216: lock.acquire()
									πF.SetLineno(216)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
							µf = πTemp001
							// line 217: b = Bunch(f, 1, True)
							πF.SetLineno(217)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp004[0] = µf
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp004[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µb = πTemp005
							// line 218: try:
							πF.SetLineno(218)
							πF.PushCheckpoint(1)
							// line 219: self.assertRaises(RuntimeError, lock.release)
							πF.SetLineno(219)
							πTemp004 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							// line 221: b.do_finish()
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßdo_finish, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006 != nil {
								πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ßtest_different_thread.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 223: def test__is_owned(self):
					πF.SetLineno(223)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test__is_owned", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 224: lock = self.locktype()
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocktype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp002
							// line 225: self.assertFalse(lock._is_owned())
							πF.SetLineno(225)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ß_is_owned, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 226: lock.acquire()
							πF.SetLineno(226)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 227: self.assertTrue(lock._is_owned())
							πF.SetLineno(227)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ß_is_owned, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 228: lock.acquire()
							πF.SetLineno(228)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 229: self.assertTrue(lock._is_owned())
							πF.SetLineno(229)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ß_is_owned, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 230: result = []
							πF.SetLineno(230)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µresult = πTemp001
							// line 231: def f():
							πF.SetLineno(231)
							πTemp004 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 232: result.append(lock._is_owned())
									πF.SetLineno(232)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µlock, ß_is_owned, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp003
									if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
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
							µf = πTemp001
							// line 233: Bunch(f, 1).wait_for_finished()
							πF.SetLineno(233)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[0] = µf
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 234: self.assertFalse(result[0])
							πF.SetLineno(234)
							πTemp003 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µresult, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 235: lock.release()
							πF.SetLineno(235)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 236: self.assertTrue(lock._is_owned())
							πF.SetLineno(236)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ß_is_owned, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 237: lock.release()
							πF.SetLineno(237)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 238: self.assertFalse(lock._is_owned())
							πF.SetLineno(238)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ß_is_owned, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest__is_owned.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("RLockTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRLockTests.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 241: class EventTests(BaseTestCase):
			πF.SetLineno(241)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßBaseTestCase); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("EventTests", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 242: """
					πF.SetLineno(242)
					// line 242: """
					πF.SetLineno(242)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Tests for Event objects.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 246: def test_is_set(self):
					πF.SetLineno(246)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_is_set", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µevt *πg.Object = πg.UnboundLocal; _ = µevt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 247: evt = self.eventtype()
							πF.SetLineno(247)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßeventtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µevt = πTemp002
							// line 248: self.assertFalse(evt.is_set())
							πF.SetLineno(248)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 249: evt.set()
							πF.SetLineno(249)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 250: self.assertTrue(evt.is_set())
							πF.SetLineno(250)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 251: evt.set()
							πF.SetLineno(251)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 252: self.assertTrue(evt.is_set())
							πF.SetLineno(252)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 253: evt.clear()
							πF.SetLineno(253)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßclear, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 254: self.assertFalse(evt.is_set())
							πF.SetLineno(254)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 255: evt.clear()
							πF.SetLineno(255)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßclear, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 256: self.assertFalse(evt.is_set())
							πF.SetLineno(256)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_is_set.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 258: def _check_notify(self, evt):
					πF.SetLineno(258)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "evt", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_check_notify", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µevt *πg.Object = πArgs[1]; _ = µevt
						var µN *πg.Object = πg.UnboundLocal; _ = µN
						var µresults1 *πg.Object = πg.UnboundLocal; _ = µresults1
						var µresults2 *πg.Object = πg.UnboundLocal; _ = µresults2
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
							// line 260: N = 5
							πF.SetLineno(260)
							µN = πg.NewInt(5).ToObject()
							// line 261: results1 = []
							πF.SetLineno(261)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresults1 = πTemp002
							// line 262: results2 = []
							πF.SetLineno(262)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresults2 = πTemp002
							// line 263: def f():
							πF.SetLineno(263)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 264: results1.append(evt.wait())
									πF.SetLineno(264)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µevt, ßwait, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp003
									if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µresults1, ßappend, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 265: results2.append(evt.wait())
									πF.SetLineno(265)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µevt, ßwait, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp003
									if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µresults2, ßappend, nil); πE != nil {
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
							µf = πTemp002
							// line 266: b = Bunch(f, N)
							πF.SetLineno(266)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[0] = µf
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							πTemp001[1] = µN
							if πTemp004, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µb = πTemp005
							// line 267: b.wait_for_started()
							πF.SetLineno(267)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µb, ßwait_for_started, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 268: _wait()
							πF.SetLineno(268)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 269: self.assertEqual(len(results1), 0)
							πF.SetLineno(269)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp006[0] = µresults1
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp005
							πTemp001[1] = πg.NewInt(0).ToObject()
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
							// line 270: evt.set()
							πF.SetLineno(270)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µevt, ßset, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 271: b.wait_for_finished()
							πF.SetLineno(271)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µb, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 272: self.assertEqual(results1, [True] * N)
							πF.SetLineno(272)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							πTemp006 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							πTemp005 = πg.NewList(πTemp006...).ToObject()
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, µN); πE != nil {
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
							// line 273: self.assertEqual(results2, [True] * N)
							πF.SetLineno(273)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							πTemp001[0] = µresults2
							πTemp006 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							πTemp005 = πg.NewList(πTemp006...).ToObject()
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, µN); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_check_notify.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 275: def test_notify(self):
					πF.SetLineno(275)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_notify", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µevt *πg.Object = πg.UnboundLocal; _ = µevt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 276: evt = self.eventtype()
							πF.SetLineno(276)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßeventtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µevt = πTemp002
							// line 277: self._check_notify(evt)
							πF.SetLineno(277)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							πTemp003[0] = µevt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_check_notify, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 279: evt.set()
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 280: evt.clear()
							πF.SetLineno(280)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ßclear, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 281: self._check_notify(evt)
							πF.SetLineno(281)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							πTemp003[0] = µevt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_check_notify, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_notify.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 283: def test_timeout(self):
					πF.SetLineno(283)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_timeout", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µevt *πg.Object = πg.UnboundLocal; _ = µevt
						var µresults1 *πg.Object = πg.UnboundLocal; _ = µresults1
						var µresults2 *πg.Object = πg.UnboundLocal; _ = µresults2
						var µN *πg.Object = πg.UnboundLocal; _ = µN
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var µdt *πg.Object = πg.UnboundLocal; _ = µdt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							// line 284: evt = self.eventtype()
							πF.SetLineno(284)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßeventtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µevt = πTemp002
							// line 285: results1 = []
							πF.SetLineno(285)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µresults1 = πTemp001
							// line 286: results2 = []
							πF.SetLineno(286)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µresults2 = πTemp001
							// line 287: N = 5
							πF.SetLineno(287)
							µN = πg.NewInt(5).ToObject()
							// line 288: def f():
							πF.SetLineno(288)
							πTemp004 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µt1 *πg.Object = πg.UnboundLocal; _ = µt1
								var µr *πg.Object = πg.UnboundLocal; _ = µr
								var µt2 *πg.Object = πg.UnboundLocal; _ = µt2
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
									// line 289: results1.append(evt.wait(0.0))
									πF.SetLineno(289)
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πF.MakeArgs(1)
									πTemp002[0] = πg.NewFloat(0.0).ToObject()
									if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µevt, ßwait, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									πTemp001[0] = πTemp004
									if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µresults1, ßappend, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 290: t1 = time.time()
									πF.SetLineno(290)
									if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
										continue
									}
									µt1 = πTemp003
									// line 291: r = evt.wait(0.2)
									πF.SetLineno(291)
									πTemp001 = πF.MakeArgs(1)
									πTemp001[0] = πg.NewFloat(0.2).ToObject()
									if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µevt, ßwait, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µr = πTemp004
									// line 292: t2 = time.time()
									πF.SetLineno(292)
									if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
										continue
									}
									µt2 = πTemp003
									// line 293: results2.append((r, t2 - t1))
									πF.SetLineno(293)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µt2, "t2"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µt1, "t1"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.Sub(πF, µt2, µt1); πE != nil {
										continue
									}
									πTemp003 = πg.NewTuple2(µr, πTemp004).ToObject()
									πTemp001[0] = πTemp003
									if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µresults2, ßappend, nil); πE != nil {
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
							µf = πTemp001
							// line 294: Bunch(f, N).wait_for_finished()
							πF.SetLineno(294)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[0] = µf
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							πTemp003[1] = µN
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 295: self.assertEqual(results1, [False] * N)
							πF.SetLineno(295)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp003[0] = µresults1
							πTemp006 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							πTemp005 = πg.NewList(πTemp006...).ToObject()
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp005, µN); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µresults2); πE != nil {
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
							if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, πTemp005); πE != nil {
									continue
								}
								µr = πTemp009
								µdt = πTemp010
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 297: self.assertFalse(r)
							πF.SetLineno(297)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							πTemp003[0] = µr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 298: self.assertTrue(dt >= 0.2, dt)
							πF.SetLineno(298)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GE(πF, µdt, πg.NewFloat(0.2).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							πTemp003[1] = µdt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 300: results1 = []
							πF.SetLineno(300)
							πTemp003 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp003...).ToObject()
							µresults1 = πTemp002
							// line 301: results2 = []
							πF.SetLineno(301)
							πTemp003 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp003...).ToObject()
							µresults2 = πTemp002
							// line 302: evt.set()
							πF.SetLineno(302)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µevt, ßset, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 303: Bunch(f, N).wait_for_finished()
							πF.SetLineno(303)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[0] = µf
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							πTemp003[1] = µN
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 304: self.assertEqual(results1, [True] * N)
							πF.SetLineno(304)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp003[0] = µresults1
							πTemp006 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp006[0] = πTemp005
							πTemp005 = πg.NewList(πTemp006...).ToObject()
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp005, µN); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µresults2); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp007 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, πTemp005); πE != nil {
									continue
								}
								µr = πTemp009
								µdt = πTemp010
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 306: self.assertTrue(r)
							πF.SetLineno(306)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							πTemp003[0] = µr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_timeout.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 309: def test_reset_internal_locks(self):
					πF.SetLineno(309)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_reset_internal_locks", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µevt *πg.Object = πg.UnboundLocal; _ = µevt
						var µold_lock *πg.Object = πg.UnboundLocal; _ = µold_lock
						var µnew_lock *πg.Object = πg.UnboundLocal; _ = µnew_lock
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
							// line 310: evt = self.eventtype()
							πF.SetLineno(310)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßeventtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µevt = πTemp002
							// line 311: old_lock = evt._Event__cond._Condition__lock
							πF.SetLineno(311)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ß_Event__cond, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_Condition__lock, nil); πE != nil {
								continue
							}
							µold_lock = πTemp002
							// line 312: evt._reset_internal_locks()
							πF.SetLineno(312)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ß_reset_internal_locks, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 313: new_lock = evt._Event__cond._Condition__lock
							πF.SetLineno(313)
							if πE = πg.CheckLocal(πF, µevt, "evt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µevt, ß_Event__cond, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_Condition__lock, nil); πE != nil {
								continue
							}
							µnew_lock = πTemp002
							// line 314: self.assertIsNot(new_lock, old_lock)
							πF.SetLineno(314)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnew_lock, "new_lock"); πE != nil {
								continue
							}
							πTemp003[0] = µnew_lock
							if πE = πg.CheckLocal(πF, µold_lock, "old_lock"); πE != nil {
								continue
							}
							πTemp003[1] = µold_lock
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertIsNot, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 315: self.assertIs(type(new_lock), type(old_lock))
							πF.SetLineno(315)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnew_lock, "new_lock"); πE != nil {
								continue
							}
							πTemp004[0] = µnew_lock
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µold_lock, "old_lock"); πE != nil {
								continue
							}
							πTemp004[0] = µold_lock
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_reset_internal_locks.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 308: @unittest.skip('grumpy')
					πF.SetLineno(308)
					πTemp007 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßtest_reset_internal_locks); πE != nil {
						continue
					}
					πTemp007[0] = πTemp008
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = ßgrumpy.ToObject()
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßskip, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp010.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp010, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πClass.SetItem(πF, ßtest_reset_internal_locks.ToObject(), πTemp010); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("EventTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßEventTests.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 318: class ConditionTests(BaseTestCase):
			πF.SetLineno(318)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßBaseTestCase); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ConditionTests", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 319: """
					πF.SetLineno(319)
					// line 319: """
					πF.SetLineno(319)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Tests for condition variables.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 323: def test_acquire(self):
					πF.SetLineno(323)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_acquire", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcond *πg.Object = πg.UnboundLocal; _ = µcond
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πTemp010 *πg.Type
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 324: cond = self.condtype()
							πF.SetLineno(324)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcondtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcond = πTemp002
							// line 327: cond.acquire()
							πF.SetLineno(327)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 328: cond.acquire()
							πF.SetLineno(328)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 329: cond.release()
							πF.SetLineno(329)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 330: cond.release()
							πF.SetLineno(330)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 331: lock = threading.Lock()
							πF.SetLineno(331)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßLock, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp001
							// line 332: cond = self.condtype(lock)
							πF.SetLineno(332)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							πTemp003[0] = µlock
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcondtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µcond = πTemp002
							// line 333: cond.acquire()
							πF.SetLineno(333)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 334: self.assertFalse(lock.acquire(False))
							πF.SetLineno(334)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 335: cond.release()
							πF.SetLineno(335)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 336: self.assertTrue(lock.acquire(False))
							πF.SetLineno(336)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 337: self.assertFalse(cond.acquire(False))
							πF.SetLineno(337)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 338: lock.release()
							πF.SetLineno(338)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 339: with cond:
							πF.SetLineno(339)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcond.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp002.Call(πF, πg.Args{µcond}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 340: self.assertFalse(lock.acquire(False))
							πF.SetLineno(340)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.PopCheckpoint()
						Label1:
							πTemp008, πTemp009 = nil, nil
							if πE != nil {
								πTemp008, πTemp009 = πF.ExcInfo()
							}
							if πTemp008 != nil {
								πTemp010 = πTemp008.Type()
								if πTemp005, πE = πTemp001.Call(πF, πg.Args{µcond, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp005, πE = πTemp001.Call(πF, πg.Args{µcond, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp008 != nil && πTemp007 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ßtest_acquire.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 342: def test_unacquired_wait(self):
					πF.SetLineno(342)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_unacquired_wait", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcond *πg.Object = πg.UnboundLocal; _ = µcond
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 343: cond = self.condtype()
							πF.SetLineno(343)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcondtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcond = πTemp002
							// line 344: self.assertRaises(RuntimeError, cond.wait)
							πF.SetLineno(344)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond, ßwait, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_unacquired_wait.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 346: def test_unacquired_notify(self):
					πF.SetLineno(346)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_unacquired_notify", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcond *πg.Object = πg.UnboundLocal; _ = µcond
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 347: cond = self.condtype()
							πF.SetLineno(347)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcondtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcond = πTemp002
							// line 348: self.assertRaises(RuntimeError, cond.notify)
							πF.SetLineno(348)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcond, ßnotify, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_unacquired_notify.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 350: def _check_notify(self, cond):
					πF.SetLineno(350)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "cond", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_check_notify", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcond *πg.Object = πArgs[1]; _ = µcond
						var µN *πg.Object = πg.UnboundLocal; _ = µN
						var µready *πg.Object = πg.UnboundLocal; _ = µready
						var µresults1 *πg.Object = πg.UnboundLocal; _ = µresults1
						var µresults2 *πg.Object = πg.UnboundLocal; _ = µresults2
						var µphase_num *πg.Object = πg.UnboundLocal; _ = µphase_num
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
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
							case 7: goto Label7
							case 8: goto Label8
							case 10: goto Label10
							case 11: goto Label11
							case 13: goto Label13
							case 14: goto Label14
							case 16: goto Label16
							case 17: goto Label17
							default: panic("unexpected function state")
							}
							// line 362: N = 5
							πF.SetLineno(362)
							µN = πg.NewInt(5).ToObject()
							// line 363: ready = []
							πF.SetLineno(363)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µready = πTemp002
							// line 364: results1 = []
							πF.SetLineno(364)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresults1 = πTemp002
							// line 365: results2 = []
							πF.SetLineno(365)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresults2 = πTemp002
							// line 366: phase_num = 0
							πF.SetLineno(366)
							µphase_num = πg.NewInt(0).ToObject()
							// line 367: def f():
							πF.SetLineno(367)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 []*πg.Object
								_ = πTemp003
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 368: cond.acquire()
									πF.SetLineno(368)
									if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 369: ready.append(phase_num)
									πF.SetLineno(369)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µphase_num, "phase_num"); πE != nil {
										continue
									}
									πTemp003[0] = µphase_num
									if πE = πg.CheckLocal(πF, µready, "ready"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µready, ßappend, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 370: cond.wait()
									πF.SetLineno(370)
									if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcond, ßwait, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 371: cond.release()
									πF.SetLineno(371)
									if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcond, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 372: results1.append(phase_num)
									πF.SetLineno(372)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µphase_num, "phase_num"); πE != nil {
										continue
									}
									πTemp003[0] = µphase_num
									if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µresults1, ßappend, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 373: cond.acquire()
									πF.SetLineno(373)
									if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 374: ready.append(phase_num)
									πF.SetLineno(374)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µphase_num, "phase_num"); πE != nil {
										continue
									}
									πTemp003[0] = µphase_num
									if πE = πg.CheckLocal(πF, µready, "ready"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µready, ßappend, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 375: cond.wait()
									πF.SetLineno(375)
									if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcond, ßwait, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 376: cond.release()
									πF.SetLineno(376)
									if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcond, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 377: results2.append(phase_num)
									πF.SetLineno(377)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µphase_num, "phase_num"); πE != nil {
										continue
									}
									πTemp003[0] = µphase_num
									if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µresults2, ßappend, nil); πE != nil {
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
							µf = πTemp002
							// line 378: b = Bunch(f, N)
							πF.SetLineno(378)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[0] = µf
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							πTemp001[1] = µN
							if πTemp004, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µb = πTemp005
							// line 379: b.wait_for_started()
							πF.SetLineno(379)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µb, ßwait_for_started, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 382: while len(ready) < 5:
							πF.SetLineno(382)
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µready, "ready"); πE != nil {
								continue
							}
							πTemp001[0] = µready
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.LT(πF, πTemp008, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 383: _wait()
							πF.SetLineno(383)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 384: ready = []
							πF.SetLineno(384)
							πTemp001 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp001...).ToObject()
							µready = πTemp004
							// line 385: self.assertEqual(results1, [])
							πF.SetLineno(385)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							πTemp009 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp009...).ToObject()
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
							// line 387: cond.acquire()
							πF.SetLineno(387)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 388: cond.notify(3)
							πF.SetLineno(388)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcond, ßnotify, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 389: _wait()
							πF.SetLineno(389)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 390: phase_num = 1
							πF.SetLineno(390)
							µphase_num = πg.NewInt(1).ToObject()
							// line 391: cond.release()
							πF.SetLineno(391)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcond, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 392: while len(results1) < 3:
							πF.SetLineno(392)
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.LT(πF, πTemp008, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 393: _wait()
							πF.SetLineno(393)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 394: self.assertEqual(results1, [1] * 3)
							πF.SetLineno(394)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = πg.NewInt(1).ToObject()
							πTemp005 = πg.NewList(πTemp009...).ToObject()
							if πTemp004, πE = πg.Mul(πF, πTemp005, πg.NewInt(3).ToObject()); πE != nil {
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
							// line 395: self.assertEqual(results2, [])
							πF.SetLineno(395)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							πTemp001[0] = µresults2
							πTemp009 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp009...).ToObject()
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
							// line 397: while len(ready) < 3:
							πF.SetLineno(397)
							πF.PushCheckpoint(8)
							πTemp006 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label9
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µready, "ready"); πE != nil {
								continue
							}
							πTemp001[0] = µready
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.LT(πF, πTemp008, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 398: _wait()
							πF.SetLineno(398)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							// line 400: cond.acquire()
							πF.SetLineno(400)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 401: cond.notify(5)
							πF.SetLineno(401)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcond, ßnotify, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 402: _wait()
							πF.SetLineno(402)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 403: phase_num = 2
							πF.SetLineno(403)
							µphase_num = πg.NewInt(2).ToObject()
							// line 404: cond.release()
							πF.SetLineno(404)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcond, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 405: while len(results1) + len(results2) < 8:
							πF.SetLineno(405)
							πF.PushCheckpoint(11)
							πTemp006 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label12
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							πTemp001[0] = µresults2
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.Add(πF, πTemp010, πTemp011); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, πTemp005, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(10)            
							// line 406: _wait()
							πF.SetLineno(406)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							// line 407: self.assertEqual(results1, [1] * 3 + [2] * 2)
							πF.SetLineno(407)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = πg.NewInt(1).ToObject()
							πTemp008 = πg.NewList(πTemp009...).ToObject()
							if πTemp005, πE = πg.Mul(πF, πTemp008, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = πg.NewInt(2).ToObject()
							πTemp010 = πg.NewList(πTemp009...).ToObject()
							if πTemp008, πE = πg.Mul(πF, πTemp010, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
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
							// line 408: self.assertEqual(results2, [2] * 3)
							πF.SetLineno(408)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							πTemp001[0] = µresults2
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = πg.NewInt(2).ToObject()
							πTemp005 = πg.NewList(πTemp009...).ToObject()
							if πTemp004, πE = πg.Mul(πF, πTemp005, πg.NewInt(3).ToObject()); πE != nil {
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
							// line 410: while len(ready) < 5:
							πF.SetLineno(410)
							πF.PushCheckpoint(14)
							πTemp006 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label15
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µready, "ready"); πE != nil {
								continue
							}
							πTemp001[0] = µready
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.LT(πF, πTemp008, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 411: _wait()
							πF.SetLineno(411)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							// line 413: cond.acquire()
							πF.SetLineno(413)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 414: cond.notify_all()
							πF.SetLineno(414)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcond, ßnotify_all, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 415: _wait()
							πF.SetLineno(415)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 416: phase_num = 3
							πF.SetLineno(416)
							µphase_num = πg.NewInt(3).ToObject()
							// line 417: cond.release()
							πF.SetLineno(417)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcond, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 418: while len(results2) < 5:
							πF.SetLineno(418)
							πF.PushCheckpoint(17)
							πTemp006 = false
						Label16:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label18
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							πTemp001[0] = µresults2
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.LT(πF, πTemp008, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(16)            
							// line 419: _wait()
							πF.SetLineno(419)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label17:
							if πE != nil || πR != nil {
								continue
							}
						Label18:
							// line 420: self.assertEqual(results1, [1] * 3 + [2] * 2)
							πF.SetLineno(420)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = πg.NewInt(1).ToObject()
							πTemp008 = πg.NewList(πTemp009...).ToObject()
							if πTemp005, πE = πg.Mul(πF, πTemp008, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = πg.NewInt(2).ToObject()
							πTemp010 = πg.NewList(πTemp009...).ToObject()
							if πTemp008, πE = πg.Mul(πF, πTemp010, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
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
							// line 421: self.assertEqual(results2, [2] * 3 + [3] * 2)
							πF.SetLineno(421)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							πTemp001[0] = µresults2
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = πg.NewInt(2).ToObject()
							πTemp008 = πg.NewList(πTemp009...).ToObject()
							if πTemp005, πE = πg.Mul(πF, πTemp008, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp009 = make([]*πg.Object, 1)
							πTemp009[0] = πg.NewInt(3).ToObject()
							πTemp010 = πg.NewList(πTemp009...).ToObject()
							if πTemp008, πE = πg.Mul(πF, πTemp010, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
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
							// line 422: b.wait_for_finished()
							πF.SetLineno(422)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µb, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_check_notify.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 424: def test_notify(self):
					πF.SetLineno(424)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_notify", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcond *πg.Object = πg.UnboundLocal; _ = µcond
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 425: cond = self.condtype()
							πF.SetLineno(425)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcondtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcond = πTemp002
							// line 426: self._check_notify(cond)
							πF.SetLineno(426)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							πTemp003[0] = µcond
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_check_notify, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 428: self._check_notify(cond)
							πF.SetLineno(428)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
								continue
							}
							πTemp003[0] = µcond
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_check_notify, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_notify.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 430: def test_timeout(self):
					πF.SetLineno(430)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_timeout", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcond *πg.Object = πg.UnboundLocal; _ = µcond
						var µresults *πg.Object = πg.UnboundLocal; _ = µresults
						var µN *πg.Object = πg.UnboundLocal; _ = µN
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µdt *πg.Object = πg.UnboundLocal; _ = µdt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							// line 431: cond = self.condtype()
							πF.SetLineno(431)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcondtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcond = πTemp002
							// line 432: results = []
							πF.SetLineno(432)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							µresults = πTemp001
							// line 433: N = 5
							πF.SetLineno(433)
							µN = πg.NewInt(5).ToObject()
							// line 434: def f():
							πF.SetLineno(434)
							πTemp004 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µt1 *πg.Object = πg.UnboundLocal; _ = µt1
								var µt2 *πg.Object = πg.UnboundLocal; _ = µt2
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 []*πg.Object
								_ = πTemp003
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 435: cond.acquire()
									πF.SetLineno(435)
									if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcond, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 436: t1 = time.time()
									πF.SetLineno(436)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									µt1 = πTemp001
									// line 437: cond.wait(0.2)
									πF.SetLineno(437)
									πTemp003 = πF.MakeArgs(1)
									πTemp003[0] = πg.NewFloat(0.2).ToObject()
									if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcond, ßwait, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 438: t2 = time.time()
									πF.SetLineno(438)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtime, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									µt2 = πTemp001
									// line 439: cond.release()
									πF.SetLineno(439)
									if πE = πg.CheckLocal(πF, µcond, "cond"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcond, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 440: results.append(t2 - t1)
									πF.SetLineno(440)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µt2, "t2"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µt1, "t1"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Sub(πF, µt2, µt1); πE != nil {
										continue
									}
									πTemp003[0] = πTemp001
									if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µresults, ßappend, nil); πE != nil {
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
							µf = πTemp001
							// line 441: Bunch(f, N).wait_for_finished()
							πF.SetLineno(441)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[0] = µf
							if πE = πg.CheckLocal(πF, µN, "N"); πE != nil {
								continue
							}
							πTemp003[1] = µN
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 442: self.assertEqual(len(results), 5)
							πF.SetLineno(442)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							πTemp006[0] = µresults
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µresults); πE != nil {
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
							if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µdt = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 444: self.assertTrue(dt >= 0.2, dt)
							πF.SetLineno(444)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GE(πF, µdt, πg.NewFloat(0.2).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µdt, "dt"); πE != nil {
								continue
							}
							πTemp003[1] = µdt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßtest_timeout.ToObject(), πTemp007); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("ConditionTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßConditionTests.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 447: class BaseSemaphoreTests(BaseTestCase):
			πF.SetLineno(447)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßBaseTestCase); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BaseSemaphoreTests", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 448: """
					πF.SetLineno(448)
					// line 448: """
					πF.SetLineno(448)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Common tests for {bounded, unbounded} semaphore objects.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 452: def test_constructor(self):
					πF.SetLineno(452)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_constructor", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 453: self.assertRaises(ValueError, self.semtype, value = -1)
							πF.SetLineno(453)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"value", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 454: self.assertRaises(ValueError, self.semtype, value = -sys.maxint)
							πF.SetLineno(454)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmaxint, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"value", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_constructor.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 456: def test_acquire(self):
					πF.SetLineno(456)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_acquire", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsem *πg.Object = πg.UnboundLocal; _ = µsem
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
							// line 457: sem = self.semtype(1)
							πF.SetLineno(457)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsem = πTemp003
							// line 458: sem.acquire()
							πF.SetLineno(458)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 459: sem.release()
							πF.SetLineno(459)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 460: sem = self.semtype(2)
							πF.SetLineno(460)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsem = πTemp003
							// line 461: sem.acquire()
							πF.SetLineno(461)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 462: sem.acquire()
							πF.SetLineno(462)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 463: sem.release()
							πF.SetLineno(463)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 464: sem.release()
							πF.SetLineno(464)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_acquire.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 466: def test_acquire_destroy(self):
					πF.SetLineno(466)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_acquire_destroy", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsem *πg.Object = πg.UnboundLocal; _ = µsem
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
							// line 467: sem = self.semtype()
							πF.SetLineno(467)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsem = πTemp002
							// line 468: sem.acquire()
							πF.SetLineno(468)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 469: del sem
							πF.SetLineno(469)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							µsem = πg.UnboundLocal
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_acquire_destroy.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 471: def test_acquire_contended(self):
					πF.SetLineno(471)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_acquire_contended", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsem *πg.Object = πg.UnboundLocal; _ = µsem
						var µN *πg.Object = πg.UnboundLocal; _ = µN
						var µresults1 *πg.Object = πg.UnboundLocal; _ = µresults1
						var µresults2 *πg.Object = πg.UnboundLocal; _ = µresults2
						var µphase_num *πg.Object = πg.UnboundLocal; _ = µphase_num
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
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
							case 7: goto Label7
							case 8: goto Label8
							case 10: goto Label10
							case 11: goto Label11
							case 13: goto Label13
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							// line 472: sem = self.semtype(7)
							πF.SetLineno(472)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(7).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsem = πTemp003
							// line 473: sem.acquire()
							πF.SetLineno(473)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 474: N = 10
							πF.SetLineno(474)
							µN = πg.NewInt(10).ToObject()
							// line 475: results1 = []
							πF.SetLineno(475)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresults1 = πTemp002
							// line 476: results2 = []
							πF.SetLineno(476)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresults2 = πTemp002
							// line 477: phase_num = 0
							πF.SetLineno(477)
							µphase_num = πg.NewInt(0).ToObject()
							// line 478: def f():
							πF.SetLineno(478)
							πTemp004 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 []*πg.Object
								_ = πTemp003
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 479: sem.acquire()
									πF.SetLineno(479)
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 480: results1.append(phase_num)
									πF.SetLineno(480)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µphase_num, "phase_num"); πE != nil {
										continue
									}
									πTemp003[0] = µphase_num
									if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µresults1, ßappend, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 481: sem.acquire()
									πF.SetLineno(481)
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 482: results2.append(phase_num)
									πF.SetLineno(482)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µphase_num, "phase_num"); πE != nil {
										continue
									}
									πTemp003[0] = µphase_num
									if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µresults2, ßappend, nil); πE != nil {
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
							µf = πTemp002
							// line 483: b = Bunch(f, 10)
							πF.SetLineno(483)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[0] = µf
							πTemp001[1] = πg.NewInt(10).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µb = πTemp005
							// line 484: b.wait_for_started()
							πF.SetLineno(484)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb, ßwait_for_started, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 485: while len(results1) + len(results2) < 6:
							πF.SetLineno(485)
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							πTemp001[0] = µresults2
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.Add(πF, πTemp009, πTemp010); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, πTemp005, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 486: _wait()
							πF.SetLineno(486)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 487: self.assertEqual(results1 + results2, [0] * 6)
							πF.SetLineno(487)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µresults1, µresults2); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp011 = make([]*πg.Object, 1)
							πTemp011[0] = πg.NewInt(0).ToObject()
							πTemp005 = πg.NewList(πTemp011...).ToObject()
							if πTemp003, πE = πg.Mul(πF, πTemp005, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 488: phase_num = 1
							πF.SetLineno(488)
							µphase_num = πg.NewInt(1).ToObject()
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(7).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
							if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µi = πTemp005
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 490: sem.release()
							πF.SetLineno(490)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 491: while len(results1) + len(results2) < 13:
							πF.SetLineno(491)
							πF.PushCheckpoint(8)
							πTemp006 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label9
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							πTemp001[0] = µresults2
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.Add(πF, πTemp009, πTemp010); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, πTemp005, πg.NewInt(13).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 492: _wait()
							πF.SetLineno(492)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							// line 493: self.assertEqual(sorted(results1 + results2), [0] * 6 + [1] * 7)
							πF.SetLineno(493)
							πTemp001 = πF.MakeArgs(2)
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µresults1, µresults2); πE != nil {
								continue
							}
							πTemp011[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp001[0] = πTemp005
							πTemp011 = make([]*πg.Object, 1)
							πTemp011[0] = πg.NewInt(0).ToObject()
							πTemp008 = πg.NewList(πTemp011...).ToObject()
							if πTemp005, πE = πg.Mul(πF, πTemp008, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							πTemp011 = make([]*πg.Object, 1)
							πTemp011[0] = πg.NewInt(1).ToObject()
							πTemp009 = πg.NewList(πTemp011...).ToObject()
							if πTemp008, πE = πg.Mul(πF, πTemp009, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 494: phase_num = 2
							πF.SetLineno(494)
							µphase_num = πg.NewInt(2).ToObject()
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(6).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp006 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label12
							}
							if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µi = πTemp005
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(10)            
							// line 496: sem.release()
							πF.SetLineno(496)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							// line 497: while len(results1) + len(results2) < 19:
							πF.SetLineno(497)
							πF.PushCheckpoint(14)
							πTemp006 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label15
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							πTemp001[0] = µresults1
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							πTemp001[0] = µresults2
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.Add(πF, πTemp009, πTemp010); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, πTemp005, πg.NewInt(19).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 498: _wait()
							πF.SetLineno(498)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							// line 499: self.assertEqual(sorted(results1 + results2), [0] * 6 + [1] * 7 + [2] * 6)
							πF.SetLineno(499)
							πTemp001 = πF.MakeArgs(2)
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults1, "results1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresults2, "results2"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µresults1, µresults2); πE != nil {
								continue
							}
							πTemp011[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp001[0] = πTemp005
							πTemp011 = make([]*πg.Object, 1)
							πTemp011[0] = πg.NewInt(0).ToObject()
							πTemp009 = πg.NewList(πTemp011...).ToObject()
							if πTemp008, πE = πg.Mul(πF, πTemp009, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							πTemp011 = make([]*πg.Object, 1)
							πTemp011[0] = πg.NewInt(1).ToObject()
							πTemp010 = πg.NewList(πTemp011...).ToObject()
							if πTemp009, πE = πg.Mul(πF, πTemp010, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							πTemp011 = make([]*πg.Object, 1)
							πTemp011[0] = πg.NewInt(2).ToObject()
							πTemp009 = πg.NewList(πTemp011...).ToObject()
							if πTemp008, πE = πg.Mul(πF, πTemp009, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 501: self.assertFalse(sem.acquire(False))
							πF.SetLineno(501)
							πTemp001 = πF.MakeArgs(1)
							πTemp011 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp011[0] = πTemp003
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 503: sem.release()
							πF.SetLineno(503)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 504: b.wait_for_finished()
							πF.SetLineno(504)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µb, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_acquire_contended.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 506: def test_try_acquire(self):
					πF.SetLineno(506)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_try_acquire", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsem *πg.Object = πg.UnboundLocal; _ = µsem
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
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
							// line 507: sem = self.semtype(2)
							πF.SetLineno(507)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsem = πTemp003
							// line 508: self.assertTrue(sem.acquire(False))
							πF.SetLineno(508)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
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
							// line 509: self.assertTrue(sem.acquire(False))
							πF.SetLineno(509)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
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
							// line 510: self.assertFalse(sem.acquire(False))
							πF.SetLineno(510)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 511: sem.release()
							πF.SetLineno(511)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 512: self.assertTrue(sem.acquire(False))
							πF.SetLineno(512)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_try_acquire.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 514: def test_try_acquire_contended(self):
					πF.SetLineno(514)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_try_acquire_contended", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsem *πg.Object = πg.UnboundLocal; _ = µsem
						var µresults *πg.Object = πg.UnboundLocal; _ = µresults
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 515: sem = self.semtype(4)
							πF.SetLineno(515)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsem = πTemp003
							// line 516: sem.acquire()
							πF.SetLineno(516)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 517: results = []
							πF.SetLineno(517)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresults = πTemp002
							// line 518: def f():
							πF.SetLineno(518)
							πTemp004 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 519: results.append(sem.acquire(False))
									πF.SetLineno(519)
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πF.MakeArgs(1)
									if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									πTemp001[0] = πTemp004
									if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µresults, ßappend, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 520: results.append(sem.acquire(False))
									πF.SetLineno(520)
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πF.MakeArgs(1)
									if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
										continue
									}
									πTemp002[0] = πTemp003
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									πTemp001[0] = πTemp004
									if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µresults, ßappend, nil); πE != nil {
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
							µf = πTemp002
							// line 521: Bunch(f, 5).wait_for_finished()
							πF.SetLineno(521)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[0] = µf
							πTemp001[1] = πg.NewInt(5).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 525: self.assertEqual(sorted(results), [False] * 7 + [True] *  3 )
							πF.SetLineno(525)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresults, "results"); πE != nil {
								continue
							}
							πTemp006[0] = µresults
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp005
							πTemp006 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp006[0] = πTemp007
							πTemp007 = πg.NewList(πTemp006...).ToObject()
							if πTemp005, πE = πg.Mul(πF, πTemp007, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							πTemp006 = make([]*πg.Object, 1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp006[0] = πTemp008
							πTemp008 = πg.NewList(πTemp006...).ToObject()
							if πTemp007, πE = πg.Mul(πF, πTemp008, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_try_acquire_contended.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 527: def test_default_value(self):
					πF.SetLineno(527)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_default_value", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsem *πg.Object = πg.UnboundLocal; _ = µsem
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 529: sem = self.semtype()
							πF.SetLineno(529)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsem = πTemp002
							// line 530: sem.acquire()
							πF.SetLineno(530)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 531: def f():
							πF.SetLineno(531)
							πTemp003 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/lock_tests.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 532: sem.acquire()
									πF.SetLineno(532)
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 533: sem.release()
									πF.SetLineno(533)
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
							µf = πTemp001
							// line 534: b = Bunch(f, 1)
							πF.SetLineno(534)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp004[0] = µf
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBunch); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µb = πTemp005
							// line 535: b.wait_for_started()
							πF.SetLineno(535)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßwait_for_started, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 536: _wait()
							πF.SetLineno(536)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_wait); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 537: self.assertFalse(b.finished)
							πF.SetLineno(537)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßfinished, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 538: sem.release()
							πF.SetLineno(538)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 539: b.wait_for_finished()
							πF.SetLineno(539)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µb, ßwait_for_finished, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_default_value.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 541: def test_with(self):
					πF.SetLineno(541)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_with", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsem *πg.Object = πg.UnboundLocal; _ = µsem
						var µ_with *πg.Object = πg.UnboundLocal; _ = µ_with
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 542: sem = self.semtype(2)
							πF.SetLineno(542)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsem = πTemp003
							// line 543: def _with(err=None):
							πF.SetLineno(543)
							πTemp004 = make([]πg.Param, 1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[0] = πg.Param{Name: "err", Def: πTemp003}
							πTemp002 = πg.NewFunction(πg.NewCode("_with", "/usr/lib/python2.7/test/lock_tests.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µerr *πg.Object = πArgs[0]; _ = µerr
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 []*πg.Object
								_ = πTemp003
								var πTemp004 []*πg.Object
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
								var πTemp010 *πg.BaseException
								_ = πTemp010
								var πTemp011 *πg.Traceback
								_ = πTemp011
								var πTemp012 *πg.Type
								_ = πTemp012
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									case 1: goto Label1
									case 2: goto Label2
									default: panic("unexpected function state")
									}
									// line 544: with sem:
									πF.SetLineno(544)
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µsem.Type().ToObject(), ß__exit__, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µsem.Type().ToObject(), ß__enter__, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp002.Call(πF, πg.Args{µsem}, nil); πE != nil {
										continue
									}
									πF.PushCheckpoint(1)
									// line 545: self.assertTrue(sem.acquire(False))
									πF.SetLineno(545)
									πTemp003 = πF.MakeArgs(1)
									πTemp004 = πF.MakeArgs(1)
									if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
										continue
									}
									πTemp004[0] = πTemp005
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									πTemp003[0] = πTemp006
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									// line 546: sem.release()
									πF.SetLineno(546)
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 547: with sem:
									πF.SetLineno(547)
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetAttr(πF, µsem.Type().ToObject(), ß__exit__, nil); πE != nil {
										continue
									}
									if πTemp006, πE = πg.GetAttr(πF, µsem.Type().ToObject(), ß__enter__, nil); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp006.Call(πF, πg.Args{µsem}, nil); πE != nil {
										continue
									}
									πF.PushCheckpoint(2)
									// line 548: self.assertFalse(sem.acquire(False))
									πF.SetLineno(548)
									πTemp003 = πF.MakeArgs(1)
									πTemp004 = πF.MakeArgs(1)
									if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
										continue
									}
									πTemp004[0] = πTemp007
									if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
										continue
									}
									if πTemp007, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									πTemp003[0] = πTemp008
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp007, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
										continue
									}
									if πTemp008, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
										continue
									}
									if πTemp009, πE = πg.IsTrue(πF, µerr); πE != nil {
										continue
									}
									if πTemp009 {
										goto Label3
									}
									goto Label4
									// line 549: if err:
									πF.SetLineno(549)
								Label3:
									if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
										continue
									}
									// line 550: raise err
									πF.SetLineno(550)
									πE = πF.Raise(µerr, nil, nil)
									continue
									goto Label4
								Label4:
									πF.PopCheckpoint()
								Label2:
									πTemp010, πTemp011 = nil, nil
									if πE != nil {
										πTemp010, πTemp011 = πF.ExcInfo()
									}
									if πTemp010 != nil {
										πTemp012 = πTemp010.Type()
										if πTemp007, πE = πTemp005.Call(πF, πg.Args{µsem, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
											continue
										}
									} else {
										if πTemp007, πE = πTemp005.Call(πF, πg.Args{µsem, πg.None, πg.None, πg.None}, nil); πE != nil {
											continue
										}
									}
									if πTemp009, πE = πg.IsTrue(πF, πTemp007); πE != nil {
										continue
									}
									if πTemp010 != nil && πTemp009 != true {
										πE = πF.Raise(nil, nil, nil)
										continue
									}
									if πR != nil {
										continue
									}
									πF.PopCheckpoint()
								Label1:
									πTemp010, πTemp011 = nil, nil
									if πE != nil {
										πTemp010, πTemp011 = πF.ExcInfo()
									}
									if πTemp010 != nil {
										πTemp012 = πTemp010.Type()
										if πTemp007, πE = πTemp001.Call(πF, πg.Args{µsem, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
											continue
										}
									} else {
										if πTemp007, πE = πTemp001.Call(πF, πg.Args{µsem, πg.None, πg.None, πg.None}, nil); πE != nil {
											continue
										}
									}
									if πTemp009, πE = πg.IsTrue(πF, πTemp007); πE != nil {
										continue
									}
									if πTemp010 != nil && πTemp009 != true {
										πE = πF.Raise(nil, nil, nil)
										continue
									}
									if πR != nil {
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
							µ_with = πTemp002
							// line 551: _with()
							πF.SetLineno(551)
							if πE = πg.CheckLocal(πF, µ_with, "_with"); πE != nil {
								continue
							}
							if πTemp003, πE = µ_with.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 552: self.assertTrue(sem.acquire(False))
							πF.SetLineno(552)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 553: sem.release()
							πF.SetLineno(553)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 554: self.assertRaises(TypeError, _with, TypeError)
							πF.SetLineno(554)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µ_with, "_with"); πE != nil {
								continue
							}
							πTemp001[1] = µ_with
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 555: self.assertTrue(sem.acquire(False))
							πF.SetLineno(555)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 556: sem.release()
							πF.SetLineno(556)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_with.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("BaseSemaphoreTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBaseSemaphoreTests.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 558: class SemaphoreTests(BaseSemaphoreTests):
			πF.SetLineno(558)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßBaseSemaphoreTests); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SemaphoreTests", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 559: """
					πF.SetLineno(559)
					// line 559: """
					πF.SetLineno(559)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Tests for unbounded semaphores.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 563: def test_release_unacquired(self):
					πF.SetLineno(563)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_release_unacquired", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsem *πg.Object = πg.UnboundLocal; _ = µsem
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
							// line 565: sem = self.semtype(1)
							πF.SetLineno(565)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsem = πTemp003
							// line 566: sem.release()
							πF.SetLineno(566)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 567: sem.acquire()
							πF.SetLineno(567)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 568: sem.acquire()
							πF.SetLineno(568)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 569: sem.release()
							πF.SetLineno(569)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_release_unacquired.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("SemaphoreTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSemaphoreTests.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 572: class BoundedSemaphoreTests(BaseSemaphoreTests):
			πF.SetLineno(572)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßBaseSemaphoreTests); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BoundedSemaphoreTests", "/usr/lib/python2.7/test/lock_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 573: """
					πF.SetLineno(573)
					// line 573: """
					πF.SetLineno(573)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Tests for bounded semaphores.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 577: def test_release_unacquired(self):
					πF.SetLineno(577)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_release_unacquired", "/usr/lib/python2.7/test/lock_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsem *πg.Object = πg.UnboundLocal; _ = µsem
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 579: sem = self.semtype()
							πF.SetLineno(579)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsemtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsem = πTemp002
							// line 580: self.assertRaises(ValueError, sem.release)
							πF.SetLineno(580)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 581: sem.acquire()
							πF.SetLineno(581)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsem, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 582: sem.release()
							πF.SetLineno(582)
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 583: self.assertRaises(ValueError, sem.release)
							πF.SetLineno(583)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µsem, "sem"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsem, ßrelease, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_release_unacquired.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("BoundedSemaphoreTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBoundedSemaphoreTests.ToObject(), πTemp007); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("test.lock_tests", Code)
}

