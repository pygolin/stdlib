package _sre
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/operator"
	// _ "github.com/pygolin/stdlib/sre_constants"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßATCODES := πg.InternStr("ATCODES")
		ßCHCODES := πg.InternStr("CHCODES")
		ßCODESIZE := πg.InternStr("CODESIZE")
		ßDISPATCH_TABLE := πg.InternStr("DISPATCH_TABLE")
		ßFalse := πg.InternStr("False")
		ßIndexError := πg.InternStr("IndexError")
		ßMAGIC := πg.InternStr("MAGIC")
		ßMAXREPEAT := πg.InternStr("MAXREPEAT")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßOPCODES := πg.InternStr("OPCODES")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßSRE_FLAG_LOCALE := πg.InternStr("SRE_FLAG_LOCALE")
		ßSRE_FLAG_UNICODE := πg.InternStr("SRE_FLAG_UNICODE")
		ßSRE_INFO_LITERAL := πg.InternStr("SRE_INFO_LITERAL")
		ßSRE_INFO_PREFIX := πg.InternStr("SRE_INFO_PREFIX")
		ßSRE_Match := πg.InternStr("SRE_Match")
		ßSRE_Pattern := πg.InternStr("SRE_Pattern")
		ßSRE_Scanner := πg.InternStr("SRE_Scanner")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ß_ := πg.InternStr("_")
		ß_AtcodeDispatcher := πg.InternStr("_AtcodeDispatcher")
		ß_CharsetDispatcher := πg.InternStr("_CharsetDispatcher")
		ß_ChcodeDispatcher := πg.InternStr("_ChcodeDispatcher")
		ß_Dispatcher := πg.InternStr("_Dispatcher")
		ß_MatchContext := πg.InternStr("_MatchContext")
		ß_OpcodeDispatcher := πg.InternStr("_OpcodeDispatcher")
		ß_RepeatContext := πg.InternStr("_RepeatContext")
		ß_State := πg.InternStr("_State")
		ß__all__ := πg.InternStr("__all__")
		ß__class__ := πg.InternStr("__class__")
		ß__copy__ := πg.InternStr("__copy__")
		ß__deepcopy__ := πg.InternStr("__deepcopy__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_ascii_char_info := πg.InternStr("_ascii_char_info")
		ß_code := πg.InternStr("_code")
		ß_create_regs := πg.InternStr("_create_regs")
		ß_get_index := πg.InternStr("_get_index")
		ß_get_slice := πg.InternStr("_get_slice")
		ß_indexgroup := πg.InternStr("_indexgroup")
		ß_is_digit := πg.InternStr("_is_digit")
		ß_is_linebreak := πg.InternStr("_is_linebreak")
		ß_is_loc_word := πg.InternStr("_is_loc_word")
		ß_is_space := πg.InternStr("_is_space")
		ß_is_uni_word := πg.InternStr("_is_uni_word")
		ß_is_word := πg.InternStr("_is_word")
		ß_log := πg.InternStr("_log")
		ß_match_search := πg.InternStr("_match_search")
		ß_state := πg.InternStr("_state")
		ß_subx := πg.InternStr("_subx")
		ß_uni_linebreaks := πg.InternStr("_uni_linebreaks")
		ßappend := πg.InternStr("append")
		ßat_beginning := πg.InternStr("at_beginning")
		ßat_beginning_line := πg.InternStr("at_beginning_line")
		ßat_beginning_string := πg.InternStr("at_beginning_string")
		ßat_boundary := πg.InternStr("at_boundary")
		ßat_dispatcher := πg.InternStr("at_dispatcher")
		ßat_end := πg.InternStr("at_end")
		ßat_end_line := πg.InternStr("at_end_line")
		ßat_end_string := πg.InternStr("at_end_string")
		ßat_linebreak := πg.InternStr("at_linebreak")
		ßat_loc_boundary := πg.InternStr("at_loc_boundary")
		ßat_loc_non_boundary := πg.InternStr("at_loc_non_boundary")
		ßat_non_boundary := πg.InternStr("at_non_boundary")
		ßat_uni_boundary := πg.InternStr("at_uni_boundary")
		ßat_uni_non_boundary := πg.InternStr("at_uni_non_boundary")
		ßbuild_dispatch_table := πg.InternStr("build_dispatch_table")
		ßcallable := πg.InternStr("callable")
		ßcategory_digit := πg.InternStr("category_digit")
		ßcategory_linebreak := πg.InternStr("category_linebreak")
		ßcategory_loc_not_word := πg.InternStr("category_loc_not_word")
		ßcategory_loc_word := πg.InternStr("category_loc_word")
		ßcategory_not_digit := πg.InternStr("category_not_digit")
		ßcategory_not_linebreak := πg.InternStr("category_not_linebreak")
		ßcategory_not_space := πg.InternStr("category_not_space")
		ßcategory_not_word := πg.InternStr("category_not_word")
		ßcategory_space := πg.InternStr("category_space")
		ßcategory_uni_digit := πg.InternStr("category_uni_digit")
		ßcategory_uni_linebreak := πg.InternStr("category_uni_linebreak")
		ßcategory_uni_not_digit := πg.InternStr("category_uni_not_digit")
		ßcategory_uni_not_linebreak := πg.InternStr("category_uni_not_linebreak")
		ßcategory_uni_not_space := πg.InternStr("category_uni_not_space")
		ßcategory_uni_not_word := πg.InternStr("category_uni_not_word")
		ßcategory_uni_space := πg.InternStr("category_uni_space")
		ßcategory_uni_word := πg.InternStr("category_uni_word")
		ßcategory_word := πg.InternStr("category_word")
		ßch_dispatcher := πg.InternStr("ch_dispatcher")
		ßchar := πg.InternStr("char")
		ßcheck_charset := πg.InternStr("check_charset")
		ßchr := πg.InternStr("chr")
		ßclassmethod := πg.InternStr("classmethod")
		ßcode_position := πg.InternStr("code_position")
		ßcompile := πg.InternStr("compile")
		ßcontext_stack := πg.InternStr("context_stack")
		ßcopyright := πg.InternStr("copyright")
		ßcount := πg.InternStr("count")
		ßcount_repetitions := πg.InternStr("count_repetitions")
		ßdispatch := πg.InternStr("dispatch")
		ßend := πg.InternStr("end")
		ßendpos := πg.InternStr("endpos")
		ßeq := πg.InternStr("eq")
		ßexecuting_contexts := πg.InternStr("executing_contexts")
		ßexpand := πg.InternStr("expand")
		ßfast_search := πg.InternStr("fast_search")
		ßfindall := πg.InternStr("findall")
		ßfinditer := πg.InternStr("finditer")
		ßflags := πg.InternStr("flags")
		ßgeneral_op_groupref := πg.InternStr("general_op_groupref")
		ßgeneral_op_in := πg.InternStr("general_op_in")
		ßgeneral_op_literal := πg.InternStr("general_op_literal")
		ßget := πg.InternStr("get")
		ßget_marks := πg.InternStr("get_marks")
		ßgetattr := πg.InternStr("getattr")
		ßgetcodesize := πg.InternStr("getcodesize")
		ßgetlower := πg.InternStr("getlower")
		ßglobals := πg.InternStr("globals")
		ßgroup := πg.InternStr("group")
		ßgroupdict := πg.InternStr("groupdict")
		ßgroupindex := πg.InternStr("groupindex")
		ßgroups := πg.InternStr("groups")
		ßhas_matched := πg.InternStr("has_matched")
		ßhasattr := πg.InternStr("hasattr")
		ßid := πg.InternStr("id")
		ßinfo := πg.InternStr("info")
		ßint := πg.InternStr("int")
		ßisalnum := πg.InternStr("isalnum")
		ßisdigit := πg.InternStr("isdigit")
		ßisinstance := πg.InternStr("isinstance")
		ßisspace := πg.InternStr("isspace")
		ßitems := πg.InternStr("items")
		ßiter := πg.InternStr("iter")
		ßjoin := πg.InternStr("join")
		ßlast_position := πg.InternStr("last_position")
		ßlastgroup := πg.InternStr("lastgroup")
		ßlastindex := πg.InternStr("lastindex")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßliteral := πg.InternStr("literal")
		ßlower := πg.InternStr("lower")
		ßmarks := πg.InternStr("marks")
		ßmarks_pop := πg.InternStr("marks_pop")
		ßmarks_pop_discard := πg.InternStr("marks_pop_discard")
		ßmarks_pop_keep := πg.InternStr("marks_pop_keep")
		ßmarks_push := πg.InternStr("marks_push")
		ßmarks_stack := πg.InternStr("marks_stack")
		ßmatch := πg.InternStr("match")
		ßmaxint := πg.InternStr("maxint")
		ßname := πg.InternStr("name")
		ßne := πg.InternStr("ne")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßok := πg.InternStr("ok")
		ßop_ := πg.InternStr("op_")
		ßop_any := πg.InternStr("op_any")
		ßop_any_all := πg.InternStr("op_any_all")
		ßop_assert := πg.InternStr("op_assert")
		ßop_assert_not := πg.InternStr("op_assert_not")
		ßop_at := πg.InternStr("op_at")
		ßop_branch := πg.InternStr("op_branch")
		ßop_category := πg.InternStr("op_category")
		ßop_failure := πg.InternStr("op_failure")
		ßop_groupref := πg.InternStr("op_groupref")
		ßop_groupref_exists := πg.InternStr("op_groupref_exists")
		ßop_groupref_ignore := πg.InternStr("op_groupref_ignore")
		ßop_in := πg.InternStr("op_in")
		ßop_in_ignore := πg.InternStr("op_in_ignore")
		ßop_info := πg.InternStr("op_info")
		ßop_jump := πg.InternStr("op_jump")
		ßop_literal := πg.InternStr("op_literal")
		ßop_literal_ignore := πg.InternStr("op_literal_ignore")
		ßop_mark := πg.InternStr("op_mark")
		ßop_max_until := πg.InternStr("op_max_until")
		ßop_min_repeat_one := πg.InternStr("op_min_repeat_one")
		ßop_min_until := πg.InternStr("op_min_until")
		ßop_not_literal := πg.InternStr("op_not_literal")
		ßop_not_literal_ignore := πg.InternStr("op_not_literal_ignore")
		ßop_repeat := πg.InternStr("op_repeat")
		ßop_repeat_one := πg.InternStr("op_repeat_one")
		ßop_success := πg.InternStr("op_success")
		ßoperator := πg.InternStr("operator")
		ßord := πg.InternStr("ord")
		ßpattern := πg.InternStr("pattern")
		ßpattern_codes := πg.InternStr("pattern_codes")
		ßpeek_char := πg.InternStr("peek_char")
		ßpeek_code := πg.InternStr("peek_code")
		ßpos := πg.InternStr("pos")
		ßprevious := πg.InternStr("previous")
		ßpush_new_context := πg.InternStr("push_new_context")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßregs := πg.InternStr("regs")
		ßremaining_chars := πg.InternStr("remaining_chars")
		ßremaining_codes := πg.InternStr("remaining_codes")
		ßrepeat := πg.InternStr("repeat")
		ßreset := πg.InternStr("reset")
		ßscanner := πg.InternStr("scanner")
		ßsearch := πg.InternStr("search")
		ßset_ := πg.InternStr("set_")
		ßset_bigcharset := πg.InternStr("set_bigcharset")
		ßset_category := πg.InternStr("set_category")
		ßset_charset := πg.InternStr("set_charset")
		ßset_dispatcher := πg.InternStr("set_dispatcher")
		ßset_failure := πg.InternStr("set_failure")
		ßset_literal := πg.InternStr("set_literal")
		ßset_mark := πg.InternStr("set_mark")
		ßset_negate := πg.InternStr("set_negate")
		ßset_range := πg.InternStr("set_range")
		ßskip_char := πg.InternStr("skip_char")
		ßskip_code := πg.InternStr("skip_code")
		ßspan := πg.InternStr("span")
		ßsplit := πg.InternStr("split")
		ßsre_constants := πg.InternStr("sre_constants")
		ßstart := πg.InternStr("start")
		ßstate := πg.InternStr("state")
		ßstring := πg.InternStr("string")
		ßstring_position := πg.InternStr("string_position")
		ßsub := πg.InternStr("sub")
		ßsubn := πg.InternStr("subn")
		ßsuccess := πg.InternStr("success")
		ßsys := πg.InternStr("sys")
		ßtuple := πg.InternStr("tuple")
		ßunichr := πg.InternStr("unichr")
		ßunknown := πg.InternStr("unknown")
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
		var πTemp010 []πg.Param
		_ = πTemp010
		var πTemp011 *πg.Dict
		_ = πTemp011
		var πTemp012 *πg.Object
		_ = πTemp012
		var πTemp013 *πg.Object
		_ = πTemp013
		var πTemp014 *πg.Object
		_ = πTemp014
		var πTemp015 *πg.Object
		_ = πTemp015
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 1: goto Label1
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 2: """
			πF.SetLineno(2)
			// line 2: """
			πF.SetLineno(2)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\nA pure Python reimplementation of the _sre module from CPython 2.4\nCopyright 2005 Nik Haldimann, licensed under the MIT license\n\nThis code is based on material licensed under CNRI's Python 1.6 license and\ncopyrighted by: Copyright (c) 1997-2001 by Secret Labs AB\n").ToObject()); πE != nil {
				continue
			}
			// line 11: import sys
			πF.SetLineno(11)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: import operator
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "operator"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßoperator.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 15: import sre_constants
			πF.SetLineno(15)
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
			// line 17: globals()[name] = getattr(sre_constants, name)
			πF.SetLineno(17)
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
			// line 21: MAGIC = 20031017
			πF.SetLineno(21)
			if πE = πF.Globals().SetItem(πF, ßMAGIC.ToObject(), πg.NewInt(20031017).ToObject()); πE != nil {
				continue
			}
			// line 32: CODESIZE = 2
			πF.SetLineno(32)
			if πE = πF.Globals().SetItem(πF, ßCODESIZE.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			// line 34: copyright = "_sre.py 2.4c Copyright 2005 by Nik Haldimann"
			πF.SetLineno(34)
			if πE = πF.Globals().SetItem(πF, ßcopyright.ToObject(), πg.NewStr("_sre.py 2.4c Copyright 2005 by Nik Haldimann").ToObject()); πE != nil {
				continue
			}
			// line 36: def getcodesize():
			πF.SetLineno(36)
			πTemp010 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("getcodesize", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 37: return CODESIZE
					πF.SetLineno(37)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßgetcodesize.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 40: def compile(pattern, flags, code, groups=0, groupindex={}, indexgroup=[None]):
			πF.SetLineno(40)
			πTemp010 = make([]πg.Param, 6)
			πTemp010[0] = πg.Param{Name: "pattern", Def: nil}
			πTemp010[1] = πg.Param{Name: "flags", Def: nil}
			πTemp010[2] = πg.Param{Name: "code", Def: nil}
			πTemp010[3] = πg.Param{Name: "groups", Def: πg.NewInt(0).ToObject()}
			πTemp011 = πg.NewDict()
			πTemp004 = πTemp011.ToObject()
			πTemp010[4] = πg.Param{Name: "groupindex", Def: πTemp004}
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp002[0] = πTemp004
			πTemp004 = πg.NewList(πTemp002...).ToObject()
			πTemp010[5] = πg.Param{Name: "indexgroup", Def: πTemp004}
			πTemp003 = πg.NewFunction(πg.NewCode("compile", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µpattern *πg.Object = πArgs[0]; _ = µpattern
				var µflags *πg.Object = πArgs[1]; _ = µflags
				var µcode *πg.Object = πArgs[2]; _ = µcode
				var µgroups *πg.Object = πArgs[3]; _ = µgroups
				var µgroupindex *πg.Object = πArgs[4]; _ = µgroupindex
				var µindexgroup *πg.Object = πArgs[5]; _ = µindexgroup
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
					// line 41: """Compiles (or rather just converts) a pattern descriptor to a SRE_Pattern
					πF.SetLineno(41)
					// line 43: return SRE_Pattern(pattern, flags, code, groups, groupindex, indexgroup)
					πF.SetLineno(43)
					πTemp001 = πF.MakeArgs(6)
					if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
						continue
					}
					πTemp001[0] = µpattern
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					πTemp001[1] = µflags
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp001[2] = µcode
					if πE = πg.CheckLocal(πF, µgroups, "groups"); πE != nil {
						continue
					}
					πTemp001[3] = µgroups
					if πE = πg.CheckLocal(πF, µgroupindex, "groupindex"); πE != nil {
						continue
					}
					πTemp001[4] = µgroupindex
					if πE = πg.CheckLocal(πF, µindexgroup, "indexgroup"); πE != nil {
						continue
					}
					πTemp001[5] = µindexgroup
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_Pattern); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßcompile.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 41: """Compiles (or rather just converts) a pattern descriptor to a SRE_Pattern
			πF.SetLineno(41)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Compiles (or rather just converts) a pattern descriptor to a SRE_Pattern\n    object. Actual compilation to opcodes happens in sre_compile.").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßcompile); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 45: def getlower(char_ord, flags):
			πF.SetLineno(45)
			πTemp010 = make([]πg.Param, 2)
			πTemp010[0] = πg.Param{Name: "char_ord", Def: nil}
			πTemp010[1] = πg.Param{Name: "flags", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("getlower", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µchar_ord *πg.Object = πArgs[0]; _ = µchar_ord
				var µflags *πg.Object = πArgs[1]; _ = µflags
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
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
					if πE = πg.CheckLocal(πF, µchar_ord, "char_ord"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µchar_ord, πg.NewInt(128).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
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
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßSRE_FLAG_LOCALE); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µflags, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µchar_ord, "char_ord"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LT(πF, µchar_ord, πg.NewInt(256).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
				Label2:
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 46: if (char_ord < 128) or (flags & SRE_FLAG_UNICODE) \
					πF.SetLineno(46)
				Label3:
					// line 49: return ord(chr(char_ord).lower())
					πF.SetLineno(49)
					πTemp007 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar_ord, "char_ord"); πE != nil {
						continue
					}
					πTemp008[0] = µchar_ord
					if πTemp001, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßlower, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp007[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					πR = πTemp003
					continue
					goto Label5
				Label4:
					// line 51: return char_ord
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µchar_ord, "char_ord"); πE != nil {
						continue
					}
					πR = µchar_ord
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
			if πE = πF.Globals().SetItem(πF, ßgetlower.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 54: class SRE_Pattern(object):
			πF.SetLineno(54)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SRE_Pattern", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				var πTemp003 *πg.Dict
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
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 *πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 56: def __init__(self, pattern, flags, code, groups=0, groupindex={}, indexgroup=[None]):
					πF.SetLineno(56)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pattern", Def: nil}
					πTemp002[2] = πg.Param{Name: "flags", Def: nil}
					πTemp002[3] = πg.Param{Name: "code", Def: nil}
					πTemp002[4] = πg.Param{Name: "groups", Def: πg.NewInt(0).ToObject()}
					πTemp003 = πg.NewDict()
					πTemp004 = πTemp003.ToObject()
					πTemp002[5] = πg.Param{Name: "groupindex", Def: πTemp004}
					πTemp005 = make([]*πg.Object, 1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp005[0] = πTemp004
					πTemp004 = πg.NewList(πTemp005...).ToObject()
					πTemp002[6] = πg.Param{Name: "indexgroup", Def: πTemp004}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpattern *πg.Object = πArgs[1]; _ = µpattern
						var µflags *πg.Object = πArgs[2]; _ = µflags
						var µcode *πg.Object = πArgs[3]; _ = µcode
						var µgroups *πg.Object = πArgs[4]; _ = µgroups
						var µgroupindex *πg.Object = πArgs[5]; _ = µgroupindex
						var µindexgroup *πg.Object = πArgs[6]; _ = µindexgroup
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 57: self.pattern = pattern
							πF.SetLineno(57)
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
							// line 58: self.flags = flags
							πF.SetLineno(58)
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µflags); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßflags, πTemp001); πE != nil {
								continue
							}
							// line 59: self.groups = groups
							πF.SetLineno(59)
							if πE = πg.CheckLocal(πF, µgroups, "groups"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µgroups); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßgroups, πTemp001); πE != nil {
								continue
							}
							// line 60: self.groupindex = groupindex # Maps group names to group indices
							πF.SetLineno(60)
							if πE = πg.CheckLocal(πF, µgroupindex, "groupindex"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µgroupindex); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßgroupindex, πTemp001); πE != nil {
								continue
							}
							// line 61: self._indexgroup = indexgroup # Maps indices to group names
							πF.SetLineno(61)
							if πE = πg.CheckLocal(πF, µindexgroup, "indexgroup"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µindexgroup); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_indexgroup, πTemp001); πE != nil {
								continue
							}
							// line 62: self._code = code
							πF.SetLineno(62)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcode); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_code, πTemp001); πE != nil {
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
					// line 64: def match(self, string, pos=0, endpos=sys.maxint):
					πF.SetLineno(64)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "pos", Def: πg.NewInt(0).ToObject()}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßmaxint, nil); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "endpos", Def: πTemp007}
					πTemp004 = πg.NewFunction(πg.NewCode("match", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µpos *πg.Object = πArgs[2]; _ = µpos
						var µendpos *πg.Object = πArgs[3]; _ = µendpos
						var µstate *πg.Object = πg.UnboundLocal; _ = µstate
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
							// line 65: """If zero or more characters at the beginning of string match this
							πF.SetLineno(65)
							// line 68: state = _State(string, pos, endpos, self.flags)
							πF.SetLineno(68)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[0] = µstring
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							πTemp001[1] = µpos
							if πE = πg.CheckLocal(πF, µendpos, "endpos"); πE != nil {
								continue
							}
							πTemp001[2] = µendpos
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßflags, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_State); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstate = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_code, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßmatch, nil); πE != nil {
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
							// line 69: if state.match(self._code):
							πF.SetLineno(69)
						Label1:
							// line 70: return SRE_Match(self, state)
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp001[1] = µstate
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_Match); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label3
						Label2:
							// line 72: return None
							πF.SetLineno(72)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp002
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
					if πE = πClass.SetItem(πF, ßmatch.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 65: """If zero or more characters at the beginning of string match this
					πF.SetLineno(65)
					// line 65: """If zero or more characters at the beginning of string match this
					πF.SetLineno(65)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("If zero or more characters at the beginning of string match this\n        regular expression, return a corresponding MatchObject instance. Return\n        None if the string does not match the pattern.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("If zero or more characters at the beginning of string match this\n        regular expression, return a corresponding MatchObject instance. Return\n        None if the string does not match the pattern.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßmatch); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 74: def search(self, string, pos=0, endpos=sys.maxint):
					πF.SetLineno(74)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "pos", Def: πg.NewInt(0).ToObject()}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßmaxint, nil); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "endpos", Def: πTemp008}
					πTemp006 = πg.NewFunction(πg.NewCode("search", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µpos *πg.Object = πArgs[2]; _ = µpos
						var µendpos *πg.Object = πArgs[3]; _ = µendpos
						var µstate *πg.Object = πg.UnboundLocal; _ = µstate
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
							// line 75: """Scan through string looking for a location where this regular
							πF.SetLineno(75)
							// line 79: state = _State(string, pos, endpos, self.flags)
							πF.SetLineno(79)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[0] = µstring
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							πTemp001[1] = µpos
							if πE = πg.CheckLocal(πF, µendpos, "endpos"); πE != nil {
								continue
							}
							πTemp001[2] = µendpos
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßflags, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_State); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstate = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_code, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßsearch, nil); πE != nil {
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
							// line 80: if state.search(self._code):
							πF.SetLineno(80)
						Label1:
							// line 81: return SRE_Match(self, state)
							πF.SetLineno(81)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp001[1] = µstate
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_Match); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label3
						Label2:
							// line 83: return None
							πF.SetLineno(83)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp002
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
					if πE = πClass.SetItem(πF, ßsearch.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 75: """Scan through string looking for a location where this regular
					πF.SetLineno(75)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Scan through string looking for a location where this regular\n        expression produces a match, and return a corresponding MatchObject\n        instance. Return None if no position in the string matches the\n        pattern.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßsearch); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 85: def findall(self, string, pos=0, endpos=sys.maxint):
					πF.SetLineno(85)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "pos", Def: πg.NewInt(0).ToObject()}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßmaxint, nil); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "endpos", Def: πTemp009}
					πTemp007 = πg.NewFunction(πg.NewCode("findall", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µpos *πg.Object = πArgs[2]; _ = µpos
						var µendpos *πg.Object = πArgs[3]; _ = µendpos
						var µmatchlist *πg.Object = πg.UnboundLocal; _ = µmatchlist
						var µstate *πg.Object = πg.UnboundLocal; _ = µstate
						var µmatch *πg.Object = πg.UnboundLocal; _ = µmatch
						var µitem *πg.Object = πg.UnboundLocal; _ = µitem
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
							// line 86: """Return a list of all non-overlapping matches of pattern in string."""
							πF.SetLineno(86)
							// line 87: matchlist = []
							πF.SetLineno(87)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µmatchlist = πTemp002
							// line 88: state = _State(string, pos, endpos, self.flags)
							πF.SetLineno(88)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[0] = µstring
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							πTemp001[1] = µpos
							if πE = πg.CheckLocal(πF, µendpos, "endpos"); πE != nil {
								continue
							}
							πTemp001[2] = µendpos
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßflags, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_State); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstate = πTemp003
							// line 89: while state.start <= state.end:
							πF.SetLineno(89)
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
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µstate, ßend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 90: state.reset()
							πF.SetLineno(90)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßreset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 91: state.string_position = state.start
							πF.SetLineno(91)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstring_position, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_code, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßsearch, nil); πE != nil {
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
								goto Label4
							}
							goto Label5
							// line 92: if not state.search(self._code):
							πF.SetLineno(92)
						Label4:
							// line 93: break
							πF.SetLineno(93)
							πTemp004 = true
							continue
							goto Label5
						Label5:
							// line 94: match = SRE_Match(self, state)
							πF.SetLineno(94)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp001[1] = µstate
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_Match); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmatch = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßgroups, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßgroups, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label6:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 95: if self.groups == 0 or self.groups == 1:
							πF.SetLineno(95)
						Label7:
							// line 96: item = match.group(self.groups)
							πF.SetLineno(96)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgroups, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmatch, ßgroup, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µitem = πTemp003
							goto Label9
						Label8:
							// line 98: item = match.groups("")
							πF.SetLineno(98)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ß.ToObject()
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmatch, ßgroups, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µitem = πTemp003
							goto Label9
						Label9:
							// line 99: matchlist.append(item)
							πF.SetLineno(99)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µmatchlist, "matchlist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmatchlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label10
							}
							goto Label11
							// line 100: if state.string_position == state.start:
							πF.SetLineno(100)
						Label10:
							// line 101: state.start += 1
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstart, πTemp003); πE != nil {
								continue
							}
							goto Label12
						Label11:
							// line 103: state.start = state.string_position
							πF.SetLineno(103)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstart, πTemp003); πE != nil {
								continue
							}
							goto Label12
						Label12:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 104: return matchlist
							πF.SetLineno(104)
							if πE = πg.CheckLocal(πF, µmatchlist, "matchlist"); πE != nil {
								continue
							}
							πR = µmatchlist
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßfindall.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 86: """Return a list of all non-overlapping matches of pattern in string."""
					πF.SetLineno(86)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Return a list of all non-overlapping matches of pattern in string.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßfindall); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 106: def _subx(self, template, string, count=0, subn=False):
					πF.SetLineno(106)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "template", Def: nil}
					πTemp002[2] = πg.Param{Name: "string", Def: nil}
					πTemp002[3] = πg.Param{Name: "count", Def: πg.NewInt(0).ToObject()}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "subn", Def: πTemp009}
					πTemp008 = πg.NewFunction(πg.NewCode("_subx", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtemplate *πg.Object = πArgs[1]; _ = µtemplate
						var µstring *πg.Object = πArgs[2]; _ = µstring
						var µcount *πg.Object = πArgs[3]; _ = µcount
						var µsubn *πg.Object = πArgs[4]; _ = µsubn
						var µfilter *πg.Object = πg.UnboundLocal; _ = µfilter
						var µstate *πg.Object = πg.UnboundLocal; _ = µstate
						var µsublist *πg.Object = πg.UnboundLocal; _ = µsublist
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µlast_pos *πg.Object = πg.UnboundLocal; _ = µlast_pos
						var µitem *πg.Object = πg.UnboundLocal; _ = µitem
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 107: filter = template
							πF.SetLineno(107)
							if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
								continue
							}
							µfilter = µtemplate
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
								continue
							}
							πTemp004[0] = µtemplate
							if πTemp005, πE = πg.ResolveGlobal(πF, ßcallable); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µtemplate, "template"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µtemplate, πg.NewStr("\\").ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 108: if not callable(template) and "\\" in template:
							πF.SetLineno(108)
						Label2:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 110: raise NotImplementedError()
							πF.SetLineno(110)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label3
						Label3:
							// line 111: state = _State(string, 0, sys.maxint, self.flags)
							πF.SetLineno(111)
							πTemp004 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp004[0] = µstring
							πTemp004[1] = πg.NewInt(0).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmaxint, nil); πE != nil {
								continue
							}
							πTemp004[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßflags, nil); πE != nil {
								continue
							}
							πTemp004[3] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_State); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µstate = πTemp003
							// line 112: sublist = []
							πF.SetLineno(112)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µsublist = πTemp001
							// line 114: n = last_pos = 0
							πF.SetLineno(114)
							µn = πg.NewInt(0).ToObject()
							µlast_pos = πg.NewInt(0).ToObject()
							// line 115: while not count or n < count:
							πF.SetLineno(115)
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
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µcount); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp009).ToObject()
							πTemp001 = πTemp003
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µn, µcount); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label7:
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 116: state.reset()
							πF.SetLineno(116)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßreset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 117: state.string_position = state.start
							πF.SetLineno(117)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstring_position, πTemp003); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_code, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßsearch, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label8
							}
							goto Label9
							// line 118: if not state.search(self._code):
							πF.SetLineno(118)
						Label8:
							// line 119: break
							πF.SetLineno(119)
							πTemp002 = true
							continue
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µlast_pos, "last_pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µlast_pos, πTemp003); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label10
							}
							goto Label11
							// line 120: if last_pos < state.start:
							πF.SetLineno(120)
						Label10:
							// line 121: sublist.append(string[last_pos:state.start])
							πF.SetLineno(121)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlast_pos, "last_pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µlast_pos, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µstring, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µsublist, "sublist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsublist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µlast_pos, "last_pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, µlast_pos, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µlast_pos, "last_pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, µlast_pos, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp007 {
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GT(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
						Label12:
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label13
							}
							goto Label14
							// line 122: if not (last_pos == state.start and
							πF.SetLineno(122)
						Label13:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfilter, "filter"); πE != nil {
								continue
							}
							πTemp004[0] = µfilter
							if πTemp001, πE = πg.ResolveGlobal(πF, ßcallable); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label15
							}
							goto Label16
							// line 125: if callable(filter):
							πF.SetLineno(125)
						Label15:
							// line 126: sublist.append(filter(SRE_Match(self, state)))
							πF.SetLineno(126)
							πTemp004 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(1)
							πTemp011 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp011[0] = µself
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp011[1] = µstate
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_Match); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp010[0] = πTemp003
							if πE = πg.CheckLocal(πF, µfilter, "filter"); πE != nil {
								continue
							}
							if πTemp001, πE = µfilter.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µsublist, "sublist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsublist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label17
						Label16:
							// line 128: sublist.append(filter)
							πF.SetLineno(128)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfilter, "filter"); πE != nil {
								continue
							}
							πTemp004[0] = µfilter
							if πE = πg.CheckLocal(πF, µsublist, "sublist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsublist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label17
						Label17:
							// line 129: last_pos = state.string_position
							πF.SetLineno(129)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							µlast_pos = πTemp001
							// line 130: n += 1
							πF.SetLineno(130)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µn = πTemp001
							goto Label14
						Label14:
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label18
							}
							goto Label19
							// line 131: if state.string_position == state.start:
							πF.SetLineno(131)
						Label18:
							// line 132: state.start += 1
							πF.SetLineno(132)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstart, πTemp003); πE != nil {
								continue
							}
							goto Label20
						Label19:
							// line 134: state.start = state.string_position
							πF.SetLineno(134)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstart, πTemp003); πE != nil {
								continue
							}
							goto Label20
						Label20:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							if πE = πg.CheckLocal(πF, µlast_pos, "last_pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µlast_pos, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label21
							}
							goto Label22
							// line 136: if last_pos < state.end:
							πF.SetLineno(136)
						Label21:
							// line 137: sublist.append(string[last_pos:state.end])
							πF.SetLineno(137)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlast_pos, "last_pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µlast_pos, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µstring, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πE = πg.CheckLocal(πF, µsublist, "sublist"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µsublist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label22
						Label22:
							// line 138: item = "".join(sublist)
							πF.SetLineno(138)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsublist, "sublist"); πE != nil {
								continue
							}
							πTemp004[0] = µsublist
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µitem = πTemp003
							if πE = πg.CheckLocal(πF, µsubn, "subn"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µsubn); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label23
							}
							goto Label24
							// line 139: if subn:
							πF.SetLineno(139)
						Label23:
							// line 140: return item, n
							πF.SetLineno(140)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µitem, µn).ToObject()
							πR = πTemp001
							continue
							goto Label25
						Label24:
							// line 142: return item
							πF.SetLineno(142)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πR = µitem
							continue
							goto Label25
						Label25:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_subx.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 144: def sub(self, repl, string, count=0):
					πF.SetLineno(144)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "repl", Def: nil}
					πTemp002[2] = πg.Param{Name: "string", Def: nil}
					πTemp002[3] = πg.Param{Name: "count", Def: πg.NewInt(0).ToObject()}
					πTemp009 = πg.NewFunction(πg.NewCode("sub", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrepl *πg.Object = πArgs[1]; _ = µrepl
						var µstring *πg.Object = πArgs[2]; _ = µstring
						var µcount *πg.Object = πArgs[3]; _ = µcount
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
							// line 145: """Return the string obtained by replacing the leftmost non-overlapping
							πF.SetLineno(145)
							// line 147: return self._subx(repl, string, count, False)
							πF.SetLineno(147)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µrepl, "repl"); πE != nil {
								continue
							}
							πTemp001[0] = µrepl
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[1] = µstring
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							πTemp001[2] = µcount
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_subx, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsub.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 145: """Return the string obtained by replacing the leftmost non-overlapping
					πF.SetLineno(145)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Return the string obtained by replacing the leftmost non-overlapping\n        occurrences of pattern in string by the replacement repl.").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßsub); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 149: def subn(self, repl, string, count=0):
					πF.SetLineno(149)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "repl", Def: nil}
					πTemp002[2] = πg.Param{Name: "string", Def: nil}
					πTemp002[3] = πg.Param{Name: "count", Def: πg.NewInt(0).ToObject()}
					πTemp010 = πg.NewFunction(πg.NewCode("subn", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrepl *πg.Object = πArgs[1]; _ = µrepl
						var µstring *πg.Object = πArgs[2]; _ = µstring
						var µcount *πg.Object = πArgs[3]; _ = µcount
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
							// line 150: """Return the tuple (new_string, number_of_subs_made) found by replacing
							πF.SetLineno(150)
							// line 153: return self._subx(repl, string, count, True)
							πF.SetLineno(153)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µrepl, "repl"); πE != nil {
								continue
							}
							πTemp001[0] = µrepl
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[1] = µstring
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							πTemp001[2] = µcount
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_subx, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsubn.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 150: """Return the tuple (new_string, number_of_subs_made) found by replacing
					πF.SetLineno(150)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Return the tuple (new_string, number_of_subs_made) found by replacing\n        the leftmost non-overlapping occurrences of pattern with the replacement\n        repl.").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßsubn); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 155: def split(self, string, maxsplit=0):
					πF.SetLineno(155)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "maxsplit", Def: πg.NewInt(0).ToObject()}
					πTemp011 = πg.NewFunction(πg.NewCode("split", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µmaxsplit *πg.Object = πArgs[2]; _ = µmaxsplit
						var µsplitlist *πg.Object = πg.UnboundLocal; _ = µsplitlist
						var µstate *πg.Object = πg.UnboundLocal; _ = µstate
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µlast *πg.Object = πg.UnboundLocal; _ = µlast
						var µmatch *πg.Object = πg.UnboundLocal; _ = µmatch
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
						var πTemp007 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 156: """Split string by the occurrences of pattern."""
							πF.SetLineno(156)
							// line 157: splitlist = []
							πF.SetLineno(157)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µsplitlist = πTemp002
							// line 158: state = _State(string, 0, sys.maxint, self.flags)
							πF.SetLineno(158)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[0] = µstring
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmaxint, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßflags, nil); πE != nil {
								continue
							}
							πTemp001[3] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_State); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstate = πTemp003
							// line 159: n = 0
							πF.SetLineno(159)
							µn = πg.NewInt(0).ToObject()
							// line 160: last = state.start
							πF.SetLineno(160)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							µlast = πTemp002
							// line 161: while not maxsplit or n < maxsplit:
							πF.SetLineno(161)
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
							if πE = πg.CheckLocal(πF, µmaxsplit, "maxsplit"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µmaxsplit); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp002 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmaxsplit, "maxsplit"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µn, µmaxsplit); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label4:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 162: state.reset()
							πF.SetLineno(162)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßreset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 163: state.string_position = state.start
							πF.SetLineno(163)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstring_position, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_code, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßsearch, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 164: if not state.search(self._code):
							πF.SetLineno(164)
						Label5:
							// line 165: break
							πF.SetLineno(165)
							πTemp004 = true
							continue
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp003, πTemp008); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 166: if state.start == state.string_position: # zero-width match
							πF.SetLineno(166)
						Label7:
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µlast, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							goto Label10
							// line 167: if last == state.end:                # or end of string
							πF.SetLineno(167)
						Label9:
							// line 168: break
							πF.SetLineno(168)
							πTemp004 = true
							continue
							goto Label10
						Label10:
							// line 169: state.start += 1
							πF.SetLineno(169)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstart, πTemp003); πE != nil {
								continue
							}
							// line 170: continue
							πF.SetLineno(170)
							continue
							goto Label8
						Label8:
							// line 171: splitlist.append(string[last:state.start])
							πF.SetLineno(171)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µlast, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µstring, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µsplitlist, "splitlist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsplitlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgroups, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label11
							}
							goto Label12
							// line 173: if self.groups:
							πF.SetLineno(173)
						Label11:
							// line 174: match = SRE_Match(self, state)
							πF.SetLineno(174)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp001[1] = µstate
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_Match); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmatch = πTemp003
							// line 177: splitlist += (list(match.groups(None)))
							πF.SetLineno(177)
							if πE = πg.CheckLocal(πF, µsplitlist, "splitlist"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmatch, ßgroups, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.IAdd(πF, µsplitlist, πTemp003); πE != nil {
								continue
							}
							µsplitlist = πTemp002
							goto Label12
						Label12:
							// line 178: n += 1
							πF.SetLineno(178)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µn = πTemp002
							// line 179: last = state.start = state.string_position
							πF.SetLineno(179)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							µlast = πTemp002
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstart, πTemp003); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 180: splitlist.append(string[last:state.end])
							πF.SetLineno(180)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlast, "last"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µlast, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µstring, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µsplitlist, "splitlist"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µsplitlist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 181: return splitlist
							πF.SetLineno(181)
							if πE = πg.CheckLocal(πF, µsplitlist, "splitlist"); πE != nil {
								continue
							}
							πR = µsplitlist
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsplit.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 156: """Split string by the occurrences of pattern."""
					πF.SetLineno(156)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Split string by the occurrences of pattern.").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßsplit); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 183: def finditer(self, string, pos=0, endpos=sys.maxint):
					πF.SetLineno(183)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "pos", Def: πg.NewInt(0).ToObject()}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßmaxint, nil); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "endpos", Def: πTemp014}
					πTemp012 = πg.NewFunction(πg.NewCode("finditer", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µpos *πg.Object = πArgs[2]; _ = µpos
						var µendpos *πg.Object = πArgs[3]; _ = µendpos
						var µscanner *πg.Object = πg.UnboundLocal; _ = µscanner
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
							// line 184: """Return a list of all non-overlapping matches of pattern in string."""
							πF.SetLineno(184)
							// line 185: scanner = self.scanner(string, pos, endpos)
							πF.SetLineno(185)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[0] = µstring
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							πTemp001[1] = µpos
							if πE = πg.CheckLocal(πF, µendpos, "endpos"); πE != nil {
								continue
							}
							πTemp001[2] = µendpos
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßscanner, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µscanner = πTemp003
							// line 186: return iter(scanner.search, None)
							πF.SetLineno(186)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µscanner, "scanner"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µscanner, ßsearch, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
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
					if πE = πClass.SetItem(πF, ßfinditer.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 184: """Return a list of all non-overlapping matches of pattern in string."""
					πF.SetLineno(184)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Return a list of all non-overlapping matches of pattern in string.").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßfinditer); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 188: def scanner(self, string, start=0, end=sys.maxint):
					πF.SetLineno(188)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "start", Def: πg.NewInt(0).ToObject()}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßsys); πE != nil {
						continue
					}
					if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßmaxint, nil); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "end", Def: πTemp015}
					πTemp013 = πg.NewFunction(πg.NewCode("scanner", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µstart *πg.Object = πArgs[2]; _ = µstart
						var µend *πg.Object = πArgs[3]; _ = µend
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
							// line 189: return SRE_Scanner(self, string, start, end)
							πF.SetLineno(189)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp001[1] = µstring
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp001[2] = µstart
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp001[3] = µend
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_Scanner); πE != nil {
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
					if πE = πClass.SetItem(πF, ßscanner.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 191: def __copy__(self):
					πF.SetLineno(191)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("__copy__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							// line 192: raise TypeError, "cannot copy this pattern object"
							πF.SetLineno(192)
							πE = πF.Raise(πTemp001, πg.NewStr("cannot copy this pattern object").ToObject(), nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__copy__.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 194: def __deepcopy__(self):
					πF.SetLineno(194)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("__deepcopy__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							// line 195: raise TypeError, "cannot copy this pattern object"
							πF.SetLineno(195)
							πE = πF.Raise(πTemp001, πg.NewStr("cannot copy this pattern object").ToObject(), nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__deepcopy__.ToObject(), πTemp015); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("SRE_Pattern").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSRE_Pattern.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 198: class SRE_Scanner(object):
			πF.SetLineno(198)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SRE_Scanner", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
					// line 199: """Undocumented scanner interface of sre."""
					πF.SetLineno(199)
					// line 199: """Undocumented scanner interface of sre."""
					πF.SetLineno(199)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Undocumented scanner interface of sre.").ToObject()); πE != nil {
						continue
					}
					// line 201: def __init__(self, pattern, string, start, end):
					πF.SetLineno(201)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pattern", Def: nil}
					πTemp002[2] = πg.Param{Name: "string", Def: nil}
					πTemp002[3] = πg.Param{Name: "start", Def: nil}
					πTemp002[4] = πg.Param{Name: "end", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpattern *πg.Object = πArgs[1]; _ = µpattern
						var µstring *πg.Object = πArgs[2]; _ = µstring
						var µstart *πg.Object = πArgs[3]; _ = µstart
						var µend *πg.Object = πArgs[4]; _ = µend
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							// line 202: self.pattern = pattern
							πF.SetLineno(202)
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
							// line 203: self._state = _State(string, start, end, self.pattern.flags)
							πF.SetLineno(203)
							πTemp002 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp002[0] = µstring
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp002[1] = µstart
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp002[2] = µend
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpattern, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßflags, nil); πE != nil {
								continue
							}
							πTemp002[3] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_State); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_state, πTemp001); πE != nil {
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
					// line 205: def _match_search(self, matcher):
					πF.SetLineno(205)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "matcher", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_match_search", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmatcher *πg.Object = πArgs[1]; _ = µmatcher
						var µstate *πg.Object = πg.UnboundLocal; _ = µstate
						var µmatch *πg.Object = πg.UnboundLocal; _ = µmatch
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 206: state = self._state
							πF.SetLineno(206)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_state, nil); πE != nil {
								continue
							}
							µstate = πTemp001
							// line 207: state.reset()
							πF.SetLineno(207)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßreset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 208: state.string_position = state.start
							πF.SetLineno(208)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstring_position, πTemp002); πE != nil {
								continue
							}
							// line 209: match = None
							πF.SetLineno(209)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µmatch = πTemp001
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpattern, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß_code, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µmatcher, "matcher"); πE != nil {
								continue
							}
							if πTemp001, πE = µmatcher.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 210: if matcher(self.pattern._code):
							πF.SetLineno(210)
						Label1:
							// line 211: match = SRE_Match(self.pattern, state)
							πF.SetLineno(211)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpattern, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp003[1] = µstate
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSRE_Match); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmatch = πTemp002
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µmatch == πTemp005).ToObject()
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label3:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 212: if match is None or state.string_position == state.start:
							πF.SetLineno(212)
						Label4:
							// line 213: state.start += 1
							πF.SetLineno(213)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstart, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label5:
							// line 215: state.start = state.string_position
							πF.SetLineno(215)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µstate, ßstart, πTemp002); πE != nil {
								continue
							}
							goto Label6
						Label6:
							// line 216: return match
							πF.SetLineno(216)
							if πE = πg.CheckLocal(πF, µmatch, "match"); πE != nil {
								continue
							}
							πR = µmatch
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_match_search.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 218: def match(self):
					πF.SetLineno(218)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("match", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 219: return self._match_search(self._state.match)
							πF.SetLineno(219)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_state, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmatch, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_match_search, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmatch.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 221: def search(self):
					πF.SetLineno(221)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("search", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 222: return self._match_search(self._state.search)
							πF.SetLineno(222)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_state, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsearch, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_match_search, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsearch.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("SRE_Scanner").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSRE_Scanner.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 225: class SRE_Match(object):
			πF.SetLineno(225)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SRE_Match", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
					// line 227: def __init__(self, pattern, state):
					πF.SetLineno(227)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pattern", Def: nil}
					πTemp002[2] = πg.Param{Name: "state", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpattern *πg.Object = πArgs[1]; _ = µpattern
						var µstate *πg.Object = πArgs[2]; _ = µstate
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
							default: panic("unexpected function state")
							}
							// line 228: self.re = pattern
							πF.SetLineno(228)
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µpattern); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßre, πTemp001); πE != nil {
								continue
							}
							// line 229: self.string = state.string
							πF.SetLineno(229)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßstring, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstring, πTemp002); πE != nil {
								continue
							}
							// line 230: self.pos = state.pos
							πF.SetLineno(230)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßpos, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp002); πE != nil {
								continue
							}
							// line 231: self.endpos = state.end
							πF.SetLineno(231)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßend, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßendpos, πTemp002); πE != nil {
								continue
							}
							// line 232: self.lastindex = state.lastindex
							πF.SetLineno(232)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßlastindex, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlastindex, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlastindex, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 233: if self.lastindex < 0:
							πF.SetLineno(233)
						Label1:
							// line 234: self.lastindex = None
							πF.SetLineno(234)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlastindex, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 235: self.regs = self._create_regs(state)
							πF.SetLineno(235)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp004[0] = µstate
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_create_regs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßregs, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µpattern, ß_indexgroup, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßlastindex, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, πg.NewInt(0).ToObject(), πTemp005); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label4
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µpattern, ß_indexgroup, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp007
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp002, πE = πg.LT(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
						Label4:
							πTemp001 = πTemp002
						Label3:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 236: if pattern._indexgroup and 0 <= self.lastindex < len(pattern._indexgroup):
							πF.SetLineno(236)
						Label5:
							// line 242: self.lastgroup = pattern._indexgroup[self.lastindex]
							πF.SetLineno(242)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßlastindex, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µpattern, "pattern"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µpattern, ß_indexgroup, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlastgroup, πTemp001); πE != nil {
								continue
							}
							goto Label7
						Label6:
							// line 244: self.lastgroup = None
							πF.SetLineno(244)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlastgroup, πTemp002); πE != nil {
								continue
							}
							goto Label7
						Label7:
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
					// line 246: def _create_regs(self, state):
					πF.SetLineno(246)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "state", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_create_regs", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstate *πg.Object = πArgs[1]; _ = µstate
						var µregs *πg.Object = πg.UnboundLocal; _ = µregs
						var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
						var µmark_index *πg.Object = πg.UnboundLocal; _ = µmark_index
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
							// line 247: """Creates a tuple of index pairs representing matched groups."""
							πF.SetLineno(247)
							// line 248: regs = [(state.start, state.string_position)]
							πF.SetLineno(248)
							πTemp001 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µstate, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µregs = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßre, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßgroups, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
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
								µgroup = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 250: mark_index = 2 * group
							πF.SetLineno(250)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), µgroup); πE != nil {
								continue
							}
							µmark_index = πTemp003
							if πE = πg.CheckLocal(πF, µmark_index, "mark_index"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, µmark_index, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µstate, ßmarks, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.LT(πF, πTemp007, πTemp009); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µmark_index, "mark_index"); πE != nil {
								continue
							}
							πTemp007 = µmark_index
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µstate, ßmarks, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp007); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp008 != πTemp007).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µmark_index, "mark_index"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, µmark_index, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µstate, ßmarks, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp007); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp008 != πTemp007).ToObject()
							πTemp003 = πTemp004
						Label4:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 251: if mark_index + 1 < len(state.marks) \
							πF.SetLineno(251)
						Label5:
							// line 254: regs.append((state.marks[mark_index], state.marks[mark_index + 1]))
							πF.SetLineno(254)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmark_index, "mark_index"); πE != nil {
								continue
							}
							πTemp004 = µmark_index
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µstate, ßmarks, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmark_index, "mark_index"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, µmark_index, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp008
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µstate, ßmarks, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µregs, "regs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µregs, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label7
						Label6:
							// line 256: regs.append((-1, -1))
							πF.SetLineno(256)
							πTemp001 = πF.MakeArgs(1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µregs, "regs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µregs, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label7
						Label7:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 257: return tuple(regs)
							πF.SetLineno(257)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µregs, "regs"); πE != nil {
								continue
							}
							πTemp001[0] = µregs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_create_regs.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 247: """Creates a tuple of index pairs representing matched groups."""
					πF.SetLineno(247)
					// line 247: """Creates a tuple of index pairs representing matched groups."""
					πF.SetLineno(247)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Creates a tuple of index pairs representing matched groups.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Creates a tuple of index pairs representing matched groups.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß_create_regs); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 259: def _get_index(self, group):
					πF.SetLineno(259)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "group", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_get_index", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgroup *πg.Object = πArgs[1]; _ = µgroup
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp001[0] = µgroup
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
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
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßre, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßgroupindex, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp005, µgroup); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 260: if isinstance(group, int):
							πF.SetLineno(260)
						Label1:
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, µgroup, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßre, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßgroups, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LE(πF, µgroup, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label4:
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 261: if group >= 0 and group <= self.re.groups:
							πF.SetLineno(261)
						Label5:
							// line 262: return group
							πF.SetLineno(262)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πR = µgroup
							continue
							goto Label6
						Label6:
							goto Label3
							// line 264: if group in self.re.groupindex:
							πF.SetLineno(264)
						Label2:
							// line 265: return self.re.groupindex[group]
							πF.SetLineno(265)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp002 = µgroup
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßre, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßgroupindex, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label3
						Label3:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("no such group").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 266: raise IndexError("no such group")
							πF.SetLineno(266)
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
					if πE = πClass.SetItem(πF, ß_get_index.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 268: def _get_slice(self, group, default):
					πF.SetLineno(268)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "group", Def: nil}
					πTemp002[2] = πg.Param{Name: "default", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_get_slice", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgroup *πg.Object = πArgs[1]; _ = µgroup
						var µdefault *πg.Object = πArgs[2]; _ = µdefault
						var µgroup_indices *πg.Object = πg.UnboundLocal; _ = µgroup_indices
						var πTemp001 *πg.Object
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
							// line 269: group_indices = self.regs[group]
							πF.SetLineno(269)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp001 = µgroup
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßregs, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µgroup_indices = πTemp002
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µgroup_indices, "group_indices"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µgroup_indices, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, πTemp003, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 270: if group_indices[0] >= 0:
							πF.SetLineno(270)
						Label1:
							// line 271: return self.string[group_indices[0]:group_indices[1]]
							πF.SetLineno(271)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µgroup_indices, "group_indices"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µgroup_indices, πTemp002); πE != nil {
								continue
							}
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µgroup_indices, "group_indices"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µgroup_indices, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πTemp005, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstring, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 273: return default
							πF.SetLineno(273)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πR = µdefault
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
					if πE = πClass.SetItem(πF, ß_get_slice.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 275: def start(self, group=0):
					πF.SetLineno(275)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "group", Def: πg.NewInt(0).ToObject()}
					πTemp006 = πg.NewFunction(πg.NewCode("start", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgroup *πg.Object = πArgs[1]; _ = µgroup
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
							// line 276: """Returns the indices of the start of the substring matched by group;
							πF.SetLineno(276)
							// line 279: return self.regs[self._get_index(group)][0]
							πF.SetLineno(279)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp004[0] = µgroup
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_get_index, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßregs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßstart.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 276: """Returns the indices of the start of the substring matched by group;
					πF.SetLineno(276)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Returns the indices of the start of the substring matched by group;\n        group defaults to zero (meaning the whole matched substring). Returns -1\n        if group exists but did not contribute to the match.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßstart); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 281: def end(self, group=0):
					πF.SetLineno(281)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "group", Def: πg.NewInt(0).ToObject()}
					πTemp007 = πg.NewFunction(πg.NewCode("end", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgroup *πg.Object = πArgs[1]; _ = µgroup
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
							// line 282: """Returns the indices of the end of the substring matched by group;
							πF.SetLineno(282)
							// line 285: return self.regs[self._get_index(group)][1]
							πF.SetLineno(285)
							πTemp001 = πg.NewInt(1).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp004[0] = µgroup
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_get_index, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßregs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßend.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 282: """Returns the indices of the end of the substring matched by group;
					πF.SetLineno(282)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Returns the indices of the end of the substring matched by group;\n        group defaults to zero (meaning the whole matched substring). Returns -1\n        if group exists but did not contribute to the match.").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßend); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 287: def span(self, group=0):
					πF.SetLineno(287)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "group", Def: πg.NewInt(0).ToObject()}
					πTemp008 = πg.NewFunction(πg.NewCode("span", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgroup *πg.Object = πArgs[1]; _ = µgroup
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 288: """Returns the 2-tuple (m.start(group), m.end(group))."""
							πF.SetLineno(288)
							// line 289: return self.start(group), self.end(group)
							πF.SetLineno(289)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp002[0] = µgroup
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstart, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp002[0] = µgroup
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
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
					if πE = πClass.SetItem(πF, ßspan.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 288: """Returns the 2-tuple (m.start(group), m.end(group))."""
					πF.SetLineno(288)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Returns the 2-tuple (m.start(group), m.end(group)).").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßspan); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 291: def expand(self, template):
					πF.SetLineno(291)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "template", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("expand", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtemplate *πg.Object = πArgs[1]; _ = µtemplate
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 292: """Return the string obtained by doing backslash substitution and
							πF.SetLineno(292)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							// line 294: raise NotImplementedError
							πF.SetLineno(294)
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
					if πE = πClass.SetItem(πF, ßexpand.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 292: """Return the string obtained by doing backslash substitution and
					πF.SetLineno(292)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Return the string obtained by doing backslash substitution and\n        resolving group references on template.").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßexpand); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 296: def groups(self, default=None):
					πF.SetLineno(296)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "default", Def: πTemp011}
					πTemp010 = πg.NewFunction(πg.NewCode("groups", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdefault *πg.Object = πArgs[1]; _ = µdefault
						var µgroups *πg.Object = πg.UnboundLocal; _ = µgroups
						var µindices *πg.Object = πg.UnboundLocal; _ = µindices
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
							// line 297: """Returns a tuple containing all the subgroups of the match. The
							πF.SetLineno(297)
							// line 300: groups = []
							πF.SetLineno(300)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µgroups = πTemp002
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßregs, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
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
								µindices = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µindices, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							goto Label5
							// line 302: if indices[0] >= 0:
							πF.SetLineno(302)
						Label4:
							// line 303: groups.append(self.string[indices[0]:indices[1]])
							πF.SetLineno(303)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µindices, πTemp004); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µindices, "indices"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µindices, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp005, πTemp008, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstring, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µgroups, "groups"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µgroups, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label5:
							// line 305: groups.append(default)
							πF.SetLineno(305)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp001[0] = µdefault
							if πE = πg.CheckLocal(πF, µgroups, "groups"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µgroups, ßappend, nil); πE != nil {
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
							// line 306: return tuple(groups)
							πF.SetLineno(306)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgroups, "groups"); πE != nil {
								continue
							}
							πTemp001[0] = µgroups
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgroups.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 297: """Returns a tuple containing all the subgroups of the match. The
					πF.SetLineno(297)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Returns a tuple containing all the subgroups of the match. The\n        default argument is used for groups that did not participate in the\n        match (defaults to None).").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßgroups); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 308: def groupdict(self, default=None):
					πF.SetLineno(308)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "default", Def: πTemp012}
					πTemp011 = πg.NewFunction(πg.NewCode("groupdict", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdefault *πg.Object = πArgs[1]; _ = µdefault
						var µgroupdict *πg.Object = πg.UnboundLocal; _ = µgroupdict
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var πTemp001 *πg.Dict
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 309: """Return a dictionary containing all the named subgroups of the match.
							πF.SetLineno(309)
							// line 312: groupdict = {}
							πF.SetLineno(312)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µgroupdict = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßre, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßgroupindex, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßitems, nil); πE != nil {
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
							// line 314: groupdict[key] = self._get_slice(value, default)
							πF.SetLineno(314)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp008[0] = µvalue
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp008[1] = µdefault
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_get_slice, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µgroupdict, "groupdict"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007 = µkey
							if πE = πg.SetItem(πF, µgroupdict, πTemp007, πTemp003); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 315: return groupdict
							πF.SetLineno(315)
							if πE = πg.CheckLocal(πF, µgroupdict, "groupdict"); πE != nil {
								continue
							}
							πR = µgroupdict
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgroupdict.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 309: """Return a dictionary containing all the named subgroups of the match.
					πF.SetLineno(309)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Return a dictionary containing all the named subgroups of the match.\n        The default argument is used for groups that did not participate in the\n        match (defaults to None).").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßgroupdict); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 317: def group(self, *args):
					πF.SetLineno(317)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("group", "/usr/lib/python2.7/_sre.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µgrouplist *πg.Object = πg.UnboundLocal; _ = µgrouplist
						var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 318: """Returns one or more subgroups of the match. Each argument is either a
							πF.SetLineno(318)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 320: if len(args) == 0:
							πF.SetLineno(320)
						Label1:
							// line 321: args = (0,)
							πF.SetLineno(321)
							πTemp001 = πg.NewTuple1(πg.NewInt(0).ToObject()).ToObject()
							µargs = πTemp001
							goto Label2
						Label2:
							// line 322: grouplist = []
							πF.SetLineno(322)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							µgrouplist = πTemp001
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µargs); πE != nil {
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
								µgroup = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 324: grouplist.append(self._get_slice(self._get_index(group), None))
							πF.SetLineno(324)
							πTemp002 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(2)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp008[0] = µgroup
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_get_index, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp007[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_get_slice, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp002[0] = πTemp004
							if πE = πg.CheckLocal(πF, µgrouplist, "grouplist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µgrouplist, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgrouplist, "grouplist"); πE != nil {
								continue
							}
							πTemp002[0] = µgrouplist
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Eq(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 325: if len(grouplist) == 1:
							πF.SetLineno(325)
						Label6:
							// line 326: return grouplist[0]
							πF.SetLineno(326)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µgrouplist, "grouplist"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µgrouplist, πTemp001); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label8
						Label7:
							// line 328: return tuple(grouplist)
							πF.SetLineno(328)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgrouplist, "grouplist"); πE != nil {
								continue
							}
							πTemp002[0] = µgrouplist
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πR = πTemp003
							continue
							goto Label8
						Label8:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßgroup.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 318: """Returns one or more subgroups of the match. Each argument is either a
					πF.SetLineno(318)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Returns one or more subgroups of the match. Each argument is either a\n        group index or a group name.").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßgroup); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 330: def __copy__():
					πF.SetLineno(330)
					πTemp002 = make([]πg.Param, 0)
					πTemp013 = πg.NewFunction(πg.NewCode("__copy__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							// line 331: raise TypeError, "cannot copy this pattern object"
							πF.SetLineno(331)
							πE = πF.Raise(πTemp001, πg.NewStr("cannot copy this pattern object").ToObject(), nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__copy__.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 333: def __deepcopy__():
					πF.SetLineno(333)
					πTemp002 = make([]πg.Param, 0)
					πTemp014 = πg.NewFunction(πg.NewCode("__deepcopy__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							// line 334: raise TypeError, "cannot copy this pattern object"
							πF.SetLineno(334)
							πE = πF.Raise(πTemp001, πg.NewStr("cannot copy this pattern object").ToObject(), nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__deepcopy__.ToObject(), πTemp014); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("SRE_Match").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSRE_Match.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 337: class _State(object):
			πF.SetLineno(337)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_State", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 339: def __init__(self, string, start, end, flags):
					πF.SetLineno(339)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "string", Def: nil}
					πTemp002[2] = πg.Param{Name: "start", Def: nil}
					πTemp002[3] = πg.Param{Name: "end", Def: nil}
					πTemp002[4] = πg.Param{Name: "flags", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstring *πg.Object = πArgs[1]; _ = µstring
						var µstart *πg.Object = πArgs[2]; _ = µstart
						var µend *πg.Object = πArgs[3]; _ = µend
						var µflags *πg.Object = πArgs[4]; _ = µflags
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 340: self.string = string
							πF.SetLineno(340)
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
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µstart, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 341: if start < 0:
							πF.SetLineno(341)
						Label1:
							// line 342: start = 0
							πF.SetLineno(342)
							µstart = πg.NewInt(0).ToObject()
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp003[0] = µstring
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp001, πE = πg.GT(πF, µend, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 343: if end > len(string):
							πF.SetLineno(343)
						Label3:
							// line 344: end = len(string)
							πF.SetLineno(344)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstring, "string"); πE != nil {
								continue
							}
							πTemp003[0] = µstring
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µend = πTemp004
							goto Label4
						Label4:
							// line 345: self.start = start
							πF.SetLineno(345)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstart); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstart, πTemp001); πE != nil {
								continue
							}
							// line 346: self.string_position = self.start
							πF.SetLineno(346)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstart, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstring_position, πTemp004); πE != nil {
								continue
							}
							// line 347: self.end = end
							πF.SetLineno(347)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µend); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßend, πTemp001); πE != nil {
								continue
							}
							// line 348: self.pos = start
							πF.SetLineno(348)
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstart); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpos, πTemp001); πE != nil {
								continue
							}
							// line 349: self.flags = flags
							πF.SetLineno(349)
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µflags); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßflags, πTemp001); πE != nil {
								continue
							}
							// line 350: self.reset()
							πF.SetLineno(350)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreset, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
					// line 352: def reset(self):
					πF.SetLineno(352)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("reset", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 353: self.marks = []
							πF.SetLineno(353)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmarks, πTemp003); πE != nil {
								continue
							}
							// line 354: self.lastindex = -1
							πF.SetLineno(354)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlastindex, πTemp003); πE != nil {
								continue
							}
							// line 355: self.marks_stack = []
							πF.SetLineno(355)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmarks_stack, πTemp003); πE != nil {
								continue
							}
							// line 356: self.context_stack = []
							πF.SetLineno(356)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontext_stack, πTemp003); πE != nil {
								continue
							}
							// line 357: self.repeat = None
							πF.SetLineno(357)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrepeat, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreset.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 359: def match(self, pattern_codes):
					πF.SetLineno(359)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pattern_codes", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("match", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpattern_codes *πg.Object = πArgs[1]; _ = µpattern_codes
						var µdispatcher *πg.Object = πg.UnboundLocal; _ = µdispatcher
						var µhas_matched *πg.Object = πg.UnboundLocal; _ = µhas_matched
						var µcontext *πg.Object = πg.UnboundLocal; _ = µcontext
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6: goto Label6
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µpattern_codes, πTemp004); πE != nil {
								continue
							}
							πTemp004 = ßinfo.ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							πTemp003 = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µpattern_codes, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 362: if pattern_codes[0] == OPCODES["info"] and pattern_codes[3]:
							πF.SetLineno(362)
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßend, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstring_position, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µpattern_codes, πTemp004); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 363: if self.end - self.string_position < pattern_codes[3]:
							πF.SetLineno(363)
						Label4:
							// line 366: return False
							πF.SetLineno(366)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label5:
							goto Label3
						Label3:
							// line 368: dispatcher = _OpcodeDispatcher()
							πF.SetLineno(368)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_OpcodeDispatcher); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdispatcher = πTemp003
							// line 369: self.context_stack.append(_MatchContext(self, pattern_codes))
							πF.SetLineno(369)
							πTemp008 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp009[0] = µself
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							πTemp009[1] = µpattern_codes
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_MatchContext); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp008[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcontext_stack, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 370: has_matched = None
							πF.SetLineno(370)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µhas_matched = πTemp001
							// line 371: while len(self.context_stack) > 0:
							πF.SetLineno(371)
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
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcontext_stack, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp001, πE = πg.GT(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(6)            
							// line 372: context = self.context_stack[-1]
							πF.SetLineno(372)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontext_stack, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µcontext = πTemp003
							// line 373: has_matched = dispatcher.match(context)
							πF.SetLineno(373)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp008[0] = µcontext
							if πE = πg.CheckLocal(πF, µdispatcher, "dispatcher"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdispatcher, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µhas_matched = πTemp003
							if πE = πg.CheckLocal(πF, µhas_matched, "has_matched"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µhas_matched != πTemp003).ToObject()
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label9
							}
							goto Label10
							// line 374: if has_matched is not None: # don't pop if context isn't done
							πF.SetLineno(374)
						Label9:
							// line 377: self.context_stack = self.context_stack[:-1]
							πF.SetLineno(377)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcontext_stack, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcontext_stack, πTemp001); πE != nil {
								continue
							}
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 378: return has_matched
							πF.SetLineno(378)
							if πE = πg.CheckLocal(πF, µhas_matched, "has_matched"); πE != nil {
								continue
							}
							πR = µhas_matched
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
					// line 380: def search(self, pattern_codes):
					πF.SetLineno(380)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pattern_codes", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("search", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpattern_codes *πg.Object = πArgs[1]; _ = µpattern_codes
						var µflags *πg.Object = πg.UnboundLocal; _ = µflags
						var µstring_position *πg.Object = πg.UnboundLocal; _ = µstring_position
						var µcharacter *πg.Object = πg.UnboundLocal; _ = µcharacter
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8: goto Label8
							case 9: goto Label9
							case 11: goto Label11
							case 12: goto Label12
							case 21: goto Label21
							case 22: goto Label22
							default: panic("unexpected function state")
							}
							// line 381: flags = 0
							πF.SetLineno(381)
							µflags = πg.NewInt(0).ToObject()
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µpattern_codes, πTemp002); πE != nil {
								continue
							}
							πTemp002 = ßinfo.ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 382: if pattern_codes[0] == OPCODES["info"]:
							πF.SetLineno(382)
						Label1:
							πTemp003 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µpattern_codes, πTemp003); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßSRE_INFO_PREFIX); πE != nil {
								continue
							}
							if πTemp002, πE = πg.And(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label3
							}
							πTemp003 = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µpattern_codes, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label3:
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 385: if pattern_codes[2] & SRE_INFO_PREFIX and pattern_codes[5] > 1:
							πF.SetLineno(385)
						Label4:
							// line 386: return self.fast_search(pattern_codes)
							πF.SetLineno(386)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							πTemp007[0] = µpattern_codes
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfast_search, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πR = πTemp002
							continue
							goto Label5
						Label5:
							// line 387: flags = pattern_codes[2]
							πF.SetLineno(387)
							πTemp001 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							µflags = πTemp002
							// line 388: pattern_codes = pattern_codes[pattern_codes[1] + 1:]
							πF.SetLineno(388)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µpattern_codes, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							µpattern_codes = πTemp002
							goto Label2
						Label2:
							// line 390: string_position = self.start
							πF.SetLineno(390)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstart, nil); πE != nil {
								continue
							}
							µstring_position = πTemp001
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µpattern_codes, πTemp002); πE != nil {
								continue
							}
							πTemp002 = ßliteral.ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 391: if pattern_codes[0] == OPCODES["literal"]:
							πF.SetLineno(391)
						Label6:
							// line 394: character = pattern_codes[1]
							πF.SetLineno(394)
							πTemp001 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							µcharacter = πTemp002
							// line 395: while True:
							πF.SetLineno(395)
							πF.PushCheckpoint(9)
							πTemp006 = false
						Label8:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label10
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(8)            
							// line 396: while string_position < self.end \
							πF.SetLineno(396)
							πF.PushCheckpoint(12)
							πTemp008 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp008 {
								πF.PopCheckpoint()
								goto Label13
							}
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µstring_position, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label14
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							πTemp003 = µstring_position
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstring, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.CheckLocal(πF, µcharacter, "character"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp004, µcharacter); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label14:
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(11)            
							// line 398: string_position += 1
							πF.SetLineno(398)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µstring_position, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µstring_position = πTemp001
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µstring_position, πTemp002); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label15
							}
							goto Label16
							// line 399: if string_position >= self.end:
							πF.SetLineno(399)
						Label15:
							// line 400: return False
							πF.SetLineno(400)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label16
						Label16:
							// line 401: self.start = string_position
							πF.SetLineno(401)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstring_position); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstart, πTemp001); πE != nil {
								continue
							}
							// line 402: string_position += 1
							πF.SetLineno(402)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µstring_position, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µstring_position = πTemp001
							// line 403: self.string_position = string_position
							πF.SetLineno(403)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstring_position); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstring_position, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_INFO_LITERAL); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, µflags, πTemp002); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label17
							}
							goto Label18
							// line 404: if flags & SRE_INFO_LITERAL:
							πF.SetLineno(404)
						Label17:
							// line 405: return True
							πF.SetLineno(405)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label18
						Label18:
							πTemp007 = πF.MakeArgs(1)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label19
							}
							goto Label20
							// line 406: if self.match(pattern_codes[2:]):
							πF.SetLineno(406)
						Label19:
							// line 407: return True
							πF.SetLineno(407)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label20
						Label20:
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
							// line 408: return False
							πF.SetLineno(408)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label7
						Label7:
							// line 411: while string_position <= self.end:
							πF.SetLineno(411)
							πF.PushCheckpoint(22)
							πTemp006 = false
						Label21:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label23
							}
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, µstring_position, πTemp002); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(21)            
							// line 412: self.reset()
							πF.SetLineno(412)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 413: self.start = self.string_position = string_position
							πF.SetLineno(413)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstring_position); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstart, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstring_position); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstring_position, πTemp001); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							πTemp007[0] = µpattern_codes
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label24
							}
							goto Label25
							// line 414: if self.match(pattern_codes):
							πF.SetLineno(414)
						Label24:
							// line 415: return True
							πF.SetLineno(415)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label25
						Label25:
							// line 416: string_position += 1
							πF.SetLineno(416)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µstring_position, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µstring_position = πTemp001
							continue
						Label22:
							if πE != nil || πR != nil {
								continue
							}
						Label23:
							// line 417: return False
							πF.SetLineno(417)
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
					if πE = πClass.SetItem(πF, ßsearch.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 419: def fast_search(self, pattern_codes):
					πF.SetLineno(419)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pattern_codes", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("fast_search", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpattern_codes *πg.Object = πArgs[1]; _ = µpattern_codes
						var µflags *πg.Object = πg.UnboundLocal; _ = µflags
						var µprefix_len *πg.Object = πg.UnboundLocal; _ = µprefix_len
						var µprefix_skip *πg.Object = πg.UnboundLocal; _ = µprefix_skip
						var µprefix *πg.Object = πg.UnboundLocal; _ = µprefix
						var µoverlap *πg.Object = πg.UnboundLocal; _ = µoverlap
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µstring_position *πg.Object = πg.UnboundLocal; _ = µstring_position
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							// line 420: """Skips forward in a string as fast as possible using information from
							πF.SetLineno(420)
							// line 424: flags = pattern_codes[2]
							πF.SetLineno(424)
							πTemp001 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							µflags = πTemp002
							// line 425: prefix_len = pattern_codes[5]
							πF.SetLineno(425)
							πTemp001 = πg.NewInt(5).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							µprefix_len = πTemp002
							// line 426: prefix_skip = pattern_codes[6] # don't really know what this is good for
							πF.SetLineno(426)
							πTemp001 = πg.NewInt(6).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							µprefix_skip = πTemp002
							// line 427: prefix = pattern_codes[7:7 + prefix_len]
							πF.SetLineno(427)
							if πE = πg.CheckLocal(πF, µprefix_len, "prefix_len"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewInt(7).ToObject(), µprefix_len); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(7).ToObject(), πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							µprefix = πTemp002
							// line 428: overlap = pattern_codes[7 + prefix_len - 1:pattern_codes[1] + 1]
							πF.SetLineno(428)
							if πE = πg.CheckLocal(πF, µprefix_len, "prefix_len"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πg.NewInt(7).ToObject(), µprefix_len); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µpattern_codes, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							µoverlap = πTemp002
							// line 429: pattern_codes = pattern_codes[pattern_codes[1] + 1:]
							πF.SetLineno(429)
							πTemp003 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µpattern_codes, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							µpattern_codes = πTemp002
							// line 430: i = 0
							πF.SetLineno(430)
							µi = πg.NewInt(0).ToObject()
							// line 431: string_position = self.string_position
							πF.SetLineno(431)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstring_position, nil); πE != nil {
								continue
							}
							µstring_position = πTemp001
							// line 432: while string_position < self.end:
							πF.SetLineno(432)
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
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µstring_position, πTemp002); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 433: while True:
							πF.SetLineno(433)
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							πTemp002 = µstring_position
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstring, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp009[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µprefix, "prefix"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µprefix, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							goto Label8
							// line 434: if ord(self.string[string_position]) != prefix[i]:
							πF.SetLineno(434)
						Label7:
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label10
							}
							goto Label11
							// line 435: if i == 0:
							πF.SetLineno(435)
						Label10:
							// line 436: break
							πF.SetLineno(436)
							πTemp007 = true
							continue
							goto Label12
						Label11:
							// line 438: i = overlap[i]
							πF.SetLineno(438)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = µi
							if πE = πg.CheckLocal(πF, µoverlap, "overlap"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µoverlap, πTemp001); πE != nil {
								continue
							}
							µi = πTemp002
							goto Label12
						Label12:
							goto Label9
						Label8:
							// line 440: i += 1
							πF.SetLineno(440)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp001
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µprefix_len, "prefix_len"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µi, µprefix_len); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label13
							}
							goto Label14
							// line 441: if i == prefix_len:
							πF.SetLineno(441)
						Label13:
							// line 443: self.start = string_position + 1 - prefix_len
							πF.SetLineno(443)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µstring_position, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µprefix_len, "prefix_len"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, µprefix_len); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstart, πTemp002); πE != nil {
								continue
							}
							// line 444: self.string_position = string_position + 1 \
							πF.SetLineno(444)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µstring_position, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µprefix_len, "prefix_len"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, µprefix_len); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µprefix_skip, "prefix_skip"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µprefix_skip); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstring_position, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSRE_INFO_LITERAL); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, µflags, πTemp002); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label15
							}
							goto Label16
							// line 446: if flags & SRE_INFO_LITERAL:
							πF.SetLineno(446)
						Label15:
							// line 447: return True # matched all of pure literal pattern
							πF.SetLineno(447)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label16
						Label16:
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprefix_skip, "prefix_skip"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), µprefix_skip); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µpattern_codes, πTemp001); πE != nil {
								continue
							}
							πTemp009[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmatch, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label17
							}
							goto Label18
							// line 448: if self.match(pattern_codes[2 * prefix_skip:]):
							πF.SetLineno(448)
						Label17:
							// line 449: return True
							πF.SetLineno(449)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label18
						Label18:
							// line 450: i = overlap[i]
							πF.SetLineno(450)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001 = µi
							if πE = πg.CheckLocal(πF, µoverlap, "overlap"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µoverlap, πTemp001); πE != nil {
								continue
							}
							µi = πTemp002
							goto Label14
						Label14:
							// line 451: break
							πF.SetLineno(451)
							πTemp007 = true
							continue
							goto Label9
						Label9:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 452: string_position += 1
							πF.SetLineno(452)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µstring_position, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µstring_position = πTemp001
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 453: return False
							πF.SetLineno(453)
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
					if πE = πClass.SetItem(πF, ßfast_search.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 420: """Skips forward in a string as fast as possible using information from
					πF.SetLineno(420)
					// line 420: """Skips forward in a string as fast as possible using information from
					πF.SetLineno(420)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Skips forward in a string as fast as possible using information from\n        an optimization info block.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Skips forward in a string as fast as possible using information from\n        an optimization info block.").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßfast_search); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 455: def set_mark(self, mark_nr, position):
					πF.SetLineno(455)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "mark_nr", Def: nil}
					πTemp002[2] = πg.Param{Name: "position", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("set_mark", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmark_nr *πg.Object = πArgs[1]; _ = µmark_nr
						var µposition *πg.Object = πArgs[2]; _ = µposition
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
							if πE = πg.CheckLocal(πF, µmark_nr, "mark_nr"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, µmark_nr, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 456: if mark_nr & 1:
							πF.SetLineno(456)
						Label1:
							// line 458: self.lastindex = mark_nr / 2 + 1
							πF.SetLineno(458)
							if πE = πg.CheckLocal(πF, µmark_nr, "mark_nr"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Div(πF, µmark_nr, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlastindex, πTemp003); πE != nil {
								continue
							}
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µmark_nr, "mark_nr"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmarks, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.GE(πF, µmark_nr, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 459: if mark_nr >= len(self.marks):
							πF.SetLineno(459)
						Label3:
							// line 462: self.marks += ([None] * (mark_nr - len(self.marks) + 1))
							πF.SetLineno(462)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmarks, nil); πE != nil {
								continue
							}
							πTemp004 = make([]*πg.Object, 1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							πTemp005 = πg.NewList(πTemp004...).ToObject()
							if πE = πg.CheckLocal(πF, µmark_nr, "mark_nr"); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßmarks, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp008
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp007, πE = πg.Sub(πF, µmark_nr, πTemp009); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IAdd(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmarks, πTemp005); πE != nil {
								continue
							}
							goto Label4
						Label4:
							// line 463: self.marks[mark_nr] = position
							πF.SetLineno(463)
							if πE = πg.CheckLocal(πF, µposition, "position"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µposition); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmarks, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmark_nr, "mark_nr"); πE != nil {
								continue
							}
							πTemp005 = µmark_nr
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_mark.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 465: def get_marks(self, group_index):
					πF.SetLineno(465)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "group_index", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("get_marks", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µgroup_index *πg.Object = πArgs[1]; _ = µgroup_index
						var µmarks_index *πg.Object = πg.UnboundLocal; _ = µmarks_index
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
							// line 466: marks_index = 2 * group_index
							πF.SetLineno(466)
							if πE = πg.CheckLocal(πF, µgroup_index, "group_index"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πg.NewInt(2).ToObject(), µgroup_index); πE != nil {
								continue
							}
							µmarks_index = πTemp001
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmarks, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µmarks_index, "marks_index"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µmarks_index, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 467: if len(self.marks) > marks_index + 1:
							πF.SetLineno(467)
						Label1:
							// line 468: return self.marks[marks_index], self.marks[marks_index + 1]
							πF.SetLineno(468)
							if πE = πg.CheckLocal(πF, µmarks_index, "marks_index"); πE != nil {
								continue
							}
							πTemp003 = µmarks_index
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßmarks, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmarks_index, "marks_index"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, µmarks_index, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßmarks, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp004, πTemp006).ToObject()
							πR = πTemp001
							continue
							goto Label3
						Label2:
							// line 470: return None, None
							πF.SetLineno(470)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp003, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ßget_marks.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 472: def marks_push(self):
					πF.SetLineno(472)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("marks_push", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 473: self.marks_stack.append((self.marks[:], self.lastindex))
							πF.SetLineno(473)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmarks, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßlastindex, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, πTemp003).ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmarks_stack, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmarks_push.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 475: def marks_pop(self):
					πF.SetLineno(475)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("marks_pop", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 478: self.marks, self.lastindex = self.marks_stack[-1]
							πF.SetLineno(478)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmarks_stack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmarks, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlastindex, πTemp003); πE != nil {
								continue
							}
							// line 479: self.marks_stack = self.marks_stack[:-1]
							πF.SetLineno(479)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmarks_stack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmarks_stack, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmarks_pop.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 481: def marks_pop_keep(self):
					πF.SetLineno(481)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("marks_pop_keep", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 482: self.marks, self.lastindex = self.marks_stack[-1]
							πF.SetLineno(482)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmarks_stack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmarks, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlastindex, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmarks_pop_keep.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 484: def marks_pop_discard(self):
					πF.SetLineno(484)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("marks_pop_discard", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 486: self.marks_stack = self.marks_stack[:-1]
							πF.SetLineno(486)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmarks_stack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmarks_stack, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmarks_pop_discard.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 488: def lower(self, char_ord):
					πF.SetLineno(488)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "char_ord", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("lower", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µchar_ord *πg.Object = πArgs[1]; _ = µchar_ord
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
							// line 489: return getlower(char_ord, self.flags)
							πF.SetLineno(489)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchar_ord, "char_ord"); πE != nil {
								continue
							}
							πTemp001[0] = µchar_ord
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßflags, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetlower); πE != nil {
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
					if πE = πClass.SetItem(πF, ßlower.ToObject(), πTemp013); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("_State").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_State.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 492: class _MatchContext(object):
			πF.SetLineno(492)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_MatchContext", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 494: def __init__(self, state, pattern_codes):
					πF.SetLineno(494)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "state", Def: nil}
					πTemp002[2] = πg.Param{Name: "pattern_codes", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstate *πg.Object = πArgs[1]; _ = µstate
						var µpattern_codes *πg.Object = πArgs[2]; _ = µpattern_codes
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
							// line 495: self.state = state
							πF.SetLineno(495)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstate); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp001); πE != nil {
								continue
							}
							// line 496: self.pattern_codes = pattern_codes
							πF.SetLineno(496)
							if πE = πg.CheckLocal(πF, µpattern_codes, "pattern_codes"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µpattern_codes); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpattern_codes, πTemp001); πE != nil {
								continue
							}
							// line 497: self.string_position = state.string_position
							πF.SetLineno(497)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstate, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstring_position, πTemp002); πE != nil {
								continue
							}
							// line 498: self.code_position = 0
							πF.SetLineno(498)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcode_position, πTemp001); πE != nil {
								continue
							}
							// line 499: self.has_matched = None
							πF.SetLineno(499)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhas_matched, πTemp002); πE != nil {
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
					// line 501: def push_new_context(self, pattern_offset):
					πF.SetLineno(501)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pattern_offset", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("push_new_context", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpattern_offset *πg.Object = πArgs[1]; _ = µpattern_offset
						var µchild_context *πg.Object = πg.UnboundLocal; _ = µchild_context
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
							// line 502: """Creates a new child context of this context and pushes it on the
							πF.SetLineno(502)
							// line 505: child_context = _MatchContext(self.state,
							πF.SetLineno(505)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcode_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpattern_offset, "pattern_offset"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µpattern_offset); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßpattern_codes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_MatchContext); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µchild_context = πTemp003
							// line 507: self.state.context_stack.append(child_context)
							πF.SetLineno(507)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
								continue
							}
							πTemp001[0] = µchild_context
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcontext_stack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 508: return child_context
							πF.SetLineno(508)
							if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
								continue
							}
							πR = µchild_context
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßpush_new_context.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 502: """Creates a new child context of this context and pushes it on the
					πF.SetLineno(502)
					// line 502: """Creates a new child context of this context and pushes it on the
					πF.SetLineno(502)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Creates a new child context of this context and pushes it on the\n        stack. pattern_offset is the offset off the current code position to\n        start interpreting from.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Creates a new child context of this context and pushes it on the\n        stack. pattern_offset is the offset off the current code position to\n        start interpreting from.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßpush_new_context); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 510: def peek_char(self, peek=0):
					πF.SetLineno(510)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "peek", Def: πg.NewInt(0).ToObject()}
					πTemp004 = πg.NewFunction(πg.NewCode("peek_char", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpeek *πg.Object = πArgs[1]; _ = µpeek
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
							// line 511: return self.state.string[self.string_position + peek]
							πF.SetLineno(511)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpeek, "peek"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µpeek); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstring, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpeek_char.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 513: def skip_char(self, skip_count):
					πF.SetLineno(513)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "skip_count", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("skip_char", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µskip_count *πg.Object = πArgs[1]; _ = µskip_count
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
							// line 514: self.string_position += skip_count
							πF.SetLineno(514)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µskip_count, "skip_count"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, µskip_count); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstring_position, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßskip_char.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 516: def remaining_chars(self):
					πF.SetLineno(516)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("remaining_chars", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 517: return self.state.end - self.string_position
							πF.SetLineno(517)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßend, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstring_position, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp003, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßremaining_chars.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 519: def peek_code(self, peek=0):
					πF.SetLineno(519)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "peek", Def: πg.NewInt(0).ToObject()}
					πTemp007 = πg.NewFunction(πg.NewCode("peek_code", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpeek *πg.Object = πArgs[1]; _ = µpeek
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
							// line 520: return self.pattern_codes[self.code_position + peek]
							πF.SetLineno(520)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcode_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpeek, "peek"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µpeek); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpattern_codes, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßpeek_code.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 522: def skip_code(self, skip_count):
					πF.SetLineno(522)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "skip_count", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("skip_code", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µskip_count *πg.Object = πArgs[1]; _ = µskip_count
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
							// line 523: self.code_position += skip_count
							πF.SetLineno(523)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcode_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µskip_count, "skip_count"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, µskip_count); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcode_position, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßskip_code.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 525: def remaining_codes(self):
					πF.SetLineno(525)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("remaining_codes", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 526: return len(self.pattern_codes) - self.code_position
							πF.SetLineno(526)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpattern_codes, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcode_position, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp004, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßremaining_codes.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 528: def at_beginning(self):
					πF.SetLineno(528)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("at_beginning", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 529: return self.string_position == 0
							πF.SetLineno(529)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstring_position, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_beginning.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 531: def at_end(self):
					πF.SetLineno(531)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("at_end", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 532: return self.string_position == self.state.end
							πF.SetLineno(532)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_end.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 534: def at_linebreak(self):
					πF.SetLineno(534)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("at_linebreak", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 535: return not self.at_end() and _is_linebreak(self.peek_char())
							πF.SetLineno(535)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
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
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_linebreak); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001 = πTemp004
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
					if πE = πClass.SetItem(πF, ßat_linebreak.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 537: def at_boundary(self, word_checker):
					πF.SetLineno(537)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "word_checker", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("at_boundary", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µword_checker *πg.Object = πArgs[1]; _ = µword_checker
						var µthat *πg.Object = πg.UnboundLocal; _ = µthat
						var µthis *πg.Object = πg.UnboundLocal; _ = µthis
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
						var πTemp008 []*πg.Object
						_ = πTemp008
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßat_beginning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 538: if self.at_beginning() and self.at_end():
							πF.SetLineno(538)
						Label2:
							// line 539: return False
							πF.SetLineno(539)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label3:
							// line 540: that = not self.at_beginning() and word_checker(self.peek_char(-1))
							πF.SetLineno(540)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßat_beginning, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label4
							}
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp008[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µword_checker, "word_checker"); πE != nil {
								continue
							}
							if πTemp003, πE = µword_checker.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001 = πTemp003
						Label4:
							µthat = πTemp001
							// line 541: this = not self.at_end() and word_checker(self.peek_char())
							πF.SetLineno(541)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label5
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp004
							if πE = πg.CheckLocal(πF, µword_checker, "word_checker"); πE != nil {
								continue
							}
							if πTemp003, πE = µword_checker.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001 = πTemp003
						Label5:
							µthis = πTemp001
							// line 542: return this != that
							πF.SetLineno(542)
							if πE = πg.CheckLocal(πF, µthis, "this"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µthat, "that"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, µthis, µthat); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_boundary.ToObject(), πTemp013); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("_MatchContext").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_MatchContext.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 545: class _RepeatContext(_MatchContext):
			πF.SetLineno(545)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_MatchContext); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_RepeatContext", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
					// line 547: def __init__(self, context):
					πF.SetLineno(547)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "context", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcontext *πg.Object = πArgs[1]; _ = µcontext
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
							// line 548: _MatchContext.__init__(self, context.state,
							πF.SetLineno(548)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcontext, ßstate, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcontext, ßcode_position, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcontext, ßpattern_codes, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_MatchContext); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 550: self.count = -1
							πF.SetLineno(550)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcount, πTemp003); πE != nil {
								continue
							}
							// line 551: self.previous = context.state.repeat
							πF.SetLineno(551)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcontext, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrepeat, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßprevious, πTemp002); πE != nil {
								continue
							}
							// line 552: self.last_position = None
							πF.SetLineno(552)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlast_position, πTemp003); πE != nil {
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
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("_RepeatContext").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_RepeatContext.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 555: class _Dispatcher(object):
			πF.SetLineno(555)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Dispatcher", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
					// line 557: DISPATCH_TABLE = None
					πF.SetLineno(557)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßDISPATCH_TABLE.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 559: def dispatch(self, code, context):
					πF.SetLineno(559)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "code", Def: nil}
					πTemp002[2] = πg.Param{Name: "context", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("dispatch", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcode *πg.Object = πArgs[1]; _ = µcode
						var µcontext *πg.Object = πArgs[2]; _ = µcontext
						var µmethod *πg.Object = πg.UnboundLocal; _ = µmethod
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
							// line 560: method = self.DISPATCH_TABLE.get(code, self.__class__.unknown)
							πF.SetLineno(560)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[0] = µcode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunknown, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßDISPATCH_TABLE, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmethod = πTemp002
							// line 561: return method(self, context)
							πF.SetLineno(561)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp001[1] = µcontext
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							if πTemp002, πE = µmethod.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdispatch.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 563: def unknown(self, code, ctx):
					πF.SetLineno(563)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "code", Def: nil}
					πTemp002[2] = πg.Param{Name: "ctx", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("unknown", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcode *πg.Object = πArgs[1]; _ = µcode
						var µctx *πg.Object = πArgs[2]; _ = µctx
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNotImplementedError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 564: raise NotImplementedError()
							πF.SetLineno(564)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßunknown.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 566: def build_dispatch_table(cls, code_dict, method_prefix):
					πF.SetLineno(566)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "code_dict", Def: nil}
					πTemp002[2] = πg.Param{Name: "method_prefix", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("build_dispatch_table", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µcode_dict *πg.Object = πArgs[1]; _ = µcode_dict
						var µmethod_prefix *πg.Object = πArgs[2]; _ = µmethod_prefix
						var µtable *πg.Object = πg.UnboundLocal; _ = µtable
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
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 bool
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
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcls, ßDISPATCH_TABLE, nil); πE != nil {
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
							// line 567: if cls.DISPATCH_TABLE is not None:
							πF.SetLineno(567)
						Label1:
							// line 568: return
							πF.SetLineno(568)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 569: table = {}
							πF.SetLineno(569)
							πTemp005 = πg.NewDict()
							πTemp001 = πTemp005.ToObject()
							µtable = πTemp001
							if πE = πg.CheckLocal(πF, µcode_dict, "code_dict"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µcode_dict, ßitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
									continue
								}
								µkey = πTemp003
								µvalue = πTemp007
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp008[0] = µcls
							if πE = πg.CheckLocal(πF, µmethod_prefix, "method_prefix"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µmethod_prefix, µkey).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp008[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 571: if hasattr(cls, "%s%s" % (method_prefix, key)):
							πF.SetLineno(571)
						Label6:
							// line 572: table[value] = getattr(cls, "%s%s" % (method_prefix, key))
							πF.SetLineno(572)
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp008[0] = µcls
							if πE = πg.CheckLocal(πF, µmethod_prefix, "method_prefix"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µmethod_prefix, µkey).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s%s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp008[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp007 = µvalue
							if πE = πg.SetItem(πF, µtable, πTemp007, πTemp002); πE != nil {
								continue
							}
							goto Label7
						Label7:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 573: cls.DISPATCH_TABLE = table
							πF.SetLineno(573)
							if πE = πg.CheckLocal(πF, µtable, "table"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtable); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcls, ßDISPATCH_TABLE, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßbuild_dispatch_table.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 575: build_dispatch_table = classmethod(build_dispatch_table)
					πF.SetLineno(575)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßbuild_dispatch_table); πE != nil {
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
					if πE = πClass.SetItem(πF, ßbuild_dispatch_table.ToObject(), πTemp007); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("_Dispatcher").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Dispatcher.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 578: class _OpcodeDispatcher(_Dispatcher):
			πF.SetLineno(578)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_Dispatcher); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_OpcodeDispatcher", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
				var πTemp009 []πg.Param
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
				var πTemp035 *πg.Object
				_ = πTemp035
				var πTemp036 *πg.Object
				_ = πTemp036
				var πTemp037 *πg.Object
				_ = πTemp037
				var πTemp038 *πg.Object
				_ = πTemp038
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 580: def __init__(self):
					πF.SetLineno(580)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 581: self.executing_contexts = {}
							πF.SetLineno(581)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßexecuting_contexts, πTemp003); πE != nil {
								continue
							}
							// line 582: self.at_dispatcher = _AtcodeDispatcher()
							πF.SetLineno(582)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_AtcodeDispatcher); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßat_dispatcher, πTemp002); πE != nil {
								continue
							}
							// line 583: self.ch_dispatcher = _ChcodeDispatcher()
							πF.SetLineno(583)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_ChcodeDispatcher); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßch_dispatcher, πTemp002); πE != nil {
								continue
							}
							// line 584: self.set_dispatcher = _CharsetDispatcher()
							πF.SetLineno(584)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_CharsetDispatcher); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßset_dispatcher, πTemp002); πE != nil {
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
					// line 586: def match(self, context):
					πF.SetLineno(586)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "context", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("match", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcontext *πg.Object = πArgs[1]; _ = µcontext
						var µopcode *πg.Object = πg.UnboundLocal; _ = µopcode
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 bool
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
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 587: """Returns True if the current context matches, False if it doesn't and
							πF.SetLineno(587)
							// line 590: while context.remaining_codes() > 0 and context.has_matched is None:
							πF.SetLineno(590)
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
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µcontext, ßremaining_codes, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GT(πF, πTemp007, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µcontext, ßhas_matched, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp006 == πTemp007).ToObject()
							πTemp003 = πTemp005
						Label4:
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 591: opcode = context.peek_code()
							πF.SetLineno(591)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcontext, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µopcode = πTemp005
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µopcode, "opcode"); πE != nil {
								continue
							}
							πTemp008[0] = µopcode
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp008[1] = µcontext
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp002, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 592: if not self.dispatch(opcode, context):
							πF.SetLineno(592)
						Label5:
							// line 593: return None
							πF.SetLineno(593)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label6
						Label6:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcontext, ßhas_matched, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005 == πTemp006).ToObject()
							if πTemp001, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label7
							}
							goto Label8
							// line 594: if context.has_matched is None:
							πF.SetLineno(594)
						Label7:
							// line 595: context.has_matched = False
							πF.SetLineno(595)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µcontext, ßhas_matched, πTemp005); πE != nil {
								continue
							}
							goto Label8
						Label8:
							// line 596: return context.has_matched
							πF.SetLineno(596)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µcontext, ßhas_matched, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmatch.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 587: """Returns True if the current context matches, False if it doesn't and
					πF.SetLineno(587)
					// line 587: """Returns True if the current context matches, False if it doesn't and
					πF.SetLineno(587)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Returns True if the current context matches, False if it doesn't and\n        None if matching is not finished, ie must be resumed after child\n        contexts have been matched.").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Returns True if the current context matches, False if it doesn't and\n        None if matching is not finished, ie must be resumed after child\n        contexts have been matched.").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßmatch); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 598: def dispatch(self, opcode, context):
					πF.SetLineno(598)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "opcode", Def: nil}
					πTemp002[2] = πg.Param{Name: "context", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("dispatch", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopcode *πg.Object = πArgs[1]; _ = µopcode
						var µcontext *πg.Object = πArgs[2]; _ = µcontext
						var µgenerator *πg.Object = πg.UnboundLocal; _ = µgenerator
						var µhas_finished *πg.Object = πg.UnboundLocal; _ = µhas_finished
						var µmethod *πg.Object = πg.UnboundLocal; _ = µmethod
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
							// line 599: """Dispatches a context on a given opcode. Returns True if the context
							πF.SetLineno(599)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[0] = µcontext
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßexecuting_contexts, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 601: if id(context) in self.executing_contexts:
							πF.SetLineno(601)
						Label1:
							// line 602: generator = self.executing_contexts[id(context)]
							πF.SetLineno(602)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[0] = µcontext
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
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
							if πTemp004, πE = πg.GetAttr(πF, µself, ßexecuting_contexts, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µgenerator = πTemp003
							// line 603: del self.executing_contexts[id(context)]
							πF.SetLineno(603)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßexecuting_contexts, nil); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[0] = µcontext
							if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp006
							if πE = πg.DelItem(πF, πTemp001, πTemp003); πE != nil {
								continue
							}
							// line 604: has_finished = generator.next()
							πF.SetLineno(604)
							if πE = πg.CheckLocal(πF, µgenerator, "generator"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µgenerator, ßnext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µhas_finished = πTemp003
							goto Label3
						Label2:
							// line 606: method = self.DISPATCH_TABLE.get(opcode, _OpcodeDispatcher.unknown)
							πF.SetLineno(606)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µopcode, "opcode"); πE != nil {
								continue
							}
							πTemp002[0] = µopcode
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_OpcodeDispatcher); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßunknown, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßDISPATCH_TABLE, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µmethod = πTemp001
							// line 607: has_finished = method(self, context)
							πF.SetLineno(607)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[1] = µcontext
							if πE = πg.CheckLocal(πF, µmethod, "method"); πE != nil {
								continue
							}
							if πTemp001, πE = µmethod.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µhas_finished = πTemp001
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhas_finished, "has_finished"); πE != nil {
								continue
							}
							πTemp002[0] = µhas_finished
							πTemp002[1] = ßnext.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 608: if hasattr(has_finished, "next"): # avoid using the types module
							πF.SetLineno(608)
						Label4:
							// line 609: generator = has_finished
							πF.SetLineno(609)
							if πE = πg.CheckLocal(πF, µhas_finished, "has_finished"); πE != nil {
								continue
							}
							µgenerator = µhas_finished
							// line 610: has_finished = generator.next()
							πF.SetLineno(610)
							if πE = πg.CheckLocal(πF, µgenerator, "generator"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µgenerator, ßnext, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µhas_finished = πTemp003
							goto Label5
						Label5:
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µhas_finished, "has_finished"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µhas_finished); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 611: if not has_finished:
							πF.SetLineno(611)
						Label6:
							// line 612: self.executing_contexts[id(context)] = generator
							πF.SetLineno(612)
							if πE = πg.CheckLocal(πF, µgenerator, "generator"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µgenerator); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßexecuting_contexts, nil); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							πTemp002[0] = µcontext
							if πTemp006, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp004 = πTemp007
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
								continue
							}
							goto Label7
						Label7:
							// line 613: return has_finished
							πF.SetLineno(613)
							if πE = πg.CheckLocal(πF, µhas_finished, "has_finished"); πE != nil {
								continue
							}
							πR = µhas_finished
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdispatch.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 599: """Dispatches a context on a given opcode. Returns True if the context
					πF.SetLineno(599)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Dispatches a context on a given opcode. Returns True if the context\n        is done matching, False if it must be resumed when next encountered.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 615: def op_success(self, ctx):
					πF.SetLineno(615)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("op_success", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 618: ctx.state.string_position = ctx.string_position
							πF.SetLineno(618)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, πTemp003, ßstring_position, πTemp002); πE != nil {
								continue
							}
							// line 619: ctx.has_matched = True
							πF.SetLineno(619)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp002); πE != nil {
								continue
							}
							// line 620: return True
							πF.SetLineno(620)
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
					if πE = πClass.SetItem(πF, ßop_success.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 622: def op_failure(self, ctx):
					πF.SetLineno(622)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("op_failure", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 625: ctx.has_matched = False
							πF.SetLineno(625)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp002); πE != nil {
								continue
							}
							// line 626: return True
							πF.SetLineno(626)
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
					if πE = πClass.SetItem(πF, ßop_failure.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 628: def general_op_literal(self, ctx, compare, decorate=lambda x: x):
					πF.SetLineno(628)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp002[2] = πg.Param{Name: "compare", Def: nil}
					πTemp009 = make([]πg.Param, 1)
					πTemp009[0] = πg.Param{Name: "x", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sre.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 628: def general_op_literal(self, ctx, compare, decorate=lambda x: x):
							πF.SetLineno(628)
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
					πTemp002[3] = πg.Param{Name: "decorate", Def: πTemp008}
					πTemp007 = πg.NewFunction(πg.NewCode("general_op_literal", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µcompare *πg.Object = πArgs[2]; _ = µcompare
						var µdecorate *πg.Object = πArgs[3]; _ = µdecorate
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
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
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp005 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp008
							if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp008
							if πE = πg.CheckLocal(πF, µdecorate, "decorate"); πE != nil {
								continue
							}
							if πTemp004, πE = µdecorate.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp004
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							πTemp007[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp006[0] = πTemp008
							if πE = πg.CheckLocal(πF, µdecorate, "decorate"); πE != nil {
								continue
							}
							if πTemp004, πE = µdecorate.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[1] = πTemp004
							if πE = πg.CheckLocal(πF, µcompare, "compare"); πE != nil {
								continue
							}
							if πTemp004, πE = µcompare.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp009).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 629: if ctx.at_end() or not compare(decorate(ord(ctx.peek_char())),
							πF.SetLineno(629)
						Label2:
							// line 631: ctx.has_matched = False
							πF.SetLineno(631)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp003); πE != nil {
								continue
							}
							goto Label3
						Label3:
							// line 632: ctx.skip_code(2)
							πF.SetLineno(632)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 633: ctx.skip_char(1)
							πF.SetLineno(633)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgeneral_op_literal.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 635: def op_literal(self, ctx):
					πF.SetLineno(635)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("op_literal", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 639: self.general_op_literal(ctx, operator.eq)
							πF.SetLineno(639)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[0] = µctx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgeneral_op_literal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 640: return True
							πF.SetLineno(640)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_literal.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 642: def op_not_literal(self, ctx):
					πF.SetLineno(642)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("op_not_literal", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 646: self.general_op_literal(ctx, operator.ne)
							πF.SetLineno(646)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[0] = µctx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgeneral_op_literal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 647: return True
							πF.SetLineno(647)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_not_literal.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 649: def op_literal_ignore(self, ctx):
					πF.SetLineno(649)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("op_literal_ignore", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 653: self.general_op_literal(ctx, operator.eq, ctx.state.lower)
							πF.SetLineno(653)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[0] = µctx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßeq, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlower, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgeneral_op_literal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 654: return True
							πF.SetLineno(654)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_literal_ignore.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 656: def op_not_literal_ignore(self, ctx):
					πF.SetLineno(656)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("op_not_literal_ignore", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 660: self.general_op_literal(ctx, operator.ne, ctx.state.lower)
							πF.SetLineno(660)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[0] = µctx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßne, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlower, nil); πE != nil {
								continue
							}
							πTemp001[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgeneral_op_literal, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 661: return True
							πF.SetLineno(661)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_not_literal_ignore.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 663: def op_at(self, ctx):
					πF.SetLineno(663)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("op_at", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(2)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp005
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp002[1] = µctx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßat_dispatcher, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 667: if not self.at_dispatcher.dispatch(ctx.peek_code(1), ctx):
							πF.SetLineno(667)
						Label1:
							// line 668: ctx.has_matched = False
							πF.SetLineno(668)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp004); πE != nil {
								continue
							}
							// line 669: return True
							πF.SetLineno(669)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 670: ctx.skip_code(2)
							πF.SetLineno(670)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 671: return True
							πF.SetLineno(671)
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
					if πE = πClass.SetItem(πF, ßop_at.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 673: def op_category(self, ctx):
					πF.SetLineno(673)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("op_category", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp005 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp007
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp005[1] = µctx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßch_dispatcher, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp004, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp008).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 677: if ctx.at_end() or not self.ch_dispatcher.dispatch(ctx.peek_code(1), ctx):
							πF.SetLineno(677)
						Label2:
							// line 678: ctx.has_matched = False
							πF.SetLineno(678)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp003); πE != nil {
								continue
							}
							// line 679: return True
							πF.SetLineno(679)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label3:
							// line 680: ctx.skip_code(2)
							πF.SetLineno(680)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 681: ctx.skip_char(1)
							πF.SetLineno(681)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 682: return True
							πF.SetLineno(682)
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
					if πE = πClass.SetItem(πF, ßop_category.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 684: def op_any(self, ctx):
					πF.SetLineno(684)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("op_any", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_linebreak, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 688: if ctx.at_end() or ctx.at_linebreak():
							πF.SetLineno(688)
						Label2:
							// line 689: ctx.has_matched = False
							πF.SetLineno(689)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp003); πE != nil {
								continue
							}
							// line 690: return True
							πF.SetLineno(690)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label3:
							// line 691: ctx.skip_code(1)
							πF.SetLineno(691)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 692: ctx.skip_char(1)
							πF.SetLineno(692)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 693: return True
							πF.SetLineno(693)
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
					if πE = πClass.SetItem(πF, ßop_any.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 695: def op_any_all(self, ctx):
					πF.SetLineno(695)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("op_any_all", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 699: if ctx.at_end():
							πF.SetLineno(699)
						Label1:
							// line 700: ctx.has_matched = False
							πF.SetLineno(700)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp002); πE != nil {
								continue
							}
							// line 701: return True
							πF.SetLineno(701)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label2
						Label2:
							// line 702: ctx.skip_code(1)
							πF.SetLineno(702)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 703: ctx.skip_char(1)
							πF.SetLineno(703)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 704: return True
							πF.SetLineno(704)
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
					if πE = πClass.SetItem(πF, ßop_any_all.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 706: def general_op_in(self, ctx, decorate=lambda x: x):
					πF.SetLineno(706)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp009 = make([]πg.Param, 1)
					πTemp009[0] = πg.Param{Name: "x", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sre.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 706: def general_op_in(self, ctx, decorate=lambda x: x):
							πF.SetLineno(706)
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
					πTemp002[2] = πg.Param{Name: "decorate", Def: πTemp018}
					πTemp017 = πg.NewFunction(πg.NewCode("general_op_in", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µdecorate *πg.Object = πArgs[2]; _ = µdecorate
						var µskip *πg.Object = πg.UnboundLocal; _ = µskip
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
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 708: if ctx.at_end():
							πF.SetLineno(708)
						Label1:
							// line 709: ctx.has_matched = False
							πF.SetLineno(709)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp002); πE != nil {
								continue
							}
							// line 710: return
							πF.SetLineno(710)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 711: skip = ctx.peek_code(1)
							πF.SetLineno(711)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µskip = πTemp002
							// line 712: ctx.skip_code(2) # set op pointer to the set code
							πF.SetLineno(712)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp004[0] = µctx
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp007
							if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp007
							if πE = πg.CheckLocal(πF, µdecorate, "decorate"); πE != nil {
								continue
							}
							if πTemp002, πE = µdecorate.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheck_charset, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.IsTrue(πF, πTemp007); πE != nil {
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
							// line 713: if not self.check_charset(ctx, decorate(ord(ctx.peek_char()))):
							πF.SetLineno(713)
						Label3:
							// line 714: ctx.has_matched = False
							πF.SetLineno(714)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp002); πE != nil {
								continue
							}
							// line 715: return
							πF.SetLineno(715)
							πR = πg.None
							continue
							goto Label4
						Label4:
							// line 716: ctx.skip_code(skip - 1)
							πF.SetLineno(716)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µskip, "skip"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µskip, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 717: ctx.skip_char(1)
							πF.SetLineno(717)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgeneral_op_in.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 719: def op_in(self, ctx):
					πF.SetLineno(719)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("op_in", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 723: self.general_op_in(ctx)
							πF.SetLineno(723)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[0] = µctx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgeneral_op_in, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 724: return True
							πF.SetLineno(724)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_in.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 726: def op_in_ignore(self, ctx):
					πF.SetLineno(726)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("op_in_ignore", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 730: self.general_op_in(ctx, ctx.state.lower)
							πF.SetLineno(730)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[0] = µctx
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlower, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgeneral_op_in, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 731: return True
							πF.SetLineno(731)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_in_ignore.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 733: def op_jump(self, ctx):
					πF.SetLineno(733)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("op_jump", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 737: ctx.skip_code(ctx.peek_code(1) + 1)
							πF.SetLineno(737)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 738: return True
							πF.SetLineno(738)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_jump.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 742: op_info = op_jump
					πF.SetLineno(742)
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßop_jump); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßop_info.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 744: def op_mark(self, ctx):
					πF.SetLineno(744)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("op_mark", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 748: ctx.state.set_mark(ctx.peek_code(1), ctx.string_position)
							πF.SetLineno(748)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßset_mark, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 749: ctx.skip_code(2)
							πF.SetLineno(749)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 750: return True
							πF.SetLineno(750)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_mark.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 752: def op_branch(self, ctx):
					πF.SetLineno(752)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("op_branch", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µcurrent_branch_length *πg.Object = πg.UnboundLocal; _ = µcurrent_branch_length
						var µchild_context *πg.Object = πg.UnboundLocal; _ = µchild_context
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
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 8: goto Label8
								case 1: goto Label1
								case 2: goto Label2
								case 11: goto Label11
								case 12: goto Label12
								default: panic("unexpected function state")
								}
								// line 756: ctx.state.marks_push()
								πF.SetLineno(756)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmarks_push, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
									continue
								}
								// line 757: ctx.skip_code(1)
								πF.SetLineno(757)
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								// line 758: current_branch_length = ctx.peek_code(0)
								πF.SetLineno(758)
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								µcurrent_branch_length = πTemp002
								// line 759: while current_branch_length:
								πF.SetLineno(759)
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
								if πE = πg.CheckLocal(πF, µcurrent_branch_length, "current_branch_length"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, µcurrent_branch_length); πE != nil {
									continue
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								πTemp007 = ßliteral.ToObject()
								if πTemp010, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp007); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Eq(πF, πTemp008, πTemp009); πE != nil {
									continue
								}
								πTemp002 = πTemp006
								if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if !πTemp005 {
									goto Label4
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp006 = πTemp008
								if πTemp011, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp011 {
									goto Label5
								}
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								πTemp003 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp003[0] = πTemp010
								if πTemp008, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								if πTemp007, πE = πg.NE(πF, πTemp009, πTemp010); πE != nil {
									continue
								}
								πTemp006 = πTemp007
							Label5:
								πTemp002 = πTemp006
							Label4:
								if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(!πTemp005).ToObject()
								if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label6
								}
								goto Label7
								// line 762: if not (ctx.peek_code(1) == OPCODES["literal"] and \
								πF.SetLineno(762)
							Label6:
								// line 764: ctx.state.string_position = ctx.string_position
								πF.SetLineno(764)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp006, ßstring_position, πTemp002); πE != nil {
									continue
								}
								// line 765: child_context = ctx.push_new_context(1)
								πF.SetLineno(765)
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								µchild_context = πTemp002
								// line 766: yield False
								πF.SetLineno(766)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(8)
								return πTemp001, nil
							Label8:
								πTemp002 = πSent
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label9
								}
								goto Label10
								// line 767: if child_context.has_matched:
								πF.SetLineno(767)
							Label9:
								// line 768: ctx.has_matched = True
								πF.SetLineno(768)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp006); πE != nil {
									continue
								}
								// line 769: yield True
								πF.SetLineno(769)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(11)
								return πTemp002, nil
							Label11:
								πTemp006 = πSent
								goto Label10
							Label10:
								// line 770: ctx.state.marks_pop_keep()
								πF.SetLineno(770)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßmarks_pop_keep, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								goto Label7
							Label7:
								// line 771: ctx.skip_code(current_branch_length)
								πF.SetLineno(771)
								πTemp003 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µcurrent_branch_length, "current_branch_length"); πE != nil {
									continue
								}
								πTemp003[0] = µcurrent_branch_length
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								// line 772: current_branch_length = ctx.peek_code(0)
								πF.SetLineno(772)
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(0).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								µcurrent_branch_length = πTemp007
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
								// line 773: ctx.state.marks_pop_discard()
								πF.SetLineno(773)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßmarks_pop_discard, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp007.Call(πF, nil, nil); πE != nil {
									continue
								}
								// line 774: ctx.has_matched = False
								πF.SetLineno(774)
								if πTemp006, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp006); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp007); πE != nil {
									continue
								}
								// line 775: yield True
								πF.SetLineno(775)
								if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(12)
								return πTemp006, nil
							Label12:
								πTemp007 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßop_branch.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 777: def op_repeat_one(self, ctx):
					πF.SetLineno(777)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("op_repeat_one", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µmincount *πg.Object = πg.UnboundLocal; _ = µmincount
						var µmaxcount *πg.Object = πg.UnboundLocal; _ = µmaxcount
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µchar *πg.Object = πg.UnboundLocal; _ = µchar
						var µchild_context *πg.Object = πg.UnboundLocal; _ = µchild_context
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πTemp013 bool
						_ = πTemp013
						var πTemp014 bool
						_ = πTemp014
						var πTemp015 bool
						_ = πTemp015
						var πTemp016 *πg.Object
						_ = πTemp016
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 33: goto Label33
								case 34: goto Label34
								case 3: goto Label3
								case 6: goto Label6
								case 9: goto Label9
								case 13: goto Label13
								case 14: goto Label14
								case 16: goto Label16
								case 17: goto Label17
								case 23: goto Label23
								case 26: goto Label26
								case 27: goto Label27
								case 28: goto Label28
								case 30: goto Label30
								default: panic("unexpected function state")
								}
								// line 782: mincount = ctx.peek_code(2)
								πF.SetLineno(782)
								πTemp001 = πF.MakeArgs(1)
								πTemp001[0] = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µmincount = πTemp003
								// line 783: maxcount = ctx.peek_code(3)
								πF.SetLineno(783)
								πTemp001 = πF.MakeArgs(1)
								πTemp001[0] = πg.NewInt(3).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µmaxcount = πTemp003
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µctx, ßremaining_chars, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.LT(πF, πTemp004, µmincount); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label1
								}
								goto Label2
								// line 786: if ctx.remaining_chars() < mincount:
								πF.SetLineno(786)
							Label1:
								// line 787: ctx.has_matched = False
								πF.SetLineno(787)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp003); πE != nil {
									continue
								}
								// line 788: yield True
								πF.SetLineno(788)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(3)
								return πTemp002, nil
							Label3:
								πTemp003 = πSent
								goto Label2
							Label2:
								// line 789: ctx.state.string_position = ctx.string_position
								πF.SetLineno(789)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp006, ßstring_position, πTemp004); πE != nil {
									continue
								}
								// line 790: count = self.count_repetitions(ctx, maxcount)
								πF.SetLineno(790)
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								πTemp001[0] = µctx
								if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
									continue
								}
								πTemp001[1] = µmaxcount
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßcount_repetitions, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µcount = πTemp004
								// line 791: ctx.skip_char(count)
								πF.SetLineno(791)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								πTemp001[0] = µcount
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.LT(πF, µcount, µmincount); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label4
								}
								goto Label5
								// line 792: if count < mincount:
								πF.SetLineno(792)
							Label4:
								// line 793: ctx.has_matched = False
								πF.SetLineno(793)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp004); πE != nil {
									continue
								}
								// line 794: yield True
								πF.SetLineno(794)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return πTemp003, nil
							Label6:
								πTemp004 = πSent
								goto Label5
							Label5:
								πTemp001 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp007[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp006, πE = πg.Add(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp006
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πTemp006 = ßsuccess.ToObject()
								if πTemp010, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp006); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Eq(πF, πTemp008, πTemp009); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label7
								}
								goto Label8
								// line 795: if ctx.peek_code(ctx.peek_code(1) + 1) == OPCODES["success"]:
								πF.SetLineno(795)
							Label7:
								// line 797: ctx.state.string_position = ctx.string_position
								πF.SetLineno(797)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp004); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp008, ßstring_position, πTemp006); πE != nil {
									continue
								}
								// line 798: ctx.has_matched = True
								πF.SetLineno(798)
								if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp004); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp006); πE != nil {
									continue
								}
								// line 799: yield True
								πF.SetLineno(799)
								if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(9)
								return πTemp004, nil
							Label9:
								πTemp006 = πSent
								goto Label8
							Label8:
								// line 801: ctx.state.marks_push()
								πF.SetLineno(801)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßmarks_push, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp008.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp001 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp007[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp009.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp008, πE = πg.Add(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp008
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πTemp008 = ßliteral.ToObject()
								if πTemp011, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetItem(πF, πTemp011, πTemp008); πE != nil {
									continue
								}
								if πTemp006, πE = πg.Eq(πF, πTemp009, πTemp010); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label10
								}
								goto Label11
								// line 802: if ctx.peek_code(ctx.peek_code(1) + 1) == OPCODES["literal"]:
								πF.SetLineno(802)
							Label10:
								// line 805: char = ctx.peek_code(ctx.peek_code(1) + 2)
								πF.SetLineno(805)
								πTemp001 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp007[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp006, πE = πg.Add(πF, πTemp009, πg.NewInt(2).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp006
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µchar = πTemp008
								// line 806: while True:
								πF.SetLineno(806)
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
								if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πE != nil || !πTemp012 {
									continue
								}
								πF.PushCheckpoint(13)            
								// line 807: while count >= mincount and \
								πF.SetLineno(807)
								πF.PushCheckpoint(17)
								πTemp012 = false
							Label16:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp012 {
									πF.PopCheckpoint()
									goto Label18
								}
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GE(πF, µcount, µmincount); πE != nil {
									continue
								}
								πTemp006 = πTemp008
								if πTemp014, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if !πTemp014 {
									goto Label19
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp008 = πTemp010
								if πTemp015, πE = πg.IsTrue(πF, πTemp008); πE != nil {
									continue
								}
								if πTemp015 {
									goto Label20
								}
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
									continue
								}
								if πTemp011, πE = πTemp010.Call(πF, nil, nil); πE != nil {
									continue
								}
								πTemp001[0] = πTemp011
								if πTemp010, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
									continue
								}
								if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.NE(πF, πTemp011, µchar); πE != nil {
									continue
								}
								πTemp008 = πTemp009
							Label20:
								πTemp006 = πTemp008
							Label19:
								if πTemp013, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πE != nil || !πTemp013 {
									continue
								}
								πF.PushCheckpoint(16)            
								// line 809: ctx.skip_char(-1)
								πF.SetLineno(809)
								πTemp001 = πF.MakeArgs(1)
								if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp006
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								// line 810: count -= 1
								πF.SetLineno(810)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.ISub(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µcount = πTemp006
								continue
							Label17:
								if πE != nil || πR != nil {
									continue
								}
							Label18:
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.LT(πF, µcount, µmincount); πE != nil {
									continue
								}
								if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp012 {
									goto Label21
								}
								goto Label22
								// line 811: if count < mincount:
								πF.SetLineno(811)
							Label21:
								// line 812: break
								πF.SetLineno(812)
								πTemp005 = true
								continue
								goto Label22
							Label22:
								// line 813: ctx.state.string_position = ctx.string_position
								πF.SetLineno(813)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp006); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp009, ßstring_position, πTemp008); πE != nil {
									continue
								}
								// line 814: child_context = ctx.push_new_context(ctx.peek_code(1) + 1)
								πF.SetLineno(814)
								πTemp001 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp007[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp006, πE = πg.Add(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp006
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µchild_context = πTemp008
								// line 815: yield False
								πF.SetLineno(815)
								if πTemp006, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(23)
								return πTemp006, nil
							Label23:
								πTemp008 = πSent
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp012, πE = πg.IsTrue(πF, πTemp008); πE != nil {
									continue
								}
								if πTemp012 {
									goto Label24
								}
								goto Label25
								// line 816: if child_context.has_matched:
								πF.SetLineno(816)
							Label24:
								// line 817: ctx.has_matched = True
								πF.SetLineno(817)
								if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp008); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp009); πE != nil {
									continue
								}
								// line 818: yield True
								πF.SetLineno(818)
								if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(26)
								return πTemp008, nil
							Label26:
								πTemp009 = πSent
								goto Label25
							Label25:
								// line 819: ctx.skip_char(-1)
								πF.SetLineno(819)
								πTemp001 = πF.MakeArgs(1)
								if πTemp009, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp009
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								// line 820: count -= 1
								πF.SetLineno(820)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ISub(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µcount = πTemp009
								// line 821: ctx.state.marks_pop_keep()
								πF.SetLineno(821)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßmarks_pop_keep, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp010.Call(πF, nil, nil); πE != nil {
									continue
								}
								continue
							Label14:
								if πE != nil || πR != nil {
									continue
								}
							Label15:
								goto Label12
							Label11:
								// line 825: while count >= mincount:
								πF.SetLineno(825)
								πF.PushCheckpoint(28)
								πTemp005 = false
							Label27:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp005 {
									πF.PopCheckpoint()
									goto Label29
								}
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GE(πF, µcount, µmincount); πE != nil {
									continue
								}
								if πTemp012, πE = πg.IsTrue(πF, πTemp009); πE != nil {
									continue
								}
								if πE != nil || !πTemp012 {
									continue
								}
								πF.PushCheckpoint(27)            
								// line 826: ctx.state.string_position = ctx.string_position
								πF.SetLineno(826)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp011, ßstring_position, πTemp010); πE != nil {
									continue
								}
								// line 827: child_context = ctx.push_new_context(ctx.peek_code(1) + 1)
								πF.SetLineno(827)
								πTemp001 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp007[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp011, πE = πTemp010.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp009, πE = πg.Add(πF, πTemp011, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp009
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µchild_context = πTemp010
								// line 828: yield False
								πF.SetLineno(828)
								if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(30)
								return πTemp009, nil
							Label30:
								πTemp010 = πSent
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp012, πE = πg.IsTrue(πF, πTemp010); πE != nil {
									continue
								}
								if πTemp012 {
									goto Label31
								}
								goto Label32
								// line 829: if child_context.has_matched:
								πF.SetLineno(829)
							Label31:
								// line 830: ctx.has_matched = True
								πF.SetLineno(830)
								if πTemp010, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp010); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp011); πE != nil {
									continue
								}
								// line 831: yield True
								πF.SetLineno(831)
								if πTemp010, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(33)
								return πTemp010, nil
							Label33:
								πTemp011 = πSent
								goto Label32
							Label32:
								// line 832: ctx.skip_char(-1)
								πF.SetLineno(832)
								πTemp001 = πF.MakeArgs(1)
								if πTemp011, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp011
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
									continue
								}
								if πTemp016, πE = πTemp011.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								// line 833: count -= 1
								πF.SetLineno(833)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.ISub(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µcount = πTemp011
								// line 834: ctx.state.marks_pop_keep()
								πF.SetLineno(834)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp016, πE = πg.GetAttr(πF, πTemp011, ßmarks_pop_keep, nil); πE != nil {
									continue
								}
								if πTemp011, πE = πTemp016.Call(πF, nil, nil); πE != nil {
									continue
								}
								continue
							Label28:
								if πE != nil || πR != nil {
									continue
								}
							Label29:
								goto Label12
							Label12:
								// line 836: ctx.state.marks_pop_discard()
								πF.SetLineno(836)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp016, πE = πg.GetAttr(πF, πTemp011, ßmarks_pop_discard, nil); πE != nil {
									continue
								}
								if πTemp011, πE = πTemp016.Call(πF, nil, nil); πE != nil {
									continue
								}
								// line 837: ctx.has_matched = False
								πF.SetLineno(837)
								if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp011); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp016); πE != nil {
									continue
								}
								// line 838: yield True
								πF.SetLineno(838)
								if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(34)
								return πTemp011, nil
							Label34:
								πTemp016 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßop_repeat_one.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 840: def op_min_repeat_one(self, ctx):
					πF.SetLineno(840)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("op_min_repeat_one", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µmincount *πg.Object = πg.UnboundLocal; _ = µmincount
						var µmaxcount *πg.Object = πg.UnboundLocal; _ = µmaxcount
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µchild_context *πg.Object = πg.UnboundLocal; _ = µchild_context
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 3: goto Label3
								case 9: goto Label9
								case 12: goto Label12
								case 13: goto Label13
								case 14: goto Label14
								case 17: goto Label17
								case 20: goto Label20
								case 23: goto Label23
								default: panic("unexpected function state")
								}
								// line 843: mincount = ctx.peek_code(2)
								πF.SetLineno(843)
								πTemp001 = πF.MakeArgs(1)
								πTemp001[0] = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µmincount = πTemp003
								// line 844: maxcount = ctx.peek_code(3)
								πF.SetLineno(844)
								πTemp001 = πF.MakeArgs(1)
								πTemp001[0] = πg.NewInt(3).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µmaxcount = πTemp003
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µctx, ßremaining_chars, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.LT(πF, πTemp004, µmincount); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label1
								}
								goto Label2
								// line 847: if ctx.remaining_chars() < mincount:
								πF.SetLineno(847)
							Label1:
								// line 848: ctx.has_matched = False
								πF.SetLineno(848)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp003); πE != nil {
									continue
								}
								// line 849: yield True
								πF.SetLineno(849)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(3)
								return πTemp002, nil
							Label3:
								πTemp003 = πSent
								goto Label2
							Label2:
								// line 850: ctx.state.string_position = ctx.string_position
								πF.SetLineno(850)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp006, ßstring_position, πTemp004); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Eq(πF, µmincount, πg.NewInt(0).ToObject()); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label4
								}
								goto Label5
								// line 851: if mincount == 0:
								πF.SetLineno(851)
							Label4:
								// line 852: count = 0
								πF.SetLineno(852)
								µcount = πg.NewInt(0).ToObject()
								goto Label6
							Label5:
								// line 854: count = self.count_repetitions(ctx, mincount)
								πF.SetLineno(854)
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								πTemp001[0] = µctx
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								πTemp001[1] = µmincount
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µself, ßcount_repetitions, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µcount = πTemp004
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.LT(πF, µcount, µmincount); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label7
								}
								goto Label8
								// line 855: if count < mincount:
								πF.SetLineno(855)
							Label7:
								// line 856: ctx.has_matched = False
								πF.SetLineno(856)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp004); πE != nil {
									continue
								}
								// line 857: yield True
								πF.SetLineno(857)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(9)
								return πTemp003, nil
							Label9:
								πTemp004 = πSent
								goto Label8
							Label8:
								// line 858: ctx.skip_char(count)
								πF.SetLineno(858)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								πTemp001[0] = µcount
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								goto Label6
							Label6:
								πTemp001 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp007[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp006, πE = πg.Add(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp006
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								πTemp006 = ßsuccess.ToObject()
								if πTemp010, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetItem(πF, πTemp010, πTemp006); πE != nil {
									continue
								}
								if πTemp004, πE = πg.Eq(πF, πTemp008, πTemp009); πE != nil {
									continue
								}
								if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label10
								}
								goto Label11
								// line 859: if ctx.peek_code(ctx.peek_code(1) + 1) == OPCODES["success"]:
								πF.SetLineno(859)
							Label10:
								// line 861: ctx.state.string_position = ctx.string_position
								πF.SetLineno(861)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp004); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp008, ßstring_position, πTemp006); πE != nil {
									continue
								}
								// line 862: ctx.has_matched = True
								πF.SetLineno(862)
								if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp004); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp006); πE != nil {
									continue
								}
								// line 863: yield True
								πF.SetLineno(863)
								if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(12)
								return πTemp004, nil
							Label12:
								πTemp006 = πSent
								goto Label11
							Label11:
								// line 865: ctx.state.marks_push()
								πF.SetLineno(865)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßmarks_push, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp008.Call(πF, nil, nil); πE != nil {
									continue
								}
								// line 866: while maxcount == MAXREPEAT or count <= maxcount:
								πF.SetLineno(866)
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
								if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Eq(πF, µmaxcount, πTemp009); πE != nil {
									continue
								}
								πTemp006 = πTemp008
								if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp012 {
									goto Label16
								}
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.LE(πF, µcount, µmaxcount); πE != nil {
									continue
								}
								πTemp006 = πTemp008
							Label16:
								if πTemp011, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πE != nil || !πTemp011 {
									continue
								}
								πF.PushCheckpoint(13)            
								// line 867: ctx.state.string_position = ctx.string_position
								πF.SetLineno(867)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp006); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp009, ßstring_position, πTemp008); πE != nil {
									continue
								}
								// line 868: child_context = ctx.push_new_context(ctx.peek_code(1) + 1)
								πF.SetLineno(868)
								πTemp001 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp007[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp006, πE = πg.Add(πF, πTemp009, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp006
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µchild_context = πTemp008
								// line 869: yield False
								πF.SetLineno(869)
								if πTemp006, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(17)
								return πTemp006, nil
							Label17:
								πTemp008 = πSent
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
									continue
								}
								if πTemp011 {
									goto Label18
								}
								goto Label19
								// line 870: if child_context.has_matched:
								πF.SetLineno(870)
							Label18:
								// line 871: ctx.has_matched = True
								πF.SetLineno(871)
								if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp008); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp009); πE != nil {
									continue
								}
								// line 872: yield True
								πF.SetLineno(872)
								if πTemp008, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(20)
								return πTemp008, nil
							Label20:
								πTemp009 = πSent
								goto Label19
							Label19:
								// line 873: ctx.state.string_position = ctx.string_position
								πF.SetLineno(873)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp013, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp013, ßstring_position, πTemp010); πE != nil {
									continue
								}
								πTemp001 = πF.MakeArgs(2)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								πTemp001[0] = µctx
								πTemp001[1] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, µself, ßcount_repetitions, nil); πE != nil {
									continue
								}
								if πTemp013, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								if πTemp009, πE = πg.Eq(πF, πTemp013, πg.NewInt(0).ToObject()); πE != nil {
									continue
								}
								if πTemp011, πE = πg.IsTrue(πF, πTemp009); πE != nil {
									continue
								}
								if πTemp011 {
									goto Label21
								}
								goto Label22
								// line 874: if self.count_repetitions(ctx, 1) == 0:
								πF.SetLineno(874)
							Label21:
								// line 875: break
								πF.SetLineno(875)
								πTemp005 = true
								continue
								goto Label22
							Label22:
								// line 876: ctx.skip_char(1)
								πF.SetLineno(876)
								πTemp001 = πF.MakeArgs(1)
								πTemp001[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								// line 877: count += 1
								πF.SetLineno(877)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.IAdd(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µcount = πTemp009
								// line 878: ctx.state.marks_pop_keep()
								πF.SetLineno(878)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßmarks_pop_keep, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp010.Call(πF, nil, nil); πE != nil {
									continue
								}
								continue
							Label14:
								if πE != nil || πR != nil {
									continue
								}
							Label15:
								// line 880: ctx.state.marks_pop_discard()
								πF.SetLineno(880)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßmarks_pop_discard, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp010.Call(πF, nil, nil); πE != nil {
									continue
								}
								// line 881: ctx.has_matched = False
								πF.SetLineno(881)
								if πTemp009, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp010); πE != nil {
									continue
								}
								// line 882: yield True
								πF.SetLineno(882)
								if πTemp009, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(23)
								return πTemp009, nil
							Label23:
								πTemp010 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßop_min_repeat_one.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 884: def op_repeat(self, ctx):
					πF.SetLineno(884)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("op_repeat", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µrepeat *πg.Object = πg.UnboundLocal; _ = µrepeat
						var µchild_context *πg.Object = πg.UnboundLocal; _ = µchild_context
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								default: panic("unexpected function state")
								}
								// line 889: repeat = _RepeatContext(ctx)
								πF.SetLineno(889)
								πTemp001 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								πTemp001[0] = µctx
								if πTemp002, πE = πg.ResolveGlobal(πF, ß_RepeatContext); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µrepeat = πTemp003
								// line 890: ctx.state.repeat = repeat
								πF.SetLineno(890)
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µrepeat); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp003, ßrepeat, πTemp002); πE != nil {
									continue
								}
								// line 891: ctx.state.string_position = ctx.string_position
								πF.SetLineno(891)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp004, ßstring_position, πTemp003); πE != nil {
									continue
								}
								// line 892: child_context = ctx.push_new_context(ctx.peek_code(1) + 1)
								πF.SetLineno(892)
								πTemp001 = πF.MakeArgs(1)
								πTemp005 = πF.MakeArgs(1)
								πTemp005[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp005)
								if πTemp002, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp001[0] = πTemp002
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp001)
								µchild_context = πTemp003
								// line 893: yield False
								πF.SetLineno(893)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(1)
								return πTemp002, nil
							Label1:
								πTemp003 = πSent
								// line 894: ctx.state.repeat = repeat.previous
								πF.SetLineno(894)
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µrepeat, ßprevious, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp006, ßrepeat, πTemp004); πE != nil {
									continue
								}
								// line 895: ctx.has_matched = child_context.has_matched
								πF.SetLineno(895)
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp004); πE != nil {
									continue
								}
								// line 896: yield True
								πF.SetLineno(896)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								return πTemp003, nil
							Label2:
								πTemp004 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßop_repeat.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 898: def op_max_until(self, ctx):
					πF.SetLineno(898)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("op_max_until", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µrepeat *πg.Object = πg.UnboundLocal; _ = µrepeat
						var µmincount *πg.Object = πg.UnboundLocal; _ = µmincount
						var µmaxcount *πg.Object = πg.UnboundLocal; _ = µmaxcount
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µchild_context *πg.Object = πg.UnboundLocal; _ = µchild_context
						var µsave_last_position *πg.Object = πg.UnboundLocal; _ = µsave_last_position
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 5: goto Label5
								case 8: goto Label8
								case 13: goto Label13
								case 16: goto Label16
								case 17: goto Label17
								case 20: goto Label20
								default: panic("unexpected function state")
								}
								// line 901: repeat = ctx.state.repeat
								πF.SetLineno(901)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrepeat, nil); πE != nil {
									continue
								}
								µrepeat = πTemp002
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µrepeat == πTemp002).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label1
								}
								goto Label2
								// line 902: if repeat is None:
								πF.SetLineno(902)
							Label1:
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewStr("Internal re error: MAX_UNTIL without REPEAT.").ToObject()
								if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								// line 903: raise RuntimeError("Internal re error: MAX_UNTIL without REPEAT.")
								πF.SetLineno(903)
								πE = πF.Raise(πTemp002, nil, nil)
								continue
								goto Label2
							Label2:
								// line 904: mincount = repeat.peek_code(2)
								πF.SetLineno(904)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µrepeat, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µmincount = πTemp002
								// line 905: maxcount = repeat.peek_code(3)
								πF.SetLineno(905)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(3).ToObject()
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µrepeat, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µmaxcount = πTemp002
								// line 906: ctx.state.string_position = ctx.string_position
								πF.SetLineno(906)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp005, ßstring_position, πTemp002); πE != nil {
									continue
								}
								// line 907: count = repeat.count + 1
								πF.SetLineno(907)
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µrepeat, ßcount, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µcount = πTemp001
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.LT(πF, µcount, µmincount); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label3
								}
								goto Label4
								// line 910: if count < mincount:
								πF.SetLineno(910)
							Label3:
								// line 912: repeat.count = count
								πF.SetLineno(912)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcount); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßcount, πTemp001); πE != nil {
									continue
								}
								// line 913: child_context = repeat.push_new_context(4)
								πF.SetLineno(913)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(4).ToObject()
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µrepeat, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchild_context = πTemp002
								// line 914: yield False
								πF.SetLineno(914)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								return πTemp001, nil
							Label5:
								πTemp002 = πSent
								// line 915: ctx.has_matched = child_context.has_matched
								πF.SetLineno(915)
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp005); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(!πTemp003).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label6
								}
								goto Label7
								// line 916: if not ctx.has_matched:
								πF.SetLineno(916)
							Label6:
								// line 917: repeat.count = count - 1
								πF.SetLineno(917)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Sub(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßcount, πTemp005); πE != nil {
									continue
								}
								// line 918: ctx.state.string_position = ctx.string_position
								πF.SetLineno(918)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp006, ßstring_position, πTemp005); πE != nil {
									continue
								}
								goto Label7
							Label7:
								// line 919: yield True
								πF.SetLineno(919)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(8)
								return πTemp002, nil
							Label8:
								πTemp005 = πSent
								goto Label4
							Label4:
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.LT(πF, µcount, µmaxcount); πE != nil {
									continue
								}
								πTemp006 = πTemp008
								if πTemp007, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp007 {
									goto Label10
								}
								if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Eq(πF, µmaxcount, πTemp009); πE != nil {
									continue
								}
								πTemp006 = πTemp008
							Label10:
								πTemp005 = πTemp006
								if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if !πTemp003 {
									goto Label9
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µrepeat, ßlast_position, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πg.NE(πF, πTemp009, πTemp008); πE != nil {
									continue
								}
								πTemp005 = πTemp006
							Label9:
								if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label11
								}
								goto Label12
								// line 921: if (count < maxcount or maxcount == MAXREPEAT) \
								πF.SetLineno(921)
							Label11:
								// line 924: repeat.count = count
								πF.SetLineno(924)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, µcount); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßcount, πTemp005); πE != nil {
									continue
								}
								// line 925: ctx.state.marks_push()
								πF.SetLineno(925)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmarks_push, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
									continue
								}
								// line 926: save_last_position = repeat.last_position # zero-width match protection
								πF.SetLineno(926)
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µrepeat, ßlast_position, nil); πE != nil {
									continue
								}
								µsave_last_position = πTemp005
								// line 927: repeat.last_position = ctx.state.string_position
								πF.SetLineno(927)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp006); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßlast_position, πTemp005); πE != nil {
									continue
								}
								// line 928: child_context = repeat.push_new_context(4)
								πF.SetLineno(928)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(4).ToObject()
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µrepeat, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchild_context = πTemp006
								// line 929: yield False
								πF.SetLineno(929)
								if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(13)
								return πTemp005, nil
							Label13:
								πTemp006 = πSent
								// line 930: repeat.last_position = save_last_position
								πF.SetLineno(930)
								if πE = πg.CheckLocal(πF, µsave_last_position, "save_last_position"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µsave_last_position); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßlast_position, πTemp006); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label14
								}
								goto Label15
								// line 931: if child_context.has_matched:
								πF.SetLineno(931)
							Label14:
								// line 932: ctx.state.marks_pop_discard()
								πF.SetLineno(932)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßmarks_pop_discard, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp008.Call(πF, nil, nil); πE != nil {
									continue
								}
								// line 933: ctx.has_matched = True
								πF.SetLineno(933)
								if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp006); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp008); πE != nil {
									continue
								}
								// line 934: yield True
								πF.SetLineno(934)
								if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(16)
								return πTemp006, nil
							Label16:
								πTemp008 = πSent
								goto Label15
							Label15:
								// line 935: ctx.state.marks_pop()
								πF.SetLineno(935)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßmarks_pop, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp009.Call(πF, nil, nil); πE != nil {
									continue
								}
								// line 936: repeat.count = count - 1
								πF.SetLineno(936)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.Sub(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp008); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßcount, πTemp009); πE != nil {
									continue
								}
								// line 937: ctx.state.string_position = ctx.string_position
								πF.SetLineno(937)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp008); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp010, ßstring_position, πTemp009); πE != nil {
									continue
								}
								goto Label12
							Label12:
								// line 940: ctx.state.repeat = repeat.previous
								πF.SetLineno(940)
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µrepeat, ßprevious, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp008); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp010, ßrepeat, πTemp009); πE != nil {
									continue
								}
								// line 941: child_context = ctx.push_new_context(1)
								πF.SetLineno(941)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchild_context = πTemp009
								// line 942: yield False
								πF.SetLineno(942)
								if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(17)
								return πTemp008, nil
							Label17:
								πTemp009 = πSent
								// line 943: ctx.has_matched = child_context.has_matched
								πF.SetLineno(943)
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp010); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, µctx, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp010); πE != nil {
									continue
								}
								πTemp009 = πg.GetBool(!πTemp003).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp009); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label18
								}
								goto Label19
								// line 944: if not ctx.has_matched:
								πF.SetLineno(944)
							Label18:
								// line 945: ctx.state.repeat = repeat
								πF.SetLineno(945)
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, µrepeat); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp010, ßrepeat, πTemp009); πE != nil {
									continue
								}
								// line 946: ctx.state.string_position = ctx.string_position
								πF.SetLineno(946)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp011, ßstring_position, πTemp010); πE != nil {
									continue
								}
								goto Label19
							Label19:
								// line 947: yield True
								πF.SetLineno(947)
								if πTemp009, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(20)
								return πTemp009, nil
							Label20:
								πTemp010 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßop_max_until.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 949: def op_min_until(self, ctx):
					πF.SetLineno(949)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("op_min_until", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µrepeat *πg.Object = πg.UnboundLocal; _ = µrepeat
						var µmincount *πg.Object = πg.UnboundLocal; _ = µmincount
						var µmaxcount *πg.Object = πg.UnboundLocal; _ = µmaxcount
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µchild_context *πg.Object = πg.UnboundLocal; _ = µchild_context
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
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
								case 5: goto Label5
								case 8: goto Label8
								case 9: goto Label9
								case 12: goto Label12
								case 16: goto Label16
								case 17: goto Label17
								case 20: goto Label20
								default: panic("unexpected function state")
								}
								// line 952: repeat = ctx.state.repeat
								πF.SetLineno(952)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßrepeat, nil); πE != nil {
									continue
								}
								µrepeat = πTemp002
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
									continue
								}
								πTemp001 = πg.GetBool(µrepeat == πTemp002).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label1
								}
								goto Label2
								// line 953: if repeat is None:
								πF.SetLineno(953)
							Label1:
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewStr("Internal re error: MIN_UNTIL without REPEAT.").ToObject()
								if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								// line 954: raise RuntimeError("Internal re error: MIN_UNTIL without REPEAT.")
								πF.SetLineno(954)
								πE = πF.Raise(πTemp002, nil, nil)
								continue
								goto Label2
							Label2:
								// line 955: mincount = repeat.peek_code(2)
								πF.SetLineno(955)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µrepeat, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µmincount = πTemp002
								// line 956: maxcount = repeat.peek_code(3)
								πF.SetLineno(956)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(3).ToObject()
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µrepeat, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µmaxcount = πTemp002
								// line 957: ctx.state.string_position = ctx.string_position
								πF.SetLineno(957)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp005, ßstring_position, πTemp002); πE != nil {
									continue
								}
								// line 958: count = repeat.count + 1
								πF.SetLineno(958)
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µrepeat, ßcount, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								µcount = πTemp001
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmincount, "mincount"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.LT(πF, µcount, µmincount); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label3
								}
								goto Label4
								// line 961: if count < mincount:
								πF.SetLineno(961)
							Label3:
								// line 963: repeat.count = count
								πF.SetLineno(963)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcount); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßcount, πTemp001); πE != nil {
									continue
								}
								// line 964: child_context = repeat.push_new_context(4)
								πF.SetLineno(964)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(4).ToObject()
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µrepeat, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchild_context = πTemp002
								// line 965: yield False
								πF.SetLineno(965)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(5)
								return πTemp001, nil
							Label5:
								πTemp002 = πSent
								// line 966: ctx.has_matched = child_context.has_matched
								πF.SetLineno(966)
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp005); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp005); πE != nil {
									continue
								}
								πTemp002 = πg.GetBool(!πTemp003).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label6
								}
								goto Label7
								// line 967: if not ctx.has_matched:
								πF.SetLineno(967)
							Label6:
								// line 968: repeat.count = count - 1
								πF.SetLineno(968)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.Sub(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßcount, πTemp005); πE != nil {
									continue
								}
								// line 969: ctx.state.string_position = ctx.string_position
								πF.SetLineno(969)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp006, ßstring_position, πTemp005); πE != nil {
									continue
								}
								goto Label7
							Label7:
								// line 970: yield True
								πF.SetLineno(970)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(8)
								return πTemp002, nil
							Label8:
								πTemp005 = πSent
								goto Label4
							Label4:
								// line 973: ctx.state.marks_push()
								πF.SetLineno(973)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmarks_push, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp006.Call(πF, nil, nil); πE != nil {
									continue
								}
								// line 974: ctx.state.repeat = repeat.previous
								πF.SetLineno(974)
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µrepeat, ßprevious, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp005); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp007, ßrepeat, πTemp006); πE != nil {
									continue
								}
								// line 975: child_context = ctx.push_new_context(1)
								πF.SetLineno(975)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchild_context = πTemp006
								// line 976: yield False
								πF.SetLineno(976)
								if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(9)
								return πTemp005, nil
							Label9:
								πTemp006 = πSent
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp006, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label10
								}
								goto Label11
								// line 977: if child_context.has_matched:
								πF.SetLineno(977)
							Label10:
								// line 978: ctx.has_matched = True
								πF.SetLineno(978)
								if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp006); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp007); πE != nil {
									continue
								}
								// line 979: yield True
								πF.SetLineno(979)
								if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(12)
								return πTemp006, nil
							Label12:
								πTemp007 = πSent
								goto Label11
							Label11:
								// line 980: ctx.state.repeat = repeat
								πF.SetLineno(980)
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, µrepeat); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp008, ßrepeat, πTemp007); πE != nil {
									continue
								}
								// line 981: ctx.state.string_position = ctx.string_position
								πF.SetLineno(981)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp007); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp009, ßstring_position, πTemp008); πE != nil {
									continue
								}
								// line 982: ctx.state.marks_pop()
								πF.SetLineno(982)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp007, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßmarks_pop, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp008.Call(πF, nil, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GE(πF, µcount, µmaxcount); πE != nil {
									continue
								}
								πTemp007 = πTemp008
								if πTemp003, πE = πg.IsTrue(πF, πTemp007); πE != nil {
									continue
								}
								if !πTemp003 {
									goto Label13
								}
								if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
									continue
								}
								if πTemp008, πE = πg.NE(πF, µmaxcount, πTemp009); πE != nil {
									continue
								}
								πTemp007 = πTemp008
							Label13:
								if πTemp003, πE = πg.IsTrue(πF, πTemp007); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label14
								}
								goto Label15
								// line 985: if count >= maxcount and maxcount != MAXREPEAT:
								πF.SetLineno(985)
							Label14:
								// line 986: ctx.has_matched = False
								πF.SetLineno(986)
								if πTemp007, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp007); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp008); πE != nil {
									continue
								}
								// line 987: yield True
								πF.SetLineno(987)
								if πTemp007, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(16)
								return πTemp007, nil
							Label16:
								πTemp008 = πSent
								goto Label15
							Label15:
								// line 988: repeat.count = count
								πF.SetLineno(988)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, µcount); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßcount, πTemp008); πE != nil {
									continue
								}
								// line 989: child_context = repeat.push_new_context(4)
								πF.SetLineno(989)
								πTemp004 = πF.MakeArgs(1)
								πTemp004[0] = πg.NewInt(4).ToObject()
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πTemp008, πE = πg.GetAttr(πF, µrepeat, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp009, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								µchild_context = πTemp009
								// line 990: yield False
								πF.SetLineno(990)
								if πTemp008, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(17)
								return πTemp008, nil
							Label17:
								πTemp009 = πSent
								// line 991: ctx.has_matched = child_context.has_matched
								πF.SetLineno(991)
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp010); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp010, πE = πg.GetAttr(πF, µctx, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πg.IsTrue(πF, πTemp010); πE != nil {
									continue
								}
								πTemp009 = πg.GetBool(!πTemp003).ToObject()
								if πTemp003, πE = πg.IsTrue(πF, πTemp009); πE != nil {
									continue
								}
								if πTemp003 {
									goto Label18
								}
								goto Label19
								// line 992: if not ctx.has_matched:
								πF.SetLineno(992)
							Label18:
								// line 993: repeat.count = count - 1
								πF.SetLineno(993)
								if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.Sub(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µrepeat, "repeat"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µrepeat, ßcount, πTemp010); πE != nil {
									continue
								}
								// line 994: ctx.state.string_position = ctx.string_position
								πF.SetLineno(994)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp009, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp011, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp011, ßstring_position, πTemp010); πE != nil {
									continue
								}
								goto Label19
							Label19:
								// line 995: yield True
								πF.SetLineno(995)
								if πTemp009, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(20)
								return πTemp009, nil
							Label20:
								πTemp010 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßop_min_until.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 997: def general_op_groupref(self, ctx, decorate=lambda x: x):
					πF.SetLineno(997)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp009 = make([]πg.Param, 1)
					πTemp009[0] = πg.Param{Name: "x", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_sre.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πArgs[0]; _ = µx
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 997: def general_op_groupref(self, ctx, decorate=lambda x: x):
							πF.SetLineno(997)
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
					πTemp002[2] = πg.Param{Name: "decorate", Def: πTemp029}
					πTemp028 = πg.NewFunction(πg.NewCode("general_op_groupref", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µdecorate *πg.Object = πArgs[2]; _ = µdecorate
						var µgroup_start *πg.Object = πg.UnboundLocal; _ = µgroup_start
						var µgroup_end *πg.Object = πg.UnboundLocal; _ = µgroup_end
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 998: group_start, group_end = ctx.state.get_marks(ctx.peek_code(1))
							πF.SetLineno(998)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget_marks, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µgroup_start = πTemp004
							µgroup_end = πTemp005
							if πE = πg.CheckLocal(πF, µgroup_start, "group_start"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µgroup_start == πTemp005).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µgroup_end, "group_end"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µgroup_end == πTemp005).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µgroup_end, "group_end"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µgroup_start, "group_start"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µgroup_end, µgroup_start); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label1:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							goto Label3
							// line 999: if group_start is None or group_end is None or group_end < group_start:
							πF.SetLineno(999)
						Label2:
							// line 1000: ctx.has_matched = False
							πF.SetLineno(1000)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp004); πE != nil {
								continue
							}
							// line 1001: return True
							πF.SetLineno(1001)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label3
						Label3:
							// line 1002: while group_start < group_end:
							πF.SetLineno(1002)
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
							if πE = πg.CheckLocal(πF, µgroup_start, "group_start"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µgroup_end, "group_end"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µgroup_start, µgroup_end); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)            
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp008
							if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp008
							if πE = πg.CheckLocal(πF, µdecorate, "decorate"); πE != nil {
								continue
							}
							if πTemp005, πE = µdecorate.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgroup_start, "group_start"); πE != nil {
								continue
							}
							πTemp008 = µgroup_start
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßstring, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, πTemp011, πTemp008); πE != nil {
								continue
							}
							πTemp002[0] = πTemp009
							if πTemp008, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp009
							if πE = πg.CheckLocal(πF, µdecorate, "decorate"); πE != nil {
								continue
							}
							if πTemp008, πE = µdecorate.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.NE(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label7:
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label8
							}
							goto Label9
							// line 1003: if ctx.at_end() or decorate(ord(ctx.peek_char())) \
							πF.SetLineno(1003)
						Label8:
							// line 1005: ctx.has_matched = False
							πF.SetLineno(1005)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp004); πE != nil {
								continue
							}
							// line 1006: return True
							πF.SetLineno(1006)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label9
						Label9:
							// line 1007: group_start += 1
							πF.SetLineno(1007)
							if πE = πg.CheckLocal(πF, µgroup_start, "group_start"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µgroup_start, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µgroup_start = πTemp003
							// line 1008: ctx.skip_char(1)
							πF.SetLineno(1008)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßskip_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 1009: ctx.skip_code(2)
							πF.SetLineno(1009)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1010: return True
							πF.SetLineno(1010)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgeneral_op_groupref.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 1012: def op_groupref(self, ctx):
					πF.SetLineno(1012)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("op_groupref", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1016: return self.general_op_groupref(ctx)
							πF.SetLineno(1016)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[0] = µctx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgeneral_op_groupref, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_groupref.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 1018: def op_groupref_ignore(self, ctx):
					πF.SetLineno(1018)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("op_groupref_ignore", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1022: return self.general_op_groupref(ctx, ctx.state.lower)
							πF.SetLineno(1022)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[0] = µctx
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßlower, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßgeneral_op_groupref, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_groupref_ignore.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 1024: def op_groupref_exists(self, ctx):
					πF.SetLineno(1024)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("op_groupref_exists", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µgroup_start *πg.Object = πg.UnboundLocal; _ = µgroup_start
						var µgroup_end *πg.Object = πg.UnboundLocal; _ = µgroup_end
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
						var πTemp006 bool
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1027: group_start, group_end = ctx.state.get_marks(ctx.peek_code(1))
							πF.SetLineno(1027)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget_marks, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
								continue
							}
							µgroup_start = πTemp004
							µgroup_end = πTemp005
							if πE = πg.CheckLocal(πF, µgroup_start, "group_start"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µgroup_start == πTemp005).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µgroup_end, "group_end"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(µgroup_end == πTemp005).ToObject()
							πTemp003 = πTemp004
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µgroup_end, "group_end"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µgroup_start, "group_start"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LT(πF, µgroup_end, µgroup_start); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label1:
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							goto Label3
							// line 1028: if group_start is None or group_end is None or group_end < group_start:
							πF.SetLineno(1028)
						Label2:
							// line 1029: ctx.skip_code(ctx.peek_code(2) + 1)
							πF.SetLineno(1029)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Add(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label3:
							// line 1031: ctx.skip_code(3)
							πF.SetLineno(1031)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label4:
							// line 1032: return True
							πF.SetLineno(1032)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
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
					if πE = πClass.SetItem(πF, ßop_groupref_exists.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 1034: def op_assert(self, ctx):
					πF.SetLineno(1034)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("op_assert", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µchild_context *πg.Object = πg.UnboundLocal; _ = µchild_context
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
								case 8: goto Label8
								case 3: goto Label3
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								// line 1038: ctx.state.string_position = ctx.string_position - ctx.peek_code(2)
								πF.SetLineno(1038)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp005); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp004, ßstring_position, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßstring_position, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.LT(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label1
								}
								goto Label2
								// line 1039: if ctx.state.string_position < 0:
								πF.SetLineno(1039)
							Label1:
								// line 1040: ctx.has_matched = False
								πF.SetLineno(1040)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp002); πE != nil {
									continue
								}
								// line 1041: yield True
								πF.SetLineno(1041)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(3)
								return πTemp001, nil
							Label3:
								πTemp002 = πSent
								goto Label2
							Label2:
								// line 1042: child_context = ctx.push_new_context(3)
								πF.SetLineno(1042)
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(3).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								µchild_context = πTemp004
								// line 1043: yield False
								πF.SetLineno(1043)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp002, nil
							Label4:
								πTemp004 = πSent
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label5
								}
								goto Label6
								// line 1044: if child_context.has_matched:
								πF.SetLineno(1044)
							Label5:
								// line 1045: ctx.skip_code(ctx.peek_code(1) + 1)
								πF.SetLineno(1045)
								πTemp003 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp007[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp004, πE = πg.Add(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp003[0] = πTemp004
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								goto Label7
							Label6:
								// line 1047: ctx.has_matched = False
								πF.SetLineno(1047)
								if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp005); πE != nil {
									continue
								}
								goto Label7
							Label7:
								// line 1048: yield True
								πF.SetLineno(1048)
								if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(8)
								return πTemp004, nil
							Label8:
								πTemp005 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßop_assert.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 1050: def op_assert_not(self, ctx):
					πF.SetLineno(1050)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("op_assert_not", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µchild_context *πg.Object = πg.UnboundLocal; _ = µchild_context
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
								case 3: goto Label3
								case 6: goto Label6
								case 7: goto Label7
								default: panic("unexpected function state")
								}
								// line 1054: ctx.state.string_position = ctx.string_position - ctx.peek_code(2)
								πF.SetLineno(1054)
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
									continue
								}
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(2).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp005); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, πTemp004, ßstring_position, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßstring_position, nil); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GE(πF, πTemp004, πg.NewInt(0).ToObject()); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label1
								}
								goto Label2
								// line 1055: if ctx.state.string_position >= 0:
								πF.SetLineno(1055)
							Label1:
								// line 1056: child_context = ctx.push_new_context(3)
								πF.SetLineno(1056)
								πTemp003 = πF.MakeArgs(1)
								πTemp003[0] = πg.NewInt(3).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.GetAttr(πF, µctx, ßpush_new_context, nil); πE != nil {
									continue
								}
								if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								µchild_context = πTemp002
								// line 1057: yield False
								πF.SetLineno(1057)
								if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								πF.PushCheckpoint(3)
								return πTemp001, nil
							Label3:
								πTemp002 = πSent
								if πE = πg.CheckLocal(πF, µchild_context, "child_context"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µchild_context, ßhas_matched, nil); πE != nil {
									continue
								}
								if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
									continue
								}
								if πTemp006 {
									goto Label4
								}
								goto Label5
								// line 1058: if child_context.has_matched:
								πF.SetLineno(1058)
							Label4:
								// line 1059: ctx.has_matched = False
								πF.SetLineno(1059)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
									continue
								}
								if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp004); πE != nil {
									continue
								}
								// line 1060: yield True
								πF.SetLineno(1060)
								if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return πTemp002, nil
							Label6:
								πTemp004 = πSent
								goto Label5
							Label5:
								goto Label2
							Label2:
								// line 1061: ctx.skip_code(ctx.peek_code(1) + 1)
								πF.SetLineno(1061)
								πTemp003 = πF.MakeArgs(1)
								πTemp007 = πF.MakeArgs(1)
								πTemp007[0] = πg.NewInt(1).ToObject()
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp005, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
									continue
								}
								if πTemp008, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp007)
								if πTemp004, πE = πg.Add(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
									continue
								}
								πTemp003[0] = πTemp004
								if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp003)
								// line 1062: yield True
								πF.SetLineno(1062)
								if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								return πTemp004, nil
							Label7:
								πTemp005 = πSent
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßop_assert_not.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 1064: def unknown(self, ctx):
					πF.SetLineno(1064)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("unknown", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Internal re error. Unknown opcode: %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1066: raise RuntimeError("Internal re error. Unknown opcode: %s" % ctx.peek_code())
							πF.SetLineno(1066)
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
					if πE = πClass.SetItem(πF, ßunknown.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 1068: def check_charset(self, ctx, char):
					πF.SetLineno(1068)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp002[2] = πg.Param{Name: "char", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("check_charset", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µchar *πg.Object = πArgs[2]; _ = µchar
						var µsave_position *πg.Object = πg.UnboundLocal; _ = µsave_position
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 1069: """Checks whether a character matches set of arbitrary length. Assumes
							πF.SetLineno(1069)
							// line 1071: self.set_dispatcher.reset(char)
							πF.SetLineno(1071)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp001[0] = µchar
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_dispatcher, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßreset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1072: save_position = ctx.code_position
							πF.SetLineno(1072)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßcode_position, nil); πE != nil {
								continue
							}
							µsave_position = πTemp002
							// line 1073: result = None
							πF.SetLineno(1073)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µresult = πTemp002
							// line 1074: while result is None:
							πF.SetLineno(1074)
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
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µresult == πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1075: result = self.set_dispatcher.dispatch(ctx.peek_code(), ctx)
							πF.SetLineno(1075)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[1] = µctx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_dispatcher, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp002
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1076: ctx.code_position = save_position
							πF.SetLineno(1076)
							if πE = πg.CheckLocal(πF, µsave_position, "save_position"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µsave_position); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßcode_position, πTemp002); πE != nil {
								continue
							}
							// line 1077: return result
							πF.SetLineno(1077)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πR = µresult
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcheck_charset.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 1069: """Checks whether a character matches set of arbitrary length. Assumes
					πF.SetLineno(1069)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πg.NewStr("Checks whether a character matches set of arbitrary length. Assumes\n        the code pointer is at the first member of the set.").ToObject()); πE != nil {
						continue
					}
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_charset); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp037, ß__doc__, πTemp036); πE != nil {
						continue
					}
					// line 1079: def count_repetitions(self, ctx, maxcount):
					πF.SetLineno(1079)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp002[2] = πg.Param{Name: "maxcount", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("count_repetitions", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µmaxcount *πg.Object = πArgs[2]; _ = µmaxcount
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µreal_maxcount *πg.Object = πg.UnboundLocal; _ = µreal_maxcount
						var µcode_position *πg.Object = πg.UnboundLocal; _ = µcode_position
						var µstring_position *πg.Object = πg.UnboundLocal; _ = µstring_position
						var µreset_position *πg.Object = πg.UnboundLocal; _ = µreset_position
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 1080: """Returns the number of repetitions of a single item, starting from the
							πF.SetLineno(1080)
							// line 1083: count = 0
							πF.SetLineno(1083)
							µcount = πg.NewInt(0).ToObject()
							// line 1084: real_maxcount = ctx.state.end - ctx.string_position
							πF.SetLineno(1084)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßstate, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßend, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							µreal_maxcount = πTemp001
							if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µreal_maxcount, "real_maxcount"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µmaxcount, µreal_maxcount); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßMAXREPEAT); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, µmaxcount, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label1:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 1085: if maxcount < real_maxcount and maxcount != MAXREPEAT:
							πF.SetLineno(1085)
						Label2:
							// line 1086: real_maxcount = maxcount
							πF.SetLineno(1086)
							if πE = πg.CheckLocal(πF, µmaxcount, "maxcount"); πE != nil {
								continue
							}
							µreal_maxcount = µmaxcount
							goto Label3
						Label3:
							// line 1090: code_position = ctx.code_position
							πF.SetLineno(1090)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßcode_position, nil); πE != nil {
								continue
							}
							µcode_position = πTemp001
							// line 1091: string_position = ctx.string_position
							πF.SetLineno(1091)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßstring_position, nil); πE != nil {
								continue
							}
							µstring_position = πTemp001
							// line 1092: ctx.skip_code(4)
							πF.SetLineno(1092)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 1093: reset_position = ctx.code_position
							πF.SetLineno(1093)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßcode_position, nil); πE != nil {
								continue
							}
							µreset_position = πTemp001
							// line 1094: while count < real_maxcount:
							πF.SetLineno(1094)
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
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µreal_maxcount, "real_maxcount"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µcount, µreal_maxcount); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 1097: ctx.code_position = reset_position
							πF.SetLineno(1097)
							if πE = πg.CheckLocal(πF, µreset_position, "reset_position"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µreset_position); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßcode_position, πTemp001); πE != nil {
								continue
							}
							// line 1098: self.dispatch(ctx.peek_code(), ctx)
							πF.SetLineno(1098)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp005[1] = µctx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßhas_matched, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 1099: if ctx.has_matched is False: # could be None as well
							πF.SetLineno(1099)
						Label7:
							// line 1100: break
							πF.SetLineno(1100)
							πTemp004 = true
							continue
							goto Label8
						Label8:
							// line 1101: count += 1
							πF.SetLineno(1101)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µcount, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µcount = πTemp001
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 1102: ctx.has_matched = None
							πF.SetLineno(1102)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßhas_matched, πTemp002); πE != nil {
								continue
							}
							// line 1103: ctx.code_position = code_position
							πF.SetLineno(1103)
							if πE = πg.CheckLocal(πF, µcode_position, "code_position"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcode_position); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßcode_position, πTemp001); πE != nil {
								continue
							}
							// line 1104: ctx.string_position = string_position
							πF.SetLineno(1104)
							if πE = πg.CheckLocal(πF, µstring_position, "string_position"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstring_position); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µctx, ßstring_position, πTemp001); πE != nil {
								continue
							}
							// line 1105: return count
							πF.SetLineno(1105)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							πR = µcount
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßcount_repetitions.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 1080: """Returns the number of repetitions of a single item, starting from the
					πF.SetLineno(1080)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp037}, πg.NewStr("Returns the number of repetitions of a single item, starting from the\n        current string position. The code pointer is expected to point to a\n        REPEAT_ONE operation (with the repeated 4 ahead).").ToObject()); πE != nil {
						continue
					}
					if πTemp038, πE = πg.ResolveClass(πF, πClass, nil, ßcount_repetitions); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp038, ß__doc__, πTemp037); πE != nil {
						continue
					}
					// line 1107: def _log(self, context, opname, *args):
					πF.SetLineno(1107)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "context", Def: nil}
					πTemp002[2] = πg.Param{Name: "opname", Def: nil}
					πTemp037 = πg.NewFunction(πg.NewCode("_log", "/usr/lib/python2.7/_sre.py", πTemp002, πg.CodeFlagVarArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µcontext *πg.Object = πArgs[1]; _ = µcontext
						var µopname *πg.Object = πArgs[2]; _ = µopname
						var µargs *πg.Object = πArgs[3]; _ = µargs
						var µarg_string *πg.Object = πg.UnboundLocal; _ = µarg_string
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
							// line 1108: arg_string = ("%s " * len(args)) % args
							πF.SetLineno(1108)
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
							if πTemp002, πE = πg.Mul(πF, πg.NewStr("%s ").ToObject(), πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp002, µargs); πE != nil {
								continue
							}
							µarg_string = πTemp001
							// line 1109: _log("|%s|%s|%s %s" % (context.pattern_codes,
							πF.SetLineno(1109)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcontext, ßpattern_codes, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcontext, "context"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µcontext, ßstring_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopname, "opname"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µarg_string, "arg_string"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(πTemp004, πTemp005, µopname, µarg_string).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("|%s|%s|%s %s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_log); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_log.ToObject(), πTemp037); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("_OpcodeDispatcher").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_OpcodeDispatcher.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1112: _OpcodeDispatcher.build_dispatch_table(OPCODES, "op_")
			πF.SetLineno(1112)
			πTemp002 = πF.MakeArgs(2)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp002[1] = ßop_.ToObject()
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_OpcodeDispatcher); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßbuild_dispatch_table, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 1115: class _CharsetDispatcher(_Dispatcher):
			πF.SetLineno(1115)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_Dispatcher); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_CharsetDispatcher", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
					// line 1117: def __init__(self):
					πF.SetLineno(1117)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1118: self.ch_dispatcher = _ChcodeDispatcher()
							πF.SetLineno(1118)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_ChcodeDispatcher); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßch_dispatcher, πTemp001); πE != nil {
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
					// line 1120: def reset(self, char):
					πF.SetLineno(1120)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "char", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("reset", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µchar *πg.Object = πArgs[1]; _ = µchar
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
							// line 1121: self.char = char
							πF.SetLineno(1121)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µchar); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßchar, πTemp001); πE != nil {
								continue
							}
							// line 1122: self.ok = True
							πF.SetLineno(1122)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßok, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßreset.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1124: def set_failure(self, ctx):
					πF.SetLineno(1124)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("set_failure", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1125: return not self.ok
							πF.SetLineno(1125)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßok, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_failure.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1126: def set_literal(self, ctx):
					πF.SetLineno(1126)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("set_literal", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßchar, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1128: if ctx.peek_code(1) == self.char:
							πF.SetLineno(1128)
						Label1:
							// line 1129: return self.ok
							πF.SetLineno(1129)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßok, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label2:
							// line 1131: ctx.skip_code(2)
							πF.SetLineno(1131)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßset_literal.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1132: def set_category(self, ctx):
					πF.SetLineno(1132)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("set_category", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var πTemp001 []*πg.Object
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
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							πTemp001[1] = µctx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßch_dispatcher, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 1134: if self.ch_dispatcher.dispatch(ctx.peek_code(1), ctx):
							πF.SetLineno(1134)
						Label1:
							// line 1135: return self.ok
							πF.SetLineno(1135)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßok, nil); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label3
						Label2:
							// line 1137: ctx.skip_code(2)
							πF.SetLineno(1137)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_category.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1138: def set_charset(self, ctx):
					πF.SetLineno(1138)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("set_charset", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µchar_code *πg.Object = πg.UnboundLocal; _ = µchar_code
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1140: char_code = self.char
							πF.SetLineno(1140)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßchar, nil); πE != nil {
								continue
							}
							µchar_code = πTemp001
							// line 1141: ctx.skip_code(1) # point to beginning of bitmap
							πF.SetLineno(1141)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1142: if CODESIZE == 2:
							πF.SetLineno(1142)
						Label1:
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µchar_code, πg.NewInt(256).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.RShift(πF, µchar_code, πg.NewInt(4).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.And(πF, µchar_code, πg.NewInt(15).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label4:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 1143: if char_code < 256 and ctx.peek_code(char_code >> 4) \
							πF.SetLineno(1143)
						Label5:
							// line 1145: return self.ok
							πF.SetLineno(1145)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßok, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label6
						Label6:
							// line 1146: ctx.skip_code(16) # skip bitmap
							πF.SetLineno(1146)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(16).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label3
						Label2:
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LT(πF, µchar_code, πg.NewInt(256).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label7
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.RShift(πF, µchar_code, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp005
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.And(πF, µchar_code, πg.NewInt(31).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label7:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 1148: if char_code < 256 and ctx.peek_code(char_code >> 5) \
							πF.SetLineno(1148)
						Label8:
							// line 1150: return self.ok
							πF.SetLineno(1150)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßok, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label9
						Label9:
							// line 1151: ctx.skip_code(8) # skip bitmap
							πF.SetLineno(1151)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(8).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
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
					if πE = πClass.SetItem(πF, ßset_charset.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1152: def set_range(self, ctx):
					πF.SetLineno(1152)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("set_range", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßchar, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.LE(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
						Label1:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 1154: if ctx.peek_code(1) <= self.char <= ctx.peek_code(2):
							πF.SetLineno(1154)
						Label2:
							// line 1155: return self.ok
							πF.SetLineno(1155)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßok, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label3:
							// line 1156: ctx.skip_code(3)
							πF.SetLineno(1156)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_range.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1157: def set_negate(self, ctx):
					πF.SetLineno(1157)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("set_negate", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1158: self.ok = not self.ok
							πF.SetLineno(1158)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßok, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßok, πTemp002); πE != nil {
								continue
							}
							// line 1159: ctx.skip_code(1)
							πF.SetLineno(1159)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_negate.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1160: def set_bigcharset(self, ctx):
					πF.SetLineno(1160)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("set_bigcharset", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var µchar_code *πg.Object = πg.UnboundLocal; _ = µchar_code
						var µcount *πg.Object = πg.UnboundLocal; _ = µcount
						var µblock_index *πg.Object = πg.UnboundLocal; _ = µblock_index
						var µa *πg.Object = πg.UnboundLocal; _ = µa
						var µblock *πg.Object = πg.UnboundLocal; _ = µblock
						var µblock_value *πg.Object = πg.UnboundLocal; _ = µblock_value
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
						var πTemp010 bool
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
							default: panic("unexpected function state")
							}
							// line 1162: char_code = self.char
							πF.SetLineno(1162)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßchar, nil); πE != nil {
								continue
							}
							µchar_code = πTemp001
							// line 1163: count = ctx.peek_code(1)
							πF.SetLineno(1163)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µcount = πTemp003
							// line 1164: ctx.skip_code(2)
							πF.SetLineno(1164)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µchar_code, πg.NewInt(65536).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1165: if char_code < 65536:
							πF.SetLineno(1165)
						Label1:
							// line 1166: block_index = char_code >> 8
							πF.SetLineno(1166)
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.RShift(πF, µchar_code, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							µblock_index = πTemp001
							// line 1169: a = []
							πF.SetLineno(1169)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							µa = πTemp001
							// line 1172: a += [ctx.peek_code(block_index // CODESIZE)]
							πF.SetLineno(1172)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							πTemp002 = make([]*πg.Object, 1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock_index, "block_index"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.FloorDiv(πF, µblock_index, πTemp003); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002[0] = πTemp003
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πTemp003, πE = πg.IAdd(πF, µa, πTemp001); πE != nil {
								continue
							}
							µa = πTemp003
							// line 1173: block = a[block_index % CODESIZE]
							πF.SetLineno(1173)
							if πE = πg.CheckLocal(πF, µblock_index, "block_index"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, µblock_index, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µa, πTemp001); πE != nil {
								continue
							}
							µblock = πTemp003
							// line 1174: ctx.skip_code(256 / CODESIZE) # skip block indices
							πF.SetLineno(1174)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Div(πF, πg.NewInt(256).ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1175: block_value = ctx.peek_code(block * (32 / CODESIZE)
							πF.SetLineno(1175)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µblock, "block"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Div(πF, πg.NewInt(32).ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, µblock, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.And(πF, µchar_code, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							if πTemp012, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
								continue
							}
							if πTemp011, πE = πg.Eq(πF, πTemp012, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp009 = πTemp011
							if πTemp010, πE = πg.IsTrue(πF, πTemp009); πE != nil {
								continue
							}
							if !πTemp010 {
								goto Label5
							}
							πTemp009 = πg.NewInt(4).ToObject()
						Label5:
							πTemp008 = πTemp009
							if πTemp004, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							πTemp008 = πg.NewInt(5).ToObject()
						Label4:
							if πTemp006, πE = πg.RShift(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßpeek_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µblock_value = πTemp003
							if πE = πg.CheckLocal(πF, µblock_value, "block_value"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchar_code, "char_code"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mul(πF, πg.NewInt(8).ToObject(), πTemp009); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Sub(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.And(πF, µchar_code, πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp001, πE = πg.And(πF, µblock_value, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 1177: if block_value & (1 << (char_code & ((8 * CODESIZE) - 1))):
							πF.SetLineno(1177)
						Label6:
							// line 1178: return self.ok
							πF.SetLineno(1178)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßok, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label7
						Label7:
							goto Label3
						Label2:
							// line 1180: ctx.skip_code(256 / CODESIZE) # skip block indices
							πF.SetLineno(1180)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Div(πF, πg.NewInt(256).ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label3
						Label3:
							// line 1181: ctx.skip_code(count * (32 / CODESIZE)) # skip blocks
							πF.SetLineno(1181)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcount, "count"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßCODESIZE); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Div(πF, πg.NewInt(32).ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µcount, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßskip_code, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_bigcharset.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1182: def unknown(self, ctx):
					πF.SetLineno(1182)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("unknown", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1183: return False
							πF.SetLineno(1183)
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
					if πE = πClass.SetItem(πF, ßunknown.ToObject(), πTemp011); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("_CharsetDispatcher").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_CharsetDispatcher.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1185: _CharsetDispatcher.build_dispatch_table(OPCODES, "set_")
			πF.SetLineno(1185)
			πTemp002 = πF.MakeArgs(2)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßOPCODES); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp002[1] = ßset_.ToObject()
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_CharsetDispatcher); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßbuild_dispatch_table, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 1188: class _AtcodeDispatcher(_Dispatcher):
			πF.SetLineno(1188)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_Dispatcher); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_AtcodeDispatcher", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1190: def at_beginning(self, ctx):
					πF.SetLineno(1190)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("at_beginning", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1191: return ctx.at_beginning()
							πF.SetLineno(1191)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßat_beginning, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_beginning.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1192: at_beginning_string = at_beginning
					πF.SetLineno(1192)
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßat_beginning); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßat_beginning_string.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1193: def at_beginning_line(self, ctx):
					πF.SetLineno(1193)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("at_beginning_line", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1194: return ctx.at_beginning() or _is_linebreak(ctx.peek_char(-1))
							πF.SetLineno(1194)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_beginning, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_linebreak); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001 = πTemp004
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
					if πE = πClass.SetItem(πF, ßat_beginning_line.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1195: def at_end(self, ctx):
					πF.SetLineno(1195)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("at_end", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1196: return (ctx.remaining_chars() == 1 and ctx.at_linebreak()) or ctx.at_end()
							πF.SetLineno(1196)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µctx, ßremaining_chars, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µctx, ßat_linebreak, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003 = πTemp006
						Label2:
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp005
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
					if πE = πClass.SetItem(πF, ßat_end.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1197: def at_end_line(self, ctx):
					πF.SetLineno(1197)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("at_end_line", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
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
							// line 1198: return ctx.at_linebreak() or ctx.at_end()
							πF.SetLineno(1198)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_linebreak, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
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
					if πE = πClass.SetItem(πF, ßat_end_line.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1199: def at_end_string(self, ctx):
					πF.SetLineno(1199)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("at_end_string", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1200: return ctx.at_end()
							πF.SetLineno(1200)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßat_end, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_end_string.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1201: def at_boundary(self, ctx):
					πF.SetLineno(1201)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("at_boundary", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1202: return ctx.at_boundary(_is_word)
							πF.SetLineno(1202)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_word); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßat_boundary, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_boundary.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1203: def at_non_boundary(self, ctx):
					πF.SetLineno(1203)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("at_non_boundary", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1204: return not ctx.at_boundary(_is_word)
							πF.SetLineno(1204)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_word); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_boundary, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_non_boundary.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1205: def at_loc_boundary(self, ctx):
					πF.SetLineno(1205)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("at_loc_boundary", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1206: return ctx.at_boundary(_is_loc_word)
							πF.SetLineno(1206)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_loc_word); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßat_boundary, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_loc_boundary.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1207: def at_loc_non_boundary(self, ctx):
					πF.SetLineno(1207)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("at_loc_non_boundary", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1208: return not ctx.at_boundary(_is_loc_word)
							πF.SetLineno(1208)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_loc_word); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_boundary, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_loc_non_boundary.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1209: def at_uni_boundary(self, ctx):
					πF.SetLineno(1209)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("at_uni_boundary", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1210: return ctx.at_boundary(_is_uni_word)
							πF.SetLineno(1210)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_uni_word); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßat_boundary, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_uni_boundary.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1211: def at_uni_non_boundary(self, ctx):
					πF.SetLineno(1211)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("at_uni_non_boundary", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1212: return not ctx.at_boundary(_is_uni_word)
							πF.SetLineno(1212)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_uni_word); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßat_boundary, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßat_uni_non_boundary.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1213: def unknown(self, ctx):
					πF.SetLineno(1213)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("unknown", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1214: return False
							πF.SetLineno(1214)
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
					if πE = πClass.SetItem(πF, ßunknown.ToObject(), πTemp013); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("_AtcodeDispatcher").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_AtcodeDispatcher.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1216: _AtcodeDispatcher.build_dispatch_table(ATCODES, "")
			πF.SetLineno(1216)
			πTemp002 = πF.MakeArgs(2)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßATCODES); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp002[1] = ß.ToObject()
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_AtcodeDispatcher); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßbuild_dispatch_table, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 1219: class _ChcodeDispatcher(_Dispatcher):
			πF.SetLineno(1219)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ß_Dispatcher); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp011 = πg.NewDict()
			if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp011.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_ChcodeDispatcher", "/usr/lib/python2.7/_sre.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp011
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1221: def category_digit(self, ctx):
					πF.SetLineno(1221)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("category_digit", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1222: return _is_digit(ctx.peek_char())
							πF.SetLineno(1222)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_digit); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_digit.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 1223: def category_not_digit(self, ctx):
					πF.SetLineno(1223)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("category_not_digit", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1224: return not _is_digit(ctx.peek_char())
							πF.SetLineno(1224)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_digit); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_not_digit.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1225: def category_space(self, ctx):
					πF.SetLineno(1225)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("category_space", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1226: return _is_space(ctx.peek_char())
							πF.SetLineno(1226)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_space); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_space.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1227: def category_not_space(self, ctx):
					πF.SetLineno(1227)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("category_not_space", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1228: return not _is_space(ctx.peek_char())
							πF.SetLineno(1228)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_space); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_not_space.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1229: def category_word(self, ctx):
					πF.SetLineno(1229)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("category_word", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1230: return _is_word(ctx.peek_char())
							πF.SetLineno(1230)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_word); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_word.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1231: def category_not_word(self, ctx):
					πF.SetLineno(1231)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("category_not_word", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1232: return not _is_word(ctx.peek_char())
							πF.SetLineno(1232)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_word); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_not_word.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1233: def category_linebreak(self, ctx):
					πF.SetLineno(1233)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("category_linebreak", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1234: return _is_linebreak(ctx.peek_char())
							πF.SetLineno(1234)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_linebreak); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_linebreak.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1235: def category_not_linebreak(self, ctx):
					πF.SetLineno(1235)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("category_not_linebreak", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1236: return not _is_linebreak(ctx.peek_char())
							πF.SetLineno(1236)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_linebreak); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_not_linebreak.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1237: def category_loc_word(self, ctx):
					πF.SetLineno(1237)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("category_loc_word", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1238: return _is_loc_word(ctx.peek_char())
							πF.SetLineno(1238)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_loc_word); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_loc_word.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1239: def category_loc_not_word(self, ctx):
					πF.SetLineno(1239)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("category_loc_not_word", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1240: return not _is_loc_word(ctx.peek_char())
							πF.SetLineno(1240)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_loc_word); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_loc_not_word.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1241: def category_uni_digit(self, ctx):
					πF.SetLineno(1241)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("category_uni_digit", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1242: return ctx.peek_char().isdigit()
							πF.SetLineno(1242)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßisdigit, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_uni_digit.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1243: def category_uni_not_digit(self, ctx):
					πF.SetLineno(1243)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("category_uni_not_digit", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1244: return not ctx.peek_char().isdigit()
							πF.SetLineno(1244)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßisdigit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ßcategory_uni_not_digit.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 1245: def category_uni_space(self, ctx):
					πF.SetLineno(1245)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("category_uni_space", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1246: return ctx.peek_char().isspace()
							πF.SetLineno(1246)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßisspace, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_uni_space.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 1247: def category_uni_not_space(self, ctx):
					πF.SetLineno(1247)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("category_uni_not_space", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1248: return not ctx.peek_char().isspace()
							πF.SetLineno(1248)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßisspace, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ßcategory_uni_not_space.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 1249: def category_uni_word(self, ctx):
					πF.SetLineno(1249)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("category_uni_word", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1250: return _is_uni_word(ctx.peek_char())
							πF.SetLineno(1250)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_uni_word); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_uni_word.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 1251: def category_uni_not_word(self, ctx):
					πF.SetLineno(1251)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("category_uni_not_word", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1252: return not _is_uni_word(ctx.peek_char())
							πF.SetLineno(1252)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_is_uni_word); πE != nil {
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
					if πE = πClass.SetItem(πF, ßcategory_uni_not_word.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 1253: def category_uni_linebreak(self, ctx):
					πF.SetLineno(1253)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("category_uni_linebreak", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1254: return ord(ctx.peek_char()) in _uni_linebreaks
							πF.SetLineno(1254)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_uni_linebreaks); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
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
					if πE = πClass.SetItem(πF, ßcategory_uni_linebreak.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 1255: def category_uni_not_linebreak(self, ctx):
					πF.SetLineno(1255)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("category_uni_not_linebreak", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
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
							// line 1256: return ord(ctx.peek_char()) not in _uni_linebreaks
							πF.SetLineno(1256)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µctx, "ctx"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µctx, ßpeek_char, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_uni_linebreaks); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp005).ToObject()
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
					if πE = πClass.SetItem(πF, ßcategory_uni_not_linebreak.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 1257: def unknown(self, ctx):
					πF.SetLineno(1257)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "ctx", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("unknown", "/usr/lib/python2.7/_sre.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µctx *πg.Object = πArgs[1]; _ = µctx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1258: return False
							πF.SetLineno(1258)
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
					if πE = πClass.SetItem(πF, ßunknown.ToObject(), πTemp020); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp011.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("_ChcodeDispatcher").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp011.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_ChcodeDispatcher.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1260: _ChcodeDispatcher.build_dispatch_table(CHCODES, "")
			πF.SetLineno(1260)
			πTemp002 = πF.MakeArgs(2)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßCHCODES); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp002[1] = ß.ToObject()
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_ChcodeDispatcher); πE != nil {
				continue
			}
			if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßbuild_dispatch_table, nil); πE != nil {
				continue
			}
			if πTemp007, πE = πTemp008.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 1263: _ascii_char_info = [ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 6, 2,
			πF.SetLineno(1263)
			πTemp002 = make([]*πg.Object, 128)
			πTemp002[0] = πg.NewInt(0).ToObject()
			πTemp002[1] = πg.NewInt(0).ToObject()
			πTemp002[2] = πg.NewInt(0).ToObject()
			πTemp002[3] = πg.NewInt(0).ToObject()
			πTemp002[4] = πg.NewInt(0).ToObject()
			πTemp002[5] = πg.NewInt(0).ToObject()
			πTemp002[6] = πg.NewInt(0).ToObject()
			πTemp002[7] = πg.NewInt(0).ToObject()
			πTemp002[8] = πg.NewInt(0).ToObject()
			πTemp002[9] = πg.NewInt(2).ToObject()
			πTemp002[10] = πg.NewInt(6).ToObject()
			πTemp002[11] = πg.NewInt(2).ToObject()
			πTemp002[12] = πg.NewInt(2).ToObject()
			πTemp002[13] = πg.NewInt(2).ToObject()
			πTemp002[14] = πg.NewInt(0).ToObject()
			πTemp002[15] = πg.NewInt(0).ToObject()
			πTemp002[16] = πg.NewInt(0).ToObject()
			πTemp002[17] = πg.NewInt(0).ToObject()
			πTemp002[18] = πg.NewInt(0).ToObject()
			πTemp002[19] = πg.NewInt(0).ToObject()
			πTemp002[20] = πg.NewInt(0).ToObject()
			πTemp002[21] = πg.NewInt(0).ToObject()
			πTemp002[22] = πg.NewInt(0).ToObject()
			πTemp002[23] = πg.NewInt(0).ToObject()
			πTemp002[24] = πg.NewInt(0).ToObject()
			πTemp002[25] = πg.NewInt(0).ToObject()
			πTemp002[26] = πg.NewInt(0).ToObject()
			πTemp002[27] = πg.NewInt(0).ToObject()
			πTemp002[28] = πg.NewInt(0).ToObject()
			πTemp002[29] = πg.NewInt(0).ToObject()
			πTemp002[30] = πg.NewInt(0).ToObject()
			πTemp002[31] = πg.NewInt(0).ToObject()
			πTemp002[32] = πg.NewInt(2).ToObject()
			πTemp002[33] = πg.NewInt(0).ToObject()
			πTemp002[34] = πg.NewInt(0).ToObject()
			πTemp002[35] = πg.NewInt(0).ToObject()
			πTemp002[36] = πg.NewInt(0).ToObject()
			πTemp002[37] = πg.NewInt(0).ToObject()
			πTemp002[38] = πg.NewInt(0).ToObject()
			πTemp002[39] = πg.NewInt(0).ToObject()
			πTemp002[40] = πg.NewInt(0).ToObject()
			πTemp002[41] = πg.NewInt(0).ToObject()
			πTemp002[42] = πg.NewInt(0).ToObject()
			πTemp002[43] = πg.NewInt(0).ToObject()
			πTemp002[44] = πg.NewInt(0).ToObject()
			πTemp002[45] = πg.NewInt(0).ToObject()
			πTemp002[46] = πg.NewInt(0).ToObject()
			πTemp002[47] = πg.NewInt(0).ToObject()
			πTemp002[48] = πg.NewInt(25).ToObject()
			πTemp002[49] = πg.NewInt(25).ToObject()
			πTemp002[50] = πg.NewInt(25).ToObject()
			πTemp002[51] = πg.NewInt(25).ToObject()
			πTemp002[52] = πg.NewInt(25).ToObject()
			πTemp002[53] = πg.NewInt(25).ToObject()
			πTemp002[54] = πg.NewInt(25).ToObject()
			πTemp002[55] = πg.NewInt(25).ToObject()
			πTemp002[56] = πg.NewInt(25).ToObject()
			πTemp002[57] = πg.NewInt(25).ToObject()
			πTemp002[58] = πg.NewInt(0).ToObject()
			πTemp002[59] = πg.NewInt(0).ToObject()
			πTemp002[60] = πg.NewInt(0).ToObject()
			πTemp002[61] = πg.NewInt(0).ToObject()
			πTemp002[62] = πg.NewInt(0).ToObject()
			πTemp002[63] = πg.NewInt(0).ToObject()
			πTemp002[64] = πg.NewInt(0).ToObject()
			πTemp002[65] = πg.NewInt(24).ToObject()
			πTemp002[66] = πg.NewInt(24).ToObject()
			πTemp002[67] = πg.NewInt(24).ToObject()
			πTemp002[68] = πg.NewInt(24).ToObject()
			πTemp002[69] = πg.NewInt(24).ToObject()
			πTemp002[70] = πg.NewInt(24).ToObject()
			πTemp002[71] = πg.NewInt(24).ToObject()
			πTemp002[72] = πg.NewInt(24).ToObject()
			πTemp002[73] = πg.NewInt(24).ToObject()
			πTemp002[74] = πg.NewInt(24).ToObject()
			πTemp002[75] = πg.NewInt(24).ToObject()
			πTemp002[76] = πg.NewInt(24).ToObject()
			πTemp002[77] = πg.NewInt(24).ToObject()
			πTemp002[78] = πg.NewInt(24).ToObject()
			πTemp002[79] = πg.NewInt(24).ToObject()
			πTemp002[80] = πg.NewInt(24).ToObject()
			πTemp002[81] = πg.NewInt(24).ToObject()
			πTemp002[82] = πg.NewInt(24).ToObject()
			πTemp002[83] = πg.NewInt(24).ToObject()
			πTemp002[84] = πg.NewInt(24).ToObject()
			πTemp002[85] = πg.NewInt(24).ToObject()
			πTemp002[86] = πg.NewInt(24).ToObject()
			πTemp002[87] = πg.NewInt(24).ToObject()
			πTemp002[88] = πg.NewInt(24).ToObject()
			πTemp002[89] = πg.NewInt(24).ToObject()
			πTemp002[90] = πg.NewInt(24).ToObject()
			πTemp002[91] = πg.NewInt(0).ToObject()
			πTemp002[92] = πg.NewInt(0).ToObject()
			πTemp002[93] = πg.NewInt(0).ToObject()
			πTemp002[94] = πg.NewInt(0).ToObject()
			πTemp002[95] = πg.NewInt(16).ToObject()
			πTemp002[96] = πg.NewInt(0).ToObject()
			πTemp002[97] = πg.NewInt(24).ToObject()
			πTemp002[98] = πg.NewInt(24).ToObject()
			πTemp002[99] = πg.NewInt(24).ToObject()
			πTemp002[100] = πg.NewInt(24).ToObject()
			πTemp002[101] = πg.NewInt(24).ToObject()
			πTemp002[102] = πg.NewInt(24).ToObject()
			πTemp002[103] = πg.NewInt(24).ToObject()
			πTemp002[104] = πg.NewInt(24).ToObject()
			πTemp002[105] = πg.NewInt(24).ToObject()
			πTemp002[106] = πg.NewInt(24).ToObject()
			πTemp002[107] = πg.NewInt(24).ToObject()
			πTemp002[108] = πg.NewInt(24).ToObject()
			πTemp002[109] = πg.NewInt(24).ToObject()
			πTemp002[110] = πg.NewInt(24).ToObject()
			πTemp002[111] = πg.NewInt(24).ToObject()
			πTemp002[112] = πg.NewInt(24).ToObject()
			πTemp002[113] = πg.NewInt(24).ToObject()
			πTemp002[114] = πg.NewInt(24).ToObject()
			πTemp002[115] = πg.NewInt(24).ToObject()
			πTemp002[116] = πg.NewInt(24).ToObject()
			πTemp002[117] = πg.NewInt(24).ToObject()
			πTemp002[118] = πg.NewInt(24).ToObject()
			πTemp002[119] = πg.NewInt(24).ToObject()
			πTemp002[120] = πg.NewInt(24).ToObject()
			πTemp002[121] = πg.NewInt(24).ToObject()
			πTemp002[122] = πg.NewInt(24).ToObject()
			πTemp002[123] = πg.NewInt(0).ToObject()
			πTemp002[124] = πg.NewInt(0).ToObject()
			πTemp002[125] = πg.NewInt(0).ToObject()
			πTemp002[126] = πg.NewInt(0).ToObject()
			πTemp002[127] = πg.NewInt(0).ToObject()
			πTemp007 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_ascii_char_info.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 1271: def _is_digit(char):
			πF.SetLineno(1271)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "char", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("_is_digit", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µchar *πg.Object = πArgs[0]; _ = µchar
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1272: code = ord(char)
					πF.SetLineno(1272)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp001[0] = µchar
					if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp003
					// line 1273: return code < 128 and _ascii_char_info[code] & 1
					πF.SetLineno(1273)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µcode, πg.NewInt(128).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp005 = µcode
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_ascii_char_info); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label1:
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
			if πE = πF.Globals().SetItem(πF, ß_is_digit.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 1275: def _is_space(char):
			πF.SetLineno(1275)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "char", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_is_space", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µchar *πg.Object = πArgs[0]; _ = µchar
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1276: code = ord(char)
					πF.SetLineno(1276)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp001[0] = µchar
					if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp003
					// line 1277: return code < 128 and _ascii_char_info[code] & 2
					πF.SetLineno(1277)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µcode, πg.NewInt(128).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp005 = µcode
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_ascii_char_info); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, πTemp006, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label1:
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
			if πE = πF.Globals().SetItem(πF, ß_is_space.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 1279: def _is_word(char):
			πF.SetLineno(1279)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "char", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("_is_word", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µchar *πg.Object = πArgs[0]; _ = µchar
				var µcode *πg.Object = πg.UnboundLocal; _ = µcode
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1281: code = ord(char)
					πF.SetLineno(1281)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp001[0] = µchar
					if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcode = πTemp003
					// line 1282: return code < 128 and _ascii_char_info[code] & 16
					πF.SetLineno(1282)
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.LT(πF, µcode, πg.NewInt(128).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
						continue
					}
					πTemp005 = µcode
					if πTemp007, πE = πg.ResolveGlobal(πF, ß_ascii_char_info); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp005); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, πTemp006, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label1:
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
			if πE = πF.Globals().SetItem(πF, ß_is_word.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 1284: def _is_loc_word(char):
			πF.SetLineno(1284)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "char", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("_is_loc_word", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µchar *πg.Object = πArgs[0]; _ = µchar
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 []*πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1285: return (not (ord(char) & ~255) and char.isalnum()) or char == '_'
					πF.SetLineno(1285)
					πTemp007 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp007[0] = µchar
					if πTemp008, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					if πTemp008, πE = πg.Invert(πF, πg.NewInt(255).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.And(πF, πTemp009, πTemp008); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp010).ToObject()
					πTemp003 = πTemp005
					if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µchar, ßisalnum, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp003 = πTemp006
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
			if πE = πF.Globals().SetItem(πF, ß_is_loc_word.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 1287: def _is_uni_word(char):
			πF.SetLineno(1287)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "char", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_is_uni_word", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µchar *πg.Object = πArgs[0]; _ = µchar
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
					// line 1288: return unichr(ord(char)).isalnum() or char == '_'
					πF.SetLineno(1288)
					πTemp003 = πF.MakeArgs(1)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp004[0] = µchar
					if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[0] = πTemp006
					if πTemp005, πE = πg.ResolveGlobal(πF, ßunichr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßisalnum, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp006
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Eq(πF, µchar, ß_.ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp005
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
			if πE = πF.Globals().SetItem(πF, ß_is_uni_word.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 1290: def _is_linebreak(char):
			πF.SetLineno(1290)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "char", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("_is_linebreak", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µchar *πg.Object = πArgs[0]; _ = µchar
				var πTemp001 *πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1291: return char == "\n"
					πF.SetLineno(1291)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µchar, πg.NewStr("\n").ToObject()); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_is_linebreak.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 1294: _uni_linebreaks = [10, 13, 28, 29, 30, 133, 8232, 8233]
			πF.SetLineno(1294)
			πTemp002 = make([]*πg.Object, 8)
			πTemp002[0] = πg.NewInt(10).ToObject()
			πTemp002[1] = πg.NewInt(13).ToObject()
			πTemp002[2] = πg.NewInt(28).ToObject()
			πTemp002[3] = πg.NewInt(29).ToObject()
			πTemp002[4] = πg.NewInt(30).ToObject()
			πTemp002[5] = πg.NewInt(133).ToObject()
			πTemp002[6] = πg.NewInt(8232).ToObject()
			πTemp002[7] = πg.NewInt(8233).ToObject()
			πTemp015 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_uni_linebreaks.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 1296: def _log(message):
			πF.SetLineno(1296)
			πTemp010 = make([]πg.Param, 1)
			πTemp010[0] = πg.Param{Name: "message", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("_log", "/usr/lib/python2.7/_sre.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmessage *πg.Object = πArgs[0]; _ = µmessage
				var πTemp001 bool
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πTemp001, πE = πg.IsTrue(πF, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp001 {
						goto Label1
					}
					goto Label2
					// line 1297: if 0:
					πF.SetLineno(1297)
				Label1:
					// line 1298: print message
					πF.SetLineno(1298)
					πTemp002 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πTemp002[0] = µmessage
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_log.ToObject(), πTemp015); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_sre", Code)
}

