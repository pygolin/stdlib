package string_tests
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/UserList"
	// _ "github.com/pygolin/stdlib/test/test_support"
	// _ "github.com/pygolin/stdlib/string"
	// _ "github.com/pygolin/stdlib/unittest"
	// _ "github.com/pygolin/stdlib/_struct"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß0123456789 := πg.InternStr("0123456789")
		ß0123456789a := πg.InternStr("0123456789a")
		ß10 := πg.InternStr("10")
		ß123 := πg.InternStr("123")
		ß123abc456 := πg.InternStr("123abc456")
		ß42 := πg.InternStr("42")
		ß68656c6c6f20776f726c64 := πg.InternStr("68656c6c6f20776f726c64")
		ßA := πg.InternStr("A")
		ßAA := πg.InternStr("AA")
		ßABC := πg.InternStr("ABC")
		ßAbC := πg.InternStr("AbC")
		ßBadSeq1 := πg.InternStr("BadSeq1")
		ßBadSeq2 := πg.InternStr("BadSeq2")
		ßCommonTest := πg.InternStr("CommonTest")
		ßDNSSEC := πg.InternStr("DNSSEC")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßGetint := πg.InternStr("Getint")
		ßL := πg.InternStr("L")
		ßMixinStrStringUserStringTest := πg.InternStr("MixinStrStringUserStringTest")
		ßMixinStrUnicodeTest := πg.InternStr("MixinStrUnicodeTest")
		ßMixinStrUnicodeUserStringTest := πg.InternStr("MixinStrUnicodeUserStringTest")
		ßMixinStrUserStringTest := πg.InternStr("MixinStrUserStringTest")
		ßNOT := πg.InternStr("NOT")
		ßNonStringModuleTest := πg.InternStr("NonStringModuleTest")
		ßNone := πg.InternStr("None")
		ßOverflowError := πg.InternStr("OverflowError")
		ßP := πg.InternStr("P")
		ßSequence := πg.InternStr("Sequence")
		ßTestCase := πg.InternStr("TestCase")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUserList := πg.InternStr("UserList")
		ßValueError := πg.InternStr("ValueError")
		ß_UserList := πg.InternStr("_UserList")
		ß__class__ := πg.InternStr("__class__")
		ß__contains__ := πg.InternStr("__contains__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__getslice__ := πg.InternStr("__getslice__")
		ß__init__ := πg.InternStr("__init__")
		ß__len__ := πg.InternStr("__len__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__mod__ := πg.InternStr("__mod__")
		ß__module__ := πg.InternStr("__module__")
		ß__mul__ := πg.InternStr("__mul__")
		ß__name__ := πg.InternStr("__name__")
		ßa := πg.InternStr("a")
		ßa1b3c := πg.InternStr("a1b3c")
		ßaBc := πg.InternStr("aBc")
		ßaBc123 := πg.InternStr("aBc123")
		ßab := πg.InternStr("ab")
		ßabc := πg.InternStr("abc")
		ßabcabcabc := πg.InternStr("abcabcabc")
		ßabcd := πg.InternStr("abcd")
		ßac := πg.InternStr("ac")
		ßargs := πg.InternStr("args")
		ßascii_letters := πg.InternStr("ascii_letters")
		ßasd := πg.InternStr("asd")
		ßasdf := πg.InternStr("asdf")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertIs := πg.InternStr("assertIs")
		ßassertNotEqual := πg.InternStr("assertNotEqual")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertRaisesRegexp := πg.InternStr("assertRaisesRegexp")
		ßassertTrue := πg.InternStr("assertTrue")
		ßb := πg.InternStr("b")
		ßbar := πg.InternStr("bar")
		ßbase64 := πg.InternStr("base64")
		ßbc := πg.InternStr("bc")
		ßbd := πg.InternStr("bd")
		ßc := πg.InternStr("c")
		ßcalcsize := πg.InternStr("calcsize")
		ßcapitalize := πg.InternStr("capitalize")
		ßcheck_py3k_warnings := πg.InternStr("check_py3k_warnings")
		ßcheckcall := πg.InternStr("checkcall")
		ßcheckequal := πg.InternStr("checkequal")
		ßcheckraises := πg.InternStr("checkraises")
		ßchr := πg.InternStr("chr")
		ßcount := πg.InternStr("count")
		ßd := πg.InternStr("d")
		ßdecode := πg.InternStr("decode")
		ßdef := πg.InternStr("def")
		ßdict := πg.InternStr("dict")
		ßdigits := πg.InternStr("digits")
		ßell := πg.InternStr("ell")
		ßello := πg.InternStr("ello")
		ßencode := πg.InternStr("encode")
		ßendswith := πg.InternStr("endswith")
		ßexception := πg.InternStr("exception")
		ßfail := πg.InternStr("fail")
		ßfind := πg.InternStr("find")
		ßfixtype := πg.InternStr("fixtype")
		ßfloat := πg.InternStr("float")
		ßfoo := πg.InternStr("foo")
		ßgetInt := πg.InternStr("getInt")
		ßgetattr := πg.InternStr("getattr")
		ßghi := πg.InternStr("ghi")
		ßh := πg.InternStr("h")
		ßha := πg.InternStr("ha")
		ßhash := πg.InternStr("hash")
		ßhave_unicode := πg.InternStr("have_unicode")
		ßhe := πg.InternStr("he")
		ßhel := πg.InternStr("hel")
		ßhell := πg.InternStr("hell")
		ßhello := πg.InternStr("hello")
		ßhellowo := πg.InternStr("hellowo")
		ßhelloworld := πg.InternStr("helloworld")
		ßhellox := πg.InternStr("hellox")
		ßhex := πg.InternStr("hex")
		ßhttp := πg.InternStr("http")
		ßindex := πg.InternStr("index")
		ßisalnum := πg.InternStr("isalnum")
		ßisalpha := πg.InternStr("isalpha")
		ßisdigit := πg.InternStr("isdigit")
		ßisinstance := πg.InternStr("isinstance")
		ßislower := πg.InternStr("islower")
		ßisspace := πg.InternStr("isspace")
		ßistitle := πg.InternStr("istitle")
		ßisupper := πg.InternStr("isupper")
		ßiteritems := πg.InternStr("iteritems")
		ßjoin := πg.InternStr("join")
		ßl := πg.InternStr("l")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßllo := πg.InternStr("llo")
		ßlo := πg.InternStr("lo")
		ßlower := πg.InternStr("lower")
		ßlowo := πg.InternStr("lowo")
		ßlstrip := πg.InternStr("lstrip")
		ßmaketrans := πg.InternStr("maketrans")
		ßmap := πg.InternStr("map")
		ßmaxint := πg.InternStr("maxint")
		ßo := πg.InternStr("o")
		ßobject := πg.InternStr("object")
		ßorg := πg.InternStr("org")
		ßpartition := πg.InternStr("partition")
		ßreplace := πg.InternStr("replace")
		ßrfind := πg.InternStr("rfind")
		ßrindex := πg.InternStr("rindex")
		ßrld := πg.InternStr("rld")
		ßrot13 := πg.InternStr("rot13")
		ßrpartition := πg.InternStr("rpartition")
		ßrstrip := πg.InternStr("rstrip")
		ßseq := πg.InternStr("seq")
		ßskipIf := πg.InternStr("skipIf")
		ßskipUnless := πg.InternStr("skipUnless")
		ßslice := πg.InternStr("slice")
		ßsplitlines := πg.InternStr("splitlines")
		ßstartswith := πg.InternStr("startswith")
		ßstr := πg.InternStr("str")
		ßstring := πg.InternStr("string")
		ßstrip := πg.InternStr("strip")
		ßstruct := πg.InternStr("struct")
		ßswapcase := πg.InternStr("swapcase")
		ßsys := πg.InternStr("sys")
		ßtest___contains__ := πg.InternStr("test___contains__")
		ßtest_bug1001011 := πg.InternStr("test_bug1001011")
		ßtest_encoding_decoding := πg.InternStr("test_encoding_decoding")
		ßtest_endswith := πg.InternStr("test_endswith")
		ßtest_extended_getslice := πg.InternStr("test_extended_getslice")
		ßtest_find_etc_raise_correct_error_messages := πg.InternStr("test_find_etc_raise_correct_error_messages")
		ßtest_fixtype := πg.InternStr("test_fixtype")
		ßtest_floatformatting := πg.InternStr("test_floatformatting")
		ßtest_formatting := πg.InternStr("test_formatting")
		ßtest_hash := πg.InternStr("test_hash")
		ßtest_inplace_rewrites := πg.InternStr("test_inplace_rewrites")
		ßtest_isalnum := πg.InternStr("test_isalnum")
		ßtest_isalpha := πg.InternStr("test_isalpha")
		ßtest_isdigit := πg.InternStr("test_isdigit")
		ßtest_islower := πg.InternStr("test_islower")
		ßtest_isspace := πg.InternStr("test_isspace")
		ßtest_istitle := πg.InternStr("test_istitle")
		ßtest_isupper := πg.InternStr("test_isupper")
		ßtest_join := πg.InternStr("test_join")
		ßtest_maketrans := πg.InternStr("test_maketrans")
		ßtest_mul := πg.InternStr("test_mul")
		ßtest_none_arguments := πg.InternStr("test_none_arguments")
		ßtest_partition := πg.InternStr("test_partition")
		ßtest_replace_overflow := πg.InternStr("test_replace_overflow")
		ßtest_rpartition := πg.InternStr("test_rpartition")
		ßtest_slice := πg.InternStr("test_slice")
		ßtest_splitlines := πg.InternStr("test_splitlines")
		ßtest_startswith := πg.InternStr("test_startswith")
		ßtest_strip_whitespace := πg.InternStr("test_strip_whitespace")
		ßtest_subscript := πg.InternStr("test_subscript")
		ßtest_support := πg.InternStr("test_support")
		ßtest_title := πg.InternStr("test_title")
		ßti := πg.InternStr("ti")
		ßtitle := πg.InternStr("title")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßtype2test := πg.InternStr("type2test")
		ßunicode := πg.InternStr("unicode")
		ßunittest := πg.InternStr("unittest")
		ßupper := πg.InternStr("upper")
		ßuu := πg.InternStr("uu")
		ßworl := πg.InternStr("worl")
		ßworld := πg.InternStr("world")
		ßwxyz := πg.InternStr("wxyz")
		ßx := πg.InternStr("x")
		ßxrange := πg.InternStr("xrange")
		ßxyz := πg.InternStr("xyz")
		ßxyzw := πg.InternStr("xyzw")
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """
			πF.SetLineno(1)
			// line 1: """
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nCommon tests shared by test_str, test_unicode, test_userstring and test_string.\n").ToObject()); πE != nil {
				continue
			}
			// line 6: import unittest, string, sys
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "unittest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßunittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "string"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstring.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import _struct as struct
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "_struct"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstruct.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: from test import test_support
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: import UserList as _UserList
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "UserList"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_UserList.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 11: UserList = _UserList.UserList
			πF.SetLineno(11)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_UserList); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßUserList, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUserList.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 13: class Sequence(object):
			πF.SetLineno(13)
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
			_, πE = πg.NewCode("Sequence", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 14: def __init__(self, seq='wxyz'): self.seq = seq
					πF.SetLineno(14)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					// line 14: def __init__(self, seq='wxyz'): self.seq = seq
					πF.SetLineno(14)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), ßwxyz.ToObject()); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "seq", Def: ßwxyz.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseq *πg.Object = πArgs[1]; _ = µseq
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 14: def __init__(self, seq='wxyz'): self.seq = seq
							πF.SetLineno(14)
							if πE = πg.CheckLocal(πF, µseq, "seq"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µseq); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseq, πTemp001); πE != nil {
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
					// line 15: def __len__(self): return len(self.seq)
					πF.SetLineno(15)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 15: def __len__(self): return len(self.seq)
							πF.SetLineno(15)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßseq, nil); πE != nil {
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
					// line 16: def __getitem__(self, i): return self.seq[i]
					πF.SetLineno(16)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 16: def __getitem__(self, i): return self.seq[i]
							πF.SetLineno(16)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßseq, nil); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Sequence").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSequence.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 18: class BadSeq1(Sequence):
			πF.SetLineno(18)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßSequence); πE != nil {
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
			_, πE = πg.NewCode("BadSeq1", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 19: def __init__(self): self.seq = [7, 'hello', 123L]
					πF.SetLineno(19)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 19: def __init__(self): self.seq = [7, 'hello', 123L]
							πF.SetLineno(19)
							πTemp001 = make([]*πg.Object, 3)
							πTemp001[0] = πg.NewInt(7).ToObject()
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = πg.NewLongFromBytes([]byte{0x7b,}).ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseq, πTemp003); πE != nil {
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
			if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadSeq1").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBadSeq1.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 21: class BadSeq2(Sequence):
			πF.SetLineno(21)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßSequence); πE != nil {
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
			_, πE = πg.NewCode("BadSeq2", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 22: def __init__(self): self.seq = ['a', 'b', 'c']
					πF.SetLineno(22)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 22: def __init__(self): self.seq = ['a', 'b', 'c']
							πF.SetLineno(22)
							πTemp001 = make([]*πg.Object, 3)
							πTemp001[0] = ßa.ToObject()
							πTemp001[1] = ßb.ToObject()
							πTemp001[2] = ßc.ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßseq, πTemp003); πE != nil {
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
					// line 23: def __len__(self): return 8
					πF.SetLineno(23)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 23: def __len__(self): return 8
							πF.SetLineno(23)
							πR = πg.NewInt(8).ToObject()
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BadSeq2").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBadSeq2.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 25: class CommonTest(unittest.TestCase):
			πF.SetLineno(25)
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
			_, πE = πg.NewCode("CommonTest", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πTemp013 bool
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 []*πg.Object
				_ = πTemp017
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 32: type2test = None
					πF.SetLineno(32)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßtype2test.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 37: def fixtype(self, obj):
					πF.SetLineno(37)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("fixtype", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πArgs[1]; _ = µobj
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
								goto Label2
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
								goto Label3
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
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
								goto Label4
							}
							goto Label5
							// line 38: if isinstance(obj, str):
							πF.SetLineno(38)
						Label1:
							// line 39: return self.__class__.type2test(obj)
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							goto Label6
							// line 40: elif isinstance(obj, list):
							πF.SetLineno(40)
						Label2:
							// line 41: return [self.fixtype(x) for x in obj]
							πF.SetLineno(41)
							πTemp005 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/string_tests.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πg.UnboundLocal; _ = µx
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
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µobj); πE != nil {
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
											µx = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 41: return [self.fixtype(x) for x in obj]
										πF.SetLineno(41)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										πTemp005[0] = µx
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp006, nil
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
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label6
							// line 42: elif isinstance(obj, tuple):
							πF.SetLineno(42)
						Label3:
							// line 43: return tuple([self.fixtype(x) for x in obj])
							πF.SetLineno(43)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = make([]πg.Param, 0)
							πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/string_tests.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πg.UnboundLocal; _ = µx
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
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µobj); πE != nil {
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
											µx = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 43: return tuple([self.fixtype(x) for x in obj])
										πF.SetLineno(43)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										πTemp005[0] = µx
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp006, nil
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
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp007
							continue
							goto Label6
							// line 44: elif isinstance(obj, dict):
							πF.SetLineno(44)
						Label4:
							// line 45: return dict([
							πF.SetLineno(45)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = make([]πg.Param, 0)
							πTemp007 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/string_tests.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µobj, ßiteritems, nil); πE != nil {
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
										// line 45: return dict([
										πF.SetLineno(45)
										πTemp007 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
											continue
										}
										πTemp007[0] = µkey
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
										πTemp007 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										πTemp007[0] = µvalue
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
											continue
										}
										if πTemp008, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp007)
										πTemp002 = πg.NewTuple2(πTemp006, πTemp008).ToObject()
										πF.PushCheckpoint(4)
										return πTemp002, nil
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
							if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp008
							continue
							goto Label6
						Label5:
							// line 50: return obj
							πF.SetLineno(50)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πR = µobj
							continue
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
					if πE = πClass.SetItem(πF, ßfixtype.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 52: def test_fixtype(self):
					πF.SetLineno(52)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_fixtype", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
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
							// line 53: self.assertIs(type(self.fixtype("123")), self.type2test)
							πF.SetLineno(53)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ß123.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_fixtype.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 56: def checkequal(self, result, object, methodname, *args):
					πF.SetLineno(56)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					πTemp002[2] = πg.Param{Name: "object", Def: nil}
					πTemp002[3] = πg.Param{Name: "methodname", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("checkequal", "/usr/lib/python2.7/test/string_tests.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var µobject *πg.Object = πArgs[2]; _ = µobject
						var µmethodname *πg.Object = πArgs[3]; _ = µmethodname
						var µargs *πg.Object = πArgs[4]; _ = µargs
						var µrealresult *πg.Object = πg.UnboundLocal; _ = µrealresult
						var µsubtype *πg.Object = πg.UnboundLocal; _ = µsubtype
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
							// line 57: result = self.fixtype(result)
							πF.SetLineno(57)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp003
							// line 58: object = self.fixtype(object)
							πF.SetLineno(58)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							πTemp001[0] = µobject
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µobject = πTemp003
							// line 59: args = self.fixtype(args)
							πF.SetLineno(59)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µargs = πTemp003
							// line 60: realresult = getattr(object, methodname)(*args)
							πF.SetLineno(60)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							πTemp001[0] = µobject
							if πE = πg.CheckLocal(πF, µmethodname, "methodname"); πE != nil {
								continue
							}
							πTemp001[1] = µmethodname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Invoke(πF, πTemp003, nil, µargs, nil, nil); πE != nil {
								continue
							}
							µrealresult = πTemp002
							// line 61: self.assertEqual(
							πF.SetLineno(61)
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrealresult, "realresult"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µobject, µrealresult); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 67: if object == realresult:
							πF.SetLineno(67)
						Label1:
							// line 68: class subtype(self.__class__.type2test):
							πF.SetLineno(68)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßtype2test, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp007
							πTemp005 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("subtype", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp005
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 69: pass
									πF.SetLineno(69)
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
							if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("subtype").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µsubtype = πTemp006
							// line 70: object = subtype(object)
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							πTemp001[0] = µobject
							if πE = πg.CheckLocal(πF, µsubtype, "subtype"); πE != nil {
								continue
							}
							if πTemp002, πE = µsubtype.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µobject = πTemp002
							// line 71: realresult = getattr(object, methodname)(*args)
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							πTemp001[0] = µobject
							if πE = πg.CheckLocal(πF, µmethodname, "methodname"); πE != nil {
								continue
							}
							πTemp001[1] = µmethodname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Invoke(πF, πTemp003, nil, µargs, nil, nil); πE != nil {
								continue
							}
							µrealresult = πTemp002
							// line 72: self.assertTrue(object is not realresult)
							πF.SetLineno(72)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrealresult, "realresult"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µobject != µrealresult).ToObject()
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
					if πE = πClass.SetItem(πF, ßcheckequal.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 75: def checkraises(self, exc, obj, methodname, *args):
					πF.SetLineno(75)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "exc", Def: nil}
					πTemp002[2] = πg.Param{Name: "obj", Def: nil}
					πTemp002[3] = πg.Param{Name: "methodname", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("checkraises", "/usr/lib/python2.7/test/string_tests.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							default: panic("unexpected function state")
							}
							// line 76: obj = self.fixtype(obj)
							πF.SetLineno(76)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µobj = πTemp003
							// line 77: args = self.fixtype(args)
							πF.SetLineno(77)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µargs = πTemp003
							// line 78: with self.assertRaises(exc) as cm:
							πF.SetLineno(78)
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
							// line 79: getattr(obj, methodname)(*args)
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µmethodname, "methodname"); πE != nil {
								continue
							}
							πTemp001[1] = µmethodname
							if πTemp005, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.Invoke(πF, πTemp006, nil, µargs, nil, nil); πE != nil {
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
								if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πTemp010.ToObject(), πTemp008.ToObject(), πTemp009.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp005, πE = πTemp002.Call(πF, πg.Args{πTemp003, πg.None, πg.None, πg.None}, nil); πE != nil {
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
							// line 80: self.assertNotEqual(cm.exception.args[0], '')
							πF.SetLineno(80)
							πTemp001 = πF.MakeArgs(2)
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcm, "cm"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µcm, ßexception, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßargs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp012, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							πTemp001[1] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßassertNotEqual, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcheckraises.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 83: def checkcall(self, object, methodname, *args):
					πF.SetLineno(83)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "object", Def: nil}
					πTemp002[2] = πg.Param{Name: "methodname", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("checkcall", "/usr/lib/python2.7/test/string_tests.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobject *πg.Object = πArgs[1]; _ = µobject
						var µmethodname *πg.Object = πArgs[2]; _ = µmethodname
						var µargs *πg.Object = πArgs[3]; _ = µargs
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
							// line 84: object = self.fixtype(object)
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							πTemp001[0] = µobject
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µobject = πTemp003
							// line 85: args = self.fixtype(args)
							πF.SetLineno(85)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfixtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µargs = πTemp003
							// line 86: getattr(object, methodname)(*args)
							πF.SetLineno(86)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobject, "object"); πE != nil {
								continue
							}
							πTemp001[0] = µobject
							if πE = πg.CheckLocal(πF, µmethodname, "methodname"); πE != nil {
								continue
							}
							πTemp001[1] = µmethodname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Invoke(πF, πTemp003, nil, µargs, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcheckcall.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 88: def test_hash(self):
					πF.SetLineno(88)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_hash", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µc *πg.Object = πg.UnboundLocal; _ = µc
						var πTemp001 []*πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 90: a = self.type2test('DNSSEC')
							πF.SetLineno(90)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßDNSSEC.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp003
							// line 91: b = self.type2test('')
							πF.SetLineno(91)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µb = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µa); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µc = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 93: b += c
							πF.SetLineno(93)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µb, µc); πE != nil {
								continue
							}
							µb = πTemp003
							// line 94: hash(b)
							πF.SetLineno(94)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp001[0] = µb
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 95: self.assertEqual(hash(a), hash(b))
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp007[0] = µa
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp007[0] = µb
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_hash.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 538: def test_strip_whitespace(self):
					πF.SetLineno(538)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_strip_whitespace", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µb *πg.Object = πg.UnboundLocal; _ = µb
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
							// line 539: self.checkequal('hello', '   hello   ', 'strip')
							πF.SetLineno(539)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßhello.ToObject()
							πTemp001[1] = πg.NewStr("   hello   ").ToObject()
							πTemp001[2] = ßstrip.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 540: self.checkequal('hello   ', '   hello   ', 'lstrip')
							πF.SetLineno(540)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("hello   ").ToObject()
							πTemp001[1] = πg.NewStr("   hello   ").ToObject()
							πTemp001[2] = ßlstrip.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 541: self.checkequal('   hello', '   hello   ', 'rstrip')
							πF.SetLineno(541)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("   hello").ToObject()
							πTemp001[1] = πg.NewStr("   hello   ").ToObject()
							πTemp001[2] = ßrstrip.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 542: self.checkequal('hello', 'hello', 'strip')
							πF.SetLineno(542)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßhello.ToObject()
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstrip.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 544: b = ' \t\n\r\f\vabc \t\n\r\f\v'
							πF.SetLineno(544)
							µb = πg.NewStr(" \t\n\r\x0c\x0babc \t\n\r\x0c\x0b").ToObject()
							// line 545: self.checkequal('abc', b, 'strip')
							πF.SetLineno(545)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßabc.ToObject()
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp001[1] = µb
							πTemp001[2] = ßstrip.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 546: self.checkequal('abc \t\n\r\f\v', b, 'lstrip')
							πF.SetLineno(546)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("abc \t\n\r\x0c\x0b").ToObject()
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp001[1] = µb
							πTemp001[2] = ßlstrip.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 547: self.checkequal(' \t\n\r\f\vabc', b, 'rstrip')
							πF.SetLineno(547)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr(" \t\n\r\x0c\x0babc").ToObject()
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp001[1] = µb
							πTemp001[2] = ßrstrip.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 550: self.checkequal('hello', '   hello   ', 'strip', None)
							πF.SetLineno(550)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßhello.ToObject()
							πTemp001[1] = πg.NewStr("   hello   ").ToObject()
							πTemp001[2] = ßstrip.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 551: self.checkequal('hello   ', '   hello   ', 'lstrip', None)
							πF.SetLineno(551)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("hello   ").ToObject()
							πTemp001[1] = πg.NewStr("   hello   ").ToObject()
							πTemp001[2] = ßlstrip.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 552: self.checkequal('   hello', '   hello   ', 'rstrip', None)
							πF.SetLineno(552)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("   hello").ToObject()
							πTemp001[1] = πg.NewStr("   hello   ").ToObject()
							πTemp001[2] = ßrstrip.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 553: self.checkequal('hello', 'hello', 'strip', None)
							πF.SetLineno(553)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßhello.ToObject()
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstrip.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_strip_whitespace.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 806: 'only applies to 32-bit platforms')
					πF.SetLineno(806)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_replace_overflow", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µA2_16 *πg.Object = πg.UnboundLocal; _ = µA2_16
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
							// line 809: A2_16 = "A" * (2**16)
							πF.SetLineno(809)
							if πTemp002, πE = πg.Pow(πF, πg.NewInt(2).ToObject(), πg.NewInt(16).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, ßA.ToObject(), πTemp002); πE != nil {
								continue
							}
							µA2_16 = πTemp001
							// line 810: self.checkraises(OverflowError, A2_16, "replace", "", A2_16)
							πF.SetLineno(810)
							πTemp003 = πF.MakeArgs(5)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µA2_16, "A2_16"); πE != nil {
								continue
							}
							πTemp003[1] = µA2_16
							πTemp003[2] = ßreplace.ToObject()
							πTemp003[3] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µA2_16, "A2_16"); πE != nil {
								continue
							}
							πTemp003[4] = µA2_16
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 811: self.checkraises(OverflowError, A2_16, "replace", "A", A2_16)
							πF.SetLineno(811)
							πTemp003 = πF.MakeArgs(5)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µA2_16, "A2_16"); πE != nil {
								continue
							}
							πTemp003[1] = µA2_16
							πTemp003[2] = ßreplace.ToObject()
							πTemp003[3] = ßA.ToObject()
							if πE = πg.CheckLocal(πF, µA2_16, "A2_16"); πE != nil {
								continue
							}
							πTemp003[4] = µA2_16
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 812: self.checkraises(OverflowError, A2_16, "replace", "AA", A2_16+A2_16)
							πF.SetLineno(812)
							πTemp003 = πF.MakeArgs(5)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µA2_16, "A2_16"); πE != nil {
								continue
							}
							πTemp003[1] = µA2_16
							πTemp003[2] = ßreplace.ToObject()
							πTemp003[3] = ßAA.ToObject()
							if πE = πg.CheckLocal(πF, µA2_16, "A2_16"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µA2_16, "A2_16"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µA2_16, µA2_16); πE != nil {
								continue
							}
							πTemp003[4] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_replace_overflow.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 805: @unittest.skipIf(sys.maxint > (1 << 32) or struct.calcsize('P') != 4,
					πF.SetLineno(805)
					πTemp010 = πF.MakeArgs(1)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtest_replace_overflow); πE != nil {
						continue
					}
					πTemp010[0] = πTemp011
					πTemp012 = πF.MakeArgs(2)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßmaxint, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GT(πF, πTemp016, πTemp015); πE != nil {
						continue
					}
					πTemp011 = πTemp014
					if πTemp013, πE = πg.IsTrue(πF, πTemp011); πE != nil {
						continue
					}
					if πTemp013 {
						goto Label1
					}
					πTemp017 = πF.MakeArgs(1)
					πTemp017[0] = ßP.ToObject()
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßcalcsize, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp016.Call(πF, πTemp017, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp017)
					if πTemp014, πE = πg.NE(πF, πTemp015, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp011 = πTemp014
				Label1:
					πTemp012[0] = πTemp011
					πTemp012[1] = πg.NewStr("only applies to 32-bit platforms").ToObject()
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp011, ßskipIf, nil); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp014.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					if πTemp014, πE = πTemp011.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßtest_replace_overflow.ToObject(), πTemp014); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("CommonTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßCommonTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 831: class NonStringModuleTest(object):
			πF.SetLineno(831)
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
			_, πE = πg.NewCode("NonStringModuleTest", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 835: def test_islower(self):
					πF.SetLineno(835)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_islower", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 836: self.checkequal(False, '', 'islower')
							πF.SetLineno(836)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 837: self.checkequal(True, 'a', 'islower')
							πF.SetLineno(837)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 838: self.checkequal(False, 'A', 'islower')
							πF.SetLineno(838)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßA.ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 839: self.checkequal(False, '\n', 'islower')
							πF.SetLineno(839)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\n").ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 840: self.checkequal(True, 'abc', 'islower')
							πF.SetLineno(840)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 841: self.checkequal(False, 'aBc', 'islower')
							πF.SetLineno(841)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßaBc.ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 842: self.checkequal(True, 'abc\n', 'islower')
							πF.SetLineno(842)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("abc\n").ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 843: self.checkraises(TypeError, 'abc', 'islower', 42)
							πF.SetLineno(843)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßislower.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_islower.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 845: def test_isupper(self):
					πF.SetLineno(845)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_isupper", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 846: self.checkequal(False, '', 'isupper')
							πF.SetLineno(846)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßisupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 847: self.checkequal(False, 'a', 'isupper')
							πF.SetLineno(847)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßisupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 848: self.checkequal(True, 'A', 'isupper')
							πF.SetLineno(848)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßA.ToObject()
							πTemp001[2] = ßisupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 849: self.checkequal(False, '\n', 'isupper')
							πF.SetLineno(849)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\n").ToObject()
							πTemp001[2] = ßisupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 850: self.checkequal(True, 'ABC', 'isupper')
							πF.SetLineno(850)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßABC.ToObject()
							πTemp001[2] = ßisupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 851: self.checkequal(False, 'AbC', 'isupper')
							πF.SetLineno(851)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßAbC.ToObject()
							πTemp001[2] = ßisupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 852: self.checkequal(True, 'ABC\n', 'isupper')
							πF.SetLineno(852)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("ABC\n").ToObject()
							πTemp001[2] = ßisupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 853: self.checkraises(TypeError, 'abc', 'isupper', 42)
							πF.SetLineno(853)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßisupper.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_isupper.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 855: def test_istitle(self):
					πF.SetLineno(855)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_istitle", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 856: self.checkequal(False, '', 'istitle')
							πF.SetLineno(856)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 857: self.checkequal(False, 'a', 'istitle')
							πF.SetLineno(857)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 858: self.checkequal(True, 'A', 'istitle')
							πF.SetLineno(858)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßA.ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 859: self.checkequal(False, '\n', 'istitle')
							πF.SetLineno(859)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\n").ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 860: self.checkequal(True, 'A Titlecased Line', 'istitle')
							πF.SetLineno(860)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("A Titlecased Line").ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 861: self.checkequal(True, 'A\nTitlecased Line', 'istitle')
							πF.SetLineno(861)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("A\nTitlecased Line").ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 862: self.checkequal(True, 'A Titlecased, Line', 'istitle')
							πF.SetLineno(862)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("A Titlecased, Line").ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 863: self.checkequal(False, 'Not a capitalized String', 'istitle')
							πF.SetLineno(863)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("Not a capitalized String").ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 864: self.checkequal(False, 'Not\ta Titlecase String', 'istitle')
							πF.SetLineno(864)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("Not\ta Titlecase String").ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 865: self.checkequal(False, 'Not--a Titlecase String', 'istitle')
							πF.SetLineno(865)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("Not--a Titlecase String").ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 866: self.checkequal(False, 'NOT', 'istitle')
							πF.SetLineno(866)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßNOT.ToObject()
							πTemp001[2] = ßistitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 867: self.checkraises(TypeError, 'abc', 'istitle', 42)
							πF.SetLineno(867)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßistitle.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_istitle.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 869: def test_isspace(self):
					πF.SetLineno(869)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_isspace", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 870: self.checkequal(False, '', 'isspace')
							πF.SetLineno(870)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßisspace.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 871: self.checkequal(False, 'a', 'isspace')
							πF.SetLineno(871)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßisspace.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 872: self.checkequal(True, ' ', 'isspace')
							πF.SetLineno(872)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr(" ").ToObject()
							πTemp001[2] = ßisspace.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 873: self.checkequal(True, '\t', 'isspace')
							πF.SetLineno(873)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\t").ToObject()
							πTemp001[2] = ßisspace.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 874: self.checkequal(True, '\r', 'isspace')
							πF.SetLineno(874)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\r").ToObject()
							πTemp001[2] = ßisspace.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 875: self.checkequal(True, '\n', 'isspace')
							πF.SetLineno(875)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\n").ToObject()
							πTemp001[2] = ßisspace.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 876: self.checkequal(True, ' \t\r\n', 'isspace')
							πF.SetLineno(876)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr(" \t\r\n").ToObject()
							πTemp001[2] = ßisspace.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 877: self.checkequal(False, ' \t\r\na', 'isspace')
							πF.SetLineno(877)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr(" \t\r\na").ToObject()
							πTemp001[2] = ßisspace.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 878: self.checkraises(TypeError, 'abc', 'isspace', 42)
							πF.SetLineno(878)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßisspace.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_isspace.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 880: def test_isalpha(self):
					πF.SetLineno(880)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_isalpha", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 881: self.checkequal(False, '', 'isalpha')
							πF.SetLineno(881)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßisalpha.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 882: self.checkequal(True, 'a', 'isalpha')
							πF.SetLineno(882)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßisalpha.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 883: self.checkequal(True, 'A', 'isalpha')
							πF.SetLineno(883)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßA.ToObject()
							πTemp001[2] = ßisalpha.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 884: self.checkequal(False, '\n', 'isalpha')
							πF.SetLineno(884)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\n").ToObject()
							πTemp001[2] = ßisalpha.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 885: self.checkequal(True, 'abc', 'isalpha')
							πF.SetLineno(885)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßisalpha.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 886: self.checkequal(False, 'aBc123', 'isalpha')
							πF.SetLineno(886)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßaBc123.ToObject()
							πTemp001[2] = ßisalpha.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 887: self.checkequal(False, 'abc\n', 'isalpha')
							πF.SetLineno(887)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("abc\n").ToObject()
							πTemp001[2] = ßisalpha.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 888: self.checkraises(TypeError, 'abc', 'isalpha', 42)
							πF.SetLineno(888)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßisalpha.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_isalpha.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 890: def test_isalnum(self):
					πF.SetLineno(890)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_isalnum", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 891: self.checkequal(False, '', 'isalnum')
							πF.SetLineno(891)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßisalnum.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 892: self.checkequal(True, 'a', 'isalnum')
							πF.SetLineno(892)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßisalnum.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 893: self.checkequal(True, 'A', 'isalnum')
							πF.SetLineno(893)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßA.ToObject()
							πTemp001[2] = ßisalnum.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 894: self.checkequal(False, '\n', 'isalnum')
							πF.SetLineno(894)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\n").ToObject()
							πTemp001[2] = ßisalnum.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 895: self.checkequal(True, '123abc456', 'isalnum')
							πF.SetLineno(895)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß123abc456.ToObject()
							πTemp001[2] = ßisalnum.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 896: self.checkequal(True, 'a1b3c', 'isalnum')
							πF.SetLineno(896)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa1b3c.ToObject()
							πTemp001[2] = ßisalnum.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 897: self.checkequal(False, 'aBc000 ', 'isalnum')
							πF.SetLineno(897)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("aBc000 ").ToObject()
							πTemp001[2] = ßisalnum.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 898: self.checkequal(False, 'abc\n', 'isalnum')
							πF.SetLineno(898)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("abc\n").ToObject()
							πTemp001[2] = ßisalnum.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 899: self.checkraises(TypeError, 'abc', 'isalnum', 42)
							πF.SetLineno(899)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßisalnum.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_isalnum.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 901: def test_isdigit(self):
					πF.SetLineno(901)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_isdigit", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 902: self.checkequal(False, '', 'isdigit')
							πF.SetLineno(902)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßisdigit.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 903: self.checkequal(False, 'a', 'isdigit')
							πF.SetLineno(903)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßisdigit.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 904: self.checkequal(True, '0', 'isdigit')
							πF.SetLineno(904)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß0.ToObject()
							πTemp001[2] = ßisdigit.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 905: self.checkequal(True, '0123456789', 'isdigit')
							πF.SetLineno(905)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß0123456789.ToObject()
							πTemp001[2] = ßisdigit.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 906: self.checkequal(False, '0123456789a', 'isdigit')
							πF.SetLineno(906)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß0123456789a.ToObject()
							πTemp001[2] = ßisdigit.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 908: self.checkraises(TypeError, 'abc', 'isdigit', 42)
							πF.SetLineno(908)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßisdigit.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_isdigit.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 910: def test_title(self):
					πF.SetLineno(910)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_title", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 911: self.checkequal(' Hello ', ' hello ', 'title')
							πF.SetLineno(911)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr(" Hello ").ToObject()
							πTemp001[1] = πg.NewStr(" hello ").ToObject()
							πTemp001[2] = ßtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 912: self.checkequal('Hello ', 'hello ', 'title')
							πF.SetLineno(912)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("Hello ").ToObject()
							πTemp001[1] = πg.NewStr("hello ").ToObject()
							πTemp001[2] = ßtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 913: self.checkequal('Hello ', 'Hello ', 'title')
							πF.SetLineno(913)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("Hello ").ToObject()
							πTemp001[1] = πg.NewStr("Hello ").ToObject()
							πTemp001[2] = ßtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 914: self.checkequal('Format This As Title String', "fOrMaT thIs aS titLe String", 'title')
							πF.SetLineno(914)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("Format This As Title String").ToObject()
							πTemp001[1] = πg.NewStr("fOrMaT thIs aS titLe String").ToObject()
							πTemp001[2] = ßtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 915: self.checkequal('Format,This-As*Title;String', "fOrMaT,thIs-aS*titLe;String", 'title', )
							πF.SetLineno(915)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("Format,This-As*Title;String").ToObject()
							πTemp001[1] = πg.NewStr("fOrMaT,thIs-aS*titLe;String").ToObject()
							πTemp001[2] = ßtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 916: self.checkequal('Getint', "getInt", 'title')
							πF.SetLineno(916)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßGetint.ToObject()
							πTemp001[1] = ßgetInt.ToObject()
							πTemp001[2] = ßtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 917: self.checkraises(TypeError, 'hello', 'title', 42)
							πF.SetLineno(917)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßtitle.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_title.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 919: def test_splitlines(self):
					πF.SetLineno(919)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_splitlines", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 920: self.checkequal(['abc', 'def', '', 'ghi'], "abc\ndef\n\rghi", 'splitlines')
							πF.SetLineno(920)
							πTemp001 = πF.MakeArgs(3)
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = ßabc.ToObject()
							πTemp002[1] = ßdef.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp002[3] = ßghi.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("abc\ndef\n\rghi").ToObject()
							πTemp001[2] = ßsplitlines.ToObject()
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
							// line 921: self.checkequal(['abc', 'def', '', 'ghi'], "abc\ndef\n\r\nghi", 'splitlines')
							πF.SetLineno(921)
							πTemp001 = πF.MakeArgs(3)
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = ßabc.ToObject()
							πTemp002[1] = ßdef.ToObject()
							πTemp002[2] = ß.ToObject()
							πTemp002[3] = ßghi.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("abc\ndef\n\r\nghi").ToObject()
							πTemp001[2] = ßsplitlines.ToObject()
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
							// line 922: self.checkequal(['abc', 'def', 'ghi'], "abc\ndef\r\nghi", 'splitlines')
							πF.SetLineno(922)
							πTemp001 = πF.MakeArgs(3)
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = ßabc.ToObject()
							πTemp002[1] = ßdef.ToObject()
							πTemp002[2] = ßghi.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("abc\ndef\r\nghi").ToObject()
							πTemp001[2] = ßsplitlines.ToObject()
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
							// line 923: self.checkequal(['abc', 'def', 'ghi'], "abc\ndef\r\nghi\n", 'splitlines')
							πF.SetLineno(923)
							πTemp001 = πF.MakeArgs(3)
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = ßabc.ToObject()
							πTemp002[1] = ßdef.ToObject()
							πTemp002[2] = ßghi.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("abc\ndef\r\nghi\n").ToObject()
							πTemp001[2] = ßsplitlines.ToObject()
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
							// line 924: self.checkequal(['abc', 'def', 'ghi', ''], "abc\ndef\r\nghi\n\r", 'splitlines')
							πF.SetLineno(924)
							πTemp001 = πF.MakeArgs(3)
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = ßabc.ToObject()
							πTemp002[1] = ßdef.ToObject()
							πTemp002[2] = ßghi.ToObject()
							πTemp002[3] = ß.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("abc\ndef\r\nghi\n\r").ToObject()
							πTemp001[2] = ßsplitlines.ToObject()
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
							// line 925: self.checkequal(['', 'abc', 'def', 'ghi', ''], "\nabc\ndef\r\nghi\n\r", 'splitlines')
							πF.SetLineno(925)
							πTemp001 = πF.MakeArgs(3)
							πTemp002 = make([]*πg.Object, 5)
							πTemp002[0] = ß.ToObject()
							πTemp002[1] = ßabc.ToObject()
							πTemp002[2] = ßdef.ToObject()
							πTemp002[3] = ßghi.ToObject()
							πTemp002[4] = ß.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("\nabc\ndef\r\nghi\n\r").ToObject()
							πTemp001[2] = ßsplitlines.ToObject()
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
							// line 926: self.checkequal(['\n', 'abc\n', 'def\r\n', 'ghi\n', '\r'], "\nabc\ndef\r\nghi\n\r", 'splitlines', 1)
							πF.SetLineno(926)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = make([]*πg.Object, 5)
							πTemp002[0] = πg.NewStr("\n").ToObject()
							πTemp002[1] = πg.NewStr("abc\n").ToObject()
							πTemp002[2] = πg.NewStr("def\r\n").ToObject()
							πTemp002[3] = πg.NewStr("ghi\n").ToObject()
							πTemp002[4] = πg.NewStr("\r").ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("\nabc\ndef\r\nghi\n\r").ToObject()
							πTemp001[2] = ßsplitlines.ToObject()
							πTemp001[3] = πg.NewInt(1).ToObject()
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
							// line 928: self.checkraises(TypeError, 'abc', 'splitlines', 42, 42)
							πF.SetLineno(928)
							πTemp001 = πF.MakeArgs(5)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ßsplitlines.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							πTemp001[4] = πg.NewInt(42).ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_splitlines.ToObject(), πTemp010); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("NonStringModuleTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNonStringModuleTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 931: class MixinStrUnicodeUserStringTest(NonStringModuleTest):
			πF.SetLineno(931)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNonStringModuleTest); πE != nil {
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
			_, πE = πg.NewCode("MixinStrUnicodeUserStringTest", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 936: def test_startswith(self):
					πF.SetLineno(936)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_startswith", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 937: self.checkequal(True, 'hello', 'startswith', 'he')
							πF.SetLineno(937)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßhe.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 938: self.checkequal(True, 'hello', 'startswith', 'hello')
							πF.SetLineno(938)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßhello.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 939: self.checkequal(False, 'hello', 'startswith', 'hello world')
							πF.SetLineno(939)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = πg.NewStr("hello world").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 940: self.checkequal(True, 'hello', 'startswith', '')
							πF.SetLineno(940)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 941: self.checkequal(False, 'hello', 'startswith', 'ello')
							πF.SetLineno(941)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßello.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 942: self.checkequal(True, 'hello', 'startswith', 'ello', 1)
							πF.SetLineno(942)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßello.ToObject()
							πTemp001[4] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 943: self.checkequal(True, 'hello', 'startswith', 'o', 4)
							πF.SetLineno(943)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßo.ToObject()
							πTemp001[4] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 944: self.checkequal(False, 'hello', 'startswith', 'o', 5)
							πF.SetLineno(944)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßo.ToObject()
							πTemp001[4] = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 945: self.checkequal(True, 'hello', 'startswith', '', 5)
							πF.SetLineno(945)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ß.ToObject()
							πTemp001[4] = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 946: self.checkequal(False, 'hello', 'startswith', 'lo', 6)
							πF.SetLineno(946)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßlo.ToObject()
							πTemp001[4] = πg.NewInt(6).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 947: self.checkequal(True, 'helloworld', 'startswith', 'lowo', 3)
							πF.SetLineno(947)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 948: self.checkequal(True, 'helloworld', 'startswith', 'lowo', 3, 7)
							πF.SetLineno(948)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							πTemp001[5] = πg.NewInt(7).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 949: self.checkequal(False, 'helloworld', 'startswith', 'lowo', 3, 6)
							πF.SetLineno(949)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							πTemp001[5] = πg.NewInt(6).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 952: self.checkequal(True, 'hello', 'startswith', 'he', 0, -1)
							πF.SetLineno(952)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßhe.ToObject()
							πTemp001[4] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 953: self.checkequal(True, 'hello', 'startswith', 'he', -53, -1)
							πF.SetLineno(953)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßhe.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(53).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 954: self.checkequal(False, 'hello', 'startswith', 'hello', 0, -1)
							πF.SetLineno(954)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßhello.ToObject()
							πTemp001[4] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 955: self.checkequal(False, 'hello', 'startswith', 'hello world', -1, -10)
							πF.SetLineno(955)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = πg.NewStr("hello world").ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 956: self.checkequal(False, 'hello', 'startswith', 'ello', -5)
							πF.SetLineno(956)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßello.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 957: self.checkequal(True, 'hello', 'startswith', 'ello', -4)
							πF.SetLineno(957)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßello.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 958: self.checkequal(False, 'hello', 'startswith', 'o', -2)
							πF.SetLineno(958)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßo.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 959: self.checkequal(True, 'hello', 'startswith', 'o', -1)
							πF.SetLineno(959)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßo.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 960: self.checkequal(True, 'hello', 'startswith', '', -3, -3)
							πF.SetLineno(960)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ß.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 961: self.checkequal(False, 'hello', 'startswith', 'lo', -9)
							πF.SetLineno(961)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßlo.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(9).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 963: self.checkraises(TypeError, 'hello', 'startswith')
							πF.SetLineno(963)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 964: self.checkraises(TypeError, 'hello', 'startswith', 42)
							πF.SetLineno(964)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 967: self.checkequal(True, 'hello', 'startswith', ('he', 'ha'))
							πF.SetLineno(967)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple2(ßhe.ToObject(), ßha.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 968: self.checkequal(False, 'hello', 'startswith', ('lo', 'llo'))
							πF.SetLineno(968)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple2(ßlo.ToObject(), ßllo.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 969: self.checkequal(True, 'hello', 'startswith', ('hellox', 'hello'))
							πF.SetLineno(969)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple2(ßhellox.ToObject(), ßhello.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 970: self.checkequal(False, 'hello', 'startswith', ())
							πF.SetLineno(970)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple0().ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 971: self.checkequal(True, 'helloworld', 'startswith', ('hellowo',
							πF.SetLineno(971)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple3(ßhellowo.ToObject(), ßrld.ToObject(), ßlowo.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 973: self.checkequal(False, 'helloworld', 'startswith', ('hellowo', 'ello',
							πF.SetLineno(973)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple3(ßhellowo.ToObject(), ßello.ToObject(), ßrld.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 975: self.checkequal(True, 'hello', 'startswith', ('lo', 'he'), 0, -1)
							πF.SetLineno(975)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple2(ßlo.ToObject(), ßhe.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 976: self.checkequal(False, 'hello', 'startswith', ('he', 'hel'), 0, 1)
							πF.SetLineno(976)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple2(ßhe.ToObject(), ßhel.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(0).ToObject()
							πTemp001[5] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 977: self.checkequal(True, 'hello', 'startswith', ('he', 'hel'), 0, 2)
							πF.SetLineno(977)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple2(ßhe.ToObject(), ßhel.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(0).ToObject()
							πTemp001[5] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 979: self.checkraises(TypeError, 'hello', 'startswith', (42,))
							πF.SetLineno(979)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßstartswith.ToObject()
							πTemp002 = πg.NewTuple1(πg.NewInt(42).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_startswith.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 981: def test_endswith(self):
					πF.SetLineno(981)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_endswith", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 982: self.checkequal(True, 'hello', 'endswith', 'lo')
							πF.SetLineno(982)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlo.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 983: self.checkequal(False, 'hello', 'endswith', 'he')
							πF.SetLineno(983)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßhe.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 984: self.checkequal(True, 'hello', 'endswith', '')
							πF.SetLineno(984)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 985: self.checkequal(False, 'hello', 'endswith', 'hello world')
							πF.SetLineno(985)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = πg.NewStr("hello world").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 986: self.checkequal(False, 'helloworld', 'endswith', 'worl')
							πF.SetLineno(986)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßworl.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 987: self.checkequal(True, 'helloworld', 'endswith', 'worl', 3, 9)
							πF.SetLineno(987)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßworl.ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							πTemp001[5] = πg.NewInt(9).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 988: self.checkequal(True, 'helloworld', 'endswith', 'world', 3, 12)
							πF.SetLineno(988)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßworld.ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							πTemp001[5] = πg.NewInt(12).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 989: self.checkequal(True, 'helloworld', 'endswith', 'lowo', 1, 7)
							πF.SetLineno(989)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							πTemp001[4] = πg.NewInt(1).ToObject()
							πTemp001[5] = πg.NewInt(7).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 990: self.checkequal(True, 'helloworld', 'endswith', 'lowo', 2, 7)
							πF.SetLineno(990)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							πTemp001[4] = πg.NewInt(2).ToObject()
							πTemp001[5] = πg.NewInt(7).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 991: self.checkequal(True, 'helloworld', 'endswith', 'lowo', 3, 7)
							πF.SetLineno(991)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							πTemp001[5] = πg.NewInt(7).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 992: self.checkequal(False, 'helloworld', 'endswith', 'lowo', 4, 7)
							πF.SetLineno(992)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							πTemp001[4] = πg.NewInt(4).ToObject()
							πTemp001[5] = πg.NewInt(7).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 993: self.checkequal(False, 'helloworld', 'endswith', 'lowo', 3, 8)
							πF.SetLineno(993)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							πTemp001[5] = πg.NewInt(8).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 994: self.checkequal(False, 'ab', 'endswith', 'ab', 0, 1)
							πF.SetLineno(994)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßab.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßab.ToObject()
							πTemp001[4] = πg.NewInt(0).ToObject()
							πTemp001[5] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 995: self.checkequal(False, 'ab', 'endswith', 'ab', 0, 0)
							πF.SetLineno(995)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßab.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßab.ToObject()
							πTemp001[4] = πg.NewInt(0).ToObject()
							πTemp001[5] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 998: self.checkequal(True, 'hello', 'endswith', 'lo', -2)
							πF.SetLineno(998)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlo.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 999: self.checkequal(False, 'hello', 'endswith', 'he', -2)
							πF.SetLineno(999)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßhe.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1000: self.checkequal(True, 'hello', 'endswith', '', -3, -3)
							πF.SetLineno(1000)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ß.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1001: self.checkequal(False, 'hello', 'endswith', 'hello world', -10, -2)
							πF.SetLineno(1001)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = πg.NewStr("hello world").ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(10).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1002: self.checkequal(False, 'helloworld', 'endswith', 'worl', -6)
							πF.SetLineno(1002)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßworl.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1003: self.checkequal(True, 'helloworld', 'endswith', 'worl', -5, -1)
							πF.SetLineno(1003)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßworl.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1004: self.checkequal(True, 'helloworld', 'endswith', 'worl', -5, 9)
							πF.SetLineno(1004)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßworl.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							πTemp001[5] = πg.NewInt(9).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1005: self.checkequal(True, 'helloworld', 'endswith', 'world', -7, 12)
							πF.SetLineno(1005)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßworld.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							πTemp001[5] = πg.NewInt(12).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1006: self.checkequal(True, 'helloworld', 'endswith', 'lowo', -99, -3)
							πF.SetLineno(1006)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(99).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1007: self.checkequal(True, 'helloworld', 'endswith', 'lowo', -8, -3)
							πF.SetLineno(1007)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1008: self.checkequal(True, 'helloworld', 'endswith', 'lowo', -7, -3)
							πF.SetLineno(1008)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(7).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1009: self.checkequal(False, 'helloworld', 'endswith', 'lowo', 3, -4)
							πF.SetLineno(1009)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1010: self.checkequal(False, 'helloworld', 'endswith', 'lowo', -8, -2)
							πF.SetLineno(1010)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlowo.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1012: self.checkraises(TypeError, 'hello', 'endswith')
							πF.SetLineno(1012)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1013: self.checkraises(TypeError, 'hello', 'endswith', 42)
							πF.SetLineno(1013)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1016: self.checkequal(False, 'hello', 'endswith', ('he', 'ha'))
							πF.SetLineno(1016)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple2(ßhe.ToObject(), ßha.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1017: self.checkequal(True, 'hello', 'endswith', ('lo', 'llo'))
							πF.SetLineno(1017)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple2(ßlo.ToObject(), ßllo.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1018: self.checkequal(True, 'hello', 'endswith', ('hellox', 'hello'))
							πF.SetLineno(1018)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple2(ßhellox.ToObject(), ßhello.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1019: self.checkequal(False, 'hello', 'endswith', ())
							πF.SetLineno(1019)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple0().ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1020: self.checkequal(True, 'helloworld', 'endswith', ('hellowo',
							πF.SetLineno(1020)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple3(ßhellowo.ToObject(), ßrld.ToObject(), ßlowo.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1022: self.checkequal(False, 'helloworld', 'endswith', ('hellowo', 'ello',
							πF.SetLineno(1022)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhelloworld.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple3(ßhellowo.ToObject(), ßello.ToObject(), ßrld.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(3).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1024: self.checkequal(True, 'hello', 'endswith', ('hell', 'ell'), 0, -1)
							πF.SetLineno(1024)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple2(ßhell.ToObject(), ßell.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1025: self.checkequal(False, 'hello', 'endswith', ('he', 'hel'), 0, 1)
							πF.SetLineno(1025)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple2(ßhe.ToObject(), ßhel.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(0).ToObject()
							πTemp001[5] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1026: self.checkequal(True, 'hello', 'endswith', ('he', 'hell'), 0, 4)
							πF.SetLineno(1026)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple2(ßhe.ToObject(), ßhell.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp001[4] = πg.NewInt(0).ToObject()
							πTemp001[5] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1028: self.checkraises(TypeError, 'hello', 'endswith', (42,))
							πF.SetLineno(1028)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßhello.ToObject()
							πTemp001[2] = ßendswith.ToObject()
							πTemp002 = πg.NewTuple1(πg.NewInt(42).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_endswith.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1030: def test___contains__(self):
					πF.SetLineno(1030)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test___contains__", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1031: self.checkequal(True, '', '__contains__', '')
							πF.SetLineno(1031)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ß__contains__.ToObject()
							πTemp001[3] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1032: self.checkequal(True, 'abc', '__contains__', '')
							πF.SetLineno(1032)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__contains__.ToObject()
							πTemp001[3] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1033: self.checkequal(False, 'abc', '__contains__', '\0')
							πF.SetLineno(1033)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__contains__.ToObject()
							πTemp001[3] = πg.NewStr("\x00").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1034: self.checkequal(True, '\0abc', '__contains__', '\0')
							πF.SetLineno(1034)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\x00abc").ToObject()
							πTemp001[2] = ß__contains__.ToObject()
							πTemp001[3] = πg.NewStr("\x00").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1035: self.checkequal(True, 'abc\0', '__contains__', '\0')
							πF.SetLineno(1035)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("abc\x00").ToObject()
							πTemp001[2] = ß__contains__.ToObject()
							πTemp001[3] = πg.NewStr("\x00").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1036: self.checkequal(True, '\0abc', '__contains__', 'a')
							πF.SetLineno(1036)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\x00abc").ToObject()
							πTemp001[2] = ß__contains__.ToObject()
							πTemp001[3] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1037: self.checkequal(True, 'asdf', '__contains__', 'asdf')
							πF.SetLineno(1037)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßasdf.ToObject()
							πTemp001[2] = ß__contains__.ToObject()
							πTemp001[3] = ßasdf.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1038: self.checkequal(False, 'asd', '__contains__', 'asdf')
							πF.SetLineno(1038)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßasd.ToObject()
							πTemp001[2] = ß__contains__.ToObject()
							πTemp001[3] = ßasdf.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1039: self.checkequal(False, '', '__contains__', 'asdf')
							πF.SetLineno(1039)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ß__contains__.ToObject()
							πTemp001[3] = ßasdf.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest___contains__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1041: def test_subscript(self):
					πF.SetLineno(1041)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_subscript", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1042: self.checkequal(u'a', 'abc', '__getitem__', 0)
							πF.SetLineno(1042)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewUnicode("a").ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getitem__.ToObject()
							πTemp001[3] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1043: self.checkequal(u'c', 'abc', '__getitem__', -1)
							πF.SetLineno(1043)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewUnicode("c").ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getitem__.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1044: self.checkequal(u'a', 'abc', '__getitem__', 0L)
							πF.SetLineno(1044)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewUnicode("a").ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getitem__.ToObject()
							πTemp001[3] = πg.NewLongFromBytes([]byte{}).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1045: self.checkequal(u'abc', 'abc', '__getitem__', slice(0, 3))
							πF.SetLineno(1045)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewUnicode("abc").ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getitem__.ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							πTemp004[1] = πg.NewInt(3).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[3] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1046: self.checkequal(u'abc', 'abc', '__getitem__', slice(0, 1000))
							πF.SetLineno(1046)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewUnicode("abc").ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getitem__.ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							πTemp004[1] = πg.NewInt(1000).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[3] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1047: self.checkequal(u'a', 'abc', '__getitem__', slice(0, 1))
							πF.SetLineno(1047)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewUnicode("a").ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getitem__.ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[3] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1048: self.checkequal(u'', 'abc', '__getitem__', slice(0, 0))
							πF.SetLineno(1048)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewUnicode("").ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getitem__.ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[3] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1050: self.checkraises(TypeError, 'abc', '__getitem__', 'def')
							πF.SetLineno(1050)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getitem__.ToObject()
							πTemp001[3] = ßdef.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_subscript.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1052: def test_slice(self):
					πF.SetLineno(1052)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_slice", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1053: self.checkequal('abc', 'abc', '__getslice__', 0, 1000)
							πF.SetLineno(1053)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ßabc.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = πg.NewInt(0).ToObject()
							πTemp001[4] = πg.NewInt(1000).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1054: self.checkequal('abc', 'abc', '__getslice__', 0, 3)
							πF.SetLineno(1054)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ßabc.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = πg.NewInt(0).ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1055: self.checkequal('ab', 'abc', '__getslice__', 0, 2)
							πF.SetLineno(1055)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ßab.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = πg.NewInt(0).ToObject()
							πTemp001[4] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1056: self.checkequal('bc', 'abc', '__getslice__', 1, 3)
							πF.SetLineno(1056)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ßbc.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = πg.NewInt(1).ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1057: self.checkequal('b', 'abc', '__getslice__', 1, 2)
							πF.SetLineno(1057)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ßb.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = πg.NewInt(1).ToObject()
							πTemp001[4] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1058: self.checkequal('', 'abc', '__getslice__', 2, 2)
							πF.SetLineno(1058)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ß.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = πg.NewInt(2).ToObject()
							πTemp001[4] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1059: self.checkequal('', 'abc', '__getslice__', 1000, 1000)
							πF.SetLineno(1059)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ß.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = πg.NewInt(1000).ToObject()
							πTemp001[4] = πg.NewInt(1000).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1060: self.checkequal('', 'abc', '__getslice__', 2000, 1000)
							πF.SetLineno(1060)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ß.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = πg.NewInt(2000).ToObject()
							πTemp001[4] = πg.NewInt(1000).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1061: self.checkequal('', 'abc', '__getslice__', 2, 1)
							πF.SetLineno(1061)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ß.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = πg.NewInt(2).ToObject()
							πTemp001[4] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1063: self.checkraises(TypeError, 'abc', '__getslice__', 'def')
							πF.SetLineno(1063)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__getslice__.ToObject()
							πTemp001[3] = ßdef.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_slice.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1065: def test_extended_getslice(self):
					πF.SetLineno(1065)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_extended_getslice", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µindices *πg.Object = πg.UnboundLocal; _ = µindices
						var µstart *πg.Object = πg.UnboundLocal; _ = µstart
						var µstop *πg.Object = πg.UnboundLocal; _ = µstop
						var µstep *πg.Object = πg.UnboundLocal; _ = µstep
						var µL *πg.Object = πg.UnboundLocal; _ = µL
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
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
							default: panic("unexpected function state")
							}
							// line 1067: s = string.ascii_letters + string.digits
							πF.SetLineno(1067)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßascii_letters, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßdigits, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µs = πTemp001
							// line 1068: indices = (0, None, 1, 3, 41, -1, -2, -37)
							πF.SetLineno(1068)
							πTemp005 = make([]*πg.Object, 8)
							πTemp005[0] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(41).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp005[5] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp005[6] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(37).ToObject()); πE != nil {
								continue
							}
							πTemp005[7] = πTemp002
							πTemp001 = πg.NewTuple(πTemp005...).ToObject()
							µindices = πTemp001
							if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µindices); πE != nil {
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
								µstart = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µindices); πE != nil {
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
								µstop = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µindices, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, πTemp009); πE != nil {
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
							if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
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
								µstep = πTemp004
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 1073: L = list(s)[start:stop:step]
							πF.SetLineno(1073)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstep, "step"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{µstart, µstop, µstep}, nil); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp005[0] = µs
							if πTemp011, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp009, πE = πg.GetItem(πF, πTemp012, πTemp004); πE != nil {
								continue
							}
							µL = πTemp009
							// line 1074: self.checkequal(u"".join(L), s, '__getitem__',
							πF.SetLineno(1074)
							πTemp005 = πF.MakeArgs(4)
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µL, "L"); πE != nil {
								continue
							}
							πTemp013[0] = µL
							if πTemp004, πE = πg.GetAttr(πF, πg.NewUnicode("").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp005[0] = πTemp009
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp005[1] = µs
							πTemp005[2] = ß__getitem__.ToObject()
							πTemp013 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp013[0] = µstart
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							πTemp013[1] = µstop
							if πE = πg.CheckLocal(πF, µstep, "step"); πE != nil {
								continue
							}
							πTemp013[2] = µstep
							if πTemp004, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							πTemp005[3] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
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
					if πE = πClass.SetItem(πF, ßtest_extended_getslice.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1077: def test_mul(self):
					πF.SetLineno(1077)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_mul", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1078: self.checkequal('', 'abc', '__mul__', -1)
							πF.SetLineno(1078)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ß.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__mul__.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1079: self.checkequal('', 'abc', '__mul__', 0)
							πF.SetLineno(1079)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ß.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__mul__.ToObject()
							πTemp001[3] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1080: self.checkequal('abc', 'abc', '__mul__', 1)
							πF.SetLineno(1080)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßabc.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__mul__.ToObject()
							πTemp001[3] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1081: self.checkequal('abcabcabc', 'abc', '__mul__', 3)
							πF.SetLineno(1081)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßabcabcabc.ToObject()
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__mul__.ToObject()
							πTemp001[3] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1082: self.checkraises(TypeError, 'abc', '__mul__')
							πF.SetLineno(1082)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__mul__.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1083: self.checkraises(TypeError, 'abc', '__mul__', '')
							πF.SetLineno(1083)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__mul__.ToObject()
							πTemp001[3] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_mul.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1088: def test_join(self):
					πF.SetLineno(1088)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_join", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1093: self.checkequal('a b c d', ' ', 'join', ['a', 'b', 'c', 'd'])
							πF.SetLineno(1093)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("a b c d").ToObject()
							πTemp001[1] = πg.NewStr(" ").ToObject()
							πTemp001[2] = ßjoin.ToObject()
							πTemp002 = make([]*πg.Object, 4)
							πTemp002[0] = ßa.ToObject()
							πTemp002[1] = ßb.ToObject()
							πTemp002[2] = ßc.ToObject()
							πTemp002[3] = ßd.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[3] = πTemp003
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
							// line 1094: self.checkequal('abcd', '', 'join', ('a', 'b', 'c', 'd'))
							πF.SetLineno(1094)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßabcd.ToObject()
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßjoin.ToObject()
							πTemp003 = πg.NewTuple4(ßa.ToObject(), ßb.ToObject(), ßc.ToObject(), ßd.ToObject()).ToObject()
							πTemp001[3] = πTemp003
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
							// line 1095: self.checkequal('bd', '', 'join', ('', 'b', '', 'd'))
							πF.SetLineno(1095)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßbd.ToObject()
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßjoin.ToObject()
							πTemp003 = πg.NewTuple4(ß.ToObject(), ßb.ToObject(), ß.ToObject(), ßd.ToObject()).ToObject()
							πTemp001[3] = πTemp003
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
							// line 1096: self.checkequal('ac', '', 'join', ('a', '', 'c', ''))
							πF.SetLineno(1096)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßac.ToObject()
							πTemp001[1] = ß.ToObject()
							πTemp001[2] = ßjoin.ToObject()
							πTemp003 = πg.NewTuple4(ßa.ToObject(), ß.ToObject(), ßc.ToObject(), ß.ToObject()).ToObject()
							πTemp001[3] = πTemp003
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
							// line 1097: self.checkequal('w x y z', ' ', 'join', Sequence())
							πF.SetLineno(1097)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("w x y z").ToObject()
							πTemp001[1] = πg.NewStr(" ").ToObject()
							πTemp001[2] = ßjoin.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSequence); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
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
							// line 1098: self.checkequal('abc', 'a', 'join', ('abc',))
							πF.SetLineno(1098)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßabc.ToObject()
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßjoin.ToObject()
							πTemp003 = πg.NewTuple1(ßabc.ToObject()).ToObject()
							πTemp001[3] = πTemp003
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
							// line 1099: self.checkequal('z', 'a', 'join', UserList(['z']))
							πF.SetLineno(1099)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßz.ToObject()
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßjoin.ToObject()
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
							// line 1100: if test_support.have_unicode:
							πF.SetLineno(1100)
						Label1:
							// line 1101: self.checkequal(unicode('a.b.c'), unicode('.'), 'join', ['a', 'b', 'c'])
							πF.SetLineno(1101)
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
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr(".").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							πTemp001[2] = ßjoin.ToObject()
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = ßa.ToObject()
							πTemp002[1] = ßb.ToObject()
							πTemp002[2] = ßc.ToObject()
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[3] = πTemp003
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
							// line 1102: self.checkequal(unicode('a.b.c'), '.', 'join', [unicode('a'), 'b', 'c'])
							πF.SetLineno(1102)
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
							πTemp001[1] = πg.NewStr(".").ToObject()
							πTemp001[2] = ßjoin.ToObject()
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
							πTemp001[3] = πTemp003
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
							// line 1103: self.checkequal(unicode('a.b.c'), '.', 'join', ['a', unicode('b'), 'c'])
							πF.SetLineno(1103)
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
							πTemp001[1] = πg.NewStr(".").ToObject()
							πTemp001[2] = ßjoin.ToObject()
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
							πTemp001[3] = πTemp003
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
							// line 1104: self.checkequal(unicode('a.b.c'), '.', 'join', ['a', 'b', unicode('c')])
							πF.SetLineno(1104)
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
							πTemp001[1] = πg.NewStr(".").ToObject()
							πTemp001[2] = ßjoin.ToObject()
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
							πTemp001[3] = πTemp003
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
							// line 1105: self.checkraises(TypeError, '.', 'join', ['a', unicode('b'), 3])
							πF.SetLineno(1105)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr(".").ToObject()
							πTemp001[2] = ßjoin.ToObject()
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
							πTemp001[3] = πTemp003
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
							// line 1107: self.checkequal(((('a' * i) + '-') * i)[:-1], '-', 'join',
							πF.SetLineno(1107)
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
							πTemp001[1] = πg.NewStr("-").ToObject()
							πTemp001[2] = ßjoin.ToObject()
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
							πTemp001[3] = πTemp004
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
							// line 1109: self.checkequal(((('a' * i) + '-') * i)[:-1], '-', 'join',
							πF.SetLineno(1109)
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
							πTemp001[1] = πg.NewStr("-").ToObject()
							πTemp001[2] = ßjoin.ToObject()
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
							πTemp001[3] = πTemp004
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
							// line 1112: self.checkraises(TypeError, ' ', 'join', BadSeq1())
							πF.SetLineno(1112)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr(" ").ToObject()
							πTemp001[2] = ßjoin.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBadSeq1); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp004
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
							// line 1113: self.checkequal('a b c', ' ', 'join', BadSeq2())
							πF.SetLineno(1113)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("a b c").ToObject()
							πTemp001[1] = πg.NewStr(" ").ToObject()
							πTemp001[2] = ßjoin.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBadSeq2); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
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
							// line 1115: self.checkraises(TypeError, ' ', 'join')
							πF.SetLineno(1115)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr(" ").ToObject()
							πTemp001[2] = ßjoin.ToObject()
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
							// line 1116: self.checkraises(TypeError, ' ', 'join', None)
							πF.SetLineno(1116)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr(" ").ToObject()
							πTemp001[2] = ßjoin.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[3] = πTemp003
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
							// line 1117: self.checkraises(TypeError, ' ', 'join', 7)
							πF.SetLineno(1117)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr(" ").ToObject()
							πTemp001[2] = ßjoin.ToObject()
							πTemp001[3] = πg.NewInt(7).ToObject()
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
							// line 1118: self.checkraises(TypeError, ' ', 'join', Sequence([7, 'hello', 123L]))
							πF.SetLineno(1118)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr(" ").ToObject()
							πTemp001[2] = ßjoin.ToObject()
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = make([]*πg.Object, 3)
							πTemp005[0] = πg.NewInt(7).ToObject()
							πTemp005[1] = ßhello.ToObject()
							πTemp005[2] = πg.NewLongFromBytes([]byte{0x7b,}).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSequence); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1119: try:
							πF.SetLineno(1119)
							πF.PushCheckpoint(7)
							// line 1120: def f():
							πF.SetLineno(1120)
							πTemp012 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("f", "/usr/lib/python2.7/test/string_tests.py", πTemp012, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										// line 1121: yield 4 + ""
										πF.SetLineno(1121)
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
							// line 1122: self.fixtype(' ').join(f())
							πF.SetLineno(1122)
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
							// line 1127: self.fail('exception not raised')
							πF.SetLineno(1127)
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
							// line 1123: except TypeError, e:
							πF.SetLineno(1123)
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
							// line 1124: if '+' not in str(e):
							πF.SetLineno(1124)
						Label9:
							// line 1125: self.fail('join() ate exception message')
							πF.SetLineno(1125)
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
					if πE = πClass.SetItem(πF, ßtest_join.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1129: def test_formatting(self):
					πF.SetLineno(1129)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_formatting", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µordinal *πg.Object = πg.UnboundLocal; _ = µordinal
						var µlongvalue *πg.Object = πg.UnboundLocal; _ = µlongvalue
						var µslongvalue *πg.Object = πg.UnboundLocal; _ = µslongvalue
						var µX *πg.Object = πg.UnboundLocal; _ = µX
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.Dict
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
							// line 1130: self.checkequal('+hello+', '+%s+', '__mod__', 'hello')
							πF.SetLineno(1130)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("+hello+").ToObject()
							πTemp001[1] = πg.NewStr("+%s+").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = ßhello.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1131: self.checkequal('+10+', '+%d+', '__mod__', 10)
							πF.SetLineno(1131)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("+10+").ToObject()
							πTemp001[1] = πg.NewStr("+%d+").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1132: self.checkequal('a', "%c", '__mod__', "a")
							πF.SetLineno(1132)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßa.ToObject()
							πTemp001[1] = πg.NewStr("%c").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1133: self.checkequal('a', "%c", '__mod__', "a")
							πF.SetLineno(1133)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßa.ToObject()
							πTemp001[1] = πg.NewStr("%c").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1134: self.checkequal('"', "%c", '__mod__', 34)
							πF.SetLineno(1134)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("\"").ToObject()
							πTemp001[1] = πg.NewStr("%c").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewInt(34).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1135: self.checkequal('$', "%c", '__mod__', 36)
							πF.SetLineno(1135)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("$").ToObject()
							πTemp001[1] = πg.NewStr("%c").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewInt(36).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1136: self.checkequal('10', "%d", '__mod__', 10)
							πF.SetLineno(1136)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ß10.ToObject()
							πTemp001[1] = πg.NewStr("%d").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewInt(10).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1137: self.checkequal('\x7f', "%c", '__mod__', 0x7f)
							πF.SetLineno(1137)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("\x7f").ToObject()
							πTemp001[1] = πg.NewStr("%c").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewInt(127).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(100).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πg.NewInt(2097152).ToObject()).ToObject()
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
								µordinal = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1141: self.checkraises((ValueError, OverflowError), '%c', '__mod__', ordinal)
							πF.SetLineno(1141)
							πTemp001 = πF.MakeArgs(4)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
							πTemp001[0] = πTemp003
							πTemp001[1] = πg.NewStr("%c").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							if πE = πg.CheckLocal(πF, µordinal, "ordinal"); πE != nil {
								continue
							}
							πTemp001[3] = µordinal
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
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1143: longvalue = sys.maxint + 10L
							πF.SetLineno(1143)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmaxint, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewLongFromBytes([]byte{0xa,}).ToObject()); πE != nil {
								continue
							}
							µlongvalue = πTemp002
							// line 1144: slongvalue = str(longvalue)
							πF.SetLineno(1144)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlongvalue, "longvalue"); πE != nil {
								continue
							}
							πTemp001[0] = µlongvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µslongvalue = πTemp003
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µslongvalue, "slongvalue"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µslongvalue, πTemp003); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(ßL.ToObject(), ßl.ToObject()).ToObject()
							if πTemp005, πE = πg.Contains(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 1145: if slongvalue[-1] in ("L","l"): slongvalue = slongvalue[:-1]
							πF.SetLineno(1145)
						Label4:
							// line 1145: if slongvalue[-1] in ("L","l"): slongvalue = slongvalue[:-1]
							πF.SetLineno(1145)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µslongvalue, "slongvalue"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µslongvalue, πTemp002); πE != nil {
								continue
							}
							µslongvalue = πTemp003
							goto Label5
						Label5:
							// line 1146: self.checkequal(' 42', '%3ld', '__mod__', 42)
							πF.SetLineno(1146)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr(" 42").ToObject()
							πTemp001[1] = πg.NewStr("%3ld").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1147: self.checkequal('42', '%d', '__mod__', 42L)
							πF.SetLineno(1147)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ß42.ToObject()
							πTemp001[1] = πg.NewStr("%d").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewLongFromBytes([]byte{0x2a,}).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1148: self.checkequal('42', '%d', '__mod__', 42.0)
							πF.SetLineno(1148)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ß42.ToObject()
							πTemp001[1] = πg.NewStr("%d").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewFloat(42.0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1149: self.checkequal(slongvalue, '%d', '__mod__', longvalue)
							πF.SetLineno(1149)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µslongvalue, "slongvalue"); πE != nil {
								continue
							}
							πTemp001[0] = µslongvalue
							πTemp001[1] = πg.NewStr("%d").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							if πE = πg.CheckLocal(πF, µlongvalue, "longvalue"); πE != nil {
								continue
							}
							πTemp001[3] = µlongvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1150: self.checkcall('%d', '__mod__', float(longvalue))
							πF.SetLineno(1150)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("%d").ToObject()
							πTemp001[1] = ß__mod__.ToObject()
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlongvalue, "longvalue"); πE != nil {
								continue
							}
							πTemp008[0] = µlongvalue
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckcall, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1151: self.checkequal('0042.00', '%07.2f', '__mod__', 42)
							πF.SetLineno(1151)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("0042.00").ToObject()
							πTemp001[1] = πg.NewStr("%07.2f").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1152: self.checkequal('0042.00', '%07.2F', '__mod__', 42)
							πF.SetLineno(1152)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("0042.00").ToObject()
							πTemp001[1] = πg.NewStr("%07.2F").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1154: self.checkraises(TypeError, 'abc', '__mod__')
							πF.SetLineno(1154)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1155: self.checkraises(TypeError, '%(foo)s', '__mod__', 42)
							πF.SetLineno(1155)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%(foo)s").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1156: self.checkraises(TypeError, '%s%s', '__mod__', (42,))
							πF.SetLineno(1156)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%s%s").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp002 = πg.NewTuple1(πg.NewInt(42).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1157: self.checkraises(TypeError, '%c', '__mod__', (None,))
							πF.SetLineno(1157)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%c").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple1(πTemp003).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1158: self.checkraises(ValueError, '%(foo', '__mod__', {})
							πF.SetLineno(1158)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%(foo").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp009 = πg.NewDict()
							πTemp002 = πTemp009.ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1159: self.checkraises(TypeError, '%(foo)s %(bar)s', '__mod__', ('foo', 42))
							πF.SetLineno(1159)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%(foo)s %(bar)s").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp002 = πg.NewTuple2(ßfoo.ToObject(), πg.NewInt(42).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1160: self.checkraises(TypeError, '%d', '__mod__', "42") # not numeric
							πF.SetLineno(1160)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%d").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp001[3] = ß42.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1164: self.checkequal('bar', '%((foo))s', '__mod__', {'(foo)': 'bar'})
							πF.SetLineno(1164)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = ßbar.ToObject()
							πTemp001[1] = πg.NewStr("%((foo))s").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp009 = πg.NewDict()
							if πE = πTemp009.SetItem(πF, πg.NewStr("(foo)").ToObject(), ßbar.ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp009.ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1167: self.checkequal(103*'a'+'x', '%sx', '__mod__', 103*'a')
							πF.SetLineno(1167)
							πTemp001 = πF.MakeArgs(4)
							if πTemp003, πE = πg.Mul(πF, πg.NewInt(103).ToObject(), ßa.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, ßx.ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%sx").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							if πTemp002, πE = πg.Mul(πF, πg.NewInt(103).ToObject(), ßa.ToObject()); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1169: self.checkraises(TypeError, '%*s', '__mod__', ('foo', 'bar'))
							πF.SetLineno(1169)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%*s").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp002 = πg.NewTuple2(ßfoo.ToObject(), ßbar.ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1170: self.checkraises(TypeError, '%10.*f', '__mod__', ('foo', 42.))
							πF.SetLineno(1170)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%10.*f").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp002 = πg.NewTuple2(ßfoo.ToObject(), πg.NewFloat(42.0).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1171: self.checkraises(ValueError, '%10', '__mod__', (42,))
							πF.SetLineno(1171)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("%10").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							πTemp002 = πg.NewTuple1(πg.NewInt(42).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1173: class X(object): pass
							πF.SetLineno(1173)
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
							_, πE = πg.NewCode("X", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp009
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 1173: class X(object): pass
									πF.SetLineno(1173)
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
							if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("X").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
								continue
							}
							µX = πTemp004
							// line 1174: self.checkraises(TypeError, 'abc', '__mod__', X())
							πF.SetLineno(1174)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßabc.ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp002, πE = µX.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1175: class X(Exception):
							πF.SetLineno(1175)
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
							_, πE = πg.NewCode("X", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 1176: def __getitem__(self, k):
									πF.SetLineno(1176)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "k", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µk *πg.Object = πArgs[1]; _ = µk
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 1177: return k
											πF.SetLineno(1177)
											if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
												continue
											}
											πR = µk
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
							if πTemp003, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp003 == nil {
								πTemp003 = πg.TypeType.ToObject()
							}
							if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("X").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
								continue
							}
							µX = πTemp004
							// line 1178: self.checkequal('melon apple', '%(melon)s %(apple)s', '__mod__', X())
							πF.SetLineno(1178)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("melon apple").ToObject()
							πTemp001[1] = πg.NewStr("%(melon)s %(apple)s").ToObject()
							πTemp001[2] = ß__mod__.ToObject()
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp002, πE = µX.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_formatting.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1198: def test_floatformatting(self):
					πF.SetLineno(1198)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_floatformatting", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µprec *πg.Object = πg.UnboundLocal; _ = µprec
						var µformat *πg.Object = πg.UnboundLocal; _ = µformat
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(100).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
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
								µprec = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1201: format = '%%.%if' % prec
							πF.SetLineno(1201)
							if πE = πg.CheckLocal(πF, µprec, "prec"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%%.%if").ToObject(), µprec); πE != nil {
								continue
							}
							µformat = πTemp003
							// line 1202: value = 0.01
							πF.SetLineno(1202)
							µvalue = πg.NewFloat(0.01).ToObject()
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(60).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
								µx = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 1204: value = value * 3.14159265359 / 3.0 * 10.0
							πF.SetLineno(1204)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Mul(πF, µvalue, πg.NewFloat(3.14159265359).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Div(πF, πTemp009, πg.NewFloat(3.0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp007, πg.NewFloat(10.0).ToObject()); πE != nil {
								continue
							}
							µvalue = πTemp004
							// line 1205: self.checkcall(format, "__mod__", value)
							πF.SetLineno(1205)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µformat, "format"); πE != nil {
								continue
							}
							πTemp002[0] = µformat
							πTemp002[1] = ß__mod__.ToObject()
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[2] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcheckcall, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßtest_floatformatting.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1207: def test_inplace_rewrites(self):
					πF.SetLineno(1207)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_inplace_rewrites", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1209: self.checkequal('a', 'A', 'lower')
							πF.SetLineno(1209)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßa.ToObject()
							πTemp001[1] = ßA.ToObject()
							πTemp001[2] = ßlower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1210: self.checkequal(True, 'A', 'isupper')
							πF.SetLineno(1210)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßA.ToObject()
							πTemp001[2] = ßisupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1211: self.checkequal('A', 'a', 'upper')
							πF.SetLineno(1211)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßA.ToObject()
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1212: self.checkequal(True, 'a', 'islower')
							πF.SetLineno(1212)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1214: self.checkequal('a', 'A', 'replace', 'A', 'a')
							πF.SetLineno(1214)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = ßa.ToObject()
							πTemp001[1] = ßA.ToObject()
							πTemp001[2] = ßreplace.ToObject()
							πTemp001[3] = ßA.ToObject()
							πTemp001[4] = ßa.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1215: self.checkequal(True, 'A', 'isupper')
							πF.SetLineno(1215)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßA.ToObject()
							πTemp001[2] = ßisupper.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1217: self.checkequal('A', 'a', 'capitalize')
							πF.SetLineno(1217)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßA.ToObject()
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßcapitalize.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1218: self.checkequal(True, 'a', 'islower')
							πF.SetLineno(1218)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1220: self.checkequal('A', 'a', 'swapcase')
							πF.SetLineno(1220)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßA.ToObject()
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßswapcase.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1221: self.checkequal(True, 'a', 'islower')
							πF.SetLineno(1221)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1223: self.checkequal('A', 'a', 'title')
							πF.SetLineno(1223)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßA.ToObject()
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßtitle.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1224: self.checkequal(True, 'a', 'islower')
							πF.SetLineno(1224)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßa.ToObject()
							πTemp001[2] = ßislower.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_inplace_rewrites.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1226: def test_partition(self):
					πF.SetLineno(1226)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_partition", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µS *πg.Object = πg.UnboundLocal; _ = µS
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
							// line 1228: self.checkequal(('this is the par', 'ti', 'tion method'),
							πF.SetLineno(1228)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(πg.NewStr("this is the par").ToObject(), ßti.ToObject(), πg.NewStr("tion method").ToObject()).ToObject()
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("this is the partition method").ToObject()
							πTemp001[2] = ßpartition.ToObject()
							πTemp001[3] = ßti.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1232: S = 'http://www.python.org'
							πF.SetLineno(1232)
							µS = πg.NewStr("http://www.python.org").ToObject()
							// line 1233: self.checkequal(('http', '://', 'www.python.org'), S, 'partition', '://')
							πF.SetLineno(1233)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(ßhttp.ToObject(), πg.NewStr("://").ToObject(), πg.NewStr("www.python.org").ToObject()).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßpartition.ToObject()
							πTemp001[3] = πg.NewStr("://").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1234: self.checkequal(('http://www.python.org', '', ''), S, 'partition', '?')
							πF.SetLineno(1234)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(πg.NewStr("http://www.python.org").ToObject(), ß.ToObject(), ß.ToObject()).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßpartition.ToObject()
							πTemp001[3] = πg.NewStr("?").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1235: self.checkequal(('', 'http://', 'www.python.org'), S, 'partition', 'http://')
							πF.SetLineno(1235)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(ß.ToObject(), πg.NewStr("http://").ToObject(), πg.NewStr("www.python.org").ToObject()).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßpartition.ToObject()
							πTemp001[3] = πg.NewStr("http://").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1236: self.checkequal(('http://www.python.', 'org', ''), S, 'partition', 'org')
							πF.SetLineno(1236)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(πg.NewStr("http://www.python.").ToObject(), ßorg.ToObject(), ß.ToObject()).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßpartition.ToObject()
							πTemp001[3] = ßorg.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1238: self.checkraises(ValueError, S, 'partition', '')
							πF.SetLineno(1238)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßpartition.ToObject()
							πTemp001[3] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1239: self.checkraises(TypeError, S, 'partition', None)
							πF.SetLineno(1239)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßpartition.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1242: self.assertEqual('a/b/c'.partition(u'/'), ('a', '/', 'b/c'))
							πF.SetLineno(1242)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewUnicode("/").ToObject()
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("a/b/c").ToObject(), ßpartition, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp002 = πg.NewTuple3(ßa.ToObject(), πg.NewStr("/").ToObject(), πg.NewStr("b/c").ToObject()).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_partition.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 1244: def test_rpartition(self):
					πF.SetLineno(1244)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("test_rpartition", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µS *πg.Object = πg.UnboundLocal; _ = µS
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
							// line 1246: self.checkequal(('this is the rparti', 'ti', 'on method'),
							πF.SetLineno(1246)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(πg.NewStr("this is the rparti").ToObject(), ßti.ToObject(), πg.NewStr("on method").ToObject()).ToObject()
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("this is the rpartition method").ToObject()
							πTemp001[2] = ßrpartition.ToObject()
							πTemp001[3] = ßti.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1250: S = 'http://www.python.org'
							πF.SetLineno(1250)
							µS = πg.NewStr("http://www.python.org").ToObject()
							// line 1251: self.checkequal(('http', '://', 'www.python.org'), S, 'rpartition', '://')
							πF.SetLineno(1251)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(ßhttp.ToObject(), πg.NewStr("://").ToObject(), πg.NewStr("www.python.org").ToObject()).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßrpartition.ToObject()
							πTemp001[3] = πg.NewStr("://").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1252: self.checkequal(('', '', 'http://www.python.org'), S, 'rpartition', '?')
							πF.SetLineno(1252)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(ß.ToObject(), ß.ToObject(), πg.NewStr("http://www.python.org").ToObject()).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßrpartition.ToObject()
							πTemp001[3] = πg.NewStr("?").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1253: self.checkequal(('', 'http://', 'www.python.org'), S, 'rpartition', 'http://')
							πF.SetLineno(1253)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(ß.ToObject(), πg.NewStr("http://").ToObject(), πg.NewStr("www.python.org").ToObject()).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßrpartition.ToObject()
							πTemp001[3] = πg.NewStr("http://").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1254: self.checkequal(('http://www.python.', 'org', ''), S, 'rpartition', 'org')
							πF.SetLineno(1254)
							πTemp001 = πF.MakeArgs(4)
							πTemp002 = πg.NewTuple3(πg.NewStr("http://www.python.").ToObject(), ßorg.ToObject(), ß.ToObject()).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßrpartition.ToObject()
							πTemp001[3] = ßorg.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1256: self.checkraises(ValueError, S, 'rpartition', '')
							πF.SetLineno(1256)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßrpartition.ToObject()
							πTemp001[3] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1257: self.checkraises(TypeError, S, 'rpartition', None)
							πF.SetLineno(1257)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µS, "S"); πE != nil {
								continue
							}
							πTemp001[1] = µS
							πTemp001[2] = ßrpartition.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1260: self.assertEqual('a/b/c'.rpartition(u'/'), ('a/b', '/', 'c'))
							πF.SetLineno(1260)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewUnicode("/").ToObject()
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("a/b/c").ToObject(), ßrpartition, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp002 = πg.NewTuple3(πg.NewStr("a/b").ToObject(), πg.NewStr("/").ToObject(), ßc.ToObject()).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_rpartition.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 1262: def test_none_arguments(self):
					πF.SetLineno(1262)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_none_arguments", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
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
							// line 1264: s = 'hello'
							πF.SetLineno(1264)
							µs = ßhello.ToObject()
							// line 1265: self.checkequal(2, s, 'find', 'l', None)
							πF.SetLineno(1265)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßfind.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1266: self.checkequal(3, s, 'find', 'l', -2, None)
							πF.SetLineno(1266)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßfind.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1267: self.checkequal(2, s, 'find', 'l', None, -2)
							πF.SetLineno(1267)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßfind.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1268: self.checkequal(0, s, 'find', 'h', None, None)
							πF.SetLineno(1268)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßfind.ToObject()
							πTemp001[3] = ßh.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1270: self.checkequal(3, s, 'rfind', 'l', None)
							πF.SetLineno(1270)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßrfind.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1271: self.checkequal(3, s, 'rfind', 'l', -2, None)
							πF.SetLineno(1271)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßrfind.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1272: self.checkequal(2, s, 'rfind', 'l', None, -2)
							πF.SetLineno(1272)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßrfind.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1273: self.checkequal(0, s, 'rfind', 'h', None, None)
							πF.SetLineno(1273)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßrfind.ToObject()
							πTemp001[3] = ßh.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1275: self.checkequal(2, s, 'index', 'l', None)
							πF.SetLineno(1275)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßindex.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1276: self.checkequal(3, s, 'index', 'l', -2, None)
							πF.SetLineno(1276)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßindex.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1277: self.checkequal(2, s, 'index', 'l', None, -2)
							πF.SetLineno(1277)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßindex.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1278: self.checkequal(0, s, 'index', 'h', None, None)
							πF.SetLineno(1278)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßindex.ToObject()
							πTemp001[3] = ßh.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1280: self.checkequal(3, s, 'rindex', 'l', None)
							πF.SetLineno(1280)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßrindex.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1281: self.checkequal(3, s, 'rindex', 'l', -2, None)
							πF.SetLineno(1281)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßrindex.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1282: self.checkequal(2, s, 'rindex', 'l', None, -2)
							πF.SetLineno(1282)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßrindex.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1283: self.checkequal(0, s, 'rindex', 'h', None, None)
							πF.SetLineno(1283)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßrindex.ToObject()
							πTemp001[3] = ßh.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1285: self.checkequal(2, s, 'count', 'l', None)
							πF.SetLineno(1285)
							πTemp001 = πF.MakeArgs(5)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßcount.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1286: self.checkequal(1, s, 'count', 'l', -2, None)
							πF.SetLineno(1286)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßcount.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1287: self.checkequal(1, s, 'count', 'l', None, -2)
							πF.SetLineno(1287)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßcount.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1288: self.checkequal(0, s, 'count', 'x', None, None)
							πF.SetLineno(1288)
							πTemp001 = πF.MakeArgs(6)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßcount.ToObject()
							πTemp001[3] = ßx.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1290: self.checkequal(True, s, 'endswith', 'o', None)
							πF.SetLineno(1290)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßo.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1291: self.checkequal(True, s, 'endswith', 'lo', -2, None)
							πF.SetLineno(1291)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßlo.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1292: self.checkequal(True, s, 'endswith', 'l', None, -2)
							πF.SetLineno(1292)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1293: self.checkequal(False, s, 'endswith', 'x', None, None)
							πF.SetLineno(1293)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßendswith.ToObject()
							πTemp001[3] = ßx.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1295: self.checkequal(True, s, 'startswith', 'h', None)
							πF.SetLineno(1295)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßh.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1296: self.checkequal(True, s, 'startswith', 'l', -2, None)
							πF.SetLineno(1296)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßl.ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1297: self.checkequal(True, s, 'startswith', 'h', None, -2)
							πF.SetLineno(1297)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßh.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1298: self.checkequal(False, s, 'startswith', 'x', None, None)
							πF.SetLineno(1298)
							πTemp001 = πF.MakeArgs(6)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[1] = µs
							πTemp001[2] = ßstartswith.ToObject()
							πTemp001[3] = ßx.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_none_arguments.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 1300: def test_find_etc_raise_correct_error_messages(self):
					πF.SetLineno(1300)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_find_etc_raise_correct_error_messages", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πg.UnboundLocal; _ = µs
						var µx *πg.Object = πg.UnboundLocal; _ = µx
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
							// line 1302: s = 'hello'
							πF.SetLineno(1302)
							µs = ßhello.ToObject()
							// line 1303: x = 'x'
							πF.SetLineno(1303)
							µx = ßx.ToObject()
							// line 1304: self.assertRaisesRegexp(TypeError, r'\bfind\b', s.find,
							πF.SetLineno(1304)
							πTemp001 = πF.MakeArgs(7)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\\bfind\\b").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßfind, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[3] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaisesRegexp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1306: self.assertRaisesRegexp(TypeError, r'\brfind\b', s.rfind,
							πF.SetLineno(1306)
							πTemp001 = πF.MakeArgs(7)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\\brfind\\b").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßrfind, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[3] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaisesRegexp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1308: self.assertRaisesRegexp(TypeError, r'\bindex\b', s.index,
							πF.SetLineno(1308)
							πTemp001 = πF.MakeArgs(7)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\\bindex\\b").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßindex, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[3] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaisesRegexp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1310: self.assertRaisesRegexp(TypeError, r'\brindex\b', s.rindex,
							πF.SetLineno(1310)
							πTemp001 = πF.MakeArgs(7)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("\\brindex\\b").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßrindex, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[3] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaisesRegexp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1312: self.assertRaisesRegexp(TypeError, r'^count\(', s.count,
							πF.SetLineno(1312)
							πTemp001 = πF.MakeArgs(7)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("^count\\(").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßcount, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[3] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaisesRegexp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1314: self.assertRaisesRegexp(TypeError, r'^startswith\(', s.startswith,
							πF.SetLineno(1314)
							πTemp001 = πF.MakeArgs(7)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("^startswith\\(").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßstartswith, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[3] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaisesRegexp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1316: self.assertRaisesRegexp(TypeError, r'^endswith\(', s.endswith,
							πF.SetLineno(1316)
							πTemp001 = πF.MakeArgs(7)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewStr("^endswith\\(").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßendswith, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[3] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[5] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[6] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertRaisesRegexp, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_find_etc_raise_correct_error_messages.ToObject(), πTemp016); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("MixinStrUnicodeUserStringTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMixinStrUnicodeUserStringTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 1319: class MixinStrStringUserStringTest(object):
			πF.SetLineno(1319)
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
			_, πE = πg.NewCode("MixinStrStringUserStringTest", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 1323: def test_maketrans(self):
					πF.SetLineno(1323)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_maketrans", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1324: self.assertEqual(
							πF.SetLineno(1324)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßabc.ToObject()
							πTemp002[1] = ßxyz.ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(256).ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[1] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp007
							if πTemp005, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp007
							πTemp002 = πF.MakeArgs(2)
							πTemp002[0] = ßabc.ToObject()
							πTemp002[1] = ßxyz.ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßmaketrans, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßassertEqual, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1328: self.assertRaises(ValueError, string.maketrans, 'abc', 'xyzw')
							πF.SetLineno(1328)
							πTemp001 = πF.MakeArgs(4)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßstring); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßmaketrans, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp007
							πTemp001[2] = ßabc.ToObject()
							πTemp001[3] = ßxyzw.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_maketrans.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("MixinStrStringUserStringTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMixinStrStringUserStringTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 1345: class MixinStrUserStringTest(object):
			πF.SetLineno(1345)
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
			_, πE = πg.NewCode("MixinStrUserStringTest", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1350: def test_encoding_decoding(self):
					πF.SetLineno(1350)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_encoding_decoding", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcodecs *πg.Object = πg.UnboundLocal; _ = µcodecs
						var µencoding *πg.Object = πg.UnboundLocal; _ = µencoding
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
						var πTemp001 []*πg.Object
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
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
							// line 1351: codecs = [('rot13', 'uryyb jbeyq'),
							πF.SetLineno(1351)
							πTemp001 = make([]*πg.Object, 4)
							πTemp002 = πg.NewTuple2(ßrot13.ToObject(), πg.NewStr("uryyb jbeyq").ToObject()).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewTuple2(ßbase64.ToObject(), πg.NewStr("aGVsbG8gd29ybGQ=\n").ToObject()).ToObject()
							πTemp001[1] = πTemp002
							πTemp002 = πg.NewTuple2(ßhex.ToObject(), ß68656c6c6f20776f726c64.ToObject()).ToObject()
							πTemp001[2] = πTemp002
							πTemp002 = πg.NewTuple2(ßuu.ToObject(), πg.NewStr("begin 666 <data>\n+:&5L;&\\@=V]R;&0 \n \nend\n").ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µcodecs = πTemp002
							if πE = πg.CheckLocal(πF, µcodecs, "codecs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µcodecs); πE != nil {
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
							if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp005); πE != nil {
									continue
								}
								µencoding = πTemp006
								µdata = πTemp007
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1356: with test_support.check_py3k_warnings():
							πF.SetLineno(1356)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßcheck_py3k_warnings, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp007.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(4)
							// line 1357: self.checkequal(data, 'hello world', 'encode', encoding)
							πF.SetLineno(1357)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							πTemp001[1] = πg.NewStr("hello world").ToObject()
							πTemp001[2] = ßencode.ToObject()
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp001[3] = µencoding
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
						Label4:
							πTemp010, πTemp011 = nil, nil
							if πE != nil {
								πTemp010, πTemp011 = πF.ExcInfo()
							}
							if πTemp010 != nil {
								πTemp012 = πTemp010.Type()
								if πTemp008, πE = πTemp006.Call(πF, πg.Args{πTemp005, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp008, πE = πTemp006.Call(πF, πg.Args{πTemp005, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp010 != nil && πTemp004 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 1358: with test_support.check_py3k_warnings():
							πF.SetLineno(1358)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßtest_support); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßcheck_py3k_warnings, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp009.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, πTemp008.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp013.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							// line 1359: self.checkequal('hello world', data, 'decode', encoding)
							πF.SetLineno(1359)
							πTemp001 = πF.MakeArgs(4)
							πTemp001[0] = πg.NewStr("hello world").ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[1] = µdata
							πTemp001[2] = ßdecode.ToObject()
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp001[3] = µencoding
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetAttr(πF, µself, ßcheckequal, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp014.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
						Label5:
							πTemp010, πTemp011 = nil, nil
							if πE != nil {
								πTemp010, πTemp011 = πF.ExcInfo()
							}
							if πTemp010 != nil {
								πTemp012 = πTemp010.Type()
								if πTemp014, πE = πTemp009.Call(πF, πg.Args{πTemp008, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp014, πE = πTemp009.Call(πF, πg.Args{πTemp008, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp014); πE != nil {
								continue
							}
							if πTemp010 != nil && πTemp004 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1372: self.checkraises(TypeError, 'xyz', 'decode', 42)
							πF.SetLineno(1372)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßxyz.ToObject()
							πTemp001[2] = ßdecode.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1373: self.checkraises(TypeError, 'xyz', 'encode', 42)
							πF.SetLineno(1373)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = ßxyz.ToObject()
							πTemp001[2] = ßencode.ToObject()
							πTemp001[3] = πg.NewInt(42).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheckraises, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_encoding_decoding.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1349: @unittest.skipUnless(test_support.have_unicode, 'no unicode support')
					πF.SetLineno(1349)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtest_encoding_decoding); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					πTemp005 = πF.MakeArgs(2)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtest_support); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßhave_unicode, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					πTemp005[1] = πg.NewStr("no unicode support").ToObject()
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßskipUnless, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßtest_encoding_decoding.ToObject(), πTemp006); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("MixinStrUserStringTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMixinStrUserStringTest.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 1376: class MixinStrUnicodeTest(object):
			πF.SetLineno(1376)
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
			_, πE = πg.NewCode("MixinStrUnicodeTest", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 1379: def test_bug1001011(self):
					πF.SetLineno(1379)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_bug1001011", "/usr/lib/python2.7/test/string_tests.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µt *πg.Object = πg.UnboundLocal; _ = µt
						var µsubclass *πg.Object = πg.UnboundLocal; _ = µsubclass
						var µs1 *πg.Object = πg.UnboundLocal; _ = µs1
						var µs2 *πg.Object = πg.UnboundLocal; _ = µs2
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
						var πTemp006 []*πg.Object
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
							// line 1384: t = self.type2test
							πF.SetLineno(1384)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtype2test, nil); πE != nil {
								continue
							}
							µt = πTemp001
							// line 1385: class subclass(t):
							πF.SetLineno(1385)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp003[0] = µt
							πTemp002 = πg.NewDict()
							if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp002.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
								continue
							}
							_, πE = πg.NewCode("subclass", "/usr/lib/python2.7/test/string_tests.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp002
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 1386: pass
									πF.SetLineno(1386)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("subclass").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp002.ToObject()}, nil); πE != nil {
								continue
							}
							µsubclass = πTemp005
							// line 1387: s1 = subclass("abcd")
							πF.SetLineno(1387)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßabcd.ToObject()
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							if πTemp001, πE = µsubclass.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs1 = πTemp001
							// line 1388: s2 = t().join([s1])
							πF.SetLineno(1388)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp006[0] = µs1
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = µt.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs2 = πTemp001
							// line 1389: self.assertTrue(s1 is not s2)
							πF.SetLineno(1389)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µs1 != µs2).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1390: self.assertTrue(type(s2) is t)
							πF.SetLineno(1390)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp006[0] = µs2
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005 == µt).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1392: s1 = t("abcd")
							πF.SetLineno(1392)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßabcd.ToObject()
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = µt.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs1 = πTemp001
							// line 1393: s2 = t().join([s1])
							πF.SetLineno(1393)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp006[0] = µs1
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = µt.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs2 = πTemp001
							// line 1394: self.assertTrue(s1 is s2)
							πF.SetLineno(1394)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µs1 == µs2).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µt == πTemp004).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µt == πTemp004).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label2
							}
							goto Label3
							// line 1397: if t is unicode:
							πF.SetLineno(1397)
						Label1:
							// line 1398: s1 = subclass("abcd")
							πF.SetLineno(1398)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßabcd.ToObject()
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							if πTemp001, πE = µsubclass.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs1 = πTemp001
							// line 1399: s2 = "".join([s1])
							πF.SetLineno(1399)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp006[0] = µs1
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs2 = πTemp004
							// line 1400: self.assertTrue(s1 is not s2)
							πF.SetLineno(1400)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µs1 != µs2).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1401: self.assertTrue(type(s2) is t)
							πF.SetLineno(1401)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp006[0] = µs2
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005 == µt).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1403: s1 = t("abcd")
							πF.SetLineno(1403)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßabcd.ToObject()
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = µt.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs1 = πTemp001
							// line 1404: s2 = "".join([s1])
							πF.SetLineno(1404)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp006[0] = µs1
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs2 = πTemp004
							// line 1405: self.assertTrue(s1 is s2)
							πF.SetLineno(1405)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µs1 == µs2).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label4
							// line 1407: elif t is str:
							πF.SetLineno(1407)
						Label2:
							// line 1408: s1 = subclass("abcd")
							πF.SetLineno(1408)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßabcd.ToObject()
							if πE = πg.CheckLocal(πF, µsubclass, "subclass"); πE != nil {
								continue
							}
							if πTemp001, πE = µsubclass.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs1 = πTemp001
							// line 1409: s2 = u"".join([s1])
							πF.SetLineno(1409)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp006[0] = µs1
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewUnicode("").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs2 = πTemp004
							// line 1410: self.assertTrue(s1 is not s2)
							πF.SetLineno(1410)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µs1 != µs2).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1411: self.assertTrue(type(s2) is unicode) # promotes!
							πF.SetLineno(1411)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp006[0] = µs2
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005 == πTemp004).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1413: s1 = t("abcd")
							πF.SetLineno(1413)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßabcd.ToObject()
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = µt.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs1 = πTemp001
							// line 1414: s2 = u"".join([s1])
							πF.SetLineno(1414)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							πTemp006[0] = µs1
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, πg.NewUnicode("").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs2 = πTemp004
							// line 1415: self.assertTrue(s1 is not s2)
							πF.SetLineno(1415)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µs1 != µs2).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1416: self.assertTrue(type(s2) is unicode) # promotes!
							πF.SetLineno(1416)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
								continue
							}
							πTemp006[0] = µs2
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005 == πTemp004).ToObject()
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label4
						Label3:
							// line 1419: self.fail("unexpected type for MixinStrUnicodeTest %r" % t)
							πF.SetLineno(1419)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("unexpected type for MixinStrUnicodeTest %r").ToObject(), µt); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label4
						Label4:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_bug1001011.ToObject(), πTemp001); πE != nil {
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("MixinStrUnicodeTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßMixinStrUnicodeTest.ToObject(), πTemp005); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("test.string_tests", Code)
}

