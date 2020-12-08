package mutex
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/collections"
	// _ "github.com/pygolin/stdlib/warnings"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/mutex.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßFalse := πg.InternStr("False")
		ßTrue := πg.InternStr("True")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßappend := πg.InternStr("append")
		ßdeque := πg.InternStr("deque")
		ßlock := πg.InternStr("lock")
		ßlocked := πg.InternStr("locked")
		ßmutex := πg.InternStr("mutex")
		ßobject := πg.InternStr("object")
		ßpopleft := πg.InternStr("popleft")
		ßqueue := πg.InternStr("queue")
		ßtest := πg.InternStr("test")
		ßtestandset := πg.InternStr("testandset")
		ßunlock := πg.InternStr("unlock")
		ßwarnpy3k := πg.InternStr("warnpy3k")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 πg.KWArgs
		_ = πTemp004
		var πTemp005 *πg.Dict
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Mutual exclusion -- for use with module sched
			πF.SetLineno(1)
			// line 1: """Mutual exclusion -- for use with module sched
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Mutual exclusion -- for use with module sched\n\nA mutex has two pieces of state -- a 'locked' bit and a queue.\nWhen the mutex is not locked, the queue is empty.\nOtherwise, the queue contains 0 or more (function, argument) pairs\nrepresenting functions (or methods) waiting to acquire the lock.\nWhen the mutex is unlocked while the queue is not empty,\nthe first queue entry is removed and its function(argument) pair called,\nimplying it now has the lock.\n\nOf course, no multi-threading is implied -- hence the funny interface\nfor lock, where a function is called once the lock is acquired.\n").ToObject()); πE != nil {
				continue
			}
			// line 14: from warnings import warnpy3k
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßwarnpy3k); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßwarnpy3k.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 15: warnpy3k("the mutex module has been removed in Python 3.0", stacklevel=2)
			πF.SetLineno(15)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("the mutex module has been removed in Python 3.0").ToObject()
			πTemp004 = πg.KWArgs{
				{"stacklevel", πg.NewInt(2).ToObject()},
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnpy3k); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, πTemp004); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 16: del warnpy3k
			πF.SetLineno(16)
			if πE = πg.DelVar(πF, πF.Globals(), ßwarnpy3k); πE != nil {
				continue
			}
			// line 18: from collections import deque
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "collections"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdeque); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdeque.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: class mutex(object):
			πF.SetLineno(20)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp005 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("mutex", "/usr/lib/python2.7/mutex.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 21: def __init__(self):
					πF.SetLineno(21)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/mutex.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 22: """Create a new mutex -- initially unlocked."""
							πF.SetLineno(22)
							// line 23: self.locked = False
							πF.SetLineno(23)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlocked, πTemp002); πE != nil {
								continue
							}
							// line 24: self.queue = deque()
							πF.SetLineno(24)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdeque); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßqueue, πTemp001); πE != nil {
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
					// line 22: """Create a new mutex -- initially unlocked."""
					πF.SetLineno(22)
					// line 22: """Create a new mutex -- initially unlocked."""
					πF.SetLineno(22)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Create a new mutex -- initially unlocked.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Create a new mutex -- initially unlocked.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 26: def test(self):
					πF.SetLineno(26)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test", "/usr/lib/python2.7/mutex.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 27: """Test the locked bit of the mutex."""
							πF.SetLineno(27)
							// line 28: return self.locked
							πF.SetLineno(28)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlocked, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 27: """Test the locked bit of the mutex."""
					πF.SetLineno(27)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Test the locked bit of the mutex.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßtest); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 30: def testandset(self):
					πF.SetLineno(30)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("testandset", "/usr/lib/python2.7/mutex.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 31: """Atomic test-and-set -- grab the lock if it is not set,
							πF.SetLineno(31)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlocked, nil); πE != nil {
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
							// line 33: if not self.locked:
							πF.SetLineno(33)
						Label1:
							// line 34: self.locked = True
							πF.SetLineno(34)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlocked, πTemp002); πE != nil {
								continue
							}
							// line 35: return True
							πF.SetLineno(35)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label2:
							// line 37: return False
							πF.SetLineno(37)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
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
					if πE = πClass.SetItem(πF, ßtestandset.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 31: """Atomic test-and-set -- grab the lock if it is not set,
					πF.SetLineno(31)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Atomic test-and-set -- grab the lock if it is not set,\n        return True if it succeeded.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßtestandset); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 39: def lock(self, function, argument):
					πF.SetLineno(39)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "function", Def: nil}
					πTemp002[2] = πg.Param{Name: "argument", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("lock", "/usr/lib/python2.7/mutex.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfunction *πg.Object = πArgs[1]; _ = µfunction
						var µargument *πg.Object = πArgs[2]; _ = µargument
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
							// line 40: """Lock a mutex, call the function with supplied argument
							πF.SetLineno(40)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtestandset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 43: if self.testandset():
							πF.SetLineno(43)
						Label1:
							// line 44: function(argument)
							πF.SetLineno(44)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πTemp004[0] = µargument
							if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
								continue
							}
							if πTemp001, πE = µfunction.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label3
						Label2:
							// line 46: self.queue.append((function, argument))
							πF.SetLineno(46)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µfunction, µargument).ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßlock.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 40: """Lock a mutex, call the function with supplied argument
					πF.SetLineno(40)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Lock a mutex, call the function with supplied argument\n        when it is acquired.  If the mutex is already locked, place\n        function and argument in the queue.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßlock); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 48: def unlock(self):
					πF.SetLineno(48)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("unlock", "/usr/lib/python2.7/mutex.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfunction *πg.Object = πg.UnboundLocal; _ = µfunction
						var µargument *πg.Object = πg.UnboundLocal; _ = µargument
						var πTemp001 *πg.Object
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
							default: panic("unexpected function state")
							}
							// line 49: """Unlock a mutex.  If the queue is not empty, call the next
							πF.SetLineno(49)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 51: if self.queue:
							πF.SetLineno(51)
						Label1:
							// line 52: function, argument = self.queue.popleft()
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßqueue, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpopleft, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µfunction = πTemp003
							µargument = πTemp004
							// line 53: function(argument)
							πF.SetLineno(53)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargument, "argument"); πE != nil {
								continue
							}
							πTemp005[0] = µargument
							if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
								continue
							}
							if πTemp001, πE = µfunction.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label3
						Label2:
							// line 55: self.locked = False
							πF.SetLineno(55)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlocked, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßunlock.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 49: """Unlock a mutex.  If the queue is not empty, call the next
					πF.SetLineno(49)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Unlock a mutex.  If the queue is not empty, call the next\n        function with its argument.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßunlock); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
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
			if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("mutex").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmutex.ToObject(), πTemp006); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("mutex", Code)
}

