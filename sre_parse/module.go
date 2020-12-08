package sre_parse
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/sre_constants"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/sre_parse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß01234567 := πg.InternStr("01234567")
		ß0123456789 := πg.InternStr("0123456789")
		ß0123456789abcdefABCDEF := πg.InternStr("0123456789abcdefABCDEF")
		ß9 := πg.InternStr("9")
		ßA := πg.InternStr("A")
		ßANY := πg.InternStr("ANY")
		ßASSERT := πg.InternStr("ASSERT")
		ßASSERT_NOT := πg.InternStr("ASSERT_NOT")
		ßAT := πg.InternStr("AT")
		ßAT_BEGINNING := πg.InternStr("AT_BEGINNING")
		ßAT_BEGINNING_STRING := πg.InternStr("AT_BEGINNING_STRING")
		ßAT_BOUNDARY := πg.InternStr("AT_BOUNDARY")
		ßAT_END := πg.InternStr("AT_END")
		ßAT_END_STRING := πg.InternStr("AT_END_STRING")
		ßAT_NON_BOUNDARY := πg.InternStr("AT_NON_BOUNDARY")
		ßBRANCH := πg.InternStr("BRANCH")
		ßCALL := πg.InternStr("CALL")
		ßCATEGORIES := πg.InternStr("CATEGORIES")
		ßCATEGORY := πg.InternStr("CATEGORY")
		ßCATEGORY_DIGIT := πg.InternStr("CATEGORY_DIGIT")
		ßCATEGORY_NOT_DIGIT := πg.InternStr("CATEGORY_NOT_DIGIT")
		ßCATEGORY_NOT_SPACE := πg.InternStr("CATEGORY_NOT_SPACE")
		ßCATEGORY_NOT_WORD := πg.InternStr("CATEGORY_NOT_WORD")
		ßCATEGORY_SPACE := πg.InternStr("CATEGORY_SPACE")
		ßCATEGORY_WORD := πg.InternStr("CATEGORY_WORD")
		ßDIGITS := πg.InternStr("DIGITS")
		ßESCAPES := πg.InternStr("ESCAPES")
		ßFLAGS := πg.InternStr("FLAGS")
		ßFalse := πg.InternStr("False")
		ßGROUPREF := πg.InternStr("GROUPREF")
		ßGROUPREF_EXISTS := πg.InternStr("GROUPREF_EXISTS")
		ßHEXDIGITS := πg.InternStr("HEXDIGITS")
		ßIN := πg.InternStr("IN")
		ßIndexError := πg.InternStr("IndexError")
		ßKeyError := πg.InternStr("KeyError")
		ßL := πg.InternStr("L")
		ßLITERAL := πg.InternStr("LITERAL")
		ßMARK := πg.InternStr("MARK")
		ßMAXREPEAT := πg.InternStr("MAXREPEAT")
		ßMAX_REPEAT := πg.InternStr("MAX_REPEAT")
		ßMIN_REPEAT := πg.InternStr("MIN_REPEAT")
		ßNEGATE := πg.InternStr("NEGATE")
		ßNOT_LITERAL := πg.InternStr("NOT_LITERAL")
		ßNone := πg.InternStr("None")
		ßOCTDIGITS := πg.InternStr("OCTDIGITS")
		ßOverflowError := πg.InternStr("OverflowError")
		ßP := πg.InternStr("P")
		ßPattern := πg.InternStr("Pattern")
		ßRANGE := πg.InternStr("RANGE")
		ßREPEAT_CHARS := πg.InternStr("REPEAT_CHARS")
		ßSPECIAL_CHARS := πg.InternStr("SPECIAL_CHARS")
		ßSRE_FLAG_DEBUG := πg.InternStr("SRE_FLAG_DEBUG")
		ßSRE_FLAG_DOTALL := πg.InternStr("SRE_FLAG_DOTALL")
		ßSRE_FLAG_IGNORECASE := πg.InternStr("SRE_FLAG_IGNORECASE")
		ßSRE_FLAG_LOCALE := πg.InternStr("SRE_FLAG_LOCALE")
		ßSRE_FLAG_MULTILINE := πg.InternStr("SRE_FLAG_MULTILINE")
		ßSRE_FLAG_TEMPLATE := πg.InternStr("SRE_FLAG_TEMPLATE")
		ßSRE_FLAG_UNICODE := πg.InternStr("SRE_FLAG_UNICODE")
		ßSRE_FLAG_VERBOSE := πg.InternStr("SRE_FLAG_VERBOSE")
		ßSUBPATTERN := πg.InternStr("SUBPATTERN")
		ßSUCCESS := πg.InternStr("SUCCESS")
		ßSubPattern := πg.InternStr("SubPattern")
		ßTokenizer := πg.InternStr("Tokenizer")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ßWHITESPACE := πg.InternStr("WHITESPACE")
		ßZ := πg.InternStr("Z")
		ß_ := πg.InternStr("_")
		ß_ASSERTCHARS := πg.InternStr("_ASSERTCHARS")
		ß_LOOKBEHINDASSERTCHARS := πg.InternStr("_LOOKBEHINDASSERTCHARS")
		ß_PATTERNENDERS := πg.InternStr("_PATTERNENDERS")
		ß_REPEATCODES := πg.InternStr("_REPEATCODES")
		ß__all__ := πg.InternStr("__all__")
		ß__delitem__ := πg.InternStr("__delitem__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__init__ := πg.InternStr("__init__")
		ß__len__ := πg.InternStr("__len__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__next := πg.InternStr("__next")
		ß__repr__ := πg.InternStr("__repr__")
		ß__setitem__ := πg.InternStr("__setitem__")
		ß_class_escape := πg.InternStr("_class_escape")
		ß_escape := πg.InternStr("_escape")
		ß_parse := πg.InternStr("_parse")
		ß_parse_sub := πg.InternStr("_parse_sub")
		ß_parse_sub_cond := πg.InternStr("_parse_sub_cond")
		ßa := πg.InternStr("a")
		ßappend := πg.InternStr("append")
		ßcheckgroup := πg.InternStr("checkgroup")
		ßchr := πg.InternStr("chr")
		ßclosegroup := πg.InternStr("closegroup")
		ßdata := πg.InternStr("data")
		ßdump := πg.InternStr("dump")
		ßelse := πg.InternStr("else")
		ßenumerate := πg.InternStr("enumerate")
		ßerror := πg.InternStr("error")
		ßexpand_template := πg.InternStr("expand_template")
		ßflags := πg.InternStr("flags")
		ßg := πg.InternStr("g")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßgetwidth := πg.InternStr("getwidth")
		ßglobals := πg.InternStr("globals")
		ßgroup := πg.InternStr("group")
		ßgroupdict := πg.InternStr("groupdict")
		ßgroupindex := πg.InternStr("groupindex")
		ßgroups := πg.InternStr("groups")
		ßi := πg.InternStr("i")
		ßindex := πg.InternStr("index")
		ßinsert := πg.InternStr("insert")
		ßint := πg.InternStr("int")
		ßisdigit := πg.InternStr("isdigit")
		ßisident := πg.InternStr("isident")
		ßisinstance := πg.InternStr("isinstance")
		ßisname := πg.InternStr("isname")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßlookbehind := πg.InternStr("lookbehind")
		ßm := πg.InternStr("m")
		ßmatch := πg.InternStr("match")
		ßmax := πg.InternStr("max")
		ßmin := πg.InternStr("min")
		ßname := πg.InternStr("name")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßopen := πg.InternStr("open")
		ßopengroup := πg.InternStr("opengroup")
		ßor := πg.InternStr("or")
		ßord := πg.InternStr("ord")
		ßparse := πg.InternStr("parse")
		ßparse_template := πg.InternStr("parse_template")
		ßpattern := πg.InternStr("pattern")
		ßrange := πg.InternStr("range")
		ßrepr := πg.InternStr("repr")
		ßs := πg.InternStr("s")
		ßseek := πg.InternStr("seek")
		ßset := πg.InternStr("set")
		ßslice := πg.InternStr("slice")
		ßsre_constants := πg.InternStr("sre_constants")
		ßstr := πg.InternStr("str")
		ßstring := πg.InternStr("string")
		ßsys := πg.InternStr("sys")
		ßt := πg.InternStr("t")
		ßtell := πg.InternStr("tell")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßu := πg.InternStr("u")
		ßunichr := πg.InternStr("unichr")
		ßwidth := πg.InternStr("width")
		ßx := πg.InternStr("x")
		ßz := πg.InternStr("z")
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
		var πTemp010 *πg.Dict
		_ = πTemp010
		var πTemp011 []πg.Param
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 []*πg.Object
		_ = πTemp015
		var πTemp016 *πg.Object
		_ = πTemp016
		var πTemp017 *πg.Object
		_ = πTemp017
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 11: """Internal support module for sre"""
			πF.SetLineno(11)
			// line 11: """Internal support module for sre"""
			πF.SetLineno(11)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Internal support module for sre").ToObject()); πE != nil {
				continue
			}
			// line 15: import sys
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: import sre_constants
			πF.SetLineno(18)
			if πTemp002, πE = πg.ImportModule(πF, "sre_constants"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsre_constants.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsre_constants); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__all__, nil); πE != nil {
				continue
			}
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
				if πE = πF.Globals().SetItem(πF, ßname.ToObject(), πTemp003); πE != nil {
					continue
				}
			}
			if πE != nil || !πTemp006 {
				continue
			}
			πF.PushCheckpoint(1)            
			// line 20: globals()[name] = getattr(sre_constants, name)
			πF.SetLineno(20)
			πTemp002 = πF.MakeArgs(2)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßsre_constants); πE != nil {
				continue
			}
			πTemp002[0] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßname); πE != nil {
				continue
			}
			πTemp002[1] = πTemp003
			if πTemp003, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßglobals); πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßname); πE != nil {
				continue
			}
			πTemp007 = πTemp009
			if πE = πg.SetItem(πF, πTemp008, πTemp007, πTemp003); πE != nil {
				continue
			}
			continue
		Label2:
			if πE != nil || πR != nil {
				continue
			}
		Label3:
			// line 22: SPECIAL_CHARS = ".\\[{()*+?^$|"
			πF.SetLineno(22)
			if πE = πF.Globals().SetItem(πF, ßSPECIAL_CHARS.ToObject(), πg.NewStr(".\\[{()*+?^$|").ToObject()); πE != nil {
				continue
			}
			// line 23: REPEAT_CHARS = "*+?{"
			πF.SetLineno(23)
			if πE = πF.Globals().SetItem(πF, ßREPEAT_CHARS.ToObject(), πg.NewStr("*+?{").ToObject()); πE != nil {
				continue
			}
			// line 25: DIGITS = set("0123456789")
			πF.SetLineno(25)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ß0123456789.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßDIGITS.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 27: OCTDIGITS = set("01234567")
			πF.SetLineno(27)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ß01234567.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßOCTDIGITS.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 28: HEXDIGITS = set("0123456789abcdefABCDEF")
			πF.SetLineno(28)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ß0123456789abcdefABCDEF.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßHEXDIGITS.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 30: WHITESPACE = set(" \t\n\r\v\f")
			πF.SetLineno(30)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr(" \t\n\r\x0b\x0c").ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßWHITESPACE.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 32: ESCAPES = {
			πF.SetLineno(32)
			πTemp010 = πg.NewDict()
			if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("\x07").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\a").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("\x08").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\b").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("\x0c").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\f").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("\n").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\n").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("\r").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\r").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("\t").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\t").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("\x0b").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\v").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("\\").ToObject()
			if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp001 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\\\").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πTemp010.ToObject()
			if πE = πF.Globals().SetItem(πF, ßESCAPES.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 43: CATEGORIES = {
			πF.SetLineno(43)
			πTemp010 = πg.NewDict()
			if πTemp003, πE = πg.ResolveGlobal(πF, ßAT); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_BEGINNING_STRING); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\A").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßAT); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_BOUNDARY); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\b").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßAT); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_NON_BOUNDARY); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\B").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
				continue
			}
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßCATEGORY_DIGIT); πE != nil {
				continue
			}
			πTemp004 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
			πTemp002[0] = πTemp004
			πTemp004 = πg.NewList(πTemp002...).ToObject()
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\d").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
				continue
			}
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_DIGIT); πE != nil {
				continue
			}
			πTemp004 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
			πTemp002[0] = πTemp004
			πTemp004 = πg.NewList(πTemp002...).ToObject()
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\D").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
				continue
			}
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßCATEGORY_SPACE); πE != nil {
				continue
			}
			πTemp004 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
			πTemp002[0] = πTemp004
			πTemp004 = πg.NewList(πTemp002...).ToObject()
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\s").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
				continue
			}
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_SPACE); πE != nil {
				continue
			}
			πTemp004 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
			πTemp002[0] = πTemp004
			πTemp004 = πg.NewList(πTemp002...).ToObject()
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\S").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
				continue
			}
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßCATEGORY_WORD); πE != nil {
				continue
			}
			πTemp004 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
			πTemp002[0] = πTemp004
			πTemp004 = πg.NewList(πTemp002...).ToObject()
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\w").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
				continue
			}
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßCATEGORY_NOT_WORD); πE != nil {
				continue
			}
			πTemp004 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
			πTemp002[0] = πTemp004
			πTemp004 = πg.NewList(πTemp002...).ToObject()
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\W").ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßAT); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_END_STRING); πE != nil {
				continue
			}
			πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
			if πE = πTemp010.SetItem(πF, πg.NewStr("\\Z").ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πTemp010.ToObject()
			if πE = πF.Globals().SetItem(πF, ßCATEGORIES.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 56: FLAGS = {
			πF.SetLineno(56)
			πTemp010 = πg.NewDict()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_IGNORECASE); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ßi.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_LOCALE); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ßL.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_MULTILINE); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ßm.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_DOTALL); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ßs.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_VERBOSE); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ßx.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_TEMPLATE); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ßt.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_UNICODE); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ßu.ToObject(), πTemp001); πE != nil {
				continue
			}
			πTemp001 = πTemp010.ToObject()
			if πE = πF.Globals().SetItem(πF, ßFLAGS.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 68: class Pattern(object):
			πF.SetLineno(68)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp010 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Pattern", "/usr/lib/python2.7/sre_parse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 70: def __init__(self):
					πF.SetLineno(70)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 71: self.flags = 0
							πF.SetLineno(71)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßflags, πTemp001); πE != nil {
								continue
							}
							// line 72: self.open = []
							πF.SetLineno(72)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopen, πTemp003); πE != nil {
								continue
							}
							// line 73: self.groups = 1
							πF.SetLineno(73)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßgroups, πTemp001); πE != nil {
								continue
							}
							// line 74: self.groupdict = {}
							πF.SetLineno(74)
							πTemp004 = πg.NewDict()
							πTemp001 = πTemp004.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßgroupdict, πTemp003); πE != nil {
								continue
							}
							// line 75: self.lookbehind = 0
							πF.SetLineno(75)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlookbehind, πTemp001); πE != nil {
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
					// line 77: def opengroup(self, name=None):
					πF.SetLineno(77)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "name", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("opengroup", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µname *πg.Object = πArgs[1]; _ = µname
						var µgid *πg.Object = πg.UnboundLocal; _ = µgid
						var µogid *πg.Object = πg.UnboundLocal; _ = µogid
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
							// line 78: gid = self.groups
							πF.SetLineno(78)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgroups, nil); πE != nil {
								continue
							}
							µgid = πTemp001
							// line 79: self.groups = gid + 1
							πF.SetLineno(79)
							if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µgid, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßgroups, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µname != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 80: if name is not None:
							πF.SetLineno(80)
						Label1:
							// line 81: ogid = self.groupdict.get(name, None)
							πF.SetLineno(81)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[0] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßgroupdict, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µogid = πTemp001
							if πE = πg.CheckLocal(πF, µogid, "ogid"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µogid != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 82: if ogid is not None:
							πF.SetLineno(82)
						Label3:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[0] = µname
							if πTemp006, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µogid, "ogid"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple3(πTemp007, µgid, µogid).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("redefinition of group name %s as group %d; was group %d").ToObject(), πTemp005); πE != nil {
								continue
							}
							// line 83: raise error, ("redefinition of group name %s as group %d; "
							πF.SetLineno(83)
							πE = πF.Raise(πTemp001, πTemp002, nil)
							continue
							goto Label4
						Label4:
							// line 85: self.groupdict[name] = gid
							πF.SetLineno(85)
							if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µgid); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgroupdict, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005 = µname
							if πE = πg.SetItem(πF, πTemp002, πTemp005, πTemp001); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 86: self.open.append(gid)
							πF.SetLineno(86)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
								continue
							}
							πTemp004[0] = µgid
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßopen, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 87: return gid
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
								continue
							}
							πR = µgid
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßopengroup.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 88: def closegroup(self, gid):
					πF.SetLineno(88)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "gid", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("closegroup", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgid *πg.Object = πArgs[1]; _ = µgid
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
							// line 90: self.open = [x for x in self.open if x != gid]
							πF.SetLineno(90)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/sre_parse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πg.UnboundLocal; _ = µx
								var πTemp001 *πg.Object
								_ = πTemp001
								var πTemp002 *πg.Object
								_ = πTemp002
								var πTemp003 bool
								_ = πTemp003
								var πTemp004 bool
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
										if πTemp002, πE = πg.GetAttr(πF, µself, ßopen, nil); πE != nil {
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
											µx = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.NE(πF, µx, µgid); πE != nil {
											continue
										}
										if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
											continue
										}
										if πTemp004 {
											goto Label4
										}
										goto Label5
										// line 90: self.open = [x for x in self.open if x != gid]
										πF.SetLineno(90)
									Label4:
										// line 90: self.open = [x for x in self.open if x != gid]
										πF.SetLineno(90)
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µx, nil
									Label6:
										πTemp002 = πSent
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
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopen, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßclosegroup.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 91: def checkgroup(self, gid):
					πF.SetLineno(91)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "gid", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("checkgroup", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgid *πg.Object = πArgs[1]; _ = µgid
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 92: return gid < self.groups and gid not in self.open
							πF.SetLineno(92)
							if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßgroups, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µgid, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßopen, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp004, µgid); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							πTemp001 = πTemp003
						Label1:
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
					if πE = πClass.SetItem(πF, ßcheckgroup.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Pattern").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPattern.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 94: class SubPattern(object):
			πF.SetLineno(94)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp010 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SubPattern", "/usr/lib/python2.7/sre_parse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 96: def __init__(self, pattern, data=None):
					πF.SetLineno(96)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pattern", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "data", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpattern *πg.Object = πArgs[1]; _ = µpattern
						var µdata *πg.Object = πArgs[2]; _ = µdata
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
							// line 97: self.pattern = pattern
							πF.SetLineno(97)
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µpattern); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpattern, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdata == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 98: if data is None:
							πF.SetLineno(98)
						Label1:
							// line 99: data = []
							πF.SetLineno(99)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µdata = πTemp001
							goto Label2
						Label2:
							// line 100: self.data = data
							πF.SetLineno(100)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdata); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp001); πE != nil {
								continue
							}
							// line 101: self.width = None
							πF.SetLineno(101)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwidth, πTemp002); πE != nil {
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
					// line 102: def dump(self, level=0):
					πF.SetLineno(102)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "level", Def: πg.NewInt(0).ToObject()}
					πTemp003 = πg.NewFunction(πg.NewCode("dump", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlevel *πg.Object = πArgs[1]; _ = µlevel
						var µseqtypes *πg.Object = πg.UnboundLocal; _ = µseqtypes
						var µop *πg.Object = πg.UnboundLocal; _ = µop
						var µav *πg.Object = πg.UnboundLocal; _ = µav
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µcondgroup *πg.Object = πg.UnboundLocal; _ = µcondgroup
						var µitem_yes *πg.Object = πg.UnboundLocal; _ = µitem_yes
						var µitem_no *πg.Object = πg.UnboundLocal; _ = µitem_no
						var µnl *πg.Object = πg.UnboundLocal; _ = µnl
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
							case 10: goto Label10
							case 11: goto Label11
							case 13: goto Label13
							case 14: goto Label14
							case 20: goto Label20
							case 21: goto Label21
							default: panic("unexpected function state")
							}
							// line 103: seqtypes = (tuple, list)
							πF.SetLineno(103)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
							µseqtypes = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
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
								µop = πTemp003
								µav = πTemp006
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 105: print level*"  " + op,
							πF.SetLineno(105)
							πTemp007 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µlevel, πg.NewStr("  ").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µop); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.Print(πF, πTemp007, false); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µop, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBRANCH); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µop, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßGROUPREF_EXISTS); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µop, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							πTemp007[0] = µav
							if πE = πg.CheckLocal(πF, µseqtypes, "seqtypes"); πE != nil {
								continue
							}
							πTemp007[1] = µseqtypes
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 106: if op == IN:
							πF.SetLineno(106)
						Label4:
							// line 108: print
							πF.SetLineno(108)
							πTemp007 = make([]*πg.Object, 0)
							if πE = πg.Print(πF, πTemp007, true); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µav); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp005 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label12
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
									continue
								}
								µop = πTemp006
								µa = πTemp009
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(10)            
							// line 110: print (level+1)*"  " + op, a
							πF.SetLineno(110)
							πTemp007 = make([]*πg.Object, 2)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Add(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mul(πF, πTemp009, πg.NewStr("  ").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, µop); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp007[1] = µa
							if πE = πg.Print(πF, πTemp007, true); πE != nil {
								continue
							}
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							goto Label9
							// line 111: elif op == BRANCH:
							πF.SetLineno(111)
						Label5:
							// line 112: print
							πF.SetLineno(112)
							πTemp007 = make([]*πg.Object, 0)
							if πE = πg.Print(πF, πTemp007, true); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µav, πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp006
							if πTemp003, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(14)
							πTemp005 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label15
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp009}}}, πTemp003); πE != nil {
									continue
								}
								µi = πTemp006
								µa = πTemp009
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(13)            
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µi); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label16
							}
							goto Label17
							// line 114: if i:
							πF.SetLineno(114)
						Label16:
							// line 115: print level*"  " + "or"
							πF.SetLineno(115)
							πTemp007 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mul(πF, µlevel, πg.NewStr("  ").ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, ßor.ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πE = πg.Print(πF, πTemp007, true); πE != nil {
								continue
							}
							goto Label17
						Label17:
							// line 116: a.dump(level+1)
							πF.SetLineno(116)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßdump, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							goto Label9
							// line 117: elif op == GROUPREF_EXISTS:
							πF.SetLineno(117)
						Label6:
							// line 118: condgroup, item_yes, item_no = av
							πF.SetLineno(118)
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, µav); πE != nil {
								continue
							}
							µcondgroup = πTemp002
							µitem_yes = πTemp003
							µitem_no = πTemp006
							// line 119: print condgroup
							πF.SetLineno(119)
							πTemp007 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µcondgroup, "condgroup"); πE != nil {
								continue
							}
							πTemp007[0] = µcondgroup
							if πE = πg.Print(πF, πTemp007, true); πE != nil {
								continue
							}
							// line 120: item_yes.dump(level+1)
							πF.SetLineno(120)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µitem_yes, "item_yes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µitem_yes, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.CheckLocal(πF, µitem_no, "item_no"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µitem_no); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label18
							}
							goto Label19
							// line 121: if item_no:
							πF.SetLineno(121)
						Label18:
							// line 122: print level*"  " + "else"
							πF.SetLineno(122)
							πTemp007 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µlevel, πg.NewStr("  ").ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, ßelse.ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.Print(πF, πTemp007, true); πE != nil {
								continue
							}
							// line 123: item_no.dump(level+1)
							πF.SetLineno(123)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µitem_no, "item_no"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µitem_no, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label19
						Label19:
							goto Label9
							// line 124: elif isinstance(av, seqtypes):
							πF.SetLineno(124)
						Label7:
							// line 125: nl = 0
							πF.SetLineno(125)
							µnl = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µav); πE != nil {
								continue
							}
							πF.PushCheckpoint(21)
							πTemp005 = false
						Label20:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp005 {
								πF.PopCheckpoint()
								goto Label22
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
								µa = πTemp003
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(20)            
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp007[0] = µa
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSubPattern); πE != nil {
								continue
							}
							πTemp007[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label23
							}
							goto Label24
							// line 127: if isinstance(a, SubPattern):
							πF.SetLineno(127)
						Label23:
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, µnl); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp008).ToObject()
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label26
							}
							goto Label27
							// line 128: if not nl:
							πF.SetLineno(128)
						Label26:
							// line 129: print
							πF.SetLineno(129)
							πTemp007 = make([]*πg.Object, 0)
							if πE = πg.Print(πF, πTemp007, true); πE != nil {
								continue
							}
							goto Label27
						Label27:
							// line 130: a.dump(level+1)
							πF.SetLineno(130)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlevel, "level"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µlevel, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007[0] = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µa, ßdump, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 131: nl = 1
							πF.SetLineno(131)
							µnl = πg.NewInt(1).ToObject()
							goto Label25
						Label24:
							// line 133: print a,
							πF.SetLineno(133)
							πTemp007 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp007[0] = µa
							if πE = πg.Print(πF, πTemp007, false); πE != nil {
								continue
							}
							// line 134: nl = 0
							πF.SetLineno(134)
							µnl = πg.NewInt(0).ToObject()
							goto Label25
						Label25:
							continue
						Label21:
							if πE != nil || πR != nil {
								continue
							}
						Label22:
							if πE = πg.CheckLocal(πF, µnl, "nl"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µnl); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label28
							}
							goto Label29
							// line 135: if not nl:
							πF.SetLineno(135)
						Label28:
							// line 136: print
							πF.SetLineno(136)
							πTemp007 = make([]*πg.Object, 0)
							if πE = πg.Print(πF, πTemp007, true); πE != nil {
								continue
							}
							goto Label29
						Label29:
							goto Label9
						Label8:
							// line 138: print av
							πF.SetLineno(138)
							πTemp007 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							πTemp007[0] = µav
							if πE = πg.Print(πF, πTemp007, true); πE != nil {
								continue
							}
							goto Label9
						Label9:
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
					if πE = πClass.SetItem(πF, ßdump.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 139: def __repr__(self):
					πF.SetLineno(139)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 140: return repr(self.data)
							πF.SetLineno(140)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 141: def __len__(self):
					πF.SetLineno(141)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("__len__", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 142: return len(self.data)
							πF.SetLineno(142)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__len__.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 143: def __delitem__(self, index):
					πF.SetLineno(143)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("__delitem__", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
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
							// line 145: self.data = self.data[:index] + self.data[index+1:]
							πF.SetLineno(145)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µindex, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µindex, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdata, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__delitem__.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 146: def __getitem__(self, index):
					πF.SetLineno(146)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("__getitem__", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001[0] = µindex
							if πTemp002, πE = πg.ResolveGlobal(πF, ßslice); πE != nil {
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
							// line 147: if isinstance(index, slice):
							πF.SetLineno(147)
						Label1:
							// line 148: return SubPattern(self.pattern, self.data[index])
							πF.SetLineno(148)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpattern, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp002 = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSubPattern); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label2
						Label2:
							// line 149: return self.data[index]
							πF.SetLineno(149)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp002 = µindex
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__getitem__.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 150: def __setitem__(self, index, code):
					πF.SetLineno(150)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp002[2] = πg.Param{Name: "code", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__setitem__", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var µcode *πg.Object = πArgs[2]; _ = µcode
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
							// line 151: self.data[index] = code
							πF.SetLineno(151)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp003 = µindex
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
					if πE = πClass.SetItem(πF, ß__setitem__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 152: def insert(self, index, code):
					πF.SetLineno(152)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp002[2] = πg.Param{Name: "code", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("insert", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
						var µcode *πg.Object = πArgs[2]; _ = µcode
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
							// line 153: self.data.insert(index, code)
							πF.SetLineno(153)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							πTemp001[0] = µindex
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[1] = µcode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinsert, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßinsert.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 154: def append(self, code):
					πF.SetLineno(154)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "code", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("append", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcode *πg.Object = πArgs[1]; _ = µcode
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
							// line 155: self.data.append(code)
							πF.SetLineno(155)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[0] = µcode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßappend.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 156: def getwidth(self):
					πF.SetLineno(156)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("getwidth", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlo *πg.Object = πg.UnboundLocal; _ = µlo
						var µhi *πg.Object = πg.UnboundLocal; _ = µhi
						var µUNITCODES *πg.Object = πg.UnboundLocal; _ = µUNITCODES
						var µREPEATCODES *πg.Object = πg.UnboundLocal; _ = µREPEATCODES
						var µop *πg.Object = πg.UnboundLocal; _ = µop
						var µav *πg.Object = πg.UnboundLocal; _ = µav
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µj *πg.Object = πg.UnboundLocal; _ = µj
						var µl *πg.Object = πg.UnboundLocal; _ = µl
						var µh *πg.Object = πg.UnboundLocal; _ = µh
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							case 13: goto Label13
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwidth, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 158: if self.width:
							πF.SetLineno(158)
						Label1:
							// line 159: return self.width
							πF.SetLineno(159)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwidth, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 160: lo = hi = 0
							πF.SetLineno(160)
							µlo = πg.NewInt(0).ToObject()
							µhi = πg.NewInt(0).ToObject()
							// line 161: UNITCODES = (ANY, RANGE, IN, LITERAL, NOT_LITERAL, CATEGORY)
							πF.SetLineno(161)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßANY); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßRANGE); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNOT_LITERAL); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple6(πTemp003, πTemp004, πTemp005, πTemp006, πTemp007, πTemp008).ToObject()
							µUNITCODES = πTemp001
							// line 162: REPEATCODES = (MIN_REPEAT, MAX_REPEAT)
							πF.SetLineno(162)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßMIN_REPEAT); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßMAX_REPEAT); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							µREPEATCODES = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdata, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
									continue
								}
								µop = πTemp004
								µav = πTemp005
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(3)            
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßBRANCH); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßCALL); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßSUBPATTERN); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µREPEATCODES, "REPEATCODES"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, µREPEATCODES, µop); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µUNITCODES, "UNITCODES"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, µUNITCODES, µop); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label10
							}
							if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßSUCCESS); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µop, πTemp004); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label11
							}
							goto Label12
							// line 164: if op is BRANCH:
							πF.SetLineno(164)
						Label6:
							// line 165: i = MAXREPEAT - 1
							πF.SetLineno(165)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp003
							// line 166: j = 0
							πF.SetLineno(166)
							µj = πg.NewInt(0).ToObject()
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(14)
							πTemp009 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label15
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
								µav = πTemp004
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 168: l, h = av.getwidth()
							πF.SetLineno(168)
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µav, ßgetwidth, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp006}}}, πTemp005); πE != nil {
								continue
							}
							µl = πTemp004
							µh = πTemp006
							// line 169: i = min(i, l)
							πF.SetLineno(169)
							πTemp011 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp011[0] = µi
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp011[1] = µl
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							µi = πTemp005
							// line 170: j = max(j, h)
							πF.SetLineno(170)
							πTemp011 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							πTemp011[0] = µj
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							πTemp011[1] = µh
							if πTemp004, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							µj = πTemp005
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							// line 171: lo = lo + i
							πF.SetLineno(171)
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µlo, µi); πE != nil {
								continue
							}
							µlo = πTemp003
							// line 172: hi = hi + j
							πF.SetLineno(172)
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µhi, µj); πE != nil {
								continue
							}
							µhi = πTemp003
							goto Label12
							// line 173: elif op is CALL:
							πF.SetLineno(173)
						Label7:
							// line 174: i, j = av.getwidth()
							πF.SetLineno(174)
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µav, ßgetwidth, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
								continue
							}
							µi = πTemp003
							µj = πTemp005
							// line 175: lo = lo + i
							πF.SetLineno(175)
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µlo, µi); πE != nil {
								continue
							}
							µlo = πTemp003
							// line 176: hi = hi + j
							πF.SetLineno(176)
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µhi, µj); πE != nil {
								continue
							}
							µhi = πTemp003
							goto Label12
							// line 177: elif op is SUBPATTERN:
							πF.SetLineno(177)
						Label8:
							// line 178: i, j = av[1].getwidth()
							πF.SetLineno(178)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µav, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßgetwidth, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
								continue
							}
							µi = πTemp003
							µj = πTemp005
							// line 179: lo = lo + i
							πF.SetLineno(179)
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µlo, µi); πE != nil {
								continue
							}
							µlo = πTemp003
							// line 180: hi = hi + j
							πF.SetLineno(180)
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µhi, µj); πE != nil {
								continue
							}
							µhi = πTemp003
							goto Label12
							// line 181: elif op in REPEATCODES:
							πF.SetLineno(181)
						Label9:
							// line 182: i, j = av[2].getwidth()
							πF.SetLineno(182)
							πTemp003 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µav, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßgetwidth, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}}}, πTemp004); πE != nil {
								continue
							}
							µi = πTemp003
							µj = πTemp005
							// line 183: lo = lo + i * av[0]
							πF.SetLineno(183)
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µav, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, µi, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µlo, πTemp004); πE != nil {
								continue
							}
							µlo = πTemp003
							// line 184: hi = hi + j * av[1]
							πF.SetLineno(184)
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µav, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, µj, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µhi, πTemp004); πE != nil {
								continue
							}
							µhi = πTemp003
							goto Label12
							// line 185: elif op in UNITCODES:
							πF.SetLineno(185)
						Label10:
							// line 186: lo = lo + 1
							πF.SetLineno(186)
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µlo, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µlo = πTemp003
							// line 187: hi = hi + 1
							πF.SetLineno(187)
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µhi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µhi = πTemp003
							goto Label12
							// line 188: elif op == SUCCESS:
							πF.SetLineno(188)
						Label11:
							// line 189: break
							πF.SetLineno(189)
							πTemp002 = true
							continue
							goto Label12
						Label12:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 190: self.width = min(lo, MAXREPEAT - 1), min(hi, MAXREPEAT)
							πF.SetLineno(190)
							πTemp011 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							πTemp011[0] = µlo
							if πTemp004, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp011[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp011 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							πTemp011[0] = µhi
							if πTemp003, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
								continue
							}
							πTemp011[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwidth, πTemp003); πE != nil {
								continue
							}
							// line 191: return self.width
							πF.SetLineno(191)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwidth, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgetwidth.ToObject(), πTemp011); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("SubPattern").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSubPattern.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 193: class Tokenizer(object):
			πF.SetLineno(193)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp010 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp010.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Tokenizer", "/usr/lib/python2.7/sre_parse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
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
					// line 194: def __init__(self, string):
					πF.SetLineno(194)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
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
							// line 195: self.string = string
							πF.SetLineno(195)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstring); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstring, πTemp001); πE != nil {
								continue
							}
							// line 196: self.index = 0
							πF.SetLineno(196)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindex, πTemp001); πE != nil {
								continue
							}
							// line 197: self.__next()
							πF.SetLineno(197)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__next, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 198: def __next(self):
					πF.SetLineno(198)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__next", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µchar *πg.Object = πg.UnboundLocal; _ = µchar
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindex, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstring, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GE(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 199: if self.index >= len(self.string):
							πF.SetLineno(199)
						Label1:
							// line 200: self.next = None
							πF.SetLineno(200)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnext, πTemp002); πE != nil {
								continue
							}
							// line 201: return
							πF.SetLineno(201)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 202: char = self.string[self.index]
							πF.SetLineno(202)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindex, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstring, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µchar = πTemp002
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µchar, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewStr("\\").ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 203: if char[0] == "\\":
							πF.SetLineno(203)
						Label3:
							// line 204: try:
							πF.SetLineno(204)
							πF.PushCheckpoint(6)
							// line 205: c = self.string[self.index + 1]
							πF.SetLineno(205)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßindex, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstring, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µc = πTemp002
							πF.PopCheckpoint()
							goto Label5
						Label6:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 206: except IndexError:
							πF.SetLineno(206)
						Label7:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
								continue
							}
							// line 207: raise error, "bogus escape (end of line)"
							πF.SetLineno(207)
							πE = πF.Raise(πTemp001, πg.NewStr("bogus escape (end of line)").ToObject(), nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
							// line 208: char = char + c
							πF.SetLineno(208)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µchar, µc); πE != nil {
								continue
							}
							µchar = πTemp001
							goto Label4
						Label4:
							// line 209: self.index = self.index + len(char)
							πF.SetLineno(209)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindex, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp003[0] = µchar
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindex, πTemp002); πE != nil {
								continue
							}
							// line 210: self.next = char
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µchar); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnext, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__next.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 211: def match(self, char, skip=1):
					πF.SetLineno(211)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "char", Def: nil}
					πTemp002[2] = πg.Param{Name: "skip", Def: πg.NewInt(1).ToObject()}
					πTemp004 = πg.NewFunction(πg.NewCode("match", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µchar *πg.Object = πArgs[1]; _ = µchar
						var µskip *πg.Object = πArgs[2]; _ = µskip
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
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnext, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µchar, πTemp002); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 212: if char == self.next:
							πF.SetLineno(212)
						Label1:
							if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µskip); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 213: if skip:
							πF.SetLineno(213)
						Label3:
							// line 214: self.__next()
							πF.SetLineno(214)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__next, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 215: return 1
							πF.SetLineno(215)
							πR = πg.NewInt(1).ToObject()
							continue
							goto Label2
						Label2:
							// line 216: return 0
							πF.SetLineno(216)
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
					if πE = πClass.SetItem(πF, ßmatch.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 217: def get(self):
					πF.SetLineno(217)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("get", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µthis *πg.Object = πg.UnboundLocal; _ = µthis
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
							// line 218: this = self.next
							πF.SetLineno(218)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnext, nil); πE != nil {
								continue
							}
							µthis = πTemp001
							// line 219: self.__next()
							πF.SetLineno(219)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß__next, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 220: return this
							πF.SetLineno(220)
							if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
								continue
							}
							πR = µthis
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 221: def tell(self):
					πF.SetLineno(221)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("tell", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 222: return self.index, self.next
							πF.SetLineno(222)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindex, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßnext, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßtell.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 223: def seek(self, index):
					πF.SetLineno(223)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "index", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("seek", "/usr/lib/python2.7/sre_parse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindex *πg.Object = πArgs[1]; _ = µindex
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
							// line 224: self.index, self.next = index
							πF.SetLineno(224)
							if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µindex); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindex, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnext, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßseek.ToObject(), πTemp007); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp010.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Tokenizer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTokenizer.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 226: def isident(char):
			πF.SetLineno(226)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "char", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("isident", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µchar *πg.Object = πArgs[0]; _ = µchar
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
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
					// line 227: return "a" <= char <= "z" or "A" <= char <= "Z" or char == "_"
					πF.SetLineno(227)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, ßa.ToObject(), µchar); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label2
					}
					if πTemp003, πE = πg.LE(πF, µchar, ßz.ToObject()); πE != nil {
						continue
					}
				Label2:
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LE(πF, ßA.ToObject(), µchar); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label3
					}
					if πTemp003, πE = πg.LE(πF, µchar, ßZ.ToObject()); πE != nil {
						continue
					}
				Label3:
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µchar, ß_.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label1:
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
			if πE = πF.Globals().SetItem(πF, ßisident.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 229: def isdigit(char):
			πF.SetLineno(229)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "char", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("isdigit", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µchar *πg.Object = πArgs[0]; _ = µchar
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
					// line 230: return "0" <= char <= "9"
					πF.SetLineno(230)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, ß0.ToObject(), µchar); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πTemp001, πE = πg.LE(πF, µchar, ß9.ToObject()); πE != nil {
						continue
					}
				Label1:
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
			if πE = πF.Globals().SetItem(πF, ßisdigit.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 232: def isname(name):
			πF.SetLineno(232)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "name", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("isname", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µchar *πg.Object = πg.UnboundLocal; _ = µchar
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
				var πTemp009 bool
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µname, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisident); πE != nil {
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
					// line 234: if not isident(name[0]):
					πF.SetLineno(234)
				Label1:
					// line 235: return False
					πF.SetLineno(235)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp001
					continue
					goto Label2
				Label2:
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µname, πTemp003); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(4)
					πTemp005 = false
				Label3:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label5
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
						µchar = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(3)            
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp002[0] = µchar
					if πTemp007, πE = πg.ResolveGlobal(πF, ßisident); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp009).ToObject()
					πTemp003 = πTemp004
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label6
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp002[0] = µchar
					if πTemp007, πE = πg.ResolveGlobal(πF, ßisdigit); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp009).ToObject()
					πTemp003 = πTemp004
				Label6:
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					goto Label8
					// line 237: if not isident(char) and not isdigit(char):
					πF.SetLineno(237)
				Label7:
					// line 238: return False
					πF.SetLineno(238)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label8
				Label8:
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 239: return True
					πF.SetLineno(239)
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
			if πE = πF.Globals().SetItem(πF, ßisname.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 241: def _class_escape(source, escape):
			πF.SetLineno(241)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "source", Def: nil}
			πTemp011[1] = πg.Param{Name: "escape", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_class_escape", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]; _ = µsource
				var µescape *πg.Object = πArgs[1]; _ = µescape
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var µc *πg.Object = πg.UnboundLocal; _ = µc
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
				var πTemp007 bool
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 18: goto Label18
					case 19: goto Label19
					case 12: goto Label12
					case 13: goto Label13
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 243: code = ESCAPES.get(escape)
					πF.SetLineno(243)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp002, πE = πg.ResolveGlobal(πF, ßESCAPES); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp002
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcode); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 244: if code:
					πF.SetLineno(244)
				Label1:
					// line 245: return code
					πF.SetLineno(245)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πR = µcode
					continue
					goto Label2
				Label2:
					// line 246: code = CATEGORIES.get(escape)
					πF.SetLineno(246)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORIES); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp002
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp002 = µcode
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label3
					}
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µcode, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πTemp005); πE != nil {
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
					// line 247: if code and code[0] == IN:
					πF.SetLineno(247)
				Label4:
					// line 248: return code
					πF.SetLineno(248)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πR = µcode
					continue
					goto Label5
				Label5:
					// line 249: try:
					πF.SetLineno(249)
					πF.PushCheckpoint(7)
					// line 250: c = escape[1:2]
					πF.SetLineno(250)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µescape, πTemp002); πE != nil {
						continue
					}
					µc = πTemp003
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, ßx.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp003, µc); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßDIGITS); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp003, µc); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label10
					}
					goto Label11
					// line 251: if c == "x":
					πF.SetLineno(251)
				Label8:
					// line 253: while source.next in HEXDIGITS and len(escape) < 4:
					πF.SetLineno(253)
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
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßHEXDIGITS); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					πTemp002 = πTemp003
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label15
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.LT(πF, πTemp006, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label15:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(12)            
					// line 254: escape = escape + source.get()
					πF.SetLineno(254)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßget, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µescape, πTemp005); πE != nil {
						continue
					}
					µescape = πTemp002
					continue
				Label13:
					if πE != nil || πR != nil {
						continue
					}
				Label14:
					// line 255: escape = escape[2:]
					πF.SetLineno(255)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µescape, πTemp002); πE != nil {
						continue
					}
					µescape = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.NE(πF, πTemp005, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label16
					}
					goto Label17
					// line 256: if len(escape) != 2:
					πF.SetLineno(256)
				Label16:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, πg.NewStr("\\").ToObject(), µescape); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("bogus escape: %s").ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 257: raise error, "bogus escape: %s" % repr("\\" + escape)
					πF.SetLineno(257)
					πE = πF.Raise(πTemp002, πTemp003, nil)
					continue
					goto Label17
				Label17:
					// line 258: return LITERAL, int(escape, 16) & 0xff
					πF.SetLineno(258)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					πTemp001[1] = πg.NewInt(16).ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.And(πF, πTemp010, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πR = πTemp002
					continue
					goto Label11
					// line 259: elif c in OCTDIGITS:
					πF.SetLineno(259)
				Label9:
					// line 261: while source.next in OCTDIGITS and len(escape) < 4:
					πF.SetLineno(261)
					πF.PushCheckpoint(19)
					πTemp004 = false
				Label18:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label20
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					πTemp002 = πTemp003
					if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label21
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.LT(πF, πTemp006, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label21:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(18)            
					// line 262: escape = escape + source.get()
					πF.SetLineno(262)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßget, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µescape, πTemp005); πE != nil {
						continue
					}
					µescape = πTemp002
					continue
				Label19:
					if πE != nil || πR != nil {
						continue
					}
				Label20:
					// line 263: escape = escape[1:]
					πF.SetLineno(263)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µescape, πTemp002); πE != nil {
						continue
					}
					µescape = πTemp003
					// line 264: return LITERAL, int(escape, 8) & 0xff
					πF.SetLineno(264)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					πTemp001[1] = πg.NewInt(8).ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.And(πF, πTemp010, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πR = πTemp002
					continue
					goto Label11
					// line 265: elif c in DIGITS:
					πF.SetLineno(265)
				Label10:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp005, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("bogus escape: %s").ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 266: raise error, "bogus escape: %s" % repr(escape)
					πF.SetLineno(266)
					πE = πF.Raise(πTemp002, πTemp003, nil)
					continue
					goto Label11
				Label11:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label22
					}
					goto Label23
					// line 267: if len(escape) == 2:
					πF.SetLineno(267)
				Label22:
					// line 268: return LITERAL, ord(escape[1])
					πF.SetLineno(268)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µescape, πTemp005); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πR = πTemp002
					continue
					goto Label23
				Label23:
					πF.PopCheckpoint()
					goto Label6
				Label7:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label24
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 269: except ValueError:
					πF.SetLineno(269)
				Label24:
					// line 270: pass
					πF.SetLineno(270)
					πF.RestoreExc(nil, nil)
					goto Label6
				Label6:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp005, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("bogus escape: %s").ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 271: raise error, "bogus escape: %s" % repr(escape)
					πF.SetLineno(271)
					πE = πF.Raise(πTemp002, πTemp003, nil)
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_class_escape.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 273: def _escape(source, escape, state):
			πF.SetLineno(273)
			πTemp011 = make([]πg.Param, 3)
			πTemp011[0] = πg.Param{Name: "source", Def: nil}
			πTemp011[1] = πg.Param{Name: "escape", Def: nil}
			πTemp011[2] = πg.Param{Name: "state", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_escape", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]; _ = µsource
				var µescape *πg.Object = πArgs[1]; _ = µescape
				var µstate *πg.Object = πArgs[2]; _ = µstate
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 17: goto Label17
					case 18: goto Label18
					case 11: goto Label11
					case 12: goto Label12
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 275: code = CATEGORIES.get(escape)
					πF.SetLineno(275)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp002, πE = πg.ResolveGlobal(πF, ßCATEGORIES); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp002
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcode); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 276: if code:
					πF.SetLineno(276)
				Label1:
					// line 277: return code
					πF.SetLineno(277)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πR = µcode
					continue
					goto Label2
				Label2:
					// line 278: code = ESCAPES.get(escape)
					πF.SetLineno(278)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp002, πE = πg.ResolveGlobal(πF, ßESCAPES); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp002
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcode); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 279: if code:
					πF.SetLineno(279)
				Label3:
					// line 280: return code
					πF.SetLineno(280)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πR = µcode
					continue
					goto Label4
				Label4:
					// line 281: try:
					πF.SetLineno(281)
					πF.PushCheckpoint(6)
					// line 282: c = escape[1:2]
					πF.SetLineno(282)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µescape, πTemp002); πE != nil {
						continue
					}
					µc = πTemp003
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, ßx.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µc, ß0.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßDIGITS); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp003, µc); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 283: if c == "x":
					πF.SetLineno(283)
				Label7:
					// line 285: while source.next in HEXDIGITS and len(escape) < 4:
					πF.SetLineno(285)
					πF.PushCheckpoint(12)
					πTemp004 = false
				Label11:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßHEXDIGITS); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp008, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label14
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.LT(πF, πTemp008, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label14:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(11)            
					// line 286: escape = escape + source.get()
					πF.SetLineno(286)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßget, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µescape, πTemp007); πE != nil {
						continue
					}
					µescape = πTemp002
					continue
				Label12:
					if πE != nil || πR != nil {
						continue
					}
				Label13:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.NE(πF, πTemp007, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label15
					}
					goto Label16
					// line 287: if len(escape) != 4:
					πF.SetLineno(287)
				Label15:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					// line 288: raise ValueError
					πF.SetLineno(288)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label16
				Label16:
					// line 289: return LITERAL, int(escape[2:], 16) & 0xff
					πF.SetLineno(289)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(2)
					if πTemp008, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µescape, πTemp008); πE != nil {
						continue
					}
					πTemp001[0] = πTemp010
					πTemp001[1] = πg.NewInt(16).ToObject()
					if πTemp008, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.And(πF, πTemp010, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
					πR = πTemp002
					continue
					goto Label10
					// line 290: elif c == "0":
					πF.SetLineno(290)
				Label8:
					// line 292: while source.next in OCTDIGITS and len(escape) < 4:
					πF.SetLineno(292)
					πF.PushCheckpoint(18)
					πTemp004 = false
				Label17:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label19
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp008, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label20
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.LT(πF, πTemp008, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label20:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(17)            
					// line 293: escape = escape + source.get()
					πF.SetLineno(293)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßget, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µescape, πTemp007); πE != nil {
						continue
					}
					µescape = πTemp002
					continue
				Label18:
					if πE != nil || πR != nil {
						continue
					}
				Label19:
					// line 294: return LITERAL, int(escape[1:], 8) & 0xff
					πF.SetLineno(294)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(2)
					if πTemp008, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µescape, πTemp008); πE != nil {
						continue
					}
					πTemp001[0] = πTemp010
					πTemp001[1] = πg.NewInt(8).ToObject()
					if πTemp008, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.And(πF, πTemp010, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
					πR = πTemp002
					continue
					goto Label10
					// line 295: elif c in DIGITS:
					πF.SetLineno(295)
				Label9:
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßDIGITS); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label21
					}
					goto Label22
					// line 297: if source.next in DIGITS:
					πF.SetLineno(297)
				Label21:
					// line 298: escape = escape + source.get()
					πF.SetLineno(298)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßget, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µescape, πTemp007); πE != nil {
						continue
					}
					µescape = πTemp002
					πTemp007 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µescape, πTemp007); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, πTemp007, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp005).ToObject()
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label23
					}
					πTemp007 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µescape, πTemp007); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, πTemp007, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp005).ToObject()
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label23
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, πTemp008, πTemp007); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp005).ToObject()
					πTemp002 = πTemp003
				Label23:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label24
					}
					goto Label25
					// line 299: if (escape[1] in OCTDIGITS and escape[2] in OCTDIGITS and
					πF.SetLineno(299)
				Label24:
					// line 302: escape = escape + source.get()
					πF.SetLineno(302)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßget, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µescape, πTemp007); πE != nil {
						continue
					}
					µescape = πTemp002
					// line 303: return LITERAL, int(escape[1:], 8) & 0xff
					πF.SetLineno(303)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(2)
					if πTemp008, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µescape, πTemp008); πE != nil {
						continue
					}
					πTemp001[0] = πTemp010
					πTemp001[1] = πg.NewInt(8).ToObject()
					if πTemp008, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.And(πF, πTemp010, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
					πR = πTemp002
					continue
					goto Label25
				Label25:
					goto Label22
				Label22:
					// line 305: group = int(escape[1:])
					πF.SetLineno(305)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µescape, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgroup = πTemp003
					if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstate, ßgroups, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µgroup, πTemp003); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label26
					}
					goto Label27
					// line 306: if group < state.groups:
					πF.SetLineno(306)
				Label26:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
						continue
					}
					πTemp001[0] = µgroup
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstate, ßcheckgroup, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label28
					}
					goto Label29
					// line 307: if not state.checkgroup(group):
					πF.SetLineno(307)
				Label28:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 308: raise error, "cannot refer to open group"
					πF.SetLineno(308)
					πE = πF.Raise(πTemp002, πg.NewStr("cannot refer to open group").ToObject(), nil)
					continue
					goto Label29
				Label29:
					// line 314: return GROUPREF, group
					πF.SetLineno(314)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßGROUPREF); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, µgroup).ToObject()
					πR = πTemp002
					continue
					goto Label27
				Label27:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					// line 315: raise ValueError
					πF.SetLineno(315)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label10
				Label10:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Eq(πF, πTemp007, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label30
					}
					goto Label31
					// line 316: if len(escape) == 2:
					πF.SetLineno(316)
				Label30:
					// line 317: return LITERAL, ord(escape[1])
					πF.SetLineno(317)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp007 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µescape, πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp008
					if πTemp007, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp008).ToObject()
					πR = πTemp002
					continue
					goto Label31
				Label31:
					πF.PopCheckpoint()
					goto Label5
				Label6:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label32
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 318: except ValueError:
					πF.SetLineno(318)
				Label32:
					// line 319: pass
					πF.SetLineno(319)
					πF.RestoreExc(nil, nil)
					goto Label5
				Label5:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µescape, "escape"); πE != nil {
						continue
					}
					πTemp001[0] = µescape
					if πTemp007, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("bogus escape: %s").ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 320: raise error, "bogus escape: %s" % repr(escape)
					πF.SetLineno(320)
					πE = πF.Raise(πTemp002, πTemp003, nil)
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_escape.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 322: def _parse_sub(source, state, nested=1):
			πF.SetLineno(322)
			πTemp011 = make([]πg.Param, 3)
			πTemp011[0] = πg.Param{Name: "source", Def: nil}
			πTemp011[1] = πg.Param{Name: "state", Def: nil}
			πTemp011[2] = πg.Param{Name: "nested", Def: πg.NewInt(1).ToObject()}
			πTemp009 = πg.NewFunction(πg.NewCode("_parse_sub", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]; _ = µsource
				var µstate *πg.Object = πArgs[1]; _ = µstate
				var µnested *πg.Object = πArgs[2]; _ = µnested
				var µitems *πg.Object = πg.UnboundLocal; _ = µitems
				var µitemsappend *πg.Object = πg.UnboundLocal; _ = µitemsappend
				var µsourcematch *πg.Object = πg.UnboundLocal; _ = µsourcematch
				var µsubpattern *πg.Object = πg.UnboundLocal; _ = µsubpattern
				var µsubpatternappend *πg.Object = πg.UnboundLocal; _ = µsubpatternappend
				var µprefix *πg.Object = πg.UnboundLocal; _ = µprefix
				var µcommon *πg.Object = πg.UnboundLocal; _ = µcommon
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µset *πg.Object = πg.UnboundLocal; _ = µset
				var µsetappend *πg.Object = πg.UnboundLocal; _ = µsetappend
				var πTemp001 []*πg.Object
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
				var πTemp013 *πg.Object
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 36: goto Label36
					case 37: goto Label37
					case 14: goto Label14
					case 15: goto Label15
					case 17: goto Label17
					case 18: goto Label18
					case 25: goto Label25
					case 26: goto Label26
					case 30: goto Label30
					case 31: goto Label31
					default: panic("unexpected function state")
					}
					// line 325: items = []
					πF.SetLineno(325)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µitems = πTemp002
					// line 326: itemsappend = items.append
					πF.SetLineno(326)
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µitems, ßappend, nil); πE != nil {
						continue
					}
					µitemsappend = πTemp002
					// line 327: sourcematch = source.match
					πF.SetLineno(327)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource, ßmatch, nil); πE != nil {
						continue
					}
					µsourcematch = πTemp002
					// line 328: while 1:
					πF.SetLineno(328)
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
					if πTemp004, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 329: itemsappend(_parse(source, state))
					πF.SetLineno(329)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp005[0] = µsource
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp005[1] = µstate
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µitemsappend, "itemsappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µitemsappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("|").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 330: if sourcematch("|"):
					πF.SetLineno(330)
				Label4:
					// line 331: continue
					πF.SetLineno(331)
					continue
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µnested, "nested"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µnested); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					goto Label7
					// line 332: if not nested:
					πF.SetLineno(332)
				Label6:
					// line 333: break
					πF.SetLineno(333)
					πTemp003 = true
					continue
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp007); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(!πTemp008).ToObject()
					πTemp002 = πTemp006
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr(")").ToObject()
					πTemp001[1] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp006, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πTemp006
				Label8:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 334: if not source.next or sourcematch(")", 0):
					πF.SetLineno(334)
				Label9:
					// line 335: break
					πF.SetLineno(335)
					πTemp003 = true
					continue
					goto Label11
				Label10:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 337: raise error, "pattern not properly closed"
					πF.SetLineno(337)
					πE = πF.Raise(πTemp002, πg.NewStr("pattern not properly closed").ToObject(), nil)
					continue
					goto Label11
				Label11:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp001[0] = µitems
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label12
					}
					goto Label13
					// line 339: if len(items) == 1:
					πF.SetLineno(339)
				Label12:
					// line 340: return items[0]
					πF.SetLineno(340)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µitems, πTemp002); πE != nil {
						continue
					}
					πR = πTemp006
					continue
					goto Label13
				Label13:
					// line 342: subpattern = SubPattern(state)
					πF.SetLineno(342)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[0] = µstate
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSubPattern); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsubpattern = πTemp006
					// line 343: subpatternappend = subpattern.append
					πF.SetLineno(343)
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsubpattern, ßappend, nil); πE != nil {
						continue
					}
					µsubpatternappend = πTemp002
					// line 346: while 1:
					πF.SetLineno(346)
					πF.PushCheckpoint(15)
					πTemp003 = false
				Label14:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label16
					}
					if πTemp004, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(14)            
					// line 347: prefix, common = None, False
					πF.SetLineno(347)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
						continue
					}
					µprefix = πTemp006
					µcommon = πTemp007
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µitems); πE != nil {
						continue
					}
					πF.PushCheckpoint(18)
					πTemp004 = false
				Label17:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label19
					}
					if πTemp006, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µitem = πTemp006
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(17)            
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µitem); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(!πTemp008).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label20
					}
					goto Label21
					// line 349: if not item:
					πF.SetLineno(349)
				Label20:
					// line 350: break
					πF.SetLineno(350)
					πTemp004 = true
					continue
					goto Label21
				Label21:
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(µprefix == πTemp007).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label22
					}
					πTemp007 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µitem, πTemp007); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.NE(πF, πTemp009, µprefix); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label23
					}
					goto Label24
					// line 351: if prefix is None:
					πF.SetLineno(351)
				Label22:
					// line 352: prefix = item[0]
					πF.SetLineno(352)
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µitem, πTemp006); πE != nil {
						continue
					}
					µprefix = πTemp007
					goto Label24
					// line 353: elif item[0] != prefix:
					πF.SetLineno(353)
				Label23:
					// line 354: break
					πF.SetLineno(354)
					πTemp004 = true
					continue
					goto Label24
				Label24:
					continue
				Label18:
					if πE != nil || πR != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp005[0] = µitems
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp009
					if πTemp007, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.Iter(πF, πTemp009); πE != nil {
						continue
					}
					πF.PushCheckpoint(26)
					πTemp008 = false
				Label25:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label27
					}
					if πTemp007, πE = πg.Next(πF, πTemp006); πE != nil {
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
					πF.PushCheckpoint(25)            
					// line 362: items[i] = items[i][1:]
					πF.SetLineno(362)
					if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp011 = µi
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µitems, πTemp011); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, πTemp012, πTemp007); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp009); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp011 = µi
					if πE = πg.SetItem(πF, µitems, πTemp011, πTemp007); πE != nil {
						continue
					}
					continue
				Label26:
					if πE != nil || πR != nil {
						continue
					}
				Label27:
					// line 363: subpatternappend(prefix)
					πF.SetLineno(363)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp001[0] = µprefix
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp006, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 365: common = True
					πF.SetLineno(365)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µcommon = πTemp006
				Label19:
					if πE = πg.CheckLocal(πF, µcommon, "common"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcommon); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label28
					}
					goto Label29
					// line 366: if common:
					πF.SetLineno(366)
				Label28:
					// line 367: continue
					πF.SetLineno(367)
					continue
					goto Label29
				Label29:
					// line 368: break
					πF.SetLineno(368)
					πTemp003 = true
					continue
					continue
				Label15:
					if πE != nil || πR != nil {
						continue
					}
				Label16:
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µitems); πE != nil {
						continue
					}
					πF.PushCheckpoint(31)
					πTemp003 = false
				Label30:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp003 {
						πF.PopCheckpoint()
						goto Label32
					}
					if πTemp006, πE = πg.Next(πF, πTemp002); πE != nil {
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
						µitem = πTemp006
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(30)            
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.NE(πF, πTemp011, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πTemp007
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label33
					}
					πTemp009 = πg.NewInt(0).ToObject()
					πTemp012 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetItem(πF, µitem, πTemp012); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, πTemp013, πTemp009); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					if πTemp007, πE = πg.NE(πF, πTemp011, πTemp009); πE != nil {
						continue
					}
					πTemp006 = πTemp007
				Label33:
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label34
					}
					goto Label35
					// line 372: if len(item) != 1 or item[0][0] != LITERAL:
					πF.SetLineno(372)
				Label34:
					// line 373: break
					πF.SetLineno(373)
					πTemp003 = true
					continue
					goto Label35
				Label35:
					continue
				Label31:
					if πE != nil || πR != nil {
						continue
					}
					// line 377: set = []
					πF.SetLineno(377)
					πTemp001 = make([]*πg.Object, 0)
					πTemp006 = πg.NewList(πTemp001...).ToObject()
					µset = πTemp006
					// line 378: setappend = set.append
					πF.SetLineno(378)
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µset, ßappend, nil); πE != nil {
						continue
					}
					µsetappend = πTemp006
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Iter(πF, µitems); πE != nil {
						continue
					}
					πF.PushCheckpoint(37)
					πTemp004 = false
				Label36:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label38
					}
					if πTemp007, πE = πg.Next(πF, πTemp006); πE != nil {
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
						µitem = πTemp007
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(36)            
					// line 380: setappend(item[0])
					πF.SetLineno(380)
					πTemp001 = πF.MakeArgs(1)
					πTemp007 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µitem, πTemp007); πE != nil {
						continue
					}
					πTemp001[0] = πTemp009
					if πE = πg.CheckLocal(πF, µsetappend, "setappend"); πE != nil {
						continue
					}
					if πTemp007, πE = µsetappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label37:
					if πE != nil || πR != nil {
						continue
					}
				Label38:
					// line 381: subpatternappend((IN, set))
					πF.SetLineno(381)
					πTemp001 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(πTemp007, µset).ToObject()
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp006, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 382: return subpattern
					πF.SetLineno(382)
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					πR = µsubpattern
					continue
				Label32:
					// line 384: subpattern.append((BRANCH, (None, items)))
					πF.SetLineno(384)
					πTemp001 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßBRANCH); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
						continue
					}
					πTemp007 = πg.NewTuple2(πTemp009, µitems).ToObject()
					πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsubpattern, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 385: return subpattern
					πF.SetLineno(385)
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					πR = µsubpattern
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_parse_sub.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 387: def _parse_sub_cond(source, state, condgroup):
			πF.SetLineno(387)
			πTemp011 = make([]πg.Param, 3)
			πTemp011[0] = πg.Param{Name: "source", Def: nil}
			πTemp011[1] = πg.Param{Name: "state", Def: nil}
			πTemp011[2] = πg.Param{Name: "condgroup", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("_parse_sub_cond", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]; _ = µsource
				var µstate *πg.Object = πArgs[1]; _ = µstate
				var µcondgroup *πg.Object = πArgs[2]; _ = µcondgroup
				var µitem_yes *πg.Object = πg.UnboundLocal; _ = µitem_yes
				var µitem_no *πg.Object = πg.UnboundLocal; _ = µitem_no
				var µsubpattern *πg.Object = πg.UnboundLocal; _ = µsubpattern
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
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 388: item_yes = _parse(source, state)
					πF.SetLineno(388)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[1] = µstate
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µitem_yes = πTemp003
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("|").ToObject()
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource, ßmatch, nil); πE != nil {
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
					// line 389: if source.match("|"):
					πF.SetLineno(389)
				Label1:
					// line 390: item_no = _parse(source, state)
					πF.SetLineno(390)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[1] = µstate
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µitem_no = πTemp003
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("|").ToObject()
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource, ßmatch, nil); πE != nil {
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
					// line 391: if source.match("|"):
					πF.SetLineno(391)
				Label4:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 392: raise error, "conditional backref with more than two branches"
					πF.SetLineno(392)
					πE = πF.Raise(πTemp002, πg.NewStr("conditional backref with more than two branches").ToObject(), nil)
					continue
					goto Label5
				Label5:
					goto Label3
				Label2:
					// line 394: item_no = None
					πF.SetLineno(394)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µitem_no = πTemp002
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label6
					}
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr(")").ToObject()
					πTemp001[1] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µsource, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp002 = πTemp003
				Label6:
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label7
					}
					goto Label8
					// line 395: if source.next and not source.match(")", 0):
					πF.SetLineno(395)
				Label7:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 396: raise error, "pattern not properly closed"
					πF.SetLineno(396)
					πE = πF.Raise(πTemp002, πg.NewStr("pattern not properly closed").ToObject(), nil)
					continue
					goto Label8
				Label8:
					// line 397: subpattern = SubPattern(state)
					πF.SetLineno(397)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[0] = µstate
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSubPattern); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsubpattern = πTemp003
					// line 398: subpattern.append((GROUPREF_EXISTS, (condgroup, item_yes, item_no)))
					πF.SetLineno(398)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßGROUPREF_EXISTS); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcondgroup, "condgroup"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem_yes, "item_yes"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem_no, "item_no"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple3(µcondgroup, µitem_yes, µitem_no).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp005).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsubpattern, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 399: return subpattern
					πF.SetLineno(399)
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					πR = µsubpattern
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_parse_sub_cond.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 401: _PATTERNENDERS = set("|)")
			πF.SetLineno(401)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("|)").ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_PATTERNENDERS.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 402: _ASSERTCHARS = set("=!<")
			πF.SetLineno(402)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("=!<").ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_ASSERTCHARS.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 403: _LOOKBEHINDASSERTCHARS = set("=!")
			πF.SetLineno(403)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("=!").ToObject()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_LOOKBEHINDASSERTCHARS.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 404: _REPEATCODES = set([MIN_REPEAT, MAX_REPEAT])
			πF.SetLineno(404)
			πTemp002 = πF.MakeArgs(1)
			πTemp015 = make([]*πg.Object, 2)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßMIN_REPEAT); πE != nil {
				continue
			}
			πTemp015[0] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßMAX_REPEAT); πE != nil {
				continue
			}
			πTemp015[1] = πTemp013
			πTemp013 = πg.NewList(πTemp015...).ToObject()
			πTemp002[0] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_REPEATCODES.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 406: def _parse(source, state):
			πF.SetLineno(406)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "source", Def: nil}
			πTemp011[1] = πg.Param{Name: "state", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_parse", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]; _ = µsource
				var µstate *πg.Object = πArgs[1]; _ = µstate
				var µsubpattern *πg.Object = πg.UnboundLocal; _ = µsubpattern
				var µsubpatternappend *πg.Object = πg.UnboundLocal; _ = µsubpatternappend
				var µsourceget *πg.Object = πg.UnboundLocal; _ = µsourceget
				var µsourcematch *πg.Object = πg.UnboundLocal; _ = µsourcematch
				var µ_len *πg.Object = πg.UnboundLocal; _ = µ_len
				var µPATTERNENDERS *πg.Object = πg.UnboundLocal; _ = µPATTERNENDERS
				var µASSERTCHARS *πg.Object = πg.UnboundLocal; _ = µASSERTCHARS
				var µLOOKBEHINDASSERTCHARS *πg.Object = πg.UnboundLocal; _ = µLOOKBEHINDASSERTCHARS
				var µREPEATCODES *πg.Object = πg.UnboundLocal; _ = µREPEATCODES
				var µthis *πg.Object = πg.UnboundLocal; _ = µthis
				var µset *πg.Object = πg.UnboundLocal; _ = µset
				var µsetappend *πg.Object = πg.UnboundLocal; _ = µsetappend
				var µstart *πg.Object = πg.UnboundLocal; _ = µstart
				var µcode1 *πg.Object = πg.UnboundLocal; _ = µcode1
				var µcode2 *πg.Object = πg.UnboundLocal; _ = µcode2
				var µlo *πg.Object = πg.UnboundLocal; _ = µlo
				var µhi *πg.Object = πg.UnboundLocal; _ = µhi
				var µmin *πg.Object = πg.UnboundLocal; _ = µmin
				var µmax *πg.Object = πg.UnboundLocal; _ = µmax
				var µhere *πg.Object = πg.UnboundLocal; _ = µhere
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
				var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µcondgroup *πg.Object = πg.UnboundLocal; _ = µcondgroup
				var µchar *πg.Object = πg.UnboundLocal; _ = µchar
				var µgid *πg.Object = πg.UnboundLocal; _ = µgid
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var µdir *πg.Object = πg.UnboundLocal; _ = µdir
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µcondname *πg.Object = πg.UnboundLocal; _ = µcondname
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
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
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 bool
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.BaseException
				_ = πTemp014
				var πTemp015 *πg.Traceback
				_ = πTemp015
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					case 134: goto Label134
					case 135: goto Label135
					case 14: goto Label14
					case 15: goto Label15
					case 149: goto Label149
					case 150: goto Label150
					case 34: goto Label34
					case 35: goto Label35
					case 168: goto Label168
					case 169: goto Label169
					case 183: goto Label183
					case 187: goto Label187
					case 188: goto Label188
					case 203: goto Label203
					case 204: goto Label204
					case 77: goto Label77
					case 78: goto Label78
					case 83: goto Label83
					case 84: goto Label84
					case 123: goto Label123
					case 124: goto Label124
					default: panic("unexpected function state")
					}
					// line 408: subpattern = SubPattern(state)
					πF.SetLineno(408)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[0] = µstate
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSubPattern); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsubpattern = πTemp003
					// line 411: subpatternappend = subpattern.append
					πF.SetLineno(411)
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsubpattern, ßappend, nil); πE != nil {
						continue
					}
					µsubpatternappend = πTemp002
					// line 412: sourceget = source.get
					πF.SetLineno(412)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource, ßget, nil); πE != nil {
						continue
					}
					µsourceget = πTemp002
					// line 413: sourcematch = source.match
					πF.SetLineno(413)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource, ßmatch, nil); πE != nil {
						continue
					}
					µsourcematch = πTemp002
					// line 414: _len = len
					πF.SetLineno(414)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					µ_len = πTemp002
					// line 415: PATTERNENDERS = _PATTERNENDERS
					πF.SetLineno(415)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_PATTERNENDERS); πE != nil {
						continue
					}
					µPATTERNENDERS = πTemp002
					// line 416: ASSERTCHARS = _ASSERTCHARS
					πF.SetLineno(416)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_ASSERTCHARS); πE != nil {
						continue
					}
					µASSERTCHARS = πTemp002
					// line 417: LOOKBEHINDASSERTCHARS = _LOOKBEHINDASSERTCHARS
					πF.SetLineno(417)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_LOOKBEHINDASSERTCHARS); πE != nil {
						continue
					}
					µLOOKBEHINDASSERTCHARS = πTemp002
					// line 418: REPEATCODES = _REPEATCODES
					πF.SetLineno(418)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_REPEATCODES); πE != nil {
						continue
					}
					µREPEATCODES = πTemp002
					// line 420: while 1:
					πF.SetLineno(420)
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
					if πTemp005, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µPATTERNENDERS, "PATTERNENDERS"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µPATTERNENDERS, πTemp003); πE != nil {
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
					// line 422: if source.next in PATTERNENDERS:
					πF.SetLineno(422)
				Label4:
					// line 423: break # end of subpattern
					πF.SetLineno(423)
					πTemp004 = true
					continue
					goto Label5
				Label5:
					// line 424: this = sourceget()
					πF.SetLineno(424)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µthis = πTemp002
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µthis == πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 425: if this is None:
					πF.SetLineno(425)
				Label6:
					// line 426: break # end of pattern
					πF.SetLineno(426)
					πTemp004 = true
					continue
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstate, ßflags, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_VERBOSE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, πTemp003, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 428: if state.flags & SRE_FLAG_VERBOSE:
					πF.SetLineno(428)
				Label8:
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßWHITESPACE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, πTemp003, µthis); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label10
					}
					goto Label11
					// line 430: if this in WHITESPACE:
					πF.SetLineno(430)
				Label10:
					// line 431: continue
					πF.SetLineno(431)
					continue
					goto Label11
				Label11:
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("#").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label12
					}
					goto Label13
					// line 432: if this == "#":
					πF.SetLineno(432)
				Label12:
					// line 433: while 1:
					πF.SetLineno(433)
					πF.PushCheckpoint(15)
					πTemp005 = false
				Label14:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label16
					}
					if πTemp007, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(14)            
					// line 434: this = sourceget()
					πF.SetLineno(434)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µthis = πTemp002
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp006, πg.NewStr("\n").ToObject()).ToObject()
					if πTemp007, πE = πg.Contains(πF, πTemp003, µthis); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label17
					}
					goto Label18
					// line 435: if this in (None, "\n"):
					πF.SetLineno(435)
				Label17:
					// line 436: break
					πF.SetLineno(436)
					πTemp005 = true
					continue
					goto Label18
				Label18:
					continue
				Label15:
					if πE != nil || πR != nil {
						continue
					}
				Label16:
					// line 437: continue
					πF.SetLineno(437)
					continue
					goto Label13
				Label13:
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp002 = µthis
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label19
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µthis, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßSPECIAL_CHARS); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, πTemp006, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp002 = πTemp003
				Label19:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label20
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("[").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label21
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp002 = µthis
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label22
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µthis, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßREPEAT_CHARS); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, πTemp006, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp007).ToObject()
					πTemp002 = πTemp003
				Label22:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label23
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label24
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("(").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label25
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("^").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label26
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("$").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label27
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp002 = µthis
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label28
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µthis, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp008, πg.NewStr("\\").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label28:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label29
					}
					goto Label30
					// line 439: if this and this[0] not in SPECIAL_CHARS:
					πF.SetLineno(439)
				Label20:
					// line 440: subpatternappend((LITERAL, ord(this)))
					πF.SetLineno(440)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp009[0] = µthis
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp008).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label31
					// line 442: elif this == "[":
					πF.SetLineno(442)
				Label21:
					// line 444: set = []
					πF.SetLineno(444)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µset = πTemp002
					// line 445: setappend = set.append
					πF.SetLineno(445)
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µset, ßappend, nil); πE != nil {
						continue
					}
					µsetappend = πTemp002
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("^").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label32
					}
					goto Label33
					// line 448: if sourcematch("^"):
					πF.SetLineno(448)
				Label32:
					// line 449: setappend((NEGATE, None))
					πF.SetLineno(449)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNEGATE); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsetappend, "setappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsetappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label33
				Label33:
					// line 451: start = set[:]
					πF.SetLineno(451)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µset, πTemp002); πE != nil {
						continue
					}
					µstart = πTemp003
					// line 452: while 1:
					πF.SetLineno(452)
					πF.PushCheckpoint(35)
					πTemp005 = false
				Label34:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label36
					}
					if πTemp007, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(34)            
					// line 453: this = sourceget()
					πF.SetLineno(453)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µthis = πTemp002
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µthis, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label37
					}
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, µset, µstart); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label37:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label38
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp002 = µthis
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label39
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µthis, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp008, πg.NewStr("\\").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label39:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label40
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µthis); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label41
					}
					goto Label42
					// line 454: if this == "]" and set != start:
					πF.SetLineno(454)
				Label38:
					// line 455: break
					πF.SetLineno(455)
					πTemp005 = true
					continue
					goto Label43
					// line 456: elif this and this[0] == "\\":
					πF.SetLineno(456)
				Label40:
					// line 457: code1 = _class_escape(source, this)
					πF.SetLineno(457)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp001[1] = µthis
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_class_escape); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode1 = πTemp003
					goto Label43
					// line 458: elif this:
					πF.SetLineno(458)
				Label41:
					// line 459: code1 = LITERAL, ord(this)
					πF.SetLineno(459)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp001[0] = µthis
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp008).ToObject()
					µcode1 = πTemp002
					goto Label43
				Label42:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 461: raise error, "unexpected end of regular expression"
					πF.SetLineno(461)
					πE = πF.Raise(πTemp002, πg.NewStr("unexpected end of regular expression").ToObject(), nil)
					continue
					goto Label43
				Label43:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("-").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label44
					}
					goto Label45
					// line 462: if sourcematch("-"):
					πF.SetLineno(462)
				Label44:
					// line 464: this = sourceget()
					πF.SetLineno(464)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µthis = πTemp002
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label47
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µthis); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label48
					}
					goto Label49
					// line 465: if this == "]":
					πF.SetLineno(465)
				Label47:
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcode1, "code1"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µcode1, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006 == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label51
					}
					goto Label52
					// line 466: if code1[0] is IN:
					πF.SetLineno(466)
				Label51:
					// line 467: code1 = code1[1][0]
					πF.SetLineno(467)
					πTemp002 = πg.NewInt(0).ToObject()
					πTemp006 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µcode1, "code1"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µcode1, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
						continue
					}
					µcode1 = πTemp003
					goto Label52
				Label52:
					// line 468: setappend(code1)
					πF.SetLineno(468)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode1, "code1"); πE != nil {
						continue
					}
					πTemp001[0] = µcode1
					if πE = πg.CheckLocal(πF, µsetappend, "setappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsetappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 469: setappend((LITERAL, ord("-")))
					πF.SetLineno(469)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp009 = πF.MakeArgs(1)
					πTemp009[0] = πg.NewStr("-").ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp008).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsetappend, "setappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsetappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 470: break
					πF.SetLineno(470)
					πTemp005 = true
					continue
					goto Label50
					// line 471: elif this:
					πF.SetLineno(471)
				Label48:
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µthis, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewStr("\\").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label53
					}
					goto Label54
					// line 472: if this[0] == "\\":
					πF.SetLineno(472)
				Label53:
					// line 473: code2 = _class_escape(source, this)
					πF.SetLineno(473)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp001[1] = µthis
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_class_escape); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode2 = πTemp003
					goto Label55
				Label54:
					// line 475: code2 = LITERAL, ord(this)
					πF.SetLineno(475)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp001[0] = µthis
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp008).ToObject()
					µcode2 = πTemp002
					goto Label55
				Label55:
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcode1, "code1"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µcode1, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp008, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label56
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcode2, "code2"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µcode2, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp008, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label56:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label57
					}
					goto Label58
					// line 476: if code1[0] != LITERAL or code2[0] != LITERAL:
					πF.SetLineno(476)
				Label57:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 477: raise error, "bad character range"
					πF.SetLineno(477)
					πE = πF.Raise(πTemp002, πg.NewStr("bad character range").ToObject(), nil)
					continue
					goto Label58
				Label58:
					// line 478: lo = code1[1]
					πF.SetLineno(478)
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µcode1, "code1"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µcode1, πTemp002); πE != nil {
						continue
					}
					µlo = πTemp003
					// line 479: hi = code2[1]
					πF.SetLineno(479)
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µcode2, "code2"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µcode2, πTemp002); πE != nil {
						continue
					}
					µhi = πTemp003
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µhi, µlo); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label59
					}
					goto Label60
					// line 480: if hi < lo:
					πF.SetLineno(480)
				Label59:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 481: raise error, "bad character range"
					πF.SetLineno(481)
					πE = πF.Raise(πTemp002, πg.NewStr("bad character range").ToObject(), nil)
					continue
					goto Label60
				Label60:
					// line 482: setappend((RANGE, (lo, hi)))
					πF.SetLineno(482)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßRANGE); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(µlo, µhi).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsetappend, "setappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsetappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label50
				Label49:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 484: raise error, "unexpected end of regular expression"
					πF.SetLineno(484)
					πE = πF.Raise(πTemp002, πg.NewStr("unexpected end of regular expression").ToObject(), nil)
					continue
					goto Label50
				Label50:
					goto Label46
				Label45:
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µcode1, "code1"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µcode1, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006 == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label61
					}
					goto Label62
					// line 486: if code1[0] is IN:
					πF.SetLineno(486)
				Label61:
					// line 487: code1 = code1[1][0]
					πF.SetLineno(487)
					πTemp002 = πg.NewInt(0).ToObject()
					πTemp006 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µcode1, "code1"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µcode1, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
						continue
					}
					µcode1 = πTemp003
					goto Label62
				Label62:
					// line 488: setappend(code1)
					πF.SetLineno(488)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode1, "code1"); πE != nil {
						continue
					}
					πTemp001[0] = µcode1
					if πE = πg.CheckLocal(πF, µsetappend, "setappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsetappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label46
				Label46:
					continue
				Label35:
					if πE != nil || πR != nil {
						continue
					}
				Label36:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					πTemp001[0] = µset
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp006, πE = µ_len.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label63
					}
					πTemp006 = πg.NewInt(0).ToObject()
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µset, πTemp010); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp011, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008 == πTemp006).ToObject()
					πTemp002 = πTemp003
				Label63:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label64
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					πTemp001[0] = µset
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp006, πE = µ_len.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label65
					}
					πTemp006 = πg.NewInt(0).ToObject()
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µset, πTemp010); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp011, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNEGATE); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008 == πTemp006).ToObject()
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label65
					}
					πTemp006 = πg.NewInt(0).ToObject()
					πTemp010 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µset, πTemp010); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp011, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008 == πTemp006).ToObject()
					πTemp002 = πTemp003
				Label65:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label66
					}
					goto Label67
					// line 491: if _len(set)==1 and set[0][0] is LITERAL:
					πF.SetLineno(491)
				Label64:
					// line 492: subpatternappend(set[0]) # optimization
					πF.SetLineno(492)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µset, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label68
					// line 493: elif _len(set)==2 and set[0][0] is NEGATE and set[1][0] is LITERAL:
					πF.SetLineno(493)
				Label66:
					// line 494: subpatternappend((NOT_LITERAL, set[1][1])) # optimization
					πF.SetLineno(494)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNOT_LITERAL); πE != nil {
						continue
					}
					πTemp006 = πg.NewInt(1).ToObject()
					πTemp010 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetItem(πF, µset, πTemp010); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp011, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp008).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label68
				Label67:
					// line 497: subpatternappend((IN, set))
					πF.SetLineno(497)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µset, "set"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, µset).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label68
				Label68:
					goto Label31
					// line 499: elif this and this[0] in REPEAT_CHARS:
					πF.SetLineno(499)
				Label23:
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("?").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label69
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("*").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label70
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("+").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label71
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µthis, πg.NewStr("{").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label72
					}
					goto Label73
					// line 501: if this == "?":
					πF.SetLineno(501)
				Label69:
					// line 502: min, max = 0, 1
					πF.SetLineno(502)
					πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πg.NewInt(1).ToObject()).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
						continue
					}
					µmin = πTemp003
					µmax = πTemp006
					goto Label74
					// line 503: elif this == "*":
					πF.SetLineno(503)
				Label70:
					// line 504: min, max = 0, MAXREPEAT
					πF.SetLineno(504)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp003).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
						continue
					}
					µmin = πTemp003
					µmax = πTemp006
					goto Label74
					// line 506: elif this == "+":
					πF.SetLineno(506)
				Label71:
					// line 507: min, max = 1, MAXREPEAT
					πF.SetLineno(507)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πg.NewInt(1).ToObject(), πTemp003).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
						continue
					}
					µmin = πTemp003
					µmax = πTemp006
					goto Label74
					// line 508: elif this == "{":
					πF.SetLineno(508)
				Label72:
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp003, πg.NewStr("}").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label75
					}
					goto Label76
					// line 509: if source.next == "}":
					πF.SetLineno(509)
				Label75:
					// line 510: subpatternappend((LITERAL, ord(this)))
					πF.SetLineno(510)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp009[0] = µthis
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp008).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 511: continue
					πF.SetLineno(511)
					continue
					goto Label76
				Label76:
					// line 512: here = source.tell()
					πF.SetLineno(512)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource, ßtell, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µhere = πTemp003
					// line 513: min, max = 0, MAXREPEAT
					πF.SetLineno(513)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp003).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
						continue
					}
					µmin = πTemp003
					µmax = πTemp006
					// line 514: lo = hi = ""
					πF.SetLineno(514)
					µlo = ß.ToObject()
					µhi = ß.ToObject()
					// line 515: while source.next in DIGITS:
					πF.SetLineno(515)
					πF.PushCheckpoint(78)
					πTemp005 = false
				Label77:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label79
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßDIGITS); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Contains(πF, πTemp006, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp012).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(77)            
					// line 516: lo = lo + source.get()
					πF.SetLineno(516)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßget, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µlo, πTemp006); πE != nil {
						continue
					}
					µlo = πTemp002
					continue
				Label78:
					if πE != nil || πR != nil {
						continue
					}
				Label79:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(",").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label80
					}
					goto Label81
					// line 517: if sourcematch(","):
					πF.SetLineno(517)
				Label80:
					// line 518: while source.next in DIGITS:
					πF.SetLineno(518)
					πF.PushCheckpoint(84)
					πTemp005 = false
				Label83:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label85
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßDIGITS); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Contains(πF, πTemp006, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp012).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(83)            
					// line 519: hi = hi + sourceget()
					πF.SetLineno(519)
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp003, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µhi, πTemp003); πE != nil {
						continue
					}
					µhi = πTemp002
					continue
				Label84:
					if πE != nil || πR != nil {
						continue
					}
				Label85:
					goto Label82
				Label81:
					// line 521: hi = lo
					πF.SetLineno(521)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					µhi = µlo
					goto Label82
				Label82:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("}").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp003, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label86
					}
					goto Label87
					// line 522: if not sourcematch("}"):
					πF.SetLineno(522)
				Label86:
					// line 523: subpatternappend((LITERAL, ord(this)))
					πF.SetLineno(523)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp009[0] = µthis
					if πTemp006, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp002 = πg.NewTuple2(πTemp003, πTemp008).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 524: source.seek(here)
					πF.SetLineno(524)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhere, "here"); πE != nil {
						continue
					}
					πTemp001[0] = µhere
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource, ßseek, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 525: continue
					πF.SetLineno(525)
					continue
					goto Label87
				Label87:
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µlo); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label88
					}
					goto Label89
					// line 526: if lo:
					πF.SetLineno(526)
				Label88:
					// line 527: min = int(lo)
					πF.SetLineno(527)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp001[0] = µlo
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmin = πTemp003
					if πE = πg.CheckLocal(πF, µmin, "min"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µmin, πTemp003); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label90
					}
					goto Label91
					// line 528: if min >= MAXREPEAT:
					πF.SetLineno(528)
				Label90:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("the repetition number is too large").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 529: raise OverflowError("the repetition number is too large")
					πF.SetLineno(529)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label91
				Label91:
					goto Label89
				Label89:
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µhi); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label92
					}
					goto Label93
					// line 530: if hi:
					πF.SetLineno(530)
				Label92:
					// line 531: max = int(hi)
					πF.SetLineno(531)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					πTemp001[0] = µhi
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µmax = πTemp003
					if πE = πg.CheckLocal(πF, µmax, "max"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, µmax, πTemp003); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label94
					}
					goto Label95
					// line 532: if max >= MAXREPEAT:
					πF.SetLineno(532)
				Label94:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("the repetition number is too large").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 533: raise OverflowError("the repetition number is too large")
					πF.SetLineno(533)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label95
				Label95:
					if πE = πg.CheckLocal(πF, µmax, "max"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmin, "min"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µmax, µmin); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label96
					}
					goto Label97
					// line 534: if max < min:
					πF.SetLineno(534)
				Label96:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("bad repeat interval").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 535: raise error("bad repeat interval")
					πF.SetLineno(535)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label97
				Label97:
					goto Label93
				Label93:
					goto Label74
				Label73:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 537: raise error, "not supported"
					πF.SetLineno(537)
					πE = πF.Raise(πTemp002, πg.NewStr("not supported").ToObject(), nil)
					continue
					goto Label74
				Label74:
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µsubpattern); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label98
					}
					goto Label99
					// line 539: if subpattern:
					πF.SetLineno(539)
				Label98:
					// line 540: item = subpattern[-1:]
					πF.SetLineno(540)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µsubpattern, πTemp002); πE != nil {
						continue
					}
					µitem = πTemp003
					goto Label100
				Label99:
					// line 542: item = None
					πF.SetLineno(542)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µitem = πTemp002
					goto Label100
				Label100:
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µitem); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label101
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp008, πE = µ_len.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.Eq(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp006
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label102
					}
					πTemp008 = πg.NewInt(0).ToObject()
					πTemp011 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp013, πE = πg.GetItem(πF, µitem, πTemp011); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, πTemp013, πTemp008); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßAT); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, πTemp010, πTemp008); πE != nil {
						continue
					}
					πTemp003 = πTemp006
				Label102:
					πTemp002 = πTemp003
				Label101:
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label103
					}
					goto Label104
					// line 543: if not item or (_len(item) == 1 and item[0][0] == AT):
					πF.SetLineno(543)
				Label103:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 544: raise error, "nothing to repeat"
					πF.SetLineno(544)
					πE = πF.Raise(πTemp002, πg.NewStr("nothing to repeat").ToObject(), nil)
					continue
					goto Label104
				Label104:
					πTemp003 = πg.NewInt(0).ToObject()
					πTemp008 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µitem, πTemp008); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp010, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µREPEATCODES, "REPEATCODES"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µREPEATCODES, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label105
					}
					goto Label106
					// line 545: if item[0][0] in REPEATCODES:
					πF.SetLineno(545)
				Label105:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 546: raise error, "multiple repeat"
					πF.SetLineno(546)
					πE = πF.Raise(πTemp002, πg.NewStr("multiple repeat").ToObject(), nil)
					continue
					goto Label106
				Label106:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("?").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label107
					}
					goto Label108
					// line 547: if sourcematch("?"):
					πF.SetLineno(547)
				Label107:
					// line 548: subpattern[-1] = (MIN_REPEAT, (min, max, item))
					πF.SetLineno(548)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMIN_REPEAT); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmin, "min"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmax, "max"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple3(µmin, µmax, µitem).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πTemp008
					if πE = πg.SetItem(πF, µsubpattern, πTemp006, πTemp003); πE != nil {
						continue
					}
					goto Label109
				Label108:
					// line 550: subpattern[-1] = (MAX_REPEAT, (min, max, item))
					πF.SetLineno(550)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMAX_REPEAT); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmin, "min"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmax, "max"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple3(µmin, µmax, µitem).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006 = πTemp008
					if πE = πg.SetItem(πF, µsubpattern, πTemp006, πTemp003); πE != nil {
						continue
					}
					goto Label109
				Label109:
					goto Label31
					// line 552: elif this == ".":
					πF.SetLineno(552)
				Label24:
					// line 553: subpatternappend((ANY, None))
					πF.SetLineno(553)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßANY); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label31
					// line 555: elif this == "(":
					πF.SetLineno(555)
				Label25:
					// line 556: group = 1
					πF.SetLineno(556)
					µgroup = πg.NewInt(1).ToObject()
					// line 557: name = None
					πF.SetLineno(557)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µname = πTemp002
					// line 558: condgroup = None
					πF.SetLineno(558)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µcondgroup = πTemp002
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("?").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label110
					}
					goto Label111
					// line 559: if sourcematch("?"):
					πF.SetLineno(559)
				Label110:
					// line 560: group = 0
					πF.SetLineno(560)
					µgroup = πg.NewInt(0).ToObject()
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ßP.ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label112
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label113
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("#").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label114
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µASSERTCHARS, "ASSERTCHARS"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µASSERTCHARS, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label115
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("(").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label116
					}
					goto Label117
					// line 562: if sourcematch("P"):
					πF.SetLineno(562)
				Label112:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("<").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label119
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("=").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label120
					}
					goto Label121
					// line 564: if sourcematch("<"):
					πF.SetLineno(564)
				Label119:
					// line 566: name = ""
					πF.SetLineno(566)
					µname = ß.ToObject()
					// line 567: while 1:
					πF.SetLineno(567)
					πF.PushCheckpoint(124)
					πTemp005 = false
				Label123:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label125
					}
					if πTemp007, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(123)            
					// line 568: char = sourceget()
					πF.SetLineno(568)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchar = πTemp002
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µchar == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label126
					}
					goto Label127
					// line 569: if char is None:
					πF.SetLineno(569)
				Label126:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 570: raise error, "unterminated name"
					πF.SetLineno(570)
					πE = πF.Raise(πTemp002, πg.NewStr("unterminated name").ToObject(), nil)
					continue
					goto Label127
				Label127:
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µchar, πg.NewStr(">").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label128
					}
					goto Label129
					// line 571: if char == ">":
					πF.SetLineno(571)
				Label128:
					// line 572: break
					πF.SetLineno(572)
					πTemp005 = true
					continue
					goto Label129
				Label129:
					// line 573: name = name + char
					πF.SetLineno(573)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µname, µchar); πE != nil {
						continue
					}
					µname = πTemp002
					continue
				Label124:
					if πE != nil || πR != nil {
						continue
					}
				Label125:
					// line 574: group = 1
					πF.SetLineno(574)
					µgroup = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µname); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label130
					}
					goto Label131
					// line 575: if not name:
					πF.SetLineno(575)
				Label130:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("missing group name").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 576: raise error("missing group name")
					πF.SetLineno(576)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label131
				Label131:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisname); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label132
					}
					goto Label133
					// line 577: if not isname(name):
					πF.SetLineno(577)
				Label132:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("bad character in group name %r").ToObject(), µname); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 578: raise error("bad character in group name %r" %
					πF.SetLineno(578)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label133
				Label133:
					goto Label122
					// line 580: elif sourcematch("="):
					πF.SetLineno(580)
				Label120:
					// line 582: name = ""
					πF.SetLineno(582)
					µname = ß.ToObject()
					// line 583: while 1:
					πF.SetLineno(583)
					πF.PushCheckpoint(135)
					πTemp005 = false
				Label134:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label136
					}
					if πTemp007, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(134)            
					// line 584: char = sourceget()
					πF.SetLineno(584)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchar = πTemp002
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µchar == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label137
					}
					goto Label138
					// line 585: if char is None:
					πF.SetLineno(585)
				Label137:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 586: raise error, "unterminated name"
					πF.SetLineno(586)
					πE = πF.Raise(πTemp002, πg.NewStr("unterminated name").ToObject(), nil)
					continue
					goto Label138
				Label138:
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µchar, πg.NewStr(")").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label139
					}
					goto Label140
					// line 587: if char == ")":
					πF.SetLineno(587)
				Label139:
					// line 588: break
					πF.SetLineno(588)
					πTemp005 = true
					continue
					goto Label140
				Label140:
					// line 589: name = name + char
					πF.SetLineno(589)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µname, µchar); πE != nil {
						continue
					}
					µname = πTemp002
					continue
				Label135:
					if πE != nil || πR != nil {
						continue
					}
				Label136:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µname); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label141
					}
					goto Label142
					// line 590: if not name:
					πF.SetLineno(590)
				Label141:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("missing group name").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 591: raise error("missing group name")
					πF.SetLineno(591)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label142
				Label142:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πTemp003, πE = πg.ResolveGlobal(πF, ßisname); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label143
					}
					goto Label144
					// line 592: if not isname(name):
					πF.SetLineno(592)
				Label143:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("bad character in backref group name %r").ToObject(), µname); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 593: raise error("bad character in backref group name "
					πF.SetLineno(593)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label144
				Label144:
					// line 595: gid = state.groupdict.get(name)
					πF.SetLineno(595)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µstate, ßgroupdict, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgid = πTemp002
					if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µgid == πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label145
					}
					goto Label146
					// line 596: if gid is None:
					πF.SetLineno(596)
				Label145:
					// line 598: msg = "unknown group name: %s" % (name)
					πF.SetLineno(598)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("unknown group name: %s").ToObject(), µname); πE != nil {
						continue
					}
					µmsg = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[0] = µmsg
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 599: raise error(msg)
					πF.SetLineno(599)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label146
				Label146:
					// line 605: subpatternappend((GROUPREF, gid))
					πF.SetLineno(605)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßGROUPREF); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgid, "gid"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, µgid).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 606: continue
					πF.SetLineno(606)
					continue
					goto Label122
				Label121:
					// line 608: char = sourceget()
					πF.SetLineno(608)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchar = πTemp002
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µchar == πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label147
					}
					goto Label148
					// line 609: if char is None:
					πF.SetLineno(609)
				Label147:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 610: raise error, "unexpected end of pattern"
					πF.SetLineno(610)
					πE = πF.Raise(πTemp002, πg.NewStr("unexpected end of pattern").ToObject(), nil)
					continue
					goto Label148
				Label148:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("unknown specifier: ?P%s").ToObject(), µchar); πE != nil {
						continue
					}
					// line 611: raise error, "unknown specifier: ?P%s" % char
					πF.SetLineno(611)
					πE = πF.Raise(πTemp002, πTemp003, nil)
					continue
					goto Label122
				Label122:
					goto Label118
					// line 612: elif sourcematch(":"):
					πF.SetLineno(612)
				Label113:
					// line 614: group = 2
					πF.SetLineno(614)
					µgroup = πg.NewInt(2).ToObject()
					goto Label118
					// line 615: elif sourcematch("#"):
					πF.SetLineno(615)
				Label114:
					// line 617: while 1:
					πF.SetLineno(617)
					πF.PushCheckpoint(150)
					πTemp005 = false
				Label149:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label151
					}
					if πTemp007, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(149)            
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006 == πTemp008).ToObject()
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label152
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewStr(")").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label152:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label153
					}
					goto Label154
					// line 618: if source.next is None or source.next == ")":
					πF.SetLineno(618)
				Label153:
					// line 619: break
					πF.SetLineno(619)
					πTemp005 = true
					continue
					goto Label154
				Label154:
					// line 620: sourceget()
					πF.SetLineno(620)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					continue
				Label150:
					if πE != nil || πR != nil {
						continue
					}
				Label151:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(")").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp003, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label155
					}
					goto Label156
					// line 621: if not sourcematch(")"):
					πF.SetLineno(621)
				Label155:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 622: raise error, "unbalanced parenthesis"
					πF.SetLineno(622)
					πE = πF.Raise(πTemp002, πg.NewStr("unbalanced parenthesis").ToObject(), nil)
					continue
					goto Label156
				Label156:
					// line 623: continue
					πF.SetLineno(623)
					continue
					goto Label118
					// line 624: elif source.next in ASSERTCHARS:
					πF.SetLineno(624)
				Label115:
					// line 626: char = sourceget()
					πF.SetLineno(626)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchar = πTemp002
					// line 627: dir = 1
					πF.SetLineno(627)
					µdir = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µchar, πg.NewStr("<").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label157
					}
					goto Label158
					// line 628: if char == "<":
					πF.SetLineno(628)
				Label157:
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µLOOKBEHINDASSERTCHARS, "LOOKBEHINDASSERTCHARS"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µLOOKBEHINDASSERTCHARS, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label159
					}
					goto Label160
					// line 629: if source.next not in LOOKBEHINDASSERTCHARS:
					πF.SetLineno(629)
				Label159:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 630: raise error, "syntax error"
					πF.SetLineno(630)
					πE = πF.Raise(πTemp002, πg.NewStr("syntax error").ToObject(), nil)
					continue
					goto Label160
				Label160:
					// line 631: dir = -1 # lookbehind
					πF.SetLineno(631)
					if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µdir = πTemp002
					// line 632: char = sourceget()
					πF.SetLineno(632)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchar = πTemp002
					// line 633: state.lookbehind += 1
					πF.SetLineno(633)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µstate, ßlookbehind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µstate, ßlookbehind, πTemp003); πE != nil {
						continue
					}
					goto Label158
				Label158:
					// line 634: p = _parse_sub(source, state)
					πF.SetLineno(634)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[1] = µstate
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse_sub); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp003
					if πE = πg.CheckLocal(πF, µdir, "dir"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LT(πF, µdir, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label161
					}
					goto Label162
					// line 635: if dir < 0:
					πF.SetLineno(635)
				Label161:
					// line 636: state.lookbehind -= 1
					πF.SetLineno(636)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µstate, ßlookbehind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ISub(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µstate, ßlookbehind, πTemp003); πE != nil {
						continue
					}
					goto Label162
				Label162:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(")").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp003, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label163
					}
					goto Label164
					// line 637: if not sourcematch(")"):
					πF.SetLineno(637)
				Label163:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 638: raise error, "unbalanced parenthesis"
					πF.SetLineno(638)
					πE = πF.Raise(πTemp002, πg.NewStr("unbalanced parenthesis").ToObject(), nil)
					continue
					goto Label164
				Label164:
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µchar, πg.NewStr("=").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label165
					}
					goto Label166
					// line 639: if char == "=":
					πF.SetLineno(639)
				Label165:
					// line 640: subpatternappend((ASSERT, (dir, p)))
					πF.SetLineno(640)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßASSERT); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdir, "dir"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(µdir, µp).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label167
				Label166:
					// line 642: subpatternappend((ASSERT_NOT, (dir, p)))
					πF.SetLineno(642)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßASSERT_NOT); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdir, "dir"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(µdir, µp).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label167
				Label167:
					// line 643: continue
					πF.SetLineno(643)
					continue
					goto Label118
					// line 644: elif sourcematch("("):
					πF.SetLineno(644)
				Label116:
					// line 646: condname = ""
					πF.SetLineno(646)
					µcondname = ß.ToObject()
					// line 647: while 1:
					πF.SetLineno(647)
					πF.PushCheckpoint(169)
					πTemp005 = false
				Label168:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label170
					}
					if πTemp007, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(168)            
					// line 648: char = sourceget()
					πF.SetLineno(648)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchar = πTemp002
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µchar == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label171
					}
					goto Label172
					// line 649: if char is None:
					πF.SetLineno(649)
				Label171:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 650: raise error, "unterminated name"
					πF.SetLineno(650)
					πE = πF.Raise(πTemp002, πg.NewStr("unterminated name").ToObject(), nil)
					continue
					goto Label172
				Label172:
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µchar, πg.NewStr(")").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label173
					}
					goto Label174
					// line 651: if char == ")":
					πF.SetLineno(651)
				Label173:
					// line 652: break
					πF.SetLineno(652)
					πTemp005 = true
					continue
					goto Label174
				Label174:
					// line 653: condname = condname + char
					πF.SetLineno(653)
					if πE = πg.CheckLocal(πF, µcondname, "condname"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µcondname, µchar); πE != nil {
						continue
					}
					µcondname = πTemp002
					continue
				Label169:
					if πE != nil || πR != nil {
						continue
					}
				Label170:
					// line 654: group = 2
					πF.SetLineno(654)
					µgroup = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µcondname, "condname"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µcondname); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label175
					}
					goto Label176
					// line 655: if not condname:
					πF.SetLineno(655)
				Label175:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("missing group name").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 656: raise error("missing group name")
					πF.SetLineno(656)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label176
				Label176:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcondname, "condname"); πE != nil {
						continue
					}
					πTemp001[0] = µcondname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisname); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label177
					}
					goto Label178
					// line 657: if isname(condname):
					πF.SetLineno(657)
				Label177:
					// line 658: condgroup = state.groupdict.get(condname)
					πF.SetLineno(658)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcondname, "condname"); πE != nil {
						continue
					}
					πTemp001[0] = µcondname
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µstate, ßgroupdict, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcondgroup = πTemp002
					if πE = πg.CheckLocal(πF, µcondgroup, "condgroup"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µcondgroup == πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label180
					}
					goto Label181
					// line 659: if condgroup is None:
					πF.SetLineno(659)
				Label180:
					// line 661: msg = "unknown group name: %s" % (condname)
					πF.SetLineno(661)
					if πE = πg.CheckLocal(πF, µcondname, "condname"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mod(πF, πg.NewStr("unknown group name: %s").ToObject(), µcondname); πE != nil {
						continue
					}
					µmsg = πTemp002
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[0] = µmsg
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 662: raise error(msg)
					πF.SetLineno(662)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label181
				Label181:
					goto Label179
				Label178:
					// line 664: try:
					πF.SetLineno(664)
					πF.PushCheckpoint(183)
					// line 665: condgroup = int(condname)
					πF.SetLineno(665)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcondname, "condname"); πE != nil {
						continue
					}
					πTemp001[0] = µcondname
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcondgroup = πTemp003
					πF.PopCheckpoint()
					goto Label182
				Label183:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp014, πTemp015 = πF.ExcInfo()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp014.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label184
					}
					πE = πF.Raise(πTemp014.ToObject(), nil, πTemp015.ToObject())
					continue
					// line 666: except ValueError:
					πF.SetLineno(666)
				Label184:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 667: raise error, "bad character in group name"
					πF.SetLineno(667)
					πE = πF.Raise(πTemp002, πg.NewStr("bad character in group name").ToObject(), nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label182
				Label182:
					goto Label179
				Label179:
					goto Label118
				Label117:
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveGlobal(πF, ßFLAGS); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, πTemp008, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label185
					}
					goto Label186
					// line 675: if not source.next in FLAGS:
					πF.SetLineno(675)
				Label185:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 676: raise error, "unexpected end of pattern"
					πF.SetLineno(676)
					πE = πF.Raise(πTemp002, πg.NewStr("unexpected end of pattern").ToObject(), nil)
					continue
					goto Label186
				Label186:
					// line 677: while source.next in FLAGS:
					πF.SetLineno(677)
					πF.PushCheckpoint(188)
					πTemp005 = false
				Label187:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label189
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µsource, ßnext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßFLAGS); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Contains(πF, πTemp006, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp012).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(187)            
					// line 678: state.flags = state.flags | FLAGS[sourceget()]
					πF.SetLineno(678)
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µstate, ßflags, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp008, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006 = πTemp008
					if πTemp010, πE = πg.ResolveGlobal(πF, ßFLAGS); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp006); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, πTemp003, πTemp008); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µstate, ßflags, πTemp003); πE != nil {
						continue
					}
					continue
				Label188:
					if πE != nil || πR != nil {
						continue
					}
				Label189:
					goto Label118
				Label118:
					goto Label111
				Label111:
					if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µgroup); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label190
					}
					goto Label191
					// line 679: if group:
					πF.SetLineno(679)
				Label190:
					if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µgroup, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label193
					}
					goto Label194
					// line 681: if group == 2:
					πF.SetLineno(681)
				Label193:
					// line 683: group = None
					πF.SetLineno(683)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µgroup = πTemp002
					goto Label195
				Label194:
					// line 685: group = state.opengroup(name)
					πF.SetLineno(685)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µstate, ßopengroup, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µgroup = πTemp003
					goto Label195
				Label195:
					if πE = πg.CheckLocal(πF, µcondgroup, "condgroup"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µcondgroup); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label196
					}
					goto Label197
					// line 686: if condgroup:
					πF.SetLineno(686)
				Label196:
					// line 687: p = _parse_sub_cond(source, state, condgroup)
					πF.SetLineno(687)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[1] = µstate
					if πE = πg.CheckLocal(πF, µcondgroup, "condgroup"); πE != nil {
						continue
					}
					πTemp001[2] = µcondgroup
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse_sub_cond); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp003
					goto Label198
				Label197:
					// line 689: p = _parse_sub(source, state)
					πF.SetLineno(689)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[1] = µstate
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse_sub); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp003
					goto Label198
				Label198:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(")").ToObject()
					if πE = πg.CheckLocal(πF, µsourcematch, "sourcematch"); πE != nil {
						continue
					}
					if πTemp003, πE = µsourcematch.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label199
					}
					goto Label200
					// line 690: if not sourcematch(")"):
					πF.SetLineno(690)
				Label199:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 691: raise error, "unbalanced parenthesis"
					πF.SetLineno(691)
					πE = πF.Raise(πTemp002, πg.NewStr("unbalanced parenthesis").ToObject(), nil)
					continue
					goto Label200
				Label200:
					if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µgroup != πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label201
					}
					goto Label202
					// line 692: if group is not None:
					πF.SetLineno(692)
				Label201:
					// line 693: state.closegroup(group)
					πF.SetLineno(693)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
						continue
					}
					πTemp001[0] = µgroup
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µstate, ßclosegroup, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label202
				Label202:
					// line 694: subpatternappend((SUBPATTERN, (group, p)))
					πF.SetLineno(694)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSUBPATTERN); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(µgroup, µp).ToObject()
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label192
				Label191:
					// line 696: while 1:
					πF.SetLineno(696)
					πF.PushCheckpoint(204)
					πTemp005 = false
				Label203:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label205
					}
					if πTemp007, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(203)            
					// line 697: char = sourceget()
					πF.SetLineno(697)
					if πE = πg.CheckLocal(πF, µsourceget, "sourceget"); πE != nil {
						continue
					}
					if πTemp002, πE = µsourceget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchar = πTemp002
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µchar == πTemp003).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label206
					}
					goto Label207
					// line 698: if char is None:
					πF.SetLineno(698)
				Label206:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 699: raise error, "unexpected end of pattern"
					πF.SetLineno(699)
					πE = πF.Raise(πTemp002, πg.NewStr("unexpected end of pattern").ToObject(), nil)
					continue
					goto Label207
				Label207:
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µchar, πg.NewStr(")").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label208
					}
					goto Label209
					// line 700: if char == ")":
					πF.SetLineno(700)
				Label208:
					// line 701: break
					πF.SetLineno(701)
					πTemp005 = true
					continue
					goto Label209
				Label209:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 702: raise error, "unknown extension"
					πF.SetLineno(702)
					πE = πF.Raise(πTemp002, πg.NewStr("unknown extension").ToObject(), nil)
					continue
					continue
				Label204:
					if πE != nil || πR != nil {
						continue
					}
				Label205:
					goto Label192
				Label192:
					goto Label31
					// line 704: elif this == "^":
					πF.SetLineno(704)
				Label26:
					// line 705: subpatternappend((AT, AT_BEGINNING))
					πF.SetLineno(705)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAT); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßAT_BEGINNING); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label31
					// line 707: elif this == "$":
					πF.SetLineno(707)
				Label27:
					// line 708: subpattern.append((AT, AT_END))
					πF.SetLineno(708)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßAT); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßAT_END); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp003, πTemp006).ToObject()
					πTemp001[0] = πTemp002
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsubpattern, ßappend, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label31
					// line 710: elif this and this[0] == "\\":
					πF.SetLineno(710)
				Label29:
					// line 711: code = _escape(source, this, state)
					πF.SetLineno(711)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp001[1] = µthis
					if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
						continue
					}
					πTemp001[2] = µstate
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_escape); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp003
					// line 712: subpatternappend(code)
					πF.SetLineno(712)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[0] = µcode
					if πE = πg.CheckLocal(πF, µsubpatternappend, "subpatternappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µsubpatternappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label31
				Label30:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 715: raise error, "parser error"
					πF.SetLineno(715)
					πE = πF.Raise(πTemp002, πg.NewStr("parser error").ToObject(), nil)
					continue
					goto Label31
				Label31:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 717: return subpattern
					πF.SetLineno(717)
					if πE = πg.CheckLocal(πF, µsubpattern, "subpattern"); πE != nil {
						continue
					}
					πR = µsubpattern
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_parse.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 719: def parse(str, flags=0, pattern=None):
			πF.SetLineno(719)
			πTemp011 = make([]πg.Param, 3)
			πTemp011[0] = πg.Param{Name: "str", Def: nil}
			πTemp011[1] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			if πTemp016, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp011[2] = πg.Param{Name: "pattern", Def: πTemp016}
			πTemp014 = πg.NewFunction(πg.NewCode("parse", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstr *πg.Object = πArgs[0]; _ = µstr
				var µflags *πg.Object = πArgs[1]; _ = µflags
				var µpattern *πg.Object = πArgs[2]; _ = µpattern
				var µsource *πg.Object = πg.UnboundLocal; _ = µsource
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µtail *πg.Object = πg.UnboundLocal; _ = µtail
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
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 722: source = Tokenizer(str)
					πF.SetLineno(722)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
						continue
					}
					πTemp001[0] = µstr
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTokenizer); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µsource = πTemp003
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µpattern == πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 724: if pattern is None:
					πF.SetLineno(724)
				Label1:
					// line 725: pattern = Pattern()
					πF.SetLineno(725)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßPattern); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µpattern = πTemp003
					goto Label2
				Label2:
					// line 726: pattern.flags = flags
					πF.SetLineno(726)
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µflags); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpattern, ßflags, πTemp002); πE != nil {
						continue
					}
					// line 727: pattern.str = str
					πF.SetLineno(727)
					if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µstr); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µpattern, ßstr, πTemp002); πE != nil {
						continue
					}
					// line 729: p = _parse_sub(source, pattern, 0)
					πF.SetLineno(729)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[1] = µpattern
					πTemp001[2] = πg.NewInt(0).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse_sub); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp003
					// line 731: tail = source.get()
					πF.SetLineno(731)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µsource, ßget, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µtail = πTemp003
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µtail, πg.NewStr(")").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µtail); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 732: if tail == ")":
					πF.SetLineno(732)
				Label3:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 733: raise error, "unbalanced parenthesis"
					πF.SetLineno(733)
					πE = πF.Raise(πTemp002, πg.NewStr("unbalanced parenthesis").ToObject(), nil)
					continue
					goto Label5
					// line 734: elif tail:
					πF.SetLineno(734)
				Label4:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 735: raise error, "bogus characters at end of regular expression"
					πF.SetLineno(735)
					πE = πF.Raise(πTemp002, πg.NewStr("bogus characters at end of regular expression").ToObject(), nil)
					continue
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_VERBOSE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.And(πF, µflags, πTemp006); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µp, ßpattern, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßflags, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_VERBOSE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, πTemp006, πTemp005); πE != nil {
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
					// line 737: if not (flags & SRE_FLAG_VERBOSE) and p.pattern.flags & SRE_FLAG_VERBOSE:
					πF.SetLineno(737)
				Label7:
					// line 740: return parse(str, p.pattern.flags)
					πF.SetLineno(740)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
						continue
					}
					πTemp001[0] = µstr
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßpattern, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßflags, nil); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßparse); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πR = πTemp003
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_DEBUG); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µflags, πTemp003); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label9
					}
					goto Label10
					// line 742: if flags & SRE_FLAG_DEBUG:
					πF.SetLineno(742)
				Label9:
					// line 743: p.dump()
					πF.SetLineno(743)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßdump, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label10
				Label10:
					// line 745: return p
					πF.SetLineno(745)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πR = µp
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßparse.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 747: def parse_template(source, pattern):
			πF.SetLineno(747)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "source", Def: nil}
			πTemp011[1] = πg.Param{Name: "pattern", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("parse_template", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µsource *πg.Object = πArgs[0]; _ = µsource
				var µpattern *πg.Object = πArgs[1]; _ = µpattern
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µsget *πg.Object = πg.UnboundLocal; _ = µsget
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µliteral *πg.Object = πg.UnboundLocal; _ = µliteral
				var µsep *πg.Object = πg.UnboundLocal; _ = µsep
				var µmakechar *πg.Object = πg.UnboundLocal; _ = µmakechar
				var µthis *πg.Object = πg.UnboundLocal; _ = µthis
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µchar *πg.Object = πg.UnboundLocal; _ = µchar
				var µindex *πg.Object = πg.UnboundLocal; _ = µindex
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var µisoctal *πg.Object = πg.UnboundLocal; _ = µisoctal
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µgroups *πg.Object = πg.UnboundLocal; _ = µgroups
				var µgroupsappend *πg.Object = πg.UnboundLocal; _ = µgroupsappend
				var µliterals *πg.Object = πg.UnboundLocal; _ = µliterals
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 *πg.BaseException
				_ = πTemp011
				var πTemp012 *πg.Traceback
				_ = πTemp012
				var πTemp013 *πg.BaseException
				_ = πTemp013
				var πTemp014 []*πg.Object
				_ = πTemp014
				var πTemp015 []*πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 4: goto Label4
					case 37: goto Label37
					case 51: goto Label51
					case 20: goto Label20
					case 21: goto Label21
					case 54: goto Label54
					case 53: goto Label53
					case 30: goto Label30
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 750: s = Tokenizer(source)
					πF.SetLineno(750)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					πTemp001[0] = µsource
					if πTemp002, πE = πg.ResolveGlobal(πF, ßTokenizer); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µs = πTemp003
					// line 751: sget = s.get
					πF.SetLineno(751)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßget, nil); πE != nil {
						continue
					}
					µsget = πTemp002
					// line 752: p = []
					πF.SetLineno(752)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µp = πTemp002
					// line 753: a = p.append
					πF.SetLineno(753)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßappend, nil); πE != nil {
						continue
					}
					µa = πTemp002
					// line 754: def literal(literal, p=p, pappend=a):
					πF.SetLineno(754)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "literal", Def: nil}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp004[1] = πg.Param{Name: "p", Def: µp}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					πTemp004[2] = πg.Param{Name: "pappend", Def: µa}
					πTemp002 = πg.NewFunction(πg.NewCode("literal", "/usr/lib/python2.7/sre_parse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µliteral *πg.Object = πArgs[0]; _ = µliteral
						var µp *πg.Object = πArgs[1]; _ = µp
						var µpappend *πg.Object = πArgs[2]; _ = µpappend
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
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							πTemp001 = µp
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µp, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005 == πTemp004).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 755: if p and p[-1][0] is LITERAL:
							πF.SetLineno(755)
						Label2:
							// line 756: p[-1] = LITERAL, p[-1][1] + literal
							πF.SetLineno(756)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(1).ToObject()
							if πTemp008, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µp, πTemp007); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µliteral, "literal"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp006, µliteral); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, µp, πTemp004, πTemp003); πE != nil {
								continue
							}
							goto Label4
						Label3:
							// line 758: pappend((LITERAL, literal))
							πF.SetLineno(758)
							πTemp009 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µliteral, "literal"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp003, µliteral).ToObject()
							πTemp009[0] = πTemp001
							if πE = πg.CheckLocal(πF, µpappend, "pappend"); πE != nil {
								continue
							}
							if πTemp001, πE = µpappend.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
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
					µliteral = πTemp002
					// line 759: sep = source[:0]
					πF.SetLineno(759)
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(0).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsource, "source"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsource, πTemp003); πE != nil {
						continue
					}
					µsep = πTemp005
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
						continue
					}
					πTemp001[0] = µsep
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = ß.ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp003 = πg.GetBool(πTemp006 == πTemp007).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label1
					}
					goto Label2
					// line 760: if type(sep) is type(""):
					πF.SetLineno(760)
				Label1:
					// line 761: makechar = chr
					πF.SetLineno(761)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					µmakechar = πTemp003
					goto Label3
				Label2:
					// line 763: makechar = unichr
					πF.SetLineno(763)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßunichr); πE != nil {
						continue
					}
					µmakechar = πTemp003
					goto Label3
				Label3:
					// line 764: while 1:
					πF.SetLineno(764)
					πF.PushCheckpoint(5)
					πTemp008 = false
				Label4:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label6
					}
					if πTemp009, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(4)            
					// line 765: this = sget()
					πF.SetLineno(765)
					if πE = πg.CheckLocal(πF, µsget, "sget"); πE != nil {
						continue
					}
					if πTemp003, πE = µsget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µthis = πTemp003
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µthis == πTemp005).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label7
					}
					goto Label8
					// line 766: if this is None:
					πF.SetLineno(766)
				Label7:
					// line 767: break # end of replacement string
					πF.SetLineno(767)
					πTemp008 = true
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp003 = µthis
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label9
					}
					πTemp006 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µthis, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp007, πg.NewStr("\\").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
				Label9:
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label10
					}
					goto Label11
					// line 768: if this and this[0] == "\\":
					πF.SetLineno(768)
				Label10:
					// line 770: c = this[1:2]
					πF.SetLineno(770)
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µthis, πTemp003); πE != nil {
						continue
					}
					µc = πTemp005
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µc, ßg.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µc, ß0.ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßDIGITS); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp005, µc); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label15
					}
					goto Label16
					// line 771: if c == "g":
					πF.SetLineno(771)
				Label13:
					// line 772: name = ""
					πF.SetLineno(772)
					µname = ß.ToObject()
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("<").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µs, ßmatch, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label18
					}
					goto Label19
					// line 773: if s.match("<"):
					πF.SetLineno(773)
				Label18:
					// line 774: while 1:
					πF.SetLineno(774)
					πF.PushCheckpoint(21)
					πTemp009 = false
				Label20:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp009 {
						πF.PopCheckpoint()
						goto Label22
					}
					if πTemp010, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(20)            
					// line 775: char = sget()
					πF.SetLineno(775)
					if πE = πg.CheckLocal(πF, µsget, "sget"); πE != nil {
						continue
					}
					if πTemp003, πE = µsget.Call(πF, nil, nil); πE != nil {
						continue
					}
					µchar = πTemp003
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µchar == πTemp005).ToObject()
					if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label23
					}
					goto Label24
					// line 776: if char is None:
					πF.SetLineno(776)
				Label23:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 777: raise error, "unterminated group name"
					πF.SetLineno(777)
					πE = πF.Raise(πTemp003, πg.NewStr("unterminated group name").ToObject(), nil)
					continue
					goto Label24
				Label24:
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µchar, πg.NewStr(">").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp010 {
						goto Label25
					}
					goto Label26
					// line 778: if char == ">":
					πF.SetLineno(778)
				Label25:
					// line 779: break
					πF.SetLineno(779)
					πTemp009 = true
					continue
					goto Label26
				Label26:
					// line 780: name = name + char
					πF.SetLineno(780)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µname, µchar); πE != nil {
						continue
					}
					µname = πTemp003
					continue
				Label21:
					if πE != nil || πR != nil {
						continue
					}
				Label22:
					goto Label19
				Label19:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µname); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label27
					}
					goto Label28
					// line 781: if not name:
					πF.SetLineno(781)
				Label27:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 782: raise error, "missing group name"
					πF.SetLineno(782)
					πE = πF.Raise(πTemp003, πg.NewStr("missing group name").ToObject(), nil)
					continue
					goto Label28
				Label28:
					// line 783: try:
					πF.SetLineno(783)
					πF.PushCheckpoint(30)
					// line 784: index = int(name)
					πF.SetLineno(784)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µindex = πTemp005
					if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µindex, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label31
					}
					goto Label32
					// line 785: if index < 0:
					πF.SetLineno(785)
				Label31:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 786: raise error, "negative group number"
					πF.SetLineno(786)
					πE = πF.Raise(πTemp003, πg.NewStr("negative group number").ToObject(), nil)
					continue
					goto Label32
				Label32:
					πF.PopCheckpoint()
					goto Label29
				Label30:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label33
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 787: except ValueError:
					πF.SetLineno(787)
				Label33:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001[0] = µname
					if πTemp005, πE = πg.ResolveGlobal(πF, ßisname); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label34
					}
					goto Label35
					// line 788: if not isname(name):
					πF.SetLineno(788)
				Label34:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 789: raise error, "bad character in group name"
					πF.SetLineno(789)
					πE = πF.Raise(πTemp003, πg.NewStr("bad character in group name").ToObject(), nil)
					continue
					goto Label35
				Label35:
					// line 790: try:
					πF.SetLineno(790)
					πF.PushCheckpoint(37)
					// line 791: index = pattern.groupindex[name]
					πF.SetLineno(791)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp003 = µname
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µpattern, ßgroupindex, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
						continue
					}
					µindex = πTemp005
					πF.PopCheckpoint()
					goto Label36
				Label37:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp013, πTemp012 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label38
					}
					πE = πF.Raise(πTemp013.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 792: except KeyError:
					πF.SetLineno(792)
				Label38:
					// line 794: msg = "unknown group name: %s" % (name)
					πF.SetLineno(794)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mod(πF, πg.NewStr("unknown group name: %s").ToObject(), µname); πE != nil {
						continue
					}
					µmsg = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[0] = µmsg
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 795: raise IndexError(msg)
					πF.SetLineno(795)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label36
				Label36:
					πF.RestoreExc(nil, nil)
					goto Label29
				Label29:
					// line 796: a((MARK, index))
					πF.SetLineno(796)
					πTemp001 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp005, µindex).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp003, πE = µa.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label17
					// line 797: elif c == "0":
					πF.SetLineno(797)
				Label14:
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µs, ßnext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label39
					}
					goto Label40
					// line 798: if s.next in OCTDIGITS:
					πF.SetLineno(798)
				Label39:
					// line 799: this = this + sget()
					πF.SetLineno(799)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsget, "sget"); πE != nil {
						continue
					}
					if πTemp005, πE = µsget.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µthis, πTemp005); πE != nil {
						continue
					}
					µthis = πTemp003
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µs, ßnext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label41
					}
					goto Label42
					// line 800: if s.next in OCTDIGITS:
					πF.SetLineno(800)
				Label41:
					// line 801: this = this + sget()
					πF.SetLineno(801)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsget, "sget"); πE != nil {
						continue
					}
					if πTemp005, πE = µsget.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µthis, πTemp005); πE != nil {
						continue
					}
					µthis = πTemp003
					goto Label42
				Label42:
					goto Label40
				Label40:
					// line 802: literal(makechar(int(this[1:], 8) & 0xff))
					πF.SetLineno(802)
					πTemp001 = πF.MakeArgs(1)
					πTemp014 = πF.MakeArgs(1)
					πTemp015 = πF.MakeArgs(2)
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µthis, πTemp005); πE != nil {
						continue
					}
					πTemp015[0] = πTemp006
					πTemp015[1] = πg.NewInt(8).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp015, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp015)
					if πTemp003, πE = πg.And(πF, πTemp006, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp014[0] = πTemp003
					if πE = πg.CheckLocal(πF, µmakechar, "makechar"); πE != nil {
						continue
					}
					if πTemp003, πE = µmakechar.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µliteral, "literal"); πE != nil {
						continue
					}
					if πTemp003, πE = µliteral.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label17
					// line 803: elif c in DIGITS:
					πF.SetLineno(803)
				Label15:
					// line 804: isoctal = False
					πF.SetLineno(804)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
						continue
					}
					µisoctal = πTemp003
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µs, ßnext, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßDIGITS); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, πTemp006, πTemp005); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label43
					}
					goto Label44
					// line 805: if s.next in DIGITS:
					πF.SetLineno(805)
				Label43:
					// line 806: this = this + sget()
					πF.SetLineno(806)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsget, "sget"); πE != nil {
						continue
					}
					if πTemp005, πE = µsget.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µthis, πTemp005); πE != nil {
						continue
					}
					µthis = πTemp003
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp006, µc); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp010).ToObject()
					πTemp003 = πTemp005
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label45
					}
					πTemp006 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µthis, πTemp006); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp006, πTemp007); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp010).ToObject()
					πTemp003 = πTemp005
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp009 {
						goto Label45
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µs, ßnext, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßOCTDIGITS); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp007, πTemp006); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp010).ToObject()
					πTemp003 = πTemp005
				Label45:
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label46
					}
					goto Label47
					// line 807: if (c in OCTDIGITS and this[2] in OCTDIGITS and
					πF.SetLineno(807)
				Label46:
					// line 809: this = this + sget()
					πF.SetLineno(809)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsget, "sget"); πE != nil {
						continue
					}
					if πTemp005, πE = µsget.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µthis, πTemp005); πE != nil {
						continue
					}
					µthis = πTemp003
					// line 810: isoctal = True
					πF.SetLineno(810)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					µisoctal = πTemp003
					// line 811: literal(makechar(int(this[1:], 8) & 0xff))
					πF.SetLineno(811)
					πTemp001 = πF.MakeArgs(1)
					πTemp014 = πF.MakeArgs(1)
					πTemp015 = πF.MakeArgs(2)
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µthis, πTemp005); πE != nil {
						continue
					}
					πTemp015[0] = πTemp006
					πTemp015[1] = πg.NewInt(8).ToObject()
					if πTemp005, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp015, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp015)
					if πTemp003, πE = πg.And(πF, πTemp006, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					πTemp014[0] = πTemp003
					if πE = πg.CheckLocal(πF, µmakechar, "makechar"); πE != nil {
						continue
					}
					if πTemp003, πE = µmakechar.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µliteral, "literal"); πE != nil {
						continue
					}
					if πTemp003, πE = µliteral.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label47
				Label47:
					goto Label44
				Label44:
					if πE = πg.CheckLocal(πF, µisoctal, "isoctal"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µisoctal); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label48
					}
					goto Label49
					// line 812: if not isoctal:
					πF.SetLineno(812)
				Label48:
					// line 813: a((MARK, int(this[1:])))
					πF.SetLineno(813)
					πTemp001 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
						continue
					}
					πTemp014 = πF.MakeArgs(1)
					if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µthis, πTemp006); πE != nil {
						continue
					}
					πTemp014[0] = πTemp007
					if πTemp006, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					πTemp003 = πg.NewTuple2(πTemp005, πTemp007).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πTemp003, πE = µa.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label49
				Label49:
					goto Label17
				Label16:
					// line 815: try:
					πF.SetLineno(815)
					πF.PushCheckpoint(51)
					// line 816: this = makechar(ESCAPES[this][1])
					πF.SetLineno(816)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp006 = µthis
					if πTemp016, πE = πg.ResolveGlobal(πF, ßESCAPES); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp016, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µmakechar, "makechar"); πE != nil {
						continue
					}
					if πTemp003, πE = µmakechar.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µthis = πTemp003
					πF.PopCheckpoint()
					goto Label50
				Label51:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label52
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 817: except KeyError:
					πF.SetLineno(817)
				Label52:
					// line 818: pass
					πF.SetLineno(818)
					πF.RestoreExc(nil, nil)
					goto Label50
				Label50:
					// line 819: literal(this)
					πF.SetLineno(819)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp001[0] = µthis
					if πE = πg.CheckLocal(πF, µliteral, "literal"); πE != nil {
						continue
					}
					if πTemp003, πE = µliteral.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label17
				Label17:
					goto Label12
				Label11:
					// line 821: literal(this)
					πF.SetLineno(821)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
						continue
					}
					πTemp001[0] = µthis
					if πE = πg.CheckLocal(πF, µliteral, "literal"); πE != nil {
						continue
					}
					if πTemp003, πE = µliteral.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label12
				Label12:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
				Label6:
					// line 823: i = 0
					πF.SetLineno(823)
					µi = πg.NewInt(0).ToObject()
					// line 824: groups = []
					πF.SetLineno(824)
					πTemp001 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					µgroups = πTemp003
					// line 825: groupsappend = groups.append
					πF.SetLineno(825)
					if πE = πg.CheckLocal(πF, µgroups, "groups"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µgroups, ßappend, nil); πE != nil {
						continue
					}
					µgroupsappend = πTemp003
					// line 826: literals = [None] * len(p)
					πF.SetLineno(826)
					πTemp001 = make([]*πg.Object, 1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					πTemp005 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp001[0] = µp
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Mul(πF, πTemp005, πTemp007); πE != nil {
						continue
					}
					µliterals = πTemp003
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, µp); πE != nil {
						continue
					}
					πF.PushCheckpoint(54)
					πTemp008 = false
				Label53:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label55
					}
					if πTemp005, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp005); πE != nil {
							continue
						}
						µc = πTemp006
						µs = πTemp007
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(53)            
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(µc == πTemp006).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label56
					}
					goto Label57
					// line 828: if c is MARK:
					πF.SetLineno(828)
				Label56:
					// line 829: groupsappend((i, s))
					πF.SetLineno(829)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µi, µs).ToObject()
					πTemp001[0] = πTemp005
					if πE = πg.CheckLocal(πF, µgroupsappend, "groupsappend"); πE != nil {
						continue
					}
					if πTemp005, πE = µgroupsappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label58
				Label57:
					// line 832: literals[i] = s
					πF.SetLineno(832)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, µs); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µliterals, "literals"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp006 = µi
					if πE = πg.SetItem(πF, µliterals, πTemp006, πTemp005); πE != nil {
						continue
					}
					goto Label58
				Label58:
					// line 833: i = i + 1
					πF.SetLineno(833)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp005
					continue
				Label54:
					if πE != nil || πR != nil {
						continue
					}
				Label55:
					// line 834: return groups, literals
					πF.SetLineno(834)
					if πE = πg.CheckLocal(πF, µgroups, "groups"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µliterals, "literals"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µgroups, µliterals).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßparse_template.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 836: def expand_template(template, match):
			πF.SetLineno(836)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "template", Def: nil}
			πTemp011[1] = πg.Param{Name: "match", Def: nil}
			πTemp017 = πg.NewFunction(πg.NewCode("expand_template", "/usr/lib/python2.7/sre_parse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µtemplate *πg.Object = πArgs[0]; _ = µtemplate
				var µmatch *πg.Object = πArgs[1]; _ = µmatch
				var µg *πg.Object = πg.UnboundLocal; _ = µg
				var µsep *πg.Object = πg.UnboundLocal; _ = µsep
				var µgroups *πg.Object = πg.UnboundLocal; _ = µgroups
				var µliterals *πg.Object = πg.UnboundLocal; _ = µliterals
				var µindex *πg.Object = πg.UnboundLocal; _ = µindex
				var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
				var µs *πg.Object = πg.UnboundLocal; _ = µs
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
					// line 837: g = match.group
					πF.SetLineno(837)
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
						continue
					}
					µg = πTemp001
					// line 838: sep = match.string[:0]
					πF.SetLineno(838)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(0).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µmatch, ßstring, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					µsep = πTemp002
					// line 839: groups, literals = template
					πF.SetLineno(839)
					if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µtemplate); πE != nil {
						continue
					}
					µgroups = πTemp001
					µliterals = πTemp002
					// line 840: literals = literals[:]
					πF.SetLineno(840)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µliterals, "literals"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µliterals, πTemp001); πE != nil {
						continue
					}
					µliterals = πTemp002
					// line 841: try:
					πF.SetLineno(841)
					πF.PushCheckpoint(2)
					if πE = πg.CheckLocal(πF, µgroups, "groups"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µgroups); πE != nil {
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
						πTemp005 = !isStop
					} else {
						πTemp005 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
							continue
						}
						µindex = πTemp003
						µgroup = πTemp006
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 843: literals[index] = s = g(group)
					πF.SetLineno(843)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
						continue
					}
					πTemp007[0] = µgroup
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πTemp002, πE = µg.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µliterals, "literals"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindex, "index"); πE != nil {
						continue
					}
					πTemp006 = µindex
					if πE = πg.SetItem(πF, µliterals, πTemp006, πTemp003); πE != nil {
						continue
					}
					µs = πTemp002
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µs == πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 844: if s is None:
					πF.SetLineno(844)
				Label6:
					if πTemp002, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 845: raise error, "unmatched group"
					πF.SetLineno(845)
					πE = πF.Raise(πTemp002, πg.NewStr("unmatched group").ToObject(), nil)
					continue
					goto Label7
				Label7:
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
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 846: except IndexError:
					πF.SetLineno(846)
				Label8:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 847: raise error, "invalid group reference"
					πF.SetLineno(847)
					πE = πF.Raise(πTemp001, πg.NewStr("invalid group reference").ToObject(), nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 848: return sep.join(literals)
					πF.SetLineno(848)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µliterals, "literals"); πE != nil {
						continue
					}
					πTemp007[0] = µliterals
					if πE = πg.CheckLocal(πF, µsep, "sep"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsep, ßjoin, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
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
			if πE = πF.Globals().SetItem(πF, ßexpand_template.ToObject(), πTemp017); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("sre_parse", Code)
}

