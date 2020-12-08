package test_operator
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/operator"
	// _ "github.com/pygolin/stdlib/unittest"
	// _ "github.com/pygolin/stdlib/test/test_support"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß10 := πg.InternStr("10")
		ß2 := πg.InternStr("2")
		ß5 := πg.InternStr("5")
		ßABCDE := πg.InternStr("ABCDE")
		ßAttributeError := πg.InternStr("AttributeError")
		ßC := πg.InternStr("C")
		ßIndexError := πg.InternStr("IndexError")
		ßKeyError := πg.InternStr("KeyError")
		ßLookupError := πg.InternStr("LookupError")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßOperatorTestCase := πg.InternStr("OperatorTestCase")
		ßSeq1 := πg.InternStr("Seq1")
		ßSeq2 := πg.InternStr("Seq2")
		ßSyntaxError := πg.InternStr("SyntaxError")
		ßTestCase := πg.InternStr("TestCase")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ßX := πg.InternStr("X")
		ßY := πg.InternStr("Y")
		ßZ := πg.InternStr("Z")
		ß_ := πg.InternStr("_")
		ß__ := πg.InternStr("__")
		ß__add__ := πg.InternStr("__add__")
		ß__bool__ := πg.InternStr("__bool__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__getattr__ := πg.InternStr("__getattr__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__iadd__ := πg.InternStr("__iadd__")
		ß__iand__ := πg.InternStr("__iand__")
		ß__ifloordiv__ := πg.InternStr("__ifloordiv__")
		ß__ilshift__ := πg.InternStr("__ilshift__")
		ß__imod__ := πg.InternStr("__imod__")
		ß__imul__ := πg.InternStr("__imul__")
		ß__init__ := πg.InternStr("__init__")
		ß__ior__ := πg.InternStr("__ior__")
		ß__ipow__ := πg.InternStr("__ipow__")
		ß__irshift__ := πg.InternStr("__irshift__")
		ß__isub__ := πg.InternStr("__isub__")
		ß__itruediv__ := πg.InternStr("__itruediv__")
		ß__ixor__ := πg.InternStr("__ixor__")
		ß__len__ := πg.InternStr("__len__")
		ß__length_hint__ := πg.InternStr("__length_hint__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__mul__ := πg.InternStr("__mul__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__rmul__ := πg.InternStr("__rmul__")
		ßabc := πg.InternStr("abc")
		ßabs := πg.InternStr("abs")
		ßadd := πg.InternStr("add")
		ßand_ := πg.InternStr("and_")
		ßapple := πg.InternStr("apple")
		ßarthur := πg.InternStr("arthur")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertIs := πg.InternStr("assertIs")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertTrue := πg.InternStr("assertTrue")
		ßattrgetter := πg.InternStr("attrgetter")
		ßbanana := πg.InternStr("banana")
		ßbar := πg.InternStr("bar")
		ßbaz := πg.InternStr("baz")
		ßchild := πg.InternStr("child")
		ßconcat := πg.InternStr("concat")
		ßcontains := πg.InternStr("contains")
		ßcountOf := πg.InternStr("countOf")
		ßdelitem := πg.InternStr("delitem")
		ßdict := πg.InternStr("dict")
		ßdir := πg.InternStr("dir")
		ßeggs := πg.InternStr("eggs")
		ßeq := πg.InternStr("eq")
		ßexpectedFailure := πg.InternStr("expectedFailure")
		ßfloordiv := πg.InternStr("floordiv")
		ßfoo := πg.InternStr("foo")
		ßge := πg.InternStr("ge")
		ßgetattr := πg.InternStr("getattr")
		ßgetitem := πg.InternStr("getitem")
		ßgt := πg.InternStr("gt")
		ßiadd := πg.InternStr("iadd")
		ßiand := πg.InternStr("iand")
		ßiconcat := πg.InternStr("iconcat")
		ßifloordiv := πg.InternStr("ifloordiv")
		ßilshift := πg.InternStr("ilshift")
		ßimod := πg.InternStr("imod")
		ßimul := πg.InternStr("imul")
		ßindexOf := πg.InternStr("indexOf")
		ßinv := πg.InternStr("inv")
		ßinvert := πg.InternStr("invert")
		ßior := πg.InternStr("ior")
		ßipow := πg.InternStr("ipow")
		ßirshift := πg.InternStr("irshift")
		ßis_ := πg.InternStr("is_")
		ßis_not := πg.InternStr("is_not")
		ßisub := πg.InternStr("isub")
		ßitemgetter := πg.InternStr("itemgetter")
		ßiter := πg.InternStr("iter")
		ßitruediv := πg.InternStr("itruediv")
		ßixor := πg.InternStr("ixor")
		ßjohnson := πg.InternStr("johnson")
		ßkey := πg.InternStr("key")
		ßle := πg.InternStr("le")
		ßlen := πg.InternStr("len")
		ßlength_hint := πg.InternStr("length_hint")
		ßlist := πg.InternStr("list")
		ßlshift := πg.InternStr("lshift")
		ßlst := πg.InternStr("lst")
		ßlt := πg.InternStr("lt")
		ßmap := πg.InternStr("map")
		ßmethodcaller := πg.InternStr("methodcaller")
		ßmod := πg.InternStr("mod")
		ßmul := πg.InternStr("mul")
		ßname := πg.InternStr("name")
		ßne := πg.InternStr("ne")
		ßneg := πg.InternStr("neg")
		ßnonkey := πg.InternStr("nonkey")
		ßobject := πg.InternStr("object")
		ßoperator := πg.InternStr("operator")
		ßor_ := πg.InternStr("or_")
		ßorange := πg.InternStr("orange")
		ßpear := πg.InternStr("pear")
		ßpos := πg.InternStr("pos")
		ßpow := πg.InternStr("pow")
		ßpy := πg.InternStr("py")
		ßpython := πg.InternStr("python")
		ßrange := πg.InternStr("range")
		ßrank := πg.InternStr("rank")
		ßrshift := πg.InternStr("rshift")
		ßrun_unittest := πg.InternStr("run_unittest")
		ßself := πg.InternStr("self")
		ßsetitem := πg.InternStr("setitem")
		ßsorted := πg.InternStr("sorted")
		ßspam := πg.InternStr("spam")
		ßstartswith := πg.InternStr("startswith")
		ßstr := πg.InternStr("str")
		ßstrip := πg.InternStr("strip")
		ßsub := πg.InternStr("sub")
		ßtest_abs := πg.InternStr("test_abs")
		ßtest_add := πg.InternStr("test_add")
		ßtest_attrgetter := πg.InternStr("test_attrgetter")
		ßtest_bitwise_and := πg.InternStr("test_bitwise_and")
		ßtest_bitwise_or := πg.InternStr("test_bitwise_or")
		ßtest_bitwise_xor := πg.InternStr("test_bitwise_xor")
		ßtest_complex_operator := πg.InternStr("test_complex_operator")
		ßtest_concat := πg.InternStr("test_concat")
		ßtest_contains := πg.InternStr("test_contains")
		ßtest_countOf := πg.InternStr("test_countOf")
		ßtest_delitem := πg.InternStr("test_delitem")
		ßtest_dunder_is_original := πg.InternStr("test_dunder_is_original")
		ßtest_eq := πg.InternStr("test_eq")
		ßtest_floordiv := πg.InternStr("test_floordiv")
		ßtest_ge := πg.InternStr("test_ge")
		ßtest_getitem := πg.InternStr("test_getitem")
		ßtest_gt := πg.InternStr("test_gt")
		ßtest_indexOf := πg.InternStr("test_indexOf")
		ßtest_inplace := πg.InternStr("test_inplace")
		ßtest_invert := πg.InternStr("test_invert")
		ßtest_is := πg.InternStr("test_is")
		ßtest_is_not := πg.InternStr("test_is_not")
		ßtest_itemgetter := πg.InternStr("test_itemgetter")
		ßtest_le := πg.InternStr("test_le")
		ßtest_length_hint := πg.InternStr("test_length_hint")
		ßtest_lshift := πg.InternStr("test_lshift")
		ßtest_lt := πg.InternStr("test_lt")
		ßtest_main := πg.InternStr("test_main")
		ßtest_methodcaller := πg.InternStr("test_methodcaller")
		ßtest_mod := πg.InternStr("test_mod")
		ßtest_mul := πg.InternStr("test_mul")
		ßtest_ne := πg.InternStr("test_ne")
		ßtest_neg := πg.InternStr("test_neg")
		ßtest_pos := πg.InternStr("test_pos")
		ßtest_pow := πg.InternStr("test_pow")
		ßtest_rshift := πg.InternStr("test_rshift")
		ßtest_setitem := πg.InternStr("test_setitem")
		ßtest_sub := πg.InternStr("test_sub")
		ßtest_support := πg.InternStr("test_support")
		ßtest_truediv := πg.InternStr("test_truediv")
		ßtest_truth := πg.InternStr("test_truth")
		ßthomas := πg.InternStr("thomas")
		ßthon := πg.InternStr("thon")
		ßtruediv := πg.InternStr("truediv")
		ßtruth := πg.InternStr("truth")
		ßtype := πg.InternStr("type")
		ßunittest := πg.InternStr("unittest")
		ßval := πg.InternStr("val")
		ßvalue := πg.InternStr("value")
		ßx := πg.InternStr("x")
		ßxor := πg.InternStr("xor")
		ßxyzpdq := πg.InternStr("xyzpdq")
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
			// line 2: import operator
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "operator"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßoperator.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: from test import test_support
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "test.test_support"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[1]
			if πE = πF.Globals().SetItem(πF, ßtest_support.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: class Seq1(object):
			πF.SetLineno(5)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
			_, πE = πg.NewCode("Seq1", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 6: def __init__(self, lst):
					πF.SetLineno(6)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "lst", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlst *πg.Object = πArgs[1]; _ = µlst
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 7: self.lst = lst
							πF.SetLineno(7)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlst); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlst, πTemp001); πE != nil {
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
					// line 8: def __len__(self):
					πF.SetLineno(8)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 9: return len(self.lst)
							πF.SetLineno(9)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
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
					// line 10: def __getitem__(self, i):
					πF.SetLineno(10)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 11: return self.lst[i]
							πF.SetLineno(11)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
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
					// line 12: def __add__(self, other):
					πF.SetLineno(12)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__add__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 13: return self.lst + other.lst
							πF.SetLineno(13)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßlst, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 14: def __mul__(self, other):
					πF.SetLineno(14)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__mul__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 15: return self.lst * other
							πF.SetLineno(15)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, µother); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__mul__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 16: def __rmul__(self, other):
					πF.SetLineno(16)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__rmul__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 17: return other * self.lst
							πF.SetLineno(17)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µother, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__rmul__.ToObject(), πTemp007); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Seq1").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSeq1.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 19: class Seq2(object):
			πF.SetLineno(19)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
			_, πE = πg.NewCode("Seq2", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 20: def __init__(self, lst):
					πF.SetLineno(20)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "lst", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlst *πg.Object = πArgs[1]; _ = µlst
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 21: self.lst = lst
							πF.SetLineno(21)
							if πE = πg.CheckLocal(πF, µlst, "lst"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlst); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlst, πTemp001); πE != nil {
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
					// line 22: def __len__(self):
					πF.SetLineno(22)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 23: return len(self.lst)
							πF.SetLineno(23)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
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
					// line 24: def __getitem__(self, i):
					πF.SetLineno(24)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 25: return self.lst[i]
							πF.SetLineno(25)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = µi
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
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
					// line 26: def __add__(self, other):
					πF.SetLineno(26)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__add__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 27: return self.lst + other.lst
							πF.SetLineno(27)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µother, ßlst, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__add__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 28: def __mul__(self, other):
					πF.SetLineno(28)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__mul__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 29: return self.lst * other
							πF.SetLineno(29)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, µother); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__mul__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 30: def __rmul__(self, other):
					πF.SetLineno(30)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__rmul__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							// line 31: return other * self.lst
							πF.SetLineno(31)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlst, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µother, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__rmul__.ToObject(), πTemp007); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Seq2").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSeq2.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 33: class OperatorTestCase(unittest.TestCase):
			πF.SetLineno(33)
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
			_, πE = πg.NewCode("OperatorTestCase", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp014 []*πg.Object
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
				var πTemp036 *πg.Object
				_ = πTemp036
				var πTemp037 *πg.Object
				_ = πTemp037
				var πTemp038 *πg.Object
				_ = πTemp038
				var πTemp039 *πg.Object
				_ = πTemp039
				var πTemp040 *πg.Object
				_ = πTemp040
				var πTemp041 *πg.Object
				_ = πTemp041
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 34: def test_lt(self):
					πF.SetLineno(34)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("test_lt", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 36: self.assertRaises(TypeError, operator.lt)
							πF.SetLineno(36)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlt, nil); πE != nil {
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
							// line 37: self.assertFalse(operator.lt(1, 0))
							πF.SetLineno(37)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 38: self.assertFalse(operator.lt(1, 0.0))
							πF.SetLineno(38)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(0.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 39: self.assertFalse(operator.lt(1, 1))
							πF.SetLineno(39)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 40: self.assertFalse(operator.lt(1, 1.0))
							πF.SetLineno(40)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(1.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 41: self.assertTrue(operator.lt(1, 2))
							πF.SetLineno(41)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 42: self.assertTrue(operator.lt(1, 2.0))
							πF.SetLineno(42)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(2.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßtest_lt.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 44: def test_le(self):
					πF.SetLineno(44)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("test_le", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 46: self.assertRaises(TypeError, operator.le)
							πF.SetLineno(46)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßle, nil); πE != nil {
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
							// line 47: self.assertFalse(operator.le(1, 0))
							πF.SetLineno(47)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 48: self.assertFalse(operator.le(1, 0.0))
							πF.SetLineno(48)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(0.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 49: self.assertTrue(operator.le(1, 1))
							πF.SetLineno(49)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 50: self.assertTrue(operator.le(1, 1.0))
							πF.SetLineno(50)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(1.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 51: self.assertTrue(operator.le(1, 2))
							πF.SetLineno(51)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 52: self.assertTrue(operator.le(1, 2.0))
							πF.SetLineno(52)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(2.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßle, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßtest_le.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 54: def test_eq(self):
					πF.SetLineno(54)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("test_eq", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 56: class C(object):
							πF.SetLineno(56)
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
							_, πE = πg.NewCode("C", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 57: def __eq__(self, other):
									πF.SetLineno(57)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
												continue
											}
											// line 58: raise SyntaxError
											πF.SetLineno(58)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("C").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µC = πTemp005
							// line 59: self.assertRaises(TypeError, operator.eq)
							πF.SetLineno(59)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
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
							// line 60: self.assertRaises(SyntaxError, operator.eq, C(), C())
							πF.SetLineno(60)
							πTemp003 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp002, πE = µC.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp002
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp002, πE = µC.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[3] = πTemp002
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
							// line 61: self.assertFalse(operator.eq(1, 0))
							πF.SetLineno(61)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 62: self.assertFalse(operator.eq(1, 0.0))
							πF.SetLineno(62)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewFloat(0.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 63: self.assertTrue(operator.eq(1, 1))
							πF.SetLineno(63)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 64: self.assertTrue(operator.eq(1, 1.0))
							πF.SetLineno(64)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewFloat(1.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 65: self.assertFalse(operator.eq(1, 2))
							πF.SetLineno(65)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 66: self.assertFalse(operator.eq(1, 2.0))
							πF.SetLineno(66)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewFloat(2.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_eq.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 68: def test_ne(self):
					πF.SetLineno(68)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("test_ne", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 70: class C(object):
							πF.SetLineno(70)
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
							_, πE = πg.NewCode("C", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 71: def __ne__(self, other):
									πF.SetLineno(71)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
												continue
											}
											// line 72: raise SyntaxError
											πF.SetLineno(72)
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
									if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("C").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µC = πTemp005
							// line 73: self.assertRaises(TypeError, operator.ne)
							πF.SetLineno(73)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
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
							// line 74: self.assertRaises(SyntaxError, operator.ne, C(), C())
							πF.SetLineno(74)
							πTemp003 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp002, πE = µC.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp002
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp002, πE = µC.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[3] = πTemp002
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
							// line 75: self.assertTrue(operator.ne(1, 0))
							πF.SetLineno(75)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 76: self.assertTrue(operator.ne(1, 0.0))
							πF.SetLineno(76)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewFloat(0.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 77: self.assertFalse(operator.ne(1, 1))
							πF.SetLineno(77)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 78: self.assertFalse(operator.ne(1, 1.0))
							πF.SetLineno(78)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewFloat(1.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 79: self.assertTrue(operator.ne(1, 2))
							πF.SetLineno(79)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 80: self.assertTrue(operator.ne(1, 2.0))
							πF.SetLineno(80)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewFloat(2.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_ne.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 82: def test_ge(self):
					πF.SetLineno(82)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("test_ge", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 84: self.assertRaises(TypeError, operator.ge)
							πF.SetLineno(84)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßge, nil); πE != nil {
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
							// line 85: self.assertTrue(operator.ge(1, 0))
							πF.SetLineno(85)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßge, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 86: self.assertTrue(operator.ge(1, 0.0))
							πF.SetLineno(86)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(0.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßge, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 87: self.assertTrue(operator.ge(1, 1))
							πF.SetLineno(87)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßge, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 88: self.assertTrue(operator.ge(1, 1.0))
							πF.SetLineno(88)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(1.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßge, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 89: self.assertFalse(operator.ge(1, 2))
							πF.SetLineno(89)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßge, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 90: self.assertFalse(operator.ge(1, 2.0))
							πF.SetLineno(90)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(2.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßge, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_ge.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 92: def test_gt(self):
					πF.SetLineno(92)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("test_gt", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 94: self.assertRaises(TypeError, operator.gt)
							πF.SetLineno(94)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgt, nil); πE != nil {
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
							// line 95: self.assertTrue(operator.gt(1, 0))
							πF.SetLineno(95)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 96: self.assertTrue(operator.gt(1, 0.0))
							πF.SetLineno(96)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(0.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 97: self.assertFalse(operator.gt(1, 1))
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 98: self.assertFalse(operator.gt(1, 1.0))
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(1.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 99: self.assertFalse(operator.gt(1, 2))
							πF.SetLineno(99)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 100: self.assertFalse(operator.gt(1, 2.0))
							πF.SetLineno(100)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewFloat(2.0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_gt.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 102: def test_abs(self):
					πF.SetLineno(102)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("test_abs", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 104: self.assertRaises(TypeError, operator.abs)
							πF.SetLineno(104)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßabs, nil); πE != nil {
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
							// line 105: self.assertRaises(TypeError, operator.abs, None)
							πF.SetLineno(105)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßabs, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
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
							// line 106: self.assertEqual(operator.abs(-1), 1)
							πF.SetLineno(106)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
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
							// line 107: self.assertEqual(operator.abs(1), 1)
							πF.SetLineno(107)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßabs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
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
					if πE = πClass.SetItem(πF, ßtest_abs.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 109: def test_add(self):
					πF.SetLineno(109)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("test_add", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 111: self.assertRaises(TypeError, operator.add)
							πF.SetLineno(111)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßadd, nil); πE != nil {
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
							// line 112: self.assertRaises(TypeError, operator.add, None, None)
							πF.SetLineno(112)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßadd, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 113: self.assertTrue(operator.add(3, 4) == 7)
							πF.SetLineno(113)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(3).ToObject()
							πTemp004[1] = πg.NewInt(4).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßadd, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(7).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_add.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 115: def test_bitwise_and(self):
					πF.SetLineno(115)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("test_bitwise_and", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 117: self.assertRaises(TypeError, operator.and_)
							πF.SetLineno(117)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßand_, nil); πE != nil {
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
							// line 118: self.assertRaises(TypeError, operator.and_, None, None)
							πF.SetLineno(118)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßand_, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 119: self.assertTrue(operator.and_(0xf, 0xa) == 0xa)
							πF.SetLineno(119)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(15).ToObject()
							πTemp004[1] = πg.NewInt(10).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßand_, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(10).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_bitwise_and.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 121: def test_concat(self):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("test_concat", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 123: self.assertRaises(TypeError, operator.concat)
							πF.SetLineno(123)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßconcat, nil); πE != nil {
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
							// line 124: self.assertRaises(TypeError, operator.concat, None, None)
							πF.SetLineno(124)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßconcat, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 125: self.assertTrue(operator.concat('py', 'thon') == 'python')
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = ßpy.ToObject()
							πTemp004[1] = ßthon.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßconcat, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, ßpython.ToObject()); πE != nil {
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
							// line 126: self.assertTrue(operator.concat([1, 2], [3, 4]) == [1, 2, 3, 4])
							πF.SetLineno(126)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp006 = make([]*πg.Object, 2)
							πTemp006[0] = πg.NewInt(1).ToObject()
							πTemp006[1] = πg.NewInt(2).ToObject()
							πTemp003 = πg.NewList(πTemp006...).ToObject()
							πTemp004[0] = πTemp003
							πTemp006 = make([]*πg.Object, 2)
							πTemp006[0] = πg.NewInt(3).ToObject()
							πTemp006[1] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp006...).ToObject()
							πTemp004[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßconcat, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = make([]*πg.Object, 4)
							πTemp004[0] = πg.NewInt(1).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							πTemp004[2] = πg.NewInt(3).ToObject()
							πTemp004[3] = πg.NewInt(4).ToObject()
							πTemp005 = πg.NewList(πTemp004...).ToObject()
							if πTemp002, πE = πg.Eq(πF, πTemp003, πTemp005); πE != nil {
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
							// line 127: self.assertTrue(operator.concat(Seq1([5, 6]), Seq1([7])) == [5, 6, 7])
							πF.SetLineno(127)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 2)
							πTemp007[0] = πg.NewInt(5).ToObject()
							πTemp007[1] = πg.NewInt(6).ToObject()
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSeq1); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp005
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 1)
							πTemp007[0] = πg.NewInt(7).ToObject()
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSeq1); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßconcat, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = make([]*πg.Object, 3)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(6).ToObject()
							πTemp004[2] = πg.NewInt(7).ToObject()
							πTemp005 = πg.NewList(πTemp004...).ToObject()
							if πTemp002, πE = πg.Eq(πF, πTemp003, πTemp005); πE != nil {
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
							// line 128: self.assertTrue(operator.concat(Seq2([5, 6]), Seq2([7])) == [5, 6, 7])
							πF.SetLineno(128)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 2)
							πTemp007[0] = πg.NewInt(5).ToObject()
							πTemp007[1] = πg.NewInt(6).ToObject()
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSeq2); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp005
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 1)
							πTemp007[0] = πg.NewInt(7).ToObject()
							πTemp003 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSeq2); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[1] = πTemp005
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßconcat, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = make([]*πg.Object, 3)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(6).ToObject()
							πTemp004[2] = πg.NewInt(7).ToObject()
							πTemp005 = πg.NewList(πTemp004...).ToObject()
							if πTemp002, πE = πg.Eq(πF, πTemp003, πTemp005); πE != nil {
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
							// line 129: self.assertRaises(TypeError, operator.concat, 13, 29)
							πF.SetLineno(129)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßconcat, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(13).ToObject()
							πTemp001[3] = πg.NewInt(29).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_concat.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 131: def test_countOf(self):
					πF.SetLineno(131)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("test_countOf", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 133: self.assertRaises(TypeError, operator.countOf)
							πF.SetLineno(133)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcountOf, nil); πE != nil {
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
							// line 134: self.assertRaises(TypeError, operator.countOf, None, None)
							πF.SetLineno(134)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcountOf, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 135: self.assertTrue(operator.countOf([1, 2, 1, 3, 1, 4], 3) == 1)
							πF.SetLineno(135)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = make([]*πg.Object, 6)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(1).ToObject()
							πTemp005[5] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewInt(3).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßcountOf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
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
							// line 136: self.assertTrue(operator.countOf([1, 2, 1, 3, 1, 4], 5) == 0)
							πF.SetLineno(136)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = make([]*πg.Object, 6)
							πTemp005[0] = πg.NewInt(1).ToObject()
							πTemp005[1] = πg.NewInt(2).ToObject()
							πTemp005[2] = πg.NewInt(1).ToObject()
							πTemp005[3] = πg.NewInt(3).ToObject()
							πTemp005[4] = πg.NewInt(1).ToObject()
							πTemp005[5] = πg.NewInt(4).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewInt(5).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßcountOf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_countOf.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 139: def test_delitem(self):
					πF.SetLineno(139)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("test_delitem", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 141: a = [4, 3, 2, 1]
							πF.SetLineno(141)
							πTemp001 = make([]*πg.Object, 4)
							πTemp001[0] = πg.NewInt(4).ToObject()
							πTemp001[1] = πg.NewInt(3).ToObject()
							πTemp001[2] = πg.NewInt(2).ToObject()
							πTemp001[3] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µa = πTemp002
							// line 142: self.assertRaises(TypeError, operator.delitem, a)
							πF.SetLineno(142)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdelitem, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
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
							// line 143: self.assertRaises(TypeError, operator.delitem, a, None)
							πF.SetLineno(143)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdelitem, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 144: self.assertTrue(operator.delitem(a, 1) is None)
							πF.SetLineno(144)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004[0] = µa
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßdelitem, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp003 == πTemp005).ToObject()
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
							// line 145: self.assertTrue(a == [4, 2, 1])
							πF.SetLineno(145)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 3)
							πTemp004[0] = πg.NewInt(4).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							πTemp004[2] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp004...).ToObject()
							if πTemp002, πE = πg.Eq(πF, µa, πTemp003); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_delitem.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 138: @unittest.expectedFailure
					πF.SetLineno(138)
					πTemp014 = πF.MakeArgs(1)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßtest_delitem); πE != nil {
						continue
					}
					πTemp014[0] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp016.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πE = πClass.SetItem(πF, ßtest_delitem.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 147: def test_floordiv(self):
					πF.SetLineno(147)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("test_floordiv", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 149: self.assertRaises(TypeError, operator.floordiv, 5)
							πF.SetLineno(149)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfloordiv, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(5).ToObject()
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
							// line 150: self.assertRaises(TypeError, operator.floordiv, None, None)
							πF.SetLineno(150)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfloordiv, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 151: self.assertTrue(operator.floordiv(5, 2) == 2)
							πF.SetLineno(151)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßfloordiv, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_floordiv.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 153: def test_truediv(self):
					πF.SetLineno(153)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("test_truediv", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 155: self.assertRaises(TypeError, operator.truediv, 5)
							πF.SetLineno(155)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtruediv, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(5).ToObject()
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
							// line 156: self.assertRaises(TypeError, operator.truediv, None, None)
							πF.SetLineno(156)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtruediv, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 157: self.assertTrue(operator.truediv(5, 2) == 2.5)
							πF.SetLineno(157)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßtruediv, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewFloat(2.5).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_truediv.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 159: def test_getitem(self):
					πF.SetLineno(159)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("test_getitem", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 161: a = range(10)
							πF.SetLineno(161)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(10).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp003
							// line 162: self.assertRaises(TypeError, operator.getitem)
							πF.SetLineno(162)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetitem, nil); πE != nil {
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
							// line 163: self.assertRaises(TypeError, operator.getitem, a, None)
							πF.SetLineno(163)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetitem, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 164: self.assertTrue(operator.getitem(a, 2) == 2)
							πF.SetLineno(164)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004[0] = µa
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßgetitem, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_getitem.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 166: def test_indexOf(self):
					πF.SetLineno(166)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("test_indexOf", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 168: self.assertRaises(TypeError, operator.indexOf)
							πF.SetLineno(168)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßindexOf, nil); πE != nil {
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
							// line 169: self.assertRaises(TypeError, operator.indexOf, None, None)
							πF.SetLineno(169)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßindexOf, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 170: self.assertTrue(operator.indexOf([4, 3, 2, 1], 3) == 1)
							πF.SetLineno(170)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = make([]*πg.Object, 4)
							πTemp005[0] = πg.NewInt(4).ToObject()
							πTemp005[1] = πg.NewInt(3).ToObject()
							πTemp005[2] = πg.NewInt(2).ToObject()
							πTemp005[3] = πg.NewInt(1).ToObject()
							πTemp003 = πg.NewList(πTemp005...).ToObject()
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewInt(3).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßindexOf, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
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
							// line 171: self.assertRaises(ValueError, operator.indexOf, [4, 3, 2, 1], 0)
							πF.SetLineno(171)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßindexOf, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp004 = make([]*πg.Object, 4)
							πTemp004[0] = πg.NewInt(4).ToObject()
							πTemp004[1] = πg.NewInt(3).ToObject()
							πTemp004[2] = πg.NewInt(2).ToObject()
							πTemp004[3] = πg.NewInt(1).ToObject()
							πTemp002 = πg.NewList(πTemp004...).ToObject()
							πTemp001[2] = πTemp002
							πTemp001[3] = πg.NewInt(0).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_indexOf.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 173: def test_invert(self):
					πF.SetLineno(173)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("test_invert", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 175: self.assertRaises(TypeError, operator.invert)
							πF.SetLineno(175)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinvert, nil); πE != nil {
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
							// line 176: self.assertRaises(TypeError, operator.invert, None)
							πF.SetLineno(176)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinvert, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
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
							// line 177: self.assertEqual(operator.inv(4), -5)
							πF.SetLineno(177)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(4).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinv, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßtest_invert.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 179: def test_lshift(self):
					πF.SetLineno(179)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("test_lshift", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 181: self.assertRaises(TypeError, operator.lshift)
							πF.SetLineno(181)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlshift, nil); πE != nil {
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
							// line 182: self.assertRaises(TypeError, operator.lshift, None, 42)
							πF.SetLineno(182)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlshift, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							πTemp001[3] = πg.NewInt(42).ToObject()
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
							// line 183: self.assertTrue(operator.lshift(5, 1) == 10)
							πF.SetLineno(183)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßlshift, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(10).ToObject()); πE != nil {
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
							// line 184: self.assertTrue(operator.lshift(5, 0) == 5)
							πF.SetLineno(184)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßlshift, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(5).ToObject()); πE != nil {
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
							// line 185: self.assertRaises(ValueError, operator.lshift, 2, -1)
							πF.SetLineno(185)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlshift, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßtest_lshift.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 187: def test_mod(self):
					πF.SetLineno(187)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("test_mod", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 189: self.assertRaises(TypeError, operator.mod)
							πF.SetLineno(189)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmod, nil); πE != nil {
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
							// line 190: self.assertRaises(TypeError, operator.mod, None, 42)
							πF.SetLineno(190)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmod, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							πTemp001[3] = πg.NewInt(42).ToObject()
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
							// line 191: self.assertTrue(operator.mod(5, 2) == 1)
							πF.SetLineno(191)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßmod, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_mod.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 193: def test_mul(self):
					πF.SetLineno(193)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("test_mul", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 195: self.assertRaises(TypeError, operator.mul)
							πF.SetLineno(195)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmul, nil); πE != nil {
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
							// line 196: self.assertRaises(TypeError, operator.mul, None, None)
							πF.SetLineno(196)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmul, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 197: self.assertTrue(operator.mul(5, 2) == 10)
							πF.SetLineno(197)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßmul, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(10).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_mul.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 199: def test_neg(self):
					πF.SetLineno(199)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("test_neg", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 201: self.assertRaises(TypeError, operator.neg)
							πF.SetLineno(201)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßneg, nil); πE != nil {
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
							// line 202: self.assertRaises(TypeError, operator.neg, None)
							πF.SetLineno(202)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßneg, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
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
							// line 203: self.assertEqual(operator.neg(5), -5)
							πF.SetLineno(203)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßneg, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
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
							// line 204: self.assertEqual(operator.neg(-5), 5)
							πF.SetLineno(204)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßneg, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(5).ToObject()
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
							// line 205: self.assertEqual(operator.neg(0), 0)
							πF.SetLineno(205)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßneg, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(0).ToObject()
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
							// line 206: self.assertEqual(operator.neg(-0), 0)
							πF.SetLineno(206)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßneg, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(0).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_neg.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 208: def test_bitwise_or(self):
					πF.SetLineno(208)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("test_bitwise_or", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 210: self.assertRaises(TypeError, operator.or_)
							πF.SetLineno(210)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßor_, nil); πE != nil {
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
							// line 211: self.assertRaises(TypeError, operator.or_, None, None)
							πF.SetLineno(211)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßor_, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 212: self.assertTrue(operator.or_(0xa, 0x5) == 0xf)
							πF.SetLineno(212)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(10).ToObject()
							πTemp004[1] = πg.NewInt(5).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßor_, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(15).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_bitwise_or.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 214: def test_pos(self):
					πF.SetLineno(214)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("test_pos", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 216: self.assertRaises(TypeError, operator.pos)
							πF.SetLineno(216)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpos, nil); πE != nil {
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
							// line 217: self.assertRaises(TypeError, operator.pos, None)
							πF.SetLineno(217)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpos, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
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
							// line 218: self.assertEqual(operator.pos(5), 5)
							πF.SetLineno(218)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(5).ToObject()
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
							// line 219: self.assertEqual(operator.pos(-5), -5)
							πF.SetLineno(219)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
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
							// line 220: self.assertEqual(operator.pos(0), 0)
							πF.SetLineno(220)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(0).ToObject()
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
							// line 221: self.assertEqual(operator.pos(-0), 0)
							πF.SetLineno(221)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(0).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_pos.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 223: def test_pow(self):
					πF.SetLineno(223)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("test_pow", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 225: self.assertRaises(TypeError, operator.pow)
							πF.SetLineno(225)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpow, nil); πE != nil {
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
							// line 226: self.assertRaises(TypeError, operator.pow, None, None)
							πF.SetLineno(226)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpow, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 227: self.assertEqual(operator.pow(3,5), 3**5)
							πF.SetLineno(227)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(3).ToObject()
							πTemp004[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpow, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.Pow(πF, πg.NewInt(3).ToObject(), πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
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
							// line 228: self.assertRaises(TypeError, operator.pow, 1)
							πF.SetLineno(228)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpow, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(1).ToObject()
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
							// line 229: self.assertRaises(TypeError, operator.pow, 1, 2, 3)
							πF.SetLineno(229)
							πTemp001 = πF.MakeArgs(5)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßpow, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(1).ToObject()
							πTemp001[3] = πg.NewInt(2).ToObject()
							πTemp001[4] = πg.NewInt(3).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_pow.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 231: def test_rshift(self):
					πF.SetLineno(231)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("test_rshift", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 233: self.assertRaises(TypeError, operator.rshift)
							πF.SetLineno(233)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrshift, nil); πE != nil {
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
							// line 234: self.assertRaises(TypeError, operator.rshift, None, 42)
							πF.SetLineno(234)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrshift, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							πTemp001[3] = πg.NewInt(42).ToObject()
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
							// line 235: self.assertTrue(operator.rshift(5, 1) == 2)
							πF.SetLineno(235)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßrshift, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
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
							// line 236: self.assertTrue(operator.rshift(5, 0) == 5)
							πF.SetLineno(236)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßrshift, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(5).ToObject()); πE != nil {
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
							// line 237: self.assertRaises(ValueError, operator.rshift, 2, -1)
							πF.SetLineno(237)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrshift, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßtest_rshift.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 239: def test_contains(self):
					πF.SetLineno(239)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("test_contains", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 241: self.assertRaises(TypeError, operator.contains)
							πF.SetLineno(241)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcontains, nil); πE != nil {
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
							// line 242: self.assertRaises(TypeError, operator.contains, None, None)
							πF.SetLineno(242)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcontains, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 243: self.assertTrue(operator.contains(range(4), 2))
							πF.SetLineno(243)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(4).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcontains, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
							// line 244: self.assertFalse(operator.contains(range(4), 5))
							πF.SetLineno(244)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(4).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp003
							πTemp004[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcontains, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_contains.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 246: def test_setitem(self):
					πF.SetLineno(246)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("test_setitem", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 248: a = list(range(3))
							πF.SetLineno(248)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(3).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 249: self.assertRaises(TypeError, operator.setitem, a)
							πF.SetLineno(249)
							πTemp001 = πF.MakeArgs(3)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsetitem, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
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
							// line 250: self.assertRaises(TypeError, operator.setitem, a, None, None)
							πF.SetLineno(250)
							πTemp001 = πF.MakeArgs(5)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsetitem, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[3] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[4] = πTemp003
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
							// line 251: self.assertTrue(operator.setitem(a, 0, 2) is None)
							πF.SetLineno(251)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp002[0] = µa
							πTemp002[1] = πg.NewInt(0).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsetitem, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 == πTemp005).ToObject()
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
							// line 252: self.assertTrue(a == [2, 1, 2])
							πF.SetLineno(252)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = πg.NewInt(2).ToObject()
							πTemp002[1] = πg.NewInt(1).ToObject()
							πTemp002[2] = πg.NewInt(2).ToObject()
							πTemp004 = πg.NewList(πTemp002...).ToObject()
							if πTemp003, πE = πg.Eq(πF, µa, πTemp004); πE != nil {
								continue
							}
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
							// line 253: self.assertRaises(IndexError, operator.setitem, a, 4, 2)
							πF.SetLineno(253)
							πTemp001 = πF.MakeArgs(5)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßsetitem, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp004
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
							πTemp001[3] = πg.NewInt(4).ToObject()
							πTemp001[4] = πg.NewInt(2).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_setitem.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 255: def test_sub(self):
					πF.SetLineno(255)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("test_sub", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 257: self.assertRaises(TypeError, operator.sub)
							πF.SetLineno(257)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
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
							// line 258: self.assertRaises(TypeError, operator.sub, None, None)
							πF.SetLineno(258)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsub, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 259: self.assertTrue(operator.sub(5, 2) == 3)
							πF.SetLineno(259)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(5).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßsub, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(3).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_sub.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 262: def test_truth(self):
					πF.SetLineno(262)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("test_truth", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 []*πg.Object
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
							// line 264: class C(object):
							πF.SetLineno(264)
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
							_, πE = πg.NewCode("C", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 265: def __bool__(self):
									πF.SetLineno(265)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__bool__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
												continue
											}
											// line 266: raise SyntaxError
											πF.SetLineno(266)
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
									if πE = πClass.SetItem(πF, ß__bool__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("C").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µC = πTemp005
							// line 267: self.assertRaises(TypeError, operator.truth)
							πF.SetLineno(267)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtruth, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
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
							// line 268: self.assertRaises(SyntaxError, operator.truth, C())
							πF.SetLineno(268)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtruth, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp002, πE = µC.Call(πF, nil, nil); πE != nil {
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
							// line 269: self.assertTrue(operator.truth(5))
							πF.SetLineno(269)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtruth, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 270: self.assertTrue(operator.truth([0]))
							πF.SetLineno(270)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 1)
							πTemp007[0] = πg.NewInt(0).ToObject()
							πTemp002 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtruth, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 271: self.assertFalse(operator.truth(0))
							πF.SetLineno(271)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtruth, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
							// line 272: self.assertFalse(operator.truth([]))
							πF.SetLineno(272)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßtruth, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_truth.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 261: @unittest.expectedFailure
					πF.SetLineno(261)
					πTemp014 = πF.MakeArgs(1)
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßtest_truth); πE != nil {
						continue
					}
					πTemp014[0] = πTemp032
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp033, πE = πg.GetAttr(πF, πTemp032, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp032, πE = πTemp033.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πE = πClass.SetItem(πF, ßtest_truth.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 274: def test_bitwise_xor(self):
					πF.SetLineno(274)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("test_bitwise_xor", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 276: self.assertRaises(TypeError, operator.xor)
							πF.SetLineno(276)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßxor, nil); πE != nil {
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
							// line 277: self.assertRaises(TypeError, operator.xor, None, None)
							πF.SetLineno(277)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßxor, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
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
							// line 278: self.assertTrue(operator.xor(0xb, 0xc) == 0x7)
							πF.SetLineno(278)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(11).ToObject()
							πTemp004[1] = πg.NewInt(12).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßxor, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewInt(7).ToObject()); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_bitwise_xor.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 280: def test_is(self):
					πF.SetLineno(280)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("test_is", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
							// line 282: a = b = 'xyzpdq'
							πF.SetLineno(282)
							µa = ßxyzpdq.ToObject()
							µb = ßxyzpdq.ToObject()
							// line 283: c = a[:3] + b[3:]
							πF.SetLineno(283)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µa, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µb, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µc = πTemp001
							// line 284: self.assertRaises(TypeError, operator.is_)
							πF.SetLineno(284)
							πTemp005 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßis_, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
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
							// line 285: self.assertTrue(operator.is_(a, b))
							πF.SetLineno(285)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp006[1] = µb
							if πTemp001, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßis_, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_is.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 289: def test_is_not(self):
					πF.SetLineno(289)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("test_is_not", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µb *πg.Object = πg.UnboundLocal; _ = µb
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
							// line 291: a = b = 'xyzpdq'
							πF.SetLineno(291)
							µa = ßxyzpdq.ToObject()
							µb = ßxyzpdq.ToObject()
							// line 292: c = a[:3] + b[3:]
							πF.SetLineno(292)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(3).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µa, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(3).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µb, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µc = πTemp001
							// line 293: self.assertRaises(TypeError, operator.is_not)
							πF.SetLineno(293)
							πTemp005 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßis_not, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
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
							// line 294: self.assertFalse(operator.is_not(a, b))
							πF.SetLineno(294)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp006[1] = µb
							if πTemp001, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßis_not, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertFalse, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 295: self.assertTrue(operator.is_not(a,c))
							πF.SetLineno(295)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[1] = µc
							if πTemp001, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßis_not, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßassertTrue, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_is_not.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 288: @unittest.expectedFailure
					πF.SetLineno(288)
					πTemp014 = πF.MakeArgs(1)
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßtest_is_not); πE != nil {
						continue
					}
					πTemp014[0] = πTemp035
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp036, πE = πg.GetAttr(πF, πTemp035, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp035, πE = πTemp036.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πE = πClass.SetItem(πF, ßtest_is_not.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 298: def test_attrgetter(self):
					πF.SetLineno(298)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("test_attrgetter", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µA *πg.Object = πg.UnboundLocal; _ = µA
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µrecord *πg.Object = πg.UnboundLocal; _ = µrecord
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
						var πTemp006 []*πg.Object
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
							// line 300: class A(object):
							πF.SetLineno(300)
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
							_, πE = πg.NewCode("A", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp001
								_ = πClass
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 301: pass
									πF.SetLineno(301)
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("A").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µA = πTemp005
							// line 302: a = A()
							πF.SetLineno(302)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							if πTemp002, πE = µA.Call(πF, nil, nil); πE != nil {
								continue
							}
							µa = πTemp002
							// line 303: a.name = 'arthur'
							πF.SetLineno(303)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßarthur.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µa, ßname, πTemp002); πE != nil {
								continue
							}
							// line 304: f = operator.attrgetter('name')
							πF.SetLineno(304)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßname.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µf = πTemp002
							// line 305: self.assertEqual(f(a), 'arthur')
							πF.SetLineno(305)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßarthur.ToObject()
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
							// line 306: f = operator.attrgetter('rank')
							πF.SetLineno(306)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßrank.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µf = πTemp002
							// line 307: self.assertRaises(AttributeError, f, a)
							πF.SetLineno(307)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[1] = µf
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp003[2] = µa
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
							// line 308: self.assertRaises(TypeError, operator.attrgetter, 2)
							πF.SetLineno(308)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							πTemp003[2] = πg.NewInt(2).ToObject()
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
							// line 309: self.assertRaises(TypeError, operator.attrgetter)
							πF.SetLineno(309)
							πTemp003 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
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
							// line 312: record = A()
							πF.SetLineno(312)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							if πTemp002, πE = µA.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrecord = πTemp002
							// line 313: record.x = 'X'
							πF.SetLineno(313)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßX.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrecord, "record"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µrecord, ßx, πTemp002); πE != nil {
								continue
							}
							// line 314: record.y = 'Y'
							πF.SetLineno(314)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßY.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrecord, "record"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µrecord, ßy, πTemp002); πE != nil {
								continue
							}
							// line 315: record.z = 'Z'
							πF.SetLineno(315)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßZ.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrecord, "record"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µrecord, ßz, πTemp002); πE != nil {
								continue
							}
							// line 316: self.assertEqual(operator.attrgetter('x','z','y')(record), ('X', 'Z', 'Y'))
							πF.SetLineno(316)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrecord, "record"); πE != nil {
								continue
							}
							πTemp006[0] = µrecord
							πTemp007 = πF.MakeArgs(3)
							πTemp007[0] = ßx.ToObject()
							πTemp007[1] = ßz.ToObject()
							πTemp007[2] = ßy.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp004
							πTemp002 = πg.NewTuple3(ßX.ToObject(), ßZ.ToObject(), ßY.ToObject()).ToObject()
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
							// line 317: self.assertRaises(TypeError, operator.attrgetter, ('x', (), 'y'))
							πF.SetLineno(317)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp004
							πTemp004 = πg.NewTuple0().ToObject()
							πTemp002 = πg.NewTuple3(ßx.ToObject(), πTemp004, ßy.ToObject()).ToObject()
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
							// line 319: class C(object):
							πF.SetLineno(319)
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
							_, πE = πg.NewCode("C", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 320: def __getattr__(self, name):
									πF.SetLineno(320)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "name", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__getattr__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
												continue
											}
											// line 321: raise SyntaxError
											πF.SetLineno(321)
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
									if πE = πClass.SetItem(πF, ß__getattr__.ToObject(), πTemp001); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("C").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µC = πTemp005
							// line 322: self.assertRaises(SyntaxError, operator.attrgetter('foo'), C())
							πF.SetLineno(322)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ßfoo.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[1] = πTemp002
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp002, πE = µC.Call(πF, nil, nil); πE != nil {
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
							// line 325: a = A()
							πF.SetLineno(325)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							if πTemp002, πE = µA.Call(πF, nil, nil); πE != nil {
								continue
							}
							µa = πTemp002
							// line 326: a.name = 'arthur'
							πF.SetLineno(326)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßarthur.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µa, ßname, πTemp002); πE != nil {
								continue
							}
							// line 327: a.child = A()
							πF.SetLineno(327)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							if πTemp002, πE = µA.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µa, ßchild, πTemp004); πE != nil {
								continue
							}
							// line 328: a.child.name = 'thomas'
							πF.SetLineno(328)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßthomas.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µa, ßchild, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp004, ßname, πTemp002); πE != nil {
								continue
							}
							// line 329: f = operator.attrgetter('child.name')
							πF.SetLineno(329)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("child.name").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µf = πTemp002
							// line 330: self.assertEqual(f(a), 'thomas')
							πF.SetLineno(330)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßthomas.ToObject()
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
							// line 331: self.assertRaises(AttributeError, f, a.child)
							πF.SetLineno(331)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[1] = µf
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µa, ßchild, nil); πE != nil {
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
							// line 332: f = operator.attrgetter('name', 'child.name')
							πF.SetLineno(332)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = ßname.ToObject()
							πTemp003[1] = πg.NewStr("child.name").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µf = πTemp002
							// line 333: self.assertEqual(f(a), ('arthur', 'thomas'))
							πF.SetLineno(333)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp002 = πg.NewTuple2(ßarthur.ToObject(), ßthomas.ToObject()).ToObject()
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
							// line 334: f = operator.attrgetter('name', 'child.name', 'child.child.name')
							πF.SetLineno(334)
							πTemp003 = πF.MakeArgs(3)
							πTemp003[0] = ßname.ToObject()
							πTemp003[1] = πg.NewStr("child.name").ToObject()
							πTemp003[2] = πg.NewStr("child.child.name").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µf = πTemp002
							// line 335: self.assertRaises(AttributeError, f, a)
							πF.SetLineno(335)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[1] = µf
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp003[2] = µa
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
							// line 336: f = operator.attrgetter('child.')
							πF.SetLineno(336)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("child.").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µf = πTemp002
							// line 337: self.assertRaises(AttributeError, f, a)
							πF.SetLineno(337)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[1] = µf
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp003[2] = µa
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
							// line 338: f = operator.attrgetter('.child')
							πF.SetLineno(338)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr(".child").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µf = πTemp002
							// line 339: self.assertRaises(AttributeError, f, a)
							πF.SetLineno(339)
							πTemp003 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp003[1] = µf
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp003[2] = µa
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
							// line 341: a.child.child = A()
							πF.SetLineno(341)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							if πTemp002, πE = µA.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µa, ßchild, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßchild, πTemp004); πE != nil {
								continue
							}
							// line 342: a.child.child.name = 'johnson'
							πF.SetLineno(342)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, ßjohnson.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µa, ßchild, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßchild, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp005, ßname, πTemp002); πE != nil {
								continue
							}
							// line 343: f = operator.attrgetter('child.child.name')
							πF.SetLineno(343)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("child.child.name").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µf = πTemp002
							// line 344: self.assertEqual(f(a), 'johnson')
							πF.SetLineno(344)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßjohnson.ToObject()
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
							// line 345: f = operator.attrgetter('name', 'child.name', 'child.child.name')
							πF.SetLineno(345)
							πTemp003 = πF.MakeArgs(3)
							πTemp003[0] = ßname.ToObject()
							πTemp003[1] = πg.NewStr("child.name").ToObject()
							πTemp003[2] = πg.NewStr("child.child.name").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßattrgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µf = πTemp002
							// line 346: self.assertEqual(f(a), ('arthur', 'thomas', 'johnson'))
							πF.SetLineno(346)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp002 = πg.NewTuple3(ßarthur.ToObject(), ßthomas.ToObject(), ßjohnson.ToObject()).ToObject()
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtest_attrgetter.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 297: @unittest.expectedFailure
					πF.SetLineno(297)
					πTemp014 = πF.MakeArgs(1)
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßtest_attrgetter); πE != nil {
						continue
					}
					πTemp014[0] = πTemp036
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp037, πE = πg.GetAttr(πF, πTemp036, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp036, πE = πTemp037.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πE = πClass.SetItem(πF, ßtest_attrgetter.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 349: def test_itemgetter(self):
					πF.SetLineno(349)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("test_itemgetter", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µf *πg.Object = πg.UnboundLocal; _ = µf
						var µC *πg.Object = πg.UnboundLocal; _ = µC
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µinventory *πg.Object = πg.UnboundLocal; _ = µinventory
						var µgetcount *πg.Object = πg.UnboundLocal; _ = µgetcount
						var µdata *πg.Object = πg.UnboundLocal; _ = µdata
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 πg.KWArgs
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
							// line 351: a = 'ABCDE'
							πF.SetLineno(351)
							µa = ßABCDE.ToObject()
							// line 352: f = operator.itemgetter(2)
							πF.SetLineno(352)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 353: self.assertEqual(f(a), 'C')
							πF.SetLineno(353)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp001[1] = ßC.ToObject()
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
							// line 354: f = operator.itemgetter(10)
							πF.SetLineno(354)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(10).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 355: self.assertRaises(IndexError, f, a)
							πF.SetLineno(355)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[1] = µf
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
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
							// line 357: class C(object):
							πF.SetLineno(357)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							πTemp005 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp005.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("C", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 358: def __getitem__(self, name):
									πF.SetLineno(358)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "name", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											if πTemp001, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
												continue
											}
											// line 359: raise SyntaxError
											πF.SetLineno(359)
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
							if πTemp006, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("C").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp005.ToObject()}, nil); πE != nil {
								continue
							}
							µC = πTemp006
							// line 360: self.assertRaises(SyntaxError, operator.itemgetter(42), C())
							πF.SetLineno(360)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSyntaxError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(42).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp002, πE = µC.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
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
							// line 362: f = operator.itemgetter('name')
							πF.SetLineno(362)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßname.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 363: self.assertRaises(TypeError, f, a)
							πF.SetLineno(363)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[1] = µf
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
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
							// line 364: self.assertRaises(TypeError, operator.itemgetter)
							πF.SetLineno(364)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
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
							// line 366: d = dict(key='val')
							πF.SetLineno(366)
							πTemp007 = πg.KWArgs{
								{"key", ßval.ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, πTemp007); πE != nil {
								continue
							}
							µd = πTemp003
							// line 367: f = operator.itemgetter('key')
							πF.SetLineno(367)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßkey.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 368: self.assertEqual(f(d), 'val')
							πF.SetLineno(368)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp004[0] = µd
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp002
							πTemp001[1] = ßval.ToObject()
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
							// line 369: f = operator.itemgetter('nonkey')
							πF.SetLineno(369)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßnonkey.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 370: self.assertRaises(KeyError, f, d)
							πF.SetLineno(370)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[1] = µf
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[2] = µd
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
							// line 373: inventory = [('apple', 3), ('banana', 2), ('pear', 5), ('orange', 1)]
							πF.SetLineno(373)
							πTemp001 = make([]*πg.Object, 4)
							πTemp002 = πg.NewTuple2(ßapple.ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewTuple2(ßbanana.ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp001[1] = πTemp002
							πTemp002 = πg.NewTuple2(ßpear.ToObject(), πg.NewInt(5).ToObject()).ToObject()
							πTemp001[2] = πTemp002
							πTemp002 = πg.NewTuple2(ßorange.ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp001[3] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µinventory = πTemp002
							// line 374: getcount = operator.itemgetter(1)
							πF.SetLineno(374)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µgetcount = πTemp002
							// line 375: self.assertEqual(list(map(getcount, inventory)), [3, 2, 5, 1])
							πF.SetLineno(375)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µgetcount, "getcount"); πE != nil {
								continue
							}
							πTemp008[0] = µgetcount
							if πE = πg.CheckLocal(πF, µinventory, "inventory"); πE != nil {
								continue
							}
							πTemp008[1] = µinventory
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp004 = make([]*πg.Object, 4)
							πTemp004[0] = πg.NewInt(3).ToObject()
							πTemp004[1] = πg.NewInt(2).ToObject()
							πTemp004[2] = πg.NewInt(5).ToObject()
							πTemp004[3] = πg.NewInt(1).ToObject()
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
							// line 376: self.assertEqual(sorted(inventory, key=getcount),
							πF.SetLineno(376)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µinventory, "inventory"); πE != nil {
								continue
							}
							πTemp004[0] = µinventory
							if πE = πg.CheckLocal(πF, µgetcount, "getcount"); πE != nil {
								continue
							}
							πTemp007 = πg.KWArgs{
								{"key", µgetcount},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp004 = make([]*πg.Object, 4)
							πTemp002 = πg.NewTuple2(ßorange.ToObject(), πg.NewInt(1).ToObject()).ToObject()
							πTemp004[0] = πTemp002
							πTemp002 = πg.NewTuple2(ßbanana.ToObject(), πg.NewInt(2).ToObject()).ToObject()
							πTemp004[1] = πTemp002
							πTemp002 = πg.NewTuple2(ßapple.ToObject(), πg.NewInt(3).ToObject()).ToObject()
							πTemp004[2] = πTemp002
							πTemp002 = πg.NewTuple2(ßpear.ToObject(), πg.NewInt(5).ToObject()).ToObject()
							πTemp004[3] = πTemp002
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
							// line 380: data = list(map(str, range(20)))
							πF.SetLineno(380)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewInt(20).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdata = πTemp003
							// line 381: self.assertEqual(operator.itemgetter(2,10,5)(data), ('2', '10', '5'))
							πF.SetLineno(381)
							πTemp001 = πF.MakeArgs(2)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp004[0] = µdata
							πTemp008 = πF.MakeArgs(3)
							πTemp008[0] = πg.NewInt(2).ToObject()
							πTemp008[1] = πg.NewInt(10).ToObject()
							πTemp008[2] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							πTemp002 = πg.NewTuple3(ß2.ToObject(), ß10.ToObject(), ß5.ToObject()).ToObject()
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
							// line 382: self.assertRaises(TypeError, operator.itemgetter(2, 'x', 5), data)
							πF.SetLineno(382)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp004 = πF.MakeArgs(3)
							πTemp004[0] = πg.NewInt(2).ToObject()
							πTemp004[1] = ßx.ToObject()
							πTemp004[2] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßitemgetter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[2] = µdata
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
					if πE = πClass.SetItem(πF, ßtest_itemgetter.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 348: @unittest.expectedFailure
					πF.SetLineno(348)
					πTemp014 = πF.MakeArgs(1)
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßtest_itemgetter); πE != nil {
						continue
					}
					πTemp014[0] = πTemp037
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp038, πE = πg.GetAttr(πF, πTemp037, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp037, πE = πTemp038.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πE = πClass.SetItem(πF, ßtest_itemgetter.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 384: def test_methodcaller(self):
					πF.SetLineno(384)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp037 = πg.NewFunction(πg.NewCode("test_methodcaller", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µA *πg.Object = πg.UnboundLocal; _ = µA
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µf *πg.Object = πg.UnboundLocal; _ = µf
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
						var πTemp007 πg.KWArgs
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 386: self.assertRaises(TypeError, operator.methodcaller)
							πF.SetLineno(386)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmethodcaller, nil); πE != nil {
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
							// line 387: class A(object):
							πF.SetLineno(387)
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
							_, πE = πg.NewCode("A", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 388: def foo(self, *args, **kwds):
									πF.SetLineno(388)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("foo", "/usr/lib/python2.7/test/test_operator.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µargs *πg.Object = πArgs[1]; _ = µargs
										var µkwds *πg.Object = πArgs[2]; _ = µkwds
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
											// line 389: return args[0] + args[1]
											πF.SetLineno(389)
											πTemp002 = πg.NewInt(0).ToObject()
											if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
												continue
											}
											πTemp002 = πg.NewInt(1).ToObject()
											if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
												continue
											}
											if πTemp004, πE = πg.GetItem(πF, µargs, πTemp002); πE != nil {
												continue
											}
											if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
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
									if πE = πClass.SetItem(πF, ßfoo.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 390: def bar(self, f=42):
									πF.SetLineno(390)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "f", Def: πg.NewInt(42).ToObject()}
									πTemp003 = πg.NewFunction(πg.NewCode("bar", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µf *πg.Object = πArgs[1]; _ = µf
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 391: return f
											πF.SetLineno(391)
											if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
												continue
											}
											πR = µf
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ßbar.ToObject(), πTemp003); πE != nil {
										continue
									}
									// line 392: def baz(*args, **kwds):
									πF.SetLineno(392)
									πTemp002 = make([]πg.Param, 0)
									πTemp004 = πg.NewFunction(πg.NewCode("baz", "/usr/lib/python2.7/test/test_operator.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µargs *πg.Object = πArgs[0]; _ = µargs
										var µkwds *πg.Object = πArgs[1]; _ = µkwds
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
											// line 393: return kwds['name'], kwds['self']
											πF.SetLineno(393)
											πTemp002 = ßname.ToObject()
											if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetItem(πF, µkwds, πTemp002); πE != nil {
												continue
											}
											πTemp002 = ßself.ToObject()
											if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
												continue
											}
											if πTemp004, πE = πg.GetItem(πF, µkwds, πTemp002); πE != nil {
												continue
											}
											πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
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
									if πE = πClass.SetItem(πF, ßbaz.ToObject(), πTemp004); πE != nil {
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
							if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("A").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
								continue
							}
							µA = πTemp005
							// line 394: a = A()
							πF.SetLineno(394)
							if πE = πg.CheckLocal(πF, µA, "A"); πE != nil {
								continue
							}
							if πTemp002, πE = µA.Call(πF, nil, nil); πE != nil {
								continue
							}
							µa = πTemp002
							// line 395: f = operator.methodcaller('foo')
							πF.SetLineno(395)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßfoo.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmethodcaller, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 396: self.assertRaises(IndexError, f, a)
							πF.SetLineno(396)
							πTemp001 = πF.MakeArgs(3)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[1] = µf
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
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
							// line 397: f = operator.methodcaller('foo', 1, 2)
							πF.SetLineno(397)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = ßfoo.ToObject()
							πTemp001[1] = πg.NewInt(1).ToObject()
							πTemp001[2] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmethodcaller, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 398: self.assertEqual(f(a), 3)
							πF.SetLineno(398)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp002
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
							// line 399: f = operator.methodcaller('bar')
							πF.SetLineno(399)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßbar.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmethodcaller, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 400: self.assertEqual(f(a), 42)
							πF.SetLineno(400)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(42).ToObject()
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
							// line 401: self.assertRaises(TypeError, f, a, a)
							πF.SetLineno(401)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp001[1] = µf
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[2] = µa
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp001[3] = µa
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
							// line 402: f = operator.methodcaller('bar', f=5)
							πF.SetLineno(402)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßbar.ToObject()
							πTemp007 = πg.KWArgs{
								{"f", πg.NewInt(5).ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmethodcaller, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 403: self.assertEqual(f(a), 5)
							πF.SetLineno(403)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(5).ToObject()
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
							// line 404: f = operator.methodcaller('baz', name='spam', self='eggs')
							πF.SetLineno(404)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßbaz.ToObject()
							πTemp007 = πg.KWArgs{
								{"name", ßspam.ToObject()},
								{"self", ßeggs.ToObject()},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmethodcaller, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp007); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							// line 405: self.assertEqual(f(a), ('spam', 'eggs'))
							πF.SetLineno(405)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp006[0] = µa
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewTuple2(ßspam.ToObject(), ßeggs.ToObject()).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_methodcaller.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 408: def test_inplace(self):
					πF.SetLineno(408)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp038 = πg.NewFunction(πg.NewCode("test_inplace", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µC *πg.Object = πg.UnboundLocal; _ = µC
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 410: class C(object):
							πF.SetLineno(410)
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
							_, πE = πg.NewCode("C", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 411: def __iadd__     (self, other): return "iadd"
									πF.SetLineno(411)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__iadd__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 411: def __iadd__     (self, other): return "iadd"
											πF.SetLineno(411)
											πR = ßiadd.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__iadd__.ToObject(), πTemp001); πE != nil {
										continue
									}
									// line 412: def __iand__     (self, other): return "iand"
									πF.SetLineno(412)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__iand__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 412: def __iand__     (self, other): return "iand"
											πF.SetLineno(412)
											πR = ßiand.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__iand__.ToObject(), πTemp003); πE != nil {
										continue
									}
									// line 413: def __ifloordiv__(self, other): return "ifloordiv"
									πF.SetLineno(413)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp004 = πg.NewFunction(πg.NewCode("__ifloordiv__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 413: def __ifloordiv__(self, other): return "ifloordiv"
											πF.SetLineno(413)
											πR = ßifloordiv.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__ifloordiv__.ToObject(), πTemp004); πE != nil {
										continue
									}
									// line 414: def __ilshift__  (self, other): return "ilshift"
									πF.SetLineno(414)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp005 = πg.NewFunction(πg.NewCode("__ilshift__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 414: def __ilshift__  (self, other): return "ilshift"
											πF.SetLineno(414)
											πR = ßilshift.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__ilshift__.ToObject(), πTemp005); πE != nil {
										continue
									}
									// line 415: def __imod__     (self, other): return "imod"
									πF.SetLineno(415)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp006 = πg.NewFunction(πg.NewCode("__imod__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 415: def __imod__     (self, other): return "imod"
											πF.SetLineno(415)
											πR = ßimod.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__imod__.ToObject(), πTemp006); πE != nil {
										continue
									}
									// line 416: def __imul__     (self, other): return "imul"
									πF.SetLineno(416)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp007 = πg.NewFunction(πg.NewCode("__imul__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 416: def __imul__     (self, other): return "imul"
											πF.SetLineno(416)
											πR = ßimul.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__imul__.ToObject(), πTemp007); πE != nil {
										continue
									}
									// line 417: def __ior__      (self, other): return "ior"
									πF.SetLineno(417)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp008 = πg.NewFunction(πg.NewCode("__ior__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 417: def __ior__      (self, other): return "ior"
											πF.SetLineno(417)
											πR = ßior.ToObject()
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
									// line 418: def __ipow__     (self, other): return "ipow"
									πF.SetLineno(418)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp009 = πg.NewFunction(πg.NewCode("__ipow__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 418: def __ipow__     (self, other): return "ipow"
											πF.SetLineno(418)
											πR = ßipow.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__ipow__.ToObject(), πTemp009); πE != nil {
										continue
									}
									// line 419: def __irshift__  (self, other): return "irshift"
									πF.SetLineno(419)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp010 = πg.NewFunction(πg.NewCode("__irshift__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 419: def __irshift__  (self, other): return "irshift"
											πF.SetLineno(419)
											πR = ßirshift.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__irshift__.ToObject(), πTemp010); πE != nil {
										continue
									}
									// line 420: def __isub__     (self, other): return "isub"
									πF.SetLineno(420)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp011 = πg.NewFunction(πg.NewCode("__isub__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 420: def __isub__     (self, other): return "isub"
											πF.SetLineno(420)
											πR = ßisub.ToObject()
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
									// line 421: def __itruediv__ (self, other): return "itruediv"
									πF.SetLineno(421)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp012 = πg.NewFunction(πg.NewCode("__itruediv__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 421: def __itruediv__ (self, other): return "itruediv"
											πF.SetLineno(421)
											πR = ßitruediv.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__itruediv__.ToObject(), πTemp012); πE != nil {
										continue
									}
									// line 422: def __ixor__     (self, other): return "ixor"
									πF.SetLineno(422)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp013 = πg.NewFunction(πg.NewCode("__ixor__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 422: def __ixor__     (self, other): return "ixor"
											πF.SetLineno(422)
											πR = ßixor.ToObject()
											continue
										}
										if πE != nil {
											πR = nil
										} else if πR == nil {
											πR = πg.None
										}
										return πR, πE
									}), πF.Globals()).ToObject()
									if πE = πClass.SetItem(πF, ß__ixor__.ToObject(), πTemp013); πE != nil {
										continue
									}
									// line 423: def __getitem__(self, other): return 5  # so that C is a sequence
									πF.SetLineno(423)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "other", Def: nil}
									πTemp014 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
										var µself *πg.Object = πArgs[0]; _ = µself
										var µother *πg.Object = πArgs[1]; _ = µother
										var πR *πg.Object; _ = πR
										var πE *πg.BaseException; _ = πE
										for ; πF.State() >= 0; πF.PopCheckpoint() {
											switch πF.State() {
											case 0:
											default: panic("unexpected function state")
											}
											// line 423: def __getitem__(self, other): return 5  # so that C is a sequence
											πF.SetLineno(423)
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
									if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp014); πE != nil {
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
							if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("C").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
								continue
							}
							µC = πTemp005
							// line 424: c = C()
							πF.SetLineno(424)
							if πE = πg.CheckLocal(πF, µC, "C"); πE != nil {
								continue
							}
							if πTemp002, πE = µC.Call(πF, nil, nil); πE != nil {
								continue
							}
							µc = πTemp002
							// line 425: self.assertEqual(operator.iadd     (c, 5), "iadd")
							πF.SetLineno(425)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßiadd, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßiadd.ToObject()
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
							// line 426: self.assertEqual(operator.iand     (c, 5), "iand")
							πF.SetLineno(426)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßiand, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßiand.ToObject()
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
							// line 427: self.assertEqual(operator.ifloordiv(c, 5), "ifloordiv")
							πF.SetLineno(427)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßifloordiv, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßifloordiv.ToObject()
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
							// line 428: self.assertEqual(operator.ilshift  (c, 5), "ilshift")
							πF.SetLineno(428)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßilshift, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßilshift.ToObject()
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
							// line 429: self.assertEqual(operator.imod     (c, 5), "imod")
							πF.SetLineno(429)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßimod, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßimod.ToObject()
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
							// line 430: self.assertEqual(operator.imul     (c, 5), "imul")
							πF.SetLineno(430)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßimul, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßimul.ToObject()
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
							// line 431: self.assertEqual(operator.ior      (c, 5), "ior")
							πF.SetLineno(431)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßior, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßior.ToObject()
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
							// line 432: self.assertEqual(operator.ipow     (c, 5), "ipow")
							πF.SetLineno(432)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßipow, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßipow.ToObject()
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
							// line 433: self.assertEqual(operator.irshift  (c, 5), "irshift")
							πF.SetLineno(433)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßirshift, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßirshift.ToObject()
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
							// line 434: self.assertEqual(operator.isub     (c, 5), "isub")
							πF.SetLineno(434)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßisub, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßisub.ToObject()
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
							// line 435: self.assertEqual(operator.itruediv (c, 5), "itruediv")
							πF.SetLineno(435)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßitruediv, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßitruediv.ToObject()
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
							// line 436: self.assertEqual(operator.ixor     (c, 5), "ixor")
							πF.SetLineno(436)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							πTemp006[1] = πg.NewInt(5).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßixor, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßixor.ToObject()
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
							// line 437: self.assertEqual(operator.iconcat  (c, c), "iadd")
							πF.SetLineno(437)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[1] = µc
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßiconcat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = ßiadd.ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_inplace.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 407: @unittest.expectedFailure
					πF.SetLineno(407)
					πTemp014 = πF.MakeArgs(1)
					if πTemp039, πE = πg.ResolveClass(πF, πClass, nil, ßtest_inplace); πE != nil {
						continue
					}
					πTemp014[0] = πTemp039
					if πTemp039, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp040, πE = πg.GetAttr(πF, πTemp039, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp039, πE = πTemp040.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πE = πClass.SetItem(πF, ßtest_inplace.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 440: def test_length_hint(self):
					πF.SetLineno(440)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp039 = πg.NewFunction(πg.NewCode("test_length_hint", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µX *πg.Object = πg.UnboundLocal; _ = µX
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
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
							// line 442: class X(object):
							πF.SetLineno(442)
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
							_, πE = πg.NewCode("X", "/usr/lib/python2.7/test/test_operator.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
									// line 443: def __init__(self, value):
									πF.SetLineno(443)
									πTemp002 = make([]πg.Param, 2)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp002[1] = πg.Param{Name: "value", Def: nil}
									πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											// line 444: self.value = value
											πF.SetLineno(444)
											if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
												continue
											}
											if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
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
									// line 446: def __length_hint__(self):
									πF.SetLineno(446)
									πTemp002 = make([]πg.Param, 1)
									πTemp002[0] = πg.Param{Name: "self", Def: nil}
									πTemp003 = πg.NewFunction(πg.NewCode("__length_hint__", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
											πTemp002 = πF.MakeArgs(1)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp003, πE = πg.GetAttr(πF, µself, ßvalue, nil); πE != nil {
												continue
											}
											πTemp002[0] = πTemp003
											if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
												continue
											}
											if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
												continue
											}
											πF.FreeArgs(πTemp002)
											if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
												continue
											}
											πTemp001 = πg.GetBool(πTemp004 == πTemp003).ToObject()
											if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
												continue
											}
											if πTemp005 {
												goto Label1
											}
											goto Label2
											// line 447: if type(self.value) is type:
											πF.SetLineno(447)
										Label1:
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ßvalue, nil); πE != nil {
												continue
											}
											// line 448: raise self.value
											πF.SetLineno(448)
											πE = πF.Raise(πTemp001, nil, nil)
											continue
											goto Label3
										Label2:
											// line 450: return self.value
											πF.SetLineno(450)
											if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
												continue
											}
											if πTemp001, πE = πg.GetAttr(πF, µself, ßvalue, nil); πE != nil {
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
									if πE = πClass.SetItem(πF, ß__length_hint__.ToObject(), πTemp003); πE != nil {
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
							// line 452: self.assertEqual(operator.length_hint([], 2), 0)
							πF.SetLineno(452)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							πTemp007 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp007...).ToObject()
							πTemp006[0] = πTemp002
							πTemp006[1] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßlength_hint, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(0).ToObject()
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
							// line 453: self.assertEqual(operator.length_hint(iter([1, 2, 3])), 3)
							πF.SetLineno(453)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = make([]*πg.Object, 3)
							πTemp008[0] = πg.NewInt(1).ToObject()
							πTemp008[1] = πg.NewInt(2).ToObject()
							πTemp008[2] = πg.NewInt(3).ToObject()
							πTemp002 = πg.NewList(πTemp008...).ToObject()
							πTemp007[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßlength_hint, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
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
							// line 455: self.assertEqual(operator.length_hint(X(2)), 2)
							πF.SetLineno(455)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp002, πE = µX.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßlength_hint, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
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
							// line 456: self.assertEqual(operator.length_hint(X(NotImplemented), 4), 4)
							πF.SetLineno(456)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp002, πE = µX.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp002
							πTemp006[1] = πg.NewInt(4).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßlength_hint, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
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
							// line 457: self.assertEqual(operator.length_hint(X(TypeError), 12), 12)
							πF.SetLineno(457)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp002, πE = µX.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp002
							πTemp006[1] = πg.NewInt(12).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßlength_hint, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp002
							πTemp003[1] = πg.NewInt(12).ToObject()
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
							// line 458: with self.assertRaises(TypeError):
							πF.SetLineno(458)
							πTemp003 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
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
							// line 459: operator.length_hint(X("abc"))
							πF.SetLineno(459)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ßabc.ToObject()
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp009, πE = µX.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßlength_hint, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp010.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.PopCheckpoint()
						Label1:
							πTemp012, πTemp013 = nil, nil
							if πE != nil {
								πTemp012, πTemp013 = πF.ExcInfo()
							}
							if πTemp012 != nil {
								πTemp014 = πTemp012.Type()
								if πTemp009, πE = πTemp002.Call(πF, πg.Args{πTemp004, πTemp014.ToObject(), πTemp012.ToObject(), πTemp013.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp009, πE = πTemp002.Call(πF, πg.Args{πTemp004, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if πTemp012 != nil && πTemp011 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 460: with self.assertRaises(ValueError):
							πF.SetLineno(460)
							πTemp003 = πF.MakeArgs(1)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp009
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp009, πE = πg.GetAttr(πF, πTemp010.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetAttr(πF, πTemp010.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πTemp015.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							// line 461: operator.length_hint(X(-2))
							πF.SetLineno(461)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πTemp016, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp006[0] = πTemp016
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp016, πE = µX.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp016
							if πTemp016, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßlength_hint, nil); πE != nil {
								continue
							}
							if πTemp016, πE = πTemp017.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.PopCheckpoint()
						Label2:
							πTemp012, πTemp013 = nil, nil
							if πE != nil {
								πTemp012, πTemp013 = πF.ExcInfo()
							}
							if πTemp012 != nil {
								πTemp014 = πTemp012.Type()
								if πTemp016, πE = πTemp009.Call(πF, πg.Args{πTemp010, πTemp014.ToObject(), πTemp012.ToObject(), πTemp013.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp016, πE = πTemp009.Call(πF, πg.Args{πTemp010, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp016); πE != nil {
								continue
							}
							if πTemp012 != nil && πTemp011 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							// line 462: with self.assertRaises(LookupError):
							πF.SetLineno(462)
							πTemp003 = πF.MakeArgs(1)
							if πTemp016, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
								continue
							}
							πTemp003[0] = πTemp016
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetAttr(πF, µself, ßassertRaises, nil); πE != nil {
								continue
							}
							if πTemp017, πE = πTemp016.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp016, πE = πg.GetAttr(πF, πTemp017.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp018, πE = πg.GetAttr(πF, πTemp017.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp018, πE = πTemp018.Call(πF, πg.Args{πTemp017}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							// line 463: operator.length_hint(X(LookupError))
							πF.SetLineno(463)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πTemp019, πE = πg.ResolveGlobal(πF, ßLookupError); πE != nil {
								continue
							}
							πTemp006[0] = πTemp019
							if πE = πg.CheckLocal(πF, µX, "X"); πE != nil {
								continue
							}
							if πTemp019, πE = µX.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp019
							if πTemp019, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßlength_hint, nil); πE != nil {
								continue
							}
							if πTemp019, πE = πTemp020.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.PopCheckpoint()
						Label3:
							πTemp012, πTemp013 = nil, nil
							if πE != nil {
								πTemp012, πTemp013 = πF.ExcInfo()
							}
							if πTemp012 != nil {
								πTemp014 = πTemp012.Type()
								if πTemp019, πE = πTemp016.Call(πF, πg.Args{πTemp017, πTemp014.ToObject(), πTemp012.ToObject(), πTemp013.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp019, πE = πTemp016.Call(πF, πg.Args{πTemp017, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp019); πE != nil {
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
					if πE = πClass.SetItem(πF, ßtest_length_hint.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 439: @unittest.expectedFailure
					πF.SetLineno(439)
					πTemp014 = πF.MakeArgs(1)
					if πTemp040, πE = πg.ResolveClass(πF, πClass, nil, ßtest_length_hint); πE != nil {
						continue
					}
					πTemp014[0] = πTemp040
					if πTemp040, πE = πg.ResolveClass(πF, πClass, nil, ßunittest); πE != nil {
						continue
					}
					if πTemp041, πE = πg.GetAttr(πF, πTemp040, ßexpectedFailure, nil); πE != nil {
						continue
					}
					if πTemp040, πE = πTemp041.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πE = πClass.SetItem(πF, ßtest_length_hint.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 465: def test_dunder_is_original(self):
					πF.SetLineno(465)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp040 = πg.NewFunction(πg.NewCode("test_dunder_is_original", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnames *πg.Object = πg.UnboundLocal; _ = µnames
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var µorig *πg.Object = πg.UnboundLocal; _ = µorig
						var µdunder *πg.Object = πg.UnboundLocal; _ = µdunder
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
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
						var πTemp009 []*πg.Object
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
							default: panic("unexpected function state")
							}
							// line 468: names = [name for name in dir(operator) if not name.startswith('_')]
							πF.SetLineno(468)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/test/test_operator.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µname *πg.Object = πg.UnboundLocal; _ = µname
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
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 6: goto Label6
										default: panic("unexpected function state")
										}
										πTemp002 = πF.MakeArgs(1)
										if πTemp003, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
											continue
										}
										πTemp002[0] = πTemp003
										if πTemp003, πE = πg.ResolveGlobal(πF, ßdir); πE != nil {
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
											µname = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)            
										πTemp002 = πF.MakeArgs(1)
										πTemp002[0] = ß_.ToObject()
										if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µname, ßstartswith, nil); πE != nil {
											continue
										}
										if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
										if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
											continue
										}
										πTemp003 = πg.GetBool(!πTemp006).ToObject()
										if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
											continue
										}
										if πTemp006 {
											goto Label4
										}
										goto Label5
										// line 468: names = [name for name in dir(operator) if not name.startswith('_')]
										πF.SetLineno(468)
									Label4:
										// line 468: names = [name for name in dir(operator) if not name.startswith('_')]
										πF.SetLineno(468)
										if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µname, nil
									Label6:
										πTemp003 = πSent
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
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µnames = πTemp001
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µnames); πE != nil {
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
							if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µname = πTemp004
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 470: orig = getattr(operator, name)
							πF.SetLineno(470)
							πTemp007 = πF.MakeArgs(2)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp007[1] = µname
							if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µorig = πTemp008
							// line 471: dunder = getattr(operator, '__' + name.strip('_') + '__', None)
							πF.SetLineno(471)
							πTemp007 = πF.MakeArgs(3)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							πTemp009 = πF.MakeArgs(1)
							πTemp009[0] = ß_.ToObject()
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µname, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp008, πE = πg.Add(πF, ß__.ToObject(), πTemp011); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp008, ß__.ToObject()); πE != nil {
								continue
							}
							πTemp007[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp007[2] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µdunder = πTemp008
							if πE = πg.CheckLocal(πF, µdunder, "dunder"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µdunder); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 472: if dunder:
							πF.SetLineno(472)
						Label4:
							// line 473: self.assertIs(dunder, orig)
							πF.SetLineno(473)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdunder, "dunder"); πE != nil {
								continue
							}
							πTemp007[0] = µdunder
							if πE = πg.CheckLocal(πF, µorig, "orig"); πE != nil {
								continue
							}
							πTemp007[1] = µorig
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßassertIs, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßtest_dunder_is_original.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 475: def test_complex_operator(self):
					πF.SetLineno(475)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp041 = πg.NewFunction(πg.NewCode("test_complex_operator", "/usr/lib/python2.7/test/test_operator.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 476: self.assertRaises(TypeError, operator.lt, 1j, 2j)
							πF.SetLineno(476)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlt, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewComplex(complex(0.0, 1.0)).ToObject()
							πTemp001[3] = πg.NewComplex(complex(0.0, 2.0)).ToObject()
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
							// line 477: self.assertRaises(TypeError, operator.le, 1j, 2j)
							πF.SetLineno(477)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßle, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewComplex(complex(0.0, 1.0)).ToObject()
							πTemp001[3] = πg.NewComplex(complex(0.0, 2.0)).ToObject()
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
							// line 478: self.assertRaises(TypeError, operator.ge, 1j, 2j)
							πF.SetLineno(478)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßge, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewComplex(complex(0.0, 1.0)).ToObject()
							πTemp001[3] = πg.NewComplex(complex(0.0, 2.0)).ToObject()
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
							// line 479: self.assertRaises(TypeError, operator.gt, 1j, 2j)
							πF.SetLineno(479)
							πTemp001 = πF.MakeArgs(4)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgt, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							πTemp001[2] = πg.NewComplex(complex(0.0, 1.0)).ToObject()
							πTemp001[3] = πg.NewComplex(complex(0.0, 2.0)).ToObject()
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
					if πE = πClass.SetItem(πF, ßtest_complex_operator.ToObject(), πTemp041); πE != nil {
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
			if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("OperatorTestCase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOperatorTestCase.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 480: def test_main():
			πF.SetLineno(480)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("test_main", "/usr/lib/python2.7/test/test_operator.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 481: test_support.run_unittest(OperatorTestCase)
					πF.SetLineno(481)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOperatorTestCase); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
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
			// line 483: if __name__ == "__main__":
			πF.SetLineno(483)
		Label1:
			// line 484: test_main()
			πF.SetLineno(484)
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
	πg.RegisterModule("test.test_operator", Code)
}

