package _abcoll
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/abc"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßABCMeta := πg.InternStr("ABCMeta")
		ßAttributeError := πg.InternStr("AttributeError")
		ßCallable := πg.InternStr("Callable")
		ßContainer := πg.InternStr("Container")
		ßFalse := πg.InternStr("False")
		ßHashable := πg.InternStr("Hashable")
		ßIndexError := πg.InternStr("IndexError")
		ßIterable := πg.InternStr("Iterable")
		ßIterator := πg.InternStr("Iterator")
		ßKeyError := πg.InternStr("KeyError")
		ßMapping := πg.InternStr("Mapping")
		ßMappingView := πg.InternStr("MappingView")
		ßMutableMapping := πg.InternStr("MutableMapping")
		ßMutableSequence := πg.InternStr("MutableSequence")
		ßMutableSet := πg.InternStr("MutableSet")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßSequence := πg.InternStr("Sequence")
		ßSet := πg.InternStr("Set")
		ßSized := πg.InternStr("Sized")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß__all__ := πg.InternStr("__all__")
		ß__and__ := πg.InternStr("__and__")
		ß__call__ := πg.InternStr("__call__")
		ß__class__ := πg.InternStr("__class__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__ge__ := πg.InternStr("__ge__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__gt__ := πg.InternStr("__gt__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__iadd__ := πg.InternStr("__iadd__")
		ß__iand__ := πg.InternStr("__iand__")
		ß__init__ := πg.InternStr("__init__")
		ß__ior__ := πg.InternStr("__ior__")
		ß__isub__ := πg.InternStr("__isub__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__ixor__ := πg.InternStr("__ixor__")
		ß__le__ := πg.InternStr("__le__")
		ß__len__ := πg.InternStr("__len__")
		ß__lt__ := πg.InternStr("__lt__")
		ß__marker := πg.InternStr("__marker")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__mro__ := πg.InternStr("__mro__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__or__ := πg.InternStr("__or__")
		ß__rand__ := πg.InternStr("__rand__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__reversed__ := πg.InternStr("__reversed__")
		ß__ror__ := πg.InternStr("__ror__")
		ß__rsub__ := πg.InternStr("__rsub__")
		ß__rxor__ := πg.InternStr("__rxor__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß__sub__ := πg.InternStr("__sub__")
		ß__subclasshook__ := πg.InternStr("__subclasshook__")
		ß__xor__ := πg.InternStr("__xor__")
		ß_from_iterable := πg.InternStr("_from_iterable")
		ß_hasattr := πg.InternStr("_hasattr")
		ß_hash := πg.InternStr("_hash")
		ß_mapping := πg.InternStr("_mapping")
		ßabc := πg.InternStr("abc")
		ßabstractmethod := πg.InternStr("abstractmethod")
		ßadd := πg.InternStr("add")
		ßany := πg.InternStr("any")
		ßappend := πg.InternStr("append")
		ßbasestring := πg.InternStr("basestring")
		ßclassmethod := πg.InternStr("classmethod")
		ßclear := πg.InternStr("clear")
		ßcount := πg.InternStr("count")
		ßdict := πg.InternStr("dict")
		ßdiscard := πg.InternStr("discard")
		ßenumerate := πg.InternStr("enumerate")
		ßextend := πg.InternStr("extend")
		ßfrozenset := πg.InternStr("frozenset")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßhash := πg.InternStr("hash")
		ßindex := πg.InternStr("index")
		ßinsert := πg.InternStr("insert")
		ßisdisjoint := πg.InternStr("isdisjoint")
		ßisinstance := πg.InternStr("isinstance")
		ßitems := πg.InternStr("items")
		ßiter := πg.InternStr("iter")
		ßiteritems := πg.InternStr("iteritems")
		ßiterkeys := πg.InternStr("iterkeys")
		ßitervalues := πg.InternStr("itervalues")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßmaxint := πg.InternStr("maxint")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßpop := πg.InternStr("pop")
		ßpopitem := πg.InternStr("popitem")
		ßrange := πg.InternStr("range")
		ßregister := πg.InternStr("register")
		ßremove := πg.InternStr("remove")
		ßreverse := πg.InternStr("reverse")
		ßreversed := πg.InternStr("reversed")
		ßset := πg.InternStr("set")
		ßsetdefault := πg.InternStr("setdefault")
		ßstr := πg.InternStr("str")
		ßsum := πg.InternStr("sum")
		ßsys := πg.InternStr("sys")
		ßtuple := πg.InternStr("tuple")
		ßupdate := πg.InternStr("update")
		ßvalues := πg.InternStr("values")
		ßxrange := πg.InternStr("xrange")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 4: """Abstract Base Classes (ABCs) for collections, according to PEP 3119.
			πF.SetLineno(4)
			// line 4: """Abstract Base Classes (ABCs) for collections, according to PEP 3119.
			πF.SetLineno(4)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Abstract Base Classes (ABCs) for collections, according to PEP 3119.\n\nDON'T USE THIS MODULE DIRECTLY!  The classes here should be imported\nvia collections; they are defined here only to alleviate certain\nbootstrapping issues.  Unit tests are in test_collections.\n").ToObject()); πE != nil {
				continue
			}
			// line 11: import abc
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "abc"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßabc.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: ABCMeta = abc.ABCMeta
			πF.SetLineno(12)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßabc); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßABCMeta, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßABCMeta.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 13: abstractmethod = abc.abstractmethod
			πF.SetLineno(13)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßabc); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßabstractmethod, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßabstractmethod.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 14: import sys
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: __all__ = ["Hashable", "Iterable", "Iterator",
			πF.SetLineno(16)
			πTemp002 = make([]*πg.Object, 13)
			πTemp002[0] = ßHashable.ToObject()
			πTemp002[1] = ßIterable.ToObject()
			πTemp002[2] = ßIterator.ToObject()
			πTemp002[3] = ßSized.ToObject()
			πTemp002[4] = ßContainer.ToObject()
			πTemp002[5] = ßCallable.ToObject()
			πTemp002[6] = ßSet.ToObject()
			πTemp002[7] = ßMutableSet.ToObject()
			πTemp002[8] = ßMapping.ToObject()
			πTemp002[9] = ßMutableMapping.ToObject()
			πTemp002[10] = ßMappingView.ToObject()
			πTemp002[11] = ßSequence.ToObject()
			πTemp002[12] = ßMutableSequence.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 26: def _hasattr(C, attr):
			πF.SetLineno(26)
			πTemp004 = make([]πg.Param, 2)
			πTemp004[0] = πg.Param{Name: "C", Def: nil}
			πTemp004[1] = πg.Param{Name: "attr", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_hasattr", "/usr/lib/python2.7/_abcoll.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µC *πg.Object = πArgs[0]; _ = µC
				var µattr *πg.Object = πArgs[1]; _ = µattr
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
				var πTemp006 *πg.BaseException
				_ = πTemp006
				var πTemp007 *πg.Traceback
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 27: try:
					πF.SetLineno(27)
					πF.PushCheckpoint(2)
					// line 28: return any(attr in B.__dict__ for B in C.__mro__)
					πF.SetLineno(28)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_abcoll.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µB *πg.Object = πg.UnboundLocal; _ = µB
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
								if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µC, ß__mro__, nil); πE != nil {
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
									µB = πTemp002
								}
								if πE != nil || !πTemp004 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 28: return any(attr in B.__dict__ for B in C.__mro__)
								πF.SetLineno(28)
								if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µB, ß__dict__, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Contains(πF, πTemp005, µattr); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(πTemp004).ToObject()
								πF.PushCheckpoint(4)
								return πTemp002, nil
							Label4:
								πTemp005 = πSent
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
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßany); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp005
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp006, πTemp007 = πF.ExcInfo()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label3
					}
					πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
					continue
					// line 29: except AttributeError:
					πF.SetLineno(29)
				Label3:
					// line 31: return hasattr(C, attr)
					πF.SetLineno(31)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
						continue
					}
					πTemp001[0] = µC
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp001[1] = µattr
					if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp005
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
			if πE = πF.Globals().SetItem(πF, ß_hasattr.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: class Hashable(object):
			πF.SetLineno(34)
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
			_, πE = πg.NewCode("Hashable", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 35: __metaclass__ = ABCMeta
					πF.SetLineno(35)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßABCMeta); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__metaclass__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 38: def __hash__(self):
					πF.SetLineno(38)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 39: return 0
							πF.SetLineno(39)
							πR = πg.NewInt(0).ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 37: @abstractmethod
					πF.SetLineno(37)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__hash__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 42: def __subclasshook__(cls, C):
					πF.SetLineno(42)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "C", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__subclasshook__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µC *πg.Object = πArgs[1]; _ = µC
						var µB *πg.Object = πg.UnboundLocal; _ = µB
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							case 5: goto Label5
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßHashable); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µcls == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 43: if cls is Hashable:
							πF.SetLineno(43)
						Label1:
							// line 44: try:
							πF.SetLineno(44)
							πF.PushCheckpoint(4)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µC, ß__mro__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								µB = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(5)            
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µB, ß__dict__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp005, ß__hash__.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 46: if "__hash__" in B.__dict__:
							πF.SetLineno(46)
						Label8:
							πTemp002 = ß__hash__.ToObject()
							if πE = πg.CheckLocal(πF, µB, "B"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µB, ß__dict__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label10
							}
							goto Label11
							// line 47: if B.__dict__["__hash__"]:
							πF.SetLineno(47)
						Label10:
							// line 48: return True
							πF.SetLineno(48)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label11
						Label11:
							// line 49: break
							πF.SetLineno(49)
							πTemp003 = true
							continue
							goto Label9
						Label9:
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 50: except AttributeError:
							πF.SetLineno(50)
						Label12:
							πTemp009 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							πTemp009[0] = µC
							πTemp009[1] = ß__hash__.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp009[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label13
							}
							goto Label14
							// line 52: if getattr(C, "__hash__", None):
							πF.SetLineno(52)
						Label13:
							// line 53: return True
							πF.SetLineno(53)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label14
						Label14:
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							goto Label2
						Label2:
							// line 54: return NotImplemented
							πF.SetLineno(54)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 41: @classmethod
					πF.SetLineno(41)
					πTemp003 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__subclasshook__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp006); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Hashable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHashable.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 57: class Iterable(object):
			πF.SetLineno(57)
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
			_, πE = πg.NewCode("Iterable", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 58: __metaclass__ = ABCMeta
					πF.SetLineno(58)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßABCMeta); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__metaclass__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 61: def __iter__(self):
					πF.SetLineno(61)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 *πg.Object
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
								// line 62: while False:
								πF.SetLineno(62)
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
								if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 63: yield None
								πF.SetLineno(63)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 60: @abstractmethod
					πF.SetLineno(60)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__iter__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 66: def __subclasshook__(cls, C):
					πF.SetLineno(66)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "C", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__subclasshook__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µC *πg.Object = πArgs[1]; _ = µC
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
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µcls == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 67: if cls is Iterable:
							πF.SetLineno(67)
						Label1:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							πTemp004[0] = µC
							πTemp004[1] = ß__iter__.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_hasattr); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 68: if _hasattr(C, "__iter__"):
							πF.SetLineno(68)
						Label3:
							// line 69: return True
							πF.SetLineno(69)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 70: return NotImplemented
							πF.SetLineno(70)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 65: @classmethod
					πF.SetLineno(65)
					πTemp003 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__subclasshook__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp006); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Iterable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIterable.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 72: Iterable.register(str)
			πF.SetLineno(72)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßregister, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 75: class Iterator(Iterable):
			πF.SetLineno(75)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
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
			_, πE = πg.NewCode("Iterator", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 78: def next(self):
					πF.SetLineno(78)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 79: 'Return the next item from the iterator. When exhausted, raise StopIteration'
							πF.SetLineno(79)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							// line 80: raise StopIteration
							πF.SetLineno(80)
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
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 79: 'Return the next item from the iterator. When exhausted, raise StopIteration'
					πF.SetLineno(79)
					// line 79: 'Return the next item from the iterator. When exhausted, raise StopIteration'
					πF.SetLineno(79)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Return the next item from the iterator. When exhausted, raise StopIteration").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Return the next item from the iterator. When exhausted, raise StopIteration").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßnext); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 77: @abstractmethod
					πF.SetLineno(77)
					πTemp005 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßnext); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 82: def __iter__(self):
					πF.SetLineno(82)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 83: return self
							πF.SetLineno(83)
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
					// line 86: def __subclasshook__(cls, C):
					πF.SetLineno(86)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "C", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__subclasshook__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µC *πg.Object = πArgs[1]; _ = µC
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
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIterator); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µcls == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 87: if cls is Iterator:
							πF.SetLineno(87)
						Label1:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							πTemp004[0] = µC
							πTemp004[1] = ßnext.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_hasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001 = πTemp005
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label3
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							πTemp004[0] = µC
							πTemp004[1] = ß__iter__.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_hasattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001 = πTemp005
						Label3:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 88: if _hasattr(C, "next") and _hasattr(C, "__iter__"):
							πF.SetLineno(88)
						Label4:
							// line 89: return True
							πF.SetLineno(89)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label5:
							goto Label2
						Label2:
							// line 90: return NotImplemented
							πF.SetLineno(90)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 85: @classmethod
					πF.SetLineno(85)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß__subclasshook__); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp007); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Iterator").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIterator.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 93: class Sized(object):
			πF.SetLineno(93)
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
			_, πE = πg.NewCode("Sized", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 94: __metaclass__ = ABCMeta
					πF.SetLineno(94)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßABCMeta); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__metaclass__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 97: def __len__(self):
					πF.SetLineno(97)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 98: return 0
							πF.SetLineno(98)
							πR = πg.NewInt(0).ToObject()
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
					// line 96: @abstractmethod
					πF.SetLineno(96)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__len__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 101: def __subclasshook__(cls, C):
					πF.SetLineno(101)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "C", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__subclasshook__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µC *πg.Object = πArgs[1]; _ = µC
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
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSized); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µcls == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 102: if cls is Sized:
							πF.SetLineno(102)
						Label1:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							πTemp004[0] = µC
							πTemp004[1] = ß__len__.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_hasattr); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 103: if _hasattr(C, "__len__"):
							πF.SetLineno(103)
						Label3:
							// line 104: return True
							πF.SetLineno(104)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 105: return NotImplemented
							πF.SetLineno(105)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 100: @classmethod
					πF.SetLineno(100)
					πTemp003 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__subclasshook__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp006); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Sized").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSized.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 108: class Container(object):
			πF.SetLineno(108)
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
			_, πE = πg.NewCode("Container", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 109: __metaclass__ = ABCMeta
					πF.SetLineno(109)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßABCMeta); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__metaclass__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 112: def __contains__(self, x):
					πF.SetLineno(112)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "x", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__contains__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πArgs[1]; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 113: return False
							πF.SetLineno(113)
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
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 111: @abstractmethod
					πF.SetLineno(111)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__contains__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 116: def __subclasshook__(cls, C):
					πF.SetLineno(116)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "C", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__subclasshook__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µC *πg.Object = πArgs[1]; _ = µC
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
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßContainer); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µcls == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 117: if cls is Container:
							πF.SetLineno(117)
						Label1:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							πTemp004[0] = µC
							πTemp004[1] = ß__contains__.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_hasattr); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 118: if _hasattr(C, "__contains__"):
							πF.SetLineno(118)
						Label3:
							// line 119: return True
							πF.SetLineno(119)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 120: return NotImplemented
							πF.SetLineno(120)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 115: @classmethod
					πF.SetLineno(115)
					πTemp003 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__subclasshook__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp006); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Container").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßContainer.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 123: class Callable(object):
			πF.SetLineno(123)
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
			_, πE = πg.NewCode("Callable", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 124: __metaclass__ = ABCMeta
					πF.SetLineno(124)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßABCMeta); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__metaclass__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 127: def __call__(self, *args, **kwds):
					πF.SetLineno(127)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__call__", "/usr/lib/python2.7/_abcoll.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwds *πg.Object = πArgs[2]; _ = µkwds
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 128: return False
							πF.SetLineno(128)
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 126: @abstractmethod
					πF.SetLineno(126)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__call__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 131: def __subclasshook__(cls, C):
					πF.SetLineno(131)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "C", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__subclasshook__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µC *πg.Object = πArgs[1]; _ = µC
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
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßCallable); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µcls == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 132: if cls is Callable:
							πF.SetLineno(132)
						Label1:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							πTemp004[0] = µC
							πTemp004[1] = ß__call__.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_hasattr); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 133: if _hasattr(C, "__call__"):
							πF.SetLineno(133)
						Label3:
							// line 134: return True
							πF.SetLineno(134)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 135: return NotImplemented
							πF.SetLineno(135)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 130: @classmethod
					πF.SetLineno(130)
					πTemp003 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__subclasshook__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__subclasshook__.ToObject(), πTemp006); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Callable").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCallable.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 141: class Set(Sized, Iterable, Container):
			πF.SetLineno(141)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßSized); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
				continue
			}
			πTemp002[1] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßContainer); πE != nil {
				continue
			}
			πTemp002[2] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Set", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp010 *πg.Object
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
				var πTemp018 *πg.Object
				_ = πTemp018
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 142: """A set is a finite, iterable container.
					πF.SetLineno(142)
					// line 142: """A set is a finite, iterable container.
					πF.SetLineno(142)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A set is a finite, iterable container.\n\n    This class provides concrete generic implementations of all\n    methods except for __contains__, __iter__ and __len__.\n\n    To override the comparisons (presumably for speed, as the\n    semantics are fixed), redefine __le__ and __ge__,\n    then the other operations will automatically follow suit.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 152: def __le__(self, other):
					πF.SetLineno(152)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__le__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
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
							case 5: goto Label5
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 153: if not isinstance(other, Set):
							πF.SetLineno(153)
						Label1:
							// line 154: return NotImplemented
							πF.SetLineno(154)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.GT(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 155: if len(self) > len(other):
							πF.SetLineno(155)
						Label3:
							// line 156: return False
							πF.SetLineno(156)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp005 = false
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label7
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µelem = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(5)            
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µother, µelem); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label8
							}
							goto Label9
							// line 158: if elem not in other:
							πF.SetLineno(158)
						Label8:
							// line 159: return False
							πF.SetLineno(159)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label9
						Label9:
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							// line 160: return True
							πF.SetLineno(160)
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
					if πE = πClass.SetItem(πF, ß__le__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 162: def __lt__(self, other):
					πF.SetLineno(162)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__lt__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 163: if not isinstance(other, Set):
							πF.SetLineno(163)
						Label1:
							// line 164: return NotImplemented
							πF.SetLineno(164)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 165: return len(self) < len(other) and self.__le__(other)
							πF.SetLineno(165)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.LT(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__le__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
						Label3:
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
					if πE = πClass.SetItem(πF, ß__lt__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 167: def __gt__(self, other):
					πF.SetLineno(167)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__gt__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 168: if not isinstance(other, Set):
							πF.SetLineno(168)
						Label1:
							// line 169: return NotImplemented
							πF.SetLineno(169)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 170: return len(self) > len(other) and self.__ge__(other)
							πF.SetLineno(170)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.GT(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__ge__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
						Label3:
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
					if πE = πClass.SetItem(πF, ß__gt__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 172: def __ge__(self, other):
					πF.SetLineno(172)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__ge__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µelem *πg.Object = πg.UnboundLocal; _ = µelem
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
							case 5: goto Label5
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 173: if not isinstance(other, Set):
							πF.SetLineno(173)
						Label1:
							// line 174: return NotImplemented
							πF.SetLineno(174)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.LT(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 175: if len(self) < len(other):
							πF.SetLineno(175)
						Label3:
							// line 176: return False
							πF.SetLineno(176)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µother); πE != nil {
								continue
							}
							πF.PushCheckpoint(6)
							πTemp005 = false
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label7
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µelem = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(5)            
							if πE = πg.CheckLocal(πF, µelem, "elem"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µself, µelem); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label8
							}
							goto Label9
							// line 178: if elem not in self:
							πF.SetLineno(178)
						Label8:
							// line 179: return False
							πF.SetLineno(179)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label9
						Label9:
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							// line 180: return True
							πF.SetLineno(180)
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
					if πE = πClass.SetItem(πF, ß__ge__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 182: def __eq__(self, other):
					πF.SetLineno(182)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 183: if not isinstance(other, Set):
							πF.SetLineno(183)
						Label1:
							// line 184: return NotImplemented
							πF.SetLineno(184)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 185: return len(self) == len(other) and self.__le__(other)
							πF.SetLineno(185)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Eq(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__le__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
						Label3:
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 187: def __ne__(self, other):
					πF.SetLineno(187)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 188: return not (self == other)
							πF.SetLineno(188)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 191: def _from_iterable(cls, it):
					πF.SetLineno(191)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "it", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_from_iterable", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µit *πg.Object = πArgs[1]; _ = µit
						var πTemp001 []*πg.Object
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
							// line 192: '''Construct an instance of the class from any iterable input.
							πF.SetLineno(192)
							// line 197: return cls(it)
							πF.SetLineno(197)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							πTemp001[0] = µit
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = µcls.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_from_iterable.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 192: '''Construct an instance of the class from any iterable input.
					πF.SetLineno(192)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Construct an instance of the class from any iterable input.\n\n        Must override this method if the class constructor signature\n        does not accept an iterable for an input.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_from_iterable); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 190: @classmethod
					πF.SetLineno(190)
					πTemp011 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_from_iterable); πE != nil {
						continue
					}
					πTemp011[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					if πE = πClass.SetItem(πF, ß_from_iterable.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 199: def __and__(self, other):
					πF.SetLineno(199)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__and__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 []πg.Param
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 200: if not isinstance(other, Iterable):
							πF.SetLineno(200)
						Label1:
							// line 201: return NotImplemented
							πF.SetLineno(201)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 202: return self._from_iterable(value for value in other if value in self)
							πF.SetLineno(202)
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_abcoll.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
										case 6: goto Label6
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µother); πE != nil {
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
											µvalue = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.Contains(πF, µself, µvalue); πE != nil {
											continue
										}
										πTemp004 = πg.GetBool(πTemp003).ToObject()
										if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 202: return self._from_iterable(value for value in other if value in self)
										πF.SetLineno(202)
									Label4:
										// line 202: return self._from_iterable(value for value in other if value in self)
										πF.SetLineno(202)
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µvalue, nil
									Label6:
										πTemp004 = πSent
										goto Label5
									Label5:
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
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_from_iterable, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp004
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__and__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 204: __rand__ = __and__
					πF.SetLineno(204)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß__and__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__rand__.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 206: def isdisjoint(self, other):
					πF.SetLineno(206)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("isdisjoint", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
							// line 207: 'Return True if two sets have a null intersection.'
							πF.SetLineno(207)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µother); πE != nil {
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
								µvalue = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µself, µvalue); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 209: if value in self:
							πF.SetLineno(209)
						Label4:
							// line 210: return False
							πF.SetLineno(210)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp004
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 211: return True
							πF.SetLineno(211)
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
					if πE = πClass.SetItem(πF, ßisdisjoint.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 207: 'Return True if two sets have a null intersection.'
					πF.SetLineno(207)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Return True if two sets have a null intersection.").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßisdisjoint); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 213: def __or__(self, other):
					πF.SetLineno(213)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("__or__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
						var µchain *πg.Object = πg.UnboundLocal; _ = µchain
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
						var πTemp006 []πg.Param
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 214: if not isinstance(other, Iterable):
							πF.SetLineno(214)
						Label1:
							// line 215: return NotImplemented
							πF.SetLineno(215)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 216: chain = (e for s in (self, other) for e in s)
							πF.SetLineno(216)
							πTemp006 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_abcoll.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µs *πg.Object = πg.UnboundLocal; _ = µs
								var µe *πg.Object = πg.UnboundLocal; _ = µe
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
								var πTemp006 *πg.Object
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
										case 5: goto Label5
										case 7: goto Label7
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
											continue
										}
										πTemp002 = πg.NewTuple2(µself, µother).ToObject()
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
											µs = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.Iter(πF, µs); πE != nil {
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
										if πTemp006, πE = πg.Next(πF, πTemp002); πE != nil {
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
											µe = πTemp006
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(4)            
										// line 216: chain = (e for s in (self, other) for e in s)
										πF.SetLineno(216)
										if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
											continue
										}
										πF.PushCheckpoint(7)
										return µe, nil
									Label7:
										πTemp006 = πSent
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
									return nil, πE
								}).ToObject(), nil
							}), πF.Globals()).ToObject()
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µchain = πTemp003
							// line 217: return self._from_iterable(chain)
							πF.SetLineno(217)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchain, "chain"); πE != nil {
								continue
							}
							πTemp002[0] = µchain
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_from_iterable, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp004
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__or__.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 219: __ror__ = __or__
					πF.SetLineno(219)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ß__or__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__ror__.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 221: def __sub__(self, other):
					πF.SetLineno(221)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("__sub__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 []πg.Param
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 222: if not isinstance(other, Set):
							πF.SetLineno(222)
						Label1:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 223: if not isinstance(other, Iterable):
							πF.SetLineno(223)
						Label3:
							// line 224: return NotImplemented
							πF.SetLineno(224)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							// line 225: other = self._from_iterable(other)
							πF.SetLineno(225)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_from_iterable, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µother = πTemp003
							goto Label2
						Label2:
							// line 226: return self._from_iterable(value for value in self
							πF.SetLineno(226)
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_abcoll.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
										case 6: goto Label6
										default: panic("unexpected function state")
										}
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
											µvalue = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.Contains(πF, µother, µvalue); πE != nil {
											continue
										}
										πTemp004 = πg.GetBool(!πTemp003).ToObject()
										if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 226: return self._from_iterable(value for value in self
										πF.SetLineno(226)
									Label4:
										// line 226: return self._from_iterable(value for value in self
										πF.SetLineno(226)
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µvalue, nil
									Label6:
										πTemp004 = πSent
										goto Label5
									Label5:
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
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_from_iterable, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp004
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__sub__.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 229: def __rsub__(self, other):
					πF.SetLineno(229)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("__rsub__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πTemp006 []πg.Param
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 230: if not isinstance(other, Set):
							πF.SetLineno(230)
						Label1:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 231: if not isinstance(other, Iterable):
							πF.SetLineno(231)
						Label3:
							// line 232: return NotImplemented
							πF.SetLineno(232)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							// line 233: other = self._from_iterable(other)
							πF.SetLineno(233)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_from_iterable, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µother = πTemp003
							goto Label2
						Label2:
							// line 234: return self._from_iterable(value for value in other
							πF.SetLineno(234)
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_abcoll.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
										case 6: goto Label6
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µother); πE != nil {
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
											µvalue = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.Contains(πF, µself, µvalue); πE != nil {
											continue
										}
										πTemp004 = πg.GetBool(!πTemp003).ToObject()
										if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 234: return self._from_iterable(value for value in other
										πF.SetLineno(234)
									Label4:
										// line 234: return self._from_iterable(value for value in other
										πF.SetLineno(234)
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µvalue, nil
									Label6:
										πTemp004 = πSent
										goto Label5
									Label5:
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
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_from_iterable, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp004
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__rsub__.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 237: def __xor__(self, other):
					πF.SetLineno(237)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("__xor__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 238: if not isinstance(other, Set):
							πF.SetLineno(238)
						Label1:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label3
							}
							goto Label4
							// line 239: if not isinstance(other, Iterable):
							πF.SetLineno(239)
						Label3:
							// line 240: return NotImplemented
							πF.SetLineno(240)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							// line 241: other = self._from_iterable(other)
							πF.SetLineno(241)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_from_iterable, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µother = πTemp003
							goto Label2
						Label2:
							// line 242: return (self - other) | (other - self)
							πF.SetLineno(242)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µself, µother); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µother, µself); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Or(πF, πTemp003, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__xor__.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 244: __rxor__ = __xor__
					πF.SetLineno(244)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ß__xor__); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__rxor__.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 247: __hash__ = None
					πF.SetLineno(247)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 249: def _hash(self):
					πF.SetLineno(249)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("_hash", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µMAX *πg.Object = πg.UnboundLocal; _ = µMAX
						var µMASK *πg.Object = πg.UnboundLocal; _ = µMASK
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µh *πg.Object = πg.UnboundLocal; _ = µh
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µhx *πg.Object = πg.UnboundLocal; _ = µhx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 250: """Compute the hash value of a set.
							πF.SetLineno(250)
							// line 264: MAX = sys.maxint
							πF.SetLineno(264)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmaxint, nil); πE != nil {
								continue
							}
							µMAX = πTemp002
							// line 265: MASK = 2 * MAX + 1
							πF.SetLineno(265)
							if πE = πg.CheckLocal(πF, µMAX, "MAX"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), µMAX); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µMASK = πTemp001
							// line 266: n = len(self)
							πF.SetLineno(266)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µn = πTemp002
							// line 267: h = 1927868237 * (n + 1)
							πF.SetLineno(267)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πg.NewInt(1927868237).ToObject(), πTemp002); πE != nil {
								continue
							}
							µh = πTemp001
							// line 268: h &= MASK
							πF.SetLineno(268)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µMASK, "MASK"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAnd(πF, µh, µMASK); πE != nil {
								continue
							}
							µh = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µself); πE != nil {
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
								µx = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 270: hx = hash(x)
							πF.SetLineno(270)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µhx = πTemp006
							// line 271: h ^= (hx ^ (hx << 16) ^ 89869747)  * 3644798167
							πF.SetLineno(271)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhx, "hx"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhx, "hx"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.LShift(πF, µhx, πg.NewInt(16).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Xor(πF, µhx, πTemp008); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Xor(πF, πTemp007, πg.NewInt(89869747).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp006, πg.NewInt(3644798167).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IXor(πF, µh, πTemp002); πE != nil {
								continue
							}
							µh = πTemp006
							// line 272: h &= MASK
							πF.SetLineno(272)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µMASK, "MASK"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAnd(πF, µh, µMASK); πE != nil {
								continue
							}
							µh = πTemp002
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 273: h = h * 69069 + 907133923
							πF.SetLineno(273)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, µh, πg.NewInt(69069).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewInt(907133923).ToObject()); πE != nil {
								continue
							}
							µh = πTemp001
							// line 274: h &= MASK
							πF.SetLineno(274)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µMASK, "MASK"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAnd(πF, µh, µMASK); πE != nil {
								continue
							}
							µh = πTemp001
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µMAX, "MAX"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µh, µMAX); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 275: if h > MAX:
							πF.SetLineno(275)
						Label4:
							// line 276: h -= MASK + 1
							πF.SetLineno(276)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µMASK, "MASK"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µMASK, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, µh, πTemp001); πE != nil {
								continue
							}
							µh = πTemp002
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µh, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 277: if h == -1:
							πF.SetLineno(277)
						Label6:
							// line 278: h = 590923713
							πF.SetLineno(278)
							µh = πg.NewInt(590923713).ToObject()
							goto Label7
						Label7:
							// line 279: return h
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							πR = µh
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_hash.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 250: """Compute the hash value of a set.
					πF.SetLineno(250)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Compute the hash value of a set.\n\n        Note that we don't define __hash__: not all sets are hashable.\n        But if you define a hashable set type, its __hash__ should\n        call this function.\n\n        This must be compatible __eq__.\n\n        All sets ought to compare equal if they contain the same\n        elements, regardless of how they are implemented, and\n        regardless of the order of the elements; so there's not much\n        freedom for __eq__ or __hash__.  We match the algorithm used\n        by the built-in frozenset type.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ß_hash); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp018, ß__doc__, πTemp017); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Set").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSet.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 281: Set.register(frozenset)
			πF.SetLineno(281)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßregister, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 284: class MutableSet(Set):
			πF.SetLineno(284)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
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
			_, πE = πg.NewCode("MutableSet", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 285: """A mutable set is a finite, iterable container.
					πF.SetLineno(285)
					// line 285: """A mutable set is a finite, iterable container.
					πF.SetLineno(285)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A mutable set is a finite, iterable container.\n\n    This class provides concrete generic implementations of all\n    methods except for __contains__, __iter__, __len__,\n    add(), and discard().\n\n    To override the comparisons (presumably for speed, as the\n    semantics are fixed), all you have to do is redefine __le__ and\n    then the other operations will automatically follow suit.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 297: def add(self, value):
					πF.SetLineno(297)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("add", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 298: """Add an element."""
							πF.SetLineno(298)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 299: raise NotImplementedError
							πF.SetLineno(299)
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
					if πE = πClass.SetItem(πF, ßadd.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 298: """Add an element."""
					πF.SetLineno(298)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Add an element.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßadd); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 296: @abstractmethod
					πF.SetLineno(296)
					πTemp005 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßadd); πE != nil {
						continue
					}
					πTemp005[0] = πTemp003
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßadd.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 302: def discard(self, value):
					πF.SetLineno(302)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("discard", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 303: """Remove an element.  Do not raise an exception if absent."""
							πF.SetLineno(303)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 304: raise NotImplementedError
							πF.SetLineno(304)
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
					if πE = πClass.SetItem(πF, ßdiscard.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 303: """Remove an element.  Do not raise an exception if absent."""
					πF.SetLineno(303)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Remove an element.  Do not raise an exception if absent.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdiscard); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 301: @abstractmethod
					πF.SetLineno(301)
					πTemp005 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßdiscard); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßdiscard.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 306: def remove(self, value):
					πF.SetLineno(306)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("remove", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
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
							// line 307: """Remove an element. If not a member, raise a KeyError."""
							πF.SetLineno(307)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µself, µvalue); πE != nil {
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
							// line 308: if value not in self:
							πF.SetLineno(308)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[0] = µvalue
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 309: raise KeyError(value)
							πF.SetLineno(309)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 310: self.discard(value)
							πF.SetLineno(310)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdiscard, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßremove.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 307: """Remove an element. If not a member, raise a KeyError."""
					πF.SetLineno(307)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Remove an element. If not a member, raise a KeyError.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßremove); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 312: def pop(self):
					πF.SetLineno(312)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("pop", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µit *πg.Object = πg.UnboundLocal; _ = µit
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
							// line 313: """Return the popped value.  Raise KeyError if empty."""
							πF.SetLineno(313)
							// line 314: it = iter(self)
							πF.SetLineno(314)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µit = πTemp003
							// line 315: try:
							πF.SetLineno(315)
							πF.PushCheckpoint(2)
							// line 316: value = next(it)
							πF.SetLineno(316)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							πTemp001[0] = µit
							if πTemp002, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvalue = πTemp003
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
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
							// line 317: except StopIteration:
							πF.SetLineno(317)
						Label3:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							// line 318: raise KeyError
							πF.SetLineno(318)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 319: self.discard(value)
							πF.SetLineno(319)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdiscard, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 320: return value
							πF.SetLineno(320)
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
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 313: """Return the popped value.  Raise KeyError if empty."""
					πF.SetLineno(313)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Return the popped value.  Raise KeyError if empty.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßpop); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 322: def clear(self):
					πF.SetLineno(322)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("clear", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 bool
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
							// line 323: """This is slow (creates N new iterators!) but effective."""
							πF.SetLineno(323)
							// line 324: try:
							πF.SetLineno(324)
							πF.PushCheckpoint(2)
							// line 325: while True:
							πF.SetLineno(325)
							πF.PushCheckpoint(4)
							πTemp001 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp001 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 326: self.pop()
							πF.SetLineno(326)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpop, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
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
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label6
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 327: except KeyError:
							πF.SetLineno(327)
						Label6:
							// line 328: pass
							πF.SetLineno(328)
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
					if πE = πClass.SetItem(πF, ßclear.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 323: """This is slow (creates N new iterators!) but effective."""
					πF.SetLineno(323)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("This is slow (creates N new iterators!) but effective.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßclear); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 330: def __ior__(self, it):
					πF.SetLineno(330)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "it", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__ior__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µit *πg.Object = πArgs[1]; _ = µit
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µit); πE != nil {
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
								µvalue = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 332: self.add(value)
							πF.SetLineno(332)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßadd, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 333: return self
							πF.SetLineno(333)
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
					if πE = πClass.SetItem(πF, ß__ior__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 335: def __iand__(self, it):
					πF.SetLineno(335)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "it", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__iand__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µit *πg.Object = πArgs[1]; _ = µit
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µself, µit); πE != nil {
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
								µvalue = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 337: self.discard(value)
							πF.SetLineno(337)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdiscard, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 338: return self
							πF.SetLineno(338)
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
					if πE = πClass.SetItem(πF, ß__iand__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 340: def __ixor__(self, it):
					πF.SetLineno(340)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "it", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("__ixor__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µit *πg.Object = πArgs[1]; _ = µit
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
						var πTemp006 bool
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6: goto Label6
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µit == µself).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 341: if it is self:
							πF.SetLineno(341)
						Label1:
							// line 342: self.clear()
							πF.SetLineno(342)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßclear, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label3
						Label2:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							πTemp004[0] = µit
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSet); πE != nil {
								continue
							}
							πTemp004[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 344: if not isinstance(it, Set):
							πF.SetLineno(344)
						Label4:
							// line 345: it = self._from_iterable(it)
							πF.SetLineno(345)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							πTemp004[0] = µit
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_from_iterable, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µit = πTemp003
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µit); πE != nil {
								continue
							}
							πF.PushCheckpoint(7)
							πTemp002 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label8
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
								µvalue = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(6)            
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µself, µvalue); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 347: if value in self:
							πF.SetLineno(347)
						Label9:
							// line 348: self.discard(value)
							πF.SetLineno(348)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdiscard, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label11
						Label10:
							// line 350: self.add(value)
							πF.SetLineno(350)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßadd, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label11
						Label11:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							goto Label3
						Label3:
							// line 351: return self
							πF.SetLineno(351)
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
					if πE = πClass.SetItem(πF, ß__ixor__.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 353: def __isub__(self, it):
					πF.SetLineno(353)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "it", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("__isub__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µit *πg.Object = πArgs[1]; _ = µit
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µit == µself).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 354: if it is self:
							πF.SetLineno(354)
						Label1:
							// line 355: self.clear()
							πF.SetLineno(355)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßclear, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label3
						Label2:
							if πE = πg.CheckLocal(πF, µit, "it"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µit); πE != nil {
								continue
							}
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
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µvalue = πTemp003
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 358: self.discard(value)
							πF.SetLineno(358)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp005[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdiscard, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							goto Label3
						Label3:
							// line 359: return self
							πF.SetLineno(359)
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
					if πE = πClass.SetItem(πF, ß__isub__.ToObject(), πTemp011); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("MutableSet").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMutableSet.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 361: MutableSet.register(set)
			πF.SetLineno(361)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßMutableSet); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßregister, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 367: class Mapping(Sized, Iterable, Container):
			πF.SetLineno(367)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßSized); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
				continue
			}
			πTemp002[1] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßContainer); πE != nil {
				continue
			}
			πTemp002[2] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Mapping", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
					// line 369: """A Mapping is a generic container for associating key/value
					πF.SetLineno(369)
					// line 369: """A Mapping is a generic container for associating key/value
					πF.SetLineno(369)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A Mapping is a generic container for associating key/value\n    pairs.\n\n    This class provides concrete generic implementations of all\n    methods except for __getitem__, __iter__, and __len__.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 378: def __getitem__(self, key):
					πF.SetLineno(378)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							// line 379: raise KeyError
							πF.SetLineno(379)
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 377: @abstractmethod
					πF.SetLineno(377)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__getitem__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 381: def get(self, key, default=None):
					πF.SetLineno(381)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp005}
					πTemp004 = πg.NewFunction(πg.NewCode("get", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 382: 'D.get(k[,d]) -> D[k] if k in D, else d.  d defaults to None.'
							πF.SetLineno(382)
							// line 383: try:
							πF.SetLineno(383)
							πF.PushCheckpoint(2)
							// line 384: return self[key]
							πF.SetLineno(384)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
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
							πTemp003, πTemp004 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
							// line 385: except KeyError:
							πF.SetLineno(385)
						Label3:
							// line 386: return default
							πF.SetLineno(386)
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
					if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 382: 'D.get(k[,d]) -> D[k] if k in D, else d.  d defaults to None.'
					πF.SetLineno(382)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("D.get(k[,d]) -> D[k] if k in D, else d.  d defaults to None.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßget); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 388: def __contains__(self, key):
					πF.SetLineno(388)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__contains__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 389: try:
							πF.SetLineno(389)
							πF.PushCheckpoint(2)
							// line 390: self[key]
							πF.SetLineno(390)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							// line 394: return True
							πF.SetLineno(394)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
							πTemp003, πTemp004 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
							// line 391: except KeyError:
							πF.SetLineno(391)
						Label3:
							// line 392: return False
							πF.SetLineno(392)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
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
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 396: def iterkeys(self):
					πF.SetLineno(396)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("iterkeys", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 397: 'D.iterkeys() -> an iterator over the keys of D'
							πF.SetLineno(397)
							// line 398: return iter(self)
							πF.SetLineno(398)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
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
					if πE = πClass.SetItem(πF, ßiterkeys.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 397: 'D.iterkeys() -> an iterator over the keys of D'
					πF.SetLineno(397)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("D.iterkeys() -> an iterator over the keys of D").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßiterkeys); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 400: def itervalues(self):
					πF.SetLineno(400)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("itervalues", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								// line 401: 'D.itervalues() -> an iterator over the values of D'
								πF.SetLineno(401)
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
									µkey = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 403: yield self[key]
								πF.SetLineno(403)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp004 = µkey
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp005, nil
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
					if πE = πClass.SetItem(πF, ßitervalues.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 401: 'D.itervalues() -> an iterator over the values of D'
					πF.SetLineno(401)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("D.itervalues() -> an iterator over the values of D").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßitervalues); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 405: def iteritems(self):
					πF.SetLineno(405)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("iteritems", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
						var πTemp006 *πg.Object
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
								// line 406: 'D.iteritems() -> an iterator over the (key, value) items of D'
								πF.SetLineno(406)
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
									µkey = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 408: yield (key, self[key])
								πF.SetLineno(408)
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
									continue
								}
								πTemp005 = µkey
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µself, πTemp005); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple2(µkey, πTemp006).ToObject()
								πF.PushCheckpoint(4)
								return πTemp004, nil
							Label4:
								πTemp005 = πSent
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
					if πE = πClass.SetItem(πF, ßiteritems.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 406: 'D.iteritems() -> an iterator over the (key, value) items of D'
					πF.SetLineno(406)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("D.iteritems() -> an iterator over the (key, value) items of D").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßiteritems); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 410: def keys(self):
					πF.SetLineno(410)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 411: "D.keys() -> list of D's keys"
							πF.SetLineno(411)
							// line 412: return list(self)
							πF.SetLineno(412)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
					if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 411: "D.keys() -> list of D's keys"
					πF.SetLineno(411)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("D.keys() -> list of D's keys").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßkeys); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 414: def items(self):
					πF.SetLineno(414)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("items", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
							// line 415: "D.items() -> list of D's (key, value) pairs, as 2-tuples"
							πF.SetLineno(415)
							// line 416: return [(key, self[key]) for key in self]
							πF.SetLineno(416)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_abcoll.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
								var πTemp006 *πg.Object
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
											µkey = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 416: return [(key, self[key]) for key in self]
										πF.SetLineno(416)
										if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
											continue
										}
										πTemp005 = µkey
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp006, πE = πg.GetItem(πF, µself, πTemp005); πE != nil {
											continue
										}
										πTemp004 = πg.NewTuple2(µkey, πTemp006).ToObject()
										πF.PushCheckpoint(4)
										return πTemp004, nil
									Label4:
										πTemp005 = πSent
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßitems.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 415: "D.items() -> list of D's (key, value) pairs, as 2-tuples"
					πF.SetLineno(415)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("D.items() -> list of D's (key, value) pairs, as 2-tuples").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßitems); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 418: def values(self):
					πF.SetLineno(418)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("values", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
							// line 419: "D.values() -> list of D's values"
							πF.SetLineno(419)
							// line 420: return [self[key] for key in self]
							πF.SetLineno(420)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_abcoll.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
											µkey = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 420: return [self[key] for key in self]
										πF.SetLineno(420)
										if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
											continue
										}
										πTemp004 = µkey
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp005, nil
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßvalues.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 419: "D.values() -> list of D's values"
					πF.SetLineno(419)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("D.values() -> list of D's values").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßvalues); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 423: __hash__ = None
					πF.SetLineno(423)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 425: def __eq__(self, other):
					πF.SetLineno(425)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßMapping); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
							// line 426: if not isinstance(other, Mapping):
							πF.SetLineno(426)
						Label1:
							// line 427: return NotImplemented
							πF.SetLineno(427)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 428: return dict(self.items()) == dict(other.items())
							πF.SetLineno(428)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßitems, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Eq(πF, πTemp004, πTemp006); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 430: def __ne__(self, other):
					πF.SetLineno(430)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 431: return not (self == other)
							πF.SetLineno(431)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µself, µother); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp013); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Mapping").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMapping.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 433: class MappingView(Sized):
			πF.SetLineno(433)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßSized); πE != nil {
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
			_, πE = πg.NewCode("MappingView", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 435: def __init__(self, mapping):
					πF.SetLineno(435)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "mapping", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmapping *πg.Object = πArgs[1]; _ = µmapping
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 436: self._mapping = mapping
							πF.SetLineno(436)
							if πE = πg.CheckLocal(πF, µmapping, "mapping"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmapping); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_mapping, πTemp001); πE != nil {
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
					// line 438: def __len__(self):
					πF.SetLineno(438)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 439: return len(self._mapping)
							πF.SetLineno(439)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_mapping, nil); πE != nil {
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
					// line 441: def __repr__(self):
					πF.SetLineno(441)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 443: return '%s(%r)' % (self.__class__.__name__, self._mapping)
							πF.SetLineno(443)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_mapping, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, πTemp003).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s(%r)").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp004); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("MappingView").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMappingView.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 496: class MutableMapping(Mapping):
			πF.SetLineno(496)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßMapping); πE != nil {
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
			_, πE = πg.NewCode("MutableMapping", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 498: """A MutableMapping is a generic container for associating
					πF.SetLineno(498)
					// line 498: """A MutableMapping is a generic container for associating
					πF.SetLineno(498)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A MutableMapping is a generic container for associating\n    key/value pairs.\n\n    This class provides concrete generic implementations of all\n    methods except for __getitem__, __setitem__, __delitem__,\n    __iter__, and __len__.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 508: def __setitem__(self, key, value):
					πF.SetLineno(508)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							// line 509: raise KeyError
							πF.SetLineno(509)
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 507: @abstractmethod
					πF.SetLineno(507)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__setitem__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 512: def __delitem__(self, key):
					πF.SetLineno(512)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							// line 513: raise KeyError
							πF.SetLineno(513)
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
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 511: @abstractmethod
					πF.SetLineno(511)
					πTemp003 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__delitem__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 515: __marker = object()
					πF.SetLineno(515)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßobject); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__marker.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 517: def pop(self, key, default=__marker):
					πF.SetLineno(517)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß__marker); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp006}
					πTemp005 = πg.NewFunction(πg.NewCode("pop", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 518: '''D.pop(k[,d]) -> v, remove specified key and return the corresponding value.
							πF.SetLineno(518)
							// line 521: try:
							πF.SetLineno(521)
							πF.PushCheckpoint(2)
							// line 522: value = self[key]
							πF.SetLineno(522)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							µvalue = πTemp002
							πF.PopCheckpoint()
							// line 528: del self[key]
							πF.SetLineno(528)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.DelItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							// line 529: return value
							πF.SetLineno(529)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πR = µvalue
							continue
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp003, πTemp004 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
							// line 523: except KeyError:
							πF.SetLineno(523)
						Label3:
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__marker, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdefault == πTemp002).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 524: if default is self.__marker:
							πF.SetLineno(524)
						Label4:
							// line 525: raise
							πF.SetLineno(525)
							πE = πF.Raise(nil, nil, nil)
							continue
							goto Label5
						Label5:
							// line 526: return default
							πF.SetLineno(526)
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
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 518: '''D.pop(k[,d]) -> v, remove specified key and return the corresponding value.
					πF.SetLineno(518)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("D.pop(k[,d]) -> v, remove specified key and return the corresponding value.\n          If key is not found, d is returned if given, otherwise KeyError is raised.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßpop); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 531: def popitem(self):
					πF.SetLineno(531)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("popitem", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 532: '''D.popitem() -> (k, v), remove and return some (key, value) pair
							πF.SetLineno(532)
							// line 535: try:
							πF.SetLineno(535)
							πF.PushCheckpoint(2)
							// line 536: key = next(iter(self))
							πF.SetLineno(536)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µkey = πTemp004
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 537: except StopIteration:
							πF.SetLineno(537)
						Label3:
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							// line 538: raise KeyError
							πF.SetLineno(538)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 539: value = self[key]
							πF.SetLineno(539)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µself, πTemp003); πE != nil {
								continue
							}
							µvalue = πTemp004
							// line 540: del self[key]
							πF.SetLineno(540)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.DelItem(πF, µself, πTemp003); πE != nil {
								continue
							}
							// line 541: return key, value
							πF.SetLineno(541)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µkey, µvalue).ToObject()
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
					if πE = πClass.SetItem(πF, ßpopitem.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 532: '''D.popitem() -> (k, v), remove and return some (key, value) pair
					πF.SetLineno(532)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("D.popitem() -> (k, v), remove and return some (key, value) pair\n           as a 2-tuple; but raise KeyError if D is empty.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßpopitem); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 543: def clear(self):
					πF.SetLineno(543)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("clear", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 bool
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
							// line 544: 'D.clear() -> None.  Remove all items from D.'
							πF.SetLineno(544)
							// line 545: try:
							πF.SetLineno(545)
							πF.PushCheckpoint(2)
							// line 546: while True:
							πF.SetLineno(546)
							πF.PushCheckpoint(4)
							πTemp001 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp001 {
								πF.PopCheckpoint()
								goto Label5
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 547: self.popitem()
							πF.SetLineno(547)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
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
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label6
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 548: except KeyError:
							πF.SetLineno(548)
						Label6:
							// line 549: pass
							πF.SetLineno(549)
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
					if πE = πClass.SetItem(πF, ßclear.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 544: 'D.clear() -> None.  Remove all items from D.'
					πF.SetLineno(544)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("D.clear() -> None.  Remove all items from D.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßclear); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 551: def update(*args, **kwds):
					πF.SetLineno(551)
					πTemp002 = make([]πg.Param, 0)
					πTemp008 = πg.NewFunction(πg.NewCode("update", "/usr/lib/python2.7/_abcoll.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µother *πg.Object = πg.UnboundLocal; _ = µother
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 11: goto Label11
							case 12: goto Label12
							case 14: goto Label14
							case 15: goto Label15
							case 17: goto Label17
							case 18: goto Label18
							case 20: goto Label20
							case 21: goto Label21
							default: panic("unexpected function state")
							}
							// line 552: ''' D.update([E, ]**F) -> None.  Update D from mapping/iterable E and F.
							πF.SetLineno(552)
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
							// line 557: if not args:
							πF.SetLineno(557)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("descriptor 'update' of 'MutableMapping' object needs an argument").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 558: raise TypeError("descriptor 'update' of 'MutableMapping' object "
							πF.SetLineno(558)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 560: self = args[0]
							πF.SetLineno(560)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µself = πTemp004
							// line 561: args = args[1:]
							πF.SetLineno(561)
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
							// line 562: if len(args) > 1:
							πF.SetLineno(562)
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
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("update expected at most 1 arguments, got %d").ToObject(), πTemp005); πE != nil {
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
							// line 563: raise TypeError('update expected at most 1 arguments, got %d' %
							πF.SetLineno(563)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 565: if args:
							πF.SetLineno(565)
						Label5:
							// line 566: other = args[0]
							πF.SetLineno(566)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µother = πTemp004
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp003[0] = µother
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMapping); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label7
							}
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp003[0] = µother
							πTemp003[1] = ßkeys.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
								goto Label8
							}
							goto Label9
							// line 567: if isinstance(other, Mapping):
							πF.SetLineno(567)
						Label7:
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µother); πE != nil {
								continue
							}
							πF.PushCheckpoint(12)
							πTemp002 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label13
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
								µkey = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(11)            
							// line 569: self[key] = other[key]
							πF.SetLineno(569)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004 = µkey
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µother, πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp008 = µkey
							if πE = πg.SetItem(πF, µself, πTemp008, πTemp004); πE != nil {
								continue
							}
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							goto Label10
							// line 570: elif hasattr(other, "keys"):
							πF.SetLineno(570)
						Label8:
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µother, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(15)
							πTemp002 = false
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label16
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
								µkey = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(14)            
							// line 572: self[key] = other[key]
							πF.SetLineno(572)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004 = µkey
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µother, πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp008 = µkey
							if πE = πg.SetItem(πF, µself, πTemp008, πTemp004); πE != nil {
								continue
							}
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
						Label16:
							goto Label10
						Label9:
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µother); πE != nil {
								continue
							}
							πF.PushCheckpoint(18)
							πTemp002 = false
						Label17:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label19
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
									continue
								}
								µkey = πTemp005
								µvalue = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(17)            
							// line 575: self[key] = value
							πF.SetLineno(575)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp005 = µkey
							if πE = πg.SetItem(πF, µself, πTemp005, πTemp004); πE != nil {
								continue
							}
							continue
						Label18:
							if πE != nil || πR != nil {
								continue
							}
						Label19:
							goto Label10
						Label10:
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µkwds, ßitems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(21)
							πTemp002 = false
						Label20:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label22
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
									continue
								}
								µkey = πTemp005
								µvalue = πTemp008
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(20)            
							// line 577: self[key] = value
							πF.SetLineno(577)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp005 = µkey
							if πE = πg.SetItem(πF, µself, πTemp005, πTemp004); πE != nil {
								continue
							}
							continue
						Label21:
							if πE != nil || πR != nil {
								continue
							}
						Label22:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßupdate.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 552: ''' D.update([E, ]**F) -> None.  Update D from mapping/iterable E and F.
					πF.SetLineno(552)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr(" D.update([E, ]**F) -> None.  Update D from mapping/iterable E and F.\n            If E present and has a .keys() method, does:     for k in E: D[k] = E[k]\n            If E present and lacks .keys() method, does:     for (k, v) in E: D[k] = v\n            In either case, this is followed by: for k, v in F.items(): D[k] = v\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßupdate); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 579: def setdefault(self, key, default=None):
					πF.SetLineno(579)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "key", Def: nil}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "default", Def: πTemp010}
					πTemp009 = πg.NewFunction(πg.NewCode("setdefault", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µkey *πg.Object = πArgs[1]; _ = µkey
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 580: 'D.setdefault(k[,d]) -> D.get(k,d), also set D[k]=d if k not in D'
							πF.SetLineno(580)
							// line 581: try:
							πF.SetLineno(581)
							πF.PushCheckpoint(2)
							// line 582: return self[key]
							πF.SetLineno(582)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
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
							πTemp003, πTemp004 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
							// line 583: except KeyError:
							πF.SetLineno(583)
						Label3:
							// line 584: self[key] = default
							πF.SetLineno(584)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdefault); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.SetItem(πF, µself, πTemp002, πTemp001); πE != nil {
								continue
							}
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 585: return default
							πF.SetLineno(585)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πR = µdefault
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsetdefault.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 580: 'D.setdefault(k[,d]) -> D.get(k,d), also set D[k]=d if k not in D'
					πF.SetLineno(580)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("D.setdefault(k[,d]) -> D.get(k,d), also set D[k]=d if k not in D").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßsetdefault); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("MutableMapping").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMutableMapping.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 587: MutableMapping.register(dict)
			πF.SetLineno(587)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßMutableMapping); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßregister, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 593: class Sequence(Sized, Iterable, Container):
			πF.SetLineno(593)
			πTemp002 = make([]*πg.Object, 3)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßSized); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßIterable); πE != nil {
				continue
			}
			πTemp002[1] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßContainer); πE != nil {
				continue
			}
			πTemp002[2] = πTemp007
			πTemp005 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Sequence", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
					// line 594: """All the operations on a read-only sequence.
					πF.SetLineno(594)
					// line 594: """All the operations on a read-only sequence.
					πF.SetLineno(594)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("All the operations on a read-only sequence.\n\n    Concrete subclasses must override __new__ or __init__,\n    __getitem__, and __len__.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 601: def __getitem__(self, index):
					πF.SetLineno(601)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							// line 602: raise IndexError
							πF.SetLineno(602)
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 600: @abstractmethod
					πF.SetLineno(600)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__getitem__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 604: def __iter__(self):
					πF.SetLineno(604)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µv *πg.Object = πg.UnboundLocal; _ = µv
						var πTemp001 bool
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 2: goto Label2
								case 3: goto Label3
								case 4: goto Label4
								case 6: goto Label6
								default: panic("unexpected function state")
								}
								// line 605: i = 0
								πF.SetLineno(605)
								µi = πg.NewInt(0).ToObject()
								// line 606: try:
								πF.SetLineno(606)
								πF.PushCheckpoint(2)
								// line 607: while True:
								πF.SetLineno(607)
								πF.PushCheckpoint(4)
								πTemp001 = false
							Label3:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp001 {
									πF.PopCheckpoint()
									goto Label5
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πE != nil || !πTemp002 {
									continue
								}
								πF.PushCheckpoint(3)            
								// line 608: v = self[i]
								πF.SetLineno(608)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp003 = µi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetItem(πF, µself, πTemp003); πE != nil {
									continue
								}
								µv = πTemp004
								// line 609: yield v
								πF.SetLineno(609)
								if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µv, nil
							Label6:
								πTemp003 = πSent
								// line 610: i += 1
								πF.SetLineno(610)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µi = πTemp003
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
								πTemp005, πTemp006 = πF.ExcInfo()
								if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
									continue
								}
								if πTemp001, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp003); πE != nil {
									continue
								}
								if πTemp001 {
									goto Label7
								}
								πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
								continue
								// line 611: except IndexError:
								πF.SetLineno(611)
							Label7:
								// line 612: return
								πF.SetLineno(612)
								πR = πg.None
								continue
								πF.RestoreExc(nil, nil)
								goto Label1
							Label1:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 614: def __contains__(self, value):
					πF.SetLineno(614)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__contains__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var µv *πg.Object = πg.UnboundLocal; _ = µv
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
								µv = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µv, µvalue); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 616: if v == value:
							πF.SetLineno(616)
						Label4:
							// line 617: return True
							πF.SetLineno(617)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp004
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 618: return False
							πF.SetLineno(618)
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
					if πE = πClass.SetItem(πF, ß__contains__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 620: def __reversed__(self):
					πF.SetLineno(620)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__reversed__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
						var πTemp008 bool
						_ = πTemp008
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
								πTemp003 = πF.MakeArgs(1)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								πTemp004[0] = µself
								if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πTemp003[0] = πTemp006
								if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								πTemp002[0] = πTemp006
								if πTemp005, πE = πg.ResolveGlobal(πF, ßreversed); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp001, πE = πg.Iter(πF, πTemp006); πE != nil {
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
								if πTemp005, πE = πg.Next(πF, πTemp001); πE != nil {
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
								// line 622: yield self[i]
								πF.SetLineno(622)
								if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
									continue
								}
								πTemp005 = µi
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetItem(πF, µself, πTemp005); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp006, nil
							Label4:
								πTemp005 = πSent
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
					if πE = πClass.SetItem(πF, ß__reversed__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 624: def index(self, value):
					πF.SetLineno(624)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("index", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µv *πg.Object = πg.UnboundLocal; _ = µv
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
							// line 625: '''S.index(value) -> integer -- return first index of value.
							πF.SetLineno(625)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
									continue
								}
								µi = πTemp004
								µv = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µv, µvalue); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 629: if v == value:
							πF.SetLineno(629)
						Label4:
							// line 630: return i
							πF.SetLineno(630)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πR = µi
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 631: raise ValueError
							πF.SetLineno(631)
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
					if πE = πClass.SetItem(πF, ßindex.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 625: '''S.index(value) -> integer -- return first index of value.
					πF.SetLineno(625)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("S.index(value) -> integer -- return first index of value.\n           Raises ValueError if the value is not present.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßindex); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 633: def count(self, value):
					πF.SetLineno(633)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("count", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 634: 'S.count(value) -> integer -- return number of occurrences of value'
							πF.SetLineno(634)
							// line 635: return sum(1 for v in self if v == value)
							πF.SetLineno(635)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_abcoll.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µv *πg.Object = πg.UnboundLocal; _ = µv
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
										case 6: goto Label6
										default: panic("unexpected function state")
										}
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
											µv = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Eq(πF, µv, µvalue); πE != nil {
											continue
										}
										if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 635: return sum(1 for v in self if v == value)
										πF.SetLineno(635)
									Label4:
										// line 635: return sum(1 for v in self if v == value)
										πF.SetLineno(635)
										πF.PushCheckpoint(6)
										return πg.NewInt(1).ToObject(), nil
									Label6:
										πTemp004 = πSent
										goto Label5
									Label5:
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsum); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp005
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcount.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 634: 'S.count(value) -> integer -- return number of occurrences of value'
					πF.SetLineno(634)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("S.count(value) -> integer -- return number of occurrences of value").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßcount); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Sequence").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSequence.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 637: Sequence.register(tuple)
			πF.SetLineno(637)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSequence); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßregister, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 638: Sequence.register(basestring)
			πF.SetLineno(638)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSequence); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßregister, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 640: Sequence.register(xrange)
			πF.SetLineno(640)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßSequence); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßregister, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 643: class MutableSequence(Sequence):
			πF.SetLineno(643)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßSequence); πE != nil {
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
			_, πE = πg.NewCode("MutableSequence", "/usr/lib/python2.7/_abcoll.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp005
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 645: """All the operations on a read-only sequence.
					πF.SetLineno(645)
					// line 645: """All the operations on a read-only sequence.
					πF.SetLineno(645)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("All the operations on a read-only sequence.\n\n    Concrete subclasses must provide __new__ or __init__,\n    __getitem__, __setitem__, __delitem__, __len__, and insert().\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 653: def __setitem__(self, index, value):
					πF.SetLineno(653)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							// line 654: raise IndexError
							πF.SetLineno(654)
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 652: @abstractmethod
					πF.SetLineno(652)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__setitem__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 657: def __delitem__(self, index):
					πF.SetLineno(657)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							// line 658: raise IndexError
							πF.SetLineno(658)
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
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 656: @abstractmethod
					πF.SetLineno(656)
					πTemp003 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß__delitem__); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 661: def insert(self, index, value):
					πF.SetLineno(661)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("insert", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 662: 'S.insert(index, object) -- insert object before index'
							πF.SetLineno(662)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							// line 663: raise IndexError
							πF.SetLineno(663)
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
					if πE = πClass.SetItem(πF, ßinsert.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 662: 'S.insert(index, object) -- insert object before index'
					πF.SetLineno(662)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("S.insert(index, object) -- insert object before index").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßinsert); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 660: @abstractmethod
					πF.SetLineno(660)
					πTemp003 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßinsert); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßabstractmethod); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßinsert.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 665: def append(self, value):
					πF.SetLineno(665)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("append", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
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
							// line 666: 'S.append(object) -- append object to the end of the sequence'
							πF.SetLineno(666)
							// line 667: self.insert(len(self), value)
							πF.SetLineno(667)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßinsert, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßappend.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 666: 'S.append(object) -- append object to the end of the sequence'
					πF.SetLineno(666)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("S.append(object) -- append object to the end of the sequence").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßappend); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 669: def reverse(self):
					πF.SetLineno(669)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("reverse", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
							// line 670: 'S.reverse() -- reverse *IN PLACE*'
							πF.SetLineno(670)
							// line 671: n = len(self)
							πF.SetLineno(671)
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
							µn = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.FloorDiv(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
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
								µi = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 673: self[i], self[n-i-1] = self[n-i-1], self[i]
							πF.SetLineno(673)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Sub(πF, µn, µi); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Sub(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µself, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp008 = µi
							if πE = πg.SetItem(πF, µself, πTemp008, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Sub(πF, µn, µi); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Sub(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp008
							if πE = πg.SetItem(πF, µself, πTemp004, πTemp007); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreverse.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 670: 'S.reverse() -- reverse *IN PLACE*'
					πF.SetLineno(670)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("S.reverse() -- reverse *IN PLACE*").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßreverse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 675: def extend(self, values):
					πF.SetLineno(675)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "values", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("extend", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalues *πg.Object = πArgs[1]; _ = µvalues
						var µv *πg.Object = πg.UnboundLocal; _ = µv
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 676: 'S.extend(iterable) -- extend sequence by appending elements from the iterable'
							πF.SetLineno(676)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µvalues); πE != nil {
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
								µv = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 678: self.append(v)
							πF.SetLineno(678)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp005[0] = µv
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßextend.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 676: 'S.extend(iterable) -- extend sequence by appending elements from the iterable'
					πF.SetLineno(676)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("S.extend(iterable) -- extend sequence by appending elements from the iterable").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßextend); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 680: def pop(self, index=-1):
					πF.SetLineno(680)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp010, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "index", Def: πTemp010}
					πTemp009 = πg.NewFunction(πg.NewCode("pop", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var µv *πg.Object = πg.UnboundLocal; _ = µv
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
							// line 681: '''S.pop([index]) -> item -- remove and return item at index (default last).
							πF.SetLineno(681)
							// line 684: v = self[index]
							πF.SetLineno(684)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001 = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							µv = πTemp002
							// line 685: del self[index]
							πF.SetLineno(685)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001 = µindex
							if πE = πg.DelItem(πF, µself, πTemp001); πE != nil {
								continue
							}
							// line 686: return v
							πF.SetLineno(686)
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
					if πE = πClass.SetItem(πF, ßpop.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 681: '''S.pop([index]) -> item -- remove and return item at index (default last).
					πF.SetLineno(681)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("S.pop([index]) -> item -- remove and return item at index (default last).\n           Raise IndexError if list is empty or index is out of range.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßpop); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 688: def remove(self, value):
					πF.SetLineno(688)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("remove", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalue *πg.Object = πArgs[1]; _ = µvalue
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
							// line 689: '''S.remove(value) -- remove first occurrence of value.
							πF.SetLineno(689)
							// line 692: del self[self.index(value)]
							πF.SetLineno(692)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πTemp004
							if πE = πg.DelItem(πF, µself, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßremove.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 689: '''S.remove(value) -- remove first occurrence of value.
					πF.SetLineno(689)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("S.remove(value) -- remove first occurrence of value.\n           Raise ValueError if the value is not present.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßremove); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 694: def __iadd__(self, values):
					πF.SetLineno(694)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "values", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("__iadd__", "/usr/lib/python2.7/_abcoll.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalues *πg.Object = πArgs[1]; _ = µvalues
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
							// line 695: self.extend(values)
							πF.SetLineno(695)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp001[0] = µvalues
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
							// line 696: return self
							πF.SetLineno(696)
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
					if πE = πClass.SetItem(πF, ß__iadd__.ToObject(), πTemp011); πE != nil {
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
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("MutableSequence").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMutableSequence.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 698: MutableSequence.register(list)
			πF.SetLineno(698)
			πTemp002 = πF.MakeArgs(1)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßMutableSequence); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßregister, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
		}
		return nil, πE
	})
	πg.RegisterModule("_abcoll", Code)
}

