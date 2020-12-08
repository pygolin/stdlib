package optparse
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/__builtin__"
	// _ "github.com/pygolin/stdlib/os"
	// _ "github.com/pygolin/stdlib/textwrap"
	// _ "github.com/pygolin/stdlib/types"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß0b := πg.InternStr("0b")
		ß0x := πg.InternStr("0x")
		ßACTIONS := πg.InternStr("ACTIONS")
		ßALWAYS_TYPED_ACTIONS := πg.InternStr("ALWAYS_TYPED_ACTIONS")
		ßATTRS := πg.InternStr("ATTRS")
		ßAmbiguousOptionError := πg.InternStr("AmbiguousOptionError")
		ßBadOptionError := πg.InternStr("BadOptionError")
		ßCHECK_METHODS := πg.InternStr("CHECK_METHODS")
		ßCOLUMNS := πg.InternStr("COLUMNS")
		ßCONST_ACTIONS := πg.InternStr("CONST_ACTIONS")
		ßDEFAULT := πg.InternStr("DEFAULT")
		ßDictType := πg.InternStr("DictType")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßHELP := πg.InternStr("HELP")
		ßHelpFormatter := πg.InternStr("HelpFormatter")
		ßINDENT_CHAR := πg.InternStr("INDENT_CHAR")
		ßIndentedHelpFormatter := πg.InternStr("IndentedHelpFormatter")
		ßKeyError := πg.InternStr("KeyError")
		ßListType := πg.InternStr("ListType")
		ßNO := πg.InternStr("NO")
		ßNO_DEFAULT := πg.InternStr("NO_DEFAULT")
		ßNO_DEFAULT_VALUE := πg.InternStr("NO_DEFAULT_VALUE")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßNotImplementedError := πg.InternStr("NotImplementedError")
		ßOptParseError := πg.InternStr("OptParseError")
		ßOption := πg.InternStr("Option")
		ßOptionConflictError := πg.InternStr("OptionConflictError")
		ßOptionContainer := πg.InternStr("OptionContainer")
		ßOptionError := πg.InternStr("OptionError")
		ßOptionGroup := πg.InternStr("OptionGroup")
		ßOptionParser := πg.InternStr("OptionParser")
		ßOptionValueError := πg.InternStr("OptionValueError")
		ßOptions := πg.InternStr("Options")
		ßSTORE_ACTIONS := πg.InternStr("STORE_ACTIONS")
		ßSUPPRESS := πg.InternStr("SUPPRESS")
		ßSUPPRESS_HELP := πg.InternStr("SUPPRESS_HELP")
		ßSUPPRESS_USAGE := πg.InternStr("SUPPRESS_USAGE")
		ßStringType := πg.InternStr("StringType")
		ßStringTypes := πg.InternStr("StringTypes")
		ßTYPED_ACTIONS := πg.InternStr("TYPED_ACTIONS")
		ßTYPES := πg.InternStr("TYPES")
		ßTYPE_CHECKER := πg.InternStr("TYPE_CHECKER")
		ßTitledHelpFormatter := πg.InternStr("TitledHelpFormatter")
		ßTrue := πg.InternStr("True")
		ßTupleType := πg.InternStr("TupleType")
		ßTypeError := πg.InternStr("TypeError")
		ßTypeType := πg.InternStr("TypeType")
		ßUSAGE := πg.InternStr("USAGE")
		ßUnicodeType := πg.InternStr("UnicodeType")
		ßUsage := πg.InternStr("Usage")
		ßValueError := πg.InternStr("ValueError")
		ßValues := πg.InternStr("Values")
		ß_ := πg.InternStr("_")
		ß__all__ := πg.InternStr("__all__")
		ß__call__ := πg.InternStr("__call__")
		ß__class__ := πg.InternStr("__class__")
		ß__cmp__ := πg.InternStr("__cmp__")
		ß__copyright__ := πg.InternStr("__copyright__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__import__ := πg.InternStr("__import__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__str__ := πg.InternStr("__str__")
		ß__version__ := πg.InternStr("__version__")
		ß_add_help_option := πg.InternStr("_add_help_option")
		ß_add_version_option := πg.InternStr("_add_version_option")
		ß_builtin_cvt := πg.InternStr("_builtin_cvt")
		ß_check_action := πg.InternStr("_check_action")
		ß_check_callback := πg.InternStr("_check_callback")
		ß_check_choice := πg.InternStr("_check_choice")
		ß_check_conflict := πg.InternStr("_check_conflict")
		ß_check_const := πg.InternStr("_check_const")
		ß_check_dest := πg.InternStr("_check_dest")
		ß_check_nargs := πg.InternStr("_check_nargs")
		ß_check_opt_strings := πg.InternStr("_check_opt_strings")
		ß_check_type := πg.InternStr("_check_type")
		ß_create_option_list := πg.InternStr("_create_option_list")
		ß_create_option_mappings := πg.InternStr("_create_option_mappings")
		ß_format_text := πg.InternStr("_format_text")
		ß_get_all_options := πg.InternStr("_get_all_options")
		ß_get_args := πg.InternStr("_get_args")
		ß_get_encoding := πg.InternStr("_get_encoding")
		ß_init_parsing_state := πg.InternStr("_init_parsing_state")
		ß_long_opt := πg.InternStr("_long_opt")
		ß_long_opt_fmt := πg.InternStr("_long_opt_fmt")
		ß_long_opts := πg.InternStr("_long_opts")
		ß_match_abbrev := πg.InternStr("_match_abbrev")
		ß_match_long_opt := πg.InternStr("_match_long_opt")
		ß_parse_int := πg.InternStr("_parse_int")
		ß_parse_long := πg.InternStr("_parse_long")
		ß_parse_num := πg.InternStr("_parse_num")
		ß_populate_option_list := πg.InternStr("_populate_option_list")
		ß_process_args := πg.InternStr("_process_args")
		ß_process_long_opt := πg.InternStr("_process_long_opt")
		ß_process_short_opts := πg.InternStr("_process_short_opts")
		ß_repr := πg.InternStr("_repr")
		ß_set_attrs := πg.InternStr("_set_attrs")
		ß_set_opt_strings := πg.InternStr("_set_opt_strings")
		ß_share_option_mappings := πg.InternStr("_share_option_mappings")
		ß_short_opt := πg.InternStr("_short_opt")
		ß_short_opt_fmt := πg.InternStr("_short_opt_fmt")
		ß_short_opts := πg.InternStr("_short_opts")
		ß_update := πg.InternStr("_update")
		ß_update_careful := πg.InternStr("_update_careful")
		ß_update_loose := πg.InternStr("_update_loose")
		ßaction := πg.InternStr("action")
		ßadd_option := πg.InternStr("add_option")
		ßadd_option_group := πg.InternStr("add_option_group")
		ßadd_options := πg.InternStr("add_options")
		ßallow_interspersed_args := πg.InternStr("allow_interspersed_args")
		ßappend := πg.InternStr("append")
		ßappend_const := πg.InternStr("append_const")
		ßargv := πg.InternStr("argv")
		ßbasestring := πg.InternStr("basestring")
		ßcallback := πg.InternStr("callback")
		ßcallback_args := πg.InternStr("callback_args")
		ßcallback_kwargs := πg.InternStr("callback_kwargs")
		ßcareful := πg.InternStr("careful")
		ßcheck_builtin := πg.InternStr("check_builtin")
		ßcheck_choice := πg.InternStr("check_choice")
		ßcheck_value := πg.InternStr("check_value")
		ßcheck_values := πg.InternStr("check_values")
		ßchoice := πg.InternStr("choice")
		ßchoices := πg.InternStr("choices")
		ßcmp := πg.InternStr("cmp")
		ßconflict_handler := πg.InternStr("conflict_handler")
		ßconst := πg.InternStr("const")
		ßcontainer := πg.InternStr("container")
		ßconvert_value := πg.InternStr("convert_value")
		ßcount := πg.InternStr("count")
		ßcurrent_indent := πg.InternStr("current_indent")
		ßdedent := πg.InternStr("dedent")
		ßdefault := πg.InternStr("default")
		ßdefault_tag := πg.InternStr("default_tag")
		ßdefaults := πg.InternStr("defaults")
		ßdescription := πg.InternStr("description")
		ßdest := πg.InternStr("dest")
		ßdestroy := πg.InternStr("destroy")
		ßdict := πg.InternStr("dict")
		ßdir := πg.InternStr("dir")
		ßdisable_interspersed_args := πg.InternStr("disable_interspersed_args")
		ßenable_interspersed_args := πg.InternStr("enable_interspersed_args")
		ßensure_value := πg.InternStr("ensure_value")
		ßenviron := πg.InternStr("environ")
		ßepilog := πg.InternStr("epilog")
		ßerror := πg.InternStr("error")
		ßexecfile := πg.InternStr("execfile")
		ßexit := πg.InternStr("exit")
		ßexpand_default := πg.InternStr("expand_default")
		ßexpand_prog_name := πg.InternStr("expand_prog_name")
		ßfill := πg.InternStr("fill")
		ßfloat := πg.InternStr("float")
		ßformat_description := πg.InternStr("format_description")
		ßformat_epilog := πg.InternStr("format_epilog")
		ßformat_heading := πg.InternStr("format_heading")
		ßformat_help := πg.InternStr("format_help")
		ßformat_option := πg.InternStr("format_option")
		ßformat_option_help := πg.InternStr("format_option_help")
		ßformat_option_strings := πg.InternStr("format_option_strings")
		ßformat_usage := πg.InternStr("format_usage")
		ßformatter := πg.InternStr("formatter")
		ßget := πg.InternStr("get")
		ßget_default_values := πg.InternStr("get_default_values")
		ßget_description := πg.InternStr("get_description")
		ßget_opt_string := πg.InternStr("get_opt_string")
		ßget_option := πg.InternStr("get_option")
		ßget_option_group := πg.InternStr("get_option_group")
		ßget_prog_name := πg.InternStr("get_prog_name")
		ßget_usage := πg.InternStr("get_usage")
		ßget_version := πg.InternStr("get_version")
		ßgetattr := πg.InternStr("getattr")
		ßgettext := πg.InternStr("gettext")
		ßhas_option := πg.InternStr("has_option")
		ßhasattr := πg.InternStr("hasattr")
		ßhelp := πg.InternStr("help")
		ßhelp_position := πg.InternStr("help_position")
		ßhelp_width := πg.InternStr("help_width")
		ßid := πg.InternStr("id")
		ßindent := πg.InternStr("indent")
		ßindent_increment := πg.InternStr("indent_increment")
		ßinsert := πg.InternStr("insert")
		ßint := πg.InternStr("int")
		ßinteger := πg.InternStr("integer")
		ßisbasestring := πg.InternStr("isbasestring")
		ßisinstance := πg.InternStr("isinstance")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlargs := πg.InternStr("largs")
		ßlen := πg.InternStr("len")
		ßlevel := πg.InternStr("level")
		ßlong := πg.InternStr("long")
		ßloose := πg.InternStr("loose")
		ßlower := πg.InternStr("lower")
		ßmake_option := πg.InternStr("make_option")
		ßmap := πg.InternStr("map")
		ßmax := πg.InternStr("max")
		ßmax_help_position := πg.InternStr("max_help_position")
		ßmetavar := πg.InternStr("metavar")
		ßmin := πg.InternStr("min")
		ßmodules := πg.InternStr("modules")
		ßmsg := πg.InternStr("msg")
		ßnargs := πg.InternStr("nargs")
		ßnone := πg.InternStr("none")
		ßobject := πg.InternStr("object")
		ßopt_str := πg.InternStr("opt_str")
		ßoption_class := πg.InternStr("option_class")
		ßoption_groups := πg.InternStr("option_groups")
		ßoption_id := πg.InternStr("option_id")
		ßoption_list := πg.InternStr("option_list")
		ßoption_strings := πg.InternStr("option_strings")
		ßos := πg.InternStr("os")
		ßparse_args := πg.InternStr("parse_args")
		ßparser := πg.InternStr("parser")
		ßpossibilities := πg.InternStr("possibilities")
		ßprint_help := πg.InternStr("print_help")
		ßprint_usage := πg.InternStr("print_usage")
		ßprint_version := πg.InternStr("print_version")
		ßprocess := πg.InternStr("process")
		ßprocess_default_values := πg.InternStr("process_default_values")
		ßprog := πg.InternStr("prog")
		ßrargs := πg.InternStr("rargs")
		ßread_file := πg.InternStr("read_file")
		ßread_module := πg.InternStr("read_module")
		ßremove := πg.InternStr("remove")
		ßremove_option := πg.InternStr("remove_option")
		ßrepr := πg.InternStr("repr")
		ßresolve := πg.InternStr("resolve")
		ßset_conflict_handler := πg.InternStr("set_conflict_handler")
		ßset_default := πg.InternStr("set_default")
		ßset_defaults := πg.InternStr("set_defaults")
		ßset_description := πg.InternStr("set_description")
		ßset_long_opt_delimiter := πg.InternStr("set_long_opt_delimiter")
		ßset_parser := πg.InternStr("set_parser")
		ßset_process_default_values := πg.InternStr("set_process_default_values")
		ßset_short_opt_delimiter := πg.InternStr("set_short_opt_delimiter")
		ßset_title := πg.InternStr("set_title")
		ßset_usage := πg.InternStr("set_usage")
		ßsetattr := πg.InternStr("setattr")
		ßshort_first := πg.InternStr("short_first")
		ßsort := πg.InternStr("sort")
		ßsplit := πg.InternStr("split")
		ßstandard_option_list := πg.InternStr("standard_option_list")
		ßstartswith := πg.InternStr("startswith")
		ßstderr := πg.InternStr("stderr")
		ßstdout := πg.InternStr("stdout")
		ßstore := πg.InternStr("store")
		ßstore_const := πg.InternStr("store_const")
		ßstore_false := πg.InternStr("store_false")
		ßstore_option_strings := πg.InternStr("store_option_strings")
		ßstore_true := πg.InternStr("store_true")
		ßstr := πg.InternStr("str")
		ßstring := πg.InternStr("string")
		ßsys := πg.InternStr("sys")
		ßtake_action := πg.InternStr("take_action")
		ßtakes_value := πg.InternStr("takes_value")
		ßtextwrap := πg.InternStr("textwrap")
		ßtitle := πg.InternStr("title")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		ßupdate := πg.InternStr("update")
		ßupper := πg.InternStr("upper")
		ßusage := πg.InternStr("usage")
		ßvalues := πg.InternStr("values")
		ßvars := πg.InternStr("vars")
		ßversion := πg.InternStr("version")
		ßwidth := πg.InternStr("width")
		ßwrap := πg.InternStr("wrap")
		ßwrite := πg.InternStr("write")
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
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Dict
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
		var πTemp015 *πg.BaseException
		_ = πTemp015
		var πTemp016 *πg.Traceback
		_ = πTemp016
		var πTemp017 bool
		_ = πTemp017
		var πTemp018 *πg.Object
		_ = πTemp018
		var πTemp019 *πg.Object
		_ = πTemp019
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """A powerful, extensible, and easy-to-use option parser.
			πF.SetLineno(1)
			// line 1: """A powerful, extensible, and easy-to-use option parser.
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A powerful, extensible, and easy-to-use option parser.\n\nBy Greg Ward <gward@python.net>\n\nOriginally distributed as Optik.\n\nFor support, use the optik-users@lists.sourceforge.net mailing list\n(http://lists.sourceforge.net/lists/listinfo/optik-users).\n\nSimple usage example:\n\n   from optparse import OptionParser\n\n   parser = OptionParser()\n   parser.add_option(\"-f\", \"--file\", dest=\"filename\",\n                     help=\"write report to FILE\", metavar=\"FILE\")\n   parser.add_option(\"-q\", \"--quiet\",\n                     action=\"store_false\", dest=\"verbose\", default=True,\n                     help=\"don't print status messages to stdout\")\n\n   (options, args) = parser.parse_args()\n").ToObject()); πE != nil {
				continue
			}
			// line 24: __version__ = "1.5.3"
			πF.SetLineno(24)
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πg.NewStr("1.5.3").ToObject()); πE != nil {
				continue
			}
			// line 26: __all__ = ['Option',
			πF.SetLineno(26)
			πTemp001 = make([]*πg.Object, 16)
			πTemp001[0] = ßOption.ToObject()
			πTemp001[1] = ßmake_option.ToObject()
			πTemp001[2] = ßSUPPRESS_HELP.ToObject()
			πTemp001[3] = ßSUPPRESS_USAGE.ToObject()
			πTemp001[4] = ßValues.ToObject()
			πTemp001[5] = ßOptionContainer.ToObject()
			πTemp001[6] = ßOptionGroup.ToObject()
			πTemp001[7] = ßOptionParser.ToObject()
			πTemp001[8] = ßHelpFormatter.ToObject()
			πTemp001[9] = ßIndentedHelpFormatter.ToObject()
			πTemp001[10] = ßTitledHelpFormatter.ToObject()
			πTemp001[11] = ßOptParseError.ToObject()
			πTemp001[12] = ßOptionError.ToObject()
			πTemp001[13] = ßOptionConflictError.ToObject()
			πTemp001[14] = ßOptionValueError.ToObject()
			πTemp001[15] = ßBadOptionError.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 43: __copyright__ = """
			πF.SetLineno(43)
			if πE = πF.Globals().SetItem(πF, ß__copyright__.ToObject(), πg.NewStr("\nCopyright (c) 2001-2006 Gregory P. Ward.  All rights reserved.\nCopyright (c) 2002-2006 Python Software Foundation.  All rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are\nmet:\n\n  * Redistributions of source code must retain the above copyright\n    notice, this list of conditions and the following disclaimer.\n\n  * Redistributions in binary form must reproduce the above copyright\n    notice, this list of conditions and the following disclaimer in the\n    documentation and/or other materials provided with the distribution.\n\n  * Neither the name of the author nor the names of its\n    contributors may be used to endorse or promote products derived from\n    this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS\nIS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED\nTO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A\nPARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE AUTHOR OR\nCONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,\nEXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,\nPROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR\nPROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF\nLIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING\nNEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n").ToObject()); πE != nil {
				continue
			}
			// line 75: import sys, os
			πF.SetLineno(75)
			if πTemp001, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ImportModule(πF, "os"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 76: import types
			πF.SetLineno(76)
			if πTemp001, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßtypes.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 77: import textwrap
			πF.SetLineno(77)
			if πTemp001, πE = πg.ImportModule(πF, "textwrap"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßtextwrap.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 79: def _repr(self):
			πF.SetLineno(79)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "self", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("_repr", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]; _ = µself
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 80: return "<%s at 0x%x: %s>" % (self.__class__.__name__, id(self), self)
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__name__, nil); πE != nil {
						continue
					}
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp005[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(πTemp004, πTemp006, µself).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("<%s at 0x%x: %s>").ToObject(), πTemp002); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_repr.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 82: def setattr(self, attr, value):
			πF.SetLineno(82)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "self", Def: nil}
			πTemp003[1] = πg.Param{Name: "attr", Def: nil}
			πTemp003[2] = πg.Param{Name: "value", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("setattr", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]; _ = µself
				var µattr *πg.Object = πArgs[1]; _ = µattr
				var µvalue *πg.Object = πArgs[2]; _ = µvalue
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
					// line 83: self.__dict__[attr] = value
					πF.SetLineno(83)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
						continue
					}
					πTemp003 = µattr
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
			if πE = πF.Globals().SetItem(πF, ßsetattr.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 85: INDENT_CHAR = " "
			πF.SetLineno(85)
			if πE = πF.Globals().SetItem(πF, ßINDENT_CHAR.ToObject(), πg.NewStr(" ").ToObject()); πE != nil {
				continue
			}
			// line 98: def gettext(message):
			πF.SetLineno(98)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "message", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("gettext", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmessage *πg.Object = πArgs[0]; _ = µmessage
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 99: return message
					πF.SetLineno(99)
					if πE = πg.CheckLocal(πF, µmessage, "message"); πE != nil {
						continue
					}
					πR = µmessage
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßgettext.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 100: _ = gettext
			πF.SetLineno(100)
			if πTemp006, πE = πg.ResolveGlobal(πF, ßgettext); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 103: class OptParseError (Exception):
			πF.SetLineno(103)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp001[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OptParseError", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
					// line 104: def __init__(self, msg):
					πF.SetLineno(104)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "msg", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πArgs[1]; _ = µmsg
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 105: self.msg = msg
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmsg); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmsg, πTemp001); πE != nil {
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
					// line 107: def __str__(self):
					πF.SetLineno(107)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 108: return self.msg
							πF.SetLineno(108)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmsg, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("OptParseError").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOptParseError.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 111: class OptionError (OptParseError):
			πF.SetLineno(111)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßOptParseError); πE != nil {
				continue
			}
			πTemp001[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OptionError", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
					// line 112: """
					πF.SetLineno(112)
					// line 112: """
					πF.SetLineno(112)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raised if an Option instance is created with invalid or\n    inconsistent arguments.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 117: def __init__(self, msg, option):
					πF.SetLineno(117)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "msg", Def: nil}
					πTemp002[2] = πg.Param{Name: "option", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πArgs[1]; _ = µmsg
						var µoption *πg.Object = πArgs[2]; _ = µoption
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
							// line 118: self.msg = msg
							πF.SetLineno(118)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µmsg); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmsg, πTemp001); πE != nil {
								continue
							}
							// line 119: self.option_id = str(option)
							πF.SetLineno(119)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp002[0] = µoption
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßoption_id, πTemp001); πE != nil {
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
					// line 121: def __str__(self):
					πF.SetLineno(121)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoption_id, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 122: if self.option_id:
							πF.SetLineno(122)
						Label1:
							// line 123: return "option %s: %s" % (self.option_id, self.msg)
							πF.SetLineno(123)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoption_id, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmsg, nil); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("option %s: %s").ToObject(), πTemp003); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label2:
							// line 125: return self.msg
							πF.SetLineno(125)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmsg, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("OptionError").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOptionError.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 127: class OptionConflictError (OptionError):
			πF.SetLineno(127)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
				continue
			}
			πTemp001[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OptionConflictError", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 128: """
					πF.SetLineno(128)
					// line 128: """
					πF.SetLineno(128)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raised if conflicting options are added to an OptionParser.\n    ").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("OptionConflictError").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOptionConflictError.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 132: class OptionValueError (OptParseError):
			πF.SetLineno(132)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßOptParseError); πE != nil {
				continue
			}
			πTemp001[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OptionValueError", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 133: """
					πF.SetLineno(133)
					// line 133: """
					πF.SetLineno(133)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raised if an invalid option value is encountered on the command\n    line.\n    ").ToObject()); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("OptionValueError").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOptionValueError.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 138: class BadOptionError (OptParseError):
			πF.SetLineno(138)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßOptParseError); πE != nil {
				continue
			}
			πTemp001[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("BadOptionError", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
					// line 139: """
					πF.SetLineno(139)
					// line 139: """
					πF.SetLineno(139)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raised if an invalid option is seen on the command line.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 142: def __init__(self, opt_str):
					πF.SetLineno(142)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "opt_str", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt_str *πg.Object = πArgs[1]; _ = µopt_str
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 143: self.opt_str = opt_str
							πF.SetLineno(143)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µopt_str); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßopt_str, πTemp001); πE != nil {
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
					// line 145: def __str__(self):
					πF.SetLineno(145)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 146: return _("no such option: %s") % self.opt_str
							πF.SetLineno(146)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("no such option: %s").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßopt_str, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp004, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("BadOptionError").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßBadOptionError.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 148: class AmbiguousOptionError (BadOptionError):
			πF.SetLineno(148)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBadOptionError); πE != nil {
				continue
			}
			πTemp001[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("AmbiguousOptionError", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
					// line 149: """
					πF.SetLineno(149)
					// line 149: """
					πF.SetLineno(149)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Raised if an ambiguous option is seen on the command line.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 152: def __init__(self, opt_str, possibilities):
					πF.SetLineno(152)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "opt_str", Def: nil}
					πTemp002[2] = πg.Param{Name: "possibilities", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt_str *πg.Object = πArgs[1]; _ = µopt_str
						var µpossibilities *πg.Object = πArgs[2]; _ = µpossibilities
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
							// line 153: BadOptionError.__init__(self, opt_str)
							πF.SetLineno(153)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							πTemp001[1] = µopt_str
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBadOptionError); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 154: self.possibilities = possibilities
							πF.SetLineno(154)
							if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µpossibilities); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpossibilities, πTemp002); πE != nil {
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
					// line 156: def __str__(self):
					πF.SetLineno(156)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 157: return (_("ambiguous option: %s (%s?)")
							πF.SetLineno(157)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("ambiguous option: %s (%s?)").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßopt_str, nil); πE != nil {
								continue
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpossibilities, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp006
							if πTemp006, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πg.NewTuple2(πTemp005, πTemp007).ToObject()
							if πTemp001, πE = πg.Mod(πF, πTemp004, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("AmbiguousOptionError").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßAmbiguousOptionError.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 161: class HelpFormatter(object):
			πF.SetLineno(161)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("HelpFormatter", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
					// line 163: """
					πF.SetLineno(163)
					// line 163: """
					πF.SetLineno(163)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Abstract base class for formatting option help.  OptionParser\n    instances should use one of the HelpFormatter subclasses for\n    formatting help; by default IndentedHelpFormatter is used.\n\n    Instance attributes:\n      parser : OptionParser\n        the controlling OptionParser instance\n      indent_increment : int\n        the number of columns to indent per nesting level\n      max_help_position : int\n        the maximum starting column for option help text\n      help_position : int\n        the calculated starting column for option help text;\n        initially the same as the maximum\n      width : int\n        total number of columns for output (pass None to constructor for\n        this value to be taken from the $COLUMNS environment variable)\n      level : int\n        current indentation level\n      current_indent : int\n        current indentation level (in columns)\n      help_width : int\n        number of columns available for option help text (calculated)\n      default_tag : str\n        text to replace with each option's default value, \"%default\"\n        by default.  Set to false value to disable default value expansion.\n      option_strings : { Option : str }\n        maps Option instances to the snippet of help text explaining\n        the syntax of that option, e.g. \"-h, --help\" or\n        \"-fFILE, --file=FILE\"\n      _short_opt_fmt : str\n        format string controlling how short options with values are\n        printed in help text.  Must be either \"%s%s\" (\"-fFILE\") or\n        \"%s %s\" (\"-f FILE\"), because those are the two syntaxes that\n        Optik supports.\n      _long_opt_fmt : str\n        similar but for long options; must be either \"%s %s\" (\"--file FILE\")\n        or \"%s=%s\" (\"--file=FILE\").\n    ").ToObject()); πE != nil {
						continue
					}
					// line 204: NO_DEFAULT_VALUE = "none"
					πF.SetLineno(204)
					if πE = πClass.SetItem(πF, ßNO_DEFAULT_VALUE.ToObject(), ßnone.ToObject()); πE != nil {
						continue
					}
					// line 206: def __init__(self,
					πF.SetLineno(206)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "indent_increment", Def: nil}
					πTemp002[2] = πg.Param{Name: "max_help_position", Def: nil}
					πTemp002[3] = πg.Param{Name: "width", Def: nil}
					πTemp002[4] = πg.Param{Name: "short_first", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindent_increment *πg.Object = πArgs[1]; _ = µindent_increment
						var µmax_help_position *πg.Object = πArgs[2]; _ = µmax_help_position
						var µwidth *πg.Object = πArgs[3]; _ = µwidth
						var µshort_first *πg.Object = πArgs[4]; _ = µshort_first
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
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 *πg.Dict
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 211: self.parser = None
							πF.SetLineno(211)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparser, πTemp002); πE != nil {
								continue
							}
							// line 212: self.indent_increment = indent_increment
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µindent_increment, "indent_increment"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µindent_increment); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßindent_increment, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µwidth == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 213: if width is None:
							πF.SetLineno(213)
						Label1:
							// line 214: try:
							πF.SetLineno(214)
							πF.PushCheckpoint(4)
							// line 215: width = int(os.environ['COLUMNS'])
							πF.SetLineno(215)
							πTemp004 = πF.MakeArgs(1)
							πTemp001 = ßCOLUMNS.ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßenviron, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µwidth = πTemp002
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, πTemp005).ToObject()
							if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 216: except (KeyError, ValueError):
							πF.SetLineno(216)
						Label5:
							// line 217: width = 80
							πF.SetLineno(217)
							µwidth = πg.NewInt(80).ToObject()
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 218: width -= 2
							πF.SetLineno(218)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ISub(πF, µwidth, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							µwidth = πTemp001
							goto Label2
						Label2:
							// line 219: self.width = width
							πF.SetLineno(219)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µwidth); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwidth, πTemp001); πE != nil {
								continue
							}
							// line 220: self.help_position = self.max_help_position = \
							πF.SetLineno(220)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmax_help_position, "max_help_position"); πE != nil {
								continue
							}
							πTemp004[0] = µmax_help_position
							πTemp009 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µwidth, πg.NewInt(20).ToObject()); πE != nil {
								continue
							}
							πTemp009[0] = πTemp001
							if πE = πg.CheckLocal(πF, µindent_increment, "indent_increment"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µindent_increment, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp009[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp004[1] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßhelp_position, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmax_help_position, πTemp001); πE != nil {
								continue
							}
							// line 222: self.current_indent = 0
							πF.SetLineno(222)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_indent, πTemp001); πE != nil {
								continue
							}
							// line 223: self.level = 0
							πF.SetLineno(223)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlevel, πTemp001); πE != nil {
								continue
							}
							// line 224: self.help_width = None          # computed later
							πF.SetLineno(224)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhelp_width, πTemp002); πE != nil {
								continue
							}
							// line 225: self.short_first = short_first
							πF.SetLineno(225)
							if πE = πg.CheckLocal(πF, µshort_first, "short_first"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µshort_first); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßshort_first, πTemp001); πE != nil {
								continue
							}
							// line 226: self.default_tag = "%default"
							πF.SetLineno(226)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr("%default").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdefault_tag, πTemp001); πE != nil {
								continue
							}
							// line 227: self.option_strings = {}
							πF.SetLineno(227)
							πTemp010 = πg.NewDict()
							πTemp001 = πTemp010.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoption_strings, πTemp002); πE != nil {
								continue
							}
							// line 228: self._short_opt_fmt = "%s %s"
							πF.SetLineno(228)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr("%s %s").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_short_opt_fmt, πTemp001); πE != nil {
								continue
							}
							// line 229: self._long_opt_fmt = "%s=%s"
							πF.SetLineno(229)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewStr("%s=%s").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_long_opt_fmt, πTemp001); πE != nil {
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
					// line 231: def set_parser(self, parser):
					πF.SetLineno(231)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "parser", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("set_parser", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µparser *πg.Object = πArgs[1]; _ = µparser
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 232: self.parser = parser
							πF.SetLineno(232)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µparser); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparser, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_parser.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 234: def set_short_opt_delimiter(self, delim):
					πF.SetLineno(234)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "delim", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("set_short_opt_delimiter", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdelim *πg.Object = πArgs[1]; _ = µdelim
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
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ß.ToObject(), πg.NewStr(" ").ToObject()).ToObject()
							if πTemp003, πE = πg.Contains(πF, πTemp002, µdelim); πE != nil {
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
							// line 235: if delim not in ("", " "):
							πF.SetLineno(235)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid metavar delimiter for short options: %r").ToObject(), µdelim); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 236: raise ValueError(
							πF.SetLineno(236)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 238: self._short_opt_fmt = "%s" + delim + "%s"
							πF.SetLineno(238)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("%s").ToObject(), µdelim); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr("%s").ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_short_opt_fmt, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_short_opt_delimiter.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 240: def set_long_opt_delimiter(self, delim):
					πF.SetLineno(240)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "delim", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("set_long_opt_delimiter", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdelim *πg.Object = πArgs[1]; _ = µdelim
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
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πg.NewStr("=").ToObject(), πg.NewStr(" ").ToObject()).ToObject()
							if πTemp003, πE = πg.Contains(πF, πTemp002, µdelim); πE != nil {
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
							// line 241: if delim not in ("=", " "):
							πF.SetLineno(241)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid metavar delimiter for long options: %r").ToObject(), µdelim); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 242: raise ValueError(
							πF.SetLineno(242)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 244: self._long_opt_fmt = "%s" + delim + "%s"
							πF.SetLineno(244)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("%s").ToObject(), µdelim); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr("%s").ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_long_opt_fmt, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_long_opt_delimiter.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 246: def indent(self):
					πF.SetLineno(246)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("indent", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 247: self.current_indent += self.indent_increment
							πF.SetLineno(247)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent_increment, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_indent, πTemp003); πE != nil {
								continue
							}
							// line 248: self.level += 1
							πF.SetLineno(248)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlevel, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlevel, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßindent.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 250: def dedent(self):
					πF.SetLineno(250)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("dedent", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 251: self.current_indent -= self.indent_increment
							πF.SetLineno(251)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßindent_increment, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ISub(πF, πTemp001, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßcurrent_indent, πTemp003); πE != nil {
								continue
							}
							// line 252: assert self.current_indent >= 0, "Indent decreased below 0."
							πF.SetLineno(252)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, πg.NewStr("Indent decreased below 0.").ToObject()); πE != nil {
								continue
							}
							// line 253: self.level -= 1
							πF.SetLineno(253)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßlevel, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlevel, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdedent.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 255: def format_usage(self, usage):
					πF.SetLineno(255)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "usage", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("format_usage", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µusage *πg.Object = πArgs[1]; _ = µusage
						var πTemp001 *πg.Object
						_ = πTemp001
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
							// line 256: raise NotImplementedError, "subclasses must implement"
							πF.SetLineno(256)
							πE = πF.Raise(πTemp001, πg.NewStr("subclasses must implement").ToObject(), nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßformat_usage.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 258: def format_heading(self, heading):
					πF.SetLineno(258)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "heading", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("format_heading", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µheading *πg.Object = πArgs[1]; _ = µheading
						var πTemp001 *πg.Object
						_ = πTemp001
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
							// line 259: raise NotImplementedError, "subclasses must implement"
							πF.SetLineno(259)
							πE = πF.Raise(πTemp001, πg.NewStr("subclasses must implement").ToObject(), nil)
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßformat_heading.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 261: def _format_text(self, text):
					πF.SetLineno(261)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "text", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("_format_text", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtext *πg.Object = πArgs[1]; _ = µtext
						var µtext_width *πg.Object = πg.UnboundLocal; _ = µtext_width
						var µindent *πg.Object = πg.UnboundLocal; _ = µindent
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 262: """
							πF.SetLineno(262)
							// line 266: text_width = max(self.width - self.current_indent, 11)
							πF.SetLineno(266)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwidth, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp001[1] = πg.NewInt(11).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtext_width = πTemp003
							// line 267: indent = " "*self.current_indent
							πF.SetLineno(267)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πg.NewStr(" ").ToObject(), πTemp003); πE != nil {
								continue
							}
							µindent = πTemp002
							// line 268: return textwrap.fill(text,
							πF.SetLineno(268)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µtext, "text"); πE != nil {
								continue
							}
							πTemp001[0] = µtext
							if πE = πg.CheckLocal(πF, µtext_width, "text_width"); πE != nil {
								continue
							}
							πTemp001[1] = µtext_width
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"initial_indent", µindent},
								{"subsequent_indent", µindent},
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtextwrap); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßfill, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_format_text.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 262: """
					πF.SetLineno(262)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n        Format a paragraph of free-form text for inclusion in the\n        help output at the current indentation level.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ß_format_text); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 273: def format_description(self, description):
					πF.SetLineno(273)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "description", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("format_description", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdescription *πg.Object = πArgs[1]; _ = µdescription
						var πTemp001 bool
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
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µdescription); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 274: if description:
							πF.SetLineno(274)
						Label1:
							// line 275: return self._format_text(description) + "\n"
							πF.SetLineno(275)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							πTemp003[0] = µdescription
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_format_text, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, πTemp005, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 277: return ""
							πF.SetLineno(277)
							πR = ß.ToObject()
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
					if πE = πClass.SetItem(πF, ßformat_description.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 279: def format_epilog(self, epilog):
					πF.SetLineno(279)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "epilog", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("format_epilog", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µepilog *πg.Object = πArgs[1]; _ = µepilog
						var πTemp001 bool
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
							if πE = πg.CheckLocal(πF, µepilog, "epilog"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µepilog); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 280: if epilog:
							πF.SetLineno(280)
						Label1:
							// line 281: return "\n" + self._format_text(epilog) + "\n"
							πF.SetLineno(281)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µepilog, "epilog"); πE != nil {
								continue
							}
							πTemp004[0] = µepilog
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_format_text, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Add(πF, πg.NewStr("\n").ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 283: return ""
							πF.SetLineno(283)
							πR = ß.ToObject()
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
					if πE = πClass.SetItem(πF, ßformat_epilog.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 286: def expand_default(self, option):
					πF.SetLineno(286)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "option", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("expand_default", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoption *πg.Object = πArgs[1]; _ = µoption
						var µdefault_value *πg.Object = πg.UnboundLocal; _ = µdefault_value
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
							if πTemp004, πE = πg.GetAttr(πF, µself, ßparser, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 == πTemp005).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdefault_tag, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 287: if self.parser is None or not self.default_tag:
							πF.SetLineno(287)
						Label2:
							// line 288: return option.help
							πF.SetLineno(288)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ßhelp, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label3
						Label3:
							// line 290: default_value = self.parser.defaults.get(option.dest)
							πF.SetLineno(290)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßparser, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdefaults, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µdefault_value = πTemp003
							if πE = πg.CheckLocal(πF, µdefault_value, "default_value"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNO_DEFAULT); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µdefault_value == πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µdefault_value, "default_value"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µdefault_value == πTemp004).ToObject()
							πTemp001 = πTemp003
						Label4:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 291: if default_value is NO_DEFAULT or default_value is None:
							πF.SetLineno(291)
						Label5:
							// line 292: default_value = self.NO_DEFAULT_VALUE
							πF.SetLineno(292)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßNO_DEFAULT_VALUE, nil); πE != nil {
								continue
							}
							µdefault_value = πTemp001
							goto Label6
						Label6:
							// line 295: return str(default_value).join(option.help.split(self.default_tag))
							πF.SetLineno(295)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefault_tag, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp001
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ßhelp, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp001
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdefault_value, "default_value"); πE != nil {
								continue
							}
							πTemp008[0] = µdefault_value
							if πTemp001, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
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
					if πE = πClass.SetItem(πF, ßexpand_default.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 297: def format_option(self, option):
					πF.SetLineno(297)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "option", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("format_option", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoption *πg.Object = πArgs[1]; _ = µoption
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µopts *πg.Object = πg.UnboundLocal; _ = µopts
						var µopt_width *πg.Object = πg.UnboundLocal; _ = µopt_width
						var µindent_first *πg.Object = πg.UnboundLocal; _ = µindent_first
						var µhelp_text *πg.Object = πg.UnboundLocal; _ = µhelp_text
						var µhelp_lines *πg.Object = πg.UnboundLocal; _ = µhelp_lines
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 []πg.Param
						_ = πTemp011
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 312: result = []
							πF.SetLineno(312)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresult = πTemp002
							// line 313: opts = self.option_strings[option]
							πF.SetLineno(313)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp002 = µoption
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßoption_strings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							µopts = πTemp003
							// line 314: opt_width = self.help_position - self.current_indent - 2
							πF.SetLineno(314)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßhelp_position, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, πTemp003, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							µopt_width = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							πTemp001[0] = µopts
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µopt_width, "opt_width"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GT(πF, πTemp004, µopt_width); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							goto Label2
							// line 315: if len(opts) > opt_width:
							πF.SetLineno(315)
						Label1:
							// line 317: opts = "%s%s\n" % (self.current_indent * INDENT_CHAR + "", opts)
							πF.SetLineno(317)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ßINDENT_CHAR); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Mul(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, µopts).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s%s\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							µopts = πTemp002
							// line 318: indent_first = self.help_position
							πF.SetLineno(318)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhelp_position, nil); πE != nil {
								continue
							}
							µindent_first = πTemp002
							goto Label3
						Label2:
							// line 321: opts = "%s%s  " % (self.current_indent * INDENT_CHAR, (opts + opt_width * INDENT_CHAR)[:opt_width])
							πF.SetLineno(321)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßINDENT_CHAR); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt_width, "opt_width"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µopt_width, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt_width, "opt_width"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.ResolveGlobal(πF, ßINDENT_CHAR); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Mul(πF, µopt_width, πTemp010); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, µopts, πTemp009); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s%s  ").ToObject(), πTemp003); πE != nil {
								continue
							}
							µopts = πTemp002
							// line 322: indent_first = 0
							πF.SetLineno(322)
							µindent_first = πg.NewInt(0).ToObject()
							goto Label3
						Label3:
							// line 323: result.append(opts)
							πF.SetLineno(323)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							πTemp001[0] = µopts
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßhelp, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µopts, πTemp003); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label5
							}
							goto Label6
							// line 324: if option.help:
							πF.SetLineno(324)
						Label4:
							// line 325: help_text = self.expand_default(option)
							πF.SetLineno(325)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp001[0] = µoption
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßexpand_default, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µhelp_text = πTemp003
							// line 326: help_lines = textwrap.wrap(help_text, self.help_width)
							πF.SetLineno(326)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µhelp_text, "help_text"); πE != nil {
								continue
							}
							πTemp001[0] = µhelp_text
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßhelp_width, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtextwrap); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwrap, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µhelp_lines = πTemp002
							// line 328: result.append("%s%s\n" % (indent_first * INDENT_CHAR + "", help_lines[0]))
							πF.SetLineno(328)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µindent_first, "indent_first"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßINDENT_CHAR); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Mul(πF, µindent_first, πTemp007); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, πTemp005, ß.ToObject()); πE != nil {
								continue
							}
							πTemp005 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µhelp_lines, "help_lines"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, µhelp_lines, πTemp005); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp004, πTemp007).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s%s\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 332: result += (["%s%s\n" % (self.help_position * INDENT_CHAR + "", line)
							πF.SetLineno(332)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp011 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/optparse.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µline *πg.Object = πg.UnboundLocal; _ = µline
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
										if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µhelp_lines, "help_lines"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetItem(πF, µhelp_lines, πTemp002); πE != nil {
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
											µline = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 332: result += (["%s%s\n" % (self.help_position * INDENT_CHAR + "", line)
										πF.SetLineno(332)
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp008, πE = πg.GetAttr(πF, µself, ßhelp_position, nil); πE != nil {
											continue
										}
										if πTemp009, πE = πg.ResolveGlobal(πF, ßINDENT_CHAR); πE != nil {
											continue
										}
										if πTemp007, πE = πg.Mul(πF, πTemp008, πTemp009); πE != nil {
											continue
										}
										if πTemp006, πE = πg.Add(πF, πTemp007, ß.ToObject()); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
											continue
										}
										πTemp003 = πg.NewTuple2(πTemp006, µline).ToObject()
										if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s%s\n").ToObject(), πTemp003); πE != nil {
											continue
										}
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
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IAdd(πF, µresult, πTemp002); πE != nil {
								continue
							}
							µresult = πTemp004
							goto Label6
							// line 334: elif opts[-1] != "\n":
							πF.SetLineno(334)
						Label5:
							// line 335: result.append("\n")
							πF.SetLineno(335)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label6
						Label6:
							// line 336: return "".join(result)
							πF.SetLineno(336)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat_option.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 338: def store_option_strings(self, parser):
					πF.SetLineno(338)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "parser", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("store_option_strings", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µparser *πg.Object = πArgs[1]; _ = µparser
						var µmax_len *πg.Object = πg.UnboundLocal; _ = µmax_len
						var µopt *πg.Object = πg.UnboundLocal; _ = µopt
						var µstrings *πg.Object = πg.UnboundLocal; _ = µstrings
						var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
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
							// line 339: self.indent()
							πF.SetLineno(339)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 340: max_len = 0
							πF.SetLineno(340)
							µmax_len = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparser, ßoption_list, nil); πE != nil {
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
								µopt = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 342: strings = self.format_option_strings(opt)
							πF.SetLineno(342)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp005[0] = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßformat_option_strings, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µstrings = πTemp006
							// line 343: self.option_strings[opt] = strings
							πF.SetLineno(343)
							if πE = πg.CheckLocal(πF, µstrings, "strings"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µstrings); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßoption_strings, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp007 = µopt
							if πE = πg.SetItem(πF, πTemp006, πTemp007, πTemp002); πE != nil {
								continue
							}
							// line 344: max_len = max(max_len, len(strings) + self.current_indent)
							πF.SetLineno(344)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmax_len, "max_len"); πE != nil {
								continue
							}
							πTemp005[0] = µmax_len
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstrings, "strings"); πE != nil {
								continue
							}
							πTemp008[0] = µstrings
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp007, πTemp006); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmax_len = πTemp006
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 345: self.indent()
							πF.SetLineno(345)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßindent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µparser, ßoption_groups, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp003 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
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
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								µgroup = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(4)            
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µgroup, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp004 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp006, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µopt = πTemp006
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 348: strings = self.format_option_strings(opt)
							πF.SetLineno(348)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp005[0] = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßformat_option_strings, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µstrings = πTemp007
							// line 349: self.option_strings[opt] = strings
							πF.SetLineno(349)
							if πE = πg.CheckLocal(πF, µstrings, "strings"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µstrings); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßoption_strings, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp010 = µopt
							if πE = πg.SetItem(πF, πTemp007, πTemp010, πTemp006); πE != nil {
								continue
							}
							// line 350: max_len = max(max_len, len(strings) + self.current_indent)
							πF.SetLineno(350)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmax_len, "max_len"); πE != nil {
								continue
							}
							πTemp005[0] = µmax_len
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstrings, "strings"); πE != nil {
								continue
							}
							πTemp008[0] = µstrings
							if πTemp007, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp010, πTemp007); πE != nil {
								continue
							}
							πTemp005[1] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µmax_len = πTemp007
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
							// line 351: self.dedent()
							πF.SetLineno(351)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 352: self.dedent()
							πF.SetLineno(352)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdedent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 353: self.help_position = min(max_len + 2, self.max_help_position)
							πF.SetLineno(353)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmax_len, "max_len"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µmax_len, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmax_help_position, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhelp_position, πTemp001); πE != nil {
								continue
							}
							// line 354: self.help_width = max(self.width - self.help_position, 11)
							πF.SetLineno(354)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwidth, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßhelp_position, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							πTemp005[1] = πg.NewInt(11).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßmax); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßhelp_width, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßstore_option_strings.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 356: def format_option_strings(self, option):
					πF.SetLineno(356)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "option", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("format_option_strings", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoption *πg.Object = πArgs[1]; _ = µoption
						var µmetavar *πg.Object = πg.UnboundLocal; _ = µmetavar
						var µshort_opts *πg.Object = πg.UnboundLocal; _ = µshort_opts
						var µlong_opts *πg.Object = πg.UnboundLocal; _ = µlong_opts
						var µopts *πg.Object = πg.UnboundLocal; _ = µopts
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []πg.Param
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
							// line 357: """Return a comma-separated list of option strings & metavariables."""
							πF.SetLineno(357)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ßtakes_value, nil); πE != nil {
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
							// line 358: if option.takes_value():
							πF.SetLineno(358)
						Label1:
							// line 359: metavar = option.metavar or option.dest.upper()
							πF.SetLineno(359)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßmetavar, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp002, ßupper, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label4:
							µmetavar = πTemp001
							// line 360: short_opts = [self._short_opt_fmt % (sopt, metavar)
							πF.SetLineno(360)
							πTemp005 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/optparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µsopt *πg.Object = πg.UnboundLocal; _ = µsopt
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
										if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µoption, ß_short_opts, nil); πE != nil {
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
											µsopt = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 360: short_opts = [self._short_opt_fmt % (sopt, metavar)
										πF.SetLineno(360)
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetAttr(πF, µself, ß_short_opt_fmt, nil); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µsopt, "sopt"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µmetavar, "metavar"); πE != nil {
											continue
										}
										πTemp006 = πg.NewTuple2(µsopt, µmetavar).ToObject()
										if πTemp002, πE = πg.Mod(πF, πTemp005, πTemp006); πE != nil {
											continue
										}
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
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µshort_opts = πTemp001
							// line 362: long_opts = [self._long_opt_fmt % (lopt, metavar)
							πF.SetLineno(362)
							πTemp005 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/optparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µlopt *πg.Object = πg.UnboundLocal; _ = µlopt
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
										if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µoption, ß_long_opts, nil); πE != nil {
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
											µlopt = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 362: long_opts = [self._long_opt_fmt % (lopt, metavar)
										πF.SetLineno(362)
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetAttr(πF, µself, ß_long_opt_fmt, nil); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µlopt, "lopt"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µmetavar, "metavar"); πE != nil {
											continue
										}
										πTemp006 = πg.NewTuple2(µlopt, µmetavar).ToObject()
										if πTemp002, πE = πg.Mod(πF, πTemp005, πTemp006); πE != nil {
											continue
										}
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
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							µlong_opts = πTemp001
							goto Label3
						Label2:
							// line 365: short_opts = option._short_opts
							πF.SetLineno(365)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ß_short_opts, nil); πE != nil {
								continue
							}
							µshort_opts = πTemp001
							// line 366: long_opts = option._long_opts
							πF.SetLineno(366)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ß_long_opts, nil); πE != nil {
								continue
							}
							µlong_opts = πTemp001
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßshort_first, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 368: if self.short_first:
							πF.SetLineno(368)
						Label5:
							// line 369: opts = short_opts + long_opts
							πF.SetLineno(369)
							if πE = πg.CheckLocal(πF, µshort_opts, "short_opts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlong_opts, "long_opts"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µshort_opts, µlong_opts); πE != nil {
								continue
							}
							µopts = πTemp001
							goto Label7
						Label6:
							// line 371: opts = long_opts + short_opts
							πF.SetLineno(371)
							if πE = πg.CheckLocal(πF, µlong_opts, "long_opts"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µshort_opts, "short_opts"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µlong_opts, µshort_opts); πE != nil {
								continue
							}
							µopts = πTemp001
							goto Label7
						Label7:
							// line 373: return ", ".join(opts)
							πF.SetLineno(373)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							πTemp007[0] = µopts
							if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πR = πTemp006
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßformat_option_strings.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 357: """Return a comma-separated list of option strings & metavariables."""
					πF.SetLineno(357)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πg.NewStr("Return a comma-separated list of option strings & metavariables.").ToObject()); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßformat_option_strings); πE != nil {
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
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("HelpFormatter").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßHelpFormatter.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 375: class IndentedHelpFormatter (HelpFormatter):
			πF.SetLineno(375)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßHelpFormatter); πE != nil {
				continue
			}
			πTemp001[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("IndentedHelpFormatter", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
					// line 376: """Format help with indented section bodies.
					πF.SetLineno(376)
					// line 376: """Format help with indented section bodies.
					πF.SetLineno(376)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Format help with indented section bodies.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 379: def __init__(self,
					πF.SetLineno(379)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "indent_increment", Def: πg.NewInt(2).ToObject()}
					πTemp002[2] = πg.Param{Name: "max_help_position", Def: πg.NewInt(24).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "width", Def: πTemp003}
					πTemp002[4] = πg.Param{Name: "short_first", Def: πg.NewInt(1).ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindent_increment *πg.Object = πArgs[1]; _ = µindent_increment
						var µmax_help_position *πg.Object = πArgs[2]; _ = µmax_help_position
						var µwidth *πg.Object = πArgs[3]; _ = µwidth
						var µshort_first *πg.Object = πArgs[4]; _ = µshort_first
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
							// line 384: HelpFormatter.__init__(
							πF.SetLineno(384)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µindent_increment, "indent_increment"); πE != nil {
								continue
							}
							πTemp001[1] = µindent_increment
							if πE = πg.CheckLocal(πF, µmax_help_position, "max_help_position"); πE != nil {
								continue
							}
							πTemp001[2] = µmax_help_position
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							πTemp001[3] = µwidth
							if πE = πg.CheckLocal(πF, µshort_first, "short_first"); πE != nil {
								continue
							}
							πTemp001[4] = µshort_first
							if πTemp002, πE = πg.ResolveGlobal(πF, ßHelpFormatter); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 387: def format_usage(self, usage):
					πF.SetLineno(387)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "usage", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("format_usage", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µusage *πg.Object = πArgs[1]; _ = µusage
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
							// line 388: return _("Usage: %s\n") % usage
							πF.SetLineno(388)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("Usage: %s\n").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp004, µusage); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat_usage.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 390: def format_heading(self, heading):
					πF.SetLineno(390)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "heading", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("format_heading", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µheading *πg.Object = πArgs[1]; _ = µheading
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 392: return "%s%s:\n" % (self.current_indent * INDENT_CHAR + "", heading)
							πF.SetLineno(392)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcurrent_indent, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßINDENT_CHAR); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mul(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µheading, "heading"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, µheading).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s%s:\n").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat_heading.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("IndentedHelpFormatter").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßIndentedHelpFormatter.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 395: class TitledHelpFormatter (HelpFormatter):
			πF.SetLineno(395)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßHelpFormatter); πE != nil {
				continue
			}
			πTemp001[0] = πTemp009
			πTemp007 = πg.NewDict()
			if πTemp006, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp006); πE != nil {
				continue
			}
			_, πE = πg.NewCode("TitledHelpFormatter", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
					// line 396: """Format help with underlined section headers.
					πF.SetLineno(396)
					// line 396: """Format help with underlined section headers.
					πF.SetLineno(396)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Format help with underlined section headers.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 399: def __init__(self,
					πF.SetLineno(399)
					πTemp002 = make([]πg.Param, 5)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "indent_increment", Def: πg.NewInt(0).ToObject()}
					πTemp002[2] = πg.Param{Name: "max_help_position", Def: πg.NewInt(24).ToObject()}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "width", Def: πTemp003}
					πTemp002[4] = πg.Param{Name: "short_first", Def: πg.NewInt(0).ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µindent_increment *πg.Object = πArgs[1]; _ = µindent_increment
						var µmax_help_position *πg.Object = πArgs[2]; _ = µmax_help_position
						var µwidth *πg.Object = πArgs[3]; _ = µwidth
						var µshort_first *πg.Object = πArgs[4]; _ = µshort_first
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
							// line 404: HelpFormatter.__init__ (
							πF.SetLineno(404)
							πTemp001 = πF.MakeArgs(5)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µindent_increment, "indent_increment"); πE != nil {
								continue
							}
							πTemp001[1] = µindent_increment
							if πE = πg.CheckLocal(πF, µmax_help_position, "max_help_position"); πE != nil {
								continue
							}
							πTemp001[2] = µmax_help_position
							if πE = πg.CheckLocal(πF, µwidth, "width"); πE != nil {
								continue
							}
							πTemp001[3] = µwidth
							if πE = πg.CheckLocal(πF, µshort_first, "short_first"); πE != nil {
								continue
							}
							πTemp001[4] = µshort_first
							if πTemp002, πE = πg.ResolveGlobal(πF, ßHelpFormatter); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 407: def format_usage(self, usage):
					πF.SetLineno(407)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "usage", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("format_usage", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µusage *πg.Object = πArgs[1]; _ = µusage
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
							// line 408: return "%s  %s\n" % (self.format_heading(_("Usage")), usage)
							πF.SetLineno(408)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = ßUsage.ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßformat_heading, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp006, µusage).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s  %s\n").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat_usage.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 410: def format_heading(self, heading):
					πF.SetLineno(410)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "heading", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("format_heading", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µheading *πg.Object = πArgs[1]; _ = µheading
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
							// line 411: return "%s\n%s\n" % (heading, "=-"[self.level] * len(heading))
							πF.SetLineno(411)
							if πE = πg.CheckLocal(πF, µheading, "heading"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßlevel, nil); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πTemp005, πE = πg.GetItem(πF, πg.NewStr("=-").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheading, "heading"); πE != nil {
								continue
							}
							πTemp006[0] = µheading
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp003, πE = πg.Mul(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µheading, πTemp003).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s\n%s\n").ToObject(), πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat_heading.ToObject(), πTemp004); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp008, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp008 == nil {
				πTemp008 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("TitledHelpFormatter").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßTitledHelpFormatter.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 414: def _parse_num(val, type):
			πF.SetLineno(414)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "val", Def: nil}
			πTemp003[1] = πg.Param{Name: "type", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("_parse_num", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µval *πg.Object = πArgs[0]; _ = µval
				var µtype *πg.Object = πArgs[1]; _ = µtype
				var µradix *πg.Object = πg.UnboundLocal; _ = µradix
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µval, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlower, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, ß0x.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µval, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlower, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, ß0b.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µval, πTemp002); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp003, ß0.ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 415: if val[:2].lower() == "0x":         # hexadecimal
					πF.SetLineno(415)
				Label1:
					// line 416: radix = 16
					πF.SetLineno(416)
					µradix = πg.NewInt(16).ToObject()
					goto Label5
					// line 417: elif val[:2].lower() == "0b":       # binary
					πF.SetLineno(417)
				Label2:
					// line 418: radix = 2
					πF.SetLineno(418)
					µradix = πg.NewInt(2).ToObject()
					// line 419: val = val[2:] or "0"            # have to remove "0b" prefix
					πF.SetLineno(419)
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µval, πTemp002); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					πTemp001 = ß0.ToObject()
				Label6:
					µval = πTemp001
					goto Label5
					// line 420: elif val[:1] == "0":                # octal
					πF.SetLineno(420)
				Label3:
					// line 421: radix = 8
					πF.SetLineno(421)
					µradix = πg.NewInt(8).ToObject()
					goto Label5
				Label4:
					// line 423: radix = 10
					πF.SetLineno(423)
					µradix = πg.NewInt(10).ToObject()
					goto Label5
				Label5:
					// line 425: return type(val, radix)
					πF.SetLineno(425)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πTemp005[0] = µval
					if πE = πg.CheckLocal(πF, µradix, "radix"); πE != nil {
						continue
					}
					πTemp005[1] = µradix
					if πE = πg.CheckLocal(πF, µtype, "type"); πE != nil {
						continue
					}
					if πTemp001, πE = µtype.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
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
			if πE = πF.Globals().SetItem(πF, ß_parse_num.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 427: def _parse_int(val):
			πF.SetLineno(427)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "val", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("_parse_int", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µval *πg.Object = πArgs[0]; _ = µval
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
					// line 428: return _parse_num(val, int)
					πF.SetLineno(428)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πTemp001[0] = µval
					if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse_num); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_parse_int.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 430: def _parse_long(val):
			πF.SetLineno(430)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "val", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("_parse_long", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µval *πg.Object = πArgs[0]; _ = µval
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
					// line 431: return _parse_num(val, long)
					πF.SetLineno(431)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
						continue
					}
					πTemp001[0] = µval
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse_num); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_parse_long.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 433: _builtin_cvt = { "int" : (_parse_int, _("integer")),
			πF.SetLineno(433)
			πTemp007 = πg.NewDict()
			if πTemp011, πE = πg.ResolveGlobal(πF, ß_parse_int); πE != nil {
				continue
			}
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = ßinteger.ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			πTemp010 = πg.NewTuple2(πTemp011, πTemp013).ToObject()
			if πE = πTemp007.SetItem(πF, ßint.ToObject(), πTemp010); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ß_parse_long); πE != nil {
				continue
			}
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr("long integer").ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			πTemp010 = πg.NewTuple2(πTemp011, πTemp013).ToObject()
			if πE = πTemp007.SetItem(πF, ßlong.ToObject(), πTemp010); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
				continue
			}
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewStr("floating-point").ToObject()
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
				continue
			}
			if πTemp013, πE = πTemp012.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			πTemp010 = πg.NewTuple2(πTemp011, πTemp013).ToObject()
			if πE = πTemp007.SetItem(πF, ßfloat.ToObject(), πTemp010); πE != nil {
				continue
			}
			πTemp010 = πTemp007.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_builtin_cvt.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 438: def check_builtin(option, opt, value):
			πF.SetLineno(438)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "option", Def: nil}
			πTemp003[1] = πg.Param{Name: "opt", Def: nil}
			πTemp003[2] = πg.Param{Name: "value", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("check_builtin", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µoption *πg.Object = πArgs[0]; _ = µoption
				var µopt *πg.Object = πArgs[1]; _ = µopt
				var µvalue *πg.Object = πArgs[2]; _ = µvalue
				var µcvt *πg.Object = πg.UnboundLocal; _ = µcvt
				var µwhat *πg.Object = πg.UnboundLocal; _ = µwhat
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []*πg.Object
				_ = πTemp004
				var πTemp005 *πg.BaseException
				_ = πTemp005
				var πTemp006 *πg.Traceback
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 439: (cvt, what) = _builtin_cvt[option.type]
					πF.SetLineno(439)
					if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µoption, ßtype, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_builtin_cvt); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp003}}}, πTemp002); πE != nil {
						continue
					}
					µcvt = πTemp001
					µwhat = πTemp003
					// line 440: try:
					πF.SetLineno(440)
					πF.PushCheckpoint(2)
					// line 441: return cvt(value)
					πF.SetLineno(441)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp004[0] = µvalue
					if πE = πg.CheckLocal(πF, µcvt, "cvt"); πE != nil {
						continue
					}
					if πTemp001, πE = µcvt.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πR = πTemp001
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp005, πTemp006 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label3
					}
					πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
					continue
					// line 442: except ValueError:
					πF.SetLineno(442)
				Label3:
					πTemp004 = πF.MakeArgs(1)
					πTemp008 = πF.MakeArgs(1)
					πTemp008[0] = πg.NewStr("option %s: invalid %s value: %r").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwhat, "what"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µopt, µwhat, µvalue).ToObject()
					if πTemp001, πE = πg.Mod(πF, πTemp003, πTemp002); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 443: raise OptionValueError(
					πF.SetLineno(443)
					πE = πF.Raise(πTemp002, nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ßcheck_builtin.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 446: def check_choice(option, opt, value):
			πF.SetLineno(446)
			πTemp003 = make([]πg.Param, 3)
			πTemp003[0] = πg.Param{Name: "option", Def: nil}
			πTemp003[1] = πg.Param{Name: "opt", Def: nil}
			πTemp003[2] = πg.Param{Name: "value", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("check_choice", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µoption *πg.Object = πArgs[0]; _ = µoption
				var µopt *πg.Object = πArgs[1]; _ = µopt
				var µvalue *πg.Object = πArgs[2]; _ = µvalue
				var µchoices *πg.Object = πg.UnboundLocal; _ = µchoices
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µoption, ßchoices, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Contains(πF, πTemp002, µvalue); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp003).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 447: if value in option.choices:
					πF.SetLineno(447)
				Label1:
					// line 448: return value
					πF.SetLineno(448)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πR = µvalue
					continue
					goto Label3
				Label2:
					// line 450: choices = ", ".join(map(repr, option.choices))
					πF.SetLineno(450)
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(2)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					πTemp005[0] = πTemp001
					if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µoption, ßchoices, nil); πE != nil {
						continue
					}
					πTemp005[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µchoices = πTemp002
					πTemp004 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = πg.NewStr("option %s: invalid choice: %r (choose from %s)").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µchoices, "choices"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µopt, µvalue, µchoices).ToObject()
					if πTemp001, πE = πg.Mod(πF, πTemp006, πTemp002); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 451: raise OptionValueError(
					πF.SetLineno(451)
					πE = πF.Raise(πTemp002, nil, nil)
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
			if πE = πF.Globals().SetItem(πF, ßcheck_choice.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 457: NO_DEFAULT = ("NO", "DEFAULT")
			πF.SetLineno(457)
			πTemp012 = πg.NewTuple2(ßNO.ToObject(), ßDEFAULT.ToObject()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßNO_DEFAULT.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 460: class Option(object):
			πF.SetLineno(460)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp014
			πTemp007 = πg.NewDict()
			if πTemp012, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp012); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Option", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Dict
				_ = πTemp003
				var πTemp004 []πg.Param
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 461: """
					πF.SetLineno(461)
					// line 461: """
					πF.SetLineno(461)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Instance attributes:\n      _short_opts : [string]\n      _long_opts : [string]\n\n      action : string\n      type : string\n      dest : string\n      default : any\n      nargs : int\n      const : any\n      choices : [string]\n      callback : function\n      callback_args : (any*)\n      callback_kwargs : { string : any }\n      help : string\n      metavar : string\n    ").ToObject()); πE != nil {
						continue
					}
					// line 482: ATTRS = ['action',
					πF.SetLineno(482)
					πTemp001 = make([]*πg.Object, 12)
					πTemp001[0] = ßaction.ToObject()
					πTemp001[1] = ßtype.ToObject()
					πTemp001[2] = ßdest.ToObject()
					πTemp001[3] = ßdefault.ToObject()
					πTemp001[4] = ßnargs.ToObject()
					πTemp001[5] = ßconst.ToObject()
					πTemp001[6] = ßchoices.ToObject()
					πTemp001[7] = ßcallback.ToObject()
					πTemp001[8] = ßcallback_args.ToObject()
					πTemp001[9] = ßcallback_kwargs.ToObject()
					πTemp001[10] = ßhelp.ToObject()
					πTemp001[11] = ßmetavar.ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ßATTRS.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 497: ACTIONS = ("store",
					πF.SetLineno(497)
					πTemp001 = make([]*πg.Object, 10)
					πTemp001[0] = ßstore.ToObject()
					πTemp001[1] = ßstore_const.ToObject()
					πTemp001[2] = ßstore_true.ToObject()
					πTemp001[3] = ßstore_false.ToObject()
					πTemp001[4] = ßappend.ToObject()
					πTemp001[5] = ßappend_const.ToObject()
					πTemp001[6] = ßcount.ToObject()
					πTemp001[7] = ßcallback.ToObject()
					πTemp001[8] = ßhelp.ToObject()
					πTemp001[9] = ßversion.ToObject()
					πTemp002 = πg.NewTuple(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ßACTIONS.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 511: STORE_ACTIONS = ("store",
					πF.SetLineno(511)
					πTemp001 = make([]*πg.Object, 7)
					πTemp001[0] = ßstore.ToObject()
					πTemp001[1] = ßstore_const.ToObject()
					πTemp001[2] = ßstore_true.ToObject()
					πTemp001[3] = ßstore_false.ToObject()
					πTemp001[4] = ßappend.ToObject()
					πTemp001[5] = ßappend_const.ToObject()
					πTemp001[6] = ßcount.ToObject()
					πTemp002 = πg.NewTuple(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ßSTORE_ACTIONS.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 521: TYPED_ACTIONS = ("store",
					πF.SetLineno(521)
					πTemp002 = πg.NewTuple3(ßstore.ToObject(), ßappend.ToObject(), ßcallback.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßTYPED_ACTIONS.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 527: ALWAYS_TYPED_ACTIONS = ("store",
					πF.SetLineno(527)
					πTemp002 = πg.NewTuple2(ßstore.ToObject(), ßappend.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßALWAYS_TYPED_ACTIONS.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 531: CONST_ACTIONS = ("store_const",
					πF.SetLineno(531)
					πTemp002 = πg.NewTuple2(ßstore_const.ToObject(), ßappend_const.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßCONST_ACTIONS.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 536: TYPES = ("string", "int", "long", "float", "choice") #, "complex"
					πF.SetLineno(536)
					πTemp002 = πg.NewTuple5(ßstring.ToObject(), ßint.ToObject(), ßlong.ToObject(), ßfloat.ToObject(), ßchoice.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ßTYPES.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 554: TYPE_CHECKER = { "int"    : check_builtin,
					πF.SetLineno(554)
					πTemp003 = πg.NewDict()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_builtin); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßint.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_builtin); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßlong.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_builtin); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßfloat.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_choice); πE != nil {
						continue
					}
					if πE = πTemp003.SetItem(πF, ßchoice.ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp002 = πTemp003.ToObject()
					if πE = πClass.SetItem(πF, ßTYPE_CHECKER.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 570: CHECK_METHODS = None
					πF.SetLineno(570)
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßCHECK_METHODS.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 575: def __init__(self, *opts, **attrs):
					πF.SetLineno(575)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp004, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopts *πg.Object = πArgs[1]; _ = µopts
						var µattrs *πg.Object = πArgs[2]; _ = µattrs
						var µchecker *πg.Object = πg.UnboundLocal; _ = µchecker
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
							// line 578: self._short_opts = []
							πF.SetLineno(578)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_short_opts, πTemp003); πE != nil {
								continue
							}
							// line 579: self._long_opts = []
							πF.SetLineno(579)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_long_opts, πTemp003); πE != nil {
								continue
							}
							// line 580: opts = self._check_opt_strings(opts)
							πF.SetLineno(580)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							πTemp001[0] = µopts
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_check_opt_strings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µopts = πTemp003
							// line 581: self._set_opt_strings(opts)
							πF.SetLineno(581)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							πTemp001[0] = µopts
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_set_opt_strings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 584: self._set_attrs(attrs)
							πF.SetLineno(584)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							πTemp001[0] = µattrs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_set_attrs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßCHECK_METHODS, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µchecker = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 592: checker(self)
							πF.SetLineno(592)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µchecker, "checker"); πE != nil {
								continue
							}
							if πTemp003, πE = µchecker.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 594: def _check_opt_strings(self, opts):
					πF.SetLineno(594)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "opts", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_check_opt_strings", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopts *πg.Object = πArgs[1]; _ = µopts
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 599: opts = [x for x in opts if x is not None]
							πF.SetLineno(599)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πg.UnboundLocal; _ = µx
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
										case 6: goto Label6
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µopts); πE != nil {
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
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
											continue
										}
										πTemp004 = πg.GetBool(µx != πTemp005).ToObject()
										if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
											continue
										}
										if πTemp003 {
											goto Label4
										}
										goto Label5
										// line 599: opts = [x for x in opts if x is not None]
										πF.SetLineno(599)
									Label4:
										// line 599: opts = [x for x in opts if x is not None]
										πF.SetLineno(599)
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µx, nil
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
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µopts = πTemp001
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µopts); πE != nil {
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
							// line 600: if not opts:
							πF.SetLineno(600)
						Label1:
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("at least one option string must be supplied").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 601: raise TypeError("at least one option string must be supplied")
							πF.SetLineno(601)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 602: return opts
							πF.SetLineno(602)
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							πR = µopts
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_check_opt_strings.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 604: def _set_opt_strings(self, opts):
					πF.SetLineno(604)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "opts", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_set_opt_strings", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopts *πg.Object = πArgs[1]; _ = µopts
						var µopt *πg.Object = πg.UnboundLocal; _ = µopt
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
							if πE = πg.CheckLocal(πF, µopts, "opts"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µopts); πE != nil {
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
								µopt = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp005[0] = µopt
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.LT(πF, πTemp007, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp005[0] = µopt
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.Eq(πF, πTemp007, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 606: if len(opt) < 2:
							πF.SetLineno(606)
						Label4:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("invalid option string %r: must be at least two characters long").ToObject(), µopt); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 607: raise OptionError(
							πF.SetLineno(607)
							πE = πF.Raise(πTemp006, nil, nil)
							continue
							goto Label7
							// line 610: elif len(opt) == 2:
							πF.SetLineno(610)
						Label5:
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µopt, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Eq(πF, πTemp009, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label8
							}
							πTemp008 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µopt, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.NE(πF, πTemp009, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
						Label8:
							if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 611: if not (opt[0] == "-" and opt[1] != "-"):
							πF.SetLineno(611)
						Label9:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("invalid short option string %r: must be of the form -x, (x any non-dash char)").ToObject(), µopt); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 612: raise OptionError(
							πF.SetLineno(612)
							πE = πF.Raise(πTemp006, nil, nil)
							continue
							goto Label10
						Label10:
							// line 616: self._short_opts.append(opt)
							πF.SetLineno(616)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp005[0] = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_short_opts, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label7
						Label6:
							if πTemp008, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µopt, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Eq(πF, πTemp009, πg.NewStr("--").ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
							if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label11
							}
							πTemp008 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µopt, πTemp008); πE != nil {
								continue
							}
							if πTemp007, πE = πg.NE(πF, πTemp009, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp007
						Label11:
							if πTemp003, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							goto Label13
							// line 618: if not (opt[0:2] == "--" and opt[2] != "-"):
							πF.SetLineno(618)
						Label12:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("invalid long option string %r: must start with --, followed by non-dash").ToObject(), µopt); πE != nil {
								continue
							}
							πTemp005[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp004, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 619: raise OptionError(
							πF.SetLineno(619)
							πE = πF.Raise(πTemp006, nil, nil)
							continue
							goto Label13
						Label13:
							// line 623: self._long_opts.append(opt)
							πF.SetLineno(623)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp005[0] = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label7
						Label7:
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
					if πE = πClass.SetItem(πF, ß_set_opt_strings.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 625: def _set_attrs(self, attrs):
					πF.SetLineno(625)
					πTemp004 = make([]πg.Param, 2)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "attrs", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_set_attrs", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µattrs *πg.Object = πArgs[1]; _ = µattrs
						var µattr *πg.Object = πg.UnboundLocal; _ = µattr
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßATTRS, nil); πE != nil {
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
								µattr = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µattrs, µattr); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µattr, ßdefault.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 627: if attr in attrs:
							πF.SetLineno(627)
						Label4:
							// line 628: setattr(self, attr, attrs[attr])
							πF.SetLineno(628)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp005[1] = µattr
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002 = µattr
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µattrs, πTemp002); πE != nil {
								continue
							}
							πTemp005[2] = πTemp006
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 629: del attrs[attr]
							πF.SetLineno(629)
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002 = µattr
							if πE = πg.DelItem(πF, µattrs, πTemp002); πE != nil {
								continue
							}
							goto Label7
							// line 631: if attr == 'default':
							πF.SetLineno(631)
						Label5:
							// line 632: setattr(self, attr, NO_DEFAULT)
							πF.SetLineno(632)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp005[1] = µattr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNO_DEFAULT); πE != nil {
								continue
							}
							πTemp005[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label7
						Label6:
							// line 634: setattr(self, attr, None)
							πF.SetLineno(634)
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp005[1] = µattr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[2] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label7
						Label7:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µattrs); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 635: if attrs:
							πF.SetLineno(635)
						Label8:
							// line 636: attrs = attrs.keys()
							πF.SetLineno(636)
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µattrs, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µattrs = πTemp002
							// line 637: attrs.sort()
							πF.SetLineno(637)
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µattrs, ßsort, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(2)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µattrs, "attrs"); πE != nil {
								continue
							}
							πTemp007[0] = µattrs
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid keyword arguments: %s").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 638: raise OptionError(
							πF.SetLineno(638)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label9
						Label9:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_set_attrs.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 645: def _check_action(self):
					πF.SetLineno(645)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_check_action", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßACTIONS, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							goto Label3
							// line 646: if self.action is None:
							πF.SetLineno(646)
						Label1:
							// line 647: self.action = "store"
							πF.SetLineno(647)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ßstore.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßaction, πTemp001); πE != nil {
								continue
							}
							goto Label3
							// line 648: elif self.action not in self.ACTIONS:
							πF.SetLineno(648)
						Label2:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid action: %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 649: raise OptionError("invalid action: %r" % self.action, self)
							πF.SetLineno(649)
							πE = πF.Raise(πTemp002, nil, nil)
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
					if πE = πClass.SetItem(πF, ß_check_action.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 651: def _check_type(self):
					πF.SetLineno(651)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("_check_type", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µ__builtin__ *πg.Object = πg.UnboundLocal; _ = µ__builtin__
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
						var πTemp006 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 652: if self.type is None:
							πF.SetLineno(652)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßALWAYS_TYPED_ACTIONS, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 653: if self.action in self.ALWAYS_TYPED_ACTIONS:
							πF.SetLineno(653)
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchoices, nil); πE != nil {
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
								goto Label6
							}
							goto Label7
							// line 654: if self.choices is not None:
							πF.SetLineno(654)
						Label6:
							// line 656: self.type = "choice"
							πF.SetLineno(656)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ßchoice.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtype, πTemp001); πE != nil {
								continue
							}
							goto Label8
						Label7:
							// line 659: self.type = "string"
							πF.SetLineno(659)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ßstring.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtype, πTemp001); πE != nil {
								continue
							}
							goto Label8
						Label8:
							goto Label5
						Label5:
							goto Label3
						Label2:
							// line 666: import __builtin__
							πF.SetLineno(666)
							if πTemp005, πE = πg.ImportModule(πF, "__builtin__"); πE != nil {
								continue
							}
							πTemp001 = πTemp005[0]
							µ__builtin__ = πTemp001
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßTypeType, nil); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006 == πTemp007).ToObject()
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							πTemp005[1] = ß__name__.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp002 = πTemp006
							if πTemp008, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label10
							}
							πTemp005 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µ__builtin__, "__builtin__"); πE != nil {
								continue
							}
							πTemp005[0] = µ__builtin__
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß__name__, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp005[2] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007 == πTemp006).ToObject()
							πTemp002 = πTemp003
						Label10:
							πTemp001 = πTemp002
						Label9:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 667: if ( type(self.type) is types.TypeType or
							πF.SetLineno(667)
						Label11:
							// line 670: self.type = self.type.__name__
							πF.SetLineno(670)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtype, πTemp001); πE != nil {
								continue
							}
							goto Label12
						Label12:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßstr.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							goto Label14
							// line 672: if self.type == "str":
							πF.SetLineno(672)
						Label13:
							// line 673: self.type = "string"
							πF.SetLineno(673)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ßstring.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtype, πTemp001); πE != nil {
								continue
							}
							goto Label14
						Label14:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßTYPES, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							goto Label16
							// line 675: if self.type not in self.TYPES:
							πF.SetLineno(675)
						Label15:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("invalid option type: %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 676: raise OptionError("invalid option type: %r" % self.type, self)
							πF.SetLineno(676)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label16
						Label16:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßTYPED_ACTIONS, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label17
							}
							goto Label18
							// line 677: if self.action not in self.TYPED_ACTIONS:
							πF.SetLineno(677)
						Label17:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("must not supply a type for action %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 678: raise OptionError(
							πF.SetLineno(678)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label18
						Label18:
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
					if πE = πClass.SetItem(πF, ß_check_type.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 681: def _check_choice(self):
					πF.SetLineno(681)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("_check_choice", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
						var πTemp010 []*πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßchoice.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchoices, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp004).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 682: if self.type == "choice":
							πF.SetLineno(682)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchoices, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp004).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßchoices, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßTupleType, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßListType, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							if πTemp003, πE = πg.Contains(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 683: if self.choices is None:
							πF.SetLineno(683)
						Label4:
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("must supply a list of choices for type 'choice'").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 684: raise OptionError(
							πF.SetLineno(684)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label6
							// line 686: elif type(self.choices) not in (types.TupleType, types.ListType):
							πF.SetLineno(686)
						Label5:
							πTemp005 = πF.MakeArgs(2)
							πTemp002 = πg.NewInt(1).ToObject()
							πTemp009 = πF.MakeArgs(1)
							πTemp009[0] = πg.NewStr("'").ToObject()
							πTemp010 = πF.MakeArgs(1)
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßchoices, nil); πE != nil {
								continue
							}
							πTemp011[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp010[0] = πTemp007
							if πTemp006, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp004, πE = πg.GetItem(πF, πTemp007, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("choices must be a list of strings ('%s' supplied)").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 687: raise OptionError(
							πF.SetLineno(687)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label6
						Label6:
							goto Label3
							// line 690: elif self.choices is not None:
							πF.SetLineno(690)
						Label2:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("must not supply choices for type %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 691: raise OptionError(
							πF.SetLineno(691)
							πE = πF.Raise(πTemp002, nil, nil)
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
					if πE = πClass.SetItem(πF, ß_check_choice.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 694: def _check_dest(self):
					πF.SetLineno(694)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("_check_dest", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtakes_value *πg.Object = πg.UnboundLocal; _ = µtakes_value
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
						var πTemp009 *πg.Object
						_ = πTemp009
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 697: takes_value = (self.action in self.STORE_ACTIONS or
							πF.SetLineno(697)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßSTORE_ACTIONS, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 != πTemp005).ToObject()
							πTemp001 = πTemp003
						Label1:
							µtakes_value = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdest, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 == πTemp005).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µtakes_value, "takes_value"); πE != nil {
								continue
							}
							πTemp001 = µtakes_value
						Label2:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 699: if self.dest is None and takes_value:
							πF.SetLineno(699)
						Label3:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 703: if self._long_opts:
							πF.SetLineno(703)
						Label5:
							// line 706: self.dest = '_'.join(self._long_opts[0][2:].split('-'))
							πF.SetLineno(706)
							πTemp007 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("-").ToObject()
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp007[0] = πTemp003
							if πTemp001, πE = πg.GetAttr(πF, ß_.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdest, πTemp001); πE != nil {
								continue
							}
							goto Label7
						Label6:
							// line 708: self.dest = self._short_opts[0][1]
							πF.SetLineno(708)
							πTemp001 = πg.NewInt(1).ToObject()
							πTemp004 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µself, ß_short_opts, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdest, πTemp001); πE != nil {
								continue
							}
							goto Label7
						Label7:
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
					if πE = πClass.SetItem(πF, ß_check_dest.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 710: def _check_const(self):
					πF.SetLineno(710)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("_check_const", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßCONST_ACTIONS, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp005, πTemp004); πE != nil {
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßconst, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp004 != πTemp005).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 711: if self.action not in self.CONST_ACTIONS and self.const is not None:
							πF.SetLineno(711)
						Label2:
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("'const' must not be supplied for action %r").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp007[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 712: raise OptionError(
							πF.SetLineno(712)
							πE = πF.Raise(πTemp003, nil, nil)
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
					if πE = πClass.SetItem(πF, ß_check_const.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 716: def _check_nargs(self):
					πF.SetLineno(716)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("_check_nargs", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßTYPED_ACTIONS, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnargs, nil); πE != nil {
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
								goto Label2
							}
							goto Label3
							// line 717: if self.action in self.TYPED_ACTIONS:
							πF.SetLineno(717)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnargs, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 718: if self.nargs is None:
							πF.SetLineno(718)
						Label4:
							// line 719: self.nargs = 1
							πF.SetLineno(719)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnargs, πTemp001); πE != nil {
								continue
							}
							goto Label5
						Label5:
							goto Label3
							// line 720: elif self.nargs is not None:
							πF.SetLineno(720)
						Label2:
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("'nargs' must not be supplied for action %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 721: raise OptionError(
							πF.SetLineno(721)
							πE = πF.Raise(πTemp002, nil, nil)
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
					if πE = πClass.SetItem(πF, ß_check_nargs.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 725: def _check_callback(self):
					πF.SetLineno(725)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("_check_callback", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßcallback.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 726: if self.action == "callback":
							πF.SetLineno(726)
						Label1:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcallback, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							πTemp004[1] = ß__call__.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
							// line 727: if not hasattr(self.callback, '__call__'):
							πF.SetLineno(727)
						Label4:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcallback, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("callback not callable: %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 728: raise OptionError(
							πF.SetLineno(728)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcallback_args, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005 != πTemp006).ToObject()
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label6
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcallback_args, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßTupleType, nil); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006 != πTemp007).ToObject()
							πTemp001 = πTemp002
						Label6:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 730: if (self.callback_args is not None and
							πF.SetLineno(730)
						Label7:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcallback_args, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("callback_args, if supplied, must be a tuple: not %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 732: raise OptionError(
							πF.SetLineno(732)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label8
						Label8:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcallback_kwargs, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp005 != πTemp006).ToObject()
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label9
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßcallback_kwargs, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßDictType, nil); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp006 != πTemp007).ToObject()
							πTemp001 = πTemp002
						Label9:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							goto Label11
							// line 735: if (self.callback_kwargs is not None and
							πF.SetLineno(735)
						Label10:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcallback_kwargs, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("callback_kwargs, if supplied, must be a dict: not %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 737: raise OptionError(
							πF.SetLineno(737)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label11
						Label11:
							goto Label3
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcallback, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp005).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							goto Label13
							// line 741: if self.callback is not None:
							πF.SetLineno(741)
						Label12:
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcallback, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("callback supplied (%r) for non-callback option").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 742: raise OptionError(
							πF.SetLineno(742)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcallback_args, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp005).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label14
							}
							goto Label15
							// line 745: if self.callback_args is not None:
							πF.SetLineno(745)
						Label14:
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("callback_args supplied for non-callback option").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 746: raise OptionError(
							πF.SetLineno(746)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label15
						Label15:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcallback_kwargs, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp005).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label16
							}
							goto Label17
							// line 748: if self.callback_kwargs is not None:
							πF.SetLineno(748)
						Label16:
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("callback_kwargs supplied for non-callback option").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[1] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 749: raise OptionError(
							πF.SetLineno(749)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label17
						Label17:
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
					if πE = πClass.SetItem(πF, ß_check_callback.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 753: CHECK_METHODS = [_check_action,
					πF.SetLineno(753)
					πTemp001 = make([]*πg.Object, 7)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ß_check_action); πE != nil {
						continue
					}
					πTemp001[0] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ß_check_type); πE != nil {
						continue
					}
					πTemp001[1] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ß_check_choice); πE != nil {
						continue
					}
					πTemp001[2] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ß_check_dest); πE != nil {
						continue
					}
					πTemp001[3] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ß_check_const); πE != nil {
						continue
					}
					πTemp001[4] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ß_check_nargs); πE != nil {
						continue
					}
					πTemp001[5] = πTemp015
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ß_check_callback); πE != nil {
						continue
					}
					πTemp001[6] = πTemp015
					πTemp015 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ßCHECK_METHODS.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 764: def __str__(self):
					πF.SetLineno(764)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 765: return "/".join(self._short_opts + self._long_opts)
							πF.SetLineno(765)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_short_opts, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("/").ToObject(), ßjoin, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 767: __repr__ = _repr
					πF.SetLineno(767)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ß_repr); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 769: def takes_value(self):
					πF.SetLineno(769)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("takes_value", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 770: return self.type is not None
							πF.SetLineno(770)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != πTemp003).ToObject()
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
					if πE = πClass.SetItem(πF, ßtakes_value.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 772: def get_opt_string(self):
					πF.SetLineno(772)
					πTemp004 = make([]πg.Param, 1)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("get_opt_string", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 773: if self._long_opts:
							πF.SetLineno(773)
						Label1:
							// line 774: return self._long_opts[0]
							πF.SetLineno(774)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label3
						Label2:
							// line 776: return self._short_opts[0]
							πF.SetLineno(776)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_short_opts, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πR = πTemp003
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
					if πE = πClass.SetItem(πF, ßget_opt_string.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 781: def check_value(self, opt, value):
					πF.SetLineno(781)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "opt", Def: nil}
					πTemp004[2] = πg.Param{Name: "value", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("check_value", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt *πg.Object = πArgs[1]; _ = µopt
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var µchecker *πg.Object = πg.UnboundLocal; _ = µchecker
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
							// line 782: checker = self.TYPE_CHECKER.get(self.type)
							πF.SetLineno(782)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtype, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßTYPE_CHECKER, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µchecker = πTemp002
							if πE = πg.CheckLocal(πF, µchecker, "checker"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µchecker == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 783: if checker is None:
							πF.SetLineno(783)
						Label1:
							// line 784: return value
							πF.SetLineno(784)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πR = µvalue
							continue
							goto Label3
						Label2:
							// line 786: return checker(self, opt, value)
							πF.SetLineno(786)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp001[1] = µopt
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[2] = µvalue
							if πE = πg.CheckLocal(πF, µchecker, "checker"); πE != nil {
								continue
							}
							if πTemp002, πE = µchecker.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßcheck_value.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 788: def convert_value(self, opt, value):
					πF.SetLineno(788)
					πTemp004 = make([]πg.Param, 3)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "opt", Def: nil}
					πTemp004[2] = πg.Param{Name: "value", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("convert_value", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt *πg.Object = πArgs[1]; _ = µopt
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []πg.Param
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
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µvalue != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 789: if value is not None:
							πF.SetLineno(789)
						Label1:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnargs, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 790: if self.nargs == 1:
							πF.SetLineno(790)
						Label3:
							// line 791: return self.check_value(opt, value)
							πF.SetLineno(791)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp004[0] = µopt
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcheck_value, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp002
							continue
							goto Label5
						Label4:
							// line 793: return tuple([self.check_value(opt, v) for v in value])
							πF.SetLineno(793)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/optparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µvalue); πE != nil {
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
										// line 793: return tuple([self.check_value(opt, v) for v in value])
										πF.SetLineno(793)
										πTemp005 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
											continue
										}
										πTemp005[0] = µopt
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										πTemp005[1] = µv
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp004, πE = πg.GetAttr(πF, µself, ßcheck_value, nil); πE != nil {
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
							if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp006
							continue
							goto Label5
						Label5:
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
					if πE = πClass.SetItem(πF, ßconvert_value.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 795: def process(self, opt, value, values, parser):
					πF.SetLineno(795)
					πTemp004 = make([]πg.Param, 5)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "opt", Def: nil}
					πTemp004[2] = πg.Param{Name: "value", Def: nil}
					πTemp004[3] = πg.Param{Name: "values", Def: nil}
					πTemp004[4] = πg.Param{Name: "parser", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("process", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt *πg.Object = πArgs[1]; _ = µopt
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
						var µvalues *πg.Object = πArgs[3]; _ = µvalues
						var µparser *πg.Object = πArgs[4]; _ = µparser
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
							// line 799: value = self.convert_value(opt, value)
							πF.SetLineno(799)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp001[0] = µopt
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[1] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßconvert_value, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µvalue = πTemp003
							// line 804: return self.take_action(
							πF.SetLineno(804)
							πTemp001 = πF.MakeArgs(6)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdest, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp001[2] = µopt
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[3] = µvalue
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp001[4] = µvalues
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							πTemp001[5] = µparser
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtake_action, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßprocess.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 807: def take_action(self, action, dest, opt, value, values, parser):
					πF.SetLineno(807)
					πTemp004 = make([]πg.Param, 7)
					πTemp004[0] = πg.Param{Name: "self", Def: nil}
					πTemp004[1] = πg.Param{Name: "action", Def: nil}
					πTemp004[2] = πg.Param{Name: "dest", Def: nil}
					πTemp004[3] = πg.Param{Name: "opt", Def: nil}
					πTemp004[4] = πg.Param{Name: "value", Def: nil}
					πTemp004[5] = πg.Param{Name: "values", Def: nil}
					πTemp004[6] = πg.Param{Name: "parser", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("take_action", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µaction *πg.Object = πArgs[1]; _ = µaction
						var µdest *πg.Object = πArgs[2]; _ = µdest
						var µopt *πg.Object = πArgs[3]; _ = µopt
						var µvalue *πg.Object = πArgs[4]; _ = µvalue
						var µvalues *πg.Object = πArgs[5]; _ = µvalues
						var µparser *πg.Object = πArgs[6]; _ = µparser
						var µargs *πg.Object = πg.UnboundLocal; _ = µargs
						var µkwargs *πg.Object = πg.UnboundLocal; _ = µkwargs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßstore.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßstore_const.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßstore_true.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßstore_false.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßappend.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßappend_const.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßcount.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßcallback.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßhelp.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µaction, "action"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µaction, ßversion.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label10
							}
							goto Label11
							// line 808: if action == "store":
							πF.SetLineno(808)
						Label1:
							// line 809: setattr(values, dest, value)
							πF.SetLineno(809)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp003[0] = µvalues
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp003[1] = µdest
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[2] = µvalue
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label12
							// line 810: elif action == "store_const":
							πF.SetLineno(810)
						Label2:
							// line 811: setattr(values, dest, self.const)
							πF.SetLineno(811)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp003[0] = µvalues
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp003[1] = µdest
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßconst, nil); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label12
							// line 812: elif action == "store_true":
							πF.SetLineno(812)
						Label3:
							// line 813: setattr(values, dest, True)
							πF.SetLineno(813)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp003[0] = µvalues
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp003[1] = µdest
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label12
							// line 814: elif action == "store_false":
							πF.SetLineno(814)
						Label4:
							// line 815: setattr(values, dest, False)
							πF.SetLineno(815)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp003[0] = µvalues
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp003[1] = µdest
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label12
							// line 816: elif action == "append":
							πF.SetLineno(816)
						Label5:
							// line 817: values.ensure_value(dest, []).append(value)
							πF.SetLineno(817)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[0] = µvalue
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp005[0] = µdest
							πTemp006 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µvalues, ßensure_value, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label12
							// line 818: elif action == "append_const":
							πF.SetLineno(818)
						Label6:
							// line 819: values.ensure_value(dest, []).append(self.const)
							πF.SetLineno(819)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßconst, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp005[0] = µdest
							πTemp006 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp006...).ToObject()
							πTemp005[1] = πTemp001
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µvalues, ßensure_value, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label12
							// line 820: elif action == "count":
							πF.SetLineno(820)
						Label7:
							// line 821: setattr(values, dest, values.ensure_value(dest, 0) + 1)
							πF.SetLineno(821)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp003[0] = µvalues
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp003[1] = µdest
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp005[0] = µdest
							πTemp005[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µvalues, ßensure_value, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Add(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label12
							// line 822: elif action == "callback":
							πF.SetLineno(822)
						Label8:
							// line 823: args = self.callback_args or ()
							πF.SetLineno(823)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcallback_args, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label13
							}
							πTemp004 = πg.NewTuple0().ToObject()
							πTemp001 = πTemp004
						Label13:
							µargs = πTemp001
							// line 824: kwargs = self.callback_kwargs or {}
							πF.SetLineno(824)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßcallback_kwargs, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label14
							}
							πTemp008 = πg.NewDict()
							πTemp004 = πTemp008.ToObject()
							πTemp001 = πTemp004
						Label14:
							µkwargs = πTemp001
							// line 825: self.callback(self, opt, value, parser, *args, **kwargs)
							πF.SetLineno(825)
							πTemp003 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp003[1] = µopt
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[2] = µvalue
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							πTemp003[3] = µparser
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßcallback, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Invoke(πF, πTemp001, πTemp003, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label12
							// line 826: elif action == "help":
							πF.SetLineno(826)
						Label9:
							// line 827: parser.print_help()
							πF.SetLineno(827)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ßprint_help, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 828: parser.exit()
							πF.SetLineno(828)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ßexit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label12
							// line 829: elif action == "version":
							πF.SetLineno(829)
						Label10:
							// line 830: parser.print_version()
							πF.SetLineno(830)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ßprint_version, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 831: parser.exit()
							πF.SetLineno(831)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ßexit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label12
						Label11:
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßaction, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("unknown action %r").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 833: raise ValueError("unknown action %r" % self.action)
							πF.SetLineno(833)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label12
						Label12:
							// line 835: return 1
							πF.SetLineno(835)
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
					if πE = πClass.SetItem(πF, ßtake_action.ToObject(), πTemp021); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp013, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp013 == nil {
				πTemp013 = πg.TypeType.ToObject()
			}
			if πTemp014, πE = πTemp013.Call(πF, []*πg.Object{πg.NewStr("Option").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOption.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 840: SUPPRESS_HELP = "SUPPRESS"+"HELP"
			πF.SetLineno(840)
			if πTemp012, πE = πg.Add(πF, ßSUPPRESS.ToObject(), ßHELP.ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSUPPRESS_HELP.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 841: SUPPRESS_USAGE = "SUPPRESS"+"USAGE"
			πF.SetLineno(841)
			if πTemp012, πE = πg.Add(πF, ßSUPPRESS.ToObject(), ßUSAGE.ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSUPPRESS_USAGE.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 843: try:
			πF.SetLineno(843)
			πF.PushCheckpoint(2)
			// line 844: basestring
			πF.SetLineno(844)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			// line 849: def isbasestring(x):
			πF.SetLineno(849)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("isbasestring", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 850: return isinstance(x, basestring)
					πF.SetLineno(850)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
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
			if πE = πF.Globals().SetItem(πF, ßisbasestring.ToObject(), πTemp012); πE != nil {
				continue
			}
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp015, πTemp016 = πF.ExcInfo()
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp017, πE = πg.IsInstance(πF, πTemp015.ToObject(), πTemp013); πE != nil {
				continue
			}
			if πTemp017 {
				goto Label3
			}
			πE = πF.Raise(πTemp015.ToObject(), nil, πTemp016.ToObject())
			continue
			// line 845: except NameError:
			πF.SetLineno(845)
		Label3:
			// line 846: def isbasestring(x):
			πF.SetLineno(846)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "x", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("isbasestring", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
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
					// line 847: return isinstance(x, (types.StringType, types.UnicodeType))
					πF.SetLineno(847)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßStringType, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßUnicodeType, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßisbasestring.ToObject(), πTemp013); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 852: class Values(object):
			πF.SetLineno(852)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp019
			πTemp007 = πg.NewDict()
			if πTemp014, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp014); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Values", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
					// line 854: def __init__(self, defaults=None):
					πF.SetLineno(854)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "defaults", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdefaults *πg.Object = πArgs[1]; _ = µdefaults
						var µattr *πg.Object = πg.UnboundLocal; _ = µattr
						var µval *πg.Object = πg.UnboundLocal; _ = µval
						var πTemp001 bool
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µdefaults); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 855: if defaults:
							πF.SetLineno(855)
						Label1:
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µdefaults, ßitems, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
									continue
								}
								µattr = πTemp004
								µval = πTemp006
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 857: setattr(self, attr, val)
							πF.SetLineno(857)
							πTemp007 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp007[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp007[1] = µattr
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							πTemp007[2] = µval
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 859: def __str__(self):
					πF.SetLineno(859)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__str__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 860: return str(self.__dict__)
							πF.SetLineno(860)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__str__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 862: __repr__ = _repr
					πF.SetLineno(862)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_repr); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 864: def __cmp__(self, other):
					πF.SetLineno(864)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "other", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__cmp__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µother *πg.Object = πArgs[1]; _ = µother
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
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValues); πE != nil {
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
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[0] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßDictType, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
							goto Label3
							// line 865: if isinstance(other, Values):
							πF.SetLineno(865)
						Label1:
							// line 866: return cmp(self.__dict__, other.__dict__)
							πF.SetLineno(866)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µother, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label4
							// line 867: elif isinstance(other, types.DictType):
							πF.SetLineno(867)
						Label2:
							// line 868: return cmp(self.__dict__, other)
							πF.SetLineno(868)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µother, "other"); πE != nil {
								continue
							}
							πTemp001[1] = µother
							if πTemp002, πE = πg.ResolveGlobal(πF, ßcmp); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							goto Label4
						Label3:
							// line 870: return -1
							πF.SetLineno(870)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πR = πTemp002
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
					if πE = πClass.SetItem(πF, ß__cmp__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 872: def _update_careful(self, dict):
					πF.SetLineno(872)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "dict", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_update_careful", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdict *πg.Object = πArgs[1]; _ = µdict
						var µattr *πg.Object = πg.UnboundLocal; _ = µattr
						var µdval *πg.Object = πg.UnboundLocal; _ = µdval
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
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 873: """
							πF.SetLineno(873)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
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
								µattr = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µdict, µattr); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 880: if attr in dict:
							πF.SetLineno(880)
						Label4:
							// line 881: dval = dict[attr]
							πF.SetLineno(881)
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp003 = µattr
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µdict, πTemp003); πE != nil {
								continue
							}
							µdval = πTemp004
							if πE = πg.CheckLocal(πF, µdval, "dval"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µdval != πTemp004).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 882: if dval is not None:
							πF.SetLineno(882)
						Label6:
							// line 883: setattr(self, attr, dval)
							πF.SetLineno(883)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp002[1] = µattr
							if πE = πg.CheckLocal(πF, µdval, "dval"); πE != nil {
								continue
							}
							πTemp002[2] = µdval
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label7
						Label7:
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
					if πE = πClass.SetItem(πF, ß_update_careful.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 873: """
					πF.SetLineno(873)
					// line 873: """
					πF.SetLineno(873)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n        Update the option values from an arbitrary dictionary, but only\n        use keys from dict that already have a corresponding attribute\n        in self.  Any keys in dict without a corresponding attribute\n        are silently ignored.\n        ").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        Update the option values from an arbitrary dictionary, but only\n        use keys from dict that already have a corresponding attribute\n        in self.  Any keys in dict without a corresponding attribute\n        are silently ignored.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ß_update_careful); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 885: def _update_loose(self, dict):
					πF.SetLineno(885)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "dict", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_update_loose", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdict *πg.Object = πArgs[1]; _ = µdict
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
							// line 886: """
							πF.SetLineno(886)
							// line 891: self.__dict__.update(dict)
							πF.SetLineno(891)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							πTemp001[0] = µdict
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_update_loose.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 886: """
					πF.SetLineno(886)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("\n        Update the option values from an arbitrary dictionary,\n        using all keys from the dictionary regardless of whether\n        they have a corresponding attribute in self or not.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ß_update_loose); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 893: def _update(self, dict, mode):
					πF.SetLineno(893)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "dict", Def: nil}
					πTemp002[2] = πg.Param{Name: "mode", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_update", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdict *πg.Object = πArgs[1]; _ = µdict
						var µmode *πg.Object = πArgs[2]; _ = µmode
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
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µmode, ßcareful.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µmode, ßloose.ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 894: if mode == "careful":
							πF.SetLineno(894)
						Label1:
							// line 895: self._update_careful(dict)
							πF.SetLineno(895)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							πTemp003[0] = µdict
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_update_careful, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label4
							// line 896: elif mode == "loose":
							πF.SetLineno(896)
						Label2:
							// line 897: self._update_loose(dict)
							πF.SetLineno(897)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							πTemp003[0] = µdict
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_update_loose, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label4
						Label3:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("invalid update mode: %r").ToObject(), µmode); πE != nil {
								continue
							}
							// line 899: raise ValueError, "invalid update mode: %r" % mode
							πF.SetLineno(899)
							πE = πF.Raise(πTemp001, πTemp004, nil)
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
					if πE = πClass.SetItem(πF, ß_update.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 901: def read_module(self, modname, mode="careful"):
					πF.SetLineno(901)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "modname", Def: nil}
					πTemp002[2] = πg.Param{Name: "mode", Def: ßcareful.ToObject()}
					πTemp008 = πg.NewFunction(πg.NewCode("read_module", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmodname *πg.Object = πArgs[1]; _ = µmodname
						var µmode *πg.Object = πArgs[2]; _ = µmode
						var µmod *πg.Object = πg.UnboundLocal; _ = µmod
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 902: __import__(modname)
							πF.SetLineno(902)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmodname, "modname"); πE != nil {
								continue
							}
							πTemp001[0] = µmodname
							if πTemp002, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 903: mod = sys.modules[modname]
							πF.SetLineno(903)
							if πE = πg.CheckLocal(πF, µmodname, "modname"); πE != nil {
								continue
							}
							πTemp002 = µmodname
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							µmod = πTemp003
							// line 904: self._update(vars(mod), mode)
							πF.SetLineno(904)
							πTemp001 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							πTemp006[0] = µmod
							if πTemp002, πE = πg.ResolveGlobal(πF, ßvars); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							πTemp001[1] = µmode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_update, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßread_module.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 906: def read_file(self, filename, mode="careful"):
					πF.SetLineno(906)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "filename", Def: nil}
					πTemp002[2] = πg.Param{Name: "mode", Def: ßcareful.ToObject()}
					πTemp009 = πg.NewFunction(πg.NewCode("read_file", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfilename *πg.Object = πArgs[1]; _ = µfilename
						var µmode *πg.Object = πArgs[2]; _ = µmode
						var µvars *πg.Object = πg.UnboundLocal; _ = µvars
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
							// line 907: vars = {}
							πF.SetLineno(907)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µvars = πTemp002
							// line 908: execfile(filename, vars)
							πF.SetLineno(908)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp003[0] = µfilename
							if πE = πg.CheckLocal(πF, µvars, "vars"); πE != nil {
								continue
							}
							πTemp003[1] = µvars
							if πTemp002, πE = πg.ResolveGlobal(πF, ßexecfile); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 909: self._update(vars, mode)
							πF.SetLineno(909)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvars, "vars"); πE != nil {
								continue
							}
							πTemp003[0] = µvars
							if πE = πg.CheckLocal(πF, µmode, "mode"); πE != nil {
								continue
							}
							πTemp003[1] = µmode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_update, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßread_file.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 911: def ensure_value(self, attr, value):
					πF.SetLineno(911)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "attr", Def: nil}
					πTemp002[2] = πg.Param{Name: "value", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("ensure_value", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µattr *πg.Object = πArgs[1]; _ = µattr
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp004[1] = µattr
							if πTemp005, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
							if πTemp002 {
								goto Label1
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp004[1] = µattr
							if πTemp005, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp006 == πTemp005).ToObject()
							πTemp001 = πTemp003
						Label1:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							goto Label3
							// line 912: if not hasattr(self, attr) or getattr(self, attr) is None:
							πF.SetLineno(912)
						Label2:
							// line 913: setattr(self, attr, value)
							πF.SetLineno(913)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp004[1] = µattr
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[2] = µvalue
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label3
						Label3:
							// line 914: return getattr(self, attr)
							πF.SetLineno(914)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µattr, "attr"); πE != nil {
								continue
							}
							πTemp004[1] = µattr
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßensure_value.ToObject(), πTemp010); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp018, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp018 == nil {
				πTemp018 = πg.TypeType.ToObject()
			}
			if πTemp019, πE = πTemp018.Call(πF, []*πg.Object{πg.NewStr("Values").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßValues.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 917: class OptionContainer(object):
			πF.SetLineno(917)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp019
			πTemp007 = πg.NewDict()
			if πTemp014, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp014); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OptionContainer", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 919: """
					πF.SetLineno(919)
					// line 919: """
					πF.SetLineno(919)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Abstract base class.\n\n    Class attributes:\n      standard_option_list : [Option]\n        list of standard options that will be accepted by all instances\n        of this parser class (intended to be overridden by subclasses).\n\n    Instance attributes:\n      option_list : [Option]\n        the list of Option objects contained by this OptionContainer\n      _short_opt : { string : Option }\n        dictionary mapping short option strings, eg. \"-f\" or \"-X\",\n        to the Option instances that implement them.  If an Option\n        has multiple short option strings, it will appear in this\n        dictionary multiple times. [1]\n      _long_opt : { string : Option }\n        dictionary mapping long option strings, eg. \"--file\" or\n        \"--exclude\", to the Option instances that implement them.\n        Again, a given Option can occur multiple times in this\n        dictionary. [1]\n      defaults : { string : any }\n        dictionary mapping option destination names to default\n        values for each destination [1]\n\n    [1] These mappings are common to (shared by) all components of the\n        controlling OptionParser, where they are initially created.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 949: def __init__(self, option_class, conflict_handler, description):
					πF.SetLineno(949)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "option_class", Def: nil}
					πTemp002[2] = πg.Param{Name: "conflict_handler", Def: nil}
					πTemp002[3] = πg.Param{Name: "description", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoption_class *πg.Object = πArgs[1]; _ = µoption_class
						var µconflict_handler *πg.Object = πArgs[2]; _ = µconflict_handler
						var µdescription *πg.Object = πArgs[3]; _ = µdescription
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
							// line 954: self._create_option_list()
							πF.SetLineno(954)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_create_option_list, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 956: self.option_class = option_class
							πF.SetLineno(956)
							if πE = πg.CheckLocal(πF, µoption_class, "option_class"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µoption_class); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoption_class, πTemp001); πE != nil {
								continue
							}
							// line 957: self.set_conflict_handler(conflict_handler)
							πF.SetLineno(957)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µconflict_handler, "conflict_handler"); πE != nil {
								continue
							}
							πTemp003[0] = µconflict_handler
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_conflict_handler, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 958: self.set_description(description)
							πF.SetLineno(958)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							πTemp003[0] = µdescription
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßset_description, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 960: def _create_option_mappings(self):
					πF.SetLineno(960)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_create_option_mappings", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 964: self._short_opt = {}            # single letter -> Option instance
							πF.SetLineno(964)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_short_opt, πTemp003); πE != nil {
								continue
							}
							// line 965: self._long_opt = {}             # long option -> Option instance
							πF.SetLineno(965)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_long_opt, πTemp003); πE != nil {
								continue
							}
							// line 966: self.defaults = {}              # maps option dest -> default value
							πF.SetLineno(966)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdefaults, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_create_option_mappings.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 969: def _share_option_mappings(self, parser):
					πF.SetLineno(969)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "parser", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_share_option_mappings", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µparser *πg.Object = πArgs[1]; _ = µparser
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
							// line 972: self._short_opt = parser._short_opt
							πF.SetLineno(972)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ß_short_opt, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_short_opt, πTemp002); πE != nil {
								continue
							}
							// line 973: self._long_opt = parser._long_opt
							πF.SetLineno(973)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ß_long_opt, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_long_opt, πTemp002); πE != nil {
								continue
							}
							// line 974: self.defaults = parser.defaults
							πF.SetLineno(974)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ßdefaults, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdefaults, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_share_option_mappings.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 976: def set_conflict_handler(self, handler):
					πF.SetLineno(976)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "handler", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("set_conflict_handler", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µhandler *πg.Object = πArgs[1]; _ = µhandler
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
							if πE = πg.CheckLocal(πF, µhandler, "handler"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßerror.ToObject(), ßresolve.ToObject()).ToObject()
							if πTemp003, πE = πg.Contains(πF, πTemp002, µhandler); πE != nil {
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
							// line 977: if handler not in ("error", "resolve"):
							πF.SetLineno(977)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhandler, "handler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("invalid conflict_resolution value %r").ToObject(), µhandler); πE != nil {
								continue
							}
							// line 978: raise ValueError, "invalid conflict_resolution value %r" % handler
							πF.SetLineno(978)
							πE = πF.Raise(πTemp001, πTemp002, nil)
							continue
							goto Label2
						Label2:
							// line 979: self.conflict_handler = handler
							πF.SetLineno(979)
							if πE = πg.CheckLocal(πF, µhandler, "handler"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µhandler); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßconflict_handler, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_conflict_handler.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 981: def set_description(self, description):
					πF.SetLineno(981)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "description", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("set_description", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdescription *πg.Object = πArgs[1]; _ = µdescription
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 982: self.description = description
							πF.SetLineno(982)
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdescription); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdescription, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_description.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 984: def get_description(self):
					πF.SetLineno(984)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("get_description", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 985: return self.description
							πF.SetLineno(985)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdescription, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_description.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 988: def destroy(self):
					πF.SetLineno(988)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("destroy", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 989: """see OptionParser.destroy()."""
							πF.SetLineno(989)
							// line 990: del self._short_opt
							πF.SetLineno(990)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ß_short_opt); πE != nil {
								continue
							}
							// line 991: del self._long_opt
							πF.SetLineno(991)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ß_long_opt); πE != nil {
								continue
							}
							// line 992: del self.defaults
							πF.SetLineno(992)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ßdefaults); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdestroy.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 989: """see OptionParser.destroy()."""
					πF.SetLineno(989)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("see OptionParser.destroy().").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßdestroy); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 997: def _check_conflict(self, option):
					πF.SetLineno(997)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "option", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("_check_conflict", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoption *πg.Object = πArgs[1]; _ = µoption
						var µconflict_opts *πg.Object = πg.UnboundLocal; _ = µconflict_opts
						var µopt *πg.Object = πg.UnboundLocal; _ = µopt
						var µhandler *πg.Object = πg.UnboundLocal; _ = µhandler
						var µc_option *πg.Object = πg.UnboundLocal; _ = µc_option
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
						var πTemp008 *πg.Object
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
							case 1: goto Label1
							case 2: goto Label2
							case 6: goto Label6
							case 7: goto Label7
							case 16: goto Label16
							case 17: goto Label17
							default: panic("unexpected function state")
							}
							// line 998: conflict_opts = []
							πF.SetLineno(998)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µconflict_opts = πTemp002
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ß_short_opts, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µopt = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp006, µopt); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 1000: if opt in self._short_opt:
							πF.SetLineno(1000)
						Label4:
							// line 1001: conflict_opts.append((opt, self._short_opt[opt]))
							πF.SetLineno(1001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp006 = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µopt, πTemp007).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µconflict_opts, "conflict_opts"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µconflict_opts, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								πTemp005 = !isStop
							} else {
								πTemp005 = true
								µopt = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(6)            
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp006, µopt); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							goto Label10
							// line 1003: if opt in self._long_opt:
							πF.SetLineno(1003)
						Label9:
							// line 1004: conflict_opts.append((opt, self._long_opt[opt]))
							πF.SetLineno(1004)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp006 = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µopt, πTemp007).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µconflict_opts, "conflict_opts"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µconflict_opts, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label10
						Label10:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							if πE = πg.CheckLocal(πF, µconflict_opts, "conflict_opts"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µconflict_opts); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label11
							}
							goto Label12
							// line 1006: if conflict_opts:
							πF.SetLineno(1006)
						Label11:
							// line 1007: handler = self.conflict_handler
							πF.SetLineno(1007)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßconflict_handler, nil); πE != nil {
								continue
							}
							µhandler = πTemp002
							if πE = πg.CheckLocal(πF, µhandler, "handler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µhandler, ßerror.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							if πE = πg.CheckLocal(πF, µhandler, "handler"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µhandler, ßresolve.ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 1008: if handler == "error":
							πF.SetLineno(1008)
						Label13:
							πTemp001 = πF.MakeArgs(2)
							πTemp009 = πF.MakeArgs(1)
							πTemp010 = make([]πg.Param, 0)
							πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/optparse.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µco *πg.Object = πg.UnboundLocal; _ = µco
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
										if πE = πg.CheckLocal(πF, µconflict_opts, "conflict_opts"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µconflict_opts); πE != nil {
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
											µco = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 1011: % ", ".join([co[0] for co in conflict_opts]),
										πF.SetLineno(1011)
										πTemp004 = πg.NewInt(0).ToObject()
										if πE = πg.CheckLocal(πF, µco, "co"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetItem(πF, µco, πTemp004); πE != nil {
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
							if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp007}, nil); πE != nil {
								continue
							}
							πTemp009[0] = πTemp003
							if πTemp003, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("conflicting option string(s): %s").ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp001[1] = µoption
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOptionConflictError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1009: raise OptionConflictError(
							πF.SetLineno(1009)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label15
							// line 1013: elif handler == "resolve":
							πF.SetLineno(1013)
						Label14:
							if πE = πg.CheckLocal(πF, µconflict_opts, "conflict_opts"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µconflict_opts); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}}}, πTemp003); πE != nil {
									continue
								}
								µopt = πTemp007
								µc_option = πTemp008
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(16)            
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("--").ToObject()
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µopt, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label19
							}
							goto Label20
							// line 1015: if opt.startswith("--"):
							πF.SetLineno(1015)
						Label19:
							// line 1016: c_option._long_opts.remove(opt)
							πF.SetLineno(1016)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp001[0] = µopt
							if πE = πg.CheckLocal(πF, µc_option, "c_option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µc_option, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1017: del self._long_opt[opt]
							πF.SetLineno(1017)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp007 = µopt
							if πE = πg.DelItem(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							goto Label21
						Label20:
							// line 1019: c_option._short_opts.remove(opt)
							πF.SetLineno(1019)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp001[0] = µopt
							if πE = πg.CheckLocal(πF, µc_option, "c_option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µc_option, ß_short_opts, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1020: del self._short_opt[opt]
							πF.SetLineno(1020)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp007 = µopt
							if πE = πg.DelItem(πF, πTemp003, πTemp007); πE != nil {
								continue
							}
							goto Label21
						Label21:
							if πE = πg.CheckLocal(πF, µc_option, "c_option"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µc_option, ß_short_opts, nil); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label22
							}
							if πE = πg.CheckLocal(πF, µc_option, "c_option"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µc_option, ß_long_opts, nil); πE != nil {
								continue
							}
							πTemp007 = πTemp008
						Label22:
							if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label23
							}
							goto Label24
							// line 1021: if not (c_option._short_opts or c_option._long_opts):
							πF.SetLineno(1021)
						Label23:
							// line 1022: c_option.container.option_list.remove(c_option)
							πF.SetLineno(1022)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc_option, "c_option"); πE != nil {
								continue
							}
							πTemp001[0] = µc_option
							if πE = πg.CheckLocal(πF, µc_option, "c_option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µc_option, ßcontainer, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp007, ßremove, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label24
						Label24:
							continue
						Label17:
							if πE != nil || πR != nil {
								continue
							}
						Label18:
							goto Label15
						Label15:
							goto Label12
						Label12:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_check_conflict.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1024: def add_option(self, *args, **kwargs):
					πF.SetLineno(1024)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("add_option", "/usr/lib/python2.7/optparse.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
						var µoption *πg.Object = πg.UnboundLocal; _ = µoption
						var µopt *πg.Object = πg.UnboundLocal; _ = µopt
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8: goto Label8
							case 9: goto Label9
							case 11: goto Label11
							case 12: goto Label12
							default: panic("unexpected function state")
							}
							// line 1025: """add_option(Option)
							πF.SetLineno(1025)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßStringTypes, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp005, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Eq(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µkwargs); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
						Label2:
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 1028: if type(args[0]) in types.StringTypes:
							πF.SetLineno(1028)
						Label1:
							// line 1029: option = self.option_class(*args, **kwargs)
							πF.SetLineno(1029)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoption_class, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwargs); πE != nil {
								continue
							}
							µoption = πTemp003
							goto Label5
							// line 1030: elif len(args) == 1 and not kwargs:
							πF.SetLineno(1030)
						Label3:
							// line 1031: option = args[0]
							πF.SetLineno(1031)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µoption = πTemp003
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp002[0] = µoption
							if πTemp003, πE = πg.ResolveGlobal(πF, ßOption); πE != nil {
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
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 1032: if not isinstance(option, Option):
							πF.SetLineno(1032)
						Label6:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("not an Option instance: %r").ToObject(), µoption); πE != nil {
								continue
							}
							// line 1033: raise TypeError, "not an Option instance: %r" % option
							πF.SetLineno(1033)
							πE = πF.Raise(πTemp001, πTemp003, nil)
							continue
							goto Label7
						Label7:
							goto Label5
						Label4:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							// line 1035: raise TypeError, "invalid arguments"
							πF.SetLineno(1035)
							πE = πF.Raise(πTemp001, πg.NewStr("invalid arguments").ToObject(), nil)
							continue
							goto Label5
						Label5:
							// line 1037: self._check_conflict(option)
							πF.SetLineno(1037)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp002[0] = µoption
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_check_conflict, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1039: self.option_list.append(option)
							πF.SetLineno(1039)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp002[0] = µoption
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1040: option.container = self
							πF.SetLineno(1040)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µself); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µoption, ßcontainer, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ß_short_opts, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
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
								µopt = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(8)            
							// line 1042: self._short_opt[opt] = option
							πF.SetLineno(1042)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µoption); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp005 = µopt
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(12)
							πTemp006 = false
						Label11:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp006 {
								πF.PopCheckpoint()
								goto Label13
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
								µopt = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(11)            
							// line 1044: self._long_opt[opt] = option
							πF.SetLineno(1044)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, µoption); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp005 = µopt
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 != πTemp004).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label14
							}
							goto Label15
							// line 1046: if option.dest is not None:     # option has a dest, we need a default
							πF.SetLineno(1046)
						Label14:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ßdefault, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNO_DEFAULT); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 != πTemp004).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label16
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, πTemp004, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label17
							}
							goto Label18
							// line 1047: if option.default is not NO_DEFAULT:
							πF.SetLineno(1047)
						Label16:
							// line 1048: self.defaults[option.dest] = option.default
							πF.SetLineno(1048)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ßdefault, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp008
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							goto Label18
							// line 1049: elif option.dest not in self.defaults:
							πF.SetLineno(1049)
						Label17:
							// line 1050: self.defaults[option.dest] = None
							πF.SetLineno(1050)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							πTemp005 = πTemp008
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
								continue
							}
							goto Label18
						Label18:
							goto Label15
						Label15:
							// line 1052: return option
							πF.SetLineno(1052)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πR = µoption
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßadd_option.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1025: """add_option(Option)
					πF.SetLineno(1025)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("add_option(Option)\n           add_option(opt_str, ..., kwarg=val, ...)\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßadd_option); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
						continue
					}
					// line 1054: def add_options(self, option_list):
					πF.SetLineno(1054)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "option_list", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("add_options", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoption_list *πg.Object = πArgs[1]; _ = µoption_list
						var µoption *πg.Object = πg.UnboundLocal; _ = µoption
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
							if πE = πg.CheckLocal(πF, µoption_list, "option_list"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µoption_list); πE != nil {
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
								µoption = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1056: self.add_option(option)
							πF.SetLineno(1056)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp005[0] = µoption
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßadd_option, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßadd_options.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1060: def get_option(self, opt_str):
					πF.SetLineno(1060)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "opt_str", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("get_option", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt_str *πg.Object = πArgs[1]; _ = µopt_str
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
							// line 1061: return (self._short_opt.get(opt_str) or
							πF.SetLineno(1061)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							πTemp003[0] = µopt_str
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							πTemp003[0] = µopt_str
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßget_option.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1064: def has_option(self, opt_str):
					πF.SetLineno(1064)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "opt_str", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("has_option", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt_str *πg.Object = πArgs[1]; _ = µopt_str
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
							// line 1065: return (opt_str in self._short_opt or
							πF.SetLineno(1065)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp004, µopt_str); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005).ToObject()
							πTemp001 = πTemp003
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πTemp004, µopt_str); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp005).ToObject()
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
					if πE = πClass.SetItem(πF, ßhas_option.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 1068: def remove_option(self, opt_str):
					πF.SetLineno(1068)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "opt_str", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("remove_option", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt_str *πg.Object = πArgs[1]; _ = µopt_str
						var µoption *πg.Object = πg.UnboundLocal; _ = µoption
						var µopt *πg.Object = πg.UnboundLocal; _ = µopt
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
							case 8: goto Label8
							case 9: goto Label9
							case 5: goto Label5
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							// line 1069: option = self._short_opt.get(opt_str)
							πF.SetLineno(1069)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							πTemp001[0] = µopt_str
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µoption = πTemp002
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µoption == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1070: if option is None:
							πF.SetLineno(1070)
						Label1:
							// line 1071: option = self._long_opt.get(opt_str)
							πF.SetLineno(1071)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							πTemp001[0] = µopt_str
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µoption = πTemp002
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µoption == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 1072: if option is None:
							πF.SetLineno(1072)
						Label3:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("no such option %r").ToObject(), µopt_str); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1073: raise ValueError("no such option %r" % opt_str)
							πF.SetLineno(1073)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ß_short_opts, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µopt = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(5)            
							// line 1076: del self._short_opt[opt]
							πF.SetLineno(1076)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp006 = µopt
							if πE = πg.DelItem(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µoption, ß_long_opts, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
								continue
							}
							πF.PushCheckpoint(9)
							πTemp004 = false
						Label8:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label10
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
								µopt = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(8)            
							// line 1078: del self._long_opt[opt]
							πF.SetLineno(1078)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp006 = µopt
							if πE = πg.DelItem(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
							// line 1079: option.container.option_list.remove(option)
							πF.SetLineno(1079)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp001[0] = µoption
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßcontainer, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßremove, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßremove_option.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 1084: def format_option_help(self, formatter):
					πF.SetLineno(1084)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "formatter", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("format_option_help", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformatter *πg.Object = πArgs[1]; _ = µformatter
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µoption *πg.Object = πg.UnboundLocal; _ = µoption
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 bool
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
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoption_list, nil); πE != nil {
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
							// line 1085: if not self.option_list:
							πF.SetLineno(1085)
						Label1:
							// line 1086: return ""
							πF.SetLineno(1086)
							πR = ß.ToObject()
							continue
							goto Label2
						Label2:
							// line 1087: result = []
							πF.SetLineno(1087)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µresult = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
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
								πTemp005 = !isStop
							} else {
								πTemp005 = true
								µoption = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)            
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µoption, ßhelp, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ßSUPPRESS_HELP); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(πTemp007 == πTemp008).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 1089: if not option.help is SUPPRESS_HELP:
							πF.SetLineno(1089)
						Label6:
							// line 1090: result.append(formatter.format_option(option))
							πF.SetLineno(1090)
							πTemp004 = πF.MakeArgs(1)
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp009[0] = µoption
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µformatter, ßformat_option, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							πTemp004[0] = πTemp006
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label7
						Label7:
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 1091: return "".join(result)
							πF.SetLineno(1091)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp004[0] = µresult
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßformat_option_help.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 1093: def format_description(self, formatter):
					πF.SetLineno(1093)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "formatter", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("format_description", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformatter *πg.Object = πArgs[1]; _ = µformatter
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
							// line 1094: return formatter.format_description(self.get_description())
							πF.SetLineno(1094)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_description, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µformatter, ßformat_description, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat_description.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 1096: def format_help(self, formatter):
					πF.SetLineno(1096)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "formatter", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("format_help", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformatter *πg.Object = πArgs[1]; _ = µformatter
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var πTemp001 []*πg.Object
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
							// line 1097: result = []
							πF.SetLineno(1097)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µresult = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdescription, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1098: if self.description:
							πF.SetLineno(1098)
						Label1:
							// line 1099: result.append(self.format_description(formatter))
							πF.SetLineno(1099)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							πTemp004[0] = µformatter
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßformat_description, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1100: if self.option_list:
							πF.SetLineno(1100)
						Label3:
							// line 1101: result.append(self.format_option_help(formatter))
							πF.SetLineno(1101)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							πTemp004[0] = µformatter
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßformat_option_help, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label4
						Label4:
							// line 1102: return "\n".join(result)
							πF.SetLineno(1102)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("\n").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat_help.ToObject(), πTemp017); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp018, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp018 == nil {
				πTemp018 = πg.TypeType.ToObject()
			}
			if πTemp019, πE = πTemp018.Call(πF, []*πg.Object{πg.NewStr("OptionContainer").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOptionContainer.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 1105: class OptionGroup (OptionContainer):
			πF.SetLineno(1105)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßOptionContainer); πE != nil {
				continue
			}
			πTemp001[0] = πTemp019
			πTemp007 = πg.NewDict()
			if πTemp014, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp014); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OptionGroup", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
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
					// line 1107: def __init__(self, parser, title, description=None):
					πF.SetLineno(1107)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "parser", Def: nil}
					πTemp002[2] = πg.Param{Name: "title", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "description", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µparser *πg.Object = πArgs[1]; _ = µparser
						var µtitle *πg.Object = πArgs[2]; _ = µtitle
						var µdescription *πg.Object = πArgs[3]; _ = µdescription
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
							// line 1108: self.parser = parser
							πF.SetLineno(1108)
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µparser); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparser, πTemp001); πE != nil {
								continue
							}
							// line 1109: OptionContainer.__init__(
							πF.SetLineno(1109)
							πTemp002 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ßoption_class, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp001
							if πE = πg.CheckLocal(πF, µparser, "parser"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µparser, ßconflict_handler, nil); πE != nil {
								continue
							}
							πTemp002[2] = πTemp001
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							πTemp002[3] = µdescription
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionContainer); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1111: self.title = title
							πF.SetLineno(1111)
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtitle); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtitle, πTemp001); πE != nil {
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
					// line 1113: def _create_option_list(self):
					πF.SetLineno(1113)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_create_option_list", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1114: self.option_list = []
							πF.SetLineno(1114)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoption_list, πTemp003); πE != nil {
								continue
							}
							// line 1115: self._share_option_mappings(self.parser)
							πF.SetLineno(1115)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßparser, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_share_option_mappings, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_create_option_list.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 1117: def set_title(self, title):
					πF.SetLineno(1117)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "title", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("set_title", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µtitle *πg.Object = πArgs[1]; _ = µtitle
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1118: self.title = title
							πF.SetLineno(1118)
							if πE = πg.CheckLocal(πF, µtitle, "title"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µtitle); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßtitle, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_title.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1120: def destroy(self):
					πF.SetLineno(1120)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("destroy", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1121: """see OptionParser.destroy()."""
							πF.SetLineno(1121)
							// line 1122: OptionContainer.destroy(self)
							πF.SetLineno(1122)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOptionContainer); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdestroy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1123: del self.option_list
							πF.SetLineno(1123)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ßoption_list); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdestroy.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1121: """see OptionParser.destroy()."""
					πF.SetLineno(1121)
					// line 1121: """see OptionParser.destroy()."""
					πF.SetLineno(1121)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("see OptionParser.destroy().").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("see OptionParser.destroy().").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßdestroy); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 1127: def format_help(self, formatter):
					πF.SetLineno(1127)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "formatter", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("format_help", "/usr/lib/python2.7/optparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformatter *πg.Object = πArgs[1]; _ = µformatter
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
							// line 1128: result = formatter.format_heading(self.title)
							πF.SetLineno(1128)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßtitle, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µformatter, ßformat_heading, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp003
							// line 1129: formatter.indent()
							πF.SetLineno(1129)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µformatter, ßindent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1130: result += OptionContainer.format_help(self, formatter)
							πF.SetLineno(1130)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							πTemp001[1] = µformatter
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOptionContainer); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßformat_help, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.IAdd(πF, µresult, πTemp002); πE != nil {
								continue
							}
							µresult = πTemp003
							// line 1131: formatter.dedent()
							πF.SetLineno(1131)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µformatter, ßdedent, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1132: return result
							πF.SetLineno(1132)
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
					if πE = πClass.SetItem(πF, ßformat_help.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp018, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp018 == nil {
				πTemp018 = πg.TypeType.ToObject()
			}
			if πTemp019, πE = πTemp018.Call(πF, []*πg.Object{πg.NewStr("OptionGroup").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOptionGroup.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 1135: class OptionParser (OptionContainer):
			πF.SetLineno(1135)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp019, πE = πg.ResolveGlobal(πF, ßOptionContainer); πE != nil {
				continue
			}
			πTemp001[0] = πTemp019
			πTemp007 = πg.NewDict()
			if πTemp014, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp014); πE != nil {
				continue
			}
			_, πE = πg.NewCode("OptionParser", "/usr/lib/python2.7/optparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 1137: """
					πF.SetLineno(1137)
					// line 1137: """
					πF.SetLineno(1137)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    Class attributes:\n      standard_option_list : [Option]\n        list of standard options that will be accepted by all instances\n        of this parser class (intended to be overridden by subclasses).\n\n    Instance attributes:\n      usage : string\n        a usage string for your program.  Before it is displayed\n        to the user, \"%prog\" will be expanded to the name of\n        your program (self.prog or os.path.basename(sys.argv[0])).\n      prog : string\n        the name of the current program (to override\n        os.path.basename(sys.argv[0])).\n      description : string\n        A paragraph of text giving a brief overview of your program.\n        optparse reformats this paragraph to fit the current terminal\n        width and prints it when the user requests help (after usage,\n        but before the list of options).\n      epilog : string\n        paragraph of help text to print after option help\n\n      option_groups : [OptionGroup]\n        list of option groups in this parser (option groups are\n        irrelevant for parsing the command-line, but very useful\n        for generating help)\n\n      allow_interspersed_args : bool = true\n        if true, positional arguments may be interspersed with options.\n        Assuming -a and -b each take a single argument, the command-line\n          -ablah foo bar -bboo baz\n        will be interpreted the same as\n          -ablah -bboo -- foo bar baz\n        If this flag were false, that command line would be interpreted as\n          -ablah -- foo bar -bboo baz\n        -- ie. we stop processing options as soon as we see the first\n        non-option argument.  (This is the tradition followed by\n        Python's getopt module, Perl's Getopt::Std, and other argument-\n        parsing libraries, but it is generally annoying to users.)\n\n      process_default_values : bool = true\n        if true, option default values are processed similarly to option\n        values from the command line: that is, they are passed to the\n        type-checking function for the option's type (as long as the\n        default value is a string).  (This really only matters if you\n        have defined custom types; see SF bug #955889.)  Set it to false\n        to restore the behaviour of Optik 1.4.1 and earlier.\n\n      rargs : [string]\n        the argument list currently being parsed.  Only set when\n        parse_args() is active, and continually trimmed down as\n        we consume arguments.  Mainly there for the benefit of\n        callback options.\n      largs : [string]\n        the list of leftover arguments that we have skipped while\n        parsing options.  If allow_interspersed_args is false, this\n        list is always empty.\n      values : Values\n        the set of option values currently being accumulated.  Only\n        set when parse_args() is active.  Also mainly for callbacks.\n\n    Because of the 'rargs', 'largs', and 'values' attributes,\n    OptionParser is not thread-safe.  If, for some perverse reason, you\n    need to parse command-line arguments simultaneously in different\n    threads, use different OptionParser instances.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 1205: standard_option_list = []
					πF.SetLineno(1205)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ßstandard_option_list.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 1207: def __init__(self,
					πF.SetLineno(1207)
					πTemp003 = make([]πg.Param, 11)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "usage", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πg.Param{Name: "option_list", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßOption); πE != nil {
						continue
					}
					πTemp003[3] = πg.Param{Name: "option_class", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[4] = πg.Param{Name: "version", Def: πTemp004}
					πTemp003[5] = πg.Param{Name: "conflict_handler", Def: ßerror.ToObject()}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[6] = πg.Param{Name: "description", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[7] = πg.Param{Name: "formatter", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp003[8] = πg.Param{Name: "add_help_option", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[9] = πg.Param{Name: "prog", Def: πTemp004}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[10] = πg.Param{Name: "epilog", Def: πTemp004}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µusage *πg.Object = πArgs[1]; _ = µusage
						var µoption_list *πg.Object = πArgs[2]; _ = µoption_list
						var µoption_class *πg.Object = πArgs[3]; _ = µoption_class
						var µversion *πg.Object = πArgs[4]; _ = µversion
						var µconflict_handler *πg.Object = πArgs[5]; _ = µconflict_handler
						var µdescription *πg.Object = πArgs[6]; _ = µdescription
						var µformatter *πg.Object = πArgs[7]; _ = µformatter
						var µadd_help_option *πg.Object = πArgs[8]; _ = µadd_help_option
						var µprog *πg.Object = πArgs[9]; _ = µprog
						var µepilog *πg.Object = πArgs[10]; _ = µepilog
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
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
							// line 1218: OptionContainer.__init__(
							πF.SetLineno(1218)
							πTemp001 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µoption_class, "option_class"); πE != nil {
								continue
							}
							πTemp001[1] = µoption_class
							if πE = πg.CheckLocal(πF, µconflict_handler, "conflict_handler"); πE != nil {
								continue
							}
							πTemp001[2] = µconflict_handler
							if πE = πg.CheckLocal(πF, µdescription, "description"); πE != nil {
								continue
							}
							πTemp001[3] = µdescription
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOptionContainer); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__init__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1220: self.set_usage(usage)
							πF.SetLineno(1220)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							πTemp001[0] = µusage
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßset_usage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1221: self.prog = prog
							πF.SetLineno(1221)
							if πE = πg.CheckLocal(πF, µprog, "prog"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µprog); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßprog, πTemp002); πE != nil {
								continue
							}
							// line 1222: self.version = version
							πF.SetLineno(1222)
							if πE = πg.CheckLocal(πF, µversion, "version"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µversion); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßversion, πTemp002); πE != nil {
								continue
							}
							// line 1223: self.allow_interspersed_args = True
							πF.SetLineno(1223)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßallow_interspersed_args, πTemp003); πE != nil {
								continue
							}
							// line 1224: self.process_default_values = True
							πF.SetLineno(1224)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßprocess_default_values, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µformatter == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1225: if formatter is None:
							πF.SetLineno(1225)
						Label1:
							// line 1226: formatter = IndentedHelpFormatter()
							πF.SetLineno(1226)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßIndentedHelpFormatter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µformatter = πTemp003
							goto Label2
						Label2:
							// line 1227: self.formatter = formatter
							πF.SetLineno(1227)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µformatter); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßformatter, πTemp002); πE != nil {
								continue
							}
							// line 1228: self.formatter.set_parser(self)
							πF.SetLineno(1228)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßformatter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßset_parser, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1229: self.epilog = epilog
							πF.SetLineno(1229)
							if πE = πg.CheckLocal(πF, µepilog, "epilog"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µepilog); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßepilog, πTemp002); πE != nil {
								continue
							}
							// line 1235: self._populate_option_list(option_list,
							πF.SetLineno(1235)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption_list, "option_list"); πE != nil {
								continue
							}
							πTemp001[0] = µoption_list
							if πE = πg.CheckLocal(πF, µadd_help_option, "add_help_option"); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"add_help", µadd_help_option},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_populate_option_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1238: self._init_parsing_state()
							πF.SetLineno(1238)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_init_parsing_state, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 1241: def destroy(self):
					πF.SetLineno(1241)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("destroy", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 1242: """
							πF.SetLineno(1242)
							// line 1248: OptionContainer.destroy(self)
							πF.SetLineno(1248)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßOptionContainer); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdestroy, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoption_groups, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
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
								µgroup = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1250: group.destroy()
							πF.SetLineno(1250)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µgroup, ßdestroy, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1251: del self.option_list
							πF.SetLineno(1251)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ßoption_list); πE != nil {
								continue
							}
							// line 1252: del self.option_groups
							πF.SetLineno(1252)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ßoption_groups); πE != nil {
								continue
							}
							// line 1253: del self.formatter
							πF.SetLineno(1253)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.DelAttr(πF, µself, ßformatter); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdestroy.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 1242: """
					πF.SetLineno(1242)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Declare that you are done with this OptionParser.  This cleans up\n        reference cycles so the OptionParser (and all objects referenced by\n        it) can be garbage-collected promptly.  After calling destroy(), the\n        OptionParser is unusable.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdestroy); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 1259: def _create_option_list(self):
					πF.SetLineno(1259)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_create_option_list", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1260: self.option_list = []
							πF.SetLineno(1260)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoption_list, πTemp003); πE != nil {
								continue
							}
							// line 1261: self.option_groups = []
							πF.SetLineno(1261)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßoption_groups, πTemp003); πE != nil {
								continue
							}
							// line 1262: self._create_option_mappings()
							πF.SetLineno(1262)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_create_option_mappings, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_create_option_list.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 1264: def _add_help_option(self):
					πF.SetLineno(1264)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_add_help_option", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 1265: self.add_option("-h", "--help",
							πF.SetLineno(1265)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewStr("-h").ToObject()
							πTemp001[1] = πg.NewStr("--help").ToObject()
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("show this help message and exit").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp005 = πg.KWArgs{
								{"action", ßhelp.ToObject()},
								{"help", πTemp004},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßadd_option, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_add_help_option.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 1269: def _add_version_option(self):
					πF.SetLineno(1269)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_add_version_option", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
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
							// line 1270: self.add_option("--version",
							πF.SetLineno(1270)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("--version").ToObject()
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("show program's version number and exit").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp005 = πg.KWArgs{
								{"action", ßversion.ToObject()},
								{"help", πTemp004},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßadd_option, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp005); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_add_version_option.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 1274: def _populate_option_list(self, option_list, add_help=True):
					πF.SetLineno(1274)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "option_list", Def: nil}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp003[2] = πg.Param{Name: "add_help", Def: πTemp009}
					πTemp008 = πg.NewFunction(πg.NewCode("_populate_option_list", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoption_list *πg.Object = πArgs[1]; _ = µoption_list
						var µadd_help *πg.Object = πArgs[2]; _ = µadd_help
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstandard_option_list, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 1275: if self.standard_option_list:
							πF.SetLineno(1275)
						Label1:
							// line 1276: self.add_options(self.standard_option_list)
							πF.SetLineno(1276)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstandard_option_list, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_options, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µoption_list, "option_list"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µoption_list); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 1277: if option_list:
							πF.SetLineno(1277)
						Label3:
							// line 1278: self.add_options(option_list)
							πF.SetLineno(1278)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption_list, "option_list"); πE != nil {
								continue
							}
							πTemp003[0] = µoption_list
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßadd_options, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßversion, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 1279: if self.version:
							πF.SetLineno(1279)
						Label5:
							// line 1280: self._add_version_option()
							πF.SetLineno(1280)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_add_version_option, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µadd_help, "add_help"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µadd_help); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 1281: if add_help:
							πF.SetLineno(1281)
						Label7:
							// line 1282: self._add_help_option()
							πF.SetLineno(1282)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_add_help_option, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ß_populate_option_list.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 1284: def _init_parsing_state(self):
					πF.SetLineno(1284)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("_init_parsing_state", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1286: self.rargs = None
							πF.SetLineno(1286)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrargs, πTemp002); πE != nil {
								continue
							}
							// line 1287: self.largs = None
							πF.SetLineno(1287)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlargs, πTemp002); πE != nil {
								continue
							}
							// line 1288: self.values = None
							πF.SetLineno(1288)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßvalues, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_init_parsing_state.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 1293: def set_usage(self, usage):
					πF.SetLineno(1293)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "usage", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("set_usage", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µusage *πg.Object = πArgs[1]; _ = µusage
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
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µusage == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSUPPRESS_USAGE); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µusage == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("usage: ").ToObject()
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µusage, ßlower, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßstartswith, nil); πE != nil {
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
							// line 1294: if usage is None:
							πF.SetLineno(1294)
						Label1:
							// line 1295: self.usage = _("%prog [options]")
							πF.SetLineno(1295)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("%prog [options]").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßusage, πTemp001); πE != nil {
								continue
							}
							goto Label5
							// line 1296: elif usage is SUPPRESS_USAGE:
							πF.SetLineno(1296)
						Label2:
							// line 1297: self.usage = None
							πF.SetLineno(1297)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßusage, πTemp002); πE != nil {
								continue
							}
							goto Label5
							// line 1299: elif usage.lower().startswith("usage: "):
							πF.SetLineno(1299)
						Label3:
							// line 1300: self.usage = usage[7:]
							πF.SetLineno(1300)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(7).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µusage, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßusage, πTemp001); πE != nil {
								continue
							}
							goto Label5
						Label4:
							// line 1302: self.usage = usage
							πF.SetLineno(1302)
							if πE = πg.CheckLocal(πF, µusage, "usage"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µusage); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßusage, πTemp001); πE != nil {
								continue
							}
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
					if πE = πClass.SetItem(πF, ßset_usage.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 1304: def enable_interspersed_args(self):
					πF.SetLineno(1304)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("enable_interspersed_args", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1305: """Set parsing to not stop on the first non-option, allowing
							πF.SetLineno(1305)
							// line 1310: self.allow_interspersed_args = True
							πF.SetLineno(1310)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßallow_interspersed_args, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßenable_interspersed_args.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 1305: """Set parsing to not stop on the first non-option, allowing
					πF.SetLineno(1305)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Set parsing to not stop on the first non-option, allowing\n        interspersing switches with command arguments. This is the\n        default behavior. See also disable_interspersed_args() and the\n        class documentation description of the attribute\n        allow_interspersed_args.").ToObject()); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßenable_interspersed_args); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
						continue
					}
					// line 1312: def disable_interspersed_args(self):
					πF.SetLineno(1312)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("disable_interspersed_args", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1313: """Set parsing to stop on the first non-option. Use this if
							πF.SetLineno(1313)
							// line 1318: self.allow_interspersed_args = False
							πF.SetLineno(1318)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßallow_interspersed_args, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdisable_interspersed_args.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 1313: """Set parsing to stop on the first non-option. Use this if
					πF.SetLineno(1313)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Set parsing to stop on the first non-option. Use this if\n        you have a command processor which runs another command that\n        has options of its own and you want to make sure these options\n        don't get confused.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßdisable_interspersed_args); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
						continue
					}
					// line 1320: def set_process_default_values(self, process):
					πF.SetLineno(1320)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "process", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("set_process_default_values", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µprocess *πg.Object = πArgs[1]; _ = µprocess
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1321: self.process_default_values = process
							πF.SetLineno(1321)
							if πE = πg.CheckLocal(πF, µprocess, "process"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µprocess); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßprocess_default_values, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_process_default_values.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 1323: def set_default(self, dest, value):
					πF.SetLineno(1323)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "dest", Def: nil}
					πTemp003[2] = πg.Param{Name: "value", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("set_default", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdest *πg.Object = πArgs[1]; _ = µdest
						var µvalue *πg.Object = πArgs[2]; _ = µvalue
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
							// line 1324: self.defaults[dest] = value
							πF.SetLineno(1324)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdest, "dest"); πE != nil {
								continue
							}
							πTemp003 = µdest
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
					if πE = πClass.SetItem(πF, ßset_default.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 1326: def set_defaults(self, **kwargs):
					πF.SetLineno(1326)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("set_defaults", "/usr/lib/python2.7/optparse.py", πTemp003, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 1327: self.defaults.update(kwargs)
							πF.SetLineno(1327)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp001[0] = µkwargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßupdate, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßset_defaults.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 1329: def _get_all_options(self):
					πF.SetLineno(1329)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("_get_all_options", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µoptions *πg.Object = πg.UnboundLocal; _ = µoptions
						var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 1330: options = self.option_list[:]
							πF.SetLineno(1330)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µoptions = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoption_groups, nil); πE != nil {
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
								µgroup = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1333: options += (group.option_list)
							πF.SetLineno(1333)
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µgroup, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µoptions, πTemp002); πE != nil {
								continue
							}
							µoptions = πTemp003
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1334: return options
							πF.SetLineno(1334)
							if πE = πg.CheckLocal(πF, µoptions, "options"); πE != nil {
								continue
							}
							πR = µoptions
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_get_all_options.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 1336: def get_default_values(self):
					πF.SetLineno(1336)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("get_default_values", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdefaults *πg.Object = πg.UnboundLocal; _ = µdefaults
						var µoption *πg.Object = πg.UnboundLocal; _ = µoption
						var µdefault *πg.Object = πg.UnboundLocal; _ = µdefault
						var µopt_str *πg.Object = πg.UnboundLocal; _ = µopt_str
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
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßprocess_default_values, nil); πE != nil {
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
							// line 1337: if not self.process_default_values:
							πF.SetLineno(1337)
						Label1:
							// line 1339: return Values(self.defaults)
							πF.SetLineno(1339)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValues); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πR = πTemp002
							continue
							goto Label2
						Label2:
							// line 1342: defaults = dict(self.defaults)
							πF.SetLineno(1342)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdefaults, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µdefaults = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_get_all_options, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
								πTemp006 = !isStop
							} else {
								πTemp006 = true
								µoption = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 1344: default = defaults.get(option.dest)
							πF.SetLineno(1344)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdefaults, ßget, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µdefault = πTemp005
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp004[0] = µdefault
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisbasestring); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp006, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 1345: if isbasestring(default):
							πF.SetLineno(1345)
						Label6:
							// line 1346: opt_str = option.get_opt_string()
							πF.SetLineno(1346)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßget_opt_string, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µopt_str = πTemp005
							// line 1347: defaults[option.dest] = option.check_value(opt_str, default)
							πF.SetLineno(1347)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							πTemp004[0] = µopt_str
							if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
								continue
							}
							πTemp004[1] = µdefault
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßcheck_value, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µoption, ßdest, nil); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.SetItem(πF, µdefaults, πTemp007, πTemp002); πE != nil {
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
							// line 1349: return Values(defaults)
							πF.SetLineno(1349)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdefaults, "defaults"); πE != nil {
								continue
							}
							πTemp004[0] = µdefaults
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValues); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßget_default_values.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 1354: def add_option_group(self, *args, **kwargs):
					πF.SetLineno(1354)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("add_option_group", "/usr/lib/python2.7/optparse.py", πTemp003, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
						var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
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
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µargs, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßStringType, nil); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004 == πTemp005).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label1
							}
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Eq(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µkwargs); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp007).ToObject()
							πTemp001 = πTemp003
						Label2:
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							goto Label4
							// line 1356: if type(args[0]) is types.StringType:
							πF.SetLineno(1356)
						Label1:
							// line 1357: group = OptionGroup(self, *args, **kwargs)
							πF.SetLineno(1357)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionGroup); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp001, πTemp002, µargs, nil, µkwargs); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µgroup = πTemp003
							goto Label5
							// line 1358: elif len(args) == 1 and not kwargs:
							πF.SetLineno(1358)
						Label3:
							// line 1359: group = args[0]
							πF.SetLineno(1359)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µgroup = πTemp003
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp002[0] = µgroup
							if πTemp003, πE = πg.ResolveGlobal(πF, ßOptionGroup); πE != nil {
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
							if πTemp006, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							goto Label7
							// line 1360: if not isinstance(group, OptionGroup):
							πF.SetLineno(1360)
						Label6:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("not an OptionGroup instance: %r").ToObject(), µgroup); πE != nil {
								continue
							}
							// line 1361: raise TypeError, "not an OptionGroup instance: %r" % group
							πF.SetLineno(1361)
							πE = πF.Raise(πTemp001, πTemp003, nil)
							continue
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µgroup, ßparser, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003 != µself).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label8
							}
							goto Label9
							// line 1362: if group.parser is not self:
							πF.SetLineno(1362)
						Label8:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 1363: raise ValueError, "invalid OptionGroup (wrong parser)"
							πF.SetLineno(1363)
							πE = πF.Raise(πTemp001, πg.NewStr("invalid OptionGroup (wrong parser)").ToObject(), nil)
							continue
							goto Label9
						Label9:
							goto Label5
						Label4:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							// line 1365: raise TypeError, "invalid arguments"
							πF.SetLineno(1365)
							πE = πF.Raise(πTemp001, πg.NewStr("invalid arguments").ToObject(), nil)
							continue
							goto Label5
						Label5:
							// line 1367: self.option_groups.append(group)
							πF.SetLineno(1367)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πTemp002[0] = µgroup
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoption_groups, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 1368: return group
							πF.SetLineno(1368)
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							πR = µgroup
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßadd_option_group.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 1370: def get_option_group(self, opt_str):
					πF.SetLineno(1370)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "opt_str", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("get_option_group", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt_str *πg.Object = πArgs[1]; _ = µopt_str
						var µoption *πg.Object = πg.UnboundLocal; _ = µoption
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
							// line 1371: option = (self._short_opt.get(opt_str) or
							πF.SetLineno(1371)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							πTemp003[0] = µopt_str
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt_str, "opt_str"); πE != nil {
								continue
							}
							πTemp003[0] = µopt_str
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001 = πTemp004
						Label1:
							µoption = πTemp001
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							πTemp001 = µoption
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µoption, ßcontainer, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(πTemp005 != µself).ToObject()
							πTemp001 = πTemp004
						Label2:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 1373: if option and option.container is not self:
							πF.SetLineno(1373)
						Label3:
							// line 1374: return option.container
							πF.SetLineno(1374)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ßcontainer, nil); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							// line 1375: return None
							πF.SetLineno(1375)
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
					if πE = πClass.SetItem(πF, ßget_option_group.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 1380: def _get_args(self, args):
					πF.SetLineno(1380)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "args", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("_get_args", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
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
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µargs == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1381: if args is None:
							πF.SetLineno(1381)
						Label1:
							// line 1382: return sys.argv[1:]
							πF.SetLineno(1382)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßargv, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label3
						Label2:
							// line 1384: return args[:]              # don't modify caller's list
							πF.SetLineno(1384)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_get_args.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 1386: def parse_args(self, args=None, values=None):
					πF.SetLineno(1386)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "args", Def: πTemp022}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πg.Param{Name: "values", Def: πTemp022}
					πTemp021 = πg.NewFunction(πg.NewCode("parse_args", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µargs *πg.Object = πArgs[1]; _ = µargs
						var µvalues *πg.Object = πArgs[2]; _ = µvalues
						var µrargs *πg.Object = πg.UnboundLocal; _ = µrargs
						var µlargs *πg.Object = πg.UnboundLocal; _ = µlargs
						var µstop *πg.Object = πg.UnboundLocal; _ = µstop
						var µerr *πg.Object = πg.UnboundLocal; _ = µerr
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							// line 1387: """
							πF.SetLineno(1387)
							// line 1400: rargs = self._get_args(args)
							πF.SetLineno(1400)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_get_args, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrargs = πTemp003
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µvalues == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1401: if values is None:
							πF.SetLineno(1401)
						Label1:
							// line 1402: values = self.get_default_values()
							πF.SetLineno(1402)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_default_values, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvalues = πTemp003
							goto Label2
						Label2:
							// line 1413: self.rargs = rargs
							πF.SetLineno(1413)
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µrargs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrargs, πTemp002); πE != nil {
								continue
							}
							// line 1414: self.largs = largs = []
							πF.SetLineno(1414)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßlargs, πTemp003); πE != nil {
								continue
							}
							µlargs = πTemp002
							// line 1415: self.values = values
							πF.SetLineno(1415)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalues); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßvalues, πTemp002); πE != nil {
								continue
							}
							// line 1417: try:
							πF.SetLineno(1417)
							πF.PushCheckpoint(4)
							// line 1418: stop = self._process_args(largs, rargs, values)
							πF.SetLineno(1418)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µlargs, "largs"); πE != nil {
								continue
							}
							πTemp001[0] = µlargs
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							πTemp001[1] = µrargs
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp001[2] = µvalues
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_process_args, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstop = πTemp003
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBadOptionError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.ResolveGlobal(πF, ßOptionValueError); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp003, πTemp007).ToObject()
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 1419: except (BadOptionError, OptionValueError), err:
							πF.SetLineno(1419)
						Label5:
							µerr = πTemp005.ToObject()
							// line 1420: self.error(str(err))
							πF.SetLineno(1420)
							πTemp001 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp008[0] = µerr
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 1422: args = largs + rargs
							πF.SetLineno(1422)
							if πE = πg.CheckLocal(πF, µlargs, "largs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µlargs, µrargs); πE != nil {
								continue
							}
							µargs = πTemp002
							// line 1423: return self.check_values(values, args)
							πF.SetLineno(1423)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp001[0] = µvalues
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[1] = µargs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßcheck_values, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßparse_args.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 1387: """
					πF.SetLineno(1387)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("\n        parse_args(args : [string] = sys.argv[1:],\n                   values : Values = None)\n        -> (values : Values, args : [string])\n\n        Parse the command-line options found in 'args' (default:\n        sys.argv[1:]).  Any errors result in a call to 'error()', which\n        by default prints the usage message to stderr and calls\n        sys.exit() with an error message.  On success returns a pair\n        (values, args) where 'values' is a Values instance (with all\n        your option values) and 'args' is the list of arguments left\n        over after parsing options.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßparse_args); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
						continue
					}
					// line 1425: def check_values(self, values, args):
					πF.SetLineno(1425)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "values", Def: nil}
					πTemp003[2] = πg.Param{Name: "args", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("check_values", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µvalues *πg.Object = πArgs[1]; _ = µvalues
						var µargs *πg.Object = πArgs[2]; _ = µargs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1426: """
							πF.SetLineno(1426)
							// line 1436: return (values, args)
							πF.SetLineno(1436)
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µvalues, µargs).ToObject()
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
					if πE = πClass.SetItem(πF, ßcheck_values.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 1426: """
					πF.SetLineno(1426)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πg.NewStr("\n        check_values(values : Values, args : [string])\n        -> (values : Values, args : [string])\n\n        Check that the supplied option values and leftover arguments are\n        valid.  Returns the option values and leftover arguments\n        (possibly adjusted, possibly completely new -- whatever you\n        like).  Default implementation just returns the passed-in\n        values; subclasses may override as desired.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßcheck_values); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp024, ß__doc__, πTemp023); πE != nil {
						continue
					}
					// line 1438: def _process_args(self, largs, rargs, values):
					πF.SetLineno(1438)
					πTemp003 = make([]πg.Param, 4)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "largs", Def: nil}
					πTemp003[2] = πg.Param{Name: "rargs", Def: nil}
					πTemp003[3] = πg.Param{Name: "values", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("_process_args", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µlargs *πg.Object = πArgs[1]; _ = µlargs
						var µrargs *πg.Object = πArgs[2]; _ = µrargs
						var µvalues *πg.Object = πArgs[3]; _ = µvalues
						var µarg *πg.Object = πg.UnboundLocal; _ = µarg
						var πTemp001 bool
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
							// line 1439: """_process_args(largs : [string],
							πF.SetLineno(1439)
							// line 1448: while rargs:
							πF.SetLineno(1448)
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
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, µrargs); πE != nil {
								continue
							}
							if πE != nil || !πTemp002 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1449: arg = rargs[0]
							πF.SetLineno(1449)
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µrargs, πTemp003); πE != nil {
								continue
							}
							µarg = πTemp004
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µarg, πg.NewStr("--").ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µarg, πTemp004); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp005, πg.NewStr("--").ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µarg, πTemp005); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, πTemp006, πg.NewStr("-").ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label6
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							πTemp007[0] = µarg
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp004, πE = πg.GT(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label6:
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßallow_interspersed_args, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label8
							}
							goto Label9
							// line 1453: if arg == "--":
							πF.SetLineno(1453)
						Label4:
							// line 1454: del rargs[0]
							πF.SetLineno(1454)
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.DelItem(πF, µrargs, πTemp003); πE != nil {
								continue
							}
							// line 1455: return
							πF.SetLineno(1455)
							πR = πg.None
							continue
							goto Label10
							// line 1456: elif arg[0:2] == "--":
							πF.SetLineno(1456)
						Label5:
							// line 1458: self._process_long_opt(rargs, values)
							πF.SetLineno(1458)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							πTemp007[0] = µrargs
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp007[1] = µvalues
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_process_long_opt, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label10
							// line 1459: elif arg[:1] == "-" and len(arg) > 1:
							πF.SetLineno(1459)
						Label7:
							// line 1462: self._process_short_opts(rargs, values)
							πF.SetLineno(1462)
							πTemp007 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							πTemp007[0] = µrargs
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp007[1] = µvalues
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_process_short_opts, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label10
							// line 1463: elif self.allow_interspersed_args:
							πF.SetLineno(1463)
						Label8:
							// line 1464: largs.append(arg)
							πF.SetLineno(1464)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							πTemp007[0] = µarg
							if πE = πg.CheckLocal(πF, µlargs, "largs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µlargs, ßappend, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 1465: del rargs[0]
							πF.SetLineno(1465)
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							πTemp003 = πg.NewInt(0).ToObject()
							if πE = πg.DelItem(πF, µrargs, πTemp003); πE != nil {
								continue
							}
							goto Label10
						Label9:
							// line 1467: return                  # stop now, leave this arg in rargs
							πF.SetLineno(1467)
							πR = πg.None
							continue
							goto Label10
						Label10:
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
					if πE = πClass.SetItem(πF, ß_process_args.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 1439: """_process_args(largs : [string],
					πF.SetLineno(1439)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πg.NewStr("_process_args(largs : [string],\n                         rargs : [string],\n                         values : Values)\n\n        Process command-line arguments and populate 'values', consuming\n        options and arguments from 'rargs'.  If 'allow_interspersed_args' is\n        false, stop at the first non-option argument.  If true, accumulate any\n        interspersed non-option arguments in 'largs'.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ß_process_args); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp025, ß__doc__, πTemp024); πE != nil {
						continue
					}
					// line 1489: def _match_long_opt(self, opt):
					πF.SetLineno(1489)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "opt", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("_match_long_opt", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µopt *πg.Object = πArgs[1]; _ = µopt
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
							// line 1490: """_match_long_opt(opt : string) -> string
							πF.SetLineno(1490)
							// line 1496: return _match_abbrev(opt, self._long_opt)
							πF.SetLineno(1496)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp001[0] = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_match_abbrev); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_match_long_opt.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 1490: """_match_long_opt(opt : string) -> string
					πF.SetLineno(1490)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πg.NewStr("_match_long_opt(opt : string) -> string\n\n        Determine which long option string 'opt' matches, ie. which one\n        it is an unambiguous abbreviation for.  Raises BadOptionError if\n        'opt' doesn't unambiguously match any long option string.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ß_match_long_opt); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp026, ß__doc__, πTemp025); πE != nil {
						continue
					}
					// line 1498: def _process_long_opt(self, rargs, values):
					πF.SetLineno(1498)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "rargs", Def: nil}
					πTemp003[2] = πg.Param{Name: "values", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("_process_long_opt", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrargs *πg.Object = πArgs[1]; _ = µrargs
						var µvalues *πg.Object = πArgs[2]; _ = µvalues
						var µarg *πg.Object = πg.UnboundLocal; _ = µarg
						var µopt *πg.Object = πg.UnboundLocal; _ = µopt
						var µnext_arg *πg.Object = πg.UnboundLocal; _ = µnext_arg
						var µhad_explicit_value *πg.Object = πg.UnboundLocal; _ = µhad_explicit_value
						var µoption *πg.Object = πg.UnboundLocal; _ = µoption
						var µnargs *πg.Object = πg.UnboundLocal; _ = µnargs
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1500: arg = rargs[0]
							πF.SetLineno(1500)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrargs, πTemp001); πE != nil {
								continue
							}
							µarg = πTemp002
							// line 1501: rargs = rargs[1:]
							πF.SetLineno(1501)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrargs, πTemp001); πE != nil {
								continue
							}
							µrargs = πTemp002
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µarg, πg.NewStr("=").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1505: if "=" in arg:
							πF.SetLineno(1505)
						Label1:
							// line 1506: (opt, next_arg) = arg.split("=", 1)
							πF.SetLineno(1506)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("=").ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µarg, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µopt = πTemp001
							µnext_arg = πTemp005
							// line 1507: rargs.insert(0, next_arg)
							πF.SetLineno(1507)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µnext_arg, "next_arg"); πE != nil {
								continue
							}
							πTemp004[1] = µnext_arg
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrargs, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1508: had_explicit_value = True
							πF.SetLineno(1508)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µhad_explicit_value = πTemp001
							goto Label3
						Label2:
							// line 1510: opt = arg
							πF.SetLineno(1510)
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							µopt = µarg
							// line 1511: had_explicit_value = False
							πF.SetLineno(1511)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µhad_explicit_value = πTemp001
							goto Label3
						Label3:
							// line 1513: opt = self._match_long_opt(opt)
							πF.SetLineno(1513)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp004[0] = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_match_long_opt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µopt = πTemp002
							// line 1514: option = self._long_opt[opt]
							πF.SetLineno(1514)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp001 = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_long_opt, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							µoption = πTemp002
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ßtakes_value, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µhad_explicit_value, "had_explicit_value"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µhad_explicit_value); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 1515: if option.takes_value():
							πF.SetLineno(1515)
						Label4:
							// line 1516: nargs = option.nargs
							πF.SetLineno(1516)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ßnargs, nil); πE != nil {
								continue
							}
							µnargs = πTemp001
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							πTemp004[0] = µrargs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, πTemp005, µnargs); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µnargs, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							goto Label10
							// line 1517: if len(rargs) < nargs:
							πF.SetLineno(1517)
						Label8:
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µnargs, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							goto Label13
							// line 1518: if nargs == 1:
							πF.SetLineno(1518)
						Label12:
							// line 1519: self.error(_("%s option requires an argument") % opt)
							πF.SetLineno(1519)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("%s option requires an argument").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp005, µopt); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label14
						Label13:
							// line 1521: self.error(_("%s option requires %d arguments")
							πF.SetLineno(1521)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("%s option requires %d arguments").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µopt, µnargs).ToObject()
							if πTemp001, πE = πg.Mod(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label14
						Label14:
							goto Label11
							// line 1523: elif nargs == 1:
							πF.SetLineno(1523)
						Label9:
							// line 1525: value = rargs[0]
							πF.SetLineno(1525)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrargs, πTemp001); πE != nil {
								continue
							}
							µvalue = πTemp002
							// line 1526: rargs = rargs[1:]
							πF.SetLineno(1526)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrargs, πTemp001); πE != nil {
								continue
							}
							µrargs = πTemp002
							goto Label11
						Label10:
							// line 1528: value = tuple(rargs[0:nargs])
							πF.SetLineno(1528)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), µnargs, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrargs, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µvalue = πTemp002
							// line 1529: del rargs[0:nargs]
							πF.SetLineno(1529)
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), µnargs, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µrargs, πTemp001); πE != nil {
								continue
							}
							goto Label11
						Label11:
							goto Label7
							// line 1531: elif had_explicit_value:
							πF.SetLineno(1531)
						Label5:
							// line 1532: self.error(_("%s option does not take a value") % opt)
							πF.SetLineno(1532)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("%s option does not take a value").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πTemp005, µopt); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label7
						Label6:
							// line 1535: value = None
							πF.SetLineno(1535)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µvalue = πTemp001
							goto Label7
						Label7:
							// line 1537: option.process(opt, value, values, self)
							πF.SetLineno(1537)
							πTemp004 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp004[0] = µopt
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp004[1] = µvalue
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp004[2] = µvalues
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[3] = µself
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µoption, ßprocess, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_process_long_opt.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 1539: def _process_short_opts(self, rargs, values):
					πF.SetLineno(1539)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "rargs", Def: nil}
					πTemp003[2] = πg.Param{Name: "values", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("_process_short_opts", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrargs *πg.Object = πArgs[1]; _ = µrargs
						var µvalues *πg.Object = πArgs[2]; _ = µvalues
						var µarg *πg.Object = πg.UnboundLocal; _ = µarg
						var µstop *πg.Object = πg.UnboundLocal; _ = µstop
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µch *πg.Object = πg.UnboundLocal; _ = µch
						var µopt *πg.Object = πg.UnboundLocal; _ = µopt
						var µoption *πg.Object = πg.UnboundLocal; _ = µoption
						var µnargs *πg.Object = πg.UnboundLocal; _ = µnargs
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
						var πTemp006 []*πg.Object
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
							// line 1541: arg = rargs[-1]
							πF.SetLineno(1541)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrargs, πTemp001); πE != nil {
								continue
							}
							µarg = πTemp002
							// line 1542: rargs = rargs[:-1]
							πF.SetLineno(1542)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrargs, πTemp001); πE != nil {
								continue
							}
							µrargs = πTemp002
							// line 1543: stop = False
							πF.SetLineno(1543)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µstop = πTemp001
							// line 1544: i = 1
							πF.SetLineno(1544)
							µi = πg.NewInt(1).ToObject()
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µarg, πTemp002); πE != nil {
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
								µch = πTemp002
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 1546: opt = "-" + ch
							πF.SetLineno(1546)
							if πE = πg.CheckLocal(πF, µch, "ch"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πg.NewStr("-").ToObject(), µch); πE != nil {
								continue
							}
							µopt = πTemp002
							// line 1547: option = self._short_opt.get(opt)
							πF.SetLineno(1547)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp006[0] = µopt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_short_opt, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µoption = πTemp002
							// line 1548: i += 1                      # we have consumed a character
							πF.SetLineno(1548)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µi = πTemp002
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µoption); πE != nil {
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
							// line 1550: if not option:
							πF.SetLineno(1550)
						Label4:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp006[0] = µopt
							if πTemp002, πE = πg.ResolveGlobal(πF, ßBadOptionError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 1551: raise BadOptionError(opt)
							πF.SetLineno(1551)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßtakes_value, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							goto Label7
							// line 1552: if option.takes_value():
							πF.SetLineno(1552)
						Label6:
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							πTemp006[0] = µarg
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.LT(πF, µi, πTemp007); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							goto Label10
							// line 1555: if i < len(arg):
							πF.SetLineno(1555)
						Label9:
							// line 1556: rargs.insert(0, arg[i:])
							πF.SetLineno(1556)
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µi, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µarg, πTemp002); πE != nil {
								continue
							}
							πTemp006[1] = πTemp003
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µrargs, ßinsert, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 1557: stop = True
							πF.SetLineno(1557)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µstop = πTemp002
							goto Label10
						Label10:
							// line 1559: nargs = option.nargs
							πF.SetLineno(1559)
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßnargs, nil); πE != nil {
								continue
							}
							µnargs = πTemp002
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							πTemp006[0] = µrargs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, πTemp007, µnargs); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label11
							}
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µnargs, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label12
							}
							goto Label13
							// line 1560: if len(rargs) < nargs:
							πF.SetLineno(1560)
						Label11:
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µnargs, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label15
							}
							goto Label16
							// line 1561: if nargs == 1:
							πF.SetLineno(1561)
						Label15:
							// line 1562: self.error(_("%s option requires an argument") % opt)
							πF.SetLineno(1562)
							πTemp006 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("%s option requires an argument").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πTemp007, µopt); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label17
						Label16:
							// line 1564: self.error(_("%s option requires %d arguments")
							πF.SetLineno(1564)
							πTemp006 = πF.MakeArgs(1)
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("%s option requires %d arguments").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(µopt, µnargs).ToObject()
							if πTemp002, πE = πg.Mod(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßerror, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label17
						Label17:
							goto Label14
							// line 1566: elif nargs == 1:
							πF.SetLineno(1566)
						Label12:
							// line 1568: value = rargs[0]
							πF.SetLineno(1568)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µrargs, πTemp002); πE != nil {
								continue
							}
							µvalue = πTemp003
							// line 1569: rargs = rargs[1:]
							πF.SetLineno(1569)
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µrargs, πTemp002); πE != nil {
								continue
							}
							µrargs = πTemp003
							goto Label14
						Label13:
							// line 1571: value = tuple(rargs[0:nargs])
							πF.SetLineno(1571)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), µnargs, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µrargs, πTemp002); πE != nil {
								continue
							}
							πTemp006[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µvalue = πTemp003
							// line 1572: del rargs[0:nargs]
							πF.SetLineno(1572)
							if πE = πg.CheckLocal(πF, µrargs, "rargs"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnargs, "nargs"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(0).ToObject(), µnargs, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µrargs, πTemp002); πE != nil {
								continue
							}
							goto Label14
						Label14:
							goto Label8
						Label7:
							// line 1575: value = None
							πF.SetLineno(1575)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µvalue = πTemp002
							goto Label8
						Label8:
							// line 1577: option.process(opt, value, values, self)
							πF.SetLineno(1577)
							πTemp006 = πF.MakeArgs(4)
							if πE = πg.CheckLocal(πF, µopt, "opt"); πE != nil {
								continue
							}
							πTemp006[0] = µopt
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp006[1] = µvalue
							if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
								continue
							}
							πTemp006[2] = µvalues
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp006[3] = µself
							if πE = πg.CheckLocal(πF, µoption, "option"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µoption, ßprocess, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µstop, "stop"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µstop); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label18
							}
							goto Label19
							// line 1579: if stop:
							πF.SetLineno(1579)
						Label18:
							// line 1580: break
							πF.SetLineno(1580)
							πTemp004 = true
							continue
							goto Label19
						Label19:
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
					if πE = πClass.SetItem(πF, ß_process_short_opts.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 1585: def get_prog_name(self):
					πF.SetLineno(1585)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("get_prog_name", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µpaths *πg.Object = πg.UnboundLocal; _ = µpaths
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
						var πTemp006 *πg.Object
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßprog, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1586: if self.prog is None:
							πF.SetLineno(1586)
						Label1:
							// line 1588: paths = sys.argv[0].split('/')
							πF.SetLineno(1588)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("/").ToObject()
							πTemp001 = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßargv, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpaths = πTemp002
							// line 1589: return paths[-1] if paths else ''
							πF.SetLineno(1589)
							if πE = πg.CheckLocal(πF, µpaths, "paths"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µpaths); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label4
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µpaths, "paths"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µpaths, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							goto Label5
						Label4:
							πTemp001 = ß.ToObject()
						Label5:
							πR = πTemp001
							continue
							goto Label3
						Label2:
							// line 1591: return self.prog
							πF.SetLineno(1591)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßprog, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_prog_name.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 1593: def expand_prog_name(self, s):
					πF.SetLineno(1593)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "s", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("expand_prog_name", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πArgs[1]; _ = µs
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
							// line 1595: return self.get_prog_name().join(s.split("%prog"))
							πF.SetLineno(1595)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("%prog").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µs, ßsplit, nil); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget_prog_name, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjoin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßexpand_prog_name.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 1597: def get_description(self):
					πF.SetLineno(1597)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("get_description", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1598: return self.expand_prog_name(self.description)
							πF.SetLineno(1598)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdescription, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßexpand_prog_name, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_description.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 1600: def exit(self, status=0, msg=None):
					πF.SetLineno(1600)
					πTemp003 = make([]πg.Param, 3)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "status", Def: πg.NewInt(0).ToObject()}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[2] = πg.Param{Name: "msg", Def: πTemp031}
					πTemp030 = πg.NewFunction(πg.NewCode("exit", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µstatus *πg.Object = πArgs[1]; _ = µstatus
						var µmsg *πg.Object = πArgs[2]; _ = µmsg
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
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µmsg); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 1601: if msg:
							πF.SetLineno(1601)
						Label1:
							// line 1602: sys.stderr.write(msg)
							πF.SetLineno(1602)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp002[0] = µmsg
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstderr, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label2
						Label2:
							// line 1603: sys.exit(status)
							πF.SetLineno(1603)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstatus, "status"); πE != nil {
								continue
							}
							πTemp002[0] = µstatus
							if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßexit, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßexit.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 1605: def error(self, msg):
					πF.SetLineno(1605)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "msg", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("error", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µmsg *πg.Object = πArgs[1]; _ = µmsg
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
							// line 1606: """error(msg : string)
							πF.SetLineno(1606)
							// line 1612: self.print_usage(sys.stderr)
							πF.SetLineno(1612)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstderr, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßprint_usage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1613: self.exit(2, "%s: error: %s\n" % (self.get_prog_name(), msg))
							πF.SetLineno(1613)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßget_prog_name, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp005, µmsg).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s: error: %s\n").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßexit, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßerror.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 1606: """error(msg : string)
					πF.SetLineno(1606)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp032}, πg.NewStr("error(msg : string)\n\n        Print a usage message incorporating 'msg' to stderr and exit.\n        If you override this in a subclass, it should not return -- it\n        should either exit or raise an exception.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp033, πE = πg.ResolveClass(πF, πClass, nil, ßerror); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp033, ß__doc__, πTemp032); πE != nil {
						continue
					}
					// line 1615: def get_usage(self):
					πF.SetLineno(1615)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("get_usage", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßusage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 1616: if self.usage:
							πF.SetLineno(1616)
						Label1:
							// line 1617: return self.formatter.format_usage(
							πF.SetLineno(1617)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßusage, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßexpand_prog_name, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßformatter, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßformat_usage, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πR = πTemp001
							continue
							goto Label3
						Label2:
							// line 1620: return ""
							πF.SetLineno(1620)
							πR = ß.ToObject()
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
					if πE = πClass.SetItem(πF, ßget_usage.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 1622: def print_usage(self, file=None):
					πF.SetLineno(1622)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "file", Def: πTemp034}
					πTemp033 = πg.NewFunction(πg.NewCode("print_usage", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfile *πg.Object = πArgs[1]; _ = µfile
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
							// line 1623: """print_usage(file : file = stdout)
							πF.SetLineno(1623)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßusage, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 1631: if self.usage:
							πF.SetLineno(1631)
						Label1:
							// line 1632: print >>file, self.get_usage()
							πF.SetLineno(1632)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_usage, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.Print(πF, πTemp003, true); πE != nil {
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
					if πE = πClass.SetItem(πF, ßprint_usage.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 1623: """print_usage(file : file = stdout)
					πF.SetLineno(1623)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp034}, πg.NewStr("print_usage(file : file = stdout)\n\n        Print the usage message for the current program (self.usage) to\n        'file' (default stdout).  Any occurrence of the string \"%prog\" in\n        self.usage is replaced with the name of the current program\n        (basename of sys.argv[0]).  Does nothing if self.usage is empty\n        or not defined.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßprint_usage); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp035, ß__doc__, πTemp034); πE != nil {
						continue
					}
					// line 1634: def get_version(self):
					πF.SetLineno(1634)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("get_version", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßversion, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 1635: if self.version:
							πF.SetLineno(1635)
						Label1:
							// line 1636: return self.expand_prog_name(self.version)
							πF.SetLineno(1636)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßversion, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßexpand_prog_name, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πR = πTemp004
							continue
							goto Label3
						Label2:
							// line 1638: return ""
							πF.SetLineno(1638)
							πR = ß.ToObject()
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
					if πE = πClass.SetItem(πF, ßget_version.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 1640: def print_version(self, file=None):
					πF.SetLineno(1640)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "file", Def: πTemp036}
					πTemp035 = πg.NewFunction(πg.NewCode("print_version", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfile *πg.Object = πArgs[1]; _ = µfile
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
							// line 1641: """print_version(file : file = stdout)
							πF.SetLineno(1641)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßversion, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 1648: if self.version:
							πF.SetLineno(1648)
						Label1:
							// line 1649: print >>file, self.get_version()
							πF.SetLineno(1649)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget_version, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp004
							if πE = πg.Print(πF, πTemp003, true); πE != nil {
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
					if πE = πClass.SetItem(πF, ßprint_version.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 1641: """print_version(file : file = stdout)
					πF.SetLineno(1641)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πg.NewStr("print_version(file : file = stdout)\n\n        Print the version message for this program (self.version) to\n        'file' (default stdout).  As with print_usage(), any occurrence\n        of \"%prog\" in self.version is replaced by the current program's\n        name.  Does nothing if self.version is empty or undefined.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßprint_version); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp037, ß__doc__, πTemp036); πE != nil {
						continue
					}
					// line 1651: def format_option_help(self, formatter=None):
					πF.SetLineno(1651)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "formatter", Def: πTemp037}
					πTemp036 = πg.NewFunction(πg.NewCode("format_option_help", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformatter *πg.Object = πArgs[1]; _ = µformatter
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
						var µgroup *πg.Object = πg.UnboundLocal; _ = µgroup
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 5: goto Label5
							case 6: goto Label6
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µformatter == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1652: if formatter is None:
							πF.SetLineno(1652)
						Label1:
							// line 1653: formatter = self.formatter
							πF.SetLineno(1653)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßformatter, nil); πE != nil {
								continue
							}
							µformatter = πTemp001
							goto Label2
						Label2:
							// line 1654: formatter.store_option_strings(self)
							πF.SetLineno(1654)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µformatter, ßstore_option_strings, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1655: result = []
							πF.SetLineno(1655)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µresult = πTemp001
							// line 1656: result.append(formatter.format_heading(_("Options")))
							πF.SetLineno(1656)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = ßOptions.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µformatter, ßformat_heading, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1657: formatter.indent()
							πF.SetLineno(1657)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µformatter, ßindent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßoption_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1658: if self.option_list:
							πF.SetLineno(1658)
						Label3:
							// line 1659: result.append(OptionContainer.format_option_help(self, formatter))
							πF.SetLineno(1659)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							πTemp005[1] = µformatter
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOptionContainer); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßformat_option_help, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1660: result.append("\n")
							πF.SetLineno(1660)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßoption_groups, nil); πE != nil {
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
								πTemp007 = !isStop
							} else {
								πTemp007 = true
								µgroup = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(5)            
							// line 1662: result.append(group.format_help(formatter))
							πF.SetLineno(1662)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							πTemp005[0] = µformatter
							if πE = πg.CheckLocal(πF, µgroup, "group"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µgroup, ßformat_help, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp008
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1663: result.append("\n")
							πF.SetLineno(1663)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							continue
						Label6:
							if πE != nil || πR != nil {
								continue
							}
						Label7:
							// line 1664: formatter.dedent()
							πF.SetLineno(1664)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µformatter, ßdedent, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 1666: return "".join(result[:-1])
							πF.SetLineno(1666)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µresult, πTemp001); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßformat_option_help.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 1668: def format_epilog(self, formatter):
					πF.SetLineno(1668)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "formatter", Def: nil}
					πTemp037 = πg.NewFunction(πg.NewCode("format_epilog", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformatter *πg.Object = πArgs[1]; _ = µformatter
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
							// line 1669: return formatter.format_epilog(self.epilog)
							πF.SetLineno(1669)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßepilog, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µformatter, ßformat_epilog, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßformat_epilog.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 1671: def format_help(self, formatter=None):
					πF.SetLineno(1671)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp039, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "formatter", Def: πTemp039}
					πTemp038 = πg.NewFunction(πg.NewCode("format_help", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µformatter *πg.Object = πArgs[1]; _ = µformatter
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µformatter == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1672: if formatter is None:
							πF.SetLineno(1672)
						Label1:
							// line 1673: formatter = self.formatter
							πF.SetLineno(1673)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßformatter, nil); πE != nil {
								continue
							}
							µformatter = πTemp001
							goto Label2
						Label2:
							// line 1674: result = []
							πF.SetLineno(1674)
							πTemp004 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp004...).ToObject()
							µresult = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßusage, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 1675: if self.usage:
							πF.SetLineno(1675)
						Label3:
							// line 1676: result.append(self.get_usage() + "\n")
							πF.SetLineno(1676)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_usage, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp005, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label4
						Label4:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdescription, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 1677: if self.description:
							πF.SetLineno(1677)
						Label5:
							// line 1678: result.append(self.format_description(formatter) + "\n")
							πF.SetLineno(1678)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							πTemp006[0] = µformatter
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßformat_description, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.Add(πF, πTemp005, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label6
						Label6:
							// line 1679: result.append(self.format_option_help(formatter))
							πF.SetLineno(1679)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							πTemp006[0] = µformatter
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßformat_option_help, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1680: result.append(self.format_epilog(formatter))
							πF.SetLineno(1680)
							πTemp004 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µformatter, "formatter"); πE != nil {
								continue
							}
							πTemp006[0] = µformatter
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßformat_epilog, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µresult, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 1681: return "".join(result)
							πF.SetLineno(1681)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp004[0] = µresult
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßformat_help.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 1684: def _get_encoding(self, file):
					πF.SetLineno(1684)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "file", Def: nil}
					πTemp039 = πg.NewFunction(πg.NewCode("_get_encoding", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfile *πg.Object = πArgs[1]; _ = µfile
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 1685: return "utf-8"
							πF.SetLineno(1685)
							πR = πg.NewStr("utf-8").ToObject()
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_get_encoding.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 1691: def print_help(self, file=None):
					πF.SetLineno(1691)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					if πTemp041, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp003[1] = πg.Param{Name: "file", Def: πTemp041}
					πTemp040 = πg.NewFunction(πg.NewCode("print_help", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfile *πg.Object = πArgs[1]; _ = µfile
						var µencoding *πg.Object = πg.UnboundLocal; _ = µencoding
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
							// line 1692: """print_help(file : file = stdout)
							πF.SetLineno(1692)
							if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µfile == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 1697: if file is None:
							πF.SetLineno(1697)
						Label1:
							// line 1698: file = sys.stdout
							πF.SetLineno(1698)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstdout, nil); πE != nil {
								continue
							}
							µfile = πTemp002
							goto Label2
						Label2:
							// line 1699: encoding = self._get_encoding(file)
							πF.SetLineno(1699)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
								continue
							}
							πTemp004[0] = µfile
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_get_encoding, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µencoding = πTemp002
							// line 1701: file.write(self.format_help())
							πF.SetLineno(1701)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßformat_help, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfile, ßwrite, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßprint_help.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 1692: """print_help(file : file = stdout)
					πF.SetLineno(1692)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp041}, πg.NewStr("print_help(file : file = stdout)\n\n        Print an extended help message, listing all options and any\n        help text provided with them, to 'file' (default stdout).\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp042, πE = πg.ResolveClass(πF, πClass, nil, ßprint_help); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp042, ß__doc__, πTemp041); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp018, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp018 == nil {
				πTemp018 = πg.TypeType.ToObject()
			}
			if πTemp019, πE = πTemp018.Call(πF, []*πg.Object{πg.NewStr("OptionParser").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßOptionParser.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 1706: def _match_abbrev(s, wordmap):
			πF.SetLineno(1706)
			πTemp003 = make([]πg.Param, 2)
			πTemp003[0] = πg.Param{Name: "s", Def: nil}
			πTemp003[1] = πg.Param{Name: "wordmap", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("_match_abbrev", "/usr/lib/python2.7/optparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µwordmap *πg.Object = πArgs[1]; _ = µwordmap
				var µpossibilities *πg.Object = πg.UnboundLocal; _ = µpossibilities
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 []πg.Param
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
					// line 1707: """_match_abbrev(s : string, wordmap : {string : Option}) -> string
					πF.SetLineno(1707)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µwordmap, "wordmap"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µwordmap, µs); πE != nil {
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
					// line 1714: if s in wordmap:
					πF.SetLineno(1714)
				Label1:
					// line 1715: return s
					πF.SetLineno(1715)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πR = µs
					continue
					goto Label3
				Label2:
					// line 1718: possibilities = [word for word in wordmap.keys()
					πF.SetLineno(1718)
					πTemp004 = make([]πg.Param, 0)
					πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/optparse.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µword *πg.Object = πg.UnboundLocal; _ = µword
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
								if πE = πg.CheckLocal(πF, µwordmap, "wordmap"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µwordmap, ßkeys, nil); πE != nil {
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
									µword = πTemp002
								}
								if πE != nil || !πTemp005 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp006 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
									continue
								}
								πTemp006[0] = µs
								if πE = πg.CheckLocal(πF, µword, "word"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetAttr(πF, µword, ßstartswith, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp006)
								if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
									continue
								}
								if πTemp005 {
									goto Label4
								}
								goto Label5
								// line 1718: possibilities = [word for word in wordmap.keys()
								πF.SetLineno(1718)
							Label4:
								// line 1718: possibilities = [word for word in wordmap.keys()
								πF.SetLineno(1718)
								if πE = πg.CheckLocal(πF, µword, "word"); πE != nil {
									continue
								}
								πF.PushCheckpoint(6)
								return µword, nil
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
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
						continue
					}
					µpossibilities = πTemp001
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					πTemp006[0] = µpossibilities
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp001, πE = πg.Eq(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µpossibilities); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 1721: if len(possibilities) == 1:
					πF.SetLineno(1721)
				Label4:
					// line 1722: return possibilities[0]
					πF.SetLineno(1722)
					πTemp001 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µpossibilities, πTemp001); πE != nil {
						continue
					}
					πR = πTemp005
					continue
					goto Label7
					// line 1723: elif not possibilities:
					πF.SetLineno(1723)
				Label5:
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[0] = µs
					if πTemp001, πE = πg.ResolveGlobal(πF, ßBadOptionError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					// line 1724: raise BadOptionError(s)
					πF.SetLineno(1724)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label7
				Label6:
					// line 1727: possibilities.sort()
					πF.SetLineno(1727)
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpossibilities, ßsort, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp006 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[0] = µs
					if πE = πg.CheckLocal(πF, µpossibilities, "possibilities"); πE != nil {
						continue
					}
					πTemp006[1] = µpossibilities
					if πTemp001, πE = πg.ResolveGlobal(πF, ßAmbiguousOptionError); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					// line 1728: raise AmbiguousOptionError(s, possibilities)
					πF.SetLineno(1728)
					πE = πF.Raise(πTemp005, nil, nil)
					continue
					goto Label7
				Label7:
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
			if πE = πF.Globals().SetItem(πF, ß_match_abbrev.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 1707: """_match_abbrev(s : string, wordmap : {string : Option}) -> string
			πF.SetLineno(1707)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πg.NewStr("_match_abbrev(s : string, wordmap : {string : Option}) -> string\n\n    Return the string key in 'wordmap' for which 's' is an unambiguous\n    abbreviation.  If 's' is found to be ambiguous or doesn't match any of\n    'words', raise BadOptionError.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp019, πE = πg.ResolveGlobal(πF, ß_match_abbrev); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp019, ß__doc__, πTemp018); πE != nil {
				continue
			}
			// line 1735: make_option = Option
			πF.SetLineno(1735)
			if πTemp018, πE = πg.ResolveGlobal(πF, ßOption); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmake_option.ToObject(), πTemp018); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("optparse", Code)
}

