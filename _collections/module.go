package _collections
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/thread"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/_collections.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBLOCKSIZ := πg.InternStr("BLOCKSIZ")
		ßIndexError := πg.InternStr("IndexError")
		ßKeyError := πg.InternStr("KeyError")
		ßLFTLNK := πg.InternStr("LFTLNK")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßRGTLNK := πg.InternStr("RGTLNK")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__class__ := πg.InternStr("__class__")
		ß__copy__ := πg.InternStr("__copy__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__ge__ := πg.InternStr("__ge__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__getref := πg.InternStr("__getref")
		ß__gt__ := πg.InternStr("__gt__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__iadd__ := πg.InternStr("__iadd__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__le__ := πg.InternStr("__le__")
		ß__len__ := πg.InternStr("__len__")
		ß__lt__ := πg.InternStr("__lt__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__missing__ := πg.InternStr("__missing__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__new__ := πg.InternStr("__new__")
		ß__reduce__ := πg.InternStr("__reduce__")
		ß__reduce_ex__ := πg.InternStr("__reduce_ex__")
		ß__repr := πg.InternStr("__repr")
		ß__repr__ := πg.InternStr("__repr__")
		ß__reversed__ := πg.InternStr("__reversed__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß_gen := πg.InternStr("_gen")
		ß_iter_impl := πg.InternStr("_iter_impl")
		ß_maxlen := πg.InternStr("_maxlen")
		ß_reversed_impl := πg.InternStr("_reversed_impl")
		ß_thread_ident := πg.InternStr("_thread_ident")
		ßadd := πg.InternStr("add")
		ßappend := πg.InternStr("append")
		ßappendleft := πg.InternStr("appendleft")
		ßcallable := πg.InternStr("callable")
		ßclear := πg.InternStr("clear")
		ßcopy := πg.InternStr("copy")
		ßcount := πg.InternStr("count")
		ßcounter := πg.InternStr("counter")
		ßdefault_factory := πg.InternStr("default_factory")
		ßdefaultdict := πg.InternStr("defaultdict")
		ßdeque := πg.InternStr("deque")
		ßdeque_iterator := πg.InternStr("deque_iterator")
		ßdict := πg.InternStr("dict")
		ßextend := πg.InternStr("extend")
		ßextendleft := πg.InternStr("extendleft")
		ßget_ident := πg.InternStr("get_ident")
		ßid := πg.InternStr("id")
		ßisinstance := πg.InternStr("isinstance")
		ßiteritems := πg.InternStr("iteritems")
		ßleft := πg.InternStr("left")
		ßleftndx := πg.InternStr("leftndx")
		ßlen := πg.InternStr("len")
		ßlength := πg.InternStr("length")
		ßlist := πg.InternStr("list")
		ßmaxlen := πg.InternStr("maxlen")
		ßn := πg.InternStr("n")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßpop := πg.InternStr("pop")
		ßpopleft := πg.InternStr("popleft")
		ßproperty := πg.InternStr("property")
		ßrange := πg.InternStr("range")
		ßremove := πg.InternStr("remove")
		ßrepr := πg.InternStr("repr")
		ßreverse := πg.InternStr("reverse")
		ßreversed := πg.InternStr("reversed")
		ßright := πg.InternStr("right")
		ßrightndx := πg.InternStr("rightndx")
		ßrotate := πg.InternStr("rotate")
		ßset := πg.InternStr("set")
		ßstate := πg.InternStr("state")
		ßstr := πg.InternStr("str")
		ßsuper := πg.InternStr("super")
		ßthread := πg.InternStr("thread")
		ßtype := πg.InternStr("type")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """High performance data structures
			πF.SetLineno(1)
			// line 1: """High performance data structures
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("High performance data structures\n").ToObject()); πE != nil {
				continue
			}
			// line 17: import thread
			πF.SetLineno(17)
			if πTemp002, πE = πg.ImportModule(πF, "thread"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthread.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: _thread_ident = thread.get_ident
			πF.SetLineno(18)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßthread); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget_ident, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_thread_ident.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 20: n = 30
			πF.SetLineno(20)
			if πE = πF.Globals().SetItem(πF, ßn.ToObject(), πg.NewInt(30).ToObject()); πE != nil {
				continue
			}
			// line 21: LFTLNK = n
			πF.SetLineno(21)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßLFTLNK.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: RGTLNK = n+1
			πF.SetLineno(22)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßRGTLNK.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: BLOCKSIZ = n+2
			πF.SetLineno(23)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBLOCKSIZ.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 31: class deque(object):
			πF.SetLineno(31)
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
			_, πE = πg.NewCode("deque", "/usr/lib/python2.7/_collections.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp026 *πg.Object
				_ = πTemp026
				var πTemp027 *πg.Object
				_ = πTemp027
				var πTemp028 *πg.Object
				_ = πTemp028
				var πTemp029 *πg.Object
				_ = πTemp029
				var πTemp030 *πg.Object
				_ = πTemp030
				var πTemp031 *πg.Object
				_ = πTemp031
				var πTemp032 *πg.Object
				_ = πTemp032
				var πTemp033 *πg.Object
				_ = πTemp033
				var πTemp034 *πg.Object
				_ = πTemp034
				var πTemp035 *πg.Object
				_ = πTemp035
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 33: def __new__(cls, iterable=(), *args, **kw):
					πF.SetLineno(33)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp003 = πg.NewTuple0().ToObject()
					πTemp002[1] = πg.Param{Name: "iterable", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/_collections.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var µkw *πg.Object = πArgs[3]; _ = µkw
						var µself *πg.Object = πg.UnboundLocal; _ = µself
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
							// line 34: self = super(deque, cls).__new__(cls)
							πF.SetLineno(34)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdeque); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp002[1] = µcls
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µself = πTemp004
							// line 35: self.clear()
							πF.SetLineno(35)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßclear, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 36: return self
							πF.SetLineno(36)
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
					// line 38: def __init__(self, iterable=(), maxlen=None):
					πF.SetLineno(38)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewTuple0().ToObject()
					πTemp002[1] = πg.Param{Name: "iterable", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "maxlen", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µmaxlen *πg.Object = πArgs[2]; _ = µmaxlen
						var µadd *πg.Object = πg.UnboundLocal; _ = µadd
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
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
							// line 39: self.clear()
							πF.SetLineno(39)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßclear, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmaxlen, "maxlen"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmaxlen != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 40: if maxlen is not None:
							πF.SetLineno(40)
						Label1:
							if πE = πg.CheckLocal(πF, µmaxlen, "maxlen"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µmaxlen, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 41: if maxlen < 0:
							πF.SetLineno(41)
						Label3:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("maxlen must be non-negative").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 42: raise ValueError("maxlen must be non-negative")
							πF.SetLineno(42)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 43: self._maxlen = maxlen
							πF.SetLineno(43)
							if πE = πg.CheckLocal(πF, µmaxlen, "maxlen"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmaxlen); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_maxlen, πTemp001); πE != nil {
								continue
							}
							// line 44: add = self.append
							πF.SetLineno(44)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							µadd = πTemp001
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
								continue
							}
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
								µelem = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(5)            
							// line 46: add(elem)
							πF.SetLineno(46)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp004[0] = µelem
							if πE = πg.CheckLocal(πF, µadd, "add"); πE != nil {
								continue
							}
							if πTemp002, πE = µadd.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 49: def maxlen(self):
					πF.SetLineno(49)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("maxlen", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 50: return self._maxlen
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_maxlen, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmaxlen.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 52: maxlen = property(maxlen)
					πF.SetLineno(52)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßmaxlen); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmaxlen.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 53: def clear(self):
					πF.SetLineno(53)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("clear", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 54: self.right = self.left = [None] * BLOCKSIZ
							πF.SetLineno(54)
							πTemp002 = make([]*πg.Object, 1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßBLOCKSIZ); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßright, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßleft, πTemp003); πE != nil {
								continue
							}
							// line 55: self.rightndx = n//2   # points to last written element
							πF.SetLineno(55)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp001, πE = πg.FloorDiv(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrightndx, πTemp003); πE != nil {
								continue
							}
							// line 56: self.leftndx = n//2+1
							πF.SetLineno(56)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp003, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßleftndx, πTemp003); πE != nil {
								continue
							}
							// line 57: self.length = 0
							πF.SetLineno(57)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlength, πTemp001); πE != nil {
								continue
							}
							// line 58: self.state = 0
							πF.SetLineno(58)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp001); πE != nil {
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
					// line 60: def append(self, x):
					πF.SetLineno(60)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("append", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µnewblock *πg.Object = πg.UnboundLocal; _ = µnewblock
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 61: self.state += 1
							πF.SetLineno(61)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							// line 62: self.rightndx += 1
							πF.SetLineno(62)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrightndx, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 63: if self.rightndx == n:
							πF.SetLineno(63)
						Label1:
							// line 64: newblock = [None] * BLOCKSIZ
							πF.SetLineno(64)
							πTemp005 = make([]*πg.Object, 1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBLOCKSIZ); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							µnewblock = πTemp001
							// line 65: self.right[RGTLNK] = newblock
							πF.SetLineno(65)
							if πE = πg.CheckLocal(πF, µnewblock, "newblock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnewblock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßRGTLNK); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 66: newblock[LFTLNK] = self.right
							πF.SetLineno(66)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewblock, "newblock"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßLFTLNK); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.SetItem(πF, µnewblock, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 67: self.right = newblock
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µnewblock, "newblock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnewblock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßright, πTemp001); πE != nil {
								continue
							}
							// line 68: self.rightndx = 0
							πF.SetLineno(68)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrightndx, πTemp001); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 69: self.length += 1
							πF.SetLineno(69)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlength, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlength, πTemp002); πE != nil {
								continue
							}
							// line 70: self.right[self.rightndx] = x
							πF.SetLineno(70)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µx); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxlen, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 != πTemp006).ToObject()
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlength, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmaxlen, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 71: if self.maxlen is not None and self.length > self.maxlen:
							πF.SetLineno(71)
						Label4:
							// line 72: self.popleft()
							πF.SetLineno(72)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpopleft, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßappend.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 74: def appendleft(self, x):
					πF.SetLineno(74)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("appendleft", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var µnewblock *πg.Object = πg.UnboundLocal; _ = µnewblock
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 75: self.state += 1
							πF.SetLineno(75)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							// line 76: self.leftndx -= 1
							πF.SetLineno(76)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßleftndx, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 77: if self.leftndx == -1:
							πF.SetLineno(77)
						Label1:
							// line 78: newblock = [None] * BLOCKSIZ
							πF.SetLineno(78)
							πTemp005 = make([]*πg.Object, 1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBLOCKSIZ); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							µnewblock = πTemp001
							// line 79: self.left[LFTLNK] = newblock
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µnewblock, "newblock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnewblock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßLFTLNK); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							// line 80: newblock[RGTLNK] = self.left
							πF.SetLineno(80)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewblock, "newblock"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßRGTLNK); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.SetItem(πF, µnewblock, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 81: self.left = newblock
							πF.SetLineno(81)
							if πE = πg.CheckLocal(πF, µnewblock, "newblock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnewblock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßleft, πTemp001); πE != nil {
								continue
							}
							// line 82: self.leftndx = n-1
							πF.SetLineno(82)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßleftndx, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 83: self.length += 1
							πF.SetLineno(83)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlength, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlength, πTemp002); πE != nil {
								continue
							}
							// line 84: self.left[self.leftndx] = x
							πF.SetLineno(84)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µx); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxlen, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 != πTemp006).ToObject()
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlength, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmaxlen, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 85: if self.maxlen is not None and self.length > self.maxlen:
							πF.SetLineno(85)
						Label4:
							// line 86: self.pop()
							πF.SetLineno(86)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßappendleft.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 88: def extend(self, iterable):
					πF.SetLineno(88)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("extend", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µiterable == µself).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 89: if iterable is self:
							πF.SetLineno(89)
						Label1:
							// line 90: iterable = list(iterable)
							πF.SetLineno(90)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp003[0] = µiterable
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µiterable = πTemp004
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp002 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µelem = πTemp004
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 92: self.append(elem)
							πF.SetLineno(92)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp003[0] = µelem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßextend.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 94: def extendleft(self, iterable):
					πF.SetLineno(94)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("extendleft", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µiterable == µself).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 95: if iterable is self:
							πF.SetLineno(95)
						Label1:
							// line 96: iterable = list(iterable)
							πF.SetLineno(96)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp003[0] = µiterable
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µiterable = πTemp004
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp002 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µelem = πTemp004
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 98: self.appendleft(elem)
							πF.SetLineno(98)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							πTemp003[0] = µelem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßappendleft, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßextendleft.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 100: def pop(self):
					πF.SetLineno(100)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("pop", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µprevblock *πg.Object = πg.UnboundLocal; _ = µprevblock
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
						var πTemp006 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 == πTemp005).ToObject()
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
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 101: if self.left is self.right and self.leftndx > self.rightndx:
							πF.SetLineno(101)
						Label2:
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("pop from an empty deque").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 102: raise IndexError("pop from an empty deque")
							πF.SetLineno(102)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 103: x = self.right[self.rightndx]
							πF.SetLineno(103)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µx = πTemp003
							// line 104: self.right[self.rightndx] = None
							πF.SetLineno(104)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp007
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							// line 105: self.length -= 1
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlength, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlength, πTemp003); πE != nil {
								continue
							}
							// line 106: self.rightndx -= 1
							πF.SetLineno(106)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrightndx, πTemp003); πE != nil {
								continue
							}
							// line 107: self.state += 1
							πF.SetLineno(107)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 108: if self.rightndx == -1:
							πF.SetLineno(108)
						Label4:
							// line 109: prevblock = self.right[LFTLNK]
							πF.SetLineno(109)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßLFTLNK); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µprevblock = πTemp003
							if πE = πg.CheckLocal(πF, µprevblock, "prevblock"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µprevblock == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 110: if prevblock is None:
							πF.SetLineno(110)
						Label6:
							// line 112: self.rightndx = n//2
							πF.SetLineno(112)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp001, πE = πg.FloorDiv(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrightndx, πTemp003); πE != nil {
								continue
							}
							// line 113: self.leftndx = n//2+1
							πF.SetLineno(113)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp003, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßleftndx, πTemp003); πE != nil {
								continue
							}
							goto Label8
						Label7:
							// line 115: prevblock[RGTLNK] = None
							πF.SetLineno(115)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µprevblock, "prevblock"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßRGTLNK); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, µprevblock, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 116: self.right[LFTLNK] = None
							πF.SetLineno(116)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßLFTLNK); πE != nil {
								continue
							}
							πTemp005 = πTemp007
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							// line 117: self.right = prevblock
							πF.SetLineno(117)
							if πE = πg.CheckLocal(πF, µprevblock, "prevblock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µprevblock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßright, πTemp001); πE != nil {
								continue
							}
							// line 118: self.rightndx = n-1
							πF.SetLineno(118)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrightndx, πTemp003); πE != nil {
								continue
							}
							goto Label8
						Label8:
							goto Label5
						Label5:
							// line 119: return x
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πR = µx
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 121: def popleft(self):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("popleft", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µprevblock *πg.Object = πg.UnboundLocal; _ = µprevblock
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
						var πTemp006 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 == πTemp005).ToObject()
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
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 122: if self.left is self.right and self.leftndx > self.rightndx:
							πF.SetLineno(122)
						Label2:
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("pop from an empty deque").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 123: raise IndexError("pop from an empty deque")
							πF.SetLineno(123)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 124: x = self.left[self.leftndx]
							πF.SetLineno(124)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µx = πTemp003
							// line 125: self.left[self.leftndx] = None
							πF.SetLineno(125)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp007
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							// line 126: self.length -= 1
							πF.SetLineno(126)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlength, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlength, πTemp003); πE != nil {
								continue
							}
							// line 127: self.leftndx += 1
							πF.SetLineno(127)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßleftndx, πTemp003); πE != nil {
								continue
							}
							// line 128: self.state += 1
							πF.SetLineno(128)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 129: if self.leftndx == n:
							πF.SetLineno(129)
						Label4:
							// line 130: prevblock = self.left[RGTLNK]
							πF.SetLineno(130)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßRGTLNK); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µprevblock = πTemp003
							if πE = πg.CheckLocal(πF, µprevblock, "prevblock"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µprevblock == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 131: if prevblock is None:
							πF.SetLineno(131)
						Label6:
							// line 133: self.rightndx = n//2
							πF.SetLineno(133)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp001, πE = πg.FloorDiv(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrightndx, πTemp003); πE != nil {
								continue
							}
							// line 134: self.leftndx = n//2+1
							πF.SetLineno(134)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp003, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßleftndx, πTemp003); πE != nil {
								continue
							}
							goto Label8
						Label7:
							// line 136: prevblock[LFTLNK] = None
							πF.SetLineno(136)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µprevblock, "prevblock"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßLFTLNK); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, µprevblock, πTemp004, πTemp003); πE != nil {
								continue
							}
							// line 137: self.left[RGTLNK] = None
							πF.SetLineno(137)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßRGTLNK); πE != nil {
								continue
							}
							πTemp005 = πTemp007
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							// line 138: self.left = prevblock
							πF.SetLineno(138)
							if πE = πg.CheckLocal(πF, µprevblock, "prevblock"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µprevblock); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßleft, πTemp001); πE != nil {
								continue
							}
							// line 139: self.leftndx = 0
							πF.SetLineno(139)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßleftndx, πTemp001); πE != nil {
								continue
							}
							goto Label8
						Label8:
							goto Label5
						Label5:
							// line 140: return x
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πR = µx
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpopleft.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 142: def count(self, value):
					πF.SetLineno(142)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("count", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var µc *πg.Object = πg.UnboundLocal; _ = µc
						var µitem *πg.Object = πg.UnboundLocal; _ = µitem
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
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 143: c = 0
							πF.SetLineno(143)
							µc = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
								µitem = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µitem, µvalue); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 145: if item == value:
							πF.SetLineno(145)
						Label4:
							// line 146: c += 1
							πF.SetLineno(146)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IAdd(πF, µc, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µc = πTemp004
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 147: return c
							πF.SetLineno(147)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πR = µc
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcount.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 149: def remove(self, value):
					πF.SetLineno(149)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("remove", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
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
							// line 151: i = 0
							πF.SetLineno(151)
							µi = πg.NewInt(0).ToObject()
							// line 152: try:
							πF.SetLineno(152)
							πF.PushCheckpoint(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
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
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
							πF.PushCheckpoint(2)            
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µself, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp008, µvalue); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							goto Label6
							// line 154: if self[0] == value:
							πF.SetLineno(154)
						Label5:
							// line 155: self.popleft()
							πF.SetLineno(155)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpopleft, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 156: return
							πF.SetLineno(156)
							πR = πg.None
							continue
							goto Label6
						Label6:
							// line 157: self.append(self.popleft())
							πF.SetLineno(157)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpopleft, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label3:
							if πE != nil || πR != nil {
								continue
							}
						Label4:
							// line 158: i += 1
							πF.SetLineno(158)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("deque.remove(x): x not in deque").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 159: raise ValueError("deque.remove(x): x not in deque")
							πF.SetLineno(159)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							πF.PopCheckpoint()
							goto Label1
						Label1:
							πTemp009, πTemp010 = πF.RestoreExc(nil, nil)
							// line 161: self.rotate(i)
							πF.SetLineno(161)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002[0] = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrotate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßremove.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 163: def rotate(self, n=1):
					πF.SetLineno(163)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: πg.NewInt(1).ToObject()}
					πTemp015 = πg.NewFunction(πg.NewCode("rotate", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πArgs[1]; _ = µn
						var µlength *πg.Object = πg.UnboundLocal; _ = µlength
						var µhalflen *πg.Object = πg.UnboundLocal; _ = µhalflen
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
						var πTemp006 bool
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 9: goto Label9
							case 10: goto Label10
							case 12: goto Label12
							case 13: goto Label13
							default: panic("unexpected function state")
							}
							// line 164: length = len(self)
							πF.SetLineno(164)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlength = πTemp003
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, µlength, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 165: if length <= 1:
							πF.SetLineno(165)
						Label1:
							// line 166: return
							πF.SetLineno(166)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 167: halflen = length >> 1
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.RShift(πF, µlength, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µhalflen = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhalflen, "halflen"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µn, µhalflen); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhalflen, "halflen"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, µhalflen); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µn, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 168: if n > halflen or n < -halflen:
							πF.SetLineno(168)
						Label4:
							// line 169: n %= length
							πF.SetLineno(169)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IMod(πF, µn, µlength); πE != nil {
								continue
							}
							µn = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhalflen, "halflen"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µn, µhalflen); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhalflen, "halflen"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, µhalflen); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µn, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 170: if n > halflen:
							πF.SetLineno(170)
						Label6:
							// line 171: n -= length
							πF.SetLineno(171)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, µn, µlength); πE != nil {
								continue
							}
							µn = πTemp002
							goto Label8
							// line 172: elif n < -halflen:
							πF.SetLineno(172)
						Label7:
							// line 173: n += length
							πF.SetLineno(173)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µn, µlength); πE != nil {
								continue
							}
							µn = πTemp002
							goto Label8
						Label8:
							goto Label5
						Label5:
							// line 174: while n > 0:
							πF.SetLineno(174)
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
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(9)            
							// line 175: self.appendleft(self.pop())
							πF.SetLineno(175)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappendleft, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 176: n -= 1
							πF.SetLineno(176)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µn = πTemp002
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							// line 177: while n < 0:
							πF.SetLineno(177)
							πF.PushCheckpoint(13)
							πTemp004 = false
						Label12:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(12)            
							// line 178: self.append(self.popleft())
							πF.SetLineno(178)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpopleft, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 179: n += 1
							πF.SetLineno(179)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µn = πTemp002
							continue
						Label13:
							if πE != nil || πR != nil {
								continue
							}
						Label14:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßrotate.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 181: def reverse(self):
					πF.SetLineno(181)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("reverse", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µleftblock *πg.Object = πg.UnboundLocal; _ = µleftblock
						var µrightblock *πg.Object = πg.UnboundLocal; _ = µrightblock
						var µleftindex *πg.Object = πg.UnboundLocal; _ = µleftindex
						var µrightindex *πg.Object = πg.UnboundLocal; _ = µrightindex
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
						var πTemp007 *πg.Object
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
							default: panic("unexpected function state")
							}
							// line 182: "reverse *IN PLACE*"
							πF.SetLineno(182)
							// line 183: leftblock = self.left
							πF.SetLineno(183)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							µleftblock = πTemp001
							// line 184: rightblock = self.right
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							µrightblock = πTemp001
							// line 185: leftindex = self.leftndx
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							µleftindex = πTemp001
							// line 186: rightindex = self.rightndx
							πF.SetLineno(186)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							µrightindex = πTemp001
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßlength, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.FloorDiv(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
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
								µi = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 189: assert leftblock != rightblock or leftindex < rightindex
							πF.SetLineno(189)
							if πE = πg.CheckLocal(πF, µleftblock, "leftblock"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrightblock, "rightblock"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.NE(πF, µleftblock, µrightblock); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µleftindex, "leftindex"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrightindex, "rightindex"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µleftindex, µrightindex); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label4:
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 192: (rightblock[rightindex], leftblock[leftindex]) = (
							πF.SetLineno(192)
							if πE = πg.CheckLocal(πF, µleftindex, "leftindex"); πE != nil {
								continue
							}
							πTemp004 = µleftindex
							if πE = πg.CheckLocal(πF, µleftblock, "leftblock"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µleftblock, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrightindex, "rightindex"); πE != nil {
								continue
							}
							πTemp004 = µrightindex
							if πE = πg.CheckLocal(πF, µrightblock, "rightblock"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µrightblock, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrightblock, "rightblock"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrightindex, "rightindex"); πE != nil {
								continue
							}
							πTemp008 = µrightindex
							if πE = πg.SetItem(πF, µrightblock, πTemp008, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleftblock, "leftblock"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µleftindex, "leftindex"); πE != nil {
								continue
							}
							πTemp004 = µleftindex
							if πE = πg.SetItem(πF, µleftblock, πTemp004, πTemp007); πE != nil {
								continue
							}
							// line 196: leftindex += 1
							πF.SetLineno(196)
							if πE = πg.CheckLocal(πF, µleftindex, "leftindex"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µleftindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µleftindex = πTemp003
							if πE = πg.CheckLocal(πF, µleftindex, "leftindex"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µleftindex, πTemp004); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 197: if leftindex == n:
							πF.SetLineno(197)
						Label5:
							// line 198: leftblock = leftblock[RGTLNK]
							πF.SetLineno(198)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßRGTLNK); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µleftblock, "leftblock"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µleftblock, πTemp003); πE != nil {
								continue
							}
							µleftblock = πTemp004
							// line 199: assert leftblock is not None
							πF.SetLineno(199)
							if πE = πg.CheckLocal(πF, µleftblock, "leftblock"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µleftblock != πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 200: leftindex = 0
							πF.SetLineno(200)
							µleftindex = πg.NewInt(0).ToObject()
							goto Label6
						Label6:
							// line 203: rightindex -= 1
							πF.SetLineno(203)
							if πE = πg.CheckLocal(πF, µrightindex, "rightindex"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ISub(πF, µrightindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µrightindex = πTemp003
							if πE = πg.CheckLocal(πF, µrightindex, "rightindex"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µrightindex, πTemp004); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 204: if rightindex == -1:
							πF.SetLineno(204)
						Label7:
							// line 205: rightblock = rightblock[LFTLNK]
							πF.SetLineno(205)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßLFTLNK); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µrightblock, "rightblock"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µrightblock, πTemp003); πE != nil {
								continue
							}
							µrightblock = πTemp004
							// line 206: assert rightblock is not None
							πF.SetLineno(206)
							if πE = πg.CheckLocal(πF, µrightblock, "rightblock"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µrightblock != πTemp004).ToObject()
							if πE = πg.Assert(πF, πTemp003, nil); πE != nil {
								continue
							}
							// line 207: rightindex = n - 1
							πF.SetLineno(207)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µrightindex = πTemp003
							goto Label8
						Label8:
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
					if πE = πClass.SetItem(πF, ßreverse.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 182: "reverse *IN PLACE*"
					πF.SetLineno(182)
					// line 182: "reverse *IN PLACE*"
					πF.SetLineno(182)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("reverse *IN PLACE*").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("reverse *IN PLACE*").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßreverse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
						continue
					}
					// line 209: def __repr__(self):
					πF.SetLineno(209)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µthreadlocalattr *πg.Object = πg.UnboundLocal; _ = µthreadlocalattr
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
						var πTemp006 *πg.Object
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
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 210: threadlocalattr = '__repr' + str(_thread_ident())
							πF.SetLineno(210)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_thread_ident); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Add(πF, ß__repr.ToObject(), πTemp004); πE != nil {
								continue
							}
							µthreadlocalattr = πTemp001
							if πE = πg.CheckLocal(πF, µthreadlocalattr, "threadlocalattr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, µthreadlocalattr); πE != nil {
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
							// line 211: if threadlocalattr in self.__dict__:
							πF.SetLineno(211)
						Label1:
							// line 212: return 'deque([...])'
							πF.SetLineno(212)
							πR = πg.NewStr("deque([...])").ToObject()
							continue
							goto Label3
						Label2:
							// line 214: self.__dict__[threadlocalattr] = True
							πF.SetLineno(214)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µthreadlocalattr, "threadlocalattr"); πE != nil {
								continue
							}
							πTemp006 = µthreadlocalattr
							if πE = πg.SetItem(πF, πTemp004, πTemp006, πTemp003); πE != nil {
								continue
							}
							// line 215: try:
							πF.SetLineno(215)
							πF.PushCheckpoint(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmaxlen, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 != πTemp004).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 216: if self.maxlen is not None:
							πF.SetLineno(216)
						Label5:
							// line 217: return 'deque(%r, maxlen=%s)' % (list(self), self.maxlen)
							πF.SetLineno(217)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmaxlen, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp006, πTemp004).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("deque(%r, maxlen=%s)").ToObject(), πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label7
						Label6:
							// line 219: return 'deque(%r)' % (list(self),)
							πF.SetLineno(219)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πg.NewTuple1(πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("deque(%r)").ToObject(), πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label7
						Label7:
							πF.PopCheckpoint()
							goto Label4
						Label4:
							πTemp007, πTemp008 = πF.RestoreExc(nil, nil)
							// line 221: del self.__dict__[threadlocalattr]
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µthreadlocalattr, "threadlocalattr"); πE != nil {
								continue
							}
							πTemp003 = µthreadlocalattr
							if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							if πTemp007 != nil {
								πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
								continue
							}
							if πR != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 223: def __iter__(self):
					πF.SetLineno(223)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 224: return deque_iterator(self, self._iter_impl)
							πF.SetLineno(224)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_iter_impl, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdeque_iterator); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 226: def _iter_impl(self, original_state, giveup):
					πF.SetLineno(226)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "original_state", Def: nil}
					πTemp002[2] = πg.Param{Name: "giveup", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("_iter_impl", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoriginal_state *πg.Object = πArgs[1]; _ = µoriginal_state
						var µgiveup *πg.Object = πArgs[2]; _ = µgiveup
						var µblock *πg.Object = πg.UnboundLocal; _ = µblock
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
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
						var πTemp006 bool
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 11: goto Label11
								case 10: goto Label10
								case 3: goto Label3
								case 4: goto Label4
								case 13: goto Label13
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µoriginal_state, "original_state"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.NE(πF, πTemp002, µoriginal_state); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label1
								}
								goto Label2
								// line 227: if self.state != original_state:
								πF.SetLineno(227)
							Label1:
								// line 228: giveup()
								πF.SetLineno(228)
								if πE = πg.CheckLocal(πF, µgiveup, "giveup"); πE != nil {
									continue
								}
								if πTemp001, πE = µgiveup.Call(πF, nil, nil); πE != nil {
									continue
								}
								goto Label2
							Label2:
								// line 229: block = self.left
								πF.SetLineno(229)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
									continue
								}
								µblock = πTemp001
								// line 230: while block:
								πF.SetLineno(230)
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
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, µblock); πE != nil {
									continue
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(3)            
								// line 231: l, r = 0, n
								πF.SetLineno(231)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
									continue
								}
								πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp002).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
									continue
								}
								µl = πTemp002
								µr = πTemp005
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µblock == πTemp002).ToObject()
								if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label6
								}
								goto Label7
								// line 232: if block is self.left:
								πF.SetLineno(232)
							Label6:
								// line 233: l = self.leftndx
								πF.SetLineno(233)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
									continue
								}
								µl = πTemp001
								goto Label7
							Label7:
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µblock == πTemp002).ToObject()
								if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label8
								}
								goto Label9
								// line 234: if block is self.right:
								πF.SetLineno(234)
							Label8:
								// line 235: r = self.rightndx + 1
								πF.SetLineno(235)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µr = πTemp001
								goto Label9
							Label9:
								if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µl, µr, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µblock, πTemp002); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
									continue
								}
								πF.PushCheckpoint(11)
								πTemp004 = false
							Label10:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
									πF.PopCheckpoint()
									goto Label12
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
									µelem = πTemp002
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(10)            
								// line 237: yield elem
								πF.SetLineno(237)
								if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
									continue
								}
								πF.PushCheckpoint(13)
								return µelem, nil
							Label13:
								πTemp002 = πSent
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µoriginal_state, "original_state"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.NE(πF, πTemp005, µoriginal_state); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label14
								}
								goto Label15
								// line 238: if self.state != original_state:
								πF.SetLineno(238)
							Label14:
								// line 239: giveup()
								πF.SetLineno(239)
								if πE = πg.CheckLocal(πF, µgiveup, "giveup"); πE != nil {
									continue
								}
								if πTemp002, πE = µgiveup.Call(πF, nil, nil); πE != nil {
									continue
								}
								goto Label15
							Label15:
								continue
							Label11:
								if πE != nil || πR != nil {
									continue
								}
							Label12:
								// line 240: block = block[RGTLNK]
								πF.SetLineno(240)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßRGTLNK); πE != nil {
									continue
								}
								πTemp001 = πTemp002
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
									continue
								}
								µblock = πTemp002
								continue
							Label4:
								if πE != nil || πR != nil {
									continue
								}
							Label5:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_iter_impl.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 242: def __reversed__(self):
					πF.SetLineno(242)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("__reversed__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 243: return deque_iterator(self, self._reversed_impl)
							πF.SetLineno(243)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_reversed_impl, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdeque_iterator); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__reversed__.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 245: def _reversed_impl(self, original_state, giveup):
					πF.SetLineno(245)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "original_state", Def: nil}
					πTemp002[2] = πg.Param{Name: "giveup", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("_reversed_impl", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoriginal_state *πg.Object = πArgs[1]; _ = µoriginal_state
						var µgiveup *πg.Object = πArgs[2]; _ = µgiveup
						var µblock *πg.Object = πg.UnboundLocal; _ = µblock
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 11: goto Label11
								case 10: goto Label10
								case 3: goto Label3
								case 4: goto Label4
								case 13: goto Label13
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µoriginal_state, "original_state"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.NE(πF, πTemp002, µoriginal_state); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label1
								}
								goto Label2
								// line 246: if self.state != original_state:
								πF.SetLineno(246)
							Label1:
								// line 247: giveup()
								πF.SetLineno(247)
								if πE = πg.CheckLocal(πF, µgiveup, "giveup"); πE != nil {
									continue
								}
								if πTemp001, πE = µgiveup.Call(πF, nil, nil); πE != nil {
									continue
								}
								goto Label2
							Label2:
								// line 248: block = self.right
								πF.SetLineno(248)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
									continue
								}
								µblock = πTemp001
								// line 249: while block:
								πF.SetLineno(249)
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
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.IsTrue(πF, µblock); πE != nil {
									continue
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(3)            
								// line 250: l, r = 0, n
								πF.SetLineno(250)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
									continue
								}
								πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp002).ToObject()
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
									continue
								}
								µl = πTemp002
								µr = πTemp005
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µblock == πTemp002).ToObject()
								if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label6
								}
								goto Label7
								// line 251: if block is self.left:
								πF.SetLineno(251)
							Label6:
								// line 252: l = self.leftndx
								πF.SetLineno(252)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
									continue
								}
								µl = πTemp001
								goto Label7
							Label7:
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µblock == πTemp002).ToObject()
								if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp004 {
									goto Label8
								}
								goto Label9
								// line 253: if block is self.right:
								πF.SetLineno(253)
							Label8:
								// line 254: r = self.rightndx + 1
								πF.SetLineno(254)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µr = πTemp001
								goto Label9
							Label9:
								πTemp006 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µl, µr, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µblock, πTemp002); πE != nil {
									continue
								}
								πTemp006[0] = πTemp005
								if πTemp002, πE = πg.ResolveGlobal(πF, ßreversed); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
									continue
								}
								πF.PushCheckpoint(11)
								πTemp004 = false
							Label10:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp004 {
									πF.PopCheckpoint()
									goto Label12
								}
								if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
									µelem = πTemp002
								}
								if πE != nil || !πTemp007 {
									continue
								}
								πF.PushCheckpoint(10)            
								// line 256: yield elem
								πF.SetLineno(256)
								if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
									continue
								}
								πF.PushCheckpoint(13)
								return µelem, nil
							Label13:
								πTemp002 = πSent
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µoriginal_state, "original_state"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.NE(πF, πTemp005, µoriginal_state); πE != nil {
									continue
								}
								if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp007 {
									goto Label14
								}
								goto Label15
								// line 257: if self.state != original_state:
								πF.SetLineno(257)
							Label14:
								// line 258: giveup()
								πF.SetLineno(258)
								if πE = πg.CheckLocal(πF, µgiveup, "giveup"); πE != nil {
									continue
								}
								if πTemp002, πE = µgiveup.Call(πF, nil, nil); πE != nil {
									continue
								}
								goto Label15
							Label15:
								continue
							Label11:
								if πE != nil || πR != nil {
									continue
								}
							Label12:
								// line 259: block = block[LFTLNK]
								πF.SetLineno(259)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßLFTLNK); πE != nil {
									continue
								}
								πTemp001 = πTemp002
								if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
									continue
								}
								µblock = πTemp002
								continue
							Label4:
								if πE != nil || πR != nil {
									continue
								}
							Label5:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_reversed_impl.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 261: def __len__(self):
					πF.SetLineno(261)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 268: return self.length
							πF.SetLineno(268)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlength, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 270: def __getref(self, index):
					πF.SetLineno(270)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("__getref", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var µblock *πg.Object = πg.UnboundLocal; _ = µblock
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µr *πg.Object = πg.UnboundLocal; _ = µr
						var µspan *πg.Object = πg.UnboundLocal; _ = µspan
						var µnegative_span *πg.Object = πg.UnboundLocal; _ = µnegative_span
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
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
							case 4: goto Label4
							case 5: goto Label5
							case 14: goto Label14
							case 13: goto Label13
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µindex, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 271: if index >= 0:
							πF.SetLineno(271)
						Label1:
							// line 272: block = self.left
							πF.SetLineno(272)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							µblock = πTemp001
							// line 273: while block:
							πF.SetLineno(273)
							πF.PushCheckpoint(5)
							πTemp002 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µblock); πE != nil {
								continue
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 274: l, r = 0, n
							πF.SetLineno(274)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µl = πTemp004
							µr = πTemp005
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µblock == πTemp004).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 275: if block is self.left:
							πF.SetLineno(275)
						Label7:
							// line 276: l = self.leftndx
							πF.SetLineno(276)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							µl = πTemp001
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µblock == πTemp004).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 277: if block is self.right:
							πF.SetLineno(277)
						Label9:
							// line 278: r = self.rightndx + 1
							πF.SetLineno(278)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µr = πTemp001
							goto Label10
						Label10:
							// line 279: span = r-l
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µr, µl); πE != nil {
								continue
							}
							µspan = πTemp001
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µspan, "span"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µindex, µspan); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							goto Label12
							// line 280: if index < span:
							πF.SetLineno(280)
						Label11:
							// line 281: return block, l+index
							πF.SetLineno(281)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µl, µindex); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µblock, πTemp004).ToObject()
							πR = πTemp001
							continue
							goto Label12
						Label12:
							// line 282: index -= span
							πF.SetLineno(282)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µspan, "span"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ISub(πF, µindex, µspan); πE != nil {
								continue
							}
							µindex = πTemp001
							// line 283: block = block[RGTLNK]
							πF.SetLineno(283)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßRGTLNK); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
								continue
							}
							µblock = πTemp004
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							goto Label3
						Label2:
							// line 285: block = self.right
							πF.SetLineno(285)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							µblock = πTemp001
							// line 286: while block:
							πF.SetLineno(286)
							πF.PushCheckpoint(14)
							πTemp002 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µblock); πE != nil {
								continue
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 287: l, r = 0, n
							πF.SetLineno(287)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßn); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µl = πTemp004
							µr = πTemp005
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßleft, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µblock == πTemp004).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label16
							}
							goto Label17
							// line 288: if block is self.left:
							πF.SetLineno(288)
						Label16:
							// line 289: l = self.leftndx
							πF.SetLineno(289)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßleftndx, nil); πE != nil {
								continue
							}
							µl = πTemp001
							goto Label17
						Label17:
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßright, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µblock == πTemp004).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label18
							}
							goto Label19
							// line 290: if block is self.right:
							πF.SetLineno(290)
						Label18:
							// line 291: r = self.rightndx + 1
							πF.SetLineno(291)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßrightndx, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µr = πTemp001
							goto Label19
						Label19:
							// line 292: negative_span = l-r
							πF.SetLineno(292)
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µl, µr); πE != nil {
								continue
							}
							µnegative_span = πTemp001
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnegative_span, "negative_span"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µindex, µnegative_span); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label20
							}
							goto Label21
							// line 293: if index >= negative_span:
							πF.SetLineno(293)
						Label20:
							// line 294: return block, r+index
							πF.SetLineno(294)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µr, µindex); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µblock, πTemp004).ToObject()
							πR = πTemp001
							continue
							goto Label21
						Label21:
							// line 295: index -= negative_span
							πF.SetLineno(295)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnegative_span, "negative_span"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ISub(πF, µindex, µnegative_span); πE != nil {
								continue
							}
							µindex = πTemp001
							// line 296: block = block[LFTLNK]
							πF.SetLineno(296)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßLFTLNK); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µblock, πTemp001); πE != nil {
								continue
							}
							µblock = πTemp004
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							goto Label3
						Label3:
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("deque index out of range").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 297: raise IndexError("deque index out of range")
							πF.SetLineno(297)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__getref.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 299: def __getitem__(self, index):
					πF.SetLineno(299)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var µblock *πg.Object = πg.UnboundLocal; _ = µblock
						var πTemp001 []*πg.Object
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
							// line 300: block, index = self.__getref(index)
							πF.SetLineno(300)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001[0] = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__getref, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µblock = πTemp002
							µindex = πTemp004
							// line 301: return block[index]
							πF.SetLineno(301)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp002 = µindex
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µblock, πTemp002); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 303: def __setitem__(self, index, value):
					πF.SetLineno(303)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var µblock *πg.Object = πg.UnboundLocal; _ = µblock
						var πTemp001 []*πg.Object
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
							// line 304: block, index = self.__getref(index)
							πF.SetLineno(304)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001[0] = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__getref, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µblock = πTemp002
							µindex = πTemp004
							// line 305: block[index] = value
							πF.SetLineno(305)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp003 = µindex
							if πE = πg.SetItem(πF, µblock, πTemp003, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 307: def __delitem__(self, index):
					πF.SetLineno(307)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var µlength *πg.Object = πg.UnboundLocal; _ = µlength
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
							// line 308: length = len(self)
							πF.SetLineno(308)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlength = πTemp003
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µindex, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 309: if index >= 0:
							πF.SetLineno(309)
						Label1:
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µindex, µlength); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 310: if index >= length:
							πF.SetLineno(310)
						Label4:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("deque index out of range").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 311: raise IndexError("deque index out of range")
							πF.SetLineno(311)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							// line 312: self.rotate(-index)
							πF.SetLineno(312)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, µindex); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrotate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 313: self.popleft()
							πF.SetLineno(313)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpopleft, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 314: self.rotate(index)
							πF.SetLineno(314)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001[0] = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrotate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 316: index = ~index
							πF.SetLineno(316)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invert(πF, µindex); πE != nil {
								continue
							}
							µindex = πTemp002
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlength, "length"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µindex, µlength); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 317: if index >= length:
							πF.SetLineno(317)
						Label6:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("deque index out of range").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 318: raise IndexError("deque index out of range")
							πF.SetLineno(318)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label7
						Label7:
							// line 319: self.rotate(index)
							πF.SetLineno(319)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001[0] = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrotate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 320: self.pop()
							πF.SetLineno(320)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 321: self.rotate(-index)
							πF.SetLineno(321)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, µindex); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrotate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 323: def __reduce_ex__(self, proto):
					πF.SetLineno(323)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "proto", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("__reduce_ex__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µproto *πg.Object = πArgs[1]; _ = µproto
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 324: return type(self), (list(self), self.maxlen)
							πF.SetLineno(324)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmaxlen, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp006, πTemp005).ToObject()
							πTemp001 = πg.NewTuple2(πTemp004, πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß__reduce_ex__.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 326: __hash__ = None
					πF.SetLineno(326)
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 328: def __copy__(self):
					πF.SetLineno(328)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("__copy__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 329: return self.__class__(self, self.maxlen)
							πF.SetLineno(329)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmaxlen, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__copy__.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 332: def __eq__(self, other):
					πF.SetLineno(332)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdeque); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 333: if isinstance(other, deque):
							πF.SetLineno(333)
						Label1:
							// line 334: return list(self) == list(other)
							πF.SetLineno(334)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 336: return NotImplemented
							πF.SetLineno(336)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 338: def __ne__(self, other):
					πF.SetLineno(338)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdeque); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 339: if isinstance(other, deque):
							πF.SetLineno(339)
						Label1:
							// line 340: return list(self) != list(other)
							πF.SetLineno(340)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 342: return NotImplemented
							πF.SetLineno(342)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 344: def __lt__(self, other):
					πF.SetLineno(344)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("__lt__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdeque); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 345: if isinstance(other, deque):
							πF.SetLineno(345)
						Label1:
							// line 346: return list(self) < list(other)
							πF.SetLineno(346)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 348: return NotImplemented
							πF.SetLineno(348)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__lt__.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 350: def __le__(self, other):
					πF.SetLineno(350)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("__le__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdeque); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 351: if isinstance(other, deque):
							πF.SetLineno(351)
						Label1:
							// line 352: return list(self) <= list(other)
							πF.SetLineno(352)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LE(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 354: return NotImplemented
							πF.SetLineno(354)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__le__.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 356: def __gt__(self, other):
					πF.SetLineno(356)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("__gt__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdeque); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 357: if isinstance(other, deque):
							πF.SetLineno(357)
						Label1:
							// line 358: return list(self) > list(other)
							πF.SetLineno(358)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GT(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 360: return NotImplemented
							πF.SetLineno(360)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__gt__.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 362: def __ge__(self, other):
					πF.SetLineno(362)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("__ge__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdeque); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 363: if isinstance(other, deque):
							πF.SetLineno(363)
						Label1:
							// line 364: return list(self) >= list(other)
							πF.SetLineno(364)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GE(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 366: return NotImplemented
							πF.SetLineno(366)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__ge__.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 368: def __iadd__(self, other):
					πF.SetLineno(368)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("__iadd__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 369: self.extend(other)
							πF.SetLineno(369)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßextend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 370: return self
							πF.SetLineno(370)
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
					if πE = πClass.SetItem(πF, ß__iadd__.ToObject(), πTemp035); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("deque").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdeque.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 372: class deque_iterator(object):
			πF.SetLineno(372)
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
			_, πE = πg.NewCode("deque_iterator", "/usr/lib/python2.7/_collections.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 374: def __init__(self, deq, itergen):
					πF.SetLineno(374)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "deq", Def: nil}
					πTemp002[2] = πg.Param{Name: "itergen", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdeq *πg.Object = πArgs[1]; _ = µdeq
						var µitergen *πg.Object = πArgs[2]; _ = µitergen
						var µgiveup *πg.Object = πg.UnboundLocal; _ = µgiveup
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 375: self.counter = len(deq)
							πF.SetLineno(375)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdeq, "deq"); πE != nil {
								continue
							}
							πTemp001[0] = µdeq
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßcounter, πTemp002); πE != nil {
								continue
							}
							// line 376: def giveup():
							πF.SetLineno(376)
							πTemp004 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("giveup", "/usr/lib/python2.7/_collections.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 377: self.counter = 0
									πF.SetLineno(377)
									if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πE = πg.SetAttr(πF, µself, ßcounter, πTemp001); πE != nil {
										continue
									}
									πTemp002 = πF.MakeArgs(1)
									πTemp002[0] = πg.NewStr("deque mutated during iteration").ToObject()
									if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									// line 378: raise RuntimeError("deque mutated during iteration")
									πF.SetLineno(378)
									πE = πF.Raise(πTemp003, nil, nil)
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							µgiveup = πTemp002
							// line 379: self._gen = itergen(deq.state, giveup)
							πF.SetLineno(379)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdeq, "deq"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdeq, ßstate, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µgiveup, "giveup"); πE != nil {
								continue
							}
							πTemp001[1] = µgiveup
							if πE = πg.CheckLocal(πF, µitergen, "itergen"); πE != nil {
								continue
							}
							if πTemp003, πE = µitergen.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_gen, πTemp005); πE != nil {
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
					// line 381: def next(self):
					πF.SetLineno(381)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µres *πg.Object = πg.UnboundLocal; _ = µres
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
							// line 382: res = next(self._gen)
							πF.SetLineno(382)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_gen, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µres = πTemp003
							// line 383: self.counter -= 1
							πF.SetLineno(383)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcounter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ISub(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcounter, πTemp003); πE != nil {
								continue
							}
							// line 384: return res
							πF.SetLineno(384)
							if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
								continue
							}
							πR = µres
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
					// line 386: def __iter__(self):
					πF.SetLineno(386)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 387: return self
							πF.SetLineno(387)
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp004); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("deque_iterator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdeque_iterator.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 389: class defaultdict(dict):
			πF.SetLineno(389)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
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
			_, πE = πg.NewCode("defaultdict", "/usr/lib/python2.7/_collections.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 391: def __init__(self, *args, **kwds):
					πF.SetLineno(391)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_collections.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwds *πg.Object = πArgs[2]; _ = µkwds
						var µdefault_factory *πg.Object = πg.UnboundLocal; _ = µdefault_factory
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.GT(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 392: if len(args) > 0:
							πF.SetLineno(392)
						Label1:
							// line 393: default_factory = args[0]
							πF.SetLineno(393)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µdefault_factory = πTemp003
							// line 394: args = args[1:]
							πF.SetLineno(394)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µargs = πTemp003
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdefault_factory, "default_factory"); πE != nil {
								continue
							}
							πTemp002[0] = µdefault_factory
							if πTemp004, πE = πg.ResolveGlobal(πF, ßcallable); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µdefault_factory, "default_factory"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µdefault_factory != πTemp004).ToObject()
							πTemp001 = πTemp003
						Label4:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 395: if not callable(default_factory) and default_factory is not None:
							πF.SetLineno(395)
						Label5:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("first argument must be callable").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 396: raise TypeError("first argument must be callable")
							πF.SetLineno(396)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label6
						Label6:
							goto Label3
						Label2:
							// line 398: default_factory = None
							πF.SetLineno(398)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µdefault_factory = πTemp001
							goto Label3
						Label3:
							// line 399: self.default_factory = default_factory
							πF.SetLineno(399)
							if πE = πg.CheckLocal(πF, µdefault_factory, "default_factory"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdefault_factory); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdefault_factory, πTemp001); πE != nil {
								continue
							}
							// line 400: super(defaultdict, self).__init__(*args, **kwds)
							πF.SetLineno(400)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdefaultdict); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwds); πE != nil {
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
					// line 402: def __missing__(self, key):
					πF.SetLineno(402)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__missing__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdefault_factory, nil); πE != nil {
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
							// line 404: if self.default_factory is None:
							πF.SetLineno(404)
						Label1:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp005[0] = µkey
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 405: raise KeyError(key)
							πF.SetLineno(405)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 406: self[key] = value = self.default_factory()
							πF.SetLineno(406)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_factory, nil); πE != nil {
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.SetItem(πF, µself, πTemp003, πTemp001); πE != nil {
								continue
							}
							µvalue = πTemp002
							// line 407: return value
							πF.SetLineno(407)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πR = µvalue
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__missing__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 409: def __repr__(self, recurse=set()):
					πF.SetLineno(409)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßset); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "recurse", Def: πTemp006}
					πTemp004 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrecurse *πg.Object = πArgs[1]; _ = µrecurse
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µrecurse, πTemp004); πE != nil {
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
							// line 410: if id(self) in recurse:
							πF.SetLineno(410)
						Label1:
							// line 411: return "defaultdict(...)"
							πF.SetLineno(411)
							πR = πg.NewStr("defaultdict(...)").ToObject()
							continue
							goto Label2
						Label2:
							// line 412: try:
							πF.SetLineno(412)
							πF.PushCheckpoint(3)
							// line 413: recurse.add(id(self))
							πF.SetLineno(413)
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrecurse, ßadd, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 414: return "defaultdict(%s, %s)" % (repr(self.default_factory), super(defaultdict, self).__repr__())
							πF.SetLineno(414)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdefault_factory, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdefaultdict); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.GetAttr(πF, πTemp008, ß__repr__, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("defaultdict(%s, %s)").ToObject(), πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label3
						Label3:
							πTemp009, πTemp010 = πF.RestoreExc(nil, nil)
							// line 416: recurse.remove(id(self))
							πF.SetLineno(416)
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µrecurse, "recurse"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrecurse, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 418: def copy(self):
					πF.SetLineno(418)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 419: return type(self)(self.default_factory, self)
							πF.SetLineno(419)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdefault_factory, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[1] = µself
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 421: def __copy__(self):
					πF.SetLineno(421)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__copy__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 422: return self.copy()
							πF.SetLineno(422)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__copy__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 424: def __reduce__(self):
					πF.SetLineno(424)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__reduce__", "/usr/lib/python2.7/_collections.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 425: """
							πF.SetLineno(425)
							// line 436: return (type(self), (self.default_factory,), None, None, self.iteritems())
							πF.SetLineno(436)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdefault_factory, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple1(πTemp005).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple5(πTemp004, πTemp003, πTemp005, πTemp006, πTemp008).ToObject()
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
					if πE = πClass.SetItem(πF, ß__reduce__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 425: """
					πF.SetLineno(425)
					// line 425: """
					πF.SetLineno(425)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n        __reduce__ must return a 5-tuple as follows:\n\n           - factory function\n           - tuple of args for the factory function\n           - additional state (here None)\n           - sequence iterator (here None)\n           - dictionary iterator (yielding successive (key, value) pairs\n\n           This API is used by pickle.py and copy.py.\n        ").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("\n        __reduce__ must return a 5-tuple as follows:\n\n           - factory function\n           - tuple of args for the factory function\n           - additional state (here None)\n           - sequence iterator (here None)\n           - dictionary iterator (yielding successive (key, value) pairs\n\n           This API is used by pickle.py and copy.py.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß__reduce__); πE != nil {
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
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("defaultdict").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdefaultdict.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_collections", Code)
}

