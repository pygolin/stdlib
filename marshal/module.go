package marshal

import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/types"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/marshal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ßCodeType := πg.InternStr("CodeType")
		ßEOFError := πg.InternStr("EOFError")
		ßEllipsis := πg.InternStr("Ellipsis")
		ßF := πg.InternStr("F")
		ßFalse := πg.InternStr("False")
		ßI := πg.InternStr("I")
		ßImportError := πg.InternStr("ImportError")
		ßIndexError := πg.InternStr("IndexError")
		ßKeyError := πg.InternStr("KeyError")
		ßN := πg.InternStr("N")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßR := πg.InternStr("R")
		ßS := πg.InternStr("S")
		ßStopIteration := πg.InternStr("StopIteration")
		ßT := πg.InternStr("T")
		ßTYPE_CODE := πg.InternStr("TYPE_CODE")
		ßTYPE_COMPLEX := πg.InternStr("TYPE_COMPLEX")
		ßTYPE_DICT := πg.InternStr("TYPE_DICT")
		ßTYPE_ELLIPSIS := πg.InternStr("TYPE_ELLIPSIS")
		ßTYPE_FALSE := πg.InternStr("TYPE_FALSE")
		ßTYPE_FLOAT := πg.InternStr("TYPE_FLOAT")
		ßTYPE_FROZENSET := πg.InternStr("TYPE_FROZENSET")
		ßTYPE_INT := πg.InternStr("TYPE_INT")
		ßTYPE_INT64 := πg.InternStr("TYPE_INT64")
		ßTYPE_INTERNED := πg.InternStr("TYPE_INTERNED")
		ßTYPE_LIST := πg.InternStr("TYPE_LIST")
		ßTYPE_LONG := πg.InternStr("TYPE_LONG")
		ßTYPE_NONE := πg.InternStr("TYPE_NONE")
		ßTYPE_NULL := πg.InternStr("TYPE_NULL")
		ßTYPE_SET := πg.InternStr("TYPE_SET")
		ßTYPE_STOPITER := πg.InternStr("TYPE_STOPITER")
		ßTYPE_STRING := πg.InternStr("TYPE_STRING")
		ßTYPE_STRINGREF := πg.InternStr("TYPE_STRINGREF")
		ßTYPE_TRUE := πg.InternStr("TYPE_TRUE")
		ßTYPE_TUPLE := πg.InternStr("TYPE_TUPLE")
		ßTYPE_UNICODE := πg.InternStr("TYPE_UNICODE")
		ßTYPE_UNKNOWN := πg.InternStr("TYPE_UNKNOWN")
		ßTrue := πg.InternStr("True")
		ßValueError := πg.InternStr("ValueError")
		ß_FastUnmarshaller := πg.InternStr("_FastUnmarshaller")
		ß_Marshaller := πg.InternStr("_Marshaller")
		ß_NULL := πg.InternStr("_NULL")
		ß_StringBuffer := πg.InternStr("_StringBuffer")
		ß_Unmarshaller := πg.InternStr("_Unmarshaller")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_load_dispatch := πg.InternStr("_load_dispatch")
		ß_r_long := πg.InternStr("_r_long")
		ß_r_long64 := πg.InternStr("_r_long64")
		ß_r_short := πg.InternStr("_r_short")
		ß_read := πg.InternStr("_read")
		ß_read1 := πg.InternStr("_read1")
		ß_stringtable := πg.InternStr("_stringtable")
		ß_write := πg.InternStr("_write")
		ßappend := πg.InternStr("append")
		ßbool := πg.InternStr("bool")
		ßbufpos := πg.InternStr("bufpos")
		ßbufstr := πg.InternStr("bufstr")
		ßbuiltinify := πg.InternStr("builtinify")
		ßbytes := πg.InternStr("bytes")
		ßc := πg.InternStr("c")
		ßchr := πg.InternStr("chr")
		ßco_argcount := πg.InternStr("co_argcount")
		ßco_cellvars := πg.InternStr("co_cellvars")
		ßco_code := πg.InternStr("co_code")
		ßco_consts := πg.InternStr("co_consts")
		ßco_filename := πg.InternStr("co_filename")
		ßco_firstlineno := πg.InternStr("co_firstlineno")
		ßco_flags := πg.InternStr("co_flags")
		ßco_freevars := πg.InternStr("co_freevars")
		ßco_lnotab := πg.InternStr("co_lnotab")
		ßco_name := πg.InternStr("co_name")
		ßco_names := πg.InternStr("co_names")
		ßco_nlocals := πg.InternStr("co_nlocals")
		ßco_stacksize := πg.InternStr("co_stacksize")
		ßco_varnames := πg.InternStr("co_varnames")
		ßcomplex := πg.InternStr("complex")
		ßdecode := πg.InternStr("decode")
		ßdict := πg.InternStr("dict")
		ßdispatch := πg.InternStr("dispatch")
		ßdump := πg.InternStr("dump")
		ßdump_bool := πg.InternStr("dump_bool")
		ßdump_code := πg.InternStr("dump_code")
		ßdump_complex := πg.InternStr("dump_complex")
		ßdump_dict := πg.InternStr("dump_dict")
		ßdump_ellipsis := πg.InternStr("dump_ellipsis")
		ßdump_float := πg.InternStr("dump_float")
		ßdump_frozenset := πg.InternStr("dump_frozenset")
		ßdump_int := πg.InternStr("dump_int")
		ßdump_list := πg.InternStr("dump_list")
		ßdump_long := πg.InternStr("dump_long")
		ßdump_none := πg.InternStr("dump_none")
		ßdump_set := πg.InternStr("dump_set")
		ßdump_stopiter := πg.InternStr("dump_stopiter")
		ßdump_string := πg.InternStr("dump_string")
		ßdump_tuple := πg.InternStr("dump_tuple")
		ßdump_unicode := πg.InternStr("dump_unicode")
		ßdumps := πg.InternStr("dumps")
		ßencode := πg.InternStr("encode")
		ßf := πg.InternStr("f")
		ßfloat := πg.InternStr("float")
		ßfrozenset := πg.InternStr("frozenset")
		ßget := πg.InternStr("get")
		ßi := πg.InternStr("i")
		ßimag := πg.InternStr("imag")
		ßint := πg.InternStr("int")
		ßintern := πg.InternStr("intern")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßl := πg.InternStr("l")
		ßlen := πg.InternStr("len")
		ßlist := πg.InternStr("list")
		ßload := πg.InternStr("load")
		ßload_code := πg.InternStr("load_code")
		ßload_complex := πg.InternStr("load_complex")
		ßload_dict := πg.InternStr("load_dict")
		ßload_ellipsis := πg.InternStr("load_ellipsis")
		ßload_false := πg.InternStr("load_false")
		ßload_float := πg.InternStr("load_float")
		ßload_frozenset := πg.InternStr("load_frozenset")
		ßload_int := πg.InternStr("load_int")
		ßload_int64 := πg.InternStr("load_int64")
		ßload_interned := πg.InternStr("load_interned")
		ßload_list := πg.InternStr("load_list")
		ßload_long := πg.InternStr("load_long")
		ßload_none := πg.InternStr("load_none")
		ßload_null := πg.InternStr("load_null")
		ßload_set := πg.InternStr("load_set")
		ßload_stopiter := πg.InternStr("load_stopiter")
		ßload_string := πg.InternStr("load_string")
		ßload_stringref := πg.InternStr("load_stringref")
		ßload_true := πg.InternStr("load_true")
		ßload_tuple := πg.InternStr("load_tuple")
		ßload_unicode := πg.InternStr("load_unicode")
		ßloads := πg.InternStr("loads")
		ßlong := πg.InternStr("long")
		ßmro := πg.InternStr("mro")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßr_long := πg.InternStr("r_long")
		ßr_long64 := πg.InternStr("r_long64")
		ßr_short := πg.InternStr("r_short")
		ßrange := πg.InternStr("range")
		ßread := πg.InternStr("read")
		ßreal := πg.InternStr("real")
		ßrepr := πg.InternStr("repr")
		ßs := πg.InternStr("s")
		ßset := πg.InternStr("set")
		ßstr := πg.InternStr("str")
		ßt := πg.InternStr("t")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßtypes := πg.InternStr("types")
		ßu := πg.InternStr("u")
		ßunicode := πg.InternStr("unicode")
		ßutf8 := πg.InternStr("utf8")
		ßversion := πg.InternStr("version")
		ßw_long := πg.InternStr("w_long")
		ßw_long64 := πg.InternStr("w_long64")
		ßw_short := πg.InternStr("w_short")
		ßwrite := πg.InternStr("write")
		ßx := πg.InternStr("x")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.BaseException
		_ = πTemp003
		var πTemp004 *πg.Traceback
		_ = πTemp004
		var πTemp005 bool
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 *πg.Dict
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
			case 2:
				goto Label2
			case 5:
				goto Label5
			default:
				panic("unexpected function state")
			}
			// line 1: """Internal Python object serialization
			πF.SetLineno(1)
			// line 1: """Internal Python object serialization
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Internal Python object serialization\n\nThis module contains functions that can read and write Python values in a binary format. The format is specific to Python, but independent of machine architecture issues (e.g., you can write a Python value to a file on a PC, transport the file to a Sun, and read it back there). Details of the format may change between Python versions.\n").ToObject()); πE != nil {
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
			// line 11: try:
			πF.SetLineno(11)
			πF.PushCheckpoint(2)
			// line 12: intern
			πF.SetLineno(12)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßintern); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label3
			}
			πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
			continue
			// line 13: except NameError:
			πF.SetLineno(13)
		Label3:
			// line 14: from sys import intern
			πF.SetLineno(14)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp006, πE = πg.GetAttrImport(πF, πTemp001, ßintern); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßintern.ToObject(), πTemp006); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 16: try: from __pypy__ import builtinify
			πF.SetLineno(16)
			πF.PushCheckpoint(5)
			// line 16: try: from __pypy__ import builtinify
			πF.SetLineno(16)
			if πTemp002, πE = πg.ImportModule(πF, "__pypy__"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp006, πE = πg.GetAttrImport(πF, πTemp001, ßbuiltinify); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßbuiltinify.ToObject(), πTemp006); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label4
		Label5:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp003, πTemp004 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp005, πE = πg.IsInstance(πF, πTemp003.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp005 {
				goto Label6
			}
			πE = πF.Raise(πTemp003.ToObject(), nil, πTemp004.ToObject())
			continue
			// line 17: except ImportError: builtinify = lambda f: f
			πF.SetLineno(17)
		Label6:
			// line 17: except ImportError: builtinify = lambda f: f
			πF.SetLineno(17)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "f", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µf *πg.Object = πArgs[0]
				_ = µf
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 17: except ImportError: builtinify = lambda f: f
					πF.SetLineno(17)
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
			if πE = πF.Globals().SetItem(πF, ßbuiltinify.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
			// line 20: TYPE_NULL     = '0'
			πF.SetLineno(20)
			if πE = πF.Globals().SetItem(πF, ßTYPE_NULL.ToObject(), ß0.ToObject()); πE != nil {
				continue
			}
			// line 21: TYPE_NONE     = 'N'
			πF.SetLineno(21)
			if πE = πF.Globals().SetItem(πF, ßTYPE_NONE.ToObject(), ßN.ToObject()); πE != nil {
				continue
			}
			// line 22: TYPE_FALSE    = 'F'
			πF.SetLineno(22)
			if πE = πF.Globals().SetItem(πF, ßTYPE_FALSE.ToObject(), ßF.ToObject()); πE != nil {
				continue
			}
			// line 23: TYPE_TRUE     = 'T'
			πF.SetLineno(23)
			if πE = πF.Globals().SetItem(πF, ßTYPE_TRUE.ToObject(), ßT.ToObject()); πE != nil {
				continue
			}
			// line 24: TYPE_STOPITER = 'S'
			πF.SetLineno(24)
			if πE = πF.Globals().SetItem(πF, ßTYPE_STOPITER.ToObject(), ßS.ToObject()); πE != nil {
				continue
			}
			// line 25: TYPE_ELLIPSIS = '.'
			πF.SetLineno(25)
			if πE = πF.Globals().SetItem(πF, ßTYPE_ELLIPSIS.ToObject(), πg.NewStr(".").ToObject()); πE != nil {
				continue
			}
			// line 26: TYPE_INT      = 'i'
			πF.SetLineno(26)
			if πE = πF.Globals().SetItem(πF, ßTYPE_INT.ToObject(), ßi.ToObject()); πE != nil {
				continue
			}
			// line 27: TYPE_INT64    = 'I'
			πF.SetLineno(27)
			if πE = πF.Globals().SetItem(πF, ßTYPE_INT64.ToObject(), ßI.ToObject()); πE != nil {
				continue
			}
			// line 28: TYPE_FLOAT    = 'f'
			πF.SetLineno(28)
			if πE = πF.Globals().SetItem(πF, ßTYPE_FLOAT.ToObject(), ßf.ToObject()); πE != nil {
				continue
			}
			// line 29: TYPE_COMPLEX  = 'x'
			πF.SetLineno(29)
			if πE = πF.Globals().SetItem(πF, ßTYPE_COMPLEX.ToObject(), ßx.ToObject()); πE != nil {
				continue
			}
			// line 30: TYPE_LONG     = 'l'
			πF.SetLineno(30)
			if πE = πF.Globals().SetItem(πF, ßTYPE_LONG.ToObject(), ßl.ToObject()); πE != nil {
				continue
			}
			// line 31: TYPE_STRING   = 's'
			πF.SetLineno(31)
			if πE = πF.Globals().SetItem(πF, ßTYPE_STRING.ToObject(), ßs.ToObject()); πE != nil {
				continue
			}
			// line 32: TYPE_INTERNED = 't'
			πF.SetLineno(32)
			if πE = πF.Globals().SetItem(πF, ßTYPE_INTERNED.ToObject(), ßt.ToObject()); πE != nil {
				continue
			}
			// line 33: TYPE_STRINGREF= 'R'
			πF.SetLineno(33)
			if πE = πF.Globals().SetItem(πF, ßTYPE_STRINGREF.ToObject(), ßR.ToObject()); πE != nil {
				continue
			}
			// line 34: TYPE_TUPLE    = '('
			πF.SetLineno(34)
			if πE = πF.Globals().SetItem(πF, ßTYPE_TUPLE.ToObject(), πg.NewStr("(").ToObject()); πE != nil {
				continue
			}
			// line 35: TYPE_LIST     = '['
			πF.SetLineno(35)
			if πE = πF.Globals().SetItem(πF, ßTYPE_LIST.ToObject(), πg.NewStr("[").ToObject()); πE != nil {
				continue
			}
			// line 36: TYPE_DICT     = '{'
			πF.SetLineno(36)
			if πE = πF.Globals().SetItem(πF, ßTYPE_DICT.ToObject(), πg.NewStr("{").ToObject()); πE != nil {
				continue
			}
			// line 37: TYPE_CODE     = 'c'
			πF.SetLineno(37)
			if πE = πF.Globals().SetItem(πF, ßTYPE_CODE.ToObject(), ßc.ToObject()); πE != nil {
				continue
			}
			// line 38: TYPE_UNICODE  = 'u'
			πF.SetLineno(38)
			if πE = πF.Globals().SetItem(πF, ßTYPE_UNICODE.ToObject(), ßu.ToObject()); πE != nil {
				continue
			}
			// line 39: TYPE_UNKNOWN  = '?'
			πF.SetLineno(39)
			if πE = πF.Globals().SetItem(πF, ßTYPE_UNKNOWN.ToObject(), πg.NewStr("?").ToObject()); πE != nil {
				continue
			}
			// line 40: TYPE_SET      = '<'
			πF.SetLineno(40)
			if πE = πF.Globals().SetItem(πF, ßTYPE_SET.ToObject(), πg.NewStr("<").ToObject()); πE != nil {
				continue
			}
			// line 41: TYPE_FROZENSET= '>'
			πF.SetLineno(41)
			if πE = πF.Globals().SetItem(πF, ßTYPE_FROZENSET.ToObject(), πg.NewStr(">").ToObject()); πE != nil {
				continue
			}
			// line 43: class _Marshaller(object):
			πF.SetLineno(43)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Marshaller", "/usr/lib/python2.7/marshal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
				_ = πClass
				var πTemp001 *πg.Dict
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
				var πTemp013 []*πg.Object
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
				var πTemp021 bool
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2:
						goto Label2
					case 5:
						goto Label5
					case 8:
						goto Label8
					case 11:
						goto Label11
					case 14:
						goto Label14
					case 17:
						goto Label17
					case 20:
						goto Label20
					default:
						panic("unexpected function state")
					}
					// line 44: dispatch = {}
					πF.SetLineno(44)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßdispatch.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 46: def __init__(self, writefunc):
					πF.SetLineno(46)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "writefunc", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µwritefunc *πg.Object = πArgs[1]
						_ = µwritefunc
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 47: self._write = writefunc
							πF.SetLineno(47)
							if πE = πg.CheckLocal(πF, µwritefunc, "writefunc"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µwritefunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_write, πTemp001); πE != nil {
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
					// line 49: def dump(self, x):
					πF.SetLineno(49)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("dump", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µtp *πg.Object = πg.UnboundLocal
						_ = µtp
						var µfunc *πg.Object = πg.UnboundLocal
						_ = µfunc
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2:
								goto Label2
							case 4:
								goto Label4
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 50: try:
							πF.SetLineno(50)
							πF.PushCheckpoint(2)
							// line 51: self.dispatch[type(x)](self, x)
							πF.SetLineno(51)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[1] = µx
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp003[0] = µx
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002 = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label3
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 52: except KeyError:
							πF.SetLineno(52)
						Label3:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßmro, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
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
							if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µtp = πTemp004
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(4)
							// line 54: func = self.dispatch.get(tp)
							πF.SetLineno(54)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtp, "tp"); πE != nil {
								continue
							}
							πTemp001[0] = µtp
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßget, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfunc = πTemp004
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µfunc); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label7
							}
							goto Label8
							// line 55: if func:
							πF.SetLineno(55)
						Label7:
							// line 56: break
							πF.SetLineno(56)
							πTemp008 = true
							continue
							goto Label8
						Label8:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("unmarshallable object").ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 58: raise ValueError("unmarshallable object")
							πF.SetLineno(58)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
						Label6:
							// line 59: func(self, x)
							πF.SetLineno(59)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[1] = µx
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp002, πE = µfunc.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
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
					if πE = πClass.SetItem(πF, ßdump.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 61: def w_long64(self, x):
					πF.SetLineno(61)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("w_long64", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 62: self.w_long(x)
							πF.SetLineno(62)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 63: self.w_long(x>>32)
							πF.SetLineno(63)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.RShift(πF, µx, πg.NewInt(32).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßw_long64.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 65: def w_long(self, x):
					πF.SetLineno(65)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("w_long", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µa *πg.Object = πg.UnboundLocal
						_ = µa
						var µb *πg.Object = πg.UnboundLocal
						_ = µb
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 66: a = chr(x & 0xff)
							πF.SetLineno(66)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.And(πF, µx, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp003
							// line 67: x >>= 8
							πF.SetLineno(67)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IRShift(πF, µx, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							µx = πTemp002
							// line 68: b = chr(x & 0xff)
							πF.SetLineno(68)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.And(πF, µx, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µb = πTemp003
							// line 69: x >>= 8
							πF.SetLineno(69)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IRShift(πF, µx, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							µx = πTemp002
							// line 70: c = chr(x & 0xff)
							πF.SetLineno(70)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.And(πF, µx, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µc = πTemp003
							// line 71: x >>= 8
							πF.SetLineno(71)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IRShift(πF, µx, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							µx = πTemp002
							// line 72: d = chr(x & 0xff)
							πF.SetLineno(72)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.And(πF, µx, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 73: self._write(a + b + c + d)
							πF.SetLineno(73)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µa, µb); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, µc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, µd); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßw_long.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 75: def w_short(self, x):
					πF.SetLineno(75)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("w_short", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 76: self._write(chr((x)     & 0xff))
							πF.SetLineno(76)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µx, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 77: self._write(chr((x>> 8) & 0xff))
							πF.SetLineno(77)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.RShift(πF, µx, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, πTemp004, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßw_short.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 79: def dump_none(self, x):
					πF.SetLineno(79)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("dump_none", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 80: self._write(TYPE_NONE)
							πF.SetLineno(80)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_NONE); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_none.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 81: dispatch[type(None)] = dump_none
					πF.SetLineno(81)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßdump_none); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					πTemp013 = πF.MakeArgs(1)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp013[0] = πTemp014
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßtype); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp014.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					πTemp012 = πTemp015
					if πE = πg.SetItem(πF, πTemp011, πTemp012, πTemp010); πE != nil {
						continue
					}
					// line 83: def dump_bool(self, x):
					πF.SetLineno(83)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("dump_bool", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var πTemp001 bool
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IsTrue(πF, µx); πE != nil {
								continue
							}
							if πTemp001 {
								goto Label1
							}
							goto Label2
							// line 84: if x:
							πF.SetLineno(84)
						Label1:
							// line 85: self._write(TYPE_TRUE)
							πF.SetLineno(85)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTYPE_TRUE); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label3
						Label2:
							// line 87: self._write(TYPE_FALSE)
							πF.SetLineno(87)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTYPE_FALSE); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_bool.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 88: dispatch[bool] = dump_bool
					πF.SetLineno(88)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßdump_bool); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp010); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßbool); πE != nil {
						continue
					}
					πTemp014 = πTemp015
					if πE = πg.SetItem(πF, πTemp012, πTemp014, πTemp011); πE != nil {
						continue
					}
					// line 90: def dump_stopiter(self, x):
					πF.SetLineno(90)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("dump_stopiter", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µx != πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 91: if x is not StopIteration:
							πF.SetLineno(91)
						Label1:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("unmarshallable object").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 92: raise ValueError("unmarshallable object")
							πF.SetLineno(92)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 93: self._write(TYPE_STOPITER)
							πF.SetLineno(93)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTYPE_STOPITER); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_stopiter.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 94: dispatch[type(StopIteration)] = dump_stopiter
					πF.SetLineno(94)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdump_stopiter); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πTemp011); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					πTemp013 = πF.MakeArgs(1)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßStopIteration); πE != nil {
						continue
					}
					πTemp013[0] = πTemp016
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßtype); πE != nil {
						continue
					}
					if πTemp017, πE = πTemp016.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					πTemp015 = πTemp017
					if πE = πg.SetItem(πF, πTemp014, πTemp015, πTemp012); πE != nil {
						continue
					}
					// line 96: def dump_ellipsis(self, x):
					πF.SetLineno(96)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("dump_ellipsis", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 97: self._write(TYPE_ELLIPSIS)
							πF.SetLineno(97)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_ELLIPSIS); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_ellipsis.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 99: try:
					πF.SetLineno(99)
					πF.PushCheckpoint(2)
					// line 100: dispatch[type(Ellipsis)] = dump_ellipsis
					πF.SetLineno(100)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßdump_ellipsis); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp012); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					πTemp013 = πF.MakeArgs(1)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßEllipsis); πE != nil {
						continue
					}
					πTemp013[0] = πTemp017
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßtype); πE != nil {
						continue
					}
					if πTemp018, πE = πTemp017.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					πTemp016 = πTemp018
					if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp019, πTemp020 = πF.ExcInfo()
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNameError); πE != nil {
						continue
					}
					if πTemp021, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp012); πE != nil {
						continue
					}
					if πTemp021 {
						goto Label3
					}
					πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
					continue
					// line 101: except NameError:
					πF.SetLineno(101)
				Label3:
					// line 102: pass
					πF.SetLineno(102)
					πF.RestoreExc(nil, nil)
					goto Label1
				Label1:
					// line 105: def dump_int(self, x):
					πF.SetLineno(105)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("dump_int", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µy *πg.Object = πg.UnboundLocal
						_ = µy
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 106: y = x>>31
							πF.SetLineno(106)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.RShift(πF, µx, πg.NewInt(31).ToObject()); πE != nil {
								continue
							}
							µy = πTemp001
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							πTemp001 = µy
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µy, "y"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.NE(πF, µy, πTemp004); πE != nil {
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
							// line 107: if y and y != -1:
							πF.SetLineno(107)
						Label2:
							// line 108: self._write(TYPE_INT64)
							πF.SetLineno(108)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTYPE_INT64); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 109: self.w_long64(x)
							πF.SetLineno(109)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp005[0] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßw_long64, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label4
						Label3:
							// line 111: self._write(TYPE_INT)
							πF.SetLineno(111)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTYPE_INT); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 112: self.w_long(x)
							πF.SetLineno(112)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp005[0] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßdump_int.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 113: dispatch[int] = dump_int
					πF.SetLineno(113)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßdump_int); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πTemp014); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßint); πE != nil {
						continue
					}
					πTemp017 = πTemp018
					if πE = πg.SetItem(πF, πTemp016, πTemp017, πTemp015); πE != nil {
						continue
					}
					// line 115: def dump_long(self, x):
					πF.SetLineno(115)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("dump_long", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µsign *πg.Object = πg.UnboundLocal
						_ = µsign
						var µdigits *πg.Object = πg.UnboundLocal
						_ = µdigits
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3:
								goto Label3
							case 4:
								goto Label4
							case 6:
								goto Label6
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 116: self._write(TYPE_LONG)
							πF.SetLineno(116)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_LONG); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 117: sign = 1
							πF.SetLineno(117)
							µsign = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µx, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 118: if x < 0:
							πF.SetLineno(118)
						Label1:
							// line 119: sign = -1
							πF.SetLineno(119)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µsign = πTemp002
							// line 120: x = -x
							πF.SetLineno(120)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, µx); πE != nil {
								continue
							}
							µx = πTemp002
							goto Label2
						Label2:
							// line 121: digits = []
							πF.SetLineno(121)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µdigits = πTemp002
							// line 122: while x:
							πF.SetLineno(122)
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
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µx); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 123: digits.append(x & 0x7FFF)
							πF.SetLineno(123)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.And(πF, µx, πg.NewInt(32767).ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µdigits, "digits"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdigits, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 124: x = x>>15
							πF.SetLineno(124)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.RShift(πF, µx, πg.NewInt(15).ToObject()); πE != nil {
								continue
							}
							µx = πTemp002
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 125: self.w_long(len(digits) * sign)
							πF.SetLineno(125)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdigits, "digits"); πE != nil {
								continue
							}
							πTemp006[0] = µdigits
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, πTemp007, µsign); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µdigits, "digits"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µdigits); πE != nil {
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
								µd = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 127: self.w_short(d)
							πF.SetLineno(127)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßw_short, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßdump_long.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 128: try:
					πF.SetLineno(128)
					πF.PushCheckpoint(5)
					// line 129: long
					πF.SetLineno(129)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßlong); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					// line 133: dispatch[long] = dump_long
					πF.SetLineno(133)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßdump_long); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp015); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßlong); πE != nil {
						continue
					}
					πTemp018 = πTemp022
					if πE = πg.SetItem(πF, πTemp017, πTemp018, πTemp016); πE != nil {
						continue
					}
					goto Label4
				Label5:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp019, πTemp020 = πF.ExcInfo()
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßNameError); πE != nil {
						continue
					}
					if πTemp021, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp015); πE != nil {
						continue
					}
					if πTemp021 {
						goto Label6
					}
					πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
					continue
					// line 130: except NameError:
					πF.SetLineno(130)
				Label6:
					// line 131: dispatch[int] = dump_long
					πF.SetLineno(131)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßdump_long); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp015); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßint); πE != nil {
						continue
					}
					πTemp018 = πTemp022
					if πE = πg.SetItem(πF, πTemp017, πTemp018, πTemp016); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					// line 135: def dump_float(self, x):
					πF.SetLineno(135)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("dump_float", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 136: write = self._write
							πF.SetLineno(136)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							// line 137: write(TYPE_FLOAT)
							πF.SetLineno(137)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTYPE_FLOAT); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 138: s = repr(x)
							πF.SetLineno(138)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp002[0] = µx
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 139: write(chr(len(s)))
							πF.SetLineno(139)
							πTemp002 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp005[0] = µs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 140: write(s)
							πF.SetLineno(140)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_float.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 141: dispatch[float] = dump_float
					πF.SetLineno(141)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßdump_float); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πTemp016); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßfloat); πE != nil {
						continue
					}
					πTemp022 = πTemp023
					if πE = πg.SetItem(πF, πTemp018, πTemp022, πTemp017); πE != nil {
						continue
					}
					// line 143: def dump_complex(self, x):
					πF.SetLineno(143)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("dump_complex", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 144: write = self._write
							πF.SetLineno(144)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							// line 145: write(TYPE_COMPLEX)
							πF.SetLineno(145)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTYPE_COMPLEX); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 146: s = repr(x.real)
							πF.SetLineno(146)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µx, ßreal, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 147: write(chr(len(s)))
							πF.SetLineno(147)
							πTemp002 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp005[0] = µs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 148: write(s)
							πF.SetLineno(148)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 149: s = repr(x.imag)
							πF.SetLineno(149)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µx, ßimag, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µs = πTemp003
							// line 150: write(chr(len(s)))
							πF.SetLineno(150)
							πTemp002 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp005[0] = µs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[0] = πTemp003
							if πTemp001, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 151: write(s)
							πF.SetLineno(151)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_complex.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 152: try:
					πF.SetLineno(152)
					πF.PushCheckpoint(8)
					// line 153: dispatch[complex] = dump_complex
					πF.SetLineno(153)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßdump_complex); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πTemp017); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßcomplex); πE != nil {
						continue
					}
					πTemp023 = πTemp024
					if πE = πg.SetItem(πF, πTemp022, πTemp023, πTemp018); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label7
				Label8:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp019, πTemp020 = πF.ExcInfo()
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßNameError); πE != nil {
						continue
					}
					if πTemp021, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp017); πE != nil {
						continue
					}
					if πTemp021 {
						goto Label9
					}
					πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
					continue
					// line 154: except NameError:
					πF.SetLineno(154)
				Label9:
					// line 155: pass
					πF.SetLineno(155)
					πF.RestoreExc(nil, nil)
					goto Label7
				Label7:
					// line 157: def dump_string(self, x):
					πF.SetLineno(157)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("dump_string", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 160: self._write(TYPE_STRING)
							πF.SetLineno(160)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_STRING); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 161: self.w_long(len(x))
							πF.SetLineno(161)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 162: self._write(x)
							πF.SetLineno(162)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_string.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 163: dispatch[bytes] = dump_string
					πF.SetLineno(163)
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßdump_string); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πTemp018); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßbytes); πE != nil {
						continue
					}
					πTemp024 = πTemp025
					if πE = πg.SetItem(πF, πTemp023, πTemp024, πTemp022); πE != nil {
						continue
					}
					// line 165: def dump_unicode(self, x):
					πF.SetLineno(165)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("dump_unicode", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 166: self._write(TYPE_UNICODE)
							πF.SetLineno(166)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_UNICODE); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 167: s = x.encode('utf8')
							πF.SetLineno(167)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßutf8.ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßencode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 168: self.w_long(len(s))
							πF.SetLineno(168)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp004[0] = µs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 169: self._write(s)
							πF.SetLineno(169)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_unicode.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 170: try:
					πF.SetLineno(170)
					πF.PushCheckpoint(11)
					// line 171: unicode
					πF.SetLineno(171)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßunicode); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					// line 175: dispatch[unicode] = dump_unicode
					πF.SetLineno(175)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßdump_unicode); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πTemp022); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßunicode); πE != nil {
						continue
					}
					πTemp025 = πTemp026
					if πE = πg.SetItem(πF, πTemp024, πTemp025, πTemp023); πE != nil {
						continue
					}
					goto Label10
				Label11:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp019, πTemp020 = πF.ExcInfo()
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßNameError); πE != nil {
						continue
					}
					if πTemp021, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp022); πE != nil {
						continue
					}
					if πTemp021 {
						goto Label12
					}
					πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
					continue
					// line 172: except NameError:
					πF.SetLineno(172)
				Label12:
					// line 173: dispatch[str] = dump_unicode
					πF.SetLineno(173)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßdump_unicode); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πTemp022); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßstr); πE != nil {
						continue
					}
					πTemp025 = πTemp026
					if πE = πg.SetItem(πF, πTemp024, πTemp025, πTemp023); πE != nil {
						continue
					}
					πF.RestoreExc(nil, nil)
					goto Label10
				Label10:
					// line 177: def dump_tuple(self, x):
					πF.SetLineno(177)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("dump_tuple", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µitem *πg.Object = πg.UnboundLocal
						_ = µitem
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 178: self._write(TYPE_TUPLE)
							πF.SetLineno(178)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_TUPLE); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 179: self.w_long(len(x))
							πF.SetLineno(179)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µx); πE != nil {
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
								µitem = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 181: self.dump(item)
							πF.SetLineno(181)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_tuple.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 182: dispatch[tuple] = dump_tuple
					πF.SetLineno(182)
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßdump_tuple); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πTemp023); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßtuple); πE != nil {
						continue
					}
					πTemp026 = πTemp027
					if πE = πg.SetItem(πF, πTemp025, πTemp026, πTemp024); πE != nil {
						continue
					}
					// line 184: def dump_list(self, x):
					πF.SetLineno(184)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("dump_list", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µitem *πg.Object = πg.UnboundLocal
						_ = µitem
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 185: self._write(TYPE_LIST)
							πF.SetLineno(185)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_LIST); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 186: self.w_long(len(x))
							πF.SetLineno(186)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µx); πE != nil {
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
								µitem = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 188: self.dump(item)
							πF.SetLineno(188)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
								continue
							}
							πTemp001[0] = µitem
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_list.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 189: dispatch[list] = dump_list
					πF.SetLineno(189)
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßdump_list); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πTemp024); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßlist); πE != nil {
						continue
					}
					πTemp027 = πTemp028
					if πE = πg.SetItem(πF, πTemp026, πTemp027, πTemp025); πE != nil {
						continue
					}
					// line 191: def dump_dict(self, x):
					πF.SetLineno(191)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("dump_dict", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 192: self._write(TYPE_DICT)
							πF.SetLineno(192)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_DICT); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µx, ßitems, nil); πE != nil {
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
							// line 194: self.dump(key)
							πF.SetLineno(194)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 195: self.dump(value)
							πF.SetLineno(195)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
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
							// line 196: self._write(TYPE_NULL)
							πF.SetLineno(196)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_NULL); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_dict.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 197: dispatch[dict] = dump_dict
					πF.SetLineno(197)
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßdump_dict); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πTemp025); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßdict); πE != nil {
						continue
					}
					πTemp028 = πTemp029
					if πE = πg.SetItem(πF, πTemp027, πTemp028, πTemp026); πE != nil {
						continue
					}
					// line 199: def dump_code(self, x):
					πF.SetLineno(199)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("dump_code", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 200: self._write(TYPE_CODE)
							πF.SetLineno(200)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_CODE); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 201: self.w_long(x.co_argcount)
							πF.SetLineno(201)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_argcount, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 202: self.w_long(x.co_nlocals)
							πF.SetLineno(202)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_nlocals, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 203: self.w_long(x.co_stacksize)
							πF.SetLineno(203)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_stacksize, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 204: self.w_long(x.co_flags)
							πF.SetLineno(204)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_flags, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 205: self.dump(x.co_code)
							πF.SetLineno(205)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_code, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 206: self.dump(x.co_consts)
							πF.SetLineno(206)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_consts, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 207: self.dump(x.co_names)
							πF.SetLineno(207)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_names, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 208: self.dump(x.co_varnames)
							πF.SetLineno(208)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_varnames, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 209: self.dump(x.co_freevars)
							πF.SetLineno(209)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_freevars, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 210: self.dump(x.co_cellvars)
							πF.SetLineno(210)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_cellvars, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 211: self.dump(x.co_filename)
							πF.SetLineno(211)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_filename, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 212: self.dump(x.co_name)
							πF.SetLineno(212)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_name, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 213: self.w_long(x.co_firstlineno)
							πF.SetLineno(213)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_firstlineno, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 214: self.dump(x.co_lnotab)
							πF.SetLineno(214)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µx, ßco_lnotab, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_code.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 215: try:
					πF.SetLineno(215)
					πF.PushCheckpoint(14)
					// line 216: dispatch[types.CodeType] = dump_code
					πF.SetLineno(216)
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßdump_code); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πTemp026); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßtypes); πE != nil {
						continue
					}
					if πTemp031, πE = πg.GetAttr(πF, πTemp030, ßCodeType, nil); πE != nil {
						continue
					}
					πTemp029 = πTemp031
					if πE = πg.SetItem(πF, πTemp028, πTemp029, πTemp027); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label13
				Label14:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp019, πTemp020 = πF.ExcInfo()
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßNameError); πE != nil {
						continue
					}
					if πTemp021, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp026); πE != nil {
						continue
					}
					if πTemp021 {
						goto Label15
					}
					πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
					continue
					// line 217: except NameError:
					πF.SetLineno(217)
				Label15:
					// line 218: pass
					πF.SetLineno(218)
					πF.RestoreExc(nil, nil)
					goto Label13
				Label13:
					// line 220: def dump_set(self, x):
					πF.SetLineno(220)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("dump_set", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µeach *πg.Object = πg.UnboundLocal
						_ = µeach
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 221: self._write(TYPE_SET)
							πF.SetLineno(221)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_SET); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 222: self.w_long(len(x))
							πF.SetLineno(222)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µx); πE != nil {
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
								µeach = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 224: self.dump(each)
							πF.SetLineno(224)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µeach, "each"); πE != nil {
								continue
							}
							πTemp001[0] = µeach
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_set.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 225: try:
					πF.SetLineno(225)
					πF.PushCheckpoint(17)
					// line 226: dispatch[set] = dump_set
					πF.SetLineno(226)
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßdump_set); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp027); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßset); πE != nil {
						continue
					}
					πTemp030 = πTemp031
					if πE = πg.SetItem(πF, πTemp029, πTemp030, πTemp028); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label16
				Label17:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp019, πTemp020 = πF.ExcInfo()
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßNameError); πE != nil {
						continue
					}
					if πTemp021, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp027); πE != nil {
						continue
					}
					if πTemp021 {
						goto Label18
					}
					πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
					continue
					// line 227: except NameError:
					πF.SetLineno(227)
				Label18:
					// line 228: pass
					πF.SetLineno(228)
					πF.RestoreExc(nil, nil)
					goto Label16
				Label16:
					// line 230: def dump_frozenset(self, x):
					πF.SetLineno(230)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "x", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("dump_frozenset", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µx *πg.Object = πArgs[1]
						_ = µx
						var µeach *πg.Object = πg.UnboundLocal
						_ = µeach
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 231: self._write(TYPE_FROZENSET)
							πF.SetLineno(231)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTYPE_FROZENSET); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_write, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 232: self.w_long(len(x))
							πF.SetLineno(232)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp004[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßw_long, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µx); πE != nil {
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
								µeach = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 234: self.dump(each)
							πF.SetLineno(234)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µeach, "each"); πE != nil {
								continue
							}
							πTemp001[0] = µeach
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdump, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump_frozenset.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 235: try:
					πF.SetLineno(235)
					πF.PushCheckpoint(20)
					// line 236: dispatch[frozenset] = dump_frozenset
					πF.SetLineno(236)
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßdump_frozenset); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp029}, πTemp028); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßfrozenset); πE != nil {
						continue
					}
					πTemp031 = πTemp032
					if πE = πg.SetItem(πF, πTemp030, πTemp031, πTemp029); πE != nil {
						continue
					}
					πF.PopCheckpoint()
					goto Label19
				Label20:
					if πE == nil {
						continue
					}
					πE = nil
					πTemp019, πTemp020 = πF.ExcInfo()
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßNameError); πE != nil {
						continue
					}
					if πTemp021, πE = πg.IsInstance(πF, πTemp019.ToObject(), πTemp028); πE != nil {
						continue
					}
					if πTemp021 {
						goto Label21
					}
					πE = πF.Raise(πTemp019.ToObject(), nil, πTemp020.ToObject())
					continue
					// line 237: except NameError:
					πF.SetLineno(237)
				Label21:
					// line 238: pass
					πF.SetLineno(238)
					πF.RestoreExc(nil, nil)
					goto Label19
				Label19:
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("_Marshaller").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Marshaller.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 240: class _NULL(object):
			πF.SetLineno(240)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_NULL", "/usr/lib/python2.7/marshal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 241: pass
					πF.SetLineno(241)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("_NULL").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_NULL.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 243: class _StringBuffer(object):
			πF.SetLineno(243)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_StringBuffer", "/usr/lib/python2.7/marshal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
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
					default:
						panic("unexpected function state")
					}
					// line 244: def __init__(self, value):
					πF.SetLineno(244)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/marshal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvalue *πg.Object = πArgs[1]
						_ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 245: self.bufstr = value
							πF.SetLineno(245)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbufstr, πTemp001); πE != nil {
								continue
							}
							// line 246: self.bufpos = 0
							πF.SetLineno(246)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbufpos, πTemp001); πE != nil {
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
					// line 248: def read(self, n):
					πF.SetLineno(248)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "n", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("read", "/usr/lib/python2.7/marshal.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πArgs[1]
						_ = µn
						var µpos *πg.Object = πg.UnboundLocal
						_ = µpos
						var µnewpos *πg.Object = πg.UnboundLocal
						_ = µnewpos
						var µret *πg.Object = πg.UnboundLocal
						_ = µret
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 249: pos = self.bufpos
							πF.SetLineno(249)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbufpos, nil); πE != nil {
								continue
							}
							µpos = πTemp001
							// line 250: newpos = pos + n
							πF.SetLineno(250)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µpos, µn); πE != nil {
								continue
							}
							µnewpos = πTemp001
							// line 251: ret = self.bufstr[pos : newpos]
							πF.SetLineno(251)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µpos, µnewpos, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbufstr, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µret = πTemp002
							// line 252: self.bufpos = newpos
							πF.SetLineno(252)
							if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnewpos); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbufpos, πTemp001); πE != nil {
								continue
							}
							// line 253: return ret
							πF.SetLineno(253)
							if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
								continue
							}
							πR = µret
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßread.ToObject(), πTemp003); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("_StringBuffer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_StringBuffer.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 256: class _Unmarshaller(object):
			πF.SetLineno(256)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp009
			πTemp008 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_Unmarshaller", "/usr/lib/python2.7/marshal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
				_ = πClass
				var πTemp001 *πg.Dict
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 257: dispatch = {}
					πF.SetLineno(257)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßdispatch.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 259: def __init__(self, readfunc):
					πF.SetLineno(259)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "readfunc", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µreadfunc *πg.Object = πArgs[1]
						_ = µreadfunc
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 260: self._read = readfunc
							πF.SetLineno(260)
							if πE = πg.CheckLocal(πF, µreadfunc, "readfunc"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µreadfunc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_read, πTemp001); πE != nil {
								continue
							}
							// line 261: self._stringtable = []
							πF.SetLineno(261)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stringtable, πTemp003); πE != nil {
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
					// line 263: def load(self):
					πF.SetLineno(263)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("load", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
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
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 264: c = self._read(1)
							πF.SetLineno(264)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µc = πTemp003
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µc); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 265: if not c:
							πF.SetLineno(265)
						Label1:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEOFError); πE != nil {
								continue
							}
							// line 266: raise EOFError
							πF.SetLineno(266)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 267: try:
							πF.SetLineno(267)
							πF.PushCheckpoint(4)
							// line 268: return self.dispatch[c](self)
							πF.SetLineno(268)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp002 = µc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp002
							continue
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 269: except KeyError:
							πF.SetLineno(269)
						Label5:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp008[0] = µc
							if πTemp005, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp003 = πg.NewTuple2(µc, πTemp009).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("bad marshal code: %c (%d)").ToObject(), πTemp003); πE != nil {
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
							// line 270: raise ValueError("bad marshal code: %c (%d)" % (c, ord(c)))
							πF.SetLineno(270)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
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
					if πE = πClass.SetItem(πF, ßload.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 272: def r_short(self):
					πF.SetLineno(272)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("r_short", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlo *πg.Object = πg.UnboundLocal
						_ = µlo
						var µhi *πg.Object = πg.UnboundLocal
						_ = µhi
						var µx *πg.Object = πg.UnboundLocal
						_ = µx
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 273: lo = ord(self._read(1))
							πF.SetLineno(273)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlo = πTemp004
							// line 274: hi = ord(self._read(1))
							πF.SetLineno(274)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µhi = πTemp004
							// line 275: x = lo | (hi<<8)
							πF.SetLineno(275)
							if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LShift(πF, µhi, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Or(πF, µlo, πTemp004); πE != nil {
								continue
							}
							µx = πTemp003
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µx, πg.NewInt(32768).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 276: if x & 0x8000:
							πF.SetLineno(276)
						Label1:
							// line 277: x = x - 0x10000
							πF.SetLineno(277)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, µx, πg.NewInt(65536).ToObject()); πE != nil {
								continue
							}
							µx = πTemp003
							goto Label2
						Label2:
							// line 278: return x
							πF.SetLineno(278)
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
					if πE = πClass.SetItem(πF, ßr_short.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 280: def r_long(self):
					πF.SetLineno(280)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("r_long", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var µa *πg.Object = πg.UnboundLocal
						_ = µa
						var µb *πg.Object = πg.UnboundLocal
						_ = µb
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
						var µx *πg.Object = πg.UnboundLocal
						_ = µx
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 281: s = self._read(4)
							πF.SetLineno(281)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 282: a = ord(s[0])
							πF.SetLineno(282)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp003
							// line 283: b = ord(s[1])
							πF.SetLineno(283)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µb = πTemp003
							// line 284: c = ord(s[2])
							πF.SetLineno(284)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µc = πTemp003
							// line 285: d = ord(s[3])
							πF.SetLineno(285)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(3).ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp003
							// line 286: x = a | (b<<8) | (c<<16) | (d<<24)
							πF.SetLineno(286)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LShift(πF, µb, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Or(πF, µa, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LShift(πF, µc, πg.NewInt(16).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Or(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LShift(πF, µd, πg.NewInt(24).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Or(πF, πTemp003, πTemp004); πE != nil {
								continue
							}
							µx = πTemp002
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.And(πF, µd, πg.NewInt(128).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µx, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
						Label1:
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							goto Label3
							// line 287: if d & 0x80 and x > 0:
							πF.SetLineno(287)
						Label2:
							// line 288: x = -((1<<32) - x)
							πF.SetLineno(288)
							if πTemp004, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(32).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp004, µx); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, πTemp003); πE != nil {
								continue
							}
							µx = πTemp002
							// line 289: return int(x)
							πF.SetLineno(289)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp001[0] = µx
							if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
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
							// line 291: return x
							πF.SetLineno(291)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πR = µx
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
					if πE = πClass.SetItem(πF, ßr_long.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 293: def r_long64(self):
					πF.SetLineno(293)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("r_long64", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µa *πg.Object = πg.UnboundLocal
						_ = µa
						var µb *πg.Object = πg.UnboundLocal
						_ = µb
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
						var µe *πg.Object = πg.UnboundLocal
						_ = µe
						var µf *πg.Object = πg.UnboundLocal
						_ = µf
						var µg *πg.Object = πg.UnboundLocal
						_ = µg
						var µh *πg.Object = πg.UnboundLocal
						_ = µh
						var µx *πg.Object = πg.UnboundLocal
						_ = µx
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 294: a = ord(self._read(1))
							πF.SetLineno(294)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µa = πTemp004
							// line 295: b = ord(self._read(1))
							πF.SetLineno(295)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µb = πTemp004
							// line 296: c = ord(self._read(1))
							πF.SetLineno(296)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µc = πTemp004
							// line 297: d = ord(self._read(1))
							πF.SetLineno(297)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp004
							// line 298: e = ord(self._read(1))
							πF.SetLineno(298)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µe = πTemp004
							// line 299: f = ord(self._read(1))
							πF.SetLineno(299)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp004
							// line 300: g = ord(self._read(1))
							πF.SetLineno(300)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µg = πTemp004
							// line 301: h = ord(self._read(1))
							πF.SetLineno(301)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µh = πTemp004
							// line 302: x = a | (b<<8) | (c<<16) | (d<<24)
							πF.SetLineno(302)
							if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LShift(πF, µb, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Or(πF, µa, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LShift(πF, µc, πg.NewInt(16).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Or(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LShift(πF, µd, πg.NewInt(24).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Or(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							µx = πTemp003
							// line 303: x = x | (e<<32) | (f<<40) | (g<<48) | (h<<56)
							πF.SetLineno(303)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.LShift(πF, µe, πg.NewInt(32).ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Or(πF, µx, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.LShift(πF, µf, πg.NewInt(40).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Or(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.LShift(πF, µg, πg.NewInt(48).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Or(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LShift(πF, µh, πg.NewInt(56).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Or(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							µx = πTemp003
							if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.And(πF, µh, πg.NewInt(128).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp008 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, µx, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
						Label1:
							if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label2
							}
							goto Label3
							// line 304: if h & 0x80 and x > 0:
							πF.SetLineno(304)
						Label2:
							// line 305: x = -((1<<64) - x)
							πF.SetLineno(305)
							if πTemp005, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(64).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Sub(πF, πTemp005, µx); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πTemp004); πE != nil {
								continue
							}
							µx = πTemp003
							goto Label3
						Label3:
							// line 306: return x
							πF.SetLineno(306)
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
					if πE = πClass.SetItem(πF, ßr_long64.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 308: def load_null(self):
					πF.SetLineno(308)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("load_null", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 309: return _NULL
							πF.SetLineno(309)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_NULL); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_null.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 310: dispatch[TYPE_NULL] = load_null
					πF.SetLineno(310)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßload_null); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_NULL); πE != nil {
						continue
					}
					πTemp012 = πTemp013
					if πE = πg.SetItem(πF, πTemp011, πTemp012, πTemp010); πE != nil {
						continue
					}
					// line 312: def load_none(self):
					πF.SetLineno(312)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("load_none", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 313: return None
							πF.SetLineno(313)
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
					if πE = πClass.SetItem(πF, ßload_none.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 314: dispatch[TYPE_NONE] = load_none
					πF.SetLineno(314)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßload_none); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp010); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_NONE); πE != nil {
						continue
					}
					πTemp013 = πTemp014
					if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
						continue
					}
					// line 316: def load_true(self):
					πF.SetLineno(316)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("load_true", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 317: return True
							πF.SetLineno(317)
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
					if πE = πClass.SetItem(πF, ßload_true.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 318: dispatch[TYPE_TRUE] = load_true
					πF.SetLineno(318)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßload_true); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πTemp011); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_TRUE); πE != nil {
						continue
					}
					πTemp014 = πTemp015
					if πE = πg.SetItem(πF, πTemp013, πTemp014, πTemp012); πE != nil {
						continue
					}
					// line 320: def load_false(self):
					πF.SetLineno(320)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("load_false", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 321: return False
							πF.SetLineno(321)
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
					if πE = πClass.SetItem(πF, ßload_false.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 322: dispatch[TYPE_FALSE] = load_false
					πF.SetLineno(322)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßload_false); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πTemp012); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_FALSE); πE != nil {
						continue
					}
					πTemp015 = πTemp016
					if πE = πg.SetItem(πF, πTemp014, πTemp015, πTemp013); πE != nil {
						continue
					}
					// line 324: def load_stopiter(self):
					πF.SetLineno(324)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("load_stopiter", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 325: return StopIteration
							πF.SetLineno(325)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_stopiter.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 326: dispatch[TYPE_STOPITER] = load_stopiter
					πF.SetLineno(326)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßload_stopiter); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_STOPITER); πE != nil {
						continue
					}
					πTemp016 = πTemp017
					if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
						continue
					}
					// line 328: def load_ellipsis(self):
					πF.SetLineno(328)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("load_ellipsis", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 329: return Ellipsis
							πF.SetLineno(329)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEllipsis); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_ellipsis.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 330: dispatch[TYPE_ELLIPSIS] = load_ellipsis
					πF.SetLineno(330)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßload_ellipsis); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πTemp014); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_ELLIPSIS); πE != nil {
						continue
					}
					πTemp017 = πTemp018
					if πE = πg.SetItem(πF, πTemp016, πTemp017, πTemp015); πE != nil {
						continue
					}
					// line 332: dispatch[TYPE_INT] = r_long
					πF.SetLineno(332)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßr_long); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πTemp014); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_INT); πE != nil {
						continue
					}
					πTemp017 = πTemp018
					if πE = πg.SetItem(πF, πTemp016, πTemp017, πTemp015); πE != nil {
						continue
					}
					// line 334: dispatch[TYPE_INT64] = r_long64
					πF.SetLineno(334)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßr_long64); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πTemp014); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_INT64); πE != nil {
						continue
					}
					πTemp017 = πTemp018
					if πE = πg.SetItem(πF, πTemp016, πTemp017, πTemp015); πE != nil {
						continue
					}
					// line 336: def load_long(self):
					πF.SetLineno(336)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("load_long", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsize *πg.Object = πg.UnboundLocal
						_ = µsize
						var µsign *πg.Object = πg.UnboundLocal
						_ = µsign
						var µx *πg.Object = πg.UnboundLocal
						_ = µx
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 337: size = self.r_long()
							πF.SetLineno(337)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µsize = πTemp002
							// line 338: sign = 1
							πF.SetLineno(338)
							µsign = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µsize, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 339: if size < 0:
							πF.SetLineno(339)
						Label1:
							// line 340: sign = -1
							πF.SetLineno(340)
							if πTemp001, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µsign = πTemp001
							// line 341: size = -size
							πF.SetLineno(341)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Neg(πF, µsize); πE != nil {
								continue
							}
							µsize = πTemp001
							goto Label2
						Label2:
							// line 342: x = 0
							πF.SetLineno(342)
							µx = πg.NewInt(0).ToObject()
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp004[0] = µsize
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
								µi = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 344: d = self.r_short()
							πF.SetLineno(344)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßr_short, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µd = πTemp005
							// line 345: x = x | (d<<(i*15))
							πF.SetLineno(345)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Mul(πF, µi, πg.NewInt(15).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LShift(πF, µd, πTemp007); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Or(πF, µx, πTemp005); πE != nil {
								continue
							}
							µx = πTemp002
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 346: return x * sign
							πF.SetLineno(346)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, µx, µsign); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_long.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 347: dispatch[TYPE_LONG] = load_long
					πF.SetLineno(347)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßload_long); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp015); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_LONG); πE != nil {
						continue
					}
					πTemp018 = πTemp019
					if πE = πg.SetItem(πF, πTemp017, πTemp018, πTemp016); πE != nil {
						continue
					}
					// line 349: def load_float(self):
					πF.SetLineno(349)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("load_float", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 350: n = ord(self._read(1))
							πF.SetLineno(350)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp004
							// line 351: s = self._read(n)
							πF.SetLineno(351)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp004
							// line 352: return float(s)
							πF.SetLineno(352)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_float.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 353: dispatch[TYPE_FLOAT] = load_float
					πF.SetLineno(353)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßload_float); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πTemp016); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_FLOAT); πE != nil {
						continue
					}
					πTemp019 = πTemp020
					if πE = πg.SetItem(πF, πTemp018, πTemp019, πTemp017); πE != nil {
						continue
					}
					// line 355: def load_complex(self):
					πF.SetLineno(355)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("load_complex", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var µreal *πg.Object = πg.UnboundLocal
						_ = µreal
						var µimag *πg.Object = πg.UnboundLocal
						_ = µimag
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 356: n = ord(self._read(1))
							πF.SetLineno(356)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp004
							// line 357: s = self._read(n)
							πF.SetLineno(357)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp004
							// line 358: real = float(s)
							πF.SetLineno(358)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µreal = πTemp004
							// line 359: n = ord(self._read(1))
							πF.SetLineno(359)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp004
							// line 360: s = self._read(n)
							πF.SetLineno(360)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp004
							// line 361: imag = float(s)
							πF.SetLineno(361)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µimag = πTemp004
							// line 362: return complex(real, imag)
							πF.SetLineno(362)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µreal, "real"); πE != nil {
								continue
							}
							πTemp001[0] = µreal
							if πE = πg.CheckLocal(πF, µimag, "imag"); πE != nil {
								continue
							}
							πTemp001[1] = µimag
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_complex.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 363: dispatch[TYPE_COMPLEX] = load_complex
					πF.SetLineno(363)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßload_complex); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πTemp017); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_COMPLEX); πE != nil {
						continue
					}
					πTemp020 = πTemp021
					if πE = πg.SetItem(πF, πTemp019, πTemp020, πTemp018); πE != nil {
						continue
					}
					// line 365: def load_string(self):
					πF.SetLineno(365)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("load_string", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 366: n = self.r_long()
							πF.SetLineno(366)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µn = πTemp002
							// line 367: return self._read(n)
							πF.SetLineno(367)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp003[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßload_string.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 368: dispatch[TYPE_STRING] = load_string
					πF.SetLineno(368)
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßload_string); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πTemp018); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_STRING); πE != nil {
						continue
					}
					πTemp021 = πTemp022
					if πE = πg.SetItem(πF, πTemp020, πTemp021, πTemp019); πE != nil {
						continue
					}
					// line 370: def load_interned(self):
					πF.SetLineno(370)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("load_interned", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µret *πg.Object = πg.UnboundLocal
						_ = µret
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 371: n = self.r_long()
							πF.SetLineno(371)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µn = πTemp002
							// line 372: ret = intern(self._read(n))
							πF.SetLineno(372)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp004[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßintern); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µret = πTemp002
							// line 373: self._stringtable.append(ret)
							πF.SetLineno(373)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
								continue
							}
							πTemp003[0] = µret
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_stringtable, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 374: return ret
							πF.SetLineno(374)
							if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
								continue
							}
							πR = µret
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßload_interned.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 375: dispatch[TYPE_INTERNED] = load_interned
					πF.SetLineno(375)
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßload_interned); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πTemp019); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_INTERNED); πE != nil {
						continue
					}
					πTemp022 = πTemp023
					if πE = πg.SetItem(πF, πTemp021, πTemp022, πTemp020); πE != nil {
						continue
					}
					// line 377: def load_stringref(self):
					πF.SetLineno(377)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("load_stringref", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 378: n = self.r_long()
							πF.SetLineno(378)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µn = πTemp002
							// line 379: return self._stringtable[n]
							πF.SetLineno(379)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001 = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_stringtable, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_stringref.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 380: dispatch[TYPE_STRINGREF] = load_stringref
					πF.SetLineno(380)
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßload_stringref); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πTemp020); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_STRINGREF); πE != nil {
						continue
					}
					πTemp023 = πTemp024
					if πE = πg.SetItem(πF, πTemp022, πTemp023, πTemp021); πE != nil {
						continue
					}
					// line 382: def load_unicode(self):
					πF.SetLineno(382)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("load_unicode", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var µret *πg.Object = πg.UnboundLocal
						_ = µret
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 383: n = self.r_long()
							πF.SetLineno(383)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µn = πTemp002
							// line 384: s = self._read(n)
							πF.SetLineno(384)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp003[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_read, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µs = πTemp002
							// line 385: ret = s.decode('utf8')
							πF.SetLineno(385)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = ßutf8.ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µs, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µret = πTemp002
							// line 386: return ret
							πF.SetLineno(386)
							if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
								continue
							}
							πR = µret
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßload_unicode.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 387: dispatch[TYPE_UNICODE] = load_unicode
					πF.SetLineno(387)
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßload_unicode); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πTemp021); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_UNICODE); πE != nil {
						continue
					}
					πTemp024 = πTemp025
					if πE = πg.SetItem(πF, πTemp023, πTemp024, πTemp022); πE != nil {
						continue
					}
					// line 389: def load_tuple(self):
					πF.SetLineno(389)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("load_tuple", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 390: return tuple(self.load_list())
							πF.SetLineno(390)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
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
					if πE = πClass.SetItem(πF, ßload_tuple.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 391: dispatch[TYPE_TUPLE] = load_tuple
					πF.SetLineno(391)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßload_tuple); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πTemp022); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_TUPLE); πE != nil {
						continue
					}
					πTemp025 = πTemp026
					if πE = πg.SetItem(πF, πTemp024, πTemp025, πTemp023); πE != nil {
						continue
					}
					// line 393: def load_list(self):
					πF.SetLineno(393)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("load_list", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µlist *πg.Object = πg.UnboundLocal
						_ = µlist
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 394: n = self.r_long()
							πF.SetLineno(394)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µn = πTemp002
							// line 395: list = [self.load() for i in range(n)]
							πF.SetLineno(395)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
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
								var πR *πg.Object
								_ = πR
								var πE *πg.BaseException
								_ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1:
											goto Label1
										case 2:
											goto Label2
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
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
										// line 395: list = [self.load() for i in range(n)]
										πF.SetLineno(395)
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp004, nil
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µlist = πTemp001
							// line 396: return list
							πF.SetLineno(396)
							if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
								continue
							}
							πR = µlist
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßload_list.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 397: dispatch[TYPE_LIST] = load_list
					πF.SetLineno(397)
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßload_list); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πTemp023); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_LIST); πE != nil {
						continue
					}
					πTemp026 = πTemp027
					if πE = πg.SetItem(πF, πTemp025, πTemp026, πTemp024); πE != nil {
						continue
					}
					// line 399: def load_dict(self):
					πF.SetLineno(399)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("load_dict", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 400: d = {}
							πF.SetLineno(400)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 401: while 1:
							πF.SetLineno(401)
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
							// line 402: key = self.load()
							πF.SetLineno(402)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µkey = πTemp005
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_NULL); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µkey == πTemp005).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 403: if key is _NULL:
							πF.SetLineno(403)
						Label4:
							// line 404: break
							πF.SetLineno(404)
							πTemp003 = true
							continue
							goto Label5
						Label5:
							// line 405: value = self.load()
							πF.SetLineno(405)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp005
							// line 406: d[key] = value
							πF.SetLineno(406)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp005 = µkey
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp002); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 407: return d
							πF.SetLineno(407)
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
					if πE = πClass.SetItem(πF, ßload_dict.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 408: dispatch[TYPE_DICT] = load_dict
					πF.SetLineno(408)
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßload_dict); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πTemp024); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_DICT); πE != nil {
						continue
					}
					πTemp027 = πTemp028
					if πE = πg.SetItem(πF, πTemp026, πTemp027, πTemp025); πE != nil {
						continue
					}
					// line 410: def load_code(self):
					πF.SetLineno(410)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("load_code", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargcount *πg.Object = πg.UnboundLocal
						_ = µargcount
						var µnlocals *πg.Object = πg.UnboundLocal
						_ = µnlocals
						var µstacksize *πg.Object = πg.UnboundLocal
						_ = µstacksize
						var µflags *πg.Object = πg.UnboundLocal
						_ = µflags
						var µcode *πg.Object = πg.UnboundLocal
						_ = µcode
						var µconsts *πg.Object = πg.UnboundLocal
						_ = µconsts
						var µnames *πg.Object = πg.UnboundLocal
						_ = µnames
						var µvarnames *πg.Object = πg.UnboundLocal
						_ = µvarnames
						var µfreevars *πg.Object = πg.UnboundLocal
						_ = µfreevars
						var µcellvars *πg.Object = πg.UnboundLocal
						_ = µcellvars
						var µfilename *πg.Object = πg.UnboundLocal
						_ = µfilename
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µfirstlineno *πg.Object = πg.UnboundLocal
						_ = µfirstlineno
						var µlnotab *πg.Object = πg.UnboundLocal
						_ = µlnotab
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 411: argcount = self.r_long()
							πF.SetLineno(411)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µargcount = πTemp002
							// line 412: nlocals = self.r_long()
							πF.SetLineno(412)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnlocals = πTemp002
							// line 413: stacksize = self.r_long()
							πF.SetLineno(413)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstacksize = πTemp002
							// line 414: flags = self.r_long()
							πF.SetLineno(414)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µflags = πTemp002
							// line 415: code = self.load()
							πF.SetLineno(415)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcode = πTemp002
							// line 416: consts = self.load()
							πF.SetLineno(416)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µconsts = πTemp002
							// line 417: names = self.load()
							πF.SetLineno(417)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnames = πTemp002
							// line 418: varnames = self.load()
							πF.SetLineno(418)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvarnames = πTemp002
							// line 419: freevars = self.load()
							πF.SetLineno(419)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfreevars = πTemp002
							// line 420: cellvars = self.load()
							πF.SetLineno(420)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcellvars = πTemp002
							// line 421: filename = self.load()
							πF.SetLineno(421)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfilename = πTemp002
							// line 422: name = self.load()
							πF.SetLineno(422)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µname = πTemp002
							// line 423: firstlineno = self.r_long()
							πF.SetLineno(423)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfirstlineno = πTemp002
							// line 424: lnotab = self.load()
							πF.SetLineno(424)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlnotab = πTemp002
							// line 425: return types.CodeType(argcount, nlocals, stacksize, flags, code, consts,
							πF.SetLineno(425)
							πTemp003 = πF.MakeArgs(14)
							if πE = πg.CheckLocal(πF, µargcount, "argcount"); πE != nil {
								continue
							}
							πTemp003[0] = µargcount
							if πE = πg.CheckLocal(πF, µnlocals, "nlocals"); πE != nil {
								continue
							}
							πTemp003[1] = µnlocals
							if πE = πg.CheckLocal(πF, µstacksize, "stacksize"); πE != nil {
								continue
							}
							πTemp003[2] = µstacksize
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp003[3] = µflags
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp003[4] = µcode
							if πE = πg.CheckLocal(πF, µconsts, "consts"); πE != nil {
								continue
							}
							πTemp003[5] = µconsts
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							πTemp003[6] = µnames
							if πE = πg.CheckLocal(πF, µvarnames, "varnames"); πE != nil {
								continue
							}
							πTemp003[7] = µvarnames
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp003[8] = µfilename
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp003[9] = µname
							if πE = πg.CheckLocal(πF, µfirstlineno, "firstlineno"); πE != nil {
								continue
							}
							πTemp003[10] = µfirstlineno
							if πE = πg.CheckLocal(πF, µlnotab, "lnotab"); πE != nil {
								continue
							}
							πTemp003[11] = µlnotab
							if πE = πg.CheckLocal(πF, µfreevars, "freevars"); πE != nil {
								continue
							}
							πTemp003[12] = µfreevars
							if πE = πg.CheckLocal(πF, µcellvars, "cellvars"); πE != nil {
								continue
							}
							πTemp003[13] = µcellvars
							if πTemp001, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßCodeType, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßload_code.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 428: dispatch[TYPE_CODE] = load_code
					πF.SetLineno(428)
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßload_code); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πTemp025); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_CODE); πE != nil {
						continue
					}
					πTemp028 = πTemp029
					if πE = πg.SetItem(πF, πTemp027, πTemp028, πTemp026); πE != nil {
						continue
					}
					// line 430: def load_set(self):
					πF.SetLineno(430)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("load_set", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 431: n = self.r_long()
							πF.SetLineno(431)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µn = πTemp002
							// line 432: args = [self.load() for i in range(n)]
							πF.SetLineno(432)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
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
								var πR *πg.Object
								_ = πR
								var πE *πg.BaseException
								_ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1:
											goto Label1
										case 2:
											goto Label2
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
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
										// line 432: args = [self.load() for i in range(n)]
										πF.SetLineno(432)
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp004, nil
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µargs = πTemp001
							// line 433: return set(args)
							πF.SetLineno(433)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp005[0] = µargs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßload_set.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 434: dispatch[TYPE_SET] = load_set
					πF.SetLineno(434)
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßload_set); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πTemp026); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_SET); πE != nil {
						continue
					}
					πTemp029 = πTemp030
					if πE = πg.SetItem(πF, πTemp028, πTemp029, πTemp027); πE != nil {
						continue
					}
					// line 436: def load_frozenset(self):
					πF.SetLineno(436)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("load_frozenset", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []πg.Param
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 437: n = self.r_long()
							πF.SetLineno(437)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßr_long, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µn = πTemp002
							// line 438: args = [self.load() for i in range(n)]
							πF.SetLineno(438)
							πTemp003 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
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
								var πR *πg.Object
								_ = πR
								var πE *πg.BaseException
								_ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1:
											goto Label1
										case 2:
											goto Label2
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
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
										// line 438: args = [self.load() for i in range(n)]
										πF.SetLineno(438)
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp004, nil
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
							if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
								continue
							}
							µargs = πTemp001
							// line 439: return frozenset(args)
							πF.SetLineno(439)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp005[0] = µargs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
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
					if πE = πClass.SetItem(πF, ßload_frozenset.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 440: dispatch[TYPE_FROZENSET] = load_frozenset
					πF.SetLineno(440)
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßload_frozenset); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp027); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_FROZENSET); πE != nil {
						continue
					}
					πTemp030 = πTemp031
					if πE = πg.SetItem(πF, πTemp029, πTemp030, πTemp028); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp009, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("_Unmarshaller").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Unmarshaller.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 444: def _read(self, n):
			πF.SetLineno(444)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "self", Def: nil}
			πTemp007[1] = πg.Param{Name: "n", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_read", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]
				_ = µself
				var µn *πg.Object = πArgs[1]
				_ = µn
				var µpos *πg.Object = πg.UnboundLocal
				_ = µpos
				var µnewpos *πg.Object = πg.UnboundLocal
				_ = µnewpos
				var µret *πg.Object = πg.UnboundLocal
				_ = µret
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
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 445: pos = self.bufpos
					πF.SetLineno(445)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µself, ßbufpos, nil); πE != nil {
						continue
					}
					µpos = πTemp001
					// line 446: newpos = pos + n
					πF.SetLineno(446)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µpos, µn); πE != nil {
						continue
					}
					µnewpos = πTemp001
					if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µself, ßbufstr, nil); πE != nil {
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
					if πTemp001, πE = πg.GT(πF, µnewpos, πTemp004); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 447: if newpos > len(self.bufstr): raise EOFError
					πF.SetLineno(447)
				Label1:
					if πTemp001, πE = πg.ResolveGlobal(πF, ßEOFError); πE != nil {
						continue
					}
					// line 447: if newpos > len(self.bufstr): raise EOFError
					πF.SetLineno(447)
					πE = πF.Raise(πTemp001, nil, nil)
					continue
					goto Label2
				Label2:
					// line 448: ret = self.bufstr[pos : newpos]
					πF.SetLineno(448)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µpos, µnewpos, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µself, ßbufstr, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					µret = πTemp003
					// line 449: self.bufpos = newpos
					πF.SetLineno(449)
					if πE = πg.CheckLocal(πF, µnewpos, "newpos"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µnewpos); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µself, ßbufpos, πTemp001); πE != nil {
						continue
					}
					// line 450: return ret
					πF.SetLineno(450)
					if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
						continue
					}
					πR = µret
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_read.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 452: def _read1(self):
			πF.SetLineno(452)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "self", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("_read1", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]
				_ = µself
				var µret *πg.Object = πg.UnboundLocal
				_ = µret
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 453: ret = self.bufstr[self.bufpos]
					πF.SetLineno(453)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µself, ßbufpos, nil); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µself, ßbufstr, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					µret = πTemp002
					// line 454: self.bufpos += 1
					πF.SetLineno(454)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µself, ßbufpos, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µself, ßbufpos, πTemp002); πE != nil {
						continue
					}
					// line 455: return ret
					πF.SetLineno(455)
					if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
						continue
					}
					πR = µret
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_read1.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 457: def _r_short(self):
			πF.SetLineno(457)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "self", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("_r_short", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]
				_ = µself
				var µlo *πg.Object = πg.UnboundLocal
				_ = µlo
				var µhi *πg.Object = πg.UnboundLocal
				_ = µhi
				var µx *πg.Object = πg.UnboundLocal
				_ = µx
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
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 458: lo = ord(_read1(self))
					πF.SetLineno(458)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µlo = πTemp004
					// line 459: hi = ord(_read1(self))
					πF.SetLineno(459)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µhi = πTemp004
					// line 460: x = lo | (hi<<8)
					πF.SetLineno(460)
					if πE = πg.CheckLocal(πF, µlo, "lo"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µhi, "hi"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LShift(πF, µhi, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Or(πF, µlo, πTemp004); πE != nil {
						continue
					}
					µx = πTemp003
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µx, πg.NewInt(32768).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 461: if x & 0x8000:
					πF.SetLineno(461)
				Label1:
					// line 462: x = x - 0x10000
					πF.SetLineno(462)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, µx, πg.NewInt(65536).ToObject()); πE != nil {
						continue
					}
					µx = πTemp003
					goto Label2
				Label2:
					// line 463: return x
					πF.SetLineno(463)
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
			if πE = πF.Globals().SetItem(πF, ß_r_short.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 465: def _r_long(self):
			πF.SetLineno(465)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "self", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("_r_long", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]
				_ = µself
				var µp *πg.Object = πg.UnboundLocal
				_ = µp
				var µs *πg.Object = πg.UnboundLocal
				_ = µs
				var µa *πg.Object = πg.UnboundLocal
				_ = µa
				var µb *πg.Object = πg.UnboundLocal
				_ = µb
				var µc *πg.Object = πg.UnboundLocal
				_ = µc
				var µd *πg.Object = πg.UnboundLocal
				_ = µd
				var µx *πg.Object = πg.UnboundLocal
				_ = µx
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
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 467: p = self.bufpos
					πF.SetLineno(467)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µself, ßbufpos, nil); πE != nil {
						continue
					}
					µp = πTemp001
					// line 468: s = self.bufstr
					πF.SetLineno(468)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µself, ßbufstr, nil); πE != nil {
						continue
					}
					µs = πTemp001
					// line 469: a = ord(s[p])
					πF.SetLineno(469)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					πTemp001 = µp
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µa = πTemp003
					// line 470: b = ord(s[p+1])
					πF.SetLineno(470)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µp, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µb = πTemp003
					// line 471: c = ord(s[p+2])
					πF.SetLineno(471)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µp, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µc = πTemp003
					// line 472: d = ord(s[p+3])
					πF.SetLineno(472)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µp, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp001, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µd = πTemp003
					// line 473: self.bufpos += 4
					πF.SetLineno(473)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µself, ßbufpos, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.IAdd(πF, πTemp001, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, µself, ßbufpos, πTemp003); πE != nil {
						continue
					}
					// line 474: x = a | (b<<8) | (c<<16) | (d<<24)
					πF.SetLineno(474)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LShift(πF, µb, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Or(πF, µa, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LShift(πF, µc, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Or(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.LShift(πF, µd, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Or(πF, πTemp003, πTemp004); πE != nil {
						continue
					}
					µx = πTemp001
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.And(πF, µd, πg.NewInt(128).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GT(πF, µx, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
				Label1:
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label2
					}
					goto Label3
					// line 475: if d & 0x80 and x > 0:
					πF.SetLineno(475)
				Label2:
					// line 476: x = -((1<<32) - x)
					πF.SetLineno(476)
					if πTemp004, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Sub(πF, πTemp004, µx); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Neg(πF, πTemp003); πE != nil {
						continue
					}
					µx = πTemp001
					// line 477: return int(x)
					πF.SetLineno(477)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp002[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πR = πTemp003
					continue
					goto Label4
				Label3:
					// line 479: return x
					πF.SetLineno(479)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πR = µx
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
			if πE = πF.Globals().SetItem(πF, ß_r_long.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 481: def _r_long64(self):
			πF.SetLineno(481)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "self", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("_r_long64", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µself *πg.Object = πArgs[0]
				_ = µself
				var µa *πg.Object = πg.UnboundLocal
				_ = µa
				var µb *πg.Object = πg.UnboundLocal
				_ = µb
				var µc *πg.Object = πg.UnboundLocal
				_ = µc
				var µd *πg.Object = πg.UnboundLocal
				_ = µd
				var µe *πg.Object = πg.UnboundLocal
				_ = µe
				var µf *πg.Object = πg.UnboundLocal
				_ = µf
				var µg *πg.Object = πg.UnboundLocal
				_ = µg
				var µh *πg.Object = πg.UnboundLocal
				_ = µh
				var µx *πg.Object = πg.UnboundLocal
				_ = µx
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 482: a = ord(_read1(self))
					πF.SetLineno(482)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µa = πTemp004
					// line 483: b = ord(_read1(self))
					πF.SetLineno(483)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µb = πTemp004
					// line 484: c = ord(_read1(self))
					πF.SetLineno(484)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µc = πTemp004
					// line 485: d = ord(_read1(self))
					πF.SetLineno(485)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µd = πTemp004
					// line 486: e = ord(_read1(self))
					πF.SetLineno(486)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µe = πTemp004
					// line 487: f = ord(_read1(self))
					πF.SetLineno(487)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µf = πTemp004
					// line 488: g = ord(_read1(self))
					πF.SetLineno(488)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µg = πTemp004
					// line 489: h = ord(_read1(self))
					πF.SetLineno(489)
					πTemp001 = πF.MakeArgs(1)
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
						continue
					}
					πTemp002[0] = µself
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					πTemp001[0] = πTemp004
					if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µh = πTemp004
					// line 490: x = a | (b<<8) | (c<<16) | (d<<24)
					πF.SetLineno(490)
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.LShift(πF, µb, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Or(πF, µa, πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.LShift(πF, µc, πg.NewInt(16).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Or(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LShift(πF, µd, πg.NewInt(24).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Or(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					µx = πTemp003
					// line 491: x = x | (e<<32) | (f<<40) | (g<<48) | (h<<56)
					πF.SetLineno(491)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.LShift(πF, µe, πg.NewInt(32).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Or(πF, µx, πTemp007); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.LShift(πF, µf, πg.NewInt(40).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Or(πF, πTemp006, πTemp007); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µg, "g"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.LShift(πF, µg, πg.NewInt(48).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Or(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.LShift(πF, µh, πg.NewInt(56).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Or(πF, πTemp004, πTemp005); πE != nil {
						continue
					}
					µx = πTemp003
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.And(πF, µh, πg.NewInt(128).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp008 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GT(πF, µx, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp004
				Label1:
					if πTemp008, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label2
					}
					goto Label3
					// line 492: if h & 0x80 and x > 0:
					πF.SetLineno(492)
				Label2:
					// line 493: x = -((1<<64) - x)
					πF.SetLineno(493)
					if πTemp005, πE = πg.LShift(πF, πg.NewInt(1).ToObject(), πg.NewInt(64).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, πTemp005, µx); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Neg(πF, πTemp004); πE != nil {
						continue
					}
					µx = πTemp003
					goto Label3
				Label3:
					// line 494: return x
					πF.SetLineno(494)
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
			if πE = πF.Globals().SetItem(πF, ß_r_long64.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 496: _load_dispatch = {}
			πF.SetLineno(496)
			πTemp008 = πg.NewDict()
			πTemp012 = πTemp008.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_load_dispatch.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 498: class _FastUnmarshaller(object):
			πF.SetLineno(498)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp014
			πTemp008 = πg.NewDict()
			if πTemp012, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp012); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_FastUnmarshaller", "/usr/lib/python2.7/marshal.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
				_ = πClass
				var πTemp001 *πg.Dict
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 499: dispatch = {}
					πF.SetLineno(499)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					if πE = πClass.SetItem(πF, ßdispatch.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 501: def __init__(self, buffer):
					πF.SetLineno(501)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp003[1] = πg.Param{Name: "buffer", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µbuffer *πg.Object = πArgs[1]
						_ = µbuffer
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 502: self.bufstr = buffer
							πF.SetLineno(502)
							if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µbuffer); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbufstr, πTemp001); πE != nil {
								continue
							}
							// line 503: self.bufpos = 0
							πF.SetLineno(503)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbufpos, πTemp001); πE != nil {
								continue
							}
							// line 504: self._stringtable = []
							πF.SetLineno(504)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_stringtable, πTemp003); πE != nil {
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
					// line 506: def load(self):
					πF.SetLineno(506)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("load", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µc *πg.Object = πg.UnboundLocal
						_ = µc
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
						var πTemp009 *πg.Object
						_ = πTemp009
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 508: c = '?'
							πF.SetLineno(508)
							µc = πg.NewStr("?").ToObject()
							// line 509: try:
							πF.SetLineno(509)
							πF.PushCheckpoint(2)
							// line 510: c = self.bufstr[self.bufpos]
							πF.SetLineno(510)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbufpos, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßbufstr, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µc = πTemp002
							// line 511: self.bufpos += 1
							πF.SetLineno(511)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbufpos, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbufpos, πTemp002); πE != nil {
								continue
							}
							// line 512: return _load_dispatch[c](self)
							πF.SetLineno(512)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp001 = µc
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_load_dispatch); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label4
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 513: except KeyError:
							πF.SetLineno(513)
						Label3:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp008[0] = µc
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002 = πg.NewTuple2(µc, πTemp009).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("bad marshal code: %c (%d)").ToObject(), πTemp002); πE != nil {
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
							// line 514: raise ValueError("bad marshal code: %c (%d)" % (c, ord(c)))
							πF.SetLineno(514)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
							// line 515: except IndexError:
							πF.SetLineno(515)
						Label4:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEOFError); πE != nil {
								continue
							}
							// line 516: raise EOFError
							πF.SetLineno(516)
							πE = πF.Raise(πTemp001, nil, nil)
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
					if πE = πClass.SetItem(πF, ßload.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 518: def load_null(self):
					πF.SetLineno(518)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("load_null", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 519: return _NULL
							πF.SetLineno(519)
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_NULL); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_null.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 520: dispatch[TYPE_NULL] = load_null
					πF.SetLineno(520)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßload_null); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp006); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_NULL); πE != nil {
						continue
					}
					πTemp009 = πTemp010
					if πE = πg.SetItem(πF, πTemp008, πTemp009, πTemp007); πE != nil {
						continue
					}
					// line 522: def load_none(self):
					πF.SetLineno(522)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("load_none", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 523: return None
							πF.SetLineno(523)
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
					if πE = πClass.SetItem(πF, ßload_none.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 524: dispatch[TYPE_NONE] = load_none
					πF.SetLineno(524)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßload_none); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp007); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_NONE); πE != nil {
						continue
					}
					πTemp010 = πTemp011
					if πE = πg.SetItem(πF, πTemp009, πTemp010, πTemp008); πE != nil {
						continue
					}
					// line 526: def load_true(self):
					πF.SetLineno(526)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("load_true", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 527: return True
							πF.SetLineno(527)
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
					if πE = πClass.SetItem(πF, ßload_true.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 528: dispatch[TYPE_TRUE] = load_true
					πF.SetLineno(528)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßload_true); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp008); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_TRUE); πE != nil {
						continue
					}
					πTemp011 = πTemp012
					if πE = πg.SetItem(πF, πTemp010, πTemp011, πTemp009); πE != nil {
						continue
					}
					// line 530: def load_false(self):
					πF.SetLineno(530)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("load_false", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 531: return False
							πF.SetLineno(531)
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
					if πE = πClass.SetItem(πF, ßload_false.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 532: dispatch[TYPE_FALSE] = load_false
					πF.SetLineno(532)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßload_false); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_FALSE); πE != nil {
						continue
					}
					πTemp012 = πTemp013
					if πE = πg.SetItem(πF, πTemp011, πTemp012, πTemp010); πE != nil {
						continue
					}
					// line 534: def load_stopiter(self):
					πF.SetLineno(534)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("load_stopiter", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 535: return StopIteration
							πF.SetLineno(535)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_stopiter.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 536: dispatch[TYPE_STOPITER] = load_stopiter
					πF.SetLineno(536)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßload_stopiter); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp010); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_STOPITER); πE != nil {
						continue
					}
					πTemp013 = πTemp014
					if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
						continue
					}
					// line 538: def load_ellipsis(self):
					πF.SetLineno(538)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("load_ellipsis", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 539: return Ellipsis
							πF.SetLineno(539)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEllipsis); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_ellipsis.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 540: dispatch[TYPE_ELLIPSIS] = load_ellipsis
					πF.SetLineno(540)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßload_ellipsis); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πTemp011); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_ELLIPSIS); πE != nil {
						continue
					}
					πTemp014 = πTemp015
					if πE = πg.SetItem(πF, πTemp013, πTemp014, πTemp012); πE != nil {
						continue
					}
					// line 542: def load_int(self):
					πF.SetLineno(542)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("load_int", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 543: return _r_long(self)
							πF.SetLineno(543)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_int.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 544: dispatch[TYPE_INT] = load_int
					πF.SetLineno(544)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßload_int); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πTemp012); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_INT); πE != nil {
						continue
					}
					πTemp015 = πTemp016
					if πE = πg.SetItem(πF, πTemp014, πTemp015, πTemp013); πE != nil {
						continue
					}
					// line 546: def load_int64(self):
					πF.SetLineno(546)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("load_int64", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 547: return _r_long64(self)
							πF.SetLineno(547)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long64); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_int64.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 548: dispatch[TYPE_INT64] = load_int64
					πF.SetLineno(548)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßload_int64); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_INT64); πE != nil {
						continue
					}
					πTemp016 = πTemp017
					if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
						continue
					}
					// line 550: def load_long(self):
					πF.SetLineno(550)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("load_long", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µsize *πg.Object = πg.UnboundLocal
						_ = µsize
						var µsign *πg.Object = πg.UnboundLocal
						_ = µsign
						var µx *πg.Object = πg.UnboundLocal
						_ = µx
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 551: size = _r_long(self)
							πF.SetLineno(551)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µsize = πTemp003
							// line 552: sign = 1
							πF.SetLineno(552)
							µsign = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µsize, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 553: if size < 0:
							πF.SetLineno(553)
						Label1:
							// line 554: sign = -1
							πF.SetLineno(554)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µsign = πTemp002
							// line 555: size = -size
							πF.SetLineno(555)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Neg(πF, µsize); πE != nil {
								continue
							}
							µsize = πTemp002
							goto Label2
						Label2:
							// line 556: x = 0
							πF.SetLineno(556)
							µx = πg.NewInt(0).ToObject()
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsize, "size"); πE != nil {
								continue
							}
							πTemp001[0] = µsize
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
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
								µi = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 558: d = _r_short(self)
							πF.SetLineno(558)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_r_short); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µd = πTemp005
							// line 559: x = x | (d<<(i*15))
							πF.SetLineno(559)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Mul(πF, µi, πg.NewInt(15).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.LShift(πF, µd, πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Or(πF, µx, πTemp005); πE != nil {
								continue
							}
							µx = πTemp003
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 560: return x * sign
							πF.SetLineno(560)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µsign, "sign"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, µx, µsign); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_long.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 561: dispatch[TYPE_LONG] = load_long
					πF.SetLineno(561)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßload_long); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πTemp014); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_LONG); πE != nil {
						continue
					}
					πTemp017 = πTemp018
					if πE = πg.SetItem(πF, πTemp016, πTemp017, πTemp015); πE != nil {
						continue
					}
					// line 563: def load_float(self):
					πF.SetLineno(563)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("load_float", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 564: n = ord(_read1(self))
							πF.SetLineno(564)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp004
							// line 565: s = _read(self, n)
							πF.SetLineno(565)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[1] = µn
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_read); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp004
							// line 566: return float(s)
							πF.SetLineno(566)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_float.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 567: dispatch[TYPE_FLOAT] = load_float
					πF.SetLineno(567)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßload_float); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp015); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_FLOAT); πE != nil {
						continue
					}
					πTemp018 = πTemp019
					if πE = πg.SetItem(πF, πTemp017, πTemp018, πTemp016); πE != nil {
						continue
					}
					// line 569: def load_complex(self):
					πF.SetLineno(569)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("load_complex", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var µreal *πg.Object = πg.UnboundLocal
						_ = µreal
						var µimag *πg.Object = πg.UnboundLocal
						_ = µimag
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 570: n = ord(_read1(self))
							πF.SetLineno(570)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp004
							// line 571: s = _read(self, n)
							πF.SetLineno(571)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[1] = µn
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_read); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp004
							// line 572: real = float(s)
							πF.SetLineno(572)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µreal = πTemp004
							// line 573: n = ord(_read1(self))
							πF.SetLineno(573)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[0] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_read1); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp004
							// line 574: s = _read(self, n)
							πF.SetLineno(574)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[1] = µn
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_read); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp004
							// line 575: imag = float(s)
							πF.SetLineno(575)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µimag = πTemp004
							// line 576: return complex(real, imag)
							πF.SetLineno(576)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µreal, "real"); πE != nil {
								continue
							}
							πTemp001[0] = µreal
							if πE = πg.CheckLocal(πF, µimag, "imag"); πE != nil {
								continue
							}
							πTemp001[1] = µimag
							if πTemp003, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_complex.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 577: dispatch[TYPE_COMPLEX] = load_complex
					πF.SetLineno(577)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßload_complex); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πTemp016); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_COMPLEX); πE != nil {
						continue
					}
					πTemp019 = πTemp020
					if πE = πg.SetItem(πF, πTemp018, πTemp019, πTemp017); πE != nil {
						continue
					}
					// line 579: def load_string(self):
					πF.SetLineno(579)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("load_string", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 580: n = _r_long(self)
							πF.SetLineno(580)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							// line 581: return _read(self, n)
							πF.SetLineno(581)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[1] = µn
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_read); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_string.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 582: dispatch[TYPE_STRING] = load_string
					πF.SetLineno(582)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßload_string); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πTemp017); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_STRING); πE != nil {
						continue
					}
					πTemp020 = πTemp021
					if πE = πg.SetItem(πF, πTemp019, πTemp020, πTemp018); πE != nil {
						continue
					}
					// line 584: def load_interned(self):
					πF.SetLineno(584)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("load_interned", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µret *πg.Object = πg.UnboundLocal
						_ = µret
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 585: n = _r_long(self)
							πF.SetLineno(585)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							// line 586: ret = intern(_read(self, n))
							πF.SetLineno(586)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp004[0] = µself
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp004[1] = µn
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_read); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßintern); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µret = πTemp003
							// line 587: self._stringtable.append(ret)
							πF.SetLineno(587)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
								continue
							}
							πTemp001[0] = µret
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_stringtable, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 588: return ret
							πF.SetLineno(588)
							if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
								continue
							}
							πR = µret
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßload_interned.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 589: dispatch[TYPE_INTERNED] = load_interned
					πF.SetLineno(589)
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßload_interned); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πTemp018); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_INTERNED); πE != nil {
						continue
					}
					πTemp021 = πTemp022
					if πE = πg.SetItem(πF, πTemp020, πTemp021, πTemp019); πE != nil {
						continue
					}
					// line 591: def load_stringref(self):
					πF.SetLineno(591)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("load_stringref", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 592: n = _r_long(self)
							πF.SetLineno(592)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							// line 593: return self._stringtable[n]
							πF.SetLineno(593)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp002 = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ß_stringtable, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_stringref.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 594: dispatch[TYPE_STRINGREF] = load_stringref
					πF.SetLineno(594)
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßload_stringref); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πTemp019); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_STRINGREF); πE != nil {
						continue
					}
					πTemp022 = πTemp023
					if πE = πg.SetItem(πF, πTemp021, πTemp022, πTemp020); πE != nil {
						continue
					}
					// line 596: def load_unicode(self):
					πF.SetLineno(596)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("load_unicode", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
						var µret *πg.Object = πg.UnboundLocal
						_ = µret
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 597: n = _r_long(self)
							πF.SetLineno(597)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							// line 598: s = _read(self, n)
							πF.SetLineno(598)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[1] = µn
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_read); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µs = πTemp003
							// line 599: ret = s.decode('utf8')
							πF.SetLineno(599)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = ßutf8.ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µs, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µret = πTemp003
							// line 600: return ret
							πF.SetLineno(600)
							if πE = πg.CheckLocal(πF, µret, "ret"); πE != nil {
								continue
							}
							πR = µret
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßload_unicode.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 601: dispatch[TYPE_UNICODE] = load_unicode
					πF.SetLineno(601)
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßload_unicode); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πTemp020); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_UNICODE); πE != nil {
						continue
					}
					πTemp023 = πTemp024
					if πE = πg.SetItem(πF, πTemp022, πTemp023, πTemp021); πE != nil {
						continue
					}
					// line 603: def load_tuple(self):
					πF.SetLineno(603)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("load_tuple", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 604: return tuple(self.load_list())
							πF.SetLineno(604)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload_list, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
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
					if πE = πClass.SetItem(πF, ßload_tuple.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 605: dispatch[TYPE_TUPLE] = load_tuple
					πF.SetLineno(605)
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßload_tuple); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πTemp021); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_TUPLE); πE != nil {
						continue
					}
					πTemp024 = πTemp025
					if πE = πg.SetItem(πF, πTemp023, πTemp024, πTemp022); πE != nil {
						continue
					}
					// line 607: def load_list(self):
					πF.SetLineno(607)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("load_list", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µlist *πg.Object = πg.UnboundLocal
						_ = µlist
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 608: n = _r_long(self)
							πF.SetLineno(608)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							// line 609: list = []
							πF.SetLineno(609)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µlist = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
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
								µi = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 611: list.append(self.load())
							πF.SetLineno(611)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µlist, ßappend, nil); πE != nil {
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
							// line 612: return list
							πF.SetLineno(612)
							if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
								continue
							}
							πR = µlist
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßload_list.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 613: dispatch[TYPE_LIST] = load_list
					πF.SetLineno(613)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßload_list); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πTemp022); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_LIST); πE != nil {
						continue
					}
					πTemp025 = πTemp026
					if πE = πg.SetItem(πF, πTemp024, πTemp025, πTemp023); πE != nil {
						continue
					}
					// line 615: def load_dict(self):
					πF.SetLineno(615)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("load_dict", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var πTemp001 *πg.Dict
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 616: d = {}
							πF.SetLineno(616)
							πTemp001 = πg.NewDict()
							πTemp002 = πTemp001.ToObject()
							µd = πTemp002
							// line 617: while 1:
							πF.SetLineno(617)
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
							// line 618: key = self.load()
							πF.SetLineno(618)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µkey = πTemp005
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_NULL); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µkey == πTemp005).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 619: if key is _NULL:
							πF.SetLineno(619)
						Label4:
							// line 620: break
							πF.SetLineno(620)
							πTemp003 = true
							continue
							goto Label5
						Label5:
							// line 621: value = self.load()
							πF.SetLineno(621)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp005
							// line 622: d[key] = value
							πF.SetLineno(622)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp005 = µkey
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp002); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 623: return d
							πF.SetLineno(623)
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
					if πE = πClass.SetItem(πF, ßload_dict.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 624: dispatch[TYPE_DICT] = load_dict
					πF.SetLineno(624)
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßload_dict); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πTemp023); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_DICT); πE != nil {
						continue
					}
					πTemp026 = πTemp027
					if πE = πg.SetItem(πF, πTemp025, πTemp026, πTemp024); πE != nil {
						continue
					}
					// line 626: def load_code(self):
					πF.SetLineno(626)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("load_code", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargcount *πg.Object = πg.UnboundLocal
						_ = µargcount
						var µnlocals *πg.Object = πg.UnboundLocal
						_ = µnlocals
						var µstacksize *πg.Object = πg.UnboundLocal
						_ = µstacksize
						var µflags *πg.Object = πg.UnboundLocal
						_ = µflags
						var µcode *πg.Object = πg.UnboundLocal
						_ = µcode
						var µconsts *πg.Object = πg.UnboundLocal
						_ = µconsts
						var µnames *πg.Object = πg.UnboundLocal
						_ = µnames
						var µvarnames *πg.Object = πg.UnboundLocal
						_ = µvarnames
						var µfreevars *πg.Object = πg.UnboundLocal
						_ = µfreevars
						var µcellvars *πg.Object = πg.UnboundLocal
						_ = µcellvars
						var µfilename *πg.Object = πg.UnboundLocal
						_ = µfilename
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µfirstlineno *πg.Object = πg.UnboundLocal
						_ = µfirstlineno
						var µlnotab *πg.Object = πg.UnboundLocal
						_ = µlnotab
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 627: argcount = _r_long(self)
							πF.SetLineno(627)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µargcount = πTemp003
							// line 628: nlocals = _r_long(self)
							πF.SetLineno(628)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µnlocals = πTemp003
							// line 629: stacksize = _r_long(self)
							πF.SetLineno(629)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µstacksize = πTemp003
							// line 630: flags = _r_long(self)
							πF.SetLineno(630)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µflags = πTemp003
							// line 631: code = self.load()
							πF.SetLineno(631)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcode = πTemp003
							// line 632: consts = self.load()
							πF.SetLineno(632)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µconsts = πTemp003
							// line 633: names = self.load()
							πF.SetLineno(633)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnames = πTemp003
							// line 634: varnames = self.load()
							πF.SetLineno(634)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvarnames = πTemp003
							// line 635: freevars = self.load()
							πF.SetLineno(635)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfreevars = πTemp003
							// line 636: cellvars = self.load()
							πF.SetLineno(636)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µcellvars = πTemp003
							// line 637: filename = self.load()
							πF.SetLineno(637)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µfilename = πTemp003
							// line 638: name = self.load()
							πF.SetLineno(638)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µname = πTemp003
							// line 639: firstlineno = _r_long(self)
							πF.SetLineno(639)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfirstlineno = πTemp003
							// line 640: lnotab = self.load()
							πF.SetLineno(640)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µlnotab = πTemp003
							// line 641: return types.CodeType(argcount, nlocals, stacksize, flags, code, consts,
							πF.SetLineno(641)
							πTemp001 = πF.MakeArgs(14)
							if πE = πg.CheckLocal(πF, µargcount, "argcount"); πE != nil {
								continue
							}
							πTemp001[0] = µargcount
							if πE = πg.CheckLocal(πF, µnlocals, "nlocals"); πE != nil {
								continue
							}
							πTemp001[1] = µnlocals
							if πE = πg.CheckLocal(πF, µstacksize, "stacksize"); πE != nil {
								continue
							}
							πTemp001[2] = µstacksize
							if πE = πg.CheckLocal(πF, µflags, "flags"); πE != nil {
								continue
							}
							πTemp001[3] = µflags
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[4] = µcode
							if πE = πg.CheckLocal(πF, µconsts, "consts"); πE != nil {
								continue
							}
							πTemp001[5] = µconsts
							if πE = πg.CheckLocal(πF, µnames, "names"); πE != nil {
								continue
							}
							πTemp001[6] = µnames
							if πE = πg.CheckLocal(πF, µvarnames, "varnames"); πE != nil {
								continue
							}
							πTemp001[7] = µvarnames
							if πE = πg.CheckLocal(πF, µfilename, "filename"); πE != nil {
								continue
							}
							πTemp001[8] = µfilename
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[9] = µname
							if πE = πg.CheckLocal(πF, µfirstlineno, "firstlineno"); πE != nil {
								continue
							}
							πTemp001[10] = µfirstlineno
							if πE = πg.CheckLocal(πF, µlnotab, "lnotab"); πE != nil {
								continue
							}
							πTemp001[11] = µlnotab
							if πE = πg.CheckLocal(πF, µfreevars, "freevars"); πE != nil {
								continue
							}
							πTemp001[12] = µfreevars
							if πE = πg.CheckLocal(πF, µcellvars, "cellvars"); πE != nil {
								continue
							}
							πTemp001[13] = µcellvars
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtypes); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßCodeType, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_code.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 644: dispatch[TYPE_CODE] = load_code
					πF.SetLineno(644)
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßload_code); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πTemp024); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_CODE); πE != nil {
						continue
					}
					πTemp027 = πTemp028
					if πE = πg.SetItem(πF, πTemp026, πTemp027, πTemp025); πE != nil {
						continue
					}
					// line 646: def load_set(self):
					πF.SetLineno(646)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("load_set", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 647: n = _r_long(self)
							πF.SetLineno(647)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							// line 648: args = [self.load() for i in range(n)]
							πF.SetLineno(648)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/marshal.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
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
								var πR *πg.Object
								_ = πR
								var πE *πg.BaseException
								_ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1:
											goto Label1
										case 2:
											goto Label2
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
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
										// line 648: args = [self.load() for i in range(n)]
										πF.SetLineno(648)
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp004, nil
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							µargs = πTemp002
							// line 649: return set(args)
							πF.SetLineno(649)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßset); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_set.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 650: dispatch[TYPE_SET] = load_set
					πF.SetLineno(650)
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßload_set); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πTemp025); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_SET); πE != nil {
						continue
					}
					πTemp028 = πTemp029
					if πE = πg.SetItem(πF, πTemp027, πTemp028, πTemp026); πE != nil {
						continue
					}
					// line 652: def load_frozenset(self):
					πF.SetLineno(652)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("load_frozenset", "/usr/lib/python2.7/marshal.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
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
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default:
								panic("unexpected function state")
							}
							// line 653: n = _r_long(self)
							πF.SetLineno(653)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_r_long); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp003
							// line 654: args = [self.load() for i in range(n)]
							πF.SetLineno(654)
							πTemp004 = make([]πg.Param, 0)
							πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/marshal.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µi *πg.Object = πg.UnboundLocal
								_ = µi
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
								var πR *πg.Object
								_ = πR
								var πE *πg.BaseException
								_ = πE
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1:
											goto Label1
										case 2:
											goto Label2
										case 4:
											goto Label4
										default:
											panic("unexpected function state")
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
										// line 654: args = [self.load() for i in range(n)]
										πF.SetLineno(654)
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.GetAttr(πF, µself, ßload, nil); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
											continue
										}
										πF.PushCheckpoint(4)
										return πTemp004, nil
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
							if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							µargs = πTemp002
							// line 655: return frozenset(args)
							πF.SetLineno(655)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp001[0] = µargs
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfrozenset); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_frozenset.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 656: dispatch[TYPE_FROZENSET] = load_frozenset
					πF.SetLineno(656)
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßload_frozenset); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πTemp026); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßTYPE_FROZENSET); πE != nil {
						continue
					}
					πTemp029 = πTemp030
					if πE = πg.SetItem(πF, πTemp028, πTemp029, πTemp027); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp013, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp013 == nil {
				πTemp013 = πg.TypeType.ToObject()
			}
			if πTemp014, πE = πTemp013.Call(πF, []*πg.Object{πg.NewStr("_FastUnmarshaller").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_FastUnmarshaller.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 658: _load_dispatch = _FastUnmarshaller.dispatch
			πF.SetLineno(658)
			if πTemp012, πE = πg.ResolveGlobal(πF, ß_FastUnmarshaller); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßdispatch, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_load_dispatch.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 664: version = 1
			πF.SetLineno(664)
			if πE = πF.Globals().SetItem(πF, ßversion.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
				continue
			}
			// line 667: def dump(x, f, version=version):
			πF.SetLineno(667)
			πTemp007 = make([]πg.Param, 3)
			πTemp007[0] = πg.Param{Name: "x", Def: nil}
			πTemp007[1] = πg.Param{Name: "f", Def: nil}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßversion); πE != nil {
				continue
			}
			πTemp007[2] = πg.Param{Name: "version", Def: πTemp013}
			πTemp012 = πg.NewFunction(πg.NewCode("dump", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]
				_ = µx
				var µf *πg.Object = πArgs[1]
				_ = µf
				var µversion *πg.Object = πArgs[2]
				_ = µversion
				var µm *πg.Object = πg.UnboundLocal
				_ = µm
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 669: m = _Marshaller(f.write)
					πF.SetLineno(669)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßwrite, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_Marshaller); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µm = πTemp003
					// line 670: m.dump(x)
					πF.SetLineno(670)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µm, ßdump, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdump.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 666: @builtinify
			πF.SetLineno(666)
			πTemp002 = πF.MakeArgs(1)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßdump); πE != nil {
				continue
			}
			πTemp002[0] = πTemp013
			if πTemp013, πE = πg.ResolveGlobal(πF, ßbuiltinify); πE != nil {
				continue
			}
			if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßdump.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 673: def load(f):
			πF.SetLineno(673)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "f", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("load", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µf *πg.Object = πArgs[0]
				_ = µf
				var µum *πg.Object = πg.UnboundLocal
				_ = µum
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 674: um = _Unmarshaller(f.read)
					πF.SetLineno(674)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µf, ßread, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_Unmarshaller); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µum = πTemp003
					// line 675: return um.load()
					πF.SetLineno(675)
					if πE = πg.CheckLocal(πF, µum, "um"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µum, ßload, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßload.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 672: @builtinify
			πF.SetLineno(672)
			πTemp002 = πF.MakeArgs(1)
			if πTemp014, πE = πg.ResolveGlobal(πF, ßload); πE != nil {
				continue
			}
			πTemp002[0] = πTemp014
			if πTemp014, πE = πg.ResolveGlobal(πF, ßbuiltinify); πE != nil {
				continue
			}
			if πTemp015, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßload.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 678: def dumps(x, version=version):
			πF.SetLineno(678)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "x", Def: nil}
			if πTemp015, πE = πg.ResolveGlobal(πF, ßversion); πE != nil {
				continue
			}
			πTemp007[1] = πg.Param{Name: "version", Def: πTemp015}
			πTemp014 = πg.NewFunction(πg.NewCode("dumps", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]
				_ = µx
				var µversion *πg.Object = πArgs[1]
				_ = µversion
				var µbuffer *πg.Object = πg.UnboundLocal
				_ = µbuffer
				var µm *πg.Object = πg.UnboundLocal
				_ = µm
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 680: buffer = []
					πF.SetLineno(680)
					πTemp001 = make([]*πg.Object, 0)
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µbuffer = πTemp002
					// line 681: m = _Marshaller(buffer.append)
					πF.SetLineno(681)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µbuffer, ßappend, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_Marshaller); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µm = πTemp003
					// line 682: m.dump(x)
					πF.SetLineno(682)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µm, ßdump, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 683: return ''.join(buffer)
					πF.SetLineno(683)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbuffer, "buffer"); πE != nil {
						continue
					}
					πTemp001[0] = µbuffer
					if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdumps.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 677: @builtinify
			πF.SetLineno(677)
			πTemp002 = πF.MakeArgs(1)
			if πTemp015, πE = πg.ResolveGlobal(πF, ßdumps); πE != nil {
				continue
			}
			πTemp002[0] = πTemp015
			if πTemp015, πE = πg.ResolveGlobal(πF, ßbuiltinify); πE != nil {
				continue
			}
			if πTemp016, πE = πTemp015.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßdumps.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 686: def loads(s):
			πF.SetLineno(686)
			πTemp007 = make([]πg.Param, 1)
			πTemp007[0] = πg.Param{Name: "s", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("loads", "/usr/lib/python2.7/marshal.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]
				_ = µs
				var µum *πg.Object = πg.UnboundLocal
				_ = µum
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 687: um = _FastUnmarshaller(s)
					πF.SetLineno(687)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_FastUnmarshaller); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µum = πTemp003
					// line 688: return um.load()
					πF.SetLineno(688)
					if πE = πg.CheckLocal(πF, µum, "um"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µum, ßload, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßloads.ToObject(), πTemp015); πE != nil {
				continue
			}
			// line 685: @builtinify
			πF.SetLineno(685)
			πTemp002 = πF.MakeArgs(1)
			if πTemp016, πE = πg.ResolveGlobal(πF, ßloads); πE != nil {
				continue
			}
			πTemp002[0] = πTemp016
			if πTemp016, πE = πg.ResolveGlobal(πF, ßbuiltinify); πE != nil {
				continue
			}
			if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßloads.ToObject(), πTemp017); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("marshal", Code)
}
