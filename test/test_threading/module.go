package test_threading
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/weakref"
	// _ "github.com/pygolin/stdlib/thread"
	// _ "github.com/pygolin/stdlib/threading"
	// _ "github.com/pygolin/stdlib/unittest"
	// _ "github.com/pygolin/stdlib/random"
	// _ "github.com/pygolin/stdlib/test/lock_tests"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/re"
	// _ "github.com/pygolin/stdlib/test/test_support"
	// _ "github.com/pygolin/stdlib/time"
	// _ "github.com/pygolin/stdlib/os"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßBaseTestCase := πg.InternStr("BaseTestCase")
		ßBoundedSemaphore := πg.InternStr("BoundedSemaphore")
		ßBoundedSemaphoreTests := πg.InternStr("BoundedSemaphoreTests")
		ßCondition := πg.InternStr("Condition")
		ßConditionAsRLockTests := πg.InternStr("ConditionAsRLockTests")
		ßConditionTests := πg.InternStr("ConditionTests")
		ßCounter := πg.InternStr("Counter")
		ßEvent := πg.InternStr("Event")
		ßEventTests := πg.InternStr("EventTests")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßImportError := πg.InternStr("ImportError")
		ßLock := πg.InternStr("Lock")
		ßLockTests := πg.InternStr("LockTests")
		ßMULTILINE := πg.InternStr("MULTILINE")
		ßNone := πg.InternStr("None")
		ßPIPE := πg.InternStr("PIPE")
		ßPopen := πg.InternStr("Popen")
		ßPyThreadState_SetAsyncExc := πg.InternStr("PyThreadState_SetAsyncExc")
		ßRLock := πg.InternStr("RLock")
		ßRLockTests := πg.InternStr("RLockTests")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßSemaphore := πg.InternStr("Semaphore")
		ßSemaphoreTests := πg.InternStr("SemaphoreTests")
		ßSystemExit := πg.InternStr("SystemExit")
		ßTestCase := πg.InternStr("TestCase")
		ßTestThread := πg.InternStr("TestThread")
		ßThread := πg.InternStr("Thread")
		ßThreadJoinOnShutdown := πg.InternStr("ThreadJoinOnShutdown")
		ßThreadTests := πg.InternStr("ThreadTests")
		ßThreadingExceptionTests := πg.InternStr("ThreadingExceptionTests")
		ßTrue := πg.InternStr("True")
		ßUnboundLocalError := πg.InternStr("UnboundLocalError")
		ßValueError := πg.InternStr("ValueError")
		ßZeroDivisionError := πg.InternStr("ZeroDivisionError")
		ß_DummyThread := πg.InternStr("_DummyThread")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_active := πg.InternStr("_active")
		ß_exit := πg.InternStr("_exit")
		ß_limbo := πg.InternStr("_limbo")
		ß_run := πg.InternStr("_run")
		ß_run_and_join := πg.InternStr("_run_and_join")
		ß_start_new_thread := πg.InternStr("_start_new_thread")
		ß_testcapi := πg.InternStr("_testcapi")
		ß_threads := πg.InternStr("_threads")
		ßacquire := πg.InternStr("acquire")
		ßaddCleanup := πg.InternStr("addCleanup")
		ßappend := πg.InternStr("append")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertGreaterEqual := πg.InternStr("assertGreaterEqual")
		ßassertIn := πg.InternStr("assertIn")
		ßassertIsInstance := πg.InternStr("assertIsInstance")
		ßassertIsNone := πg.InternStr("assertIsNone")
		ßassertIsNotNone := πg.InternStr("assertIsNotNone")
		ßassertLessEqual := πg.InternStr("assertLessEqual")
		ßassertNotEqual := πg.InternStr("assertNotEqual")
		ßassertNotIn := πg.InternStr("assertNotIn")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertRegexpMatches := πg.InternStr("assertRegexpMatches")
		ßassertScriptHasOutput := πg.InternStr("assertScriptHasOutput")
		ßassertTrue := πg.InternStr("assertTrue")
		ßassert_python_ok := πg.InternStr("assert_python_ok")
		ßc_long := πg.InternStr("c_long")
		ßcall := πg.InternStr("call")
		ßcall_in_temporary_c_thread := πg.InternStr("call_in_temporary_c_thread")
		ßclose := πg.InternStr("close")
		ßcommunicate := πg.InternStr("communicate")
		ßcondtype := πg.InternStr("condtype")
		ßcpython_only := πg.InternStr("cpython_only")
		ßctypes := πg.InternStr("ctypes")
		ßcurrentThread := πg.InternStr("currentThread")
		ßcurrent_thread := πg.InternStr("current_thread")
		ßdaemon := πg.InternStr("daemon")
		ßdarwin := πg.InternStr("darwin")
		ßdec := πg.InternStr("dec")
		ßdecode := πg.InternStr("decode")
		ßdone := πg.InternStr("done")
		ßenumerate := πg.InternStr("enumerate")
		ßerror := πg.InternStr("error")
		ßeventtype := πg.InternStr("eventtype")
		ßexecutable := πg.InternStr("executable")
		ßfail := πg.InternStr("fail")
		ßfinished := πg.InternStr("finished")
		ßfork := πg.InternStr("fork")
		ßfreebsd4 := πg.InternStr("freebsd4")
		ßfreebsd5 := πg.InternStr("freebsd5")
		ßfreebsd6 := πg.InternStr("freebsd6")
		ßgen := πg.InternStr("gen")
		ßgenerator := πg.InternStr("generator")
		ßget := πg.InternStr("get")
		ßget_ident := πg.InternStr("get_ident")
		ßgetcheckinterval := πg.InternStr("getcheckinterval")
		ßgetrefcount := πg.InternStr("getrefcount")
		ßgettrace := πg.InternStr("gettrace")
		ßgrumpy := πg.InternStr("grumpy")
		ßhasattr := πg.InternStr("hasattr")
		ßid := πg.InternStr("id")
		ßident := πg.InternStr("ident")
		ßinc := πg.InternStr("inc")
		ßis_alive := πg.InternStr("is_alive")
		ßjoin := πg.InternStr("join")
		ßlock_tests := πg.InternStr("lock_tests")
		ßlocktype := πg.InternStr("locktype")
		ßmutex := πg.InternStr("mutex")
		ßname := πg.InternStr("name")
		ßnetbsd5 := πg.InternStr("netbsd5")
		ßnext := πg.InternStr("next")
		ßnrunning := πg.InternStr("nrunning")
		ßobject := πg.InternStr("object")
		ßos := πg.InternStr("os")
		ßos2emx := πg.InternStr("os2emx")
		ßplatform := πg.InternStr("platform")
		ßplatforms_to_skip := πg.InternStr("platforms_to_skip")
		ßpy_object := πg.InternStr("py_object")
		ßpythonapi := πg.InternStr("pythonapi")
		ßrandom := πg.InternStr("random")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßread := πg.InternStr("read")
		ßreap_children := πg.InternStr("reap_children")
		ßref := πg.InternStr("ref")
		ßrelease := πg.InternStr("release")
		ßreplace := πg.InternStr("replace")
		ßrepr := πg.InternStr("repr")
		ßreturncode := πg.InternStr("returncode")
		ßrun := πg.InternStr("run")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsema := πg.InternStr("sema")
		ßsemtype := πg.InternStr("semtype")
		ßset := πg.InternStr("set")
		ßsetUp := πg.InternStr("setUp")
		ßsetattr := πg.InternStr("setattr")
		ßsetcheckinterval := πg.InternStr("setcheckinterval")
		ßsettrace := πg.InternStr("settrace")
		ßshould_raise := πg.InternStr("should_raise")
		ßskip := πg.InternStr("skip")
		ßskipIf := πg.InternStr("skipIf")
		ßskipTest := πg.InternStr("skipTest")
		ßskipUnless := πg.InternStr("skipUnless")
		ßsleep := πg.InternStr("sleep")
		ßstack_size := πg.InternStr("stack_size")
		ßstart := πg.InternStr("start")
		ßstart_new_thread := πg.InternStr("start_new_thread")
		ßstaticmethod := πg.InternStr("staticmethod")
		ßstderr := πg.InternStr("stderr")
		ßstdout := πg.InternStr("stdout")
		ßstrip := πg.InternStr("strip")
		ßsub := πg.InternStr("sub")
		ßsubprocess := πg.InternStr("subprocess")
		ßsys := πg.InternStr("sys")
		ßtask := πg.InternStr("task")
		ßtearDown := πg.InternStr("tearDown")
		ßtest := πg.InternStr("test")
		ßtest_1_join_on_shutdown := πg.InternStr("test_1_join_on_shutdown")
		ßtest_2_join_in_forked_process := πg.InternStr("test_2_join_in_forked_process")
		ßtest_3_join_in_forked_from_thread := πg.InternStr("test_3_join_in_forked_from_thread")
		ßtest_4_joining_across_fork_in_worker_thread := πg.InternStr("test_4_joining_across_fork_in_worker_thread")
		ßtest_5_clear_waiter_locks_to_avoid_crash := πg.InternStr("test_5_clear_waiter_locks_to_avoid_crash")
		ßtest_BoundedSemaphore_limit := πg.InternStr("test_BoundedSemaphore_limit")
		ßtest_PyThreadState_SetAsyncExc := πg.InternStr("test_PyThreadState_SetAsyncExc")
		ßtest_daemonize_active_thread := πg.InternStr("test_daemonize_active_thread")
		ßtest_dummy_thread_after_fork := πg.InternStr("test_dummy_thread_after_fork")
		ßtest_enumerate_after_join := πg.InternStr("test_enumerate_after_join")
		ßtest_finalize_runnning_thread := πg.InternStr("test_finalize_runnning_thread")
		ßtest_finalize_with_trace := πg.InternStr("test_finalize_with_trace")
		ßtest_foreign_thread := πg.InternStr("test_foreign_thread")
		ßtest_frame_tstate_tracing := πg.InternStr("test_frame_tstate_tracing")
		ßtest_ident_of_no_threading_threads := πg.InternStr("test_ident_of_no_threading_threads")
		ßtest_is_alive_after_fork := πg.InternStr("test_is_alive_after_fork")
		ßtest_join_nondaemon_on_shutdown := πg.InternStr("test_join_nondaemon_on_shutdown")
		ßtest_joining_current_thread := πg.InternStr("test_joining_current_thread")
		ßtest_joining_inactive_thread := πg.InternStr("test_joining_inactive_thread")
		ßtest_limbo_cleanup := πg.InternStr("test_limbo_cleanup")
		ßtest_main := πg.InternStr("test_main")
		ßtest_no_refcycle_through_target := πg.InternStr("test_no_refcycle_through_target")
		ßtest_print_exception := πg.InternStr("test_print_exception")
		ßtest_print_exception_stderr_is_none_1 := πg.InternStr("test_print_exception_stderr_is_none_1")
		ßtest_print_exception_stderr_is_none_2 := πg.InternStr("test_print_exception_stderr_is_none_2")
		ßtest_recursion_limit := πg.InternStr("test_recursion_limit")
		ßtest_reinit_tls_after_fork := πg.InternStr("test_reinit_tls_after_fork")
		ßtest_start_thread_again := πg.InternStr("test_start_thread_again")
		ßtest_support := πg.InternStr("test_support")
		ßtest_various_ops := πg.InternStr("test_various_ops")
		ßtest_various_ops_large_stack := πg.InternStr("test_various_ops_large_stack")
		ßtest_various_ops_small_stack := πg.InternStr("test_various_ops_small_stack")
		ßtestcase := πg.InternStr("testcase")
		ßthread := πg.InternStr("thread")
		ßthreading := πg.InternStr("threading")
		ßthreading_cleanup := πg.InternStr("threading_cleanup")
		ßthreading_setup := πg.InternStr("threading_setup")
		ßtime := πg.InternStr("time")
		ßunittest := πg.InternStr("unittest")
		ßvalue := πg.InternStr("value")
		ßverbose := πg.InternStr("verbose")
		ßwait := πg.InternStr("wait")
		ßwaitpid := πg.InternStr("waitpid")
		ßweakref := πg.InternStr("weakref")
		ßxrange := πg.InternStr("xrange")
		ßyet_another := πg.InternStr("yet_another")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 bool
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 3: import test.test_support
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: from test.test_support import verbose, cpython_only
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßverbose); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßverbose.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßcpython_only); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcpython_only.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 7: import random
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "random"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßrandom.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import re
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: import sys
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: import thread
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "thread"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthread.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: import threading
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "threading"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: import time
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: import unittest
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import weakref
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "weakref"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweakref.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 17: import os
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: _testcapi = None
			πF.SetLineno(22)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_testcapi.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 24: from test import lock_tests
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "test.lock_tests"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßlock_tests.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: class Counter(object):
			πF.SetLineno(27)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Counter", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 28: def __init__(self):
					πF.SetLineno(28)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 29: self.value = 0
							πF.SetLineno(29)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßvalue, πTemp001); πE != nil {
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
					// line 30: def inc(self):
					πF.SetLineno(30)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("inc", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 31: self.value += 1
							πF.SetLineno(31)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßvalue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßvalue, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßinc.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 32: def dec(self):
					πF.SetLineno(32)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("dec", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 33: self.value -= 1
							πF.SetLineno(33)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßvalue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßvalue, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdec.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 34: def get(self):
					πF.SetLineno(34)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("get", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 35: return self.value
							πF.SetLineno(35)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßvalue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Counter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCounter.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 37: class TestThread(threading.Thread):
			πF.SetLineno(37)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßThread, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestThread", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 38: def __init__(self, name, testcase, sema, mutex, nrunning):
					πF.SetLineno(38)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp002[2] = πg.Param{Name: "testcase", Def: nil}
					πTemp002[3] = πg.Param{Name: "sema", Def: nil}
					πTemp002[4] = πg.Param{Name: "mutex", Def: nil}
					πTemp002[5] = πg.Param{Name: "nrunning", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µtestcase *πg.Object = πArgs[2]; _ = µtestcase
						var µsema *πg.Object = πArgs[3]; _ = µsema
						var µmutex *πg.Object = πArgs[4]; _ = µmutex
						var µnrunning *πg.Object = πArgs[5]; _ = µnrunning
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
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
							// line 39: threading.Thread.__init__(self, name=name)
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = πg.KWArgs{
								{"name", µname},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßThread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 40: self.testcase = testcase
							πF.SetLineno(40)
							if πE = πg.CheckLocal(πF, µtestcase, "testcase"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µtestcase); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtestcase, πTemp003); πE != nil {
								continue
							}
							// line 41: self.sema = sema
							πF.SetLineno(41)
							if πE = πg.CheckLocal(πF, µsema, "sema"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µsema); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßsema, πTemp003); πE != nil {
								continue
							}
							// line 42: self.mutex = mutex
							πF.SetLineno(42)
							if πE = πg.CheckLocal(πF, µmutex, "mutex"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µmutex); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmutex, πTemp003); πE != nil {
								continue
							}
							// line 43: self.nrunning = nrunning
							πF.SetLineno(43)
							if πE = πg.CheckLocal(πF, µnrunning, "nrunning"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µnrunning); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnrunning, πTemp003); πE != nil {
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
					// line 45: def run(self):
					πF.SetLineno(45)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdelay *πg.Object = πg.UnboundLocal; _ = µdelay
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
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 *πg.Type
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 9: goto Label9
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 46: delay = random.random() / 10000.0
							πF.SetLineno(46)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrandom, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Div(πF, πTemp002, πg.NewFloat(10000.0).ToObject()); πE != nil {
								continue
							}
							µdelay = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 47: if verbose:
							πF.SetLineno(47)
						Label1:
							// line 48: print 'task %s will run for %s usec' % (
							πF.SetLineno(48)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdelay, "delay"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mul(πF, µdelay, πg.NewFloat(1000000.0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("task %s will run for %s usec").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 51: with self.sema:
							πF.SetLineno(51)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsema, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp003.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							// line 52: with self.mutex:
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmutex, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp008.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							// line 53: self.nrunning.inc()
							πF.SetLineno(53)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßnrunning, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßinc, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp010.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 54: if verbose:
							πF.SetLineno(54)
						Label5:
							// line 55: print self.nrunning.get(), 'tasks are running'
							πF.SetLineno(55)
							πTemp005 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßnrunning, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßget, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp010.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp009
							πTemp005[1] = πg.NewStr("tasks are running").ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label6
						Label6:
							// line 56: self.testcase.assertLessEqual(self.nrunning.get(), 3)
							πF.SetLineno(56)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßnrunning, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßget, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp010.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp009
							πTemp005[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßtestcase, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßassertLessEqual, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp010.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
						Label4:
							πTemp011, πTemp012 = nil, nil
							if πE != nil {
								πTemp011, πTemp012 = πF.ExcInfo()
							}
							if πTemp011 != nil {
								πTemp013 = πTemp011.Type()
								if πTemp009, πE = πTemp007.Call(πF, πg.Args{πTemp006, πTemp013.ToObject(), πTemp011.ToObject(), πTemp012.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp009, πE = πTemp007.Call(πF, πg.Args{πTemp006, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if πTemp011 != nil && πTemp004 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 58: time.sleep(delay)
							πF.SetLineno(58)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelay, "delay"); πE != nil {
								continue
							}
							πTemp005[0] = µdelay
							if πTemp009, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßsleep, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp010.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 59: if verbose:
							πF.SetLineno(59)
						Label7:
							// line 60: print 'task', self.name, 'done'
							πF.SetLineno(60)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = ßtask.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp009
							πTemp005[2] = ßdone.ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label8
						Label8:
							// line 62: with self.mutex:
							πF.SetLineno(62)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßmutex, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, πTemp009.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp014.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(9)
							// line 63: self.nrunning.dec()
							πF.SetLineno(63)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetAttr(πF, µself, ßnrunning, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßdec, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp016.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 64: self.testcase.assertGreaterEqual(self.nrunning.get(), 0)
							πF.SetLineno(64)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetAttr(πF, µself, ßnrunning, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßget, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp016.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp015
							πTemp005[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetAttr(πF, µself, ßtestcase, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßassertGreaterEqual, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp016.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp015, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp015); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 65: if verbose:
							πF.SetLineno(65)
						Label10:
							// line 66: print '%s is finished. %d tasks are running' % (
							πF.SetLineno(66)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp017, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp018, πE = πg.GetAttr(πF, µself, ßnrunning, nil); πE != nil {
								continue
							}
							if πTemp019, πE = πg.GetAttr(πF, πTemp018, ßget, nil); πE != nil {
								continue
							}
							if πTemp018, πE = πTemp019.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp016 = πg.NewTuple2(πTemp017, πTemp018).ToObject()
							if πTemp015, πE = πg.Mod(πF, πg.NewStr("%s is finished. %d tasks are running").ToObject(), πTemp016); πE != nil {
								continue
							}
							πTemp005[0] = πTemp015
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label11
						Label11:
							πF.PopCheckpoint()
						Label9:
							πTemp011, πTemp012 = nil, nil
							if πE != nil {
								πTemp011, πTemp012 = πF.ExcInfo()
							}
							if πTemp011 != nil {
								πTemp013 = πTemp011.Type()
								if πTemp015, πE = πTemp010.Call(πF, πg.Args{πTemp009, πTemp013.ToObject(), πTemp011.ToObject(), πTemp012.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp015, πE = πTemp010.Call(πF, πg.Args{πTemp009, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp015); πE != nil {
								continue
							}
							if πTemp011 != nil && πTemp004 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							πF.PopCheckpoint()
						Label3:
							πTemp011, πTemp012 = nil, nil
							if πE != nil {
								πTemp011, πTemp012 = πF.ExcInfo()
							}
							if πTemp011 != nil {
								πTemp013 = πTemp011.Type()
								if πTemp015, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp013.ToObject(), πTemp011.ToObject(), πTemp012.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp015, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp015); πE != nil {
								continue
							}
							if πTemp011 != nil && πTemp004 != true {
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TestThread").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestThread.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 69: class BaseTestCase(unittest.TestCase):
			πF.SetLineno(69)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BaseTestCase", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 70: def setUp(self):
					πF.SetLineno(70)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("setUp", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 71: self._threads = test.test_support.threading_setup()
							πF.SetLineno(71)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtest_support, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßthreading_setup, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_threads, πTemp001); πE != nil {
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
					// line 73: def tearDown(self):
					πF.SetLineno(73)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("tearDown", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 74: test.test_support.threading_cleanup(*self._threads)
							πF.SetLineno(74)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_threads, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_support, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßthreading_cleanup, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, πTemp001, nil, nil); πE != nil {
								continue
							}
							// line 75: test.test_support.reap_children()
							πF.SetLineno(75)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßtest_support, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßreap_children, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtearDown.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BaseTestCase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBaseTestCase.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 78: class ThreadTests(BaseTestCase):
			πF.SetLineno(78)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseTestCase); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ThreadTests", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 []*πg.Object
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
				var πTemp018 []*πg.Object
				_ = πTemp018
				var πTemp019 *πg.Object
				_ = πTemp019
				var πTemp020 *πg.Object
				_ = πTemp020
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 82: def test_various_ops(self):
					πF.SetLineno(82)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_various_ops", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µNUMTASKS *πg.Object = πg.UnboundLocal; _ = µNUMTASKS
						var µsema *πg.Object = πg.UnboundLocal; _ = µsema
						var µmutex *πg.Object = πg.UnboundLocal; _ = µmutex
						var µnumrunning *πg.Object = πg.UnboundLocal; _ = µnumrunning
						var µthreads *πg.Object = πg.UnboundLocal; _ = µthreads
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var πTemp001 πg.KWArgs
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 6: goto Label6
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 85: NUMTASKS = 10
							πF.SetLineno(85)
							µNUMTASKS = πg.NewInt(10).ToObject()
							// line 88: sema = threading.BoundedSemaphore(value=3)
							πF.SetLineno(88)
							πTemp001 = πg.KWArgs{
								{"value", πg.NewInt(3).ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßBoundedSemaphore, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, πTemp001); πE != nil {
								continue
							}
							µsema = πTemp002
							// line 89: mutex = threading.RLock()
							πF.SetLineno(89)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßRLock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µmutex = πTemp002
							// line 90: numrunning = Counter()
							πF.SetLineno(90)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßCounter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnumrunning = πTemp003
							// line 92: threads = []
							πF.SetLineno(92)
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							µthreads = πTemp002
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µNUMTASKS, "NUMTASKS"); πE != nil {
								continue
							}
							πTemp004[0] = µNUMTASKS
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
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
							πF.PushCheckpoint(1)            
							// line 95: t = TestThread("<thread %d>"%i, self, sema, mutex, numrunning)
							πF.SetLineno(95)
							πTemp004 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("<thread %d>").ToObject(), µi); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µsema, "sema"); πE != nil {
								continue
							}
							πTemp004[2] = µsema
							if πE = πg.CheckLocal(πF, µmutex, "mutex"); πE != nil {
								continue
							}
							πTemp004[3] = µmutex
							if πE = πg.CheckLocal(πF, µnumrunning, "numrunning"); πE != nil {
								continue
							}
							πTemp004[4] = µnumrunning
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTestThread); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µt = πTemp005
							// line 96: threads.append(t)
							πF.SetLineno(96)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp004[0] = µt
							if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µthreads, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 97: self.assertIsNone(t.ident)
							πF.SetLineno(97)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßident, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIsNone, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 98: self.assertRegexpMatches(repr(t), r'^<TestThread\(.*, initial\)>$')
							πF.SetLineno(98)
							πTemp004 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp008[0] = µt
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp005
							πTemp004[1] = πg.NewStr("^<TestThread\\(.*, initial\\)>$").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRegexpMatches, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 99: t.start()
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 101: if verbose:
							πF.SetLineno(101)
						Label4:
							// line 102: print 'waiting for all tasks to complete'
							πF.SetLineno(102)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = πg.NewStr("waiting for all tasks to complete").ToObject()
							if πE = πg.Print(πF, πTemp004, true); πE != nil {
								continue
							}
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µthreads); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp006 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label8
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
								µt = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(6)            
							// line 104: t.join(NUMTASKS)
							πF.SetLineno(104)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µNUMTASKS, "NUMTASKS"); πE != nil {
								continue
							}
							πTemp004[0] = µNUMTASKS
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 105: self.assertFalse(t.is_alive())
							πF.SetLineno(105)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßis_alive, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 106: self.assertNotEqual(t.ident, 0)
							πF.SetLineno(106)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßident, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertNotEqual, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 107: self.assertIsNotNone(t.ident)
							πF.SetLineno(107)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßident, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIsNotNone, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 108: self.assertRegexpMatches(repr(t), r'^<TestThread\(.*, \w+ -?\d+\)>$')
							πF.SetLineno(108)
							πTemp004 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp008[0] = µt
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp005
							πTemp004[1] = πg.NewStr("^<TestThread\\(.*, \\w+ -?\\d+\\)>$").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRegexpMatches, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 109: if verbose:
							πF.SetLineno(109)
						Label9:
							// line 110: print 'all tasks done'
							πF.SetLineno(110)
							πTemp004 = make([]*πg.Object, 1)
							πTemp004[0] = πg.NewStr("all tasks done").ToObject()
							if πE = πg.Print(πF, πTemp004, true); πE != nil {
								continue
							}
							goto Label10
						Label10:
							// line 111: self.assertEqual(numrunning.get(), 0)
							πF.SetLineno(111)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µnumrunning, "numrunning"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µnumrunning, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_various_ops.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 113: def test_ident_of_no_threading_threads(self):
					πF.SetLineno(113)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_ident_of_no_threading_threads", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µdone *πg.Object = πg.UnboundLocal; _ = µdone
						var µident *πg.Object = πg.UnboundLocal; _ = µident
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
							// line 115: self.assertIsNotNone(threading.currentThread().ident)
							πF.SetLineno(115)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcurrentThread, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßident, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsNotNone, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 116: def f():
							πF.SetLineno(116)
							πTemp004 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/test_threading.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 117: ident.append(threading.currentThread().ident)
									πF.SetLineno(117)
									πTemp001 = πF.MakeArgs(1)
									if πTemp002, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcurrentThread, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßident, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp003
									if πE = πg.CheckLocal(πF, µident, "ident"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µident, ßappend, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 118: done.set()
									πF.SetLineno(118)
									if πE = πg.CheckLocal(πF, µdone, "done"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µdone, ßset, nil); πE != nil {
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
							µf = πTemp002
							// line 119: done = threading.Event()
							πF.SetLineno(119)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßEvent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdone = πTemp003
							// line 120: ident = []
							πF.SetLineno(120)
							πTemp001 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp001...).ToObject()
							µident = πTemp003
							// line 121: thread.start_new_thread(f, ())
							πF.SetLineno(121)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[0] = µf
							πTemp003 = πg.NewTuple0().ToObject()
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßstart_new_thread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 122: done.wait()
							πF.SetLineno(122)
							if πE = πg.CheckLocal(πF, µdone, "done"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdone, ßwait, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 123: self.assertIsNotNone(ident[0])
							πF.SetLineno(123)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µident, "ident"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µident, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIsNotNone, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 125: del threading._active[ident[0]]
							πF.SetLineno(125)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ß_active, nil); πE != nil {
								continue
							}
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µident, "ident"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µident, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πTemp007
							if πE = πg.DelItem(πF, πTemp005, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_ident_of_no_threading_threads.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 128: def test_various_ops_small_stack(self):
					πF.SetLineno(128)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_various_ops_small_stack", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 129: if verbose:
							πF.SetLineno(129)
						Label1:
							// line 130: print 'with 256kB thread stack size...'
							πF.SetLineno(130)
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewStr("with 256kB thread stack size...").ToObject()
							if πE = πg.Print(πF, πTemp003, true); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 131: try:
							πF.SetLineno(131)
							πF.PushCheckpoint(4)
							// line 132: threading.stack_size(262144)
							πF.SetLineno(132)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(262144).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßstack_size, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp004); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 133: except thread.error:
							πF.SetLineno(133)
						Label5:
							// line 134: self.skipTest('platform does not support changing thread stack size')
							πF.SetLineno(134)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("platform does not support changing thread stack size").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßskipTest, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 135: self.test_various_ops()
							πF.SetLineno(135)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtest_various_ops, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 136: threading.stack_size(0)
							πF.SetLineno(136)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßstack_size, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_various_ops_small_stack.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 139: def test_various_ops_large_stack(self):
					πF.SetLineno(139)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_various_ops_large_stack", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 140: if verbose:
							πF.SetLineno(140)
						Label1:
							// line 141: print 'with 1MB thread stack size...'
							πF.SetLineno(141)
							πTemp003 = make([]*πg.Object, 1)
							πTemp003[0] = πg.NewStr("with 1MB thread stack size...").ToObject()
							if πE = πg.Print(πF, πTemp003, true); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 142: try:
							πF.SetLineno(142)
							πF.PushCheckpoint(4)
							// line 143: threading.stack_size(0x100000)
							πF.SetLineno(143)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1048576).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßstack_size, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp004); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 144: except thread.error:
							πF.SetLineno(144)
						Label5:
							// line 145: self.skipTest('platform does not support changing thread stack size')
							πF.SetLineno(145)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("platform does not support changing thread stack size").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßskipTest, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 146: self.test_various_ops()
							πF.SetLineno(146)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtest_various_ops, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 147: threading.stack_size(0)
							πF.SetLineno(147)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßstack_size, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_various_ops_large_stack.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 149: def test_foreign_thread(self):
					πF.SetLineno(149)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_foreign_thread", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µmutex *πg.Object = πg.UnboundLocal; _ = µmutex
						var µtid *πg.Object = πg.UnboundLocal; _ = µtid
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []πg.Param
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 151: def f(mutex):
							πF.SetLineno(151)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "mutex", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µmutex *πg.Object = πArgs[0]; _ = µmutex
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
									// line 154: threading.current_thread()
									πF.SetLineno(154)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcurrent_thread, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 155: mutex.release()
									πF.SetLineno(155)
									if πE = πg.CheckLocal(πF, µmutex, "mutex"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µmutex, ßrelease, nil); πE != nil {
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
							// line 157: mutex = threading.Lock()
							πF.SetLineno(157)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßLock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							µmutex = πTemp003
							// line 158: mutex.acquire()
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µmutex, "mutex"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmutex, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 159: tid = thread.start_new_thread(f, (mutex,))
							πF.SetLineno(159)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp005[0] = µf
							if πE = πg.CheckLocal(πF, µmutex, "mutex"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple1(µmutex).ToObject()
							πTemp005[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstart_new_thread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µtid = πTemp003
							// line 161: mutex.acquire()
							πF.SetLineno(161)
							if πE = πg.CheckLocal(πF, µmutex, "mutex"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmutex, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 162: self.assertIn(tid, threading._active)
							πF.SetLineno(162)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtid, "tid"); πE != nil {
								continue
							}
							πTemp005[0] = µtid
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_active, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 163: self.assertIsInstance(threading._active[tid], threading._DummyThread)
							πF.SetLineno(163)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtid, "tid"); πE != nil {
								continue
							}
							πTemp003 = µtid
							if πTemp006, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß_active, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_DummyThread, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 164: del threading._active[tid]
							πF.SetLineno(164)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_active, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtid, "tid"); πE != nil {
								continue
							}
							πTemp003 = µtid
							if πE = πg.DelItem(πF, πTemp004, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_foreign_thread.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 169: def test_PyThreadState_SetAsyncExc(self):
					πF.SetLineno(169)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_PyThreadState_SetAsyncExc", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µset_async_exc *πg.Object = πg.UnboundLocal; _ = µset_async_exc
						var µAsyncExc *πg.Object = πg.UnboundLocal; _ = µAsyncExc
						var µexception *πg.Object = πg.UnboundLocal; _ = µexception
						var µtid *πg.Object = πg.UnboundLocal; _ = µtid
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µworker_started *πg.Object = πg.UnboundLocal; _ = µworker_started
						var µworker_saw_exception *πg.Object = πg.UnboundLocal; _ = µworker_saw_exception
						var µWorker *πg.Object = πg.UnboundLocal; _ = µWorker
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var µret *πg.Object = πg.UnboundLocal; _ = µret
						var πTemp001 *πg.BaseException
						_ = πTemp001
						var πTemp002 *πg.Traceback
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Dict
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 πg.KWArgs
						_ = πTemp012
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							case 11: goto Label11
							case 5: goto Label5
							case 6: goto Label6
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 170: try:
							πF.SetLineno(170)
							πF.PushCheckpoint(2)
							// line 172: pass
							πF.SetLineno(172)
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp001, πTemp002 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp001.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							πE = πF.Raise(πTemp001.ToObject(), nil, πTemp002.ToObject())
							continue
							// line 173: except ImportError:
							πF.SetLineno(173)
						Label3:
							// line 174: self.skipTest('requires ctypes')
							πF.SetLineno(174)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("requires ctypes").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßskipTest, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 176: set_async_exc = ctypes.pythonapi.PyThreadState_SetAsyncExc
							πF.SetLineno(176)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpythonapi, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßPyThreadState_SetAsyncExc, nil); πE != nil {
								continue
							}
							µset_async_exc = πTemp003
							// line 178: class AsyncExc(Exception):
							πF.SetLineno(178)
							πTemp005 = make([]*πg.Object, 1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp005[0] = πTemp008
							πTemp007 = πg.NewDict()
							if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
								continue
							}
							_, πE = πg.NewCode("AsyncExc", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp007
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 179: pass
									πF.SetLineno(179)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp006, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp006 == nil {
								πTemp006 = πg.TypeType.ToObject()
							}
							if πTemp008, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("AsyncExc").ToObject(), πg.NewTuple(πTemp005...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
								continue
							}
							µAsyncExc = πTemp008
							// line 181: exception = ctypes.py_object(AsyncExc)
							πF.SetLineno(181)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µAsyncExc, "AsyncExc"); πE != nil {
								continue
							}
							πTemp005[0] = µAsyncExc
							if πTemp003, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßpy_object, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µexception = πTemp003
							// line 184: tid = thread.get_ident()
							πF.SetLineno(184)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßget_ident, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							µtid = πTemp003
							// line 186: try:
							πF.SetLineno(186)
							πF.PushCheckpoint(5)
							// line 187: result = set_async_exc(ctypes.c_long(tid), exception)
							πF.SetLineno(187)
							πTemp005 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtid, "tid"); πE != nil {
								continue
							}
							πTemp009[0] = µtid
							if πTemp003, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßc_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							πTemp005[1] = µexception
							if πE = πg.CheckLocal(πF, µset_async_exc, "set_async_exc"); πE != nil {
								continue
							}
							if πTemp003, πE = µset_async_exc.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µresult = πTemp003
							// line 190: while True:
							πF.SetLineno(190)
							πF.PushCheckpoint(7)
							πTemp004 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(6)            
							// line 191: pass
							πF.SetLineno(191)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							πF.PopCheckpoint()
							// line 197: self.fail("AsyncExc not raised")
							πF.SetLineno(197)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("AsyncExc not raised").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp001, πTemp002 = πF.ExcInfo()
							if πE = πg.CheckLocal(πF, µAsyncExc, "AsyncExc"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp001.ToObject(), µAsyncExc); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							πE = πF.Raise(πTemp001.ToObject(), nil, πTemp002.ToObject())
							continue
							// line 192: except AsyncExc:
							πF.SetLineno(192)
						Label9:
							// line 193: pass
							πF.SetLineno(193)
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							// line 198: try:
							πF.SetLineno(198)
							πF.PushCheckpoint(11)
							// line 199: self.assertEqual(result, 1) # one thread state modified
							πF.SetLineno(199)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp005[0] = µresult
							πTemp005[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label10
						Label11:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp001, πTemp002 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßUnboundLocalError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp001.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							πE = πF.Raise(πTemp001.ToObject(), nil, πTemp002.ToObject())
							continue
							// line 200: except UnboundLocalError:
							πF.SetLineno(200)
						Label12:
							// line 202: pass
							πF.SetLineno(202)
							πF.RestoreExc(nil, nil)
							goto Label10
						Label10:
							// line 208: worker_started = threading.Event()
							πF.SetLineno(208)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßEvent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							µworker_started = πTemp003
							// line 209: worker_saw_exception = threading.Event()
							πF.SetLineno(209)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßEvent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							µworker_saw_exception = πTemp003
							// line 211: class Worker(threading.Thread):
							πF.SetLineno(211)
							πTemp005 = make([]*πg.Object, 1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp008, ßThread, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp011
							πTemp007 = πg.NewDict()
							if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Worker", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp007
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
									// line 212: def run(self):
									πF.SetLineno(212)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 *πg.Object
										_ = πTemp002
										var πTemp003 bool
										_ = πTemp003
										var πTemp004 bool
										_ = πTemp004
										var πTemp005 []*πg.Object
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
											case 2: goto Label2
											case 3: goto Label3
											case 4: goto Label4
											default: panic("unexpected function state")
											}
											// line 213: self.id = thread.get_ident()
											πF.SetLineno(213)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget_ident, nil); πE != nil {
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
											if πE = πg.SetAttr(πF, µself, ßid, πTemp002); πE != nil {
												continue
											}
											// line 214: self.finished = False
											πF.SetLineno(214)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßfinished, πTemp002); πE != nil {
												continue
											}
											// line 216: try:
											πF.SetLineno(216)
											πF.PushCheckpoint(2)
											// line 217: while True:
											πF.SetLineno(217)
											πF.PushCheckpoint(4)
											πTemp003 = false
										Label3:
											if πE != nil || πR != nil {
												continue
											}
											if πTemp003 {
												πF.PopCheckpoint()
												goto Label5
											}
											if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
												continue
											}
											if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πE != nil || !πTemp004 {
												continue
											}
											πF.PushCheckpoint(3)            
											// line 218: worker_started.set()
											πF.SetLineno(218)
											if πE = πg.CheckLocal(πF, µworker_started, "worker_started"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µworker_started, ßset, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 219: time.sleep(0.1)
											πF.SetLineno(219)
											πTemp005 = πF.MakeArgs(1)
											πTemp005[0] = πg.NewFloat(0.1).ToObject()
											if πTemp001, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsleep, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
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
											πTemp006, πTemp007 = πF.ExcInfo()
											if πE = πg.CheckLocal(πF, µAsyncExc, "AsyncExc"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), µAsyncExc); πE != nil {
												continue
											}
											if πTemp003 {
												goto Label6
											}
											πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
											continue
											// line 220: except AsyncExc:
											πF.SetLineno(220)
										Label6:
											// line 221: self.finished = True
											πF.SetLineno(221)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßfinished, πTemp002); πE != nil {
												continue
											}
											// line 222: worker_saw_exception.set()
											πF.SetLineno(222)
											if πE = πg.CheckLocal(πF, µworker_saw_exception, "worker_saw_exception"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µworker_saw_exception, ßset, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
												continue
											}
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
									if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp006, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp006 == nil {
								πTemp006 = πg.TypeType.ToObject()
							}
							if πTemp008, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Worker").ToObject(), πg.NewTuple(πTemp005...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
								continue
							}
							µWorker = πTemp008
							// line 224: t = Worker()
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µWorker, "Worker"); πE != nil {
								continue
							}
							if πTemp003, πE = µWorker.Call(πF, nil, nil); πE != nil {
								continue
							}
							µt = πTemp003
							// line 225: t.daemon = True # so if this fails, we don't hang Python at shutdown
							πF.SetLineno(225)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µt, ßdaemon, πTemp006); πE != nil {
								continue
							}
							// line 226: t.start()
							πF.SetLineno(226)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							goto Label14
							// line 227: if verbose:
							πF.SetLineno(227)
						Label13:
							// line 228: print "    started worker thread"
							πF.SetLineno(228)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewStr("    started worker thread").ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label14
						Label14:
							if πTemp003, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							goto Label16
							// line 231: if verbose:
							πF.SetLineno(231)
						Label15:
							// line 232: print "    trying nonsensical thread id"
							πF.SetLineno(232)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewStr("    trying nonsensical thread id").ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label16
						Label16:
							// line 233: result = set_async_exc(ctypes.c_long(-1), exception)
							πF.SetLineno(233)
							πTemp005 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp009[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßc_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							πTemp005[1] = µexception
							if πE = πg.CheckLocal(πF, µset_async_exc, "set_async_exc"); πE != nil {
								continue
							}
							if πTemp003, πE = µset_async_exc.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µresult = πTemp003
							// line 234: self.assertEqual(result, 0)  # no thread states modified
							πF.SetLineno(234)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp005[0] = µresult
							πTemp005[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label17
							}
							goto Label18
							// line 237: if verbose:
							πF.SetLineno(237)
						Label17:
							// line 238: print "    waiting for worker thread to get started"
							πF.SetLineno(238)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewStr("    waiting for worker thread to get started").ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label18
						Label18:
							// line 239: ret = worker_started.wait()
							πF.SetLineno(239)
							if πE = πg.CheckLocal(πF, µworker_started, "worker_started"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µworker_started, ßwait, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µret = πTemp006
							// line 240: self.assertTrue(ret)
							πF.SetLineno(240)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
								continue
							}
							πTemp005[0] = µret
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label19
							}
							goto Label20
							// line 241: if verbose:
							πF.SetLineno(241)
						Label19:
							// line 242: print "    verifying worker hasn't exited"
							πF.SetLineno(242)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewStr("    verifying worker hasn't exited").ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label20
						Label20:
							// line 243: self.assertFalse(t.finished)
							πF.SetLineno(243)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßfinished, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label21
							}
							goto Label22
							// line 244: if verbose:
							πF.SetLineno(244)
						Label21:
							// line 245: print "    attempting to raise asynch exception in worker"
							πF.SetLineno(245)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewStr("    attempting to raise asynch exception in worker").ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label22
						Label22:
							// line 246: result = set_async_exc(ctypes.c_long(t.id), exception)
							πF.SetLineno(246)
							πTemp005 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßid, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßctypes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßc_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µexception, "exception"); πE != nil {
								continue
							}
							πTemp005[1] = µexception
							if πE = πg.CheckLocal(πF, µset_async_exc, "set_async_exc"); πE != nil {
								continue
							}
							if πTemp003, πE = µset_async_exc.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µresult = πTemp003
							// line 247: self.assertEqual(result, 1) # one thread state modified
							πF.SetLineno(247)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp005[0] = µresult
							πTemp005[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label23
							}
							goto Label24
							// line 248: if verbose:
							πF.SetLineno(248)
						Label23:
							// line 249: print "    waiting for worker to say it caught the exception"
							πF.SetLineno(249)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewStr("    waiting for worker to say it caught the exception").ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label24
						Label24:
							// line 250: worker_saw_exception.wait(timeout=10)
							πF.SetLineno(250)
							πTemp012 = πg.KWArgs{
								{"timeout", πg.NewInt(10).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µworker_saw_exception, "worker_saw_exception"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µworker_saw_exception, ßwait, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, πTemp012); πE != nil {
								continue
							}
							// line 251: self.assertTrue(t.finished)
							πF.SetLineno(251)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßfinished, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßverbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label25
							}
							goto Label26
							// line 252: if verbose:
							πF.SetLineno(252)
						Label25:
							// line 253: print "    all OK -- joining worker"
							πF.SetLineno(253)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewStr("    all OK -- joining worker").ToObject()
							if πE = πg.Print(πF, πTemp005, true); πE != nil {
								continue
							}
							goto Label26
						Label26:
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßfinished, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label27
							}
							goto Label28
							// line 254: if t.finished:
							πF.SetLineno(254)
						Label27:
							// line 255: t.join()
							πF.SetLineno(255)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label28
						Label28:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_PyThreadState_SetAsyncExc.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 168: @unittest.skip('grumpy')
					πF.SetLineno(168)
					πTemp008 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßtest_PyThreadState_SetAsyncExc); πE != nil {
						continue
					}
					πTemp008[0] = πTemp009
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp009, ßskip, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp011.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp011, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_PyThreadState_SetAsyncExc.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 258: def test_limbo_cleanup(self):
					πF.SetLineno(258)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_limbo_cleanup", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfail_new_thread *πg.Object = πg.UnboundLocal; _ = µfail_new_thread
						var µ_start_new_thread *πg.Object = πg.UnboundLocal; _ = µ_start_new_thread
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []πg.Param
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 260: def fail_new_thread(*args):
							πF.SetLineno(260)
							πTemp002 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("fail_new_thread", "/usr/lib/python2.7/test/test_threading.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µargs *πg.Object = πArgs[0]; _ = µargs
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
									if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 261: raise thread.error()
									πF.SetLineno(261)
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
							µfail_new_thread = πTemp001
							// line 262: _start_new_thread = threading._start_new_thread
							πF.SetLineno(262)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß_start_new_thread, nil); πE != nil {
								continue
							}
							µ_start_new_thread = πTemp004
							// line 263: threading._start_new_thread = fail_new_thread
							πF.SetLineno(263)
							if πE = πg.CheckLocal(πF, µfail_new_thread, "fail_new_thread"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µfail_new_thread); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ß_start_new_thread, πTemp003); πE != nil {
								continue
							}
							// line 264: try:
							πF.SetLineno(264)
							πF.PushCheckpoint(1)
							// line 265: t = threading.Thread(target=lambda: None)
							πF.SetLineno(265)
							πTemp002 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var πTemp001 *πg.Object
								_ = πTemp001
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 265: t = threading.Thread(target=lambda: None)
									πF.SetLineno(265)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
							πTemp005 = πg.KWArgs{
								{"target", πTemp003},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßThread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							µt = πTemp003
							// line 266: self.assertRaises(thread.error, t.start)
							πF.SetLineno(266)
							πTemp006 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßerror, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 267: self.assertFalse(
							πF.SetLineno(267)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ß_limbo, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, πTemp007, µt); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp008).ToObject()
							πTemp006[0] = πTemp003
							πTemp006[1] = πg.NewStr("Failed to cleanup _limbo map on failure of Thread.start().").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp009, πTemp010 = πF.RestoreExc(nil, nil)
							// line 271: threading._start_new_thread = _start_new_thread
							πF.SetLineno(271)
							if πE = πg.CheckLocal(πF, µ_start_new_thread, "_start_new_thread"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µ_start_new_thread); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ß_start_new_thread, πTemp003); πE != nil {
								continue
							}
							if πTemp009 != nil {
								πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_limbo_cleanup.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 274: def test_finalize_runnning_thread(self):
					πF.SetLineno(274)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_finalize_runnning_thread", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrc *πg.Object = πg.UnboundLocal; _ = µrc
						var πTemp001 *πg.BaseException
						_ = πTemp001
						var πTemp002 *πg.Traceback
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
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
							// line 278: try:
							πF.SetLineno(278)
							πF.PushCheckpoint(2)
							// line 280: pass
							πF.SetLineno(280)
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp001, πTemp002 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp001.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							πE = πF.Raise(πTemp001.ToObject(), nil, πTemp002.ToObject())
							continue
							// line 281: except ImportError:
							πF.SetLineno(281)
						Label3:
							// line 282: self.skipTest('requires ctypes')
							πF.SetLineno(282)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("requires ctypes").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßskipTest, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 284: rc = subprocess.call([sys.executable, "-c", """if 1:
							πF.SetLineno(284)
							πTemp005 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßexecutable, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp006
							πTemp007[1] = πg.NewStr("-c").ToObject()
							πTemp007[2] = πg.NewStr("if 1:\n            import ctypes, sys, time, thread\n\n            # This lock is used as a simple event variable.\n            ready = thread.allocate_lock()\n            ready.acquire()\n\n            # Module globals are cleared before __del__ is run\n            # So we save the functions in class dict\n            class C:\n                ensure = ctypes.pythonapi.PyGILState_Ensure\n                release = ctypes.pythonapi.PyGILState_Release\n                def __del__(self):\n                    state = self.ensure()\n                    self.release(state)\n\n            def waitingThread():\n                x = C()\n                ready.release()\n                time.sleep(100)\n\n            thread.start_new_thread(waitingThread, ())\n            ready.acquire()  # Be sure the other thread is waiting.\n            sys.exit(42)\n            ").ToObject()
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							πTemp005[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßcall, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µrc = πTemp003
							// line 309: self.assertEqual(rc, 42)
							πF.SetLineno(309)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
								continue
							}
							πTemp005[0] = µrc
							πTemp005[1] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_finalize_runnning_thread.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 273: @unittest.skip('grumpy')
					πF.SetLineno(273)
					πTemp008 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßtest_finalize_runnning_thread); πE != nil {
						continue
					}
					πTemp008[0] = πTemp012
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßskip, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp013.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp013, πE = πTemp012.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_finalize_runnning_thread.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 312: def test_finalize_with_trace(self):
					πF.SetLineno(312)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_finalize_with_trace", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µstdout *πg.Object = πg.UnboundLocal; _ = µstdout
						var µstderr *πg.Object = πg.UnboundLocal; _ = µstderr
						var µrc *πg.Object = πg.UnboundLocal; _ = µrc
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 315: p = subprocess.Popen([sys.executable, "-c", """if 1:
							πF.SetLineno(315)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexecutable, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewStr("-c").ToObject()
							πTemp002[2] = πg.NewStr("if 1:\n            import sys, threading\n\n            # A deadlock-killer, to prevent the\n            # testsuite to hang forever\n            def killer():\n                import os, time\n                time.sleep(2)\n                print 'program blocked; aborting'\n                os._exit(2)\n            t = threading.Thread(target=killer)\n            t.daemon = True\n            t.start()\n\n            # This is the trace function\n            def func(frame, event, arg):\n                threading.current_thread()\n                return func\n\n            sys.settrace(func)\n            ").ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"stdout", πTemp004},
								{"stderr", πTemp005},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPopen, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µp = πTemp003
							// line 338: self.addCleanup(p.stdout.close)
							πF.SetLineno(338)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßaddCleanup, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 339: self.addCleanup(p.stderr.close)
							πF.SetLineno(339)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßstderr, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßaddCleanup, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 340: stdout, stderr = p.communicate()
							πF.SetLineno(340)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßcommunicate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
								continue
							}
							µstdout = πTemp003
							µstderr = πTemp005
							// line 341: rc = p.returncode
							πF.SetLineno(341)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßreturncode, nil); πE != nil {
								continue
							}
							µrc = πTemp003
							// line 342: self.assertFalse(rc == 2, "interpreted was blocked")
							πF.SetLineno(342)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µrc, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("interpreted was blocked").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 343: self.assertTrue(rc == 0,
							πF.SetLineno(343)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µrc, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstderr, "stderr"); πE != nil {
								continue
							}
							πTemp002[0] = µstderr
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Add(πF, πg.NewStr("Unexpected error: ").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_finalize_with_trace.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 311: @unittest.skip('grumpy')
					πF.SetLineno(311)
					πTemp008 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßtest_finalize_with_trace); πE != nil {
						continue
					}
					πTemp008[0] = πTemp013
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßskip, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp014.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp014, πE = πTemp013.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_finalize_with_trace.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 347: def test_join_nondaemon_on_shutdown(self):
					πF.SetLineno(347)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_join_nondaemon_on_shutdown", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µstdout *πg.Object = πg.UnboundLocal; _ = µstdout
						var µstderr *πg.Object = πg.UnboundLocal; _ = µstderr
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
						var πTemp006 πg.KWArgs
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 350: p = subprocess.Popen([sys.executable, "-c", """if 1:
							πF.SetLineno(350)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexecutable, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewStr("-c").ToObject()
							πTemp002[2] = πg.NewStr("if 1:\n                import threading\n                from time import sleep\n\n                def child():\n                    sleep(1)\n                    # As a non-daemon thread we SHOULD wake up and nothing\n                    # should be torn down yet\n                    print \"Woke up, sleep function is:\", sleep\n\n                threading.Thread(target=child).start()\n                raise SystemExit\n            ").ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"stdout", πTemp004},
								{"stderr", πTemp005},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPopen, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µp = πTemp003
							// line 365: self.addCleanup(p.stdout.close)
							πF.SetLineno(365)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßaddCleanup, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 366: self.addCleanup(p.stderr.close)
							πF.SetLineno(366)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßstderr, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßclose, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßaddCleanup, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 367: stdout, stderr = p.communicate()
							πF.SetLineno(367)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßcommunicate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
								continue
							}
							µstdout = πTemp003
							µstderr = πTemp005
							// line 368: self.assertEqual(stdout.strip(),
							πF.SetLineno(368)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstdout, "stdout"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstdout, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewStr("Woke up, sleep function is: <built-in function sleep>").ToObject()
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
							// line 370: stderr = re.sub(r"^\[\d+ refs\]", "", stderr, re.MULTILINE).strip()
							πF.SetLineno(370)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("^\\[\\d+ refs\\]").ToObject()
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µstderr, "stderr"); πE != nil {
								continue
							}
							πTemp001[2] = µstderr
							if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßMULTILINE, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsub, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstderr = πTemp003
							// line 371: self.assertEqual(stderr, "")
							πF.SetLineno(371)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstderr, "stderr"); πE != nil {
								continue
							}
							πTemp001[0] = µstderr
							πTemp001[1] = ß.ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_join_nondaemon_on_shutdown.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 346: @unittest.skip('grumpy')
					πF.SetLineno(346)
					πTemp008 = πF.MakeArgs(1)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßtest_join_nondaemon_on_shutdown); πE != nil {
						continue
					}
					πTemp008[0] = πTemp014
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßskip, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp015.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp015, πE = πTemp014.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_join_nondaemon_on_shutdown.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 374: def test_enumerate_after_join(self):
					πF.SetLineno(374)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_enumerate_after_join", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µenum *πg.Object = πg.UnboundLocal; _ = µenum
						var µold_interval *πg.Object = πg.UnboundLocal; _ = µold_interval
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []πg.Param
						_ = πTemp007
						var πTemp008 πg.KWArgs
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
							case 1: goto Label1
							case 2: goto Label2
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							// line 377: enum = threading.enumerate
							πF.SetLineno(377)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßenumerate, nil); πE != nil {
								continue
							}
							µenum = πTemp002
							// line 378: old_interval = sys.getcheckinterval()
							πF.SetLineno(378)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßgetcheckinterval, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µold_interval = πTemp001
							// line 379: try:
							πF.SetLineno(379)
							πF.PushCheckpoint(1)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewInt(1).ToObject()
							πTemp003[1] = πg.NewInt(100).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							πTemp005 = false
						Label2:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label4
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(2)            
							// line 383: sys.setcheckinterval(i // 5)
							πF.SetLineno(383)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.FloorDiv(πF, µi, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßsetcheckinterval, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 384: t = threading.Thread(target=lambda: None)
							πF.SetLineno(384)
							πTemp007 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/test/test_threading.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var πTemp001 *πg.Object
								_ = πTemp001
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 384: t = threading.Thread(target=lambda: None)
									πF.SetLineno(384)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
							πTemp008 = πg.KWArgs{
								{"target", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßThread, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							µt = πTemp002
							// line 385: t.start()
							πF.SetLineno(385)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 386: t.join()
							πF.SetLineno(386)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 387: l = enum()
							πF.SetLineno(387)
							if πE = πg.CheckLocal(πF, µenum, "enum"); πE != nil {
								continue
							}
							if πTemp002, πE = µenum.Call(πF, nil, nil); πE != nil {
								continue
							}
							µl = πTemp002
							// line 388: self.assertNotIn(t, l,
							πF.SetLineno(388)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp003[0] = µt
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp003[1] = µl
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(µi, µl).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("#1703448 triggered after %d trials: %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp003[2] = πTemp002
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
							continue
						Label3:
							if πE != nil || πR != nil {
								continue
							}
						Label4:
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp009, πTemp010 = πF.RestoreExc(nil, nil)
							// line 391: sys.setcheckinterval(old_interval)
							πF.SetLineno(391)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µold_interval, "old_interval"); πE != nil {
								continue
							}
							πTemp003[0] = µold_interval
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsetcheckinterval, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp009 != nil {
								πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_enumerate_after_join.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 373: @unittest.skip('grumpy')
					πF.SetLineno(373)
					πTemp008 = πF.MakeArgs(1)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßtest_enumerate_after_join); πE != nil {
						continue
					}
					πTemp008[0] = πTemp015
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßskip, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp016.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp016, πE = πTemp015.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_enumerate_after_join.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 394: def test_no_refcycle_through_target(self):
					πF.SetLineno(394)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_no_refcycle_through_target", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µRunSelfFunction *πg.Object = πg.UnboundLocal; _ = µRunSelfFunction
						var µcyclic_object *πg.Object = πg.UnboundLocal; _ = µcyclic_object
						var µweak_cyclic_object *πg.Object = πg.UnboundLocal; _ = µweak_cyclic_object
						var µraising_cyclic_object *πg.Object = πg.UnboundLocal; _ = µraising_cyclic_object
						var µweak_raising_cyclic_object *πg.Object = πg.UnboundLocal; _ = µweak_raising_cyclic_object
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
						var πTemp006 πg.KWArgs
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
							// line 395: class RunSelfFunction(object):
							πF.SetLineno(395)
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
							_, πE = πg.NewCode("RunSelfFunction", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 396: def __init__(self, should_raise):
									πF.SetLineno(396)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "should_raise", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µshould_raise *πg.Object = πArgs[1]; _ = µshould_raise
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 *πg.Object
										_ = πTemp002
										var πTemp003 *πg.Dict
										_ = πTemp003
										var πTemp004 *πg.Object
										_ = πTemp004
										var πTemp005 πg.KWArgs
										_ = πTemp005
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 399: self.should_raise = should_raise
											πF.SetLineno(399)
											if πE = πg.CheckLocal(πF, µshould_raise, "should_raise"); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µshould_raise); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßshould_raise, πTemp001); πE != nil {
												continue
											}
											// line 400: self.thread = threading.Thread(target=self._run,
											πF.SetLineno(400)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ß_run, nil); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											πTemp002 = πg.NewTuple1(µself).ToObject()
											πTemp003 = πg.NewDict()
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πTemp003.SetItem(πF, ßyet_another.ToObject(), µself); πE != nil {
												continue
											}
											πTemp004 = πTemp003.ToObject()
											πTemp005 = πg.KWArgs{
												{"target", πTemp001},
												{"args", πTemp002},
												{"kwargs", πTemp004},
											}
											if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßThread, nil); πE != nil {
												continue
											}
											if πTemp001, πE = πTemp002.Call(πF, nil, πTemp005); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßthread, πTemp002); πE != nil {
												continue
											}
											// line 403: self.thread.start()
											πF.SetLineno(403)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ßthread, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstart, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 405: def _run(self, other_ref, yet_another):
									πF.SetLineno(405)
									πTemp002 = make([]πg.Param, 3)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other_ref", Def: nil}
									πTemp002[2] = πg.Param{Name: "yet_another", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("_run", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother_ref *πg.Object = πArgs[1]; _ = µother_ref
										var µyet_another *πg.Object = πArgs[2]; _ = µyet_another
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
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ßshould_raise, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label1
											}
											goto Label2
											// line 406: if self.should_raise:
											πF.SetLineno(406)
										Label1:
											if πTemp001, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
												continue
											}
											// line 407: raise SystemExit
											πF.SetLineno(407)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
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
									if πE = πClass.SetItem(πF, ß_run.ToObject(), πTemp003); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("RunSelfFunction").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µRunSelfFunction = πTemp005
							// line 409: cyclic_object = RunSelfFunction(should_raise=False)
							πF.SetLineno(409)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"should_raise", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µRunSelfFunction, "RunSelfFunction"); πE != nil {
								continue
							}
							if πTemp002, πE = µRunSelfFunction.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							µcyclic_object = πTemp002
							// line 410: weak_cyclic_object = weakref.ref(cyclic_object)
							πF.SetLineno(410)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcyclic_object, "cyclic_object"); πE != nil {
								continue
							}
							πTemp003[0] = µcyclic_object
							if πTemp002, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßref, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µweak_cyclic_object = πTemp002
							// line 411: cyclic_object.thread.join()
							πF.SetLineno(411)
							if πE = πg.CheckLocal(πF, µcyclic_object, "cyclic_object"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcyclic_object, ßthread, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 412: del cyclic_object
							πF.SetLineno(412)
							if πE = πg.CheckLocal(πF, µcyclic_object, "cyclic_object"); πE != nil {
								continue
							}
							µcyclic_object = πg.UnboundLocal
							// line 413: self.assertEqual(None, weak_cyclic_object(),
							πF.SetLineno(413)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µweak_cyclic_object, "weak_cyclic_object"); πE != nil {
								continue
							}
							if πTemp002, πE = µweak_cyclic_object.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µweak_cyclic_object, "weak_cyclic_object"); πE != nil {
								continue
							}
							if πTemp004, πE = µweak_cyclic_object.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßgetrefcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%d references still around").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"msg", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 417: raising_cyclic_object = RunSelfFunction(should_raise=True)
							πF.SetLineno(417)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"should_raise", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µRunSelfFunction, "RunSelfFunction"); πE != nil {
								continue
							}
							if πTemp002, πE = µRunSelfFunction.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							µraising_cyclic_object = πTemp002
							// line 418: weak_raising_cyclic_object = weakref.ref(raising_cyclic_object)
							πF.SetLineno(418)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µraising_cyclic_object, "raising_cyclic_object"); πE != nil {
								continue
							}
							πTemp003[0] = µraising_cyclic_object
							if πTemp002, πE = πg.ResolveGlobal(πF, ßweakref); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßref, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µweak_raising_cyclic_object = πTemp002
							// line 419: raising_cyclic_object.thread.join()
							πF.SetLineno(419)
							if πE = πg.CheckLocal(πF, µraising_cyclic_object, "raising_cyclic_object"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µraising_cyclic_object, ßthread, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 420: del raising_cyclic_object
							πF.SetLineno(420)
							if πE = πg.CheckLocal(πF, µraising_cyclic_object, "raising_cyclic_object"); πE != nil {
								continue
							}
							µraising_cyclic_object = πg.UnboundLocal
							// line 421: self.assertEqual(None, weak_raising_cyclic_object(),
							πF.SetLineno(421)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µweak_raising_cyclic_object, "weak_raising_cyclic_object"); πE != nil {
								continue
							}
							if πTemp002, πE = µweak_raising_cyclic_object.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µweak_raising_cyclic_object, "weak_raising_cyclic_object"); πE != nil {
								continue
							}
							if πTemp004, πE = µweak_raising_cyclic_object.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßgetrefcount, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%d references still around").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp006 = πg.KWArgs{
								{"msg", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, πTemp006); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_no_refcycle_through_target.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 393: @unittest.skip('grumpy')
					πF.SetLineno(393)
					πTemp008 = πF.MakeArgs(1)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßtest_no_refcycle_through_target); πE != nil {
						continue
					}
					πTemp008[0] = πTemp016
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßskip, nil); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp017.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp017, πE = πTemp016.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_no_refcycle_through_target.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 427: def test_dummy_thread_after_fork(self):
					πF.SetLineno(427)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_dummy_thread_after_fork", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcode *πg.Object = πg.UnboundLocal; _ = µcode
						var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
						var µout *πg.Object = πg.UnboundLocal; _ = µout
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 430: code = """if 1:
							πF.SetLineno(430)
							µcode = πg.NewStr("if 1:\n            import thread, threading, os, time\n\n            def background_thread(evt):\n                # Creates and registers the _DummyThread instance\n                threading.current_thread()\n                evt.set()\n                time.sleep(10)\n\n            evt = threading.Event()\n            thread.start_new_thread(background_thread, (evt,))\n            evt.wait()\n            assert threading.active_count() == 2, threading.active_count()\n            if os.fork() == 0:\n                assert threading.active_count() == 1, threading.active_count()\n                os._exit(0)\n            else:\n                os.wait()\n        ").ToObject()
							// line 449: _, out, err = assert_python_ok("-c", code)
							πF.SetLineno(449)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("-c").ToObject()
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[1] = µcode
							if πTemp002, πE = πg.ResolveGlobal(πF, ßassert_python_ok); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µ_ = πTemp002
							µout = πTemp004
							µerr = πTemp005
							// line 450: self.assertEqual(out, '')
							πF.SetLineno(450)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[0] = µout
							πTemp001[1] = ß.ToObject()
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
							// line 451: self.assertEqual(err, '')
							πF.SetLineno(451)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[0] = µerr
							πTemp001[1] = ß.ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_dummy_thread_after_fork.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 425: @unittest.skip('grumpy')
					πF.SetLineno(425)
					πTemp008 = πF.MakeArgs(1)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßtest_dummy_thread_after_fork); πE != nil {
						continue
					}
					πTemp008[0] = πTemp017
					πTemp010 = πF.MakeArgs(2)
					πTemp018 = πF.MakeArgs(2)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					πTemp018[0] = πTemp017
					πTemp018[1] = ßfork.ToObject()
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πTemp019, πE = πTemp017.Call(πF, πTemp018, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp018)
					πTemp010[0] = πTemp019
					πTemp010[1] = πg.NewStr("test needs fork()").ToObject()
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp019, πE = πg.GetAttr(πF, πTemp017, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp017, πE = πTemp019.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp019, πE = πTemp017.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_dummy_thread_after_fork.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 425: @unittest.skip('grumpy')
					πF.SetLineno(425)
					πTemp008 = πF.MakeArgs(1)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßtest_dummy_thread_after_fork); πE != nil {
						continue
					}
					πTemp008[0] = πTemp017
					πTemp010 = πF.MakeArgs(1)
					πTemp010[0] = ßgrumpy.ToObject()
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp019, πE = πg.GetAttr(πF, πTemp017, ßskip, nil); πE != nil {
						continue
					}
					if πTemp017, πE = πTemp019.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp019, πE = πTemp017.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_dummy_thread_after_fork.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 454: def test_is_alive_after_fork(self):
					πF.SetLineno(454)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("test_is_alive_after_fork", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µold_interval *πg.Object = πg.UnboundLocal; _ = µold_interval
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var µpid *πg.Object = πg.UnboundLocal; _ = µpid
						var µstatus *πg.Object = πg.UnboundLocal; _ = µstatus
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []πg.Param
						_ = πTemp007
						var πTemp008 πg.KWArgs
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.BaseException
						_ = πTemp010
						var πTemp011 *πg.Traceback
						_ = πTemp011
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
							// line 457: old_interval = sys.getcheckinterval()
							πF.SetLineno(457)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßgetcheckinterval, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µold_interval = πTemp001
							// line 460: sys.setcheckinterval(10)
							πF.SetLineno(460)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(10).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsetcheckinterval, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 462: try:
							πF.SetLineno(462)
							πF.PushCheckpoint(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(20).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							πTemp005 = false
						Label2:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label4
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(2)            
							// line 464: t = threading.Thread(target=lambda: None)
							πF.SetLineno(464)
							πTemp007 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/test/test_threading.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var πTemp001 *πg.Object
								_ = πTemp001
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 464: t = threading.Thread(target=lambda: None)
									πF.SetLineno(464)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
							πTemp008 = πg.KWArgs{
								{"target", πTemp002},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßThread, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, πTemp008); πE != nil {
								continue
							}
							µt = πTemp002
							// line 465: t.start()
							πF.SetLineno(465)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 466: pid = os.fork()
							πF.SetLineno(466)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßfork, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							µpid = πTemp002
							if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µpid, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 467: if pid == 0:
							πF.SetLineno(467)
						Label5:
							// line 468: os._exit(1 if t.is_alive() else 0)
							πF.SetLineno(468)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µt, ßis_alive, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label8
							}
							πTemp002 = πg.NewInt(1).ToObject()
							goto Label9
						Label8:
							πTemp002 = πg.NewInt(0).ToObject()
						Label9:
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ß_exit, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label7
						Label6:
							// line 470: t.join()
							πF.SetLineno(470)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 471: pid, status = os.waitpid(pid, 0)
							πF.SetLineno(471)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
								continue
							}
							πTemp003[0] = µpid
							πTemp003[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßwaitpid, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp009}}}, πTemp002); πE != nil {
								continue
							}
							µpid = πTemp004
							µstatus = πTemp009
							// line 472: self.assertEqual(0, status)
							πF.SetLineno(472)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µstatus, "status"); πE != nil {
								continue
							}
							πTemp003[1] = µstatus
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
							goto Label7
						Label7:
							continue
						Label3:
							if πE != nil || πR != nil {
								continue
							}
						Label4:
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp010, πTemp011 = πF.RestoreExc(nil, nil)
							// line 474: sys.setcheckinterval(old_interval)
							πF.SetLineno(474)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µold_interval, "old_interval"); πE != nil {
								continue
							}
							πTemp003[0] = µold_interval
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsetcheckinterval, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp010 != nil {
								πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_is_alive_after_fork.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 453: @unittest.skipUnless(hasattr(os, 'fork'), "needs os.fork()")
					πF.SetLineno(453)
					πTemp008 = πF.MakeArgs(1)
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßtest_is_alive_after_fork); πE != nil {
						continue
					}
					πTemp008[0] = πTemp019
					πTemp010 = πF.MakeArgs(2)
					πTemp018 = πF.MakeArgs(2)
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					πTemp018[0] = πTemp019
					πTemp018[1] = ßfork.ToObject()
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πTemp020, πE = πTemp019.Call(πF, πTemp018, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp018)
					πTemp010[0] = πTemp020
					πTemp010[1] = πg.NewStr("needs os.fork()").ToObject()
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp019, πE = πTemp020.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp020, πE = πTemp019.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πClass.SetItem(πF, ßtest_is_alive_after_fork.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 476: def test_BoundedSemaphore_limit(self):
					πF.SetLineno(476)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("test_BoundedSemaphore_limit", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlimit *πg.Object = πg.UnboundLocal; _ = µlimit
						var µbs *πg.Object = πg.UnboundLocal; _ = µbs
						var µthreads *πg.Object = πg.UnboundLocal; _ = µthreads
						var µt *πg.Object = πg.UnboundLocal; _ = µt
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
						var πTemp009 bool
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
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewInt(1).ToObject()
							πTemp002[1] = πg.NewInt(10).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
								µlimit = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 479: bs = threading.BoundedSemaphore(limit)
							πF.SetLineno(479)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
								continue
							}
							πTemp002[0] = µlimit
							if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßBoundedSemaphore, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µbs = πTemp003
							// line 480: threads = [threading.Thread(target=bs.acquire)
							πF.SetLineno(480)
							πTemp007 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/test_threading.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
								var πTemp007 πg.KWArgs
								_ = πTemp007
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
										if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
											continue
										}
										πTemp002[0] = µlimit
										if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
											µ_ = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 480: threads = [threading.Thread(target=bs.acquire)
										πF.SetLineno(480)
										if πE = πg.CheckLocal(πF, µbs, "bs"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µbs, ßacquire, nil); πE != nil {
											continue
										}
										πTemp007 = πg.KWArgs{
											{"target", πTemp003},
										}
										if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßThread, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp004.Call(πF, nil, πTemp007); πE != nil {
											continue
										}
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
							µthreads = πTemp003
							if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µthreads); πE != nil {
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
							if πTemp008, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µt = πTemp008
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 483: t.start()
							πF.SetLineno(483)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µthreads); πE != nil {
								continue
							}
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
							if πTemp008, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µt = πTemp008
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 485: t.join()
							πF.SetLineno(485)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							// line 486: threads = [threading.Thread(target=bs.release)
							πF.SetLineno(486)
							πTemp007 = make([]πg.Param, 0)
							πTemp008 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/test_threading.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
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
								var πTemp007 πg.KWArgs
								_ = πTemp007
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
										if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
											continue
										}
										πTemp002[0] = µlimit
										if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
											µ_ = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 486: threads = [threading.Thread(target=bs.release)
										πF.SetLineno(486)
										if πE = πg.CheckLocal(πF, µbs, "bs"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µbs, ßrelease, nil); πE != nil {
											continue
										}
										πTemp007 = πg.KWArgs{
											{"target", πTemp003},
										}
										if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßThread, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πTemp004.Call(πF, nil, πTemp007); πE != nil {
											continue
										}
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
							if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
								continue
							}
							µthreads = πTemp003
							if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µthreads); πE != nil {
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
							if πTemp010, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µt = πTemp010
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(10)            
							// line 489: t.start()
							πF.SetLineno(489)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µthreads); πE != nil {
								continue
							}
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
							if πTemp010, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µt = πTemp010
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 491: t.join()
							πF.SetLineno(491)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							// line 492: self.assertRaises(ValueError, bs.release)
							πF.SetLineno(492)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µbs, "bs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µbs, ßrelease, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_BoundedSemaphore_limit.ToObject(), πTemp019); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ThreadTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßThreadTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 494: class ThreadJoinOnShutdown(BaseTestCase):
			πF.SetLineno(494)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseTestCase); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ThreadJoinOnShutdown", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 []*πg.Object
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 500: platforms_to_skip = ('freebsd4', 'freebsd5', 'freebsd6', 'netbsd5',
					πF.SetLineno(500)
					πTemp001 = πg.NewTuple5(ßfreebsd4.ToObject(), ßfreebsd5.ToObject(), ßfreebsd6.ToObject(), ßnetbsd5.ToObject(), ßos2emx.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßplatforms_to_skip.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 503: def _run_and_join(self, script):
					πF.SetLineno(503)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "script", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_run_and_join", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πArgs[1]; _ = µscript
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µrc *πg.Object = πg.UnboundLocal; _ = µrc
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 504: script = """if 1:
							πF.SetLineno(504)
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πg.NewStr("if 1:\n            import sys, os, time, threading\n\n            # a thread, which waits for the main program to terminate\n            def joiningfunc(mainthread):\n                mainthread.join()\n                print 'end of thread'\n        \n").ToObject(), µscript); πE != nil {
								continue
							}
							µscript = πTemp001
							// line 513: p = subprocess.Popen([sys.executable, "-c", script], stdout=subprocess.PIPE)
							πF.SetLineno(513)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = make([]*πg.Object, 3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßexecutable, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewStr("-c").ToObject()
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp003[2] = µscript
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßPIPE, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"stdout", πTemp004},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßPopen, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µp = πTemp001
							// line 514: rc = p.wait()
							πF.SetLineno(514)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µp, ßwait, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrc = πTemp004
							// line 515: data = p.stdout.read().replace('\r', '')
							πF.SetLineno(515)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewStr("\r").ToObject()
							πTemp002[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µp, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdata = πTemp001
							// line 516: p.stdout.close()
							πF.SetLineno(516)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µp, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßclose, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 517: self.assertEqual(data, "end of main\nend of thread\n")
							πF.SetLineno(517)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp002[0] = µdata
							πTemp002[1] = πg.NewStr("end of main\nend of thread\n").ToObject()
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
							// line 518: self.assertFalse(rc == 2, "interpreter was blocked")
							πF.SetLineno(518)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µrc, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = πg.NewStr("interpreter was blocked").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 519: self.assertTrue(rc == 0, "Unexpected error")
							πF.SetLineno(519)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µrc, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							πTemp002[1] = πg.NewStr("Unexpected error").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_run_and_join.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 522: def test_1_join_on_shutdown(self):
					πF.SetLineno(522)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_1_join_on_shutdown", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πg.UnboundLocal; _ = µscript
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
							// line 524: script = """if 1:
							πF.SetLineno(524)
							µscript = πg.NewStr("if 1:\n            import os\n            t = threading.Thread(target=joiningfunc,\n                                 args=(threading.current_thread(),))\n            t.start()\n            time.sleep(0.1)\n            print 'end of main'\n            ").ToObject()
							// line 532: self._run_and_join(script)
							πF.SetLineno(532)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp001[0] = µscript
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_run_and_join, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_1_join_on_shutdown.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 521: @unittest.skip('grumpy')
					πF.SetLineno(521)
					πTemp004 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßtest_1_join_on_shutdown); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßgrumpy.ToObject()
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßskip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_1_join_on_shutdown.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 538: def test_2_join_in_forked_process(self):
					πF.SetLineno(538)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_2_join_in_forked_process", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πg.UnboundLocal; _ = µscript
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
							// line 540: script = """if 1:
							πF.SetLineno(540)
							µscript = πg.NewStr("if 1:\n            childpid = os.fork()\n            if childpid != 0:\n                os.waitpid(childpid, 0)\n                sys.exit(0)\n\n            t = threading.Thread(target=joiningfunc,\n                                 args=(threading.current_thread(),))\n            t.start()\n            print 'end of main'\n            ").ToObject()
							// line 551: self._run_and_join(script)
							πF.SetLineno(551)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp001[0] = µscript
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_run_and_join, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_2_join_in_forked_process.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 535: @unittest.skip('grumpy')
					πF.SetLineno(535)
					πTemp004 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_2_join_in_forked_process); πE != nil {
						continue
					}
					πTemp004[0] = πTemp007
					πTemp006 = πF.MakeArgs(2)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßplatform, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßplatforms_to_skip); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp008, πTemp009); πE != nil {
						continue
					}
					πTemp007 = πg.GetBool(πTemp010).ToObject()
					πTemp006[0] = πTemp007
					πTemp006[1] = πg.NewStr("due to known OS bug").ToObject()
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßskipIf, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_2_join_in_forked_process.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 535: @unittest.skip('grumpy')
					πF.SetLineno(535)
					πTemp004 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_2_join_in_forked_process); πE != nil {
						continue
					}
					πTemp004[0] = πTemp007
					πTemp006 = πF.MakeArgs(2)
					πTemp011 = πF.MakeArgs(2)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					πTemp011[0] = πTemp007
					πTemp011[1] = ßfork.ToObject()
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp006[0] = πTemp008
					πTemp006[1] = πg.NewStr("needs os.fork()").ToObject()
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_2_join_in_forked_process.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 535: @unittest.skip('grumpy')
					πF.SetLineno(535)
					πTemp004 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_2_join_in_forked_process); πE != nil {
						continue
					}
					πTemp004[0] = πTemp007
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßgrumpy.ToObject()
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßskip, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_2_join_in_forked_process.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 556: def test_3_join_in_forked_from_thread(self):
					πF.SetLineno(556)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_3_join_in_forked_from_thread", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πg.UnboundLocal; _ = µscript
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
							// line 559: script = """if 1:
							πF.SetLineno(559)
							µscript = πg.NewStr("if 1:\n            main_thread = threading.current_thread()\n            def worker():\n                childpid = os.fork()\n                if childpid != 0:\n                    os.waitpid(childpid, 0)\n                    sys.exit(0)\n\n                t = threading.Thread(target=joiningfunc,\n                                     args=(main_thread,))\n                print 'end of main'\n                t.start()\n                t.join() # Should not block: main_thread is already stopped\n\n            w = threading.Thread(target=worker)\n            w.start()\n            ").ToObject()
							// line 576: self._run_and_join(script)
							πF.SetLineno(576)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp001[0] = µscript
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_run_and_join, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_3_join_in_forked_from_thread.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 553: @unittest.skip('grumpy')
					πF.SetLineno(553)
					πTemp004 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßtest_3_join_in_forked_from_thread); πE != nil {
						continue
					}
					πTemp004[0] = πTemp008
					πTemp006 = πF.MakeArgs(2)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp009, ßplatform, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßplatforms_to_skip); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp009, πTemp012); πE != nil {
						continue
					}
					πTemp008 = πg.GetBool(πTemp010).ToObject()
					πTemp006[0] = πTemp008
					πTemp006[1] = πg.NewStr("due to known OS bug").ToObject()
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßskipIf, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp009, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_3_join_in_forked_from_thread.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 553: @unittest.skip('grumpy')
					πF.SetLineno(553)
					πTemp004 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßtest_3_join_in_forked_from_thread); πE != nil {
						continue
					}
					πTemp004[0] = πTemp008
					πTemp006 = πF.MakeArgs(2)
					πTemp011 = πF.MakeArgs(2)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					πTemp011[0] = πTemp008
					πTemp011[1] = ßfork.ToObject()
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp006[0] = πTemp009
					πTemp006[1] = πg.NewStr("needs os.fork()").ToObject()
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp009, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_3_join_in_forked_from_thread.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 553: @unittest.skip('grumpy')
					πF.SetLineno(553)
					πTemp004 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßtest_3_join_in_forked_from_thread); πE != nil {
						continue
					}
					πTemp004[0] = πTemp008
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßgrumpy.ToObject()
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßskip, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp009, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_3_join_in_forked_from_thread.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 578: def assertScriptHasOutput(self, script, expected_output):
					πF.SetLineno(578)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "script", Def: nil}
					πTemp002[2] = πg.Param{Name: "expected_output", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("assertScriptHasOutput", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πArgs[1]; _ = µscript
						var µexpected_output *πg.Object = πArgs[2]; _ = µexpected_output
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µrc *πg.Object = πg.UnboundLocal; _ = µrc
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 579: p = subprocess.Popen([sys.executable, "-c", script],
							πF.SetLineno(579)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexecutable, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewStr("-c").ToObject()
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp002[2] = µscript
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"stdout", πTemp004},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPopen, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µp = πTemp003
							// line 581: rc = p.wait()
							πF.SetLineno(581)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßwait, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrc = πTemp004
							// line 582: data = p.stdout.read().decode().replace('\r', '')
							πF.SetLineno(582)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("\r").ToObject()
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßstdout, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßread, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdata = πTemp003
							// line 583: self.assertEqual(rc, 0, "Unexpected error")
							πF.SetLineno(583)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
								continue
							}
							πTemp001[0] = µrc
							πTemp001[1] = πg.NewInt(0).ToObject()
							πTemp001[2] = πg.NewStr("Unexpected error").ToObject()
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
							// line 584: self.assertEqual(data, expected_output)
							πF.SetLineno(584)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πE = πg.CheckLocal(πF, µexpected_output, "expected_output"); πE != nil {
								continue
							}
							πTemp001[1] = µexpected_output
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
					if πE = πClass.SetItem(πF, ßassertScriptHasOutput.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 589: def test_4_joining_across_fork_in_worker_thread(self):
					πF.SetLineno(589)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_4_joining_across_fork_in_worker_thread", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πg.UnboundLocal; _ = µscript
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
							// line 610: script = """if 1:
							πF.SetLineno(610)
							µscript = πg.NewStr("if 1:\n            import os, time, threading\n\n            finish_join = False\n            start_fork = False\n\n            def worker():\n                # Wait until this thread's lock is acquired before forking to\n                # create the deadlock.\n                global finish_join\n                while not start_fork:\n                    time.sleep(0.01)\n                # LOCK HELD: Main thread holds lock across this call.\n                childpid = os.fork()\n                finish_join = True\n                if childpid != 0:\n                    # Parent process just waits for child.\n                    os.waitpid(childpid, 0)\n                # Child process should just return.\n\n            w = threading.Thread(target=worker)\n\n            # Stub out the private condition variable's lock acquire method.\n            # This acquires the lock and then waits until the child has forked\n            # before returning, which will release the lock soon after.  If\n            # someone else tries to fix this test case by acquiring this lock\n            # before forking instead of resetting it, the test case will\n            # deadlock when it shouldn't.\n            condition = w._block\n            orig_acquire = condition.acquire\n            call_count_lock = threading.Lock()\n            call_count = 0\n            def my_acquire():\n                global call_count\n                global start_fork\n                orig_acquire()  # LOCK ACQUIRED HERE\n                start_fork = True\n                if call_count == 0:\n                    while not finish_join:\n                        time.sleep(0.01)  # WORKER THREAD FORKS HERE\n                with call_count_lock:\n                    call_count += 1\n            condition.acquire = my_acquire\n\n            w.start()\n            w.join()\n            print('end of main')\n            ").ToObject()
							// line 658: self.assertScriptHasOutput(script, "end of main\n")
							πF.SetLineno(658)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp001[0] = µscript
							πTemp001[1] = πg.NewStr("end of main\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertScriptHasOutput, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_4_joining_across_fork_in_worker_thread.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 586: @unittest.skip('grumpy')
					πF.SetLineno(586)
					πTemp004 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßtest_4_joining_across_fork_in_worker_thread); πE != nil {
						continue
					}
					πTemp004[0] = πTemp012
					πTemp006 = πF.MakeArgs(2)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßplatform, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßplatforms_to_skip); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp013, πTemp014); πE != nil {
						continue
					}
					πTemp012 = πg.GetBool(πTemp010).ToObject()
					πTemp006[0] = πTemp012
					πTemp006[1] = πg.NewStr("due to known OS bug").ToObject()
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßskipIf, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp013.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp013, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_4_joining_across_fork_in_worker_thread.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 586: @unittest.skip('grumpy')
					πF.SetLineno(586)
					πTemp004 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßtest_4_joining_across_fork_in_worker_thread); πE != nil {
						continue
					}
					πTemp004[0] = πTemp012
					πTemp006 = πF.MakeArgs(2)
					πTemp011 = πF.MakeArgs(2)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					πTemp011[0] = πTemp012
					πTemp011[1] = ßfork.ToObject()
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp012.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp006[0] = πTemp013
					πTemp006[1] = πg.NewStr("needs os.fork()").ToObject()
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp013.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp013, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_4_joining_across_fork_in_worker_thread.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 586: @unittest.skip('grumpy')
					πF.SetLineno(586)
					πTemp004 = πF.MakeArgs(1)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßtest_4_joining_across_fork_in_worker_thread); πE != nil {
						continue
					}
					πTemp004[0] = πTemp012
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßgrumpy.ToObject()
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßskip, nil); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp013.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp013, πE = πTemp012.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_4_joining_across_fork_in_worker_thread.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 663: def test_5_clear_waiter_locks_to_avoid_crash(self):
					πF.SetLineno(663)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_5_clear_waiter_locks_to_avoid_crash", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πg.UnboundLocal; _ = µscript
						var µoutput *πg.Object = πg.UnboundLocal; _ = µoutput
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
							// line 675: script = """if True:
							πF.SetLineno(675)
							µscript = πg.NewStr("if True:\n            import os, time, threading\n\n            start_fork = False\n\n            def worker():\n                # Wait until the main thread has attempted to join this thread\n                # before continuing.\n                while not start_fork:\n                    time.sleep(0.01)\n                childpid = os.fork()\n                if childpid != 0:\n                    # Parent process just waits for child.\n                    (cpid, rc) = os.waitpid(childpid, 0)\n                    assert cpid == childpid\n                    assert rc == 0\n                    print('end of worker thread')\n                else:\n                    # Child process should just return.\n                    pass\n\n            w = threading.Thread(target=worker)\n\n            # Stub out the private condition variable's _release_save method.\n            # This releases the condition's lock and flips the global that\n            # causes the worker to fork.  At this point, the problematic waiter\n            # lock has been acquired once by the waiter and has been put onto\n            # the waiters list.\n            condition = w._block\n            orig_release_save = condition._release_save\n            def my_release_save():\n                global start_fork\n                orig_release_save()\n                # Waiter lock held here, condition lock released.\n                start_fork = True\n            condition._release_save = my_release_save\n\n            w.start()\n            w.join()\n            print('end of main thread')\n            ").ToObject()
							// line 716: output = "end of worker thread\nend of main thread\n"
							πF.SetLineno(716)
							µoutput = πg.NewStr("end of worker thread\nend of main thread\n").ToObject()
							// line 717: self.assertScriptHasOutput(script, output)
							πF.SetLineno(717)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp001[0] = µscript
							if πE = πg.CheckLocal(πF, µoutput, "output"); πE != nil {
								continue
							}
							πTemp001[1] = µoutput
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertScriptHasOutput, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_5_clear_waiter_locks_to_avoid_crash.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 660: @unittest.skip('grumpy')
					πF.SetLineno(660)
					πTemp004 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßtest_5_clear_waiter_locks_to_avoid_crash); πE != nil {
						continue
					}
					πTemp004[0] = πTemp013
					πTemp006 = πF.MakeArgs(2)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßplatform, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßplatforms_to_skip); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp014, πTemp015); πE != nil {
						continue
					}
					πTemp013 = πg.GetBool(πTemp010).ToObject()
					πTemp006[0] = πTemp013
					πTemp006[1] = πg.NewStr("due to known OS bug").ToObject()
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßskipIf, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp014.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp014, πE = πTemp013.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_5_clear_waiter_locks_to_avoid_crash.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 660: @unittest.skip('grumpy')
					πF.SetLineno(660)
					πTemp004 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßtest_5_clear_waiter_locks_to_avoid_crash); πE != nil {
						continue
					}
					πTemp004[0] = πTemp013
					πTemp006 = πF.MakeArgs(2)
					πTemp011 = πF.MakeArgs(2)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					πTemp011[0] = πTemp013
					πTemp011[1] = ßfork.ToObject()
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp013.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp006[0] = πTemp014
					πTemp006[1] = πg.NewStr("needs os.fork()").ToObject()
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp014.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp014, πE = πTemp013.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_5_clear_waiter_locks_to_avoid_crash.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 660: @unittest.skip('grumpy')
					πF.SetLineno(660)
					πTemp004 = πF.MakeArgs(1)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßtest_5_clear_waiter_locks_to_avoid_crash); πE != nil {
						continue
					}
					πTemp004[0] = πTemp013
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = ßgrumpy.ToObject()
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßskip, nil); πE != nil {
						continue
					}
					if πTemp013, πE = πTemp014.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp014, πE = πTemp013.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_5_clear_waiter_locks_to_avoid_crash.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 721: def test_reinit_tls_after_fork(self):
					πF.SetLineno(721)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_reinit_tls_after_fork", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdo_fork_and_wait *πg.Object = πg.UnboundLocal; _ = µdo_fork_and_wait
						var µthreads *πg.Object = πg.UnboundLocal; _ = µthreads
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []πg.Param
						_ = πTemp002
						var πTemp003 []*πg.Object
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
						var πTemp009 πg.KWArgs
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
							// line 725: def do_fork_and_wait():
							πF.SetLineno(725)
							πTemp002 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("do_fork_and_wait", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µpid *πg.Object = πg.UnboundLocal; _ = µpid
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 bool
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
									// line 727: pid = os.fork()
									πF.SetLineno(727)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßfork, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									µpid = πTemp001
									if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GT(πF, µpid, πg.NewInt(0).ToObject()); πE != nil {
										continue
									}
									if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label1
									}
									goto Label2
									// line 728: if pid > 0:
									πF.SetLineno(728)
								Label1:
									// line 729: os.waitpid(pid, 0)
									πF.SetLineno(729)
									πTemp004 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
										continue
									}
									πTemp004[0] = µpid
									πTemp004[1] = πg.NewInt(0).ToObject()
									if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwaitpid, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									goto Label3
								Label2:
									// line 731: os._exit(0)
									πF.SetLineno(731)
									πTemp004 = πF.MakeArgs(1)
									πTemp004[0] = πg.NewInt(0).ToObject()
									if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_exit, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
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
							µdo_fork_and_wait = πTemp001
							// line 734: threads = []
							πF.SetLineno(734)
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							µthreads = πTemp004
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(16).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.Iter(πF, πTemp006); πE != nil {
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
							if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µi = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 736: t = threading.Thread(target=do_fork_and_wait)
							πF.SetLineno(736)
							if πE = πg.CheckLocal(πF, µdo_fork_and_wait, "do_fork_and_wait"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"target", µdo_fork_and_wait},
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßThread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, πTemp009); πE != nil {
								continue
							}
							µt = πTemp005
							// line 737: threads.append(t)
							πF.SetLineno(737)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp003[0] = µt
							if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µthreads, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 738: t.start()
							πF.SetLineno(738)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, µthreads); πE != nil {
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
							if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µt = πTemp005
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 741: t.join()
							πF.SetLineno(741)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßtest_reinit_tls_after_fork.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 719: @unittest.skipUnless(hasattr(os, 'fork'), "needs os.fork()")
					πF.SetLineno(719)
					πTemp004 = πF.MakeArgs(1)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßtest_reinit_tls_after_fork); πE != nil {
						continue
					}
					πTemp004[0] = πTemp014
					πTemp006 = πF.MakeArgs(2)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßplatform, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßplatforms_to_skip); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp015, πTemp016); πE != nil {
						continue
					}
					πTemp014 = πg.GetBool(πTemp010).ToObject()
					πTemp006[0] = πTemp014
					πTemp006[1] = πg.NewStr("due to known OS bug").ToObject()
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßskipIf, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp015.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp015, πE = πTemp014.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_reinit_tls_after_fork.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 719: @unittest.skipUnless(hasattr(os, 'fork'), "needs os.fork()")
					πF.SetLineno(719)
					πTemp004 = πF.MakeArgs(1)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßtest_reinit_tls_after_fork); πE != nil {
						continue
					}
					πTemp004[0] = πTemp014
					πTemp006 = πF.MakeArgs(2)
					πTemp011 = πF.MakeArgs(2)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßos); πE != nil {
						continue
					}
					πTemp011[0] = πTemp014
					πTemp011[1] = ßfork.ToObject()
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßhasattr); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp006[0] = πTemp015
					πTemp006[1] = πg.NewStr("needs os.fork()").ToObject()
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp014, πE = πTemp015.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp015, πE = πTemp014.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_reinit_tls_after_fork.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 745: def test_frame_tstate_tracing(self):
					πF.SetLineno(745)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_frame_tstate_tracing", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnoop_trace *πg.Object = πg.UnboundLocal; _ = µnoop_trace
						var µgenerator *πg.Object = πg.UnboundLocal; _ = µgenerator
						var µcallback *πg.Object = πg.UnboundLocal; _ = µcallback
						var µold_trace *πg.Object = πg.UnboundLocal; _ = µold_trace
						var µtest *πg.Object = πg.UnboundLocal; _ = µtest
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
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
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
							// line 752: def noop_trace(frame, event, arg):
							πF.SetLineno(752)
							πTemp002 = make([]πg.Param, 3)
							πTemp002[0] = πg.Param{Name: "frame", Def: nil}
							πTemp002[1] = πg.Param{Name: "event", Def: nil}
							πTemp002[2] = πg.Param{Name: "arg", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("noop_trace", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µframe *πg.Object = πArgs[0]; _ = µframe
								var µevent *πg.Object = πArgs[1]; _ = µevent
								var µarg *πg.Object = πArgs[2]; _ = µarg
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 754: return noop_trace
									πF.SetLineno(754)
									if πE = πg.CheckLocal(πF, µnoop_trace, "noop_trace"); πE != nil {
										continue
									}
									πR = µnoop_trace
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µnoop_trace = πTemp001
							// line 756: def generator():
							πF.SetLineno(756)
							πTemp002 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("generator", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var πTemp001 bool
								_ = πTemp001
								var πTemp002 bool
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
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
										// line 757: while 1:
										πF.SetLineno(757)
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
										if πTemp002, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
											continue
										}
										if πE != nil || !πTemp002 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 758: yield "generator"
										πF.SetLineno(758)
										πF.PushCheckpoint(4)
										return ßgenerator.ToObject(), nil
									Label4:
										πTemp003 = πSent
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
							µgenerator = πTemp003
							// line 760: def callback():
							πF.SetLineno(760)
							πTemp002 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("callback", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µcallback, ßgen, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
									if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp004 {
										goto Label1
									}
									goto Label2
									// line 761: if callback.gen is None:
									πF.SetLineno(761)
								Label1:
									// line 762: callback.gen = generator()
									πF.SetLineno(762)
									if πE = πg.CheckLocal(πF, µgenerator, "generator"); πE != nil {
										continue
									}
									if πTemp001, πE = µgenerator.Call(πF, nil, nil); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µcallback, ßgen, πTemp002); πE != nil {
										continue
									}
									goto Label2
								Label2:
									// line 763: return next(callback.gen)
									πF.SetLineno(763)
									πTemp005 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µcallback, ßgen, nil); πE != nil {
										continue
									}
									πTemp005[0] = πTemp001
									if πTemp001, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp005)
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
							µcallback = πTemp004
							// line 764: callback.gen = None
							πF.SetLineno(764)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcallback, ßgen, πTemp006); πE != nil {
								continue
							}
							// line 766: old_trace = sys.gettrace()
							πF.SetLineno(766)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßgettrace, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							µold_trace = πTemp005
							// line 767: sys.settrace(noop_trace)
							πF.SetLineno(767)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnoop_trace, "noop_trace"); πE != nil {
								continue
							}
							πTemp007[0] = µnoop_trace
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsettrace, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 768: try:
							πF.SetLineno(768)
							πF.PushCheckpoint(1)
							// line 770: threading.settrace(noop_trace)
							πF.SetLineno(770)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnoop_trace, "noop_trace"); πE != nil {
								continue
							}
							πTemp007[0] = µnoop_trace
							if πTemp005, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsettrace, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 773: _testcapi.call_in_temporary_c_thread(callback)
							πF.SetLineno(773)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
								continue
							}
							πTemp007[0] = µcallback
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_testcapi); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßcall_in_temporary_c_thread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewInt(3).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp005, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							πTemp009 = false
						Label2:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label4
							}
							if πTemp006, πE = πg.Next(πF, πTemp005); πE != nil {
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
								µtest = πTemp006
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(2)            
							// line 779: callback()
							πF.SetLineno(779)
							if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
								continue
							}
							if πTemp006, πE = µcallback.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label3:
							if πE != nil || πR != nil {
								continue
							}
						Label4:
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp011, πTemp012 = πF.RestoreExc(nil, nil)
							// line 781: sys.settrace(old_trace)
							πF.SetLineno(781)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µold_trace, "old_trace"); πE != nil {
								continue
							}
							πTemp007[0] = µold_trace
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßsettrace, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp011 != nil {
								πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
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
					if πE = πClass.SetItem(πF, ßtest_frame_tstate_tracing.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 743: @cpython_only
					πF.SetLineno(743)
					πTemp004 = πF.MakeArgs(1)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßtest_frame_tstate_tracing); πE != nil {
						continue
					}
					πTemp004[0] = πTemp015
					πTemp006 = πF.MakeArgs(2)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ß_testcapi); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp015 = πg.GetBool(πTemp016 == πTemp017).ToObject()
					πTemp006[0] = πTemp015
					πTemp006[1] = πg.NewStr("need _testcapi module").ToObject()
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßskipIf, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp016.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp016, πE = πTemp015.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_frame_tstate_tracing.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 743: @cpython_only
					πF.SetLineno(743)
					πTemp004 = πF.MakeArgs(1)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßtest_frame_tstate_tracing); πE != nil {
						continue
					}
					πTemp004[0] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßcpython_only); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp015.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_frame_tstate_tracing.ToObject(), πTemp016); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ThreadJoinOnShutdown").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßThreadJoinOnShutdown.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 784: class ThreadingExceptionTests(BaseTestCase):
			πF.SetLineno(784)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseTestCase); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ThreadingExceptionTests", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 787: def test_start_thread_again(self):
					πF.SetLineno(787)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_start_thread_again", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µthread *πg.Object = πg.UnboundLocal; _ = µthread
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
							// line 788: thread = threading.Thread()
							πF.SetLineno(788)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßThread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µthread = πTemp001
							// line 789: thread.start()
							πF.SetLineno(789)
							if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µthread, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 790: self.assertRaises(RuntimeError, thread.start)
							πF.SetLineno(790)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µthread, ßstart, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_start_thread_again.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 792: def test_joining_current_thread(self):
					πF.SetLineno(792)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_joining_current_thread", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcurrent_thread *πg.Object = πg.UnboundLocal; _ = µcurrent_thread
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
							// line 793: current_thread = threading.current_thread()
							πF.SetLineno(793)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcurrent_thread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcurrent_thread = πTemp001
							// line 794: self.assertRaises(RuntimeError, current_thread.join);
							πF.SetLineno(794)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µcurrent_thread, "current_thread"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcurrent_thread, ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_joining_current_thread.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 796: def test_joining_inactive_thread(self):
					πF.SetLineno(796)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_joining_inactive_thread", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µthread *πg.Object = πg.UnboundLocal; _ = µthread
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
							// line 797: thread = threading.Thread()
							πF.SetLineno(797)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßThread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µthread = πTemp001
							// line 798: self.assertRaises(RuntimeError, thread.join)
							πF.SetLineno(798)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µthread, ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_joining_inactive_thread.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 800: def test_daemonize_active_thread(self):
					πF.SetLineno(800)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_daemonize_active_thread", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µthread *πg.Object = πg.UnboundLocal; _ = µthread
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
							// line 801: thread = threading.Thread()
							πF.SetLineno(801)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßThread, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µthread = πTemp001
							// line 802: thread.start()
							πF.SetLineno(802)
							if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µthread, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 803: self.assertRaises(RuntimeError, setattr, thread, "daemon", True)
							πF.SetLineno(803)
							πTemp003 = πF.MakeArgs(5)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
								continue
							}
							πTemp003[2] = µthread
							πTemp003[3] = ßdaemon.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp003[4] = πTemp001
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
					if πE = πClass.SetItem(πF, ßtest_daemonize_active_thread.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 806: def test_print_exception(self):
					πF.SetLineno(806)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_print_exception", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πg.UnboundLocal; _ = µscript
						var µrc *πg.Object = πg.UnboundLocal; _ = µrc
						var µout *πg.Object = πg.UnboundLocal; _ = µout
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 807: script = r"""if 1:
							πF.SetLineno(807)
							µscript = πg.NewStr("if 1:\n            import threading\n            import time\n\n            running = False\n            def run():\n                global running\n                running = True\n                while running:\n                    time.sleep(0.01)\n                1.0/0.0\n            t = threading.Thread(target=run)\n            t.start()\n            while not running:\n                time.sleep(0.01)\n            running = False\n            t.join()\n            ").ToObject()
							// line 825: rc, out, err = assert_python_ok("-c", script)
							πF.SetLineno(825)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("-c").ToObject()
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp001[1] = µscript
							if πTemp002, πE = πg.ResolveGlobal(πF, ßassert_python_ok); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µrc = πTemp002
							µout = πTemp004
							µerr = πTemp005
							// line 826: self.assertEqual(out, '')
							πF.SetLineno(826)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[0] = µout
							πTemp001[1] = ß.ToObject()
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
							// line 827: self.assertIn("Exception in thread", err)
							πF.SetLineno(827)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("Exception in thread").ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 828: self.assertIn("Traceback (most recent call last):", err)
							πF.SetLineno(828)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("Traceback (most recent call last):").ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 829: self.assertIn("ZeroDivisionError", err)
							πF.SetLineno(829)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßZeroDivisionError.ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 830: self.assertNotIn("Unhandled exception", err)
							πF.SetLineno(830)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("Unhandled exception").ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_print_exception.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 805: @unittest.skip('grumpy')
					πF.SetLineno(805)
					πTemp007 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßtest_print_exception); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_print_exception.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 833: def test_print_exception_stderr_is_none_1(self):
					πF.SetLineno(833)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_print_exception_stderr_is_none_1", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πg.UnboundLocal; _ = µscript
						var µrc *πg.Object = πg.UnboundLocal; _ = µrc
						var µout *πg.Object = πg.UnboundLocal; _ = µout
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 834: script = r"""if 1:
							πF.SetLineno(834)
							µscript = πg.NewStr("if 1:\n            import sys\n            import threading\n            import time\n\n            running = False\n            def run():\n                global running\n                running = True\n                while running:\n                    time.sleep(0.01)\n                1.0/0.0\n            t = threading.Thread(target=run)\n            t.start()\n            while not running:\n                time.sleep(0.01)\n            sys.stderr = None\n            running = False\n            t.join()\n            ").ToObject()
							// line 854: rc, out, err = assert_python_ok("-c", script)
							πF.SetLineno(854)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("-c").ToObject()
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp001[1] = µscript
							if πTemp002, πE = πg.ResolveGlobal(πF, ßassert_python_ok); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µrc = πTemp002
							µout = πTemp004
							µerr = πTemp005
							// line 855: self.assertEqual(out, '')
							πF.SetLineno(855)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[0] = µout
							πTemp001[1] = ß.ToObject()
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
							// line 856: self.assertIn("Exception in thread", err)
							πF.SetLineno(856)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("Exception in thread").ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 857: self.assertIn("Traceback (most recent call last):", err)
							πF.SetLineno(857)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("Traceback (most recent call last):").ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 858: self.assertIn("ZeroDivisionError", err)
							πF.SetLineno(858)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßZeroDivisionError.ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 859: self.assertNotIn("Unhandled exception", err)
							πF.SetLineno(859)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("Unhandled exception").ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_print_exception_stderr_is_none_1.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 832: @unittest.skip('grumpy')
					πF.SetLineno(832)
					πTemp007 = πF.MakeArgs(1)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßtest_print_exception_stderr_is_none_1); πE != nil {
						continue
					}
					πTemp007[0] = πTemp010
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = ßgrumpy.ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßskip, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp011, πE = πTemp010.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πClass.SetItem(πF, ßtest_print_exception_stderr_is_none_1.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 862: def test_print_exception_stderr_is_none_2(self):
					πF.SetLineno(862)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_print_exception_stderr_is_none_2", "/usr/lib/python2.7/test/test_threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πg.UnboundLocal; _ = µscript
						var µrc *πg.Object = πg.UnboundLocal; _ = µrc
						var µout *πg.Object = πg.UnboundLocal; _ = µout
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 863: script = r"""if 1:
							πF.SetLineno(863)
							µscript = πg.NewStr("if 1:\n            import sys\n            import threading\n            import time\n\n            running = False\n            def run():\n                global running\n                running = True\n                while running:\n                    time.sleep(0.01)\n                1.0/0.0\n            sys.stderr = None\n            t = threading.Thread(target=run)\n            t.start()\n            while not running:\n                time.sleep(0.01)\n            running = False\n            t.join()\n            ").ToObject()
							// line 883: rc, out, err = assert_python_ok("-c", script)
							πF.SetLineno(883)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("-c").ToObject()
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp001[1] = µscript
							if πTemp002, πE = πg.ResolveGlobal(πF, ßassert_python_ok); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µrc = πTemp002
							µout = πTemp004
							µerr = πTemp005
							// line 884: self.assertEqual(out, '')
							πF.SetLineno(884)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
								continue
							}
							πTemp001[0] = µout
							πTemp001[1] = ß.ToObject()
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
							// line 885: self.assertNotIn("Unhandled exception", err)
							πF.SetLineno(885)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("Unhandled exception").ToObject()
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[1] = µerr
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_print_exception_stderr_is_none_2.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 861: @unittest.skip('grumpy')
					πF.SetLineno(861)
					πTemp007 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtest_print_exception_stderr_is_none_2); πE != nil {
						continue
					}
					πTemp007[0] = πTemp011
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = ßgrumpy.ToObject()
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßskip, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp012.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp012, πE = πTemp011.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πClass.SetItem(πF, ßtest_print_exception_stderr_is_none_2.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ThreadingExceptionTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßThreadingExceptionTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 888: class LockTests(lock_tests.LockTests):
			πF.SetLineno(888)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlock_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßLockTests, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("LockTests", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 889: locktype = staticmethod(threading.Lock)
					πF.SetLineno(889)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßthreading); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßLock, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßstaticmethod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßlocktype.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("LockTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLockTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 891: class RLockTests(lock_tests.RLockTests):
			πF.SetLineno(891)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlock_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßRLockTests, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("RLockTests", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 892: locktype = staticmethod(threading.RLock)
					πF.SetLineno(892)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßthreading); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßRLock, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßstaticmethod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßlocktype.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("RLockTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRLockTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 894: class EventTests(lock_tests.EventTests):
			πF.SetLineno(894)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlock_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßEventTests, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("EventTests", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 895: eventtype = staticmethod(threading.Event)
					πF.SetLineno(895)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßthreading); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßEvent, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßstaticmethod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßeventtype.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("EventTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßEventTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 897: class ConditionAsRLockTests(lock_tests.RLockTests):
			πF.SetLineno(897)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlock_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßRLockTests, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ConditionAsRLockTests", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 899: locktype = staticmethod(threading.Condition)
					πF.SetLineno(899)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßthreading); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßCondition, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßstaticmethod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßlocktype.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ConditionAsRLockTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßConditionAsRLockTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 901: class ConditionTests(lock_tests.ConditionTests):
			πF.SetLineno(901)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlock_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßConditionTests, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ConditionTests", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 902: condtype = staticmethod(threading.Condition)
					πF.SetLineno(902)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßthreading); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßCondition, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßstaticmethod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßcondtype.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ConditionTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßConditionTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 904: class SemaphoreTests(lock_tests.SemaphoreTests):
			πF.SetLineno(904)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlock_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßSemaphoreTests, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SemaphoreTests", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 905: semtype = staticmethod(threading.Semaphore)
					πF.SetLineno(905)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßthreading); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßSemaphore, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßstaticmethod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßsemtype.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SemaphoreTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSemaphoreTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 907: class BoundedSemaphoreTests(lock_tests.BoundedSemaphoreTests):
			πF.SetLineno(907)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßlock_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßBoundedSemaphoreTests, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BoundedSemaphoreTests", "/usr/lib/python2.7/test/test_threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
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
				var πTemp007 *πg.Object
				_ = πTemp007
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 908: semtype = staticmethod(threading.BoundedSemaphore)
					πF.SetLineno(908)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßthreading); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßBoundedSemaphore, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßstaticmethod); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßsemtype.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 912: def test_recursion_limit(self):
					πF.SetLineno(912)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("test_recursion_limit", "/usr/lib/python2.7/test/test_threading.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µscript *πg.Object = πg.UnboundLocal; _ = µscript
						var µexpected_output *πg.Object = πg.UnboundLocal; _ = µexpected_output
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µstdout *πg.Object = πg.UnboundLocal; _ = µstdout
						var µstderr *πg.Object = πg.UnboundLocal; _ = µstderr
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
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
							// line 918: script = """if True:
							πF.SetLineno(918)
							µscript = πg.NewStr("if True:\n            import threading\n\n            def recurse():\n                return recurse()\n\n            def outer():\n                try:\n                    recurse()\n                except RuntimeError:\n                    pass\n\n            w = threading.Thread(target=outer)\n            w.start()\n            w.join()\n            print('end of main thread')\n            ").ToObject()
							// line 935: expected_output = "end of main thread\n"
							πF.SetLineno(935)
							µexpected_output = πg.NewStr("end of main thread\n").ToObject()
							// line 936: p = subprocess.Popen([sys.executable, "-c", script],
							πF.SetLineno(936)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexecutable, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewStr("-c").ToObject()
							if πE = πg.CheckLocal(πF, µscript, "script"); πE != nil {
								continue
							}
							πTemp002[2] = µscript
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPIPE, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"stdout", πTemp004},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßPopen, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µp = πTemp003
							// line 938: stdout, stderr = p.communicate()
							πF.SetLineno(938)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßcommunicate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
								continue
							}
							µstdout = πTemp003
							µstderr = πTemp006
							// line 939: data = stdout.decode().replace('\r', '')
							πF.SetLineno(939)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("\r").ToObject()
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µstdout, "stdout"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstdout, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdata = πTemp004
							// line 940: self.assertEqual(p.returncode, 0, "Unexpected error")
							πF.SetLineno(940)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µp, ßreturncode, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(0).ToObject()
							πTemp001[2] = πg.NewStr("Unexpected error").ToObject()
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
							// line 941: self.assertEqual(data, expected_output)
							πF.SetLineno(941)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πE = πg.CheckLocal(πF, µexpected_output, "expected_output"); πE != nil {
								continue
							}
							πTemp001[1] = µexpected_output
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
					if πE = πClass.SetItem(πF, ßtest_recursion_limit.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 910: @unittest.skip('grumpy')
					πF.SetLineno(910)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßtest_recursion_limit); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp005 = πF.MakeArgs(2)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßplatform, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp007, ßdarwin.ToObject()); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					πTemp005[1] = πg.NewStr("test macosx problem").ToObject()
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßtest_recursion_limit.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 910: @unittest.skip('grumpy')
					πF.SetLineno(910)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßtest_recursion_limit); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ßgrumpy.ToObject()
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßskip, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßtest_recursion_limit.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BoundedSemaphoreTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBoundedSemaphoreTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 943: def test_main():
			πF.SetLineno(943)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "/usr/lib/python2.7/test/test_threading.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 944: test.test_support.run_unittest(LockTests, RLockTests, EventTests,
					πF.SetLineno(944)
					πTemp001 = πF.MakeArgs(10)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLockTests); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßRLockTests); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßEventTests); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßConditionAsRLockTests); πE != nil {
						continue
					}
					πTemp001[3] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßConditionTests); πE != nil {
						continue
					}
					πTemp001[4] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSemaphoreTests); πE != nil {
						continue
					}
					πTemp001[5] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßBoundedSemaphoreTests); πE != nil {
						continue
					}
					πTemp001[6] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßThreadTests); πE != nil {
						continue
					}
					πTemp001[7] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßThreadJoinOnShutdown); πE != nil {
						continue
					}
					πTemp001[8] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßThreadingExceptionTests); πE != nil {
						continue
					}
					πTemp001[9] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_support, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßrun_unittest, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Eq(πF, πTemp005, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label1
			}
			goto Label2
			// line 952: if __name__ == "__main__":
			πF.SetLineno(952)
		Label1:
			// line 953: test_main()
			πF.SetLineno(953)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_threading", Code)
}

