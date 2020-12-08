package weakref
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/__go__/pygolin"
	// _ "github.com/pygolin/stdlib/UserDict"
	// _ "github.com/pygolin/stdlib/copy"
	// _ "github.com/pygolin/stdlib/exceptions"
	// _ "github.com/pygolin/stdlib/_weakrefset"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/weakref.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßFalse := πg.InternStr("False")
		ßKeyError := πg.InternStr("KeyError")
		ßKeyedRef := πg.InternStr("KeyedRef")
		ßNone := πg.InternStr("None")
		ßReferenceError := πg.InternStr("ReferenceError")
		ßReferenceType := πg.InternStr("ReferenceType")
		ßTypeError := πg.InternStr("TypeError")
		ßUserDict := πg.InternStr("UserDict")
		ßWeakKeyDictionary := πg.InternStr("WeakKeyDictionary")
		ßWeakRefType := πg.InternStr("WeakRefType")
		ßWeakSet := πg.InternStr("WeakSet")
		ßWeakValueDictionary := πg.InternStr("WeakValueDictionary")
		ß_IterationGuard := πg.InternStr("_IterationGuard")
		ß__class__ := πg.InternStr("__class__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__copy__ := πg.InternStr("__copy__")
		ß__deepcopy__ := πg.InternStr("__deepcopy__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß__slots__ := πg.InternStr("__slots__")
		ß_commit_removals := πg.InternStr("_commit_removals")
		ß_iterating := πg.InternStr("_iterating")
		ß_pending_removals := πg.InternStr("_pending_removals")
		ß_remove := πg.InternStr("_remove")
		ß_weakrefset := πg.InternStr("_weakrefset")
		ßappend := πg.InternStr("append")
		ßclear := πg.InternStr("clear")
		ßcopy := πg.InternStr("copy")
		ßdata := πg.InternStr("data")
		ßdeepcopy := πg.InternStr("deepcopy")
		ßexceptions := πg.InternStr("exceptions")
		ßget := πg.InternStr("get")
		ßhas_key := πg.InternStr("has_key")
		ßhasattr := πg.InternStr("hasattr")
		ßid := πg.InternStr("id")
		ßitems := πg.InternStr("items")
		ßiteritems := πg.InternStr("iteritems")
		ßiterkeyrefs := πg.InternStr("iterkeyrefs")
		ßiterkeys := πg.InternStr("iterkeys")
		ßitervaluerefs := πg.InternStr("itervaluerefs")
		ßitervalues := πg.InternStr("itervalues")
		ßkey := πg.InternStr("key")
		ßkeyrefs := πg.InternStr("keyrefs")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßpop := πg.InternStr("pop")
		ßpopitem := πg.InternStr("popitem")
		ßref := πg.InternStr("ref")
		ßset := πg.InternStr("set")
		ßsetdefault := πg.InternStr("setdefault")
		ßsuper := πg.InternStr("super")
		ßtype := πg.InternStr("type")
		ßupdate := πg.InternStr("update")
		ßvaluerefs := πg.InternStr("valuerefs")
		ßvalues := πg.InternStr("values")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Weak reference support for Python.
			πF.SetLineno(1)
			// line 1: """Weak reference support for Python.
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Weak reference support for Python.\n\nThis module is an implementation of PEP 205:\n\nhttp://www.python.org/dev/peps/pep-0205/\n").ToObject()); πE != nil {
				continue
			}
			// line 12: import UserDict
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "UserDict"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßUserDict.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 23: from '__go__/pygolin' import WeakRefType as ReferenceType
			πF.SetLineno(23)
			if πTemp002, πE = πg.ImportModule(πF, "__go__/pygolin"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßWeakRefType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReferenceType.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 24: ref = ReferenceType
			πF.SetLineno(24)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßReferenceType); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßref.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 26: import _weakrefset
			πF.SetLineno(26)
			if πTemp002, πE = πg.ImportModule(πF, "_weakrefset"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_weakrefset.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: WeakSet = _weakrefset.WeakSet
			πF.SetLineno(27)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_weakrefset); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßWeakSet, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWeakSet.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 28: _IterationGuard = _weakrefset._IterationGuard
			πF.SetLineno(28)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_weakrefset); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß_IterationGuard, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_IterationGuard.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 30: import exceptions
			πF.SetLineno(30)
			if πTemp002, πE = πg.ImportModule(πF, "exceptions"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßexceptions.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 31: ReferenceError = exceptions.ReferenceError
			πF.SetLineno(31)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßexceptions); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßReferenceError, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReferenceError.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 41: class WeakValueDictionary(UserDict.UserDict):
			πF.SetLineno(41)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßUserDict); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßUserDict, nil); πE != nil {
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
			_, πE = πg.NewCode("WeakValueDictionary", "/usr/lib/python2.7/weakref.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 42: """Mapping class that references values weakly.
					πF.SetLineno(42)
					// line 42: """Mapping class that references values weakly.
					πF.SetLineno(42)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Mapping class that references values weakly.\n\n    Entries in the dictionary will be discarded when no strong\n    reference to the value exists anymore\n    ").ToObject()); πE != nil {
						continue
					}
					// line 53: def __init__(*args, **kw):
					πF.SetLineno(53)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/weakref.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkw *πg.Object = πArgs[1]; _ = µkw
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µremove *πg.Object = πg.UnboundLocal; _ = µremove
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []πg.Param
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 54: if not args:
							πF.SetLineno(54)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor '__init__' of 'WeakValueDictionary' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 55: raise TypeError("descriptor '__init__' of 'WeakValueDictionary' "
							πF.SetLineno(55)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 57: self = args[0]
							πF.SetLineno(57)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							// line 58: args = args[1:]
							πF.SetLineno(58)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µargs = πTemp004
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 59: if len(args) > 1:
							πF.SetLineno(59)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp006[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("expected at most 1 arguments, got %d").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 60: raise TypeError('expected at most 1 arguments, got %d' % len(args))
							πF.SetLineno(60)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							// line 61: def remove(wr, selfref=ref(self)):
							πF.SetLineno(61)
							πTemp007 = make([]πg.Param, 2)
							πTemp007[0] = πg.Param{Name: "wr", Def: nil}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp007[1] = πg.Param{Name: "selfref", Def: πTemp005}
							πTemp001 = πg.NewFunction(πg.NewCode("remove", "/usr/lib/python2.7/weakref.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µwr *πg.Object = πArgs[0]; _ = µwr
								var µselfref *πg.Object = πArgs[1]; _ = µselfref
								var µself *πg.Object = πg.UnboundLocal; _ = µself
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
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 62: self = selfref()
									πF.SetLineno(62)
									if πE = πg.CheckLocal(πF, µselfref, "selfref"); πE != nil {
										continue
									}
									if πTemp001, πE = µselfref.Call(πF, nil, nil); πE != nil {
										continue
									}
									µself = πTemp001
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(µself != πTemp002).ToObject()
									if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label1
									}
									goto Label2
									// line 63: if self is not None:
									πF.SetLineno(63)
								Label1:
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ß_iterating, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label3
									}
									goto Label4
									// line 64: if self._iterating:
									πF.SetLineno(64)
								Label3:
									// line 65: self._pending_removals.append(wr.key)
									πF.SetLineno(65)
									πTemp004 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µwr, ßkey, nil); πE != nil {
										continue
									}
									πTemp004[0] = πTemp001
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									goto Label5
								Label4:
									// line 67: del self.data[wr.key]
									πF.SetLineno(67)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetAttr(πF, µwr, ßkey, nil); πE != nil {
										continue
									}
									πTemp002 = πTemp005
									if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
										continue
									}
									goto Label5
								Label5:
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
							µremove = πTemp001
							// line 68: self._remove = remove
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µremove, "remove"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µremove); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_remove, πTemp004); πE != nil {
								continue
							}
							// line 70: self._pending_removals = []
							πF.SetLineno(70)
							πTemp003 = make([]*πg.Object, 0)
							πTemp004 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_pending_removals, πTemp005); πE != nil {
								continue
							}
							// line 71: self._iterating = set()
							πF.SetLineno(71)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_iterating, πTemp004); πE != nil {
								continue
							}
							// line 72: UserDict.UserDict.__init__(self, *args, **kw)
							πF.SetLineno(72)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßUserDict); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßUserDict, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Invoke(πF, πTemp004, πTemp003, µargs, nil, µkw); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 74: def _commit_removals(self):
					πF.SetLineno(74)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_commit_removals", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 75: l = self._pending_removals
							πF.SetLineno(75)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
								continue
							}
							µl = πTemp001
							// line 76: d = self.data
							πF.SetLineno(76)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							µd = πTemp001
							// line 79: while l:
							πF.SetLineno(79)
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
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µl); πE != nil {
								continue
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 80: del d[l.pop()]
							πF.SetLineno(80)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µl, ßpop, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							if πE = πg.DelItem(πF, µd, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_commit_removals.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 82: def __getitem__(self, key):
					πF.SetLineno(82)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µo *πg.Object = πg.UnboundLocal; _ = µo
						var πTemp001 *πg.Object
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
							// line 83: o = self.data[key]()
							πF.SetLineno(83)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp001
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µo == πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 84: if o is None:
							πF.SetLineno(84)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							// line 85: raise KeyError, key
							πF.SetLineno(85)
							πE = πF.Raise(πTemp001, µkey, nil)
							continue
							goto Label3
						Label2:
							// line 87: return o
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πR = µo
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 89: def __delitem__(self, key):
					πF.SetLineno(89)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 90: if self._pending_removals:
							πF.SetLineno(90)
						Label1:
							// line 91: self._commit_removals()
							πF.SetLineno(91)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_commit_removals, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 92: del self.data[key]
							πF.SetLineno(92)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 94: def __contains__(self, key):
					πF.SetLineno(94)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__contains__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
							// line 95: try:
							πF.SetLineno(95)
							πF.PushCheckpoint(2)
							// line 96: o = self.data[key]()
							πF.SetLineno(96)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp001
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
							// line 97: except KeyError:
							πF.SetLineno(97)
						Label3:
							// line 98: return False
							πF.SetLineno(98)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 99: return o is not None
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µo != πTemp002).ToObject()
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
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 101: def has_key(self, key):
					πF.SetLineno(101)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("has_key", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
							// line 102: try:
							πF.SetLineno(102)
							πF.PushCheckpoint(2)
							// line 103: o = self.data[key]()
							πF.SetLineno(103)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp001
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
							// line 104: except KeyError:
							πF.SetLineno(104)
						Label3:
							// line 105: return False
							πF.SetLineno(105)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 106: return o is not None
							πF.SetLineno(106)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µo != πTemp002).ToObject()
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
					if πE = πClass.SetItem(πF, ßhas_key.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 108: def __repr__(self):
					πF.SetLineno(108)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 109: return "<WeakValueDictionary at %s>" % id(self)
							πF.SetLineno(109)
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
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<WeakValueDictionary at %s>").ToObject(), πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 111: def __setitem__(self, key, value):
					πF.SetLineno(111)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var πTemp001 *πg.Object
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
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 112: if self._pending_removals:
							πF.SetLineno(112)
						Label1:
							// line 113: self._commit_removals()
							πF.SetLineno(113)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_commit_removals, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 114: self.data[key] = KeyedRef(value, self._remove, key)
							πF.SetLineno(114)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_remove, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004[2] = µkey
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyedRef); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp006 = µkey
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 116: def clear(self):
					πF.SetLineno(116)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("clear", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 117: if self._pending_removals:
							πF.SetLineno(117)
						Label1:
							// line 118: self._commit_removals()
							πF.SetLineno(118)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_commit_removals, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 119: self.data.clear()
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßclear.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 121: def copy(self):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnew *πg.Object = πg.UnboundLocal; _ = µnew
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 122: new = WeakValueDictionary()
							πF.SetLineno(122)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßWeakValueDictionary); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnew = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µkey = πTemp003
								µwr = πTemp006
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 124: o = wr()
							πF.SetLineno(124)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πTemp002, πE = µwr.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp002
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µo != πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 125: if o is not None:
							πF.SetLineno(125)
						Label4:
							// line 126: new[key] = o
							πF.SetLineno(126)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.SetItem(πF, µnew, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 127: return new
							πF.SetLineno(127)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πR = µnew
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 129: __copy__ = copy
					πF.SetLineno(129)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßcopy); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__copy__.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 131: def __deepcopy__(self, memo):
					πF.SetLineno(131)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "memo", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("__deepcopy__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmemo *πg.Object = πArgs[1]; _ = µmemo
						var µcopy *πg.Object = πg.UnboundLocal; _ = µcopy
						var µnew *πg.Object = πg.UnboundLocal; _ = µnew
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
							// line 132: import copy
							πF.SetLineno(132)
							if πTemp002, πE = πg.ImportModule(πF, "copy"); πE != nil {
								continue
							}
							πTemp001 = πTemp002[0]
							µcopy = πTemp001
							// line 133: new = self.__class__()
							πF.SetLineno(133)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnew = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µkey = πTemp004
								µwr = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 135: o = wr()
							πF.SetLineno(135)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πTemp003, πE = µwr.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp003
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µo != πTemp004).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 136: if o is not None:
							πF.SetLineno(136)
						Label4:
							// line 137: new[copy.deepcopy(key, memo)] = o
							πF.SetLineno(137)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µo); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002[0] = µkey
							if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
								continue
							}
							πTemp002[1] = µmemo
							if πE = πg.CheckLocal(πF, µcopy, "copy"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µcopy, ßdeepcopy, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp004 = πTemp008
							if πE = πg.SetItem(πF, µnew, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 138: return new
							πF.SetLineno(138)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πR = µnew
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__deepcopy__.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 140: def get(self, key, default=None):
					πF.SetLineno(140)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp014}
					πTemp013 = πg.NewFunction(πg.NewCode("get", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µo *πg.Object = πg.UnboundLocal; _ = µo
						var πTemp001 *πg.Object
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
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 141: try:
							πF.SetLineno(141)
							πF.PushCheckpoint(2)
							// line 142: wr = self.data[key]
							πF.SetLineno(142)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µwr = πTemp002
							πF.PopCheckpoint()
							// line 146: o = wr()
							πF.SetLineno(146)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πTemp001, πE = µwr.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp001
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µo == πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 147: if o is None:
							πF.SetLineno(147)
						Label3:
							// line 149: return default
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πR = µdefault
							continue
							goto Label5
						Label4:
							// line 151: return o
							πF.SetLineno(151)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πR = µo
							continue
							goto Label5
						Label5:
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 143: except KeyError:
							πF.SetLineno(143)
						Label6:
							// line 144: return default
							πF.SetLineno(144)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πR = µdefault
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
					if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 153: def items(self):
					πF.SetLineno(153)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("items", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µL *πg.Object = πg.UnboundLocal; _ = µL
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
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
							// line 154: L = []
							πF.SetLineno(154)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µL = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µkey = πTemp004
								µwr = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 156: o = wr()
							πF.SetLineno(156)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πTemp003, πE = µwr.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp003
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µo != πTemp004).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 157: if o is not None:
							πF.SetLineno(157)
						Label4:
							// line 158: L.append((key, o))
							πF.SetLineno(158)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µkey, µo).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µL, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 159: return L
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							πR = µL
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßitems.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 161: def iteritems(self):
					πF.SetLineno(161)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("iteritems", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
						var πTemp008 bool
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 3: goto Label3
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								// line 162: with _IterationGuard(self):
								πF.SetLineno(162)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πTemp002, πE = πg.ResolveGlobal(πF, ß_IterationGuard); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßitervalues, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
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
									µwr = πTemp006
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(2)            
								// line 164: value = wr()
								πF.SetLineno(164)
								if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
									continue
								}
								if πTemp006, πE = µwr.Call(πF, nil, nil); πE != nil {
									continue
								}
								µvalue = πTemp006
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp006 = πg.GetBool(µvalue != πTemp007).ToObject()
								if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp009 {
									goto Label5
								}
								goto Label6
								// line 165: if value is not None:
								πF.SetLineno(165)
							Label5:
								// line 166: yield wr.key, value
								πF.SetLineno(166)
								if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µwr, ßkey, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp006 = πg.NewTuple2(πTemp007, µvalue).ToObject()
								πF.PushCheckpoint(7)
								return πTemp006, nil
							Label7:
								πTemp007 = πSent
								goto Label6
							Label6:
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
								πF.PopCheckpoint()
							Label1:
								πTemp010, πTemp011 = nil, nil
								if πE != nil {
									πTemp010, πTemp011 = πF.ExcInfo()
								}
								if πTemp010 != nil {
									πTemp012 = πTemp010.Type()
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
										continue
									}
								} else {
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
										continue
									}
								}
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp010 != nil && πTemp008 != true {
									πE = πF.Raise(nil, nil, nil)
									continue
								}
								if πR != nil {
									continue
								}
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßiteritems.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 168: def iterkeys(self):
					πF.SetLineno(168)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("iterkeys", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µk *πg.Object = πg.UnboundLocal; _ = µk
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
						var πTemp008 bool
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 3: goto Label3
								case 5: goto Label5
								default: panic("unexpected function state")
								}
								// line 169: with _IterationGuard(self):
								πF.SetLineno(169)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πTemp002, πE = πg.ResolveGlobal(πF, ß_IterationGuard); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßiterkeys, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
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
									µk = πTemp006
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(2)            
								// line 171: yield k
								πF.SetLineno(171)
								if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								return µk, nil
							Label5:
								πTemp006 = πSent
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
								πF.PopCheckpoint()
							Label1:
								πTemp010, πTemp011 = nil, nil
								if πE != nil {
									πTemp010, πTemp011 = πF.ExcInfo()
								}
								if πTemp010 != nil {
									πTemp012 = πTemp010.Type()
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
										continue
									}
								} else {
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
										continue
									}
								}
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp010 != nil && πTemp008 != true {
									πE = πF.Raise(nil, nil, nil)
									continue
								}
								if πR != nil {
									continue
								}
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßiterkeys.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 173: __iter__ = iterkeys
					πF.SetLineno(173)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßiterkeys); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 175: def itervaluerefs(self):
					πF.SetLineno(175)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("itervaluerefs", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
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
						var πTemp008 bool
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 3: goto Label3
								case 5: goto Label5
								default: panic("unexpected function state")
								}
								// line 176: """Return an iterator that yields the weak references to the values.
								πF.SetLineno(176)
								// line 185: with _IterationGuard(self):
								πF.SetLineno(185)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πTemp002, πE = πg.ResolveGlobal(πF, ß_IterationGuard); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßitervalues, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
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
									µwr = πTemp006
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(2)            
								// line 187: yield wr
								πF.SetLineno(187)
								if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								return µwr, nil
							Label5:
								πTemp006 = πSent
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
								πF.PopCheckpoint()
							Label1:
								πTemp010, πTemp011 = nil, nil
								if πE != nil {
									πTemp010, πTemp011 = πF.ExcInfo()
								}
								if πTemp010 != nil {
									πTemp012 = πTemp010.Type()
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
										continue
									}
								} else {
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
										continue
									}
								}
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp010 != nil && πTemp008 != true {
									πE = πF.Raise(nil, nil, nil)
									continue
								}
								if πR != nil {
									continue
								}
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßitervaluerefs.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 176: """Return an iterator that yields the weak references to the values.
					πF.SetLineno(176)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("Return an iterator that yields the weak references to the values.\n\n        The references are not guaranteed to be 'live' at the time\n        they are used, so the result of calling the references needs\n        to be checked before being used.  This can be used to avoid\n        creating references that will cause the garbage collector to\n        keep the values around longer than needed.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßitervaluerefs); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
						continue
					}
					// line 189: def itervalues(self):
					πF.SetLineno(189)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("itervalues", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
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
						var πTemp008 bool
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 3: goto Label3
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								// line 190: with _IterationGuard(self):
								πF.SetLineno(190)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πTemp002, πE = πg.ResolveGlobal(πF, ß_IterationGuard); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßitervalues, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
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
									µwr = πTemp006
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(2)            
								// line 192: obj = wr()
								πF.SetLineno(192)
								if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
									continue
								}
								if πTemp006, πE = µwr.Call(πF, nil, nil); πE != nil {
									continue
								}
								µobj = πTemp006
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp006 = πg.GetBool(µobj != πTemp007).ToObject()
								if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp009 {
									goto Label5
								}
								goto Label6
								// line 193: if obj is not None:
								πF.SetLineno(193)
							Label5:
								// line 194: yield obj
								πF.SetLineno(194)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								return µobj, nil
							Label7:
								πTemp006 = πSent
								goto Label6
							Label6:
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
								πF.PopCheckpoint()
							Label1:
								πTemp010, πTemp011 = nil, nil
								if πE != nil {
									πTemp010, πTemp011 = πF.ExcInfo()
								}
								if πTemp010 != nil {
									πTemp012 = πTemp010.Type()
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
										continue
									}
								} else {
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
										continue
									}
								}
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp010 != nil && πTemp008 != true {
									πE = πF.Raise(nil, nil, nil)
									continue
								}
								if πR != nil {
									continue
								}
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßitervalues.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 196: def popitem(self):
					πF.SetLineno(196)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("popitem", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µo *πg.Object = πg.UnboundLocal; _ = µo
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 197: if self._pending_removals:
							πF.SetLineno(197)
						Label1:
							// line 198: self._commit_removals()
							πF.SetLineno(198)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_commit_removals, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 199: while 1:
							πF.SetLineno(199)
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
							if πTemp004, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 200: key, wr = self.data.popitem()
							πF.SetLineno(200)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp001); πE != nil {
								continue
							}
							µkey = πTemp003
							µwr = πTemp005
							// line 201: o = wr()
							πF.SetLineno(201)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πTemp001, πE = µwr.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp001
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µo != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 202: if o is not None:
							πF.SetLineno(202)
						Label6:
							// line 203: return key, o
							πF.SetLineno(203)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µkey, µo).ToObject()
							πR = πTemp001
							continue
							goto Label7
						Label7:
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
					if πE = πClass.SetItem(πF, ßpopitem.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 205: def pop(self, key, *args):
					πF.SetLineno(205)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("pop", "/usr/lib/python2.7/weakref.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var µo *πg.Object = πg.UnboundLocal; _ = µo
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 206: if self._pending_removals:
							πF.SetLineno(206)
						Label1:
							// line 207: self._commit_removals()
							πF.SetLineno(207)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_commit_removals, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 208: try:
							πF.SetLineno(208)
							πF.PushCheckpoint(4)
							// line 209: o = self.data.pop(key)()
							πF.SetLineno(209)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004[0] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp003
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
							// line 210: except KeyError:
							πF.SetLineno(210)
						Label5:
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 211: if args:
							πF.SetLineno(211)
						Label6:
							// line 212: return args[0]
							πF.SetLineno(212)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label7
						Label7:
							// line 213: raise
							πF.SetLineno(213)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µo == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label8
							}
							goto Label9
							// line 214: if o is None:
							πF.SetLineno(214)
						Label8:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							// line 215: raise KeyError, key
							πF.SetLineno(215)
							πE = πF.Raise(πTemp001, µkey, nil)
							continue
							goto Label10
						Label9:
							// line 217: return o
							πF.SetLineno(217)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πR = µo
							continue
							goto Label10
						Label10:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 219: def setdefault(self, key, default=None):
					πF.SetLineno(219)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp022}
					πTemp021 = πg.NewFunction(πg.NewCode("setdefault", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 220: try:
							πF.SetLineno(220)
							πF.PushCheckpoint(2)
							// line 221: wr = self.data[key]
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µwr = πTemp002
							πF.PopCheckpoint()
							// line 228: return wr()
							πF.SetLineno(228)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πTemp001, πE = µwr.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
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
							// line 222: except KeyError:
							πF.SetLineno(222)
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 223: if self._pending_removals:
							πF.SetLineno(223)
						Label4:
							// line 224: self._commit_removals()
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_commit_removals, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label5
						Label5:
							// line 225: self.data[key] = KeyedRef(default, self._remove, key)
							πF.SetLineno(225)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp007[0] = µdefault
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_remove, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007[2] = µkey
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyedRef); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp008 = µkey
							if πE = πg.SetItem(πF, πTemp003, πTemp008, πTemp001); πE != nil {
								continue
							}
							// line 226: return default
							πF.SetLineno(226)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πR = µdefault
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
					if πE = πClass.SetItem(πF, ßsetdefault.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 230: def update(*args, **kwargs):
					πF.SetLineno(230)
					πTemp002 = make([]πg.Param, 0)
					πTemp022 = πg.NewFunction(πg.NewCode("update", "/usr/lib/python2.7/weakref.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µdict *πg.Object = πg.UnboundLocal; _ = µdict
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Dict
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
							case 13: goto Label13
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 231: if not args:
							πF.SetLineno(231)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor 'update' of 'WeakValueDictionary' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 232: raise TypeError("descriptor 'update' of 'WeakValueDictionary' "
							πF.SetLineno(232)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 234: self = args[0]
							πF.SetLineno(234)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							// line 235: args = args[1:]
							πF.SetLineno(235)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µargs = πTemp004
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp003[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 236: if len(args) > 1:
							πF.SetLineno(236)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp006[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("expected at most 1 arguments, got %d").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 237: raise TypeError('expected at most 1 arguments, got %d' % len(args))
							πF.SetLineno(237)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							// line 238: dict = args[0] if args else None
							πF.SetLineno(238)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label5
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µargs, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							goto Label6
						Label5:
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label6:
							µdict = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 239: if self._pending_removals:
							πF.SetLineno(239)
						Label7:
							// line 240: self._commit_removals()
							πF.SetLineno(240)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_commit_removals, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label8
						Label8:
							// line 241: d = self.data
							πF.SetLineno(241)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							µd = πTemp001
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdict != πTemp004).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 242: if dict is not None:
							πF.SetLineno(242)
						Label9:
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							πTemp003[0] = µdict
							πTemp003[1] = ßitems.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label11
							}
							goto Label12
							// line 243: if not hasattr(dict, "items"):
							πF.SetLineno(243)
						Label11:
							// line 244: dict = type({})(dict)
							πF.SetLineno(244)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							πTemp003[0] = µdict
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							πTemp001 = πTemp007.ToObject()
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µdict = πTemp001
							goto Label12
						Label12:
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µdict, ßitems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
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
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
									continue
								}
								µkey = πTemp005
								µo = πTemp009
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 246: d[key] = KeyedRef(o, self._remove, key)
							πF.SetLineno(246)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp003[0] = µo
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_remove, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003[2] = µkey
							if πTemp004, πE = πg.ResolveGlobal(πF, ßKeyedRef); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp009 = µkey
							if πE = πg.SetItem(πF, µd, πTemp009, πTemp004); πE != nil {
								continue
							}
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							goto Label10
						Label10:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp003[0] = µkwargs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label16
							}
							goto Label17
							// line 247: if len(kwargs):
							πF.SetLineno(247)
						Label16:
							// line 248: self.update(kwargs)
							πF.SetLineno(248)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp003[0] = µkwargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label17
						Label17:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßupdate.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 250: def valuerefs(self):
					πF.SetLineno(250)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("valuerefs", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 251: """Return a list of weak references to the values.
							πF.SetLineno(251)
							// line 260: return self.data.values()
							πF.SetLineno(260)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßvalues, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvaluerefs.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 251: """Return a list of weak references to the values.
					πF.SetLineno(251)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("Return a list of weak references to the values.\n\n        The references are not guaranteed to be 'live' at the time\n        they are used, so the result of calling the references needs\n        to be checked before being used.  This can be used to avoid\n        creating references that will cause the garbage collector to\n        keep the values around longer than needed.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßvaluerefs); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
						continue
					}
					// line 262: def values(self):
					πF.SetLineno(262)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("values", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µL *πg.Object = πg.UnboundLocal; _ = µL
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
						var πTemp006 bool
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
							// line 263: L = []
							πF.SetLineno(263)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µL = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µwr = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 265: o = wr()
							πF.SetLineno(265)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πTemp003, πE = µwr.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp003
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µo != πTemp004).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 266: if o is not None:
							πF.SetLineno(266)
						Label4:
							// line 267: L.append(o)
							πF.SetLineno(267)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µL, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 268: return L
							πF.SetLineno(268)
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							πR = µL
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßvalues.ToObject(), πTemp024); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("WeakValueDictionary").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWeakValueDictionary.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 271: class KeyedRef(ref):
			πF.SetLineno(271)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
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
			_, πE = πg.NewCode("KeyedRef", "/usr/lib/python2.7/weakref.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 272: """Specialized reference that includes a key corresponding to the value.
					πF.SetLineno(272)
					// line 272: """Specialized reference that includes a key corresponding to the value.
					πF.SetLineno(272)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Specialized reference that includes a key corresponding to the value.\n\n    This is used in the WeakValueDictionary to avoid having to create\n    a function object for each key stored in the mapping.  A shared\n    callback object can use the 'key' attribute of a KeyedRef instead\n    of getting a reference to the key from an enclosing scope.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 281: __slots__ = "key",
					πF.SetLineno(281)
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), ßkey.ToObject()); πE != nil {
						continue
					}
					// line 283: def __new__(type, ob, callback, key):
					πF.SetLineno(283)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "type", Def: nil}
					πTemp002[1] = πg.Param{Name: "ob", Def: nil}
					πTemp002[2] = πg.Param{Name: "callback", Def: nil}
					πTemp002[3] = πg.Param{Name: "key", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µtype *πg.Object = πArgs[0]; _ = µtype
						var µob *πg.Object = πArgs[1]; _ = µob
						var µcallback *πg.Object = πArgs[2]; _ = µcallback
						var µkey *πg.Object = πArgs[3]; _ = µkey
						var µself *πg.Object = πg.UnboundLocal; _ = µself
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
							// line 284: self = ref.__new__(type, ob, callback)
							πF.SetLineno(284)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
								continue
							}
							πTemp001[0] = µtype
							if πE = πg.CheckLocal(πF, µob, "ob"); πE != nil {
								continue
							}
							πTemp001[1] = µob
							if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
								continue
							}
							πTemp001[2] = µcallback
							if πTemp002, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
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
							// line 285: self.key = key
							πF.SetLineno(285)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µkey); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßkey, πTemp002); πE != nil {
								continue
							}
							// line 286: return self
							πF.SetLineno(286)
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
					// line 288: def __init__(self, ob, callback, key):
					πF.SetLineno(288)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ob", Def: nil}
					πTemp002[2] = πg.Param{Name: "callback", Def: nil}
					πTemp002[3] = πg.Param{Name: "key", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µob *πg.Object = πArgs[1]; _ = µob
						var µcallback *πg.Object = πArgs[2]; _ = µcallback
						var µkey *πg.Object = πArgs[3]; _ = µkey
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
							// line 289: super(KeyedRef,  self).__init__(ob, callback)
							πF.SetLineno(289)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µob, "ob"); πE != nil {
								continue
							}
							πTemp001[0] = µob
							if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
								continue
							}
							πTemp001[1] = µcallback
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyedRef); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ß__init__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp003); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("KeyedRef").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßKeyedRef.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 292: class WeakKeyDictionary(UserDict.UserDict):
			πF.SetLineno(292)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßUserDict); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßUserDict, nil); πE != nil {
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
			_, πE = πg.NewCode("WeakKeyDictionary", "/usr/lib/python2.7/weakref.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 293: """ Mapping class that references keys weakly.
					πF.SetLineno(293)
					// line 293: """ Mapping class that references keys weakly.
					πF.SetLineno(293)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr(" Mapping class that references keys weakly.\n\n    Entries in the dictionary will be discarded when there is no\n    longer a strong reference to the key. This can be used to\n    associate additional data with an object owned by other parts of\n    an application without adding attributes to those objects. This\n    can be especially useful with objects that override attribute\n    accesses.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 303: def __init__(self, dict=None):
					πF.SetLineno(303)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "dict", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdict *πg.Object = πArgs[1]; _ = µdict
						var µremove *πg.Object = πg.UnboundLocal; _ = µremove
						var πTemp001 *πg.Dict
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
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 304: self.data = {}
							πF.SetLineno(304)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp003); πE != nil {
								continue
							}
							// line 305: def remove(k, selfref=ref(self)):
							πF.SetLineno(305)
							πTemp004 = make([]πg.Param, 2)
							πTemp004[0] = πg.Param{Name: "k", Def: nil}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[1] = πg.Param{Name: "selfref", Def: πTemp006}
							πTemp002 = πg.NewFunction(πg.NewCode("remove", "/usr/lib/python2.7/weakref.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µk *πg.Object = πArgs[0]; _ = µk
								var µselfref *πg.Object = πArgs[1]; _ = µselfref
								var µself *πg.Object = πg.UnboundLocal; _ = µself
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
									// line 306: self = selfref()
									πF.SetLineno(306)
									if πE = πg.CheckLocal(πF, µselfref, "selfref"); πE != nil {
										continue
									}
									if πTemp001, πE = µselfref.Call(πF, nil, nil); πE != nil {
										continue
									}
									µself = πTemp001
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(µself != πTemp002).ToObject()
									if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label1
									}
									goto Label2
									// line 307: if self is not None:
									πF.SetLineno(307)
								Label1:
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ß_iterating, nil); πE != nil {
										continue
									}
									if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label3
									}
									goto Label4
									// line 308: if self._iterating:
									πF.SetLineno(308)
								Label3:
									// line 309: self._pending_removals.append(k)
									πF.SetLineno(309)
									πTemp004 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
										continue
									}
									πTemp004[0] = µk
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
										continue
									}
									if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									goto Label5
								Label4:
									// line 311: del self.data[k]
									πF.SetLineno(311)
									if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
										continue
									}
									πTemp002 = µk
									if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
										continue
									}
									goto Label5
								Label5:
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
							µremove = πTemp002
							// line 312: self._remove = remove
							πF.SetLineno(312)
							if πE = πg.CheckLocal(πF, µremove, "remove"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µremove); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_remove, πTemp003); πE != nil {
								continue
							}
							// line 314: self._pending_removals = []
							πF.SetLineno(314)
							πTemp005 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_pending_removals, πTemp006); πE != nil {
								continue
							}
							// line 315: self._iterating = set()
							πF.SetLineno(315)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_iterating, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µdict != πTemp006).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 316: if dict is not None:
							πF.SetLineno(316)
						Label1:
							// line 317: self.update(dict)
							πF.SetLineno(317)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							πTemp005[0] = µdict
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 319: def _commit_removals(self):
					πF.SetLineno(319)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_commit_removals", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							case 2: goto Label2
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 324: l = self._pending_removals
							πF.SetLineno(324)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_pending_removals, nil); πE != nil {
								continue
							}
							µl = πTemp001
							// line 325: d = self.data
							πF.SetLineno(325)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							µd = πTemp001
							// line 326: while l:
							πF.SetLineno(326)
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
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µl); πE != nil {
								continue
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 327: try:
							πF.SetLineno(327)
							πF.PushCheckpoint(5)
							// line 328: del d[l.pop()]
							πF.SetLineno(328)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µl, ßpop, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp005
							if πE = πg.DelItem(πF, µd, πTemp001); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 329: except KeyError:
							πF.SetLineno(329)
						Label6:
							// line 330: pass
							πF.SetLineno(330)
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
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
					if πE = πClass.SetItem(πF, ß_commit_removals.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 332: def __delitem__(self, key):
					πF.SetLineno(332)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 333: del self.data[ref(key)]
							πF.SetLineno(333)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003[0] = µkey
							if πTemp004, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πTemp005
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 335: def __getitem__(self, key):
					πF.SetLineno(335)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 336: return self.data[ref(key)]
							πF.SetLineno(336)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002[0] = µkey
							if πTemp003, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 338: def __repr__(self):
					πF.SetLineno(338)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 339: return "<WeakKeyDictionary at %s>" % id(self)
							πF.SetLineno(339)
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
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<WeakKeyDictionary at %s>").ToObject(), πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 341: def __setitem__(self, key, value):
					πF.SetLineno(341)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
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
							// line 342: self.data[ref(key, self._remove)] = value
							πF.SetLineno(342)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004[0] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_remove, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πTemp006
							if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 344: def copy(self):
					πF.SetLineno(344)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("copy", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnew *πg.Object = πg.UnboundLocal; _ = µnew
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 345: new = WeakKeyDictionary()
							πF.SetLineno(345)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßWeakKeyDictionary); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnew = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µkey = πTemp003
								µvalue = πTemp006
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 347: o = key()
							πF.SetLineno(347)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp002, πE = µkey.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp002
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µo != πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 348: if o is not None:
							πF.SetLineno(348)
						Label4:
							// line 349: new[o] = value
							πF.SetLineno(349)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp003 = µo
							if πE = πg.SetItem(πF, µnew, πTemp003, πTemp002); πE != nil {
								continue
							}
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 350: return new
							πF.SetLineno(350)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πR = µnew
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcopy.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 352: __copy__ = copy
					πF.SetLineno(352)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßcopy); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__copy__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 354: def __deepcopy__(self, memo):
					πF.SetLineno(354)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "memo", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__deepcopy__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmemo *πg.Object = πArgs[1]; _ = µmemo
						var µcopy *πg.Object = πg.UnboundLocal; _ = µcopy
						var µnew *πg.Object = πg.UnboundLocal; _ = µnew
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 355: import copy
							πF.SetLineno(355)
							if πTemp002, πE = πg.ImportModule(πF, "copy"); πE != nil {
								continue
							}
							πTemp001 = πTemp002[0]
							µcopy = πTemp001
							// line 356: new = self.__class__()
							πF.SetLineno(356)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnew = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µkey = πTemp004
								µvalue = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 358: o = key()
							πF.SetLineno(358)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp003, πE = µkey.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp003
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µo != πTemp004).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 359: if o is not None:
							πF.SetLineno(359)
						Label4:
							// line 360: new[o] = copy.deepcopy(value, memo)
							πF.SetLineno(360)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[0] = µvalue
							if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
								continue
							}
							πTemp002[1] = µmemo
							if πE = πg.CheckLocal(πF, µcopy, "copy"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcopy, ßdeepcopy, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp007 = µo
							if πE = πg.SetItem(πF, µnew, πTemp007, πTemp003); πE != nil {
								continue
							}
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 361: return new
							πF.SetLineno(361)
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							πR = µnew
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__deepcopy__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 363: def get(self, key, default=None):
					πF.SetLineno(363)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp011}
					πTemp010 = πg.NewFunction(πg.NewCode("get", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
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
							// line 364: return self.data.get(ref(key),default)
							πF.SetLineno(364)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002[0] = µkey
							if πTemp003, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp001[1] = µdefault
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 366: def has_key(self, key):
					πF.SetLineno(366)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("has_key", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 367: try:
							πF.SetLineno(367)
							πF.PushCheckpoint(2)
							// line 368: wr = ref(key)
							πF.SetLineno(368)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µwr = πTemp003
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 369: except TypeError:
							πF.SetLineno(369)
						Label3:
							// line 370: return 0
							πF.SetLineno(370)
							πR = πg.NewInt(0).ToObject()
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 371: return wr in self.data
							πF.SetLineno(371)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp003, µwr); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006).ToObject()
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
					if πE = πClass.SetItem(πF, ßhas_key.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 373: def __contains__(self, key):
					πF.SetLineno(373)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("__contains__", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 374: try:
							πF.SetLineno(374)
							πF.PushCheckpoint(2)
							// line 375: wr = ref(key)
							πF.SetLineno(375)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πTemp002, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µwr = πTemp003
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 376: except TypeError:
							πF.SetLineno(376)
						Label3:
							// line 377: return 0
							πF.SetLineno(377)
							πR = πg.NewInt(0).ToObject()
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 378: return wr in self.data
							πF.SetLineno(378)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp003, µwr); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006).ToObject()
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
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 380: def items(self):
					πF.SetLineno(380)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("items", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µL *πg.Object = πg.UnboundLocal; _ = µL
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
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
							// line 381: L = []
							πF.SetLineno(381)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µL = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µkey = πTemp004
								µvalue = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 383: o = key()
							πF.SetLineno(383)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp003, πE = µkey.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp003
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µo != πTemp004).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 384: if o is not None:
							πF.SetLineno(384)
						Label4:
							// line 385: L.append((o, value))
							πF.SetLineno(385)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µo, µvalue).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µL, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 386: return L
							πF.SetLineno(386)
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							πR = µL
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßitems.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 388: def iteritems(self):
					πF.SetLineno(388)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("iteritems", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.BaseException
						_ = πTemp011
						var πTemp012 *πg.Traceback
						_ = πTemp012
						var πTemp013 *πg.Type
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 3: goto Label3
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								// line 389: with _IterationGuard(self):
								πF.SetLineno(389)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πTemp002, πE = πg.ResolveGlobal(πF, ß_IterationGuard); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßiteritems, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
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
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp010}}}, πTemp006); πE != nil {
										continue
									}
									µwr = πTemp007
									µvalue = πTemp010
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(2)            
								// line 391: key = wr()
								πF.SetLineno(391)
								if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
									continue
								}
								if πTemp006, πE = µwr.Call(πF, nil, nil); πE != nil {
									continue
								}
								µkey = πTemp006
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp006 = πg.GetBool(µkey != πTemp007).ToObject()
								if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp009 {
									goto Label5
								}
								goto Label6
								// line 392: if key is not None:
								πF.SetLineno(392)
							Label5:
								// line 393: yield key, value
								πF.SetLineno(393)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πTemp006 = πg.NewTuple2(µkey, µvalue).ToObject()
								πF.PushCheckpoint(7)
								return πTemp006, nil
							Label7:
								πTemp007 = πSent
								goto Label6
							Label6:
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
								πF.PopCheckpoint()
							Label1:
								πTemp011, πTemp012 = nil, nil
								if πE != nil {
									πTemp011, πTemp012 = πF.ExcInfo()
								}
								if πTemp011 != nil {
									πTemp013 = πTemp011.Type()
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp013.ToObject(), πTemp011.ToObject(), πTemp012.ToObject()}, nil); πE != nil {
										continue
									}
								} else {
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
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
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßiteritems.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 395: def iterkeyrefs(self):
					πF.SetLineno(395)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("iterkeyrefs", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
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
						var πTemp008 bool
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 3: goto Label3
								case 5: goto Label5
								default: panic("unexpected function state")
								}
								// line 396: """Return an iterator that yields the weak references to the keys.
								πF.SetLineno(396)
								// line 405: with _IterationGuard(self):
								πF.SetLineno(405)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πTemp002, πE = πg.ResolveGlobal(πF, ß_IterationGuard); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßiterkeys, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
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
									µwr = πTemp006
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(2)            
								// line 407: yield wr
								πF.SetLineno(407)
								if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								return µwr, nil
							Label5:
								πTemp006 = πSent
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
								πF.PopCheckpoint()
							Label1:
								πTemp010, πTemp011 = nil, nil
								if πE != nil {
									πTemp010, πTemp011 = πF.ExcInfo()
								}
								if πTemp010 != nil {
									πTemp012 = πTemp010.Type()
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
										continue
									}
								} else {
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
										continue
									}
								}
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp010 != nil && πTemp008 != true {
									πE = πF.Raise(nil, nil, nil)
									continue
								}
								if πR != nil {
									continue
								}
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßiterkeyrefs.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 396: """Return an iterator that yields the weak references to the keys.
					πF.SetLineno(396)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πg.NewStr("Return an iterator that yields the weak references to the keys.\n\n        The references are not guaranteed to be 'live' at the time\n        they are used, so the result of calling the references needs\n        to be checked before being used.  This can be used to avoid\n        creating references that will cause the garbage collector to\n        keep the keys around longer than needed.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßiterkeyrefs); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp017, ß__doc__, πTemp016); πE != nil {
						continue
					}
					// line 409: def iterkeys(self):
					πF.SetLineno(409)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("iterkeys", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
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
						var πTemp008 bool
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 3: goto Label3
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								// line 410: with _IterationGuard(self):
								πF.SetLineno(410)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πTemp002, πE = πg.ResolveGlobal(πF, ß_IterationGuard); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßiterkeys, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
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
									µwr = πTemp006
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(2)            
								// line 412: obj = wr()
								πF.SetLineno(412)
								if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
									continue
								}
								if πTemp006, πE = µwr.Call(πF, nil, nil); πE != nil {
									continue
								}
								µobj = πTemp006
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp006 = πg.GetBool(µobj != πTemp007).ToObject()
								if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp009 {
									goto Label5
								}
								goto Label6
								// line 413: if obj is not None:
								πF.SetLineno(413)
							Label5:
								// line 414: yield obj
								πF.SetLineno(414)
								if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								return µobj, nil
							Label7:
								πTemp006 = πSent
								goto Label6
							Label6:
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
								πF.PopCheckpoint()
							Label1:
								πTemp010, πTemp011 = nil, nil
								if πE != nil {
									πTemp010, πTemp011 = πF.ExcInfo()
								}
								if πTemp010 != nil {
									πTemp012 = πTemp010.Type()
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
										continue
									}
								} else {
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
										continue
									}
								}
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp010 != nil && πTemp008 != true {
									πE = πF.Raise(nil, nil, nil)
									continue
								}
								if πR != nil {
									continue
								}
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßiterkeys.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 416: __iter__ = iterkeys
					πF.SetLineno(416)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßiterkeys); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 418: def itervalues(self):
					πF.SetLineno(418)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("itervalues", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
						var πTemp008 bool
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 3: goto Label3
								case 5: goto Label5
								default: panic("unexpected function state")
								}
								// line 419: with _IterationGuard(self):
								πF.SetLineno(419)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp001[0] = µself
								if πTemp002, πE = πg.ResolveGlobal(πF, ß_IterationGuard); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßitervalues, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
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
									µvalue = πTemp006
								}
								if πE != nil || !πTemp009 {
									continue
								}
								πF.PushCheckpoint(2)            
								// line 421: yield value
								πF.SetLineno(421)
								if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								return µvalue, nil
							Label5:
								πTemp006 = πSent
								continue
							Label3:
								if πE != nil || πR != nil {
									continue
								}
							Label4:
								πF.PopCheckpoint()
							Label1:
								πTemp010, πTemp011 = nil, nil
								if πE != nil {
									πTemp010, πTemp011 = πF.ExcInfo()
								}
								if πTemp010 != nil {
									πTemp012 = πTemp010.Type()
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
										continue
									}
								} else {
									if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
										continue
									}
								}
								if πTemp008, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp010 != nil && πTemp008 != true {
									πE = πF.Raise(nil, nil, nil)
									continue
								}
								if πR != nil {
									continue
								}
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßitervalues.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 423: def keyrefs(self):
					πF.SetLineno(423)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("keyrefs", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 424: """Return a list of weak references to the keys.
							πF.SetLineno(424)
							// line 433: return self.data.keys()
							πF.SetLineno(433)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßkeys, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßkeyrefs.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 424: """Return a list of weak references to the keys.
					πF.SetLineno(424)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πg.NewStr("Return a list of weak references to the keys.\n\n        The references are not guaranteed to be 'live' at the time\n        they are used, so the result of calling the references needs\n        to be checked before being used.  This can be used to avoid\n        creating references that will cause the garbage collector to\n        keep the keys around longer than needed.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßkeyrefs); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp020, ß__doc__, πTemp019); πE != nil {
						continue
					}
					// line 435: def keys(self):
					πF.SetLineno(435)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µL *πg.Object = πg.UnboundLocal; _ = µL
						var µwr *πg.Object = πg.UnboundLocal; _ = µwr
						var µo *πg.Object = πg.UnboundLocal; _ = µo
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
						var πTemp006 bool
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
							// line 436: L = []
							πF.SetLineno(436)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µL = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µwr = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 438: o = wr()
							πF.SetLineno(438)
							if πE = πg.CheckLocal(πF, µwr, "wr"); πE != nil {
								continue
							}
							if πTemp003, πE = µwr.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp003
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µo != πTemp004).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 439: if o is not None:
							πF.SetLineno(439)
						Label4:
							// line 440: L.append(o)
							πF.SetLineno(440)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp001[0] = µo
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µL, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 441: return L
							πF.SetLineno(441)
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							πR = µL
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 443: def popitem(self):
					πF.SetLineno(443)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("popitem", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µo *πg.Object = πg.UnboundLocal; _ = µo
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 444: while 1:
							πF.SetLineno(444)
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
							// line 445: key, value = self.data.popitem()
							πF.SetLineno(445)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µkey = πTemp004
							µvalue = πTemp005
							// line 446: o = key()
							πF.SetLineno(446)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp003, πE = µkey.Call(πF, nil, nil); πE != nil {
								continue
							}
							µo = πTemp003
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µo != πTemp004).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 447: if o is not None:
							πF.SetLineno(447)
						Label4:
							// line 448: return o, value
							πF.SetLineno(448)
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µo, µvalue).ToObject()
							πR = πTemp003
							continue
							goto Label5
						Label5:
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
					if πE = πClass.SetItem(πF, ßpopitem.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 450: def pop(self, key, *args):
					πF.SetLineno(450)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("pop", "/usr/lib/python2.7/weakref.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µargs *πg.Object = πArgs[2]; _ = µargs
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
							// line 451: return self.data.pop(ref(key), *args)
							πF.SetLineno(451)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002[0] = µkey
							if πTemp003, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp004, πTemp001, µargs, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 453: def setdefault(self, key, default=None):
					πF.SetLineno(453)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp023}
					πTemp022 = πg.NewFunction(πg.NewCode("setdefault", "/usr/lib/python2.7/weakref.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
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
							// line 454: return self.data.setdefault(ref(key, self._remove),default)
							πF.SetLineno(454)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002[0] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_remove, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp001[1] = µdefault
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsetdefault.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 456: def update(self, dict=None, **kwargs):
					πF.SetLineno(456)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "dict", Def: πTemp024}
					πTemp023 = πg.NewFunction(πg.NewCode("update", "/usr/lib/python2.7/weakref.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdict *πg.Object = πArgs[1]; _ = µdict
						var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Dict
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
							case 5: goto Label5
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							// line 457: d = self.data
							πF.SetLineno(457)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							µd = πTemp001
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdict != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 458: if dict is not None:
							πF.SetLineno(458)
						Label1:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							πTemp004[0] = µdict
							πTemp004[1] = ßitems.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
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
							// line 459: if not hasattr(dict, "items"):
							πF.SetLineno(459)
						Label3:
							// line 460: dict = type({})(dict)
							πF.SetLineno(460)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							πTemp004[0] = µdict
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							πTemp001 = πTemp007.ToObject()
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µdict = πTemp001
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdict, ßitems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp009}}}, πTemp002); πE != nil {
									continue
								}
								µkey = πTemp005
								µvalue = πTemp009
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(5)            
							// line 462: d[ref(key, self._remove)] = value
							πF.SetLineno(462)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004[0] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ß_remove, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßref); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp005 = πTemp010
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp002); πE != nil {
								continue
							}
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							goto Label2
						Label2:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp004[0] = µkwargs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 463: if len(kwargs):
							πF.SetLineno(463)
						Label8:
							// line 464: self.update(kwargs)
							πF.SetLineno(464)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp004[0] = µkwargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label9
						Label9:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßupdate.ToObject(), πTemp023); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("WeakKeyDictionary").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWeakKeyDictionary.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("weakref", Code)
}

