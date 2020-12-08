package test_dict
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/test/test_support"
	// _ "github.com/pygolin/stdlib/string"
	// _ "github.com/pygolin/stdlib/UserDict"
	// _ "github.com/pygolin/stdlib/unittest"
	// _ "github.com/pygolin/stdlib/random"
	// _ "github.com/pygolin/stdlib/test/mapping_tests"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBasicTestMappingProtocol := πg.InternStr("BasicTestMappingProtocol")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßDict := πg.InternStr("Dict")
		ßDictTest := πg.InternStr("DictTest")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßGeneralMappingTests := πg.InternStr("GeneralMappingTests")
		ßKeyError := πg.InternStr("KeyError")
		ßNone := πg.InternStr("None")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßStopIteration := πg.InternStr("StopIteration")
		ßSubclassMappingTests := πg.InternStr("SubclassMappingTests")
		ßTestCase := πg.InternStr("TestCase")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUserDict := πg.InternStr("UserDict")
		ßValueError := πg.InternStr("ValueError")
		ßZeroDivisionError := πg.InternStr("ZeroDivisionError")
		ß__contains__ := πg.InternStr("__contains__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__missing__ := πg.InternStr("__missing__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß_not_tracked := πg.InternStr("_not_tracked")
		ß_tracked := πg.InternStr("_tracked")
		ßa := πg.InternStr("a")
		ßab := πg.InternStr("ab")
		ßabc := πg.InternStr("abc")
		ßanything := πg.InternStr("anything")
		ßappend := πg.InternStr("append")
		ßargs := πg.InternStr("args")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertIn := πg.InternStr("assertIn")
		ßassertIs := πg.InternStr("assertIs")
		ßassertIsInstance := πg.InternStr("assertIsInstance")
		ßassertIsNot := πg.InternStr("assertIsNot")
		ßassertNotIn := πg.InternStr("assertNotIn")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßb := πg.InternStr("b")
		ßbool := πg.InternStr("bool")
		ßc := πg.InternStr("c")
		ßcheck_py3k_warnings := πg.InternStr("check_py3k_warnings")
		ßchr := πg.InternStr("chr")
		ßclear := πg.InternStr("clear")
		ßcopy := πg.InternStr("copy")
		ßcpython_only := πg.InternStr("cpython_only")
		ßd := πg.InternStr("d")
		ßdef := πg.InternStr("def")
		ßdict := πg.InternStr("dict")
		ße := πg.InternStr("e")
		ßeq_count := πg.InternStr("eq_count")
		ßeval := πg.InternStr("eval")
		ßexception := πg.InternStr("exception")
		ßexpectedFailure := πg.InternStr("expectedFailure")
		ßf := πg.InternStr("f")
		ßfail := πg.InternStr("fail")
		ßformat := πg.InternStr("format")
		ßfromkeys := πg.InternStr("fromkeys")
		ßg := πg.InternStr("g")
		ßget := πg.InternStr("get")
		ßghi := πg.InternStr("ghi")
		ßh := πg.InternStr("h")
		ßhas_key := πg.InternStr("has_key")
		ßhasattr := πg.InternStr("hasattr")
		ßhash_count := πg.InternStr("hash_count")
		ßi := πg.InternStr("i")
		ßid := πg.InternStr("id")
		ßint := πg.InternStr("int")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßkey := πg.InternStr("key")
		ßkey0 := πg.InternStr("key0")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßletters := πg.InternStr("letters")
		ßmapping_tests := πg.InternStr("mapping_tests")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßpop := πg.InternStr("pop")
		ßpopitem := πg.InternStr("popitem")
		ßrandom := πg.InternStr("random")
		ßrange := πg.InternStr("range")
		ßrepr := πg.InternStr("repr")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsample := πg.InternStr("sample")
		ßset := πg.InternStr("set")
		ßsetdefault := πg.InternStr("setdefault")
		ßshuffle := πg.InternStr("shuffle")
		ßsort := πg.InternStr("sort")
		ßstring := πg.InternStr("string")
		ßtest_bool := πg.InternStr("test_bool")
		ßtest_clear := πg.InternStr("test_clear")
		ßtest_constructor := πg.InternStr("test_constructor")
		ßtest_contains := πg.InternStr("test_contains")
		ßtest_copy := πg.InternStr("test_copy")
		ßtest_empty_presized_dict_in_freelist := πg.InternStr("test_empty_presized_dict_in_freelist")
		ßtest_fromkeys := πg.InternStr("test_fromkeys")
		ßtest_get := πg.InternStr("test_get")
		ßtest_getitem := πg.InternStr("test_getitem")
		ßtest_has_key := πg.InternStr("test_has_key")
		ßtest_items := πg.InternStr("test_items")
		ßtest_keys := πg.InternStr("test_keys")
		ßtest_le := πg.InternStr("test_le")
		ßtest_len := πg.InternStr("test_len")
		ßtest_literal_constructor := πg.InternStr("test_literal_constructor")
		ßtest_main := πg.InternStr("test_main")
		ßtest_missing := πg.InternStr("test_missing")
		ßtest_mutatingiteration := πg.InternStr("test_mutatingiteration")
		ßtest_pop := πg.InternStr("test_pop")
		ßtest_popitem := πg.InternStr("test_popitem")
		ßtest_repr := πg.InternStr("test_repr")
		ßtest_resize1 := πg.InternStr("test_resize1")
		ßtest_resize2 := πg.InternStr("test_resize2")
		ßtest_setdefault := πg.InternStr("test_setdefault")
		ßtest_setdefault_atomic := πg.InternStr("test_setdefault_atomic")
		ßtest_support := πg.InternStr("test_support")
		ßtest_track_dynamic := πg.InternStr("test_track_dynamic")
		ßtest_track_literals := πg.InternStr("test_track_literals")
		ßtest_track_subtypes := πg.InternStr("test_track_subtypes")
		ßtest_tuple_keyerror := πg.InternStr("test_tuple_keyerror")
		ßtest_update := πg.InternStr("test_update")
		ßtest_values := πg.InternStr("test_values")
		ßtype2test := πg.InternStr("type2test")
		ßunittest := πg.InternStr("unittest")
		ßupdate := πg.InternStr("update")
		ßvalues := πg.InternStr("values")
		ßz := πg.InternStr("z")
		ßzip := πg.InternStr("zip")
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
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 bool
		_ = πTemp008
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: import unittest
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 2: from test import test_support
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: import UserDict, random, string
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "UserDict"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßUserDict.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "random"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßrandom.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "string"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstring.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: class DictTest(unittest.TestCase):
			πF.SetLineno(8)
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
			_, πE = πg.NewCode("DictTest", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 9: def test_constructor(self):
					πF.SetLineno(9)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_constructor", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 11: self.assertEqual(dict(), {})
							πF.SetLineno(11)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πg.NewDict()
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
							// line 12: self.assertIsNot(dict(), {})
							πF.SetLineno(12)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp004 = πg.NewDict()
							πTemp002 = πTemp004.ToObject()
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIsNot, nil); πE != nil {
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
					// line 15: def test_literal_constructor(self):
					πF.SetLineno(15)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_literal_constructor", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µitems *πg.Object = πg.UnboundLocal; _ = µitems
						var µformatted_items *πg.Object = πg.UnboundLocal; _ = µformatted_items
						var µdictliteral *πg.Object = πg.UnboundLocal; _ = µdictliteral
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
						var πTemp006 []πg.Param
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
						var πTemp012 []*πg.Object
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
							πTemp002 = πg.NewTuple5(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject(), πg.NewInt(6).ToObject(), πg.NewInt(256).ToObject(), πg.NewInt(400).ToObject()).ToObject()
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
								µn = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 19: items = [(''.join(random.sample(string.letters, 8)), i)
							πF.SetLineno(19)
							πTemp006 = make([]πg.Param, 0)
							πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/test_dict.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								var πTemp007 []*πg.Object
								_ = πTemp007
								var πTemp008 *πg.Object
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
										if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
											continue
										}
										πTemp002[0] = µn
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
										// line 19: items = [(''.join(random.sample(string.letters, 8)), i)
										πF.SetLineno(19)
										πTemp002 = πF.MakeArgs(1)
										πTemp007 = πF.MakeArgs(2)
										if πTemp004, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
											continue
										}
										if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßletters, nil); πE != nil {
											continue
										}
										πTemp007[0] = πTemp008
										πTemp007[1] = πg.NewInt(8).ToObject()
										if πTemp004, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
											continue
										}
										if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßsample, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
										πTemp002[0] = πTemp004
										if πTemp004, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
											continue
										}
										if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										πTemp003 = πg.NewTuple2(πTemp008, µi).ToObject()
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
							if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							µitems = πTemp002
							// line 21: random.shuffle(items)
							πF.SetLineno(21)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp008[0] = µitems
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrandom); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp002, ßshuffle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 22: formatted_items = ('{!r}: {:d}'.format(k, v) for k, v in items)
							πF.SetLineno(22)
							πTemp006 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/test_dict.py", πTemp006, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µk *πg.Object = πg.UnboundLocal; _ = µk
								var µv *πg.Object = πg.UnboundLocal; _ = µv
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
								var πTemp007 []*πg.Object
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
										if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µitems); πE != nil {
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
											if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
												continue
											}
											µk = πTemp005
											µv = πTemp006
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 22: formatted_items = ('{!r}: {:d}'.format(k, v) for k, v in items)
										πF.SetLineno(22)
										πTemp007 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
											continue
										}
										πTemp007[0] = µk
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										πTemp007[1] = µv
										if πTemp004, πE = πg.GetAttr(πF, πg.NewStr("{!r}: {:d}").ToObject(), ßformat, nil); πE != nil {
											continue
										}
										if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
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
							if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µformatted_items = πTemp007
							// line 23: dictliteral = '{' + ', '.join(formatted_items) + '}'
							πF.SetLineno(23)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformatted_items, "formatted_items"); πE != nil {
								continue
							}
							πTemp008[0] = µformatted_items
							if πTemp010, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp009, πE = πg.Add(πF, πg.NewStr("{").ToObject(), πTemp011); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, πTemp009, πg.NewStr("}").ToObject()); πE != nil {
								continue
							}
							µdictliteral = πTemp007
							// line 24: self.assertEqual(eval(dictliteral), dict(items))
							πF.SetLineno(24)
							πTemp008 = πF.MakeArgs(2)
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdictliteral, "dictliteral"); πE != nil {
								continue
							}
							πTemp012[0] = µdictliteral
							if πTemp007, πE = πg.ResolveGlobal(πF, ßeval); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp008[0] = πTemp009
							πTemp012 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp012[0] = µitems
							if πTemp007, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp012, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp012)
							πTemp008[1] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
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
					if πE = πClass.SetItem(πF, ßtest_literal_constructor.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 14: @unittest.expectedFailure
					πF.SetLineno(14)
					πTemp004 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßtest_literal_constructor); πE != nil {
						continue
					}
					πTemp004[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_literal_constructor.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 26: def test_bool(self):
					πF.SetLineno(26)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_bool", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Dict
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
							// line 27: self.assertIs(not {}, True)
							πF.SetLineno(27)
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewDict()
							πTemp004 = πTemp003.ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 28: self.assertTrue({1: 2})
							πF.SetLineno(28)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewDict()
							if πE = πTemp003.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003.ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 29: self.assertIs(bool({}), False)
							πF.SetLineno(29)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp003 = πg.NewDict()
							πTemp002 = πTemp003.ToObject()
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 30: self.assertIs(bool({1: 2}), True)
							πF.SetLineno(30)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp003 = πg.NewDict()
							if πE = πTemp003.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003.ToObject()
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_bool.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 33: def test_keys(self):
					πF.SetLineno(33)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_keys", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 34: d = {}
							πF.SetLineno(34)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 35: self.assertEqual(d.keys(), [])
							πF.SetLineno(35)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
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
							// line 36: d = {'a': 1, 'b': 2}
							πF.SetLineno(36)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 37: k = d.keys()
							πF.SetLineno(37)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µk = πTemp004
							// line 39: self.assertIn('a', k)
							πF.SetLineno(39)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[1] = µk
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
							// line 40: self.assertIn('b', k)
							πF.SetLineno(40)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßb.ToObject()
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[1] = µk
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
							// line 41: self.assertTrue(d.has_key('a'))
							πF.SetLineno(41)
							πTemp003 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßhas_key, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 42: self.assertTrue(d.has_key('b'))
							πF.SetLineno(42)
							πTemp003 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßb.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßhas_key, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 43: self.assertRaises(TypeError, d.keys, None)
							πF.SetLineno(43)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_keys.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 32: @unittest.expectedFailure
					πF.SetLineno(32)
					πTemp004 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_keys); πE != nil {
						continue
					}
					πTemp004[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_keys.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 45: def test_values(self):
					πF.SetLineno(45)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_values", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 46: d = {}
							πF.SetLineno(46)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 47: self.assertEqual(d.values(), [])
							πF.SetLineno(47)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
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
							// line 48: d = {1:2}
							πF.SetLineno(48)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 49: self.assertEqual(d.values(), [2])
							πF.SetLineno(49)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßvalues, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = πg.NewInt(2).ToObject()
							πTemp002 = πg.NewList(πTemp005...).ToObject()
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
							// line 51: self.assertRaises(TypeError, d.values, None)
							πF.SetLineno(51)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßvalues, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_values.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 53: def test_items(self):
					πF.SetLineno(53)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_items", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 54: d = {}
							πF.SetLineno(54)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 55: self.assertEqual(d.items(), [])
							πF.SetLineno(55)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
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
							// line 57: d = {1:2}
							πF.SetLineno(57)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 58: self.assertEqual(d.items(), [(1, 2)])
							πF.SetLineno(58)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp005 = make([]*πg.Object, 1)
							πTemp002 = πg.NewTuple2(πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp005[0] = πTemp002
							πTemp002 = πg.NewList(πTemp005...).ToObject()
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
							// line 60: self.assertRaises(TypeError, d.items, None)
							πF.SetLineno(60)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßitems, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_items.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 63: def test_has_key(self):
					πF.SetLineno(63)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_has_key", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 64: d = {}
							πF.SetLineno(64)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 65: self.assertFalse(d.has_key('a'))
							πF.SetLineno(65)
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
							// line 66: d = {'a': 1, 'b': 2}
							πF.SetLineno(66)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 67: k = d.keys()
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µk = πTemp005
							// line 68: k.sort()
							πF.SetLineno(68)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µk, ßsort, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 69: self.assertEqual(k, ['a', 'b'])
							πF.SetLineno(69)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[0] = µk
							πTemp004 = make([]*πg.Object, 2)
							πTemp004[0] = ßa.ToObject()
							πTemp004[1] = ßb.ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
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
							// line 71: self.assertRaises(TypeError, d.has_key)
							πF.SetLineno(71)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßhas_key, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_has_key.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 62: @unittest.expectedFailure
					πF.SetLineno(62)
					πTemp004 = πF.MakeArgs(1)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßtest_has_key); πE != nil {
						continue
					}
					πTemp004[0] = πTemp010
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_has_key.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 73: def test_contains(self):
					πF.SetLineno(73)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_contains", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 74: d = {}
							πF.SetLineno(74)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 75: self.assertNotIn('a', d)
							πF.SetLineno(75)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßa.ToObject()
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
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 76: self.assertFalse('a' in d)
							πF.SetLineno(76)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µd, ßa.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 77: self.assertTrue('a' not in d)
							πF.SetLineno(77)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, µd, ßa.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 78: d = {'a': 1, 'b': 2}
							πF.SetLineno(78)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 79: self.assertIn('a', d)
							πF.SetLineno(79)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßa.ToObject()
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
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 80: self.assertIn('b', d)
							πF.SetLineno(80)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßb.ToObject()
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
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 81: self.assertNotIn('c', d)
							πF.SetLineno(81)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßc.ToObject()
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
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 83: self.assertRaises(TypeError, d.__contains__)
							πF.SetLineno(83)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ß__contains__, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_contains.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 85: def test_len(self):
					πF.SetLineno(85)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_len", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 86: d = {}
							πF.SetLineno(86)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 87: self.assertEqual(len(d), 0)
							πF.SetLineno(87)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewInt(0).ToObject()
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
							// line 88: d = {'a': 1, 'b': 2}
							πF.SetLineno(88)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 89: self.assertEqual(len(d), 2)
							πF.SetLineno(89)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßtest_len.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 91: def test_getitem(self):
					πF.SetLineno(91)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_getitem", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µBadEq *πg.Object = πg.UnboundLocal; _ = µBadEq
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadHash *πg.Object = πg.UnboundLocal; _ = µBadHash
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
							// line 92: d = {'a': 1, 'b': 2}
							πF.SetLineno(92)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 93: self.assertEqual(d['a'], 1)
							πF.SetLineno(93)
							πTemp003 = πF.MakeArgs(2)
							πTemp002 = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
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
							// line 94: self.assertEqual(d['b'], 2)
							πF.SetLineno(94)
							πTemp003 = πF.MakeArgs(2)
							πTemp002 = ßb.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(2).ToObject()
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
							// line 95: d['c'] = 3
							πF.SetLineno(95)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004 = ßc.ToObject()
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 96: d['a'] = 4
							πF.SetLineno(96)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004 = ßa.ToObject()
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 97: self.assertEqual(d['c'], 3)
							πF.SetLineno(97)
							πTemp003 = πF.MakeArgs(2)
							πTemp002 = ßc.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(3).ToObject()
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
							// line 98: self.assertEqual(d['a'], 4)
							πF.SetLineno(98)
							πTemp003 = πF.MakeArgs(2)
							πTemp002 = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewInt(4).ToObject()
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
							// line 99: del d['b']
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002 = ßb.ToObject()
							if πE = πg.DelItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							// line 100: self.assertEqual(d, {'a': 4, 'c': 3})
							πF.SetLineno(100)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßa.ToObject(), πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ßc.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
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
							// line 102: self.assertRaises(TypeError, d.__getitem__)
							πF.SetLineno(102)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
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
							// line 104: class BadEq(object):
							πF.SetLineno(104)
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
							_, πE = πg.NewCode("BadEq", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 105: def __eq__(self, other):
									πF.SetLineno(105)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 106: raise Exc()
											πF.SetLineno(106)
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
									// line 107: def __hash__(self):
									πF.SetLineno(107)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 108: return 24
											πF.SetLineno(108)
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
							if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadEq").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µBadEq = πTemp005
							// line 110: d = {}
							πF.SetLineno(110)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 111: d[BadEq()] = 42
							πF.SetLineno(111)
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
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 112: self.assertRaises(KeyError, d.__getitem__, 23)
							πF.SetLineno(112)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							πTemp003[2] = πg.NewInt(23).ToObject()
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
							// line 114: class Exc(Exception): pass
							πF.SetLineno(114)
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
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 114: class Exc(Exception): pass
									πF.SetLineno(114)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp005
							// line 116: class BadHash(object):
							πF.SetLineno(116)
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
							_, πE = πg.NewCode("BadHash", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 117: fail = False
									πF.SetLineno(117)
									if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßfail.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 118: def __hash__(self):
									πF.SetLineno(118)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 119: if self.fail:
											πF.SetLineno(119)
										Label1:
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 120: raise Exc()
											πF.SetLineno(120)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label3
										Label2:
											// line 122: return 42
											πF.SetLineno(122)
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
							if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadHash").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µBadHash = πTemp005
							// line 124: x = BadHash()
							πF.SetLineno(124)
							if πE = πg.CheckLocal(πF, µBadHash, "BadHash"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadHash.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp002
							// line 125: d[x] = 42
							πF.SetLineno(125)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(42).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004 = µx
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 126: x.fail = True
							πF.SetLineno(126)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µx, ßfail, πTemp004); πE != nil {
								continue
							}
							// line 127: self.assertRaises(Exc, d.__getitem__, x)
							πF.SetLineno(127)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp003[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003[2] = µx
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
					if πE = πClass.SetItem(πF, ßtest_getitem.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 129: def test_clear(self):
					πF.SetLineno(129)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_clear", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Dict
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
							// line 130: d = {1:1, 2:2, 3:3}
							πF.SetLineno(130)
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
							µd = πTemp002
							// line 131: d.clear()
							πF.SetLineno(131)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 132: self.assertEqual(d, {})
							πF.SetLineno(132)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							πTemp004[1] = πTemp002
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
							// line 134: self.assertRaises(TypeError, d.clear, None)
							πF.SetLineno(134)
							πTemp004 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[2] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_clear.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 137: def test_update(self):
					πF.SetLineno(137)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_update", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µSimpleUserDict *πg.Object = πg.UnboundLocal; _ = µSimpleUserDict
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µFailingUserDict *πg.Object = πg.UnboundLocal; _ = µFailingUserDict
						var µbadseq *πg.Object = πg.UnboundLocal; _ = µbadseq
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
							// line 138: d = {}
							πF.SetLineno(138)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 139: d.update({1:100})
							πF.SetLineno(139)
							πTemp003 = πF.MakeArgs(1)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 140: d.update({2:20})
							πF.SetLineno(140)
							πTemp003 = πF.MakeArgs(1)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, πg.NewInt(2).ToObject(), πg.NewInt(20).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 141: d.update({1:1, 2:2, 3:3})
							πF.SetLineno(141)
							πTemp003 = πF.MakeArgs(1)
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
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 142: self.assertEqual(d, {1:1, 2:2, 3:3})
							πF.SetLineno(142)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
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
							// line 144: d.update()
							πF.SetLineno(144)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 145: self.assertEqual(d, {1:1, 2:2, 3:3})
							πF.SetLineno(145)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
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
							// line 147: self.assertRaises((TypeError, AttributeError), d.update, None)
							πF.SetLineno(147)
							πTemp003 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
							// line 149: class SimpleUserDict(object):
							πF.SetLineno(149)
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
							_, πE = πg.NewCode("SimpleUserDict", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 150: def __init__(self):
									πF.SetLineno(150)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 151: self.d = {1:1, 2:2, 3:3}
											πF.SetLineno(151)
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
									// line 152: def keys(self):
									πF.SetLineno(152)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 153: return self.d.keys()
											πF.SetLineno(153)
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
									// line 154: def __getitem__(self, i):
									πF.SetLineno(154)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "i", Def: nil}
									πTemp004 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 155: return self.d[i]
											πF.SetLineno(155)
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
							if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("SimpleUserDict").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µSimpleUserDict = πTemp005
							// line 156: d.clear()
							πF.SetLineno(156)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 157: d.update(SimpleUserDict())
							πF.SetLineno(157)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µSimpleUserDict, "SimpleUserDict"); πE != nil {
								continue
							}
							if πTemp002, πE = µSimpleUserDict.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 158: self.assertEqual(d, {1:1, 2:2, 3:3})
							πF.SetLineno(158)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
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
							// line 160: class Exc(Exception): pass
							πF.SetLineno(160)
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
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 160: class Exc(Exception): pass
									πF.SetLineno(160)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp005
							// line 162: d.clear()
							πF.SetLineno(162)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 163: class FailingUserDict(object):
							πF.SetLineno(163)
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
							_, πE = πg.NewCode("FailingUserDict", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 164: def keys(self):
									πF.SetLineno(164)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 165: raise Exc
											πF.SetLineno(165)
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
							if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FailingUserDict").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µFailingUserDict = πTemp005
							// line 166: self.assertRaises(Exc, d.update, FailingUserDict())
							πF.SetLineno(166)
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
							// line 168: class FailingUserDict(object):
							πF.SetLineno(168)
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
							_, πE = πg.NewCode("FailingUserDict", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 169: def keys(self):
									πF.SetLineno(169)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 170: class BogonIter(object):
											πF.SetLineno(170)
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
											_, πE = πg.NewCode("BogonIter", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
													// line 171: def __init__(self):
													πF.SetLineno(171)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
															// line 172: self.i = 1
															πF.SetLineno(172)
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
													// line 173: def __iter__(self):
													πF.SetLineno(173)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
														var µself *πg.Object = πArgs[0]; _ = µself
														var πR *πg.Object; _ = πR
														var πE *πg.BaseException; _ = πE
														for ; πF.State() >= 0; πF.PopCheckpoint() {
															switch πF.State() {
															case 0:
															default: panic("unexpected function state")
															}
															// line 174: return self
															πF.SetLineno(174)
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
													// line 175: def next(self):
													πF.SetLineno(175)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp004 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
															// line 176: if self.i:
															πF.SetLineno(176)
														Label1:
															// line 177: self.i = 0
															πF.SetLineno(177)
															if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
																continue
															}
															if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
																continue
															}
															if πE = πg.SetAttr(πF, µself, ßi, πTemp001); πE != nil {
																continue
															}
															// line 178: return 'a'
															πF.SetLineno(178)
															πR = ßa.ToObject()
															continue
															goto Label2
														Label2:
															if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
																continue
															}
															// line 179: raise Exc
															πF.SetLineno(179)
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
											// line 180: return BogonIter()
											πF.SetLineno(180)
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
									// line 181: def __getitem__(self, key):
									πF.SetLineno(181)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µkey *πg.Object = πArgs[1]; _ = µkey
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 182: return key
											πF.SetLineno(182)
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
							if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FailingUserDict").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µFailingUserDict = πTemp005
							// line 183: self.assertRaises(Exc, d.update, FailingUserDict())
							πF.SetLineno(183)
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
							// line 185: class FailingUserDict(object):
							πF.SetLineno(185)
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
							_, πE = πg.NewCode("FailingUserDict", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 186: def keys(self):
									πF.SetLineno(186)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("keys", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 187: class BogonIter(object):
											πF.SetLineno(187)
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
											_, πE = πg.NewCode("BogonIter", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
													// line 188: def __init__(self):
													πF.SetLineno(188)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
															// line 189: self.i = ord('a')
															πF.SetLineno(189)
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
													// line 190: def __iter__(self):
													πF.SetLineno(190)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
														var µself *πg.Object = πArgs[0]; _ = µself
														var πR *πg.Object; _ = πR
														var πE *πg.BaseException; _ = πE
														for ; πF.State() >= 0; πF.PopCheckpoint() {
															switch πF.State() {
															case 0:
															default: panic("unexpected function state")
															}
															// line 191: return self
															πF.SetLineno(191)
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
													// line 192: def next(self):
													πF.SetLineno(192)
													πTemp002 = make([]πg.Param, 1)
													πTemp002[0] = πg.Param{Name: "self", Def: nil}
													πTemp004 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
															// line 193: if self.i <= ord('z'):
															πF.SetLineno(193)
														Label1:
															// line 194: rtn = chr(self.i)
															πF.SetLineno(194)
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
															// line 195: self.i += 1
															πF.SetLineno(195)
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
															// line 196: return rtn
															πF.SetLineno(196)
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
															// line 197: raise StopIteration
															πF.SetLineno(197)
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
											// line 198: return BogonIter()
											πF.SetLineno(198)
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
									// line 199: def __getitem__(self, key):
									πF.SetLineno(199)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 200: raise Exc
											πF.SetLineno(200)
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
							if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("FailingUserDict").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µFailingUserDict = πTemp005
							// line 201: self.assertRaises(Exc, d.update, FailingUserDict())
							πF.SetLineno(201)
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
							// line 203: class badseq(object):
							πF.SetLineno(203)
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
							_, πE = πg.NewCode("badseq", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 204: def __iter__(self):
									πF.SetLineno(204)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 205: return self
											πF.SetLineno(205)
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
									// line 206: def next(self):
									πF.SetLineno(206)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 207: raise Exc()
											πF.SetLineno(207)
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
							if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("badseq").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µbadseq = πTemp005
							// line 209: self.assertRaises(Exc, {}.update, badseq())
							πF.SetLineno(209)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp003[0] = µExc
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
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
							// line 211: self.assertRaises(ValueError, {}.update, [(1, 2, 3)])
							πF.SetLineno(211)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							πTemp006 = make([]*πg.Object, 1)
							πTemp002 = πg.NewTuple3(πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp006[0] = πTemp002
							πTemp002 = πg.NewList(πTemp006...).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_update.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 136: @unittest.expectedFailure
					πF.SetLineno(136)
					πTemp004 = πF.MakeArgs(1)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßtest_update); πE != nil {
						continue
					}
					πTemp004[0] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp016.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_update.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 214: def test_fromkeys(self):
					πF.SetLineno(214)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_fromkeys", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var µbaddict3 *πg.Object = πg.UnboundLocal; _ = µbaddict3
						var µres *πg.Object = πg.UnboundLocal; _ = µres
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []πg.Param
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
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
							default: panic("unexpected function state")
							}
							// line 215: self.assertEqual(dict.fromkeys('abc'), {'a':None, 'b':None, 'c':None})
							πF.SetLineno(215)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßabc.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
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
							// line 216: d = {}
							πF.SetLineno(216)
							πTemp005 = πg.NewDict()
							πTemp003 = πTemp005.ToObject()
							µd = πTemp003
							// line 217: self.assertIsNot(d.fromkeys('abc'), d)
							πF.SetLineno(217)
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
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[1] = µd
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
							// line 218: self.assertEqual(d.fromkeys('abc'), {'a':None, 'b':None, 'c':None})
							πF.SetLineno(218)
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
							// line 219: self.assertEqual(d.fromkeys((4,5),0), {4:0, 5:0})
							πF.SetLineno(219)
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
							// line 220: self.assertEqual(d.fromkeys([]), {})
							πF.SetLineno(220)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp006...).ToObject()
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
							// line 221: def g():
							πF.SetLineno(221)
							πTemp007 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("g", "/usr/lib/python2.7/test/test_dict.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										// line 222: yield 1
										πF.SetLineno(222)
										πF.PushCheckpoint(1)
										return πg.NewInt(1).ToObject(), nil
									Label1:
										πTemp001 = πSent
									}
									return nil, πE
								}).ToObject(), nil
							}), πF.Globals()).ToObject()
							µg = πTemp003
							// line 223: self.assertEqual(d.fromkeys(g()), {1:None})
							πF.SetLineno(223)
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
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp008
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
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 224: self.assertRaises(TypeError, {}.fromkeys, 3)
							πF.SetLineno(224)
							πTemp001 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp005 = πg.NewDict()
							πTemp004 = πTemp005.ToObject()
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp008
							πTemp001[2] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 225: class dictlike(dict): pass
							πF.SetLineno(225)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("dictlike", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 225: class dictlike(dict): pass
									πF.SetLineno(225)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp008 == nil {
								πTemp008 = πg.TypeType.ToObject()
							}
							if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("dictlike").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µdictlike = πTemp009
							// line 226: self.assertEqual(dictlike.fromkeys('a'), {'a':None})
							πF.SetLineno(226)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µdictlike, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp008
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
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 227: self.assertEqual(dictlike().fromkeys('a'), {'a':None})
							πF.SetLineno(227)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							if πTemp004, πE = µdictlike.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
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
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 228: self.assertIsInstance(dictlike.fromkeys('a'), dictlike)
							πF.SetLineno(228)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µdictlike, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp008
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							πTemp001[1] = µdictlike
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 229: self.assertIsInstance(dictlike().fromkeys('a'), dictlike)
							πF.SetLineno(229)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							if πTemp004, πE = µdictlike.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µdictlike, "dictlike"); πE != nil {
								continue
							}
							πTemp001[1] = µdictlike
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 230: class mydict(dict):
							πF.SetLineno(230)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("mydict", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 231: def __new__(cls):
									πF.SetLineno(231)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "cls", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 232: return UserDict.UserDict()
											πF.SetLineno(232)
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
							if πTemp008, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp008 == nil {
								πTemp008 = πg.TypeType.ToObject()
							}
							if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("mydict").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µmydict = πTemp009
							// line 233: ud = mydict.fromkeys('ab')
							πF.SetLineno(233)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßab.ToObject()
							if πE = πg.CheckLocal(πF, µmydict, "mydict"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µmydict, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µud = πTemp008
							// line 234: self.assertEqual(ud, {'a':None, 'b':None})
							πF.SetLineno(234)
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
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 235: self.assertIsInstance(ud, UserDict.UserDict)
							πF.SetLineno(235)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µud, "ud"); πE != nil {
								continue
							}
							πTemp001[0] = µud
							if πTemp004, πE = πg.ResolveGlobal(πF, ßUserDict); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßUserDict, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertIsInstance, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 236: self.assertRaises(TypeError, dict.fromkeys)
							πF.SetLineno(236)
							πTemp001 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 238: class Exc(Exception): pass
							πF.SetLineno(238)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 238: class Exc(Exception): pass
									πF.SetLineno(238)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp008 == nil {
								πTemp008 = πg.TypeType.ToObject()
							}
							if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp009
							// line 240: class baddict1(dict):
							πF.SetLineno(240)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("baddict1", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 241: def __init__(self):
									πF.SetLineno(241)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 242: raise Exc()
											πF.SetLineno(242)
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
							if πTemp008, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp008 == nil {
								πTemp008 = πg.TypeType.ToObject()
							}
							if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("baddict1").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µbaddict1 = πTemp009
							// line 244: self.assertRaises(Exc, baddict1.fromkeys, [1])
							πF.SetLineno(244)
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
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 246: class BadSeq(object):
							πF.SetLineno(246)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadSeq", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 247: def __iter__(self):
									πF.SetLineno(247)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 248: return self
											πF.SetLineno(248)
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
									// line 249: def next(self):
									πF.SetLineno(249)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 250: raise Exc()
											πF.SetLineno(250)
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
							if πTemp008, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp008 == nil {
								πTemp008 = πg.TypeType.ToObject()
							}
							if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("BadSeq").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µBadSeq = πTemp009
							// line 252: self.assertRaises(Exc, dict.fromkeys, BadSeq())
							πF.SetLineno(252)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp008
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
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 254: class baddict2(dict):
							πF.SetLineno(254)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("baddict2", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 255: def __setitem__(self, key, value):
									πF.SetLineno(255)
									πTemp002 = make([]πg.Param, 3)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp002[2] = πg.Param{Name: "value", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 256: raise Exc()
											πF.SetLineno(256)
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
							if πTemp008, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp008 == nil {
								πTemp008 = πg.TypeType.ToObject()
							}
							if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("baddict2").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µbaddict2 = πTemp009
							// line 258: self.assertRaises(Exc, baddict2.fromkeys, [1])
							πF.SetLineno(258)
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
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 261: d = dict(zip(range(6), range(6)))
							πF.SetLineno(261)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(6).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[0] = πTemp008
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(6).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[1] = πTemp008
							if πTemp004, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp008
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp008
							// line 262: self.assertEqual(dict.fromkeys(d, 0), dict(zip(range(6), [0]*6)))
							πF.SetLineno(262)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002[0] = µd
							πTemp002[1] = πg.NewInt(0).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp004, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(1)
							πTemp010[0] = πg.NewInt(6).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp006[0] = πTemp008
							πTemp010 = make([]*πg.Object, 1)
							πTemp010[0] = πg.NewInt(0).ToObject()
							πTemp008 = πg.NewList(πTemp010...).ToObject()
							if πTemp004, πE = πg.Mul(πF, πTemp008, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							πTemp006[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002[0] = πTemp008
							if πTemp004, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 264: class baddict3(dict):
							πF.SetLineno(264)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							πTemp005 = πg.NewDict()
							if πTemp004, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp004); πE != nil {
								continue
							}
							_, πE = πg.NewCode("baddict3", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 265: def __new__(cls):
									πF.SetLineno(265)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "cls", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µcls *πg.Object = πArgs[0]; _ = µcls
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 266: return d
											πF.SetLineno(266)
											if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
												continue
											}
											πR = µd
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
							if πTemp008, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp008 == nil {
								πTemp008 = πg.TypeType.ToObject()
							}
							if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("baddict3").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µbaddict3 = πTemp009
							// line 267: d = {i : i for i in range(10)}
							πF.SetLineno(267)
							πTemp007 = make([]πg.Param, 0)
							πTemp008 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/test_dict.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										πTemp002[0] = πg.NewInt(10).ToObject()
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
										// line 267: d = {i : i for i in range(10)}
										πF.SetLineno(267)
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
											continue
										}
										πTemp003 = πg.NewTuple2(µi, µi).ToObject()
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
							if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.DictType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							µd = πTemp004
							// line 268: res = d.copy()
							πF.SetLineno(268)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							µres = πTemp009
							// line 269: res.update(a=None, b=None, c=None)
							πF.SetLineno(269)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp012 = πg.KWArgs{
								{"a", πTemp004},
								{"b", πTemp009},
								{"c", πTemp011},
							}
							if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µres, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, nil, πTemp012); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_fromkeys.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 213: @unittest.expectedFailure
					πF.SetLineno(213)
					πTemp004 = πF.MakeArgs(1)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßtest_fromkeys); πE != nil {
						continue
					}
					πTemp004[0] = πTemp016
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp017.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_fromkeys.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 273: def test_copy(self):
					πF.SetLineno(273)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_copy", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Dict
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
							// line 274: d = {1:1, 2:2, 3:3}
							πF.SetLineno(274)
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
							µd = πTemp002
							// line 275: self.assertEqual(d.copy(), {1:1, 2:2, 3:3})
							πF.SetLineno(275)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
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
							// line 276: self.assertEqual({}.copy(), {})
							πF.SetLineno(276)
							πTemp003 = πF.MakeArgs(2)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
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
							// line 277: self.assertRaises(TypeError, d.copy, None)
							πF.SetLineno(277)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_copy.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 272: @unittest.expectedFailure
					πF.SetLineno(272)
					πTemp004 = πF.MakeArgs(1)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßtest_copy); πE != nil {
						continue
					}
					πTemp004[0] = πTemp017
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp017, πE = πTemp018.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_copy.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 279: def test_get(self):
					πF.SetLineno(279)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("test_get", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
							// line 280: d = {}
							πF.SetLineno(280)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 281: self.assertIs(d.get('c'), None)
							πF.SetLineno(281)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßc.ToObject()
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
							πTemp003[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 282: self.assertEqual(d.get('c', 3), 3)
							πF.SetLineno(282)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßc.ToObject()
							πTemp004[1] = πg.NewInt(3).ToObject()
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
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewInt(3).ToObject()
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
							// line 283: d = {'a': 1, 'b': 2}
							πF.SetLineno(283)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßa.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 284: self.assertIs(d.get('c'), None)
							πF.SetLineno(284)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßc.ToObject()
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
							πTemp003[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 285: self.assertEqual(d.get('c', 3), 3)
							πF.SetLineno(285)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßc.ToObject()
							πTemp004[1] = πg.NewInt(3).ToObject()
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
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewInt(3).ToObject()
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
							// line 286: self.assertEqual(d.get('a'), 1)
							πF.SetLineno(286)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßa.ToObject()
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
							// line 287: self.assertEqual(d.get('a', 3), 1)
							πF.SetLineno(287)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßa.ToObject()
							πTemp004[1] = πg.NewInt(3).ToObject()
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
							// line 288: self.assertRaises(TypeError, d.get)
							πF.SetLineno(288)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 289: self.assertRaises(TypeError, d.get, None, None, None)
							πF.SetLineno(289)
							πTemp003 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßget, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_get.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 291: def test_setdefault(self):
					πF.SetLineno(291)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("test_setdefault", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadHash *πg.Object = πg.UnboundLocal; _ = µBadHash
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 293: d = {}
							πF.SetLineno(293)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 294: self.assertIs(d.setdefault('key0'), None)
							πF.SetLineno(294)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßkey0.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 295: d.setdefault('key0', [])
							πF.SetLineno(295)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßkey0.ToObject()
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 296: self.assertIs(d.setdefault('key0'), None)
							πF.SetLineno(296)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßkey0.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 297: d.setdefault('key', []).append(3)
							πF.SetLineno(297)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(3).ToObject()
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
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 298: self.assertEqual(d['key'][0], 3)
							πF.SetLineno(298)
							πTemp003 = πF.MakeArgs(2)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp007 = ßkey.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µd, πTemp007); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewInt(3).ToObject()
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
							// line 299: d.setdefault('key', []).append(4)
							πF.SetLineno(299)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(4).ToObject()
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
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.GetAttr(πF, πTemp005, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 300: self.assertEqual(len(d['key']), 2)
							πF.SetLineno(300)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp002 = ßkey.ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µd, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 301: self.assertRaises(TypeError, d.setdefault)
							πF.SetLineno(301)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 303: class Exc(Exception): pass
							πF.SetLineno(303)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp003[0] = πTemp007
							πTemp001 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 303: class Exc(Exception): pass
									πF.SetLineno(303)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp005 == nil {
								πTemp005 = πg.TypeType.ToObject()
							}
							if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp007
							// line 305: class BadHash(object):
							πF.SetLineno(305)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp003[0] = πTemp007
							πTemp001 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadHash", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 306: fail = False
									πF.SetLineno(306)
									if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßfail.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 307: def __hash__(self):
									πF.SetLineno(307)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 308: if self.fail:
											πF.SetLineno(308)
										Label1:
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 309: raise Exc()
											πF.SetLineno(309)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label3
										Label2:
											// line 311: return 42
											πF.SetLineno(311)
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
							if πTemp005, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp005 == nil {
								πTemp005 = πg.TypeType.ToObject()
							}
							if πTemp007, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("BadHash").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µBadHash = πTemp007
							// line 313: x = BadHash()
							πF.SetLineno(313)
							if πE = πg.CheckLocal(πF, µBadHash, "BadHash"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadHash.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp002
							// line 314: d[x] = 42
							πF.SetLineno(314)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(42).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp005 = µx
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 315: x.fail = True
							πF.SetLineno(315)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µx, ßfail, πTemp005); πE != nil {
								continue
							}
							// line 316: self.assertRaises(Exc, d.setdefault, x, [])
							πF.SetLineno(316)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp003[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßsetdefault, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003[2] = µx
							πTemp004 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp003[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_setdefault.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 318: def test_setdefault_atomic(self):
					πF.SetLineno(318)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("test_setdefault_atomic", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µHashed *πg.Object = πg.UnboundLocal; _ = µHashed
						var µhashed1 *πg.Object = πg.UnboundLocal; _ = µhashed1
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µhashed2 *πg.Object = πg.UnboundLocal; _ = µhashed2
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
							// line 320: class Hashed(object):
							πF.SetLineno(320)
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
							_, πE = πg.NewCode("Hashed", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 321: def __init__(self):
									πF.SetLineno(321)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 322: self.hash_count = 0
											πF.SetLineno(322)
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßhash_count, πTemp001); πE != nil {
												continue
											}
											// line 323: self.eq_count = 0
											πF.SetLineno(323)
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßeq_count, πTemp001); πE != nil {
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
									// line 324: def __hash__(self):
									πF.SetLineno(324)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 325: self.hash_count += 1
											πF.SetLineno(325)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ßhash_count, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßhash_count, πTemp002); πE != nil {
												continue
											}
											// line 326: return 42
											πF.SetLineno(326)
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
									// line 327: def __eq__(self, other):
									πF.SetLineno(327)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp004 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
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
											// line 328: self.eq_count += 1
											πF.SetLineno(328)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ßeq_count, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßeq_count, πTemp002); πE != nil {
												continue
											}
											// line 329: return id(self) == id(other)
											πF.SetLineno(329)
											πTemp003 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											πTemp003[0] = µself
											if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
												continue
											}
											if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp003)
											πTemp003 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
												continue
											}
											πTemp003[0] = µother
											if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
												continue
											}
											if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp003)
											if πTemp001, πE = πg.Eq(πF, πTemp004, πTemp005); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp004); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Hashed").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µHashed = πTemp005
							// line 330: hashed1 = Hashed()
							πF.SetLineno(330)
							if πE = πg.CheckLocal(πF, µHashed, "Hashed"); πE != nil {
								continue
							}
							if πTemp002, πE = µHashed.Call(πF, nil, nil); πE != nil {
								continue
							}
							µhashed1 = πTemp002
							// line 331: y = {hashed1: 5}
							πF.SetLineno(331)
							πTemp001 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µhashed1, "hashed1"); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, µhashed1, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µy = πTemp002
							// line 332: hashed2 = Hashed()
							πF.SetLineno(332)
							if πE = πg.CheckLocal(πF, µHashed, "Hashed"); πE != nil {
								continue
							}
							if πTemp002, πE = µHashed.Call(πF, nil, nil); πE != nil {
								continue
							}
							µhashed2 = πTemp002
							// line 333: y.setdefault(hashed2, [])
							πF.SetLineno(333)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhashed2, "hashed2"); πE != nil {
								continue
							}
							πTemp003[0] = µhashed2
							πTemp006 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp006...).ToObject()
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µy, ßsetdefault, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 334: self.assertEqual(hashed1.hash_count, 1)
							πF.SetLineno(334)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhashed1, "hashed1"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µhashed1, ßhash_count, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
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
							// line 335: self.assertEqual(hashed2.hash_count, 1)
							πF.SetLineno(335)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhashed2, "hashed2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µhashed2, ßhash_count, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
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
							// line 336: self.assertEqual(hashed1.eq_count + hashed2.eq_count, 1)
							πF.SetLineno(336)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhashed1, "hashed1"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µhashed1, ßeq_count, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhashed2, "hashed2"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µhashed2, ßeq_count, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_setdefault_atomic.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 339: def test_popitem(self):
					πF.SetLineno(339)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("test_popitem", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Dict
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.Object
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
							case 14: goto Label14
							case 15: goto Label15
							default: panic("unexpected function state")
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πg.NewInt(1).ToObject()).ToObject()
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
								µcopymode = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(12).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Iter(πF, πTemp007); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µlog2size = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 346: size = 2**log2size
							πF.SetLineno(346)
							if πE = πg.CheckLocal(πF, µlog2size, "log2size"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Pow(πF, πg.NewInt(2).ToObject(), µlog2size); πE != nil {
								continue
							}
							µsize = πTemp003
							// line 347: a = {}
							πF.SetLineno(347)
							πTemp009 = πg.NewDict()
							πTemp003 = πTemp009.ToObject()
							µa = πTemp003
							// line 348: b = {}
							πF.SetLineno(348)
							πTemp009 = πg.NewDict()
							πTemp003 = πTemp009.ToObject()
							µb = πTemp003
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp006[0] = µsize
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.Iter(πF, πTemp010); πE != nil {
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
							if πTemp007, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µi = πTemp007
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 350: a[repr(i)] = i
							πF.SetLineno(350)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, µi); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp006[0] = µi
							if πTemp012, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp010 = πTemp013
							if πE = πg.SetItem(πF, µa, πTemp010, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcopymode, "copymode"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.LT(πF, µcopymode, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label10
							}
							goto Label11
							// line 351: if copymode < 0:
							πF.SetLineno(351)
						Label10:
							// line 352: b[repr(i)] = i
							πF.SetLineno(352)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, µi); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp006[0] = µi
							if πTemp012, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp010 = πTemp013
							if πE = πg.SetItem(πF, µb, πTemp010, πTemp007); πE != nil {
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
							if πTemp003, πE = πg.GT(πF, µcopymode, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label12
							}
							goto Label13
							// line 353: if copymode > 0:
							πF.SetLineno(353)
						Label12:
							// line 354: b = a.copy()
							πF.SetLineno(354)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µb = πTemp007
							goto Label13
						Label13:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp006[0] = µsize
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.Iter(πF, πTemp010); πE != nil {
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
							if πTemp007, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µi = πTemp007
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(14)            
							// line 356: ka, va = ta = a.popitem()
							πF.SetLineno(356)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µa, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp012}}}, πTemp010); πE != nil {
								continue
							}
							µka = πTemp007
							µva = πTemp012
							µta = πTemp010
							// line 357: self.assertEqual(va, int(ka))
							πF.SetLineno(357)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µva, "va"); πE != nil {
								continue
							}
							πTemp006[0] = µva
							πTemp014 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µka, "ka"); πE != nil {
								continue
							}
							πTemp014[0] = µka
							if πTemp007, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp006[1] = πTemp010
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 358: kb, vb = tb = b.popitem()
							πF.SetLineno(358)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µb, ßpopitem, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp012}}}, πTemp010); πE != nil {
								continue
							}
							µkb = πTemp007
							µvb = πTemp012
							µtb = πTemp010
							// line 359: self.assertEqual(vb, int(kb))
							πF.SetLineno(359)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvb, "vb"); πE != nil {
								continue
							}
							πTemp006[0] = µvb
							πTemp014 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkb, "kb"); πE != nil {
								continue
							}
							πTemp014[0] = µkb
							if πTemp007, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp014, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp014)
							πTemp006[1] = πTemp010
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 360: self.assertFalse(copymode < 0 and ta != tb)
							πF.SetLineno(360)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcopymode, "copymode"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.LT(πF, µcopymode, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp010
							if πTemp011, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if !πTemp011 {
								goto Label17
							}
							if πE = πg.CheckLocal(πF, µta, "ta"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtb, "tb"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.NE(πF, µta, µtb); πE != nil {
								continue
							}
							πTemp007 = πTemp010
						Label17:
							πTemp006[0] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
						Label16:
							// line 361: self.assertFalse(a)
							πF.SetLineno(361)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 362: self.assertFalse(b)
							πF.SetLineno(362)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp006[0] = µb
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 364: d = {}
							πF.SetLineno(364)
							πTemp009 = πg.NewDict()
							πTemp001 = πTemp009.ToObject()
							µd = πTemp001
							// line 365: self.assertRaises(KeyError, d.popitem)
							πF.SetLineno(365)
							πTemp006 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßpopitem, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
					// line 338: @unittest.expectedFailure
					πF.SetLineno(338)
					πTemp004 = πF.MakeArgs(1)
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßtest_popitem); πE != nil {
						continue
					}
					πTemp004[0] = πTemp021
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp022, πE = πg.GetAttr(πF, πTemp021, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp021, πE = πTemp022.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_popitem.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 368: def test_pop(self):
					πF.SetLineno(368)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("test_pop", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var µv *πg.Object = πg.UnboundLocal; _ = µv
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µh *πg.Object = πg.UnboundLocal; _ = µh
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadHash *πg.Object = πg.UnboundLocal; _ = µBadHash
						var πTemp001 *πg.Dict
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
							// line 370: d = {}
							πF.SetLineno(370)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 371: k, v = 'abc', 'def'
							πF.SetLineno(371)
							πTemp002 = πg.NewTuple2(ßabc.ToObject(), ßdef.ToObject()).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
								continue
							}
							µk = πTemp003
							µv = πTemp004
							// line 372: d[k] = v
							πF.SetLineno(372)
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
							// line 373: self.assertRaises(KeyError, d.pop, 'ghi')
							πF.SetLineno(373)
							πTemp005 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							πTemp005[2] = ßghi.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 375: self.assertEqual(d.pop(k), v)
							πF.SetLineno(375)
							πTemp005 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp006[0] = µk
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
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp005[1] = µv
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 376: self.assertEqual(len(d), 0)
							πF.SetLineno(376)
							πTemp005 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp006[0] = µd
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp003
							πTemp005[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 378: self.assertRaises(KeyError, d.pop, k)
							πF.SetLineno(378)
							πTemp005 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp005[2] = µk
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 382: x = 4503599627370496L
							πF.SetLineno(382)
							µx = πg.NewLongFromBytes([]byte{0x10,0x0,0x0,0x0,0x0,0x0,0x0,}).ToObject()
							// line 383: y = 4503599627370496
							πF.SetLineno(383)
							µy = πg.NewInt(4503599627370496).ToObject()
							// line 384: h = {x: 'anything', y: 'something else'}
							πF.SetLineno(384)
							πTemp001 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, µx, ßanything.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, µy, πg.NewStr("something else").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µh = πTemp002
							// line 385: self.assertEqual(h[x], h[y])
							πF.SetLineno(385)
							πTemp005 = πF.MakeArgs(2)
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
							πTemp005[0] = πTemp003
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
							πTemp005[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 387: self.assertEqual(d.pop(k, v), v)
							πF.SetLineno(387)
							πTemp005 = πF.MakeArgs(2)
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
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp005[1] = µv
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 388: d[k] = v
							πF.SetLineno(388)
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
							// line 389: self.assertEqual(d.pop(k, 1), v)
							πF.SetLineno(389)
							πTemp005 = πF.MakeArgs(2)
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
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp005[1] = µv
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 391: self.assertRaises(TypeError, d.pop)
							πF.SetLineno(391)
							πTemp005 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 393: class Exc(Exception): pass
							πF.SetLineno(393)
							πTemp005 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							πTemp001 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 393: class Exc(Exception): pass
									πF.SetLineno(393)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp003 == nil {
								πTemp003 = πg.TypeType.ToObject()
							}
							if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp005...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp004
							// line 395: class BadHash(object):
							πF.SetLineno(395)
							πTemp005 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							πTemp001 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadHash", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 396: fail = False
									πF.SetLineno(396)
									if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßfail.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 397: def __hash__(self):
									πF.SetLineno(397)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 398: if self.fail:
											πF.SetLineno(398)
										Label1:
											if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
												continue
											}
											if πTemp001, πE = µExc.Call(πF, nil, nil); πE != nil {
												continue
											}
											// line 399: raise Exc()
											πF.SetLineno(399)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label3
										Label2:
											// line 401: return 42
											πF.SetLineno(401)
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
							if πTemp003, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp003 == nil {
								πTemp003 = πg.TypeType.ToObject()
							}
							if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadHash").ToObject(), πg.NewTuple(πTemp005...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µBadHash = πTemp004
							// line 403: x = BadHash()
							πF.SetLineno(403)
							if πE = πg.CheckLocal(πF, µBadHash, "BadHash"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadHash.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp002
							// line 404: d[x] = 42
							πF.SetLineno(404)
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
							// line 405: x.fail = True
							πF.SetLineno(405)
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
							// line 406: self.assertRaises(Exc, d.pop, x)
							πF.SetLineno(406)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp005[0] = µExc
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßpop, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp005[2] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
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
					// line 367: @unittest.expectedFailure
					πF.SetLineno(367)
					πTemp004 = πF.MakeArgs(1)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßtest_pop); πE != nil {
						continue
					}
					πTemp004[0] = πTemp022
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp023, πE = πg.GetAttr(πF, πTemp022, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp022, πE = πTemp023.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_pop.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 408: def test_mutatingiteration(self):
					πF.SetLineno(408)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("test_mutatingiteration", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 *πg.Dict
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
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
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
							case 2: goto Label2
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							// line 410: d = {}
							πF.SetLineno(410)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 411: d[1] = 1
							πF.SetLineno(411)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp003, πTemp002); πE != nil {
								continue
							}
							// line 412: with self.assertRaises(RuntimeError):
							πF.SetLineno(412)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp003}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Iter(πF, µd); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							πTemp007 = false
						Label2:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label4
							}
							if πTemp009, πE = πg.Next(πF, πTemp006); πE != nil {
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
								µi = πTemp009
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(2)            
							// line 414: d[i+1] = 1
							πF.SetLineno(414)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp010 = πTemp011
							if πE = πg.SetItem(πF, µd, πTemp010, πTemp009); πE != nil {
								continue
							}
							continue
						Label3:
							if πE != nil || πR != nil {
								continue
							}
						Label4:
							πF.PopCheckpoint()
						Label1:
							πTemp012, πTemp013 = nil, nil
							if πE != nil {
								πTemp012, πTemp013 = πF.ExcInfo()
							}
							if πTemp012 != nil {
								πTemp014 = πTemp012.Type()
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp014.ToObject(), πTemp012.ToObject(), πTemp013.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp012 != nil && πTemp007 != true {
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
					if πE = πClass.SetItem(πF, ßtest_mutatingiteration.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 416: def test_repr(self):
					πF.SetLineno(416)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("test_repr", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadRepr *πg.Object = πg.UnboundLocal; _ = µBadRepr
						var πTemp001 *πg.Dict
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
							// line 417: d = {}
							πF.SetLineno(417)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 418: self.assertEqual(repr(d), '{}')
							πF.SetLineno(418)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewStr("{}").ToObject()
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
							// line 419: d[1] = 2
							πF.SetLineno(419)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 420: self.assertEqual(repr(d), '{1: 2}')
							πF.SetLineno(420)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewStr("{1: 2}").ToObject()
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
							// line 421: d = {}
							πF.SetLineno(421)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 422: d[1] = d
							πF.SetLineno(422)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µd); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 423: self.assertEqual(repr(d), '{1: {...}}')
							πF.SetLineno(423)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp005
							πTemp003[1] = πg.NewStr("{1: {...}}").ToObject()
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
							// line 425: class Exc(Exception): pass
							πF.SetLineno(425)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp003[0] = πTemp006
							πTemp001 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 425: class Exc(Exception): pass
									πF.SetLineno(425)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp005 == nil {
								πTemp005 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp006
							// line 427: class BadRepr(object):
							πF.SetLineno(427)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp003[0] = πTemp006
							πTemp001 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadRepr", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 428: def __repr__(self):
									πF.SetLineno(428)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 429: raise Exc()
											πF.SetLineno(429)
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
							if πTemp005, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp005 == nil {
								πTemp005 = πg.TypeType.ToObject()
							}
							if πTemp006, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("BadRepr").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µBadRepr = πTemp006
							// line 431: d = {1: BadRepr()}
							πF.SetLineno(431)
							πTemp001 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µBadRepr, "BadRepr"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadRepr.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πTemp001.SetItem(πF, πg.NewInt(1).ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 432: self.assertRaises(Exc, repr, d)
							πF.SetLineno(432)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp003[0] = µExc
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[2] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_repr.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 435: def test_le(self):
					πF.SetLineno(435)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("test_le", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µExc *πg.Object = πg.UnboundLocal; _ = µExc
						var µBadCmp *πg.Object = πg.UnboundLocal; _ = µBadCmp
						var µd1 *πg.Object = πg.UnboundLocal; _ = µd1
						var µd2 *πg.Object = πg.UnboundLocal; _ = µd2
						var πTemp001 []*πg.Object
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
							// line 436: self.assertFalse({} < {})
							πF.SetLineno(436)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewDict()
							πTemp004 = πTemp003.ToObject()
							πTemp003 = πg.NewDict()
							πTemp005 = πTemp003.ToObject()
							if πTemp002, πE = πg.LT(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 437: self.assertFalse({1: 2} < {1L: 2L})
							πF.SetLineno(437)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πg.NewDict()
							if πE = πTemp003.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp003.ToObject()
							πTemp003 = πg.NewDict()
							if πE = πTemp003.SetItem(πF, πg.NewLongFromBytes([]byte{0x1,}).ToObject(), πg.NewLongFromBytes([]byte{0x2,}).ToObject()); πE != nil {
								continue
							}
							πTemp005 = πTemp003.ToObject()
							if πTemp002, πE = πg.LT(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 439: class Exc(Exception): pass
							πF.SetLineno(439)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp003 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("Exc", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp003
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 439: class Exc(Exception): pass
									πF.SetLineno(439)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Exc").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
								continue
							}
							µExc = πTemp005
							// line 441: class BadCmp(object):
							πF.SetLineno(441)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							πTemp003 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BadCmp", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp003
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
									// line 442: def __eq__(self, other):
									πF.SetLineno(442)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 443: raise Exc()
											πF.SetLineno(443)
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
									// line 444: def __hash__(self):
									πF.SetLineno(444)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 445: return 42
											πF.SetLineno(445)
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
							if πTemp004, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadCmp").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
								continue
							}
							µBadCmp = πTemp005
							// line 447: d1 = {BadCmp(): 1}
							πF.SetLineno(447)
							πTemp003 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µBadCmp, "BadCmp"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadCmp.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πTemp003.SetItem(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003.ToObject()
							µd1 = πTemp002
							// line 448: d2 = {1: 1}
							πF.SetLineno(448)
							πTemp003 = πg.NewDict()
							if πE = πTemp003.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003.ToObject()
							µd2 = πTemp002
							// line 450: with self.assertRaises(Exc):
							πF.SetLineno(450)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µExc, "Exc"); πE != nil {
								continue
							}
							πTemp001[0] = µExc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							// line 451: d1 < d2
							πF.SetLineno(451)
							if πE = πg.CheckLocal(πF, µd1, "d1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd2, "d2"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LT(πF, µd1, µd2); πE != nil {
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
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp004, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp004, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_le.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 434: @unittest.expectedFailure
					πF.SetLineno(434)
					πTemp004 = πF.MakeArgs(1)
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßtest_le); πE != nil {
						continue
					}
					πTemp004[0] = πTemp025
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp026, πE = πg.GetAttr(πF, πTemp025, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp025, πE = πTemp026.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_le.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 454: def test_missing(self):
					πF.SetLineno(454)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("test_missing", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µD *πg.Object = πg.UnboundLocal; _ = µD
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µE *πg.Object = πg.UnboundLocal; _ = µE
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µc *πg.Object = πg.UnboundLocal; _ = µc
						var µF *πg.Object = πg.UnboundLocal; _ = µF
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µG *πg.Object = πg.UnboundLocal; _ = µG
						var µg *πg.Object = πg.UnboundLocal; _ = µg
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
							// line 456: self.assertFalse(hasattr(dict, "__missing__"))
							πF.SetLineno(456)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = ß__missing__.ToObject()
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 457: self.assertFalse(hasattr({}, "__missing__"))
							πF.SetLineno(457)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πg.NewDict()
							πTemp003 = πTemp005.ToObject()
							πTemp002[0] = πTemp003
							πTemp002[1] = ß__missing__.ToObject()
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 463: class D(dict):
							πF.SetLineno(463)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
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
							_, πE = πg.NewCode("D", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 464: def __missing__(self, key):
									πF.SetLineno(464)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__missing__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µkey *πg.Object = πArgs[1]; _ = µkey
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 465: return 42
											πF.SetLineno(465)
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
									if πE = πClass.SetItem(πF, ß__missing__.ToObject(), πTemp001); πE != nil {
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
							if πTemp006, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("D").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µD = πTemp006
							// line 466: d = D({1: 2, 3: 4})
							πF.SetLineno(466)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πg.NewDict()
							if πE = πTemp005.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, πg.NewInt(3).ToObject(), πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005.ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µD, "D"); πE != nil {
								continue
							}
							if πTemp003, πE = µD.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 467: self.assertEqual(d[1], 2)
							πF.SetLineno(467)
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µd, πTemp003); πE != nil {
								continue
							}
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
							// line 468: self.assertEqual(d[3], 4)
							πF.SetLineno(468)
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µd, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(4).ToObject()
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
							// line 469: self.assertNotIn(2, d)
							πF.SetLineno(469)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[1] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 470: self.assertNotIn(2, d.keys())
							πF.SetLineno(470)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µd, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertNotIn, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 471: self.assertEqual(d[2], 42)
							πF.SetLineno(471)
							πTemp001 = πF.MakeArgs(2)
							πTemp003 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µd, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp001[1] = πg.NewInt(42).ToObject()
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
							// line 473: class E(dict):
							πF.SetLineno(473)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
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
							_, πE = πg.NewCode("E", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 474: def __missing__(self, key):
									πF.SetLineno(474)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__missing__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µkey *πg.Object = πArgs[1]; _ = µkey
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
											πTemp001 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
												continue
											}
											πTemp001[0] = µkey
											if πTemp002, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp001)
											// line 475: raise RuntimeError(key)
											πF.SetLineno(475)
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
									if πE = πClass.SetItem(πF, ß__missing__.ToObject(), πTemp001); πE != nil {
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
							if πTemp006, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("E").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µE = πTemp006
							// line 476: e = E()
							πF.SetLineno(476)
							if πE = πg.CheckLocal(πF, µE, "E"); πE != nil {
								continue
							}
							if πTemp003, πE = µE.Call(πF, nil, nil); πE != nil {
								continue
							}
							µe = πTemp003
							// line 477: with self.assertRaises(RuntimeError) as c:
							πF.SetLineno(477)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
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
							if πTemp003, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp006.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							µc = πTemp006
							// line 478: e[42]
							πF.SetLineno(478)
							πTemp007 = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µe, πTemp007); πE != nil {
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
								if πTemp007, πE = πTemp003.Call(πF, πg.Args{πTemp004, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp007, πE = πTemp003.Call(πF, πg.Args{πTemp004, πg.None, πg.None, πg.None}, nil); πE != nil {
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
							// line 479: self.assertEqual(c.exception.args, (42,))
							πF.SetLineno(479)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µc, ßexception, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßargs, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							πTemp007 = πg.NewTuple1(πg.NewInt(42).ToObject()).ToObject()
							πTemp001[1] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 481: class F(dict):
							πF.SetLineno(481)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp013, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001[0] = πTemp013
							πTemp005 = πg.NewDict()
							if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
								continue
							}
							_, πE = πg.NewCode("F", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 482: def __init__(self):
									πF.SetLineno(482)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 []πg.Param
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
											// line 484: self.__missing__ = lambda key: None
											πF.SetLineno(484)
											πTemp002 = make([]πg.Param, 1)
											πTemp002[0] = πg.Param{Name: "key", Def: nil}
											πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
												var µkey *πg.Object = πArgs[0]; _ = µkey
												var πTemp001 *πg.Object
												_ = πTemp001
												var πR *πg.Object; _ = πR
												var πE *πg.BaseException; _ = πE
												for ; πF.State() >= 0; πF.PopCheckpoint() {
													switch πF.State() {
													case 0:
													default: panic("unexpected function state")
													}
													// line 484: self.__missing__ = lambda key: None
													πF.SetLineno(484)
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
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ß__missing__, πTemp003); πE != nil {
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
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp008 == nil {
								πTemp008 = πg.TypeType.ToObject()
							}
							if πTemp013, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("F").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µF = πTemp013
							// line 485: f = F()
							πF.SetLineno(485)
							if πE = πg.CheckLocal(πF, µF, "F"); πE != nil {
								continue
							}
							if πTemp007, πE = µF.Call(πF, nil, nil); πE != nil {
								continue
							}
							µf = πTemp007
							// line 486: with self.assertRaises(KeyError) as c:
							πF.SetLineno(486)
							πTemp001 = πF.MakeArgs(1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp007, πE = πg.GetAttr(πF, πTemp008.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp008.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp013.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							µc = πTemp013
							// line 487: f[42]
							πF.SetLineno(487)
							πTemp014 = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetItem(πF, µf, πTemp014); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label2:
							πTemp010, πTemp011 = nil, nil
							if πE != nil {
								πTemp010, πTemp011 = πF.ExcInfo()
							}
							if πTemp010 != nil {
								πTemp012 = πTemp010.Type()
								if πTemp014, πE = πTemp007.Call(πF, πg.Args{πTemp008, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp014, πE = πTemp007.Call(πF, πg.Args{πTemp008, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp014); πE != nil {
								continue
							}
							if πTemp010 != nil && πTemp009 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 488: self.assertEqual(c.exception.args, (42,))
							πF.SetLineno(488)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, µc, ßexception, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßargs, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp015
							πTemp014 = πg.NewTuple1(πg.NewInt(42).ToObject()).ToObject()
							πTemp001[1] = πTemp014
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp014.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 490: class G(dict):
							πF.SetLineno(490)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp016, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001[0] = πTemp016
							πTemp005 = πg.NewDict()
							if πTemp014, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp014); πE != nil {
								continue
							}
							_, πE = πg.NewCode("G", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 491: pass
									πF.SetLineno(491)
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp015, πE = πTemp005.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp015 == nil {
								πTemp015 = πg.TypeType.ToObject()
							}
							if πTemp016, πE = πTemp015.Call(πF, []*πg.Object{πg.NewStr("G").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µG = πTemp016
							// line 492: g = G()
							πF.SetLineno(492)
							if πE = πg.CheckLocal(πF, µG, "G"); πE != nil {
								continue
							}
							if πTemp014, πE = µG.Call(πF, nil, nil); πE != nil {
								continue
							}
							µg = πTemp014
							// line 493: with self.assertRaises(KeyError) as c:
							πF.SetLineno(493)
							πTemp001 = πF.MakeArgs(1)
							if πTemp014, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp014
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp014.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp014, πE = πg.GetAttr(πF, πTemp015.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetAttr(πF, πTemp015.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp016.Call(πF, πg.Args{πTemp015}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							µc = πTemp016
							// line 494: g[42]
							πF.SetLineno(494)
							πTemp017 = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
								continue
							}
							if πTemp018, πE = πg.GetItem(πF, µg, πTemp017); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label3:
							πTemp010, πTemp011 = nil, nil
							if πE != nil {
								πTemp010, πTemp011 = πF.ExcInfo()
							}
							if πTemp010 != nil {
								πTemp012 = πTemp010.Type()
								if πTemp017, πE = πTemp014.Call(πF, πg.Args{πTemp015, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp017, πE = πTemp014.Call(πF, πg.Args{πTemp015, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp017); πE != nil {
								continue
							}
							if πTemp010 != nil && πTemp009 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 495: self.assertEqual(c.exception.args, (42,))
							πF.SetLineno(495)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp017, πE = πg.GetAttr(πF, µc, ßexception, nil); πE != nil {
								continue
							}
							if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßargs, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp018
							πTemp017 = πg.NewTuple1(πg.NewInt(42).ToObject()).ToObject()
							πTemp001[1] = πTemp017
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp017, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp018, πE = πTemp017.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_missing.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 453: @unittest.expectedFailure
					πF.SetLineno(453)
					πTemp004 = πF.MakeArgs(1)
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßtest_missing); πE != nil {
						continue
					}
					πTemp004[0] = πTemp026
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp027, πE = πg.GetAttr(πF, πTemp026, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp026, πE = πTemp027.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_missing.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 498: def test_tuple_keyerror(self):
					πF.SetLineno(498)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("test_tuple_keyerror", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
							// line 500: d = {}
							πF.SetLineno(500)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 501: with self.assertRaises(KeyError) as c:
							πF.SetLineno(501)
							πTemp003 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
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
							if πTemp002, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp005.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							µc = πTemp005
							// line 502: d[(1,)]
							πF.SetLineno(502)
							πTemp007 = πg.NewTuple1(πg.NewInt(1).ToObject()).ToObject()
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µd, πTemp006); πE != nil {
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
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp004, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp004, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp009 != nil && πTemp008 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 503: self.assertEqual(c.exception.args, ((1,),))
							πF.SetLineno(503)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µc, ßexception, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßargs, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp007
							πTemp007 = πg.NewTuple1(πg.NewInt(1).ToObject()).ToObject()
							πTemp006 = πg.NewTuple1(πTemp007).ToObject()
							πTemp003[1] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_tuple_keyerror.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 497: @unittest.expectedFailure
					πF.SetLineno(497)
					πTemp004 = πF.MakeArgs(1)
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßtest_tuple_keyerror); πE != nil {
						continue
					}
					πTemp004[0] = πTemp027
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp028, πE = πg.GetAttr(πF, πTemp027, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp027, πE = πTemp028.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_tuple_keyerror.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 534: def test_resize1(self):
					πF.SetLineno(534)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("test_resize1", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µi *πg.Object = πg.UnboundLocal; _ = µi
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
							// line 542: d = {}
							πF.SetLineno(542)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(5).ToObject()
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
							// line 544: d[i] = i
							πF.SetLineno(544)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µi); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = µi
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp004); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(5).ToObject()
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
								µi = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 546: del d[i]
							πF.SetLineno(546)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = µi
							if πE = πg.DelItem(πF, µd, πTemp004); πE != nil {
								continue
							}
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewInt(5).ToObject()
							πTemp003[1] = πg.NewInt(9).ToObject()
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
								µi = πTemp004
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 548: d[i] = i
							πF.SetLineno(548)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, µi); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = µi
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp004); πE != nil {
								continue
							}
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_resize1.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 550: def test_resize2(self):
					πF.SetLineno(550)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("test_resize2", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µX *πg.Object = πg.UnboundLocal; _ = µX
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µresizing *πg.Object = πg.UnboundLocal; _ = µresizing
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
							// line 554: class X(object):
							πF.SetLineno(554)
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
							_, πE = πg.NewCode("X", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 555: def __hash__(self):
									πF.SetLineno(555)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 556: return 5
											πF.SetLineno(556)
											πR = πg.NewInt(5).ToObject()
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
									// line 557: def __eq__(self, other):
									πF.SetLineno(557)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πTemp001 bool
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
											if πE = πg.CheckLocal(πF, µresizing, "resizing"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.IsTrue(πF, µresizing); πE != nil {
												continue
											}
											if πTemp001 {
												goto Label1
											}
											goto Label2
											// line 558: if resizing:
											πF.SetLineno(558)
										Label1:
											// line 559: d.clear()
											πF.SetLineno(559)
											if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µd, ßclear, nil); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
												continue
											}
											goto Label2
										Label2:
											// line 560: return False
											πF.SetLineno(560)
											if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp003); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("X").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µX = πTemp005
							// line 561: d = {}
							πF.SetLineno(561)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 562: resizing = False
							πF.SetLineno(562)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µresizing = πTemp002
							// line 563: d[X()] = 1
							πF.SetLineno(563)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp005, πE = µX.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 564: d[X()] = 2
							πF.SetLineno(564)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp005, πE = µX.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 565: d[X()] = 3
							πF.SetLineno(565)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp005, πE = µX.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 566: d[X()] = 4
							πF.SetLineno(566)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp005, πE = µX.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 567: d[X()] = 5
							πF.SetLineno(567)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp005, πE = µX.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 569: resizing = True
							πF.SetLineno(569)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µresizing = πTemp002
							// line 570: d[9] = 6
							πF.SetLineno(570)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(9).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_resize2.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 572: def test_empty_presized_dict_in_freelist(self):
					πF.SetLineno(572)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("test_empty_presized_dict_in_freelist", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd *πg.Object = πg.UnboundLocal; _ = µd
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
							// line 575: with self.assertRaises(ZeroDivisionError):
							πF.SetLineno(575)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßZeroDivisionError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							// line 576: d = {'a': 1 // 0, 'b': None, 'c': None, 'd': None, 'e': None,
							πF.SetLineno(576)
							πTemp005 = πg.NewDict()
							if πTemp006, πE = πg.FloorDiv(πF, πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßa.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßb.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßc.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßd.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ße.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßf.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßg.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ßh.ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp006 = πTemp005.ToObject()
							µd = πTemp006
							πF.PopCheckpoint()
						Label1:
							πTemp008, πTemp009 = nil, nil
							if πE != nil {
								πTemp008, πTemp009 = πF.ExcInfo()
							}
							if πTemp008 != nil {
								πTemp010 = πTemp008.Type()
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp008 != nil && πTemp007 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 578: d = {}
							πF.SetLineno(578)
							πTemp005 = πg.NewDict()
							πTemp006 = πTemp005.ToObject()
							µd = πTemp006
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_empty_presized_dict_in_freelist.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 607: def test_track_literals(self):
					πF.SetLineno(607)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("test_track_literals", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µz *πg.Object = πg.UnboundLocal; _ = µz
						var µw *πg.Object = πg.UnboundLocal; _ = µw
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
						var πTemp007 *πg.Dict
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.Dict
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 609: x, y, z, w = 1.5, "a", (1, None), []
							πF.SetLineno(609)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewInt(1).ToObject(), πTemp003).ToObject()
							πTemp004 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp004...).ToObject()
							πTemp001 = πg.NewTuple4(πg.NewFloat(1.5).ToObject(), ßa.ToObject(), πTemp002, πTemp003).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
								continue
							}
							µx = πTemp002
							µy = πTemp003
							µz = πTemp005
							µw = πTemp006
							// line 611: self._not_tracked({})
							πF.SetLineno(611)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							πTemp001 = πTemp007.ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 612: self._not_tracked({x:(), y:x, z:1})
							πF.SetLineno(612)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple0().ToObject()
							if πE = πTemp007.SetItem(πF, µx, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, µy, µx); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, µz, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 613: self._not_tracked({1: "a", "b": 2})
							πF.SetLineno(613)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							if πE = πTemp007.SetItem(πF, πg.NewInt(1).ToObject(), ßa.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, ßb.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 614: self._not_tracked({1: 2, (None, True, False, ()): int})
							πF.SetLineno(614)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							if πE = πTemp007.SetItem(πF, πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple0().ToObject()
							πTemp001 = πg.NewTuple4(πTemp002, πTemp003, πTemp005, πTemp006).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 615: self._not_tracked({1: object()})
							πF.SetLineno(615)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, πg.NewInt(1).ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 619: self._tracked({1: []})
							πF.SetLineno(619)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							πTemp008 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp008...).ToObject()
							if πE = πTemp007.SetItem(πF, πg.NewInt(1).ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 620: self._tracked({1: ([],)})
							πF.SetLineno(620)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							πTemp008 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp008...).ToObject()
							πTemp001 = πg.NewTuple1(πTemp002).ToObject()
							if πE = πTemp007.SetItem(πF, πg.NewInt(1).ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 621: self._tracked({1: {}})
							πF.SetLineno(621)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							πTemp009 = πg.NewDict()
							πTemp001 = πTemp009.ToObject()
							if πE = πTemp007.SetItem(πF, πg.NewInt(1).ToObject(), πTemp001); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 622: self._tracked({1: set()})
							πF.SetLineno(622)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πg.NewDict()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πTemp007.SetItem(πF, πg.NewInt(1).ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp007.ToObject()
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_track_literals.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 606: @test_support.cpython_only
					πF.SetLineno(606)
					πTemp004 = πF.MakeArgs(1)
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßtest_track_literals); πE != nil {
						continue
					}
					πTemp004[0] = πTemp031
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßtest_support); πE != nil {
						continue
					}
					if πTemp032, πE = πg.GetAttr(πF, πTemp031, ßcpython_only, nil); πE != nil {
						continue
					}
					if πTemp031, πE = πTemp032.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_track_literals.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 625: def test_track_dynamic(self):
					πF.SetLineno(625)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("test_track_dynamic", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µMyObject *πg.Object = πg.UnboundLocal; _ = µMyObject
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var µy *πg.Object = πg.UnboundLocal; _ = µy
						var µz *πg.Object = πg.UnboundLocal; _ = µz
						var µw *πg.Object = πg.UnboundLocal; _ = µw
						var µo *πg.Object = πg.UnboundLocal; _ = µo
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µdd *πg.Object = πg.UnboundLocal; _ = µdd
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 πg.KWArgs
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 627: class MyObject(object):
							πF.SetLineno(627)
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
							_, πE = πg.NewCode("MyObject", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 628: pass
									πF.SetLineno(628)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MyObject").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µMyObject = πTemp005
							// line 629: x, y, z, w, o = 1.5, "a", (1, object()), [], MyObject()
							πF.SetLineno(629)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πg.NewInt(1).ToObject(), πTemp006).ToObject()
							πTemp003 = make([]*πg.Object, 0)
							πTemp005 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.CheckLocal(πF, µMyObject, "MyObject"); πE != nil {
								continue
							}
							if πTemp006, πE = µMyObject.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple5(πg.NewFloat(1.5).ToObject(), ßa.ToObject(), πTemp004, πTemp005, πTemp006).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
								continue
							}
							µx = πTemp004
							µy = πTemp005
							µz = πTemp006
							µw = πTemp007
							µo = πTemp008
							// line 631: d = dict()
							πF.SetLineno(631)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp004
							// line 632: self._not_tracked(d)
							πF.SetLineno(632)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 633: d[1] = "a"
							πF.SetLineno(633)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßa.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 634: self._not_tracked(d)
							πF.SetLineno(634)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 635: d[y] = 2
							πF.SetLineno(635)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp004 = µy
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 636: self._not_tracked(d)
							πF.SetLineno(636)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 637: d[z] = 3
							πF.SetLineno(637)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							πTemp004 = µz
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 638: self._not_tracked(d)
							πF.SetLineno(638)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 639: self._not_tracked(d.copy())
							πF.SetLineno(639)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 640: d[4] = w
							πF.SetLineno(640)
							if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µw); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(4).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 641: self._tracked(d)
							πF.SetLineno(641)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 642: self._tracked(d.copy())
							πF.SetLineno(642)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 643: d[4] = None
							πF.SetLineno(643)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(4).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp004); πE != nil {
								continue
							}
							// line 644: self._not_tracked(d)
							πF.SetLineno(644)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 645: self._not_tracked(d.copy())
							πF.SetLineno(645)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßcopy, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 649: d = dict()
							πF.SetLineno(649)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp004
							// line 650: dd = dict()
							πF.SetLineno(650)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdd = πTemp004
							// line 651: d[1] = dd
							πF.SetLineno(651)
							if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdd); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 652: self._not_tracked(dd)
							πF.SetLineno(652)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
								continue
							}
							πTemp003[0] = µdd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 653: self._tracked(d)
							πF.SetLineno(653)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 654: dd[1] = d
							πF.SetLineno(654)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µd); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.SetItem(πF, µdd, πTemp004, πTemp002); πE != nil {
								continue
							}
							// line 655: self._tracked(dd)
							πF.SetLineno(655)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
								continue
							}
							πTemp003[0] = µdd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 657: d = dict.fromkeys([x, y, z])
							πF.SetLineno(657)
							πTemp003 = πF.MakeArgs(1)
							πTemp009 = make([]*πg.Object, 3)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp009[0] = µx
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp009[1] = µy
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							πTemp009[2] = µz
							πTemp002 = πg.NewList(πTemp009...).ToObject()
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µd = πTemp002
							// line 658: self._not_tracked(d)
							πF.SetLineno(658)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 659: dd = dict()
							πF.SetLineno(659)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdd = πTemp004
							// line 660: dd.update(d)
							πF.SetLineno(660)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 661: self._not_tracked(dd)
							πF.SetLineno(661)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
								continue
							}
							πTemp003[0] = µdd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 662: d = dict.fromkeys([x, y, z, o])
							πF.SetLineno(662)
							πTemp003 = πF.MakeArgs(1)
							πTemp009 = make([]*πg.Object, 4)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp009[0] = µx
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp009[1] = µy
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							πTemp009[2] = µz
							if πE = πg.CheckLocal(πF, µo, "o"); πE != nil {
								continue
							}
							πTemp009[3] = µo
							πTemp002 = πg.NewList(πTemp009...).ToObject()
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßfromkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µd = πTemp002
							// line 663: self._tracked(d)
							πF.SetLineno(663)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 664: dd = dict()
							πF.SetLineno(664)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdd = πTemp004
							// line 665: dd.update(d)
							πF.SetLineno(665)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 666: self._tracked(dd)
							πF.SetLineno(666)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdd, "dd"); πE != nil {
								continue
							}
							πTemp003[0] = µdd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 668: d = dict(x=x, y=y, z=z)
							πF.SetLineno(668)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"x", µx},
								{"y", µy},
								{"z", µz},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							µd = πTemp004
							// line 669: self._not_tracked(d)
							πF.SetLineno(669)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 670: d = dict(x=x, y=y, z=z, w=w)
							πF.SetLineno(670)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"x", µx},
								{"y", µy},
								{"z", µz},
								{"w", µw},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							µd = πTemp004
							// line 671: self._tracked(d)
							πF.SetLineno(671)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 672: d = dict()
							πF.SetLineno(672)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp004
							// line 673: d.update(x=x, y=y, z=z)
							πF.SetLineno(673)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"x", µx},
								{"y", µy},
								{"z", µz},
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							// line 674: self._not_tracked(d)
							πF.SetLineno(674)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 675: d.update(w=w)
							πF.SetLineno(675)
							if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
								continue
							}
							πTemp010 = πg.KWArgs{
								{"w", µw},
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, πTemp010); πE != nil {
								continue
							}
							// line 676: self._tracked(d)
							πF.SetLineno(676)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 678: d = dict([(x, y), (z, 1)])
							πF.SetLineno(678)
							πTemp003 = πF.MakeArgs(1)
							πTemp009 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µx, µy).ToObject()
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µz, πg.NewInt(1).ToObject()).ToObject()
							πTemp009[1] = πTemp002
							πTemp002 = πg.NewList(πTemp009...).ToObject()
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µd = πTemp004
							// line 679: self._not_tracked(d)
							πF.SetLineno(679)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 680: d = dict([(x, y), (z, w)])
							πF.SetLineno(680)
							πTemp003 = πF.MakeArgs(1)
							πTemp009 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µx, µy).ToObject()
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µz, µw).ToObject()
							πTemp009[1] = πTemp002
							πTemp002 = πg.NewList(πTemp009...).ToObject()
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µd = πTemp004
							// line 681: self._tracked(d)
							πF.SetLineno(681)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 682: d = dict()
							πF.SetLineno(682)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp004
							// line 683: d.update([(x, y), (z, 1)])
							πF.SetLineno(683)
							πTemp003 = πF.MakeArgs(1)
							πTemp009 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µx, µy).ToObject()
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µz, πg.NewInt(1).ToObject()).ToObject()
							πTemp009[1] = πTemp002
							πTemp002 = πg.NewList(πTemp009...).ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 684: self._not_tracked(d)
							πF.SetLineno(684)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_not_tracked, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 685: d.update([(x, y), (z, w)])
							πF.SetLineno(685)
							πTemp003 = πF.MakeArgs(1)
							πTemp009 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µx, µy).ToObject()
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µz, "z"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µw, "w"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µz, µw).ToObject()
							πTemp009[1] = πTemp002
							πTemp002 = πg.NewList(πTemp009...).ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 686: self._tracked(d)
							πF.SetLineno(686)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_track_dynamic.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 624: @test_support.cpython_only
					πF.SetLineno(624)
					πTemp004 = πF.MakeArgs(1)
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßtest_track_dynamic); πE != nil {
						continue
					}
					πTemp004[0] = πTemp032
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßtest_support); πE != nil {
						continue
					}
					if πTemp033, πE = πg.GetAttr(πF, πTemp032, ßcpython_only, nil); πE != nil {
						continue
					}
					if πTemp032, πE = πTemp033.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_track_dynamic.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 689: def test_track_subtypes(self):
					πF.SetLineno(689)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("test_track_subtypes", "/usr/lib/python2.7/test/test_dict.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µMyDict *πg.Object = πg.UnboundLocal; _ = µMyDict
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
							// line 691: class MyDict(dict):
							πF.SetLineno(691)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
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
							_, πE = πg.NewCode("MyDict", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 692: pass
									πF.SetLineno(692)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MyDict").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µMyDict = πTemp005
							// line 693: self._tracked(MyDict())
							πF.SetLineno(693)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µMyDict, "MyDict"); πE != nil {
								continue
							}
							if πTemp002, πE = µMyDict.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tracked, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_track_subtypes.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 688: @test_support.cpython_only
					πF.SetLineno(688)
					πTemp004 = πF.MakeArgs(1)
					if πTemp033, πE = πg.ResolveClass(πF, πClass, nil, ßtest_track_subtypes); πE != nil {
						continue
					}
					πTemp004[0] = πTemp033
					if πTemp033, πE = πg.ResolveClass(πF, πClass, nil, ßtest_support); πE != nil {
						continue
					}
					if πTemp034, πE = πg.GetAttr(πF, πTemp033, ßcpython_only, nil); πE != nil {
						continue
					}
					if πTemp033, πE = πTemp034.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πClass.SetItem(πF, ßtest_track_subtypes.ToObject(), πTemp033); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("DictTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDictTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 705: from test import mapping_tests
			πF.SetLineno(705)
			if πTemp002, πE = πg.ImportModule(πF, "test.mapping_tests"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßmapping_tests.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 707: class GeneralMappingTests(mapping_tests.BasicTestMappingProtocol):
			πF.SetLineno(707)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßmapping_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßBasicTestMappingProtocol, nil); πE != nil {
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
			_, πE = πg.NewCode("GeneralMappingTests", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 708: type2test = dict
					πF.SetLineno(708)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdict); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("GeneralMappingTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßGeneralMappingTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 710: class Dict(dict):
			πF.SetLineno(710)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
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
			_, πE = πg.NewCode("Dict", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 711: pass
					πF.SetLineno(711)
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Dict").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDict.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 713: class SubclassMappingTests(mapping_tests.BasicTestMappingProtocol):
			πF.SetLineno(713)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßmapping_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßBasicTestMappingProtocol, nil); πE != nil {
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
			_, πE = πg.NewCode("SubclassMappingTests", "/usr/lib/python2.7/test/test_dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp003
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 714: type2test = Dict
					πF.SetLineno(714)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßDict); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("SubclassMappingTests").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSubclassMappingTests.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 716: def test_main():
			πF.SetLineno(716)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "/usr/lib/python2.7/test/test_dict.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 717: with test_support.check_py3k_warnings(
					πF.SetLineno(717)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πg.NewStr("dict(.has_key..| inequality comparisons) not supported in 3.x").ToObject(), πTemp003).ToObject()
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcheck_py3k_warnings, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__exit__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp002.Type().ToObject(), ß__enter__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp002}, nil); πE != nil {
						continue
					}
					πF.PushCheckpoint(1)
					// line 720: test_support.run_unittest(
					πF.SetLineno(720)
					πTemp001 = πF.MakeArgs(3)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßDictTest); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßGeneralMappingTests); πE != nil {
						continue
					}
					πTemp001[1] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSubclassMappingTests); πE != nil {
						continue
					}
					πTemp001[2] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßrun_unittest, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
				Label1:
					πTemp008, πTemp009 = nil, nil
					if πE != nil {
						πTemp008, πTemp009 = πF.ExcInfo()
					}
					if πTemp008 != nil {
						πTemp010 = πTemp008.Type()
						if πTemp005, πE = πTemp003.Call(πF, πg.Args{πTemp002, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
							continue
						}
					} else {
						if πTemp005, πE = πTemp003.Call(πF, πg.Args{πTemp002, πg.None, πg.None, πg.None}, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßtest_main.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp004, πE = πg.Eq(πF, πTemp005, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label1
			}
			goto Label2
			// line 726: if __name__ == "__main__":
			πF.SetLineno(726)
		Label1:
			// line 727: test_main()
			πF.SetLineno(727)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßtest_main); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("test.test_dict", Code)
}

