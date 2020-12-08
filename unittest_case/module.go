package unittest_case
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/warnings"
	// _ "github.com/pygolin/stdlib/unittest_util"
	// _ "github.com/pygolin/stdlib/functools"
	// _ "github.com/pygolin/stdlib/pprint"
	// _ "github.com/pygolin/stdlib/difflib"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/re"
	// _ "github.com/pygolin/stdlib/collections"
	// _ "github.com/pygolin/stdlib/unittest_result"
	// _ "github.com/pygolin/stdlib/types"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/unittest_case.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßAssertionError := πg.InternStr("AssertionError")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBaseException := πg.InternStr("BaseException")
		ßClassType := πg.InternStr("ClassType")
		ßCounter := πg.InternStr("Counter")
		ßDIFF_OMITTED := πg.InternStr("DIFF_OMITTED")
		ßDeprecationWarning := πg.InternStr("DeprecationWarning")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßFunctionTestCase := πg.InternStr("FunctionTestCase")
		ßIndexError := πg.InternStr("IndexError")
		ßKeyboardInterrupt := πg.InternStr("KeyboardInterrupt")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßNotImplemented := πg.InternStr("NotImplemented")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßPendingDeprecationWarning := πg.InternStr("PendingDeprecationWarning")
		ßRuntimeWarning := πg.InternStr("RuntimeWarning")
		ßSkipTest := πg.InternStr("SkipTest")
		ßTestCase := πg.InternStr("TestCase")
		ßTestResult := πg.InternStr("TestResult")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßUnicodeDecodeError := πg.InternStr("UnicodeDecodeError")
		ßValueError := πg.InternStr("ValueError")
		ß_AssertRaisesContext := πg.InternStr("_AssertRaisesContext")
		ß_ExpectedFailure := πg.InternStr("_ExpectedFailure")
		ß_UnexpectedSuccess := πg.InternStr("_UnexpectedSuccess")
		ß__call__ := πg.InternStr("__call__")
		ß__class__ := πg.InternStr("__class__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__enter__ := πg.InternStr("__enter__")
		ß__eq__ := πg.InternStr("__eq__")
		ß__exit__ := πg.InternStr("__exit__")
		ß__hash__ := πg.InternStr("__hash__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__ne__ := πg.InternStr("__ne__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__str__ := πg.InternStr("__str__")
		ß__unittest := πg.InternStr("__unittest")
		ß__unittest_skip__ := πg.InternStr("__unittest_skip__")
		ß__unittest_skip_why__ := πg.InternStr("__unittest_skip_why__")
		ß_addSkip := πg.InternStr("_addSkip")
		ß_baseAssertEqual := πg.InternStr("_baseAssertEqual")
		ß_classSetupFailed := πg.InternStr("_classSetupFailed")
		ß_cleanups := πg.InternStr("_cleanups")
		ß_count_diff_all_purpose := πg.InternStr("_count_diff_all_purpose")
		ß_count_diff_hashable := πg.InternStr("_count_diff_hashable")
		ß_deprecate := πg.InternStr("_deprecate")
		ß_description := πg.InternStr("_description")
		ß_diffThreshold := πg.InternStr("_diffThreshold")
		ß_formatMessage := πg.InternStr("_formatMessage")
		ß_getAssertEqualityFunc := πg.InternStr("_getAssertEqualityFunc")
		ß_id := πg.InternStr("_id")
		ß_resultForDoCleanups := πg.InternStr("_resultForDoCleanups")
		ß_setUpFunc := πg.InternStr("_setUpFunc")
		ß_tearDownFunc := πg.InternStr("_tearDownFunc")
		ß_testFunc := πg.InternStr("_testFunc")
		ß_testMethodName := πg.InternStr("_testMethodName")
		ß_truncateMessage := πg.InternStr("_truncateMessage")
		ß_type_equality_funcs := πg.InternStr("_type_equality_funcs")
		ß_util := πg.InternStr("_util")
		ßabs := πg.InternStr("abs")
		ßaddCleanup := πg.InternStr("addCleanup")
		ßaddError := πg.InternStr("addError")
		ßaddExpectedFailure := πg.InternStr("addExpectedFailure")
		ßaddFailure := πg.InternStr("addFailure")
		ßaddSkip := πg.InternStr("addSkip")
		ßaddSuccess := πg.InternStr("addSuccess")
		ßaddTypeEqualityFunc := πg.InternStr("addTypeEqualityFunc")
		ßaddUnexpectedSuccess := πg.InternStr("addUnexpectedSuccess")
		ßappend := πg.InternStr("append")
		ßassertAlmostEqual := πg.InternStr("assertAlmostEqual")
		ßassertAlmostEquals := πg.InternStr("assertAlmostEquals")
		ßassertDictContainsSubset := πg.InternStr("assertDictContainsSubset")
		ßassertDictEqual := πg.InternStr("assertDictEqual")
		ßassertEqual := πg.InternStr("assertEqual")
		ßassertEquals := πg.InternStr("assertEquals")
		ßassertFalse := πg.InternStr("assertFalse")
		ßassertGreater := πg.InternStr("assertGreater")
		ßassertGreaterEqual := πg.InternStr("assertGreaterEqual")
		ßassertIn := πg.InternStr("assertIn")
		ßassertIs := πg.InternStr("assertIs")
		ßassertIsInstance := πg.InternStr("assertIsInstance")
		ßassertIsNone := πg.InternStr("assertIsNone")
		ßassertIsNot := πg.InternStr("assertIsNot")
		ßassertIsNotNone := πg.InternStr("assertIsNotNone")
		ßassertItemsEqual := πg.InternStr("assertItemsEqual")
		ßassertLess := πg.InternStr("assertLess")
		ßassertLessEqual := πg.InternStr("assertLessEqual")
		ßassertListEqual := πg.InternStr("assertListEqual")
		ßassertMultiLineEqual := πg.InternStr("assertMultiLineEqual")
		ßassertNotAlmostEqual := πg.InternStr("assertNotAlmostEqual")
		ßassertNotAlmostEquals := πg.InternStr("assertNotAlmostEquals")
		ßassertNotEqual := πg.InternStr("assertNotEqual")
		ßassertNotEquals := πg.InternStr("assertNotEquals")
		ßassertNotIn := πg.InternStr("assertNotIn")
		ßassertNotIsInstance := πg.InternStr("assertNotIsInstance")
		ßassertNotRegexpMatches := πg.InternStr("assertNotRegexpMatches")
		ßassertRaises := πg.InternStr("assertRaises")
		ßassertRaisesRegexp := πg.InternStr("assertRaisesRegexp")
		ßassertRegexpMatches := πg.InternStr("assertRegexpMatches")
		ßassertSequenceEqual := πg.InternStr("assertSequenceEqual")
		ßassertSetEqual := πg.InternStr("assertSetEqual")
		ßassertTrue := πg.InternStr("assertTrue")
		ßassertTupleEqual := πg.InternStr("assertTupleEqual")
		ßassert_ := πg.InternStr("assert_")
		ßbasestring := πg.InternStr("basestring")
		ßcapitalize := πg.InternStr("capitalize")
		ßcatch_warnings := πg.InternStr("catch_warnings")
		ßclassmethod := πg.InternStr("classmethod")
		ßcollections := πg.InternStr("collections")
		ßcompile := πg.InternStr("compile")
		ßcountTestCases := πg.InternStr("countTestCases")
		ßdebug := πg.InternStr("debug")
		ßdefaultTestResult := πg.InternStr("defaultTestResult")
		ßdict := πg.InternStr("dict")
		ßdifference := πg.InternStr("difference")
		ßdifflib := πg.InternStr("difflib")
		ßdoCleanups := πg.InternStr("doCleanups")
		ßdoc := πg.InternStr("doc")
		ßend := πg.InternStr("end")
		ßexc_info := πg.InternStr("exc_info")
		ßexception := πg.InternStr("exception")
		ßexpected := πg.InternStr("expected")
		ßexpectedFailure := πg.InternStr("expectedFailure")
		ßexpected_regexp := πg.InternStr("expected_regexp")
		ßfail := πg.InternStr("fail")
		ßfailIf := πg.InternStr("failIf")
		ßfailIfAlmostEqual := πg.InternStr("failIfAlmostEqual")
		ßfailIfEqual := πg.InternStr("failIfEqual")
		ßfailUnless := πg.InternStr("failUnless")
		ßfailUnlessAlmostEqual := πg.InternStr("failUnlessAlmostEqual")
		ßfailUnlessEqual := πg.InternStr("failUnlessEqual")
		ßfailUnlessRaises := πg.InternStr("failUnlessRaises")
		ßfailureException := πg.InternStr("failureException")
		ßfilterwarnings := πg.InternStr("filterwarnings")
		ßfrozenset := πg.InternStr("frozenset")
		ßfunctools := πg.InternStr("functools")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßhash := πg.InternStr("hash")
		ßid := πg.InternStr("id")
		ßignore := πg.InternStr("ignore")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßiteritems := πg.InternStr("iteritems")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßlongMessage := πg.InternStr("longMessage")
		ßmaxDiff := πg.InternStr("maxDiff")
		ßmin := πg.InternStr("min")
		ßndiff := πg.InternStr("ndiff")
		ßobject := πg.InternStr("object")
		ßpattern := πg.InternStr("pattern")
		ßpformat := πg.InternStr("pformat")
		ßpop := πg.InternStr("pop")
		ßpprint := πg.InternStr("pprint")
		ßpy3kwarning := πg.InternStr("py3kwarning")
		ßre := πg.InternStr("re")
		ßrepr := πg.InternStr("repr")
		ßresult := πg.InternStr("result")
		ßround := πg.InternStr("round")
		ßrun := πg.InternStr("run")
		ßrunTest := πg.InternStr("runTest")
		ßsafe_repr := πg.InternStr("safe_repr")
		ßsearch := πg.InternStr("search")
		ßsequence := πg.InternStr("sequence")
		ßset := πg.InternStr("set")
		ßsetUp := πg.InternStr("setUp")
		ßsetUpClass := πg.InternStr("setUpClass")
		ßshortDescription := πg.InternStr("shortDescription")
		ßskip := πg.InternStr("skip")
		ßskipIf := πg.InternStr("skipIf")
		ßskipTest := πg.InternStr("skipTest")
		ßskipUnless := πg.InternStr("skipUnless")
		ßsplit := πg.InternStr("split")
		ßsplitlines := πg.InternStr("splitlines")
		ßstart := πg.InternStr("start")
		ßstartTest := πg.InternStr("startTest")
		ßstartTestRun := πg.InternStr("startTestRun")
		ßstopTest := πg.InternStr("stopTest")
		ßstopTestRun := πg.InternStr("stopTestRun")
		ßstr := πg.InternStr("str")
		ßstrclass := πg.InternStr("strclass")
		ßstrip := πg.InternStr("strip")
		ßsuper := πg.InternStr("super")
		ßsys := πg.InternStr("sys")
		ßtearDown := πg.InternStr("tearDown")
		ßtearDownClass := πg.InternStr("tearDownClass")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		ßunicode := πg.InternStr("unicode")
		ßunorderable_list_difference := πg.InternStr("unorderable_list_difference")
		ßwarn := πg.InternStr("warn")
		ßwarnings := πg.InternStr("warnings")
		ßwraps := πg.InternStr("wraps")
		ßxrange := πg.InternStr("xrange")
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
		var πTemp009 *πg.Dict
		_ = πTemp009
		var πTemp010 []πg.Param
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: """Test case implementation"""
			πF.SetLineno(1)
			// line 1: """Test case implementation"""
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Test case implementation").ToObject()); πE != nil {
				continue
			}
			// line 3: import collections
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "collections"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcollections.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 4: import sys
			πF.SetLineno(4)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: import functools
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßfunctools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: import difflib
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "difflib"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßdifflib.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import pprint
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "pprint"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßpprint.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import re
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 9: import types
			πF.SetLineno(9)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 10: import warnings
			πF.SetLineno(10)
			if πTemp002, πE = πg.ImportModule(πF, "warnings"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßwarnings.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 13: import unittest_result as result
			πF.SetLineno(13)
			if πTemp002, πE = πg.ImportModule(πF, "unittest_result"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßresult.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: import unittest_util as _util
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "unittest_util"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_util.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: strclass, safe_repr, unorderable_list_difference, _count_diff_all_purpose, \
			πF.SetLineno(19)
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_util); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstrclass, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_util); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßsafe_repr, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_util); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßunorderable_list_difference, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_util); πE != nil {
				continue
			}
			if πTemp007, πE = πg.GetAttr(πF, πTemp003, ß_count_diff_all_purpose, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_util); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp003, ß_count_diff_hashable, nil); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple5(πTemp004, πTemp005, πTemp006, πTemp007, πTemp008).ToObject()
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp001); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßstrclass.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßsafe_repr.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßunorderable_list_difference.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_count_diff_all_purpose.ToObject(), πTemp006); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_count_diff_hashable.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 24: class KeyboardInterrupt(BaseException):
			πF.SetLineno(24)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßBaseException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp009 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("KeyboardInterrupt", "/usr/lib/python2.7/unittest_case.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 25: pass
					πF.SetLineno(25)
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("KeyboardInterrupt").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßKeyboardInterrupt.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 28: __unittest = True
			πF.SetLineno(28)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__unittest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 31: DIFF_OMITTED = ('\nDiff is %s characters long. '
			πF.SetLineno(31)
			if πE = πF.Globals().SetItem(πF, ßDIFF_OMITTED.ToObject(), πg.NewStr("\nDiff is %s characters long. Set self.maxDiff to None to see it.").ToObject()); πE != nil {
				continue
			}
			// line 34: class SkipTest(Exception):
			πF.SetLineno(34)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp009 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SkipTest", "/usr/lib/python2.7/unittest_case.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 35: """
					πF.SetLineno(35)
					// line 35: """
					πF.SetLineno(35)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raise this exception in a test to skip it.\n\n    Usually you can use TestCase.skipTest() or one of the skipping decorators\n    instead of raising this directly.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 41: pass
					πF.SetLineno(41)
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SkipTest").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSkipTest.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 43: class _ExpectedFailure(Exception):
			πF.SetLineno(43)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp009 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_ExpectedFailure", "/usr/lib/python2.7/unittest_case.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 44: """
					πF.SetLineno(44)
					// line 44: """
					πF.SetLineno(44)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raise this when a test is expected to fail.\n\n    This is an implementation detail.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 50: def __init__(self, exc_info):
					πF.SetLineno(50)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "exc_info", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexc_info *πg.Object = πArgs[1]; _ = µexc_info
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
							// line 51: super(_ExpectedFailure, self).__init__()
							πF.SetLineno(51)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_ExpectedFailure); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[1] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 52: self.exc_info = exc_info
							πF.SetLineno(52)
							if πE = πg.CheckLocal(πF, µexc_info, "exc_info"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µexc_info); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßexc_info, πTemp002); πE != nil {
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
			if πTemp003, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("_ExpectedFailure").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_ExpectedFailure.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 54: class _UnexpectedSuccess(Exception):
			πF.SetLineno(54)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp009 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_UnexpectedSuccess", "/usr/lib/python2.7/unittest_case.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp009
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 55: """
					πF.SetLineno(55)
					// line 55: """
					πF.SetLineno(55)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    The test was supposed to fail, but it didn't!\n    ").ToObject()); πE != nil {
						continue
					}
					// line 58: pass
					πF.SetLineno(58)
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("_UnexpectedSuccess").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_UnexpectedSuccess.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 60: def _id(obj):
			πF.SetLineno(60)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "obj", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_id", "/usr/lib/python2.7/unittest_case.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 61: return obj
					πF.SetLineno(61)
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
			if πE = πF.Globals().SetItem(πF, ß_id.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 63: def skip(reason):
			πF.SetLineno(63)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "reason", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("skip", "/usr/lib/python2.7/unittest_case.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µreason *πg.Object = πArgs[0]; _ = µreason
				var µdecorator *πg.Object = πg.UnboundLocal; _ = µdecorator
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 64: """
					πF.SetLineno(64)
					// line 67: def decorator(test_item):
					πF.SetLineno(67)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "test_item", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("decorator", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µtest_item *πg.Object = πArgs[0]; _ = µtest_item
						var µskip_wrapper *πg.Object = πg.UnboundLocal; _ = µskip_wrapper
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
						var πTemp007 bool
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
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtest_item, "test_item"); πE != nil {
								continue
							}
							πTemp002[0] = µtest_item
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßClassType, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 68: if not isinstance(test_item, (type, types.ClassType)):
							πF.SetLineno(68)
						Label1:
							// line 70: def skip_wrapper(*args, **kwargs):
							πF.SetLineno(70)
							πTemp008 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("skip_wrapper", "/usr/lib/python2.7/unittest_case.py", πTemp008, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µargs *πg.Object = πArgs[0]; _ = µargs
								var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
									if πE = πg.CheckLocal(πF, µreason, "reason"); πE != nil {
										continue
									}
									πTemp001[0] = µreason
									if πTemp002, πE = πg.ResolveGlobal(πF, ßSkipTest); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 71: raise SkipTest(reason)
									πF.SetLineno(71)
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
							µskip_wrapper = πTemp001
							// line 72: skip_wrapper = functools.wraps(test_item)(skip_wrapper)
							πF.SetLineno(72)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µskip_wrapper, "skip_wrapper"); πE != nil {
								continue
							}
							πTemp002[0] = µskip_wrapper
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtest_item, "test_item"); πE != nil {
								continue
							}
							πTemp009[0] = µtest_item
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwraps, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µskip_wrapper = πTemp004
							// line 73: test_item = skip_wrapper
							πF.SetLineno(73)
							if πE = πg.CheckLocal(πF, µskip_wrapper, "skip_wrapper"); πE != nil {
								continue
							}
							µtest_item = µskip_wrapper
							goto Label2
						Label2:
							// line 75: test_item.__unittest_skip__ = True
							πF.SetLineno(75)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtest_item, "test_item"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtest_item, ß__unittest_skip__, πTemp004); πE != nil {
								continue
							}
							// line 76: test_item.__unittest_skip_why__ = reason
							πF.SetLineno(76)
							if πE = πg.CheckLocal(πF, µreason, "reason"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µreason); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtest_item, "test_item"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µtest_item, ß__unittest_skip_why__, πTemp003); πE != nil {
								continue
							}
							// line 77: return test_item
							πF.SetLineno(77)
							if πE = πg.CheckLocal(πF, µtest_item, "test_item"); πE != nil {
								continue
							}
							πR = µtest_item
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					µdecorator = πTemp001
					// line 78: return decorator
					πF.SetLineno(78)
					if πE = πg.CheckLocal(πF, µdecorator, "decorator"); πE != nil {
						continue
					}
					πR = µdecorator
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßskip.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 64: """
			πF.SetLineno(64)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n    Unconditionally skip a test.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßskip); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 80: def skipIf(condition, reason):
			πF.SetLineno(80)
			πTemp010 = make([]πg.Param, 2)
			πTemp010[0] = πg.Param{Name: "condition", Def: nil}
			πTemp010[1] = πg.Param{Name: "reason", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("skipIf", "/usr/lib/python2.7/unittest_case.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcondition *πg.Object = πArgs[0]; _ = µcondition
				var µreason *πg.Object = πArgs[1]; _ = µreason
				var πTemp001 bool
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
					// line 81: """
					πF.SetLineno(81)
					if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IsTrue(πF, µcondition); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 84: if condition:
					πF.SetLineno(84)
				Label1:
					// line 85: return skip(reason)
					πF.SetLineno(85)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µreason, "reason"); πE != nil {
						continue
					}
					πTemp002[0] = µreason
					if πTemp003, πE = πg.ResolveGlobal(πF, ßskip); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp004
					continue
					goto Label2
				Label2:
					// line 86: return _id
					πF.SetLineno(86)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_id); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßskipIf.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 81: """
			πF.SetLineno(81)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n    Skip a test if the condition is true.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßskipIf); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 88: def skipUnless(condition, reason):
			πF.SetLineno(88)
			πTemp010 = make([]πg.Param, 2)
			πTemp010[0] = πg.Param{Name: "condition", Def: nil}
			πTemp010[1] = πg.Param{Name: "reason", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("skipUnless", "/usr/lib/python2.7/unittest_case.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcondition *πg.Object = πArgs[0]; _ = µcondition
				var µreason *πg.Object = πArgs[1]; _ = µreason
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
					// line 89: """
					πF.SetLineno(89)
					if πE = πg.CheckLocal(πF, µcondition, "condition"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µcondition); πE != nil {
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
					// line 92: if not condition:
					πF.SetLineno(92)
				Label1:
					// line 93: return skip(reason)
					πF.SetLineno(93)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µreason, "reason"); πE != nil {
						continue
					}
					πTemp003[0] = µreason
					if πTemp001, πE = πg.ResolveGlobal(πF, ßskip); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label2
				Label2:
					// line 94: return _id
					πF.SetLineno(94)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_id); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßskipUnless.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 89: """
			πF.SetLineno(89)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n    Skip a test unless the condition is true.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßskipUnless); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 97: def expectedFailure(func):
			πF.SetLineno(97)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "func", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("expectedFailure", "/usr/lib/python2.7/unittest_case.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]; _ = µfunc
				var µwrapper *πg.Object = πg.UnboundLocal; _ = µwrapper
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
					// line 99: def wrapper(*args, **kwargs):
					πF.SetLineno(99)
					πTemp002 = make([]πg.Param, 0)
					πTemp001 = πg.NewFunction(πg.NewCode("wrapper", "/usr/lib/python2.7/unittest_case.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µargs *πg.Object = πArgs[0]; _ = µargs
						var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.BaseException
						_ = πTemp002
						var πTemp003 *πg.Traceback
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
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 100: try:
							πF.SetLineno(100)
							πF.PushCheckpoint(2)
							// line 101: func(*args, **kwargs)
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, µfunc, nil, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp002, πTemp003 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp002.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							πE = πF.Raise(πTemp002.ToObject(), nil, πTemp003.ToObject())
							continue
							// line 102: except Exception:
							πF.SetLineno(102)
						Label3:
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_ExpectedFailure); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 103: raise _ExpectedFailure(sys.exc_info())
							πF.SetLineno(103)
							πE = πF.Raise(πTemp006, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_UnexpectedSuccess); πE != nil {
								continue
							}
							// line 104: raise _UnexpectedSuccess
							πF.SetLineno(104)
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
					µwrapper = πTemp001
					// line 105: wrapper = functools.wraps(func)(wrapper)
					πF.SetLineno(105)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µwrapper, "wrapper"); πE != nil {
						continue
					}
					πTemp003[0] = µwrapper
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp004[0] = µfunc
					if πTemp005, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßwraps, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µwrapper = πTemp006
					// line 106: return wrapper
					πF.SetLineno(106)
					if πE = πg.CheckLocal(πF, µwrapper, "wrapper"); πE != nil {
						continue
					}
					πR = µwrapper
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßexpectedFailure.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 109: class _AssertRaisesContext(object):
			πF.SetLineno(109)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp011
			πTemp009 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_AssertRaisesContext", "/usr/lib/python2.7/unittest_case.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 110: """A context manager used to implement TestCase.assertRaises* methods."""
					πF.SetLineno(110)
					// line 110: """A context manager used to implement TestCase.assertRaises* methods."""
					πF.SetLineno(110)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A context manager used to implement TestCase.assertRaises* methods.").ToObject()); πE != nil {
						continue
					}
					// line 112: def __init__(self, expected, test_case, expected_regexp=None):
					πF.SetLineno(112)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "expected", Def: nil}
					πTemp002[2] = πg.Param{Name: "test_case", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "expected_regexp", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpected *πg.Object = πArgs[1]; _ = µexpected
						var µtest_case *πg.Object = πArgs[2]; _ = µtest_case
						var µexpected_regexp *πg.Object = πArgs[3]; _ = µexpected_regexp
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
							// line 113: self.expected = expected
							πF.SetLineno(113)
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µexpected); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßexpected, πTemp001); πE != nil {
								continue
							}
							// line 114: self.failureException = test_case.failureException
							πF.SetLineno(114)
							if πE = πg.CheckLocal(πF, µtest_case, "test_case"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µtest_case, ßfailureException, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfailureException, πTemp002); πE != nil {
								continue
							}
							// line 115: self.expected_regexp = expected_regexp
							πF.SetLineno(115)
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µexpected_regexp); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßexpected_regexp, πTemp001); πE != nil {
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
					// line 117: def __enter__(self):
					πF.SetLineno(117)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__enter__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 118: return self
							πF.SetLineno(118)
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
					if πE = πClass.SetItem(πF, ß__enter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 120: def __exit__(self, exc_type, exc_value, tb):
					πF.SetLineno(120)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "exc_type", Def: nil}
					πTemp002[2] = πg.Param{Name: "exc_value", Def: nil}
					πTemp002[3] = πg.Param{Name: "tb", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__exit__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexc_type *πg.Object = πArgs[1]; _ = µexc_type
						var µexc_value *πg.Object = πArgs[2]; _ = µexc_value
						var µtb *πg.Object = πArgs[3]; _ = µtb
						var µexc_name *πg.Object = πg.UnboundLocal; _ = µexc_name
						var µexpected_regexp *πg.Object = πg.UnboundLocal; _ = µexpected_regexp
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
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
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µexc_type, "exc_type"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µexc_type == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 121: if exc_type is None:
							πF.SetLineno(121)
						Label1:
							// line 122: try:
							πF.SetLineno(122)
							πF.PushCheckpoint(4)
							// line 123: exc_name = self.expected.__name__
							πF.SetLineno(123)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßexpected, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__name__, nil); πE != nil {
								continue
							}
							µexc_name = πTemp002
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 124: except AttributeError:
							πF.SetLineno(124)
						Label5:
							// line 125: exc_name = str(self.expected)
							πF.SetLineno(125)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßexpected, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µexc_name = πTemp002
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexc_name, "exc_name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s not raised").ToObject(), µexc_name); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 126: raise self.failureException(
							πF.SetLineno(126)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µexc_type, "exc_type"); πE != nil {
								continue
							}
							πTemp006[0] = µexc_type
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßexpected, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 129: if not issubclass(exc_type, self.expected):
							πF.SetLineno(129)
						Label6:
							// line 131: return False
							πF.SetLineno(131)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label7
						Label7:
							// line 132: self.exception = exc_value # store for later retrieval
							πF.SetLineno(132)
							if πE = πg.CheckLocal(πF, µexc_value, "exc_value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µexc_value); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßexception, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßexpected_regexp, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp007).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 133: if self.expected_regexp is None:
							πF.SetLineno(133)
						Label8:
							// line 134: return True
							πF.SetLineno(134)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label9
						Label9:
							// line 136: expected_regexp = self.expected_regexp
							πF.SetLineno(136)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßexpected_regexp, nil); πE != nil {
								continue
							}
							µexpected_regexp = πTemp001
							πTemp006 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexc_value, "exc_value"); πE != nil {
								continue
							}
							πTemp008[0] = µexc_value
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp006[0] = πTemp007
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µexpected_regexp, ßsearch, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							goto Label11
							// line 137: if not expected_regexp.search(str(exc_value)):
							πF.SetLineno(137)
						Label10:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µexpected_regexp, ßpattern, nil); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexc_value, "exc_value"); πE != nil {
								continue
							}
							πTemp008[0] = µexc_value
							if πTemp009, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002 = πg.NewTuple2(πTemp007, πTemp010).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\"%s\" does not match \"%s\"").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 138: raise self.failureException('"%s" does not match "%s"' %
							πF.SetLineno(138)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label11
						Label11:
							// line 140: return True
							πF.SetLineno(140)
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
					if πE = πClass.SetItem(πF, ß__exit__.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("_AssertRaisesContext").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_AssertRaisesContext.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 143: class TestCase(object):
			πF.SetLineno(143)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp011
			πTemp009 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TestCase", "/usr/lib/python2.7/unittest_case.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp042 *πg.Object
				_ = πTemp042
				var πTemp043 *πg.Object
				_ = πTemp043
				var πTemp044 *πg.Object
				_ = πTemp044
				var πTemp045 *πg.Object
				_ = πTemp045
				var πTemp046 *πg.Object
				_ = πTemp046
				var πTemp047 *πg.Object
				_ = πTemp047
				var πTemp048 *πg.Object
				_ = πTemp048
				var πTemp049 *πg.Object
				_ = πTemp049
				var πTemp050 *πg.Object
				_ = πTemp050
				var πTemp051 *πg.Object
				_ = πTemp051
				var πTemp052 *πg.Object
				_ = πTemp052
				var πTemp053 *πg.Object
				_ = πTemp053
				var πTemp054 *πg.Object
				_ = πTemp054
				var πTemp055 *πg.Object
				_ = πTemp055
				var πTemp056 *πg.Object
				_ = πTemp056
				var πTemp057 *πg.Object
				_ = πTemp057
				var πTemp058 *πg.Object
				_ = πTemp058
				var πTemp059 *πg.Object
				_ = πTemp059
				var πTemp060 *πg.Object
				_ = πTemp060
				var πTemp061 *πg.Object
				_ = πTemp061
				var πTemp062 *πg.Object
				_ = πTemp062
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 144: """A class whose instances are single test cases.
					πF.SetLineno(144)
					// line 144: """A class whose instances are single test cases.
					πF.SetLineno(144)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A class whose instances are single test cases.\n\n    By default, the test code itself should be placed in a method named\n    'runTest'.\n\n    If the fixture may be used for many test cases, create as\n    many test methods as are needed. When instantiating such a TestCase\n    subclass, specify in the constructor arguments the name of the test method\n    that the instance is to execute.\n\n    Test authors should subclass TestCase for their own tests. Construction\n    and deconstruction of the test's environment ('fixture') can be\n    implemented by overriding the 'setUp' and 'tearDown' methods respectively.\n\n    If it is necessary to override the __init__ method, the base class\n    __init__ method must always be called. It is important that subclasses\n    should not change the signature of their __init__ method, since instances\n    of the classes are instantiated automatically by parts of the framework\n    in order to be run.\n\n    When subclassing TestCase, you can set these attributes:\n    * failureException: determines which exception will be raised when\n        the instance's assertion methods fail; test methods raising this\n        exception will be deemed to have 'failed' rather than 'errored'.\n    * longMessage: determines whether long messages (including repr of\n        objects used in assert methods) will be printed on failure in *addition*\n        to any explicit message passed.\n    * maxDiff: sets the maximum length of a diff in failure messages\n        by assert methods using difflib. It is looked up as an instance\n        attribute so can be configured by individual tests if required.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 176: failureException = AssertionError
					πF.SetLineno(176)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßAssertionError); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßfailureException.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 178: longMessage = False
					πF.SetLineno(178)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßlongMessage.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 180: maxDiff = 80*8
					πF.SetLineno(180)
					if πTemp001, πE = πg.Mul(πF, πg.NewInt(80).ToObject(), πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßmaxDiff.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 185: _diffThreshold = 1<<16
					πF.SetLineno(185)
					if πTemp001, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_diffThreshold.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 189: _classSetupFailed = False
					πF.SetLineno(189)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_classSetupFailed.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 191: def __init__(self, methodName='runTest'):
					πF.SetLineno(191)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "methodName", Def: ßrunTest.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmethodName *πg.Object = πArgs[1]; _ = µmethodName
						var µtestMethod *πg.Object = πg.UnboundLocal; _ = µtestMethod
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Dict
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 192: """Create an instance of the class that will use the named test
							πF.SetLineno(192)
							// line 196: self._testMethodName = methodName
							πF.SetLineno(196)
							if πE = πg.CheckLocal(πF, µmethodName, "methodName"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmethodName); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_testMethodName, πTemp001); πE != nil {
								continue
							}
							// line 197: self._resultForDoCleanups = None
							πF.SetLineno(197)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_resultForDoCleanups, πTemp002); πE != nil {
								continue
							}
							// line 198: try:
							πF.SetLineno(198)
							πF.PushCheckpoint(2)
							// line 199: testMethod = getattr(self, methodName)
							πF.SetLineno(199)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µmethodName, "methodName"); πE != nil {
								continue
							}
							πTemp003[1] = µmethodName
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µtestMethod = πTemp002
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
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
							// line 200: except AttributeError:
							πF.SetLineno(200)
						Label3:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmethodName, "methodName"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp007, µmethodName).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("no such test method in %s: %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 201: raise ValueError("no such test method in %s: %s" %
							πF.SetLineno(201)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 204: self._cleanups = []
							πF.SetLineno(204)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_cleanups, πTemp002); πE != nil {
								continue
							}
							// line 209: self._type_equality_funcs = {}
							πF.SetLineno(209)
							πTemp008 = πg.NewDict()
							πTemp001 = πTemp008.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_type_equality_funcs, πTemp002); πE != nil {
								continue
							}
							// line 210: self.addTypeEqualityFunc(dict, 'assertDictEqual')
							πF.SetLineno(210)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = ßassertDictEqual.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßaddTypeEqualityFunc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 211: self.addTypeEqualityFunc(list, 'assertListEqual')
							πF.SetLineno(211)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = ßassertListEqual.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßaddTypeEqualityFunc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 212: self.addTypeEqualityFunc(tuple, 'assertTupleEqual')
							πF.SetLineno(212)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = ßassertTupleEqual.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßaddTypeEqualityFunc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 213: self.addTypeEqualityFunc(set, 'assertSetEqual')
							πF.SetLineno(213)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = ßassertSetEqual.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßaddTypeEqualityFunc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 214: self.addTypeEqualityFunc(frozenset, 'assertSetEqual')
							πF.SetLineno(214)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = ßassertSetEqual.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßaddTypeEqualityFunc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 215: try:
							πF.SetLineno(215)
							πF.PushCheckpoint(5)
							// line 216: self.addTypeEqualityFunc(unicode, 'assertMultiLineEqual')
							πF.SetLineno(216)
							πTemp003 = πF.MakeArgs(2)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp003[1] = ßassertMultiLineEqual.ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßaddTypeEqualityFunc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 217: except NameError:
							πF.SetLineno(217)
						Label6:
							// line 219: pass
							πF.SetLineno(219)
							πF.RestoreExc(nil, nil)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 192: """Create an instance of the class that will use the named test
					πF.SetLineno(192)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Create an instance of the class that will use the named test\n           method when executed. Raises a ValueError if the instance does\n           not have a method with the specified name.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 221: def addTypeEqualityFunc(self, typeobj, function):
					πF.SetLineno(221)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "typeobj", Def: nil}
					πTemp002[2] = πg.Param{Name: "function", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("addTypeEqualityFunc", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtypeobj *πg.Object = πArgs[1]; _ = µtypeobj
						var µfunction *πg.Object = πArgs[2]; _ = µfunction
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
							// line 222: """Add a type specific assertEqual style function to compare a type.
							πF.SetLineno(222)
							// line 234: self._type_equality_funcs[typeobj] = function
							πF.SetLineno(234)
							if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfunction); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_type_equality_funcs, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtypeobj, "typeobj"); πE != nil {
								continue
							}
							πTemp003 = µtypeobj
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
					if πE = πClass.SetItem(πF, ßaddTypeEqualityFunc.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 222: """Add a type specific assertEqual style function to compare a type.
					πF.SetLineno(222)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Add a type specific assertEqual style function to compare a type.\n\n        This method is for use by TestCase subclasses that need to register\n        their own type equality functions to provide nicer error messages.\n\n        Args:\n            typeobj: The data type to call this function on when both values\n                    are of the same type in assertEqual().\n            function: The callable taking two arguments and an optional\n                    msg= argument that raises self.failureException with a\n                    useful error message when the two arguments are not equal.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßaddTypeEqualityFunc); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 236: def addCleanup(self, function, *args, **kwargs):
					πF.SetLineno(236)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "function", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("addCleanup", "/usr/lib/python2.7/unittest_case.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfunction *πg.Object = πArgs[1]; _ = µfunction
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var µkwargs *πg.Object = πArgs[3]; _ = µkwargs
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
							// line 237: """Add a function, with arguments, to be called when the test is
							πF.SetLineno(237)
							// line 242: self._cleanups.append((function, args, kwargs))
							πF.SetLineno(242)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µfunction, µargs, µkwargs).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_cleanups, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßaddCleanup.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 237: """Add a function, with arguments, to be called when the test is
					πF.SetLineno(237)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Add a function, with arguments, to be called when the test is\n        completed. Functions added are called on a LIFO basis and are\n        called after tearDown on test failure or success.\n\n        Cleanup items are called even if setUp fails (unlike tearDown).").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßaddCleanup); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 244: def setUp(self):
					πF.SetLineno(244)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("setUp", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 245: "Hook method for setting up the test fixture before exercising it."
							πF.SetLineno(245)
							// line 246: pass
							πF.SetLineno(246)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsetUp.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 245: "Hook method for setting up the test fixture before exercising it."
					πF.SetLineno(245)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Hook method for setting up the test fixture before exercising it.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßsetUp); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 248: def tearDown(self):
					πF.SetLineno(248)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("tearDown", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 249: "Hook method for deconstructing the test fixture after testing it."
							πF.SetLineno(249)
							// line 250: pass
							πF.SetLineno(250)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtearDown.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 249: "Hook method for deconstructing the test fixture after testing it."
					πF.SetLineno(249)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Hook method for deconstructing the test fixture after testing it.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßtearDown); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 253: def setUpClass(cls):
					πF.SetLineno(253)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("setUpClass", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 254: "Hook method for setting up class fixture before running tests in the class."
							πF.SetLineno(254)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsetUpClass.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 254: "Hook method for setting up class fixture before running tests in the class."
					πF.SetLineno(254)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Hook method for setting up class fixture before running tests in the class.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßsetUpClass); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 255: setUpClass = classmethod(setUpClass)
					πF.SetLineno(255)
					πTemp010 = πF.MakeArgs(1)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßsetUpClass); πE != nil {
						continue
					}
					πTemp010[0] = πTemp008
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßsetUpClass.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 258: def tearDownClass(cls):
					πF.SetLineno(258)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("tearDownClass", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 259: "Hook method for deconstructing the class fixture after running all tests in the class."
							πF.SetLineno(259)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßtearDownClass.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 259: "Hook method for deconstructing the class fixture after running all tests in the class."
					πF.SetLineno(259)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Hook method for deconstructing the class fixture after running all tests in the class.").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßtearDownClass); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 260: tearDownClass = classmethod(tearDownClass)
					πF.SetLineno(260)
					πTemp010 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßtearDownClass); πE != nil {
						continue
					}
					πTemp010[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp009.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßtearDownClass.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 262: def countTestCases(self):
					πF.SetLineno(262)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("countTestCases", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 263: return 1
							πF.SetLineno(263)
							πR = πg.NewInt(1).ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcountTestCases.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 265: def defaultTestResult(self):
					πF.SetLineno(265)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("defaultTestResult", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 266: return result.TestResult()
							πF.SetLineno(266)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßresult); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßTestResult, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdefaultTestResult.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 268: def shortDescription(self):
					πF.SetLineno(268)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("shortDescription", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 269: """Returns a one-line description of the test, or None if no
							πF.SetLineno(269)
							// line 277: return ''
							πF.SetLineno(277)
							πR = ß.ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßshortDescription.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 269: """Returns a one-line description of the test, or None if no
					πF.SetLineno(269)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Returns a one-line description of the test, or None if no\n        description has been provided.\n\n        The default implementation of this method returns the first line of\n        the specified test method's docstring.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßshortDescription); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 280: def id(self):
					πF.SetLineno(280)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("id", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 281: return "%s.%s" % (strclass(self.__class__), self._testMethodName)
							πF.SetLineno(281)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstrclass); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_testMethodName, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp005, πTemp004).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s.%s").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßid.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 283: def __eq__(self, other):
					πF.SetLineno(283)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp002[0] = µother
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πg.GetBool(πTemp004 != πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 284: if type(self) is not type(other):
							πF.SetLineno(284)
						Label1:
							// line 285: return NotImplemented
							πF.SetLineno(285)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 287: return self._testMethodName == other._testMethodName
							πF.SetLineno(287)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_testMethodName, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µother, ß_testMethodName, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 289: def __ne__(self, other):
					πF.SetLineno(289)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 290: return not self == other
							πF.SetLineno(290)
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 292: def __hash__(self):
					πF.SetLineno(292)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
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
							// line 293: return hash((type(self), self._testMethodName))
							πF.SetLineno(293)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_testMethodName, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp005, πTemp004).ToObject()
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 295: def __str__(self):
					πF.SetLineno(295)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 296: return "%s (%s)" % (self._testMethodName, strclass(self.__class__))
							πF.SetLineno(296)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_testMethodName, nil); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßstrclass); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s (%s)").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 298: def __repr__(self):
					πF.SetLineno(298)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 299: return "<%s testMethod=%s>" % \
							πF.SetLineno(299)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstrclass); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_testMethodName, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp005, πTemp004).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s testMethod=%s>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 302: def _addSkip(self, result, reason):
					πF.SetLineno(302)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "result", Def: nil}
					πTemp002[2] = πg.Param{Name: "reason", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("_addSkip", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var µreason *πg.Object = πArgs[2]; _ = µreason
						var µaddSkip *πg.Object = πg.UnboundLocal; _ = µaddSkip
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
							// line 303: addSkip = getattr(result, 'addSkip', None)
							πF.SetLineno(303)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							πTemp001[1] = ßaddSkip.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µaddSkip = πTemp003
							if πE = πg.CheckLocal(πF, µaddSkip, "addSkip"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µaddSkip != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 304: if addSkip is not None:
							πF.SetLineno(304)
						Label1:
							// line 305: addSkip(self, reason)
							πF.SetLineno(305)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µreason, "reason"); πE != nil {
								continue
							}
							πTemp001[1] = µreason
							if πE = πg.CheckLocal(πF, µaddSkip, "addSkip"); πE != nil {
								continue
							}
							if πTemp002, πE = µaddSkip.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label3
						Label2:
							// line 307: warnings.warn("TestResult has no addSkip method, skips not reported",
							πF.SetLineno(307)
							πTemp001 = πF.MakeArgs(3)
							πTemp001[0] = πg.NewStr("TestResult has no addSkip method, skips not reported").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßRuntimeWarning); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwarn, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 309: result.addSuccess(self)
							πF.SetLineno(309)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßaddSuccess, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_addSkip.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 311: def run(self, result=None):
					πF.SetLineno(311)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "result", Def: πTemp021}
					πTemp020 = πg.NewFunction(πg.NewCode("run", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πArgs[1]; _ = µresult
						var µorig_result *πg.Object = πg.UnboundLocal; _ = µorig_result
						var µstartTestRun *πg.Object = πg.UnboundLocal; _ = µstartTestRun
						var µtestMethod *πg.Object = πg.UnboundLocal; _ = µtestMethod
						var µskip_why *πg.Object = πg.UnboundLocal; _ = µskip_why
						var µsuccess *πg.Object = πg.UnboundLocal; _ = µsuccess
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µaddExpectedFailure *πg.Object = πg.UnboundLocal; _ = µaddExpectedFailure
						var µaddUnexpectedSuccess *πg.Object = πg.UnboundLocal; _ = µaddUnexpectedSuccess
						var µcleanUpSuccess *πg.Object = πg.UnboundLocal; _ = µcleanUpSuccess
						var µstopTestRun *πg.Object = πg.UnboundLocal; _ = µstopTestRun
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8: goto Label8
							case 28: goto Label28
							case 10: goto Label10
							case 12: goto Label12
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							// line 312: orig_result = result
							πF.SetLineno(312)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							µorig_result = µresult
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µresult == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 313: if result is None:
							πF.SetLineno(313)
						Label1:
							// line 314: result = self.defaultTestResult()
							πF.SetLineno(314)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefaultTestResult, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µresult = πTemp002
							// line 315: startTestRun = getattr(result, 'startTestRun', None)
							πF.SetLineno(315)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp004[0] = µresult
							πTemp004[1] = ßstartTestRun.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µstartTestRun = πTemp002
							if πE = πg.CheckLocal(πF, µstartTestRun, "startTestRun"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µstartTestRun != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 316: if startTestRun is not None:
							πF.SetLineno(316)
						Label3:
							// line 317: startTestRun()
							πF.SetLineno(317)
							if πE = πg.CheckLocal(πF, µstartTestRun, "startTestRun"); πE != nil {
								continue
							}
							if πTemp001, πE = µstartTestRun.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 319: self._resultForDoCleanups = result
							πF.SetLineno(319)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µresult); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_resultForDoCleanups, πTemp001); πE != nil {
								continue
							}
							// line 320: result.startTest(self)
							πF.SetLineno(320)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßstartTest, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 322: testMethod = getattr(self, self._testMethodName)
							πF.SetLineno(322)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_testMethodName, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µtestMethod = πTemp002
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp004[1] = ß__unittest_skip__.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
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
							if πTemp003 {
								goto Label5
							}
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtestMethod, "testMethod"); πE != nil {
								continue
							}
							πTemp004[0] = µtestMethod
							πTemp004[1] = ß__unittest_skip__.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp004[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001 = πTemp005
						Label5:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 323: if (getattr(self.__class__, "__unittest_skip__", False) or
							πF.SetLineno(323)
						Label6:
							// line 326: try:
							πF.SetLineno(326)
							πF.PushCheckpoint(8)
							// line 327: skip_why = (getattr(self.__class__, '__unittest_skip_why__', '')
							πF.SetLineno(327)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp004[1] = ß__unittest_skip_why__.ToObject()
							πTemp004[2] = ß.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
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
							if πTemp003 {
								goto Label9
							}
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtestMethod, "testMethod"); πE != nil {
								continue
							}
							πTemp004[0] = µtestMethod
							πTemp004[1] = ß__unittest_skip_why__.ToObject()
							πTemp004[2] = ß.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001 = πTemp005
						Label9:
							µskip_why = πTemp001
							// line 329: self._addSkip(result, skip_why)
							πF.SetLineno(329)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp004[0] = µresult
							if πE = πg.CheckLocal(πF, µskip_why, "skip_why"); πE != nil {
								continue
							}
							πTemp004[1] = µskip_why
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_addSkip, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.PopCheckpoint()
							goto Label8
						Label8:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							// line 331: result.stopTest(self)
							πF.SetLineno(331)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßstopTest, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp006 != nil {
								πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
								continue
							}
							if πR != nil {
								continue
							}
							// line 332: return
							πF.SetLineno(332)
							πR = πg.None
							continue
							goto Label7
						Label7:
							// line 333: try:
							πF.SetLineno(333)
							πF.PushCheckpoint(10)
							// line 334: success = False
							πF.SetLineno(334)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µsuccess = πTemp001
							// line 335: try:
							πF.SetLineno(335)
							πF.PushCheckpoint(12)
							// line 336: self.setUp()
							πF.SetLineno(336)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsetUp, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							// line 344: try:
							πF.SetLineno(344)
							πF.PushCheckpoint(14)
							// line 345: testMethod()
							πF.SetLineno(345)
							if πE = πg.CheckLocal(πF, µtestMethod, "testMethod"); πE != nil {
								continue
							}
							if πTemp001, πE = µtestMethod.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							// line 371: success = True
							πF.SetLineno(371)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µsuccess = πTemp001
							goto Label13
						Label14:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label16
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_ExpectedFailure); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label17
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_UnexpectedSuccess); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label18
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSkipTest); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label19
							}
							goto Label20
							// line 346: except KeyboardInterrupt:
							πF.SetLineno(346)
						Label15:
							// line 347: raise
							πF.SetLineno(347)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label13
							// line 348: except self.failureException:
							πF.SetLineno(348)
						Label16:
							// line 349: result.addFailure(self, sys.exc_info())
							πF.SetLineno(349)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßaddFailure, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.RestoreExc(nil, nil)
							goto Label13
							// line 350: except _ExpectedFailure as e:
							πF.SetLineno(350)
						Label17:
							µe = πTemp006.ToObject()
							// line 351: addExpectedFailure = getattr(result, 'addExpectedFailure', None)
							πF.SetLineno(351)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp004[0] = µresult
							πTemp004[1] = ßaddExpectedFailure.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µaddExpectedFailure = πTemp002
							if πE = πg.CheckLocal(πF, µaddExpectedFailure, "addExpectedFailure"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µaddExpectedFailure != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label21
							}
							goto Label22
							// line 352: if addExpectedFailure is not None:
							πF.SetLineno(352)
						Label21:
							// line 353: addExpectedFailure(self, e.exc_info)
							πF.SetLineno(353)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µe, ßexc_info, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µaddExpectedFailure, "addExpectedFailure"); πE != nil {
								continue
							}
							if πTemp001, πE = µaddExpectedFailure.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label23
						Label22:
							// line 355: warnings.warn("TestResult has no addExpectedFailure method, reporting as passes",
							πF.SetLineno(355)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("TestResult has no addExpectedFailure method, reporting as passes").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeWarning); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwarn, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 357: result.addSuccess(self)
							πF.SetLineno(357)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßaddSuccess, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label23
						Label23:
							πF.RestoreExc(nil, nil)
							goto Label13
							// line 358: except _UnexpectedSuccess:
							πF.SetLineno(358)
						Label18:
							// line 359: addUnexpectedSuccess = getattr(result, 'addUnexpectedSuccess', None)
							πF.SetLineno(359)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp004[0] = µresult
							πTemp004[1] = ßaddUnexpectedSuccess.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µaddUnexpectedSuccess = πTemp002
							if πE = πg.CheckLocal(πF, µaddUnexpectedSuccess, "addUnexpectedSuccess"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µaddUnexpectedSuccess != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label24
							}
							goto Label25
							// line 360: if addUnexpectedSuccess is not None:
							πF.SetLineno(360)
						Label24:
							// line 361: addUnexpectedSuccess(self)
							πF.SetLineno(361)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µaddUnexpectedSuccess, "addUnexpectedSuccess"); πE != nil {
								continue
							}
							if πTemp001, πE = µaddUnexpectedSuccess.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label26
						Label25:
							// line 363: warnings.warn("TestResult has no addUnexpectedSuccess method, reporting as failures",
							πF.SetLineno(363)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("TestResult has no addUnexpectedSuccess method, reporting as failures").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeWarning); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßwarn, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 365: result.addFailure(self, sys.exc_info())
							πF.SetLineno(365)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßaddFailure, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label26
						Label26:
							πF.RestoreExc(nil, nil)
							goto Label13
							// line 366: except SkipTest as e:
							πF.SetLineno(366)
						Label19:
							µe = πTemp006.ToObject()
							// line 367: self._addSkip(result, str(e))
							πF.SetLineno(367)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp004[0] = µresult
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp008[0] = µe
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_addSkip, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.RestoreExc(nil, nil)
							goto Label13
							// line 368: except:
							πF.SetLineno(368)
						Label20:
							// line 369: result.addError(self, sys.exc_info())
							πF.SetLineno(369)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßaddError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.RestoreExc(nil, nil)
							goto Label13
						Label13:
							// line 373: try:
							πF.SetLineno(373)
							πF.PushCheckpoint(28)
							// line 374: self.tearDown()
							πF.SetLineno(374)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtearDown, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label27
						Label28:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label29
							}
							goto Label30
							// line 375: except KeyboardInterrupt:
							πF.SetLineno(375)
						Label29:
							// line 376: raise
							πF.SetLineno(376)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label27
							// line 377: except:
							πF.SetLineno(377)
						Label30:
							// line 378: result.addError(self, sys.exc_info())
							πF.SetLineno(378)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßaddError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 379: success = False
							πF.SetLineno(379)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µsuccess = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label27
						Label27:
							goto Label11
						Label12:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSkipTest); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label31
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label32
							}
							goto Label33
							// line 337: except SkipTest as e:
							πF.SetLineno(337)
						Label31:
							µe = πTemp006.ToObject()
							// line 338: self._addSkip(result, str(e))
							πF.SetLineno(338)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp004[0] = µresult
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp008[0] = µe
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_addSkip, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.RestoreExc(nil, nil)
							goto Label11
							// line 339: except KeyboardInterrupt:
							πF.SetLineno(339)
						Label32:
							// line 340: raise
							πF.SetLineno(340)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label11
							// line 341: except:
							πF.SetLineno(341)
						Label33:
							// line 342: result.addError(self, sys.exc_info())
							πF.SetLineno(342)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßaddError, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.RestoreExc(nil, nil)
							goto Label11
						Label11:
							// line 381: cleanUpSuccess = self.doCleanups()
							πF.SetLineno(381)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdoCleanups, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcleanUpSuccess = πTemp002
							// line 382: success = success and cleanUpSuccess
							πF.SetLineno(382)
							if πE = πg.CheckLocal(πF, µsuccess, "success"); πE != nil {
								continue
							}
							πTemp001 = µsuccess
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label34
							}
							if πE = πg.CheckLocal(πF, µcleanUpSuccess, "cleanUpSuccess"); πE != nil {
								continue
							}
							πTemp001 = µcleanUpSuccess
						Label34:
							µsuccess = πTemp001
							if πE = πg.CheckLocal(πF, µsuccess, "success"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µsuccess); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label35
							}
							goto Label36
							// line 383: if success:
							πF.SetLineno(383)
						Label35:
							// line 384: result.addSuccess(self)
							πF.SetLineno(384)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßaddSuccess, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label36
						Label36:
							πF.PopCheckpoint()
							goto Label10
						Label10:
							πTemp006, πTemp007 = πF.RestoreExc(nil, nil)
							// line 386: result.stopTest(self)
							πF.SetLineno(386)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßstopTest, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µorig_result, "orig_result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µorig_result == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label37
							}
							goto Label38
							// line 387: if orig_result is None:
							πF.SetLineno(387)
						Label37:
							// line 388: stopTestRun = getattr(result, 'stopTestRun', None)
							πF.SetLineno(388)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp004[0] = µresult
							πTemp004[1] = ßstopTestRun.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µstopTestRun = πTemp002
							if πE = πg.CheckLocal(πF, µstopTestRun, "stopTestRun"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µstopTestRun != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label39
							}
							goto Label40
							// line 389: if stopTestRun is not None:
							πF.SetLineno(389)
						Label39:
							// line 390: stopTestRun()
							πF.SetLineno(390)
							if πE = πg.CheckLocal(πF, µstopTestRun, "stopTestRun"); πE != nil {
								continue
							}
							if πTemp001, πE = µstopTestRun.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label40
						Label40:
							goto Label38
						Label38:
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
					if πE = πClass.SetItem(πF, ßrun.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 392: def doCleanups(self):
					πF.SetLineno(392)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("doCleanups", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µok *πg.Object = πg.UnboundLocal; _ = µok
						var µfunction *πg.Object = πg.UnboundLocal; _ = µfunction
						var µargs *πg.Object = πg.UnboundLocal; _ = µargs
						var µkwargs *πg.Object = πg.UnboundLocal; _ = µkwargs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
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
							// line 393: """Execute all cleanup functions. Normally called for you after
							πF.SetLineno(393)
							// line 395: result = self._resultForDoCleanups
							πF.SetLineno(395)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_resultForDoCleanups, nil); πE != nil {
								continue
							}
							µresult = πTemp001
							// line 396: ok = True
							πF.SetLineno(396)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µok = πTemp001
							// line 397: while self._cleanups:
							πF.SetLineno(397)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_cleanups, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 398: function, args, kwargs = self._cleanups.pop(-1)
							πF.SetLineno(398)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_cleanups, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp001); πE != nil {
								continue
							}
							µfunction = πTemp005
							µargs = πTemp006
							µkwargs = πTemp007
							// line 399: try:
							πF.SetLineno(399)
							πF.PushCheckpoint(5)
							// line 400: function(*args, **kwargs)
							πF.SetLineno(400)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, µfunction, nil, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyboardInterrupt); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							goto Label7
							// line 401: except KeyboardInterrupt:
							πF.SetLineno(401)
						Label6:
							// line 402: raise
							πF.SetLineno(402)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label4
							// line 403: except:
							πF.SetLineno(403)
						Label7:
							// line 404: ok = False
							πF.SetLineno(404)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µok = πTemp001
							// line 405: result.addError(self, sys.exc_info())
							πF.SetLineno(405)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßaddError, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 406: return ok
							πF.SetLineno(406)
							if πE = πg.CheckLocal(πF, µok, "ok"); πE != nil {
								continue
							}
							πR = µok
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdoCleanups.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 393: """Execute all cleanup functions. Normally called for you after
					πF.SetLineno(393)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("Execute all cleanup functions. Normally called for you after\n        tearDown.").ToObject()); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßdoCleanups); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
						continue
					}
					// line 408: def __call__(self, *args, **kwds):
					πF.SetLineno(408)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("__call__", "/usr/lib/python2.7/unittest_case.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwds *πg.Object = πArgs[2]; _ = µkwds
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
							// line 409: return self.run(*args, **kwds)
							πF.SetLineno(409)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrun, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwds); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__call__.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 411: def debug(self):
					πF.SetLineno(411)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("debug", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfunction *πg.Object = πg.UnboundLocal; _ = µfunction
						var µargs *πg.Object = πg.UnboundLocal; _ = µargs
						var µkwargs *πg.Object = πg.UnboundLocal; _ = µkwargs
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 412: """Run the test without collecting errors in a TestResult"""
							πF.SetLineno(412)
							// line 413: self.setUp()
							πF.SetLineno(413)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsetUp, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 414: getattr(self, self._testMethodName)()
							πF.SetLineno(414)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_testMethodName, nil); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 415: self.tearDown()
							πF.SetLineno(415)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtearDown, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 416: while self._cleanups:
							πF.SetLineno(416)
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_cleanups, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 417: function, args, kwargs = self._cleanups.pop(-1)
							πF.SetLineno(417)
							πTemp003 = πF.MakeArgs(1)
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_cleanups, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp001); πE != nil {
								continue
							}
							µfunction = πTemp002
							µargs = πTemp006
							µkwargs = πTemp007
							// line 418: function(*args, **kwargs)
							πF.SetLineno(418)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfunction, "function"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, µfunction, nil, µargs, nil, µkwargs); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdebug.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 412: """Run the test without collecting errors in a TestResult"""
					πF.SetLineno(412)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("Run the test without collecting errors in a TestResult").ToObject()); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßdebug); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
						continue
					}
					// line 420: def skipTest(self, reason):
					πF.SetLineno(420)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "reason", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("skipTest", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µreason *πg.Object = πArgs[1]; _ = µreason
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
							// line 421: """Skip this test."""
							πF.SetLineno(421)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreason, "reason"); πE != nil {
								continue
							}
							πTemp001[0] = µreason
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSkipTest); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 422: raise SkipTest(reason)
							πF.SetLineno(422)
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
					if πE = πClass.SetItem(πF, ßskipTest.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 421: """Skip this test."""
					πF.SetLineno(421)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("Skip this test.").ToObject()); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßskipTest); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
						continue
					}
					// line 424: def fail(self, msg=None):
					πF.SetLineno(424)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "msg", Def: πTemp026}
					πTemp025 = πg.NewFunction(πg.NewCode("fail", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πArgs[1]; _ = µmsg
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
							// line 425: """Fail immediately, with the given message."""
							πF.SetLineno(425)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 426: raise self.failureException(msg)
							πF.SetLineno(426)
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
					if πE = πClass.SetItem(πF, ßfail.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 425: """Fail immediately, with the given message."""
					πF.SetLineno(425)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πg.NewStr("Fail immediately, with the given message.").ToObject()); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßfail); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp027, ß__doc__, πTemp026); πE != nil {
						continue
					}
					// line 428: def assertFalse(self, expr, msg=None):
					πF.SetLineno(428)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "expr", Def: nil}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "msg", Def: πTemp027}
					πTemp026 = πg.NewFunction(πg.NewCode("assertFalse", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpr *πg.Object = πArgs[1]; _ = µexpr
						var µmsg *πg.Object = πArgs[2]; _ = µmsg
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 429: """Check that the expression is false."""
							πF.SetLineno(429)
							if πE = πg.CheckLocal(πF, µexpr, "expr"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µexpr); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 430: if expr:
							πF.SetLineno(430)
						Label1:
							// line 431: msg = self._formatMessage(msg, "%s is not false" % safe_repr(expr))
							πF.SetLineno(431)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[0] = µmsg
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpr, "expr"); πE != nil {
								continue
							}
							πTemp004[0] = µexpr
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s is not false").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µmsg = πTemp005
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 432: raise self.failureException(msg)
							πF.SetLineno(432)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßassertFalse.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 429: """Check that the expression is false."""
					πF.SetLineno(429)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πg.NewStr("Check that the expression is false.").ToObject()); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßassertFalse); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp028, ß__doc__, πTemp027); πE != nil {
						continue
					}
					// line 434: def assertTrue(self, expr, msg=None):
					πF.SetLineno(434)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "expr", Def: nil}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "msg", Def: πTemp028}
					πTemp027 = πg.NewFunction(πg.NewCode("assertTrue", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpr *πg.Object = πArgs[1]; _ = µexpr
						var µmsg *πg.Object = πArgs[2]; _ = µmsg
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 435: """Check that the expression is true."""
							πF.SetLineno(435)
							if πE = πg.CheckLocal(πF, µexpr, "expr"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µexpr); πE != nil {
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
							// line 436: if not expr:
							πF.SetLineno(436)
						Label1:
							// line 437: msg = self._formatMessage(msg, "%s is not true" % safe_repr(expr))
							πF.SetLineno(437)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp003[0] = µmsg
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpr, "expr"); πE != nil {
								continue
							}
							πTemp004[0] = µexpr
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s is not true").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp003[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmsg = πTemp005
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp003[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 438: raise self.failureException(msg)
							πF.SetLineno(438)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßassertTrue.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 435: """Check that the expression is true."""
					πF.SetLineno(435)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πg.NewStr("Check that the expression is true.").ToObject()); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßassertTrue); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp029, ß__doc__, πTemp028); πE != nil {
						continue
					}
					// line 440: def _formatMessage(self, msg, standardMsg):
					πF.SetLineno(440)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "msg", Def: nil}
					πTemp002[2] = πg.Param{Name: "standardMsg", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("_formatMessage", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πArgs[1]; _ = µmsg
						var µstandardMsg *πg.Object = πArgs[2]; _ = µstandardMsg
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.BaseException
						_ = πTemp004
						var πTemp005 *πg.Traceback
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 441: """Honour the longMessage attribute when generating failure messages.
							πF.SetLineno(441)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlongMessage, nil); πE != nil {
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
							// line 450: if not self.longMessage:
							πF.SetLineno(450)
						Label1:
							// line 451: return msg or standardMsg
							πF.SetLineno(451)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001 = µmsg
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp001 = µstandardMsg
						Label3:
							πR = πTemp001
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmsg == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 452: if msg is None:
							πF.SetLineno(452)
						Label4:
							// line 453: return standardMsg
							πF.SetLineno(453)
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πR = µstandardMsg
							continue
							goto Label5
						Label5:
							// line 454: try:
							πF.SetLineno(454)
							πF.PushCheckpoint(7)
							// line 457: return '%s : %s' % (standardMsg, msg)
							πF.SetLineno(457)
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µstandardMsg, µmsg).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s : %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.PopCheckpoint()
							goto Label6
						Label7:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeDecodeError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 458: except UnicodeDecodeError:
							πF.SetLineno(458)
						Label8:
							// line 459: return  '%s : %s' % (safe_repr(standardMsg), safe_repr(msg))
							πF.SetLineno(459)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp006[0] = µstandardMsg
							if πTemp007, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp006[0] = µmsg
							if πTemp007, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002 = πg.NewTuple2(πTemp008, πTemp009).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s : %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πR = πTemp001
							continue
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
					if πE = πClass.SetItem(πF, ß_formatMessage.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 441: """Honour the longMessage attribute when generating failure messages.
					πF.SetLineno(441)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp029}, πg.NewStr("Honour the longMessage attribute when generating failure messages.\n        If longMessage is False this means:\n        * Use only an explicit message if it is provided\n        * Otherwise use the standard message for the assert\n\n        If longMessage is True:\n        * Use the standard message\n        * If an explicit message is provided, plus ' : ' and the explicit message\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ß_formatMessage); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp030, ß__doc__, πTemp029); πE != nil {
						continue
					}
					// line 462: def assertRaises(self, excClass, callableObj=None, *args, **kwargs):
					πF.SetLineno(462)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "excClass", Def: nil}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "callableObj", Def: πTemp030}
					πTemp029 = πg.NewFunction(πg.NewCode("assertRaises", "/usr/lib/python2.7/unittest_case.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexcClass *πg.Object = πArgs[1]; _ = µexcClass
						var µcallableObj *πg.Object = πArgs[2]; _ = µcallableObj
						var µargs *πg.Object = πArgs[3]; _ = µargs
						var µkwargs *πg.Object = πArgs[4]; _ = µkwargs
						var µcontext *πg.Object = πg.UnboundLocal; _ = µcontext
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 *πg.Type
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							default: panic("unexpected function state")
							}
							// line 463: """Fail unless an exception of class excClass is raised
							πF.SetLineno(463)
							// line 485: context = _AssertRaisesContext(excClass, self)
							πF.SetLineno(485)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µexcClass, "excClass"); πE != nil {
								continue
							}
							πTemp001[0] = µexcClass
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[1] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_AssertRaisesContext); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcontext = πTemp003
							if πE = πg.CheckLocal(πF, µcallableObj, "callableObj"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µcallableObj == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 486: if callableObj is None:
							πF.SetLineno(486)
						Label1:
							// line 487: return context
							πF.SetLineno(487)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πR = µcontext
							continue
							goto Label2
						Label2:
							// line 488: with context:
							πF.SetLineno(488)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcontext.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcontext.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp003.Call(πF, πg.Args{µcontext}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(3)
							// line 489: callableObj(*args, **kwargs)
							πF.SetLineno(489)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcallableObj, "callableObj"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Invoke(πF, µcallableObj, nil, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label3:
							πTemp006, πTemp007 = nil, nil
							if πE != nil {
								πTemp006, πTemp007 = πF.ExcInfo()
							}
							if πTemp006 != nil {
								πTemp008 = πTemp006.Type()
								if πTemp005, πE = πTemp002.Call(πF, πg.Args{µcontext, πTemp008.ToObject(), πTemp006.ToObject(), πTemp007.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp005, πE = πTemp002.Call(πF, πg.Args{µcontext, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp006 != nil && πTemp004 != true {
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
					if πE = πClass.SetItem(πF, ßassertRaises.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 463: """Fail unless an exception of class excClass is raised
					πF.SetLineno(463)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp030}, πg.NewStr("Fail unless an exception of class excClass is raised\n           by callableObj when invoked with arguments args and keyword\n           arguments kwargs. If a different type of exception is\n           raised, it will not be caught, and the test case will be\n           deemed to have suffered an error, exactly as for an\n           unexpected exception.\n\n           If called with callableObj omitted or None, will return a\n           context object used like this::\n\n                with self.assertRaises(SomeException):\n                    do_something()\n\n           The context manager keeps a reference to the exception as\n           the 'exception' attribute. This allows you to inspect the\n           exception after the assertion::\n\n               with self.assertRaises(SomeException) as cm:\n                   do_something()\n               the_exception = cm.exception\n               self.assertEqual(the_exception.error_code, 3)\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßassertRaises); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp031, ß__doc__, πTemp030); πE != nil {
						continue
					}
					// line 491: def _getAssertEqualityFunc(self, first, second):
					πF.SetLineno(491)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "first", Def: nil}
					πTemp002[2] = πg.Param{Name: "second", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("_getAssertEqualityFunc", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfirst *πg.Object = πArgs[1]; _ = µfirst
						var µsecond *πg.Object = πArgs[2]; _ = µsecond
						var µasserter *πg.Object = πg.UnboundLocal; _ = µasserter
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
						var πTemp006 bool
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
							// line 492: """Get a detailed comparison function for the types of the two args.
							πF.SetLineno(492)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp002[0] = µfirst
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp002[0] = µsecond
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πg.GetBool(πTemp004 == πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 508: if type(first) is type(second):
							πF.SetLineno(508)
						Label1:
							// line 509: asserter = self._type_equality_funcs.get(type(first))
							πF.SetLineno(509)
							πTemp002 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp007[0] = µfirst
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_type_equality_funcs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µasserter = πTemp001
							if πE = πg.CheckLocal(πF, µasserter, "asserter"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µasserter != πTemp003).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 510: if asserter is not None:
							πF.SetLineno(510)
						Label3:
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µasserter, "asserter"); πE != nil {
								continue
							}
							πTemp002[0] = µasserter
							if πTemp001, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 511: if isinstance(asserter, basestring):
							πF.SetLineno(511)
						Label5:
							// line 512: asserter = getattr(self, asserter)
							πF.SetLineno(512)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µasserter, "asserter"); πE != nil {
								continue
							}
							πTemp002[1] = µasserter
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µasserter = πTemp003
							goto Label6
						Label6:
							// line 513: return asserter
							πF.SetLineno(513)
							if πE = πg.CheckLocal(πF, µasserter, "asserter"); πE != nil {
								continue
							}
							πR = µasserter
							continue
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 515: return self._baseAssertEqual
							πF.SetLineno(515)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_baseAssertEqual, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_getAssertEqualityFunc.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 492: """Get a detailed comparison function for the types of the two args.
					πF.SetLineno(492)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp031}, πg.NewStr("Get a detailed comparison function for the types of the two args.\n\n        Returns: A callable accepting (first, second, msg=None) that will\n        raise a failure exception if first != second with a useful human\n        readable error message for those types.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ß_getAssertEqualityFunc); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp032, ß__doc__, πTemp031); πE != nil {
						continue
					}
					// line 517: def _baseAssertEqual(self, first, second, msg=None):
					πF.SetLineno(517)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "first", Def: nil}
					πTemp002[2] = πg.Param{Name: "second", Def: nil}
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp032}
					πTemp031 = πg.NewFunction(πg.NewCode("_baseAssertEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfirst *πg.Object = πArgs[1]; _ = µfirst
						var µsecond *πg.Object = πArgs[2]; _ = µsecond
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 518: """The default assertEqual implementation, not type specific."""
							πF.SetLineno(518)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µfirst, µsecond); πE != nil {
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
							// line 519: if not first == second:
							πF.SetLineno(519)
						Label1:
							// line 520: standardMsg = '%s != %s' % (safe_repr(first), safe_repr(second))
							πF.SetLineno(520)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp004[0] = µfirst
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp004[0] = µsecond
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s != %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 521: msg = self._formatMessage(msg, standardMsg)
							πF.SetLineno(521)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp004[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp004[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µmsg = πTemp002
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp004[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 522: raise self.failureException(msg)
							πF.SetLineno(522)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ß_baseAssertEqual.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 518: """The default assertEqual implementation, not type specific."""
					πF.SetLineno(518)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp032}, πg.NewStr("The default assertEqual implementation, not type specific.").ToObject()); πE != nil {
						continue
					}
					if πTemp033, πE = πg.ResolveClass(πF, πClass, nil, ß_baseAssertEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp033, ß__doc__, πTemp032); πE != nil {
						continue
					}
					// line 524: def assertEqual(self, first, second, msg=None):
					πF.SetLineno(524)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "first", Def: nil}
					πTemp002[2] = πg.Param{Name: "second", Def: nil}
					if πTemp033, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp033}
					πTemp032 = πg.NewFunction(πg.NewCode("assertEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfirst *πg.Object = πArgs[1]; _ = µfirst
						var µsecond *πg.Object = πArgs[2]; _ = µsecond
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µassertion_func *πg.Object = πg.UnboundLocal; _ = µassertion_func
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 525: """Fail if the two objects are unequal as determined by the '=='
							πF.SetLineno(525)
							// line 528: assertion_func = self._getAssertEqualityFunc(first, second)
							πF.SetLineno(528)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp001[0] = µfirst
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp001[1] = µsecond
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_getAssertEqualityFunc, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µassertion_func = πTemp003
							// line 529: assertion_func(first, second, msg=msg)
							πF.SetLineno(529)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp001[0] = µfirst
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp001[1] = µsecond
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp004 = πg.KWArgs{
								{"msg", µmsg},
							}
							if πE = πg.CheckLocal(πF, µassertion_func, "assertion_func"); πE != nil {
								continue
							}
							if πTemp002, πE = µassertion_func.Call(πF, πTemp001, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertEqual.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 525: """Fail if the two objects are unequal as determined by the '=='
					πF.SetLineno(525)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp033}, πg.NewStr("Fail if the two objects are unequal as determined by the '=='\n           operator.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßassertEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp034, ß__doc__, πTemp033); πE != nil {
						continue
					}
					// line 531: def assertNotEqual(self, first, second, msg=None):
					πF.SetLineno(531)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "first", Def: nil}
					πTemp002[2] = πg.Param{Name: "second", Def: nil}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp034}
					πTemp033 = πg.NewFunction(πg.NewCode("assertNotEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfirst *πg.Object = πArgs[1]; _ = µfirst
						var µsecond *πg.Object = πArgs[2]; _ = µsecond
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
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
							// line 532: """Fail if the two objects are equal as determined by the '!='
							πF.SetLineno(532)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µfirst, µsecond); πE != nil {
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
							// line 535: if not first != second:
							πF.SetLineno(535)
						Label1:
							// line 536: msg = self._formatMessage(msg, '%s == %s' % (safe_repr(first),
							πF.SetLineno(536)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp004[0] = µmsg
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp005[0] = µfirst
							if πTemp006, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp005[0] = µsecond
							if πTemp006, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s == %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µmsg = πTemp002
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp004[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 538: raise self.failureException(msg)
							πF.SetLineno(538)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßassertNotEqual.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 532: """Fail if the two objects are equal as determined by the '!='
					πF.SetLineno(532)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp034}, πg.NewStr("Fail if the two objects are equal as determined by the '!='\n           operator.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßassertNotEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp035, ß__doc__, πTemp034); πE != nil {
						continue
					}
					// line 541: def assertAlmostEqual(self, first, second, places=None, msg=None, delta=None):
					πF.SetLineno(541)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "first", Def: nil}
					πTemp002[2] = πg.Param{Name: "second", Def: nil}
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "places", Def: πTemp035}
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "msg", Def: πTemp035}
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "delta", Def: πTemp035}
					πTemp034 = πg.NewFunction(πg.NewCode("assertAlmostEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfirst *πg.Object = πArgs[1]; _ = µfirst
						var µsecond *πg.Object = πArgs[2]; _ = µsecond
						var µplaces *πg.Object = πArgs[3]; _ = µplaces
						var µmsg *πg.Object = πArgs[4]; _ = µmsg
						var µdelta *πg.Object = πArgs[5]; _ = µdelta
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
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
							// line 542: """Fail if the two objects are unequal as determined by their
							πF.SetLineno(542)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µfirst, µsecond); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 553: if first == second:
							πF.SetLineno(553)
						Label1:
							// line 555: return
							πF.SetLineno(555)
							πR = πg.None
							continue
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µdelta != πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µplaces, "places"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µplaces != πTemp004).ToObject()
							πTemp001 = πTemp003
						Label3:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 556: if delta is not None and places is not None:
							πF.SetLineno(556)
						Label4:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("specify delta or places not both").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 557: raise TypeError("specify delta or places not both")
							πF.SetLineno(557)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdelta != πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 559: if delta is not None:
							πF.SetLineno(559)
						Label6:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µfirst, µsecond); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßabs); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, πTemp004, µdelta); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							goto Label10
							// line 560: if abs(first - second) <= delta:
							πF.SetLineno(560)
						Label9:
							// line 561: return
							πF.SetLineno(561)
							πR = πg.None
							continue
							goto Label10
						Label10:
							// line 563: standardMsg = '%s != %s within %s delta' % (safe_repr(first),
							πF.SetLineno(563)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp005[0] = µfirst
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp005[0] = µsecond
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							πTemp005[0] = µdelta
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp003 = πg.NewTuple3(πTemp006, πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s != %s within %s delta").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							goto Label8
						Label7:
							if πE = πg.CheckLocal(πF, µplaces, "places"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µplaces == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label11
							}
							goto Label12
							// line 567: if places is None:
							πF.SetLineno(567)
						Label11:
							// line 568: places = 7
							πF.SetLineno(568)
							µplaces = πg.NewInt(7).ToObject()
							goto Label12
						Label12:
							πTemp005 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µsecond, µfirst); πE != nil {
								continue
							}
							πTemp009[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßabs); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µplaces, "places"); πE != nil {
								continue
							}
							πTemp005[1] = µplaces
							if πTemp003, πE = πg.ResolveGlobal(πF, ßround); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label13
							}
							goto Label14
							// line 570: if round(abs(second-first), places) == 0:
							πF.SetLineno(570)
						Label13:
							// line 571: return
							πF.SetLineno(571)
							πR = πg.None
							continue
							goto Label14
						Label14:
							// line 573: standardMsg = '%s != %s within %r places' % (safe_repr(first),
							πF.SetLineno(573)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp005[0] = µfirst
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp005[0] = µsecond
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µplaces, "places"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(πTemp006, πTemp007, µplaces).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s != %s within %r places").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							goto Label8
						Label8:
							// line 576: msg = self._formatMessage(msg, standardMsg)
							πF.SetLineno(576)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp005[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp005[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmsg = πTemp003
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp005[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 577: raise self.failureException(msg)
							πF.SetLineno(577)
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
					if πE = πClass.SetItem(πF, ßassertAlmostEqual.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 542: """Fail if the two objects are unequal as determined by their
					πF.SetLineno(542)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp035}, πg.NewStr("Fail if the two objects are unequal as determined by their\n           difference rounded to the given number of decimal places\n           (default 7) and comparing to zero, or by comparing that the\n           between the two objects is more than the given delta.\n\n           Note that decimal places (from zero) are usually not the same\n           as significant digits (measured from the most signficant digit).\n\n           If the two objects compare equal then they will automatically\n           compare almost equal.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßassertAlmostEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp036, ß__doc__, πTemp035); πE != nil {
						continue
					}
					// line 579: def assertNotAlmostEqual(self, first, second, places=None, msg=None, delta=None):
					πF.SetLineno(579)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "first", Def: nil}
					πTemp002[2] = πg.Param{Name: "second", Def: nil}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "places", Def: πTemp036}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "msg", Def: πTemp036}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "delta", Def: πTemp036}
					πTemp035 = πg.NewFunction(πg.NewCode("assertNotAlmostEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfirst *πg.Object = πArgs[1]; _ = µfirst
						var µsecond *πg.Object = πArgs[2]; _ = µsecond
						var µplaces *πg.Object = πArgs[3]; _ = µplaces
						var µmsg *πg.Object = πArgs[4]; _ = µmsg
						var µdelta *πg.Object = πArgs[5]; _ = µdelta
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 580: """Fail if the two objects are equal as determined by their
							πF.SetLineno(580)
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µdelta != πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µplaces, "places"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µplaces != πTemp004).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 590: if delta is not None and places is not None:
							πF.SetLineno(590)
						Label2:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("specify delta or places not both").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 591: raise TypeError("specify delta or places not both")
							πF.SetLineno(591)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdelta != πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 592: if delta is not None:
							πF.SetLineno(592)
						Label4:
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µfirst, µsecond); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label7
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µfirst, µsecond); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßabs); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp007, µdelta); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label7:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label8
							}
							goto Label9
							// line 593: if not (first == second) and abs(first - second) > delta:
							πF.SetLineno(593)
						Label8:
							// line 594: return
							πF.SetLineno(594)
							πR = πg.None
							continue
							goto Label9
						Label9:
							// line 595: standardMsg = '%s == %s within %s delta' % (safe_repr(first),
							πF.SetLineno(595)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp005[0] = µfirst
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp005[0] = µsecond
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelta, "delta"); πE != nil {
								continue
							}
							πTemp005[0] = µdelta
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp003 = πg.NewTuple3(πTemp007, πTemp008, πTemp009).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s == %s within %s delta").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							goto Label6
						Label5:
							if πE = πg.CheckLocal(πF, µplaces, "places"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µplaces == πTemp003).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label10
							}
							goto Label11
							// line 599: if places is None:
							πF.SetLineno(599)
						Label10:
							// line 600: places = 7
							πF.SetLineno(600)
							µplaces = πg.NewInt(7).ToObject()
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µfirst, µsecond); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label12
							}
							πTemp005 = πF.MakeArgs(2)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, µsecond, µfirst); πE != nil {
								continue
							}
							πTemp010[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßabs); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp005[0] = πTemp007
							if πE = πg.CheckLocal(πF, µplaces, "places"); πE != nil {
								continue
							}
							πTemp005[1] = µplaces
							if πTemp004, πE = πg.ResolveGlobal(πF, ßround); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.NE(πF, πTemp007, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label12:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label13
							}
							goto Label14
							// line 601: if not (first == second) and round(abs(second-first), places) != 0:
							πF.SetLineno(601)
						Label13:
							// line 602: return
							πF.SetLineno(602)
							πR = πg.None
							continue
							goto Label14
						Label14:
							// line 603: standardMsg = '%s == %s within %r places' % (safe_repr(first),
							πF.SetLineno(603)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp005[0] = µfirst
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp005[0] = µsecond
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µplaces, "places"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(πTemp007, πTemp008, µplaces).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s == %s within %r places").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							goto Label6
						Label6:
							// line 607: msg = self._formatMessage(msg, standardMsg)
							πF.SetLineno(607)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp005[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp005[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmsg = πTemp003
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp005[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 608: raise self.failureException(msg)
							πF.SetLineno(608)
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
					if πE = πClass.SetItem(πF, ßassertNotAlmostEqual.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 580: """Fail if the two objects are equal as determined by their
					πF.SetLineno(580)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πg.NewStr("Fail if the two objects are equal as determined by their\n           difference rounded to the given number of decimal places\n           (default 7) and comparing to zero, or by comparing that the\n           between the two objects is less than the given delta.\n\n           Note that decimal places (from zero) are usually not the same\n           as significant digits (measured from the most signficant digit).\n\n           Objects that are equal automatically fail.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßassertNotAlmostEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp037, ß__doc__, πTemp036); πE != nil {
						continue
					}
					// line 615: assertEquals = assertEqual
					πF.SetLineno(615)
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßassertEqual); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßassertEquals.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 616: assertNotEquals = assertNotEqual
					πF.SetLineno(616)
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßassertNotEqual); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßassertNotEquals.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 617: assertAlmostEquals = assertAlmostEqual
					πF.SetLineno(617)
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßassertAlmostEqual); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßassertAlmostEquals.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 618: assertNotAlmostEquals = assertNotAlmostEqual
					πF.SetLineno(618)
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßassertNotAlmostEqual); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßassertNotAlmostEquals.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 619: assert_ = assertTrue
					πF.SetLineno(619)
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßassertTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßassert_.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 623: def _deprecate(original_func):
					πF.SetLineno(623)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "original_func", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("_deprecate", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µoriginal_func *πg.Object = πArgs[0]; _ = µoriginal_func
						var µdeprecated_func *πg.Object = πg.UnboundLocal; _ = µdeprecated_func
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []πg.Param
						_ = πTemp002
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 624: def deprecated_func(*args, **kwargs):
							πF.SetLineno(624)
							πTemp002 = make([]πg.Param, 0)
							πTemp001 = πg.NewFunction(πg.NewCode("deprecated_func", "/usr/lib/python2.7/unittest_case.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µargs *πg.Object = πArgs[0]; _ = µargs
								var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
									// line 625: warnings.warn(
									πF.SetLineno(625)
									πTemp001 = πF.MakeArgs(3)
									if πE = πg.CheckLocal(πF, µoriginal_func, "original_func"); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, µoriginal_func, ß__name__, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Mod(πF, πg.NewStr("Please use %s instead.").ToObject(), πTemp003); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßPendingDeprecationWarning); πE != nil {
										continue
									}
									πTemp001[1] = πTemp002
									πTemp001[2] = πg.NewInt(2).ToObject()
									if πTemp002, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
										continue
									}
									if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwarn, nil); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 629: return original_func(*args, **kwargs)
									πF.SetLineno(629)
									if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µoriginal_func, "original_func"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Invoke(πF, µoriginal_func, nil, µargs, nil, µkwargs); πE != nil {
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
							µdeprecated_func = πTemp001
							// line 630: return deprecated_func
							πF.SetLineno(630)
							if πE = πg.CheckLocal(πF, µdeprecated_func, "deprecated_func"); πE != nil {
								continue
							}
							πR = µdeprecated_func
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_deprecate.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 632: failUnlessEqual = _deprecate(assertEqual)
					πF.SetLineno(632)
					πTemp010 = πF.MakeArgs(1)
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßassertEqual); πE != nil {
						continue
					}
					πTemp010[0] = πTemp037
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ß_deprecate); πE != nil {
						continue
					}
					if πTemp038, πE = πTemp037.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßfailUnlessEqual.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 633: failIfEqual = _deprecate(assertNotEqual)
					πF.SetLineno(633)
					πTemp010 = πF.MakeArgs(1)
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßassertNotEqual); πE != nil {
						continue
					}
					πTemp010[0] = πTemp037
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ß_deprecate); πE != nil {
						continue
					}
					if πTemp038, πE = πTemp037.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßfailIfEqual.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 634: failUnlessAlmostEqual = _deprecate(assertAlmostEqual)
					πF.SetLineno(634)
					πTemp010 = πF.MakeArgs(1)
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßassertAlmostEqual); πE != nil {
						continue
					}
					πTemp010[0] = πTemp037
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ß_deprecate); πE != nil {
						continue
					}
					if πTemp038, πE = πTemp037.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßfailUnlessAlmostEqual.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 635: failIfAlmostEqual = _deprecate(assertNotAlmostEqual)
					πF.SetLineno(635)
					πTemp010 = πF.MakeArgs(1)
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßassertNotAlmostEqual); πE != nil {
						continue
					}
					πTemp010[0] = πTemp037
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ß_deprecate); πE != nil {
						continue
					}
					if πTemp038, πE = πTemp037.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßfailIfAlmostEqual.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 636: failUnless = _deprecate(assertTrue)
					πF.SetLineno(636)
					πTemp010 = πF.MakeArgs(1)
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßassertTrue); πE != nil {
						continue
					}
					πTemp010[0] = πTemp037
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ß_deprecate); πE != nil {
						continue
					}
					if πTemp038, πE = πTemp037.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßfailUnless.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 637: failUnlessRaises = _deprecate(assertRaises)
					πF.SetLineno(637)
					πTemp010 = πF.MakeArgs(1)
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßassertRaises); πE != nil {
						continue
					}
					πTemp010[0] = πTemp037
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ß_deprecate); πE != nil {
						continue
					}
					if πTemp038, πE = πTemp037.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßfailUnlessRaises.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 638: failIf = _deprecate(assertFalse)
					πF.SetLineno(638)
					πTemp010 = πF.MakeArgs(1)
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßassertFalse); πE != nil {
						continue
					}
					πTemp010[0] = πTemp037
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ß_deprecate); πE != nil {
						continue
					}
					if πTemp038, πE = πTemp037.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					if πE = πClass.SetItem(πF, ßfailIf.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 640: def assertSequenceEqual(self, seq1, seq2, msg=None, seq_type=None):
					πF.SetLineno(640)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "seq1", Def: nil}
					πTemp002[2] = πg.Param{Name: "seq2", Def: nil}
					if πTemp038, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp038}
					if πTemp038, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "seq_type", Def: πTemp038}
					πTemp037 = πg.NewFunction(πg.NewCode("assertSequenceEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µseq1 *πg.Object = πArgs[1]; _ = µseq1
						var µseq2 *πg.Object = πArgs[2]; _ = µseq2
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µseq_type *πg.Object = πArgs[4]; _ = µseq_type
						var µseq_type_name *πg.Object = πg.UnboundLocal; _ = µseq_type_name
						var µdiffering *πg.Object = πg.UnboundLocal; _ = µdiffering
						var µlen1 *πg.Object = πg.UnboundLocal; _ = µlen1
						var µlen2 *πg.Object = πg.UnboundLocal; _ = µlen2
						var µseq1_repr *πg.Object = πg.UnboundLocal; _ = µseq1_repr
						var µseq2_repr *πg.Object = πg.UnboundLocal; _ = µseq2_repr
						var µelements *πg.Object = πg.UnboundLocal; _ = µelements
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µitem1 *πg.Object = πg.UnboundLocal; _ = µitem1
						var µitem2 *πg.Object = πg.UnboundLocal; _ = µitem2
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
						var µdiffMsg *πg.Object = πg.UnboundLocal; _ = µdiffMsg
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
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
							case 9: goto Label9
							case 42: goto Label42
							case 45: goto Label45
							case 14: goto Label14
							case 24: goto Label24
							case 25: goto Label25
							case 28: goto Label28
							case 31: goto Label31
							default: panic("unexpected function state")
							}
							// line 641: """An equality assertion for ordered sequences (like lists and tuples).
							πF.SetLineno(641)
							if πE = πg.CheckLocal(πF, µseq_type, "seq_type"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µseq_type != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 654: if seq_type is not None:
							πF.SetLineno(654)
						Label1:
							// line 655: seq_type_name = seq_type.__name__
							πF.SetLineno(655)
							if πE = πg.CheckLocal(πF, µseq_type, "seq_type"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µseq_type, ß__name__, nil); πE != nil {
								continue
							}
							µseq_type_name = πTemp001
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µseq1, "seq1"); πE != nil {
								continue
							}
							πTemp004[0] = µseq1
							if πE = πg.CheckLocal(πF, µseq_type, "seq_type"); πE != nil {
								continue
							}
							πTemp004[1] = µseq_type
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label4
							}
							goto Label5
							// line 656: if not isinstance(seq1, seq_type):
							πF.SetLineno(656)
						Label4:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq1, "seq1"); πE != nil {
								continue
							}
							πTemp006[0] = µseq1
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002 = πg.NewTuple2(µseq_type_name, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("First sequence is not a %s: %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 657: raise self.failureException('First sequence is not a %s: %s'
							πF.SetLineno(657)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label5
						Label5:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µseq2, "seq2"); πE != nil {
								continue
							}
							πTemp004[0] = µseq2
							if πE = πg.CheckLocal(πF, µseq_type, "seq_type"); πE != nil {
								continue
							}
							πTemp004[1] = µseq_type
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
								goto Label6
							}
							goto Label7
							// line 659: if not isinstance(seq2, seq_type):
							πF.SetLineno(659)
						Label6:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq2, "seq2"); πE != nil {
								continue
							}
							πTemp006[0] = µseq2
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp002 = πg.NewTuple2(µseq_type_name, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Second sequence is not a %s: %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 660: raise self.failureException('Second sequence is not a %s: %s'
							πF.SetLineno(660)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label7
						Label7:
							goto Label3
						Label2:
							// line 663: seq_type_name = "sequence"
							πF.SetLineno(663)
							µseq_type_name = ßsequence.ToObject()
							goto Label3
						Label3:
							// line 665: differing = None
							πF.SetLineno(665)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µdiffering = πTemp001
							// line 666: try:
							πF.SetLineno(666)
							πF.PushCheckpoint(9)
							// line 667: len1 = len(seq1)
							πF.SetLineno(667)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq1, "seq1"); πE != nil {
								continue
							}
							πTemp004[0] = µseq1
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µlen1 = πTemp002
							πF.PopCheckpoint()
							goto Label8
						Label9:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp005).ToObject()
							if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 668: except (TypeError, NotImplementedError):
							πF.SetLineno(668)
						Label10:
							// line 669: differing = 'First %s has no length.    Non-sequence?' % (
							πF.SetLineno(669)
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("First %s has no length.    Non-sequence?").ToObject(), µseq_type_name); πE != nil {
								continue
							}
							µdiffering = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdiffering == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							goto Label12
							// line 672: if differing is None:
							πF.SetLineno(672)
						Label11:
							// line 673: try:
							πF.SetLineno(673)
							πF.PushCheckpoint(14)
							// line 674: len2 = len(seq2)
							πF.SetLineno(674)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq2, "seq2"); πE != nil {
								continue
							}
							πTemp004[0] = µseq2
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µlen2 = πTemp002
							πF.PopCheckpoint()
							goto Label13
						Label14:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp005).ToObject()
							if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label15
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 675: except (TypeError, NotImplementedError):
							πF.SetLineno(675)
						Label15:
							// line 676: differing = 'Second %s has no length.    Non-sequence?' % (
							πF.SetLineno(676)
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Second %s has no length.    Non-sequence?").ToObject(), µseq_type_name); πE != nil {
								continue
							}
							µdiffering = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label13
						Label13:
							goto Label12
						Label12:
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdiffering == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label16
							}
							goto Label17
							// line 679: if differing is None:
							πF.SetLineno(679)
						Label16:
							if πE = πg.CheckLocal(πF, µseq1, "seq1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq2, "seq2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µseq1, µseq2); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label18
							}
							goto Label19
							// line 680: if seq1 == seq2:
							πF.SetLineno(680)
						Label18:
							// line 681: return
							πF.SetLineno(681)
							πR = πg.None
							continue
							goto Label19
						Label19:
							// line 683: seq1_repr = safe_repr(seq1)
							πF.SetLineno(683)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq1, "seq1"); πE != nil {
								continue
							}
							πTemp004[0] = µseq1
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µseq1_repr = πTemp002
							// line 684: seq2_repr = safe_repr(seq2)
							πF.SetLineno(684)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq2, "seq2"); πE != nil {
								continue
							}
							πTemp004[0] = µseq2
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µseq2_repr = πTemp002
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq1_repr, "seq1_repr"); πE != nil {
								continue
							}
							πTemp004[0] = µseq1_repr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(30).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label20
							}
							goto Label21
							// line 685: if len(seq1_repr) > 30:
							πF.SetLineno(685)
						Label20:
							// line 686: seq1_repr = seq1_repr[:30] + '...'
							πF.SetLineno(686)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(30).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq1_repr, "seq1_repr"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µseq1_repr, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp005, πg.NewStr("...").ToObject()); πE != nil {
								continue
							}
							µseq1_repr = πTemp001
							goto Label21
						Label21:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq2_repr, "seq2_repr"); πE != nil {
								continue
							}
							πTemp004[0] = µseq2_repr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.GT(πF, πTemp005, πg.NewInt(30).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label22
							}
							goto Label23
							// line 687: if len(seq2_repr) > 30:
							πF.SetLineno(687)
						Label22:
							// line 688: seq2_repr = seq2_repr[:30] + '...'
							πF.SetLineno(688)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(30).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq2_repr, "seq2_repr"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µseq2_repr, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp005, πg.NewStr("...").ToObject()); πE != nil {
								continue
							}
							µseq2_repr = πTemp001
							goto Label23
						Label23:
							// line 689: elements = (seq_type_name.capitalize(), seq1_repr, seq2_repr)
							πF.SetLineno(689)
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µseq_type_name, ßcapitalize, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq1_repr, "seq1_repr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq2_repr, "seq2_repr"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp005, µseq1_repr, µseq2_repr).ToObject()
							µelements = πTemp001
							// line 690: differing = '%ss differ: %s != %s\n' % elements
							πF.SetLineno(690)
							if πE = πg.CheckLocal(πF, µelements, "elements"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%ss differ: %s != %s\n").ToObject(), µelements); πE != nil {
								continue
							}
							µdiffering = πTemp001
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlen1, "len1"); πE != nil {
								continue
							}
							πTemp006[0] = µlen1
							if πE = πg.CheckLocal(πF, µlen2, "len2"); πE != nil {
								continue
							}
							πTemp006[1] = µlen2
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp005
							if πTemp002, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(25)
							πTemp003 = false
						Label24:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label26
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(24)            
							// line 693: try:
							πF.SetLineno(693)
							πF.PushCheckpoint(28)
							// line 694: item1 = seq1[i]
							πF.SetLineno(694)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µseq1, "seq1"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µseq1, πTemp002); πE != nil {
								continue
							}
							µitem1 = πTemp005
							πF.PopCheckpoint()
							goto Label27
						Label28:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp005, πTemp007, πTemp011).ToObject()
							if πTemp010, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label29
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 695: except (TypeError, IndexError, NotImplementedError):
							πF.SetLineno(695)
						Label29:
							// line 696: differing += ('\nUnable to index element %d of first %s\n' %
							πF.SetLineno(696)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µi, µseq_type_name).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\nUnable to index element %d of first %s\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µdiffering, πTemp002); πE != nil {
								continue
							}
							µdiffering = πTemp005
							// line 698: break
							πF.SetLineno(698)
							πTemp003 = true
							continue
							πF.RestoreExc(nil, nil)
							goto Label27
						Label27:
							// line 700: try:
							πF.SetLineno(700)
							πF.PushCheckpoint(31)
							// line 701: item2 = seq2[i]
							πF.SetLineno(701)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µseq2, "seq2"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µseq2, πTemp002); πE != nil {
								continue
							}
							µitem2 = πTemp005
							πF.PopCheckpoint()
							goto Label30
						Label31:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp005, πTemp007, πTemp011).ToObject()
							if πTemp010, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label32
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 702: except (TypeError, IndexError, NotImplementedError):
							πF.SetLineno(702)
						Label32:
							// line 703: differing += ('\nUnable to index element %d of second %s\n' %
							πF.SetLineno(703)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(µi, µseq_type_name).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\nUnable to index element %d of second %s\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µdiffering, πTemp002); πE != nil {
								continue
							}
							µdiffering = πTemp005
							// line 705: break
							πF.SetLineno(705)
							πTemp003 = true
							continue
							πF.RestoreExc(nil, nil)
							goto Label30
						Label30:
							if πE = πg.CheckLocal(πF, µitem1, "item1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µitem2, "item2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µitem1, µitem2); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label33
							}
							goto Label34
							// line 707: if item1 != item2:
							πF.SetLineno(707)
						Label33:
							// line 708: differing += ('\nFirst differing element %d:\n%s\n%s\n' %
							πF.SetLineno(708)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem1, "item1"); πE != nil {
								continue
							}
							πTemp004[0] = µitem1
							if πTemp007, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem2, "item2"); πE != nil {
								continue
							}
							πTemp004[0] = µitem2
							if πTemp007, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp005 = πg.NewTuple3(µi, πTemp011, πTemp012).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("\nFirst differing element %d:\n%s\n%s\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, µdiffering, πTemp002); πE != nil {
								continue
							}
							µdiffering = πTemp005
							// line 710: break
							πF.SetLineno(710)
							πTemp003 = true
							continue
							goto Label34
						Label34:
							continue
						Label25:
							if πE != nil || πR != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen1, "len1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen2, "len2"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, µlen1, µlen2); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label35
							}
							if πE = πg.CheckLocal(πF, µseq_type, "seq_type"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(µseq_type == πTemp007).ToObject()
							πTemp002 = πTemp005
							if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label35
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq1, "seq1"); πE != nil {
								continue
							}
							πTemp004[0] = µseq1
							if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq2, "seq2"); πE != nil {
								continue
							}
							πTemp004[0] = µseq2
							if πTemp007, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.NE(πF, πTemp011, πTemp012); πE != nil {
								continue
							}
							πTemp002 = πTemp005
						Label35:
							if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label36
							}
							goto Label37
							// line 712: if (len1 == len2 and seq_type is None and
							πF.SetLineno(712)
						Label36:
							// line 715: return
							πF.SetLineno(715)
							πR = πg.None
							continue
							goto Label37
						Label37:
						Label26:
							if πE = πg.CheckLocal(πF, µlen1, "len1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen2, "len2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µlen1, µlen2); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label38
							}
							if πE = πg.CheckLocal(πF, µlen1, "len1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen2, "len2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µlen1, µlen2); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label39
							}
							goto Label40
							// line 717: if len1 > len2:
							πF.SetLineno(717)
						Label38:
							// line 718: differing += ('\nFirst %s contains %d additional '
							πF.SetLineno(718)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen1, "len1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen2, "len2"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Sub(πF, µlen1, µlen2); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µseq_type_name, πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\nFirst %s contains %d additional elements.\n").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µdiffering, πTemp001); πE != nil {
								continue
							}
							µdiffering = πTemp002
							// line 720: try:
							πF.SetLineno(720)
							πF.PushCheckpoint(42)
							// line 721: differing += ('First extra element %d:\n%s\n' %
							πF.SetLineno(721)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen2, "len2"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlen2, "len2"); πE != nil {
								continue
							}
							πTemp005 = µlen2
							if πE = πg.CheckLocal(πF, µseq1, "seq1"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µseq1, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πg.NewTuple2(µlen2, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("First extra element %d:\n%s\n").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µdiffering, πTemp001); πE != nil {
								continue
							}
							µdiffering = πTemp002
							πF.PopCheckpoint()
							goto Label41
						Label42:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp005, πTemp007).ToObject()
							if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label43
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 723: except (TypeError, IndexError, NotImplementedError):
							πF.SetLineno(723)
						Label43:
							// line 724: differing += ('Unable to index element %d '
							πF.SetLineno(724)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen2, "len2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µlen2, µseq_type_name).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Unable to index element %d of first %s\n").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µdiffering, πTemp001); πE != nil {
								continue
							}
							µdiffering = πTemp002
							πF.RestoreExc(nil, nil)
							goto Label41
						Label41:
							goto Label40
							// line 726: elif len1 < len2:
							πF.SetLineno(726)
						Label39:
							// line 727: differing += ('\nSecond %s contains %d additional '
							πF.SetLineno(727)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen2, "len2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen1, "len1"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Sub(πF, µlen2, µlen1); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µseq_type_name, πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("\nSecond %s contains %d additional elements.\n").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µdiffering, πTemp001); πE != nil {
								continue
							}
							µdiffering = πTemp002
							// line 729: try:
							πF.SetLineno(729)
							πF.PushCheckpoint(45)
							// line 730: differing += ('First extra element %d:\n%s\n' %
							πF.SetLineno(730)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen1, "len1"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlen1, "len1"); πE != nil {
								continue
							}
							πTemp005 = µlen1
							if πE = πg.CheckLocal(πF, µseq2, "seq2"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µseq2, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp007
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πg.NewTuple2(µlen1, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("First extra element %d:\n%s\n").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µdiffering, πTemp001); πE != nil {
								continue
							}
							µdiffering = πTemp002
							πF.PopCheckpoint()
							goto Label44
						Label45:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp005, πTemp007).ToObject()
							if πTemp003, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label46
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 732: except (TypeError, IndexError, NotImplementedError):
							πF.SetLineno(732)
						Label46:
							// line 733: differing += ('Unable to index element %d '
							πF.SetLineno(733)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlen1, "len1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µseq_type_name, "seq_type_name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µlen1, µseq_type_name).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Unable to index element %d of second %s\n").ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µdiffering, πTemp001); πE != nil {
								continue
							}
							µdiffering = πTemp002
							πF.RestoreExc(nil, nil)
							goto Label44
						Label44:
							goto Label40
						Label40:
							goto Label17
						Label17:
							// line 735: standardMsg = differing
							πF.SetLineno(735)
							if πE = πg.CheckLocal(πF, µdiffering, "differing"); πE != nil {
								continue
							}
							µstandardMsg = µdiffering
							// line 736: diffMsg = '\n' + '\n'.join(
							πF.SetLineno(736)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq1, "seq1"); πE != nil {
								continue
							}
							πTemp013[0] = µseq1
							if πTemp002, πE = πg.ResolveGlobal(πF, ßpprint); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							πTemp013 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µseq2, "seq2"); πE != nil {
								continue
							}
							πTemp013[0] = µseq2
							if πTemp002, πE = πg.ResolveGlobal(πF, ßpprint); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp013, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp013)
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdifflib); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßndiff, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							µdiffMsg = πTemp001
							// line 739: standardMsg = self._truncateMessage(standardMsg, diffMsg)
							πF.SetLineno(739)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp004[0] = µstandardMsg
							if πE = πg.CheckLocal(πF, µdiffMsg, "diffMsg"); πE != nil {
								continue
							}
							πTemp004[1] = µdiffMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_truncateMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µstandardMsg = πTemp002
							// line 740: msg = self._formatMessage(msg, standardMsg)
							πF.SetLineno(740)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp004[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp004[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µmsg = πTemp002
							// line 741: self.fail(msg)
							πF.SetLineno(741)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp004[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertSequenceEqual.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 641: """An equality assertion for ordered sequences (like lists and tuples).
					πF.SetLineno(641)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp038}, πg.NewStr("An equality assertion for ordered sequences (like lists and tuples).\n\n        For the purposes of this function, a valid ordered sequence type is one\n        which can be indexed, has a length, and has an equality operator.\n\n        Args:\n            seq1: The first sequence to compare.\n            seq2: The second sequence to compare.\n            seq_type: The expected datatype of the sequences, or None if no\n                    datatype should be enforced.\n            msg: Optional message to use on failure instead of a list of\n                    differences.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp039, πE = πg.ResolveClass(πF, πClass, nil, ßassertSequenceEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp039, ß__doc__, πTemp038); πE != nil {
						continue
					}
					// line 743: def _truncateMessage(self, message, diff):
					πF.SetLineno(743)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "message", Def: nil}
					πTemp002[2] = πg.Param{Name: "diff", Def: nil}
					πTemp038 = πg.NewFunction(πg.NewCode("_truncateMessage", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmessage *πg.Object = πArgs[1]; _ = µmessage
						var µdiff *πg.Object = πArgs[2]; _ = µdiff
						var µmax_diff *πg.Object = πg.UnboundLocal; _ = µmax_diff
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
							// line 744: max_diff = self.maxDiff
							πF.SetLineno(744)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmaxDiff, nil); πE != nil {
								continue
							}
							µmax_diff = πTemp001
							if πE = πg.CheckLocal(πF, µmax_diff, "max_diff"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µmax_diff == πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
								continue
							}
							πTemp005[0] = µdiff
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µmax_diff, "max_diff"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LE(πF, πTemp006, µmax_diff); πE != nil {
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
							// line 745: if max_diff is None or len(diff) <= max_diff:
							πF.SetLineno(745)
						Label2:
							// line 746: return message + diff
							πF.SetLineno(746)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µmessage, µdiff); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label3:
							// line 747: return message + (DIFF_OMITTED % len(diff))
							πF.SetLineno(747)
							if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßDIFF_OMITTED); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
								continue
							}
							πTemp005[0] = µdiff
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.Mod(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µmessage, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_truncateMessage.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 749: def assertListEqual(self, list1, list2, msg=None):
					πF.SetLineno(749)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "list1", Def: nil}
					πTemp002[2] = πg.Param{Name: "list2", Def: nil}
					if πTemp040, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp040}
					πTemp039 = πg.NewFunction(πg.NewCode("assertListEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlist1 *πg.Object = πArgs[1]; _ = µlist1
						var µlist2 *πg.Object = πArgs[2]; _ = µlist2
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
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
							// line 750: """A list-specific equality assertion.
							πF.SetLineno(750)
							// line 759: self.assertSequenceEqual(list1, list2, msg, seq_type=list)
							πF.SetLineno(759)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µlist1, "list1"); πE != nil {
								continue
							}
							πTemp001[0] = µlist1
							if πE = πg.CheckLocal(πF, µlist2, "list2"); πE != nil {
								continue
							}
							πTemp001[1] = µlist2
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[2] = µmsg
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"seq_type", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertSequenceEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertListEqual.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 750: """A list-specific equality assertion.
					πF.SetLineno(750)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp040}, πg.NewStr("A list-specific equality assertion.\n\n        Args:\n            list1: The first list to compare.\n            list2: The second list to compare.\n            msg: Optional message to use on failure instead of a list of\n                    differences.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp041, πE = πg.ResolveClass(πF, πClass, nil, ßassertListEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp041, ß__doc__, πTemp040); πE != nil {
						continue
					}
					// line 761: def assertTupleEqual(self, tuple1, tuple2, msg=None):
					πF.SetLineno(761)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "tuple1", Def: nil}
					πTemp002[2] = πg.Param{Name: "tuple2", Def: nil}
					if πTemp041, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp041}
					πTemp040 = πg.NewFunction(πg.NewCode("assertTupleEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtuple1 *πg.Object = πArgs[1]; _ = µtuple1
						var µtuple2 *πg.Object = πArgs[2]; _ = µtuple2
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 πg.KWArgs
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
							// line 762: """A tuple-specific equality assertion.
							πF.SetLineno(762)
							// line 770: self.assertSequenceEqual(tuple1, tuple2, msg, seq_type=tuple)
							πF.SetLineno(770)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µtuple1, "tuple1"); πE != nil {
								continue
							}
							πTemp001[0] = µtuple1
							if πE = πg.CheckLocal(πF, µtuple2, "tuple2"); πE != nil {
								continue
							}
							πTemp001[1] = µtuple2
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[2] = µmsg
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp003 = πg.KWArgs{
								{"seq_type", πTemp002},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßassertSequenceEqual, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertTupleEqual.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 762: """A tuple-specific equality assertion.
					πF.SetLineno(762)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp041}, πg.NewStr("A tuple-specific equality assertion.\n\n        Args:\n            tuple1: The first tuple to compare.\n            tuple2: The second tuple to compare.\n            msg: Optional message to use on failure instead of a list of\n                    differences.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp042, πE = πg.ResolveClass(πF, πClass, nil, ßassertTupleEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp042, ß__doc__, πTemp041); πE != nil {
						continue
					}
					// line 772: def assertSetEqual(self, set1, set2, msg=None):
					πF.SetLineno(772)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "set1", Def: nil}
					πTemp002[2] = πg.Param{Name: "set2", Def: nil}
					if πTemp042, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp042}
					πTemp041 = πg.NewFunction(πg.NewCode("assertSetEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µset1 *πg.Object = πArgs[1]; _ = µset1
						var µset2 *πg.Object = πArgs[2]; _ = µset2
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µdifference1 *πg.Object = πg.UnboundLocal; _ = µdifference1
						var µe *πg.Object = πg.UnboundLocal; _ = µe
						var µdifference2 *πg.Object = πg.UnboundLocal; _ = µdifference2
						var µlines *πg.Object = πg.UnboundLocal; _ = µlines
						var µitem *πg.Object = πg.UnboundLocal; _ = µitem
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							case 6: goto Label6
							case 14: goto Label14
							case 15: goto Label15
							case 19: goto Label19
							case 20: goto Label20
							default: panic("unexpected function state")
							}
							// line 773: """A set-specific equality assertion.
							πF.SetLineno(773)
							// line 785: try:
							πF.SetLineno(785)
							πF.PushCheckpoint(2)
							// line 786: difference1 = set1.difference(set2)
							πF.SetLineno(786)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µset2, "set2"); πE != nil {
								continue
							}
							πTemp001[0] = µset2
							if πE = πg.CheckLocal(πF, µset1, "set1"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µset1, ßdifference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdifference1 = πTemp003
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
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 787: except TypeError, e:
							πF.SetLineno(787)
						Label3:
							µe = πTemp004.ToObject()
							// line 788: self.fail('invalid type when attempting set difference: %s' % e)
							πF.SetLineno(788)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("invalid type when attempting set difference: %s").ToObject(), µe); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							πF.RestoreExc(nil, nil)
							goto Label1
							// line 789: except AttributeError, e:
							πF.SetLineno(789)
						Label4:
							µe = πTemp004.ToObject()
							// line 790: self.fail('first argument does not support set difference: %s' % e)
							πF.SetLineno(790)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("first argument does not support set difference: %s").ToObject(), µe); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 792: try:
							πF.SetLineno(792)
							πF.PushCheckpoint(6)
							// line 793: difference2 = set2.difference(set1)
							πF.SetLineno(793)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µset1, "set1"); πE != nil {
								continue
							}
							πTemp001[0] = µset1
							if πE = πg.CheckLocal(πF, µset2, "set2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µset2, ßdifference, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdifference2 = πTemp003
							πF.PopCheckpoint()
							goto Label5
						Label6:
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
								goto Label7
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 794: except TypeError, e:
							πF.SetLineno(794)
						Label7:
							µe = πTemp004.ToObject()
							// line 795: self.fail('invalid type when attempting set difference: %s' % e)
							πF.SetLineno(795)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("invalid type when attempting set difference: %s").ToObject(), µe); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							πF.RestoreExc(nil, nil)
							goto Label5
							// line 796: except AttributeError, e:
							πF.SetLineno(796)
						Label8:
							µe = πTemp004.ToObject()
							// line 797: self.fail('second argument does not support set difference: %s' % e)
							πF.SetLineno(797)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("second argument does not support set difference: %s").ToObject(), µe); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
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
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µdifference1, "difference1"); πE != nil {
								continue
							}
							πTemp003 = µdifference1
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µdifference2, "difference2"); πE != nil {
								continue
							}
							πTemp003 = µdifference2
						Label9:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label10
							}
							goto Label11
							// line 799: if not (difference1 or difference2):
							πF.SetLineno(799)
						Label10:
							// line 800: return
							πF.SetLineno(800)
							πR = πg.None
							continue
							goto Label11
						Label11:
							// line 802: lines = []
							πF.SetLineno(802)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µlines = πTemp002
							if πE = πg.CheckLocal(πF, µdifference1, "difference1"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µdifference1); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label12
							}
							goto Label13
							// line 803: if difference1:
							πF.SetLineno(803)
						Label12:
							// line 804: lines.append('Items in the first set but not the second:')
							πF.SetLineno(804)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Items in the first set but not the second:").ToObject()
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µdifference1, "difference1"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µdifference1); πE != nil {
								continue
							}
							πF.PushCheckpoint(15)
							πTemp006 = false
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label16
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
								µitem = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(14)            
							// line 806: lines.append(repr(item))
							πF.SetLineno(806)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp008[0] = µitem
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp009
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
						Label16:
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µdifference2, "difference2"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µdifference2); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label17
							}
							goto Label18
							// line 807: if difference2:
							πF.SetLineno(807)
						Label17:
							// line 808: lines.append('Items in the second set but not the first:')
							πF.SetLineno(808)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("Items in the second set but not the first:").ToObject()
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µdifference2, "difference2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µdifference2); πE != nil {
								continue
							}
							πF.PushCheckpoint(20)
							πTemp006 = false
						Label19:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label21
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
								µitem = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(19)            
							// line 810: lines.append(repr(item))
							πF.SetLineno(810)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp008[0] = µitem
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp009
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µlines, ßappend, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label20:
							if πE != nil || πR != nil {
								continue
							}
						Label21:
							goto Label18
						Label18:
							// line 812: standardMsg = '\n'.join(lines)
							πF.SetLineno(812)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp001[0] = µlines
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstandardMsg = πTemp003
							// line 813: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(813)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp003
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßassertSetEqual.ToObject(), πTemp041); πE != nil {
						continue
					}
					// line 773: """A set-specific equality assertion.
					πF.SetLineno(773)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp042}, πg.NewStr("A set-specific equality assertion.\n\n        Args:\n            set1: The first set to compare.\n            set2: The second set to compare.\n            msg: Optional message to use on failure instead of a list of\n                    differences.\n\n        assertSetEqual uses ducktyping to support different types of sets, and\n        is optimized for sets specifically (parameters must support a\n        difference method).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp043, πE = πg.ResolveClass(πF, πClass, nil, ßassertSetEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp043, ß__doc__, πTemp042); πE != nil {
						continue
					}
					// line 815: def assertIn(self, member, container, msg=None):
					πF.SetLineno(815)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "member", Def: nil}
					πTemp002[2] = πg.Param{Name: "container", Def: nil}
					if πTemp043, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp043}
					πTemp042 = πg.NewFunction(πg.NewCode("assertIn", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmember *πg.Object = πArgs[1]; _ = µmember
						var µcontainer *πg.Object = πArgs[2]; _ = µcontainer
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp007 *πg.Object
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
							// line 816: """Just like self.assertTrue(a in b), but with a nicer default message."""
							πF.SetLineno(816)
							if πE = πg.CheckLocal(πF, µmember, "member"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontainer, "container"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µcontainer, µmember); πE != nil {
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
							// line 817: if member not in container:
							πF.SetLineno(817)
						Label1:
							// line 818: standardMsg = '%s not found in %s' % (safe_repr(member),
							πF.SetLineno(818)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmember, "member"); πE != nil {
								continue
							}
							πTemp004[0] = µmember
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontainer, "container"); πE != nil {
								continue
							}
							πTemp004[0] = µcontainer
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s not found in %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 820: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(820)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertIn.ToObject(), πTemp042); πE != nil {
						continue
					}
					// line 816: """Just like self.assertTrue(a in b), but with a nicer default message."""
					πF.SetLineno(816)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp043}, πg.NewStr("Just like self.assertTrue(a in b), but with a nicer default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp044, πE = πg.ResolveClass(πF, πClass, nil, ßassertIn); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp044, ß__doc__, πTemp043); πE != nil {
						continue
					}
					// line 822: def assertNotIn(self, member, container, msg=None):
					πF.SetLineno(822)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "member", Def: nil}
					πTemp002[2] = πg.Param{Name: "container", Def: nil}
					if πTemp044, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp044}
					πTemp043 = πg.NewFunction(πg.NewCode("assertNotIn", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmember *πg.Object = πArgs[1]; _ = µmember
						var µcontainer *πg.Object = πArgs[2]; _ = µcontainer
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp007 *πg.Object
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
							// line 823: """Just like self.assertTrue(a not in b), but with a nicer default message."""
							πF.SetLineno(823)
							if πE = πg.CheckLocal(πF, µmember, "member"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontainer, "container"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µcontainer, µmember); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 824: if member in container:
							πF.SetLineno(824)
						Label1:
							// line 825: standardMsg = '%s unexpectedly found in %s' % (safe_repr(member),
							πF.SetLineno(825)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmember, "member"); πE != nil {
								continue
							}
							πTemp004[0] = µmember
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontainer, "container"); πE != nil {
								continue
							}
							πTemp004[0] = µcontainer
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s unexpectedly found in %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 827: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(827)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertNotIn.ToObject(), πTemp043); πE != nil {
						continue
					}
					// line 823: """Just like self.assertTrue(a not in b), but with a nicer default message."""
					πF.SetLineno(823)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp044}, πg.NewStr("Just like self.assertTrue(a not in b), but with a nicer default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp045, πE = πg.ResolveClass(πF, πClass, nil, ßassertNotIn); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp045, ß__doc__, πTemp044); πE != nil {
						continue
					}
					// line 829: def assertIs(self, expr1, expr2, msg=None):
					πF.SetLineno(829)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "expr1", Def: nil}
					πTemp002[2] = πg.Param{Name: "expr2", Def: nil}
					if πTemp045, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp045}
					πTemp044 = πg.NewFunction(πg.NewCode("assertIs", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpr1 *πg.Object = πArgs[1]; _ = µexpr1
						var µexpr2 *πg.Object = πArgs[2]; _ = µexpr2
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp007 *πg.Object
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
							// line 830: """Just like self.assertTrue(a is b), but with a nicer default message."""
							πF.SetLineno(830)
							if πE = πg.CheckLocal(πF, µexpr1, "expr1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexpr2, "expr2"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µexpr1 != µexpr2).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 831: if expr1 is not expr2:
							πF.SetLineno(831)
						Label1:
							// line 832: standardMsg = '%s is not %s' % (safe_repr(expr1),
							πF.SetLineno(832)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpr1, "expr1"); πE != nil {
								continue
							}
							πTemp004[0] = µexpr1
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpr2, "expr2"); πE != nil {
								continue
							}
							πTemp004[0] = µexpr2
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s is not %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 834: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(834)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertIs.ToObject(), πTemp044); πE != nil {
						continue
					}
					// line 830: """Just like self.assertTrue(a is b), but with a nicer default message."""
					πF.SetLineno(830)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp045}, πg.NewStr("Just like self.assertTrue(a is b), but with a nicer default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp046, πE = πg.ResolveClass(πF, πClass, nil, ßassertIs); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp046, ß__doc__, πTemp045); πE != nil {
						continue
					}
					// line 836: def assertIsNot(self, expr1, expr2, msg=None):
					πF.SetLineno(836)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "expr1", Def: nil}
					πTemp002[2] = πg.Param{Name: "expr2", Def: nil}
					if πTemp046, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp046}
					πTemp045 = πg.NewFunction(πg.NewCode("assertIsNot", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpr1 *πg.Object = πArgs[1]; _ = µexpr1
						var µexpr2 *πg.Object = πArgs[2]; _ = µexpr2
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 837: """Just like self.assertTrue(a is not b), but with a nicer default message."""
							πF.SetLineno(837)
							if πE = πg.CheckLocal(πF, µexpr1, "expr1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexpr2, "expr2"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µexpr1 == µexpr2).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 838: if expr1 is expr2:
							πF.SetLineno(838)
						Label1:
							// line 839: standardMsg = 'unexpectedly identical: %s' % (safe_repr(expr1),)
							πF.SetLineno(839)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpr1, "expr1"); πE != nil {
								continue
							}
							πTemp004[0] = µexpr1
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πg.NewTuple1(πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("unexpectedly identical: %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 840: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(840)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp007[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp007[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertIsNot.ToObject(), πTemp045); πE != nil {
						continue
					}
					// line 837: """Just like self.assertTrue(a is not b), but with a nicer default message."""
					πF.SetLineno(837)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp046}, πg.NewStr("Just like self.assertTrue(a is not b), but with a nicer default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp047, πE = πg.ResolveClass(πF, πClass, nil, ßassertIsNot); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp047, ß__doc__, πTemp046); πE != nil {
						continue
					}
					// line 842: def assertDictEqual(self, d1, d2, msg=None):
					πF.SetLineno(842)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "d1", Def: nil}
					πTemp002[2] = πg.Param{Name: "d2", Def: nil}
					if πTemp047, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp047}
					πTemp046 = πg.NewFunction(πg.NewCode("assertDictEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µd1 *πg.Object = πArgs[1]; _ = µd1
						var µd2 *πg.Object = πArgs[2]; _ = µd2
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
						var µdiff *πg.Object = πg.UnboundLocal; _ = µdiff
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
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
							// line 843: self.assertIsInstance(d1, dict, 'First argument is not a dictionary')
							πF.SetLineno(843)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µd1, "d1"); πE != nil {
								continue
							}
							πTemp001[0] = µd1
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewStr("First argument is not a dictionary").ToObject()
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
							// line 844: self.assertIsInstance(d2, dict, 'Second argument is not a dictionary')
							πF.SetLineno(844)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µd2, "d2"); πE != nil {
								continue
							}
							πTemp001[0] = µd2
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewStr("Second argument is not a dictionary").ToObject()
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
							if πE = πg.CheckLocal(πF, µd1, "d1"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd2, "d2"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µd1, µd2); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 846: if d1 != d2:
							πF.SetLineno(846)
						Label1:
							// line 847: standardMsg = '%s != %s' % (safe_repr(d1, True), safe_repr(d2, True))
							πF.SetLineno(847)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd1, "d1"); πE != nil {
								continue
							}
							πTemp001[0] = µd1
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µd2, "d2"); πE != nil {
								continue
							}
							πTemp001[0] = µd2
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s != %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp002
							// line 848: diff = ('\n' + '\n'.join(difflib.ndiff(
							πF.SetLineno(848)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd1, "d1"); πE != nil {
								continue
							}
							πTemp009[0] = µd1
							if πTemp003, πE = πg.ResolveGlobal(πF, ßpprint); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp003
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd2, "d2"); πE != nil {
								continue
							}
							πTemp009[0] = µd2
							if πTemp003, πE = πg.ResolveGlobal(πF, ßpprint); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßpformat, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp008[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdifflib); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßndiff, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							µdiff = πTemp002
							// line 851: standardMsg = self._truncateMessage(standardMsg, diff)
							πF.SetLineno(851)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp001[0] = µstandardMsg
							if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
								continue
							}
							πTemp001[1] = µdiff
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_truncateMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstandardMsg = πTemp003
							// line 852: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(852)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp003
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
					if πE = πClass.SetItem(πF, ßassertDictEqual.ToObject(), πTemp046); πE != nil {
						continue
					}
					// line 854: def assertDictContainsSubset(self, expected, actual, msg=None):
					πF.SetLineno(854)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "expected", Def: nil}
					πTemp002[2] = πg.Param{Name: "actual", Def: nil}
					if πTemp048, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp048}
					πTemp047 = πg.NewFunction(πg.NewCode("assertDictContainsSubset", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpected *πg.Object = πArgs[1]; _ = µexpected
						var µactual *πg.Object = πArgs[2]; _ = µactual
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µmissing *πg.Object = πg.UnboundLocal; _ = µmissing
						var µmismatched *πg.Object = πg.UnboundLocal; _ = µmismatched
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 []πg.Param
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
							// line 855: """Checks whether actual is a superset of expected."""
							πF.SetLineno(855)
							// line 856: missing = []
							πF.SetLineno(856)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µmissing = πTemp002
							// line 857: mismatched = []
							πF.SetLineno(857)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µmismatched = πTemp002
							if πE = πg.CheckLocal(πF, µexpected, "expected"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µexpected, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
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
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µactual, µkey); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp004 = µkey
							if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µactual, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, µvalue, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 859: if key not in actual:
							πF.SetLineno(859)
						Label4:
							// line 860: missing.append(key)
							πF.SetLineno(860)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmissing, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
							// line 861: elif value != actual[key]:
							πF.SetLineno(861)
						Label5:
							// line 862: mismatched.append('%s, expected: %s, actual: %s' %
							πF.SetLineno(862)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp008[0] = µkey
							if πTemp007, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp008[0] = µvalue
							if πTemp007, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007 = µkey
							if πE = πg.CheckLocal(πF, µactual, "actual"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetItem(πF, µactual, πTemp007); πE != nil {
								continue
							}
							πTemp008[0] = πTemp011
							if πTemp007, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004 = πg.NewTuple3(πTemp009, πTemp010, πTemp011).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("%s, expected: %s, actual: %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µmismatched, "mismatched"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µmismatched, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
								continue
							}
							πTemp003 = µmissing
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µmismatched, "mismatched"); πE != nil {
								continue
							}
							πTemp003 = µmismatched
						Label7:
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label8
							}
							goto Label9
							// line 866: if not (missing or mismatched):
							πF.SetLineno(866)
						Label8:
							// line 867: return
							πF.SetLineno(867)
							πR = πg.None
							continue
							goto Label9
						Label9:
							// line 869: standardMsg = ''
							πF.SetLineno(869)
							µstandardMsg = ß.ToObject()
							if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µmissing); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 870: if missing:
							πF.SetLineno(870)
						Label10:
							// line 871: standardMsg = 'Missing: %s' % ','.join(safe_repr(m) for m in
							πF.SetLineno(871)
							πTemp001 = πF.MakeArgs(1)
							πTemp012 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/unittest_case.py", πTemp012, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µm *πg.Object = πg.UnboundLocal; _ = µm
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
										if πE = πg.CheckLocal(πF, µmissing, "missing"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µmissing); πE != nil {
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
											µm = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 871: standardMsg = 'Missing: %s' % ','.join(safe_repr(m) for m in
										πF.SetLineno(871)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
											continue
										}
										πTemp005[0] = µm
										if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
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
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πTemp004, πE = πg.GetAttr(πF, πg.NewStr(",").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Missing: %s").ToObject(), πTemp007); πE != nil {
								continue
							}
							µstandardMsg = πTemp002
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µmismatched, "mismatched"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µmismatched); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label12
							}
							goto Label13
							// line 873: if mismatched:
							πF.SetLineno(873)
						Label12:
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µstandardMsg); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label14
							}
							goto Label15
							// line 874: if standardMsg:
							πF.SetLineno(874)
						Label14:
							// line 875: standardMsg += '; '
							πF.SetLineno(875)
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µstandardMsg, πg.NewStr("; ").ToObject()); πE != nil {
								continue
							}
							µstandardMsg = πTemp002
							goto Label15
						Label15:
							// line 876: standardMsg += 'Mismatched values: %s' % ','.join(mismatched)
							πF.SetLineno(876)
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmismatched, "mismatched"); πE != nil {
								continue
							}
							πTemp001[0] = µmismatched
							if πTemp004, πE = πg.GetAttr(πF, πg.NewStr(",").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Mismatched values: %s").ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IAdd(πF, µstandardMsg, πTemp002); πE != nil {
								continue
							}
							µstandardMsg = πTemp004
							goto Label13
						Label13:
							// line 878: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(878)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertDictContainsSubset.ToObject(), πTemp047); πE != nil {
						continue
					}
					// line 855: """Checks whether actual is a superset of expected."""
					πF.SetLineno(855)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp048}, πg.NewStr("Checks whether actual is a superset of expected.").ToObject()); πE != nil {
						continue
					}
					if πTemp049, πE = πg.ResolveClass(πF, πClass, nil, ßassertDictContainsSubset); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp049, ß__doc__, πTemp048); πE != nil {
						continue
					}
					// line 880: def assertItemsEqual(self, expected_seq, actual_seq, msg=None):
					πF.SetLineno(880)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "expected_seq", Def: nil}
					πTemp002[2] = πg.Param{Name: "actual_seq", Def: nil}
					if πTemp049, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp049}
					πTemp048 = πg.NewFunction(πg.NewCode("assertItemsEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpected_seq *πg.Object = πArgs[1]; _ = µexpected_seq
						var µactual_seq *πg.Object = πArgs[2]; _ = µactual_seq
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µfirst_seq *πg.Object = πg.UnboundLocal; _ = µfirst_seq
						var µsecond_seq *πg.Object = πg.UnboundLocal; _ = µsecond_seq
						var µ_msg *πg.Object = πg.UnboundLocal; _ = µ_msg
						var µfirst *πg.Object = πg.UnboundLocal; _ = µfirst
						var µsecond *πg.Object = πg.UnboundLocal; _ = µsecond
						var µdifferences *πg.Object = πg.UnboundLocal; _ = µdifferences
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
						var µlines *πg.Object = πg.UnboundLocal; _ = µlines
						var µdiffMsg *πg.Object = πg.UnboundLocal; _ = µdiffMsg
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.BaseException
						_ = πTemp010
						var πTemp011 *πg.Traceback
						_ = πTemp011
						var πTemp012 *πg.Type
						_ = πTemp012
						var πTemp013 []πg.Param
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8: goto Label8
							case 1: goto Label1
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 881: """An unordered sequence specific comparison. It asserts that
							πF.SetLineno(881)
							// line 893: first_seq, second_seq = list(expected_seq), list(actual_seq)
							πF.SetLineno(893)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpected_seq, "expected_seq"); πE != nil {
								continue
							}
							πTemp002[0] = µexpected_seq
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µactual_seq, "actual_seq"); πE != nil {
								continue
							}
							πTemp002[0] = µactual_seq
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp001); πE != nil {
								continue
							}
							µfirst_seq = πTemp003
							µsecond_seq = πTemp004
							// line 894: with warnings.catch_warnings():
							πF.SetLineno(894)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßcatch_warnings, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp001.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp004.Call(πF, πg.Args{πTemp001}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßpy3kwarning, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label2
							}
							goto Label3
							// line 895: if sys.py3kwarning:
							πF.SetLineno(895)
						Label2:
							πTemp002 = make([]*πg.Object, 3)
							πTemp002[0] = πg.NewStr("(code|dict|type) inequality comparisons").ToObject()
							πTemp002[1] = πg.NewStr("builtin_function_or_method order comparisons").ToObject()
							πTemp002[2] = πg.NewStr("comparing unequal types").ToObject()
							πTemp006 = πg.NewList(πTemp002...).ToObject()
							if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
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
							if πTemp006, πE = πg.Next(πF, πTemp005); πE != nil {
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
								µ_msg = πTemp006
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 900: warnings.filterwarnings("ignore", _msg, DeprecationWarning)
							πF.SetLineno(900)
							πTemp002 = πF.MakeArgs(3)
							πTemp002[0] = ßignore.ToObject()
							if πE = πg.CheckLocal(πF, µ_msg, "_msg"); πE != nil {
								continue
							}
							πTemp002[1] = µ_msg
							if πTemp006, πE = πg.ResolveGlobal(πF, ßDeprecationWarning); πE != nil {
								continue
							}
							πTemp002[2] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßwarnings); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp006, ßfilterwarnings, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp009.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							goto Label3
						Label3:
							// line 901: try:
							πF.SetLineno(901)
							πF.PushCheckpoint(8)
							// line 902: first = collections.Counter(first_seq)
							πF.SetLineno(902)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst_seq, "first_seq"); πE != nil {
								continue
							}
							πTemp002[0] = µfirst_seq
							if πTemp005, πE = πg.ResolveGlobal(πF, ßcollections); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßCounter, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µfirst = πTemp005
							// line 903: second = collections.Counter(second_seq)
							πF.SetLineno(903)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond_seq, "second_seq"); πE != nil {
								continue
							}
							πTemp002[0] = µsecond_seq
							if πTemp005, πE = πg.ResolveGlobal(πF, ßcollections); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßCounter, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µsecond = πTemp005
							πF.PopCheckpoint()
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, µfirst, µsecond); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label9
							}
							goto Label10
							// line 908: if first == second:
							πF.SetLineno(908)
						Label9:
							// line 909: return
							πF.SetLineno(909)
							πR = πg.None
							continue
							goto Label10
						Label10:
							// line 910: differences = _count_diff_hashable(first_seq, second_seq)
							πF.SetLineno(910)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfirst_seq, "first_seq"); πE != nil {
								continue
							}
							πTemp002[0] = µfirst_seq
							if πE = πg.CheckLocal(πF, µsecond_seq, "second_seq"); πE != nil {
								continue
							}
							πTemp002[1] = µsecond_seq
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_count_diff_hashable); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdifferences = πTemp006
							goto Label7
						Label8:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp005); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label11
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 904: except TypeError:
							πF.SetLineno(904)
						Label11:
							// line 906: differences = _count_diff_all_purpose(first_seq, second_seq)
							πF.SetLineno(906)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfirst_seq, "first_seq"); πE != nil {
								continue
							}
							πTemp002[0] = µfirst_seq
							if πE = πg.CheckLocal(πF, µsecond_seq, "second_seq"); πE != nil {
								continue
							}
							πTemp002[1] = µsecond_seq
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_count_diff_all_purpose); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdifferences = πTemp006
							πF.RestoreExc(nil, nil)
							goto Label7
						Label7:
							πF.PopCheckpoint()
						Label1:
							πTemp010, πTemp011 = nil, nil
							if πE != nil {
								πTemp010, πTemp011 = πF.ExcInfo()
							}
							if πTemp010 != nil {
								πTemp012 = πTemp010.Type()
								if πTemp005, πE = πTemp003.Call(πF, πg.Args{πTemp001, πTemp012.ToObject(), πTemp010.ToObject(), πTemp011.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp005, πE = πTemp003.Call(πF, πg.Args{πTemp001, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp010 != nil && πTemp007 != true {
								πE = πF.Raise(nil, nil, nil)
								continue
							}
							if πR != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdifferences, "differences"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µdifferences); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label12
							}
							goto Label13
							// line 912: if differences:
							πF.SetLineno(912)
						Label12:
							// line 913: standardMsg = 'Element counts were not equal:\n'
							πF.SetLineno(913)
							µstandardMsg = πg.NewStr("Element counts were not equal:\n").ToObject()
							// line 914: lines = ['First has %d, Second has %d:  %r' % diff for diff in differences]
							πF.SetLineno(914)
							πTemp013 = make([]πg.Param, 0)
							πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/unittest_case.py", πTemp013, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µdiff *πg.Object = πg.UnboundLocal; _ = µdiff
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
										if πE = πg.CheckLocal(πF, µdifferences, "differences"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µdifferences); πE != nil {
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
											µdiff = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 914: lines = ['First has %d, Second has %d:  %r' % diff for diff in differences]
										πF.SetLineno(914)
										if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.Mod(πF, πg.NewStr("First has %d, Second has %d:  %r").ToObject(), µdiff); πE != nil {
											continue
										}
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
							if πTemp009, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							µlines = πTemp005
							// line 915: diffMsg = '\n'.join(lines)
							πF.SetLineno(915)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlines, "lines"); πE != nil {
								continue
							}
							πTemp002[0] = µlines
							if πTemp005, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µdiffMsg = πTemp009
							// line 916: standardMsg = self._truncateMessage(standardMsg, diffMsg)
							πF.SetLineno(916)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp002[0] = µstandardMsg
							if πE = πg.CheckLocal(πF, µdiffMsg, "diffMsg"); πE != nil {
								continue
							}
							πTemp002[1] = µdiffMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_truncateMessage, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µstandardMsg = πTemp009
							// line 917: msg = self._formatMessage(msg, standardMsg)
							πF.SetLineno(917)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp002[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µmsg = πTemp009
							// line 918: self.fail(msg)
							πF.SetLineno(918)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label13
						Label13:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßassertItemsEqual.ToObject(), πTemp048); πE != nil {
						continue
					}
					// line 881: """An unordered sequence specific comparison. It asserts that
					πF.SetLineno(881)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp049}, πg.NewStr("An unordered sequence specific comparison. It asserts that\n        actual_seq and expected_seq have the same element counts.\n        Equivalent to::\n\n            self.assertEqual(Counter(iter(actual_seq)),\n                             Counter(iter(expected_seq)))\n\n        Asserts that each element has the same count in both sequences.\n        Example:\n            - [0, 1, 1] and [1, 0, 1] compare equal.\n            - [0, 0, 1] and [0, 1] compare unequal.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp050, πE = πg.ResolveClass(πF, πClass, nil, ßassertItemsEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp050, ß__doc__, πTemp049); πE != nil {
						continue
					}
					// line 920: def assertMultiLineEqual(self, first, second, msg=None):
					πF.SetLineno(920)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "first", Def: nil}
					πTemp002[2] = πg.Param{Name: "second", Def: nil}
					if πTemp050, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp050}
					πTemp049 = πg.NewFunction(πg.NewCode("assertMultiLineEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfirst *πg.Object = πArgs[1]; _ = µfirst
						var µsecond *πg.Object = πArgs[2]; _ = µsecond
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µfirstlines *πg.Object = πg.UnboundLocal; _ = µfirstlines
						var µsecondlines *πg.Object = πg.UnboundLocal; _ = µsecondlines
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
						var µdiff *πg.Object = πg.UnboundLocal; _ = µdiff
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
						var πTemp007 *πg.Object
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
							// line 921: """Assert that two multi-line strings are equal."""
							πF.SetLineno(921)
							// line 922: self.assertIsInstance(first, basestring,
							πF.SetLineno(922)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp001[0] = µfirst
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewStr("First argument is not a string").ToObject()
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
							// line 924: self.assertIsInstance(second, basestring,
							πF.SetLineno(924)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp001[0] = µsecond
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							πTemp001[2] = πg.NewStr("Second argument is not a string").ToObject()
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
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µfirst, µsecond); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 927: if first != second:
							πF.SetLineno(927)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp001[0] = µfirst
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_diffThreshold, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp001[0] = µsecond
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_diffThreshold, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, πTemp006, πTemp005); πE != nil {
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
							// line 929: if (len(first) > self._diffThreshold or
							πF.SetLineno(929)
						Label4:
							// line 931: self._baseAssertEqual(first, second, msg)
							πF.SetLineno(931)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp001[0] = µfirst
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp001[1] = µsecond
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[2] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_baseAssertEqual, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							// line 932: firstlines = first.splitlines(True)
							πF.SetLineno(932)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µfirst, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfirstlines = πTemp003
							// line 933: secondlines = second.splitlines(True)
							πF.SetLineno(933)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsecond, ßsplitlines, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsecondlines = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfirstlines, "firstlines"); πE != nil {
								continue
							}
							πTemp001[0] = µfirstlines
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label6
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\r\n").ToObject()
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µfirst, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, µfirst); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label6:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 934: if len(firstlines) == 1 and first.strip('\r\n') == first:
							πF.SetLineno(934)
						Label7:
							// line 935: firstlines = [first + '\n']
							πF.SetLineno(935)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µfirst, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µfirstlines = πTemp002
							// line 936: secondlines = [second + '\n']
							πF.SetLineno(936)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µsecond, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µsecondlines = πTemp002
							goto Label8
						Label8:
							// line 937: standardMsg = '%s != %s' % (safe_repr(first, True),
							πF.SetLineno(937)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfirst, "first"); πE != nil {
								continue
							}
							πTemp001[0] = µfirst
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsecond, "second"); πE != nil {
								continue
							}
							πTemp001[0] = µsecond
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[1] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp003 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s != %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp002
							// line 939: diff = '\n' + ''.join(difflib.ndiff(firstlines, secondlines))
							πF.SetLineno(939)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfirstlines, "firstlines"); πE != nil {
								continue
							}
							πTemp008[0] = µfirstlines
							if πE = πg.CheckLocal(πF, µsecondlines, "secondlines"); πE != nil {
								continue
							}
							πTemp008[1] = µsecondlines
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdifflib); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßndiff, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp005); πE != nil {
								continue
							}
							µdiff = πTemp002
							// line 940: standardMsg = self._truncateMessage(standardMsg, diff)
							πF.SetLineno(940)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp001[0] = µstandardMsg
							if πE = πg.CheckLocal(πF, µdiff, "diff"); πE != nil {
								continue
							}
							πTemp001[1] = µdiff
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_truncateMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstandardMsg = πTemp003
							// line 941: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(941)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001[0] = πTemp003
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
					if πE = πClass.SetItem(πF, ßassertMultiLineEqual.ToObject(), πTemp049); πE != nil {
						continue
					}
					// line 921: """Assert that two multi-line strings are equal."""
					πF.SetLineno(921)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp050}, πg.NewStr("Assert that two multi-line strings are equal.").ToObject()); πE != nil {
						continue
					}
					if πTemp051, πE = πg.ResolveClass(πF, πClass, nil, ßassertMultiLineEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp051, ß__doc__, πTemp050); πE != nil {
						continue
					}
					// line 943: def assertLess(self, a, b, msg=None):
					πF.SetLineno(943)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "b", Def: nil}
					if πTemp051, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp051}
					πTemp050 = πg.NewFunction(πg.NewCode("assertLess", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µb *πg.Object = πArgs[2]; _ = µb
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 944: """Just like self.assertTrue(a < b), but with a nicer default message."""
							πF.SetLineno(944)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µa, µb); πE != nil {
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
							// line 945: if not a < b:
							πF.SetLineno(945)
						Label1:
							// line 946: standardMsg = '%s not less than %s' % (safe_repr(a), safe_repr(b))
							πF.SetLineno(946)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004[0] = µa
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp004[0] = µb
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s not less than %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 947: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(947)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertLess.ToObject(), πTemp050); πE != nil {
						continue
					}
					// line 944: """Just like self.assertTrue(a < b), but with a nicer default message."""
					πF.SetLineno(944)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp051}, πg.NewStr("Just like self.assertTrue(a < b), but with a nicer default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp052, πE = πg.ResolveClass(πF, πClass, nil, ßassertLess); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp052, ß__doc__, πTemp051); πE != nil {
						continue
					}
					// line 949: def assertLessEqual(self, a, b, msg=None):
					πF.SetLineno(949)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "b", Def: nil}
					if πTemp052, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp052}
					πTemp051 = πg.NewFunction(πg.NewCode("assertLessEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µb *πg.Object = πArgs[2]; _ = µb
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 950: """Just like self.assertTrue(a <= b), but with a nicer default message."""
							πF.SetLineno(950)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, µa, µb); πE != nil {
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
							// line 951: if not a <= b:
							πF.SetLineno(951)
						Label1:
							// line 952: standardMsg = '%s not less than or equal to %s' % (safe_repr(a), safe_repr(b))
							πF.SetLineno(952)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004[0] = µa
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp004[0] = µb
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s not less than or equal to %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 953: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(953)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertLessEqual.ToObject(), πTemp051); πE != nil {
						continue
					}
					// line 950: """Just like self.assertTrue(a <= b), but with a nicer default message."""
					πF.SetLineno(950)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp052}, πg.NewStr("Just like self.assertTrue(a <= b), but with a nicer default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp053, πE = πg.ResolveClass(πF, πClass, nil, ßassertLessEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp053, ß__doc__, πTemp052); πE != nil {
						continue
					}
					// line 955: def assertGreater(self, a, b, msg=None):
					πF.SetLineno(955)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "b", Def: nil}
					if πTemp053, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp053}
					πTemp052 = πg.NewFunction(πg.NewCode("assertGreater", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µb *πg.Object = πArgs[2]; _ = µb
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 956: """Just like self.assertTrue(a > b), but with a nicer default message."""
							πF.SetLineno(956)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, µa, µb); πE != nil {
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
							// line 957: if not a > b:
							πF.SetLineno(957)
						Label1:
							// line 958: standardMsg = '%s not greater than %s' % (safe_repr(a), safe_repr(b))
							πF.SetLineno(958)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004[0] = µa
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp004[0] = µb
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s not greater than %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 959: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(959)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertGreater.ToObject(), πTemp052); πE != nil {
						continue
					}
					// line 956: """Just like self.assertTrue(a > b), but with a nicer default message."""
					πF.SetLineno(956)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp053}, πg.NewStr("Just like self.assertTrue(a > b), but with a nicer default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp054, πE = πg.ResolveClass(πF, πClass, nil, ßassertGreater); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp054, ß__doc__, πTemp053); πE != nil {
						continue
					}
					// line 961: def assertGreaterEqual(self, a, b, msg=None):
					πF.SetLineno(961)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "a", Def: nil}
					πTemp002[2] = πg.Param{Name: "b", Def: nil}
					if πTemp054, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp054}
					πTemp053 = πg.NewFunction(πg.NewCode("assertGreaterEqual", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µa *πg.Object = πArgs[1]; _ = µa
						var µb *πg.Object = πArgs[2]; _ = µb
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 962: """Just like self.assertTrue(a >= b), but with a nicer default message."""
							πF.SetLineno(962)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GE(πF, µa, µb); πE != nil {
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
							// line 963: if not a >= b:
							πF.SetLineno(963)
						Label1:
							// line 964: standardMsg = '%s not greater than or equal to %s' % (safe_repr(a), safe_repr(b))
							πF.SetLineno(964)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp004[0] = µa
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							πTemp004[0] = µb
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s not greater than or equal to %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 965: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(965)
							πTemp004 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp008[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp008[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertGreaterEqual.ToObject(), πTemp053); πE != nil {
						continue
					}
					// line 962: """Just like self.assertTrue(a >= b), but with a nicer default message."""
					πF.SetLineno(962)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp054}, πg.NewStr("Just like self.assertTrue(a >= b), but with a nicer default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp055, πE = πg.ResolveClass(πF, πClass, nil, ßassertGreaterEqual); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp055, ß__doc__, πTemp054); πE != nil {
						continue
					}
					// line 967: def assertIsNone(self, obj, msg=None):
					πF.SetLineno(967)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					if πTemp055, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "msg", Def: πTemp055}
					πTemp054 = πg.NewFunction(πg.NewCode("assertIsNone", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πArgs[1]; _ = µobj
						var µmsg *πg.Object = πArgs[2]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 968: """Same as self.assertTrue(obj is None), with a nicer default message."""
							πF.SetLineno(968)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µobj != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 969: if obj is not None:
							πF.SetLineno(969)
						Label1:
							// line 970: standardMsg = '%s is not None' % (safe_repr(obj),)
							πF.SetLineno(970)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πg.NewTuple1(πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s is not None").ToObject(), πTemp002); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 971: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(971)
							πTemp004 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp007[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp007[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertIsNone.ToObject(), πTemp054); πE != nil {
						continue
					}
					// line 968: """Same as self.assertTrue(obj is None), with a nicer default message."""
					πF.SetLineno(968)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp055}, πg.NewStr("Same as self.assertTrue(obj is None), with a nicer default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp056, πE = πg.ResolveClass(πF, πClass, nil, ßassertIsNone); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp056, ß__doc__, πTemp055); πE != nil {
						continue
					}
					// line 973: def assertIsNotNone(self, obj, msg=None):
					πF.SetLineno(973)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					if πTemp056, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "msg", Def: πTemp056}
					πTemp055 = πg.NewFunction(πg.NewCode("assertIsNotNone", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πArgs[1]; _ = µobj
						var µmsg *πg.Object = πArgs[2]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
							// line 974: """Included for symmetry with assertIsNone."""
							πF.SetLineno(974)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µobj == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 975: if obj is None:
							πF.SetLineno(975)
						Label1:
							// line 976: standardMsg = 'unexpectedly None'
							πF.SetLineno(976)
							µstandardMsg = πg.NewStr("unexpectedly None").ToObject()
							// line 977: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(977)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp005[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp005[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßassertIsNotNone.ToObject(), πTemp055); πE != nil {
						continue
					}
					// line 974: """Included for symmetry with assertIsNone."""
					πF.SetLineno(974)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp056}, πg.NewStr("Included for symmetry with assertIsNone.").ToObject()); πE != nil {
						continue
					}
					if πTemp057, πE = πg.ResolveClass(πF, πClass, nil, ßassertIsNotNone); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp057, ß__doc__, πTemp056); πE != nil {
						continue
					}
					// line 979: def assertIsInstance(self, obj, cls, msg=None):
					πF.SetLineno(979)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp002[2] = πg.Param{Name: "cls", Def: nil}
					if πTemp057, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp057}
					πTemp056 = πg.NewFunction(πg.NewCode("assertIsInstance", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πArgs[1]; _ = µobj
						var µcls *πg.Object = πArgs[2]; _ = µcls
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 980: """Same as self.assertTrue(isinstance(obj, cls)), with a nicer
							πF.SetLineno(980)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp002[1] = µcls
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
							// line 982: if not isinstance(obj, cls):
							πF.SetLineno(982)
						Label1:
							// line 983: standardMsg = '%s is not an instance of %r' % (safe_repr(obj), cls)
							πF.SetLineno(983)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp006, µcls).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s is not an instance of %r").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp001
							// line 984: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(984)
							πTemp002 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp007[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp007[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfail, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßassertIsInstance.ToObject(), πTemp056); πE != nil {
						continue
					}
					// line 980: """Same as self.assertTrue(isinstance(obj, cls)), with a nicer
					πF.SetLineno(980)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp057}, πg.NewStr("Same as self.assertTrue(isinstance(obj, cls)), with a nicer\n        default message.").ToObject()); πE != nil {
						continue
					}
					if πTemp058, πE = πg.ResolveClass(πF, πClass, nil, ßassertIsInstance); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp058, ß__doc__, πTemp057); πE != nil {
						continue
					}
					// line 986: def assertNotIsInstance(self, obj, cls, msg=None):
					πF.SetLineno(986)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp002[2] = πg.Param{Name: "cls", Def: nil}
					if πTemp058, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp058}
					πTemp057 = πg.NewFunction(πg.NewCode("assertNotIsInstance", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µobj *πg.Object = πArgs[1]; _ = µobj
						var µcls *πg.Object = πArgs[2]; _ = µcls
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µstandardMsg *πg.Object = πg.UnboundLocal; _ = µstandardMsg
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 987: """Included for symmetry with assertIsInstance."""
							πF.SetLineno(987)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[1] = µcls
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
							// line 988: if isinstance(obj, cls):
							πF.SetLineno(988)
						Label1:
							// line 989: standardMsg = '%s is an instance of %r' % (safe_repr(obj), cls)
							πF.SetLineno(989)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsafe_repr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp006, µcls).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s is an instance of %r").ToObject(), πTemp003); πE != nil {
								continue
							}
							µstandardMsg = πTemp002
							// line 990: self.fail(self._formatMessage(msg, standardMsg))
							πF.SetLineno(990)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp007[0] = µmsg
							if πE = πg.CheckLocal(πF, µstandardMsg, "standardMsg"); πE != nil {
								continue
							}
							πTemp007[1] = µstandardMsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_formatMessage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
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
					if πE = πClass.SetItem(πF, ßassertNotIsInstance.ToObject(), πTemp057); πE != nil {
						continue
					}
					// line 987: """Included for symmetry with assertIsInstance."""
					πF.SetLineno(987)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp058}, πg.NewStr("Included for symmetry with assertIsInstance.").ToObject()); πE != nil {
						continue
					}
					if πTemp059, πE = πg.ResolveClass(πF, πClass, nil, ßassertNotIsInstance); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp059, ß__doc__, πTemp058); πE != nil {
						continue
					}
					// line 992: def assertRaisesRegexp(self, expected_exception, expected_regexp,
					πF.SetLineno(992)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "expected_exception", Def: nil}
					πTemp002[2] = πg.Param{Name: "expected_regexp", Def: nil}
					if πTemp059, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "callable_obj", Def: πTemp059}
					πTemp058 = πg.NewFunction(πg.NewCode("assertRaisesRegexp", "/usr/lib/python2.7/unittest_case.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µexpected_exception *πg.Object = πArgs[1]; _ = µexpected_exception
						var µexpected_regexp *πg.Object = πArgs[2]; _ = µexpected_regexp
						var µcallable_obj *πg.Object = πArgs[3]; _ = µcallable_obj
						var µargs *πg.Object = πArgs[4]; _ = µargs
						var µkwargs *πg.Object = πArgs[5]; _ = µkwargs
						var µcontext *πg.Object = πg.UnboundLocal; _ = µcontext
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 *πg.Type
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 994: """Asserts that the message in a raised exception matches a regexp.
							πF.SetLineno(994)
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µexpected_regexp != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1004: if expected_regexp is not None:
							πF.SetLineno(1004)
						Label1:
							// line 1005: expected_regexp = re.compile(expected_regexp)
							πF.SetLineno(1005)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							πTemp004[0] = µexpected_regexp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcompile, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µexpected_regexp = πTemp001
							goto Label2
						Label2:
							// line 1006: context = _AssertRaisesContext(expected_exception, self, expected_regexp)
							πF.SetLineno(1006)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µexpected_exception, "expected_exception"); πE != nil {
								continue
							}
							πTemp004[0] = µexpected_exception
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							πTemp004[2] = µexpected_regexp
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_AssertRaisesContext); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µcontext = πTemp002
							if πE = πg.CheckLocal(πF, µcallable_obj, "callable_obj"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µcallable_obj == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1007: if callable_obj is None:
							πF.SetLineno(1007)
						Label3:
							// line 1008: return context
							πF.SetLineno(1008)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πR = µcontext
							continue
							goto Label4
						Label4:
							// line 1009: with context:
							πF.SetLineno(1009)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcontext.Type().ToObject(), ß__exit__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcontext.Type().ToObject(), ß__enter__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp002.Call(πF, πg.Args{µcontext}, nil); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							// line 1010: callable_obj(*args, **kwargs)
							πF.SetLineno(1010)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcallable_obj, "callable_obj"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Invoke(πF, µcallable_obj, nil, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.PopCheckpoint()
						Label5:
							πTemp006, πTemp007 = nil, nil
							if πE != nil {
								πTemp006, πTemp007 = πF.ExcInfo()
							}
							if πTemp006 != nil {
								πTemp008 = πTemp006.Type()
								if πTemp005, πE = πTemp001.Call(πF, πg.Args{µcontext, πTemp008.ToObject(), πTemp006.ToObject(), πTemp007.ToObject()}, nil); πE != nil {
									continue
								}
							} else {
								if πTemp005, πE = πTemp001.Call(πF, πg.Args{µcontext, πg.None, πg.None, πg.None}, nil); πE != nil {
									continue
								}
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp006 != nil && πTemp003 != true {
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
					if πE = πClass.SetItem(πF, ßassertRaisesRegexp.ToObject(), πTemp058); πE != nil {
						continue
					}
					// line 994: """Asserts that the message in a raised exception matches a regexp.
					πF.SetLineno(994)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp059}, πg.NewStr("Asserts that the message in a raised exception matches a regexp.\n\n        Args:\n            expected_exception: Exception class expected to be raised.\n            expected_regexp: Regexp (re pattern object or string) expected\n                    to be found in error message.\n            callable_obj: Function to be called.\n            args: Extra args.\n            kwargs: Extra kwargs.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp060, πE = πg.ResolveClass(πF, πClass, nil, ßassertRaisesRegexp); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp060, ß__doc__, πTemp059); πE != nil {
						continue
					}
					// line 1012: def assertRegexpMatches(self, text, expected_regexp, msg=None):
					πF.SetLineno(1012)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "text", Def: nil}
					πTemp002[2] = πg.Param{Name: "expected_regexp", Def: nil}
					if πTemp060, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp060}
					πTemp059 = πg.NewFunction(πg.NewCode("assertRegexpMatches", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtext *πg.Object = πArgs[1]; _ = µtext
						var µexpected_regexp *πg.Object = πArgs[2]; _ = µexpected_regexp
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1013: """Fail the test unless the text matches the regular expression."""
							πF.SetLineno(1013)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							πTemp001[0] = µexpected_regexp
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
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
							// line 1014: if isinstance(expected_regexp, basestring):
							πF.SetLineno(1014)
						Label1:
							// line 1015: expected_regexp = re.compile(expected_regexp)
							πF.SetLineno(1015)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							πTemp001[0] = µexpected_regexp
							if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µexpected_regexp = πTemp002
							goto Label2
						Label2:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µexpected_regexp, ßsearch, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1016: if not expected_regexp.search(text):
							πF.SetLineno(1016)
						Label3:
							// line 1017: msg = msg or "Regexp didn't match"
							πF.SetLineno(1017)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002 = µmsg
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πTemp002 = πg.NewStr("Regexp didn't match").ToObject()
						Label5:
							µmsg = πTemp002
							// line 1018: msg = '%s: %r not found in %r' % (msg, expected_regexp.pattern, text)
							πF.SetLineno(1018)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µexpected_regexp, "expected_regexp"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µexpected_regexp, ßpattern, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple3(µmsg, πTemp005, µtext).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s: %r not found in %r").ToObject(), πTemp003); πE != nil {
								continue
							}
							µmsg = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1019: raise self.failureException(msg)
							πF.SetLineno(1019)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßassertRegexpMatches.ToObject(), πTemp059); πE != nil {
						continue
					}
					// line 1013: """Fail the test unless the text matches the regular expression."""
					πF.SetLineno(1013)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp060}, πg.NewStr("Fail the test unless the text matches the regular expression.").ToObject()); πE != nil {
						continue
					}
					if πTemp061, πE = πg.ResolveClass(πF, πClass, nil, ßassertRegexpMatches); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp061, ß__doc__, πTemp060); πE != nil {
						continue
					}
					// line 1021: def assertNotRegexpMatches(self, text, unexpected_regexp, msg=None):
					πF.SetLineno(1021)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "text", Def: nil}
					πTemp002[2] = πg.Param{Name: "unexpected_regexp", Def: nil}
					if πTemp061, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "msg", Def: πTemp061}
					πTemp060 = πg.NewFunction(πg.NewCode("assertNotRegexpMatches", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtext *πg.Object = πArgs[1]; _ = µtext
						var µunexpected_regexp *πg.Object = πArgs[2]; _ = µunexpected_regexp
						var µmsg *πg.Object = πArgs[3]; _ = µmsg
						var µmatch *πg.Object = πg.UnboundLocal; _ = µmatch
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
							// line 1022: """Fail the test if the text matches the regular expression."""
							πF.SetLineno(1022)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µunexpected_regexp, "unexpected_regexp"); πE != nil {
								continue
							}
							πTemp001[0] = µunexpected_regexp
							if πTemp002, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
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
							// line 1023: if isinstance(unexpected_regexp, basestring):
							πF.SetLineno(1023)
						Label1:
							// line 1024: unexpected_regexp = re.compile(unexpected_regexp)
							πF.SetLineno(1024)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µunexpected_regexp, "unexpected_regexp"); πE != nil {
								continue
							}
							πTemp001[0] = µunexpected_regexp
							if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µunexpected_regexp = πTemp002
							goto Label2
						Label2:
							// line 1025: match = unexpected_regexp.search(text)
							πF.SetLineno(1025)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µunexpected_regexp, "unexpected_regexp"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µunexpected_regexp, ßsearch, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmatch = πTemp003
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmatch); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1026: if match:
							πF.SetLineno(1026)
						Label3:
							// line 1027: msg = msg or "Regexp matched"
							πF.SetLineno(1027)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002 = µmsg
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πTemp002 = πg.NewStr("Regexp matched").ToObject()
						Label5:
							µmsg = πTemp002
							// line 1028: msg = '%s: %r matches %r in %r' % (msg,
							πF.SetLineno(1028)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µmatch, ßstart, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µmatch, ßend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πTemp007, πTemp008, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µtext, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µunexpected_regexp, "unexpected_regexp"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µunexpected_regexp, ßpattern, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple4(µmsg, πTemp006, πTemp005, µtext).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s: %r matches %r in %r").ToObject(), πTemp003); πE != nil {
								continue
							}
							µmsg = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp001[0] = µmsg
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfailureException, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1032: raise self.failureException(msg)
							πF.SetLineno(1032)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
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
					if πE = πClass.SetItem(πF, ßassertNotRegexpMatches.ToObject(), πTemp060); πE != nil {
						continue
					}
					// line 1022: """Fail the test if the text matches the regular expression."""
					πF.SetLineno(1022)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp061}, πg.NewStr("Fail the test if the text matches the regular expression.").ToObject()); πE != nil {
						continue
					}
					if πTemp062, πE = πg.ResolveClass(πF, πClass, nil, ßassertNotRegexpMatches); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp062, ß__doc__, πTemp061); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("TestCase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTestCase.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 1035: class FunctionTestCase(TestCase):
			πF.SetLineno(1035)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßTestCase); πE != nil {
				continue
			}
			πTemp002[0] = πTemp011
			πTemp009 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp009.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("FunctionTestCase", "/usr/lib/python2.7/unittest_case.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1036: """A test case that wraps a test function.
					πF.SetLineno(1036)
					// line 1036: """A test case that wraps a test function.
					πF.SetLineno(1036)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A test case that wraps a test function.\n\n    This is useful for slipping pre-existing test functions into the\n    unittest framework. Optionally, set-up and tidy-up functions can be\n    supplied. As with TestCase, the tidy-up ('tearDown') function will\n    always be called if the set-up ('setUp') function ran successfully.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1044: def __init__(self, testFunc, setUp=None, tearDown=None, description=None):
					πF.SetLineno(1044)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "testFunc", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "setUp", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "tearDown", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "description", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtestFunc *πg.Object = πArgs[1]; _ = µtestFunc
						var µsetUp *πg.Object = πArgs[2]; _ = µsetUp
						var µtearDown *πg.Object = πArgs[3]; _ = µtearDown
						var µdescription *πg.Object = πArgs[4]; _ = µdescription
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
							// line 1045: super(FunctionTestCase, self).__init__()
							πF.SetLineno(1045)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFunctionTestCase); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[1] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsuper); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1046: self._setUpFunc = setUp
							πF.SetLineno(1046)
							if πE = πg.CheckLocal(πF, µsetUp, "setUp"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µsetUp); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_setUpFunc, πTemp002); πE != nil {
								continue
							}
							// line 1047: self._tearDownFunc = tearDown
							πF.SetLineno(1047)
							if πE = πg.CheckLocal(πF, µtearDown, "tearDown"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µtearDown); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_tearDownFunc, πTemp002); πE != nil {
								continue
							}
							// line 1048: self._testFunc = testFunc
							πF.SetLineno(1048)
							if πE = πg.CheckLocal(πF, µtestFunc, "testFunc"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µtestFunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_testFunc, πTemp002); πE != nil {
								continue
							}
							// line 1049: self._description = description
							πF.SetLineno(1049)
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdescription); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_description, πTemp002); πE != nil {
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
					// line 1051: def setUp(self):
					πF.SetLineno(1051)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("setUp", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_setUpFunc, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1052: if self._setUpFunc is not None:
							πF.SetLineno(1052)
						Label1:
							// line 1053: self._setUpFunc()
							πF.SetLineno(1053)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_setUpFunc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßsetUp.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1055: def tearDown(self):
					πF.SetLineno(1055)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("tearDown", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_tearDownFunc, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1056: if self._tearDownFunc is not None:
							πF.SetLineno(1056)
						Label1:
							// line 1057: self._tearDownFunc()
							πF.SetLineno(1057)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_tearDownFunc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßtearDown.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1059: def runTest(self):
					πF.SetLineno(1059)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("runTest", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1060: self._testFunc()
							πF.SetLineno(1060)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_testFunc, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßrunTest.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1062: def id(self):
					πF.SetLineno(1062)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("id", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1063: return self._testFunc.__name__
							πF.SetLineno(1063)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_testFunc, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__name__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßid.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1065: def __eq__(self, other):
					πF.SetLineno(1065)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__eq__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
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
							// line 1066: if not isinstance(other, self.__class__):
							πF.SetLineno(1066)
						Label1:
							// line 1067: return NotImplemented
							πF.SetLineno(1067)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplemented); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1069: return self._setUpFunc == other._setUpFunc and \
							πF.SetLineno(1069)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_setUpFunc, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µother, ß_setUpFunc, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_tearDownFunc, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µother, ß_tearDownFunc, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_testFunc, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µother, ß_testFunc, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_description, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µother, ß_description, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp003
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
					if πE = πClass.SetItem(πF, ß__eq__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1074: def __ne__(self, other):
					πF.SetLineno(1074)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__ne__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1075: return not self == other
							πF.SetLineno(1075)
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
					if πE = πClass.SetItem(πF, ß__ne__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1077: def __hash__(self):
					πF.SetLineno(1077)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__hash__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1078: return hash((type(self), self._setUpFunc, self._tearDownFunc,
							πF.SetLineno(1078)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_setUpFunc, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_tearDownFunc, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ß_testFunc, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ß_description, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple5(πTemp005, πTemp004, πTemp006, πTemp007, πTemp008).ToObject()
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__hash__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1081: def __str__(self):
					πF.SetLineno(1081)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1082: return "%s (%s)" % (strclass(self.__class__),
							πF.SetLineno(1082)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstrclass); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_testFunc, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ß__name__, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s (%s)").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1085: def __repr__(self):
					πF.SetLineno(1085)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 1086: return "<%s tec=%s>" % (strclass(self.__class__),
							πF.SetLineno(1086)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßstrclass); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_testFunc, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp005, πTemp004).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s tec=%s>").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1089: def shortDescription(self):
					πF.SetLineno(1089)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("shortDescription", "/usr/lib/python2.7/unittest_case.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πTemp009 *πg.Object
						_ = πTemp009
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_description, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1090: if self._description is not None:
							πF.SetLineno(1090)
						Label1:
							// line 1091: return self._description
							πF.SetLineno(1091)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_description, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 1093: return doc and doc.split("\n")[0].strip() or None
							πF.SetLineno(1093)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdoc); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label4
							}
							πTemp003 = πg.NewInt(0).ToObject()
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewStr("\n").ToObject()
							if πTemp008, πE = πg.ResolveGlobal(πF, ßdoc); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp009.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßstrip, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002 = πTemp006
						Label4:
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πTemp002
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
					if πE = πClass.SetItem(πF, ßshortDescription.ToObject(), πTemp012); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp009.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("FunctionTestCase").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp009.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFunctionTestCase.ToObject(), πTemp011); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("unittest_case", Code)
}

