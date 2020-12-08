package mapping_tests
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/UserDict"
	// _ "github.com/pygolin/stdlib/unittest"
	// _ "github.com/pygolin/stdlib/test/test_support"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßAttributeError := πg.InternStr("AttributeError")
		ßBasicTestMappingProtocol := πg.InternStr("BasicTestMappingProtocol")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßKeyError := πg.InternStr("KeyError")
		ßNone := πg.InternStr("None")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTestCase := πg.InternStr("TestCase")
		ßTestHashMappingProtocol := πg.InternStr("TestHashMappingProtocol")
		ßTestMappingProtocol := πg.InternStr("TestMappingProtocol")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUserDict := πg.InternStr("UserDict")
		ßValueError := πg.InternStr("ValueError")
		ß__class__ := πg.InternStr("__class__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß_empty_mapping := πg.InternStr("_empty_mapping")
		ß_full_mapping := πg.InternStr("_full_mapping")
		ß_reference := πg.InternStr("_reference")
		ßa := πg.InternStr("a")
		ßab := πg.InternStr("ab")
		ßabc := πg.InternStr("abc")
		ßanything := πg.InternStr("anything")
		ßappend := πg.InternStr("append")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertIn := πg.InternStr("assertIn")
		ßassertIsInstance := πg.InternStr("assertIsInstance")
		ßassertNotIn := πg.InternStr("assertNotIn")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßb := πg.InternStr("b")
		ßbool := πg.InternStr("bool")
		ßc := πg.InternStr("c")
		ßcheck_py3k_warnings := πg.InternStr("check_py3k_warnings")
		ßchr := πg.InternStr("chr")
		ßclear := πg.InternStr("clear")
		ßcmp := πg.InternStr("cmp")
		ßcopy := πg.InternStr("copy")
		ßd := πg.InternStr("d")
		ßdef := πg.InternStr("def")
		ßdict := πg.InternStr("dict")
		ßfail := πg.InternStr("fail")
		ßfromkeys := πg.InternStr("fromkeys")
		ßget := πg.InternStr("get")
		ßgrumpy := πg.InternStr("grumpy")
		ßhas_key := πg.InternStr("has_key")
		ßhasattr := πg.InternStr("hasattr")
		ßi := πg.InternStr("i")
		ßid := πg.InternStr("id")
		ßinmapping := πg.InternStr("inmapping")
		ßint := πg.InternStr("int")
		ßitems := πg.InternStr("items")
		ßiter := πg.InternStr("iter")
		ßiteritems := πg.InternStr("iteritems")
		ßiterkeys := πg.InternStr("iterkeys")
		ßitervalues := πg.InternStr("itervalues")
		ßkey := πg.InternStr("key")
		ßkey0 := πg.InternStr("key0")
		ßkey1 := πg.InternStr("key1")
		ßkey2 := πg.InternStr("key2")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßother := πg.InternStr("other")
		ßpop := πg.InternStr("pop")
		ßpopitem := πg.InternStr("popitem")
		ßrange := πg.InternStr("range")
		ßreference := πg.InternStr("reference")
		ßrepr := πg.InternStr("repr")
		ßset := πg.InternStr("set")
		ßsetdefault := πg.InternStr("setdefault")
		ßskip := πg.InternStr("skip")
		ßsort := πg.InternStr("sort")
		ßtest_bool := πg.InternStr("test_bool")
		ßtest_clear := πg.InternStr("test_clear")
		ßtest_constructor := πg.InternStr("test_constructor")
		ßtest_contains := πg.InternStr("test_contains")
		ßtest_copy := πg.InternStr("test_copy")
		ßtest_fromkeys := πg.InternStr("test_fromkeys")
		ßtest_get := πg.InternStr("test_get")
		ßtest_getitem := πg.InternStr("test_getitem")
		ßtest_has_key := πg.InternStr("test_has_key")
		ßtest_items := πg.InternStr("test_items")
		ßtest_keys := πg.InternStr("test_keys")
		ßtest_le := πg.InternStr("test_le")
		ßtest_len := πg.InternStr("test_len")
		ßtest_mutatingiteration := πg.InternStr("test_mutatingiteration")
		ßtest_pop := πg.InternStr("test_pop")
		ßtest_popitem := πg.InternStr("test_popitem")
		ßtest_read := πg.InternStr("test_read")
		ßtest_repr := πg.InternStr("test_repr")
		ßtest_setdefault := πg.InternStr("test_setdefault")
		ßtest_support := πg.InternStr("test_support")
		ßtest_update := πg.InternStr("test_update")
		ßtest_values := πg.InternStr("test_values")
		ßtest_write := πg.InternStr("test_write")
		ßtype := πg.InternStr("type")
		ßtype2test := πg.InternStr("type2test")
		ßunittest := πg.InternStr("unittest")
		ßupdate := πg.InternStr("update")
		ßvalue1 := πg.InternStr("value1")
		ßvalues := πg.InternStr("values")
		ßx := πg.InternStr("x")
		ßy := πg.InternStr("y")
		ßz := πg.InternStr("z")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Dict
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
			// line 2: import unittest
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: import UserDict
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "UserDict"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßUserDict.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: from test import test_support
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: class BasicTestMappingProtocol(unittest.TestCase):
			πF.SetLineno(7)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßTestCase, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BasicTestMappingProtocol", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 13: type2test = None # which class is being tested (overwrite in subclasses)
					πF.SetLineno(13)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 15: def _reference(self):
					πF.SetLineno(15)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("_reference", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Dict
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
							// line 16: """Return a dictionary of values which are invariant by storage
							πF.SetLineno(16)
							// line 18: return {1:2, "key1":"value1", "key2":(1,2,3)}
							πF.SetLineno(18)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ßkey1.ToObject(), ßvalue1.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							if πE = πTemp001.SetItem(πF, ßkey2.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
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
					if πE = πClass.SetItem(πF, ß_reference.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 16: """Return a dictionary of values which are invariant by storage
					πF.SetLineno(16)
					// line 16: """Return a dictionary of values which are invariant by storage
					πF.SetLineno(16)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Return a dictionary of values which are invariant by storage\n        in the object under test.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Return a dictionary of values which are invariant by storage\n        in the object under test.").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_reference); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 19: def _empty_mapping(self):
					πF.SetLineno(19)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_empty_mapping", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 20: """Return an empty mapping object"""
							πF.SetLineno(20)
							// line 21: return self.type2test()
							πF.SetLineno(21)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_empty_mapping.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 20: """Return an empty mapping object"""
					πF.SetLineno(20)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Return an empty mapping object").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß_empty_mapping); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 22: def _full_mapping(self, data):
					πF.SetLineno(22)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_full_mapping", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdata *πg.Object = πArgs[1]; _ = µdata
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
							// line 23: """Return a mapping object with the value contained in data
							πF.SetLineno(23)
							// line 25: x = self._empty_mapping()
							πF.SetLineno(25)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp002
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdata, ßitems, nil); πE != nil {
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
							// line 27: x[key] = value
							πF.SetLineno(27)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.SetItem(πF, µx, πTemp003, πTemp002); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 28: return x
							πF.SetLineno(28)
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
					if πE = πClass.SetItem(πF, ß_full_mapping.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 23: """Return a mapping object with the value contained in data
					πF.SetLineno(23)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return a mapping object with the value contained in data\n        dictionary").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß_full_mapping); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 30: def __init__(self, *args, **kw):
					πF.SetLineno(30)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkw *πg.Object = πArgs[2]; _ = µkw
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 31: unittest.TestCase.__init__(self, *args, **kw)
							πF.SetLineno(31)
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunittest); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßTestCase, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, µargs, nil, µkw); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 32: self.reference = self._reference().copy()
							πF.SetLineno(32)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_reference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßcopy, nil); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßreference, πTemp002); πE != nil {
								continue
							}
							// line 35: key, value = self.reference.popitem()
							πF.SetLineno(35)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							µkey = πTemp003
							µvalue = πTemp004
							// line 36: self.other = {key:value}
							πF.SetLineno(36)
							πTemp005 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, µkey, µvalue); πE != nil {
								continue
							}
							πTemp002 = πTemp005.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßother, πTemp003); πE != nil {
								continue
							}
							// line 39: key, value = self.reference.popitem()
							πF.SetLineno(39)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							µkey = πTemp003
							µvalue = πTemp004
							// line 40: self.inmapping = {key:value}
							πF.SetLineno(40)
							πTemp005 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, µkey, µvalue); πE != nil {
								continue
							}
							πTemp002 = πTemp005.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinmapping, πTemp003); πE != nil {
								continue
							}
							// line 41: self.reference[key] = value
							πF.SetLineno(41)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004 = µkey
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 44: def test_read(self):
					πF.SetLineno(44)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_read", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µp1 *πg.Object = πg.UnboundLocal; _ = µp1
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µknownkey *πg.Object = πg.UnboundLocal; _ = µknownkey
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var µcheck_iterandlist *πg.Object = πg.UnboundLocal; _ = µcheck_iterandlist
						var µknownvalue *πg.Object = πg.UnboundLocal; _ = µknownvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []πg.Param
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.BaseException
						_ = πTemp013
						var πTemp014 *πg.Traceback
						_ = πTemp014
						var πTemp015 *πg.Type
						_ = πTemp015
						var πTemp016 *πg.Object
						_ = πTemp016
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							case 6: goto Label6
							case 7: goto Label7
							case 9: goto Label9
							case 10: goto Label10
							case 12: goto Label12
							case 13: goto Label13
							case 14: goto Label14
							case 16: goto Label16
							case 17: goto Label17
							default: panic("unexpected function state")
							}
							// line 46: p = self._empty_mapping()
							πF.SetLineno(46)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µp = πTemp002
							// line 47: p1 = dict(p) #workaround for singleton objects
							πF.SetLineno(47)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp003[0] = µp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µp1 = πTemp002
							// line 48: d = self._full_mapping(self.reference)
							πF.SetLineno(48)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µd = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µd == µp).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 49: if d is p:
							πF.SetLineno(49)
						Label1:
							// line 50: p = p1
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µp1, "p1"); πE != nil {
								continue
							}
							µp = µp1
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
									continue
								}
								µkey = πTemp005
								µvalue = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 53: self.assertEqual(d[key], value)
							πF.SetLineno(53)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[1] = µvalue
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
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 54: knownkey = self.other.keys()[0]
							πF.SetLineno(54)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							µknownkey = πTemp002
							// line 55: self.assertRaises(KeyError, lambda:d[knownkey])
							πF.SetLineno(55)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp008 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/test/mapping_tests.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 55: self.assertRaises(KeyError, lambda:d[knownkey])
									πF.SetLineno(55)
									if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
										continue
									}
									πTemp001 = µknownkey
									if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetItem(πF, µd, πTemp001); πE != nil {
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
							// line 57: self.assertEqual(len(p), 0)
							πF.SetLineno(57)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp009[0] = µp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(0).ToObject()
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
							// line 58: self.assertEqual(len(d), len(self.reference))
							πF.SetLineno(58)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp009[0] = µd
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp002
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[1] = πTemp002
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
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
								µk = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(6)            
							// line 61: self.assertIn(k, d)
							πF.SetLineno(61)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[0] = µk
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µk = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(9)            
							// line 63: self.assertNotIn(k, d)
							πF.SetLineno(63)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[0] = µk
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							// line 65: with test_support.check_py3k_warnings(quiet=True):
							πF.SetLineno(65)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"quiet", πTemp001},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcheck_py3k_warnings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, πTemp010); πE != nil {
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
							πF.PushCheckpoint(12)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Iter(πF, πTemp011); πE != nil {
								continue
							}
							πF.PushCheckpoint(14)
							πTemp004 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label15
							}
							if πTemp011, πE = πg.Next(πF, πTemp007); πE != nil {
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
								µk = πTemp011
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 67: self.assertTrue(d.has_key(k))
							πF.SetLineno(67)
							πTemp003 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp009[0] = µk
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßhas_key, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp012
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Iter(πF, πTemp011); πE != nil {
								continue
							}
							πF.PushCheckpoint(17)
							πTemp004 = false
						Label16:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label18
							}
							if πTemp011, πE = πg.Next(πF, πTemp007); πE != nil {
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
								µk = πTemp011
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(16)            
							// line 69: self.assertFalse(d.has_key(k))
							πF.SetLineno(69)
							πTemp003 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp009[0] = µk
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßhas_key, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp012
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label17:
							if πE != nil || πR != nil {
								continue
							}
						Label18:
							πF.PopCheckpoint()
						Label12:
							πTemp013, πTemp014 = nil, nil
							if πE != nil {
								πTemp013, πTemp014 = πF.ExcInfo()
							}
							if πTemp013 != nil {
								πTemp015 = πTemp013.Type()
								if πTemp007, πE = πTemp002.Call(πF, πg.Args{πTemp001, πTemp015.ToObject(), πTemp013.ToObject(), πTemp014.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp007, πE = πTemp002.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp013 != nil && πTemp004 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 71: self.assertEqual(cmp(p,p), 0)
							πF.SetLineno(71)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp009[0] = µp
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp009[1] = µp
							if πTemp007, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp011
							πTemp003[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 72: self.assertEqual(cmp(d,d), 0)
							πF.SetLineno(72)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp009[0] = µd
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp009[1] = µd
							if πTemp007, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp011
							πTemp003[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 73: self.assertEqual(cmp(p,d), -1)
							πF.SetLineno(73)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp009[0] = µp
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp009[1] = µd
							if πTemp007, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp011
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[1] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 74: self.assertEqual(cmp(d,p), 1)
							πF.SetLineno(74)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp009[0] = µd
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp009[1] = µp
							if πTemp007, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp011
							πTemp003[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µp); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label19
							}
							goto Label20
							// line 76: if p: self.fail("Empty mapping must compare to False")
							πF.SetLineno(76)
						Label19:
							// line 76: if p: self.fail("Empty mapping must compare to False")
							πF.SetLineno(76)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Empty mapping must compare to False").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label20
						Label20:
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µd); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label21
							}
							goto Label22
							// line 77: if not d: self.fail("Full mapping must compare to True")
							πF.SetLineno(77)
						Label21:
							// line 77: if not d: self.fail("Full mapping must compare to True")
							πF.SetLineno(77)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("Full mapping must compare to True").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label22
						Label22:
							// line 79: def check_iterandlist(iter, lst, ref):
							πF.SetLineno(79)
							πTemp008 = make([]πg.Param, 3)
							πTemp008[0] = πg.Param{Name: "iter", Def: nil}
							πTemp008[1] = πg.Param{Name: "lst", Def: nil}
							πTemp008[2] = πg.Param{Name: "ref", Def: nil}
							πTemp007 = πg.NewFunction(πg.NewCode("check_iterandlist", "/usr/lib/python2.7/test/mapping_tests.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µiter *πg.Object = πArgs[0]; _ = µiter
								var µlst *πg.Object = πArgs[1]; _ = µlst
								var µref *πg.Object = πArgs[2]; _ = µref
								var µx *πg.Object = πg.UnboundLocal; _ = µx
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
									// line 80: self.assertTrue(hasattr(iter, 'next'))
									πF.SetLineno(80)
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µiter, "iter"); πE != nil {
										continue
									}
									πTemp002[0] = µiter
									πTemp002[1] = ßnext.ToObject()
									if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
									if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 81: self.assertTrue(hasattr(iter, '__iter__'))
									πF.SetLineno(81)
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µiter, "iter"); πE != nil {
										continue
									}
									πTemp002[0] = µiter
									πTemp002[1] = ß__iter__.ToObject()
									if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
									if πTemp003, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 82: x = list(iter)
									πF.SetLineno(82)
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µiter, "iter"); πE != nil {
										continue
									}
									πTemp001[0] = µiter
									if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									µx = πTemp004
									// line 83: self.assertTrue(set(x)==set(lst)==set(ref))
									πF.SetLineno(83)
									πTemp001 = πF.MakeArgs(1)
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp002[0] = µx
									if πTemp004, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
										continue
									}
									πTemp002[0] = µlst
									if πTemp004, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp003, πE = πg.Eq(πF, πTemp005, πTemp006); πE != nil {
										continue
									}
									if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									if !πTemp007 {
										goto Label1
									}
									πTemp002 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µref, "ref"); πE != nil {
										continue
									}
									πTemp002[0] = µref
									if πTemp004, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp003, πE = πg.Eq(πF, πTemp006, πTemp005); πE != nil {
										continue
									}
								Label1:
									πTemp001[0] = πTemp003
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
							µcheck_iterandlist = πTemp007
							// line 84: check_iterandlist(d.iterkeys(), d.keys(), self.reference.keys())
							πF.SetLineno(84)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßiterkeys, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp012
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp012
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp011
							if πE = πg.CheckLocal(πF, µcheck_iterandlist, "check_iterandlist"); πE != nil {
								continue
							}
							if πTemp011, πE = µcheck_iterandlist.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 85: check_iterandlist(iter(d), d.keys(), self.reference.keys())
							πF.SetLineno(85)
							πTemp003 = πF.MakeArgs(3)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp009[0] = µd
							if πTemp011, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp012
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp012
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp011
							if πE = πg.CheckLocal(πF, µcheck_iterandlist, "check_iterandlist"); πE != nil {
								continue
							}
							if πTemp011, πE = µcheck_iterandlist.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 86: check_iterandlist(d.itervalues(), d.values(), self.reference.values())
							πF.SetLineno(86)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßitervalues, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp012
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp012
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp011
							if πE = πg.CheckLocal(πF, µcheck_iterandlist, "check_iterandlist"); πE != nil {
								continue
							}
							if πTemp011, πE = µcheck_iterandlist.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 87: check_iterandlist(d.iteritems(), d.items(), self.reference.items())
							πF.SetLineno(87)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp012
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp012
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßitems, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp011
							if πE = πg.CheckLocal(πF, µcheck_iterandlist, "check_iterandlist"); πE != nil {
								continue
							}
							if πTemp011, πE = µcheck_iterandlist.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 89: key, value = d.iteritems().next()
							πF.SetLineno(89)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp012, ßnext, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp011}, πg.TieTarget{Target: &πTemp016}}}, πTemp012); πE != nil {
								continue
							}
							µkey = πTemp011
							µvalue = πTemp016
							// line 90: knownkey, knownvalue = self.other.iteritems().next()
							πF.SetLineno(90)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßnext, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp012.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp012}, πg.TieTarget{Target: &πTemp016}}}, πTemp011); πE != nil {
								continue
							}
							µknownkey = πTemp012
							µknownvalue = πTemp016
							// line 91: self.assertEqual(d.get(key, knownvalue), value)
							πF.SetLineno(91)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp009[0] = µkey
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							πTemp009[1] = µknownvalue
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp012
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 92: self.assertEqual(d.get(knownkey, knownvalue), knownvalue)
							πF.SetLineno(92)
							πTemp003 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp009[0] = µknownkey
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							πTemp009[1] = µknownvalue
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp003[0] = πTemp012
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							πTemp003[1] = µknownvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 93: self.assertNotIn(knownkey, d)
							πF.SetLineno(93)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp003[0] = µknownkey
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_read.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 43: @unittest.skip('grumpy')
					πF.SetLineno(43)
					πTemp007 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßtest_read); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_read.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 95: def test_write(self):
					πF.SetLineno(95)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_write", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µp *πg.Object = πg.UnboundLocal; _ = µp
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µitems *πg.Object = πg.UnboundLocal; _ = µitems
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µknownkey *πg.Object = πg.UnboundLocal; _ = µknownkey
						var µknownvalue *πg.Object = πg.UnboundLocal; _ = µknownvalue
						var µdefault *πg.Object = πg.UnboundLocal; _ = µdefault
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []πg.Param
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
							// line 97: p = self._empty_mapping()
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µp = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
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
							// line 100: p[key] = value
							πF.SetLineno(100)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πE = πg.SetItem(πF, µp, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 101: self.assertEqual(p[key], value)
							πF.SetLineno(101)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µp, πTemp002); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								µkey = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 103: del p[key]
							πF.SetLineno(103)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.DelItem(πF, µp, πTemp002); πE != nil {
								continue
							}
							// line 104: self.assertRaises(KeyError, lambda:p[key])
							πF.SetLineno(104)
							πTemp007 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							πTemp008 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/test/mapping_tests.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 104: self.assertRaises(KeyError, lambda:p[key])
									πF.SetLineno(104)
									if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
										continue
									}
									πTemp001 = µkey
									if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.GetItem(πF, µp, πTemp001); πE != nil {
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
							πTemp007[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 105: p = self._empty_mapping()
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µp = πTemp002
							// line 107: p.update(self.reference)
							πF.SetLineno(107)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µp, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 108: self.assertEqual(dict(p), self.reference)
							πF.SetLineno(108)
							πTemp007 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp009[0] = µp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 109: items = p.items()
							πF.SetLineno(109)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µp, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µitems = πTemp002
							// line 110: p = self._empty_mapping()
							πF.SetLineno(110)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µp = πTemp002
							// line 111: p.update(items)
							πF.SetLineno(111)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp007[0] = µitems
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µp, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 112: self.assertEqual(dict(p), self.reference)
							πF.SetLineno(112)
							πTemp007 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp009[0] = µp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 113: d = self._full_mapping(self.reference)
							πF.SetLineno(113)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µd = πTemp002
							// line 115: key, value = d.iteritems().next()
							πF.SetLineno(115)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßnext, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µkey = πTemp001
							µvalue = πTemp003
							// line 116: knownkey, knownvalue = self.other.iteritems().next()
							πF.SetLineno(116)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnext, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
								continue
							}
							µknownkey = πTemp002
							µknownvalue = πTemp003
							// line 117: self.assertEqual(d.setdefault(key, knownvalue), value)
							πF.SetLineno(117)
							πTemp007 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp009[0] = µkey
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							πTemp009[1] = µknownvalue
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 118: self.assertEqual(d[key], value)
							πF.SetLineno(118)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µd, πTemp001); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 119: self.assertEqual(d.setdefault(knownkey, knownvalue), knownvalue)
							πF.SetLineno(119)
							πTemp007 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp009[0] = µknownkey
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							πTemp009[1] = µknownvalue
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							πTemp007[1] = µknownvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 120: self.assertEqual(d[knownkey], knownvalue)
							πF.SetLineno(120)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp001 = µknownkey
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µd, πTemp001); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							πTemp007[1] = µknownvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 122: self.assertEqual(d.pop(knownkey), knownvalue)
							πF.SetLineno(122)
							πTemp007 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp009[0] = µknownkey
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							πTemp007[1] = µknownvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 123: self.assertNotIn(knownkey, d)
							πF.SetLineno(123)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp007[0] = µknownkey
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp007[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 124: self.assertRaises(KeyError, d.pop, knownkey)
							πF.SetLineno(124)
							πTemp007 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp007[2] = µknownkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 125: default = 909
							πF.SetLineno(125)
							µdefault = πg.NewInt(909).ToObject()
							// line 126: d[knownkey] = knownvalue
							πF.SetLineno(126)
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µknownvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp002 = µknownkey
							if πE = πg.SetItem(πF, µd, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 127: self.assertEqual(d.pop(knownkey, default), knownvalue)
							πF.SetLineno(127)
							πTemp007 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp009[0] = µknownkey
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp009[1] = µdefault
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µknownvalue, "knownvalue"); πE != nil {
								continue
							}
							πTemp007[1] = µknownvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 128: self.assertNotIn(knownkey, d)
							πF.SetLineno(128)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp007[0] = µknownkey
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp007[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 129: self.assertEqual(d.pop(knownkey, default), default)
							πF.SetLineno(129)
							πTemp007 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µknownkey, "knownkey"); πE != nil {
								continue
							}
							πTemp009[0] = µknownkey
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp009[1] = µdefault
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp007[1] = µdefault
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 131: key, value = d.popitem()
							πF.SetLineno(131)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µkey = πTemp001
							µvalue = πTemp003
							// line 132: self.assertNotIn(key, d)
							πF.SetLineno(132)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007[0] = µkey
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp007[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 133: self.assertEqual(value, self.reference[key])
							πF.SetLineno(133)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007[0] = µvalue
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πTemp007[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 134: p=self._empty_mapping()
							πF.SetLineno(134)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µp = πTemp002
							// line 135: self.assertRaises(KeyError, p.popitem)
							πF.SetLineno(135)
							πTemp007 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µp, ßpopitem, nil); πE != nil {
								continue
							}
							πTemp007[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_write.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 137: def test_constructor(self):
					πF.SetLineno(137)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_constructor", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 138: self.assertEqual(self._empty_mapping(), self._empty_mapping())
							πF.SetLineno(138)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_constructor.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 140: def test_bool(self):
					πF.SetLineno(140)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_bool", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 141: self.assertTrue(not self._empty_mapping())
							πF.SetLineno(141)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							πTemp001[0] = πTemp002
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
							// line 142: self.assertTrue(self.reference)
							πF.SetLineno(142)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							// line 143: self.assertTrue(bool(self._empty_mapping()) is False)
							πF.SetLineno(143)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004 == πTemp003).ToObject()
							πTemp001[0] = πTemp002
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
							// line 144: self.assertTrue(bool(self.reference) is True)
							πF.SetLineno(144)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004 == πTemp003).ToObject()
							πTemp001[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_bool.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 146: def test_keys(self):
					πF.SetLineno(146)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_keys", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 147: d = self._empty_mapping()
							πF.SetLineno(147)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 148: self.assertEqual(d.keys(), [])
							πF.SetLineno(148)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
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
							// line 149: d = self.reference
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							µd = πTemp001
							// line 150: self.assertIn(self.inmapping.keys()[0], d.keys())
							πF.SetLineno(150)
							πTemp003 = πF.MakeArgs(2)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßinmapping, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 151: self.assertNotIn(self.other.keys()[0], d.keys())
							πF.SetLineno(151)
							πTemp003 = πF.MakeArgs(2)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 152: self.assertRaises(TypeError, d.keys, None)
							πF.SetLineno(152)
							πTemp003 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
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
					if πE = πClass.SetItem(πF, ßtest_keys.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 154: def test_values(self):
					πF.SetLineno(154)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_values", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							// line 155: d = self._empty_mapping()
							πF.SetLineno(155)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 156: self.assertEqual(d.values(), [])
							πF.SetLineno(156)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
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
							// line 158: self.assertRaises(TypeError, d.values, None)
							πF.SetLineno(158)
							πTemp003 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßvalues, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
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
					if πE = πClass.SetItem(πF, ßtest_values.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 160: def test_items(self):
					πF.SetLineno(160)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_items", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							// line 161: d = self._empty_mapping()
							πF.SetLineno(161)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 162: self.assertEqual(d.items(), [])
							πF.SetLineno(162)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
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
							// line 164: self.assertRaises(TypeError, d.items, None)
							πF.SetLineno(164)
							πTemp003 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
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
					if πE = πClass.SetItem(πF, ßtest_items.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 166: def test_len(self):
					πF.SetLineno(166)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_len", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							// line 167: d = self._empty_mapping()
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 168: self.assertEqual(len(d), 0)
							πF.SetLineno(168)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(0).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_len.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 170: def test_getitem(self):
					πF.SetLineno(170)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_getitem", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							// line 171: d = self.reference
							πF.SetLineno(171)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							µd = πTemp001
							// line 172: self.assertEqual(d[self.inmapping.keys()[0]], self.inmapping.values()[0])
							πF.SetLineno(172)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßinmapping, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µd, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßinmapping, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 174: self.assertRaises(TypeError, d.__getitem__)
							πF.SetLineno(174)
							πTemp002 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_getitem.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 177: def test_update(self):
					πF.SetLineno(177)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("test_update", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µouterself *πg.Object = πg.UnboundLocal; _ = µouterself
						var µSimpleUserDict *πg.Object = πg.UnboundLocal; _ = µSimpleUserDict
						var µi1 *πg.Object = πg.UnboundLocal; _ = µi1
						var µi2 *πg.Object = πg.UnboundLocal; _ = µi2
						var µsafe_sort_key *πg.Object = πg.UnboundLocal; _ = µsafe_sort_key
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µFailingUserDict *πg.Object = πg.UnboundLocal; _ = µFailingUserDict
						var µbadseq *πg.Object = πg.UnboundLocal; _ = µbadseq
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 []πg.Param
						_ = πTemp006
						var πTemp007 πg.KWArgs
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
							default: panic("unexpected function state")
							}
							// line 179: d = self._empty_mapping()
							πF.SetLineno(179)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 180: d.update(self.other)
							πF.SetLineno(180)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 181: self.assertEqual(d.items(), self.other.items())
							πF.SetLineno(181)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßitems, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
							// line 184: d = self._empty_mapping()
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 185: d.update()
							πF.SetLineno(185)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 186: self.assertEqual(d, self._empty_mapping())
							πF.SetLineno(186)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
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
							// line 189: d = self._empty_mapping()
							πF.SetLineno(189)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 190: d.update(self.other.items())
							πF.SetLineno(190)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßitems, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 191: self.assertEqual(d.items(), self.other.items())
							πF.SetLineno(191)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßitems, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
							// line 194: d = self._empty_mapping()
							πF.SetLineno(194)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 195: d.update(self.other.iteritems())
							πF.SetLineno(195)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 196: self.assertEqual(d.items(), self.other.items())
							πF.SetLineno(196)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßitems, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
							// line 200: self.assertRaises((TypeError, AttributeError), d.update, 42)
							πF.SetLineno(200)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp004).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = πg.NewInt(42).ToObject()
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
							// line 202: outerself = self
							πF.SetLineno(202)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							µouterself = µself
							// line 203: class SimpleUserDict(object):
							πF.SetLineno(203)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp005 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("SimpleUserDict", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 204: def __init__(self):
									πF.SetLineno(204)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 205: self.d = outerself.reference
											πF.SetLineno(205)
											if πE = πg.CheckLocal(πF, µouterself, "outerself"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µouterself, ßreference, nil); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßd, πTemp002); πE != nil {
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
									// line 206: def keys(self):
									πF.SetLineno(206)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 207: return self.d.keys()
											πF.SetLineno(207)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ßd, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp003); πE != nil {
										continue
									}
									// line 208: def __getitem__(self, i):
									πF.SetLineno(208)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "i", Def: nil}
									πTemp004 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 209: return self.d[i]
											πF.SetLineno(209)
											if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
												continue
											}
											πTemp001 = µi
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetAttr(πF, µself, ßd, nil); πE != nil {
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
							if πTemp002, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp002 == nil {
								πTemp002 = πg.TypeType.ToObject()
							}
							if πTemp004, πE = πTemp002.Call(πF, []*πg.Object{πg.NewStr("SimpleUserDict").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µSimpleUserDict = πTemp004
							// line 210: d.clear()
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 211: d.update(SimpleUserDict())
							πF.SetLineno(211)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µSimpleUserDict, "SimpleUserDict"); πE != nil {
								continue
							}
							if πTemp001, πE = µSimpleUserDict.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 212: i1 = d.items()
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µi1 = πTemp002
							// line 213: i2 = self.reference.items()
							πF.SetLineno(213)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßitems, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µi2 = πTemp001
							// line 215: def safe_sort_key(kv):
							πF.SetLineno(215)
							πTemp006 = make([]πg.Param, 1)
							πTemp006[0] = πg.Param{Name: "kv", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("safe_sort_key", "/usr/lib/python2.7/test/mapping_tests.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µkv *πg.Object = πArgs[0]; _ = µkv
								var µk *πg.Object = πg.UnboundLocal; _ = µk
								var µv *πg.Object = πg.UnboundLocal; _ = µv
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
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 216: k, v = kv
									πF.SetLineno(216)
									if πE = πg.CheckLocal(πF, µkv, "kv"); πE != nil {
										continue
									}
									if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µkv); πE != nil {
										continue
									}
									µk = πTemp001
									µv = πTemp002
									// line 217: return id(type(k)), id(type(v)), k, v
									πF.SetLineno(217)
									πTemp003 = πF.MakeArgs(1)
									πTemp004 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
										continue
									}
									πTemp004[0] = µk
									if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									πTemp003[0] = πTemp005
									if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									πTemp003 = πF.MakeArgs(1)
									πTemp004 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
										continue
									}
									πTemp004[0] = µv
									if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									πTemp003[0] = πTemp006
									if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
										continue
									}
									πTemp001 = πg.NewTuple4(πTemp005, πTemp006, µk, µv).ToObject()
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
							µsafe_sort_key = πTemp001
							// line 218: i1.sort(key=safe_sort_key)
							πF.SetLineno(218)
							if πE = πg.CheckLocal(πF, µsafe_sort_key, "safe_sort_key"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"key", µsafe_sort_key},
							}
							if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µi1, ßsort, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, πTemp007); πE != nil {
								continue
							}
							// line 219: i2.sort(key=safe_sort_key)
							πF.SetLineno(219)
							if πE = πg.CheckLocal(πF, µsafe_sort_key, "safe_sort_key"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"key", µsafe_sort_key},
							}
							if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µi2, ßsort, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, πTemp007); πE != nil {
								continue
							}
							// line 220: self.assertEqual(i1, i2)
							πF.SetLineno(220)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi1, "i1"); πE != nil {
								continue
							}
							πTemp003[0] = µi1
							if πE = πg.CheckLocal(πF, µi2, "i2"); πE != nil {
								continue
							}
							πTemp003[1] = µi2
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
							// line 222: class Exc(Exception): pass
							πF.SetLineno(222)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp005 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 222: class Exc(Exception): pass
									πF.SetLineno(222)
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
							if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp008
							// line 224: d = self._empty_mapping()
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp004
							// line 225: class FailingUserDict(object):
							πF.SetLineno(225)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp005 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("FailingUserDict", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 226: def keys(self):
									πF.SetLineno(226)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											// line 227: raise Exc
											πF.SetLineno(227)
											πE = πF.Raise(µExc, nil, nil)
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp001); πE != nil {
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
							if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FailingUserDict").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µFailingUserDict = πTemp008
							// line 228: self.assertRaises(Exc, d.update, FailingUserDict())
							πF.SetLineno(228)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp003[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µFailingUserDict, "FailingUserDict"); πE != nil {
								continue
							}
							if πTemp002, πE = µFailingUserDict.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp002
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
							// line 230: d.clear()
							πF.SetLineno(230)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 232: class FailingUserDict(object):
							πF.SetLineno(232)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp005 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("FailingUserDict", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 233: def keys(self):
									πF.SetLineno(233)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µBogonIter *πg.Object = πg.UnboundLocal; _ = µBogonIter
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
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 234: class BogonIter(object):
											πF.SetLineno(234)
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
											_, πE = πg.NewCode("BogonIter", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
													// line 235: def __init__(self):
													πF.SetLineno(235)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
															// line 236: self.i = 1
															πF.SetLineno(236)
															if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
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
													// line 237: def __iter__(self):
													πF.SetLineno(237)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
														var µself *πg.Object = πArgs[0]; _ = µself
														var πR *πg.Object; _ = πR
														var πE *πg.BaseException; _ = πE
														for ; πF.State() >= 0; πF.PopCheckpoint() {
															switch πF.State() {
															case 0:
															default: panic("unexpected function state")
															}
															// line 238: return self
															πF.SetLineno(238)
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
													// line 239: def next(self):
													πF.SetLineno(239)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp004 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
														var µself *πg.Object = πArgs[0]; _ = µself
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
															if πTemp001, πE = πg.GetAttr(πF, µself, ßi, nil); πE != nil {
																continue
															}
															if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
																continue
															}
															if πTemp002 {
																goto Label1
															}
															goto Label2
															// line 240: if self.i:
															πF.SetLineno(240)
														Label1:
															// line 241: self.i = 0
															πF.SetLineno(241)
															if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
																continue
															}
															if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
																continue
															}
															if πE = πg.SetAttr(πF, µself, ßi, πTemp001); πE != nil {
																continue
															}
															// line 242: return 'a'
															πF.SetLineno(242)
															πR = ßa.ToObject()
															continue
															goto Label2
														Label2:
															if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
																continue
															}
															// line 243: raise Exc
															πF.SetLineno(243)
															πE = πF.Raise(µExc, nil, nil)
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
											if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
												continue
											}
											if πTemp004 == nil {
												πTemp004 = πg.TypeType.ToObject()
											}
											if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BogonIter").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
												continue
											}
											µBogonIter = πTemp005
											// line 244: return BogonIter()
											πF.SetLineno(244)
											if πE = πg.CheckLocal(πF, µBogonIter, "BogonIter"); πE != nil {
												continue
											}
											if πTemp002, πE = µBogonIter.Call(πF, nil, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 245: def __getitem__(self, key):
									πF.SetLineno(245)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µkey *πg.Object = πArgs[1]; _ = µkey
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 246: return key
											πF.SetLineno(246)
											if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
												continue
											}
											πR = µkey
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
							if πTemp004, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FailingUserDict").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µFailingUserDict = πTemp008
							// line 247: self.assertRaises(Exc, d.update, FailingUserDict())
							πF.SetLineno(247)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp003[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µFailingUserDict, "FailingUserDict"); πE != nil {
								continue
							}
							if πTemp002, πE = µFailingUserDict.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp002
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
							// line 249: class FailingUserDict(object):
							πF.SetLineno(249)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp005 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("FailingUserDict", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 250: def keys(self):
									πF.SetLineno(250)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µBogonIter *πg.Object = πg.UnboundLocal; _ = µBogonIter
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
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 251: class BogonIter(object):
											πF.SetLineno(251)
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
											_, πE = πg.NewCode("BogonIter", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
													// line 252: def __init__(self):
													πF.SetLineno(252)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
															// line 253: self.i = ord('a')
															πF.SetLineno(253)
															πTemp001 = πF.MakeArgs(1)
															πTemp001[0] = ßa.ToObject()
															if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
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
															if πE = πg.SetAttr(πF, µself, ßi, πTemp002); πE != nil {
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
													// line 254: def __iter__(self):
													πF.SetLineno(254)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
														var µself *πg.Object = πArgs[0]; _ = µself
														var πR *πg.Object; _ = πR
														var πE *πg.BaseException; _ = πE
														for ; πF.State() >= 0; πF.PopCheckpoint() {
															switch πF.State() {
															case 0:
															default: panic("unexpected function state")
															}
															// line 255: return self
															πF.SetLineno(255)
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
													// line 256: def next(self):
													πF.SetLineno(256)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp004 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
														var µself *πg.Object = πArgs[0]; _ = µself
														var µrtn *πg.Object = πg.UnboundLocal; _ = µrtn
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
															πTemp003[0] = ßz.ToObject()
															if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
																continue
															}
															if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
																continue
															}
															πF.FreeArgs(πTemp003)
															if πTemp001, πE = πg.LE(πF, πTemp002, πTemp005); πE != nil {
																continue
															}
															if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
																continue
															}
															if πTemp006 {
																goto Label1
															}
															goto Label2
															// line 257: if self.i <= ord('z'):
															πF.SetLineno(257)
														Label1:
															// line 258: rtn = chr(self.i)
															πF.SetLineno(258)
															πTemp003 = πF.MakeArgs(1)
															if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
																continue
															}
															if πTemp001, πE = πg.GetAttr(πF, µself, ßi, nil); πE != nil {
																continue
															}
															πTemp003[0] = πTemp001
															if πTemp001, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
																continue
															}
															if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
																continue
															}
															πF.FreeArgs(πTemp003)
															µrtn = πTemp002
															// line 259: self.i += 1
															πF.SetLineno(259)
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
															// line 260: return rtn
															πF.SetLineno(260)
															if πE = πg.CheckLocal(πF, µrtn, "rtn"); πE != nil {
																continue
															}
															πR = µrtn
															continue
															goto Label2
														Label2:
															if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
																continue
															}
															// line 261: raise StopIteration
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
													if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp004); πE != nil {
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
											if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BogonIter").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
												continue
											}
											µBogonIter = πTemp005
											// line 262: return BogonIter()
											πF.SetLineno(262)
											if πE = πg.CheckLocal(πF, µBogonIter, "BogonIter"); πE != nil {
												continue
											}
											if πTemp002, πE = µBogonIter.Call(πF, nil, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 263: def __getitem__(self, key):
									πF.SetLineno(263)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µkey *πg.Object = πArgs[1]; _ = µkey
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											// line 264: raise Exc
											πF.SetLineno(264)
											πE = πF.Raise(µExc, nil, nil)
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
							if πTemp004, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FailingUserDict").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µFailingUserDict = πTemp008
							// line 265: self.assertRaises(Exc, d.update, FailingUserDict())
							πF.SetLineno(265)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp003[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µFailingUserDict, "FailingUserDict"); πE != nil {
								continue
							}
							if πTemp002, πE = µFailingUserDict.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp002
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
							// line 267: d = self._empty_mapping()
							πF.SetLineno(267)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp004
							// line 268: class badseq(object):
							πF.SetLineno(268)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp005 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("badseq", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 269: def __iter__(self):
									πF.SetLineno(269)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 270: return self
											πF.SetLineno(270)
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
									if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 271: def next(self):
									πF.SetLineno(271)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 272: raise Exc()
											πF.SetLineno(272)
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
									if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp003); πE != nil {
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
							if πTemp008, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("badseq").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µbadseq = πTemp008
							// line 274: self.assertRaises(Exc, d.update, badseq())
							πF.SetLineno(274)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp003[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µbadseq, "badseq"); πE != nil {
								continue
							}
							if πTemp002, πE = µbadseq.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp002
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
							// line 276: self.assertRaises(ValueError, d.update, [(1, 2, 3)])
							πF.SetLineno(276)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							πTemp009 = make([]*πg.Object, 1)
							πTemp002 = πg.NewTuple3(πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp009[0] = πTemp002
							πTemp002 = πg.NewList(πTemp009...).ToObject()
							πTemp003[2] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_update.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 176: @unittest.skip('grumpy')
					πF.SetLineno(176)
					πTemp007 = πF.MakeArgs(1)
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßtest_update); πE != nil {
						continue
					}
					πTemp007[0] = πTemp018
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = ßgrumpy.ToObject()
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp019, πE = πg.GetAttr(πF, πTemp018, ßskip, nil); πE != nil {
						continue
					}
					if πTemp018, πE = πTemp019.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					if πTemp019, πE = πTemp018.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πClass.SetItem(πF, ßtest_update.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 280: def test_get(self):
					πF.SetLineno(280)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("test_get", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 281: d = self._empty_mapping()
							πF.SetLineno(281)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 282: self.assertTrue(d.get(self.other.keys()[0]) is None)
							πF.SetLineno(282)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005 == πTemp002).ToObject()
							πTemp003[0] = πTemp001
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
							// line 283: self.assertEqual(d.get(self.other.keys()[0], 3), 3)
							πF.SetLineno(283)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp004[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(3).ToObject()
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
							// line 284: d = self.reference
							πF.SetLineno(284)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreference, nil); πE != nil {
								continue
							}
							µd = πTemp001
							// line 285: self.assertTrue(d.get(self.other.keys()[0]) is None)
							πF.SetLineno(285)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005 == πTemp002).ToObject()
							πTemp003[0] = πTemp001
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
							// line 286: self.assertEqual(d.get(self.other.keys()[0], 3), 3)
							πF.SetLineno(286)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp004[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(3).ToObject()
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
							// line 287: self.assertEqual(d.get(self.inmapping.keys()[0]), self.inmapping.values()[0])
							πF.SetLineno(287)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßinmapping, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßinmapping, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
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
							// line 288: self.assertEqual(d.get(self.inmapping.keys()[0], 3), self.inmapping.values()[0])
							πF.SetLineno(288)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßinmapping, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp004[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßinmapping, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
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
							// line 289: self.assertRaises(TypeError, d.get)
							πF.SetLineno(289)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
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
							// line 290: self.assertRaises(TypeError, d.get, None, None, None)
							πF.SetLineno(290)
							πTemp003 = πF.MakeArgs(5)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_get.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 292: def test_setdefault(self):
					πF.SetLineno(292)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("test_setdefault", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							// line 293: d = self._empty_mapping()
							πF.SetLineno(293)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 294: self.assertRaises(TypeError, d.setdefault)
							πF.SetLineno(294)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_setdefault.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 296: def test_popitem(self):
					πF.SetLineno(296)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("test_popitem", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							// line 297: d = self._empty_mapping()
							πF.SetLineno(297)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 298: self.assertRaises(KeyError, d.popitem)
							πF.SetLineno(298)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpopitem, nil); πE != nil {
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
							// line 299: self.assertRaises(TypeError, d.popitem, 42)
							πF.SetLineno(299)
							πTemp003 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpopitem, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = πg.NewInt(42).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_popitem.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 301: def test_pop(self):
					πF.SetLineno(301)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("test_pop", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var µv *πg.Object = πg.UnboundLocal; _ = µv
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 302: d = self._empty_mapping()
							πF.SetLineno(302)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 303: k, v = self.inmapping.items()[0]
							πF.SetLineno(303)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßinmapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							µk = πTemp001
							µv = πTemp003
							// line 304: d[k] = v
							πF.SetLineno(304)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µv); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp002 = µk
							if πE = πg.SetItem(πF, µd, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 305: self.assertRaises(KeyError, d.pop, self.other.keys()[0])
							πF.SetLineno(305)
							πTemp005 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßother, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πTemp005[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 307: self.assertEqual(d.pop(k), v)
							πF.SetLineno(307)
							πTemp005 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp006[0] = µk
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp005[1] = µv
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 308: self.assertEqual(len(d), 0)
							πF.SetLineno(308)
							πTemp005 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp006[0] = µd
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							πTemp005[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 310: self.assertRaises(KeyError, d.pop, k)
							πF.SetLineno(310)
							πTemp005 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp005[2] = µk
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_pop.ToObject(), πTemp021); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BasicTestMappingProtocol").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBasicTestMappingProtocol.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 313: class TestMappingProtocol(BasicTestMappingProtocol):
			πF.SetLineno(313)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestMappingProtocol", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 314: def test_constructor(self):
					πF.SetLineno(314)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_constructor", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 πg.KWArgs
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
							// line 315: BasicTestMappingProtocol.test_constructor(self)
							πF.SetLineno(315)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_constructor, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 316: self.assertTrue(self._empty_mapping() is not self._empty_mapping())
							πF.SetLineno(316)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004 != πTemp005).ToObject()
							πTemp001[0] = πTemp002
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
							// line 317: self.assertEqual(self.type2test(x=1, y=2), {"x": 1, "y": 2})
							πF.SetLineno(317)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πg.KWArgs{
								{"x", πg.NewInt(1).ToObject()},
								{"y", πg.NewInt(2).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp007 = πg.NewDict()
							if πE = πTemp007.SetItem(πF, ßx.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ßy.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp007.ToObject()
							πTemp001[1] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_constructor.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 319: def test_bool(self):
					πF.SetLineno(319)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_bool", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Dict
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 320: BasicTestMappingProtocol.test_bool(self)
							πF.SetLineno(320)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_bool, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 321: self.assertTrue(not self._empty_mapping())
							πF.SetLineno(321)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							πTemp001[0] = πTemp002
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
							// line 322: self.assertTrue(self._full_mapping({"x": "y"}))
							πF.SetLineno(322)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							if πE = πTemp007.SetItem(πF, ßx.ToObject(), ßy.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp007.ToObject()
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 323: self.assertTrue(bool(self._empty_mapping()) is False)
							πF.SetLineno(323)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004 == πTemp003).ToObject()
							πTemp001[0] = πTemp002
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
							// line 324: self.assertTrue(bool(self._full_mapping({"x": "y"})) is True)
							πF.SetLineno(324)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							if πE = πTemp007.SetItem(πF, ßx.ToObject(), ßy.ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp007.ToObject()
							πTemp008[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp006[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004 == πTemp003).ToObject()
							πTemp001[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_bool.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 326: def test_keys(self):
					πF.SetLineno(326)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_keys", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µk *πg.Object = πg.UnboundLocal; _ = µk
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 327: BasicTestMappingProtocol.test_keys(self)
							πF.SetLineno(327)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_keys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 328: d = self._empty_mapping()
							πF.SetLineno(328)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 329: self.assertEqual(d.keys(), [])
							πF.SetLineno(329)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[1] = πTemp002
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
							// line 330: d = self._full_mapping({'a': 1, 'b': 2})
							πF.SetLineno(330)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πg.NewDict()
							if πE = πTemp005.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp005.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 331: k = d.keys()
							πF.SetLineno(331)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µk = πTemp003
							// line 332: self.assertIn('a', k)
							πF.SetLineno(332)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp001[1] = µk
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
							// line 333: self.assertIn('b', k)
							πF.SetLineno(333)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßb.ToObject()
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp001[1] = µk
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
							// line 334: self.assertNotIn('c', k)
							πF.SetLineno(334)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp001[1] = µk
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
					if πE = πClass.SetItem(πF, ßtest_keys.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 336: def test_values(self):
					πF.SetLineno(336)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_values", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 337: BasicTestMappingProtocol.test_values(self)
							πF.SetLineno(337)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_values, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 338: d = self._full_mapping({1:2})
							πF.SetLineno(338)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 339: self.assertEqual(d.values(), [2])
							πF.SetLineno(339)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(2).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp001[1] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_values.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 341: def test_items(self):
					πF.SetLineno(341)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_items", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 342: BasicTestMappingProtocol.test_items(self)
							πF.SetLineno(342)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_items, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 344: d = self._full_mapping({1:2})
							πF.SetLineno(344)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 345: self.assertEqual(d.items(), [(1, 2)])
							πF.SetLineno(345)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp005 = make([]*πg.Object, 1)
							πTemp002 = πg.NewTuple2(πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp005[0] = πTemp002
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							πTemp001[1] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_items.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 347: def test_has_key(self):
					πF.SetLineno(347)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_has_key", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µk *πg.Object = πg.UnboundLocal; _ = µk
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Dict
						_ = πTemp007
						var πTemp008 []πg.Param
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 348: d = self._empty_mapping()
							πF.SetLineno(348)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 349: self.assertTrue(not d.has_key('a'))
							πF.SetLineno(349)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßhas_key, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							πTemp003[0] = πTemp001
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
							// line 350: d = self._full_mapping({'a': 1, 'b': 2})
							πF.SetLineno(350)
							πTemp003 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							if πE = πTemp007.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µd = πTemp002
							// line 351: k = d.keys()
							πF.SetLineno(351)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µk = πTemp002
							// line 352: k.sort(key=lambda k: (id(type(k)), k))
							πF.SetLineno(352)
							πTemp008 = make([]πg.Param, 1)
							πTemp008[0] = πg.Param{Name: "k", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/test/mapping_tests.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µk *πg.Object = πArgs[0]; _ = µk
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
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 352: k.sort(key=lambda k: (id(type(k)), k))
									πF.SetLineno(352)
									πTemp002 = πF.MakeArgs(1)
									πTemp003 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
										continue
									}
									πTemp003[0] = µk
									if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp003)
									πTemp002[0] = πTemp005
									if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
										continue
									}
									if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
										continue
									}
									πTemp001 = πg.NewTuple2(πTemp005, µk).ToObject()
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
							πTemp009 = πg.KWArgs{
								{"key", πTemp001},
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µk, ßsort, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp009); πE != nil {
								continue
							}
							// line 353: self.assertEqual(k, ['a', 'b'])
							πF.SetLineno(353)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[0] = µk
							πTemp004 = make([]*πg.Object, 2)
							πTemp004[0] = ßa.ToObject()
							πTemp004[1] = ßb.ToObject()
							πTemp001 = πg.NewList(πTemp004...).ToObject()
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
							// line 355: self.assertRaises(TypeError, d.has_key)
							πF.SetLineno(355)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßhas_key, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_has_key.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 357: def test_contains(self):
					πF.SetLineno(357)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_contains", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 358: d = self._empty_mapping()
							πF.SetLineno(358)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 359: self.assertNotIn('a', d)
							πF.SetLineno(359)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 360: self.assertTrue(not ('a' in d))
							πF.SetLineno(360)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µd, ßa.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							πTemp003[0] = πTemp001
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
							// line 361: self.assertTrue('a' not in d)
							πF.SetLineno(361)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µd, ßa.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							πTemp003[0] = πTemp001
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
							// line 362: d = self._full_mapping({'a': 1, 'b': 2})
							πF.SetLineno(362)
							πTemp003 = πF.MakeArgs(1)
							πTemp005 = πg.NewDict()
							if πE = πTemp005.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp005.ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µd = πTemp002
							// line 363: self.assertIn('a', d)
							πF.SetLineno(363)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 364: self.assertIn('b', d)
							πF.SetLineno(364)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßb.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 365: self.assertNotIn('c', d)
							πF.SetLineno(365)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 367: self.assertRaises(TypeError, d.__contains__)
							πF.SetLineno(367)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ß__contains__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_contains.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 369: def test_len(self):
					πF.SetLineno(369)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_len", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 370: BasicTestMappingProtocol.test_len(self)
							πF.SetLineno(370)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_len, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 371: d = self._full_mapping({'a': 1, 'b': 2})
							πF.SetLineno(371)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 372: self.assertEqual(len(d), 2)
							πF.SetLineno(372)
							πTemp001 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(2).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_len.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 374: def test_getitem(self):
					πF.SetLineno(374)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_getitem", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 375: BasicTestMappingProtocol.test_getitem(self)
							πF.SetLineno(375)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_getitem, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 376: d = self._full_mapping({'a': 1, 'b': 2})
							πF.SetLineno(376)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 377: self.assertEqual(d['a'], 1)
							πF.SetLineno(377)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(1).ToObject()
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
							// line 378: self.assertEqual(d['b'], 2)
							πF.SetLineno(378)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = ßb.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(2).ToObject()
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
							// line 379: d['c'] = 3
							πF.SetLineno(379)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003 = ßc.ToObject()
							if πE = πg.SetItem(πF, µd, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 380: d['a'] = 4
							πF.SetLineno(380)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003 = ßa.ToObject()
							if πE = πg.SetItem(πF, µd, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 381: self.assertEqual(d['c'], 3)
							πF.SetLineno(381)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(3).ToObject()
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
							// line 382: self.assertEqual(d['a'], 4)
							πF.SetLineno(382)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(4).ToObject()
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
							// line 383: del d['b']
							πF.SetLineno(383)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002 = ßb.ToObject()
							if πE = πg.DelItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							// line 384: self.assertEqual(d, {'a': 4, 'c': 3})
							πF.SetLineno(384)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, ßa.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßc.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
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
							// line 386: self.assertRaises(TypeError, d.__getitem__)
							πF.SetLineno(386)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ß__getitem__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_getitem.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 388: def test_clear(self):
					πF.SetLineno(388)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_clear", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Dict
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
							// line 389: d = self._full_mapping({1:1, 2:2, 3:3})
							πF.SetLineno(389)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewDict()
							if πE = πTemp002.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp002.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp002.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp002.ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp004
							// line 390: d.clear()
							πF.SetLineno(390)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 391: self.assertEqual(d, {})
							πF.SetLineno(391)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp002 = πg.NewDict()
							πTemp003 = πTemp002.ToObject()
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
							// line 393: self.assertRaises(TypeError, d.clear, None)
							πF.SetLineno(393)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_clear.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 395: def test_update(self):
					πF.SetLineno(395)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_update", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µSimpleUserDict *πg.Object = πg.UnboundLocal; _ = µSimpleUserDict
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
						_ = πTemp004
						var πTemp005 πg.KWArgs
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
							// line 396: BasicTestMappingProtocol.test_update(self)
							πF.SetLineno(396)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_update, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 398: d = self._empty_mapping()
							πF.SetLineno(398)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 399: d.update({1:100})
							πF.SetLineno(399)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 400: d.update({2:20})
							πF.SetLineno(400)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(20).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 401: d.update({1:1, 2:2, 3:3})
							πF.SetLineno(401)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 402: self.assertEqual(d, {1:1, 2:2, 3:3})
							πF.SetLineno(402)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
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
							// line 405: d.update()
							πF.SetLineno(405)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 406: self.assertEqual(d, {1:1, 2:2, 3:3})
							πF.SetLineno(406)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
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
							// line 409: d = self._empty_mapping()
							πF.SetLineno(409)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 410: d.update(x=100)
							πF.SetLineno(410)
							πTemp005 = πg.KWArgs{
								{"x", πg.NewInt(100).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							// line 411: d.update(y=20)
							πF.SetLineno(411)
							πTemp005 = πg.KWArgs{
								{"y", πg.NewInt(20).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							// line 412: d.update(x=1, y=2, z=3)
							πF.SetLineno(412)
							πTemp005 = πg.KWArgs{
								{"x", πg.NewInt(1).ToObject()},
								{"y", πg.NewInt(2).ToObject()},
								{"z", πg.NewInt(3).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							// line 413: self.assertEqual(d, {"x":1, "y":2, "z":3})
							πF.SetLineno(413)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, ßx.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßy.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßz.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
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
							// line 416: d = self._empty_mapping()
							πF.SetLineno(416)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 417: d.update([("x", 100), ("y", 20)])
							πF.SetLineno(417)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 2)
							πTemp002 = πg.NewTuple2(ßx.ToObject(), πg.NewInt(100).ToObject()).ToObject()
							πTemp006[0] = πTemp002
							πTemp002 = πg.NewTuple2(ßy.ToObject(), πg.NewInt(20).ToObject()).ToObject()
							πTemp006[1] = πTemp002
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 418: self.assertEqual(d, {"x":100, "y":20})
							πF.SetLineno(418)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, ßx.ToObject(), πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßy.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
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
							// line 421: d = self._empty_mapping()
							πF.SetLineno(421)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 422: d.update([("x", 100), ("y", 20)], x=1, y=2)
							πF.SetLineno(422)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 2)
							πTemp002 = πg.NewTuple2(ßx.ToObject(), πg.NewInt(100).ToObject()).ToObject()
							πTemp006[0] = πTemp002
							πTemp002 = πg.NewTuple2(ßy.ToObject(), πg.NewInt(20).ToObject()).ToObject()
							πTemp006[1] = πTemp002
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp001[0] = πTemp002
							πTemp005 = πg.KWArgs{
								{"x", πg.NewInt(1).ToObject()},
								{"y", πg.NewInt(2).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 423: self.assertEqual(d, {"x":1, "y":2})
							πF.SetLineno(423)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, ßx.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßy.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
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
							// line 426: d = self._full_mapping({1:3, 2:4})
							πF.SetLineno(426)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 427: d.update(self._full_mapping({1:2, 3:4, 5:6}).iteritems())
							πF.SetLineno(427)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(5).ToObject(), πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 428: self.assertEqual(d, {1:2, 2:4, 3:4, 5:6})
							πF.SetLineno(428)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(5).ToObject(), πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
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
							// line 430: class SimpleUserDict(object):
							πF.SetLineno(430)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("SimpleUserDict", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 431: def __init__(self):
									πF.SetLineno(431)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πTemp001 *πg.Dict
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
											// line 432: self.d = {1:1, 2:2, 3:3}
											πF.SetLineno(432)
											πTemp001 = πg.NewDict()
											if πE = πTemp001.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
												continue
											}
											if πE = πTemp001.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
												continue
											}
											if πE = πTemp001.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
												continue
											}
											πTemp002 = πTemp001.ToObject()
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßd, πTemp003); πE != nil {
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
									// line 433: def keys(self):
									πF.SetLineno(433)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 434: return self.d.keys()
											πF.SetLineno(434)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ßd, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ßkeys.ToObject(), πTemp003); πE != nil {
										continue
									}
									// line 435: def __getitem__(self, i):
									πF.SetLineno(435)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "i", Def: nil}
									πTemp004 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 436: return self.d[i]
											πF.SetLineno(436)
											if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
												continue
											}
											πTemp001 = µi
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetAttr(πF, µself, ßd, nil); πE != nil {
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
							if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp003 == nil {
								πTemp003 = πg.TypeType.ToObject()
							}
							if πTemp007, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SimpleUserDict").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µSimpleUserDict = πTemp007
							// line 437: d.clear()
							πF.SetLineno(437)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 438: d.update(SimpleUserDict())
							πF.SetLineno(438)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µSimpleUserDict, "SimpleUserDict"); πE != nil {
								continue
							}
							if πTemp002, πE = µSimpleUserDict.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 439: self.assertEqual(d, {1:1, 2:2, 3:3})
							πF.SetLineno(439)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp004 = πg.NewDict()
							if πE = πTemp004.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_update.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 441: def test_fromkeys(self):
					πF.SetLineno(441)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_fromkeys", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µg *πg.Object = πg.UnboundLocal; _ = µg
						var µdictlike *πg.Object = πg.UnboundLocal; _ = µdictlike
						var µmydict *πg.Object = πg.UnboundLocal; _ = µmydict
						var µud *πg.Object = πg.UnboundLocal; _ = µud
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µbaddict1 *πg.Object = πg.UnboundLocal; _ = µbaddict1
						var µBadSeq *πg.Object = πg.UnboundLocal; _ = µBadSeq
						var µbaddict2 *πg.Object = πg.UnboundLocal; _ = µbaddict2
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 []πg.Param
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 442: self.assertEqual(self.type2test.fromkeys('abc'), {'a':None, 'b':None, 'c':None})
							πF.SetLineno(442)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßabc.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp005 = πg.NewDict()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßa.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßb.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßc.ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp003 = πTemp005.ToObject()
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
							// line 443: d = self._empty_mapping()
							πF.SetLineno(443)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp004
							// line 444: self.assertTrue(not(d.fromkeys('abc') is d))
							πF.SetLineno(444)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßabc.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µd, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp007 == µd).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp008).ToObject()
							πTemp001[0] = πTemp003
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
							// line 445: self.assertEqual(d.fromkeys('abc'), {'a':None, 'b':None, 'c':None})
							πF.SetLineno(445)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßabc.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp005 = πg.NewDict()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßa.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßb.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßc.ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp003 = πTemp005.ToObject()
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
							// line 446: self.assertEqual(d.fromkeys((4,5),0), {4:0, 5:0})
							πF.SetLineno(446)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πg.NewTuple2(πg.NewInt(4).ToObject(), πg.NewInt(5).ToObject()).ToObject()
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp005 = πg.NewDict()
							if πE = πTemp005.SetItem(πF, πg.NewInt(4).ToObject(), πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, πg.NewInt(5).ToObject(), πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005.ToObject()
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
							// line 447: self.assertEqual(d.fromkeys([]), {})
							πF.SetLineno(447)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp009 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp009...).ToObject()
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp005 = πg.NewDict()
							πTemp003 = πTemp005.ToObject()
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
							// line 448: def g():
							πF.SetLineno(448)
							πTemp010 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("g", "/usr/lib/python2.7/test/mapping_tests.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										// line 449: yield 1
										πF.SetLineno(449)
										πF.PushCheckpoint(1)
										return πg.NewInt(1).ToObject(), nil
									Label1:
										πTemp001 = πSent
									}
									return nil, πE
								}).ToObject(), nil
							}), πF.Globals()).ToObject()
							µg = πTemp003
							// line 450: self.assertEqual(d.fromkeys(g()), {1:None})
							πF.SetLineno(450)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
								continue
							}
							if πTemp004, πE = µg.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µd, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, πg.NewInt(1).ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp004 = πTemp005.ToObject()
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 451: self.assertRaises(TypeError, {}.fromkeys, 3)
							πF.SetLineno(451)
							πTemp001 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp005 = πg.NewDict()
							πTemp004 = πTemp005.ToObject()
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							πTemp001[2] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 452: class dictlike(self.type2test): pass
							πF.SetLineno(452)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("dictlike", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 452: class dictlike(self.type2test): pass
									πF.SetLineno(452)
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
							if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("dictlike").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µdictlike = πTemp007
							// line 453: self.assertEqual(dictlike.fromkeys('a'), {'a':None})
							πF.SetLineno(453)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µdictlike, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßa.ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp004 = πTemp005.ToObject()
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 454: self.assertEqual(dictlike().fromkeys('a'), {'a':None})
							πF.SetLineno(454)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							if πTemp004, πE = µdictlike.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßa.ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp004 = πTemp005.ToObject()
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 455: self.assertTrue(dictlike.fromkeys('a').__class__ is dictlike)
							πF.SetLineno(455)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µdictlike, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp006 == µdictlike).ToObject()
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 456: self.assertTrue(dictlike().fromkeys('a').__class__ is dictlike)
							πF.SetLineno(456)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							if πTemp006, πE = µdictlike.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp007 == µdictlike).ToObject()
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 459: class mydict(self.type2test):
							πF.SetLineno(459)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("mydict", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 460: def __new__(cls):
									πF.SetLineno(460)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "cls", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µcls *πg.Object = πArgs[0]; _ = µcls
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
											// line 461: return UserDict.UserDict()
											πF.SetLineno(461)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßUserDict); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßUserDict, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
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
							if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("mydict").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µmydict = πTemp007
							// line 462: ud = mydict.fromkeys('ab')
							πF.SetLineno(462)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßab.ToObject()
							if πE = πg.CheckLocal(πF, µmydict, "mydict"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µmydict, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µud = πTemp006
							// line 463: self.assertEqual(ud, {'a':None, 'b':None})
							πF.SetLineno(463)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µud, "ud"); πE != nil {
								continue
							}
							πTemp001[0] = µud
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßa.ToObject(), πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßb.ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp004 = πTemp005.ToObject()
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 466: self.assertRaises(TypeError, dict.fromkeys)
							πF.SetLineno(466)
							πTemp001 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 468: class Exc(Exception): pass
							πF.SetLineno(468)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 468: class Exc(Exception): pass
									πF.SetLineno(468)
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
							if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp007
							// line 470: class baddict1(self.type2test):
							πF.SetLineno(470)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("baddict1", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 471: def __init__(self):
									πF.SetLineno(471)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 472: raise Exc()
											πF.SetLineno(472)
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
									if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
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
							if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("baddict1").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µbaddict1 = πTemp007
							// line 474: self.assertRaises(Exc, baddict1.fromkeys, [1])
							πF.SetLineno(474)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πE = πg.CheckLocal(πF, µbaddict1, "baddict1"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µbaddict1, ßfromkeys, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							πTemp001[2] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 476: class BadSeq(object):
							πF.SetLineno(476)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadSeq", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 477: def __iter__(self):
									πF.SetLineno(477)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 478: return self
											πF.SetLineno(478)
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
									if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 479: def next(self):
									πF.SetLineno(479)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 480: raise Exc()
											πF.SetLineno(480)
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
									if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp003); πE != nil {
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
							if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("BadSeq").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µBadSeq = πTemp007
							// line 482: self.assertRaises(Exc, self.type2test.fromkeys, BadSeq())
							πF.SetLineno(482)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp006
							if πE = πg.CheckLocal(πF, µBadSeq, "BadSeq"); πE != nil {
								continue
							}
							if πTemp004, πE = µBadSeq.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 484: class baddict2(self.type2test):
							πF.SetLineno(484)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("baddict2", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 485: def __setitem__(self, key, value):
									πF.SetLineno(485)
									πTemp002 = make([]πg.Param, 3)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp002[2] = πg.Param{Name: "value", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 486: raise Exc()
											πF.SetLineno(486)
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
							if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("baddict2").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µbaddict2 = πTemp007
							// line 488: self.assertRaises(Exc, baddict2.fromkeys, [1])
							πF.SetLineno(488)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πE = πg.CheckLocal(πF, µbaddict2, "baddict2"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µbaddict2, ßfromkeys, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp002 = make([]*πg.Object, 1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							πTemp001[2] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_fromkeys.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 490: def test_copy(self):
					πF.SetLineno(490)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_copy", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Dict
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
							// line 491: d = self._full_mapping({1:1, 2:2, 3:3})
							πF.SetLineno(491)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewDict()
							if πE = πTemp002.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp002.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp002.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp002.ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp004
							// line 492: self.assertEqual(d.copy(), {1:1, 2:2, 3:3})
							πF.SetLineno(492)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp002 = πg.NewDict()
							if πE = πTemp002.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp002.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp002.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp002.ToObject()
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
							// line 493: d = self._empty_mapping()
							πF.SetLineno(493)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp004
							// line 494: self.assertEqual(d.copy(), d)
							πF.SetLineno(494)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[1] = µd
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
							// line 495: self.assertIsInstance(d.copy(), d.__class__)
							πF.SetLineno(495)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ß__class__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 496: self.assertRaises(TypeError, d.copy, None)
							πF.SetLineno(496)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_copy.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 498: def test_get(self):
					πF.SetLineno(498)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_get", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πTemp006 *πg.Dict
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 499: BasicTestMappingProtocol.test_get(self)
							πF.SetLineno(499)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_get, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 500: d = self._empty_mapping()
							πF.SetLineno(500)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 501: self.assertTrue(d.get('c') is None)
							πF.SetLineno(501)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005 == πTemp003).ToObject()
							πTemp001[0] = πTemp002
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
							// line 502: self.assertEqual(d.get('c', 3), 3)
							πF.SetLineno(502)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßc.ToObject()
							πTemp004[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(3).ToObject()
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
							// line 503: d = self._full_mapping({'a' : 1, 'b' : 2})
							πF.SetLineno(503)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πg.NewDict()
							if πE = πTemp006.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp006.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 504: self.assertTrue(d.get('c') is None)
							πF.SetLineno(504)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005 == πTemp003).ToObject()
							πTemp001[0] = πTemp002
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
							// line 505: self.assertEqual(d.get('c', 3), 3)
							πF.SetLineno(505)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßc.ToObject()
							πTemp004[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(3).ToObject()
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
							// line 506: self.assertEqual(d.get('a'), 1)
							πF.SetLineno(506)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(1).ToObject()
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
							// line 507: self.assertEqual(d.get('a', 3), 1)
							πF.SetLineno(507)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßa.ToObject()
							πTemp004[1] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(1).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_get.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 509: def test_setdefault(self):
					πF.SetLineno(509)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_setdefault", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							// line 510: BasicTestMappingProtocol.test_setdefault(self)
							πF.SetLineno(510)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_setdefault, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 511: d = self._empty_mapping()
							πF.SetLineno(511)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 512: self.assertTrue(d.setdefault('key0') is None)
							πF.SetLineno(512)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßkey0.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005 == πTemp003).ToObject()
							πTemp001[0] = πTemp002
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
							// line 513: d.setdefault('key0', [])
							πF.SetLineno(513)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßkey0.ToObject()
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 514: self.assertTrue(d.setdefault('key0') is None)
							πF.SetLineno(514)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßkey0.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005 == πTemp003).ToObject()
							πTemp001[0] = πTemp002
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
							// line 515: d.setdefault('key', []).append(3)
							πF.SetLineno(515)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(3).ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßkey.ToObject()
							πTemp006 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 516: self.assertEqual(d['key'][0], 3)
							πF.SetLineno(516)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp005 = ßkey.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µd, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(3).ToObject()
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
							// line 517: d.setdefault('key', []).append(4)
							πF.SetLineno(517)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(4).ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßkey.ToObject()
							πTemp006 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 518: self.assertEqual(len(d['key']), 2)
							πF.SetLineno(518)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp002 = ßkey.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewInt(2).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_setdefault.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 520: def test_popitem(self):
					πF.SetLineno(520)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("test_popitem", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcopymode *πg.Object = πg.UnboundLocal; _ = µcopymode
						var µlog2size *πg.Object = πg.UnboundLocal; _ = µlog2size
						var µsize *πg.Object = πg.UnboundLocal; _ = µsize
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µka *πg.Object = πg.UnboundLocal; _ = µka
						var µva *πg.Object = πg.UnboundLocal; _ = µva
						var µta *πg.Object = πg.UnboundLocal; _ = µta
						var µkb *πg.Object = πg.UnboundLocal; _ = µkb
						var µvb *πg.Object = πg.UnboundLocal; _ = µvb
						var µtb *πg.Object = πg.UnboundLocal; _ = µtb
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
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 []*πg.Object
						_ = πTemp013
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
							case 14: goto Label14
							case 15: goto Label15
							default: panic("unexpected function state")
							}
							// line 521: BasicTestMappingProtocol.test_popitem(self)
							πF.SetLineno(521)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_popitem, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πg.NewInt(1).ToObject()).ToObject()
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
								µcopymode = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(12).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µlog2size = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 527: size = 2**log2size
							πF.SetLineno(527)
							if πE = πg.CheckLocal(πF, µlog2size, "log2size"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Pow(πF, πg.NewInt(2).ToObject(), µlog2size); πE != nil {
								continue
							}
							µsize = πTemp004
							// line 528: a = self._empty_mapping()
							πF.SetLineno(528)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							µa = πTemp007
							// line 529: b = self._empty_mapping()
							πF.SetLineno(529)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							µb = πTemp007
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp001[0] = µsize
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Iter(πF, πTemp009); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp008 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp007, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µi = πTemp007
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 531: a[repr(i)] = i
							πF.SetLineno(531)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, µi); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πTemp011, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp009 = πTemp012
							if πE = πg.SetItem(πF, µa, πTemp009, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcopymode, "copymode"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.LT(πF, µcopymode, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label10
							}
							goto Label11
							// line 532: if copymode < 0:
							πF.SetLineno(532)
						Label10:
							// line 533: b[repr(i)] = i
							πF.SetLineno(533)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, µi); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πTemp011, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp009 = πTemp012
							if πE = πg.SetItem(πF, µb, πTemp009, πTemp007); πE != nil {
								continue
							}
							goto Label11
						Label11:
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							if πE = πg.CheckLocal(πF, µcopymode, "copymode"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, µcopymode, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label12
							}
							goto Label13
							// line 534: if copymode > 0:
							πF.SetLineno(534)
						Label12:
							// line 535: b = a.copy()
							πF.SetLineno(535)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µa, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							µb = πTemp007
							goto Label13
						Label13:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp001[0] = µsize
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Iter(πF, πTemp009); πE != nil {
								continue
							}
							πF.PushCheckpoint(15)
							πTemp008 = false
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label16
							}
							if πTemp007, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µi = πTemp007
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(14)            
							// line 537: ka, va = ta = a.popitem()
							πF.SetLineno(537)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µa, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp011}}}, πTemp009); πE != nil {
								continue
							}
							µka = πTemp007
							µva = πTemp011
							µta = πTemp009
							// line 538: self.assertEqual(va, int(ka))
							πF.SetLineno(538)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µva, "va"); πE != nil {
								continue
							}
							πTemp001[0] = µva
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µka, "ka"); πE != nil {
								continue
							}
							πTemp013[0] = µka
							if πTemp007, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp001[1] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 539: kb, vb = tb = b.popitem()
							πF.SetLineno(539)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µb, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp011}}}, πTemp009); πE != nil {
								continue
							}
							µkb = πTemp007
							µvb = πTemp011
							µtb = πTemp009
							// line 540: self.assertEqual(vb, int(kb))
							πF.SetLineno(540)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvb, "vb"); πE != nil {
								continue
							}
							πTemp001[0] = µvb
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkb, "kb"); πE != nil {
								continue
							}
							πTemp013[0] = µkb
							if πTemp007, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp001[1] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 541: self.assertTrue(not(copymode < 0 and ta != tb))
							πF.SetLineno(541)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcopymode, "copymode"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.LT(πF, µcopymode, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp009 = πTemp011
							if πTemp010, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label17
							}
							if πE = πg.CheckLocal(πF, µta, "ta"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.NE(πF, µta, µtb); πE != nil {
								continue
							}
							πTemp009 = πTemp011
						Label17:
							if πTemp010, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(!πTemp010).ToObject()
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
						Label16:
							// line 542: self.assertTrue(not a)
							πF.SetLineno(542)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µa); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 543: self.assertTrue(not b)
							πF.SetLineno(543)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µb); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp008).ToObject()
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_popitem.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 545: def test_pop(self):
					πF.SetLineno(545)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("test_pop", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var µv *πg.Object = πg.UnboundLocal; _ = µv
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µh *πg.Object = πg.UnboundLocal; _ = µh
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
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
							// line 546: BasicTestMappingProtocol.test_pop(self)
							πF.SetLineno(546)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBasicTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_pop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 549: d = self._empty_mapping()
							πF.SetLineno(549)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 550: k, v = 'abc', 'def'
							πF.SetLineno(550)
							πTemp002 = πg.NewTuple2(ßabc.ToObject(), ßdef.ToObject()).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							µk = πTemp003
							µv = πTemp004
							// line 554: x = 4503599627370496L
							πF.SetLineno(554)
							µx = πg.NewLongFromBytes([]byte{0x10,0x0,0x0,0x0,0x0,0x0,0x0,}).ToObject()
							// line 555: y = 4503599627370496
							πF.SetLineno(555)
							µy = πg.NewInt(4503599627370496).ToObject()
							// line 556: h = self._full_mapping({x: 'anything', y: 'something else'})
							πF.SetLineno(556)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, µx, ßanything.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, µy, πg.NewStr("something else").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp005.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µh = πTemp003
							// line 557: self.assertEqual(h[x], h[y])
							πF.SetLineno(557)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp002 = µx
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µh, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp002 = µy
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µh, πTemp002); πE != nil {
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
							// line 559: self.assertEqual(d.pop(k, v), v)
							πF.SetLineno(559)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp006[0] = µk
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp006[1] = µv
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp001[1] = µv
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
							// line 560: d[k] = v
							πF.SetLineno(560)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µv); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003 = µk
							if πE = πg.SetItem(πF, µd, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 561: self.assertEqual(d.pop(k, 1), v)
							πF.SetLineno(561)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp006[0] = µk
							πTemp006[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp001[1] = µv
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
					if πE = πClass.SetItem(πF, ßtest_pop.ToObject(), πTemp018); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("TestMappingProtocol").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestMappingProtocol.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 564: class TestHashMappingProtocol(TestMappingProtocol):
			πF.SetLineno(564)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTestMappingProtocol); πE != nil {
				continue
			}
			πTemp002[0] = πTemp005
			πTemp003 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestHashMappingProtocol", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
					// line 566: def test_getitem(self):
					πF.SetLineno(566)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_getitem", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadEq *πg.Object = πg.UnboundLocal; _ = µBadEq
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µBadHash *πg.Object = πg.UnboundLocal; _ = µBadHash
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 567: TestMappingProtocol.test_getitem(self)
							πF.SetLineno(567)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_getitem, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 568: class Exc(Exception): pass
							πF.SetLineno(568)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 568: class Exc(Exception): pass
									πF.SetLineno(568)
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp005
							// line 570: class BadEq(object):
							πF.SetLineno(570)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadEq", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 571: def __eq__(self, other):
									πF.SetLineno(571)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 572: raise Exc()
											πF.SetLineno(572)
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 573: def __hash__(self):
									πF.SetLineno(573)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 574: return 24
											πF.SetLineno(574)
											πR = πg.NewInt(24).ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp003); πE != nil {
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadEq").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µBadEq = πTemp005
							// line 576: d = self._empty_mapping()
							πF.SetLineno(576)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 577: d[BadEq()] = 42
							πF.SetLineno(577)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(42).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µBadEq, "BadEq"); πE != nil {
								continue
							}
							if πTemp005, πE = µBadEq.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πE = πg.SetItem(πF, µd, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 578: self.assertRaises(KeyError, d.__getitem__, 23)
							πF.SetLineno(578)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewInt(23).ToObject()
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
							// line 580: class BadHash(object):
							πF.SetLineno(580)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadHash", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
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
									// line 581: fail = False
									πF.SetLineno(581)
									if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßfail.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 582: def __hash__(self):
									πF.SetLineno(582)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
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
											if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label1
											}
											goto Label2
											// line 583: if self.fail:
											πF.SetLineno(583)
										Label1:
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 584: raise Exc()
											πF.SetLineno(584)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label3
										Label2:
											// line 586: return 42
											πF.SetLineno(586)
											πR = πg.NewInt(42).ToObject()
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
									if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadHash").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µBadHash = πTemp005
							// line 588: d = self._empty_mapping()
							πF.SetLineno(588)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 589: x = BadHash()
							πF.SetLineno(589)
							if πE = πg.CheckLocal(πF, µBadHash, "BadHash"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadHash.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp002
							// line 590: d[x] = 42
							πF.SetLineno(590)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(42).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003 = µx
							if πE = πg.SetItem(πF, µd, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 591: x.fail = True
							πF.SetLineno(591)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µx, ßfail, πTemp003); πE != nil {
								continue
							}
							// line 592: self.assertRaises(Exc, d.__getitem__, x)
							πF.SetLineno(592)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[2] = µx
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
					if πE = πClass.SetItem(πF, ßtest_getitem.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 594: def test_fromkeys(self):
					πF.SetLineno(594)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_fromkeys", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmydict *πg.Object = πg.UnboundLocal; _ = µmydict
						var µud *πg.Object = πg.UnboundLocal; _ = µud
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 595: TestMappingProtocol.test_fromkeys(self)
							πF.SetLineno(595)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_fromkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 596: class mydict(self.type2test):
							πF.SetLineno(596)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("mydict", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
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
									// line 597: def __new__(cls):
									πF.SetLineno(597)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "cls", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µcls *πg.Object = πArgs[0]; _ = µcls
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
											// line 598: return UserDict.UserDict()
											πF.SetLineno(598)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßUserDict); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßUserDict, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("mydict").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µmydict = πTemp005
							// line 599: ud = mydict.fromkeys('ab')
							πF.SetLineno(599)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßab.ToObject()
							if πE = πg.CheckLocal(πF, µmydict, "mydict"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmydict, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µud = πTemp003
							// line 600: self.assertEqual(ud, {'a':None, 'b':None})
							πF.SetLineno(600)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µud, "ud"); πE != nil {
								continue
							}
							πTemp001[0] = µud
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßa.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ßb.ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
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
							// line 601: self.assertIsInstance(ud, UserDict.UserDict)
							πF.SetLineno(601)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µud, "ud"); πE != nil {
								continue
							}
							πTemp001[0] = µud
							if πTemp002, πE = πg.ResolveGlobal(πF, ßUserDict); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßUserDict, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_fromkeys.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 603: def test_pop(self):
					πF.SetLineno(603)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_pop", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadHash *πg.Object = πg.UnboundLocal; _ = µBadHash
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 604: TestMappingProtocol.test_pop(self)
							πF.SetLineno(604)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_pop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 606: class Exc(Exception): pass
							πF.SetLineno(606)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 606: class Exc(Exception): pass
									πF.SetLineno(606)
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp005
							// line 608: class BadHash(object):
							πF.SetLineno(608)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadHash", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
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
									// line 609: fail = False
									πF.SetLineno(609)
									if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßfail.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 610: def __hash__(self):
									πF.SetLineno(610)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
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
											if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label1
											}
											goto Label2
											// line 611: if self.fail:
											πF.SetLineno(611)
										Label1:
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 612: raise Exc()
											πF.SetLineno(612)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label3
										Label2:
											// line 614: return 42
											πF.SetLineno(614)
											πR = πg.NewInt(42).ToObject()
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
									if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadHash").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µBadHash = πTemp005
							// line 616: d = self._empty_mapping()
							πF.SetLineno(616)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 617: x = BadHash()
							πF.SetLineno(617)
							if πE = πg.CheckLocal(πF, µBadHash, "BadHash"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadHash.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp002
							// line 618: d[x] = 42
							πF.SetLineno(618)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(42).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003 = µx
							if πE = πg.SetItem(πF, µd, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 619: x.fail = True
							πF.SetLineno(619)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µx, ßfail, πTemp003); πE != nil {
								continue
							}
							// line 620: self.assertRaises(Exc, d.pop, x)
							πF.SetLineno(620)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[2] = µx
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
					if πE = πClass.SetItem(πF, ßtest_pop.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 622: def test_mutatingiteration(self):
					πF.SetLineno(622)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_mutatingiteration", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
						var πTemp007 []*πg.Object
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
							case 2: goto Label2
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 623: d = self._empty_mapping()
							πF.SetLineno(623)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 624: d[1] = 1
							πF.SetLineno(624)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 625: try:
							πF.SetLineno(625)
							πF.PushCheckpoint(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µd); πE != nil {
								continue
							}
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
								µi = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 627: d[i+1] = 1
							πF.SetLineno(627)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp006
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp002); πE != nil {
								continue
							}
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							πF.PopCheckpoint()
							// line 631: self.fail("changing dict size during iteration doesn't raise Error")
							πF.SetLineno(631)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("changing dict size during iteration doesn't raise Error").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 628: except RuntimeError:
							πF.SetLineno(628)
						Label6:
							// line 629: pass
							πF.SetLineno(629)
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
					if πE = πClass.SetItem(πF, ßtest_mutatingiteration.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 633: def test_repr(self):
					πF.SetLineno(633)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_repr", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadRepr *πg.Object = πg.UnboundLocal; _ = µBadRepr
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 634: d = self._empty_mapping()
							πF.SetLineno(634)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 635: self.assertEqual(repr(d), '{}')
							πF.SetLineno(635)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("{}").ToObject()
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
							// line 636: d[1] = 2
							πF.SetLineno(636)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 637: self.assertEqual(repr(d), '{1: 2}')
							πF.SetLineno(637)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("{1: 2}").ToObject()
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
							// line 638: d = self._empty_mapping()
							πF.SetLineno(638)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp002
							// line 639: d[1] = d
							πF.SetLineno(639)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µd); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp002, πTemp001); πE != nil {
								continue
							}
							// line 640: self.assertEqual(repr(d), '{1: {...}}')
							πF.SetLineno(640)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("{1: {...}}").ToObject()
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
							// line 642: class Exc(Exception): pass
							πF.SetLineno(642)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp003[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 642: class Exc(Exception): pass
									πF.SetLineno(642)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp002 == nil {
								πTemp002 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp002.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp006
							// line 644: class BadRepr(object):
							πF.SetLineno(644)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp003[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadRepr", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 645: def __repr__(self):
									πF.SetLineno(645)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 646: raise Exc()
											πF.SetLineno(646)
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
									if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp002 == nil {
								πTemp002 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp002.Call(πF, []*πg.Object{πg.NewStr("BadRepr").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µBadRepr = πTemp006
							// line 648: d = self._full_mapping({1: BadRepr()})
							πF.SetLineno(648)
							πTemp003 = πF.MakeArgs(1)
							πTemp005 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µBadRepr, "BadRepr"); πE != nil {
								continue
							}
							if πTemp001, πE = µBadRepr.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, πg.NewInt(1).ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp005.ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µd = πTemp002
							// line 649: self.assertRaises(Exc, repr, d)
							πF.SetLineno(649)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp003[0] = µExc
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[2] = µd
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
					if πE = πClass.SetItem(πF, ßtest_repr.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 651: def test_le(self):
					πF.SetLineno(651)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_le", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadCmp *πg.Object = πg.UnboundLocal; _ = µBadCmp
						var µd1 *πg.Object = πg.UnboundLocal; _ = µd1
						var µd2 *πg.Object = πg.UnboundLocal; _ = µd2
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.Dict
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
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 652: self.assertTrue(not (self._empty_mapping() < self._empty_mapping()))
							πF.SetLineno(652)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							πTemp001[0] = πTemp002
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
							// line 653: self.assertTrue(not (self._full_mapping({1: 2}) < self._full_mapping({1L: 2L})))
							πF.SetLineno(653)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							πTemp009 = πg.NewDict()
							if πE = πTemp009.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp009.ToObject()
							πTemp008[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp008 = πF.MakeArgs(1)
							πTemp009 = πg.NewDict()
							if πE = πTemp009.SetItem(πF, πg.NewLongFromBytes([]byte{0x1,}).ToObject(), πg.NewLongFromBytes([]byte{0x2,}).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp009.ToObject()
							πTemp008[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp003, πE = πg.LT(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							πTemp001[0] = πTemp002
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
							// line 655: class Exc(Exception): pass
							πF.SetLineno(655)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp009 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp009
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 655: class Exc(Exception): pass
									πF.SetLineno(655)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp003, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp003 == nil {
								πTemp003 = πg.TypeType.ToObject()
							}
							if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp004
							// line 657: class BadCmp(object):
							πF.SetLineno(657)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp009 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadCmp", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 658: def __eq__(self, other):
									πF.SetLineno(658)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 659: raise Exc()
											πF.SetLineno(659)
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 660: def __hash__(self):
									πF.SetLineno(660)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 661: return 42
											πF.SetLineno(661)
											πR = πg.NewInt(42).ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp003); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp003, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp003 == nil {
								πTemp003 = πg.TypeType.ToObject()
							}
							if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadCmp").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
								continue
							}
							µBadCmp = πTemp004
							// line 663: d1 = self._full_mapping({BadCmp(): 1})
							πF.SetLineno(663)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πTemp009.SetItem(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp009.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd1 = πTemp003
							// line 664: d2 = self._full_mapping({1: 1})
							πF.SetLineno(664)
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = πg.NewDict()
							if πE = πTemp009.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp009.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_full_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd2 = πTemp003
							// line 665: try:
							πF.SetLineno(665)
							πF.PushCheckpoint(2)
							// line 666: d1 < d2
							πF.SetLineno(666)
							if πE = πg.CheckLocal(πF, µd1, "d1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd2, "d2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µd1, µd2); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							// line 670: self.fail("< didn't raise Exc")
							πF.SetLineno(670)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("< didn't raise Exc").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp010.ToObject(), µExc); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 667: except Exc:
							πF.SetLineno(667)
						Label3:
							// line 668: pass
							πF.SetLineno(668)
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
					if πE = πClass.SetItem(πF, ßtest_le.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 672: def test_setdefault(self):
					πF.SetLineno(672)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_setdefault", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadHash *πg.Object = πg.UnboundLocal; _ = µBadHash
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Dict
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
							// line 673: TestMappingProtocol.test_setdefault(self)
							πF.SetLineno(673)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTestMappingProtocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtest_setdefault, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 675: class Exc(Exception): pass
							πF.SetLineno(675)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 675: class Exc(Exception): pass
									πF.SetLineno(675)
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp005
							// line 677: class BadHash(object):
							πF.SetLineno(677)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp004 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadHash", "/usr/lib/python2.7/test/mapping_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp004
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
									// line 678: fail = False
									πF.SetLineno(678)
									if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßfail.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 679: def __hash__(self):
									πF.SetLineno(679)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/mapping_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
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
											if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label1
											}
											goto Label2
											// line 680: if self.fail:
											πF.SetLineno(680)
										Label1:
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 681: raise Exc()
											πF.SetLineno(681)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label3
										Label2:
											// line 683: return 42
											πF.SetLineno(683)
											πR = πg.NewInt(42).ToObject()
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
									if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadHash").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µBadHash = πTemp005
							// line 685: d = self._empty_mapping()
							πF.SetLineno(685)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_empty_mapping, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp003
							// line 686: x = BadHash()
							πF.SetLineno(686)
							if πE = πg.CheckLocal(πF, µBadHash, "BadHash"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadHash.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp002
							// line 687: d[x] = 42
							πF.SetLineno(687)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(42).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003 = µx
							if πE = πg.SetItem(πF, µd, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 688: x.fail = True
							πF.SetLineno(688)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µx, ßfail, πTemp003); πE != nil {
								continue
							}
							// line 689: self.assertRaises(Exc, d.setdefault, x, [])
							πF.SetLineno(689)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[2] = µx
							πTemp006 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp001[3] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_setdefault.ToObject(), πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp004 == nil {
				πTemp004 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("TestHashMappingProtocol").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestHashMappingProtocol.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("test.mapping_tests", Code)
}

