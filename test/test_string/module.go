package test_string
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/UserList"
	// _ "github.com/pygolin/stdlib/unittest"
	// _ "github.com/pygolin/stdlib/test/test_support"
	// _ "github.com/pygolin/stdlib/string"
	// _ "github.com/pygolin/stdlib/test/string_tests"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß10 := πg.InternStr("10")
		ß10100 := πg.InternStr("10100")
		ß1010020 := πg.InternStr("1010020")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBadSeq1 := πg.InternStr("BadSeq1")
		ßBadSeq2 := πg.InternStr("BadSeq2")
		ßBag := πg.InternStr("Bag")
		ßBytesAliasTest := πg.InternStr("BytesAliasTest")
		ßCleveland := πg.InternStr("Cleveland")
		ßCommonTest := πg.InternStr("CommonTest")
		ßFormatter := πg.InternStr("Formatter")
		ßKeyError := πg.InternStr("KeyError")
		ßMapping := πg.InternStr("Mapping")
		ßMixinStrStringUserStringTest := πg.InternStr("MixinStrStringUserStringTest")
		ßModuleTest := πg.InternStr("ModuleTest")
		ßNone := πg.InternStr("None")
		ßSequence := πg.InternStr("Sequence")
		ßStringTest := πg.InternStr("StringTest")
		ßTemplate := πg.InternStr("Template")
		ßTestCase := πg.InternStr("TestCase")
		ßTestTemplate := πg.InternStr("TestTemplate")
		ßTypeError := πg.InternStr("TypeError")
		ßUserList := πg.InternStr("UserList")
		ßValueError := πg.InternStr("ValueError")
		ß_UserList := πg.InternStr("_UserList")
		ß__doc__ := πg.InternStr("__doc__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ßa := πg.InternStr("a")
		ßabc := πg.InternStr("abc")
		ßabcd := πg.InternStr("abcd")
		ßargs := πg.InternStr("args")
		ßassertAlmostEqual := πg.InternStr("assertAlmostEqual")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertNotEqual := πg.InternStr("assertNotEqual")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßatof := πg.InternStr("atof")
		ßatoi := πg.InternStr("atoi")
		ßatol := πg.InternStr("atol")
		ßb := πg.InternStr("b")
		ßbag := πg.InternStr("bag")
		ßbar := πg.InternStr("bar")
		ßbaz := πg.InternStr("baz")
		ßbozo := πg.InternStr("bozo")
		ßbud := πg.InternStr("bud")
		ßbytes := πg.InternStr("bytes")
		ßc := πg.InternStr("c")
		ßcapwords := πg.InternStr("capwords")
		ßcheck_unused_args := πg.InternStr("check_unused_args")
		ßcheckcall := πg.InternStr("checkcall")
		ßcheckequal := πg.InternStr("checkequal")
		ßcheckraises := πg.InternStr("checkraises")
		ßconvert_field := πg.InternStr("convert_field")
		ßd := πg.InternStr("d")
		ßdelimiter := πg.InternStr("delimiter")
		ßdict := πg.InternStr("dict")
		ßdigits := πg.InternStr("digits")
		ßdinner := πg.InternStr("dinner")
		ßexception := πg.InternStr("exception")
		ßexpectedFailure := πg.InternStr("expectedFailure")
		ßfail := πg.InternStr("fail")
		ßfixtype := πg.InternStr("fixtype")
		ßfoo := πg.InternStr("foo")
		ßfoobar := πg.InternStr("foobar")
		ßformat := πg.InternStr("format")
		ßformat_field := πg.InternStr("format_field")
		ßfred := πg.InternStr("fred")
		ßget_value := πg.InternStr("get_value")
		ßgetattr := πg.InternStr("getattr")
		ßgreeting := πg.InternStr("greeting")
		ßham := πg.InternStr("ham")
		ßhave_unicode := πg.InternStr("have_unicode")
		ßhello := πg.InternStr("hello")
		ßhexdigits := πg.InternStr("hexdigits")
		ßidpattern := πg.InternStr("idpattern")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßletters := πg.InternStr("letters")
		ßlocation := πg.InternStr("location")
		ßlowercase := πg.InternStr("lowercase")
		ßmaketrans := πg.InternStr("maketrans")
		ßnamespace := πg.InternStr("namespace")
		ßnone := πg.InternStr("none")
		ßobject := πg.InternStr("object")
		ßoctdigits := πg.InternStr("octdigits")
		ßone := πg.InternStr("one")
		ßparse := πg.InternStr("parse")
		ßpartition := πg.InternStr("partition")
		ßpattern := πg.InternStr("pattern")
		ßprintable := πg.InternStr("printable")
		ßpunctuation := πg.InternStr("punctuation")
		ßrange := πg.InternStr("range")
		ßremove := πg.InternStr("remove")
		ßresult := πg.InternStr("result")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßsafe_substitute := πg.InternStr("safe_substitute")
		ßset := πg.InternStr("set")
		ßspam := πg.InternStr("spam")
		ßsplit := πg.InternStr("split")
		ßstr := πg.InternStr("str")
		ßstring := πg.InternStr("string")
		ßstring_tests := πg.InternStr("string_tests")
		ßsubstitute := πg.InternStr("substitute")
		ßsuper := πg.InternStr("super")
		ßtest := πg.InternStr("test")
		ßtest_SafeTemplate := πg.InternStr("test_SafeTemplate")
		ßtest_atof := πg.InternStr("test_atof")
		ßtest_atoi := πg.InternStr("test_atoi")
		ßtest_atol := πg.InternStr("test_atol")
		ßtest_attrs := πg.InternStr("test_attrs")
		ßtest_braced_override := πg.InternStr("test_braced_override")
		ßtest_braced_override_safe := πg.InternStr("test_braced_override_safe")
		ßtest_builtin := πg.InternStr("test_builtin")
		ßtest_capwords := πg.InternStr("test_capwords")
		ßtest_delimiter_override := πg.InternStr("test_delimiter_override")
		ßtest_escapes := πg.InternStr("test_escapes")
		ßtest_format_keyword_arguments := πg.InternStr("test_format_keyword_arguments")
		ßtest_formatter := πg.InternStr("test_formatter")
		ßtest_idpattern_override := πg.InternStr("test_idpattern_override")
		ßtest_invalid_placeholders := πg.InternStr("test_invalid_placeholders")
		ßtest_join := πg.InternStr("test_join")
		ßtest_keyword_arguments := πg.InternStr("test_keyword_arguments")
		ßtest_keyword_arguments_safe := πg.InternStr("test_keyword_arguments_safe")
		ßtest_main := πg.InternStr("test_main")
		ßtest_maketrans := πg.InternStr("test_maketrans")
		ßtest_pattern_override := πg.InternStr("test_pattern_override")
		ßtest_percents := πg.InternStr("test_percents")
		ßtest_regular_templates := πg.InternStr("test_regular_templates")
		ßtest_regular_templates_with_braces := πg.InternStr("test_regular_templates_with_braces")
		ßtest_stringification := πg.InternStr("test_stringification")
		ßtest_support := πg.InternStr("test_support")
		ßtest_syntax := πg.InternStr("test_syntax")
		ßtest_tupleargs := πg.InternStr("test_tupleargs")
		ßtest_unicode_values := πg.InternStr("test_unicode_values")
		ßtim := πg.InternStr("tim")
		ßtwo := πg.InternStr("two")
		ßtype := πg.InternStr("type")
		ßtype2test := πg.InternStr("type2test")
		ßunicode := πg.InternStr("unicode")
		ßunittest := πg.InternStr("unittest")
		ßupdate := πg.InternStr("update")
		ßuppercase := πg.InternStr("uppercase")
		ßwhat := πg.InternStr("what")
		ßwhitespace := πg.InternStr("whitespace")
		ßwho := πg.InternStr("who")
		ßx := πg.InternStr("x")
		ßxyz := πg.InternStr("xyz")
		ßxyzq := πg.InternStr("xyzq")
		ßyou := πg.InternStr("you")
		ßz := πg.InternStr("z")
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
			// line 1: import unittest
			πF.SetLineno(1)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 2: import string
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "string"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstring.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: Template = string.Template
			πF.SetLineno(4)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßTemplate, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTemplate.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 5: from test import test_support, string_tests
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "test.string_tests"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßstring_tests.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import UserList as _UserList
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "UserList"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_UserList.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: UserList = _UserList.UserList
			πF.SetLineno(8)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_UserList); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßUserList, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUserList.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 10: class StringTest(
			πF.SetLineno(10)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßstring_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßCommonTest, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp006
			if πTemp005, πE = πg.ResolveGlobal(πF, ßstring_tests); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßMixinStrStringUserStringTest, nil); πE != nil {
				continue
			}
			πTemp002[1] = πTemp006
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("StringTest", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 []*πg.Object
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
					// line 15: type2test = str
					πF.SetLineno(15)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßstr); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 17: def checkequal(self, result, object, methodname, *args):
					πF.SetLineno(17)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					πTemp002[2] = πg.Param{Name: "object", Def: nil}
					πTemp002[3] = πg.Param{Name: "methodname", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("checkequal", "/usr/lib/python2.7/test/test_string.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var µobject *πg.Object = πArgs[2]; _ = µobject
						var µmethodname *πg.Object = πArgs[3]; _ = µmethodname
						var µargs *πg.Object = πArgs[4]; _ = µargs
						var µrealresult *πg.Object = πg.UnboundLocal; _ = µrealresult
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
							// line 18: realresult = getattr(string, methodname)(object, *args)
							πF.SetLineno(18)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							πTemp001[0] = µobject
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µmethodname, "methodname"); πE != nil {
								continue
							}
							πTemp002[1] = µmethodname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Invoke(πF, πTemp004, πTemp001, µargs, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrealresult = πTemp003
							// line 19: self.assertEqual(
							πF.SetLineno(19)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µrealresult, "realresult"); πE != nil {
								continue
							}
							πTemp001[1] = µrealresult
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
					if πE = πClass.SetItem(πF, ßcheckequal.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 24: def checkraises(self, exc, obj, methodname, *args):
					πF.SetLineno(24)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "exc", Def: nil}
					πTemp002[2] = πg.Param{Name: "obj", Def: nil}
					πTemp002[3] = πg.Param{Name: "methodname", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("checkraises", "/usr/lib/python2.7/test/test_string.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexc *πg.Object = πArgs[1]; _ = µexc
						var µobj *πg.Object = πArgs[2]; _ = µobj
						var µmethodname *πg.Object = πArgs[3]; _ = µmethodname
						var µargs *πg.Object = πArgs[4]; _ = µargs
						var µcm *πg.Object = πg.UnboundLocal; _ = µcm
						var πTemp001 []*πg.Object
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
						var πTemp008 bool
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 25: with self.assertRaises(exc) as cm:
							πF.SetLineno(25)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexc, "exc"); πE != nil {
								continue
							}
							πTemp001[0] = µexc
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
							µcm = πTemp004
							// line 26: getattr(string, methodname)(obj, *args)
							πF.SetLineno(26)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(2)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πE = πg.CheckLocal(πF, µmethodname, "methodname"); πE != nil {
								continue
							}
							πTemp005[1] = µmethodname
							if πTemp006, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp006, πE = πg.Invoke(πF, πTemp007, πTemp001, µargs, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
						Label1:
							πTemp009, πTemp010 = nil, nil
							if πE != nil {
								πTemp009, πTemp010 = πF.ExcInfo()
							}
							if πTemp009 != nil {
								πTemp011 = πTemp009.Type()
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp011.ToObject(), πTemp009.ToObject(), πTemp010.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp006, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
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
							// line 27: self.assertNotEqual(cm.exception.args[0], '')
							πF.SetLineno(27)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcm, "cm"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µcm, ßexception, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßargs, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp013, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßassertNotEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcheckraises.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 29: def checkcall(self, object, methodname, *args):
					πF.SetLineno(29)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "object", Def: nil}
					πTemp002[2] = πg.Param{Name: "methodname", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("checkcall", "/usr/lib/python2.7/test/test_string.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobject *πg.Object = πArgs[1]; _ = µobject
						var µmethodname *πg.Object = πArgs[2]; _ = µmethodname
						var µargs *πg.Object = πArgs[3]; _ = µargs
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
							// line 30: getattr(string, methodname)(object, *args)
							πF.SetLineno(30)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							πTemp001[0] = µobject
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(2)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µmethodname, "methodname"); πE != nil {
								continue
							}
							πTemp002[1] = µmethodname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Invoke(πF, πTemp004, πTemp001, µargs, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcheckcall.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 33: def test_join(self):
					πF.SetLineno(33)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_join", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µe *πg.Object = πg.UnboundLocal; _ = µe
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
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 []πg.Param
						_ = πTemp012
						var πTemp013 *πg.BaseException
						_ = πTemp013
						var πTemp014 *πg.Traceback
						_ = πTemp014
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 36: self.checkequal('a b c d', ['a', 'b', 'c', 'd'], 'join', ' ')
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("a b c d").ToObject()
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = ßa.ToObject()
							πTemp002[1] = ßb.ToObject()
							πTemp002[2] = ßc.ToObject()
							πTemp002[3] = ßd.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 37: self.checkequal('abcd', ('a', 'b', 'c', 'd'), 'join', '')
							πF.SetLineno(37)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßabcd.ToObject()
							πTemp003 = πg.NewTuple4(ßa.ToObject(), ßb.ToObject(), ßc.ToObject(), ßd.ToObject()).ToObject()
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 38: self.checkequal('w x y z', string_tests.Sequence(), 'join', ' ')
							πF.SetLineno(38)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("w x y z").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring_tests); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßSequence, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 39: self.checkequal('abc', ('abc',), 'join', 'a')
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßabc.ToObject()
							πTemp003 = πg.NewTuple1(ßabc.ToObject()).ToObject()
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 40: self.checkequal('z', UserList(['z']), 'join', 'a')
							πF.SetLineno(40)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßz.ToObject()
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 1)
							πTemp005[0] = ßz.ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßUserList); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßhave_unicode, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 41: if test_support.have_unicode:
							πF.SetLineno(41)
						Label1:
							// line 42: self.checkequal(unicode('a.b.c'), ['a', 'b', 'c'], 'join', unicode('.'))
							πF.SetLineno(42)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("a.b.c").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = ßa.ToObject()
							πTemp002[1] = ßb.ToObject()
							πTemp002[2] = ßc.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(".").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[3] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 43: self.checkequal(unicode('a.b.c'), [unicode('a'), 'b', 'c'], 'join', '.')
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("a.b.c").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = make([]*πg.Object, 3)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßa.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp004
							πTemp002[1] = ßb.ToObject()
							πTemp002[2] = ßc.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 44: self.checkequal(unicode('a.b.c'), ['a', unicode('b'), 'c'], 'join', '.')
							πF.SetLineno(44)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("a.b.c").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = ßa.ToObject()
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßb.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[1] = πTemp004
							πTemp002[2] = ßc.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 45: self.checkequal(unicode('a.b.c'), ['a', 'b', unicode('c')], 'join', '.')
							πF.SetLineno(45)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("a.b.c").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = ßa.ToObject()
							πTemp002[1] = ßb.ToObject()
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßc.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[2] = πTemp004
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 46: self.checkraises(TypeError, ['a', unicode('b'), 3], 'join', '.')
							πF.SetLineno(46)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = ßa.ToObject()
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = ßb.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[1] = πTemp004
							πTemp002[2] = πg.NewInt(3).ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							πTemp001 = make([]*πg.Object, 3)
							πTemp001[0] = πg.NewInt(5).ToObject()
							πTemp001[1] = πg.NewInt(25).ToObject()
							πTemp001[2] = πg.NewInt(125).ToObject()
							πTemp004 = πg.NewList(πTemp001...).ToObject()
							if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							πTemp006 = false
						Label3:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label5
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
							πF.PushCheckpoint(3)            
							// line 48: self.checkequal(
							πF.SetLineno(48)
							πTemp001 = πF.MakeArgs(4)
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp008, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Mul(πF, ßa.ToObject(), µi); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Add(πF, πTemp011, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Mul(πF, πTemp010, µi); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							πTemp002 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mul(πF, ßa.ToObject(), µi); πE != nil {
								continue
							}
							πTemp002[0] = πTemp008
							πTemp008 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp008, µi); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr("-").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 51: self.checkequal(
							πF.SetLineno(51)
							πTemp001 = πF.MakeArgs(4)
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp008, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Mul(πF, ßa.ToObject(), µi); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Add(πF, πTemp011, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Mul(πF, πTemp010, µi); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Mul(πF, ßa.ToObject(), µi); πE != nil {
								continue
							}
							πTemp008 = πg.NewTuple1(πTemp009).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp008, µi); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr("-").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 55: self.checkraises(TypeError, string_tests.BadSeq1(), 'join', ' ')
							πF.SetLineno(55)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring_tests); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßBadSeq1, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 56: self.checkequal('a b c', string_tests.BadSeq2(), 'join', ' ')
							πF.SetLineno(56)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("a b c").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring_tests); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßBadSeq2, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 57: try:
							πF.SetLineno(57)
							πF.PushCheckpoint(7)
							// line 58: def f():
							πF.SetLineno(58)
							πTemp012 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/test_string.py", πTemp012, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										default: panic("unexpected function state")
										}
										// line 59: yield 4 + ""
										πF.SetLineno(59)
										if πTemp001, πE = πg.Add(πF, πg.NewInt(4).ToObject(), ß.ToObject()); πE != nil {
											continue
										}
										πF.PushCheckpoint(1)
										return πTemp001, nil
									Label1:
										πTemp002 = πSent
									}
									return nil, πE
								}).ToObject(), nil
							}), πF.Globals()).ToObject()
							µf = πTemp003
							// line 60: self.fixtype(' ').join(f())
							πF.SetLineno(60)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp004, πE = µf.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(" ").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.GetAttr(πF, πTemp008, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							// line 65: self.fail('exception not raised')
							πF.SetLineno(65)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("exception not raised").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label7:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp013, πTemp014 = πF.ExcInfo()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							πE = πF.Raise(πTemp013.ToObject(), nil, πTemp014.ToObject())
							continue
							// line 61: except TypeError, e:
							πF.SetLineno(61)
						Label8:
							µe = πTemp013.ToObject()
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp001[0] = µe
							if πTemp008, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.Contains(πF, πTemp009, πg.NewStr("+").ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							goto Label10
							// line 62: if '+' not in str(e):
							πF.SetLineno(62)
						Label9:
							// line 63: self.fail('join() ate exception message')
							πF.SetLineno(63)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("join() ate exception message").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label10
						Label10:
							πF.RestoreExc(nil, nil)
							goto Label6
						Label6:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_join.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 32: @unittest.expectedFailure
					πF.SetLineno(32)
					πTemp006 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_join); πE != nil {
						continue
					}
					πTemp006[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_join.ToObject(), πTemp007); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("StringTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 68: class ModuleTest(unittest.TestCase):
			πF.SetLineno(68)
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
			_, πE = πg.NewCode("ModuleTest", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 []*πg.Object
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
					// line 70: def test_attrs(self):
					πF.SetLineno(70)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_attrs", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 71: string.whitespace
							πF.SetLineno(71)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwhitespace, nil); πE != nil {
								continue
							}
							// line 72: string.lowercase
							πF.SetLineno(72)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlowercase, nil); πE != nil {
								continue
							}
							// line 73: string.uppercase
							πF.SetLineno(73)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßuppercase, nil); πE != nil {
								continue
							}
							// line 74: string.letters
							πF.SetLineno(74)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßletters, nil); πE != nil {
								continue
							}
							// line 75: string.digits
							πF.SetLineno(75)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdigits, nil); πE != nil {
								continue
							}
							// line 76: string.hexdigits
							πF.SetLineno(76)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßhexdigits, nil); πE != nil {
								continue
							}
							// line 77: string.octdigits
							πF.SetLineno(77)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßoctdigits, nil); πE != nil {
								continue
							}
							// line 78: string.punctuation
							πF.SetLineno(78)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpunctuation, nil); πE != nil {
								continue
							}
							// line 79: string.printable
							πF.SetLineno(79)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßprintable, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_attrs.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 81: def test_atoi(self):
					πF.SetLineno(81)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_atoi", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 82: self.assertEqual(string.atoi(" 1 "), 1)
							πF.SetLineno(82)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(" 1 ").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßatoi, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
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
							// line 83: self.assertRaises(ValueError, string.atoi, " 1x")
							πF.SetLineno(83)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßatoi, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = πg.NewStr(" 1x").ToObject()
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
							// line 84: self.assertRaises(ValueError, string.atoi, " x1 ")
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßatoi, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = πg.NewStr(" x1 ").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_atoi.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 86: def test_atol(self):
					πF.SetLineno(86)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_atol", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 87: self.assertEqual(string.atol("  1  "), 1L)
							πF.SetLineno(87)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("  1  ").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßatol, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewLongFromBytes([]byte{0x1,}).ToObject()
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
							// line 88: self.assertRaises(ValueError, string.atol, "  1x ")
							πF.SetLineno(88)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßatol, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = πg.NewStr("  1x ").ToObject()
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
							// line 89: self.assertRaises(ValueError, string.atol, "  x1 ")
							πF.SetLineno(89)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßatol, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = πg.NewStr("  x1 ").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_atol.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 92: def test_atof(self):
					πF.SetLineno(92)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_atof", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 93: self.assertAlmostEqual(string.atof("  1  "), 1.0)
							πF.SetLineno(93)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("  1  ").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßatof, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewFloat(1.0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßassertAlmostEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 94: self.assertRaises(ValueError, string.atof, "  1x ")
							πF.SetLineno(94)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßatof, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = πg.NewStr("  1x ").ToObject()
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
							// line 95: self.assertRaises(ValueError, string.atof, "  x1 ")
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßatof, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = πg.NewStr("  x1 ").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_atof.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 91: @unittest.expectedFailure
					πF.SetLineno(91)
					πTemp006 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßtest_atof); πE != nil {
						continue
					}
					πTemp006[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_atof.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 97: def test_maketrans(self):
					πF.SetLineno(97)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_maketrans", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtranstable *πg.Object = πg.UnboundLocal; _ = µtranstable
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
							// line 98: transtable = '\000\001\002\003\004\005\006\007\010\011\012\013\014\015\016\017\020\021\022\023\024\025\026\027\030\031\032\033\034\035\036\037 !"#$%&\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`xyzdefghijklmnopqrstuvwxyz{|}~\177\200\201\202\203\204\205\206\207\210\211\212\213\214\215\216\217\220\221\222\223\224\225\226\227\230\231\232\233\234\235\236\237\240\241\242\243\244\245\246\247\250\251\252\253\254\255\256\257\260\261\262\263\264\265\266\267\270\271\272\273\274\275\276\277\300\301\302\303\304\305\306\307\310\311\312\313\314\315\316\317\320\321\322\323\324\325\326\327\330\331\332\333\334\335\336\337\340\341\342\343\344\345\346\347\350\351\352\353\354\355\356\357\360\361\362\363\364\365\366\367\370\371\372\373\374\375\376\377'
							πF.SetLineno(98)
							µtranstable = πg.NewStr("\x00\x01\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`xyzdefghijklmnopqrstuvwxyz{|}~\x7f\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff").ToObject()
							// line 100: self.assertEqual(string.maketrans('abc', 'xyz'), transtable)
							πF.SetLineno(100)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßabc.ToObject()
							πTemp002[1] = ßxyz.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmaketrans, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µtranstable, "transtable"); πE != nil {
								continue
							}
							πTemp001[1] = µtranstable
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
							// line 101: self.assertRaises(ValueError, string.maketrans, 'abc', 'xyzq')
							πF.SetLineno(101)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmaketrans, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							πTemp001[2] = ßabc.ToObject()
							πTemp001[3] = ßxyzq.ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_maketrans.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 104: def test_capwords(self):
					πF.SetLineno(104)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_capwords", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 105: self.assertEqual(string.capwords('abc def ghi'), 'Abc Def Ghi')
							πF.SetLineno(105)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("abc def ghi").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcapwords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Abc Def Ghi").ToObject()
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
							// line 106: self.assertEqual(string.capwords('abc\tdef\nghi'), 'Abc Def Ghi')
							πF.SetLineno(106)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("abc\tdef\nghi").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcapwords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Abc Def Ghi").ToObject()
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
							// line 107: self.assertEqual(string.capwords('abc\t   def  \nghi'), 'Abc Def Ghi')
							πF.SetLineno(107)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("abc\t   def  \nghi").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcapwords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Abc Def Ghi").ToObject()
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
							// line 108: self.assertEqual(string.capwords('ABC DEF GHI'), 'Abc Def Ghi')
							πF.SetLineno(108)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("ABC DEF GHI").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcapwords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Abc Def Ghi").ToObject()
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
							// line 109: self.assertEqual(string.capwords('ABC-DEF-GHI', '-'), 'Abc-Def-Ghi')
							πF.SetLineno(109)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewStr("ABC-DEF-GHI").ToObject()
							πTemp002[1] = πg.NewStr("-").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcapwords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Abc-Def-Ghi").ToObject()
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
							// line 110: self.assertEqual(string.capwords('ABC-def DEF-ghi GHI'), 'Abc-def Def-ghi Ghi')
							πF.SetLineno(110)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("ABC-def DEF-ghi GHI").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcapwords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Abc-def Def-ghi Ghi").ToObject()
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
							// line 111: self.assertEqual(string.capwords('   aBc  DeF   '), 'Abc Def')
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("   aBc  DeF   ").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcapwords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Abc Def").ToObject()
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
							// line 112: self.assertEqual(string.capwords('\taBc\tDeF\t'), 'Abc Def')
							πF.SetLineno(112)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("\taBc\tDeF\t").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcapwords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("Abc Def").ToObject()
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
							// line 113: self.assertEqual(string.capwords('\taBc\tDeF\t', '\t'), '\tAbc\tDef\t')
							πF.SetLineno(113)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = πg.NewStr("\taBc\tDeF\t").ToObject()
							πTemp002[1] = πg.NewStr("\t").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcapwords, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("\tAbc\tDef\t").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_capwords.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 103: @unittest.expectedFailure
					πF.SetLineno(103)
					πTemp006 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßtest_capwords); πE != nil {
						continue
					}
					πTemp006[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp010.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_capwords.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 116: def test_formatter(self):
					πF.SetLineno(116)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_formatter", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfmt *πg.Object = πg.UnboundLocal; _ = µfmt
						var µNamespaceFormatter *πg.Object = πg.UnboundLocal; _ = µNamespaceFormatter
						var µCallFormatter *πg.Object = πg.UnboundLocal; _ = µCallFormatter
						var µXFormatter *πg.Object = πg.UnboundLocal; _ = µXFormatter
						var µBarFormatter *πg.Object = πg.UnboundLocal; _ = µBarFormatter
						var µCheckAllUsedFormatter *πg.Object = πg.UnboundLocal; _ = µCheckAllUsedFormatter
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 πg.KWArgs
						_ = πTemp005
						var πTemp006 *πg.Dict
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []πg.Param
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 117: fmt = string.Formatter()
							πF.SetLineno(117)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßFormatter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfmt = πTemp001
							// line 118: self.assertEqual(fmt.format("foo"), "foo")
							πF.SetLineno(118)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßfoo.ToObject()
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
							// line 120: self.assertEqual(fmt.format("foo{0}", "bar"), "foobar")
							πF.SetLineno(120)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("foo{0}").ToObject()
							πTemp004[1] = ßbar.ToObject()
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßfoobar.ToObject()
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
							// line 121: self.assertEqual(fmt.format("foo{1}{0}-{1}", "bar", 6), "foo6bar-6")
							πF.SetLineno(121)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(3)
							πTemp004[0] = πg.NewStr("foo{1}{0}-{1}").ToObject()
							πTemp004[1] = ßbar.ToObject()
							πTemp004[2] = πg.NewInt(6).ToObject()
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("foo6bar-6").ToObject()
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
							// line 122: self.assertEqual(fmt.format("-{arg!r}-", arg='test'), "-'test'-")
							πF.SetLineno(122)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("-{arg!r}-").ToObject()
							πTemp005 = πg.KWArgs{
								{"arg", ßtest.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("-'test'-").ToObject()
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
							// line 125: class NamespaceFormatter(string.Formatter):
							πF.SetLineno(125)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßFormatter, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp006 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("NamespaceFormatter", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp006
								_ = πClass
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 []πg.Param
								_ = πTemp002
								var πTemp003 *πg.Dict
								_ = πTemp003
								var πTemp004 *πg.Object
								_ = πTemp004
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 126: def __init__(self, namespace={}):
									πF.SetLineno(126)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewDict()
									πTemp004 = πTemp003.ToObject()
									πTemp002[1] = πg.Param{Name: "namespace", Def: πTemp004}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µnamespace *πg.Object = πArgs[1]; _ = µnamespace
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
											// line 127: string.Formatter.__init__(self)
											πF.SetLineno(127)
											πTemp001 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											πTemp001[0] = µself
											if πTemp002, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßFormatter, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp001)
											// line 128: self.namespace = namespace
											πF.SetLineno(128)
											if πE = πg.CheckLocal(πF, µnamespace, "namespace"); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µnamespace); πE != nil {
												continue
											}
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πE = πg.SetAttr(πF, µself, ßnamespace, πTemp002); πE != nil {
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
									// line 130: def get_value(self, key, args, kwds):
									πF.SetLineno(130)
									πTemp002 = make([]πg.Param, 4)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "key", Def: nil}
									πTemp002[2] = πg.Param{Name: "args", Def: nil}
									πTemp002[3] = πg.Param{Name: "kwds", Def: nil}
									πTemp004 = πg.NewFunction(πg.NewCode("get_value", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µkey *πg.Object = πArgs[1]; _ = µkey
										var µargs *πg.Object = πArgs[2]; _ = µargs
										var µkwds *πg.Object = πArgs[3]; _ = µkwds
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
										var πTemp007 *πg.Object
										_ = πTemp007
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											case 5: goto Label5
											default: panic("unexpected function state")
											}
											πTemp001 = πF.MakeArgs(2)
											if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
												continue
											}
											πTemp001[0] = µkey
											if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
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
											// line 131: if isinstance(key, str):
											πF.SetLineno(131)
										Label1:
											// line 132: try:
											πF.SetLineno(132)
											πF.PushCheckpoint(5)
											// line 134: return kwds[key]
											πF.SetLineno(134)
											if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
												continue
											}
											πTemp002 = µkey
											if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetItem(πF, µkwds, πTemp002); πE != nil {
												continue
											}
											πR = πTemp003
											continue
											πF.PopCheckpoint()
											goto Label4
										Label5:
											if πE == nil {
											  continue
											}
											πE = nil
											πTemp005, πTemp006 = πF.ExcInfo()
											if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
												continue
											}
											if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
												continue
											}
											if πTemp004 {
												goto Label6
											}
											πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
											continue
											// line 135: except KeyError:
											πF.SetLineno(135)
										Label6:
											// line 136: return self.namespace[key]
											πF.SetLineno(136)
											if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
												continue
											}
											πTemp002 = µkey
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp007, πE = πg.GetAttr(πF, µself, ßnamespace, nil); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
												continue
											}
											πR = πTemp003
											continue
											πF.RestoreExc(nil, nil)
											goto Label4
										Label4:
											goto Label3
										Label2:
											// line 138: string.Formatter.get_value(key, args, kwds)
											πF.SetLineno(138)
											πTemp001 = πF.MakeArgs(3)
											if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
												continue
											}
											πTemp001[0] = µkey
											if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
												continue
											}
											πTemp001[1] = µargs
											if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
												continue
											}
											πTemp001[2] = µkwds
											if πTemp002, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßFormatter, nil); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßget_value, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ßget_value.ToObject(), πTemp004); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp002 == nil {
								πTemp002 = πg.TypeType.ToObject()
							}
							if πTemp007, πE = πTemp002.Call(πF, []*πg.Object{πg.NewStr("NamespaceFormatter").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
								continue
							}
							µNamespaceFormatter = πTemp007
							// line 140: fmt = NamespaceFormatter({'greeting':'hello'})
							πF.SetLineno(140)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πg.NewDict()
							if πE = πTemp006.SetItem(πF, ßgreeting.ToObject(), ßhello.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp006.ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µNamespaceFormatter, "NamespaceFormatter"); πE != nil {
								continue
							}
							if πTemp001, πE = µNamespaceFormatter.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µfmt = πTemp001
							// line 141: self.assertEqual(fmt.format("{greeting}, world!"), 'hello, world!')
							πF.SetLineno(141)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("{greeting}, world!").ToObject()
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("hello, world!").ToObject()
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
							// line 145: class CallFormatter(string.Formatter):
							πF.SetLineno(145)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßFormatter, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp006 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("CallFormatter", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 146: def format_field(self, value, format_spec):
									πF.SetLineno(146)
									πTemp002 = make([]πg.Param, 3)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "value", Def: nil}
									πTemp002[2] = πg.Param{Name: "format_spec", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("format_field", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µvalue *πg.Object = πArgs[1]; _ = µvalue
										var µformat_spec *πg.Object = πArgs[2]; _ = µformat_spec
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
											// line 147: return format(value(), format_spec)
											πF.SetLineno(147)
											πTemp001 = πF.MakeArgs(2)
											if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
												continue
											}
											if πTemp002, πE = µvalue.Call(πF, nil, nil); πE != nil {
												continue
											}
											πTemp001[0] = πTemp002
											if πE = πg.CheckLocal(πF, µformat_spec, "format_spec"); πE != nil {
												continue
											}
											πTemp001[1] = µformat_spec
											if πTemp002, πE = πg.ResolveGlobal(πF, ßformat); πE != nil {
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
									if πE = πClass.SetItem(πF, ßformat_field.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp002 == nil {
								πTemp002 = πg.TypeType.ToObject()
							}
							if πTemp007, πE = πTemp002.Call(πF, []*πg.Object{πg.NewStr("CallFormatter").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
								continue
							}
							µCallFormatter = πTemp007
							// line 149: fmt = CallFormatter()
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µCallFormatter, "CallFormatter"); πE != nil {
								continue
							}
							if πTemp001, πE = µCallFormatter.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfmt = πTemp001
							// line 150: self.assertEqual(fmt.format('*{0}*', lambda : 'result'), '*result*')
							πF.SetLineno(150)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("*{0}*").ToObject()
							πTemp009 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/test/test_string.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 150: self.assertEqual(fmt.format('*{0}*', lambda : 'result'), '*result*')
									πF.SetLineno(150)
									πR = ßresult.ToObject()
									continue
								}
								if πE != nil {
									πR = nil
								} else if πR == nil {
									πR = πg.None
								}
								return πR, πE
							}), πF.Globals()).ToObject()
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("*result*").ToObject()
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
							// line 154: class XFormatter(string.Formatter):
							πF.SetLineno(154)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßFormatter, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp006 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("XFormatter", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 155: def convert_field(self, value, conversion):
									πF.SetLineno(155)
									πTemp002 = make([]πg.Param, 3)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "value", Def: nil}
									πTemp002[2] = πg.Param{Name: "conversion", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("convert_field", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µvalue *πg.Object = πArgs[1]; _ = µvalue
										var µconversion *πg.Object = πArgs[2]; _ = µconversion
										var πTemp001 *πg.Object
										_ = πTemp001
										var πTemp002 bool
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
											if πE = πg.CheckLocal(πF, µconversion, "conversion"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Eq(πF, µconversion, ßx.ToObject()); πE != nil {
												continue
											}
											if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp002 {
												goto Label1
											}
											goto Label2
											// line 156: if conversion == 'x':
											πF.SetLineno(156)
										Label1:
											// line 157: return None
											πF.SetLineno(157)
											if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
												continue
											}
											πR = πTemp001
											continue
											goto Label2
										Label2:
											// line 158: return super(XFormatter, self).convert_field(value, conversion)
											πF.SetLineno(158)
											πTemp003 = πF.MakeArgs(2)
											if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
												continue
											}
											πTemp003[0] = µvalue
											if πE = πg.CheckLocal(πF, µconversion, "conversion"); πE != nil {
												continue
											}
											πTemp003[1] = µconversion
											πTemp004 = πF.MakeArgs(2)
											if πE = πg.CheckLocal(πF, µXFormatter, "XFormatter"); πE != nil {
												continue
											}
											πTemp004[0] = µXFormatter
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											πTemp004[1] = µself
											if πTemp001, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
												continue
											}
											if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp004)
											if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßconvert_field, nil); πE != nil {
												continue
											}
											if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp003)
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
									if πE = πClass.SetItem(πF, ßconvert_field.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp002 == nil {
								πTemp002 = πg.TypeType.ToObject()
							}
							if πTemp007, πE = πTemp002.Call(πF, []*πg.Object{πg.NewStr("XFormatter").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
								continue
							}
							µXFormatter = πTemp007
							// line 160: fmt = XFormatter()
							πF.SetLineno(160)
							if πE = πg.CheckLocal(πF, µXFormatter, "XFormatter"); πE != nil {
								continue
							}
							if πTemp001, πE = µXFormatter.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfmt = πTemp001
							// line 161: self.assertEqual(fmt.format("{0!r}:{0!x}", 'foo', 'foo'), "'foo':None")
							πF.SetLineno(161)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(3)
							πTemp004[0] = πg.NewStr("{0!r}:{0!x}").ToObject()
							πTemp004[1] = ßfoo.ToObject()
							πTemp004[2] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("'foo':None").ToObject()
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
							// line 165: class BarFormatter(string.Formatter):
							πF.SetLineno(165)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßFormatter, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp006 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("BarFormatter", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 168: def parse(self, format_string):
									πF.SetLineno(168)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "format_string", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("parse", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µformat_string *πg.Object = πArgs[1]; _ = µformat_string
										var µfield *πg.Object = πg.UnboundLocal; _ = µfield
										var µfield_name *πg.Object = πg.UnboundLocal; _ = µfield_name
										var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
										var µformat_spec *πg.Object = πg.UnboundLocal; _ = µformat_spec
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
										var πTemp009 *πg.Object
										_ = πTemp009
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
											for ; πF.State() >= 0; πF.PopCheckpoint() {
												switch πF.State() {
												case 0:
												case 8: goto Label8
												case 1: goto Label1
												case 2: goto Label2
												case 7: goto Label7
												default: panic("unexpected function state")
												}
												πTemp002 = πF.MakeArgs(1)
												πTemp002[0] = πg.NewStr("|").ToObject()
												if πE = πg.CheckLocal(πF, µformat_string, "format_string"); πE != nil {
													continue
												}
												if πTemp003, πE = πg.GetAttr(πF, µformat_string, ßsplit, nil); πE != nil {
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
													µfield = πTemp003
												}
												if πE != nil || !πTemp006 {
													continue
												}
												πF.PushCheckpoint(1)            
												πTemp004 = πg.NewInt(0).ToObject()
												if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
													continue
												}
												if πTemp007, πE = πg.GetItem(πF, µfield, πTemp004); πE != nil {
													continue
												}
												if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewStr("+").ToObject()); πE != nil {
													continue
												}
												if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
													continue
												}
												if πTemp006 {
													goto Label4
												}
												goto Label5
												// line 170: if field[0] == '+':
												πF.SetLineno(170)
											Label4:
												// line 172: field_name, _, format_spec = field[1:].partition(':')
												πF.SetLineno(172)
												πTemp002 = πF.MakeArgs(1)
												πTemp002[0] = πg.NewStr(":").ToObject()
												if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
													continue
												}
												if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
													continue
												}
												if πTemp004, πE = πg.GetItem(πF, µfield, πTemp003); πE != nil {
													continue
												}
												if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßpartition, nil); πE != nil {
													continue
												}
												if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
													continue
												}
												πF.FreeArgs(πTemp002)
												if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
													continue
												}
												µfield_name = πTemp003
												µ_ = πTemp007
												µformat_spec = πTemp008
												// line 173: yield '', field_name, format_spec, None
												πF.SetLineno(173)
												if πE = πg.CheckLocal(πF, µfield_name, "field_name"); πE != nil {
													continue
												}
												if πE = πg.CheckLocal(πF, µformat_spec, "format_spec"); πE != nil {
													continue
												}
												if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
													continue
												}
												πTemp003 = πg.NewTuple4(ß.ToObject(), µfield_name, µformat_spec, πTemp004).ToObject()
												πF.PushCheckpoint(7)
												return πTemp003, nil
											Label7:
												πTemp004 = πSent
												goto Label6
											Label5:
												// line 175: yield field, None, None, None
												πF.SetLineno(175)
												if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
													continue
												}
												if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
													continue
												}
												if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
													continue
												}
												if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
													continue
												}
												πTemp004 = πg.NewTuple4(µfield, πTemp007, πTemp008, πTemp009).ToObject()
												πF.PushCheckpoint(8)
												return πTemp004, nil
											Label8:
												πTemp007 = πSent
												goto Label6
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
									if πE = πClass.SetItem(πF, ßparse.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp002 == nil {
								πTemp002 = πg.TypeType.ToObject()
							}
							if πTemp007, πE = πTemp002.Call(πF, []*πg.Object{πg.NewStr("BarFormatter").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
								continue
							}
							µBarFormatter = πTemp007
							// line 177: fmt = BarFormatter()
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µBarFormatter, "BarFormatter"); πE != nil {
								continue
							}
							if πTemp001, πE = µBarFormatter.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfmt = πTemp001
							// line 178: self.assertEqual(fmt.format('*|+0:^10s|*', 'foo'), '*   foo    *')
							πF.SetLineno(178)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("*|+0:^10s|*").ToObject()
							πTemp004[1] = ßfoo.ToObject()
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("*   foo    *").ToObject()
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
							// line 181: class CheckAllUsedFormatter(string.Formatter):
							πF.SetLineno(181)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßFormatter, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp008
							πTemp006 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp006.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("CheckAllUsedFormatter", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 182: def check_unused_args(self, used_args, args, kwargs):
									πF.SetLineno(182)
									πTemp002 = make([]πg.Param, 4)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "used_args", Def: nil}
									πTemp002[2] = πg.Param{Name: "args", Def: nil}
									πTemp002[3] = πg.Param{Name: "kwargs", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("check_unused_args", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µused_args *πg.Object = πArgs[1]; _ = µused_args
										var µargs *πg.Object = πArgs[2]; _ = µargs
										var µkwargs *πg.Object = πArgs[3]; _ = µkwargs
										var µunused_args *πg.Object = πg.UnboundLocal; _ = µunused_args
										var µarg *πg.Object = πg.UnboundLocal; _ = µarg
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
											default: panic("unexpected function state")
											}
											// line 184: unused_args = set(kwargs.keys())
											πF.SetLineno(184)
											πTemp001 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µkwargs, ßkeys, nil); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
												continue
											}
											πTemp001[0] = πTemp003
											if πTemp002, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp001)
											µunused_args = πTemp003
											// line 185: unused_args.update(range(0, len(args)))
											πF.SetLineno(185)
											πTemp001 = πF.MakeArgs(1)
											πTemp004 = πF.MakeArgs(2)
											πTemp004[0] = πg.NewInt(0).ToObject()
											πTemp005 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
												continue
											}
											πTemp005[0] = µargs
											if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp005)
											πTemp004[1] = πTemp003
											if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp004)
											πTemp001[0] = πTemp003
											if πE = πg.CheckLocal(πF, µunused_args, "unused_args"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.GetAttr(πF, µunused_args, ßupdate, nil); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp001)
											if πE = πg.CheckLocal(πF, µused_args, "used_args"); πE != nil {
												continue
											}
											if πTemp002, πE = πg.Iter(πF, µused_args); πE != nil {
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
												µarg = πTemp003
											}
											if πE != nil || !πTemp007 {
												continue
											}
											πF.PushCheckpoint(1)            
											// line 188: unused_args.remove(arg)
											πF.SetLineno(188)
											πTemp001 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
												continue
											}
											πTemp001[0] = µarg
											if πE = πg.CheckLocal(πF, µunused_args, "unused_args"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetAttr(πF, µunused_args, ßremove, nil); πE != nil {
												continue
											}
											if πTemp008, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp001)
											continue
										Label2:
											if πE != nil || πR != nil {
												continue
											}
										Label3:
											if πE = πg.CheckLocal(πF, µunused_args, "unused_args"); πE != nil {
												continue
											}
											if πTemp006, πE = πg.IsTrue(πF, µunused_args); πE != nil {
												continue
											}
											if πTemp006 {
												goto Label4
											}
											goto Label5
											// line 190: if unused_args:
											πF.SetLineno(190)
										Label4:
											πTemp001 = πF.MakeArgs(1)
											πTemp001[0] = πg.NewStr("unused arguments").ToObject()
											if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
												continue
											}
											if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp001)
											// line 191: raise ValueError("unused arguments")
											πF.SetLineno(191)
											πE = πF.Raise(πTemp003, nil, nil)
											continue
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
									if πE = πClass.SetItem(πF, ßcheck_unused_args.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp002, πE = πTemp006.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp002 == nil {
								πTemp002 = πg.TypeType.ToObject()
							}
							if πTemp007, πE = πTemp002.Call(πF, []*πg.Object{πg.NewStr("CheckAllUsedFormatter").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp006.ToObject()}, nil); πE != nil {
								continue
							}
							µCheckAllUsedFormatter = πTemp007
							// line 193: fmt = CheckAllUsedFormatter()
							πF.SetLineno(193)
							if πE = πg.CheckLocal(πF, µCheckAllUsedFormatter, "CheckAllUsedFormatter"); πE != nil {
								continue
							}
							if πTemp001, πE = µCheckAllUsedFormatter.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfmt = πTemp001
							// line 194: self.assertEqual(fmt.format("{0}", 10), "10")
							πF.SetLineno(194)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("{0}").ToObject()
							πTemp004[1] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = ß10.ToObject()
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
							// line 195: self.assertEqual(fmt.format("{0}{i}", 10, i=100), "10100")
							πF.SetLineno(195)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("{0}{i}").ToObject()
							πTemp004[1] = πg.NewInt(10).ToObject()
							πTemp005 = πg.KWArgs{
								{"i", πg.NewInt(100).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = ß10100.ToObject()
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
							// line 196: self.assertEqual(fmt.format("{0}{i}{1}", 10, 20, i=100), "1010020")
							πF.SetLineno(196)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(3)
							πTemp004[0] = πg.NewStr("{0}{i}{1}").ToObject()
							πTemp004[1] = πg.NewInt(10).ToObject()
							πTemp004[2] = πg.NewInt(20).ToObject()
							πTemp005 = πg.KWArgs{
								{"i", πg.NewInt(100).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = ß1010020.ToObject()
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
							// line 197: self.assertRaises(ValueError, fmt.format, "{0}{i}{1}", 10, 20, i=100, j=0)
							πF.SetLineno(197)
							πTemp003 = πF.MakeArgs(5)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = πg.NewStr("{0}{i}{1}").ToObject()
							πTemp003[3] = πg.NewInt(10).ToObject()
							πTemp003[4] = πg.NewInt(20).ToObject()
							πTemp005 = πg.KWArgs{
								{"i", πg.NewInt(100).ToObject()},
								{"j", πg.NewInt(0).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 198: self.assertRaises(ValueError, fmt.format, "{0}", 10, 20)
							πF.SetLineno(198)
							πTemp003 = πF.MakeArgs(5)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = πg.NewStr("{0}").ToObject()
							πTemp003[3] = πg.NewInt(10).ToObject()
							πTemp003[4] = πg.NewInt(20).ToObject()
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
							// line 199: self.assertRaises(ValueError, fmt.format, "{0}", 10, 20, i=100)
							πF.SetLineno(199)
							πTemp003 = πF.MakeArgs(5)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = πg.NewStr("{0}").ToObject()
							πTemp003[3] = πg.NewInt(10).ToObject()
							πTemp003[4] = πg.NewInt(20).ToObject()
							πTemp005 = πg.KWArgs{
								{"i", πg.NewInt(100).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 200: self.assertRaises(ValueError, fmt.format, "{i}", 10, 20, i=100)
							πF.SetLineno(200)
							πTemp003 = πF.MakeArgs(5)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = πg.NewStr("{i}").ToObject()
							πTemp003[3] = πg.NewInt(10).ToObject()
							πTemp003[4] = πg.NewInt(20).ToObject()
							πTemp005 = πg.KWArgs{
								{"i", πg.NewInt(100).ToObject()},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 203: self.assertRaises(ValueError, format, '', '#')
							πF.SetLineno(203)
							πTemp003 = πF.MakeArgs(4)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßformat); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = ß.ToObject()
							πTemp003[3] = πg.NewStr("#").ToObject()
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
							// line 204: self.assertRaises(ValueError, format, '', '#20')
							πF.SetLineno(204)
							πTemp003 = πF.MakeArgs(4)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßformat); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = ß.ToObject()
							πTemp003[3] = πg.NewStr("#20").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_formatter.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 115: @unittest.expectedFailure
					πF.SetLineno(115)
					πTemp006 = πF.MakeArgs(1)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßtest_formatter); πE != nil {
						continue
					}
					πTemp006[0] = πTemp010
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp011.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_formatter.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 207: def test_format_keyword_arguments(self):
					πF.SetLineno(207)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_format_keyword_arguments", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfmt *πg.Object = πg.UnboundLocal; _ = µfmt
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 208: fmt = string.Formatter()
							πF.SetLineno(208)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßFormatter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfmt = πTemp001
							// line 209: self.assertEqual(fmt.format("-{arg}-", arg='test'), '-test-')
							πF.SetLineno(209)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("-{arg}-").ToObject()
							πTemp005 = πg.KWArgs{
								{"arg", ßtest.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("-test-").ToObject()
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
							// line 210: self.assertRaises(KeyError, fmt.format, "-{arg}-")
							πF.SetLineno(210)
							πTemp003 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = πg.NewStr("-{arg}-").ToObject()
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
							// line 211: self.assertEqual(fmt.format("-{self}-", self='test'), '-test-')
							πF.SetLineno(211)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("-{self}-").ToObject()
							πTemp005 = πg.KWArgs{
								{"self", ßtest.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("-test-").ToObject()
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
							// line 212: self.assertRaises(KeyError, fmt.format, "-{self}-")
							πF.SetLineno(212)
							πTemp003 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = πg.NewStr("-{self}-").ToObject()
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
							// line 213: self.assertEqual(fmt.format("-{format_string}-", format_string='test'),
							πF.SetLineno(213)
							πTemp003 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("-{format_string}-").ToObject()
							πTemp005 = πg.KWArgs{
								{"format_string", ßtest.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("-test-").ToObject()
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
							// line 215: self.assertRaises(KeyError, fmt.format, "-{format_string}-")
							πF.SetLineno(215)
							πTemp003 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp003[2] = πg.NewStr("-{format_string}-").ToObject()
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
							// line 216: self.assertEqual(fmt.format(arg='test', format_string="-{arg}-"),
							πF.SetLineno(216)
							πTemp003 = πF.MakeArgs(2)
							πTemp005 = πg.KWArgs{
								{"arg", ßtest.ToObject()},
								{"format_string", πg.NewStr("-{arg}-").ToObject()},
							}
							if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfmt, ßformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewStr("-test-").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_format_keyword_arguments.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 206: @unittest.expectedFailure
					πF.SetLineno(206)
					πTemp006 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtest_format_keyword_arguments); πE != nil {
						continue
					}
					πTemp006[0] = πTemp011
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp012.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßtest_format_keyword_arguments.ToObject(), πTemp011); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("ModuleTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßModuleTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 219: class BytesAliasTest(unittest.TestCase):
			πF.SetLineno(219)
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
			_, πE = πg.NewCode("BytesAliasTest", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 222: def test_builtin(self):
					πF.SetLineno(222)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_builtin", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 223: self.assertTrue(str is bytes)
							πF.SetLineno(223)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßbytes); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_builtin.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 221: @unittest.expectedFailure
					πF.SetLineno(221)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtest_builtin); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßtest_builtin.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 225: def test_syntax(self):
					πF.SetLineno(225)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_syntax", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 226: self.assertEqual(b"spam", "spam")
							πF.SetLineno(226)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = ßspam.ToObject()
							πTemp001[1] = ßspam.ToObject()
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
							// line 227: self.assertEqual(br"egg\foo", "egg\\foo")
							πF.SetLineno(227)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("egg\\foo").ToObject()
							πTemp001[1] = πg.NewStr("egg\\foo").ToObject()
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
							// line 228: self.assertTrue(type(b""), str)
							πF.SetLineno(228)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ß.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
							// line 229: self.assertTrue(type(br""), str)
							πF.SetLineno(229)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ß.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_syntax.ToObject(), πTemp004); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BytesAliasTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBytesAliasTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 234: class Bag(object):
			πF.SetLineno(234)
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
			_, πE = πg.NewCode("Bag", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 235: pass
					πF.SetLineno(235)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Bag").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBag.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 237: class Mapping(object):
			πF.SetLineno(237)
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
			_, πE = πg.NewCode("Mapping", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 238: def __getitem__(self, name):
					πF.SetLineno(238)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "name", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var µpart *πg.Object = πg.UnboundLocal; _ = µpart
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
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
							// line 239: obj = self
							πF.SetLineno(239)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							µobj = µself
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(".").ToObject()
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µname, ßsplit, nil); πE != nil {
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
								µpart = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 241: try:
							πF.SetLineno(241)
							πF.PushCheckpoint(5)
							// line 242: obj = getattr(obj, part)
							πF.SetLineno(242)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πE = πg.CheckLocal(πF, µpart, "part"); πE != nil {
								continue
							}
							πTemp002[1] = µpart
							if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µobj = πTemp004
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 243: except AttributeError:
							πF.SetLineno(243)
						Label6:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002[0] = µname
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 244: raise KeyError(name)
							πF.SetLineno(244)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 245: return obj
							πF.SetLineno(245)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πR = µobj
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
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Mapping").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMapping.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 248: class TestTemplate(unittest.TestCase):
			πF.SetLineno(248)
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
			_, πE = πg.NewCode("TestTemplate", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp018 *πg.Object
				_ = πTemp018
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 249: def test_regular_templates(self):
					πF.SetLineno(249)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_regular_templates", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 250: s = Template('$who likes to eat a bag of $what worth $$100')
							πF.SetLineno(250)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("$who likes to eat a bag of $what worth $$100").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 251: self.assertEqual(s.substitute(dict(who='tim', what='ham')),
							πF.SetLineno(251)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("tim likes to eat a bag of ham worth $100").ToObject()
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
							// line 253: self.assertRaises(KeyError, s.substitute, dict(who='tim'))
							πF.SetLineno(253)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp005 = πg.KWArgs{
								{"who", ßtim.ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
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
							// line 254: self.assertRaises(TypeError, Template.substitute)
							πF.SetLineno(254)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
					if πE = πClass.SetItem(πF, ßtest_regular_templates.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 256: def test_regular_templates_with_braces(self):
					πF.SetLineno(256)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_regular_templates_with_braces", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							// line 257: s = Template('$who likes ${what} for ${meal}')
							πF.SetLineno(257)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("$who likes ${what} for ${meal}").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 258: d = dict(who='tim', what='ham', meal='dinner')
							πF.SetLineno(258)
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
								{"meal", ßdinner.ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							µd = πTemp003
							// line 259: self.assertEqual(s.substitute(d), 'tim likes ham for dinner')
							πF.SetLineno(259)
							πTemp001 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("tim likes ham for dinner").ToObject()
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
							// line 260: self.assertRaises(KeyError, s.substitute,
							πF.SetLineno(260)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
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
					if πE = πClass.SetItem(πF, ßtest_regular_templates_with_braces.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 263: def test_escapes(self):
					πF.SetLineno(263)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_escapes", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 264: eq = self.assertEqual
							πF.SetLineno(264)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 265: s = Template('$who likes to eat a bag of $$what worth $$100')
							πF.SetLineno(265)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("$who likes to eat a bag of $$what worth $$100").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 266: eq(s.substitute(dict(who='tim', what='ham')),
							πF.SetLineno(266)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes to eat a bag of $what worth $100").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 268: s = Template('$who likes $$')
							πF.SetLineno(268)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("$who likes $$").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 269: eq(s.substitute(dict(who='tim', what='ham')), 'tim likes $')
							πF.SetLineno(269)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes $").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_escapes.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 271: def test_percents(self):
					πF.SetLineno(271)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_percents", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							// line 272: eq = self.assertEqual
							πF.SetLineno(272)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 273: s = Template('%(foo)s $foo ${foo}')
							πF.SetLineno(273)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("%(foo)s $foo ${foo}").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 274: d = dict(foo='baz')
							πF.SetLineno(274)
							πTemp004 = πg.KWArgs{
								{"foo", ßbaz.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							µd = πTemp003
							// line 275: eq(s.substitute(d), '%(foo)s baz baz')
							πF.SetLineno(275)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("%(foo)s baz baz").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 276: eq(s.safe_substitute(d), '%(foo)s baz baz')
							πF.SetLineno(276)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("%(foo)s baz baz").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_percents.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 278: def test_stringification(self):
					πF.SetLineno(278)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_stringification", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							// line 279: eq = self.assertEqual
							πF.SetLineno(279)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 280: s = Template('tim has eaten $count bags of ham today')
							πF.SetLineno(280)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("tim has eaten $count bags of ham today").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 281: d = dict(count=7)
							πF.SetLineno(281)
							πTemp004 = πg.KWArgs{
								{"count", πg.NewInt(7).ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							µd = πTemp003
							// line 282: eq(s.substitute(d), 'tim has eaten 7 bags of ham today')
							πF.SetLineno(282)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim has eaten 7 bags of ham today").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 283: eq(s.safe_substitute(d), 'tim has eaten 7 bags of ham today')
							πF.SetLineno(283)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim has eaten 7 bags of ham today").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 284: s = Template('tim has eaten ${count} bags of ham today')
							πF.SetLineno(284)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("tim has eaten ${count} bags of ham today").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 285: eq(s.substitute(d), 'tim has eaten 7 bags of ham today')
							πF.SetLineno(285)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim has eaten 7 bags of ham today").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_stringification.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 287: def test_tupleargs(self):
					πF.SetLineno(287)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_tupleargs", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							// line 288: eq = self.assertEqual
							πF.SetLineno(288)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 289: s = Template('$who ate ${meal}')
							πF.SetLineno(289)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("$who ate ${meal}").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 290: d = dict(who=('tim', 'fred'), meal=('ham', 'kung pao'))
							πF.SetLineno(290)
							πTemp001 = πg.NewTuple2(ßtim.ToObject(), ßfred.ToObject()).ToObject()
							πTemp003 = πg.NewTuple2(ßham.ToObject(), πg.NewStr("kung pao").ToObject()).ToObject()
							πTemp004 = πg.KWArgs{
								{"who", πTemp001},
								{"meal", πTemp003},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							µd = πTemp003
							// line 291: eq(s.substitute(d), "('tim', 'fred') ate ('ham', 'kung pao')")
							πF.SetLineno(291)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("('tim', 'fred') ate ('ham', 'kung pao')").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 292: eq(s.safe_substitute(d), "('tim', 'fred') ate ('ham', 'kung pao')")
							πF.SetLineno(292)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("('tim', 'fred') ate ('ham', 'kung pao')").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_tupleargs.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 294: def test_SafeTemplate(self):
					πF.SetLineno(294)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_SafeTemplate", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
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
							// line 295: eq = self.assertEqual
							πF.SetLineno(295)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 296: s = Template('$who likes ${what} for ${meal}')
							πF.SetLineno(296)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("$who likes ${what} for ${meal}").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 297: eq(s.safe_substitute(dict(who='tim')), 'tim likes ${what} for ${meal}')
							πF.SetLineno(297)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πg.KWArgs{
								{"who", ßtim.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes ${what} for ${meal}").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 298: eq(s.safe_substitute(dict(what='ham')), '$who likes ham for ${meal}')
							πF.SetLineno(298)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πg.KWArgs{
								{"what", ßham.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("$who likes ham for ${meal}").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 299: eq(s.safe_substitute(dict(what='ham', meal='dinner')),
							πF.SetLineno(299)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πg.KWArgs{
								{"what", ßham.ToObject()},
								{"meal", ßdinner.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("$who likes ham for dinner").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 301: eq(s.safe_substitute(dict(who='tim', what='ham')),
							πF.SetLineno(301)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes ham for ${meal}").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 303: eq(s.safe_substitute(dict(who='tim', what='ham', meal='dinner')),
							πF.SetLineno(303)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
								{"meal", ßdinner.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes ham for dinner").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_SafeTemplate.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 307: def test_invalid_placeholders(self):
					πF.SetLineno(307)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_invalid_placeholders", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µraises *πg.Object = πg.UnboundLocal; _ = µraises
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 308: raises = self.assertRaises
							πF.SetLineno(308)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							µraises = πTemp001
							// line 309: s = Template('$who likes $')
							πF.SetLineno(309)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("$who likes $").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 310: raises(ValueError, s.substitute, dict(who='tim'))
							πF.SetLineno(310)
							πTemp002 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							if πE = πg.CheckLocal(πF, µraises, "raises"); πE != nil {
								continue
							}
							if πTemp001, πE = µraises.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 311: s = Template('$who likes ${what)')
							πF.SetLineno(311)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("$who likes ${what)").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 312: raises(ValueError, s.substitute, dict(who='tim'))
							πF.SetLineno(312)
							πTemp002 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							if πE = πg.CheckLocal(πF, µraises, "raises"); πE != nil {
								continue
							}
							if πTemp001, πE = µraises.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 313: s = Template('$who likes $100')
							πF.SetLineno(313)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("$who likes $100").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 314: raises(ValueError, s.substitute, dict(who='tim'))
							πF.SetLineno(314)
							πTemp002 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							if πE = πg.CheckLocal(πF, µraises, "raises"); πE != nil {
								continue
							}
							if πTemp001, πE = µraises.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_invalid_placeholders.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 306: @unittest.expectedFailure
					πF.SetLineno(306)
					πTemp010 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtest_invalid_placeholders); πE != nil {
						continue
					}
					πTemp010[0] = πTemp011
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp012.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßtest_invalid_placeholders.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 316: def test_idpattern_override(self):
					πF.SetLineno(316)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_idpattern_override", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µPathPattern *πg.Object = πg.UnboundLocal; _ = µPathPattern
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µs *πg.Object = πg.UnboundLocal; _ = µs
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
							// line 317: class PathPattern(Template):
							πF.SetLineno(317)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
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
							_, πE = πg.NewCode("PathPattern", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 318: idpattern = r'[_a-z][._a-z0-9]*'
									πF.SetLineno(318)
									// line 318: idpattern = r'[_a-z][._a-z0-9]*'
									πF.SetLineno(318)
									if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("[_a-z][._a-z0-9]*").ToObject()); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßidpattern.ToObject(), πg.NewStr("[_a-z][._a-z0-9]*").ToObject()); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("PathPattern").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µPathPattern = πTemp005
							// line 319: m = Mapping()
							πF.SetLineno(319)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßMapping); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µm = πTemp004
							// line 320: m.bag = Bag()
							πF.SetLineno(320)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBag); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µm, ßbag, πTemp002); πE != nil {
								continue
							}
							// line 321: m.bag.foo = Bag()
							πF.SetLineno(321)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBag); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µm, ßbag, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßfoo, πTemp002); πE != nil {
								continue
							}
							// line 322: m.bag.foo.who = 'tim'
							πF.SetLineno(322)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßtim.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßbag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßfoo, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßwho, πTemp002); πE != nil {
								continue
							}
							// line 323: m.bag.what = 'ham'
							πF.SetLineno(323)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßham.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßbag, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßwhat, πTemp002); πE != nil {
								continue
							}
							// line 324: s = PathPattern('$bag.foo.who likes to eat a bag of $bag.what')
							πF.SetLineno(324)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("$bag.foo.who likes to eat a bag of $bag.what").ToObject()
							if πE = πg.CheckLocal(πF, µPathPattern, "PathPattern"); πE != nil {
								continue
							}
							if πTemp002, πE = µPathPattern.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs = πTemp002
							// line 325: self.assertEqual(s.substitute(m), 'tim likes to eat a bag of ham')
							πF.SetLineno(325)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp006[0] = µm
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewStr("tim likes to eat a bag of ham").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_idpattern_override.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 327: def test_pattern_override(self):
					πF.SetLineno(327)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_pattern_override", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µMyPattern *πg.Object = πg.UnboundLocal; _ = µMyPattern
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µBadPattern *πg.Object = πg.UnboundLocal; _ = µBadPattern
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
							// line 328: class MyPattern(Template):
							πF.SetLineno(328)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
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
							_, πE = πg.NewCode("MyPattern", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 329: pattern = r"""
									πF.SetLineno(329)
									// line 329: pattern = r"""
									πF.SetLineno(329)
									if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n            (?P<escaped>@{2})                   |\n            @(?P<named>[_a-z][._a-z0-9]*)       |\n            @{(?P<braced>[_a-z][._a-z0-9]*)}    |\n            (?P<invalid>@)\n            ").ToObject()); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßpattern.ToObject(), πg.NewStr("\n            (?P<escaped>@{2})                   |\n            @(?P<named>[_a-z][._a-z0-9]*)       |\n            @{(?P<braced>[_a-z][._a-z0-9]*)}    |\n            (?P<invalid>@)\n            ").ToObject()); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MyPattern").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µMyPattern = πTemp005
							// line 335: m = Mapping()
							πF.SetLineno(335)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßMapping); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µm = πTemp004
							// line 336: m.bag = Bag()
							πF.SetLineno(336)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBag); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µm, ßbag, πTemp002); πE != nil {
								continue
							}
							// line 337: m.bag.foo = Bag()
							πF.SetLineno(337)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBag); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µm, ßbag, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßfoo, πTemp002); πE != nil {
								continue
							}
							// line 338: m.bag.foo.who = 'tim'
							πF.SetLineno(338)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßtim.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßbag, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßfoo, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßwho, πTemp002); πE != nil {
								continue
							}
							// line 339: m.bag.what = 'ham'
							πF.SetLineno(339)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßham.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µm, ßbag, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßwhat, πTemp002); πE != nil {
								continue
							}
							// line 340: s = MyPattern('@bag.foo.who likes to eat a bag of @bag.what')
							πF.SetLineno(340)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("@bag.foo.who likes to eat a bag of @bag.what").ToObject()
							if πE = πg.CheckLocal(πF, µMyPattern, "MyPattern"); πE != nil {
								continue
							}
							if πTemp002, πE = µMyPattern.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs = πTemp002
							// line 341: self.assertEqual(s.substitute(m), 'tim likes to eat a bag of ham')
							πF.SetLineno(341)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							πTemp006[0] = µm
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewStr("tim likes to eat a bag of ham").ToObject()
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
							// line 343: class BadPattern(Template):
							πF.SetLineno(343)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
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
							_, πE = πg.NewCode("BadPattern", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 344: pattern = r"""
									πF.SetLineno(344)
									// line 344: pattern = r"""
									πF.SetLineno(344)
									if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n            (?P<badname>.*)                     |\n            (?P<escaped>@{2})                   |\n            @(?P<named>[_a-z][._a-z0-9]*)       |\n            @{(?P<braced>[_a-z][._a-z0-9]*)}    |\n            (?P<invalid>@)                      |\n            ").ToObject()); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßpattern.ToObject(), πg.NewStr("\n            (?P<badname>.*)                     |\n            (?P<escaped>@{2})                   |\n            @(?P<named>[_a-z][._a-z0-9]*)       |\n            @{(?P<braced>[_a-z][._a-z0-9]*)}    |\n            (?P<invalid>@)                      |\n            ").ToObject()); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("BadPattern").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µBadPattern = πTemp005
							// line 351: s = BadPattern('@bag.foo.who likes to eat a bag of @bag.what')
							πF.SetLineno(351)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("@bag.foo.who likes to eat a bag of @bag.what").ToObject()
							if πE = πg.CheckLocal(πF, µBadPattern, "BadPattern"); πE != nil {
								continue
							}
							if πTemp002, πE = µBadPattern.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs = πTemp002
							// line 352: self.assertRaises(ValueError, s.substitute, {})
							πF.SetLineno(352)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
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
							// line 353: self.assertRaises(ValueError, s.safe_substitute, {})
							πF.SetLineno(353)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_pattern_override.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 355: def test_braced_override(self):
					πF.SetLineno(355)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_braced_override", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µMyTemplate *πg.Object = πg.UnboundLocal; _ = µMyTemplate
						var µtmpl *πg.Object = πg.UnboundLocal; _ = µtmpl
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var µval *πg.Object = πg.UnboundLocal; _ = µval
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
							// line 356: class MyTemplate(Template):
							πF.SetLineno(356)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
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
							_, πE = πg.NewCode("MyTemplate", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 357: pattern = r"""
									πF.SetLineno(357)
									// line 357: pattern = r"""
									πF.SetLineno(357)
									if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n            \\$(?:\n              (?P<escaped>$)                     |\n              (?P<named>[_a-z][_a-z0-9]*)        |\n              @@(?P<braced>[_a-z][_a-z0-9]*)@@   |\n              (?P<invalid>)                      |\n           )\n           ").ToObject()); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßpattern.ToObject(), πg.NewStr("\n            \\$(?:\n              (?P<escaped>$)                     |\n              (?P<named>[_a-z][_a-z0-9]*)        |\n              @@(?P<braced>[_a-z][_a-z0-9]*)@@   |\n              (?P<invalid>)                      |\n           )\n           ").ToObject()); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MyTemplate").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µMyTemplate = πTemp005
							// line 366: tmpl = 'PyCon in $@@location@@'
							πF.SetLineno(366)
							µtmpl = πg.NewStr("PyCon in $@@location@@").ToObject()
							// line 367: t = MyTemplate(tmpl)
							πF.SetLineno(367)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtmpl, "tmpl"); πE != nil {
								continue
							}
							πTemp003[0] = µtmpl
							if πE = πg.CheckLocal(πF, µMyTemplate, "MyTemplate"); πE != nil {
								continue
							}
							if πTemp002, πE = µMyTemplate.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µt = πTemp002
							// line 368: self.assertRaises(KeyError, t.substitute, {})
							πF.SetLineno(368)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µt, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
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
							// line 369: val = t.substitute({'location': 'Cleveland'})
							πF.SetLineno(369)
							πTemp003 = πF.MakeArgs(1)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßlocation.ToObject(), ßCleveland.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µt, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µval = πTemp004
							// line 370: self.assertEqual(val, 'PyCon in Cleveland')
							πF.SetLineno(370)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							πTemp003[0] = µval
							πTemp003[1] = πg.NewStr("PyCon in Cleveland").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_braced_override.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 372: def test_braced_override_safe(self):
					πF.SetLineno(372)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_braced_override_safe", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µMyTemplate *πg.Object = πg.UnboundLocal; _ = µMyTemplate
						var µtmpl *πg.Object = πg.UnboundLocal; _ = µtmpl
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var µval *πg.Object = πg.UnboundLocal; _ = µval
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
							// line 373: class MyTemplate(Template):
							πF.SetLineno(373)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
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
							_, πE = πg.NewCode("MyTemplate", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 374: pattern = r"""
									πF.SetLineno(374)
									// line 374: pattern = r"""
									πF.SetLineno(374)
									if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n            \\$(?:\n              (?P<escaped>$)                     |\n              (?P<named>[_a-z][_a-z0-9]*)        |\n              @@(?P<braced>[_a-z][_a-z0-9]*)@@   |\n              (?P<invalid>)                      |\n           )\n           ").ToObject()); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßpattern.ToObject(), πg.NewStr("\n            \\$(?:\n              (?P<escaped>$)                     |\n              (?P<named>[_a-z][_a-z0-9]*)        |\n              @@(?P<braced>[_a-z][_a-z0-9]*)@@   |\n              (?P<invalid>)                      |\n           )\n           ").ToObject()); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("MyTemplate").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µMyTemplate = πTemp005
							// line 383: tmpl = 'PyCon in $@@location@@'
							πF.SetLineno(383)
							µtmpl = πg.NewStr("PyCon in $@@location@@").ToObject()
							// line 384: t = MyTemplate(tmpl)
							πF.SetLineno(384)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtmpl, "tmpl"); πE != nil {
								continue
							}
							πTemp003[0] = µtmpl
							if πE = πg.CheckLocal(πF, µMyTemplate, "MyTemplate"); πE != nil {
								continue
							}
							if πTemp002, πE = µMyTemplate.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µt = πTemp002
							// line 385: self.assertEqual(t.safe_substitute(), tmpl)
							πF.SetLineno(385)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µt, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µtmpl, "tmpl"); πE != nil {
								continue
							}
							πTemp003[1] = µtmpl
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
							// line 386: val = t.safe_substitute({'location': 'Cleveland'})
							πF.SetLineno(386)
							πTemp003 = πF.MakeArgs(1)
							πTemp001 = πg.NewDict()
							if πE = πTemp001.SetItem(πF, ßlocation.ToObject(), ßCleveland.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp001.ToObject()
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µt, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µval = πTemp004
							// line 387: self.assertEqual(val, 'PyCon in Cleveland')
							πF.SetLineno(387)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							πTemp003[0] = µval
							πTemp003[1] = πg.NewStr("PyCon in Cleveland").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_braced_override_safe.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 389: def test_unicode_values(self):
					πF.SetLineno(389)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_unicode_values", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							// line 390: s = Template('$who likes $what')
							πF.SetLineno(390)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("$who likes $what").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 391: d = dict(who=u't\xffm', what=u'f\xfe\fed')
							πF.SetLineno(391)
							πTemp004 = πg.KWArgs{
								{"who", πg.NewUnicode("t\xc3\xbfm").ToObject()},
								{"what", πg.NewUnicode("f\xc3\xbe\x0ced").ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							µd = πTemp003
							// line 392: self.assertEqual(s.substitute(d), u't\xffm likes f\xfe\x0ced')
							πF.SetLineno(392)
							πTemp001 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewUnicode("t\xc3\xbfm likes f\xc3\xbe\x0ced").ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_unicode_values.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 394: def test_keyword_arguments(self):
					πF.SetLineno(394)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_keyword_arguments", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
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
							// line 395: eq = self.assertEqual
							πF.SetLineno(395)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 396: s = Template('$who likes $what')
							πF.SetLineno(396)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("$who likes $what").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 397: eq(s.substitute(who='tim', what='ham'), 'tim likes ham')
							πF.SetLineno(397)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes ham").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 398: eq(s.substitute(dict(who='tim'), what='ham'), 'tim likes ham')
							πF.SetLineno(398)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							πTemp004 = πg.KWArgs{
								{"what", ßham.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes ham").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 399: eq(s.substitute(dict(who='fred', what='kung pao'),
							πF.SetLineno(399)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp004 = πg.KWArgs{
								{"who", ßfred.ToObject()},
								{"what", πg.NewStr("kung pao").ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes ham").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 402: s = Template('the mapping is $mapping')
							πF.SetLineno(402)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("the mapping is $mapping").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 403: eq(s.substitute(dict(foo='none'), mapping='bozo'),
							πF.SetLineno(403)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp004 = πg.KWArgs{
								{"foo", ßnone.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							πTemp004 = πg.KWArgs{
								{"mapping", ßbozo.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("the mapping is bozo").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 405: eq(s.substitute(dict(mapping='one'), mapping='two'),
							πF.SetLineno(405)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp004 = πg.KWArgs{
								{"mapping", ßone.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							πTemp004 = πg.KWArgs{
								{"mapping", ßtwo.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("the mapping is two").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 408: s = Template('the self is $self')
							πF.SetLineno(408)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("the self is $self").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 409: eq(s.substitute(self='bozo'), 'the self is bozo')
							πF.SetLineno(409)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πg.KWArgs{
								{"self", ßbozo.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("the self is bozo").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_keyword_arguments.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 411: def test_keyword_arguments_safe(self):
					πF.SetLineno(411)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("test_keyword_arguments_safe", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µraises *πg.Object = πg.UnboundLocal; _ = µraises
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 πg.KWArgs
						_ = πTemp004
						var πTemp005 []*πg.Object
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
							// line 412: eq = self.assertEqual
							πF.SetLineno(412)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 413: raises = self.assertRaises
							πF.SetLineno(413)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							µraises = πTemp001
							// line 414: s = Template('$who likes $what')
							πF.SetLineno(414)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("$who likes $what").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 415: eq(s.safe_substitute(who='tim', what='ham'), 'tim likes ham')
							πF.SetLineno(415)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes ham").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 416: eq(s.safe_substitute(dict(who='tim'), what='ham'), 'tim likes ham')
							πF.SetLineno(416)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							πTemp004 = πg.KWArgs{
								{"what", ßham.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes ham").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 417: eq(s.safe_substitute(dict(who='fred', what='kung pao'),
							πF.SetLineno(417)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp004 = πg.KWArgs{
								{"who", ßfred.ToObject()},
								{"what", πg.NewStr("kung pao").ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							πTemp004 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("tim likes ham").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 420: s = Template('the mapping is $mapping')
							πF.SetLineno(420)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("the mapping is $mapping").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 421: eq(s.safe_substitute(dict(foo='none'), mapping='bozo'),
							πF.SetLineno(421)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp004 = πg.KWArgs{
								{"foo", ßnone.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							πTemp004 = πg.KWArgs{
								{"mapping", ßbozo.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("the mapping is bozo").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 423: eq(s.safe_substitute(dict(mapping='one'), mapping='two'),
							πF.SetLineno(423)
							πTemp002 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp004 = πg.KWArgs{
								{"mapping", ßone.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							πTemp004 = πg.KWArgs{
								{"mapping", ßtwo.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("the mapping is two").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 425: d = dict(mapping='one')
							πF.SetLineno(425)
							πTemp004 = πg.KWArgs{
								{"mapping", ßone.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							µd = πTemp003
							// line 426: raises(TypeError, s.substitute, d, {})
							πF.SetLineno(426)
							πTemp002 = πF.MakeArgs(4)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002[2] = µd
							πTemp006 = πg.NewDict()
							πTemp001 = πTemp006.ToObject()
							πTemp002[3] = πTemp001
							if πE = πg.CheckLocal(πF, µraises, "raises"); πE != nil {
								continue
							}
							if πTemp001, πE = µraises.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 427: raises(TypeError, s.safe_substitute, d, {})
							πF.SetLineno(427)
							πTemp002 = πF.MakeArgs(4)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp002[2] = µd
							πTemp006 = πg.NewDict()
							πTemp001 = πTemp006.ToObject()
							πTemp002[3] = πTemp001
							if πE = πg.CheckLocal(πF, µraises, "raises"); πE != nil {
								continue
							}
							if πTemp001, πE = µraises.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 429: s = Template('the self is $self')
							πF.SetLineno(429)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("the self is $self").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 430: eq(s.safe_substitute(self='bozo'), 'the self is bozo')
							πF.SetLineno(430)
							πTemp002 = πF.MakeArgs(2)
							πTemp004 = πg.KWArgs{
								{"self", ßbozo.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, πTemp004); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp002[1] = πg.NewStr("the self is bozo").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_keyword_arguments_safe.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 432: def test_delimiter_override(self):
					πF.SetLineno(432)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("test_delimiter_override", "/usr/lib/python2.7/test/test_string.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µeq *πg.Object = πg.UnboundLocal; _ = µeq
						var µraises *πg.Object = πg.UnboundLocal; _ = µraises
						var µAmpersandTemplate *πg.Object = πg.UnboundLocal; _ = µAmpersandTemplate
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µPieDelims *πg.Object = πg.UnboundLocal; _ = µPieDelims
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Dict
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
							// line 433: eq = self.assertEqual
							πF.SetLineno(433)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							µeq = πTemp001
							// line 434: raises = self.assertRaises
							πF.SetLineno(434)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							µraises = πTemp001
							// line 435: class AmpersandTemplate(Template):
							πF.SetLineno(435)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							πTemp002 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp002.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("AmpersandTemplate", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp002
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 436: delimiter = '&'
									πF.SetLineno(436)
									// line 436: delimiter = '&'
									πF.SetLineno(436)
									if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("&").ToObject()); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßdelimiter.ToObject(), πg.NewStr("&").ToObject()); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("AmpersandTemplate").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp002.ToObject()}, nil); πE != nil {
								continue
							}
							µAmpersandTemplate = πTemp005
							// line 437: s = AmpersandTemplate('this &gift is for &{who} &&')
							πF.SetLineno(437)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("this &gift is for &{who} &&").ToObject()
							if πE = πg.CheckLocal(πF, µAmpersandTemplate, "AmpersandTemplate"); πE != nil {
								continue
							}
							if πTemp001, πE = µAmpersandTemplate.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs = πTemp001
							// line 438: eq(s.substitute(gift='bud', who='you'), 'this bud is for you &')
							πF.SetLineno(438)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πg.KWArgs{
								{"gift", ßbud.ToObject()},
								{"who", ßyou.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewStr("this bud is for you &").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 439: raises(KeyError, s.substitute)
							πF.SetLineno(439)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µraises, "raises"); πE != nil {
								continue
							}
							if πTemp001, πE = µraises.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 440: eq(s.safe_substitute(gift='bud', who='you'), 'this bud is for you &')
							πF.SetLineno(440)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πg.KWArgs{
								{"gift", ßbud.ToObject()},
								{"who", ßyou.ToObject()},
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewStr("this bud is for you &").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 441: eq(s.safe_substitute(), 'this &gift is for &{who} &')
							πF.SetLineno(441)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewStr("this &gift is for &{who} &").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 442: s = AmpersandTemplate('this &gift is for &{who} &')
							πF.SetLineno(442)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("this &gift is for &{who} &").ToObject()
							if πE = πg.CheckLocal(πF, µAmpersandTemplate, "AmpersandTemplate"); πE != nil {
								continue
							}
							if πTemp001, πE = µAmpersandTemplate.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs = πTemp001
							// line 443: raises(ValueError, s.substitute, dict(gift='bud', who='you'))
							πF.SetLineno(443)
							πTemp003 = πF.MakeArgs(3)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							πTemp006 = πg.KWArgs{
								{"gift", ßbud.ToObject()},
								{"who", ßyou.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							πTemp003[2] = πTemp004
							if πE = πg.CheckLocal(πF, µraises, "raises"); πE != nil {
								continue
							}
							if πTemp001, πE = µraises.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 444: eq(s.safe_substitute(), 'this &gift is for &{who} &')
							πF.SetLineno(444)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsafe_substitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewStr("this &gift is for &{who} &").ToObject()
							if πE = πg.CheckLocal(πF, µeq, "eq"); πE != nil {
								continue
							}
							if πTemp001, πE = µeq.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 446: class PieDelims(Template):
							πF.SetLineno(446)
							πTemp003 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTemplate); πE != nil {
								continue
							}
							πTemp003[0] = πTemp005
							πTemp002 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp002.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("PieDelims", "/usr/lib/python2.7/test/test_string.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp002
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 447: delimiter = '@'
									πF.SetLineno(447)
									// line 447: delimiter = '@'
									πF.SetLineno(447)
									if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("@").ToObject()); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßdelimiter.ToObject(), πg.NewStr("@").ToObject()); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp004 == nil {
								πTemp004 = πg.TypeType.ToObject()
							}
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("PieDelims").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp002.ToObject()}, nil); πE != nil {
								continue
							}
							µPieDelims = πTemp005
							// line 448: s = PieDelims('@who likes to eat a bag of @{what} worth $100')
							πF.SetLineno(448)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("@who likes to eat a bag of @{what} worth $100").ToObject()
							if πE = πg.CheckLocal(πF, µPieDelims, "PieDelims"); πE != nil {
								continue
							}
							if πTemp001, πE = µPieDelims.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs = πTemp001
							// line 449: self.assertEqual(s.substitute(dict(who='tim', what='ham')),
							πF.SetLineno(449)
							πTemp003 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							πTemp006 = πg.KWArgs{
								{"who", ßtim.ToObject()},
								{"what", ßham.ToObject()},
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, πTemp006); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßsubstitute, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewStr("tim likes to eat a bag of ham worth $100").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_delimiter_override.ToObject(), πTemp018); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("TestTemplate").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestTemplate.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 453: def test_main():
			πF.SetLineno(453)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "/usr/lib/python2.7/test/test_string.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 454: test_support.run_unittest(StringTest, ModuleTest, BytesAliasTest, TestTemplate)
					πF.SetLineno(454)
					πTemp001 = πF.MakeArgs(4)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStringTest); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßModuleTest); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßBytesAliasTest); πE != nil {
						continue
					}
					πTemp001[2] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTestTemplate); πE != nil {
						continue
					}
					πTemp001[3] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrun_unittest, nil); πE != nil {
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
			// line 456: if __name__ == '__main__':
			πF.SetLineno(456)
		Label1:
			// line 457: test_main()
			πF.SetLineno(457)
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
	πg.RegisterModule("test.test_string", Code)
}

