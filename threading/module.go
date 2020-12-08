package threading
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/random"
	// _ "github.com/pygolin/stdlib/thread"
	// _ "github.com/pygolin/stdlib/warnings"
	// _ "github.com/pygolin/stdlib/traceback"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/itertools"
	// _ "github.com/pygolin/stdlib/collections"
	// _ "github.com/pygolin/stdlib/time"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBoundedSemaphore := πg.InternStr("BoundedSemaphore")
		ßCondition := πg.InternStr("Condition")
		ßConsumer := πg.InternStr("Consumer")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßEvent := πg.InternStr("Event")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßImportError := πg.InternStr("ImportError")
		ßKeyError := πg.InternStr("KeyError")
		ßLock := πg.InternStr("Lock")
		ßMainThread := πg.InternStr("MainThread")
		ßNone := πg.InternStr("None")
		ßProducer := πg.InternStr("Producer")
		ßRLock := πg.InternStr("RLock")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßSemaphore := πg.InternStr("Semaphore")
		ßSystemExit := πg.InternStr("SystemExit")
		ßThread := πg.InternStr("Thread")
		ßThreadError := πg.InternStr("ThreadError")
		ßTimer := πg.InternStr("Timer")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß_BoundedSemaphore := πg.InternStr("_BoundedSemaphore")
		ß_Condition := πg.InternStr("_Condition")
		ß_DummyThread := πg.InternStr("_DummyThread")
		ß_Event := πg.InternStr("_Event")
		ß_MainThread := πg.InternStr("_MainThread")
		ß_RLock := πg.InternStr("_RLock")
		ß_Semaphore := πg.InternStr("_Semaphore")
		ß_Timer := πg.InternStr("_Timer")
		ß_VERBOSE := πg.InternStr("_VERBOSE")
		ß_Verbose := πg.InternStr("_Verbose")
		ß__all__ := πg.InternStr("__all__")
		ß__args := πg.InternStr("__args")
		ß__block := πg.InternStr("__block")
		ß__bootstrap := πg.InternStr("__bootstrap")
		ß__bootstrap_inner := πg.InternStr("__bootstrap_inner")
		ß__class__ := πg.InternStr("__class__")
		ß__cond := πg.InternStr("__cond")
		ß__count := πg.InternStr("__count")
		ß__daemonic := πg.InternStr("__daemonic")
		ß__debug__ := πg.InternStr("__debug__")
		ß__del__ := πg.InternStr("__del__")
		ß__delattr__ := πg.InternStr("__delattr__")
		ß__delete := πg.InternStr("__delete")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__flag := πg.InternStr("__flag")
		ß__getattribute__ := πg.InternStr("__getattribute__")
		ß__ident := πg.InternStr("__ident")
		ß__init__ := πg.InternStr("__init__")
		ß__initialized := πg.InternStr("__initialized")
		ß__kwargs := πg.InternStr("__kwargs")
		ß__lock := πg.InternStr("__lock")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name := πg.InternStr("__name")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__owner := πg.InternStr("__owner")
		ß__repr__ := πg.InternStr("__repr__")
		ß__setattr__ := πg.InternStr("__setattr__")
		ß__slots__ := πg.InternStr("__slots__")
		ß__started := πg.InternStr("__started")
		ß__stderr := πg.InternStr("__stderr")
		ß__stop := πg.InternStr("__stop")
		ß__stopped := πg.InternStr("__stopped")
		ß__target := πg.InternStr("__target")
		ß__value := πg.InternStr("__value")
		ß__verbose := πg.InternStr("__verbose")
		ß__waiters := πg.InternStr("__waiters")
		ß_acquire_restore := πg.InternStr("_acquire_restore")
		ß_active := πg.InternStr("_active")
		ß_active_limbo_lock := πg.InternStr("_active_limbo_lock")
		ß_after_fork := πg.InternStr("_after_fork")
		ß_allocate_lock := πg.InternStr("_allocate_lock")
		ß_block := πg.InternStr("_block")
		ß_count := πg.InternStr("_count")
		ß_counter := πg.InternStr("_counter")
		ß_daemon_getter := πg.InternStr("_daemon_getter")
		ß_daemon_setter := πg.InternStr("_daemon_setter")
		ß_deque := πg.InternStr("_deque")
		ß_enumerate := πg.InternStr("_enumerate")
		ß_exitfunc := πg.InternStr("_exitfunc")
		ß_format_exc := πg.InternStr("_format_exc")
		ß_get_ident := πg.InternStr("_get_ident")
		ß_initial_value := πg.InternStr("_initial_value")
		ß_is_owned := πg.InternStr("_is_owned")
		ß_limbo := πg.InternStr("_limbo")
		ß_local__args := πg.InternStr("_local__args")
		ß_local__key := πg.InternStr("_local__key")
		ß_local__lock := πg.InternStr("_local__lock")
		ß_localbase := πg.InternStr("_localbase")
		ß_name_getter := πg.InternStr("_name_getter")
		ß_name_setter := πg.InternStr("_name_setter")
		ß_newname := πg.InternStr("_newname")
		ß_note := πg.InternStr("_note")
		ß_patch := πg.InternStr("_patch")
		ß_pickSomeNonDaemonThread := πg.InternStr("_pickSomeNonDaemonThread")
		ß_profile_hook := πg.InternStr("_profile_hook")
		ß_release_save := πg.InternStr("_release_save")
		ß_reset_internal_locks := πg.InternStr("_reset_internal_locks")
		ß_set_daemon := πg.InternStr("_set_daemon")
		ß_set_ident := πg.InternStr("_set_ident")
		ß_shutdown := πg.InternStr("_shutdown")
		ß_sleep := πg.InternStr("_sleep")
		ß_start_new_thread := πg.InternStr("_start_new_thread")
		ß_sys := πg.InternStr("_sys")
		ß_test := πg.InternStr("_test")
		ß_time := πg.InternStr("_time")
		ß_trace_hook := πg.InternStr("_trace_hook")
		ßacquire := πg.InternStr("acquire")
		ßactiveCount := πg.InternStr("activeCount")
		ßactive_count := πg.InternStr("active_count")
		ßallocate_lock := πg.InternStr("allocate_lock")
		ßappend := πg.InternStr("append")
		ßargs := πg.InternStr("args")
		ßcancel := πg.InternStr("cancel")
		ßclear := πg.InternStr("clear")
		ßco_filename := πg.InternStr("co_filename")
		ßco_name := πg.InternStr("co_name")
		ßcount := πg.InternStr("count")
		ßcurrentThread := πg.InternStr("currentThread")
		ßcurrent_thread := πg.InternStr("current_thread")
		ßdaemon := πg.InternStr("daemon")
		ßdeque := πg.InternStr("deque")
		ßdummy_threading := πg.InternStr("dummy_threading")
		ßenumerate := πg.InternStr("enumerate")
		ßerror := πg.InternStr("error")
		ßexc_clear := πg.InternStr("exc_clear")
		ßexc_info := πg.InternStr("exc_info")
		ßf_code := πg.InternStr("f_code")
		ßfilterwarnings := πg.InternStr("filterwarnings")
		ßfinished := πg.InternStr("finished")
		ßformat_exc := πg.InternStr("format_exc")
		ßfunction := πg.InternStr("function")
		ßget := πg.InternStr("get")
		ßgetName := πg.InternStr("getName")
		ßget_ident := πg.InternStr("get_ident")
		ßhasattr := πg.InternStr("hasattr")
		ßid := πg.InternStr("id")
		ßident := πg.InternStr("ident")
		ßignore := πg.InternStr("ignore")
		ßinitial := πg.InternStr("initial")
		ßinterval := πg.InternStr("interval")
		ßisAlive := πg.InternStr("isAlive")
		ßisDaemon := πg.InternStr("isDaemon")
		ßisSet := πg.InternStr("isSet")
		ßis_alive := πg.InternStr("is_alive")
		ßis_set := πg.InternStr("is_set")
		ßjoin := πg.InternStr("join")
		ßkwargs := πg.InternStr("kwargs")
		ßlen := πg.InternStr("len")
		ßlimit := πg.InternStr("limit")
		ßlocal := πg.InternStr("local")
		ßmin := πg.InternStr("min")
		ßmodules := πg.InternStr("modules")
		ßmon := πg.InternStr("mon")
		ßname := πg.InternStr("name")
		ßnext := πg.InternStr("next")
		ßnotify := πg.InternStr("notify")
		ßnotifyAll := πg.InternStr("notifyAll")
		ßnotify_all := πg.InternStr("notify_all")
		ßobject := πg.InternStr("object")
		ßpopleft := πg.InternStr("popleft")
		ßproperty := πg.InternStr("property")
		ßput := πg.InternStr("put")
		ßqueue := πg.InternStr("queue")
		ßquota := πg.InternStr("quota")
		ßrandom := πg.InternStr("random")
		ßrange := πg.InternStr("range")
		ßrc := πg.InternStr("rc")
		ßrelease := πg.InternStr("release")
		ßremove := πg.InternStr("remove")
		ßrun := πg.InternStr("run")
		ßs := πg.InternStr("s")
		ßset := πg.InternStr("set")
		ßsetDaemon := πg.InternStr("setDaemon")
		ßsetName := πg.InternStr("setName")
		ßsetprofile := πg.InternStr("setprofile")
		ßsettrace := πg.InternStr("settrace")
		ßsleep := πg.InternStr("sleep")
		ßstack_size := πg.InternStr("stack_size")
		ßstart := πg.InternStr("start")
		ßstart_new_thread := πg.InternStr("start_new_thread")
		ßstarted := πg.InternStr("started")
		ßstderr := πg.InternStr("stderr")
		ßstopped := πg.InternStr("stopped")
		ßstr := πg.InternStr("str")
		ßtb_frame := πg.InternStr("tb_frame")
		ßtb_lineno := πg.InternStr("tb_lineno")
		ßtb_next := πg.InternStr("tb_next")
		ßthread := πg.InternStr("thread")
		ßthreading := πg.InternStr("threading")
		ßtime := πg.InternStr("time")
		ßtype := πg.InternStr("type")
		ßupdate := πg.InternStr("update")
		ßvalues := πg.InternStr("values")
		ßwait := πg.InternStr("wait")
		ßwarnings := πg.InternStr("warnings")
		ßwc := πg.InternStr("wc")
		ßwrite := πg.InternStr("write")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.BaseException
		_ = πTemp003
		var πTemp004 *πg.Traceback
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 πg.KWArgs
		_ = πTemp008
		var πTemp009 *πg.Dict
		_ = πTemp009
		var πTemp010 []πg.Param
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
		var πTemp025 *πg.Object
		_ = πTemp025
		var πTemp026 *πg.Object
		_ = πTemp026
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """Thread module emulating a subset of Java's threading model."""
			πF.SetLineno(1)
			// line 1: """Thread module emulating a subset of Java's threading model."""
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Thread module emulating a subset of Java's threading model.").ToObject()); πE != nil {
				continue
			}
			// line 3: import sys as _sys
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_sys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: try:
			πF.SetLineno(5)
			πF.PushCheckpoint(2)
			// line 6: import thread
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "thread"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthread.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label3
			}
			πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
			continue
			// line 7: except ImportError:
			πF.SetLineno(7)
		Label3:
			// line 8: del _sys.modules[__name__]
			πF.SetLineno(8)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßmodules, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			πTemp001 = πTemp007
			if πE = πg.DelItem(πF, πTemp006, πTemp001); πE != nil {
				continue
			}
			// line 9: raise
			πF.SetLineno(9)
			πE = πF.Raise(nil, nil, nil)
			continue
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 11: import warnings
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: from collections import deque as _deque
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "collections"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp006, πE = πg.GetAttrImport(πF, πTemp001, ßdeque); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_deque.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 14: from itertools import count as _count
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "itertools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp006, πE = πg.GetAttrImport(πF, πTemp001, ßcount); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_count.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 15: from time import time as _time, sleep as _sleep
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp006, πE = πg.GetAttrImport(πF, πTemp001, ßtime); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_time.ToObject(), πTemp006); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp006, πE = πg.GetAttrImport(πF, πTemp001, ßsleep); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_sleep.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 16: from traceback import format_exc as _format_exc
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "traceback"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp006, πE = πg.GetAttrImport(πF, πTemp001, ßformat_exc); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_format_exc.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 30: __all__ = ['activeCount', 'active_count', 'Condition', 'currentThread',
			πF.SetLineno(30)
			πTemp002 = make([]*πg.Object, 17)
			πTemp002[0] = ßactiveCount.ToObject()
			πTemp002[1] = ßactive_count.ToObject()
			πTemp002[2] = ßCondition.ToObject()
			πTemp002[3] = ßcurrentThread.ToObject()
			πTemp002[4] = ßcurrent_thread.ToObject()
			πTemp002[5] = ßenumerate.ToObject()
			πTemp002[6] = ßEvent.ToObject()
			πTemp002[7] = ßLock.ToObject()
			πTemp002[8] = ßRLock.ToObject()
			πTemp002[9] = ßSemaphore.ToObject()
			πTemp002[10] = ßBoundedSemaphore.ToObject()
			πTemp002[11] = ßThread.ToObject()
			πTemp002[12] = ßTimer.ToObject()
			πTemp002[13] = ßsetprofile.ToObject()
			πTemp002[14] = ßsettrace.ToObject()
			πTemp002[15] = ßlocal.ToObject()
			πTemp002[16] = ßstack_size.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: _start_new_thread = thread.start_new_thread
			πF.SetLineno(35)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßstart_new_thread, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_start_new_thread.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 36: _allocate_lock = thread.allocate_lock
			πF.SetLineno(36)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßallocate_lock, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_allocate_lock.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 37: _get_ident = thread.get_ident
			πF.SetLineno(37)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßget_ident, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_get_ident.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 38: ThreadError = thread.error
			πF.SetLineno(38)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßerror, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßThreadError.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 39: del thread
			πF.SetLineno(39)
			if πE = πg.DelVar(πF, πF.Globals(), ßthread); πE != nil {
				continue
			}
			// line 44: warnings.filterwarnings('ignore', category=DeprecationWarning,
			πF.SetLineno(44)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ßignore.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
				continue
			}
			πTemp008 = πg.KWArgs{
				{"category", πTemp001},
				{"module", ßthreading.ToObject()},
				{"message", πg.NewStr("sys.exc_clear").ToObject()},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßfilterwarnings, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp006.Call(πF, πTemp002, πTemp008); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 53: _VERBOSE = False
			πF.SetLineno(53)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_VERBOSE.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label4
			}
			goto Label5
			// line 55: if __debug__:
			πF.SetLineno(55)
		Label4:
			// line 57: class _Verbose(object):
			πF.SetLineno(57)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp009 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Verbose", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 59: def __init__(self, verbose=None):
					πF.SetLineno(59)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "verbose", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µverbose *πg.Object = πArgs[1]; _ = µverbose
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µverbose == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 60: if verbose is None:
							πF.SetLineno(60)
						Label1:
							// line 61: verbose = _VERBOSE
							πF.SetLineno(61)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_VERBOSE); πE != nil {
								continue
							}
							µverbose = πTemp001
							goto Label2
						Label2:
							// line 62: self.__verbose = verbose
							πF.SetLineno(62)
							if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µverbose); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__verbose, πTemp001); πE != nil {
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
					// line 64: def _note(self, format, *args):
					πF.SetLineno(64)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "format", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_note", "/usr/lib/python2.7/threading.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformat *πg.Object = πArgs[1]; _ = µformat
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var µident *πg.Object = πg.UnboundLocal; _ = µident
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__verbose, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 65: if self.__verbose:
							πF.SetLineno(65)
						Label1:
							// line 66: format = format % args
							πF.SetLineno(66)
							if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, µformat, µargs); πE != nil {
								continue
							}
							µformat = πTemp001
							// line 69: ident = _get_ident()
							πF.SetLineno(69)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µident = πTemp003
							// line 70: try:
							πF.SetLineno(70)
							πF.PushCheckpoint(4)
							// line 71: name = _active[ident].name
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µident, "ident"); πE != nil {
								continue
							}
							πTemp001 = µident
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßname, nil); πE != nil {
								continue
							}
							µname = πTemp001
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 72: except KeyError:
							πF.SetLineno(72)
						Label5:
							// line 73: name = "<OS thread %d>" % ident
							πF.SetLineno(73)
							if πE = πg.CheckLocal(πF, µident, "ident"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<OS thread %d>").ToObject(), µident); πE != nil {
								continue
							}
							µname = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 74: format = "%s: %s\n" % (name, format)
							πF.SetLineno(74)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µname, µformat).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s: %s\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							µformat = πTemp001
							// line 75: _sys.stderr.write(format)
							πF.SetLineno(75)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
								continue
							}
							πTemp007[0] = µformat
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ß_note.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("_Verbose").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Verbose.ToObject(), πTemp007); πE != nil {
				continue
			}
			goto Label6
		Label5:
			// line 79: class _Verbose(object):
			πF.SetLineno(79)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp009 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Verbose", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 80: def __init__(self, verbose=None):
					πF.SetLineno(80)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "verbose", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µverbose *πg.Object = πArgs[1]; _ = µverbose
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 81: pass
							πF.SetLineno(81)
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
					// line 82: def _note(self, *args):
					πF.SetLineno(82)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_note", "/usr/lib/python2.7/threading.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 83: pass
							πF.SetLineno(83)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_note.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("_Verbose").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Verbose.ToObject(), πTemp007); πE != nil {
				continue
			}
			goto Label6
		Label6:
			// line 87: _profile_hook = None
			πF.SetLineno(87)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_profile_hook.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 88: _trace_hook = None
			πF.SetLineno(88)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_trace_hook.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 90: def setprofile(func):
			πF.SetLineno(90)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "func", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("setprofile", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 91: """Set a profile function for all threads started from the threading module.
					πF.SetLineno(91)
					// line 97: global _profile_hook
					πF.SetLineno(97)
					// line 98: _profile_hook = func
					πF.SetLineno(98)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_profile_hook.ToObject(), µfunc); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsetprofile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 91: """Set a profile function for all threads started from the threading module.
			πF.SetLineno(91)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Set a profile function for all threads started from the threading module.\n\n    The func will be passed to sys.setprofile() for each thread, before its\n    run() method is called.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßsetprofile); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 100: def settrace(func):
			πF.SetLineno(100)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "func", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("settrace", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 101: """Set a trace function for all threads started from the threading module.
					πF.SetLineno(101)
					// line 107: global _trace_hook
					πF.SetLineno(107)
					// line 108: _trace_hook = func
					πF.SetLineno(108)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_trace_hook.ToObject(), µfunc); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßsettrace.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 101: """Set a trace function for all threads started from the threading module.
			πF.SetLineno(101)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Set a trace function for all threads started from the threading module.\n\n    The func will be passed to sys.settrace() for each thread, before its run()\n    method is called.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßsettrace); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp007); πE != nil {
				continue
			}
			// line 112: Lock = _allocate_lock
			πF.SetLineno(112)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_allocate_lock); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLock.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 114: def RLock(*args, **kwargs):
			πF.SetLineno(114)
			πTemp010 = make([]πg.Param, 0)
			πTemp007 = πg.NewFunction(πg.NewCode("RLock", "/usr/lib/python2.7/threading.py", πTemp010, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
					// line 115: """Factory function that returns a new reentrant lock.
					πF.SetLineno(115)
					// line 123: return _RLock(*args, **kwargs)
					πF.SetLineno(123)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_RLock); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwargs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßRLock.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 115: """Factory function that returns a new reentrant lock.
			πF.SetLineno(115)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Factory function that returns a new reentrant lock.\n\n    A reentrant lock must be released by the thread that acquired it. Once a\n    thread has acquired a reentrant lock, the same thread may acquire it again\n    without blocking; the thread must release it once for each time it has\n    acquired it.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßRLock); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 125: class _RLock(_Verbose):
			πF.SetLineno(125)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			πTemp009 = πg.NewDict()
			if πTemp011, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp011); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_RLock", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 126: """A reentrant lock must be released by the thread that acquired it. Once a
					πF.SetLineno(126)
					// line 126: """A reentrant lock must be released by the thread that acquired it. Once a
					πF.SetLineno(126)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A reentrant lock must be released by the thread that acquired it. Once a\n       thread has acquired a reentrant lock, the same thread may acquire it\n       again without blocking; the thread must release it once for each time it\n       has acquired it.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 132: def __init__(self, verbose=None):
					πF.SetLineno(132)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "verbose", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µverbose *πg.Object = πArgs[1]; _ = µverbose
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
							// line 133: _Verbose.__init__(self, verbose)
							πF.SetLineno(133)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
								continue
							}
							πTemp001[1] = µverbose
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 134: self.__block = _allocate_lock()
							πF.SetLineno(134)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_allocate_lock); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__block, πTemp002); πE != nil {
								continue
							}
							// line 135: self.__owner = None
							πF.SetLineno(135)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__owner, πTemp003); πE != nil {
								continue
							}
							// line 136: self.__count = 0
							πF.SetLineno(136)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__count, πTemp002); πE != nil {
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
					// line 138: def __repr__(self):
					πF.SetLineno(138)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µowner *πg.Object = πg.UnboundLocal; _ = µowner
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 139: owner = self.__owner
							πF.SetLineno(139)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__owner, nil); πE != nil {
								continue
							}
							µowner = πTemp001
							// line 140: try:
							πF.SetLineno(140)
							πF.PushCheckpoint(2)
							// line 141: owner = _active[owner].name
							πF.SetLineno(141)
							if πE = πg.CheckLocal(πF, µowner, "owner"); πE != nil {
								continue
							}
							πTemp001 = µowner
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßname, nil); πE != nil {
								continue
							}
							µowner = πTemp001
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
							// line 142: except KeyError:
							πF.SetLineno(142)
						Label3:
							// line 143: pass
							πF.SetLineno(143)
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 144: return "<%s owner=%r count=%d>" % (
							πF.SetLineno(144)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µowner, "owner"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__count, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp007, µowner, πTemp003).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s owner=%r count=%d>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 147: def acquire(self, blocking=1):
					πF.SetLineno(147)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "blocking", Def: πg.NewInt(1).ToObject()}
					πTemp004 = πg.NewFunction(πg.NewCode("acquire", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µblocking *πg.Object = πArgs[1]; _ = µblocking
						var µme *πg.Object = πg.UnboundLocal; _ = µme
						var µrc *πg.Object = πg.UnboundLocal; _ = µrc
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
							// line 148: """Acquire a lock, blocking or non-blocking.
							πF.SetLineno(148)
							// line 168: me = _get_ident()
							πF.SetLineno(168)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µme = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__owner, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µme, "me"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, µme); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 169: if self.__owner == me:
							πF.SetLineno(169)
						Label1:
							// line 170: self.__count = self.__count + 1
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__count, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__count, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 171: if __debug__:
							πF.SetLineno(171)
						Label3:
							// line 172: self._note("%s.acquire(%s): recursive success", self, blocking)
							πF.SetLineno(172)
							πTemp004 = πF.MakeArgs(3)
							πTemp004[0] = πg.NewStr("%s.acquire(%s): recursive success").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µblocking, "blocking"); πE != nil {
								continue
							}
							πTemp004[2] = µblocking
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label4
						Label4:
							// line 173: return 1
							πF.SetLineno(173)
							πR = πg.NewInt(1).ToObject()
							continue
							goto Label2
						Label2:
							// line 174: rc = self.__block.acquire(blocking)
							πF.SetLineno(174)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblocking, "blocking"); πE != nil {
								continue
							}
							πTemp004[0] = µblocking
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µrc = πTemp001
							if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µrc); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 175: if rc:
							πF.SetLineno(175)
						Label5:
							// line 176: self.__owner = me
							πF.SetLineno(176)
							if πE = πg.CheckLocal(πF, µme, "me"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µme); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__owner, πTemp001); πE != nil {
								continue
							}
							// line 177: self.__count = 1
							πF.SetLineno(177)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__count, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 178: if __debug__:
							πF.SetLineno(178)
						Label8:
							// line 179: self._note("%s.acquire(%s): initial success", self, blocking)
							πF.SetLineno(179)
							πTemp004 = πF.MakeArgs(3)
							πTemp004[0] = πg.NewStr("%s.acquire(%s): initial success").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µblocking, "blocking"); πE != nil {
								continue
							}
							πTemp004[2] = µblocking
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label9
						Label9:
							goto Label7
							// line 181: if __debug__:
							πF.SetLineno(181)
						Label6:
							// line 182: self._note("%s.acquire(%s): failure", self, blocking)
							πF.SetLineno(182)
							πTemp004 = πF.MakeArgs(3)
							πTemp004[0] = πg.NewStr("%s.acquire(%s): failure").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µblocking, "blocking"); πE != nil {
								continue
							}
							πTemp004[2] = µblocking
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label7
						Label7:
							// line 183: return rc
							πF.SetLineno(183)
							if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
								continue
							}
							πR = µrc
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßacquire.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 148: """Acquire a lock, blocking or non-blocking.
					πF.SetLineno(148)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Acquire a lock, blocking or non-blocking.\n\n        When invoked without arguments: if this thread already owns the lock,\n        increment the recursion level by one, and return immediately. Otherwise,\n        if another thread owns the lock, block until the lock is unlocked. Once\n        the lock is unlocked (not owned by any thread), then grab ownership, set\n        the recursion level to one, and return. If more than one thread is\n        blocked waiting until the lock is unlocked, only one at a time will be\n        able to grab ownership of the lock. There is no return value in this\n        case.\n\n        When invoked with the blocking argument set to true, do the same thing\n        as when called without arguments, and return true.\n\n        When invoked with the blocking argument set to false, do not block. If a\n        call without an argument would block, return false immediately;\n        otherwise, do the same thing as when called without arguments, and\n        return true.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßacquire); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 185: __enter__ = acquire
					πF.SetLineno(185)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßacquire); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 187: def release(self):
					πF.SetLineno(187)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("release", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 188: """Release a lock, decrementing the recursion level.
							πF.SetLineno(188)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__owner, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 203: if self.__owner != _get_ident():
							πF.SetLineno(203)
						Label1:
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("cannot release un-acquired lock").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 204: raise RuntimeError("cannot release un-acquired lock")
							πF.SetLineno(204)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 205: self.__count = count = self.__count - 1
							πF.SetLineno(205)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__count, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__count, πTemp002); πE != nil {
								continue
							}
							µcount = πTemp001
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µcount); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 206: if not count:
							πF.SetLineno(206)
						Label3:
							// line 207: self.__owner = None
							πF.SetLineno(207)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__owner, πTemp002); πE != nil {
								continue
							}
							// line 208: self.__block.release()
							πF.SetLineno(208)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 209: if __debug__:
							πF.SetLineno(209)
						Label6:
							// line 210: self._note("%s.release(): final release", self)
							πF.SetLineno(210)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewStr("%s.release(): final release").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label7
						Label7:
							goto Label5
							// line 212: if __debug__:
							πF.SetLineno(212)
						Label4:
							// line 213: self._note("%s.release(): non-final release", self)
							πF.SetLineno(213)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewStr("%s.release(): non-final release").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label5
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrelease.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 188: """Release a lock, decrementing the recursion level.
					πF.SetLineno(188)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Release a lock, decrementing the recursion level.\n\n        If after the decrement it is zero, reset the lock to unlocked (not owned\n        by any thread), and if any other threads are blocked waiting for the\n        lock to become unlocked, allow exactly one of them to proceed. If after\n        the decrement the recursion level is still nonzero, the lock remains\n        locked and owned by the calling thread.\n\n        Only call this method when the calling thread owns the lock. A\n        RuntimeError is raised if this method is called when the lock is\n        unlocked.\n\n        There is no return value.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßrelease); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 215: def __exit__(self, t, v, tb):
					πF.SetLineno(215)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "t", Def: nil}
					πTemp002[2] = πg.Param{Name: "v", Def: nil}
					πTemp002[3] = πg.Param{Name: "tb", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__exit__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µt *πg.Object = πArgs[1]; _ = µt
						var µv *πg.Object = πArgs[2]; _ = µv
						var µtb *πg.Object = πArgs[3]; _ = µtb
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
							// line 216: self.release()
							πF.SetLineno(216)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrelease, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 220: def _acquire_restore(self, count_owner):
					πF.SetLineno(220)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "count_owner", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_acquire_restore", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcount_owner *πg.Object = πArgs[1]; _ = µcount_owner
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µowner *πg.Object = πg.UnboundLocal; _ = µowner
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
							// line 221: count, owner = count_owner
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µcount_owner, "count_owner"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µcount_owner); πE != nil {
								continue
							}
							µcount = πTemp001
							µowner = πTemp002
							// line 222: self.__block.acquire()
							πF.SetLineno(222)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 223: self.__count = count
							πF.SetLineno(223)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcount); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__count, πTemp001); πE != nil {
								continue
							}
							// line 224: self.__owner = owner
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µowner, "owner"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µowner); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__owner, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 225: if __debug__:
							πF.SetLineno(225)
						Label1:
							// line 226: self._note("%s._acquire_restore()", self)
							πF.SetLineno(226)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("%s._acquire_restore()").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ß_acquire_restore.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 228: def _release_save(self):
					πF.SetLineno(228)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_release_save", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µowner *πg.Object = πg.UnboundLocal; _ = µowner
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 229: if __debug__:
							πF.SetLineno(229)
						Label1:
							// line 230: self._note("%s._release_save()", self)
							πF.SetLineno(230)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("%s._release_save()").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label2
						Label2:
							// line 231: count = self.__count
							πF.SetLineno(231)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__count, nil); πE != nil {
								continue
							}
							µcount = πTemp001
							// line 232: self.__count = 0
							πF.SetLineno(232)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__count, πTemp001); πE != nil {
								continue
							}
							// line 233: owner = self.__owner
							πF.SetLineno(233)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__owner, nil); πE != nil {
								continue
							}
							µowner = πTemp001
							// line 234: self.__owner = None
							πF.SetLineno(234)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__owner, πTemp004); πE != nil {
								continue
							}
							// line 235: self.__block.release()
							πF.SetLineno(235)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 236: return (count, owner)
							πF.SetLineno(236)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µowner, "owner"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µcount, µowner).ToObject()
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
					if πE = πClass.SetItem(πF, ß_release_save.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 238: def _is_owned(self):
					πF.SetLineno(238)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("_is_owned", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 239: return self.__owner == _get_ident()
							πF.SetLineno(239)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__owner, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_is_owned.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp012, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp012 == nil {
				πTemp012 = πg.TypeType.ToObject()
			}
			if πTemp013, πE = πTemp012.Call(πF, []*πg.Object{πg.NewStr("_RLock").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_RLock.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 242: def Condition(*args, **kwargs):
			πF.SetLineno(242)
			πTemp010 = make([]πg.Param, 0)
			πTemp011 = πg.NewFunction(πg.NewCode("Condition", "/usr/lib/python2.7/threading.py", πTemp010, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
					// line 243: """Factory function that returns a new condition variable object.
					πF.SetLineno(243)
					// line 253: return _Condition(*args, **kwargs)
					πF.SetLineno(253)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_Condition); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwargs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßCondition.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 243: """Factory function that returns a new condition variable object.
			πF.SetLineno(243)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Factory function that returns a new condition variable object.\n\n    A condition variable allows one or more threads to wait until they are\n    notified by another thread.\n\n    If the lock argument is given and not None, it must be a Lock or RLock\n    object, and it is used as the underlying lock. Otherwise, a new RLock object\n    is created and used as the underlying lock.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßCondition); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 255: class _Condition(_Verbose):
			πF.SetLineno(255)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp014, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
				continue
			}
			πTemp002[0] = πTemp014
			πTemp009 = πg.NewDict()
			if πTemp012, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp012); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Condition", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 256: """Condition variables allow one or more threads to wait until they are
					πF.SetLineno(256)
					// line 256: """Condition variables allow one or more threads to wait until they are
					πF.SetLineno(256)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Condition variables allow one or more threads to wait until they are\n       notified by another thread.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 260: def __init__(self, lock=None, verbose=None):
					πF.SetLineno(260)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "lock", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "verbose", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlock *πg.Object = πArgs[1]; _ = µlock
						var µverbose *πg.Object = πArgs[2]; _ = µverbose
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							case 10: goto Label10
							case 4: goto Label4
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 261: _Verbose.__init__(self, verbose)
							πF.SetLineno(261)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
								continue
							}
							πTemp001[1] = µverbose
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µlock == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 262: if lock is None:
							πF.SetLineno(262)
						Label1:
							// line 263: lock = RLock()
							πF.SetLineno(263)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßRLock); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlock = πTemp003
							goto Label2
						Label2:
							// line 264: self.__lock = lock
							πF.SetLineno(264)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µlock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__lock, πTemp002); πE != nil {
								continue
							}
							// line 266: self.acquire = lock.acquire
							πF.SetLineno(266)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßacquire, πTemp003); πE != nil {
								continue
							}
							// line 267: self.release = lock.release
							πF.SetLineno(267)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrelease, πTemp003); πE != nil {
								continue
							}
							// line 271: try:
							πF.SetLineno(271)
							πF.PushCheckpoint(4)
							// line 272: self._release_save = lock._release_save
							πF.SetLineno(272)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ß_release_save, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_release_save, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 273: except AttributeError:
							πF.SetLineno(273)
						Label5:
							// line 274: pass
							πF.SetLineno(274)
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 275: try:
							πF.SetLineno(275)
							πF.PushCheckpoint(7)
							// line 276: self._acquire_restore = lock._acquire_restore
							πF.SetLineno(276)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ß_acquire_restore, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_acquire_restore, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 277: except AttributeError:
							πF.SetLineno(277)
						Label8:
							// line 278: pass
							πF.SetLineno(278)
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
							// line 279: try:
							πF.SetLineno(279)
							πF.PushCheckpoint(10)
							// line 280: self._is_owned = lock._is_owned
							πF.SetLineno(280)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ß_is_owned, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_is_owned, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 281: except AttributeError:
							πF.SetLineno(281)
						Label11:
							// line 282: pass
							πF.SetLineno(282)
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
							// line 283: self.__waiters = []
							πF.SetLineno(283)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__waiters, πTemp003); πE != nil {
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
					// line 285: def __enter__(self):
					πF.SetLineno(285)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__enter__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 286: return self.__lock.__enter__()
							πF.SetLineno(286)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__lock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 288: def __exit__(self, *args):
					πF.SetLineno(288)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__exit__", "/usr/lib/python2.7/threading.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
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
							// line 289: return self.__lock.__exit__(*args)
							πF.SetLineno(289)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__lock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, πTemp002, nil, µargs, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 291: def __repr__(self):
					πF.SetLineno(291)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							default: panic("unexpected function state")
							}
							// line 292: return "<Condition(%s, %d)>" % (self.__lock, len(self.__waiters))
							πF.SetLineno(292)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__lock, nil); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__waiters, nil); πE != nil {
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
							πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<Condition(%s, %d)>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 294: def _release_save(self):
					πF.SetLineno(294)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_release_save", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 295: self.__lock.release()           # No state to save
							πF.SetLineno(295)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__lock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrelease, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_release_save.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 297: def _acquire_restore(self, x):
					πF.SetLineno(297)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_acquire_restore", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
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
							// line 298: self.__lock.acquire()           # Ignore saved state
							πF.SetLineno(298)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__lock, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_acquire_restore.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 300: def _is_owned(self):
					πF.SetLineno(300)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_is_owned", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__lock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 303: if self.__lock.acquire(0):
							πF.SetLineno(303)
						Label1:
							// line 304: self.__lock.release()
							πF.SetLineno(304)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__lock, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 305: return False
							πF.SetLineno(305)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 307: return True
							πF.SetLineno(307)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp002
							continue
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
					if πE = πClass.SetItem(πF, ß_is_owned.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 309: def wait(self, timeout=None):
					πF.SetLineno(309)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "timeout", Def: πTemp010}
					πTemp009 = πg.NewFunction(πg.NewCode("wait", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtimeout *πg.Object = πArgs[1]; _ = µtimeout
						var µwaiter *πg.Object = πg.UnboundLocal; _ = µwaiter
						var µsaved_state *πg.Object = πg.UnboundLocal; _ = µsaved_state
						var µendtime *πg.Object = πg.UnboundLocal; _ = µendtime
						var µdelay *πg.Object = πg.UnboundLocal; _ = µdelay
						var µgotit *πg.Object = πg.UnboundLocal; _ = µgotit
						var µremaining *πg.Object = πg.UnboundLocal; _ = µremaining
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 9: goto Label9
							case 10: goto Label10
							case 3: goto Label3
							case 22: goto Label22
							default: panic("unexpected function state")
							}
							// line 310: """Wait until notified or until a timeout occurs.
							πF.SetLineno(310)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_is_owned, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 332: if not self._is_owned():
							πF.SetLineno(332)
						Label1:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("cannot wait on un-acquired lock").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 333: raise RuntimeError("cannot wait on un-acquired lock")
							πF.SetLineno(333)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 334: waiter = _allocate_lock()
							πF.SetLineno(334)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_allocate_lock); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µwaiter = πTemp002
							// line 335: waiter.acquire()
							πF.SetLineno(335)
							if πE = πg.CheckLocal(πF, µwaiter, "waiter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µwaiter, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 336: self.__waiters.append(waiter)
							πF.SetLineno(336)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µwaiter, "waiter"); πE != nil {
								continue
							}
							πTemp005[0] = µwaiter
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__waiters, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 337: saved_state = self._release_save()
							πF.SetLineno(337)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_release_save, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsaved_state = πTemp002
							// line 338: try:    # restore state no matter what (e.g., KeyboardInterrupt)
							πF.SetLineno(338)
							πF.PushCheckpoint(3)
							if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µtimeout == πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 339: if timeout is None:
							πF.SetLineno(339)
						Label4:
							// line 340: waiter.acquire()
							πF.SetLineno(340)
							if πE = πg.CheckLocal(πF, µwaiter, "waiter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µwaiter, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 341: if __debug__:
							πF.SetLineno(341)
						Label7:
							// line 342: self._note("%s.wait(): got it", self)
							πF.SetLineno(342)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("%s.wait(): got it").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label8
						Label8:
							goto Label6
						Label5:
							// line 349: endtime = _time() + timeout
							πF.SetLineno(349)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, µtimeout); πE != nil {
								continue
							}
							µendtime = πTemp001
							// line 350: delay = 0.0005 # 500 us -> initial delay of 1 ms
							πF.SetLineno(350)
							µdelay = πg.NewFloat(0.0005).ToObject()
							// line 351: while True:
							πF.SetLineno(351)
							πF.PushCheckpoint(10)
							πTemp004 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label11
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(9)            
							// line 352: gotit = waiter.acquire(0)
							πF.SetLineno(352)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µwaiter, "waiter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µwaiter, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µgotit = πTemp002
							if πE = πg.CheckLocal(πF, µgotit, "gotit"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µgotit); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label12
							}
							goto Label13
							// line 353: if gotit:
							πF.SetLineno(353)
						Label12:
							// line 354: break
							πF.SetLineno(354)
							πTemp004 = true
							continue
							goto Label13
						Label13:
							// line 355: remaining = endtime - _time()
							πF.SetLineno(355)
							if πE = πg.CheckLocal(πF, µendtime, "endtime"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µendtime, πTemp003); πE != nil {
								continue
							}
							µremaining = πTemp001
							if πE = πg.CheckLocal(πF, µremaining, "remaining"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, µremaining, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label14
							}
							goto Label15
							// line 356: if remaining <= 0:
							πF.SetLineno(356)
						Label14:
							// line 357: break
							πF.SetLineno(357)
							πTemp004 = true
							continue
							goto Label15
						Label15:
							// line 358: delay = min(delay * 2, remaining, .05)
							πF.SetLineno(358)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µdelay, "delay"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µdelay, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µremaining, "remaining"); πE != nil {
								continue
							}
							πTemp005[1] = µremaining
							πTemp005[2] = πg.NewFloat(0.05).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µdelay = πTemp002
							// line 359: _sleep(delay)
							πF.SetLineno(359)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelay, "delay"); πE != nil {
								continue
							}
							πTemp005[0] = µdelay
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_sleep); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							if πE = πg.CheckLocal(πF, µgotit, "gotit"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µgotit); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label17
							}
							goto Label18
							// line 360: if not gotit:
							πF.SetLineno(360)
						Label16:
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label19
							}
							goto Label20
							// line 361: if __debug__:
							πF.SetLineno(361)
						Label19:
							// line 362: self._note("%s.wait(%s): timed out", self, timeout)
							πF.SetLineno(362)
							πTemp005 = πF.MakeArgs(3)
							πTemp005[0] = πg.NewStr("%s.wait(%s): timed out").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
								continue
							}
							πTemp005[2] = µtimeout
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label20
						Label20:
							// line 363: try:
							πF.SetLineno(363)
							πF.PushCheckpoint(22)
							// line 364: self.__waiters.remove(waiter)
							πF.SetLineno(364)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µwaiter, "waiter"); πE != nil {
								continue
							}
							πTemp005[0] = µwaiter
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__waiters, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßremove, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label21
						Label22:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label23
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 365: except ValueError:
							πF.SetLineno(365)
						Label23:
							// line 366: pass
							πF.SetLineno(366)
							πF.RestoreExc(nil, nil)
							goto Label21
						Label21:
							goto Label18
							// line 368: if __debug__:
							πF.SetLineno(368)
						Label17:
							// line 369: self._note("%s.wait(%s): got it", self, timeout)
							πF.SetLineno(369)
							πTemp005 = πF.MakeArgs(3)
							πTemp005[0] = πg.NewStr("%s.wait(%s): got it").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
								continue
							}
							πTemp005[2] = µtimeout
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label18
						Label18:
							goto Label6
						Label6:
							πF.PopCheckpoint()
							goto Label3
						Label3:
							πTemp007, πTemp008 = πF.RestoreExc(nil, nil)
							// line 371: self._acquire_restore(saved_state)
							πF.SetLineno(371)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsaved_state, "saved_state"); πE != nil {
								continue
							}
							πTemp005[0] = µsaved_state
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_acquire_restore, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp007 != nil {
								πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
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
					if πE = πClass.SetItem(πF, ßwait.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 310: """Wait until notified or until a timeout occurs.
					πF.SetLineno(310)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Wait until notified or until a timeout occurs.\n\n        If the calling thread has not acquired the lock when this method is\n        called, a RuntimeError is raised.\n\n        This method releases the underlying lock, and then blocks until it is\n        awakened by a notify() or notifyAll() call for the same condition\n        variable in another thread, or until the optional timeout occurs. Once\n        awakened or timed out, it re-acquires the lock and returns.\n\n        When the timeout argument is present and not None, it should be a\n        floating point number specifying a timeout for the operation in seconds\n        (or fractions thereof).\n\n        When the underlying lock is an RLock, it is not released using its\n        release() method, since this may not actually unlock the lock when it\n        was acquired multiple times recursively. Instead, an internal interface\n        of the RLock class is used, which really unlocks it even when it has\n        been recursively acquired several times. Another internal interface is\n        then used to restore the recursion level when the lock is reacquired.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßwait); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 373: def notify(self, n=1):
					πF.SetLineno(373)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: πg.NewInt(1).ToObject()}
					πTemp010 = πg.NewFunction(πg.NewCode("notify", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
						var µ__waiters *πg.Object = πg.UnboundLocal; _ = µ__waiters
						var µwaiters *πg.Object = πg.UnboundLocal; _ = µwaiters
						var µwaiter *πg.Object = πg.UnboundLocal; _ = µwaiter
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 9: goto Label9
							case 10: goto Label10
							case 13: goto Label13
							default: panic("unexpected function state")
							}
							// line 374: """Wake up one or more threads waiting on this condition, if any.
							πF.SetLineno(374)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_is_owned, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 383: if not self._is_owned():
							πF.SetLineno(383)
						Label1:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("cannot notify on un-acquired lock").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 384: raise RuntimeError("cannot notify on un-acquired lock")
							πF.SetLineno(384)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 385: __waiters = self.__waiters
							πF.SetLineno(385)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__waiters, nil); πE != nil {
								continue
							}
							µ__waiters = πTemp001
							// line 386: waiters = __waiters[:n]
							πF.SetLineno(386)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µn, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µ__waiters, "__waiters"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µ__waiters, πTemp001); πE != nil {
								continue
							}
							µwaiters = πTemp002
							if πE = πg.CheckLocal(πF, µwaiters, "waiters"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µwaiters); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 387: if not waiters:
							πF.SetLineno(387)
						Label3:
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 388: if __debug__:
							πF.SetLineno(388)
						Label5:
							// line 389: self._note("%s.notify(): no waiters", self)
							πF.SetLineno(389)
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("%s.notify(): no waiters").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label6
						Label6:
							// line 390: return
							πF.SetLineno(390)
							πR = πg.None
							continue
							goto Label4
						Label4:
							// line 391: self._note("%s.notify(): notifying %d waiter%s", self, n,
							πF.SetLineno(391)
							πTemp005 = πF.MakeArgs(4)
							πTemp005[0] = πg.NewStr("%s.notify(): notifying %d waiter%s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp005[2] = µn
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label8
							}
							πTemp002 = ßs.ToObject()
						Label8:
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							πTemp001 = ß.ToObject()
						Label7:
							πTemp005[3] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µwaiters, "waiters"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µwaiters); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp004 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label11
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
								µwaiter = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(9)            
							// line 394: waiter.release()
							πF.SetLineno(394)
							if πE = πg.CheckLocal(πF, µwaiter, "waiter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µwaiter, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 395: try:
							πF.SetLineno(395)
							πF.PushCheckpoint(13)
							// line 396: __waiters.remove(waiter)
							πF.SetLineno(396)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µwaiter, "waiter"); πE != nil {
								continue
							}
							πTemp005[0] = µwaiter
							if πE = πg.CheckLocal(πF, µ__waiters, "__waiters"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µ__waiters, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label14
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 397: except ValueError:
							πF.SetLineno(397)
						Label14:
							// line 398: pass
							πF.SetLineno(398)
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßnotify.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 374: """Wake up one or more threads waiting on this condition, if any.
					πF.SetLineno(374)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Wake up one or more threads waiting on this condition, if any.\n\n        If the calling thread has not acquired the lock when this method is\n        called, a RuntimeError is raised.\n\n        This method wakes up at most n of the threads waiting for the condition\n        variable; it is a no-op if no threads are waiting.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßnotify); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 400: def notifyAll(self):
					πF.SetLineno(400)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("notifyAll", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 401: """Wake up all threads waiting on this condition.
							πF.SetLineno(401)
							// line 407: self.notify(len(self.__waiters))
							πF.SetLineno(407)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__waiters, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnotify, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßnotifyAll.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 401: """Wake up all threads waiting on this condition.
					πF.SetLineno(401)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Wake up all threads waiting on this condition.\n\n        If the calling thread has not acquired the lock when this method\n        is called, a RuntimeError is raised.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßnotifyAll); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 409: notify_all = notifyAll
					πF.SetLineno(409)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßnotifyAll); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßnotify_all.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp013, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp013 == nil {
				πTemp013 = πg.TypeType.ToObject()
			}
			if πTemp014, πE = πTemp013.Call(πF, []*πg.Object{πg.NewStr("_Condition").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Condition.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 412: def Semaphore(*args, **kwargs):
			πF.SetLineno(412)
			πTemp010 = make([]πg.Param, 0)
			πTemp012 = πg.NewFunction(πg.NewCode("Semaphore", "/usr/lib/python2.7/threading.py", πTemp010, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
					// line 413: """A factory function that returns a new semaphore.
					πF.SetLineno(413)
					// line 421: return _Semaphore(*args, **kwargs)
					πF.SetLineno(421)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_Semaphore); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwargs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßSemaphore.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 413: """A factory function that returns a new semaphore.
			πF.SetLineno(413)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("A factory function that returns a new semaphore.\n\n    Semaphores manage a counter representing the number of release() calls minus\n    the number of acquire() calls, plus an initial value. The acquire() method\n    blocks if necessary until it can return without making the counter\n    negative. If not given, value defaults to 1.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßSemaphore); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 423: class _Semaphore(_Verbose):
			πF.SetLineno(423)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp015, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
				continue
			}
			πTemp002[0] = πTemp015
			πTemp009 = πg.NewDict()
			if πTemp013, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp013); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Semaphore", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 424: """Semaphores manage a counter representing the number of release() calls
					πF.SetLineno(424)
					// line 424: """Semaphores manage a counter representing the number of release() calls
					πF.SetLineno(424)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Semaphores manage a counter representing the number of release() calls\n       minus the number of acquire() calls, plus an initial value. The acquire()\n       method blocks if necessary until it can return without making the counter\n       negative. If not given, value defaults to 1.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 433: def __init__(self, value=1, verbose=None):
					πF.SetLineno(433)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: πg.NewInt(1).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "verbose", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var µverbose *πg.Object = πArgs[2]; _ = µverbose
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µvalue, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 434: if value < 0:
							πF.SetLineno(434)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("semaphore initial value must be >= 0").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 435: raise ValueError("semaphore initial value must be >= 0")
							πF.SetLineno(435)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 436: _Verbose.__init__(self, verbose)
							πF.SetLineno(436)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
								continue
							}
							πTemp003[1] = µverbose
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 437: self.__cond = Condition(Lock())
							πF.SetLineno(437)
							πTemp003 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßLock); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCondition); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__cond, πTemp001); πE != nil {
								continue
							}
							// line 438: self.__value = value
							πF.SetLineno(438)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__value, πTemp001); πE != nil {
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
					// line 440: def acquire(self, blocking=1):
					πF.SetLineno(440)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "blocking", Def: πg.NewInt(1).ToObject()}
					πTemp003 = πg.NewFunction(πg.NewCode("acquire", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µblocking *πg.Object = πArgs[1]; _ = µblocking
						var µrc *πg.Object = πg.UnboundLocal; _ = µrc
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Type
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
							// line 441: """Acquire a semaphore, decrementing the internal counter by one.
							πF.SetLineno(441)
							// line 460: rc = False
							πF.SetLineno(460)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µrc = πTemp001
							// line 461: with self.__cond:
							πF.SetLineno(461)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
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
							πF.PushCheckpoint(1)
							// line 462: while self.__value == 0:
							πF.SetLineno(462)
							πF.PushCheckpoint(3)
							πTemp004 = false
						Label2:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß__value, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Eq(πF, πTemp007, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(2)            
							if πE = πg.CheckLocal(πF, µblocking, "blocking"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µblocking); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 463: if not blocking:
							πF.SetLineno(463)
						Label5:
							// line 464: break
							πF.SetLineno(464)
							πTemp004 = true
							continue
							goto Label6
						Label6:
							if πTemp006, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 465: if __debug__:
							πF.SetLineno(465)
						Label7:
							// line 466: self._note("%s.acquire(%s): blocked waiting, value=%s",
							πF.SetLineno(466)
							πTemp008 = πF.MakeArgs(4)
							πTemp008[0] = πg.NewStr("%s.acquire(%s): blocked waiting, value=%s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp008[1] = µself
							if πE = πg.CheckLocal(πF, µblocking, "blocking"); πE != nil {
								continue
							}
							πTemp008[2] = µblocking
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß__value, nil); πE != nil {
								continue
							}
							πTemp008[3] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label8
						Label8:
							// line 468: self.__cond.wait()
							πF.SetLineno(468)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßwait, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							// line 470: self.__value = self.__value - 1
							πF.SetLineno(470)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß__value, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__value, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							goto Label10
							// line 471: if __debug__:
							πF.SetLineno(471)
						Label9:
							// line 472: self._note("%s.acquire: success, value=%s",
							πF.SetLineno(472)
							πTemp008 = πF.MakeArgs(3)
							πTemp008[0] = πg.NewStr("%s.acquire: success, value=%s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp008[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß__value, nil); πE != nil {
								continue
							}
							πTemp008[2] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							goto Label10
						Label10:
							// line 474: rc = True
							πF.SetLineno(474)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µrc = πTemp006
						Label4:
							πF.PopCheckpoint()
						Label1:
							πTemp009, πTemp010 = nil, nil
							if πE != nil {
								πTemp009, πTemp010 = πF.ExcInfo()
							}
							if πTemp009 != nil {
								πTemp011 = πTemp009.Type()
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp009 != nil && πTemp004 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 475: return rc
							πF.SetLineno(475)
							if πE = πg.CheckLocal(πF, µrc, "rc"); πE != nil {
								continue
							}
							πR = µrc
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßacquire.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 441: """Acquire a semaphore, decrementing the internal counter by one.
					πF.SetLineno(441)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Acquire a semaphore, decrementing the internal counter by one.\n\n        When invoked without arguments: if the internal counter is larger than\n        zero on entry, decrement it by one and return immediately. If it is zero\n        on entry, block, waiting until some other thread has called release() to\n        make it larger than zero. This is done with proper interlocking so that\n        if multiple acquire() calls are blocked, release() will wake exactly one\n        of them up. The implementation may pick one at random, so the order in\n        which blocked threads are awakened should not be relied on. There is no\n        return value in this case.\n\n        When invoked with blocking set to true, do the same thing as when called\n        without arguments, and return true.\n\n        When invoked with blocking set to false, do not block. If a call without\n        an argument would block, return false immediately; otherwise, do the\n        same thing as when called without arguments, and return true.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßacquire); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 477: __enter__ = acquire
					πF.SetLineno(477)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßacquire); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 479: def release(self):
					πF.SetLineno(479)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("release", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 480: """Release a semaphore, incrementing the internal counter by one.
							πF.SetLineno(480)
							// line 486: with self.__cond:
							πF.SetLineno(486)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
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
							πF.PushCheckpoint(1)
							// line 487: self.__value = self.__value + 1
							πF.SetLineno(487)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__value, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__value, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							goto Label3
							// line 488: if __debug__:
							πF.SetLineno(488)
						Label2:
							// line 489: self._note("%s.release: success, value=%s",
							πF.SetLineno(489)
							πTemp007 = πF.MakeArgs(3)
							πTemp007[0] = πg.NewStr("%s.release: success, value=%s").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp007[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__value, nil); πE != nil {
								continue
							}
							πTemp007[2] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label3
						Label3:
							// line 491: self.__cond.notify()
							πF.SetLineno(491)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßnotify, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp008, πTemp009 = nil, nil
							if πE != nil {
								πTemp008, πTemp009 = πF.ExcInfo()
							}
							if πTemp008 != nil {
								πTemp010 = πTemp008.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 != nil && πTemp006 != true {
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
					if πE = πClass.SetItem(πF, ßrelease.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 480: """Release a semaphore, incrementing the internal counter by one.
					πF.SetLineno(480)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Release a semaphore, incrementing the internal counter by one.\n\n        When the counter is zero on entry and another thread is waiting for it\n        to become larger than zero again, wake up that thread.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßrelease); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 493: def __exit__(self, t, v, tb):
					πF.SetLineno(493)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "t", Def: nil}
					πTemp002[2] = πg.Param{Name: "v", Def: nil}
					πTemp002[3] = πg.Param{Name: "tb", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__exit__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µt *πg.Object = πArgs[1]; _ = µt
						var µv *πg.Object = πArgs[2]; _ = µv
						var µtb *πg.Object = πArgs[3]; _ = µtb
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
							// line 494: self.release()
							πF.SetLineno(494)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrelease, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp014, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp014 == nil {
				πTemp014 = πg.TypeType.ToObject()
			}
			if πTemp015, πE = πTemp014.Call(πF, []*πg.Object{πg.NewStr("_Semaphore").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Semaphore.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 497: def BoundedSemaphore(*args, **kwargs):
			πF.SetLineno(497)
			πTemp010 = make([]πg.Param, 0)
			πTemp013 = πg.NewFunction(πg.NewCode("BoundedSemaphore", "/usr/lib/python2.7/threading.py", πTemp010, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
					// line 498: """A factory function that returns a new bounded semaphore.
					πF.SetLineno(498)
					// line 513: return _BoundedSemaphore(*args, **kwargs)
					πF.SetLineno(513)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_BoundedSemaphore); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwargs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßBoundedSemaphore.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 498: """A factory function that returns a new bounded semaphore.
			πF.SetLineno(498)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πg.NewStr("A factory function that returns a new bounded semaphore.\n\n    A bounded semaphore checks to make sure its current value doesn't exceed its\n    initial value. If it does, ValueError is raised. In most situations\n    semaphores are used to guard resources with limited capacity.\n\n    If the semaphore is released too many times it's a sign of a bug. If not\n    given, value defaults to 1.\n\n    Like regular semaphores, bounded semaphores manage a counter representing\n    the number of release() calls minus the number of acquire() calls, plus an\n    initial value. The acquire() method blocks if necessary until it can return\n    without making the counter negative. If not given, value defaults to 1.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßBoundedSemaphore); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp015, ß__doc__, πTemp014); πE != nil {
				continue
			}
			// line 515: class _BoundedSemaphore(_Semaphore):
			πF.SetLineno(515)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ß_Semaphore); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			πTemp009 = πg.NewDict()
			if πTemp014, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp014); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_BoundedSemaphore", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 516: """A bounded semaphore checks to make sure its current value doesn't exceed
					πF.SetLineno(516)
					// line 516: """A bounded semaphore checks to make sure its current value doesn't exceed
					πF.SetLineno(516)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A bounded semaphore checks to make sure its current value doesn't exceed\n       its initial value. If it does, ValueError is raised. In most situations\n       semaphores are used to guard resources with limited capacity.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 521: def __init__(self, value=1, verbose=None):
					πF.SetLineno(521)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: πg.NewInt(1).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "verbose", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var µverbose *πg.Object = πArgs[2]; _ = µverbose
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
							// line 522: _Semaphore.__init__(self, value, verbose)
							πF.SetLineno(522)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[1] = µvalue
							if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
								continue
							}
							πTemp001[2] = µverbose
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Semaphore); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 523: self._initial_value = value
							πF.SetLineno(523)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_initial_value, πTemp002); πE != nil {
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
					// line 525: def release(self):
					πF.SetLineno(525)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("release", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Type
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 526: """Release a semaphore, incrementing the internal counter by one.
							πF.SetLineno(526)
							// line 535: with self.__cond:
							πF.SetLineno(535)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
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
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__value, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_initial_value, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GE(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label2
							}
							goto Label3
							// line 536: if self.__value >= self._initial_value:
							πF.SetLineno(536)
						Label2:
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("Semaphore released too many times").ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 537: raise ValueError("Semaphore released too many times")
							πF.SetLineno(537)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label3
						Label3:
							// line 538: self.__value += 1
							πF.SetLineno(538)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__value, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__value, πTemp005); πE != nil {
								continue
							}
							// line 539: self.__cond.notify()
							πF.SetLineno(539)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßnotify, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp009, πTemp010 = nil, nil
							if πE != nil {
								πTemp009, πTemp010 = πF.ExcInfo()
							}
							if πTemp009 != nil {
								πTemp011 = πTemp009.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp009 != nil && πTemp007 != true {
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
					if πE = πClass.SetItem(πF, ßrelease.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 526: """Release a semaphore, incrementing the internal counter by one.
					πF.SetLineno(526)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Release a semaphore, incrementing the internal counter by one.\n\n        When the counter is zero on entry and another thread is waiting for it\n        to become larger than zero again, wake up that thread.\n\n        If the number of releases exceeds the number of acquires,\n        raise a ValueError.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßrelease); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp015, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp015 == nil {
				πTemp015 = πg.TypeType.ToObject()
			}
			if πTemp016, πE = πTemp015.Call(πF, []*πg.Object{πg.NewStr("_BoundedSemaphore").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_BoundedSemaphore.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 542: def Event(*args, **kwargs):
			πF.SetLineno(542)
			πTemp010 = make([]πg.Param, 0)
			πTemp014 = πg.NewFunction(πg.NewCode("Event", "/usr/lib/python2.7/threading.py", πTemp010, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
					// line 543: """A factory function that returns a new event.
					πF.SetLineno(543)
					// line 550: return _Event(*args, **kwargs)
					πF.SetLineno(550)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_Event); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwargs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßEvent.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 543: """A factory function that returns a new event.
			πF.SetLineno(543)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("A factory function that returns a new event.\n\n    Events manage a flag that can be set to true with the set() method and reset\n    to false with the clear() method. The wait() method blocks until the flag is\n    true.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßEvent); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
				continue
			}
			// line 552: class _Event(_Verbose):
			πF.SetLineno(552)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp017, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
				continue
			}
			πTemp002[0] = πTemp017
			πTemp009 = πg.NewDict()
			if πTemp015, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp015); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Event", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 553: """A factory function that returns a new event object. An event manages a
					πF.SetLineno(553)
					// line 553: """A factory function that returns a new event object. An event manages a
					πF.SetLineno(553)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A factory function that returns a new event object. An event manages a\n       flag that can be set to true with the set() method and reset to false\n       with the clear() method. The wait() method blocks until the flag is true.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 561: def __init__(self, verbose=None):
					πF.SetLineno(561)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "verbose", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µverbose *πg.Object = πArgs[1]; _ = µverbose
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
							// line 562: _Verbose.__init__(self, verbose)
							πF.SetLineno(562)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
								continue
							}
							πTemp001[1] = µverbose
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 563: self.__cond = Condition(Lock())
							πF.SetLineno(563)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßLock); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßCondition); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__cond, πTemp002); πE != nil {
								continue
							}
							// line 564: self.__flag = False
							πF.SetLineno(564)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__flag, πTemp003); πE != nil {
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
					// line 566: def _reset_internal_locks(self):
					πF.SetLineno(566)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_reset_internal_locks", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 568: self.__cond.__init__(Lock())
							πF.SetLineno(568)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßLock); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_reset_internal_locks.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 570: def isSet(self):
					πF.SetLineno(570)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("isSet", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 571: 'Return true if and only if the internal flag is true.'
							πF.SetLineno(571)
							// line 572: return self.__flag
							πF.SetLineno(572)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__flag, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßisSet.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 571: 'Return true if and only if the internal flag is true.'
					πF.SetLineno(571)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return true if and only if the internal flag is true.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßisSet); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 574: is_set = isSet
					πF.SetLineno(574)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßisSet); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßis_set.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 576: def set(self):
					πF.SetLineno(576)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("set", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Type
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 577: """Set the internal flag to true.
							πF.SetLineno(577)
							// line 583: with self.__cond:
							πF.SetLineno(583)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
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
							πF.PushCheckpoint(1)
							// line 584: self.__flag = True
							πF.SetLineno(584)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__flag, πTemp005); πE != nil {
								continue
							}
							// line 585: self.__cond.notify_all()
							πF.SetLineno(585)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßnotify_all, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp007, πTemp008 = nil, nil
							if πE != nil {
								πTemp007, πTemp008 = πF.ExcInfo()
							}
							if πTemp007 != nil {
								πTemp009 = πTemp007.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp009.ToObject(), πTemp007.ToObject(), πTemp008.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 != nil && πTemp006 != true {
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
					if πE = πClass.SetItem(πF, ßset.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 577: """Set the internal flag to true.
					πF.SetLineno(577)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Set the internal flag to true.\n\n        All threads waiting for the flag to become true are awakened. Threads\n        that call wait() once the flag is true will not block at all.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßset); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 587: def clear(self):
					πF.SetLineno(587)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("clear", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 *πg.Type
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 588: """Reset the internal flag to false.
							πF.SetLineno(588)
							// line 594: with self.__cond:
							πF.SetLineno(594)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
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
							πF.PushCheckpoint(1)
							// line 595: self.__flag = False
							πF.SetLineno(595)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__flag, πTemp005); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp007, πTemp008 = nil, nil
							if πE != nil {
								πTemp007, πTemp008 = πF.ExcInfo()
							}
							if πTemp007 != nil {
								πTemp009 = πTemp007.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp009.ToObject(), πTemp007.ToObject(), πTemp008.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 != nil && πTemp006 != true {
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
					if πE = πClass.SetItem(πF, ßclear.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 588: """Reset the internal flag to false.
					πF.SetLineno(588)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Reset the internal flag to false.\n\n        Subsequently, threads calling wait() will block until set() is called to\n        set the internal flag to true again.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßclear); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 597: def wait(self, timeout=None):
					πF.SetLineno(597)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "timeout", Def: πTemp008}
					πTemp007 = πg.NewFunction(πg.NewCode("wait", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtimeout *πg.Object = πArgs[1]; _ = µtimeout
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							// line 598: """Block until the internal flag is true.
							πF.SetLineno(598)
							// line 612: with self.__cond:
							πF.SetLineno(612)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
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
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__flag, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							goto Label3
							// line 613: if not self.__flag:
							πF.SetLineno(613)
						Label2:
							// line 614: self.__cond.wait(timeout)
							πF.SetLineno(614)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
								continue
							}
							πTemp007[0] = µtimeout
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__cond, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßwait, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label3
						Label3:
							// line 615: return self.__flag
							πF.SetLineno(615)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__flag, nil); πE != nil {
								continue
							}
							πR = πTemp004
							continue
							πF.PopCheckpoint()
						Label1:
							πTemp008, πTemp009 = nil, nil
							if πE != nil {
								πTemp008, πTemp009 = πF.ExcInfo()
							}
							if πTemp008 != nil {
								πTemp010 = πTemp008.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 != nil && πTemp006 != true {
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
					if πE = πClass.SetItem(πF, ßwait.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 598: """Block until the internal flag is true.
					πF.SetLineno(598)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Block until the internal flag is true.\n\n        If the internal flag is true on entry, return immediately. Otherwise,\n        block until another thread calls set() to set the flag to true, or until\n        the optional timeout occurs.\n\n        When the timeout argument is present and not None, it should be a\n        floating point number specifying a timeout for the operation in seconds\n        (or fractions thereof).\n\n        This method returns the internal flag on exit, so it will always return\n        True except if a timeout is given and the operation times out.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßwait); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp016, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp016 == nil {
				πTemp016 = πg.TypeType.ToObject()
			}
			if πTemp017, πE = πTemp016.Call(πF, []*πg.Object{πg.NewStr("_Event").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Event.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 618: _counter = _count().next
			πF.SetLineno(618)
			if πTemp015, πE = πg.ResolveGlobal(πF, ß_count); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp015, πE = πg.GetAttr(πF, πTemp016, ßnext, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_counter.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 619: _counter() # Consume 0 so first non-main thread has id 1.
			πF.SetLineno(619)
			if πTemp015, πE = πg.ResolveGlobal(πF, ß_counter); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, nil, nil); πE != nil {
				continue
			}
			// line 620: def _newname(template="Thread-%d"):
			πF.SetLineno(620)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "template", Def: πg.NewStr("Thread-%d").ToObject()}
			πTemp015 = πg.NewFunction(πg.NewCode("_newname", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtemplate *πg.Object = πArgs[0]; _ = µtemplate
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
					// line 621: return template % _counter()
					πF.SetLineno(621)
					if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_counter); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mod(πF, µtemplate, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_newname.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 624: _active_limbo_lock = _allocate_lock()
			πF.SetLineno(624)
			if πTemp016, πE = πg.ResolveGlobal(πF, ß_allocate_lock); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_active_limbo_lock.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 625: _active = {}    # maps thread id to Thread object
			πF.SetLineno(625)
			πTemp009 = πg.NewDict()
			πTemp016 = πTemp009.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_active.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 626: _limbo = {}
			πF.SetLineno(626)
			πTemp009 = πg.NewDict()
			πTemp016 = πTemp009.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_limbo.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 631: class Thread(_Verbose):
			πF.SetLineno(631)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp018, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
				continue
			}
			πTemp002[0] = πTemp018
			πTemp009 = πg.NewDict()
			if πTemp016, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp016); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Thread", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
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
				var πTemp008 *πg.Object
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
				var πTemp025 *πg.Object
				_ = πTemp025
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 632: """A class that represents a thread of control.
					πF.SetLineno(632)
					// line 632: """A class that represents a thread of control.
					πF.SetLineno(632)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A class that represents a thread of control.\n\n    This class can be safely subclassed in a limited fashion.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 637: __initialized = False
					πF.SetLineno(637)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__initialized.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 639: def __init__(self, group=None, target=None, name=None,
					πF.SetLineno(639)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "group", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "target", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "name", Def: πTemp003}
					πTemp003 = πg.NewTuple0().ToObject()
					πTemp002[4] = πg.Param{Name: "args", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "kwargs", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "verbose", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgroup *πg.Object = πArgs[1]; _ = µgroup
						var µtarget *πg.Object = πArgs[2]; _ = µtarget
						var µname *πg.Object = πArgs[3]; _ = µname
						var µargs *πg.Object = πArgs[4]; _ = µargs
						var µkwargs *πg.Object = πArgs[5]; _ = µkwargs
						var µverbose *πg.Object = πArgs[6]; _ = µverbose
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Dict
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
							// line 641: """This constructor should always be called with keyword arguments. Arguments are:
							πF.SetLineno(641)
							// line 662: assert group is None, "group argument must be None for now"
							πF.SetLineno(662)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µgroup == πTemp002).ToObject()
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("group argument must be None for now").ToObject()); πE != nil {
								continue
							}
							// line 663: _Verbose.__init__(self, verbose)
							πF.SetLineno(663)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µverbose, "verbose"); πE != nil {
								continue
							}
							πTemp003[1] = µverbose
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µkwargs == πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 664: if kwargs is None:
							πF.SetLineno(664)
						Label1:
							// line 665: kwargs = {}
							πF.SetLineno(665)
							πTemp005 = πg.NewDict()
							πTemp001 = πTemp005.ToObject()
							µkwargs = πTemp001
							goto Label2
						Label2:
							// line 666: self.__target = target
							πF.SetLineno(666)
							if πE = πg.CheckLocal(πF, µtarget, "target"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtarget); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__target, πTemp001); πE != nil {
								continue
							}
							// line 667: self.__name = str(name or _newname())
							πF.SetLineno(667)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = µname
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_newname); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp006
						Label3:
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__name, πTemp001); πE != nil {
								continue
							}
							// line 668: self.__args = args
							πF.SetLineno(668)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__args, πTemp001); πE != nil {
								continue
							}
							// line 669: self.__kwargs = kwargs
							πF.SetLineno(669)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µkwargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__kwargs, πTemp001); πE != nil {
								continue
							}
							// line 670: self.__daemonic = self._set_daemon()
							πF.SetLineno(670)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_set_daemon, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ß__daemonic, πTemp001); πE != nil {
								continue
							}
							// line 671: self.__ident = None
							πF.SetLineno(671)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__ident, πTemp002); πE != nil {
								continue
							}
							// line 672: self.__started = Event()
							πF.SetLineno(672)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEvent); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ß__started, πTemp001); πE != nil {
								continue
							}
							// line 673: self.__stopped = False
							πF.SetLineno(673)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__stopped, πTemp002); πE != nil {
								continue
							}
							// line 674: self.__block = Condition(Lock())
							πF.SetLineno(674)
							πTemp003 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßLock); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßCondition); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__block, πTemp001); πE != nil {
								continue
							}
							// line 675: self.__initialized = True
							πF.SetLineno(675)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__initialized, πTemp002); πE != nil {
								continue
							}
							// line 678: self.__stderr = _sys.stderr
							πF.SetLineno(678)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstderr, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__stderr, πTemp001); πE != nil {
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
					// line 641: """This constructor should always be called with keyword arguments. Arguments are:
					πF.SetLineno(641)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("This constructor should always be called with keyword arguments. Arguments are:\n\n        *group* should be None; reserved for future extension when a ThreadGroup\n        class is implemented.\n\n        *target* is the callable object to be invoked by the run()\n        method. Defaults to None, meaning nothing is called.\n\n        *name* is the thread name. By default, a unique name is constructed of\n        the form \"Thread-N\" where N is a small decimal number.\n\n        *args* is the argument tuple for the target invocation. Defaults to ().\n\n        *kwargs* is a dictionary of keyword arguments for the target\n        invocation. Defaults to {}.\n\n        If a subclass overrides the constructor, it must make sure to invoke\n        the base class constructor (Thread.__init__()) before doing anything\n        else to the thread.\n\n").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 680: def _reset_internal_locks(self):
					πF.SetLineno(680)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_reset_internal_locks", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß__block.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 683: if hasattr(self, '__block'):  # DummyThread deletes self.__block
							πF.SetLineno(683)
						Label1:
							// line 684: self.__block.__init__()
							πF.SetLineno(684)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 685: self.__started._reset_internal_locks()
							πF.SetLineno(685)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß_reset_internal_locks, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_reset_internal_locks.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 688: def _block(self):
					πF.SetLineno(688)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_block", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 690: return self.__block
							πF.SetLineno(690)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_block.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 687: @property
					πF.SetLineno(687)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß_block); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ß_block.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 692: def _set_daemon(self):
					πF.SetLineno(692)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_set_daemon", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 694: return current_thread().daemon
							πF.SetLineno(694)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßcurrent_thread); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdaemon, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_set_daemon.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 696: def __repr__(self):
					πF.SetLineno(696)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstatus *πg.Object = πg.UnboundLocal; _ = µstatus
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
							// line 697: assert self.__initialized, "Thread.__init__() was not called"
							πF.SetLineno(697)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__initialized, nil); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("Thread.__init__() was not called").ToObject()); πE != nil {
								continue
							}
							// line 698: status = "initial"
							πF.SetLineno(698)
							µstatus = ßinitial.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 699: if self.__started.is_set():
							πF.SetLineno(699)
						Label1:
							// line 700: status = "started"
							πF.SetLineno(700)
							µstatus = ßstarted.ToObject()
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__stopped, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 701: if self.__stopped:
							πF.SetLineno(701)
						Label3:
							// line 702: status = "stopped"
							πF.SetLineno(702)
							µstatus = ßstopped.ToObject()
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__daemonic, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 703: if self.__daemonic:
							πF.SetLineno(703)
						Label5:
							// line 704: status += " daemon"
							πF.SetLineno(704)
							if πE = πg.CheckLocal(πF, µstatus, "status"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µstatus, πg.NewStr(" daemon").ToObject()); πE != nil {
								continue
							}
							µstatus = πTemp001
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__ident, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp004).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 705: if self.__ident is not None:
							πF.SetLineno(705)
						Label7:
							// line 706: status += " %s" % self.__ident
							πF.SetLineno(706)
							if πE = πg.CheckLocal(πF, µstatus, "status"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__ident, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr(" %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µstatus, πTemp001); πE != nil {
								continue
							}
							µstatus = πTemp002
							goto Label8
						Label8:
							// line 707: return "<%s(%s, %s)>" % (self.__class__.__name__, self.__name, status)
							πF.SetLineno(707)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__name, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstatus, "status"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp005, πTemp004, µstatus).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s(%s, %s)>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 709: def start(self):
					πF.SetLineno(709)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("start", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Type
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 *πg.BaseException
						_ = πTemp014
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 9: goto Label9
							case 11: goto Label11
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 710: """Start the thread's activity.
							πF.SetLineno(710)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__initialized, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 719: if not self.__initialized:
							πF.SetLineno(719)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("thread.__init__() not called").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 720: raise RuntimeError("thread.__init__() not called")
							πF.SetLineno(720)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 721: if self.__started.is_set():
							πF.SetLineno(721)
						Label3:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("threads can only be started once").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 722: raise RuntimeError("threads can only be started once")
							πF.SetLineno(722)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label4
						Label4:
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 723: if __debug__:
							πF.SetLineno(723)
						Label5:
							// line 724: self._note("%s.start(): starting thread", self)
							πF.SetLineno(724)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("%s.start(): starting thread").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label6
						Label6:
							// line 725: with _active_limbo_lock:
							πF.SetLineno(725)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							// line 726: _limbo[self] = self
							πF.SetLineno(726)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µself); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ß_limbo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp008 = µself
							if πE = πg.SetItem(πF, πTemp007, πTemp008, πTemp006); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label7:
							πTemp009, πTemp010 = nil, nil
							if πE != nil {
								πTemp009, πTemp010 = πF.ExcInfo()
							}
							if πTemp009 != nil {
								πTemp011 = πTemp009.Type()
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp009 != nil && πTemp003 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 727: try:
							πF.SetLineno(727)
							πF.PushCheckpoint(9)
							// line 728: _start_new_thread(self.__bootstrap, ())
							πF.SetLineno(728)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß__bootstrap, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp006
							πTemp006 = πg.NewTuple0().ToObject()
							πTemp004[1] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_start_new_thread); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.PopCheckpoint()
							goto Label8
						Label9:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 729: except Exception:
							πF.SetLineno(729)
						Label10:
							// line 730: with _active_limbo_lock:
							πF.SetLineno(730)
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
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
							πF.PushCheckpoint(11)
							// line 731: del _limbo[self]
							πF.SetLineno(731)
							if πTemp012, πE = πg.ResolveGlobal(πF, ß_limbo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp013 = µself
							if πE = πg.DelItem(πF, πTemp012, πTemp013); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label11:
							πTemp014, πTemp010 = nil, nil
							if πE != nil {
								πTemp014, πTemp010 = πF.ExcInfo()
							}
							if πTemp014 != nil {
								πTemp011 = πTemp014.Type()
								if πTemp012, πE = πTemp007.Call(πF, πg.Args{πTemp006, πTemp011.ToObject(), πTemp014.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp012, πE = πTemp007.Call(πF, πg.Args{πTemp006, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp012); πE != nil {
								continue
							}
							if πTemp014 != nil && πTemp003 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 732: raise
							πF.SetLineno(732)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label8
						Label8:
							// line 733: self.__started.wait()
							πF.SetLineno(733)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßwait, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp013.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßstart.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 710: """Start the thread's activity.
					πF.SetLineno(710)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Start the thread's activity.\n\n        It must be called at most once per thread object. It arranges for the\n        object's run() method to be invoked in a separate thread of control.\n\n        This method will raise a RuntimeError if called more than once on the\n        same thread object.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßstart); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 735: def run(self):
					πF.SetLineno(735)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 736: """Method representing the thread's activity.
							πF.SetLineno(736)
							// line 744: try:
							πF.SetLineno(744)
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__target, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 745: if self.__target:
							πF.SetLineno(745)
						Label2:
							// line 746: self.__target(*self.__args, **self.__kwargs)
							πF.SetLineno(746)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__args, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__kwargs, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__target, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Invoke(πF, πTemp004, nil, πTemp001, nil, πTemp003); πE != nil {
								continue
							}
							goto Label3
						Label3:
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							// line 750: del self.__target, self.__args, self.__kwargs
							πF.SetLineno(750)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ß__target); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ß__args); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ß__kwargs); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 736: """Method representing the thread's activity.
					πF.SetLineno(736)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Method representing the thread's activity.\n\n        You may override this method in a subclass. The standard run() method\n        invokes the callable object passed to the object's constructor as the\n        target argument, if any, with sequential and keyword arguments taken\n        from the args and kwargs arguments, respectively.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßrun); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 752: def __bootstrap(self):
					πF.SetLineno(752)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("__bootstrap", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.BaseException
						_ = πTemp003
						var πTemp004 *πg.Traceback
						_ = πTemp004
						var πTemp005 bool
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
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 765: try:
							πF.SetLineno(765)
							πF.PushCheckpoint(2)
							// line 766: self.__bootstrap_inner()
							πF.SetLineno(766)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__bootstrap_inner, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp003, πTemp004 = πF.ExcInfo()
							goto Label3
							// line 767: except:
							πF.SetLineno(767)
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__daemonic, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label4
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006 == πTemp007).ToObject()
							πTemp001 = πTemp002
						Label4:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 768: if self.__daemonic and _sys is None:
							πF.SetLineno(768)
						Label5:
							// line 769: return
							πF.SetLineno(769)
							πR = πg.None
							continue
							goto Label6
						Label6:
							// line 770: raise
							πF.SetLineno(770)
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
					if πE = πClass.SetItem(πF, ß__bootstrap.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 772: def _set_ident(self):
					πF.SetLineno(772)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("_set_ident", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 773: self.__ident = _get_ident()
							πF.SetLineno(773)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ß__ident, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_set_ident.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 775: def __bootstrap_inner(self):
					πF.SetLineno(775)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("__bootstrap_inner", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexc_type *πg.Object = πg.UnboundLocal; _ = µexc_type
						var µexc_value *πg.Object = πg.UnboundLocal; _ = µexc_value
						var µexc_tb *πg.Object = πg.UnboundLocal; _ = µexc_tb
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
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Type
						_ = πTemp011
						var πTemp012 []*πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 bool
						_ = πTemp014
						var πTemp015 *πg.Object
						_ = πTemp015
						var πTemp016 *πg.BaseException
						_ = πTemp016
						var πTemp017 *πg.Object
						_ = πTemp017
						var πTemp018 *πg.Traceback
						_ = πTemp018
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 9: goto Label9
							case 10: goto Label10
							case 23: goto Label23
							case 24: goto Label24
							case 25: goto Label25
							case 27: goto Label27
							case 29: goto Label29
							default: panic("unexpected function state")
							}
							// line 776: try:
							πF.SetLineno(776)
							πF.PushCheckpoint(1)
							// line 777: self._set_ident()
							πF.SetLineno(777)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_set_ident, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 778: self.__started.set()
							πF.SetLineno(778)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßset, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 779: with _active_limbo_lock:
							πF.SetLineno(779)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
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
							πF.PushCheckpoint(2)
							// line 780: _active[self.__ident] = self
							πF.SetLineno(780)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µself); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß__ident, nil); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp004); πE != nil {
								continue
							}
							// line 781: del _limbo[self]
							πF.SetLineno(781)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_limbo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005 = µself
							if πE = πg.DelItem(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label2:
							πTemp009, πTemp010 = nil, nil
							if πE != nil {
								πTemp009, πTemp010 = πF.ExcInfo()
							}
							if πTemp009 != nil {
								πTemp011 = πTemp009.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp009 != nil && πTemp008 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							goto Label4
							// line 782: if __debug__:
							πF.SetLineno(782)
						Label3:
							// line 783: self._note("%s.__bootstrap(): thread started", self)
							πF.SetLineno(783)
							πTemp012 = πF.MakeArgs(2)
							πTemp012[0] = πg.NewStr("%s.__bootstrap(): thread started").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp012[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							goto Label4
						Label4:
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_trace_hook); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label5
							}
							goto Label6
							// line 785: if _trace_hook:
							πF.SetLineno(785)
						Label5:
							// line 786: self._note("%s.__bootstrap(): registering trace hook", self)
							πF.SetLineno(786)
							πTemp012 = πF.MakeArgs(2)
							πTemp012[0] = πg.NewStr("%s.__bootstrap(): registering trace hook").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp012[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							// line 787: _sys.settrace(_trace_hook)
							πF.SetLineno(787)
							πTemp012 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_trace_hook); πE != nil {
								continue
							}
							πTemp012[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsettrace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							goto Label6
						Label6:
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_profile_hook); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							goto Label8
							// line 788: if _profile_hook:
							πF.SetLineno(788)
						Label7:
							// line 789: self._note("%s.__bootstrap(): registering profile hook", self)
							πF.SetLineno(789)
							πTemp012 = πF.MakeArgs(2)
							πTemp012[0] = πg.NewStr("%s.__bootstrap(): registering profile hook").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp012[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							// line 790: _sys.setprofile(_profile_hook)
							πF.SetLineno(790)
							πTemp012 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_profile_hook); πE != nil {
								continue
							}
							πTemp012[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsetprofile, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							goto Label8
						Label8:
							// line 792: try:
							πF.SetLineno(792)
							πF.PushCheckpoint(9)
							πF.PushCheckpoint(10)
							// line 793: self.run()
							πF.SetLineno(793)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßrun, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							if πTemp004, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label11
							}
							goto Label12
							// line 831: if __debug__:
							πF.SetLineno(831)
						Label11:
							// line 832: self._note("%s.__bootstrap(): normal return", self)
							πF.SetLineno(832)
							πTemp012 = πF.MakeArgs(2)
							πTemp012[0] = πg.NewStr("%s.__bootstrap(): normal return").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp012[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							goto Label12
						Label12:
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßSystemExit); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label13
							}
							goto Label14
							// line 794: except SystemExit:
							πF.SetLineno(794)
						Label13:
							if πTemp004, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label15
							}
							goto Label16
							// line 795: if __debug__:
							πF.SetLineno(795)
						Label15:
							// line 796: self._note("%s.__bootstrap(): raised SystemExit", self)
							πF.SetLineno(796)
							πTemp012 = πF.MakeArgs(2)
							πTemp012[0] = πg.NewStr("%s.__bootstrap(): raised SystemExit").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp012[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							goto Label16
						Label16:
							πF.RestoreExc(nil, nil)
							πF.PopCheckpoint()
							goto Label9
							// line 797: except:
							πF.SetLineno(797)
						Label14:
							if πTemp004, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label17
							}
							goto Label18
							// line 798: if __debug__:
							πF.SetLineno(798)
						Label17:
							// line 799: self._note("%s.__bootstrap(): unhandled exception", self)
							πF.SetLineno(799)
							πTemp012 = πF.MakeArgs(2)
							πTemp012[0] = πg.NewStr("%s.__bootstrap(): unhandled exception").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp012[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							goto Label18
						Label18:
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label19
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßstderr, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp007 != πTemp006).ToObject()
							πTemp004 = πTemp005
						Label19:
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label20
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__stderr, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp005 != πTemp006).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label21
							}
							goto Label22
							// line 804: if _sys and _sys.stderr is not None:
							πF.SetLineno(804)
						Label20:
							// line 805: print>>_sys.stderr, ("Exception in thread %s:\n%s" %
							πF.SetLineno(805)
							πTemp012 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ß_format_exc); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp006, πTemp013).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("Exception in thread %s:\n%s").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp012[0] = πTemp004
							if πE = πg.Print(πF, πTemp012, true); πE != nil {
								continue
							}
							goto Label22
							// line 807: elif self.__stderr is not None:
							πF.SetLineno(807)
						Label21:
							// line 811: exc_type, exc_value, exc_tb = _sys.exc_info()
							πF.SetLineno(811)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp004); πE != nil {
								continue
							}
							µexc_type = πTemp005
							µexc_value = πTemp006
							µexc_tb = πTemp007
							// line 812: try:
							πF.SetLineno(812)
							πF.PushCheckpoint(23)
							// line 813: print>>self.__stderr, (
							πF.SetLineno(813)
							πTemp012 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, πg.NewStr("Exception in thread ").ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewStr(" (most likely raised during interpreter shutdown):").ToObject()); πE != nil {
								continue
							}
							πTemp012[0] = πTemp004
							if πE = πg.Print(πF, πTemp012, true); πE != nil {
								continue
							}
							// line 816: print>>self.__stderr, (
							πF.SetLineno(816)
							πTemp012 = make([]*πg.Object, 1)
							πTemp012[0] = πg.NewStr("Traceback (most recent call last):").ToObject()
							if πE = πg.Print(πF, πTemp012, true); πE != nil {
								continue
							}
							// line 818: while exc_tb:
							πF.SetLineno(818)
							πF.PushCheckpoint(25)
							πTemp008 = false
						Label24:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label26
							}
							if πE = πg.CheckLocal(πF, µexc_tb, "exc_tb"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.IsTrue(πF, µexc_tb); πE != nil {
								continue
							}
							if πE != nil || !πTemp014 {
								continue
							}
							πF.PushCheckpoint(24)            
							// line 819: print>>self.__stderr, (
							πF.SetLineno(819)
							πTemp012 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µexc_tb, "exc_tb"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µexc_tb, ßtb_frame, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßf_code, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßco_filename, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexc_tb, "exc_tb"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µexc_tb, ßtb_lineno, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexc_tb, "exc_tb"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µexc_tb, ßtb_frame, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetAttr(πF, πTemp013, ßf_code, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp015, ßco_name, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple3(πTemp006, πTemp007, πTemp013).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("  File \"%s\", line %s, in %s").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp012[0] = πTemp004
							if πE = πg.Print(πF, πTemp012, true); πE != nil {
								continue
							}
							// line 824: exc_tb = exc_tb.tb_next
							πF.SetLineno(824)
							if πE = πg.CheckLocal(πF, µexc_tb, "exc_tb"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µexc_tb, ßtb_next, nil); πE != nil {
								continue
							}
							µexc_tb = πTemp004
							continue
						Label25:
							if πE != nil || πR != nil {
								continue
							}
						Label26:
							// line 825: print>>self.__stderr, ("%s: %s" % (exc_type, exc_value))
							πF.SetLineno(825)
							πTemp012 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µexc_type, "exc_type"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexc_value, "exc_value"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µexc_type, µexc_value).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("%s: %s").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp012[0] = πTemp004
							if πE = πg.Print(πF, πTemp012, true); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label23
						Label23:
							πTemp016, πTemp010 = πF.RestoreExc(nil, nil)
							// line 829: del exc_type, exc_value, exc_tb
							πF.SetLineno(829)
							if πE = πg.CheckLocal(πF, µexc_type, "exc_type"); πE != nil {
								continue
							}
							µexc_type = πg.UnboundLocal
							if πE = πg.CheckLocal(πF, µexc_value, "exc_value"); πE != nil {
								continue
							}
							µexc_value = πg.UnboundLocal
							if πE = πg.CheckLocal(πF, µexc_tb, "exc_tb"); πE != nil {
								continue
							}
							µexc_tb = πg.UnboundLocal
							if πTemp016 != nil {
								πE = πF.Raise(πTemp016.ToObject(), nil, πTemp010.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							goto Label22
						Label22:
							πF.RestoreExc(nil, nil)
							πF.PopCheckpoint()
							goto Label9
						Label9:
							πTemp009, πTemp010 = πF.RestoreExc(nil, nil)
							// line 838: _sys.exc_clear()
							πF.SetLineno(838)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßexc_clear, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp009 != nil {
								πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp009, πTemp010 = πF.RestoreExc(nil, nil)
							// line 840: with _active_limbo_lock:
							πF.SetLineno(840)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp006.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(27)
							// line 841: self.__stop()
							πF.SetLineno(841)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß__stop, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 842: try:
							πF.SetLineno(842)
							πF.PushCheckpoint(29)
							// line 845: del _active[_get_ident()]
							πF.SetLineno(845)
							if πTemp007, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
								continue
							}
							if πTemp015, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
								continue
							}
							if πTemp017, πE = πTemp015.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp013 = πTemp017
							if πE = πg.DelItem(πF, πTemp007, πTemp013); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label28
						Label29:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp016, πTemp018 = πF.ExcInfo()
							goto Label30
							// line 846: except:
							πF.SetLineno(846)
						Label30:
							// line 847: pass
							πF.SetLineno(847)
							πF.RestoreExc(nil, nil)
							goto Label28
						Label28:
							πF.PopCheckpoint()
						Label27:
							πTemp016, πTemp018 = nil, nil
							if πE != nil {
								πTemp016, πTemp018 = πF.ExcInfo()
							}
							if πTemp016 != nil {
								πTemp011 = πTemp016.Type()
								if πTemp007, πE = πTemp005.Call(πF, πg.Args{πTemp004, πTemp011.ToObject(), πTemp016.ToObject(), πTemp018.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp007, πE = πTemp005.Call(πF, πg.Args{πTemp004, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp016 != nil && πTemp008 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ß__bootstrap_inner.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 849: def __stop(self):
					πF.SetLineno(849)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("__stop", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							πTemp002[1] = ß__block.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 852: if not hasattr(self, '__block'):
							πF.SetLineno(852)
						Label1:
							// line 853: return
							πF.SetLineno(853)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 854: self.__block.acquire()
							πF.SetLineno(854)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 855: self.__stopped = True
							πF.SetLineno(855)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__stopped, πTemp003); πE != nil {
								continue
							}
							// line 856: self.__block.notify_all()
							πF.SetLineno(856)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßnotify_all, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 857: self.__block.release()
							πF.SetLineno(857)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__stop.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 859: def __delete(self):
					πF.SetLineno(859)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("__delete", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πTemp011 *πg.Type
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							// line 860: "Remove current thread from the dict of currently running threads."
							πF.SetLineno(860)
							// line 883: try:
							πF.SetLineno(883)
							πF.PushCheckpoint(2)
							// line 884: with _active_limbo_lock:
							πF.SetLineno(884)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
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
							// line 885: del _active[_get_ident()]
							πF.SetLineno(885)
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp007
							if πE = πg.DelItem(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label3:
							πTemp009, πTemp010 = nil, nil
							if πE != nil {
								πTemp009, πTemp010 = πF.ExcInfo()
							}
							if πTemp009 != nil {
								πTemp011 = πTemp009.Type()
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp009 != nil && πTemp008 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label4
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 890: except KeyError:
							πF.SetLineno(890)
						Label4:
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_sys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Contains(πF, πTemp006, ßdummy_threading.ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label5
							}
							goto Label6
							// line 891: if 'dummy_threading' not in _sys.modules:
							πF.SetLineno(891)
						Label5:
							// line 892: raise
							πF.SetLineno(892)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label6
						Label6:
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
					if πE = πClass.SetItem(πF, ß__delete.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 860: "Remove current thread from the dict of currently running threads."
					πF.SetLineno(860)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πg.NewStr("Remove current thread from the dict of currently running threads.").ToObject()); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ß__delete); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp016, ß__doc__, πTemp015); πE != nil {
						continue
					}
					// line 894: def join(self, timeout=None):
					πF.SetLineno(894)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "timeout", Def: πTemp016}
					πTemp015 = πg.NewFunction(πg.NewCode("join", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtimeout *πg.Object = πArgs[1]; _ = µtimeout
						var µdeadline *πg.Object = πg.UnboundLocal; _ = µdeadline
						var µdelay *πg.Object = πg.UnboundLocal; _ = µdelay
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 16: goto Label16
							case 11: goto Label11
							case 20: goto Label20
							case 21: goto Label21
							case 15: goto Label15
							default: panic("unexpected function state")
							}
							// line 895: """Wait until the thread terminates.
							πF.SetLineno(895)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__initialized, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 918: if not self.__initialized:
							πF.SetLineno(918)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("Thread.__init__() not called").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 919: raise RuntimeError("Thread.__init__() not called")
							πF.SetLineno(919)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 920: if not self.__started.is_set():
							πF.SetLineno(920)
						Label3:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("cannot join thread before it is started").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 921: raise RuntimeError("cannot join thread before it is started")
							πF.SetLineno(921)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcurrent_thread); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µself == πTemp005).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 922: if self is current_thread():
							πF.SetLineno(922)
						Label5:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("cannot join current thread").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 923: raise RuntimeError("cannot join current thread")
							πF.SetLineno(923)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label6
						Label6:
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 925: if __debug__:
							πF.SetLineno(925)
						Label7:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__stopped, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 926: if not self.__stopped:
							πF.SetLineno(926)
						Label9:
							// line 927: self._note("%s.join(): waiting until thread stops", self)
							πF.SetLineno(927)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("%s.join(): waiting until thread stops").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label10
						Label10:
							goto Label8
						Label8:
							// line 928: self.__block.acquire()
							πF.SetLineno(928)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 929: try:
							πF.SetLineno(929)
							πF.PushCheckpoint(11)
							if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µtimeout == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							goto Label13
							// line 930: if timeout is None:
							πF.SetLineno(930)
						Label12:
							// line 931: while not self.__stopped:
							πF.SetLineno(931)
							πF.PushCheckpoint(16)
							πTemp003 = false
						Label15:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label17
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__stopped, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(15)            
							// line 932: self.__block.wait()
							πF.SetLineno(932)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwait, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label16:
							if πE != nil || πR != nil {
								continue
							}
						Label17:
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label18
							}
							goto Label19
							// line 933: if __debug__:
							πF.SetLineno(933)
						Label18:
							// line 934: self._note("%s.join(): thread stopped", self)
							πF.SetLineno(934)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("%s.join(): thread stopped").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label19
						Label19:
							goto Label14
						Label13:
							// line 936: deadline = _time() + timeout
							πF.SetLineno(936)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtimeout, "timeout"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp005, µtimeout); πE != nil {
								continue
							}
							µdeadline = πTemp001
							// line 937: while not self.__stopped:
							πF.SetLineno(937)
							πF.PushCheckpoint(21)
							πTemp003 = false
						Label20:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label22
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__stopped, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(20)            
							// line 938: delay = deadline - _time()
							πF.SetLineno(938)
							if πE = πg.CheckLocal(πF, µdeadline, "deadline"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_time); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µdeadline, πTemp005); πE != nil {
								continue
							}
							µdelay = πTemp001
							if πE = πg.CheckLocal(πF, µdelay, "delay"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, µdelay, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label23
							}
							goto Label24
							// line 939: if delay <= 0:
							πF.SetLineno(939)
						Label23:
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label25
							}
							goto Label26
							// line 940: if __debug__:
							πF.SetLineno(940)
						Label25:
							// line 941: self._note("%s.join(): timed out", self)
							πF.SetLineno(941)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("%s.join(): timed out").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label26
						Label26:
							// line 942: break
							πF.SetLineno(942)
							πTemp003 = true
							continue
							goto Label24
						Label24:
							// line 943: self.__block.wait(delay)
							πF.SetLineno(943)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelay, "delay"); πE != nil {
								continue
							}
							πTemp004[0] = µdelay
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwait, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label21:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label27
							}
							goto Label28
							// line 945: if __debug__:
							πF.SetLineno(945)
						Label27:
							// line 946: self._note("%s.join(): thread stopped", self)
							πF.SetLineno(946)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("%s.join(): thread stopped").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label28
						Label28:
						Label22:
							goto Label14
						Label14:
							πF.PopCheckpoint()
							goto Label11
						Label11:
							πTemp008, πTemp009 = πF.RestoreExc(nil, nil)
							// line 948: self.__block.release()
							πF.SetLineno(948)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__block, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp008 != nil {
								πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
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
					if πE = πClass.SetItem(πF, ßjoin.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 895: """Wait until the thread terminates.
					πF.SetLineno(895)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("Wait until the thread terminates.\n\n        This blocks the calling thread until the thread whose join() method is\n        called terminates -- either normally or through an unhandled exception\n        or until the optional timeout occurs.\n\n        When the timeout argument is present and not None, it should be a\n        floating point number specifying a timeout for the operation in seconds\n        (or fractions thereof). As join() always returns None, you must call\n        isAlive() after join() to decide whether a timeout happened -- if the\n        thread is still alive, the join() call timed out.\n\n        When the timeout argument is not present or None, the operation will\n        block until the thread terminates.\n\n        A thread can be join()ed many times.\n\n        join() raises a RuntimeError if an attempt is made to join the current\n        thread as that would cause a deadlock. It is also an error to join() a\n        thread before it has been started and attempts to do so raises the same\n        exception.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßjoin); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 950: def _name_getter(self):
					πF.SetLineno(950)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("_name_getter", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 951: """A string used for identification purposes only.
							πF.SetLineno(951)
							// line 957: assert self.__initialized, "Thread.__init__() not called"
							πF.SetLineno(957)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__initialized, nil); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("Thread.__init__() not called").ToObject()); πE != nil {
								continue
							}
							// line 958: return self.__name
							πF.SetLineno(958)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__name, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_name_getter.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 951: """A string used for identification purposes only.
					πF.SetLineno(951)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("A string used for identification purposes only.\n\n        It has no semantics. Multiple threads may be given the same name. The\n        initial name is set by the constructor.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ß_name_getter); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
					// line 960: def _name_setter(self, name):
					πF.SetLineno(960)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("_name_setter", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 961: assert self.__initialized, "Thread.__init__() not called"
							πF.SetLineno(961)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__initialized, nil); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("Thread.__init__() not called").ToObject()); πE != nil {
								continue
							}
							// line 962: self.__name = str(name)
							πF.SetLineno(962)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__name, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_name_setter.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 964: name = property(_name_getter, _name_setter)
					πF.SetLineno(964)
					πTemp005 = πF.MakeArgs(2)
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ß_name_getter); πE != nil {
						continue
					}
					πTemp005[0] = πTemp018
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ß_name_setter); πE != nil {
						continue
					}
					πTemp005[1] = πTemp018
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp019, πE = πTemp018.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßname.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 967: def ident(self):
					πF.SetLineno(967)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("ident", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 968: """Thread identifier of this thread or None if it has not been started.
							πF.SetLineno(968)
							// line 975: assert self.__initialized, "Thread.__init__() not called"
							πF.SetLineno(975)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__initialized, nil); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("Thread.__init__() not called").ToObject()); πE != nil {
								continue
							}
							// line 976: return self.__ident
							πF.SetLineno(976)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__ident, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßident.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 968: """Thread identifier of this thread or None if it has not been started.
					πF.SetLineno(968)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("Thread identifier of this thread or None if it has not been started.\n\n        This is a nonzero integer. See the thread.get_ident() function. Thread\n        identifiers may be recycled when a thread exits and another thread is\n        created. The identifier is available even after the thread has exited.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßident); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
						continue
					}
					// line 966: @property
					πF.SetLineno(966)
					πTemp005 = πF.MakeArgs(1)
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßident); πE != nil {
						continue
					}
					πTemp005[0] = πTemp019
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp020, πE = πTemp019.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßident.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 978: def isAlive(self):
					πF.SetLineno(978)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("isAlive", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 979: """Return whether the thread is alive.
							πF.SetLineno(979)
							// line 986: assert self.__initialized, "Thread.__init__() not called"
							πF.SetLineno(986)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__initialized, nil); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("Thread.__init__() not called").ToObject()); πE != nil {
								continue
							}
							// line 987: return self.__started.is_set() and not self.__stopped
							πF.SetLineno(987)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__stopped, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							πTemp001 = πTemp003
						Label1:
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
					if πE = πClass.SetItem(πF, ßisAlive.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 979: """Return whether the thread is alive.
					πF.SetLineno(979)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("Return whether the thread is alive.\n\n        This method returns True just before the run() method starts until just\n        after the run() method terminates. The module function enumerate()\n        returns a list of all alive threads.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßisAlive); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
						continue
					}
					// line 989: is_alive = isAlive
					πF.SetLineno(989)
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßisAlive); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßis_alive.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 991: def _daemon_getter(self):
					πF.SetLineno(991)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("_daemon_getter", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 992: """A boolean value indicating whether this thread is a daemon thread (True) or not (False).
							πF.SetLineno(992)
							// line 1003: assert self.__initialized, "Thread.__init__() not called"
							πF.SetLineno(1003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__initialized, nil); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("Thread.__init__() not called").ToObject()); πE != nil {
								continue
							}
							// line 1004: return self.__daemonic
							πF.SetLineno(1004)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__daemonic, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_daemon_getter.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 992: """A boolean value indicating whether this thread is a daemon thread (True) or not (False).
					πF.SetLineno(992)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πg.NewStr("A boolean value indicating whether this thread is a daemon thread (True) or not (False).\n\n        This must be set before start() is called, otherwise RuntimeError is\n        raised. Its initial value is inherited from the creating thread; the\n        main thread is not a daemon thread and therefore all threads created in\n        the main thread default to daemon = False.\n\n        The entire Python program exits when no alive non-daemon threads are\n        left.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ß_daemon_getter); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp021); πE != nil {
						continue
					}
					// line 1006: def _daemon_setter(self, daemonic):
					πF.SetLineno(1006)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "daemonic", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("_daemon_setter", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdaemonic *πg.Object = πArgs[1]; _ = µdaemonic
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__initialized, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1007: if not self.__initialized:
							πF.SetLineno(1007)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("Thread.__init__() not called").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1008: raise RuntimeError("Thread.__init__() not called")
							πF.SetLineno(1008)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1009: if self.__started.is_set():
							πF.SetLineno(1009)
						Label3:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("cannot set daemon status of active thread").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1010: raise RuntimeError("cannot set daemon status of active thread");
							πF.SetLineno(1010)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label4
						Label4:
							// line 1011: self.__daemonic = daemonic
							πF.SetLineno(1011)
							if πE = πg.CheckLocal(πF, µdaemonic, "daemonic"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdaemonic); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß__daemonic, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_daemon_setter.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 1013: daemon = property(_daemon_getter, _daemon_setter)
					πF.SetLineno(1013)
					πTemp005 = πF.MakeArgs(2)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ß_daemon_getter); πE != nil {
						continue
					}
					πTemp005[0] = πTemp022
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ß_daemon_setter); πE != nil {
						continue
					}
					πTemp005[1] = πTemp022
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp023, πE = πTemp022.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßdaemon.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 1015: def isDaemon(self):
					πF.SetLineno(1015)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("isDaemon", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1016: return self.daemon
							πF.SetLineno(1016)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdaemon, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßisDaemon.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 1018: def setDaemon(self, daemonic):
					πF.SetLineno(1018)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "daemonic", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("setDaemon", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdaemonic *πg.Object = πArgs[1]; _ = µdaemonic
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1019: self.daemon = daemonic
							πF.SetLineno(1019)
							if πE = πg.CheckLocal(πF, µdaemonic, "daemonic"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdaemonic); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdaemon, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsetDaemon.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 1021: def getName(self):
					πF.SetLineno(1021)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("getName", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1022: return self.name
							πF.SetLineno(1022)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetName.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 1024: def setName(self, name):
					πF.SetLineno(1024)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("setName", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1025: self.name = name
							πF.SetLineno(1025)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µname); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßname, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsetName.ToObject(), πTemp025); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp017, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp017 == nil {
				πTemp017 = πg.TypeType.ToObject()
			}
			if πTemp018, πE = πTemp017.Call(πF, []*πg.Object{πg.NewStr("Thread").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßThread.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 1029: def Timer(*args, **kwargs):
			πF.SetLineno(1029)
			πTemp010 = make([]πg.Param, 0)
			πTemp016 = πg.NewFunction(πg.NewCode("Timer", "/usr/lib/python2.7/threading.py", πTemp010, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
					// line 1030: """Factory function to create a Timer object.
					πF.SetLineno(1030)
					// line 1039: return _Timer(*args, **kwargs)
					πF.SetLineno(1039)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_Timer); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwargs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßTimer.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 1030: """Factory function to create a Timer object.
			πF.SetLineno(1030)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Factory function to create a Timer object.\n\n    Timers call a function after a specified number of seconds:\n\n        t = Timer(30.0, f, args=[], kwargs={})\n        t.start()\n        t.cancel()     # stop the timer's action if it's still waiting\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp018, πE = πg.ResolveGlobal(πF, ßTimer); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
				continue
			}
			// line 1041: class _Timer(Thread):
			πF.SetLineno(1041)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
				continue
			}
			πTemp002[0] = πTemp019
			πTemp009 = πg.NewDict()
			if πTemp017, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp017); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Timer", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Dict
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
					// line 1042: """Call a function after a specified number of seconds:
					πF.SetLineno(1042)
					// line 1042: """Call a function after a specified number of seconds:
					πF.SetLineno(1042)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Call a function after a specified number of seconds:\n\n            t = Timer(30.0, f, args=[], kwargs={})\n            t.start()\n            t.cancel()     # stop the timer's action if it's still waiting\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1050: def __init__(self, interval, function, args=[], kwargs={}):
					πF.SetLineno(1050)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "interval", Def: nil}
					πTemp002[2] = πg.Param{Name: "function", Def: nil}
					πTemp003 = make([]*πg.Object, 0)
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					πTemp002[3] = πg.Param{Name: "args", Def: πTemp004}
					πTemp005 = πg.NewDict()
					πTemp004 = πTemp005.ToObject()
					πTemp002[4] = πg.Param{Name: "kwargs", Def: πTemp004}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µinterval *πg.Object = πArgs[1]; _ = µinterval
						var µfunction *πg.Object = πArgs[2]; _ = µfunction
						var µargs *πg.Object = πArgs[3]; _ = µargs
						var µkwargs *πg.Object = πArgs[4]; _ = µkwargs
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
							// line 1051: Thread.__init__(self)
							πF.SetLineno(1051)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1052: self.interval = interval
							πF.SetLineno(1052)
							if πE = πg.CheckLocal(πF, µinterval, "interval"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µinterval); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinterval, πTemp002); πE != nil {
								continue
							}
							// line 1053: self.function = function
							πF.SetLineno(1053)
							if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µfunction); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfunction, πTemp002); πE != nil {
								continue
							}
							// line 1054: self.args = args
							πF.SetLineno(1054)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßargs, πTemp002); πE != nil {
								continue
							}
							// line 1055: self.kwargs = kwargs
							πF.SetLineno(1055)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µkwargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßkwargs, πTemp002); πE != nil {
								continue
							}
							// line 1056: self.finished = Event()
							πF.SetLineno(1056)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEvent); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfinished, πTemp002); πE != nil {
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
					// line 1058: def cancel(self):
					πF.SetLineno(1058)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("cancel", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1059: """Stop the timer if it hasn't finished yet"""
							πF.SetLineno(1059)
							// line 1060: self.finished.set()
							πF.SetLineno(1060)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfinished, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßset, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcancel.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1059: """Stop the timer if it hasn't finished yet"""
					πF.SetLineno(1059)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Stop the timer if it hasn't finished yet").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßcancel); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 1062: def run(self):
					πF.SetLineno(1062)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
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
							// line 1063: self.finished.wait(self.interval)
							πF.SetLineno(1063)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßinterval, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfinished, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwait, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfinished, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßis_set, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1064: if not self.finished.is_set():
							πF.SetLineno(1064)
						Label1:
							// line 1065: self.function(*self.args, **self.kwargs)
							πF.SetLineno(1065)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßargs, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßkwargs, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfunction, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Invoke(πF, πTemp004, nil, πTemp002, nil, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 1066: self.finished.set()
							πF.SetLineno(1066)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfinished, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp018, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp018 == nil {
				πTemp018 = πg.TypeType.ToObject()
			}
			if πTemp019, πE = πTemp018.Call(πF, []*πg.Object{πg.NewStr("_Timer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Timer.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 1071: class _MainThread(Thread):
			πF.SetLineno(1071)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
				continue
			}
			πTemp002[0] = πTemp019
			πTemp009 = πg.NewDict()
			if πTemp017, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp017); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_MainThread", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 1073: def __init__(self):
					πF.SetLineno(1073)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 πg.KWArgs
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
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.BaseException
						_ = πTemp012
						var πTemp013 *πg.Traceback
						_ = πTemp013
						var πTemp014 *πg.Type
						_ = πTemp014
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 1074: Thread.__init__(self, name="MainThread")
							πF.SetLineno(1074)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp002 = πg.KWArgs{
								{"name", ßMainThread.ToObject()},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1075: self.__started.set()
							πF.SetLineno(1075)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1076: self._set_ident()
							πF.SetLineno(1076)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_set_ident, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1077: with _active_limbo_lock:
							πF.SetLineno(1077)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 1078: _active[_get_ident()] = self
							πF.SetLineno(1078)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µself); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πE = πg.SetItem(πF, πTemp007, πTemp008, πTemp006); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp012, πTemp013 = nil, nil
							if πE != nil {
								πTemp012, πTemp013 = πF.ExcInfo()
							}
							if πTemp012 != nil {
								πTemp014 = πTemp012.Type()
								if πTemp006, πE = πTemp004.Call(πF, πg.Args{πTemp003, πTemp014.ToObject(), πTemp012.ToObject(), πTemp013.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp004.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp012 != nil && πTemp011 != true {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1080: def _set_daemon(self):
					πF.SetLineno(1080)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_set_daemon", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1081: return False
							πF.SetLineno(1081)
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
					if πE = πClass.SetItem(πF, ß_set_daemon.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1083: def _exitfunc(self):
					πF.SetLineno(1083)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_exitfunc", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 5: goto Label5
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							// line 1084: self.__stop()
							πF.SetLineno(1084)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__stop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1085: t = _pickSomeNonDaemonThread()
							πF.SetLineno(1085)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_pickSomeNonDaemonThread); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µt = πTemp002
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µt); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1086: if t:
							πF.SetLineno(1086)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1087: if __debug__:
							πF.SetLineno(1087)
						Label3:
							// line 1088: self._note("%s: waiting for other threads", self)
							πF.SetLineno(1088)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("%s: waiting for other threads").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 1089: while t:
							πF.SetLineno(1089)
							πF.PushCheckpoint(6)
							πTemp003 = false
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µt); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(5)            
							// line 1090: t.join()
							πF.SetLineno(1090)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1091: t = _pickSomeNonDaemonThread()
							πF.SetLineno(1091)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_pickSomeNonDaemonThread); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µt = πTemp002
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__debug__); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 1092: if __debug__:
							πF.SetLineno(1092)
						Label8:
							// line 1093: self._note("%s: exiting", self)
							πF.SetLineno(1093)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("%s: exiting").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label9
						Label9:
							// line 1094: self.__delete()
							πF.SetLineno(1094)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__delete, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_exitfunc.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp018, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp018 == nil {
				πTemp018 = πg.TypeType.ToObject()
			}
			if πTemp019, πE = πTemp018.Call(πF, []*πg.Object{πg.NewStr("_MainThread").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_MainThread.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 1096: def _pickSomeNonDaemonThread():
			πF.SetLineno(1096)
			πTemp010 = make([]πg.Param, 0)
			πTemp017 = πg.NewFunction(πg.NewCode("_pickSomeNonDaemonThread", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp004 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
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
						πTemp005 = !isStop
					} else {
						πTemp005 = true
						µt = πTemp002
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µt, ßdaemon, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µt, ßis_alive, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002 = πTemp006
				Label4:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label5
					}
					goto Label6
					// line 1098: if not t.daemon and t.is_alive():
					πF.SetLineno(1098)
				Label5:
					// line 1099: return t
					πF.SetLineno(1099)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πR = µt
					continue
					goto Label6
				Label6:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 1100: return None
					πF.SetLineno(1100)
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
			if πE = πF.Globals().SetItem(πF, ß_pickSomeNonDaemonThread.ToObject(), πTemp017); πE != nil {
				continue
			}
			// line 1111: class _DummyThread(Thread):
			πF.SetLineno(1111)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp020, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
				continue
			}
			πTemp002[0] = πTemp020
			πTemp009 = πg.NewDict()
			if πTemp018, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp018); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_DummyThread", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 1113: def __init__(self):
					πF.SetLineno(1113)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πTemp013 *πg.BaseException
						_ = πTemp013
						var πTemp014 *πg.Traceback
						_ = πTemp014
						var πTemp015 *πg.Type
						_ = πTemp015
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 1114: Thread.__init__(self, name=_newname("Dummy-%d"))
							πF.SetLineno(1114)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Dummy-%d").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_newname); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp005 = πg.KWArgs{
								{"name", πTemp004},
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1119: del self.__block
							πF.SetLineno(1119)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ß__block); πE != nil {
								continue
							}
							// line 1121: self.__started.set()
							πF.SetLineno(1121)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__started, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1122: self._set_ident()
							πF.SetLineno(1122)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_set_ident, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1123: with _active_limbo_lock:
							πF.SetLineno(1123)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp006.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 1124: _active[_get_ident()] = self
							πF.SetLineno(1124)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, µself); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
								continue
							}
							if πTemp010, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp009 = πTemp011
							if πE = πg.SetItem(πF, πTemp008, πTemp009, πTemp007); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label1:
							πTemp013, πTemp014 = nil, nil
							if πE != nil {
								πTemp013, πTemp014 = πF.ExcInfo()
							}
							if πTemp013 != nil {
								πTemp015 = πTemp013.Type()
								if πTemp007, πE = πTemp004.Call(πF, πg.Args{πTemp003, πTemp015.ToObject(), πTemp013.ToObject(), πTemp014.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp007, πE = πTemp004.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp012, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp013 != nil && πTemp012 != true {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1126: def _set_daemon(self):
					πF.SetLineno(1126)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_set_daemon", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1127: return True
							πF.SetLineno(1127)
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
					if πE = πClass.SetItem(πF, ß_set_daemon.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1129: def join(self, timeout=None):
					πF.SetLineno(1129)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "timeout", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("join", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtimeout *πg.Object = πArgs[1]; _ = µtimeout
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1130: assert False, "cannot join a dummy thread"
							πF.SetLineno(1130)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("cannot join a dummy thread").ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßjoin.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp019, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp019 == nil {
				πTemp019 = πg.TypeType.ToObject()
			}
			if πTemp020, πE = πTemp019.Call(πF, []*πg.Object{πg.NewStr("_DummyThread").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_DummyThread.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 1135: def currentThread():
			πF.SetLineno(1135)
			πTemp010 = make([]πg.Param, 0)
			πTemp018 = πg.NewFunction(πg.NewCode("currentThread", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.BaseException
				_ = πTemp004
				var πTemp005 *πg.Traceback
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 1136: """Return the current Thread object, corresponding to the caller's thread of control.
					πF.SetLineno(1136)
					// line 1142: try:
					πF.SetLineno(1142)
					πF.PushCheckpoint(2)
					// line 1143: return _active[_get_ident()]
					πF.SetLineno(1143)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
					// line 1144: except KeyError:
					πF.SetLineno(1144)
				Label3:
					// line 1146: return _DummyThread()
					πF.SetLineno(1146)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_DummyThread); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πR = πTemp002
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
			if πE = πF.Globals().SetItem(πF, ßcurrentThread.ToObject(), πTemp018); πE != nil {
				continue
			}
			// line 1136: """Return the current Thread object, corresponding to the caller's thread of control.
			πF.SetLineno(1136)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("Return the current Thread object, corresponding to the caller's thread of control.\n\n    If the caller's thread of control was not created through the threading\n    module, a dummy thread object with limited functionality is returned.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp020, πE = πg.ResolveGlobal(πF, ßcurrentThread); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
				continue
			}
			// line 1148: current_thread = currentThread
			πF.SetLineno(1148)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßcurrentThread); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcurrent_thread.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 1150: def activeCount():
			πF.SetLineno(1150)
			πTemp010 = make([]πg.Param, 0)
			πTemp019 = πg.NewFunction(πg.NewCode("activeCount", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					default: panic("unexpected function state")
					}
					// line 1151: """Return the number of Thread objects currently alive.
					πF.SetLineno(1151)
					// line 1157: with _active_limbo_lock:
					πF.SetLineno(1157)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
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
					πF.PushCheckpoint(1)
					// line 1158: return len(_active) + len(_limbo)
					πF.SetLineno(1158)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_limbo); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.Add(πF, πTemp007, πTemp008); πE != nil {
						continue
					}
					πR = πTemp004
					continue
					πF.PopCheckpoint()
				Label1:
					πTemp010, πTemp011 = nil, nil
					if πE != nil {
						πTemp010, πTemp011 = πF.ExcInfo()
					}
					if πTemp010 != nil {
						πTemp012 = πTemp010.Type()
						if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßactiveCount.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 1151: """Return the number of Thread objects currently alive.
			πF.SetLineno(1151)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("Return the number of Thread objects currently alive.\n\n    The returned count is equal to the length of the list returned by\n    enumerate().\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßactiveCount); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
				continue
			}
			// line 1160: active_count = activeCount
			πF.SetLineno(1160)
			if πTemp020, πE = πg.ResolveGlobal(πF, ßactiveCount); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßactive_count.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 1162: def _enumerate():
			πF.SetLineno(1162)
			πTemp010 = make([]πg.Param, 0)
			πTemp020 = πg.NewFunction(πg.NewCode("_enumerate", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
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
					// line 1164: return _active.values() + _limbo.values()
					πF.SetLineno(1164)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_limbo); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_enumerate.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 1166: def enumerate():
			πF.SetLineno(1166)
			πTemp010 = make([]πg.Param, 0)
			πTemp021 = πg.NewFunction(πg.NewCode("enumerate", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πTemp010 *πg.Traceback
				_ = πTemp010
				var πTemp011 *πg.Type
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					default: panic("unexpected function state")
					}
					// line 1167: """Return a list of all Thread objects currently alive.
					πF.SetLineno(1167)
					// line 1174: with _active_limbo_lock:
					πF.SetLineno(1174)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
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
					πF.PushCheckpoint(1)
					// line 1175: return _active.values() + _limbo.values()
					πF.SetLineno(1175)
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_limbo); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßvalues, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πR = πTemp004
					continue
					πF.PopCheckpoint()
				Label1:
					πTemp009, πTemp010 = nil, nil
					if πE != nil {
						πTemp009, πTemp010 = πF.ExcInfo()
					}
					if πTemp009 != nil {
						πTemp011 = πTemp009.Type()
						if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp004, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp009 != nil && πTemp008 != true {
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
			if πE = πF.Globals().SetItem(πF, ßenumerate.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 1167: """Return a list of all Thread objects currently alive.
			πF.SetLineno(1167)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("Return a list of all Thread objects currently alive.\n\n    The list includes daemonic threads, dummy thread objects created by\n    current_thread(), and the main thread. It excludes terminated threads and\n    threads that have not yet been started.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp023, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
				continue
			}
			// line 1177: from thread import stack_size
			πF.SetLineno(1177)
			if πTemp002, πE = πg.ImportModule(πF, "thread"); πE != nil {
				continue
			}
			πTemp022 = πTemp002[0]
			if πTemp023, πE = πg.GetAttrImport(πF, πTemp022, ßstack_size); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstack_size.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 1183: _shutdown = _MainThread()._exitfunc
			πF.SetLineno(1183)
			if πTemp022, πE = πg.ResolveGlobal(πF, ß_MainThread); πE != nil {
				continue
			}
			if πTemp023, πE = πTemp022.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp022, πE = πg.GetAttr(πF, πTemp023, ß_exitfunc, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_shutdown.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 1191: class _localbase(object):
			πF.SetLineno(1191)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp024, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp024
			πTemp009 = πg.NewDict()
			if πTemp022, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp022); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_localbase", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 1192: __slots__ = '_local__key', '_local__args', '_local__lock'
					πF.SetLineno(1192)
					πTemp001 = πg.NewTuple3(ß_local__key.ToObject(), ß_local__args.ToObject(), ß_local__lock.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1194: def __new__(cls, *args, **kw):
					πF.SetLineno(1194)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/threading.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkw *πg.Object = πArgs[2]; _ = µkw
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µdict *πg.Object = πg.UnboundLocal; _ = µdict
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
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
							default: panic("unexpected function state")
							}
							// line 1195: self = object.__new__(cls)
							πF.SetLineno(1195)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µself = πTemp002
							// line 1196: key = '_local__key', 'thread.local.' + str(id(self))
							πF.SetLineno(1196)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp005, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp006
							if πTemp005, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Add(πF, πg.NewStr("thread.local.").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ß_local__key.ToObject(), πTemp003).ToObject()
							µkey = πTemp002
							// line 1197: object.__setattr__(self, '_local__key', key)
							πF.SetLineno(1197)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß_local__key.ToObject()
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[2] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__setattr__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1198: object.__setattr__(self, '_local__args', (args, kw))
							πF.SetLineno(1198)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß_local__args.ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µargs, µkw).ToObject()
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__setattr__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1199: object.__setattr__(self, '_local__lock', RLock())
							πF.SetLineno(1199)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß_local__lock.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßRLock); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__setattr__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003 = µargs
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
								continue
							}
							πTemp003 = µkw
						Label2:
							πTemp002 = πTemp003
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcls, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp006, ß__init__, nil); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005 == πTemp009).ToObject()
							πTemp002 = πTemp003
						Label1:
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							goto Label4
							// line 1201: if (args or kw) and (cls.__init__ is object.__init__):
							πF.SetLineno(1201)
						Label3:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Initialization arguments are not supported").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1202: raise TypeError("Initialization arguments are not supported")
							πF.SetLineno(1202)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							// line 1207: dict = object.__getattribute__(self, '__dict__')
							πF.SetLineno(1207)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß__dict__.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__getattribute__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdict = πTemp002
							// line 1208: current_thread().__dict__[key] = dict
							πF.SetLineno(1208)
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdict); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcurrent_thread); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp005, ß__dict__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp005 = µkey
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 1210: return self
							πF.SetLineno(1210)
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp023, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp023 == nil {
				πTemp023 = πg.TypeType.ToObject()
			}
			if πTemp024, πE = πTemp023.Call(πF, []*πg.Object{πg.NewStr("_localbase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_localbase.ToObject(), πTemp024); πE != nil {
				continue
			}
			// line 1212: def _patch(self):
			πF.SetLineno(1212)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "self", Def: nil}
			πTemp022 = πg.NewFunction(πg.NewCode("_patch", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]; _ = µself
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µcls *πg.Object = πg.UnboundLocal; _ = µcls
				var µargs *πg.Object = πg.UnboundLocal; _ = µargs
				var µkw *πg.Object = πg.UnboundLocal; _ = µkw
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 bool
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
					// line 1213: key = object.__getattribute__(self, '_local__key')
					πF.SetLineno(1213)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp001[0] = µself
					πTemp001[1] = ß_local__key.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__getattribute__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µkey = πTemp002
					// line 1214: d = current_thread().__dict__.get(key)
					πF.SetLineno(1214)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp001[0] = µkey
					if πTemp002, πE = πg.ResolveGlobal(πF, ßcurrent_thread); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__dict__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µd = πTemp002
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µd == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 1215: if d is None:
					πF.SetLineno(1215)
				Label1:
					// line 1216: d = {}
					πF.SetLineno(1216)
					πTemp005 = πg.NewDict()
					πTemp002 = πTemp005.ToObject()
					µd = πTemp002
					// line 1217: current_thread().__dict__[key] = d
					πF.SetLineno(1217)
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µd); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßcurrent_thread); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp006, ß__dict__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp006 = µkey
					if πE = πg.SetItem(πF, πTemp003, πTemp006, πTemp002); πE != nil {
						continue
					}
					// line 1218: object.__setattr__(self, '__dict__', d)
					πF.SetLineno(1218)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp001[0] = µself
					πTemp001[1] = ß__dict__.ToObject()
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp001[2] = µd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__setattr__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 1222: cls = type(self)
					πF.SetLineno(1222)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp001[0] = µself
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcls = πTemp003
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µcls, ß__init__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß__init__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp003 != πTemp007).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 1223: if cls.__init__ is not object.__init__:
					πF.SetLineno(1223)
				Label4:
					// line 1224: args, kw = object.__getattribute__(self, '_local__args')
					πF.SetLineno(1224)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp001[0] = µself
					πTemp001[1] = ß_local__args.ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__getattribute__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
						continue
					}
					µargs = πTemp003
					µkw = πTemp006
					// line 1225: cls.__init__(self, *args, **kw)
					πF.SetLineno(1225)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp001[0] = µself
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µcls, ß__init__, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, µkw); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label5
				Label5:
					goto Label3
				Label2:
					// line 1227: object.__setattr__(self, '__dict__', d)
					πF.SetLineno(1227)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp001[0] = µself
					πTemp001[1] = ß__dict__.ToObject()
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp001[2] = µd
					if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__setattr__, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
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
			if πE = πF.Globals().SetItem(πF, ß_patch.ToObject(), πTemp022); πE != nil {
				continue
			}
			// line 1229: class local(_localbase):
			πF.SetLineno(1229)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp025, πE = πg.ResolveGlobal(πF, ß_localbase); πE != nil {
				continue
			}
			πTemp002[0] = πTemp025
			πTemp009 = πg.NewDict()
			if πTemp023, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp023); πE != nil {
				continue
			}
			_, πE = πg.NewCode("local", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
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
					// line 1231: def __getattribute__(self, name):
					πF.SetLineno(1231)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__getattribute__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 1232: lock = object.__getattribute__(self, '_local__lock')
							πF.SetLineno(1232)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß_local__lock.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__getattribute__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlock = πTemp002
							// line 1233: lock.acquire()
							πF.SetLineno(1233)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1234: try:
							πF.SetLineno(1234)
							πF.PushCheckpoint(1)
							// line 1235: _patch(self)
							πF.SetLineno(1235)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_patch); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1236: return object.__getattribute__(self, name)
							πF.SetLineno(1236)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__getattribute__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp004, πTemp005 = πF.RestoreExc(nil, nil)
							// line 1238: lock.release()
							πF.SetLineno(1238)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__getattribute__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1240: def __setattr__(self, name, value):
					πF.SetLineno(1240)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__setattr__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µname, ß__dict__.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 1241: if name == '__dict__':
							πF.SetLineno(1241)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%r object attribute '__dict__' is read-only").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1242: raise AttributeError(
							πF.SetLineno(1242)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 1245: lock = object.__getattribute__(self, '_local__lock')
							πF.SetLineno(1245)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							πTemp003[1] = ß_local__lock.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ß__getattribute__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µlock = πTemp001
							// line 1246: lock.acquire()
							πF.SetLineno(1246)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1247: try:
							πF.SetLineno(1247)
							πF.PushCheckpoint(3)
							// line 1248: _patch(self)
							πF.SetLineno(1248)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_patch); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1249: return object.__setattr__(self, name, value)
							πF.SetLineno(1249)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp003[1] = µname
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[2] = µvalue
							if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ß__setattr__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label3
						Label3:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							// line 1251: lock.release()
							πF.SetLineno(1251)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__setattr__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1253: def __delattr__(self, name):
					πF.SetLineno(1253)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__delattr__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µlock *πg.Object = πg.UnboundLocal; _ = µlock
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µname, ß__dict__.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 1254: if name == '__dict__':
							πF.SetLineno(1254)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%r object attribute '__dict__' is read-only").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1255: raise AttributeError(
							πF.SetLineno(1255)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 1258: lock = object.__getattribute__(self, '_local__lock')
							πF.SetLineno(1258)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							πTemp003[1] = ß_local__lock.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ß__getattribute__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µlock = πTemp001
							// line 1259: lock.acquire()
							πF.SetLineno(1259)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßacquire, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1260: try:
							πF.SetLineno(1260)
							πF.PushCheckpoint(3)
							// line 1261: _patch(self)
							πF.SetLineno(1261)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_patch); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1262: return object.__delattr__(self, name)
							πF.SetLineno(1262)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp003[1] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ß__delattr__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label3
						Label3:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							// line 1264: lock.release()
							πF.SetLineno(1264)
							if πE = πg.CheckLocal(πF, µlock, "lock"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlock, ßrelease, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__delattr__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1266: def __del__(self):
					πF.SetLineno(1266)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__del__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µthreads *πg.Object = πg.UnboundLocal; _ = µthreads
						var µthread *πg.Object = πg.UnboundLocal; _ = µthread
						var µ__dict__ *πg.Object = πg.UnboundLocal; _ = µ__dict__
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8: goto Label8
							case 2: goto Label2
							case 4: goto Label4
							case 5: goto Label5
							case 13: goto Label13
							default: panic("unexpected function state")
							}
							// line 1267: key = object.__getattribute__(self, '_local__key')
							πF.SetLineno(1267)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							πTemp001[1] = ß_local__key.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__getattribute__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µkey = πTemp002
							// line 1269: try:
							πF.SetLineno(1269)
							πF.PushCheckpoint(2)
							// line 1272: threads = _enumerate()
							πF.SetLineno(1272)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_enumerate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µthreads = πTemp003
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							goto Label3
							// line 1273: except:
							πF.SetLineno(1273)
						Label3:
							// line 1277: return
							πF.SetLineno(1277)
							πR = πg.None
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µthreads); πE != nil {
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
								µthread = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 1280: try:
							πF.SetLineno(1280)
							πF.PushCheckpoint(8)
							// line 1281: __dict__ = thread.__dict__
							πF.SetLineno(1281)
							if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µthread, ß__dict__, nil); πE != nil {
								continue
							}
							µ__dict__ = πTemp003
							πF.PopCheckpoint()
							goto Label7
						Label8:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 1282: except AttributeError:
							πF.SetLineno(1282)
						Label9:
							// line 1284: continue
							πF.SetLineno(1284)
							continue
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µ__dict__, "__dict__"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µ__dict__, µkey); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							goto Label11
							// line 1286: if key in __dict__:
							πF.SetLineno(1286)
						Label10:
							// line 1287: try:
							πF.SetLineno(1287)
							πF.PushCheckpoint(13)
							// line 1288: del __dict__[key]
							πF.SetLineno(1288)
							if πE = πg.CheckLocal(πF, µ__dict__, "__dict__"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.DelItem(πF, µ__dict__, πTemp003); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label14
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 1289: except KeyError:
							πF.SetLineno(1289)
						Label14:
							// line 1290: pass # didn't have anything in this thread
							πF.SetLineno(1290)
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
							goto Label11
						Label11:
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
					if πE = πClass.SetItem(πF, ß__del__.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp024, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp024 == nil {
				πTemp024 = πg.TypeType.ToObject()
			}
			if πTemp025, πE = πTemp024.Call(πF, []*πg.Object{πg.NewStr("local").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßlocal.ToObject(), πTemp025); πE != nil {
				continue
			}
			// line 1294: def _after_fork():
			πF.SetLineno(1294)
			πTemp010 = make([]πg.Param, 0)
			πTemp023 = πg.NewFunction(πg.NewCode("_after_fork", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µnew_active *πg.Object = πg.UnboundLocal; _ = µnew_active
				var µcurrent *πg.Object = πg.UnboundLocal; _ = µcurrent
				var µthread *πg.Object = πg.UnboundLocal; _ = µthread
				var µident *πg.Object = πg.UnboundLocal; _ = µident
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Dict
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
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πTemp013 *πg.Type
				_ = πTemp013
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
					// line 1301: global _active_limbo_lock
					πF.SetLineno(1301)
					// line 1302: _active_limbo_lock = _allocate_lock()
					πF.SetLineno(1302)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_allocate_lock); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_active_limbo_lock.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 1305: new_active = {}
					πF.SetLineno(1305)
					πTemp003 = πg.NewDict()
					πTemp001 = πTemp003.ToObject()
					µnew_active = πTemp001
					// line 1306: current = current_thread()
					πF.SetLineno(1306)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßcurrent_thread); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µcurrent = πTemp002
					// line 1307: with _active_limbo_lock:
					πF.SetLineno(1307)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_active_limbo_lock); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
						continue
					}
					πF.PushCheckpoint(1)
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_enumerate); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(3)
					πTemp008 = false
				Label2:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
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
						πTemp009 = !isStop
					} else {
						πTemp009 = true
						µthread = πTemp006
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(2)            
					πTemp010 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
						continue
					}
					πTemp010[0] = µthread
					πTemp010[1] = ß_reset_internal_locks.ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp009, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label5
					}
					goto Label6
					// line 1311: if hasattr(thread, '_reset_internal_locks'):
					πF.SetLineno(1311)
				Label5:
					// line 1312: thread._reset_internal_locks()
					πF.SetLineno(1312)
					if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µthread, ß_reset_internal_locks, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label6
				Label6:
					if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcurrent, "current"); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(µthread == µcurrent).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label7
					}
					goto Label8
					// line 1313: if thread is current:
					πF.SetLineno(1313)
				Label7:
					// line 1316: ident = _get_ident()
					πF.SetLineno(1316)
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_get_ident); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					µident = πTemp007
					// line 1317: thread.__ident = ident
					πF.SetLineno(1317)
					if πE = πg.CheckLocal(πF, µident, "ident"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µident); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µthread, ß__ident, πTemp006); πE != nil {
						continue
					}
					// line 1318: new_active[ident] = thread
					πF.SetLineno(1318)
					if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µthread); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnew_active, "new_active"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µident, "ident"); πE != nil {
						continue
					}
					πTemp007 = µident
					if πE = πg.SetItem(πF, µnew_active, πTemp007, πTemp006); πE != nil {
						continue
					}
					goto Label9
				Label8:
					// line 1321: thread.__stop()
					πF.SetLineno(1321)
					if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µthread, ß__stop, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label9
				Label9:
					continue
				Label3:
					if πE != nil || πR != nil {
						continue
					}
				Label4:
					// line 1323: _limbo.clear()
					πF.SetLineno(1323)
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_limbo); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßclear, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 1324: _active.clear()
					πF.SetLineno(1324)
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßclear, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 1325: _active.update(new_active)
					πF.SetLineno(1325)
					πTemp010 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnew_active, "new_active"); πE != nil {
						continue
					}
					πTemp010[0] = µnew_active
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßupdate, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					// line 1326: assert len(_active) == 1
					πF.SetLineno(1326)
					πTemp010 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveGlobal(πF, ß_active); πE != nil {
						continue
					}
					πTemp010[0] = πTemp006
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πTemp005, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.PopCheckpoint()
				Label1:
					πTemp011, πTemp012 = nil, nil
					if πE != nil {
						πTemp011, πTemp012 = πF.ExcInfo()
					}
					if πTemp011 != nil {
						πTemp013 = πTemp011.Type()
						if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp013.ToObject(), πTemp011.ToObject(), πTemp012.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
							continue
						}
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp011 != nil && πTemp008 != true {
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
			if πE = πF.Globals().SetItem(πF, ß_after_fork.ToObject(), πTemp023); πE != nil {
				continue
			}
			// line 1331: def _test():
			πF.SetLineno(1331)
			πTemp010 = make([]πg.Param, 0)
			πTemp024 = πg.NewFunction(πg.NewCode("_test", "/usr/lib/python2.7/threading.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µBoundedQueue *πg.Object = πg.UnboundLocal; _ = µBoundedQueue
				var µProducerThread *πg.Object = πg.UnboundLocal; _ = µProducerThread
				var µConsumerThread *πg.Object = πg.UnboundLocal; _ = µConsumerThread
				var µNP *πg.Object = πg.UnboundLocal; _ = µNP
				var µQL *πg.Object = πg.UnboundLocal; _ = µQL
				var µNI *πg.Object = πg.UnboundLocal; _ = µNI
				var µQ *πg.Object = πg.UnboundLocal; _ = µQ
				var µP *πg.Object = πg.UnboundLocal; _ = µP
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var µC *πg.Object = πg.UnboundLocal; _ = µC
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
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
					// line 1333: class BoundedQueue(_Verbose):
					πF.SetLineno(1333)
					πTemp003 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
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
					_, πE = πg.NewCode("BoundedQueue", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
						πClass := πTemp001
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
							// line 1335: def __init__(self, limit):
							πF.SetLineno(1335)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp002[1] = πg.Param{Name: "limit", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µlimit *πg.Object = πArgs[1]; _ = µlimit
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
									// line 1336: _Verbose.__init__(self)
									πF.SetLineno(1336)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									πTemp001[0] = µself
									if πTemp002, πE = πg.ResolveGlobal(πF, ß_Verbose); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 1337: self.mon = RLock()
									πF.SetLineno(1337)
									if πTemp002, πE = πg.ResolveGlobal(πF, ßRLock); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßmon, πTemp002); πE != nil {
										continue
									}
									// line 1338: self.rc = Condition(self.mon)
									πF.SetLineno(1338)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßmon, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßCondition); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßrc, πTemp002); πE != nil {
										continue
									}
									// line 1339: self.wc = Condition(self.mon)
									πF.SetLineno(1339)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßmon, nil); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßCondition); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßwc, πTemp002); πE != nil {
										continue
									}
									// line 1340: self.limit = limit
									πF.SetLineno(1340)
									if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µlimit); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßlimit, πTemp002); πE != nil {
										continue
									}
									// line 1341: self.queue = _deque()
									πF.SetLineno(1341)
									if πTemp002, πE = πg.ResolveGlobal(πF, ß_deque); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßqueue, πTemp002); πE != nil {
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
							// line 1343: def put(self, item):
							πF.SetLineno(1343)
							πTemp002 = make([]πg.Param, 2)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp002[1] = πg.Param{Name: "item", Def: nil}
							πTemp003 = πg.NewFunction(πg.NewCode("put", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µitem *πg.Object = πArgs[1]; _ = µitem
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
								var πTemp006 *πg.Object
								_ = πTemp006
								var πTemp007 []*πg.Object
								_ = πTemp007
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									case 1: goto Label1
									case 2: goto Label2
									default: panic("unexpected function state")
									}
									// line 1344: self.mon.acquire()
									πF.SetLineno(1344)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßmon, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 1345: while len(self.queue) >= self.limit:
									πF.SetLineno(1345)
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
									πTemp005 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
										continue
									}
									πTemp005[0] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp005)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßlimit, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GE(πF, πTemp006, πTemp002); πE != nil {
										continue
									}
									if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πE != nil || !πTemp004 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 1346: self._note("put(%s): queue full", item)
									πF.SetLineno(1346)
									πTemp005 = πF.MakeArgs(2)
									πTemp005[0] = πg.NewStr("put(%s): queue full").ToObject()
									if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
										continue
									}
									πTemp005[1] = µitem
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp005)
									// line 1347: self.wc.wait()
									πF.SetLineno(1347)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßwc, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwait, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									continue
								Label2:
									if πE != nil || πR != nil {
										continue
									}
								Label3:
									// line 1348: self.queue.append(item)
									πF.SetLineno(1348)
									πTemp005 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
										continue
									}
									πTemp005[0] = µitem
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp005)
									// line 1349: self._note("put(%s): appended, length now %d",
									πF.SetLineno(1349)
									πTemp005 = πF.MakeArgs(3)
									πTemp005[0] = πg.NewStr("put(%s): appended, length now %d").ToObject()
									if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
										continue
									}
									πTemp005[1] = µitem
									πTemp007 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
										continue
									}
									πTemp007[0] = πTemp001
									if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp007)
									πTemp005[2] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp005)
									// line 1351: self.rc.notify()
									πF.SetLineno(1351)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßrc, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnotify, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 1352: self.mon.release()
									πF.SetLineno(1352)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßmon, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrelease, nil); πE != nil {
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
							if πE = πClass.SetItem(πF, ßput.ToObject(), πTemp003); πE != nil {
								continue
							}
							// line 1354: def get(self):
							πF.SetLineno(1354)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp004 = πg.NewFunction(πg.NewCode("get", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µitem *πg.Object = πg.UnboundLocal; _ = µitem
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 bool
								_ = πTemp004
								var πTemp005 bool
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
									case 1: goto Label1
									case 2: goto Label2
									default: panic("unexpected function state")
									}
									// line 1355: self.mon.acquire()
									πF.SetLineno(1355)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßmon, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßacquire, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 1356: while not self.queue:
									πF.SetLineno(1356)
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
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
										continue
									}
									if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(!πTemp005).ToObject()
									if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πE != nil || !πTemp004 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 1357: self._note("get(): queue empty")
									πF.SetLineno(1357)
									πTemp006 = πF.MakeArgs(1)
									πTemp006[0] = πg.NewStr("get(): queue empty").ToObject()
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp006)
									// line 1358: self.rc.wait()
									πF.SetLineno(1358)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßrc, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwait, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									continue
								Label2:
									if πE != nil || πR != nil {
										continue
									}
								Label3:
									// line 1359: item = self.queue.popleft()
									πF.SetLineno(1359)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpopleft, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									µitem = πTemp001
									// line 1360: self._note("get(): got %s, %d left", item, len(self.queue))
									πF.SetLineno(1360)
									πTemp006 = πF.MakeArgs(3)
									πTemp006[0] = πg.NewStr("get(): got %s, %d left").ToObject()
									if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
										continue
									}
									πTemp006[1] = µitem
									πTemp007 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
										continue
									}
									πTemp007[0] = πTemp001
									if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp007)
									πTemp006[2] = πTemp002
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ß_note, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp006)
									// line 1361: self.wc.notify()
									πF.SetLineno(1361)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßwc, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnotify, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 1362: self.mon.release()
									πF.SetLineno(1362)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßmon, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrelease, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
										continue
									}
									// line 1363: return item
									πF.SetLineno(1363)
									if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
										continue
									}
									πR = µitem
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp004); πE != nil {
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
					if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BoundedQueue").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					µBoundedQueue = πTemp005
					// line 1365: class ProducerThread(Thread):
					πF.SetLineno(1365)
					πTemp003 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
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
					_, πE = πg.NewCode("ProducerThread", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1367: def __init__(self, queue, quota):
							πF.SetLineno(1367)
							πTemp002 = make([]πg.Param, 3)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp002[1] = πg.Param{Name: "queue", Def: nil}
							πTemp002[2] = πg.Param{Name: "quota", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µqueue *πg.Object = πArgs[1]; _ = µqueue
								var µquota *πg.Object = πArgs[2]; _ = µquota
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
									// line 1368: Thread.__init__(self, name="Producer")
									πF.SetLineno(1368)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									πTemp001[0] = µself
									πTemp002 = πg.KWArgs{
										{"name", ßProducer.ToObject()},
									}
									if πTemp003, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp002); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 1369: self.queue = queue
									πF.SetLineno(1369)
									if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µqueue); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßqueue, πTemp003); πE != nil {
										continue
									}
									// line 1370: self.quota = quota
									πF.SetLineno(1370)
									if πE = πg.CheckLocal(πF, µquota, "quota"); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µquota); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßquota, πTemp003); πE != nil {
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
							// line 1372: def run(self):
							πF.SetLineno(1372)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp003 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µrandom *πg.Object = πg.UnboundLocal; _ = µrandom
								var µcounter *πg.Object = πg.UnboundLocal; _ = µcounter
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 []*πg.Object
								_ = πTemp002
								var πTemp003 *πg.Object
								_ = πTemp003
								var πTemp004 bool
								_ = πTemp004
								var πTemp005 bool
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
									// line 1373: from random import random
									πF.SetLineno(1373)
									if πTemp002, πE = πg.ImportModule(πF, "random"); πE != nil {
										continue
									}
									πTemp001 = πTemp002[0]
									if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßrandom); πE != nil {
										continue
									}
									µrandom = πTemp003
									// line 1374: counter = 0
									πF.SetLineno(1374)
									µcounter = πg.NewInt(0).ToObject()
									// line 1375: while counter < self.quota:
									πF.SetLineno(1375)
									πF.PushCheckpoint(2)
									πTemp004 = false
								Label1:
									if πE != nil || πR != nil {
										continue
									}
									if πTemp004 {
										πF.PopCheckpoint()
										goto Label3
									}
									if πE = πg.CheckLocal(πF, µcounter, "counter"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßquota, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πg.LT(πF, µcounter, πTemp003); πE != nil {
										continue
									}
									if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πE != nil || !πTemp005 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 1376: counter = counter + 1
									πF.SetLineno(1376)
									if πE = πg.CheckLocal(πF, µcounter, "counter"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Add(πF, µcounter, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									µcounter = πTemp001
									// line 1377: self.queue.put("%s.%d" % (self.name, counter))
									πF.SetLineno(1377)
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp006, πE = πg.GetAttr(πF, µself, ßname, nil); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µcounter, "counter"); πE != nil {
										continue
									}
									πTemp003 = πg.NewTuple2(πTemp006, µcounter).ToObject()
									if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s.%d").ToObject(), πTemp003); πE != nil {
										continue
									}
									πTemp002[0] = πTemp001
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßput, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									// line 1378: _sleep(random() * 0.00001)
									πF.SetLineno(1378)
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µrandom, "random"); πE != nil {
										continue
									}
									if πTemp003, πE = µrandom.Call(πF, nil, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Mul(πF, πTemp003, πg.NewFloat(1e-05).ToObject()); πE != nil {
										continue
									}
									πTemp002[0] = πTemp001
									if πTemp001, πE = πg.ResolveGlobal(πF, ß_sleep); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
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
							if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp003); πE != nil {
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
					if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("ProducerThread").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					µProducerThread = πTemp005
					// line 1381: class ConsumerThread(Thread):
					πF.SetLineno(1381)
					πTemp003 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
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
					_, πE = πg.NewCode("ConsumerThread", "/usr/lib/python2.7/threading.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1383: def __init__(self, queue, count):
							πF.SetLineno(1383)
							πTemp002 = make([]πg.Param, 3)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp002[1] = πg.Param{Name: "queue", Def: nil}
							πTemp002[2] = πg.Param{Name: "count", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µqueue *πg.Object = πArgs[1]; _ = µqueue
								var µcount *πg.Object = πArgs[2]; _ = µcount
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
									// line 1384: Thread.__init__(self, name="Consumer")
									πF.SetLineno(1384)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									πTemp001[0] = µself
									πTemp002 = πg.KWArgs{
										{"name", ßConsumer.ToObject()},
									}
									if πTemp003, πE = πg.ResolveGlobal(πF, ßThread); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp004.Call(πF, πTemp001, πTemp002); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 1385: self.queue = queue
									πF.SetLineno(1385)
									if πE = πg.CheckLocal(πF, µqueue, "queue"); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µqueue); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßqueue, πTemp003); πE != nil {
										continue
									}
									// line 1386: self.count = count
									πF.SetLineno(1386)
									if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µcount); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßcount, πTemp003); πE != nil {
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
							// line 1388: def run(self):
							πF.SetLineno(1388)
							πTemp002 = make([]πg.Param, 1)
							πTemp002[0] = πg.Param{Name: "self", Def: nil}
							πTemp003 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/threading.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µself *πg.Object = πArgs[0]; _ = µself
								var µitem *πg.Object = πg.UnboundLocal; _ = µitem
								var πTemp001 bool
								_ = πTemp001
								var πTemp002 bool
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
									case 1: goto Label1
									case 2: goto Label2
									default: panic("unexpected function state")
									}
									// line 1389: while self.count > 0:
									πF.SetLineno(1389)
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
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GT(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
										continue
									}
									if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									if πE != nil || !πTemp002 {
										continue
									}
									πF.PushCheckpoint(1)            
									// line 1390: item = self.queue.get()
									πF.SetLineno(1390)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
										continue
									}
									µitem = πTemp003
									// line 1391: print item
									πF.SetLineno(1391)
									πTemp005 = make([]*πg.Object, 1)
									if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
										continue
									}
									πTemp005[0] = µitem
									if πE = πg.Print(πF, πTemp005, true); πE != nil {
										continue
									}
									// line 1392: self.count = self.count - 1
									πF.SetLineno(1392)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetAttr(πF, µself, ßcount, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßcount, πTemp004); πE != nil {
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
							if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp003); πE != nil {
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
					if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("ConsumerThread").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
						continue
					}
					µConsumerThread = πTemp005
					// line 1394: NP = 3
					πF.SetLineno(1394)
					µNP = πg.NewInt(3).ToObject()
					// line 1395: QL = 4
					πF.SetLineno(1395)
					µQL = πg.NewInt(4).ToObject()
					// line 1396: NI = 5
					πF.SetLineno(1396)
					µNI = πg.NewInt(5).ToObject()
					// line 1398: Q = BoundedQueue(QL)
					πF.SetLineno(1398)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µQL, "QL"); πE != nil {
						continue
					}
					πTemp003[0] = µQL
					if πE = πg.CheckLocal(πF, µBoundedQueue, "BoundedQueue"); πE != nil {
						continue
					}
					if πTemp002, πE = µBoundedQueue.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µQ = πTemp002
					// line 1399: P = []
					πF.SetLineno(1399)
					πTemp003 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp003...).ToObject()
					µP = πTemp002
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µNP, "NP"); πE != nil {
						continue
					}
					πTemp003[0] = µNP
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
					// line 1401: t = ProducerThread(Q, NI)
					πF.SetLineno(1401)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µQ, "Q"); πE != nil {
						continue
					}
					πTemp003[0] = µQ
					if πE = πg.CheckLocal(πF, µNI, "NI"); πE != nil {
						continue
					}
					πTemp003[1] = µNI
					if πE = πg.CheckLocal(πF, µProducerThread, "ProducerThread"); πE != nil {
						continue
					}
					if πTemp004, πE = µProducerThread.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µt = πTemp004
					// line 1402: t.name = ("Producer-%d" % (i+1))
					πF.SetLineno(1402)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mod(πF, πg.NewStr("Producer-%d").ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µt, ßname, πTemp005); πE != nil {
						continue
					}
					// line 1403: P.append(t)
					πF.SetLineno(1403)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					πTemp003[0] = µt
					if πE = πg.CheckLocal(πF, µP, "P"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µP, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 1404: C = ConsumerThread(Q, NI*NP)
					πF.SetLineno(1404)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µQ, "Q"); πE != nil {
						continue
					}
					πTemp003[0] = µQ
					if πE = πg.CheckLocal(πF, µNI, "NI"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µNP, "NP"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, µNI, µNP); πE != nil {
						continue
					}
					πTemp003[1] = πTemp002
					if πE = πg.CheckLocal(πF, µConsumerThread, "ConsumerThread"); πE != nil {
						continue
					}
					if πTemp002, πE = µConsumerThread.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µC = πTemp002
					if πE = πg.CheckLocal(πF, µP, "P"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µP); πE != nil {
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
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µt = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(4)            
					// line 1406: t.start()
					πF.SetLineno(1406)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßstart, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					// line 1407: _sleep(0.000001)
					πF.SetLineno(1407)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewFloat(1e-06).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_sleep); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 1408: C.start()
					πF.SetLineno(1408)
					if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µC, ßstart, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µP, "P"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µP); πE != nil {
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
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µt = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(7)            
					// line 1410: t.join()
					πF.SetLineno(1410)
					if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µt, ßjoin, nil); πE != nil {
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
					// line 1411: C.join()
					πF.SetLineno(1411)
					if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µC, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_test.ToObject(), πTemp024); πE != nil {
				continue
			}
			if πTemp026, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp025, πE = πg.Eq(πF, πTemp026, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp025); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label7
			}
			goto Label8
			// line 1413: if __name__ == '__main__':
			πF.SetLineno(1413)
		Label7:
			// line 1414: _test()
			πF.SetLineno(1414)
			if πTemp025, πE = πg.ResolveGlobal(πF, ß_test); πE != nil {
				continue
			}
			if πTemp026, πE = πTemp025.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label8
		Label8:
		}
		return nil, πE
	})
	πg.RegisterModule("threading", Code)
}

