package sre_compile
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/_sre"
	// _ "github.com/pygolin/stdlib/sre_constants"
	// _ "github.com/pygolin/stdlib/sre_parse"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/sre_compile.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß1 := πg.InternStr("1")
		ßANY := πg.InternStr("ANY")
		ßANY_ALL := πg.InternStr("ANY_ALL")
		ßASSERT := πg.InternStr("ASSERT")
		ßASSERT_NOT := πg.InternStr("ASSERT_NOT")
		ßAT := πg.InternStr("AT")
		ßATCODES := πg.InternStr("ATCODES")
		ßAT_LOCALE := πg.InternStr("AT_LOCALE")
		ßAT_MULTILINE := πg.InternStr("AT_MULTILINE")
		ßAT_UNICODE := πg.InternStr("AT_UNICODE")
		ßAssertionError := πg.InternStr("AssertionError")
		ßBIGCHARSET := πg.InternStr("BIGCHARSET")
		ßBRANCH := πg.InternStr("BRANCH")
		ßCALL := πg.InternStr("CALL")
		ßCATEGORY := πg.InternStr("CATEGORY")
		ßCHARSET := πg.InternStr("CHARSET")
		ßCHCODES := πg.InternStr("CHCODES")
		ßCH_LOCALE := πg.InternStr("CH_LOCALE")
		ßCH_UNICODE := πg.InternStr("CH_UNICODE")
		ßCODESIZE := πg.InternStr("CODESIZE")
		ßFAILURE := πg.InternStr("FAILURE")
		ßGROUPREF := πg.InternStr("GROUPREF")
		ßGROUPREF_EXISTS := πg.InternStr("GROUPREF_EXISTS")
		ßIN := πg.InternStr("IN")
		ßINFO := πg.InternStr("INFO")
		ßIN_IGNORE := πg.InternStr("IN_IGNORE")
		ßIndexError := πg.InternStr("IndexError")
		ßJUMP := πg.InternStr("JUMP")
		ßLITERAL := πg.InternStr("LITERAL")
		ßMAGIC := πg.InternStr("MAGIC")
		ßMARK := πg.InternStr("MARK")
		ßMAXCODE := πg.InternStr("MAXCODE")
		ßMAX_REPEAT := πg.InternStr("MAX_REPEAT")
		ßMAX_UNTIL := πg.InternStr("MAX_UNTIL")
		ßMIN_REPEAT := πg.InternStr("MIN_REPEAT")
		ßMIN_REPEAT_ONE := πg.InternStr("MIN_REPEAT_ONE")
		ßMIN_UNTIL := πg.InternStr("MIN_UNTIL")
		ßNEGATE := πg.InternStr("NEGATE")
		ßNOT_LITERAL := πg.InternStr("NOT_LITERAL")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßOPCODES := πg.InternStr("OPCODES")
		ßOP_IGNORE := πg.InternStr("OP_IGNORE")
		ßRANGE := πg.InternStr("RANGE")
		ßREPEAT := πg.InternStr("REPEAT")
		ßREPEAT_ONE := πg.InternStr("REPEAT_ONE")
		ßSRE_FLAG_DOTALL := πg.InternStr("SRE_FLAG_DOTALL")
		ßSRE_FLAG_IGNORECASE := πg.InternStr("SRE_FLAG_IGNORECASE")
		ßSRE_FLAG_LOCALE := πg.InternStr("SRE_FLAG_LOCALE")
		ßSRE_FLAG_MULTILINE := πg.InternStr("SRE_FLAG_MULTILINE")
		ßSRE_FLAG_TEMPLATE := πg.InternStr("SRE_FLAG_TEMPLATE")
		ßSRE_FLAG_UNICODE := πg.InternStr("SRE_FLAG_UNICODE")
		ßSRE_INFO_CHARSET := πg.InternStr("SRE_INFO_CHARSET")
		ßSRE_INFO_LITERAL := πg.InternStr("SRE_INFO_LITERAL")
		ßSRE_INFO_PREFIX := πg.InternStr("SRE_INFO_PREFIX")
		ßSTRING_TYPES := πg.InternStr("STRING_TYPES")
		ßSUBPATTERN := πg.InternStr("SUBPATTERN")
		ßSUCCESS := πg.InternStr("SUCCESS")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ß_ASSERT_CODES := πg.InternStr("_ASSERT_CODES")
		ß_BITS_TRANS := πg.InternStr("_BITS_TRANS")
		ß_CODEBITS := πg.InternStr("_CODEBITS")
		ß_LITERAL_CODES := πg.InternStr("_LITERAL_CODES")
		ß_REPEATING_CODES := πg.InternStr("_REPEATING_CODES")
		ß_SUCCESS_CODES := πg.InternStr("_SUCCESS_CODES")
		ß__all__ := πg.InternStr("__all__")
		ß__doc__ := πg.InternStr("__doc__")
		ß_bytes_to_codes := πg.InternStr("_bytes_to_codes")
		ß_code := πg.InternStr("_code")
		ß_compile := πg.InternStr("_compile")
		ß_compile_charset := πg.InternStr("_compile_charset")
		ß_compile_info := πg.InternStr("_compile_info")
		ß_equivalences := πg.InternStr("_equivalences")
		ß_fixup_range := πg.InternStr("_fixup_range")
		ß_ignorecase_fixes := πg.InternStr("_ignorecase_fixes")
		ß_mk_bitmap := πg.InternStr("_mk_bitmap")
		ß_optimize_charset := πg.InternStr("_optimize_charset")
		ß_simple := πg.InternStr("_simple")
		ß_sre := πg.InternStr("_sre")
		ßappend := πg.InternStr("append")
		ßcompile := πg.InternStr("compile")
		ßdata := πg.InternStr("data")
		ßenumerate := πg.InternStr("enumerate")
		ßerror := πg.InternStr("error")
		ßflags := πg.InternStr("flags")
		ßget := πg.InternStr("get")
		ßgetattr := πg.InternStr("getattr")
		ßgetlower := πg.InternStr("getlower")
		ßgetwidth := πg.InternStr("getwidth")
		ßglobals := πg.InternStr("globals")
		ßgroupdict := πg.InternStr("groupdict")
		ßgroups := πg.InternStr("groups")
		ßinsert := πg.InternStr("insert")
		ßisinstance := πg.InternStr("isinstance")
		ßisstring := πg.InternStr("isstring")
		ßitems := πg.InternStr("items")
		ßl := πg.InternStr("l")
		ßlen := πg.InternStr("len")
		ßmap := πg.InternStr("map")
		ßmax := πg.InternStr("max")
		ßmin := πg.InternStr("min")
		ßname := πg.InternStr("name")
		ßparse := πg.InternStr("parse")
		ßpattern := πg.InternStr("pattern")
		ßrange := πg.InternStr("range")
		ßset := πg.InternStr("set")
		ßsre_constants := πg.InternStr("sre_constants")
		ßsre_parse := πg.InternStr("sre_parse")
		ßstr := πg.InternStr("str")
		ßsys := πg.InternStr("sys")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßunicode := πg.InternStr("unicode")
		ßxrange := πg.InternStr("xrange")
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
		var πTemp010 []*πg.Object
		_ = πTemp010
		var πTemp011 []πg.Param
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
		var πTemp019 *πg.BaseException
		_ = πTemp019
		var πTemp020 *πg.Traceback
		_ = πTemp020
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 8: goto Label8
			case 1: goto Label1
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 12: """Internal support module for sre"""
			πF.SetLineno(12)
			// line 12: """Internal support module for sre"""
			πF.SetLineno(12)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Internal support module for sre").ToObject()); πE != nil {
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
			// line 15: import _sre
			πF.SetLineno(15)
			if πTemp002, πE = πg.ImportModule(πF, "_sre"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_sre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: import sre_parse
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "sre_parse"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsre_parse.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 19: import sre_constants
			πF.SetLineno(19)
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
			// line 21: globals()[name] = getattr(sre_constants, name)
			πF.SetLineno(21)
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
			// line 23: assert _sre.MAGIC == MAGIC, "SRE module mismatch"
			πF.SetLineno(23)
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_sre); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßMAGIC, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßMAGIC); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Eq(πF, πTemp004, πTemp003); πE != nil {
				continue
			}
			if πE = πg.Assert(πF, πTemp001, πg.NewStr("SRE module mismatch").ToObject()); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_sre); πE != nil {
				continue
			}
			if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßCODESIZE, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label4
			}
			goto Label5
			// line 25: if _sre.CODESIZE == 2:
			πF.SetLineno(25)
		Label4:
			// line 26: MAXCODE = 65535
			πF.SetLineno(26)
			if πE = πF.Globals().SetItem(πF, ßMAXCODE.ToObject(), πg.NewInt(65535).ToObject()); πE != nil {
				continue
			}
			goto Label6
		Label5:
			// line 28: MAXCODE = 0xFFFFFFFFL
			πF.SetLineno(28)
			if πE = πF.Globals().SetItem(πF, ßMAXCODE.ToObject(), πg.NewLongFromBytes([]byte{0xff,0xff,0xff,0xff,}).ToObject()); πE != nil {
				continue
			}
			goto Label6
		Label6:
			// line 30: _LITERAL_CODES = set([LITERAL, NOT_LITERAL])
			πF.SetLineno(30)
			πTemp002 = πF.MakeArgs(1)
			πTemp010 = make([]*πg.Object, 2)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
				continue
			}
			πTemp010[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNOT_LITERAL); πE != nil {
				continue
			}
			πTemp010[1] = πTemp001
			πTemp001 = πg.NewList(πTemp010...).ToObject()
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_LITERAL_CODES.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 31: _REPEATING_CODES = set([REPEAT, MIN_REPEAT, MAX_REPEAT])
			πF.SetLineno(31)
			πTemp002 = πF.MakeArgs(1)
			πTemp010 = make([]*πg.Object, 3)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßREPEAT); πE != nil {
				continue
			}
			πTemp010[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßMIN_REPEAT); πE != nil {
				continue
			}
			πTemp010[1] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßMAX_REPEAT); πE != nil {
				continue
			}
			πTemp010[2] = πTemp001
			πTemp001 = πg.NewList(πTemp010...).ToObject()
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_REPEATING_CODES.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 32: _SUCCESS_CODES = set([SUCCESS, FAILURE])
			πF.SetLineno(32)
			πTemp002 = πF.MakeArgs(1)
			πTemp010 = make([]*πg.Object, 2)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßSUCCESS); πE != nil {
				continue
			}
			πTemp010[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßFAILURE); πE != nil {
				continue
			}
			πTemp010[1] = πTemp001
			πTemp001 = πg.NewList(πTemp010...).ToObject()
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_SUCCESS_CODES.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 33: _ASSERT_CODES = set([ASSERT, ASSERT_NOT])
			πF.SetLineno(33)
			πTemp002 = πF.MakeArgs(1)
			πTemp010 = make([]*πg.Object, 2)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßASSERT); πE != nil {
				continue
			}
			πTemp010[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßASSERT_NOT); πE != nil {
				continue
			}
			πTemp010[1] = πTemp001
			πTemp001 = πg.NewList(πTemp010...).ToObject()
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_ASSERT_CODES.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 36: _equivalences = (
			πF.SetLineno(36)
			πTemp002 = make([]*πg.Object, 13)
			πTemp003 = πg.NewTuple2(πg.NewInt(105).ToObject(), πg.NewInt(305).ToObject()).ToObject()
			πTemp002[0] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(115).ToObject(), πg.NewInt(383).ToObject()).ToObject()
			πTemp002[1] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(181).ToObject(), πg.NewInt(956).ToObject()).ToObject()
			πTemp002[2] = πTemp003
			πTemp003 = πg.NewTuple3(πg.NewInt(837).ToObject(), πg.NewInt(953).ToObject(), πg.NewInt(8126).ToObject()).ToObject()
			πTemp002[3] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(946).ToObject(), πg.NewInt(976).ToObject()).ToObject()
			πTemp002[4] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(949).ToObject(), πg.NewInt(1013).ToObject()).ToObject()
			πTemp002[5] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(952).ToObject(), πg.NewInt(977).ToObject()).ToObject()
			πTemp002[6] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(954).ToObject(), πg.NewInt(1008).ToObject()).ToObject()
			πTemp002[7] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(960).ToObject(), πg.NewInt(982).ToObject()).ToObject()
			πTemp002[8] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(961).ToObject(), πg.NewInt(1009).ToObject()).ToObject()
			πTemp002[9] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(962).ToObject(), πg.NewInt(963).ToObject()).ToObject()
			πTemp002[10] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(966).ToObject(), πg.NewInt(981).ToObject()).ToObject()
			πTemp002[11] = πTemp003
			πTemp003 = πg.NewTuple2(πg.NewInt(7777).ToObject(), πg.NewInt(7835).ToObject()).ToObject()
			πTemp002[12] = πTemp003
			πTemp001 = πg.NewTuple(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_equivalences.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 66: _ignorecase_fixes = {i: tuple(j for j in t if i != j)
			πF.SetLineno(66)
			πTemp011 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µt *πg.Object = πg.UnboundLocal; _ = µt
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []πg.Param
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
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
						if πTemp002, πE = πg.ResolveGlobal(πF, ß_equivalences); πE != nil {
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
							µt = πTemp002
						}
						if πE != nil || !πTemp004 {
							continue
						}
						πF.PushCheckpoint(1)            
						if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Iter(πF, µt); πE != nil {
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
							µi = πTemp006
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(4)            
						// line 66: _ignorecase_fixes = {i: tuple(j for j in t if i != j)
						πF.SetLineno(66)
						if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
							continue
						}
						πTemp007 = πF.MakeArgs(1)
						πTemp009 = make([]πg.Param, 0)
						πTemp008 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/sre_compile.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
							var µj *πg.Object = πg.UnboundLocal; _ = µj
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
									if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
										continue
									}
									if πTemp001, πE = πg.Iter(πF, µt); πE != nil {
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
										µj = πTemp004
									}
									if πE != nil || !πTemp003 {
										continue
									}
									πF.PushCheckpoint(1)            
									if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.NE(πF, µi, µj); πE != nil {
										continue
									}
									if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label4
									}
									goto Label5
									// line 66: _ignorecase_fixes = {i: tuple(j for j in t if i != j)
									πF.SetLineno(66)
								Label4:
									// line 66: _ignorecase_fixes = {i: tuple(j for j in t if i != j)
									πF.SetLineno(66)
									if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
										continue
									}
									πF.PushCheckpoint(6)
									return µj, nil
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
						if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp007[0] = πTemp010
						if πTemp010, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
							continue
						}
						if πTemp011, πE = πTemp010.Call(πF, πTemp007, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp007)
						πTemp006 = πg.NewTuple2(µi, πTemp011).ToObject()
						πF.PushCheckpoint(7)
						return πTemp006, nil
					Label7:
						πTemp010 = πSent
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
			if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.DictType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_ignorecase_fixes.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 69: def _compile(code, pattern, flags):
			πF.SetLineno(69)
			πTemp011 = make([]πg.Param, 3)
			πTemp011[0] = πg.Param{Name: "code", Def: nil}
			πTemp011[1] = πg.Param{Name: "pattern", Def: nil}
			πTemp011[2] = πg.Param{Name: "flags", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_compile", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcode *πg.Object = πArgs[0]; _ = µcode
				var µpattern *πg.Object = πArgs[1]; _ = µpattern
				var µflags *πg.Object = πArgs[2]; _ = µflags
				var µemit *πg.Object = πg.UnboundLocal; _ = µemit
				var µ_len *πg.Object = πg.UnboundLocal; _ = µ_len
				var µLITERAL_CODES *πg.Object = πg.UnboundLocal; _ = µLITERAL_CODES
				var µREPEATING_CODES *πg.Object = πg.UnboundLocal; _ = µREPEATING_CODES
				var µSUCCESS_CODES *πg.Object = πg.UnboundLocal; _ = µSUCCESS_CODES
				var µASSERT_CODES *πg.Object = πg.UnboundLocal; _ = µASSERT_CODES
				var µfixes *πg.Object = πg.UnboundLocal; _ = µfixes
				var µop *πg.Object = πg.UnboundLocal; _ = µop
				var µav *πg.Object = πg.UnboundLocal; _ = µav
				var µlo *πg.Object = πg.UnboundLocal; _ = µlo
				var µskip *πg.Object = πg.UnboundLocal; _ = µskip
				var µk *πg.Object = πg.UnboundLocal; _ = µk
				var µfixup *πg.Object = πg.UnboundLocal; _ = µfixup
				var µhi *πg.Object = πg.UnboundLocal; _ = µhi
				var µtail *πg.Object = πg.UnboundLocal; _ = µtail
				var µtailappend *πg.Object = πg.UnboundLocal; _ = µtailappend
				var µskipyes *πg.Object = πg.UnboundLocal; _ = µskipyes
				var µskipno *πg.Object = πg.UnboundLocal; _ = µskipno
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 []πg.Param
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
					case 32: goto Label32
					case 33: goto Label33
					case 66: goto Label66
					case 67: goto Label67
					case 5: goto Label5
					case 6: goto Label6
					case 70: goto Label70
					case 69: goto Label69
					default: panic("unexpected function state")
					}
					// line 71: emit = code.append
					πF.SetLineno(71)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcode, ßappend, nil); πE != nil {
						continue
					}
					µemit = πTemp001
					// line 72: _len = len
					πF.SetLineno(72)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					µ_len = πTemp001
					// line 73: LITERAL_CODES = _LITERAL_CODES
					πF.SetLineno(73)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_LITERAL_CODES); πE != nil {
						continue
					}
					µLITERAL_CODES = πTemp001
					// line 74: REPEATING_CODES = _REPEATING_CODES
					πF.SetLineno(74)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_REPEATING_CODES); πE != nil {
						continue
					}
					µREPEATING_CODES = πTemp001
					// line 75: SUCCESS_CODES = _SUCCESS_CODES
					πF.SetLineno(75)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_SUCCESS_CODES); πE != nil {
						continue
					}
					µSUCCESS_CODES = πTemp001
					// line 76: ASSERT_CODES = _ASSERT_CODES
					πF.SetLineno(76)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_ASSERT_CODES); πE != nil {
						continue
					}
					µASSERT_CODES = πTemp001
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_IGNORECASE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µflags, πTemp004); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_LOCALE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp005); πE != nil {
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
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_UNICODE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µflags, πTemp004); πE != nil {
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
					// line 77: if (flags & SRE_FLAG_IGNORECASE and
					πF.SetLineno(77)
				Label2:
					// line 80: fixes = _ignorecase_fixes
					πF.SetLineno(80)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_ignorecase_fixes); πE != nil {
						continue
					}
					µfixes = πTemp001
					goto Label4
				Label3:
					// line 82: fixes = None
					πF.SetLineno(82)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µfixes = πTemp001
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µpattern); πE != nil {
						continue
					}
					πF.PushCheckpoint(6)
					πTemp002 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
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
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
							continue
						}
						µop = πTemp004
						µav = πTemp005
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(5)            
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µLITERAL_CODES, "LITERAL_CODES"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µLITERAL_CODES, µop); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßANY); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µREPEATING_CODES, "REPEATING_CODES"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µREPEATING_CODES, µop); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSUBPATTERN); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µSUCCESS_CODES, "SUCCESS_CODES"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µSUCCESS_CODES, µop); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µASSERT_CODES, "ASSERT_CODES"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µASSERT_CODES, µop); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßCALL); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label15
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßAT); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßBRANCH); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label17
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßGROUPREF); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label19
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßGROUPREF_EXISTS); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label20
					}
					goto Label21
					// line 84: if op in LITERAL_CODES:
					πF.SetLineno(84)
				Label8:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_IGNORECASE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µflags, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label23
					}
					goto Label24
					// line 85: if flags & SRE_FLAG_IGNORECASE:
					πF.SetLineno(85)
				Label23:
					// line 86: lo = _sre.getlower(av, flags)
					πF.SetLineno(86)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[0] = µav
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[1] = µflags
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_sre); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßgetlower, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µlo = πTemp003
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					πTemp003 = µfixes
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label26
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µfixes, µlo); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp008).ToObject()
					πTemp003 = πTemp004
				Label26:
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label27
					}
					goto Label28
					// line 87: if fixes and lo in fixes:
					πF.SetLineno(87)
				Label27:
					// line 88: emit(OPCODES[IN_IGNORE])
					πF.SetLineno(88)
					πTemp007 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßIN_IGNORE); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πTemp005, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 89: skip = _len(code); emit(0)
					πF.SetLineno(89)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp003, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskip = πTemp003
					// line 89: skip = _len(code); emit(0)
					πF.SetLineno(89)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNOT_LITERAL); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label30
					}
					goto Label31
					// line 90: if op is NOT_LITERAL:
					πF.SetLineno(90)
				Label30:
					// line 91: emit(OPCODES[NEGATE])
					πF.SetLineno(91)
					πTemp007 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNEGATE); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πTemp005, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label31
				Label31:
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple1(µlo).ToObject()
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp009 = µlo
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µfixes, πTemp009); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp005, πTemp010); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(33)
					πTemp006 = false
				Label32:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label34
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
						µk = πTemp004
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(32)            
					// line 93: emit(OPCODES[LITERAL])
					πF.SetLineno(93)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 94: emit(k)
					πF.SetLineno(94)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp007[0] = µk
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					continue
				Label33:
					if πE != nil || πR != nil {
						continue
					}
				Label34:
					// line 95: emit(OPCODES[FAILURE])
					πF.SetLineno(95)
					πTemp007 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßFAILURE); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πTemp005, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 96: code[skip] = _len(code) - skip
					πF.SetLineno(96)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, πTemp004, µskip); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp005 = µskip
					if πE = πg.SetItem(πF, µcode, πTemp005, πTemp004); πE != nil {
						continue
					}
					goto Label29
				Label28:
					// line 98: emit(OPCODES[OP_IGNORE[op]])
					πF.SetLineno(98)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOP_IGNORE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 99: emit(lo)
					πF.SetLineno(99)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp007[0] = µlo
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label29
				Label29:
					goto Label25
				Label24:
					// line 101: emit(OPCODES[op])
					πF.SetLineno(101)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp003 = µop
					if πTemp005, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 102: emit(av)
					πF.SetLineno(102)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[0] = µav
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label25
				Label25:
					goto Label22
					// line 103: elif op is IN:
					πF.SetLineno(103)
				Label9:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_IGNORECASE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µflags, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label35
					}
					goto Label36
					// line 104: if flags & SRE_FLAG_IGNORECASE:
					πF.SetLineno(104)
				Label35:
					// line 105: emit(OPCODES[OP_IGNORE[op]])
					πF.SetLineno(105)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOP_IGNORE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 106: def fixup(literal, flags=flags):
					πF.SetLineno(106)
					πTemp011 = make([]πg.Param, 2)
					πTemp011[0] = πg.Param{Name: "literal", Def: nil}
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp011[1] = πg.Param{Name: "flags", Def: µflags}
					πTemp003 = πg.NewFunction(πg.NewCode("fixup", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µliteral *πg.Object = πArgs[0]; _ = µliteral
						var µflags *πg.Object = πArgs[1]; _ = µflags
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
							// line 107: return _sre.getlower(literal, flags)
							πF.SetLineno(107)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µliteral, "literal"); πE != nil {
								continue
							}
							πTemp001[0] = µliteral
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp001[1] = µflags
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_sre); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgetlower, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					µfixup = πTemp003
					goto Label37
				Label36:
					// line 109: emit(OPCODES[op])
					πF.SetLineno(109)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 110: fixup = None
					πF.SetLineno(110)
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µfixup = πTemp004
					goto Label37
				Label37:
					// line 111: skip = _len(code); emit(0)
					πF.SetLineno(111)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskip = πTemp004
					// line 111: skip = _len(code); emit(0)
					πF.SetLineno(111)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 112: _compile_charset(av, flags, code, fixup, fixes)
					πF.SetLineno(112)
					πTemp007 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[0] = µav
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[1] = µflags
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[2] = µcode
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					πTemp007[3] = µfixup
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					πTemp007[4] = µfixes
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_compile_charset); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 113: code[skip] = _len(code) - skip
					πF.SetLineno(113)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, µskip); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp009 = µskip
					if πE = πg.SetItem(πF, µcode, πTemp009, πTemp005); πE != nil {
						continue
					}
					goto Label22
					// line 114: elif op is ANY:
					πF.SetLineno(114)
				Label10:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_DOTALL); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label38
					}
					goto Label39
					// line 115: if flags & SRE_FLAG_DOTALL:
					πF.SetLineno(115)
				Label38:
					// line 116: emit(OPCODES[ANY_ALL])
					πF.SetLineno(116)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßANY_ALL); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label40
				Label39:
					// line 118: emit(OPCODES[ANY])
					πF.SetLineno(118)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßANY); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label40
				Label40:
					goto Label22
					// line 119: elif op in REPEATING_CODES:
					πF.SetLineno(119)
				Label11:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_TEMPLATE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label41
					}
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[0] = µav
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_simple); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004 = πTemp009
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label42
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveGlobal(πF, ßREPEAT); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(µop != πTemp009).ToObject()
					πTemp004 = πTemp005
				Label42:
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label43
					}
					goto Label44
					// line 120: if flags & SRE_FLAG_TEMPLATE:
					πF.SetLineno(120)
				Label41:
					if πTemp004, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 121: raise error, "internal: unsupported template operator"
					πF.SetLineno(121)
					πE = πF.Raise(πTemp004, πg.NewStr("internal: unsupported template operator").ToObject(), nil)
					continue
					// line 122: emit(OPCODES[REPEAT])
					πF.SetLineno(122)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßREPEAT); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 123: skip = _len(code); emit(0)
					πF.SetLineno(123)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskip = πTemp004
					// line 123: skip = _len(code); emit(0)
					πF.SetLineno(123)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 124: emit(av[0])
					πF.SetLineno(124)
					πTemp007 = πF.MakeArgs(1)
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 125: emit(av[1])
					πF.SetLineno(125)
					πTemp007 = πF.MakeArgs(1)
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 126: _compile(code, av[2], flags)
					πF.SetLineno(126)
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[1] = πTemp005
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[2] = µflags
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 127: emit(OPCODES[SUCCESS])
					πF.SetLineno(127)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSUCCESS); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 128: code[skip] = _len(code) - skip
					πF.SetLineno(128)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, µskip); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp009 = µskip
					if πE = πg.SetItem(πF, µcode, πTemp009, πTemp005); πE != nil {
						continue
					}
					goto Label45
					// line 129: elif _simple(av) and op is not REPEAT:
					πF.SetLineno(129)
				Label43:
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMAX_REPEAT); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(µop == πTemp005).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label46
					}
					goto Label47
					// line 130: if op is MAX_REPEAT:
					πF.SetLineno(130)
				Label46:
					// line 131: emit(OPCODES[REPEAT_ONE])
					πF.SetLineno(131)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßREPEAT_ONE); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label48
				Label47:
					// line 133: emit(OPCODES[MIN_REPEAT_ONE])
					πF.SetLineno(133)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMIN_REPEAT_ONE); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label48
				Label48:
					// line 134: skip = _len(code); emit(0)
					πF.SetLineno(134)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskip = πTemp004
					// line 134: skip = _len(code); emit(0)
					πF.SetLineno(134)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 135: emit(av[0])
					πF.SetLineno(135)
					πTemp007 = πF.MakeArgs(1)
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 136: emit(av[1])
					πF.SetLineno(136)
					πTemp007 = πF.MakeArgs(1)
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 137: _compile(code, av[2], flags)
					πF.SetLineno(137)
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[1] = πTemp005
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[2] = µflags
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 138: emit(OPCODES[SUCCESS])
					πF.SetLineno(138)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSUCCESS); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 139: code[skip] = _len(code) - skip
					πF.SetLineno(139)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, µskip); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp009 = µskip
					if πE = πg.SetItem(πF, µcode, πTemp009, πTemp005); πE != nil {
						continue
					}
					goto Label45
				Label44:
					// line 141: emit(OPCODES[REPEAT])
					πF.SetLineno(141)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßREPEAT); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 142: skip = _len(code); emit(0)
					πF.SetLineno(142)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskip = πTemp004
					// line 142: skip = _len(code); emit(0)
					πF.SetLineno(142)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 143: emit(av[0])
					πF.SetLineno(143)
					πTemp007 = πF.MakeArgs(1)
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 144: emit(av[1])
					πF.SetLineno(144)
					πTemp007 = πF.MakeArgs(1)
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 145: _compile(code, av[2], flags)
					πF.SetLineno(145)
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[1] = πTemp005
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[2] = µflags
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 146: code[skip] = _len(code) - skip
					πF.SetLineno(146)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, µskip); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp009 = µskip
					if πE = πg.SetItem(πF, µcode, πTemp009, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMAX_REPEAT); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(µop == πTemp005).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label49
					}
					goto Label50
					// line 147: if op is MAX_REPEAT:
					πF.SetLineno(147)
				Label49:
					// line 148: emit(OPCODES[MAX_UNTIL])
					πF.SetLineno(148)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMAX_UNTIL); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label51
				Label50:
					// line 150: emit(OPCODES[MIN_UNTIL])
					πF.SetLineno(150)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMIN_UNTIL); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label51
				Label51:
					goto Label45
				Label45:
					goto Label22
					// line 151: elif op is SUBPATTERN:
					πF.SetLineno(151)
				Label12:
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label52
					}
					goto Label53
					// line 152: if av[0]:
					πF.SetLineno(152)
				Label52:
					// line 153: emit(OPCODES[MARK])
					πF.SetLineno(153)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 154: emit((av[0]-1)*2)
					πF.SetLineno(154)
					πTemp007 = πF.MakeArgs(1)
					πTemp009 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µav, πTemp009); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Mul(πF, πTemp005, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label53
				Label53:
					// line 156: _compile(code, av[1], flags)
					πF.SetLineno(156)
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[1] = πTemp005
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[2] = µflags
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label54
					}
					goto Label55
					// line 157: if av[0]:
					πF.SetLineno(157)
				Label54:
					// line 158: emit(OPCODES[MARK])
					πF.SetLineno(158)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 159: emit((av[0]-1)*2+1)
					πF.SetLineno(159)
					πTemp007 = πF.MakeArgs(1)
					πTemp010 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp012, πE = πg.GetItem(πF, µav, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Sub(πF, πTemp012, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, πTemp009, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label55
				Label55:
					goto Label22
					// line 160: elif op in SUCCESS_CODES:
					πF.SetLineno(160)
				Label13:
					// line 161: emit(OPCODES[op])
					πF.SetLineno(161)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label22
					// line 162: elif op in ASSERT_CODES:
					πF.SetLineno(162)
				Label14:
					// line 163: emit(OPCODES[op])
					πF.SetLineno(163)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 164: skip = _len(code); emit(0)
					πF.SetLineno(164)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskip = πTemp004
					// line 164: skip = _len(code); emit(0)
					πF.SetLineno(164)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µav, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GE(πF, πTemp009, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label56
					}
					goto Label57
					// line 165: if av[0] >= 0:
					πF.SetLineno(165)
				Label56:
					// line 166: emit(0) # look ahead
					πF.SetLineno(166)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label58
				Label57:
					// line 168: lo, hi = av[1].getwidth()
					πF.SetLineno(168)
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßgetwidth, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp009}}}, πTemp005); πE != nil {
						continue
					}
					µlo = πTemp004
					µhi = πTemp009
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label59
					}
					goto Label60
					// line 169: if lo != hi:
					πF.SetLineno(169)
				Label59:
					if πTemp004, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 170: raise error, "look-behind requires fixed-width pattern"
					πF.SetLineno(170)
					πE = πF.Raise(πTemp004, πg.NewStr("look-behind requires fixed-width pattern").ToObject(), nil)
					continue
					goto Label60
				Label60:
					// line 171: emit(lo) # look behind
					πF.SetLineno(171)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp007[0] = µlo
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label58
				Label58:
					// line 172: _compile(code, av[1], flags)
					πF.SetLineno(172)
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[1] = πTemp005
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[2] = µflags
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 173: emit(OPCODES[SUCCESS])
					πF.SetLineno(173)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSUCCESS); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 174: code[skip] = _len(code) - skip
					πF.SetLineno(174)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, µskip); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp009 = µskip
					if πE = πg.SetItem(πF, µcode, πTemp009, πTemp005); πE != nil {
						continue
					}
					goto Label22
					// line 175: elif op is CALL:
					πF.SetLineno(175)
				Label15:
					// line 176: emit(OPCODES[op])
					πF.SetLineno(176)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 177: skip = _len(code); emit(0)
					πF.SetLineno(177)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskip = πTemp004
					// line 177: skip = _len(code); emit(0)
					πF.SetLineno(177)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 178: _compile(code, av, flags)
					πF.SetLineno(178)
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[1] = µav
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[2] = µflags
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 179: emit(OPCODES[SUCCESS])
					πF.SetLineno(179)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSUCCESS); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 180: code[skip] = _len(code) - skip
					πF.SetLineno(180)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, µskip); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp009 = µskip
					if πE = πg.SetItem(πF, µcode, πTemp009, πTemp005); πE != nil {
						continue
					}
					goto Label22
					// line 181: elif op is AT:
					πF.SetLineno(181)
				Label16:
					// line 182: emit(OPCODES[op])
					πF.SetLineno(182)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_MULTILINE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label61
					}
					goto Label62
					// line 183: if flags & SRE_FLAG_MULTILINE:
					πF.SetLineno(183)
				Label61:
					// line 184: av = AT_MULTILINE.get(av, av)
					πF.SetLineno(184)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[0] = µav
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[1] = µav
					if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_MULTILINE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µav = πTemp004
					goto Label62
				Label62:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_LOCALE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label63
					}
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_UNICODE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label64
					}
					goto Label65
					// line 185: if flags & SRE_FLAG_LOCALE:
					πF.SetLineno(185)
				Label63:
					// line 186: av = AT_LOCALE.get(av, av)
					πF.SetLineno(186)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[0] = µav
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[1] = µav
					if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_LOCALE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µav = πTemp004
					goto Label65
					// line 187: elif flags & SRE_FLAG_UNICODE:
					πF.SetLineno(187)
				Label64:
					// line 188: av = AT_UNICODE.get(av, av)
					πF.SetLineno(188)
					πTemp007 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[0] = µav
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[1] = µav
					if πTemp004, πE = πg.ResolveGlobal(πF, ßAT_UNICODE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µav = πTemp004
					goto Label65
				Label65:
					// line 189: emit(ATCODES[av])
					πF.SetLineno(189)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp004 = µav
					if πTemp009, πE = πg.ResolveGlobal(πF, ßATCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label22
					// line 190: elif op is BRANCH:
					πF.SetLineno(190)
				Label17:
					// line 191: emit(OPCODES[op])
					πF.SetLineno(191)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 192: tail = []
					πF.SetLineno(192)
					πTemp007 = make([]*πg.Object, 0)
					πTemp004 = πg.NewList(πTemp007...).ToObject()
					µtail = πTemp004
					// line 193: tailappend = tail.append
					πF.SetLineno(193)
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µtail, ßappend, nil); πE != nil {
						continue
					}
					µtailappend = πTemp004
					πTemp005 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µav, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Iter(πF, πTemp009); πE != nil {
						continue
					}
					πF.PushCheckpoint(67)
					πTemp006 = false
				Label66:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label68
					}
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
						µav = πTemp005
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(66)            
					// line 195: skip = _len(code); emit(0)
					πF.SetLineno(195)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskip = πTemp005
					// line 195: skip = _len(code); emit(0)
					πF.SetLineno(195)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp005, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 197: _compile(code, av, flags)
					πF.SetLineno(197)
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp007[1] = µav
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[2] = µflags
					if πTemp005, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 198: emit(OPCODES[JUMP])
					πF.SetLineno(198)
					πTemp007 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveGlobal(πF, ßJUMP); πE != nil {
						continue
					}
					πTemp005 = πTemp009
					if πTemp010, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
						continue
					}
					πTemp007[0] = πTemp009
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp005, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 199: tailappend(_len(code)); emit(0)
					πF.SetLineno(199)
					πTemp007 = πF.MakeArgs(1)
					πTemp013 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp013[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_len.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µtailappend, "tailappend"); πE != nil {
						continue
					}
					if πTemp005, πE = µtailappend.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 199: tailappend(_len(code)); emit(0)
					πF.SetLineno(199)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp005, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 200: code[skip] = _len(code) - skip
					πF.SetLineno(200)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp009, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp009, µskip); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp010 = µskip
					if πE = πg.SetItem(πF, µcode, πTemp010, πTemp009); πE != nil {
						continue
					}
					continue
				Label67:
					if πE != nil || πR != nil {
						continue
					}
				Label68:
					// line 201: emit(0) # end of branch
					πF.SetLineno(201)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Iter(πF, µtail); πE != nil {
						continue
					}
					πF.PushCheckpoint(70)
					πTemp006 = false
				Label69:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label71
					}
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
						µtail = πTemp005
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(69)            
					// line 203: code[tail] = _len(code) - tail
					πF.SetLineno(203)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp009, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp009, µtail); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					πTemp010 = µtail
					if πE = πg.SetItem(πF, µcode, πTemp010, πTemp009); πE != nil {
						continue
					}
					continue
				Label70:
					if πE != nil || πR != nil {
						continue
					}
				Label71:
					goto Label22
					// line 204: elif op is CATEGORY:
					πF.SetLineno(204)
				Label18:
					// line 205: emit(OPCODES[op])
					πF.SetLineno(205)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_LOCALE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label72
					}
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_UNICODE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label73
					}
					goto Label74
					// line 206: if flags & SRE_FLAG_LOCALE:
					πF.SetLineno(206)
				Label72:
					// line 207: av = CH_LOCALE[av]
					πF.SetLineno(207)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp004 = µav
					if πTemp009, πE = πg.ResolveGlobal(πF, ßCH_LOCALE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					µav = πTemp005
					goto Label74
					// line 208: elif flags & SRE_FLAG_UNICODE:
					πF.SetLineno(208)
				Label73:
					// line 209: av = CH_UNICODE[av]
					πF.SetLineno(209)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp004 = µav
					if πTemp009, πE = πg.ResolveGlobal(πF, ßCH_UNICODE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					µav = πTemp005
					goto Label74
				Label74:
					// line 210: emit(CHCODES[av])
					πF.SetLineno(210)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp004 = µav
					if πTemp009, πE = πg.ResolveGlobal(πF, ßCHCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label22
					// line 211: elif op is GROUPREF:
					πF.SetLineno(211)
				Label19:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_IGNORECASE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp005); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label75
					}
					goto Label76
					// line 212: if flags & SRE_FLAG_IGNORECASE:
					πF.SetLineno(212)
				Label75:
					// line 213: emit(OPCODES[OP_IGNORE[op]])
					πF.SetLineno(213)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp005 = µop
					if πTemp010, πE = πg.ResolveGlobal(πF, ßOP_IGNORE); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp005); πE != nil {
						continue
					}
					πTemp004 = πTemp009
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label77
				Label76:
					// line 215: emit(OPCODES[op])
					πF.SetLineno(215)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label77
				Label77:
					// line 216: emit(av-1)
					πF.SetLineno(216)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, µav, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					goto Label22
					// line 217: elif op is GROUPREF_EXISTS:
					πF.SetLineno(217)
				Label20:
					// line 218: emit(OPCODES[op])
					πF.SetLineno(218)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp004 = µop
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 219: emit(av[0]-1)
					πF.SetLineno(219)
					πTemp007 = πF.MakeArgs(1)
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetItem(πF, µav, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp007[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 220: skipyes = _len(code); emit(0)
					πF.SetLineno(220)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskipyes = πTemp004
					// line 220: skipyes = _len(code); emit(0)
					πF.SetLineno(220)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 221: _compile(code, av[1], flags)
					πF.SetLineno(221)
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[1] = πTemp005
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[2] = µflags
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label78
					}
					goto Label79
					// line 222: if av[2]:
					πF.SetLineno(222)
				Label78:
					// line 223: emit(OPCODES[JUMP])
					πF.SetLineno(223)
					πTemp007 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveGlobal(πF, ßJUMP); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πTemp009, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
						continue
					}
					πTemp007[0] = πTemp005
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 224: skipno = _len(code); emit(0)
					πF.SetLineno(224)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp004, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					µskipno = πTemp004
					// line 224: skipno = _len(code); emit(0)
					πF.SetLineno(224)
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp004, πE = µemit.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 225: code[skipyes] = _len(code) - skipyes + 1
					πF.SetLineno(225)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp009, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskipyes, "skipyes"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp009, µskipyes); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskipyes, "skipyes"); πE != nil {
						continue
					}
					πTemp009 = µskipyes
					if πE = πg.SetItem(πF, µcode, πTemp009, πTemp005); πE != nil {
						continue
					}
					// line 226: _compile(code, av[2], flags)
					πF.SetLineno(226)
					πTemp007 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µav, πTemp004); πE != nil {
						continue
					}
					πTemp007[1] = πTemp005
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp007[2] = µflags
					if πTemp004, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 227: code[skipno] = _len(code) - skipno
					πF.SetLineno(227)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp005, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskipno, "skipno"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, µskipno); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskipno, "skipno"); πE != nil {
						continue
					}
					πTemp009 = µskipno
					if πE = πg.SetItem(πF, µcode, πTemp009, πTemp005); πE != nil {
						continue
					}
					goto Label80
				Label79:
					// line 229: code[skipyes] = _len(code) - skipyes + 1
					πF.SetLineno(229)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp007[0] = µcode
					if πE = πg.CheckLocal(πF, µ_len, "_len"); πE != nil {
						continue
					}
					if πTemp009, πE = µ_len.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πE = πg.CheckLocal(πF, µskipyes, "skipyes"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Sub(πF, πTemp009, µskipyes); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskipyes, "skipyes"); πE != nil {
						continue
					}
					πTemp009 = µskipyes
					if πE = πg.SetItem(πF, µcode, πTemp009, πTemp005); πE != nil {
						continue
					}
					goto Label80
				Label80:
					goto Label22
				Label21:
					if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(πg.NewStr("unsupported operand type").ToObject(), µop).ToObject()
					// line 231: raise ValueError, ("unsupported operand type", op)
					πF.SetLineno(231)
					πE = πF.Raise(πTemp004, πTemp005, nil)
					continue
					goto Label22
				Label22:
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
			if πE = πF.Globals().SetItem(πF, ß_compile.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 233: def _compile_charset(charset, flags, code, fixup=None, fixes=None):
			πF.SetLineno(233)
			πTemp011 = make([]πg.Param, 5)
			πTemp011[0] = πg.Param{Name: "charset", Def: nil}
			πTemp011[1] = πg.Param{Name: "flags", Def: nil}
			πTemp011[2] = πg.Param{Name: "code", Def: nil}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp011[3] = πg.Param{Name: "fixup", Def: πTemp007}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp011[4] = πg.Param{Name: "fixes", Def: πTemp007}
			πTemp004 = πg.NewFunction(πg.NewCode("_compile_charset", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcharset *πg.Object = πArgs[0]; _ = µcharset
				var µflags *πg.Object = πArgs[1]; _ = µflags
				var µcode *πg.Object = πArgs[2]; _ = µcode
				var µfixup *πg.Object = πArgs[3]; _ = µfixup
				var µfixes *πg.Object = πArgs[4]; _ = µfixes
				var µemit *πg.Object = πg.UnboundLocal; _ = µemit
				var µop *πg.Object = πg.UnboundLocal; _ = µop
				var µav *πg.Object = πg.UnboundLocal; _ = µav
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
					// line 235: emit = code.append
					πF.SetLineno(235)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcode, ßappend, nil); πE != nil {
						continue
					}
					µemit = πTemp001
					πTemp002 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µcharset, "charset"); πE != nil {
						continue
					}
					πTemp002[0] = µcharset
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					πTemp002[1] = µfixup
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					πTemp002[2] = µfixes
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_UNICODE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µflags, πTemp004); πE != nil {
						continue
					}
					πTemp002[3] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_optimize_charset); πE != nil {
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
						µop = πTemp004
						µav = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 238: emit(OPCODES[op])
					πF.SetLineno(238)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					πTemp003 = µop
					if πTemp007, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNEGATE); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßRANGE); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßCHARSET); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßBIGCHARSET); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßCATEGORY); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp004).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 239: if op is NEGATE:
					πF.SetLineno(239)
				Label4:
					// line 240: pass
					πF.SetLineno(240)
					goto Label11
					// line 241: elif op is LITERAL:
					πF.SetLineno(241)
				Label5:
					// line 242: emit(av)
					πF.SetLineno(242)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp002[0] = µav
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label11
					// line 243: elif op is RANGE:
					πF.SetLineno(243)
				Label6:
					// line 244: emit(av[0])
					πF.SetLineno(244)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µav, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 245: emit(av[1])
					πF.SetLineno(245)
					πTemp002 = πF.MakeArgs(1)
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µav, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label11
					// line 246: elif op is CHARSET:
					πF.SetLineno(246)
				Label7:
					// line 248: code += (av)
					πF.SetLineno(248)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µcode, µav); πE != nil {
						continue
					}
					µcode = πTemp003
					goto Label11
					// line 249: elif op is BIGCHARSET:
					πF.SetLineno(249)
				Label8:
					// line 251: code += (av)
					πF.SetLineno(251)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µcode, µav); πE != nil {
						continue
					}
					µcode = πTemp003
					goto Label11
					// line 252: elif op is CATEGORY:
					πF.SetLineno(252)
				Label9:
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_LOCALE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µflags, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_UNICODE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µflags, πTemp004); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					goto Label14
					// line 253: if flags & SRE_FLAG_LOCALE:
					πF.SetLineno(253)
				Label12:
					// line 254: emit(CHCODES[CH_LOCALE[av]])
					πF.SetLineno(254)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp004 = µav
					if πTemp008, πE = πg.ResolveGlobal(πF, ßCH_LOCALE); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßCHCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label15
					// line 255: elif flags & SRE_FLAG_UNICODE:
					πF.SetLineno(255)
				Label13:
					// line 256: emit(CHCODES[CH_UNICODE[av]])
					πF.SetLineno(256)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp004 = µav
					if πTemp008, πE = πg.ResolveGlobal(πF, ßCH_UNICODE); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßCHCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label15
				Label14:
					// line 258: emit(CHCODES[av])
					πF.SetLineno(258)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp003 = µav
					if πTemp007, πE = πg.ResolveGlobal(πF, ßCHCODES); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
						continue
					}
					πTemp002[0] = πTemp004
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp003, πE = µemit.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label15
				Label15:
					goto Label11
				Label10:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßerror); πE != nil {
						continue
					}
					// line 260: raise error, "internal: unsupported set operator"
					πF.SetLineno(260)
					πE = πF.Raise(πTemp003, πg.NewStr("internal: unsupported set operator").ToObject(), nil)
					continue
					goto Label11
				Label11:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 261: emit(OPCODES[FAILURE])
					πF.SetLineno(261)
					πTemp002 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßFAILURE); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp004, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp002, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_compile_charset.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 263: def _optimize_charset(charset, fixup, fixes, isunicode):
			πF.SetLineno(263)
			πTemp011 = make([]πg.Param, 4)
			πTemp011[0] = πg.Param{Name: "charset", Def: nil}
			πTemp011[1] = πg.Param{Name: "fixup", Def: nil}
			πTemp011[2] = πg.Param{Name: "fixes", Def: nil}
			πTemp011[3] = πg.Param{Name: "isunicode", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_optimize_charset", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcharset *πg.Object = πArgs[0]; _ = µcharset
				var µfixup *πg.Object = πArgs[1]; _ = µfixup
				var µfixes *πg.Object = πArgs[2]; _ = µfixes
				var µisunicode *πg.Object = πArgs[3]; _ = µisunicode
				var µout *πg.Object = πg.UnboundLocal; _ = µout
				var µtail *πg.Object = πg.UnboundLocal; _ = µtail
				var µcharmap *πg.Object = πg.UnboundLocal; _ = µcharmap
				var µop *πg.Object = πg.UnboundLocal; _ = µop
				var µav *πg.Object = πg.UnboundLocal; _ = µav
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µk *πg.Object = πg.UnboundLocal; _ = µk
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µlo *πg.Object = πg.UnboundLocal; _ = µlo
				var µhi *πg.Object = πg.UnboundLocal; _ = µhi
				var µranges *πg.Object = πg.UnboundLocal; _ = µranges
				var µruns *πg.Object = πg.UnboundLocal; _ = µruns
				var µq *πg.Object = πg.UnboundLocal; _ = µq
				var µchar_find *πg.Object = πg.UnboundLocal; _ = µchar_find
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µdata *πg.Object = πg.UnboundLocal; _ = µdata
				var µcomps *πg.Object = πg.UnboundLocal; _ = µcomps
				var µmapping *πg.Object = πg.UnboundLocal; _ = µmapping
				var µblock *πg.Object = πg.UnboundLocal; _ = µblock
				var µchunk *πg.Object = πg.UnboundLocal; _ = µchunk
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 bool
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πTemp012 *πg.BaseException
				_ = πTemp012
				var πTemp013 *πg.Traceback
				_ = πTemp013
				var πTemp014 []*πg.Object
				_ = πTemp014
				var πTemp015 []πg.Param
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πTemp017 *πg.Dict
				_ = πTemp017
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 64: goto Label64
					case 1: goto Label1
					case 2: goto Label2
					case 35: goto Label35
					case 4: goto Label4
					case 5: goto Label5
					case 38: goto Label38
					case 65: goto Label65
					case 8: goto Label8
					case 75: goto Label75
					case 76: goto Label76
					case 34: goto Label34
					case 47: goto Label47
					case 48: goto Label48
					case 53: goto Label53
					case 20: goto Label20
					case 21: goto Label21
					case 54: goto Label54
					case 29: goto Label29
					case 30: goto Label30
					case 37: goto Label37
					default: panic("unexpected function state")
					}
					// line 265: out = []
					πF.SetLineno(265)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µout = πTemp002
					// line 266: tail = []
					πF.SetLineno(266)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µtail = πTemp002
					// line 268: charmap = [0] * 256
					πF.SetLineno(268)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewInt(0).ToObject()
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					if πTemp002, πE = πg.Mul(πF, πTemp003, πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					µcharmap = πTemp002
					if πE = πg.CheckLocal(πF, µcharset, "charset"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µcharset); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µop = πTemp006
						µav = πTemp007
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 270: while True:
					πF.SetLineno(270)
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
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πE != nil || !πTemp008 {
						continue
					}
					πF.PushCheckpoint(4)            
					// line 271: try:
					πF.SetLineno(271)
					πF.PushCheckpoint(8)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp006).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label9
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßRANGE); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp006).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label10
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNEGATE); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp006).ToObject()
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label11
					}
					goto Label12
					// line 272: if op is LITERAL:
					πF.SetLineno(272)
				Label9:
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µfixup); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label14
					}
					goto Label15
					// line 273: if fixup:
					πF.SetLineno(273)
				Label14:
					// line 274: i = fixup(av)
					πF.SetLineno(274)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp001[0] = µav
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					if πTemp003, πE = µfixup.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µi = πTemp003
					// line 275: charmap[i] = 1
					πF.SetLineno(275)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp006 = µi
					if πE = πg.SetItem(πF, µcharmap, πTemp006, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					πTemp003 = µfixes
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label17
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, µfixes, µi); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(πTemp009).ToObject()
					πTemp003 = πTemp006
				Label17:
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label18
					}
					goto Label19
					// line 276: if fixes and i in fixes:
					πF.SetLineno(276)
				Label18:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp006 = µi
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µfixes, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(21)
					πTemp008 = false
				Label20:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label22
					}
					if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
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
					πF.PushCheckpoint(20)            
					// line 278: charmap[k] = 1
					πF.SetLineno(278)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp007 = µk
					if πE = πg.SetItem(πF, µcharmap, πTemp007, πTemp006); πE != nil {
						continue
					}
					continue
				Label21:
					if πE != nil || πR != nil {
						continue
					}
				Label22:
					goto Label19
				Label19:
					goto Label16
				Label15:
					// line 280: charmap[av] = 1
					πF.SetLineno(280)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp006 = µav
					if πE = πg.SetItem(πF, µcharmap, πTemp006, πTemp003); πE != nil {
						continue
					}
					goto Label16
				Label16:
					goto Label13
					// line 281: elif op is RANGE:
					πF.SetLineno(281)
				Label10:
					// line 282: r = range(av[0], av[1]+1)
					πF.SetLineno(282)
					πTemp001 = πF.MakeArgs(2)
					πTemp003 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µav, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					πTemp006 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µav, πTemp006); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[1] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp006
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µfixup); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label23
					}
					goto Label24
					// line 283: if fixup:
					πF.SetLineno(283)
				Label23:
					// line 284: r = map(fixup, r)
					πF.SetLineno(284)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					πTemp001[0] = µfixup
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πTemp001[1] = µr
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µr = πTemp006
					goto Label24
				Label24:
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					πTemp003 = µfixup
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label25
					}
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					πTemp003 = µfixes
				Label25:
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label26
					}
					goto Label27
					// line 285: if fixup and fixes:
					πF.SetLineno(285)
				Label26:
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, µr); πE != nil {
						continue
					}
					πF.PushCheckpoint(30)
					πTemp008 = false
				Label29:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label31
					}
					if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
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
						µi = πTemp006
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(29)            
					// line 287: charmap[i] = 1
					πF.SetLineno(287)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.SetItem(πF, µcharmap, πTemp007, πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.Contains(πF, µfixes, µi); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(πTemp009).ToObject()
					if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label32
					}
					goto Label33
					// line 288: if i in fixes:
					πF.SetLineno(288)
				Label32:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.CheckLocal(πF, µfixes, "fixes"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µfixes, πTemp007); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Iter(πF, πTemp010); πE != nil {
						continue
					}
					πF.PushCheckpoint(35)
					πTemp009 = false
				Label34:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp009 {
						πF.PopCheckpoint()
						goto Label36
					}
					if πTemp007, πE = πg.Next(πF, πTemp006); πE != nil {
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
						µk = πTemp007
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(34)            
					// line 290: charmap[k] = 1
					πF.SetLineno(290)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp010 = µk
					if πE = πg.SetItem(πF, µcharmap, πTemp010, πTemp007); πE != nil {
						continue
					}
					continue
				Label35:
					if πE != nil || πR != nil {
						continue
					}
				Label36:
					goto Label33
				Label33:
					continue
				Label30:
					if πE != nil || πR != nil {
						continue
					}
				Label31:
					goto Label28
				Label27:
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, µr); πE != nil {
						continue
					}
					πF.PushCheckpoint(38)
					πTemp008 = false
				Label37:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label39
					}
					if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
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
						µi = πTemp006
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(37)            
					// line 293: charmap[i] = 1
					πF.SetLineno(293)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.SetItem(πF, µcharmap, πTemp007, πTemp006); πE != nil {
						continue
					}
					continue
				Label38:
					if πE != nil || πR != nil {
						continue
					}
				Label39:
					goto Label28
				Label28:
					goto Label13
					// line 294: elif op is NEGATE:
					πF.SetLineno(294)
				Label11:
					// line 295: out.append((op, av))
					πF.SetLineno(295)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µop, µav).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µout, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label13
				Label12:
					// line 297: tail.append((op, av))
					πF.SetLineno(297)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µop, µav).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µtail, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label13
				Label13:
					πF.PopCheckpoint()
					goto Label7
				Label8:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp012, πTemp013 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsInstance(πF, πTemp012.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label40
					}
					πE = πF.Raise(πTemp012.ToObject(), nil, πTemp013.ToObject())
					continue
					// line 298: except IndexError:
					πF.SetLineno(298)
				Label40:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					πTemp001[0] = µcharmap
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label41
					}
					goto Label42
					// line 299: if len(charmap) == 256:
					πF.SetLineno(299)
				Label41:
					// line 301: charmap += b'\0' * 0xff00
					πF.SetLineno(301)
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, πg.NewStr("\x00").ToObject(), πg.NewInt(65280).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IAdd(πF, µcharmap, πTemp003); πE != nil {
						continue
					}
					µcharmap = πTemp006
					// line 302: continue
					πF.SetLineno(302)
					continue
					goto Label42
				Label42:
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					πTemp003 = µfixup
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label43
					}
					if πE = πg.CheckLocal(πF, µisunicode, "isunicode"); πE != nil {
						continue
					}
					πTemp003 = µisunicode
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label43
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßRANGE); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(µop == πTemp007).ToObject()
					πTemp003 = πTemp006
				Label43:
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label44
					}
					goto Label45
					// line 304: if fixup and isunicode and op is RANGE:
					πF.SetLineno(304)
				Label44:
					// line 305: lo, hi = av
					πF.SetLineno(305)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, µav); πE != nil {
						continue
					}
					µlo = πTemp003
					µhi = πTemp006
					// line 306: ranges = [av]
					πF.SetLineno(306)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp001[0] = µav
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					µranges = πTemp003
					// line 309: _fixup_range(max(0x10000, lo), min(0x11fff, hi),
					πF.SetLineno(309)
					πTemp001 = πF.MakeArgs(4)
					πTemp014 = πF.MakeArgs(2)
					πTemp014[0] = πg.NewInt(65536).ToObject()
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp014[1] = µlo
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					πTemp001[0] = πTemp006
					πTemp014 = πF.MakeArgs(2)
					πTemp014[0] = πg.NewInt(73727).ToObject()
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					πTemp014[1] = µhi
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					πTemp001[1] = πTemp006
					if πE = πg.CheckLocal(πF, µranges, "ranges"); πE != nil {
						continue
					}
					πTemp001[2] = µranges
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					πTemp001[3] = µfixup
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_fixup_range); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µranges, "ranges"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, µranges); πE != nil {
						continue
					}
					πF.PushCheckpoint(48)
					πTemp008 = false
				Label47:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp008 {
						πF.PopCheckpoint()
						goto Label49
					}
					if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
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
						µlo = πTemp007
						µhi = πTemp010
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(47)            
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label50
					}
					goto Label51
					// line 312: if lo == hi:
					πF.SetLineno(312)
				Label50:
					// line 313: tail.append((LITERAL, hi))
					πF.SetLineno(313)
					πTemp001 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(πTemp007, µhi).ToObject()
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µtail, ßappend, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label52
				Label51:
					// line 315: tail.append((RANGE, (lo, hi)))
					πF.SetLineno(315)
					πTemp001 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveGlobal(πF, ßRANGE); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					πTemp010 = πg.NewTuple2(µlo, µhi).ToObject()
					πTemp006 = πg.NewTuple2(πTemp007, πTemp010).ToObject()
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µtail, ßappend, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label52
				Label52:
					continue
				Label48:
					if πE != nil || πR != nil {
						continue
					}
				Label49:
					goto Label46
				Label45:
					// line 317: tail.append((op, av))
					πF.SetLineno(317)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µop, µav).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µtail, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label46
				Label46:
					πF.RestoreExc(nil, nil)
					goto Label7
				Label7:
					// line 318: break
					πF.SetLineno(318)
					πTemp005 = true
					continue
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
					// line 321: runs = []
					πF.SetLineno(321)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µruns = πTemp002
					// line 322: q = 0
					πF.SetLineno(322)
					µq = πg.NewInt(0).ToObject()
					// line 323: def char_find(l, s, start):
					πF.SetLineno(323)
					πTemp015 = make([]πg.Param, 3)
					πTemp015[0] = πg.Param{Name: "l", Def: nil}
					πTemp015[1] = πg.Param{Name: "s", Def: nil}
					πTemp015[2] = πg.Param{Name: "start", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("char_find", "/usr/lib/python2.7/sre_compile.py", πTemp015, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µl *πg.Object = πArgs[0]; _ = µl
						var µs *πg.Object = πArgs[1]; _ = µs
						var µstart *πg.Object = πArgs[2]; _ = µstart
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var πTemp001 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 324: i = start
							πF.SetLineno(324)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							µi = µstart
							// line 325: while i < len(l):
							πF.SetLineno(325)
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
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp004[0] = µl
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.LT(πF, µi, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp005 = µi
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µl, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, µs); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 326: if l[i] == s:
							πF.SetLineno(326)
						Label4:
							// line 327: return i
							πF.SetLineno(327)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πR = µi
							continue
							goto Label5
						Label5:
							// line 328: i += 1
							πF.SetLineno(328)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp003
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 329: return -1
							πF.SetLineno(329)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
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
					µchar_find = πTemp002
					// line 330: while True:
					πF.SetLineno(330)
					πF.PushCheckpoint(54)
					πTemp004 = false
				Label53:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label55
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(53)            
					// line 332: p = char_find(charmap, 1, q)
					πF.SetLineno(332)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					πTemp001[0] = µcharmap
					πTemp001[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					πTemp001[2] = µq
					if πE = πg.CheckLocal(πF, µchar_find, "char_find"); πE != nil {
						continue
					}
					if πTemp003, πE = µchar_find.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp003
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µp, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label56
					}
					goto Label57
					// line 333: if p < 0:
					πF.SetLineno(333)
				Label56:
					// line 334: break
					πF.SetLineno(334)
					πTemp004 = true
					continue
					goto Label57
				Label57:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µruns, "runs"); πE != nil {
						continue
					}
					πTemp001[0] = µruns
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.GE(πF, πTemp007, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label58
					}
					goto Label59
					// line 335: if len(runs) >= 2:
					πF.SetLineno(335)
				Label58:
					// line 336: runs = None
					πF.SetLineno(336)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µruns = πTemp003
					// line 337: break
					πF.SetLineno(337)
					πTemp004 = true
					continue
					goto Label59
				Label59:
					// line 339: q = char_find(charmap, 0, p)
					πF.SetLineno(339)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					πTemp001[0] = µcharmap
					πTemp001[1] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp001[2] = µp
					if πE = πg.CheckLocal(πF, µchar_find, "char_find"); πE != nil {
						continue
					}
					if πTemp003, πE = µchar_find.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µq = πTemp003
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µq, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label60
					}
					goto Label61
					// line 340: if q < 0:
					πF.SetLineno(340)
				Label60:
					// line 341: runs.append((p, len(charmap)))
					πF.SetLineno(341)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp014 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					πTemp014[0] = µcharmap
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					πTemp003 = πg.NewTuple2(µp, πTemp007).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µruns, "runs"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µruns, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 342: break
					πF.SetLineno(342)
					πTemp004 = true
					continue
					goto Label61
				Label61:
					// line 343: runs.append((p, q))
					πF.SetLineno(343)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(µp, µq).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µruns, "runs"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µruns, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label54:
					if πE != nil || πR != nil {
						continue
					}
				Label55:
					if πE = πg.CheckLocal(πF, µruns, "runs"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µruns != πTemp006).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label62
					}
					goto Label63
					// line 344: if runs is not None:
					πF.SetLineno(344)
				Label62:
					if πE = πg.CheckLocal(πF, µruns, "runs"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Iter(πF, µruns); πE != nil {
						continue
					}
					πF.PushCheckpoint(65)
					πTemp004 = false
				Label64:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label66
					}
					if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp010}}}, πTemp006); πE != nil {
							continue
						}
						µp = πTemp007
						µq = πTemp010
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(64)            
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Sub(πF, µq, µp); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label67
					}
					goto Label68
					// line 347: if q - p == 1:
					πF.SetLineno(347)
				Label67:
					// line 348: out.append((LITERAL, p))
					πF.SetLineno(348)
					πTemp001 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp006 = πg.NewTuple2(πTemp007, µp).ToObject()
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µout, ßappend, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label69
				Label68:
					// line 350: out.append((RANGE, (p, q - 1)))
					πF.SetLineno(350)
					πTemp001 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveGlobal(πF, ßRANGE); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					if πTemp016, πE = πg.Sub(πF, µq, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp010 = πg.NewTuple2(µp, πTemp016).ToObject()
					πTemp006 = πg.NewTuple2(πTemp007, πTemp010).ToObject()
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µout, ßappend, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					goto Label69
				Label69:
					continue
				Label65:
					if πE != nil || πR != nil {
						continue
					}
				Label66:
					// line 351: out += tail
					πF.SetLineno(351)
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µout, µtail); πE != nil {
						continue
					}
					µout = πTemp003
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					πTemp003 = µfixup
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label70
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					πTemp001[0] = µout
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcharset, "charset"); πE != nil {
						continue
					}
					πTemp001[0] = µcharset
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp006, πE = πg.LT(πF, πTemp010, πTemp016); πE != nil {
						continue
					}
					πTemp003 = πTemp006
				Label70:
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label71
					}
					goto Label72
					// line 353: if fixup or len(out) < len(charset):
					πF.SetLineno(353)
				Label71:
					// line 354: return out
					πF.SetLineno(354)
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					πR = µout
					continue
					goto Label72
				Label72:
					// line 356: return charset
					πF.SetLineno(356)
					if πE = πg.CheckLocal(πF, µcharset, "charset"); πE != nil {
						continue
					}
					πR = µcharset
					continue
					goto Label63
				Label63:
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					πTemp001[0] = µcharmap
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Eq(πF, πTemp007, πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label73
					}
					goto Label74
					// line 359: if len(charmap) == 256:
					πF.SetLineno(359)
				Label73:
					// line 360: data = _mk_bitmap(charmap)
					πF.SetLineno(360)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					πTemp001[0] = µcharmap
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_mk_bitmap); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdata = πTemp006
					// line 361: out.append((CHARSET, data))
					πF.SetLineno(361)
					πTemp001 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßCHARSET); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp006, µdata).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µout, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 362: out += tail
					πF.SetLineno(362)
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µout, µtail); πE != nil {
						continue
					}
					µout = πTemp003
					// line 363: return out
					πF.SetLineno(363)
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					πR = µout
					continue
					goto Label74
				Label74:
					// line 390: charmap = str(charmap) # should be hashable
					πF.SetLineno(390)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					πTemp001[0] = µcharmap
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcharmap = πTemp006
					// line 391: comps = {}
					πF.SetLineno(391)
					πTemp017 = πg.NewDict()
					πTemp003 = πTemp017.ToObject()
					µcomps = πTemp003
					// line 393: mapping = [0] * 256
					πF.SetLineno(393)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewInt(0).ToObject()
					πTemp006 = πg.NewList(πTemp001...).ToObject()
					if πTemp003, πE = πg.Mul(πF, πTemp006, πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					µmapping = πTemp003
					// line 394: block = 0
					πF.SetLineno(394)
					µblock = πg.NewInt(0).ToObject()
					// line 396: data = []
					πF.SetLineno(396)
					πTemp001 = make([]*πg.Object, 0)
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					µdata = πTemp003
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewInt(0).ToObject()
					πTemp001[1] = πg.NewInt(65536).ToObject()
					πTemp001[2] = πg.NewInt(256).ToObject()
					if πTemp006, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
						continue
					}
					πF.PushCheckpoint(76)
					πTemp004 = false
				Label75:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label77
					}
					if πTemp006, πE = πg.Next(πF, πTemp003); πE != nil {
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
						µi = πTemp006
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(75)            
					// line 398: chunk = charmap[i: i + 256]
					πF.SetLineno(398)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, µi, πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µi, πTemp007, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcharmap, "charmap"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µcharmap, πTemp006); πE != nil {
						continue
					}
					µchunk = πTemp007
					if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcomps, "comps"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µcomps, µchunk); πE != nil {
						continue
					}
					πTemp006 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label78
					}
					goto Label79
					// line 399: if chunk in comps:
					πF.SetLineno(399)
				Label78:
					// line 400: mapping[i // 256] = comps[chunk]
					πF.SetLineno(400)
					if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
						continue
					}
					πTemp006 = µchunk
					if πE = πg.CheckLocal(πF, µcomps, "comps"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µcomps, πTemp006); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp007); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmapping, "mapping"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp016, πE = πg.FloorDiv(πF, µi, πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					πTemp010 = πTemp016
					if πE = πg.SetItem(πF, µmapping, πTemp010, πTemp006); πE != nil {
						continue
					}
					goto Label80
				Label79:
					// line 402: mapping[i // 256] = comps[chunk] = block
					πF.SetLineno(402)
					if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µblock); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmapping, "mapping"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.FloorDiv(πF, µi, πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					πTemp007 = πTemp010
					if πE = πg.SetItem(πF, µmapping, πTemp007, πTemp006); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µblock); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcomps, "comps"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
						continue
					}
					πTemp007 = µchunk
					if πE = πg.SetItem(πF, µcomps, πTemp007, πTemp006); πE != nil {
						continue
					}
					// line 403: block += 1
					πF.SetLineno(403)
					if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IAdd(πF, µblock, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µblock = πTemp006
					// line 404: data += chunk
					πF.SetLineno(404)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IAdd(πF, µdata, µchunk); πE != nil {
						continue
					}
					µdata = πTemp006
					goto Label80
				Label80:
					continue
				Label76:
					if πE != nil || πR != nil {
						continue
					}
				Label77:
					// line 405: data = _mk_bitmap(data)
					πF.SetLineno(405)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001[0] = µdata
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_mk_bitmap); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdata = πTemp006
					// line 406: data[0:0] = [block] + _bytes_to_codes(mapping)
					πF.SetLineno(406)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
						continue
					}
					πTemp001[0] = µblock
					πTemp006 = πg.NewList(πTemp001...).ToObject()
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmapping, "mapping"); πE != nil {
						continue
					}
					πTemp001[0] = µmapping
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_bytes_to_codes); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.Add(πF, πTemp006, πTemp010); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), πg.NewInt(0).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.SetItem(πF, µdata, πTemp007, πTemp006); πE != nil {
						continue
					}
					// line 407: out.append((BIGCHARSET, data))
					πF.SetLineno(407)
					πTemp001 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveGlobal(πF, ßBIGCHARSET); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp006, µdata).ToObject()
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µout, ßappend, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 408: out += tail
					πF.SetLineno(408)
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtail, "tail"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, µout, µtail); πE != nil {
						continue
					}
					µout = πTemp003
					// line 409: return out
					πF.SetLineno(409)
					if πE = πg.CheckLocal(πF, µout, "out"); πE != nil {
						continue
					}
					πR = µout
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_optimize_charset.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 411: def _fixup_range(lo, hi, ranges, fixup):
			πF.SetLineno(411)
			πTemp011 = make([]πg.Param, 4)
			πTemp011[0] = πg.Param{Name: "lo", Def: nil}
			πTemp011[1] = πg.Param{Name: "hi", Def: nil}
			πTemp011[2] = πg.Param{Name: "ranges", Def: nil}
			πTemp011[3] = πg.Param{Name: "fixup", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_fixup_range", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlo *πg.Object = πArgs[0]; _ = µlo
				var µhi *πg.Object = πArgs[1]; _ = µhi
				var µranges *πg.Object = πArgs[2]; _ = µranges
				var µfixup *πg.Object = πArgs[3]; _ = µfixup
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µk *πg.Object = πg.UnboundLocal; _ = µk
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
				var πTemp009 bool
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
					case 4: goto Label4
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfixup, "fixup"); πE != nil {
						continue
					}
					πTemp002[0] = µfixup
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp003[0] = µlo
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µhi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002[1] = πTemp005
					if πTemp004, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
					πF.PushCheckpoint(1)            
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µranges, "ranges"); πE != nil {
						continue
					}
					πTemp002[0] = µranges
					if πTemp005, πE = πg.ResolveGlobal(πF, ßenumerate); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp004, πE = πg.Iter(πF, πTemp008); πE != nil {
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
					if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp010}, πg.TieTarget{Target: &πTemp011}}}}}, πTemp005); πE != nil {
							continue
						}
						µk = πTemp008
						µlo = πTemp010
						µhi = πTemp011
					}
					if πE != nil || !πTemp009 {
						continue
					}
					πF.PushCheckpoint(4)            
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LT(πF, µi, µlo); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label7
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GT(πF, µi, µhi); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label8
					}
					goto Label9
					// line 414: if i < lo:
					πF.SetLineno(414)
				Label7:
					if πTemp008, πE = πg.ResolveGlobal(πF, ßl); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Sub(πF, µlo, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label11
					}
					goto Label12
					// line 415: if l == lo - 1:
					πF.SetLineno(415)
				Label11:
					// line 416: ranges[k] = (i, hi)
					πF.SetLineno(416)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µi, µhi).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µranges, "ranges"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp010 = µk
					if πE = πg.SetItem(πF, µranges, πTemp010, πTemp008); πE != nil {
						continue
					}
					goto Label13
				Label12:
					// line 418: ranges.insert(k, (i, i))
					πF.SetLineno(418)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp002[0] = µk
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µi, µi).ToObject()
					πTemp002[1] = πTemp005
					if πE = πg.CheckLocal(πF, µranges, "ranges"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µranges, ßinsert, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					goto Label13
				Label13:
					// line 419: break
					πF.SetLineno(419)
					πTemp007 = true
					continue
					goto Label10
					// line 420: elif i > hi:
					πF.SetLineno(420)
				Label8:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, µhi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µi, πTemp008); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
						continue
					}
					if πTemp009 {
						goto Label14
					}
					goto Label15
					// line 421: if i == hi + 1:
					πF.SetLineno(421)
				Label14:
					// line 422: ranges[k] = (lo, i)
					πF.SetLineno(422)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µlo, µi).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µranges, "ranges"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					πTemp010 = µk
					if πE = πg.SetItem(πF, µranges, πTemp010, πTemp008); πE != nil {
						continue
					}
					// line 423: break
					πF.SetLineno(423)
					πTemp007 = true
					continue
					goto Label15
				Label15:
					goto Label10
				Label9:
					// line 425: break
					πF.SetLineno(425)
					πTemp007 = true
					continue
					goto Label10
				Label10:
					continue
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					// line 427: ranges.append((i, i))
					πF.SetLineno(427)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(µi, µi).ToObject()
					πTemp002[0] = πTemp005
					if πE = πg.CheckLocal(πF, µranges, "ranges"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µranges, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
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
			if πE = πF.Globals().SetItem(πF, ß_fixup_range.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 429: _CODEBITS = _sre.CODESIZE * 8
			πF.SetLineno(429)
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_sre); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßCODESIZE, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Mul(πF, πTemp013, πg.NewInt(8).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_CODEBITS.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 430: _BITS_TRANS = b'0' + b'1' * 255
			πF.SetLineno(430)
			if πTemp012, πE = πg.Mul(πF, ß1.ToObject(), πg.NewInt(255).ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.Add(πF, ß0.ToObject(), πTemp012); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_BITS_TRANS.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 436: def _mk_bitmap(bits):
			πF.SetLineno(436)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "bits", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("_mk_bitmap", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µbits *πg.Object = πArgs[0]; _ = µbits
				var µdata *πg.Object = πg.UnboundLocal; _ = µdata
				var µdataappend *πg.Object = πg.UnboundLocal; _ = µdataappend
				var µstart *πg.Object = πg.UnboundLocal; _ = µstart
				var µm *πg.Object = πg.UnboundLocal; _ = µm
				var µv *πg.Object = πg.UnboundLocal; _ = µv
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 437: data = []
					πF.SetLineno(437)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µdata = πTemp002
					// line 438: dataappend = data.append
					πF.SetLineno(438)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µdata, ßappend, nil); πE != nil {
						continue
					}
					µdataappend = πTemp002
					// line 443: start = (1, 0)
					πF.SetLineno(443)
					πTemp002 = πg.NewTuple2(πg.NewInt(1).ToObject(), πg.NewInt(0).ToObject()).ToObject()
					µstart = πTemp002
					// line 444: m, v = start
					πF.SetLineno(444)
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, µstart); πE != nil {
						continue
					}
					µm = πTemp002
					µv = πTemp003
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, µbits); πE != nil {
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
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µc); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label4
					}
					goto Label5
					// line 446: if c:
					πF.SetLineno(446)
				Label4:
					// line 447: v = v + m
					πF.SetLineno(447)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µv, µm); πE != nil {
						continue
					}
					µv = πTemp003
					goto Label5
				Label5:
					// line 448: m = m + m
					πF.SetLineno(448)
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µm, µm); πE != nil {
						continue
					}
					µm = πTemp003
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßMAXCODE); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µm, πTemp006); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 449: if m > MAXCODE:
					πF.SetLineno(449)
				Label6:
					// line 450: dataappend(v)
					πF.SetLineno(450)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					πTemp001[0] = µv
					if πE = πg.CheckLocal(πF, µdataappend, "dataappend"); πE != nil {
						continue
					}
					if πTemp003, πE = µdataappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 451: m, v = start
					πF.SetLineno(451)
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, µstart); πE != nil {
						continue
					}
					µm = πTemp003
					µv = πTemp006
					goto Label7
				Label7:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 452: return data
					πF.SetLineno(452)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πR = µdata
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_mk_bitmap.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 454: def _bytes_to_codes(b):
			πF.SetLineno(454)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "b", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("_bytes_to_codes", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µb *πg.Object = πArgs[0]; _ = µb
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
					// line 455: return b[:]
					πF.SetLineno(455)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µb, πTemp001); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_bytes_to_codes.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 467: def _simple(av):
			πF.SetLineno(467)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "av", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_simple", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µav *πg.Object = πArgs[0]; _ = µav
				var µlo *πg.Object = πg.UnboundLocal; _ = µlo
				var µhi *πg.Object = πg.UnboundLocal; _ = µhi
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
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
					default: panic("unexpected function state")
					}
					// line 469: lo, hi = av[2].getwidth()
					πF.SetLineno(469)
					πTemp001 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µav, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßgetwidth, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µlo = πTemp001
					µhi = πTemp003
					// line 470: return lo == hi == 1 and av[2][0][0] != SUBPATTERN
					πF.SetLineno(470)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µlo, µhi); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label2
					}
					if πTemp002, πE = πg.Eq(πF, µhi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
				Label2:
					πTemp001 = πTemp002
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					πTemp003 = πg.NewInt(0).ToObject()
					πTemp007 = πg.NewInt(0).ToObject()
					πTemp009 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µav, πTemp009); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, πTemp010, πTemp007); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSUBPATTERN); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp006, πTemp003); πE != nil {
						continue
					}
					πTemp001 = πTemp002
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
			if πE = πF.Globals().SetItem(πF, ß_simple.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 472: def _compile_info(code, pattern, flags):
			πF.SetLineno(472)
			πTemp011 = make([]πg.Param, 3)
			πTemp011[0] = πg.Param{Name: "code", Def: nil}
			πTemp011[1] = πg.Param{Name: "pattern", Def: nil}
			πTemp011[2] = πg.Param{Name: "flags", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("_compile_info", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µcode *πg.Object = πArgs[0]; _ = µcode
				var µpattern *πg.Object = πArgs[1]; _ = µpattern
				var µflags *πg.Object = πArgs[2]; _ = µflags
				var µlo *πg.Object = πg.UnboundLocal; _ = µlo
				var µhi *πg.Object = πg.UnboundLocal; _ = µhi
				var µprefix *πg.Object = πg.UnboundLocal; _ = µprefix
				var µprefixappend *πg.Object = πg.UnboundLocal; _ = µprefixappend
				var µprefix_skip *πg.Object = πg.UnboundLocal; _ = µprefix_skip
				var µcharset *πg.Object = πg.UnboundLocal; _ = µcharset
				var µcharsetappend *πg.Object = πg.UnboundLocal; _ = µcharsetappend
				var µop *πg.Object = πg.UnboundLocal; _ = µop
				var µav *πg.Object = πg.UnboundLocal; _ = µav
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µcappend *πg.Object = πg.UnboundLocal; _ = µcappend
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µemit *πg.Object = πg.UnboundLocal; _ = µemit
				var µskip *πg.Object = πg.UnboundLocal; _ = µskip
				var µmask *πg.Object = πg.UnboundLocal; _ = µmask
				var µtable *πg.Object = πg.UnboundLocal; _ = µtable
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πTemp012 bool
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 64: goto Label64
					case 37: goto Label37
					case 38: goto Label38
					case 61: goto Label61
					case 6: goto Label6
					case 63: goto Label63
					case 60: goto Label60
					case 29: goto Label29
					case 30: goto Label30
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 476: lo, hi = pattern.getwidth()
					πF.SetLineno(476)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpattern, ßgetwidth, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µlo = πTemp001
					µhi = πTemp003
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µlo, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 477: if lo == 0:
					πF.SetLineno(477)
				Label1:
					// line 478: return # not worth it
					πF.SetLineno(478)
					πR = πg.None
					continue
					goto Label2
				Label2:
					// line 480: prefix = []
					πF.SetLineno(480)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					µprefix = πTemp001
					// line 481: prefixappend = prefix.append
					πF.SetLineno(481)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µprefix, ßappend, nil); πE != nil {
						continue
					}
					µprefixappend = πTemp001
					// line 482: prefix_skip = 0
					πF.SetLineno(482)
					µprefix_skip = πg.NewInt(0).ToObject()
					// line 483: charset = [] # not used
					πF.SetLineno(483)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					µcharset = πTemp001
					// line 484: charsetappend = charset.append
					πF.SetLineno(484)
					if πE = πg.CheckLocal(πF, µcharset, "charset"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcharset, ßappend, nil); πE != nil {
						continue
					}
					µcharsetappend = πTemp001
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_IGNORECASE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.And(πF, µflags, πTemp003); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 485: if not (flags & SRE_FLAG_IGNORECASE):
					πF.SetLineno(485)
				Label3:
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpattern, ßdata, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
						continue
					}
					πF.PushCheckpoint(6)
					πTemp004 = false
				Label5:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
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
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
							continue
						}
						µop = πTemp003
						µav = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(5)            
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µop == πTemp003).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveGlobal(πF, ßSUBPATTERN); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µop == πTemp007).ToObject()
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label9
					}
					πTemp005 = πF.MakeArgs(1)
					πTemp007 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µav, πTemp007); πE != nil {
						continue
					}
					πTemp005[0] = πTemp008
					if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.Eq(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label9:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label10
					}
					goto Label11
					// line 488: if op is LITERAL:
					πF.SetLineno(488)
				Label8:
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp005[0] = µprefix
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µprefix_skip, "prefix_skip"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp007, µprefix_skip); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					goto Label14
					// line 489: if len(prefix) == prefix_skip:
					πF.SetLineno(489)
				Label13:
					// line 490: prefix_skip = prefix_skip + 1
					πF.SetLineno(490)
					if πE = πg.CheckLocal(πF, µprefix_skip, "prefix_skip"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µprefix_skip, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µprefix_skip = πTemp002
					goto Label14
				Label14:
					// line 491: prefixappend(av)
					πF.SetLineno(491)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp005[0] = µav
					if πE = πg.CheckLocal(πF, µprefixappend, "prefixappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µprefixappend.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label12
					// line 492: elif op is SUBPATTERN and len(av[1]) == 1:
					πF.SetLineno(492)
				Label10:
					// line 493: op, av = av[1][0]
					πF.SetLineno(493)
					πTemp002 = πg.NewInt(0).ToObject()
					πTemp007 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µav, πTemp007); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					µop = πTemp002
					µav = πTemp007
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µop == πTemp003).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label15
					}
					goto Label16
					// line 494: if op is LITERAL:
					πF.SetLineno(494)
				Label15:
					// line 495: prefixappend(av)
					πF.SetLineno(495)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp005[0] = µav
					if πE = πg.CheckLocal(πF, µprefixappend, "prefixappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µprefixappend.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label17
				Label16:
					// line 497: break
					πF.SetLineno(497)
					πTemp004 = true
					continue
					goto Label17
				Label17:
					goto Label12
				Label11:
					// line 499: break
					πF.SetLineno(499)
					πTemp004 = true
					continue
					goto Label12
				Label12:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µprefix); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					πTemp001 = πTemp002
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label18
					}
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpattern, ßdata, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label18:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label19
					}
					goto Label20
					// line 501: if not prefix and pattern.data:
					πF.SetLineno(501)
				Label19:
					// line 502: op, av = pattern.data[0]
					πF.SetLineno(502)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µpattern, ßdata, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µop = πTemp001
					µav = πTemp003
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßSUBPATTERN); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µop == πTemp003).ToObject()
					πTemp001 = πTemp002
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label21
					}
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µav, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label21:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label22
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßBRANCH); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µop == πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label23
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßIN); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µop == πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label24
					}
					goto Label25
					// line 503: if op is SUBPATTERN and av[1]:
					πF.SetLineno(503)
				Label22:
					// line 504: op, av = av[1][0]
					πF.SetLineno(504)
					πTemp001 = πg.NewInt(0).ToObject()
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µav, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µop = πTemp001
					µav = πTemp003
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µop == πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label26
					}
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßBRANCH); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µop == πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label27
					}
					goto Label28
					// line 505: if op is LITERAL:
					πF.SetLineno(505)
				Label26:
					// line 506: charsetappend((op, av))
					πF.SetLineno(506)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µop, µav).ToObject()
					πTemp005[0] = πTemp001
					if πE = πg.CheckLocal(πF, µcharsetappend, "charsetappend"); πE != nil {
						continue
					}
					if πTemp001, πE = µcharsetappend.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label28
					// line 507: elif op is BRANCH:
					πF.SetLineno(507)
				Label27:
					// line 508: c = []
					πF.SetLineno(508)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					µc = πTemp001
					// line 509: cappend = c.append
					πF.SetLineno(509)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µc, ßappend, nil); πE != nil {
						continue
					}
					µcappend = πTemp001
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µav, πTemp002); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(30)
					πTemp004 = false
				Label29:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label31
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
						µp = πTemp002
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(29)            
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µp); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label32
					}
					goto Label33
					// line 511: if not p:
					πF.SetLineno(511)
				Label32:
					// line 512: break
					πF.SetLineno(512)
					πTemp004 = true
					continue
					goto Label33
				Label33:
					// line 513: op, av = p[0]
					πF.SetLineno(513)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µp, πTemp002); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					µop = πTemp002
					µav = πTemp007
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µop == πTemp003).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label34
					}
					goto Label35
					// line 514: if op is LITERAL:
					πF.SetLineno(514)
				Label34:
					// line 515: cappend((op, av))
					πF.SetLineno(515)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µop, µav).ToObject()
					πTemp005[0] = πTemp002
					if πE = πg.CheckLocal(πF, µcappend, "cappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µcappend.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label36
				Label35:
					// line 517: break
					πF.SetLineno(517)
					πTemp004 = true
					continue
					goto Label36
				Label36:
					continue
				Label30:
					if πE != nil || πR != nil {
						continue
					}
					// line 519: charset = c
					πF.SetLineno(519)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					µcharset = µc
				Label31:
					goto Label28
				Label28:
					goto Label25
					// line 520: elif op is BRANCH:
					πF.SetLineno(520)
				Label23:
					// line 521: c = []
					πF.SetLineno(521)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					µc = πTemp001
					// line 522: cappend = c.append
					πF.SetLineno(522)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µc, ßappend, nil); πE != nil {
						continue
					}
					µcappend = πTemp001
					πTemp002 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µav, πTemp002); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(38)
					πTemp004 = false
				Label37:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label39
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
						µp = πTemp002
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(37)            
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µp); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label40
					}
					goto Label41
					// line 524: if not p:
					πF.SetLineno(524)
				Label40:
					// line 525: break
					πF.SetLineno(525)
					πTemp004 = true
					continue
					goto Label41
				Label41:
					// line 526: op, av = p[0]
					πF.SetLineno(526)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µp, πTemp002); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
						continue
					}
					µop = πTemp002
					µav = πTemp007
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßLITERAL); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µop == πTemp003).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label42
					}
					goto Label43
					// line 527: if op is LITERAL:
					πF.SetLineno(527)
				Label42:
					// line 528: cappend((op, av))
					πF.SetLineno(528)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µop, "op"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µop, µav).ToObject()
					πTemp005[0] = πTemp002
					if πE = πg.CheckLocal(πF, µcappend, "cappend"); πE != nil {
						continue
					}
					if πTemp002, πE = µcappend.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label44
				Label43:
					// line 530: break
					πF.SetLineno(530)
					πTemp004 = true
					continue
					goto Label44
				Label44:
					continue
				Label38:
					if πE != nil || πR != nil {
						continue
					}
					// line 532: charset = c
					πF.SetLineno(532)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					µcharset = µc
				Label39:
					goto Label25
					// line 533: elif op is IN:
					πF.SetLineno(533)
				Label24:
					// line 534: charset = av
					πF.SetLineno(534)
					if πE = πg.CheckLocal(πF, µav, "av"); πE != nil {
						continue
					}
					µcharset = µav
					goto Label25
				Label25:
					goto Label20
				Label20:
					goto Label4
				Label4:
					// line 540: emit = code.append
					πF.SetLineno(540)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcode, ßappend, nil); πE != nil {
						continue
					}
					µemit = πTemp001
					// line 541: emit(OPCODES[INFO])
					πF.SetLineno(541)
					πTemp005 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßINFO); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp003, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					πTemp005[0] = πTemp002
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 542: skip = len(code); emit(0)
					πF.SetLineno(542)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp005[0] = µcode
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µskip = πTemp002
					// line 542: skip = len(code); emit(0)
					πF.SetLineno(542)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 544: mask = 0
					πF.SetLineno(544)
					µmask = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µprefix); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label45
					}
					if πE = πg.CheckLocal(πF, µcharset, "charset"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcharset); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label46
					}
					goto Label47
					// line 545: if prefix:
					πF.SetLineno(545)
				Label45:
					// line 546: mask = SRE_INFO_PREFIX
					πF.SetLineno(546)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_INFO_PREFIX); πE != nil {
						continue
					}
					µmask = πTemp001
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp005[0] = µprefix
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µprefix_skip, "prefix_skip"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, µprefix_skip); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label48
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µpattern, ßdata, nil); πE != nil {
						continue
					}
					πTemp005[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.Eq(πF, µprefix_skip, πTemp003); πE != nil {
						continue
					}
				Label48:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label49
					}
					goto Label50
					// line 547: if len(prefix) == prefix_skip == len(pattern.data):
					πF.SetLineno(547)
				Label49:
					// line 548: mask = mask + SRE_INFO_LITERAL
					πF.SetLineno(548)
					if πE = πg.CheckLocal(πF, µmask, "mask"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_INFO_LITERAL); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µmask, πTemp002); πE != nil {
						continue
					}
					µmask = πTemp001
					goto Label50
				Label50:
					goto Label47
					// line 549: elif charset:
					πF.SetLineno(549)
				Label46:
					// line 550: mask = mask + SRE_INFO_CHARSET
					πF.SetLineno(550)
					if πE = πg.CheckLocal(πF, µmask, "mask"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_INFO_CHARSET); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µmask, πTemp002); πE != nil {
						continue
					}
					µmask = πTemp001
					goto Label47
				Label47:
					// line 551: emit(mask)
					πF.SetLineno(551)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmask, "mask"); πE != nil {
						continue
					}
					πTemp005[0] = µmask
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMAXCODE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µlo, πTemp002); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label51
					}
					goto Label52
					// line 553: if lo < MAXCODE:
					πF.SetLineno(553)
				Label51:
					// line 554: emit(lo)
					πF.SetLineno(554)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					πTemp005[0] = µlo
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label53
				Label52:
					// line 556: emit(MAXCODE)
					πF.SetLineno(556)
					πTemp005 = πF.MakeArgs(1)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßMAXCODE); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 557: prefix = prefix[:MAXCODE]
					πF.SetLineno(557)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMAXCODE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µprefix, πTemp001); πE != nil {
						continue
					}
					µprefix = πTemp002
					goto Label53
				Label53:
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßMAXCODE); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µhi, πTemp002); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label54
					}
					goto Label55
					// line 558: if hi < MAXCODE:
					πF.SetLineno(558)
				Label54:
					// line 559: emit(hi)
					πF.SetLineno(559)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					πTemp005[0] = µhi
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label56
				Label55:
					// line 561: emit(0)
					πF.SetLineno(561)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label56
				Label56:
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µprefix); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label57
					}
					if πE = πg.CheckLocal(πF, µcharset, "charset"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µcharset); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label58
					}
					goto Label59
					// line 563: if prefix:
					πF.SetLineno(563)
				Label57:
					// line 564: emit(len(prefix)) # length
					πF.SetLineno(564)
					πTemp005 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp009[0] = µprefix
					if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp005[0] = πTemp002
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 565: emit(prefix_skip) # skip
					πF.SetLineno(565)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprefix_skip, "prefix_skip"); πE != nil {
						continue
					}
					πTemp005[0] = µprefix_skip
					if πE = πg.CheckLocal(πF, µemit, "emit"); πE != nil {
						continue
					}
					if πTemp001, πE = µemit.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 567: code += (prefix)
					πF.SetLineno(567)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µcode, µprefix); πE != nil {
						continue
					}
					µcode = πTemp001
					// line 569: table = [-1] + ([0]*len(prefix))
					πF.SetLineno(569)
					πTemp005 = make([]*πg.Object, 1)
					if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005[0] = πTemp002
					πTemp002 = πg.NewList(πTemp005...).ToObject()
					πTemp005 = make([]*πg.Object, 1)
					πTemp005[0] = πg.NewInt(0).ToObject()
					πTemp007 = πg.NewList(πTemp005...).ToObject()
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp005[0] = µprefix
					if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp003, πE = πg.Mul(πF, πTemp007, πTemp010); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, πTemp003); πE != nil {
						continue
					}
					µtable = πTemp001
					πTemp005 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					πTemp009[0] = µprefix
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp005[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
						continue
					}
					πF.PushCheckpoint(61)
					πTemp004 = false
				Label60:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp004 {
						πF.PopCheckpoint()
						goto Label62
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
						µi = πTemp002
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(60)            
					// line 571: table[i+1] = table[i]+1
					πF.SetLineno(571)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp003 = µi
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µtable, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp007 = πTemp008
					if πE = πg.SetItem(πF, µtable, πTemp007, πTemp003); πE != nil {
						continue
					}
					// line 572: while table[i+1] > 0 and prefix[i] != prefix[table[i+1]-1]:
					πF.SetLineno(572)
					πF.PushCheckpoint(64)
					πTemp006 = false
				Label63:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label65
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp007 = πTemp008
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µtable, πTemp007); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, πTemp008, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp012, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp012 {
						goto Label66
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp007 = µi
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µprefix, πTemp007); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp013 = πTemp014
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetItem(πF, µtable, πTemp013); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Sub(πF, πTemp014, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp007 = πTemp010
					if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µprefix, πTemp007); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp008, πTemp010); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label66:
					if πTemp011, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(63)            
					// line 573: table[i+1] = table[table[i+1]-1]+1
					πF.SetLineno(573)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp008 = πTemp010
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µtable, πTemp008); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Sub(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp007
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µtable, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp007 = πTemp008
					if πE = πg.SetItem(πF, µtable, πTemp007, πTemp003); πE != nil {
						continue
					}
					continue
				Label64:
					if πE != nil || πR != nil {
						continue
					}
				Label65:
					continue
				Label61:
					if πE != nil || πR != nil {
						continue
					}
				Label62:
					// line 575: code += (table[1:]) # don't store first entry
					πF.SetLineno(575)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µtable, πTemp001); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µcode, πTemp002); πE != nil {
						continue
					}
					µcode = πTemp001
					goto Label59
					// line 576: elif charset:
					πF.SetLineno(576)
				Label58:
					// line 577: _compile_charset(charset, flags, code)
					πF.SetLineno(577)
					πTemp005 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcharset, "charset"); πE != nil {
						continue
					}
					πTemp005[0] = µcharset
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp005[1] = µflags
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp005[2] = µcode
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_compile_charset); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label59
				Label59:
					// line 578: code[skip] = len(code) - skip
					πF.SetLineno(578)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp005[0] = µcode
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp003, µskip); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
						continue
					}
					πTemp003 = µskip
					if πE = πg.SetItem(πF, µcode, πTemp003, πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_compile_info.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 580: try:
			πF.SetLineno(580)
			πF.PushCheckpoint(8)
			// line 581: unicode
			πF.SetLineno(581)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			// line 585: STRING_TYPES = (type(""), type(unicode("")))
			πF.SetLineno(585)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ß.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp002 = πF.MakeArgs(1)
			πTemp010 = πF.MakeArgs(1)
			πTemp010[0] = ß.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp010, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp010)
			πTemp002[0] = πTemp018
			if πTemp016, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp015 = πg.NewTuple2(πTemp017, πTemp018).ToObject()
			if πE = πF.Globals().SetItem(πF, ßSTRING_TYPES.ToObject(), πTemp015); πE != nil {
				continue
			}
			goto Label7
		Label8:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp019, πTemp020 = πF.ExcInfo()
			if πTemp015, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp015); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label9
			}
			πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
			continue
			// line 582: except NameError:
			πF.SetLineno(582)
		Label9:
			// line 583: STRING_TYPES = (type(""),)
			πF.SetLineno(583)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = ß.ToObject()
			if πTemp016, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			πTemp015 = πg.NewTuple1(πTemp017).ToObject()
			if πE = πF.Globals().SetItem(πF, ßSTRING_TYPES.ToObject(), πTemp015); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label7
		Label7:
			// line 587: def isstring(obj):
			πF.SetLineno(587)
			πTemp011 = make([]πg.Param, 1)
			πTemp011[0] = πg.Param{Name: "obj", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("isstring", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µtp *πg.Object = πg.UnboundLocal; _ = µtp
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
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSTRING_TYPES); πE != nil {
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
						µtp = πTemp002
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(1)            
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp005[0] = µobj
					if πE = πg.CheckLocal(πF, µtp, "tp"); πE != nil {
						continue
					}
					πTemp005[1] = µtp
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 589: if isinstance(obj, tp):
					πF.SetLineno(589)
				Label4:
					// line 590: return 1
					πF.SetLineno(590)
					πR = πg.NewInt(1).ToObject()
					continue
					goto Label5
				Label5:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 591: return 0
					πF.SetLineno(591)
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
			if πE = πF.Globals().SetItem(πF, ßisstring.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 593: def _code(p, flags):
			πF.SetLineno(593)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "p", Def: nil}
			πTemp011[1] = πg.Param{Name: "flags", Def: nil}
			πTemp016 = πg.NewFunction(πg.NewCode("_code", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µp *πg.Object = πArgs[0]; _ = µp
				var µflags *πg.Object = πArgs[1]; _ = µflags
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var πTemp001 *πg.Object
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
					// line 595: flags = p.pattern.flags | flags
					πF.SetLineno(595)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßpattern, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßflags, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Or(πF, πTemp003, µflags); πE != nil {
						continue
					}
					µflags = πTemp001
					// line 596: code = []
					πF.SetLineno(596)
					πTemp004 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp004...).ToObject()
					µcode = πTemp001
					// line 599: _compile_info(code, p, flags)
					πF.SetLineno(599)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp004[0] = µcode
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp004[1] = µp
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp004[2] = µflags
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_compile_info); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 602: _compile(code, p.data, flags)
					πF.SetLineno(602)
					πTemp004 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp004[0] = µcode
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µp, ßdata, nil); πE != nil {
						continue
					}
					πTemp004[1] = πTemp001
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp004[2] = µflags
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_compile); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 604: code.append(OPCODES[SUCCESS])
					πF.SetLineno(604)
					πTemp004 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSUCCESS); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp003, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					πTemp004[0] = πTemp002
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µcode, ßappend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 606: return code
					πF.SetLineno(606)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πR = µcode
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_code.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 608: def compile(p, flags=0):
			πF.SetLineno(608)
			πTemp011 = make([]πg.Param, 2)
			πTemp011[0] = πg.Param{Name: "p", Def: nil}
			πTemp011[1] = πg.Param{Name: "flags", Def: πg.NewInt(0).ToObject()}
			πTemp017 = πg.NewFunction(πg.NewCode("compile", "/usr/lib/python2.7/sre_compile.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µp *πg.Object = πArgs[0]; _ = µp
				var µflags *πg.Object = πArgs[1]; _ = µflags
				var µpattern *πg.Object = πg.UnboundLocal; _ = µpattern
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
				var µgroupindex *πg.Object = πg.UnboundLocal; _ = µgroupindex
				var µindexgroup *πg.Object = πg.UnboundLocal; _ = µindexgroup
				var µk *πg.Object = πg.UnboundLocal; _ = µk
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
					case 6: goto Label6
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp001[0] = µp
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisstring); πE != nil {
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
					// line 611: if isstring(p):
					πF.SetLineno(611)
				Label1:
					// line 612: pattern = p
					πF.SetLineno(612)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					µpattern = µp
					// line 613: p = sre_parse.parse(p, flags)
					πF.SetLineno(613)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp001[0] = µp
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp001[1] = µflags
					if πTemp002, πE = πg.ResolveGlobal(πF, ßsre_parse); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßparse, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µp = πTemp002
					goto Label3
				Label2:
					// line 615: pattern = None
					πF.SetLineno(615)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					µpattern = πTemp002
					goto Label3
				Label3:
					// line 617: code = _code(p, flags)
					πF.SetLineno(617)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp001[0] = µp
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp001[1] = µflags
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_code); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp003
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßpattern, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßgroups, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, πTemp005, πg.NewInt(100).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label4
					}
					goto Label5
					// line 622: if p.pattern.groups > 100:
					πF.SetLineno(622)
				Label4:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("sorry, but this version only supports 100 named groups").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßAssertionError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 623: raise AssertionError(
					πF.SetLineno(623)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label5
				Label5:
					// line 628: groupindex = p.pattern.groupdict
					πF.SetLineno(628)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µp, ßpattern, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßgroupdict, nil); πE != nil {
						continue
					}
					µgroupindex = πTemp003
					// line 629: indexgroup = [None] * p.pattern.groups
					πF.SetLineno(629)
					πTemp001 = make([]*πg.Object, 1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp003 = πg.NewList(πTemp001...).ToObject()
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µp, ßpattern, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßgroups, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Mul(πF, πTemp003, πTemp006); πE != nil {
						continue
					}
					µindexgroup = πTemp002
					if πE = πg.CheckLocal(πF, µgroupindex, "groupindex"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µgroupindex, ßitems, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
							continue
						}
						µk = πTemp005
						µi = πTemp006
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(6)            
					// line 631: indexgroup[i] = k
					πF.SetLineno(631)
					if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µk); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindexgroup, "indexgroup"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = µi
					if πE = πg.SetItem(πF, µindexgroup, πTemp005, πTemp003); πE != nil {
						continue
					}
					continue
				Label7:
					if πE != nil || πR != nil {
						continue
					}
				Label8:
					// line 633: return _sre.compile(
					πF.SetLineno(633)
					πTemp001 = πF.MakeArgs(6)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßpattern, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßflags, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, µflags, πTemp005); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[2] = µcode
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µp, ßpattern, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßgroups, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001[3] = πTemp002
					if πE = πg.CheckLocal(πF, µgroupindex, "groupindex"); πE != nil {
						continue
					}
					πTemp001[4] = µgroupindex
					if πE = πg.CheckLocal(πF, µindexgroup, "indexgroup"); πE != nil {
						continue
					}
					πTemp001[5] = µindexgroup
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_sre); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcompile, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcompile.ToObject(), πTemp017); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("sre_compile", Code)
}

