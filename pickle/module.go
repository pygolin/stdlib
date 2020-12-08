package pickle

import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/cStringIO"
	// _ "github.com/pygolin/stdlib/copy_reg"
	// _ "github.com/pygolin/stdlib/struct_goreservedkeyword"
	// _ "github.com/pygolin/stdlib/StringIO"
	// _ "github.com/pygolin/stdlib/binascii"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/re"
	// _ "github.com/pygolin/stdlib/doctest"
	// _ "github.com/pygolin/stdlib/marshal"
	// _ "github.com/pygolin/stdlib/types"
)

var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/pickle.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object
		_ = πR
		var πE *πg.BaseException
		_ = πE
		ß := πg.InternStr("")
		ß0 := πg.InternStr("0")
		ß0x := πg.InternStr("0x")
		ß0x0 := πg.InternStr("0x0")
		ß0x00 := πg.InternStr("0x00")
		ß0xff := πg.InternStr("0xff")
		ß1 := πg.InternStr("1")
		ß2 := πg.InternStr("2")
		ßAPPEND := πg.InternStr("APPEND")
		ßAPPENDS := πg.InternStr("APPENDS")
		ßAttributeError := πg.InternStr("AttributeError")
		ßBINFLOAT := πg.InternStr("BINFLOAT")
		ßBINGET := πg.InternStr("BINGET")
		ßBININT := πg.InternStr("BININT")
		ßBININT1 := πg.InternStr("BININT1")
		ßBININT2 := πg.InternStr("BININT2")
		ßBINPERSID := πg.InternStr("BINPERSID")
		ßBINPUT := πg.InternStr("BINPUT")
		ßBINSTRING := πg.InternStr("BINSTRING")
		ßBINUNICODE := πg.InternStr("BINUNICODE")
		ßBUILD := πg.InternStr("BUILD")
		ßBuiltinFunctionType := πg.InternStr("BuiltinFunctionType")
		ßClassType := πg.InternStr("ClassType")
		ßDICT := πg.InternStr("DICT")
		ßDUP := πg.InternStr("DUP")
		ßDictionaryType := πg.InternStr("DictionaryType")
		ßEMPTY_DICT := πg.InternStr("EMPTY_DICT")
		ßEMPTY_LIST := πg.InternStr("EMPTY_LIST")
		ßEMPTY_TUPLE := πg.InternStr("EMPTY_TUPLE")
		ßEOFError := πg.InternStr("EOFError")
		ßEXT1 := πg.InternStr("EXT1")
		ßEXT2 := πg.InternStr("EXT2")
		ßEXT4 := πg.InternStr("EXT4")
		ßException := πg.InternStr("Exception")
		ßF := πg.InternStr("F")
		ßFALSE := πg.InternStr("FALSE")
		ßFLOAT := πg.InternStr("FLOAT")
		ßFalse := πg.InternStr("False")
		ßFloatType := πg.InternStr("FloatType")
		ßFunctionType := πg.InternStr("FunctionType")
		ßG := πg.InternStr("G")
		ßGET := πg.InternStr("GET")
		ßGLOBAL := πg.InternStr("GLOBAL")
		ßHIGHEST_PROTOCOL := πg.InternStr("HIGHEST_PROTOCOL")
		ßI := πg.InternStr("I")
		ßINST := πg.InternStr("INST")
		ßINT := πg.InternStr("INT")
		ßImportError := πg.InternStr("ImportError")
		ßInstanceType := πg.InternStr("InstanceType")
		ßIntType := πg.InternStr("IntType")
		ßJ := πg.InternStr("J")
		ßK := πg.InternStr("K")
		ßKeyError := πg.InternStr("KeyError")
		ßL := πg.InternStr("L")
		ßLIST := πg.InternStr("LIST")
		ßLONG := πg.InternStr("LONG")
		ßLONG1 := πg.InternStr("LONG1")
		ßLONG4 := πg.InternStr("LONG4")
		ßLONG_BINGET := πg.InternStr("LONG_BINGET")
		ßLONG_BINPUT := πg.InternStr("LONG_BINPUT")
		ßListType := πg.InternStr("ListType")
		ßLongType := πg.InternStr("LongType")
		ßM := πg.InternStr("M")
		ßMARK := πg.InternStr("MARK")
		ßModuleType := πg.InternStr("ModuleType")
		ßN := πg.InternStr("N")
		ßNEWFALSE := πg.InternStr("NEWFALSE")
		ßNEWOBJ := πg.InternStr("NEWOBJ")
		ßNEWTRUE := πg.InternStr("NEWTRUE")
		ßNONE := πg.InternStr("NONE")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßNoneType := πg.InternStr("NoneType")
		ßOBJ := πg.InternStr("OBJ")
		ßP := πg.InternStr("P")
		ßPERSID := πg.InternStr("PERSID")
		ßPOP := πg.InternStr("POP")
		ßPOP_MARK := πg.InternStr("POP_MARK")
		ßPROTO := πg.InternStr("PROTO")
		ßPUT := πg.InternStr("PUT")
		ßPickleError := πg.InternStr("PickleError")
		ßPickler := πg.InternStr("Pickler")
		ßPicklingError := πg.InternStr("PicklingError")
		ßPyStringMap := πg.InternStr("PyStringMap")
		ßQ := πg.InternStr("Q")
		ßR := πg.InternStr("R")
		ßREDUCE := πg.InternStr("REDUCE")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßS := πg.InternStr("S")
		ßSETITEM := πg.InternStr("SETITEM")
		ßSETITEMS := πg.InternStr("SETITEMS")
		ßSHORT_BINSTRING := πg.InternStr("SHORT_BINSTRING")
		ßSTOP := πg.InternStr("STOP")
		ßSTRING := πg.InternStr("STRING")
		ßStopIteration := πg.InternStr("StopIteration")
		ßStringIO := πg.InternStr("StringIO")
		ßStringType := πg.InternStr("StringType")
		ßT := πg.InternStr("T")
		ßTRUE := πg.InternStr("TRUE")
		ßTUPLE := πg.InternStr("TUPLE")
		ßTUPLE1 := πg.InternStr("TUPLE1")
		ßTUPLE2 := πg.InternStr("TUPLE2")
		ßTUPLE3 := πg.InternStr("TUPLE3")
		ßTrue := πg.InternStr("True")
		ßTupleType := πg.InternStr("TupleType")
		ßTypeError := πg.InternStr("TypeError")
		ßTypeType := πg.InternStr("TypeType")
		ßU := πg.InternStr("U")
		ßUNICODE := πg.InternStr("UNICODE")
		ßUnicodeType := πg.InternStr("UnicodeType")
		ßUnpickler := πg.InternStr("Unpickler")
		ßUnpicklingError := πg.InternStr("UnpicklingError")
		ßV := πg.InternStr("V")
		ßValueError := πg.InternStr("ValueError")
		ßX := πg.InternStr("X")
		ß_BATCHSIZE := πg.InternStr("_BATCHSIZE")
		ß_EmptyClass := πg.InternStr("_EmptyClass")
		ß_Stop := πg.InternStr("_Stop")
		ß__all__ := πg.InternStr("__all__")
		ß__call__ := πg.InternStr("__call__")
		ß__class__ := πg.InternStr("__class__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__getinitargs__ := πg.InternStr("__getinitargs__")
		ß__getstate__ := πg.InternStr("__getstate__")
		ß__import__ := πg.InternStr("__import__")
		ß__init__ := πg.InternStr("__init__")
		ß__main__ := πg.InternStr("__main__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__newobj__ := πg.InternStr("__newobj__")
		ß__reduce__ := πg.InternStr("__reduce__")
		ß__reduce_ex__ := πg.InternStr("__reduce_ex__")
		ß__setstate__ := πg.InternStr("__setstate__")
		ß__version__ := πg.InternStr("__version__")
		ß_batch_appends := πg.InternStr("_batch_appends")
		ß_batch_setitems := πg.InternStr("_batch_setitems")
		ß_binascii := πg.InternStr("_binascii")
		ß_extension_cache := πg.InternStr("_extension_cache")
		ß_extension_registry := πg.InternStr("_extension_registry")
		ß_instantiate := πg.InternStr("_instantiate")
		ß_inverted_registry := πg.InternStr("_inverted_registry")
		ß_keep_alive := πg.InternStr("_keep_alive")
		ß_pickle_maybe_moduledict := πg.InternStr("_pickle_maybe_moduledict")
		ß_test := πg.InternStr("_test")
		ß_tuplesize2code := πg.InternStr("_tuplesize2code")
		ßa := πg.InternStr("a")
		ßappend := πg.InternStr("append")
		ßb := πg.InternStr("b")
		ßbin := πg.InternStr("bin")
		ßbool := πg.InternStr("bool")
		ßc := πg.InternStr("c")
		ßchr := πg.InternStr("chr")
		ßclassmap := πg.InternStr("classmap")
		ßclear := πg.InternStr("clear")
		ßclear_memo := πg.InternStr("clear_memo")
		ßcompatible_formats := πg.InternStr("compatible_formats")
		ßcpyext := πg.InternStr("cpyext")
		ßd := πg.InternStr("d")
		ßdecode := πg.InternStr("decode")
		ßdecode_long := πg.InternStr("decode_long")
		ßdir := πg.InternStr("dir")
		ßdispatch := πg.InternStr("dispatch")
		ßdispatch_table := πg.InternStr("dispatch_table")
		ßdump := πg.InternStr("dump")
		ßdumps := πg.InternStr("dumps")
		ße := πg.InternStr("e")
		ßencode := πg.InternStr("encode")
		ßencode_long := πg.InternStr("encode_long")
		ßendswith := πg.InternStr("endswith")
		ßexc_info := πg.InternStr("exc_info")
		ßextend := πg.InternStr("extend")
		ßfast := πg.InternStr("fast")
		ßfind_class := πg.InternStr("find_class")
		ßfloat := πg.InternStr("float")
		ßformat_version := πg.InternStr("format_version")
		ßg := πg.InternStr("g")
		ßget := πg.InternStr("get")
		ßget_extension := πg.InternStr("get_extension")
		ßgetattr := πg.InternStr("getattr")
		ßgetvalue := πg.InternStr("getvalue")
		ßh := πg.InternStr("h")
		ßhasattr := πg.InternStr("hasattr")
		ßhex := πg.InternStr("hex")
		ßhexlify := πg.InternStr("hexlify")
		ßi := πg.InternStr("i")
		ßid := πg.InternStr("id")
		ßint := πg.InternStr("int")
		ßintern := πg.InternStr("intern")
		ßisinstance := πg.InternStr("isinstance")
		ßissubclass := πg.InternStr("issubclass")
		ßisunicode := πg.InternStr("isunicode")
		ßitems := πg.InternStr("items")
		ßiter := πg.InternStr("iter")
		ßiteritems := πg.InternStr("iteritems")
		ßj := πg.InternStr("j")
		ßl := πg.InternStr("l")
		ßlen := πg.InternStr("len")
		ßload := πg.InternStr("load")
		ßload_append := πg.InternStr("load_append")
		ßload_appends := πg.InternStr("load_appends")
		ßload_binfloat := πg.InternStr("load_binfloat")
		ßload_binget := πg.InternStr("load_binget")
		ßload_binint := πg.InternStr("load_binint")
		ßload_binint1 := πg.InternStr("load_binint1")
		ßload_binint2 := πg.InternStr("load_binint2")
		ßload_binpersid := πg.InternStr("load_binpersid")
		ßload_binput := πg.InternStr("load_binput")
		ßload_binstring := πg.InternStr("load_binstring")
		ßload_binunicode := πg.InternStr("load_binunicode")
		ßload_build := πg.InternStr("load_build")
		ßload_dict := πg.InternStr("load_dict")
		ßload_dup := πg.InternStr("load_dup")
		ßload_empty_dictionary := πg.InternStr("load_empty_dictionary")
		ßload_empty_list := πg.InternStr("load_empty_list")
		ßload_empty_tuple := πg.InternStr("load_empty_tuple")
		ßload_eof := πg.InternStr("load_eof")
		ßload_ext1 := πg.InternStr("load_ext1")
		ßload_ext2 := πg.InternStr("load_ext2")
		ßload_ext4 := πg.InternStr("load_ext4")
		ßload_false := πg.InternStr("load_false")
		ßload_float := πg.InternStr("load_float")
		ßload_get := πg.InternStr("load_get")
		ßload_global := πg.InternStr("load_global")
		ßload_inst := πg.InternStr("load_inst")
		ßload_int := πg.InternStr("load_int")
		ßload_list := πg.InternStr("load_list")
		ßload_long := πg.InternStr("load_long")
		ßload_long1 := πg.InternStr("load_long1")
		ßload_long4 := πg.InternStr("load_long4")
		ßload_long_binget := πg.InternStr("load_long_binget")
		ßload_long_binput := πg.InternStr("load_long_binput")
		ßload_mark := πg.InternStr("load_mark")
		ßload_newobj := πg.InternStr("load_newobj")
		ßload_none := πg.InternStr("load_none")
		ßload_obj := πg.InternStr("load_obj")
		ßload_persid := πg.InternStr("load_persid")
		ßload_pop := πg.InternStr("load_pop")
		ßload_pop_mark := πg.InternStr("load_pop_mark")
		ßload_proto := πg.InternStr("load_proto")
		ßload_put := πg.InternStr("load_put")
		ßload_reduce := πg.InternStr("load_reduce")
		ßload_setitem := πg.InternStr("load_setitem")
		ßload_setitems := πg.InternStr("load_setitems")
		ßload_short_binstring := πg.InternStr("load_short_binstring")
		ßload_stop := πg.InternStr("load_stop")
		ßload_string := πg.InternStr("load_string")
		ßload_true := πg.InternStr("load_true")
		ßload_tuple := πg.InternStr("load_tuple")
		ßload_tuple1 := πg.InternStr("load_tuple1")
		ßload_tuple2 := πg.InternStr("load_tuple2")
		ßload_tuple3 := πg.InternStr("load_tuple3")
		ßload_unicode := πg.InternStr("load_unicode")
		ßloads := πg.InternStr("loads")
		ßlong := πg.InternStr("long")
		ßmark := πg.InternStr("mark")
		ßmarker := πg.InternStr("marker")
		ßmarshal := πg.InternStr("marshal")
		ßmatch := πg.InternStr("match")
		ßmemo := πg.InternStr("memo")
		ßmemoize := πg.InternStr("memoize")
		ßmloads := πg.InternStr("mloads")
		ßmodules := πg.InternStr("modules")
		ßnext := πg.InternStr("next")
		ßo := πg.InternStr("o")
		ßobject := πg.InternStr("object")
		ßord := πg.InternStr("ord")
		ßp := πg.InternStr("p")
		ßpack := πg.InternStr("pack")
		ßpersistent_id := πg.InternStr("persistent_id")
		ßpersistent_load := πg.InternStr("persistent_load")
		ßpop := πg.InternStr("pop")
		ßproto := πg.InternStr("proto")
		ßput := πg.InternStr("put")
		ßq := πg.InternStr("q")
		ßr := πg.InternStr("r")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßread := πg.InternStr("read")
		ßreadline := πg.InternStr("readline")
		ßreplace := πg.InternStr("replace")
		ßrepr := πg.InternStr("repr")
		ßs := πg.InternStr("s")
		ßsave := πg.InternStr("save")
		ßsave_bool := πg.InternStr("save_bool")
		ßsave_dict := πg.InternStr("save_dict")
		ßsave_empty_tuple := πg.InternStr("save_empty_tuple")
		ßsave_float := πg.InternStr("save_float")
		ßsave_function := πg.InternStr("save_function")
		ßsave_global := πg.InternStr("save_global")
		ßsave_inst := πg.InternStr("save_inst")
		ßsave_int := πg.InternStr("save_int")
		ßsave_list := πg.InternStr("save_list")
		ßsave_long := πg.InternStr("save_long")
		ßsave_none := πg.InternStr("save_none")
		ßsave_pers := πg.InternStr("save_pers")
		ßsave_reduce := πg.InternStr("save_reduce")
		ßsave_string := πg.InternStr("save_string")
		ßsave_tuple := πg.InternStr("save_tuple")
		ßsave_unicode := πg.InternStr("save_unicode")
		ßsetattr := πg.InternStr("setattr")
		ßstack := πg.InternStr("stack")
		ßstartswith := πg.InternStr("startswith")
		ßstr := πg.InternStr("str")
		ßstruct := πg.InternStr("struct")
		ßsys := πg.InternStr("sys")
		ßt := πg.InternStr("t")
		ßtestmod := πg.InternStr("testmod")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßu := πg.InternStr("u")
		ßunhexlify := πg.InternStr("unhexlify")
		ßunicode := πg.InternStr("unicode")
		ßunpack := πg.InternStr("unpack")
		ßupdate := πg.InternStr("update")
		ßvalue := πg.InternStr("value")
		ßwhichmodule := πg.InternStr("whichmodule")
		ßwrite := πg.InternStr("write")
		ßx := πg.InternStr("x")
		ßxrange := πg.InternStr("xrange")
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
		var πTemp006 *πg.BaseException
		_ = πTemp006
		var πTemp007 *πg.Traceback
		_ = πTemp007
		var πTemp008 bool
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
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 8:
				goto Label8
			case 2:
				goto Label2
			case 11:
				goto Label11
			case 5:
				goto Label5
			case 14:
				goto Label14
			default:
				panic("unexpected function state")
			}
			// line 1: """Create portable serialized representations of Python objects.
			πF.SetLineno(1)
			// line 1: """Create portable serialized representations of Python objects.
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Create portable serialized representations of Python objects.\n\nSee module cPickle for a (much) faster implementation.\nSee module copy_reg for a mechanism for registering custom picklers.\nSee module pickletools source for extensive comments.\n\nClasses:\n\n    Pickler\n    Unpickler\n\nFunctions:\n\n    dump(object, file)\n    dumps(object) -> string\n    load(file) -> object\n    loads(string) -> object\n\nMisc variables:\n\n    __version__\n    format_version\n    compatible_formats\n\n").ToObject()); πE != nil {
				continue
			}
			// line 27: __version__ = "$Revision: 72223 $"       # Code version
			πF.SetLineno(27)
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πg.NewStr("$Revision: 72223 $").ToObject()); πE != nil {
				continue
			}
			// line 29: from types import *
			πF.SetLineno(29)
			if πTemp002, πE = πg.ImportModule(πF, "types"); πE != nil {
				continue
			}
			if πE = πg.LoadMembers(πF, πTemp002[0]); πE != nil {
				continue
			}
			// line 30: from copy_reg import dispatch_table
			πF.SetLineno(30)
			if πTemp002, πE = πg.ImportModule(πF, "copy_reg"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßdispatch_table); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdispatch_table.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 31: from copy_reg import _extension_registry, _inverted_registry, _extension_cache
			πF.SetLineno(31)
			if πTemp002, πE = πg.ImportModule(πF, "copy_reg"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß_extension_registry); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_extension_registry.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß_inverted_registry); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_inverted_registry.ToObject(), πTemp003); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ß_extension_cache); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_extension_cache.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 32: import marshal
			πF.SetLineno(32)
			if πTemp002, πE = πg.ImportModule(πF, "marshal"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßmarshal.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: import sys
			πF.SetLineno(33)
			if πTemp002, πE = πg.ImportModule(πF, "sys"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: import struct
			πF.SetLineno(34)
			if πTemp002, πE = πg.ImportModule(πF, "struct"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstruct.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 35: import re
			πF.SetLineno(35)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 37: __all__ = ["PickleError", "PicklingError", "UnpicklingError", "Pickler",
			πF.SetLineno(37)
			πTemp002 = make([]*πg.Object, 9)
			πTemp002[0] = ßPickleError.ToObject()
			πTemp002[1] = ßPicklingError.ToObject()
			πTemp002[2] = ßUnpicklingError.ToObject()
			πTemp002[3] = ßPickler.ToObject()
			πTemp002[4] = ßUnpickler.ToObject()
			πTemp002[5] = ßdump.ToObject()
			πTemp002[6] = ßdumps.ToObject()
			πTemp002[7] = ßload.ToObject()
			πTemp002[8] = ßloads.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 41: format_version = "2.0"                  # File format version we write
			πF.SetLineno(41)
			if πE = πF.Globals().SetItem(πF, ßformat_version.ToObject(), πg.NewStr("2.0").ToObject()); πE != nil {
				continue
			}
			// line 42: compatible_formats = ["1.0",            # Original protocol 0
			πF.SetLineno(42)
			πTemp002 = make([]*πg.Object, 5)
			πTemp002[0] = πg.NewStr("1.0").ToObject()
			πTemp002[1] = πg.NewStr("1.1").ToObject()
			πTemp002[2] = πg.NewStr("1.2").ToObject()
			πTemp002[3] = πg.NewStr("1.3").ToObject()
			πTemp002[4] = πg.NewStr("2.0").ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßcompatible_formats.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 51: HIGHEST_PROTOCOL = 2
			πF.SetLineno(51)
			if πE = πF.Globals().SetItem(πF, ßHIGHEST_PROTOCOL.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
				continue
			}
			// line 56: mloads = marshal.loads
			πF.SetLineno(56)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßmarshal); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßloads, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßmloads.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 58: class PickleError(Exception):
			πF.SetLineno(58)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("PickleError", "/usr/lib/python2.7/pickle.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 59: """A common base class for the other pickling exceptions."""
					πF.SetLineno(59)
					// line 59: """A common base class for the other pickling exceptions."""
					πF.SetLineno(59)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("A common base class for the other pickling exceptions.").ToObject()); πE != nil {
						continue
					}
					// line 60: pass
					πF.SetLineno(60)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("PickleError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPickleError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 62: class PicklingError(PickleError):
			πF.SetLineno(62)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPickleError); πE != nil {
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
			_, πE = πg.NewCode("PicklingError", "/usr/lib/python2.7/pickle.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 63: """This exception is raised when an unpicklable object is passed to the
					πF.SetLineno(63)
					// line 63: """This exception is raised when an unpicklable object is passed to the
					πF.SetLineno(63)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("This exception is raised when an unpicklable object is passed to the\n    dump() method.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 67: pass
					πF.SetLineno(67)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("PicklingError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPicklingError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 69: class UnpicklingError(PickleError):
			πF.SetLineno(69)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPickleError); πE != nil {
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
			_, πE = πg.NewCode("UnpicklingError", "/usr/lib/python2.7/pickle.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 70: """This exception is raised when there is a problem unpickling an object,
					πF.SetLineno(70)
					// line 70: """This exception is raised when there is a problem unpickling an object,
					πF.SetLineno(70)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("This exception is raised when there is a problem unpickling an object,\n    such as a security violation.\n\n    Note that other exceptions may also be raised during unpickling, including\n    (but not necessarily limited to) AttributeError, EOFError, ImportError,\n    and IndexError.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 78: pass
					πF.SetLineno(78)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("UnpicklingError").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnpicklingError.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 82: class _Stop(Exception):
			πF.SetLineno(82)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
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
			_, πE = πg.NewCode("_Stop", "/usr/lib/python2.7/pickle.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
				_ = πTemp002
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 83: def __init__(self, value):
					πF.SetLineno(83)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 84: self.value = value
							πF.SetLineno(84)
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
			if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("_Stop").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Stop.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 87: try:
			πF.SetLineno(87)
			πF.PushCheckpoint(2)
			// line 88: from org.python.core import PyStringMap
			πF.SetLineno(88)
			if πTemp002, πE = πg.ImportModule(πF, "org.python.core"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[2]
			if πTemp003, πE = πg.GetAttrImport(πF, πTemp001, ßPyStringMap); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPyStringMap.ToObject(), πTemp003); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label1
		Label2:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp006, πTemp007 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label3
			}
			πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
			continue
			// line 89: except ImportError:
			πF.SetLineno(89)
		Label3:
			// line 90: PyStringMap = None
			πF.SetLineno(90)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPyStringMap.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 93: try:
			πF.SetLineno(93)
			πF.PushCheckpoint(5)
			// line 94: UnicodeType
			πF.SetLineno(94)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßUnicodeType); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label4
		Label5:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp006, πTemp007 = πF.ExcInfo()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label6
			}
			πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
			continue
			// line 95: except NameError:
			πF.SetLineno(95)
		Label6:
			// line 96: UnicodeType = None
			πF.SetLineno(96)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnicodeType.ToObject(), πTemp001); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label4
		Label4:
			// line 102: MARK            = '('   # push special markobject on stack
			πF.SetLineno(102)
			if πE = πF.Globals().SetItem(πF, ßMARK.ToObject(), πg.NewStr("(").ToObject()); πE != nil {
				continue
			}
			// line 103: STOP            = '.'   # every pickle ends with STOP
			πF.SetLineno(103)
			if πE = πF.Globals().SetItem(πF, ßSTOP.ToObject(), πg.NewStr(".").ToObject()); πE != nil {
				continue
			}
			// line 104: POP             = '0'   # discard topmost stack item
			πF.SetLineno(104)
			if πE = πF.Globals().SetItem(πF, ßPOP.ToObject(), ß0.ToObject()); πE != nil {
				continue
			}
			// line 105: POP_MARK        = '1'   # discard stack top through topmost markobject
			πF.SetLineno(105)
			if πE = πF.Globals().SetItem(πF, ßPOP_MARK.ToObject(), ß1.ToObject()); πE != nil {
				continue
			}
			// line 106: DUP             = '2'   # duplicate top stack item
			πF.SetLineno(106)
			if πE = πF.Globals().SetItem(πF, ßDUP.ToObject(), ß2.ToObject()); πE != nil {
				continue
			}
			// line 107: FLOAT           = 'F'   # push float object; decimal string argument
			πF.SetLineno(107)
			if πE = πF.Globals().SetItem(πF, ßFLOAT.ToObject(), ßF.ToObject()); πE != nil {
				continue
			}
			// line 108: INT             = 'I'   # push integer or bool; decimal string argument
			πF.SetLineno(108)
			if πE = πF.Globals().SetItem(πF, ßINT.ToObject(), ßI.ToObject()); πE != nil {
				continue
			}
			// line 109: BININT          = 'J'   # push four-byte signed int
			πF.SetLineno(109)
			if πE = πF.Globals().SetItem(πF, ßBININT.ToObject(), ßJ.ToObject()); πE != nil {
				continue
			}
			// line 110: BININT1         = 'K'   # push 1-byte unsigned int
			πF.SetLineno(110)
			if πE = πF.Globals().SetItem(πF, ßBININT1.ToObject(), ßK.ToObject()); πE != nil {
				continue
			}
			// line 111: LONG            = 'L'   # push long; decimal string argument
			πF.SetLineno(111)
			if πE = πF.Globals().SetItem(πF, ßLONG.ToObject(), ßL.ToObject()); πE != nil {
				continue
			}
			// line 112: BININT2         = 'M'   # push 2-byte unsigned int
			πF.SetLineno(112)
			if πE = πF.Globals().SetItem(πF, ßBININT2.ToObject(), ßM.ToObject()); πE != nil {
				continue
			}
			// line 113: NONE            = 'N'   # push None
			πF.SetLineno(113)
			if πE = πF.Globals().SetItem(πF, ßNONE.ToObject(), ßN.ToObject()); πE != nil {
				continue
			}
			// line 114: PERSID          = 'P'   # push persistent object; id is taken from string arg
			πF.SetLineno(114)
			if πE = πF.Globals().SetItem(πF, ßPERSID.ToObject(), ßP.ToObject()); πE != nil {
				continue
			}
			// line 115: BINPERSID       = 'Q'   #  "       "         "  ;  "  "   "     "  stack
			πF.SetLineno(115)
			if πE = πF.Globals().SetItem(πF, ßBINPERSID.ToObject(), ßQ.ToObject()); πE != nil {
				continue
			}
			// line 116: REDUCE          = 'R'   # apply callable to argtuple, both on stack
			πF.SetLineno(116)
			if πE = πF.Globals().SetItem(πF, ßREDUCE.ToObject(), ßR.ToObject()); πE != nil {
				continue
			}
			// line 117: STRING          = 'S'   # push string; NL-terminated string argument
			πF.SetLineno(117)
			if πE = πF.Globals().SetItem(πF, ßSTRING.ToObject(), ßS.ToObject()); πE != nil {
				continue
			}
			// line 118: BINSTRING       = 'T'   # push string; counted binary string argument
			πF.SetLineno(118)
			if πE = πF.Globals().SetItem(πF, ßBINSTRING.ToObject(), ßT.ToObject()); πE != nil {
				continue
			}
			// line 119: SHORT_BINSTRING = 'U'   #  "     "   ;    "      "       "      " < 256 bytes
			πF.SetLineno(119)
			if πE = πF.Globals().SetItem(πF, ßSHORT_BINSTRING.ToObject(), ßU.ToObject()); πE != nil {
				continue
			}
			// line 120: UNICODE         = 'V'   # push Unicode string; raw-unicode-escaped'd argument
			πF.SetLineno(120)
			if πE = πF.Globals().SetItem(πF, ßUNICODE.ToObject(), ßV.ToObject()); πE != nil {
				continue
			}
			// line 121: BINUNICODE      = 'X'   #   "     "       "  ; counted UTF-8 string argument
			πF.SetLineno(121)
			if πE = πF.Globals().SetItem(πF, ßBINUNICODE.ToObject(), ßX.ToObject()); πE != nil {
				continue
			}
			// line 122: APPEND          = 'a'   # append stack top to list below it
			πF.SetLineno(122)
			if πE = πF.Globals().SetItem(πF, ßAPPEND.ToObject(), ßa.ToObject()); πE != nil {
				continue
			}
			// line 123: BUILD           = 'b'   # call __setstate__ or __dict__.update()
			πF.SetLineno(123)
			if πE = πF.Globals().SetItem(πF, ßBUILD.ToObject(), ßb.ToObject()); πE != nil {
				continue
			}
			// line 124: GLOBAL          = 'c'   # push self.find_class(modname, name); 2 string args
			πF.SetLineno(124)
			if πE = πF.Globals().SetItem(πF, ßGLOBAL.ToObject(), ßc.ToObject()); πE != nil {
				continue
			}
			// line 125: DICT            = 'd'   # build a dict from stack items
			πF.SetLineno(125)
			if πE = πF.Globals().SetItem(πF, ßDICT.ToObject(), ßd.ToObject()); πE != nil {
				continue
			}
			// line 126: EMPTY_DICT      = '}'   # push empty dict
			πF.SetLineno(126)
			if πE = πF.Globals().SetItem(πF, ßEMPTY_DICT.ToObject(), πg.NewStr("}").ToObject()); πE != nil {
				continue
			}
			// line 127: APPENDS         = 'e'   # extend list on stack by topmost stack slice
			πF.SetLineno(127)
			if πE = πF.Globals().SetItem(πF, ßAPPENDS.ToObject(), ße.ToObject()); πE != nil {
				continue
			}
			// line 128: GET             = 'g'   # push item from memo on stack; index is string arg
			πF.SetLineno(128)
			if πE = πF.Globals().SetItem(πF, ßGET.ToObject(), ßg.ToObject()); πE != nil {
				continue
			}
			// line 129: BINGET          = 'h'   #   "    "    "    "   "   "  ;   "    " 1-byte arg
			πF.SetLineno(129)
			if πE = πF.Globals().SetItem(πF, ßBINGET.ToObject(), ßh.ToObject()); πE != nil {
				continue
			}
			// line 130: INST            = 'i'   # build & push class instance
			πF.SetLineno(130)
			if πE = πF.Globals().SetItem(πF, ßINST.ToObject(), ßi.ToObject()); πE != nil {
				continue
			}
			// line 131: LONG_BINGET     = 'j'   # push item from memo on stack; index is 4-byte arg
			πF.SetLineno(131)
			if πE = πF.Globals().SetItem(πF, ßLONG_BINGET.ToObject(), ßj.ToObject()); πE != nil {
				continue
			}
			// line 132: LIST            = 'l'   # build list from topmost stack items
			πF.SetLineno(132)
			if πE = πF.Globals().SetItem(πF, ßLIST.ToObject(), ßl.ToObject()); πE != nil {
				continue
			}
			// line 133: EMPTY_LIST      = ']'   # push empty list
			πF.SetLineno(133)
			if πE = πF.Globals().SetItem(πF, ßEMPTY_LIST.ToObject(), πg.NewStr("]").ToObject()); πE != nil {
				continue
			}
			// line 134: OBJ             = 'o'   # build & push class instance
			πF.SetLineno(134)
			if πE = πF.Globals().SetItem(πF, ßOBJ.ToObject(), ßo.ToObject()); πE != nil {
				continue
			}
			// line 135: PUT             = 'p'   # store stack top in memo; index is string arg
			πF.SetLineno(135)
			if πE = πF.Globals().SetItem(πF, ßPUT.ToObject(), ßp.ToObject()); πE != nil {
				continue
			}
			// line 136: BINPUT          = 'q'   #   "     "    "   "   " ;   "    " 1-byte arg
			πF.SetLineno(136)
			if πE = πF.Globals().SetItem(πF, ßBINPUT.ToObject(), ßq.ToObject()); πE != nil {
				continue
			}
			// line 137: LONG_BINPUT     = 'r'   #   "     "    "   "   " ;   "    " 4-byte arg
			πF.SetLineno(137)
			if πE = πF.Globals().SetItem(πF, ßLONG_BINPUT.ToObject(), ßr.ToObject()); πE != nil {
				continue
			}
			// line 138: SETITEM         = 's'   # add key+value pair to dict
			πF.SetLineno(138)
			if πE = πF.Globals().SetItem(πF, ßSETITEM.ToObject(), ßs.ToObject()); πE != nil {
				continue
			}
			// line 139: TUPLE           = 't'   # build tuple from topmost stack items
			πF.SetLineno(139)
			if πE = πF.Globals().SetItem(πF, ßTUPLE.ToObject(), ßt.ToObject()); πE != nil {
				continue
			}
			// line 140: EMPTY_TUPLE     = ')'   # push empty tuple
			πF.SetLineno(140)
			if πE = πF.Globals().SetItem(πF, ßEMPTY_TUPLE.ToObject(), πg.NewStr(")").ToObject()); πE != nil {
				continue
			}
			// line 141: SETITEMS        = 'u'   # modify dict by adding topmost key+value pairs
			πF.SetLineno(141)
			if πE = πF.Globals().SetItem(πF, ßSETITEMS.ToObject(), ßu.ToObject()); πE != nil {
				continue
			}
			// line 142: BINFLOAT        = 'G'   # push float; arg is 8-byte float encoding
			πF.SetLineno(142)
			if πE = πF.Globals().SetItem(πF, ßBINFLOAT.ToObject(), ßG.ToObject()); πE != nil {
				continue
			}
			// line 144: TRUE            = 'I01\n'  # not an opcode; see INT docs in pickletools.py
			πF.SetLineno(144)
			if πE = πF.Globals().SetItem(πF, ßTRUE.ToObject(), πg.NewStr("I01\n").ToObject()); πE != nil {
				continue
			}
			// line 145: FALSE           = 'I00\n'  # not an opcode; see INT docs in pickletools.py
			πF.SetLineno(145)
			if πE = πF.Globals().SetItem(πF, ßFALSE.ToObject(), πg.NewStr("I00\n").ToObject()); πE != nil {
				continue
			}
			// line 149: PROTO           = '\x80'  # identify pickle protocol
			πF.SetLineno(149)
			if πE = πF.Globals().SetItem(πF, ßPROTO.ToObject(), πg.NewStr("\x80").ToObject()); πE != nil {
				continue
			}
			// line 150: NEWOBJ          = '\x81'  # build object by applying cls.__new__ to argtuple
			πF.SetLineno(150)
			if πE = πF.Globals().SetItem(πF, ßNEWOBJ.ToObject(), πg.NewStr("\x81").ToObject()); πE != nil {
				continue
			}
			// line 151: EXT1            = '\x82'  # push object from extension registry; 1-byte index
			πF.SetLineno(151)
			if πE = πF.Globals().SetItem(πF, ßEXT1.ToObject(), πg.NewStr("\x82").ToObject()); πE != nil {
				continue
			}
			// line 152: EXT2            = '\x83'  # ditto, but 2-byte index
			πF.SetLineno(152)
			if πE = πF.Globals().SetItem(πF, ßEXT2.ToObject(), πg.NewStr("\x83").ToObject()); πE != nil {
				continue
			}
			// line 153: EXT4            = '\x84'  # ditto, but 4-byte index
			πF.SetLineno(153)
			if πE = πF.Globals().SetItem(πF, ßEXT4.ToObject(), πg.NewStr("\x84").ToObject()); πE != nil {
				continue
			}
			// line 154: TUPLE1          = '\x85'  # build 1-tuple from stack top
			πF.SetLineno(154)
			if πE = πF.Globals().SetItem(πF, ßTUPLE1.ToObject(), πg.NewStr("\x85").ToObject()); πE != nil {
				continue
			}
			// line 155: TUPLE2          = '\x86'  # build 2-tuple from two topmost stack items
			πF.SetLineno(155)
			if πE = πF.Globals().SetItem(πF, ßTUPLE2.ToObject(), πg.NewStr("\x86").ToObject()); πE != nil {
				continue
			}
			// line 156: TUPLE3          = '\x87'  # build 3-tuple from three topmost stack items
			πF.SetLineno(156)
			if πE = πF.Globals().SetItem(πF, ßTUPLE3.ToObject(), πg.NewStr("\x87").ToObject()); πE != nil {
				continue
			}
			// line 157: NEWTRUE         = '\x88'  # push True
			πF.SetLineno(157)
			if πE = πF.Globals().SetItem(πF, ßNEWTRUE.ToObject(), πg.NewStr("\x88").ToObject()); πE != nil {
				continue
			}
			// line 158: NEWFALSE        = '\x89'  # push False
			πF.SetLineno(158)
			if πE = πF.Globals().SetItem(πF, ßNEWFALSE.ToObject(), πg.NewStr("\x89").ToObject()); πE != nil {
				continue
			}
			// line 159: LONG1           = '\x8a'  # push long from < 256 bytes
			πF.SetLineno(159)
			if πE = πF.Globals().SetItem(πF, ßLONG1.ToObject(), πg.NewStr("\x8a").ToObject()); πE != nil {
				continue
			}
			// line 160: LONG4           = '\x8b'  # push really big long
			πF.SetLineno(160)
			if πE = πF.Globals().SetItem(πF, ßLONG4.ToObject(), πg.NewStr("\x8b").ToObject()); πE != nil {
				continue
			}
			// line 162: _tuplesize2code = [EMPTY_TUPLE, TUPLE1, TUPLE2, TUPLE3]
			πF.SetLineno(162)
			πTemp002 = make([]*πg.Object, 4)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßEMPTY_TUPLE); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTUPLE1); πE != nil {
				continue
			}
			πTemp002[1] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTUPLE2); πE != nil {
				continue
			}
			πTemp002[2] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßTUPLE3); πE != nil {
				continue
			}
			πTemp002[3] = πTemp001
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß_tuplesize2code.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 165: __all__.extend([x for x in dir() if re.match("[A-Z][A-Z0-9_]+$",x)])
			πF.SetLineno(165)
			πTemp002 = πF.MakeArgs(1)
			πTemp009 = make([]πg.Param, 0)
			πTemp003 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πg.UnboundLocal
				_ = µx
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
						case 6:
							goto Label6
						default:
							panic("unexpected function state")
						}
						if πTemp002, πE = πg.ResolveGlobal(πF, ßdir); πE != nil {
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
							µx = πTemp002
						}
						if πE != nil || !πTemp005 {
							continue
						}
						πF.PushCheckpoint(1)
						πTemp006 = πF.MakeArgs(2)
						πTemp006[0] = πg.NewStr("[A-Z][A-Z0-9_]+$").ToObject()
						if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
							continue
						}
						πTemp006[1] = µx
						if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßmatch, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp005 {
							goto Label4
						}
						goto Label5
						// line 165: __all__.extend([x for x in dir() if re.match("[A-Z][A-Z0-9_]+$",x)])
						πF.SetLineno(165)
					Label4:
						// line 165: __all__.extend([x for x in dir() if re.match("[A-Z][A-Z0-9_]+$",x)])
						πF.SetLineno(165)
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
			if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ß__all__); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßextend, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 166: del x
			πF.SetLineno(166)
			if πE = πg.DelVar(πF, πF.Globals(), ßx); πE != nil {
				continue
			}
			// line 171: class Pickler(object):
			πF.SetLineno(171)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp010
			πTemp004 = πg.NewDict()
			if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Pickler", "/usr/lib/python2.7/pickle.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp013 *πg.Dict
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
				var πTemp025 bool
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 173: def __init__(self, file, protocol=None):
					πF.SetLineno(173)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "file", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "protocol", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfile *πg.Object = πArgs[1]
						_ = µfile
						var µprotocol *πg.Object = πArgs[2]
						_ = µprotocol
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
						var πTemp006 *πg.Dict
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
							// line 174: """This takes a file-like object for writing a pickle data stream.
							πF.SetLineno(174)
							if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µprotocol == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 197: if protocol is None:
							πF.SetLineno(197)
						Label1:
							// line 198: protocol = 0
							πF.SetLineno(198)
							µprotocol = πg.NewInt(0).ToObject()
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µprotocol, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µprotocol); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label4
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßHIGHEST_PROTOCOL); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LE(πF, µprotocol, πTemp004); πE != nil {
								continue
							}
						Label4:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
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
							// line 199: if protocol < 0:
							πF.SetLineno(199)
						Label3:
							// line 200: protocol = HIGHEST_PROTOCOL
							πF.SetLineno(200)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßHIGHEST_PROTOCOL); πE != nil {
								continue
							}
							µprotocol = πTemp001
							goto Label6
							// line 201: elif not 0 <= protocol <= HIGHEST_PROTOCOL:
							πF.SetLineno(201)
						Label5:
							πTemp005 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßHIGHEST_PROTOCOL); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("pickle protocol must be <= %d").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 202: raise ValueError("pickle protocol must be <= %d" % HIGHEST_PROTOCOL)
							πF.SetLineno(202)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label6
						Label6:
							// line 203: self.write = file.write
							πF.SetLineno(203)
							if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfile, ßwrite, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwrite, πTemp002); πE != nil {
								continue
							}
							// line 204: self.memo = {}
							πF.SetLineno(204)
							πTemp006 = πg.NewDict()
							πTemp001 = πTemp006.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmemo, πTemp002); πE != nil {
								continue
							}
							// line 205: self.proto = int(protocol)
							πF.SetLineno(205)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
								continue
							}
							πTemp005[0] = µprotocol
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßproto, πTemp001); πE != nil {
								continue
							}
							// line 206: self.bin = protocol >= 1
							πF.SetLineno(206)
							if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µprotocol, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßbin, πTemp002); πE != nil {
								continue
							}
							// line 207: self.fast = 0
							πF.SetLineno(207)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfast, πTemp001); πE != nil {
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
					// line 174: """This takes a file-like object for writing a pickle data stream.
					πF.SetLineno(174)
					// line 174: """This takes a file-like object for writing a pickle data stream.
					πF.SetLineno(174)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("This takes a file-like object for writing a pickle data stream.\n\n        The optional protocol argument tells the pickler to use the\n        given protocol; supported protocols are 0, 1, 2.  The default\n        protocol is 0, to be backwards compatible.  (Protocol 0 is the\n        only protocol that can be written to a file opened in text\n        mode and read back successfully.  When using a protocol higher\n        than 0, make sure the file is opened in binary mode, both when\n        pickling and unpickling.)\n\n        Protocol 1 is more efficient than protocol 0; protocol 2 is\n        more efficient than protocol 1.\n\n        Specifying a negative protocol version selects the highest\n        protocol version supported.  The higher the protocol used, the\n        more recent the version of Python needed to read the pickle\n        produced.\n\n        The file parameter must have a write() method that accepts a single\n        string argument.  It can thus be an open file object, a StringIO\n        object, or any other custom object that meets this interface.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("This takes a file-like object for writing a pickle data stream.\n\n        The optional protocol argument tells the pickler to use the\n        given protocol; supported protocols are 0, 1, 2.  The default\n        protocol is 0, to be backwards compatible.  (Protocol 0 is the\n        only protocol that can be written to a file opened in text\n        mode and read back successfully.  When using a protocol higher\n        than 0, make sure the file is opened in binary mode, both when\n        pickling and unpickling.)\n\n        Protocol 1 is more efficient than protocol 0; protocol 2 is\n        more efficient than protocol 1.\n\n        Specifying a negative protocol version selects the highest\n        protocol version supported.  The higher the protocol used, the\n        more recent the version of Python needed to read the pickle\n        produced.\n\n        The file parameter must have a write() method that accepts a single\n        string argument.  It can thus be an open file object, a StringIO\n        object, or any other custom object that meets this interface.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 209: def clear_memo(self):
					πF.SetLineno(209)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("clear_memo", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
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
							// line 210: """Clears the pickler's "memo".
							πF.SetLineno(210)
							// line 218: self.memo.clear()
							πF.SetLineno(218)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßclear, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßclear_memo.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 210: """Clears the pickler's "memo".
					πF.SetLineno(210)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Clears the pickler's \"memo\".\n\n        The memo is the data structure that remembers which objects the\n        pickler has already seen, so that shared or recursive objects are\n        pickled by reference and not by value.  This method is useful when\n        re-using picklers.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßclear_memo); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 220: def dump(self, obj):
					πF.SetLineno(220)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("dump", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
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
							// line 221: """Write a pickled representation of obj to the open file."""
							πF.SetLineno(221)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, πTemp002, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 222: if self.proto >= 2:
							πF.SetLineno(222)
						Label1:
							// line 223: self.write(PROTO + chr(self.proto))
							πF.SetLineno(223)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßPROTO); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp006
							if πTemp006, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp007); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label2
						Label2:
							// line 224: self.save(obj)
							πF.SetLineno(224)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsave, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 225: self.write(STOP)
							πF.SetLineno(225)
							πTemp004 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSTOP); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßdump.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 221: """Write a pickled representation of obj to the open file."""
					πF.SetLineno(221)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Write a pickled representation of obj to the open file.").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßdump); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 227: def memoize(self, obj):
					πF.SetLineno(227)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("memoize", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µmemo_len *πg.Object = πg.UnboundLocal
						_ = µmemo_len
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							default:
								panic("unexpected function state")
							}
							// line 228: """Store an object in the memo."""
							πF.SetLineno(228)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfast, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 242: if self.fast:
							πF.SetLineno(242)
						Label1:
							// line 243: return
							πF.SetLineno(243)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 244: assert id(obj) not in self.memo
							πF.SetLineno(244)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp003[0] = µobj
							if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp002).ToObject()
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							// line 245: memo_len = len(self.memo)
							πF.SetLineno(245)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µmemo_len = πTemp004
							// line 246: self.write(self.put(memo_len))
							πF.SetLineno(246)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmemo_len, "memo_len"); πE != nil {
								continue
							}
							πTemp006[0] = µmemo_len
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßput, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 247: self.memo[id(obj)] = memo_len, obj
							πF.SetLineno(247)
							if πE = πg.CheckLocal(πF, µmemo_len, "memo_len"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µmemo_len, µobj).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp003[0] = µobj
							if πTemp008, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp007 = πTemp009
							if πE = πg.SetItem(πF, πTemp005, πTemp007, πTemp004); πE != nil {
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
					if πE = πClass.SetItem(πF, ßmemoize.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 228: """Store an object in the memo."""
					πF.SetLineno(228)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Store an object in the memo.").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßmemoize); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 250: def put(self, i, pack=struct.pack):
					πF.SetLineno(250)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßpack, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "pack", Def: πTemp008}
					πTemp006 = πg.NewFunction(πg.NewCode("put", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
						var µpack *πg.Object = πArgs[2]
						_ = µpack
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 251: if self.bin:
							πF.SetLineno(251)
						Label1:
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µi, πg.NewInt(256).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 252: if i < 256:
							πF.SetLineno(252)
						Label3:
							// line 253: return BINPUT + chr(i)
							πF.SetLineno(253)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBINPUT); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[0] = µi
							if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label4:
							// line 255: return LONG_BINPUT + pack("<i", i)
							πF.SetLineno(255)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßLONG_BINPUT); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("<i").ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[1] = µi
							if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
								continue
							}
							if πTemp005, πE = µpack.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label5:
							goto Label2
						Label2:
							// line 257: return PUT + repr(i) + '\n'
							πF.SetLineno(257)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßPUT); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[0] = µi
							if πTemp006, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewStr("\n").ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßput.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 260: def get(self, i, pack=struct.pack):
					πF.SetLineno(260)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "i", Def: nil}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßpack, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "pack", Def: πTemp009}
					πTemp007 = πg.NewFunction(πg.NewCode("get", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πArgs[1]
						_ = µi
						var µpack *πg.Object = πArgs[2]
						_ = µpack
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 261: if self.bin:
							πF.SetLineno(261)
						Label1:
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µi, πg.NewInt(256).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 262: if i < 256:
							πF.SetLineno(262)
						Label3:
							// line 263: return BINGET + chr(i)
							πF.SetLineno(263)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßBINGET); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[0] = µi
							if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp006); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label4:
							// line 265: return LONG_BINGET + pack("<i", i)
							πF.SetLineno(265)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßLONG_BINGET); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("<i").ToObject()
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[1] = µi
							if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
								continue
							}
							if πTemp005, πE = µpack.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label5
						Label5:
							goto Label2
						Label2:
							// line 267: return GET + repr(i) + '\n'
							πF.SetLineno(267)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßGET); πE != nil {
								continue
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp004[0] = µi
							if πTemp006, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Add(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewStr("\n").ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 269: def save(self, obj):
					πF.SetLineno(269)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("save", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µpid *πg.Object = πg.UnboundLocal
						_ = µpid
						var µx *πg.Object = πg.UnboundLocal
						_ = µx
						var µt *πg.Object = πg.UnboundLocal
						_ = µt
						var µf *πg.Object = πg.UnboundLocal
						_ = µf
						var µreduce *πg.Object = πg.UnboundLocal
						_ = µreduce
						var µrv *πg.Object = πg.UnboundLocal
						_ = µrv
						var µissc *πg.Object = πg.UnboundLocal
						_ = µissc
						var µl *πg.Object = πg.UnboundLocal
						_ = µl
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 bool
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 πg.KWArgs
						_ = πTemp009
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 11:
								goto Label11
							default:
								panic("unexpected function state")
							}
							// line 271: pid = self.persistent_id(obj)
							πF.SetLineno(271)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßpersistent_id, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µpid = πTemp003
							if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µpid != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 272: if pid is not None:
							πF.SetLineno(272)
						Label1:
							// line 273: self.save_pers(pid)
							πF.SetLineno(273)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
								continue
							}
							πTemp001[0] = µpid
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsave_pers, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 274: return
							πF.SetLineno(274)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 277: x = self.memo.get(id(obj))
							πF.SetLineno(277)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp005[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µx = πTemp002
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µx); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 278: if x:
							πF.SetLineno(278)
						Label3:
							// line 279: self.write(self.get(x[0]))
							πF.SetLineno(279)
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µx, πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 280: return
							πF.SetLineno(280)
							πR = πg.None
							continue
							goto Label4
						Label4:
							// line 283: t = type(obj)
							πF.SetLineno(283)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µt = πTemp003
							// line 284: f = self.dispatch.get(t)
							πF.SetLineno(284)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp001[0] = µt
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdispatch, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µf = πTemp002
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µf); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							goto Label6
							// line 285: if f:
							πF.SetLineno(285)
						Label5:
							// line 286: f(self, obj) # Call unbound method with explicit self
							πF.SetLineno(286)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[1] = µobj
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							if πTemp002, πE = µf.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 287: return
							πF.SetLineno(287)
							πR = πg.None
							continue
							goto Label6
						Label6:
							// line 290: reduce = dispatch_table.get(t)
							πF.SetLineno(290)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp001[0] = µt
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdispatch_table); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µreduce = πTemp002
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µreduce); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 291: if reduce:
							πF.SetLineno(291)
						Label7:
							// line 292: rv = reduce(obj)
							πF.SetLineno(292)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp002, πE = µreduce.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrv = πTemp002
							goto Label9
						Label8:
							// line 295: try:
							πF.SetLineno(295)
							πF.PushCheckpoint(11)
							// line 296: issc = issubclass(t, TypeType)
							πF.SetLineno(296)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							πTemp001[0] = µt
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeType); πE != nil {
								continue
							}
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßissubclass); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µissc = πTemp003
							πF.PopCheckpoint()
							goto Label10
						Label11:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 297: except TypeError: # t is not a class (old Boost; see SF #502085)
							πF.SetLineno(297)
						Label12:
							// line 298: issc = 0
							πF.SetLineno(298)
							µissc = πg.NewInt(0).ToObject()
							πF.RestoreExc(nil, nil)
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µissc, "issc"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µissc); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label13
							}
							goto Label14
							// line 299: if issc:
							πF.SetLineno(299)
						Label13:
							// line 300: self.save_global(obj)
							πF.SetLineno(300)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsave_global, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 301: return
							πF.SetLineno(301)
							πR = πg.None
							continue
							goto Label14
						Label14:
							// line 304: reduce = getattr(obj, "__reduce_ex__", None)
							πF.SetLineno(304)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							πTemp001[1] = ß__reduce_ex__.ToObject()
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
							µreduce = πTemp003
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µreduce); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							goto Label16
							// line 305: if reduce:
							πF.SetLineno(305)
						Label15:
							// line 306: rv = reduce(self.proto)
							πF.SetLineno(306)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp002, πE = µreduce.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrv = πTemp002
							goto Label17
						Label16:
							// line 308: reduce = getattr(obj, "__reduce__", None)
							πF.SetLineno(308)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							πTemp001[1] = ß__reduce__.ToObject()
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
							µreduce = πTemp003
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µreduce); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label18
							}
							goto Label19
							// line 309: if reduce:
							πF.SetLineno(309)
						Label18:
							// line 310: rv = reduce()
							πF.SetLineno(310)
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp002, πE = µreduce.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrv = πTemp002
							goto Label20
						Label19:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µt, "t"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µt, ß__name__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp003 = πg.NewTuple2(πTemp008, µobj).ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Can't pickle %r object: %r").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 312: raise PicklingError("Can't pickle %r object: %r" %
							πF.SetLineno(312)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label20
						Label20:
							goto Label17
						Label17:
							goto Label9
						Label9:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
								continue
							}
							πTemp001[0] = µrv
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßStringType); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp008 == πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label21
							}
							goto Label22
							// line 316: if type(rv) is StringType:
							πF.SetLineno(316)
						Label21:
							// line 317: self.save_global(obj, rv)
							πF.SetLineno(317)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
								continue
							}
							πTemp001[1] = µrv
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsave_global, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 318: return
							πF.SetLineno(318)
							πR = πg.None
							continue
							goto Label22
						Label22:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
								continue
							}
							πTemp001[0] = µrv
							if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTupleType); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp008 != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label23
							}
							goto Label24
							// line 321: if type(rv) is not TupleType:
							πF.SetLineno(321)
						Label23:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("%s must return string or tuple").ToObject(), µreduce); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 322: raise PicklingError("%s must return string or tuple" % reduce)
							πF.SetLineno(322)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label24
						Label24:
							// line 325: l = len(rv)
							πF.SetLineno(325)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
								continue
							}
							πTemp001[0] = µrv
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µl = πTemp003
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LE(πF, πg.NewInt(2).ToObject(), µl); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label25
							}
							if πTemp003, πE = πg.LE(πF, µl, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
						Label25:
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label26
							}
							goto Label27
							// line 326: if not (2 <= l <= 5):
							πF.SetLineno(326)
						Label26:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Tuple returned by %s must have two to five elements").ToObject(), µreduce); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 327: raise PicklingError("Tuple returned by %s must have "
							πF.SetLineno(327)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label27
						Label27:
							// line 331: self.save_reduce(obj=obj, *rv)
							πF.SetLineno(331)
							if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp009 = πg.KWArgs{
								{"obj", µobj},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsave_reduce, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, µrv, πTemp009, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 333: def persistent_id(self, obj):
					πF.SetLineno(333)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("persistent_id", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
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
							// line 335: return None
							πF.SetLineno(335)
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
					if πE = πClass.SetItem(πF, ßpersistent_id.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 337: def save_pers(self, pid):
					πF.SetLineno(337)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "pid", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("save_pers", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpid *πg.Object = πArgs[1]
						_ = µpid
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 339: if self.bin:
							πF.SetLineno(339)
						Label1:
							// line 340: self.save(pid)
							πF.SetLineno(340)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
								continue
							}
							πTemp003[0] = µpid
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsave, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 341: self.write(BINPERSID)
							πF.SetLineno(341)
							πTemp003 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßBINPERSID); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label2:
							// line 343: self.write(PERSID + str(pid) + '\n')
							πF.SetLineno(343)
							πTemp003 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßPERSID); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
								continue
							}
							πTemp006[0] = µpid
							if πTemp007, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßsave_pers.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 345: def save_reduce(self, func, args, state=None,
					πF.SetLineno(345)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "func", Def: nil}
					πTemp002[2] = πg.Param{Name: "args", Def: nil}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "state", Def: πTemp012}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "listitems", Def: πTemp012}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "dictitems", Def: πTemp012}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "obj", Def: πTemp012}
					πTemp011 = πg.NewFunction(πg.NewCode("save_reduce", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfunc *πg.Object = πArgs[1]
						_ = µfunc
						var µargs *πg.Object = πArgs[2]
						_ = µargs
						var µstate *πg.Object = πArgs[3]
						_ = µstate
						var µlistitems *πg.Object = πArgs[4]
						_ = µlistitems
						var µdictitems *πg.Object = πArgs[5]
						_ = µdictitems
						var µobj *πg.Object = πArgs[6]
						_ = µobj
						var µsave *πg.Object = πg.UnboundLocal
						_ = µsave
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
						var µcls *πg.Object = πg.UnboundLocal
						_ = µcls
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 []*πg.Object
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
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
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTupleType); πE != nil {
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
							// line 350: if not isinstance(args, TupleType):
							πF.SetLineno(350)
						Label1:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("args from reduce() should be a tuple").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 351: raise PicklingError("args from reduce() should be a tuple")
							πF.SetLineno(351)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 354: try:
							πF.SetLineno(354)
							πF.PushCheckpoint(4)
							// line 355: func.__call__
							πF.SetLineno(355)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfunc, ß__call__, nil); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 356: except AttributeError:
							πF.SetLineno(356)
						Label5:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("func from reduce should be callable").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 357: raise PicklingError("func from reduce should be callable")
							πF.SetLineno(357)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							// line 359: save = self.save
							πF.SetLineno(359)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsave, nil); πE != nil {
								continue
							}
							µsave = πTemp001
							// line 360: write = self.write
							πF.SetLineno(360)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, πTemp004, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label6
							}
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							πTemp002[0] = µfunc
							πTemp002[1] = ß__name__.ToObject()
							πTemp002[2] = ß.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp003, πE = πg.Eq(πF, πTemp008, ß__newobj__.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label6:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 363: if self.proto >= 2 and getattr(func, "__name__", "") == "__newobj__":
							πF.SetLineno(363)
						Label7:
							// line 390: cls = args[0]
							πF.SetLineno(390)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µcls = πTemp003
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp002[0] = µcls
							πTemp002[1] = ß__new__.ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
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
								goto Label10
							}
							goto Label11
							// line 391: if not hasattr(cls, "__new__"):
							πF.SetLineno(391)
						Label10:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("args[0] from __newobj__ args has no __new__").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 392: raise PicklingError(
							πF.SetLineno(392)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label11
						Label11:
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µobj != πTemp004).ToObject()
							πTemp001 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label12
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µobj, ß__class__, nil); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µcls != πTemp004).ToObject()
							πTemp001 = πTemp003
						Label12:
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label13
							}
							goto Label14
							// line 394: if obj is not None and cls is not obj.__class__:
							πF.SetLineno(394)
						Label13:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("args[0] from __newobj__ args has the wrong class").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 395: raise PicklingError(
							πF.SetLineno(395)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label14
						Label14:
							// line 397: args = args[1:]
							πF.SetLineno(397)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µargs, πTemp001); πE != nil {
								continue
							}
							µargs = πTemp003
							// line 398: save(cls)
							πF.SetLineno(398)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp002[0] = µcls
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 399: save(args)
							πF.SetLineno(399)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 400: write(NEWOBJ)
							πF.SetLineno(400)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNEWOBJ); πE != nil {
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
							goto Label9
						Label8:
							// line 402: save(func)
							πF.SetLineno(402)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							πTemp002[0] = µfunc
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 403: save(args)
							πF.SetLineno(403)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 404: write(REDUCE)
							πF.SetLineno(404)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßREDUCE); πE != nil {
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
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µobj != πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label15
							}
							goto Label16
							// line 406: if obj is not None:
							πF.SetLineno(406)
						Label15:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
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
								goto Label17
							}
							goto Label18
							// line 410: if id(obj) in self.memo:
							πF.SetLineno(410)
						Label17:
							// line 411: write(POP + self.get(self.memo[id(obj)][0]))
							πF.SetLineno(411)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßPOP); πE != nil {
								continue
							}
							πTemp009 = πF.MakeArgs(1)
							πTemp004 = πg.NewInt(0).ToObject()
							πTemp011 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp011[0] = µobj
							if πTemp012, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp012.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp010 = πTemp013
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, πTemp013, πTemp010); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp012, πTemp004); πE != nil {
								continue
							}
							πTemp009[0] = πTemp008
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp008); πE != nil {
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
							goto Label19
						Label18:
							// line 413: self.memoize(obj)
							πF.SetLineno(413)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label19
						Label19:
							goto Label16
						Label16:
							if πE = πg.CheckLocal(πF, µlistitems, "listitems"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µlistitems != πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label20
							}
							goto Label21
							// line 420: if listitems is not None:
							πF.SetLineno(420)
						Label20:
							// line 421: self._batch_appends(listitems)
							πF.SetLineno(421)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlistitems, "listitems"); πE != nil {
								continue
							}
							πTemp002[0] = µlistitems
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_batch_appends, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label21
						Label21:
							if πE = πg.CheckLocal(πF, µdictitems, "dictitems"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdictitems != πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label22
							}
							goto Label23
							// line 423: if dictitems is not None:
							πF.SetLineno(423)
						Label22:
							// line 424: self._batch_setitems(dictitems)
							πF.SetLineno(424)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdictitems, "dictitems"); πE != nil {
								continue
							}
							πTemp002[0] = µdictitems
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_batch_setitems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label23
						Label23:
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µstate != πTemp003).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label24
							}
							goto Label25
							// line 426: if state is not None:
							πF.SetLineno(426)
						Label24:
							// line 427: save(state)
							πF.SetLineno(427)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp002[0] = µstate
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 428: write(BUILD)
							πF.SetLineno(428)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßBUILD); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_reduce.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 432: dispatch = {}
					πF.SetLineno(432)
					πTemp013 = πg.NewDict()
					πTemp012 = πTemp013.ToObject()
					if πE = πClass.SetItem(πF, ßdispatch.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 434: def save_none(self, obj):
					πF.SetLineno(434)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("save_none", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
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
							// line 435: self.write(NONE)
							πF.SetLineno(435)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNONE); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_none.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 436: dispatch[NoneType] = save_none
					πF.SetLineno(436)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßsave_none); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πTemp014); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßNoneType); πE != nil {
						continue
					}
					πTemp017 = πTemp018
					if πE = πg.SetItem(πF, πTemp016, πTemp017, πTemp015); πE != nil {
						continue
					}
					// line 438: def save_bool(self, obj):
					πF.SetLineno(438)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("save_bool", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, πTemp002, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 439: if self.proto >= 2:
							πF.SetLineno(439)
						Label1:
							// line 440: self.write(obj and NEWTRUE or NEWFALSE)
							πF.SetLineno(440)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002 = µobj
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label5
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNEWTRUE); πE != nil {
								continue
							}
							πTemp002 = πTemp006
						Label5:
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNEWFALSE); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label4:
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label3
						Label2:
							// line 442: self.write(obj and TRUE or FALSE)
							πF.SetLineno(442)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002 = µobj
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label7
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßTRUE); πE != nil {
								continue
							}
							πTemp002 = πTemp006
						Label7:
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFALSE); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label6:
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
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
					if πE = πClass.SetItem(πF, ßsave_bool.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 443: dispatch[bool] = save_bool
					πF.SetLineno(443)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßsave_bool); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp015); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßbool); πE != nil {
						continue
					}
					πTemp018 = πTemp019
					if πE = πg.SetItem(πF, πTemp017, πTemp018, πTemp016); πE != nil {
						continue
					}
					// line 445: def save_int(self, obj, pack=struct.pack):
					πF.SetLineno(445)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßpack, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "pack", Def: πTemp017}
					πTemp015 = πg.NewFunction(πg.NewCode("save_int", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µpack *πg.Object = πArgs[2]
						_ = µpack
						var µhigh_bits *πg.Object = πg.UnboundLocal
						_ = µhigh_bits
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 446: if self.bin:
							πF.SetLineno(446)
						Label1:
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, µobj, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 451: if obj >= 0:
							πF.SetLineno(451)
						Label3:
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, µobj, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label5
							}
							goto Label6
							// line 452: if obj <= 0xff:
							πF.SetLineno(452)
						Label5:
							// line 453: self.write(BININT1 + chr(obj))
							πF.SetLineno(453)
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßBININT1); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp005[0] = µobj
							if πTemp006, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Add(πF, πTemp004, πTemp007); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 454: return
							πF.SetLineno(454)
							πR = πg.None
							continue
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, µobj, πg.NewInt(65535).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label7
							}
							goto Label8
							// line 455: if obj <= 0xffff:
							πF.SetLineno(455)
						Label7:
							// line 456: self.write("%c%c%c" % (BININT2, obj&0xff, obj>>8))
							πF.SetLineno(456)
							πTemp003 = πF.MakeArgs(1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßBININT2); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.And(πF, µobj, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.RShift(πF, µobj, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple3(πTemp006, πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%c%c%c").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 457: return
							πF.SetLineno(457)
							πR = πg.None
							continue
							goto Label8
						Label8:
							goto Label4
						Label4:
							// line 459: high_bits = obj >> 31  # note that Python shift sign-extends
							πF.SetLineno(459)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.RShift(πF, µobj, πg.NewInt(31).ToObject()); πE != nil {
								continue
							}
							µhigh_bits = πTemp001
							if πE = πg.CheckLocal(πF, µhigh_bits, "high_bits"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µhigh_bits, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp004
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µhigh_bits, "high_bits"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Eq(πF, µhigh_bits, πTemp006); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label9:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label10
							}
							goto Label11
							// line 460: if high_bits == 0 or high_bits == -1:
							πF.SetLineno(460)
						Label10:
							// line 463: self.write(BININT + pack("<i", obj))
							πF.SetLineno(463)
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßBININT); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr("<i").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp005[1] = µobj
							if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
								continue
							}
							if πTemp006, πE = µpack.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 464: return
							πF.SetLineno(464)
							πR = πg.None
							continue
							goto Label11
						Label11:
							goto Label2
						Label2:
							// line 466: self.write(INT + repr(obj) + '\n')
							πF.SetLineno(466)
							πTemp003 = πF.MakeArgs(1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßINT); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp005[0] = µobj
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.Add(πF, πTemp006, πTemp008); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_int.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 467: dispatch[IntType] = save_int
					πF.SetLineno(467)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßsave_int); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πTemp016); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßIntType); πE != nil {
						continue
					}
					πTemp019 = πTemp020
					if πE = πg.SetItem(πF, πTemp018, πTemp019, πTemp017); πE != nil {
						continue
					}
					// line 469: def save_long(self, obj, pack=struct.pack):
					πF.SetLineno(469)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßpack, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "pack", Def: πTemp018}
					πTemp016 = πg.NewFunction(πg.NewCode("save_long", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µpack *πg.Object = πArgs[2]
						_ = µpack
						var µbytes *πg.Object = πg.UnboundLocal
						_ = µbytes
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
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
						var πTemp008 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, πTemp002, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 470: if self.proto >= 2:
							πF.SetLineno(470)
						Label1:
							// line 471: bytes = encode_long(obj)
							πF.SetLineno(471)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πTemp001, πE = πg.ResolveGlobal(πF, ßencode_long); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µbytes = πTemp002
							// line 472: n = len(bytes)
							πF.SetLineno(472)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbytes, "bytes"); πE != nil {
								continue
							}
							πTemp004[0] = µbytes
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µn = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µn, πg.NewInt(256).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 473: if n < 256:
							πF.SetLineno(473)
						Label3:
							// line 474: self.write(LONG1 + chr(n) + bytes)
							πF.SetLineno(474)
							πTemp004 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßLONG1); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp006[0] = µn
							if πTemp007, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbytes, "bytes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µbytes); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label5
						Label4:
							// line 476: self.write(LONG4 + pack("<i", n) + bytes)
							πF.SetLineno(476)
							πTemp004 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßLONG4); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewStr("<i").ToObject()
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp006[1] = µn
							if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
								continue
							}
							if πTemp007, πE = µpack.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Add(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µbytes, "bytes"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µbytes); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label5
						Label5:
							// line 477: return
							πF.SetLineno(477)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 478: self.write(LONG + repr(obj) + '\n')
							πF.SetLineno(478)
							πTemp004 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßLONG); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp006[0] = µobj
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp002, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_long.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 479: dispatch[LongType] = save_long
					πF.SetLineno(479)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßsave_long); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πTemp017); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßLongType); πE != nil {
						continue
					}
					πTemp020 = πTemp021
					if πE = πg.SetItem(πF, πTemp019, πTemp020, πTemp018); πE != nil {
						continue
					}
					// line 481: def save_float(self, obj, pack=struct.pack):
					πF.SetLineno(481)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp019, πE = πg.GetAttr(πF, πTemp018, ßpack, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "pack", Def: πTemp019}
					πTemp017 = πg.NewFunction(πg.NewCode("save_float", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µpack *πg.Object = πArgs[2]
						_ = µpack
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 482: if self.bin:
							πF.SetLineno(482)
						Label1:
							// line 483: self.write(BINFLOAT + pack('>d', obj))
							πF.SetLineno(483)
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßBINFLOAT); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(2)
							πTemp005[0] = πg.NewStr(">d").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp005[1] = µobj
							if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
								continue
							}
							if πTemp006, πE = µpack.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label2:
							// line 485: self.write(FLOAT + repr(obj) + '\n')
							πF.SetLineno(485)
							πTemp003 = πF.MakeArgs(1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßFLOAT); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp005[0] = µobj
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.Add(πF, πTemp006, πTemp008); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
					if πE = πClass.SetItem(πF, ßsave_float.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 486: dispatch[FloatType] = save_float
					πF.SetLineno(486)
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßsave_float); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πTemp018); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßFloatType); πE != nil {
						continue
					}
					πTemp021 = πTemp022
					if πE = πg.SetItem(πF, πTemp020, πTemp021, πTemp019); πE != nil {
						continue
					}
					// line 488: def save_string(self, obj, pack=struct.pack):
					πF.SetLineno(488)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßpack, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "pack", Def: πTemp020}
					πTemp018 = πg.NewFunction(πg.NewCode("save_string", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µpack *πg.Object = πArgs[2]
						_ = µpack
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 489: if self.bin:
							πF.SetLineno(489)
						Label1:
							// line 490: n = len(obj)
							πF.SetLineno(490)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp003[0] = µobj
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µn = πTemp004
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µn, πg.NewInt(256).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label4
							}
							goto Label5
							// line 491: if n < 256:
							πF.SetLineno(491)
						Label4:
							// line 492: self.write(SHORT_BINSTRING + chr(n) + obj)
							πF.SetLineno(492)
							πTemp003 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßSHORT_BINSTRING); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp006[0] = µn
							if πTemp007, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, µobj); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label6
						Label5:
							// line 494: self.write(BINSTRING + pack("<i", n) + obj)
							πF.SetLineno(494)
							πTemp003 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßBINSTRING); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewStr("<i").ToObject()
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp006[1] = µn
							if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
								continue
							}
							if πTemp007, πE = µpack.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, µobj); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label6
						Label6:
							goto Label3
						Label2:
							// line 496: self.write(STRING + repr(obj) + '\n')
							πF.SetLineno(496)
							πTemp003 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßSTRING); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp006[0] = µobj
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label3:
							// line 497: self.memoize(obj)
							πF.SetLineno(497)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp003[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_string.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 498: dispatch[StringType] = save_string
					πF.SetLineno(498)
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßsave_string); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πTemp019); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßStringType); πE != nil {
						continue
					}
					πTemp022 = πTemp023
					if πE = πg.SetItem(πF, πTemp021, πTemp022, πTemp020); πE != nil {
						continue
					}
					// line 500: def save_unicode(self, obj, pack=struct.pack):
					πF.SetLineno(500)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp021, πE = πg.GetAttr(πF, πTemp020, ßpack, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "pack", Def: πTemp021}
					πTemp019 = πg.NewFunction(πg.NewCode("save_unicode", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µpack *πg.Object = πArgs[2]
						_ = µpack
						var µencoding *πg.Object = πg.UnboundLocal
						_ = µencoding
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
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
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 501: if self.bin:
							πF.SetLineno(501)
						Label1:
							// line 502: encoding = obj.encode('utf-8')
							πF.SetLineno(502)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("utf-8").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ßencode, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µencoding = πTemp004
							// line 503: n = len(encoding)
							πF.SetLineno(503)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							πTemp003[0] = µencoding
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µn = πTemp004
							// line 504: self.write(BINUNICODE + pack("<i", n) + encoding)
							πF.SetLineno(504)
							πTemp003 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßBINUNICODE); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(2)
							πTemp006[0] = πg.NewStr("<i").ToObject()
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp006[1] = µn
							if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
								continue
							}
							if πTemp007, πE = µpack.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, µencoding); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label2:
							// line 506: obj = obj.replace("\\", "\\u005c")
							πF.SetLineno(506)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("\\").ToObject()
							πTemp003[1] = πg.NewStr("\\u005c").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µobj = πTemp004
							// line 507: obj = obj.replace("\n", "\\u000a")
							πF.SetLineno(507)
							πTemp003 = πF.MakeArgs(2)
							πTemp003[0] = πg.NewStr("\n").ToObject()
							πTemp003[1] = πg.NewStr("\\u000a").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µobj = πTemp004
							// line 508: self.write(UNICODE + obj.encode('raw-unicode-escape') + '\n')
							πF.SetLineno(508)
							πTemp003 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßUNICODE); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("raw-unicode-escape").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µobj, ßencode, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp004, πE = πg.Add(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label3:
							// line 509: self.memoize(obj)
							πF.SetLineno(509)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp003[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_unicode.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 510: dispatch[UnicodeType] = save_unicode
					πF.SetLineno(510)
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßsave_unicode); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πTemp020); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßUnicodeType); πE != nil {
						continue
					}
					πTemp023 = πTemp024
					if πE = πg.SetItem(πF, πTemp022, πTemp023, πTemp021); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßStringType); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßUnicodeType); πE != nil {
						continue
					}
					πTemp020 = πg.GetBool(πTemp021 == πTemp022).ToObject()
					if πTemp025, πE = πg.IsTrue(πF, πTemp020); πE != nil {
						continue
					}
					if πTemp025 {
						goto Label1
					}
					goto Label2
					// line 512: if StringType is UnicodeType:
					πF.SetLineno(512)
				Label1:
					// line 514: def save_string(self, obj, pack=struct.pack):
					πF.SetLineno(514)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp022, πE = πg.GetAttr(πF, πTemp021, ßpack, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "pack", Def: πTemp022}
					πTemp020 = πg.NewFunction(πg.NewCode("save_string", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µpack *πg.Object = πArgs[2]
						_ = µpack
						var µunicode *πg.Object = πg.UnboundLocal
						_ = µunicode
						var µl *πg.Object = πg.UnboundLocal
						_ = µl
						var µs *πg.Object = πg.UnboundLocal
						_ = µs
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							default:
								panic("unexpected function state")
							}
							// line 515: unicode = obj.isunicode()
							πF.SetLineno(515)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ßisunicode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µunicode = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µunicode, "unicode"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µunicode); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 517: if self.bin:
							πF.SetLineno(517)
						Label1:
							if πE = πg.CheckLocal(πF, µunicode, "unicode"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µunicode); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label5
							}
							goto Label6
							// line 518: if unicode:
							πF.SetLineno(518)
						Label5:
							// line 519: obj = obj.encode("utf-8")
							πF.SetLineno(519)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("utf-8").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µobj = πTemp002
							goto Label6
						Label6:
							// line 520: l = len(obj)
							πF.SetLineno(520)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µl = πTemp002
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.LT(πF, µl, πg.NewInt(256).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µunicode, "unicode"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µunicode); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							πTemp001 = πTemp002
						Label7:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label8
							}
							goto Label9
							// line 521: if l < 256 and not unicode:
							πF.SetLineno(521)
						Label8:
							// line 522: self.write(SHORT_BINSTRING + chr(l) + obj)
							πF.SetLineno(522)
							πTemp004 = πF.MakeArgs(1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßSHORT_BINSTRING); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp007[0] = µl
							if πTemp008, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.Add(πF, πTemp006, πTemp009); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µobj); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label10
						Label9:
							// line 524: s = pack("<i", l)
							πF.SetLineno(524)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("<i").ToObject()
							if πE = πg.CheckLocal(πF, µl, "l"); πE != nil {
								continue
							}
							πTemp004[1] = µl
							if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
								continue
							}
							if πTemp001, πE = µpack.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µs = πTemp001
							if πE = πg.CheckLocal(πF, µunicode, "unicode"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µunicode); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label11
							}
							goto Label12
							// line 525: if unicode:
							πF.SetLineno(525)
						Label11:
							// line 526: self.write(BINUNICODE + s + obj)
							πF.SetLineno(526)
							πTemp004 = πF.MakeArgs(1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßBINUNICODE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp006, µs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µobj); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label13
						Label12:
							// line 528: self.write(BINSTRING + s + obj)
							πF.SetLineno(528)
							πTemp004 = πF.MakeArgs(1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßBINSTRING); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp006, µs); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, µobj); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label13
						Label13:
							goto Label10
						Label10:
							goto Label4
							// line 530: if unicode:
							πF.SetLineno(530)
						Label2:
							// line 531: obj = obj.replace("\\", "\\u005c")
							πF.SetLineno(531)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("\\").ToObject()
							πTemp004[1] = πg.NewStr("\\u005c").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µobj = πTemp002
							// line 532: obj = obj.replace("\n", "\\u000a")
							πF.SetLineno(532)
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("\n").ToObject()
							πTemp004[1] = πg.NewStr("\\u000a").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µobj = πTemp002
							// line 533: obj = obj.encode('raw-unicode-escape')
							πF.SetLineno(533)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("raw-unicode-escape").ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ßencode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µobj = πTemp002
							// line 534: self.write(UNICODE + obj + '\n')
							πF.SetLineno(534)
							πTemp004 = πF.MakeArgs(1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßUNICODE); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp006, µobj); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label4
						Label3:
							// line 536: self.write(STRING + repr(obj) + '\n')
							πF.SetLineno(536)
							πTemp004 = πF.MakeArgs(1)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßSTRING); πE != nil {
								continue
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp007[0] = µobj
							if πTemp008, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp002, πE = πg.Add(πF, πTemp006, πTemp009); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label4
						Label4:
							// line 537: self.memoize(obj)
							πF.SetLineno(537)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_string.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 538: dispatch[StringType] = save_string
					πF.SetLineno(538)
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßsave_string); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πTemp021); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßStringType); πE != nil {
						continue
					}
					πTemp024 = πTemp026
					if πE = πg.SetItem(πF, πTemp023, πTemp024, πTemp022); πE != nil {
						continue
					}
					goto Label2
				Label2:
					// line 540: def save_tuple(self, obj):
					πF.SetLineno(540)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp021 = πg.NewFunction(πg.NewCode("save_tuple", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
						var µproto *πg.Object = πg.UnboundLocal
						_ = µproto
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µsave *πg.Object = πg.UnboundLocal
						_ = µsave
						var µmemo *πg.Object = πg.UnboundLocal
						_ = µmemo
						var µelement *πg.Object = πg.UnboundLocal
						_ = µelement
						var µget *πg.Object = πg.UnboundLocal
						_ = µget
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							case 16:
								goto Label16
							case 9:
								goto Label9
							case 10:
								goto Label10
							case 15:
								goto Label15
							default:
								panic("unexpected function state")
							}
							// line 541: write = self.write
							πF.SetLineno(541)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							// line 542: proto = self.proto
							πF.SetLineno(542)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							µproto = πTemp001
							// line 544: n = len(obj)
							πF.SetLineno(544)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µn = πTemp003
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µn, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 545: if n == 0:
							πF.SetLineno(545)
						Label1:
							if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µproto); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 546: if proto:
							πF.SetLineno(546)
						Label3:
							// line 547: write(EMPTY_TUPLE)
							πF.SetLineno(547)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEMPTY_TUPLE); πE != nil {
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
							goto Label5
						Label4:
							// line 549: write(MARK + TUPLE)
							πF.SetLineno(549)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßTUPLE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
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
							goto Label5
						Label5:
							// line 550: return
							πF.SetLineno(550)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 552: save = self.save
							πF.SetLineno(552)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsave, nil); πE != nil {
								continue
							}
							µsave = πTemp001
							// line 553: memo = self.memo
							πF.SetLineno(553)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							µmemo = πTemp001
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.LE(πF, µn, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GE(πF, µproto, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label6:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 554: if n <= 3 and proto >= 2:
							πF.SetLineno(554)
						Label7:
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µobj); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp004 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label11
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
								µelement = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(9)
							// line 556: save(element)
							πF.SetLineno(556)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πTemp002[0] = µelement
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp003, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µmemo, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label12
							}
							goto Label13
							// line 558: if id(obj) in memo:
							πF.SetLineno(558)
						Label12:
							// line 559: get = self.get(memo[id(obj)][0])
							πF.SetLineno(559)
							πTemp002 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp007[0] = µobj
							if πTemp008, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp005 = πTemp009
							if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µmemo, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µget = πTemp003
							// line 560: write(POP * n + get)
							πF.SetLineno(560)
							πTemp002 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßPOP); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp005, µn); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µget, "get"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, µget); πE != nil {
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
							goto Label14
						Label13:
							// line 562: write(_tuplesize2code[n])
							πF.SetLineno(562)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001 = µn
							if πTemp005, πE = πg.ResolveGlobal(πF, ß_tuplesize2code); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 563: self.memoize(obj)
							πF.SetLineno(563)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label14
						Label14:
							// line 564: return
							πF.SetLineno(564)
							πR = πg.None
							continue
							goto Label8
						Label8:
							// line 568: write(MARK)
							πF.SetLineno(568)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
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
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µobj); πE != nil {
								continue
							}
							πF.PushCheckpoint(16)
							πTemp004 = false
						Label15:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label17
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
								µelement = πTemp003
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(15)
							// line 570: save(element)
							πF.SetLineno(570)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µelement, "element"); πE != nil {
								continue
							}
							πTemp002[0] = µelement
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp003, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label16:
							if πE != nil || πR != nil {
								continue
							}
						Label17:
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πTemp003, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µmemo, πTemp005); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label18
							}
							goto Label19
							// line 572: if id(obj) in memo:
							πF.SetLineno(572)
						Label18:
							// line 580: get = self.get(memo[id(obj)][0])
							πF.SetLineno(580)
							πTemp002 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp007[0] = µobj
							if πTemp008, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp005 = πTemp009
							if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µmemo, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp008, πTemp001); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßget, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µget = πTemp003
							if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µproto); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label20
							}
							goto Label21
							// line 581: if proto:
							πF.SetLineno(581)
						Label20:
							// line 582: write(POP_MARK + get)
							πF.SetLineno(582)
							πTemp002 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßPOP_MARK); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µget, "get"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, µget); πE != nil {
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
							goto Label22
						Label21:
							// line 584: write(POP * (n+1) + get)
							πF.SetLineno(584)
							πTemp002 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßPOP); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Add(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mul(πF, πTemp005, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µget, "get"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, µget); πE != nil {
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
							goto Label22
						Label22:
							// line 585: return
							πF.SetLineno(585)
							πR = πg.None
							continue
							goto Label19
						Label19:
							// line 588: self.write(TUPLE)
							πF.SetLineno(588)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTUPLE); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 589: self.memoize(obj)
							πF.SetLineno(589)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_tuple.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 591: dispatch[TupleType] = save_tuple
					πF.SetLineno(591)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßsave_tuple); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πTemp022); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßTupleType); πE != nil {
						continue
					}
					πTemp026 = πTemp027
					if πE = πg.SetItem(πF, πTemp024, πTemp026, πTemp023); πE != nil {
						continue
					}
					// line 596: def save_empty_tuple(self, obj):
					πF.SetLineno(596)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("save_empty_tuple", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
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
							// line 597: self.write(EMPTY_TUPLE)
							πF.SetLineno(597)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEMPTY_TUPLE); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_empty_tuple.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 599: def save_list(self, obj):
					πF.SetLineno(599)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("save_list", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
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
						var πTemp006 []*πg.Object
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
							// line 600: write = self.write
							πF.SetLineno(600)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							goto Label2
							// line 602: if self.bin:
							πF.SetLineno(602)
						Label1:
							// line 603: write(EMPTY_LIST)
							πF.SetLineno(603)
							πTemp003 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEMPTY_LIST); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label2:
							// line 605: write(MARK + LIST)
							πF.SetLineno(605)
							πTemp003 = πF.MakeArgs(1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßLIST); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							goto Label3
						Label3:
							// line 607: self.memoize(obj)
							πF.SetLineno(607)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp003[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 608: self._batch_appends(iter(obj))
							πF.SetLineno(608)
							πTemp003 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp006[0] = µobj
							if πTemp001, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_batch_appends, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_list.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 610: dispatch[ListType] = save_list
					πF.SetLineno(610)
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßsave_list); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πTemp024); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßListType); πE != nil {
						continue
					}
					πTemp028 = πTemp029
					if πE = πg.SetItem(πF, πTemp027, πTemp028, πTemp026); πE != nil {
						continue
					}
					// line 614: _BATCHSIZE = 1000
					πF.SetLineno(614)
					if πE = πClass.SetItem(πF, ß_BATCHSIZE.ToObject(), πg.NewInt(1000).ToObject()); πE != nil {
						continue
					}
					// line 616: def _batch_appends(self, items):
					πF.SetLineno(616)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "items", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("_batch_appends", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitems *πg.Object = πArgs[1]
						_ = µitems
						var µsave *πg.Object = πg.UnboundLocal
						_ = µsave
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
						var µx *πg.Object = πg.UnboundLocal
						_ = µx
						var µr *πg.Object = πg.UnboundLocal
						_ = µr
						var µtmp *πg.Object = πg.UnboundLocal
						_ = µtmp
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
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
						var πTemp006 bool
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
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
							case 9:
								goto Label9
							case 10:
								goto Label10
							case 13:
								goto Label13
							case 18:
								goto Label18
							case 19:
								goto Label19
							default:
								panic("unexpected function state")
							}
							// line 618: save = self.save
							πF.SetLineno(618)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsave, nil); πE != nil {
								continue
							}
							µsave = πTemp001
							// line 619: write = self.write
							πF.SetLineno(619)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
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
							// line 621: if not self.bin:
							πF.SetLineno(621)
						Label1:
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µitems); πE != nil {
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
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								µx = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 623: save(x)
							πF.SetLineno(623)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp005[0] = µx
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp002, πE = µsave.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 624: write(APPEND)
							πF.SetLineno(624)
							πTemp005 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAPPEND); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp002, πE = µwrite.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 625: return
							πF.SetLineno(625)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 627: r = xrange(self._BATCHSIZE)
							πF.SetLineno(627)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_BATCHSIZE, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µr = πTemp002
							// line 628: while items is not None:
							πF.SetLineno(628)
							πF.PushCheckpoint(7)
							πTemp003 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µitems != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 629: tmp = []
							πF.SetLineno(629)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							µtmp = πTemp001
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µr); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp004 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label11
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
							πF.PushCheckpoint(9)
							// line 631: try:
							πF.SetLineno(631)
							πF.PushCheckpoint(13)
							// line 632: x = items.next()
							πF.SetLineno(632)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µitems, ßnext, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µx = πTemp007
							// line 633: tmp.append(x)
							πF.SetLineno(633)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp005[0] = µx
							if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtmp, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label14
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 634: except StopIteration:
							πF.SetLineno(634)
						Label14:
							// line 635: items = None
							πF.SetLineno(635)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µitems = πTemp002
							// line 636: break
							πF.SetLineno(636)
							πTemp004 = true
							continue
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							// line 637: n = len(tmp)
							πF.SetLineno(637)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
								continue
							}
							πTemp005[0] = µtmp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µn = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µn); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							goto Label17
							// line 638: if n > 1:
							πF.SetLineno(638)
						Label15:
							// line 639: write(MARK)
							πF.SetLineno(639)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtmp); πE != nil {
								continue
							}
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
								µx = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(18)
							// line 641: save(x)
							πF.SetLineno(641)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
								continue
							}
							πTemp005[0] = µx
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp002, πE = µsave.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							continue
						Label19:
							if πE != nil || πR != nil {
								continue
							}
						Label20:
							// line 642: write(APPENDS)
							πF.SetLineno(642)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAPPENDS); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label17
							// line 643: elif n:
							πF.SetLineno(643)
						Label16:
							// line 644: save(tmp[0])
							πF.SetLineno(644)
							πTemp005 = πF.MakeArgs(1)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtmp, πTemp001); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 645: write(APPEND)
							πF.SetLineno(645)
							πTemp005 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAPPEND); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							goto Label17
						Label17:
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
					if πE = πClass.SetItem(πF, ß_batch_appends.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 648: def save_dict(self, obj):
					πF.SetLineno(648)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("save_dict", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µmodict_saver *πg.Object = πg.UnboundLocal
						_ = µmodict_saver
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
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
							// line 649: modict_saver = self._pickle_maybe_moduledict(obj)
							πF.SetLineno(649)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_pickle_maybe_moduledict, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmodict_saver = πTemp003
							if πE = πg.CheckLocal(πF, µmodict_saver, "modict_saver"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µmodict_saver != πTemp003).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 650: if modict_saver is not None:
							πF.SetLineno(650)
						Label1:
							// line 651: return self.save_reduce(*modict_saver)
							πF.SetLineno(651)
							if πE = πg.CheckLocal(πF, µmodict_saver, "modict_saver"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsave_reduce, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, µmodict_saver, nil, nil); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label2
						Label2:
							// line 653: write = self.write
							πF.SetLineno(653)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							goto Label4
							// line 655: if self.bin:
							πF.SetLineno(655)
						Label3:
							// line 656: write(EMPTY_DICT)
							πF.SetLineno(656)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEMPTY_DICT); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp002, πE = µwrite.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label4:
							// line 658: write(MARK + DICT)
							πF.SetLineno(658)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßDICT); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp002, πE = µwrite.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							goto Label5
						Label5:
							// line 660: self.memoize(obj)
							πF.SetLineno(660)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 661: self._batch_setitems(obj.iteritems())
							πF.SetLineno(661)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µobj, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_batch_setitems, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_dict.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 663: dispatch[DictionaryType] = save_dict
					πF.SetLineno(663)
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßsave_dict); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp027); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßDictionaryType); πE != nil {
						continue
					}
					πTemp030 = πTemp031
					if πE = πg.SetItem(πF, πTemp029, πTemp030, πTemp028); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßPyStringMap); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp028 = πg.GetBool(πTemp029 == πTemp030).ToObject()
					if πTemp025, πE = πg.IsTrue(πF, πTemp028); πE != nil {
						continue
					}
					πTemp027 = πg.GetBool(!πTemp025).ToObject()
					if πTemp025, πE = πg.IsTrue(πF, πTemp027); πE != nil {
						continue
					}
					if πTemp025 {
						goto Label3
					}
					goto Label4
					// line 664: if not PyStringMap is None:
					πF.SetLineno(664)
				Label3:
					// line 665: dispatch[PyStringMap] = save_dict
					πF.SetLineno(665)
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßsave_dict); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp027); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßPyStringMap); πE != nil {
						continue
					}
					πTemp030 = πTemp031
					if πE = πg.SetItem(πF, πTemp029, πTemp030, πTemp028); πE != nil {
						continue
					}
					goto Label4
				Label4:
					// line 667: def _batch_setitems(self, items):
					πF.SetLineno(667)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "items", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("_batch_setitems", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µitems *πg.Object = πArgs[1]
						_ = µitems
						var µsave *πg.Object = πg.UnboundLocal
						_ = µsave
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
						var µk *πg.Object = πg.UnboundLocal
						_ = µk
						var µv *πg.Object = πg.UnboundLocal
						_ = µv
						var µr *πg.Object = πg.UnboundLocal
						_ = µr
						var µtmp *πg.Object = πg.UnboundLocal
						_ = µtmp
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
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
							case 9:
								goto Label9
							case 10:
								goto Label10
							case 13:
								goto Label13
							case 18:
								goto Label18
							case 19:
								goto Label19
							default:
								panic("unexpected function state")
							}
							// line 669: save = self.save
							πF.SetLineno(669)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsave, nil); πE != nil {
								continue
							}
							µsave = πTemp001
							// line 670: write = self.write
							πF.SetLineno(670)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
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
							// line 672: if not self.bin:
							πF.SetLineno(672)
						Label1:
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µitems); πE != nil {
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
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µk = πTemp005
								µv = πTemp006
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 674: save(k)
							πF.SetLineno(674)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp007[0] = µk
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp002, πE = µsave.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 675: save(v)
							πF.SetLineno(675)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp007[0] = µv
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp002, πE = µsave.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 676: write(SETITEM)
							πF.SetLineno(676)
							πTemp007 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßSETITEM); πE != nil {
								continue
							}
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp002, πE = µwrite.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 677: return
							πF.SetLineno(677)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 679: r = xrange(self._BATCHSIZE)
							πF.SetLineno(679)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_BATCHSIZE, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µr = πTemp002
							// line 680: while items is not None:
							πF.SetLineno(680)
							πF.PushCheckpoint(7)
							πTemp003 = false
						Label6:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
								πF.PopCheckpoint()
								goto Label8
							}
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µitems != πTemp002).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(6)
							// line 681: tmp = []
							πF.SetLineno(681)
							πTemp007 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp007...).ToObject()
							µtmp = πTemp001
							if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µr); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp004 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label11
							}
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µi = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(9)
							// line 683: try:
							πF.SetLineno(683)
							πF.PushCheckpoint(13)
							// line 684: tmp.append(items.next())
							πF.SetLineno(684)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µitems, ßnext, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp005
							if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µtmp, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πF.PopCheckpoint()
							goto Label12
						Label13:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label14
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 685: except StopIteration:
							πF.SetLineno(685)
						Label14:
							// line 686: items = None
							πF.SetLineno(686)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µitems = πTemp002
							// line 687: break
							πF.SetLineno(687)
							πTemp004 = true
							continue
							πF.RestoreExc(nil, nil)
							goto Label12
						Label12:
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							// line 688: n = len(tmp)
							πF.SetLineno(688)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
								continue
							}
							πTemp007[0] = µtmp
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							µn = πTemp002
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µn); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							goto Label17
							// line 689: if n > 1:
							πF.SetLineno(689)
						Label15:
							// line 690: write(MARK)
							πF.SetLineno(690)
							πTemp007 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µtmp); πE != nil {
								continue
							}
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µk = πTemp005
								µv = πTemp006
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(18)
							// line 692: save(k)
							πF.SetLineno(692)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp007[0] = µk
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp002, πE = µsave.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 693: save(v)
							πF.SetLineno(693)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp007[0] = µv
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp002, πE = µsave.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							continue
						Label19:
							if πE != nil || πR != nil {
								continue
							}
						Label20:
							// line 694: write(SETITEMS)
							πF.SetLineno(694)
							πTemp007 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSETITEMS); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label17
							// line 695: elif n:
							πF.SetLineno(695)
						Label16:
							// line 696: k, v = tmp[0]
							πF.SetLineno(696)
							πTemp001 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µtmp, "tmp"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µtmp, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}}}, πTemp002); πE != nil {
								continue
							}
							µk = πTemp001
							µv = πTemp005
							// line 697: save(k)
							πF.SetLineno(697)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp007[0] = µk
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 698: save(v)
							πF.SetLineno(698)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp007[0] = µv
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							// line 699: write(SETITEM)
							πF.SetLineno(699)
							πTemp007 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßSETITEM); πE != nil {
								continue
							}
							πTemp007[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							goto Label17
						Label17:
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
					if πE = πClass.SetItem(πF, ß_batch_setitems.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 702: def _pickle_maybe_moduledict(self, obj):
					πF.SetLineno(702)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("_pickle_maybe_moduledict", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µthemodule *πg.Object = πg.UnboundLocal
						_ = µthemodule
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
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
							// line 704: try:
							πF.SetLineno(704)
							πF.PushCheckpoint(2)
							// line 705: name = obj['__name__']
							πF.SetLineno(705)
							πTemp001 = ß__name__.ToObject()
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µobj, πTemp001); πE != nil {
								continue
							}
							µname = πTemp002
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp003[0] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004 != πTemp002).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label3
							}
							goto Label4
							// line 706: if type(name) is not str:
							πF.SetLineno(706)
						Label3:
							// line 707: return None
							πF.SetLineno(707)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label4
						Label4:
							// line 708: themodule = sys.modules[name]
							πF.SetLineno(708)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = µname
							if πTemp004, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µthemodule = πTemp002
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µthemodule, "themodule"); πE != nil {
								continue
							}
							πTemp003[0] = µthemodule
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßModuleType); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004 != πTemp002).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label5
							}
							goto Label6
							// line 709: if type(themodule) is not ModuleType:
							πF.SetLineno(709)
						Label5:
							// line 710: return None
							πF.SetLineno(710)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label6
						Label6:
							if πE = πg.CheckLocal(πF, µthemodule, "themodule"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µthemodule, ß__dict__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002 != µobj).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 711: if themodule.__dict__ is not obj:
							πF.SetLineno(711)
						Label7:
							// line 712: return None
							πF.SetLineno(712)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label8
						Label8:
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp004, πTemp006).ToObject()
							if πTemp005, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label9
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 713: except (AttributeError, KeyError, TypeError):
							πF.SetLineno(713)
						Label9:
							// line 714: return None
							πF.SetLineno(714)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 716: return getattr, (themodule, '__dict__')
							πF.SetLineno(716)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µthemodule, "themodule"); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(µthemodule, ß__dict__.ToObject()).ToObject()
							πTemp001 = πg.NewTuple2(πTemp002, πTemp004).ToObject()
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
					if πE = πClass.SetItem(πF, ß_pickle_maybe_moduledict.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 719: def save_inst(self, obj):
					πF.SetLineno(719)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("save_inst", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µcls *πg.Object = πg.UnboundLocal
						_ = µcls
						var µmemo *πg.Object = πg.UnboundLocal
						_ = µmemo
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
						var µsave *πg.Object = πg.UnboundLocal
						_ = µsave
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
						var µarg *πg.Object = πg.UnboundLocal
						_ = µarg
						var µgetstate *πg.Object = πg.UnboundLocal
						_ = µgetstate
						var µstuff *πg.Object = πg.UnboundLocal
						_ = µstuff
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
						var πTemp010 *πg.BaseException
						_ = πTemp010
						var πTemp011 *πg.Traceback
						_ = πTemp011
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 8:
								goto Label8
							case 10:
								goto Label10
							case 11:
								goto Label11
							case 14:
								goto Label14
							case 7:
								goto Label7
							default:
								panic("unexpected function state")
							}
							// line 720: cls = obj.__class__
							πF.SetLineno(720)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ß__class__, nil); πE != nil {
								continue
							}
							µcls = πTemp001
							// line 722: memo  = self.memo
							πF.SetLineno(722)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							µmemo = πTemp001
							// line 723: write = self.write
							πF.SetLineno(723)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							// line 724: save  = self.save
							πF.SetLineno(724)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßsave, nil); πE != nil {
								continue
							}
							µsave = πTemp001
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							πTemp002[1] = ß__getinitargs__.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 726: if hasattr(obj, '__getinitargs__'):
							πF.SetLineno(726)
						Label1:
							// line 727: args = obj.__getinitargs__()
							πF.SetLineno(727)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ß__getinitargs__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µargs = πTemp003
							// line 728: len(args) # XXX Assert it's a sequence
							πF.SetLineno(728)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 729: _keep_alive(args, memo)
							πF.SetLineno(729)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							πTemp002[0] = µargs
							if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
								continue
							}
							πTemp002[1] = µmemo
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_keep_alive); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label3
						Label2:
							// line 731: args = ()
							πF.SetLineno(731)
							πTemp001 = πg.NewTuple0().ToObject()
							µargs = πTemp001
							goto Label3
						Label3:
							// line 733: write(MARK)
							πF.SetLineno(733)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßMARK); πE != nil {
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
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßbin, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							goto Label5
							// line 735: if self.bin:
							πF.SetLineno(735)
						Label4:
							// line 736: save(cls)
							πF.SetLineno(736)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp002[0] = µcls
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µargs); πE != nil {
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
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µarg = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(7)
							// line 738: save(arg)
							πF.SetLineno(738)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							πTemp002[0] = µarg
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp003, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label8:
							if πE != nil || πR != nil {
								continue
							}
						Label9:
							// line 739: write(OBJ)
							πF.SetLineno(739)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßOBJ); πE != nil {
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
							goto Label6
						Label5:
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µargs); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp004 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label12
							}
							if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µarg = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(10)
							// line 742: save(arg)
							πF.SetLineno(742)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µarg, "arg"); πE != nil {
								continue
							}
							πTemp002[0] = µarg
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp003, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							// line 743: write(INST + cls.__module__ + '\n' + cls.__name__ + '\n')
							πF.SetLineno(743)
							πTemp002 = πF.MakeArgs(1)
							if πTemp008, πE = πg.ResolveGlobal(πF, ßINST); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetAttr(πF, µcls, ß__module__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Add(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp007, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µcls, ß__name__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp006, πTemp007); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp003, πg.NewStr("\n").ToObject()); πE != nil {
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
							goto Label6
						Label6:
							// line 745: self.memoize(obj)
							πF.SetLineno(745)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp002[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 747: try:
							πF.SetLineno(747)
							πF.PushCheckpoint(14)
							// line 748: getstate = obj.__getstate__
							πF.SetLineno(748)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ß__getstate__, nil); πE != nil {
								continue
							}
							µgetstate = πTemp001
							πF.PopCheckpoint()
							// line 752: stuff = getstate()
							πF.SetLineno(752)
							if πE = πg.CheckLocal(πF, µgetstate, "getstate"); πE != nil {
								continue
							}
							if πTemp001, πE = µgetstate.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstuff = πTemp001
							// line 753: _keep_alive(stuff, memo)
							πF.SetLineno(753)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstuff, "stuff"); πE != nil {
								continue
							}
							πTemp002[0] = µstuff
							if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
								continue
							}
							πTemp002[1] = µmemo
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_keep_alive); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							goto Label13
						Label14:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 749: except AttributeError:
							πF.SetLineno(749)
						Label15:
							// line 750: stuff = obj.__dict__
							πF.SetLineno(750)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ß__dict__, nil); πE != nil {
								continue
							}
							µstuff = πTemp001
							πF.RestoreExc(nil, nil)
							goto Label13
						Label13:
							// line 754: save(stuff)
							πF.SetLineno(754)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstuff, "stuff"); πE != nil {
								continue
							}
							πTemp002[0] = µstuff
							if πE = πg.CheckLocal(πF, µsave, "save"); πE != nil {
								continue
							}
							if πTemp001, πE = µsave.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 755: write(BUILD)
							πF.SetLineno(755)
							πTemp002 = πF.MakeArgs(1)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßBUILD); πE != nil {
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
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsave_inst.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 757: dispatch[InstanceType] = save_inst
					πF.SetLineno(757)
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßsave_inst); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp031}, πTemp030); πE != nil {
						continue
					}
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßInstanceType); πE != nil {
						continue
					}
					πTemp033 = πTemp034
					if πE = πg.SetItem(πF, πTemp032, πTemp033, πTemp031); πE != nil {
						continue
					}
					// line 759: def save_function(self, obj):
					πF.SetLineno(759)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("save_function", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µe *πg.Object = πg.UnboundLocal
						_ = µe
						var µreduce *πg.Object = πg.UnboundLocal
						_ = µreduce
						var µrv *πg.Object = πg.UnboundLocal
						_ = µrv
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 πg.KWArgs
						_ = πTemp008
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
							// line 760: try:
							πF.SetLineno(760)
							πF.PushCheckpoint(2)
							// line 761: return self.save_global(obj)
							πF.SetLineno(761)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsave_global, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πR = πTemp003
							continue
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsInstance(πF, πTemp004.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 762: except PicklingError, e:
							πF.SetLineno(762)
						Label3:
							µe = πTemp004.ToObject()
							// line 763: pass
							πF.SetLineno(763)
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 765: reduce = dispatch_table.get(type(obj))
							πF.SetLineno(765)
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp007[0] = µobj
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdispatch_table); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µreduce = πTemp002
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µreduce); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 766: if reduce:
							πF.SetLineno(766)
						Label4:
							// line 767: rv = reduce(obj)
							πF.SetLineno(767)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp002, πE = µreduce.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrv = πTemp002
							goto Label6
						Label5:
							// line 770: reduce = getattr(obj, "__reduce_ex__", None)
							πF.SetLineno(770)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							πTemp001[1] = ß__reduce_ex__.ToObject()
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
							µreduce = πTemp003
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µreduce); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 771: if reduce:
							πF.SetLineno(771)
						Label7:
							// line 772: rv = reduce(self.proto)
							πF.SetLineno(772)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp002, πE = µreduce.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrv = πTemp002
							goto Label9
						Label8:
							// line 774: reduce = getattr(obj, "__reduce__", None)
							πF.SetLineno(774)
							πTemp001 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							πTemp001[1] = ß__reduce__.ToObject()
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
							µreduce = πTemp003
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µreduce); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label10
							}
							goto Label11
							// line 775: if reduce:
							πF.SetLineno(775)
						Label10:
							// line 776: rv = reduce()
							πF.SetLineno(776)
							if πE = πg.CheckLocal(πF, µreduce, "reduce"); πE != nil {
								continue
							}
							if πTemp002, πE = µreduce.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrv = πTemp002
							goto Label12
						Label11:
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							// line 778: raise e
							πF.SetLineno(778)
							πE = πF.Raise(µe, nil, nil)
							continue
							goto Label12
						Label12:
							goto Label9
						Label9:
							goto Label6
						Label6:
							// line 779: return self.save_reduce(obj=obj, *rv)
							πF.SetLineno(779)
							if πE = πg.CheckLocal(πF, µrv, "rv"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp008 = πg.KWArgs{
								{"obj", µobj},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßsave_reduce, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, µrv, πTemp008, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_function.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 780: dispatch[FunctionType] = save_function
					πF.SetLineno(780)
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßsave_function); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp032}, πTemp031); πE != nil {
						continue
					}
					if πTemp033, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßFunctionType); πE != nil {
						continue
					}
					πTemp034 = πTemp035
					if πE = πg.SetItem(πF, πTemp033, πTemp034, πTemp032); πE != nil {
						continue
					}
					// line 782: def save_global(self, obj, name=None, pack=struct.pack):
					πF.SetLineno(782)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "obj", Def: nil}
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "name", Def: πTemp032}
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp033, πE = πg.GetAttr(πF, πTemp032, ßpack, nil); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "pack", Def: πTemp033}
					πTemp031 = πg.NewFunction(πg.NewCode("save_global", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µobj *πg.Object = πArgs[1]
						_ = µobj
						var µname *πg.Object = πArgs[2]
						_ = µname
						var µpack *πg.Object = πArgs[3]
						_ = µpack
						var µwrite *πg.Object = πg.UnboundLocal
						_ = µwrite
						var µmemo *πg.Object = πg.UnboundLocal
						_ = µmemo
						var µmodule *πg.Object = πg.UnboundLocal
						_ = µmodule
						var µmod *πg.Object = πg.UnboundLocal
						_ = µmod
						var µklass *πg.Object = πg.UnboundLocal
						_ = µklass
						var µcode *πg.Object = πg.UnboundLocal
						_ = µcode
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
						var πTemp010 *πg.Object
						_ = πTemp010
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6:
								goto Label6
							default:
								panic("unexpected function state")
							}
							// line 783: write = self.write
							πF.SetLineno(783)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwrite, nil); πE != nil {
								continue
							}
							µwrite = πTemp001
							// line 784: memo = self.memo
							πF.SetLineno(784)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							µmemo = πTemp001
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µname == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 786: if name is None:
							πF.SetLineno(786)
						Label1:
							// line 787: name = obj.__name__
							πF.SetLineno(787)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µobj, ß__name__, nil); πE != nil {
								continue
							}
							µname = πTemp001
							goto Label2
						Label2:
							// line 789: module = getattr(obj, "__module__", None)
							πF.SetLineno(789)
							πTemp004 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							πTemp004[1] = ß__module__.ToObject()
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
							µmodule = πTemp002
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µmodule == πTemp002).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 790: if module is None:
							πF.SetLineno(790)
						Label3:
							// line 791: module = whichmodule(obj, name)
							πF.SetLineno(791)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[1] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßwhichmodule); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µmodule = πTemp002
							goto Label4
						Label4:
							// line 793: try:
							πF.SetLineno(793)
							πF.PushCheckpoint(6)
							// line 794: __import__(module)
							πF.SetLineno(794)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp004[0] = µmodule
							if πTemp001, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 795: mod = sys.modules[module]
							πF.SetLineno(795)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp001 = µmodule
							if πTemp005, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßmodules, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µmod = πTemp002
							// line 796: klass = getattr(mod, name)
							πF.SetLineno(796)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							πTemp004[0] = µmod
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp004[1] = µname
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µklass = πTemp002
							πF.PopCheckpoint()
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µklass != µobj).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 802: if klass is not obj:
							πF.SetLineno(802)
						Label7:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µobj, µmodule, µname).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Can't pickle %r: it's not the same object as %s.%s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 803: raise PicklingError(
							πF.SetLineno(803)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label8
						Label8:
							goto Label5
						Label6:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßAttributeError); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple3(πTemp002, πTemp005, πTemp006).ToObject()
							if πTemp003, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label9
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 797: except (ImportError, KeyError, AttributeError):
							πF.SetLineno(797)
						Label9:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(µobj, µmodule, µname).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("Can't pickle %r: it's not found as %s.%s").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßPicklingError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 798: raise PicklingError(
							πF.SetLineno(798)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßproto, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GE(πF, πTemp002, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label10
							}
							goto Label11
							// line 807: if self.proto >= 2:
							πF.SetLineno(807)
						Label10:
							// line 808: code = _extension_registry.get((module, name))
							πF.SetLineno(808)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(µmodule, µname).ToObject()
							πTemp004[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_extension_registry); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßget, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µcode = πTemp001
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µcode); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							goto Label13
							// line 809: if code:
							πF.SetLineno(809)
						Label12:
							// line 810: assert code > 0
							πF.SetLineno(810)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µcode, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, µcode, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, µcode, πg.NewInt(65535).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label15
							}
							goto Label16
							// line 811: if code <= 0xff:
							πF.SetLineno(811)
						Label14:
							// line 812: write(EXT1 + chr(code))
							πF.SetLineno(812)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEXT1); πE != nil {
								continue
							}
							πTemp009 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp009[0] = µcode
							if πTemp005, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp006); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label17
							// line 813: elif code <= 0xffff:
							πF.SetLineno(813)
						Label15:
							// line 814: write("%c%c%c" % (EXT2, code&0xff, code>>8))
							πF.SetLineno(814)
							πTemp004 = πF.MakeArgs(1)
							if πTemp005, πE = πg.ResolveGlobal(πF, ßEXT2); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.And(πF, µcode, πg.NewInt(255).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.RShift(πF, µcode, πg.NewInt(8).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp005, πTemp006, πTemp010).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("%c%c%c").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label17
						Label16:
							// line 816: write(EXT4 + pack("<i", code))
							πF.SetLineno(816)
							πTemp004 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßEXT4); πE != nil {
								continue
							}
							πTemp009 = πF.MakeArgs(2)
							πTemp009[0] = πg.NewStr("<i").ToObject()
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp009[1] = µcode
							if πE = πg.CheckLocal(πF, µpack, "pack"); πE != nil {
								continue
							}
							if πTemp005, πE = µpack.Call(πF, πTemp009, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp009)
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label17
						Label17:
							// line 817: return
							πF.SetLineno(817)
							πR = πg.None
							continue
							goto Label13
						Label13:
							goto Label11
						Label11:
							// line 819: write(GLOBAL + module + '\n' + name + '\n')
							πF.SetLineno(819)
							πTemp004 = πF.MakeArgs(1)
							if πTemp010, πE = πg.ResolveGlobal(πF, ßGLOBAL); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Add(πF, πTemp010, µmodule); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, πTemp006, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp005, µname); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µwrite, "write"); πE != nil {
								continue
							}
							if πTemp001, πE = µwrite.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 820: self.memoize(obj)
							πF.SetLineno(820)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp004[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmemoize, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßsave_global.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 822: dispatch[ClassType] = save_global
					πF.SetLineno(822)
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßsave_global); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp033}, πTemp032); πE != nil {
						continue
					}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßClassType); πE != nil {
						continue
					}
					πTemp035 = πTemp036
					if πE = πg.SetItem(πF, πTemp034, πTemp035, πTemp033); πE != nil {
						continue
					}
					// line 823: dispatch[BuiltinFunctionType] = save_global
					πF.SetLineno(823)
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßsave_global); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp033}, πTemp032); πE != nil {
						continue
					}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßBuiltinFunctionType); πE != nil {
						continue
					}
					πTemp035 = πTemp036
					if πE = πg.SetItem(πF, πTemp034, πTemp035, πTemp033); πE != nil {
						continue
					}
					// line 824: dispatch[TypeType] = save_global
					πF.SetLineno(824)
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßsave_global); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp033}, πTemp032); πE != nil {
						continue
					}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßTypeType); πE != nil {
						continue
					}
					πTemp035 = πTemp036
					if πE = πg.SetItem(πF, πTemp034, πTemp035, πTemp033); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp005 == nil {
				πTemp005 = πg.TypeType.ToObject()
			}
			if πTemp010, πE = πTemp005.Call(πF, []*πg.Object{πg.NewStr("Pickler").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPickler.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 828: def _keep_alive(x, memo):
			πF.SetLineno(828)
			πTemp009 = make([]πg.Param, 2)
			πTemp009[0] = πg.Param{Name: "x", Def: nil}
			πTemp009[1] = πg.Param{Name: "memo", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("_keep_alive", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]
				_ = µx
				var µmemo *πg.Object = πArgs[1]
				_ = µmemo
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
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
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
					// line 829: """Keeps a reference to the object x in the memo.
					πF.SetLineno(829)
					// line 838: try:
					πF.SetLineno(838)
					πF.PushCheckpoint(2)
					// line 839: memo[id(memo)].append(x)
					πF.SetLineno(839)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp003[0] = µmemo
					if πTemp004, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp002 = πTemp005
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µmemo, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp004, ßappend, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
					// line 840: except KeyError:
					πF.SetLineno(840)
				Label3:
					// line 842: memo[id(memo)]=[x]
					πF.SetLineno(842)
					πTemp001 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µmemo, "memo"); πE != nil {
						continue
					}
					πTemp001[0] = µmemo
					if πTemp009, πE = πg.ResolveGlobal(πF, ßid); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp005 = πTemp010
					if πE = πg.SetItem(πF, µmemo, πTemp005, πTemp004); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ß_keep_alive.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 829: """Keeps a reference to the object x in the memo.
			πF.SetLineno(829)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Keeps a reference to the object x in the memo.\n\n    Because we remember objects by their id, we have\n    to assure that possibly temporary objects are kept\n    alive by referencing them.\n    We store a reference at the id of the memo, which should\n    normally not be used unless someone tries to deepcopy\n    the memo itself...\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ß_keep_alive); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 848: classmap = {} # called classmap for backwards compatibility
			πF.SetLineno(848)
			πTemp004 = πg.NewDict()
			πTemp005 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ßclassmap.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 850: def whichmodule(func, funcname):
			πF.SetLineno(850)
			πTemp009 = make([]πg.Param, 2)
			πTemp009[0] = πg.Param{Name: "func", Def: nil}
			πTemp009[1] = πg.Param{Name: "funcname", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("whichmodule", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfunc *πg.Object = πArgs[0]
				_ = µfunc
				var µfuncname *πg.Object = πArgs[1]
				_ = µfuncname
				var µmod *πg.Object = πg.UnboundLocal
				_ = µmod
				var µname *πg.Object = πg.UnboundLocal
				_ = µname
				var µmodule *πg.Object = πg.UnboundLocal
				_ = µmodule
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
				var πTemp008 *πg.Object
				_ = πTemp008
				var πR *πg.Object
				_ = πR
				var πE *πg.BaseException
				_ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5:
						goto Label5
					case 6:
						goto Label6
					default:
						panic("unexpected function state")
					}
					// line 851: """Figure out the module in which a function occurs.
					πF.SetLineno(851)
					// line 859: mod = getattr(func, "__module__", None)
					πF.SetLineno(859)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp001[0] = µfunc
					πTemp001[1] = ß__module__.ToObject()
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
					µmod = πTemp003
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µmod != πTemp003).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 860: if mod is not None:
					πF.SetLineno(860)
				Label1:
					// line 861: return mod
					πF.SetLineno(861)
					if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
						continue
					}
					πR = µmod
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßclassmap); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, πTemp003, µfunc); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 862: if func in classmap:
					πF.SetLineno(862)
				Label3:
					// line 863: return classmap[func]
					πF.SetLineno(863)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp002 = µfunc
					if πTemp005, πE = πg.ResolveGlobal(πF, ßclassmap); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
						continue
					}
					πR = πTemp003
					continue
					goto Label4
				Label4:
					if πTemp003, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßmodules, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp005, ßitems, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
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
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
							continue
						}
						µname = πTemp005
						µmodule = πTemp007
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(5)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µmodule == πTemp005).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					goto Label9
					// line 866: if module is None:
					πF.SetLineno(866)
				Label8:
					// line 867: continue # skip dummy package entries
					πF.SetLineno(867)
					continue
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.NE(πF, µname, ß__main__.ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label10
					}
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
						continue
					}
					πTemp001[0] = µmodule
					if πE = πg.CheckLocal(πF, µfuncname, "funcname"); πE != nil {
						continue
					}
					πTemp001[1] = µfuncname
					if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[2] = πTemp007
					if πTemp007, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp008 == µfunc).ToObject()
					πTemp003 = πTemp005
				Label10:
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					goto Label12
					// line 868: if name != '__main__' and getattr(module, funcname, None) is func:
					πF.SetLineno(868)
				Label11:
					// line 869: break
					πF.SetLineno(869)
					πTemp004 = true
					continue
					goto Label12
				Label12:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
					// line 871: name = '__main__'
					πF.SetLineno(871)
					µname = ß__main__.ToObject()
				Label7:
					// line 872: classmap[func] = name
					πF.SetLineno(872)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µname); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßclassmap); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
						continue
					}
					πTemp005 = µfunc
					if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
						continue
					}
					// line 873: return name
					πF.SetLineno(873)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πR = µname
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßwhichmodule.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 851: """Figure out the module in which a function occurs.
			πF.SetLineno(851)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Figure out the module in which a function occurs.\n\n    Search sys.modules for the module.\n    Cache in classmap.\n    Return a module name.\n    If the function cannot be found, return \"__main__\".\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßwhichmodule); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 878: class Unpickler(object):
			πF.SetLineno(878)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp012
			πTemp004 = πg.NewDict()
			if πTemp010, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp010); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Unpickler", "/usr/lib/python2.7/pickle.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 *πg.Dict
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
				var πTemp063 *πg.Object
				_ = πTemp063
				var πTemp064 *πg.Object
				_ = πTemp064
				var πTemp065 *πg.Object
				_ = πTemp065
				var πTemp066 *πg.Object
				_ = πTemp066
				var πTemp067 *πg.Object
				_ = πTemp067
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 880: def __init__(self, file):
					πF.SetLineno(880)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "file", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µfile *πg.Object = πArgs[1]
						_ = µfile
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Dict
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
							// line 881: """This takes a file-like object for reading a pickle data stream.
							πF.SetLineno(881)
							// line 892: self.readline = file.readline
							πF.SetLineno(892)
							if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfile, ßreadline, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreadline, πTemp002); πE != nil {
								continue
							}
							// line 893: self.read = file.read
							πF.SetLineno(893)
							if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfile, ßread, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßread, πTemp002); πE != nil {
								continue
							}
							// line 894: self.memo = {}
							πF.SetLineno(894)
							πTemp003 = πg.NewDict()
							πTemp001 = πTemp003.ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßmemo, πTemp002); πE != nil {
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
					// line 881: """This takes a file-like object for reading a pickle data stream.
					πF.SetLineno(881)
					// line 881: """This takes a file-like object for reading a pickle data stream.
					πF.SetLineno(881)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("This takes a file-like object for reading a pickle data stream.\n\n        The protocol version of the pickle is detected automatically, so no\n        proto argument is needed.\n\n        The file-like object must have two methods, a read() method that\n        takes an integer argument, and a readline() method that requires no\n        arguments.  Both methods should return a string.  Thus file-like\n        object can be a file object opened for reading, a StringIO object,\n        or any other custom object that meets this interface.\n        ").ToObject()); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("This takes a file-like object for reading a pickle data stream.\n\n        The protocol version of the pickle is detected automatically, so no\n        proto argument is needed.\n\n        The file-like object must have two methods, a read() method that\n        takes an integer argument, and a readline() method that requires no\n        arguments.  Both methods should return a string.  Thus file-like\n        object can be a file object opened for reading, a StringIO object,\n        or any other custom object that meets this interface.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 896: def load(self):
					πF.SetLineno(896)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("load", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µread *πg.Object = πg.UnboundLocal
						_ = µread
						var µdispatch *πg.Object = πg.UnboundLocal
						_ = µdispatch
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µstopinst *πg.Object = πg.UnboundLocal
						_ = µstopinst
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
						_ = πTemp007
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2:
								goto Label2
							case 3:
								goto Label3
							case 4:
								goto Label4
							default:
								panic("unexpected function state")
							}
							// line 897: """Read a pickled object representation from the open file.
							πF.SetLineno(897)
							// line 901: self.mark = object() # any new unique object
							πF.SetLineno(901)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßmark, πTemp001); πE != nil {
								continue
							}
							// line 902: self.stack = []
							πF.SetLineno(902)
							πTemp003 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstack, πTemp002); πE != nil {
								continue
							}
							// line 903: self.append = self.stack.append
							πF.SetLineno(903)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßappend, πTemp001); πE != nil {
								continue
							}
							// line 904: read = self.read
							πF.SetLineno(904)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							µread = πTemp001
							// line 905: dispatch = self.dispatch
							πF.SetLineno(905)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdispatch, nil); πE != nil {
								continue
							}
							µdispatch = πTemp001
							// line 906: try:
							πF.SetLineno(906)
							πF.PushCheckpoint(2)
							// line 907: while 1:
							πF.SetLineno(907)
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
							if πTemp005, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(3)
							// line 908: key = read(1)
							πF.SetLineno(908)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µread, "read"); πE != nil {
								continue
							}
							if πTemp001, πE = µread.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µkey = πTemp001
							// line 909: dispatch[key](self)
							πF.SetLineno(909)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp003[0] = µself
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001 = µkey
							if πE = πg.CheckLocal(πF, µdispatch, "dispatch"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µdispatch, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_Stop); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
							continue
							// line 910: except _Stop, stopinst:
							πF.SetLineno(910)
						Label6:
							µstopinst = πTemp006.ToObject()
							// line 911: return stopinst.value
							πF.SetLineno(911)
							if πE = πg.CheckLocal(πF, µstopinst, "stopinst"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstopinst, ßvalue, nil); πE != nil {
								continue
							}
							πR = πTemp001
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
					if πE = πClass.SetItem(πF, ßload.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 897: """Read a pickled object representation from the open file.
					πF.SetLineno(897)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Read a pickled object representation from the open file.\n\n        Return the reconstituted object hierarchy specified in the file.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßload); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 921: def marker(self):
					πF.SetLineno(921)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("marker", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstack *πg.Object = πg.UnboundLocal
						_ = µstack
						var µmark *πg.Object = πg.UnboundLocal
						_ = µmark
						var µk *πg.Object = πg.UnboundLocal
						_ = µk
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
							// line 922: stack = self.stack
							πF.SetLineno(922)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							µstack = πTemp001
							// line 923: mark = self.mark
							πF.SetLineno(923)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmark, nil); πE != nil {
								continue
							}
							µmark = πTemp001
							// line 924: k = len(stack)-1
							πF.SetLineno(924)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							πTemp002[0] = µstack
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp001, πE = πg.Sub(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µk = πTemp001
							// line 925: while stack[k] is not mark: k = k-1
							πF.SetLineno(925)
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
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003 = µk
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µstack, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmark, "mark"); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004 != µmark).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 925: while stack[k] is not mark: k = k-1
							πF.SetLineno(925)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µk, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µk = πTemp001
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 926: return k
							πF.SetLineno(926)
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
					if πE = πClass.SetItem(πF, ßmarker.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 928: dispatch = {}
					πF.SetLineno(928)
					πTemp006 = πg.NewDict()
					πTemp005 = πTemp006.ToObject()
					if πE = πClass.SetItem(πF, ßdispatch.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 930: def load_eof(self):
					πF.SetLineno(930)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("load_eof", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßEOFError); πE != nil {
								continue
							}
							// line 931: raise EOFError
							πF.SetLineno(931)
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
					if πE = πClass.SetItem(πF, ßload_eof.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 932: dispatch[''] = load_eof
					πF.SetLineno(932)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßload_eof); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp007); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					πTemp010 = ß.ToObject()
					if πE = πg.SetItem(πF, πTemp009, πTemp010, πTemp008); πE != nil {
						continue
					}
					// line 934: def load_proto(self):
					πF.SetLineno(934)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("load_proto", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µproto *πg.Object = πg.UnboundLocal
						_ = µproto
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
							// line 935: proto = ord(self.read(1))
							πF.SetLineno(935)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
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
							µproto = πTemp004
							if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µproto); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							if πTemp004, πE = πg.LE(πF, µproto, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
						Label1:
							if πTemp005, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 936: if not 0 <= proto <= 2:
							πF.SetLineno(936)
						Label2:
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µproto, "proto"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("unsupported pickle protocol: %d").ToObject(), µproto); πE != nil {
								continue
							}
							// line 937: raise ValueError, "unsupported pickle protocol: %d" % proto
							πF.SetLineno(937)
							πE = πF.Raise(πTemp003, πTemp004, nil)
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
					if πE = πClass.SetItem(πF, ßload_proto.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 938: dispatch[PROTO] = load_proto
					πF.SetLineno(938)
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ßload_proto); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πTemp008); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßPROTO); πE != nil {
						continue
					}
					πTemp011 = πTemp012
					if πE = πg.SetItem(πF, πTemp010, πTemp011, πTemp009); πE != nil {
						continue
					}
					// line 940: def load_persid(self):
					πF.SetLineno(940)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("load_persid", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpid *πg.Object = πg.UnboundLocal
						_ = µpid
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
							// line 941: pid = self.readline()[:-1]
							πF.SetLineno(941)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µpid = πTemp002
							// line 942: self.append(self.persistent_load(pid))
							πF.SetLineno(942)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
								continue
							}
							πTemp006[0] = µpid
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpersistent_load, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_persid.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 943: dispatch[PERSID] = load_persid
					πF.SetLineno(943)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ßload_persid); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πTemp009); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßPERSID); πE != nil {
						continue
					}
					πTemp012 = πTemp013
					if πE = πg.SetItem(πF, πTemp011, πTemp012, πTemp010); πE != nil {
						continue
					}
					// line 945: def load_binpersid(self):
					πF.SetLineno(945)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("load_binpersid", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µpid *πg.Object = πg.UnboundLocal
						_ = µpid
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
							// line 946: pid = self.stack.pop()
							πF.SetLineno(946)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µpid = πTemp001
							// line 947: self.append(self.persistent_load(pid))
							πF.SetLineno(947)
							πTemp003 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpid, "pid"); πE != nil {
								continue
							}
							πTemp004[0] = µpid
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßpersistent_load, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_binpersid.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 948: dispatch[BINPERSID] = load_binpersid
					πF.SetLineno(948)
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ßload_binpersid); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πTemp010); πE != nil {
						continue
					}
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßBINPERSID); πE != nil {
						continue
					}
					πTemp013 = πTemp014
					if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
						continue
					}
					// line 950: def load_none(self):
					πF.SetLineno(950)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp010 = πg.NewFunction(πg.NewCode("load_none", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 951: self.append(None)
							πF.SetLineno(951)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_none.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 952: dispatch[NONE] = load_none
					πF.SetLineno(952)
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ßload_none); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πTemp011); πE != nil {
						continue
					}
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßNONE); πE != nil {
						continue
					}
					πTemp014 = πTemp015
					if πE = πg.SetItem(πF, πTemp013, πTemp014, πTemp012); πE != nil {
						continue
					}
					// line 954: def load_false(self):
					πF.SetLineno(954)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp011 = πg.NewFunction(πg.NewCode("load_false", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 955: self.append(False)
							πF.SetLineno(955)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_false.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 956: dispatch[NEWFALSE] = load_false
					πF.SetLineno(956)
					if πTemp012, πE = πg.ResolveClass(πF, πClass, nil, ßload_false); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πTemp012); πE != nil {
						continue
					}
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßNEWFALSE); πE != nil {
						continue
					}
					πTemp015 = πTemp016
					if πE = πg.SetItem(πF, πTemp014, πTemp015, πTemp013); πE != nil {
						continue
					}
					// line 958: def load_true(self):
					πF.SetLineno(958)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp012 = πg.NewFunction(πg.NewCode("load_true", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 959: self.append(True)
							πF.SetLineno(959)
							πTemp001 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_true.ToObject(), πTemp012); πE != nil {
						continue
					}
					// line 960: dispatch[NEWTRUE] = load_true
					πF.SetLineno(960)
					if πTemp013, πE = πg.ResolveClass(πF, πClass, nil, ßload_true); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp014}, πTemp013); πE != nil {
						continue
					}
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßNEWTRUE); πE != nil {
						continue
					}
					πTemp016 = πTemp017
					if πE = πg.SetItem(πF, πTemp015, πTemp016, πTemp014); πE != nil {
						continue
					}
					// line 962: def load_int(self):
					πF.SetLineno(962)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp013 = πg.NewFunction(πg.NewCode("load_int", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µdata *πg.Object = πg.UnboundLocal
						_ = µdata
						var µval *πg.Object = πg.UnboundLocal
						_ = µval
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 bool
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 *πg.BaseException
						_ = πTemp007
						var πTemp008 *πg.Traceback
						_ = πTemp008
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 6:
								goto Label6
							default:
								panic("unexpected function state")
							}
							// line 963: data = self.readline()
							πF.SetLineno(963)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µdata = πTemp002
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßFALSE); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µdata, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTRUE); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µdata, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 964: if data == FALSE[1:]:
							πF.SetLineno(964)
						Label1:
							// line 965: val = False
							πF.SetLineno(965)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µval = πTemp001
							goto Label4
							// line 966: elif data == TRUE[1:]:
							πF.SetLineno(966)
						Label2:
							// line 967: val = True
							πF.SetLineno(967)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µval = πTemp001
							goto Label4
						Label3:
							// line 969: try:
							πF.SetLineno(969)
							πF.PushCheckpoint(6)
							// line 970: val = int(data)
							πF.SetLineno(970)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µval = πTemp002
							πF.PopCheckpoint()
							goto Label5
						Label6:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp007, πTemp008 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
							continue
							// line 971: except ValueError:
							πF.SetLineno(971)
						Label7:
							// line 972: val = long(data)
							πF.SetLineno(972)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp006[0] = µdata
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							µval = πTemp002
							πF.RestoreExc(nil, nil)
							goto Label5
						Label5:
							goto Label4
						Label4:
							// line 973: self.append(val)
							πF.SetLineno(973)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µval, "val"); πE != nil {
								continue
							}
							πTemp006[0] = µval
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_int.ToObject(), πTemp013); πE != nil {
						continue
					}
					// line 974: dispatch[INT] = load_int
					πF.SetLineno(974)
					if πTemp014, πE = πg.ResolveClass(πF, πClass, nil, ßload_int); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp015}, πTemp014); πE != nil {
						continue
					}
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßINT); πE != nil {
						continue
					}
					πTemp017 = πTemp018
					if πE = πg.SetItem(πF, πTemp016, πTemp017, πTemp015); πE != nil {
						continue
					}
					// line 976: def load_binint(self):
					πF.SetLineno(976)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp014 = πg.NewFunction(πg.NewCode("load_binint", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
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
							// line 977: self.append(mloads('i' + self.read(4)))
							πF.SetLineno(977)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Add(πF, ßi.ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmloads); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_binint.ToObject(), πTemp014); πE != nil {
						continue
					}
					// line 978: dispatch[BININT] = load_binint
					πF.SetLineno(978)
					if πTemp015, πE = πg.ResolveClass(πF, πClass, nil, ßload_binint); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp016}, πTemp015); πE != nil {
						continue
					}
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßBININT); πE != nil {
						continue
					}
					πTemp018 = πTemp019
					if πE = πg.SetItem(πF, πTemp017, πTemp018, πTemp016); πE != nil {
						continue
					}
					// line 980: def load_binint1(self):
					πF.SetLineno(980)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp015 = πg.NewFunction(πg.NewCode("load_binint1", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 981: self.append(ord(self.read(1)))
							πF.SetLineno(981)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp002[0] = πTemp005
							if πTemp004, πE = πg.ResolveGlobal(πF, ßord); πE != nil {
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
							if πTemp004, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_binint1.ToObject(), πTemp015); πE != nil {
						continue
					}
					// line 982: dispatch[BININT1] = load_binint1
					πF.SetLineno(982)
					if πTemp016, πE = πg.ResolveClass(πF, πClass, nil, ßload_binint1); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πTemp016); πE != nil {
						continue
					}
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßBININT1); πE != nil {
						continue
					}
					πTemp019 = πTemp020
					if πE = πg.SetItem(πF, πTemp018, πTemp019, πTemp017); πE != nil {
						continue
					}
					// line 984: def load_binint2(self):
					πF.SetLineno(984)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp016 = πg.NewFunction(πg.NewCode("load_binint2", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
						var πTemp006 *πg.Object
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
							default:
								panic("unexpected function state")
							}
							// line 985: self.append(mloads('i' + self.read(2) + '\000\000'))
							πF.SetLineno(985)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp004, πE = πg.Add(πF, ßi.ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, πTemp004, πg.NewStr("\x00\x00").ToObject()); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmloads); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_binint2.ToObject(), πTemp016); πE != nil {
						continue
					}
					// line 986: dispatch[BININT2] = load_binint2
					πF.SetLineno(986)
					if πTemp017, πE = πg.ResolveClass(πF, πClass, nil, ßload_binint2); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp018}, πTemp017); πE != nil {
						continue
					}
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßBININT2); πE != nil {
						continue
					}
					πTemp020 = πTemp021
					if πE = πg.SetItem(πF, πTemp019, πTemp020, πTemp018); πE != nil {
						continue
					}
					// line 988: def load_long(self):
					πF.SetLineno(988)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp017 = πg.NewFunction(πg.NewCode("load_long", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 989: self.append(long(self.readline()[:-1], 0))
							πF.SetLineno(989)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewInt(0).ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_long.ToObject(), πTemp017); πE != nil {
						continue
					}
					// line 990: dispatch[LONG] = load_long
					πF.SetLineno(990)
					if πTemp018, πE = πg.ResolveClass(πF, πClass, nil, ßload_long); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πTemp018); πE != nil {
						continue
					}
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßLONG); πE != nil {
						continue
					}
					πTemp021 = πTemp022
					if πE = πg.SetItem(πF, πTemp020, πTemp021, πTemp019); πE != nil {
						continue
					}
					// line 992: def load_long1(self):
					πF.SetLineno(992)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp018 = πg.NewFunction(πg.NewCode("load_long1", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µbytes *πg.Object = πg.UnboundLocal
						_ = µbytes
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
							// line 993: n = ord(self.read(1))
							πF.SetLineno(993)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
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
							// line 994: bytes = self.read(n)
							πF.SetLineno(994)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µbytes = πTemp004
							// line 995: self.append(decode_long(bytes))
							πF.SetLineno(995)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbytes, "bytes"); πE != nil {
								continue
							}
							πTemp002[0] = µbytes
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdecode_long); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_long1.ToObject(), πTemp018); πE != nil {
						continue
					}
					// line 996: dispatch[LONG1] = load_long1
					πF.SetLineno(996)
					if πTemp019, πE = πg.ResolveClass(πF, πClass, nil, ßload_long1); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πTemp019); πE != nil {
						continue
					}
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßLONG1); πE != nil {
						continue
					}
					πTemp022 = πTemp023
					if πE = πg.SetItem(πF, πTemp021, πTemp022, πTemp020); πE != nil {
						continue
					}
					// line 998: def load_long4(self):
					πF.SetLineno(998)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp019 = πg.NewFunction(πg.NewCode("load_long4", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µn *πg.Object = πg.UnboundLocal
						_ = µn
						var µbytes *πg.Object = πg.UnboundLocal
						_ = µbytes
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
							// line 999: n = mloads('i' + self.read(4))
							πF.SetLineno(999)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, ßi.ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmloads); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µn = πTemp004
							// line 1000: bytes = self.read(n)
							πF.SetLineno(1000)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp001[0] = µn
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µbytes = πTemp004
							// line 1001: self.append(decode_long(bytes))
							πF.SetLineno(1001)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µbytes, "bytes"); πE != nil {
								continue
							}
							πTemp003[0] = µbytes
							if πTemp002, πE = πg.ResolveGlobal(πF, ßdecode_long); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_long4.ToObject(), πTemp019); πE != nil {
						continue
					}
					// line 1002: dispatch[LONG4] = load_long4
					πF.SetLineno(1002)
					if πTemp020, πE = πg.ResolveClass(πF, πClass, nil, ßload_long4); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πTemp020); πE != nil {
						continue
					}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßLONG4); πE != nil {
						continue
					}
					πTemp023 = πTemp024
					if πE = πg.SetItem(πF, πTemp022, πTemp023, πTemp021); πE != nil {
						continue
					}
					// line 1004: def load_float(self):
					πF.SetLineno(1004)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp020 = πg.NewFunction(πg.NewCode("load_float", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 1005: self.append(float(self.readline()[:-1]))
							πF.SetLineno(1005)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_float.ToObject(), πTemp020); πE != nil {
						continue
					}
					// line 1006: dispatch[FLOAT] = load_float
					πF.SetLineno(1006)
					if πTemp021, πE = πg.ResolveClass(πF, πClass, nil, ßload_float); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πTemp021); πE != nil {
						continue
					}
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßFLOAT); πE != nil {
						continue
					}
					πTemp024 = πTemp025
					if πE = πg.SetItem(πF, πTemp023, πTemp024, πTemp022); πE != nil {
						continue
					}
					// line 1008: def load_binfloat(self, unpack=struct.unpack):
					πF.SetLineno(1008)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßstruct); πE != nil {
						continue
					}
					if πTemp023, πE = πg.GetAttr(πF, πTemp022, ßunpack, nil); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "unpack", Def: πTemp023}
					πTemp021 = πg.NewFunction(πg.NewCode("load_binfloat", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µunpack *πg.Object = πArgs[1]
						_ = µunpack
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
						var πTemp007 *πg.Object
						_ = πTemp007
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
							// line 1009: self.append(unpack('>d', self.read(8))[0])
							πF.SetLineno(1009)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewInt(0).ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr(">d").ToObject()
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewInt(8).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							πTemp004[1] = πTemp007
							if πE = πg.CheckLocal(πF, µunpack, "unpack"); πE != nil {
								continue
							}
							if πTemp006, πE = µunpack.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_binfloat.ToObject(), πTemp021); πE != nil {
						continue
					}
					// line 1010: dispatch[BINFLOAT] = load_binfloat
					πF.SetLineno(1010)
					if πTemp022, πE = πg.ResolveClass(πF, πClass, nil, ßload_binfloat); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp023}, πTemp022); πE != nil {
						continue
					}
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßBINFLOAT); πE != nil {
						continue
					}
					πTemp025 = πTemp026
					if πE = πg.SetItem(πF, πTemp024, πTemp025, πTemp023); πE != nil {
						continue
					}
					// line 1012: def load_string(self):
					πF.SetLineno(1012)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp022 = πg.NewFunction(πg.NewCode("load_string", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µrep *πg.Object = πg.UnboundLocal
						_ = µrep
						var µq *πg.Object = πg.UnboundLocal
						_ = µq
						var πTemp001 *πg.Object
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
						var πTemp007 []*πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 []*πg.Object
						_ = πTemp011
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
							// line 1013: rep = self.readline()[:-1]
							πF.SetLineno(1013)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µrep = πTemp002
							if πTemp001, πE = πg.Iter(πF, πg.NewStr("\"'").ToObject()); πE != nil {
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
								µq = πTemp002
							}
							if πE != nil || !πTemp006 {
								continue
							}
							πF.PushCheckpoint(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp007[0] = µq
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µrep, ßstartswith, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 1015: if rep.startswith(q):
							πF.SetLineno(1015)
						Label4:
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							πTemp007[0] = µrep
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp003, πE = πg.LT(πF, πTemp008, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label6
							}
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp007[0] = µq
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µrep, ßendswith, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp009).ToObject()
							πTemp002 = πTemp003
						Label6:
							if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label7
							}
							goto Label8
							// line 1016: if len(rep) < 2 or not rep.endswith(q):
							πF.SetLineno(1016)
						Label7:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 1017: raise ValueError, "insecure string pickle"
							πF.SetLineno(1017)
							πE = πF.Raise(πTemp002, πg.NewStr("insecure string pickle").ToObject(), nil)
							continue
							goto Label8
						Label8:
							// line 1018: rep = rep[len(q):-len(q)]
							πF.SetLineno(1018)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp007[0] = µq
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
								continue
							}
							πTemp007[0] = µq
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp003, πE = πg.Neg(πF, πTemp010); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πTemp003, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, µrep, πTemp002); πE != nil {
								continue
							}
							µrep = πTemp003
							// line 1019: break
							πF.SetLineno(1019)
							πTemp005 = true
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							// line 1021: raise ValueError, "insecure string pickle"
							πF.SetLineno(1021)
							πE = πF.Raise(πTemp002, πg.NewStr("insecure string pickle").ToObject(), nil)
							continue
						Label3:
							// line 1022: self.append(rep.decode("string-escape"))
							πF.SetLineno(1022)
							πTemp007 = πF.MakeArgs(1)
							πTemp011 = πF.MakeArgs(1)
							πTemp011[0] = πg.NewStr("string-escape").ToObject()
							if πE = πg.CheckLocal(πF, µrep, "rep"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µrep, ßdecode, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp011, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp011)
							πTemp007[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßload_string.ToObject(), πTemp022); πE != nil {
						continue
					}
					// line 1023: dispatch[STRING] = load_string
					πF.SetLineno(1023)
					if πTemp023, πE = πg.ResolveClass(πF, πClass, nil, ßload_string); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp024}, πTemp023); πE != nil {
						continue
					}
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßSTRING); πE != nil {
						continue
					}
					πTemp026 = πTemp027
					if πE = πg.SetItem(πF, πTemp025, πTemp026, πTemp024); πE != nil {
						continue
					}
					// line 1025: def load_binstring(self):
					πF.SetLineno(1025)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp023 = πg.NewFunction(πg.NewCode("load_binstring", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlen *πg.Object = πg.UnboundLocal
						_ = µlen
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
							// line 1026: len = mloads('i' + self.read(4))
							πF.SetLineno(1026)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, ßi.ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmloads); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlen = πTemp004
							// line 1027: self.append(self.read(len))
							πF.SetLineno(1027)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							πTemp003[0] = µlen
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_binstring.ToObject(), πTemp023); πE != nil {
						continue
					}
					// line 1028: dispatch[BINSTRING] = load_binstring
					πF.SetLineno(1028)
					if πTemp024, πE = πg.ResolveClass(πF, πClass, nil, ßload_binstring); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp025}, πTemp024); πE != nil {
						continue
					}
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßBINSTRING); πE != nil {
						continue
					}
					πTemp027 = πTemp028
					if πE = πg.SetItem(πF, πTemp026, πTemp027, πTemp025); πE != nil {
						continue
					}
					// line 1030: def load_unicode(self):
					πF.SetLineno(1030)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp024 = πg.NewFunction(πg.NewCode("load_unicode", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 1031: self.append(unicode(self.readline()[:-1],'raw-unicode-escape'))
							πF.SetLineno(1031)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp002[0] = πTemp004
							πTemp002[1] = πg.NewStr("raw-unicode-escape").ToObject()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_unicode.ToObject(), πTemp024); πE != nil {
						continue
					}
					// line 1032: dispatch[UNICODE] = load_unicode
					πF.SetLineno(1032)
					if πTemp025, πE = πg.ResolveClass(πF, πClass, nil, ßload_unicode); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp026}, πTemp025); πE != nil {
						continue
					}
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßUNICODE); πE != nil {
						continue
					}
					πTemp028 = πTemp029
					if πE = πg.SetItem(πF, πTemp027, πTemp028, πTemp026); πE != nil {
						continue
					}
					// line 1034: def load_binunicode(self):
					πF.SetLineno(1034)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp025 = πg.NewFunction(πg.NewCode("load_binunicode", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlen *πg.Object = πg.UnboundLocal
						_ = µlen
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
						var πTemp006 []*πg.Object
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
							// line 1035: len = mloads('i' + self.read(4))
							πF.SetLineno(1035)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, ßi.ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmloads); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µlen = πTemp004
							// line 1036: self.append(unicode(self.read(len),'utf-8'))
							πF.SetLineno(1036)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(2)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							πTemp006[0] = µlen
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp003[0] = πTemp004
							πTemp003[1] = πg.NewStr("utf-8").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_binunicode.ToObject(), πTemp025); πE != nil {
						continue
					}
					// line 1037: dispatch[BINUNICODE] = load_binunicode
					πF.SetLineno(1037)
					if πTemp026, πE = πg.ResolveClass(πF, πClass, nil, ßload_binunicode); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp027}, πTemp026); πE != nil {
						continue
					}
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßBINUNICODE); πE != nil {
						continue
					}
					πTemp029 = πTemp030
					if πE = πg.SetItem(πF, πTemp028, πTemp029, πTemp027); πE != nil {
						continue
					}
					// line 1039: def load_short_binstring(self):
					πF.SetLineno(1039)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp026 = πg.NewFunction(πg.NewCode("load_short_binstring", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µlen *πg.Object = πg.UnboundLocal
						_ = µlen
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
							// line 1040: len = ord(self.read(1))
							πF.SetLineno(1040)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
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
							µlen = πTemp004
							// line 1041: self.append(self.read(len))
							πF.SetLineno(1041)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							πTemp002[0] = µlen
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_short_binstring.ToObject(), πTemp026); πE != nil {
						continue
					}
					// line 1042: dispatch[SHORT_BINSTRING] = load_short_binstring
					πF.SetLineno(1042)
					if πTemp027, πE = πg.ResolveClass(πF, πClass, nil, ßload_short_binstring); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp028}, πTemp027); πE != nil {
						continue
					}
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßSHORT_BINSTRING); πE != nil {
						continue
					}
					πTemp030 = πTemp031
					if πE = πg.SetItem(πF, πTemp029, πTemp030, πTemp028); πE != nil {
						continue
					}
					// line 1044: def load_tuple(self):
					πF.SetLineno(1044)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp027 = πg.NewFunction(πg.NewCode("load_tuple", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µk *πg.Object = πg.UnboundLocal
						_ = µk
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
							// line 1045: k = self.marker()
							πF.SetLineno(1045)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmarker, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µk = πTemp002
							// line 1046: self.stack[k:] = [tuple(self.stack[k+1:])]
							πF.SetLineno(1046)
							πTemp003 = make([]*πg.Object, 1)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µk, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
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
							πTemp003[0] = πTemp002
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µk, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_tuple.ToObject(), πTemp027); πE != nil {
						continue
					}
					// line 1047: dispatch[TUPLE] = load_tuple
					πF.SetLineno(1047)
					if πTemp028, πE = πg.ResolveClass(πF, πClass, nil, ßload_tuple); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp029}, πTemp028); πE != nil {
						continue
					}
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßTUPLE); πE != nil {
						continue
					}
					πTemp031 = πTemp032
					if πE = πg.SetItem(πF, πTemp030, πTemp031, πTemp029); πE != nil {
						continue
					}
					// line 1049: def load_empty_tuple(self):
					πF.SetLineno(1049)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp028 = πg.NewFunction(πg.NewCode("load_empty_tuple", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1050: self.stack.append(())
							πF.SetLineno(1050)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewTuple0().ToObject()
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_empty_tuple.ToObject(), πTemp028); πE != nil {
						continue
					}
					// line 1051: dispatch[EMPTY_TUPLE] = load_empty_tuple
					πF.SetLineno(1051)
					if πTemp029, πE = πg.ResolveClass(πF, πClass, nil, ßload_empty_tuple); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp030}, πTemp029); πE != nil {
						continue
					}
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp033, πE = πg.ResolveClass(πF, πClass, nil, ßEMPTY_TUPLE); πE != nil {
						continue
					}
					πTemp032 = πTemp033
					if πE = πg.SetItem(πF, πTemp031, πTemp032, πTemp030); πE != nil {
						continue
					}
					// line 1053: def load_tuple1(self):
					πF.SetLineno(1053)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp029 = πg.NewFunction(πg.NewCode("load_tuple1", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 1054: self.stack[-1] = (self.stack[-1],)
							πF.SetLineno(1054)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple1(πTemp003).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_tuple1.ToObject(), πTemp029); πE != nil {
						continue
					}
					// line 1055: dispatch[TUPLE1] = load_tuple1
					πF.SetLineno(1055)
					if πTemp030, πE = πg.ResolveClass(πF, πClass, nil, ßload_tuple1); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp031}, πTemp030); πE != nil {
						continue
					}
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßTUPLE1); πE != nil {
						continue
					}
					πTemp033 = πTemp034
					if πE = πg.SetItem(πF, πTemp032, πTemp033, πTemp031); πE != nil {
						continue
					}
					// line 1057: def load_tuple2(self):
					πF.SetLineno(1057)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp030 = πg.NewFunction(πg.NewCode("load_tuple2", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 1058: self.stack[-2:] = [(self.stack[-2], self.stack[-1])]
							πF.SetLineno(1058)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp004, πTemp005).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_tuple2.ToObject(), πTemp030); πE != nil {
						continue
					}
					// line 1059: dispatch[TUPLE2] = load_tuple2
					πF.SetLineno(1059)
					if πTemp031, πE = πg.ResolveClass(πF, πClass, nil, ßload_tuple2); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp032}, πTemp031); πE != nil {
						continue
					}
					if πTemp033, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßTUPLE2); πE != nil {
						continue
					}
					πTemp034 = πTemp035
					if πE = πg.SetItem(πF, πTemp033, πTemp034, πTemp032); πE != nil {
						continue
					}
					// line 1061: def load_tuple3(self):
					πF.SetLineno(1061)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp031 = πg.NewFunction(πg.NewCode("load_tuple3", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
						var πTemp007 *πg.Object
						_ = πTemp007
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
							// line 1062: self.stack[-3:] = [(self.stack[-3], self.stack[-2], self.stack[-1])]
							πF.SetLineno(1062)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp006
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp004, πTemp005, πTemp006).ToObject()
							πTemp001[0] = πTemp002
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(3).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_tuple3.ToObject(), πTemp031); πE != nil {
						continue
					}
					// line 1063: dispatch[TUPLE3] = load_tuple3
					πF.SetLineno(1063)
					if πTemp032, πE = πg.ResolveClass(πF, πClass, nil, ßload_tuple3); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp033}, πTemp032); πE != nil {
						continue
					}
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßTUPLE3); πE != nil {
						continue
					}
					πTemp035 = πTemp036
					if πE = πg.SetItem(πF, πTemp034, πTemp035, πTemp033); πE != nil {
						continue
					}
					// line 1065: def load_empty_list(self):
					πF.SetLineno(1065)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp032 = πg.NewFunction(πg.NewCode("load_empty_list", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 1066: self.stack.append([])
							πF.SetLineno(1066)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = make([]*πg.Object, 0)
							πTemp003 = πg.NewList(πTemp002...).ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_empty_list.ToObject(), πTemp032); πE != nil {
						continue
					}
					// line 1067: dispatch[EMPTY_LIST] = load_empty_list
					πF.SetLineno(1067)
					if πTemp033, πE = πg.ResolveClass(πF, πClass, nil, ßload_empty_list); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp034}, πTemp033); πE != nil {
						continue
					}
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßEMPTY_LIST); πE != nil {
						continue
					}
					πTemp036 = πTemp037
					if πE = πg.SetItem(πF, πTemp035, πTemp036, πTemp034); πE != nil {
						continue
					}
					// line 1069: def load_empty_dictionary(self):
					πF.SetLineno(1069)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp033 = πg.NewFunction(πg.NewCode("load_empty_dictionary", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Dict
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
							// line 1070: self.stack.append({})
							πF.SetLineno(1070)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πg.NewDict()
							πTemp003 = πTemp002.ToObject()
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_empty_dictionary.ToObject(), πTemp033); πE != nil {
						continue
					}
					// line 1071: dispatch[EMPTY_DICT] = load_empty_dictionary
					πF.SetLineno(1071)
					if πTemp034, πE = πg.ResolveClass(πF, πClass, nil, ßload_empty_dictionary); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp035}, πTemp034); πE != nil {
						continue
					}
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp038, πE = πg.ResolveClass(πF, πClass, nil, ßEMPTY_DICT); πE != nil {
						continue
					}
					πTemp037 = πTemp038
					if πE = πg.SetItem(πF, πTemp036, πTemp037, πTemp035); πE != nil {
						continue
					}
					// line 1073: def load_list(self):
					πF.SetLineno(1073)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp034 = πg.NewFunction(πg.NewCode("load_list", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µk *πg.Object = πg.UnboundLocal
						_ = µk
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
							// line 1074: k = self.marker()
							πF.SetLineno(1074)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmarker, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µk = πTemp002
							// line 1075: self.stack[k:] = [self.stack[k+1:]]
							πF.SetLineno(1075)
							πTemp003 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µk, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp001 = πg.NewList(πTemp003...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{µk, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp004, πTemp005, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_list.ToObject(), πTemp034); πE != nil {
						continue
					}
					// line 1076: dispatch[LIST] = load_list
					πF.SetLineno(1076)
					if πTemp035, πE = πg.ResolveClass(πF, πClass, nil, ßload_list); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp036}, πTemp035); πE != nil {
						continue
					}
					if πTemp037, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp039, πE = πg.ResolveClass(πF, πClass, nil, ßLIST); πE != nil {
						continue
					}
					πTemp038 = πTemp039
					if πE = πg.SetItem(πF, πTemp037, πTemp038, πTemp036); πE != nil {
						continue
					}
					// line 1078: def load_dict(self):
					πF.SetLineno(1078)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp035 = πg.NewFunction(πg.NewCode("load_dict", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µk *πg.Object = πg.UnboundLocal
						_ = µk
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
						var µitems *πg.Object = πg.UnboundLocal
						_ = µitems
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Dict
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []*πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
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
							case 1:
								goto Label1
							case 2:
								goto Label2
							default:
								panic("unexpected function state")
							}
							// line 1079: k = self.marker()
							πF.SetLineno(1079)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmarker, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µk = πTemp002
							// line 1080: d = {}
							πF.SetLineno(1080)
							πTemp003 = πg.NewDict()
							πTemp001 = πTemp003.ToObject()
							µd = πTemp001
							// line 1081: items = self.stack[k+1:]
							πF.SetLineno(1081)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µk, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µitems = πTemp002
							πTemp005 = πF.MakeArgs(3)
							πTemp005[0] = πg.NewInt(0).ToObject()
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp006[0] = µitems
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[1] = πTemp004
							πTemp005[2] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
								continue
							}
							πF.PushCheckpoint(2)
							πTemp007 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
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
								πTemp008 = !isStop
							} else {
								πTemp008 = true
								µi = πTemp002
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 1083: key = items[i]
							πF.SetLineno(1083)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002 = µi
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µitems, πTemp002); πE != nil {
								continue
							}
							µkey = πTemp004
							// line 1084: value = items[i+1]
							πF.SetLineno(1084)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, µitems, πTemp002); πE != nil {
								continue
							}
							µvalue = πTemp004
							// line 1085: d[key] = value
							πF.SetLineno(1085)
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
							πTemp004 = µkey
							if πE = πg.SetItem(πF, µd, πTemp004, πTemp002); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1086: self.stack[k:] = [d]
							πF.SetLineno(1086)
							πTemp005 = make([]*πg.Object, 1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp005[0] = µd
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.SliceType.Call(πF, πg.Args{µk, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.SetItem(πF, πTemp004, πTemp009, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_dict.ToObject(), πTemp035); πE != nil {
						continue
					}
					// line 1087: dispatch[DICT] = load_dict
					πF.SetLineno(1087)
					if πTemp036, πE = πg.ResolveClass(πF, πClass, nil, ßload_dict); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp037}, πTemp036); πE != nil {
						continue
					}
					if πTemp038, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp040, πE = πg.ResolveClass(πF, πClass, nil, ßDICT); πE != nil {
						continue
					}
					πTemp039 = πTemp040
					if πE = πg.SetItem(πF, πTemp038, πTemp039, πTemp037); πE != nil {
						continue
					}
					// line 1094: def _instantiate(self, klass, k):
					πF.SetLineno(1094)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "klass", Def: nil}
					πTemp002[2] = πg.Param{Name: "k", Def: nil}
					πTemp036 = πg.NewFunction(πg.NewCode("_instantiate", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µklass *πg.Object = πArgs[1]
						_ = µklass
						var µk *πg.Object = πArgs[2]
						_ = µk
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
						var µinstantiated *πg.Object = πg.UnboundLocal
						_ = µinstantiated
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µerr *πg.Object = πg.UnboundLocal
						_ = µerr
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
						var πTemp008 *πg.BaseException
						_ = πTemp008
						var πTemp009 *πg.Traceback
						_ = πTemp009
						var πTemp010 *πg.Object
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 10:
								goto Label10
							case 5:
								goto Label5
							default:
								panic("unexpected function state")
							}
							// line 1095: args = tuple(self.stack[k+1:])
							πF.SetLineno(1095)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Add(πF, µk, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
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
							µargs = πTemp003
							// line 1096: del self.stack[k:]
							πF.SetLineno(1096)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µk, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							// line 1097: instantiated = 0
							πF.SetLineno(1097)
							µinstantiated = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, µargs); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							πTemp001[0] = µklass
							if πTemp004, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßClassType); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(πTemp007 == πTemp004).ToObject()
							πTemp002 = πTemp003
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label1
							}
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							πTemp001[0] = µklass
							πTemp001[1] = ß__getinitargs__.ToObject()
							if πTemp004, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(!πTemp006).ToObject()
							πTemp002 = πTemp003
						Label1:
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label2
							}
							goto Label3
							// line 1098: if (not args and
							πF.SetLineno(1098)
						Label2:
							// line 1101: try:
							πF.SetLineno(1101)
							πF.PushCheckpoint(5)
							// line 1102: value = _EmptyClass()
							πF.SetLineno(1102)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_EmptyClass); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp003
							// line 1103: value.__class__ = klass
							πF.SetLineno(1103)
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µklass); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µvalue, ß__class__, πTemp002); πE != nil {
								continue
							}
							// line 1104: instantiated = 1
							πF.SetLineno(1104)
							µinstantiated = πg.NewInt(1).ToObject()
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label6
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 1105: except RuntimeError:
							πF.SetLineno(1105)
						Label6:
							// line 1108: pass
							πF.SetLineno(1108)
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							goto Label3
						Label3:
							if πE = πg.CheckLocal(πF, µinstantiated, "instantiated"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µinstantiated); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label7
							}
							goto Label8
							// line 1109: if not instantiated:
							πF.SetLineno(1109)
						Label7:
							// line 1110: try:
							πF.SetLineno(1110)
							πF.PushCheckpoint(10)
							// line 1111: value = klass(*args)
							πF.SetLineno(1111)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, µklass, nil, µargs, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp002
							πF.PopCheckpoint()
							goto Label9
						Label10:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp008, πTemp009 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label11
							}
							πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
							continue
							// line 1112: except TypeError, err:
							πF.SetLineno(1112)
						Label11:
							µerr = πTemp008.ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µklass, ß__name__, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µerr, "err"); πE != nil {
								continue
							}
							πTemp001[0] = µerr
							if πTemp010, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp004 = πg.NewTuple2(πTemp007, πTemp011).ToObject()
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("in constructor for %s: %s").ToObject(), πTemp004); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(2).ToObject()
							if πTemp010, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßexc_info, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πTemp011.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp010, πTemp004); πE != nil {
								continue
							}
							// line 1113: raise TypeError, "in constructor for %s: %s" % (
							πF.SetLineno(1113)
							πE = πF.Raise(πTemp002, πTemp003, πTemp007)
							continue
							πF.RestoreExc(nil, nil)
							goto Label9
						Label9:
							goto Label8
						Label8:
							// line 1115: self.append(value)
							πF.SetLineno(1115)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp001[0] = µvalue
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_instantiate.ToObject(), πTemp036); πE != nil {
						continue
					}
					// line 1117: def load_inst(self):
					πF.SetLineno(1117)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp037 = πg.NewFunction(πg.NewCode("load_inst", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmodule *πg.Object = πg.UnboundLocal
						_ = µmodule
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µklass *πg.Object = πg.UnboundLocal
						_ = µklass
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
							// line 1118: module = self.readline()[:-1]
							πF.SetLineno(1118)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µmodule = πTemp002
							// line 1119: name = self.readline()[:-1]
							πF.SetLineno(1119)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µname = πTemp002
							// line 1120: klass = self.find_class(module, name)
							πF.SetLineno(1120)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp005[0] = µmodule
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005[1] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfind_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µklass = πTemp002
							// line 1121: self._instantiate(klass, self.marker())
							πF.SetLineno(1121)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							πTemp005[0] = µklass
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmarker, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005[1] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_instantiate, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_inst.ToObject(), πTemp037); πE != nil {
						continue
					}
					// line 1122: dispatch[INST] = load_inst
					πF.SetLineno(1122)
					if πTemp038, πE = πg.ResolveClass(πF, πClass, nil, ßload_inst); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp039}, πTemp038); πE != nil {
						continue
					}
					if πTemp040, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp042, πE = πg.ResolveClass(πF, πClass, nil, ßINST); πE != nil {
						continue
					}
					πTemp041 = πTemp042
					if πE = πg.SetItem(πF, πTemp040, πTemp041, πTemp039); πE != nil {
						continue
					}
					// line 1124: def load_obj(self):
					πF.SetLineno(1124)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp038 = πg.NewFunction(πg.NewCode("load_obj", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µk *πg.Object = πg.UnboundLocal
						_ = µk
						var µklass *πg.Object = πg.UnboundLocal
						_ = µklass
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
							// line 1126: k = self.marker()
							πF.SetLineno(1126)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmarker, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µk = πTemp002
							// line 1127: klass = self.stack.pop(k+1)
							πF.SetLineno(1127)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µk, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µklass = πTemp001
							// line 1128: self._instantiate(klass, k)
							πF.SetLineno(1128)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							πTemp003[0] = µklass
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[1] = µk
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_instantiate, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_obj.ToObject(), πTemp038); πE != nil {
						continue
					}
					// line 1129: dispatch[OBJ] = load_obj
					πF.SetLineno(1129)
					if πTemp039, πE = πg.ResolveClass(πF, πClass, nil, ßload_obj); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp040}, πTemp039); πE != nil {
						continue
					}
					if πTemp041, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp043, πE = πg.ResolveClass(πF, πClass, nil, ßOBJ); πE != nil {
						continue
					}
					πTemp042 = πTemp043
					if πE = πg.SetItem(πF, πTemp041, πTemp042, πTemp040); πE != nil {
						continue
					}
					// line 1131: def load_newobj(self):
					πF.SetLineno(1131)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp039 = πg.NewFunction(πg.NewCode("load_newobj", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
						var µcls *πg.Object = πg.UnboundLocal
						_ = µcls
						var µobj *πg.Object = πg.UnboundLocal
						_ = µobj
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
							// line 1132: args = self.stack.pop()
							πF.SetLineno(1132)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µargs = πTemp001
							// line 1133: cls = self.stack[-1]
							πF.SetLineno(1133)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
								continue
							}
							µcls = πTemp002
							// line 1134: obj = cls.__new__(cls, *args)
							πF.SetLineno(1134)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp004[0] = µcls
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µcls, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp001, πTemp004, µargs, nil, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µobj = πTemp002
							// line 1135: self.stack[-1] = obj
							πF.SetLineno(1135)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µobj); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp005
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
					if πE = πClass.SetItem(πF, ßload_newobj.ToObject(), πTemp039); πE != nil {
						continue
					}
					// line 1136: dispatch[NEWOBJ] = load_newobj
					πF.SetLineno(1136)
					if πTemp040, πE = πg.ResolveClass(πF, πClass, nil, ßload_newobj); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp041}, πTemp040); πE != nil {
						continue
					}
					if πTemp042, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp044, πE = πg.ResolveClass(πF, πClass, nil, ßNEWOBJ); πE != nil {
						continue
					}
					πTemp043 = πTemp044
					if πE = πg.SetItem(πF, πTemp042, πTemp043, πTemp041); πE != nil {
						continue
					}
					// line 1138: def load_global(self):
					πF.SetLineno(1138)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp040 = πg.NewFunction(πg.NewCode("load_global", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmodule *πg.Object = πg.UnboundLocal
						_ = µmodule
						var µname *πg.Object = πg.UnboundLocal
						_ = µname
						var µklass *πg.Object = πg.UnboundLocal
						_ = µklass
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
							// line 1139: module = self.readline()[:-1]
							πF.SetLineno(1139)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µmodule = πTemp002
							// line 1140: name = self.readline()[:-1]
							πF.SetLineno(1140)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp002, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µname = πTemp002
							// line 1141: klass = self.find_class(module, name)
							πF.SetLineno(1141)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp005[0] = µmodule
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005[1] = µname
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfind_class, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µklass = πTemp002
							// line 1142: self.append(klass)
							πF.SetLineno(1142)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							πTemp005[0] = µklass
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_global.ToObject(), πTemp040); πE != nil {
						continue
					}
					// line 1143: dispatch[GLOBAL] = load_global
					πF.SetLineno(1143)
					if πTemp041, πE = πg.ResolveClass(πF, πClass, nil, ßload_global); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp042}, πTemp041); πE != nil {
						continue
					}
					if πTemp043, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp045, πE = πg.ResolveClass(πF, πClass, nil, ßGLOBAL); πE != nil {
						continue
					}
					πTemp044 = πTemp045
					if πE = πg.SetItem(πF, πTemp043, πTemp044, πTemp042); πE != nil {
						continue
					}
					// line 1145: def load_ext1(self):
					πF.SetLineno(1145)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp041 = πg.NewFunction(πg.NewCode("load_ext1", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcode *πg.Object = πg.UnboundLocal
						_ = µcode
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
							// line 1146: code = ord(self.read(1))
							πF.SetLineno(1146)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
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
							µcode = πTemp004
							// line 1147: self.get_extension(code)
							πF.SetLineno(1147)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[0] = µcode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßget_extension, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_ext1.ToObject(), πTemp041); πE != nil {
						continue
					}
					// line 1148: dispatch[EXT1] = load_ext1
					πF.SetLineno(1148)
					if πTemp042, πE = πg.ResolveClass(πF, πClass, nil, ßload_ext1); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp043}, πTemp042); πE != nil {
						continue
					}
					if πTemp044, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp046, πE = πg.ResolveClass(πF, πClass, nil, ßEXT1); πE != nil {
						continue
					}
					πTemp045 = πTemp046
					if πE = πg.SetItem(πF, πTemp044, πTemp045, πTemp043); πE != nil {
						continue
					}
					// line 1150: def load_ext2(self):
					πF.SetLineno(1150)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp042 = πg.NewFunction(πg.NewCode("load_ext2", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcode *πg.Object = πg.UnboundLocal
						_ = µcode
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
						var πTemp006 *πg.Object
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
							// line 1151: code = mloads('i' + self.read(2) + '\000\000')
							πF.SetLineno(1151)
							πTemp001 = πF.MakeArgs(1)
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewInt(2).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.Add(πF, ßi.ToObject(), πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp003, πg.NewStr("\x00\x00").ToObject()); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmloads); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcode = πTemp003
							// line 1152: self.get_extension(code)
							πF.SetLineno(1152)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[0] = µcode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_extension, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_ext2.ToObject(), πTemp042); πE != nil {
						continue
					}
					// line 1153: dispatch[EXT2] = load_ext2
					πF.SetLineno(1153)
					if πTemp043, πE = πg.ResolveClass(πF, πClass, nil, ßload_ext2); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp044}, πTemp043); πE != nil {
						continue
					}
					if πTemp045, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp047, πE = πg.ResolveClass(πF, πClass, nil, ßEXT2); πE != nil {
						continue
					}
					πTemp046 = πTemp047
					if πE = πg.SetItem(πF, πTemp045, πTemp046, πTemp044); πE != nil {
						continue
					}
					// line 1155: def load_ext4(self):
					πF.SetLineno(1155)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp043 = πg.NewFunction(πg.NewCode("load_ext4", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcode *πg.Object = πg.UnboundLocal
						_ = µcode
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
							// line 1156: code = mloads('i' + self.read(4))
							πF.SetLineno(1156)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, ßi.ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmloads); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcode = πTemp004
							// line 1157: self.get_extension(code)
							πF.SetLineno(1157)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[0] = µcode
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßget_extension, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_ext4.ToObject(), πTemp043); πE != nil {
						continue
					}
					// line 1158: dispatch[EXT4] = load_ext4
					πF.SetLineno(1158)
					if πTemp044, πE = πg.ResolveClass(πF, πClass, nil, ßload_ext4); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp045}, πTemp044); πE != nil {
						continue
					}
					if πTemp046, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp048, πE = πg.ResolveClass(πF, πClass, nil, ßEXT4); πE != nil {
						continue
					}
					πTemp047 = πTemp048
					if πE = πg.SetItem(πF, πTemp046, πTemp047, πTemp045); πE != nil {
						continue
					}
					// line 1160: def get_extension(self, code):
					πF.SetLineno(1160)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "code", Def: nil}
					πTemp044 = πg.NewFunction(πg.NewCode("get_extension", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µcode *πg.Object = πArgs[1]
						_ = µcode
						var µnil *πg.Object = πg.UnboundLocal
						_ = µnil
						var µobj *πg.Object = πg.UnboundLocal
						_ = µobj
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
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
							// line 1161: nil = []
							πF.SetLineno(1161)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µnil = πTemp002
							// line 1162: obj = _extension_cache.get(code, nil)
							πF.SetLineno(1162)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[0] = µcode
							if πE = πg.CheckLocal(πF, µnil, "nil"); πE != nil {
								continue
							}
							πTemp001[1] = µnil
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_extension_cache); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µobj = πTemp002
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnil, "nil"); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(µobj != µnil).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1163: if obj is not nil:
							πF.SetLineno(1163)
						Label1:
							// line 1164: self.append(obj)
							πF.SetLineno(1164)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1165: return
							πF.SetLineno(1165)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 1166: key = _inverted_registry.get(code)
							πF.SetLineno(1166)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp001[0] = µcode
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_inverted_registry); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µkey = πTemp002
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µkey); πE != nil {
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
							// line 1167: if not key:
							πF.SetLineno(1167)
						Label3:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("unregistered extension code %d").ToObject(), µcode); πE != nil {
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
							// line 1168: raise ValueError("unregistered extension code %d" % code)
							πF.SetLineno(1168)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label4
						Label4:
							// line 1169: obj = self.find_class(*key)
							πF.SetLineno(1169)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfind_class, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp002, nil, µkey, nil, nil); πE != nil {
								continue
							}
							µobj = πTemp003
							// line 1170: _extension_cache[code] = obj
							πF.SetLineno(1170)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µobj); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_extension_cache); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcode, "code"); πE != nil {
								continue
							}
							πTemp005 = µcode
							if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
								continue
							}
							// line 1171: self.append(obj)
							πF.SetLineno(1171)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							πTemp001[0] = µobj
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßget_extension.ToObject(), πTemp044); πE != nil {
						continue
					}
					// line 1173: def find_class(self, module, name):
					πF.SetLineno(1173)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "module", Def: nil}
					πTemp002[2] = πg.Param{Name: "name", Def: nil}
					πTemp045 = πg.NewFunction(πg.NewCode("find_class", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µmodule *πg.Object = πArgs[1]
						_ = µmodule
						var µname *πg.Object = πArgs[2]
						_ = µname
						var µmod *πg.Object = πg.UnboundLocal
						_ = µmod
						var µklass *πg.Object = πg.UnboundLocal
						_ = µklass
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
							// line 1175: __import__(module)
							πF.SetLineno(1175)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp001[0] = µmodule
							if πTemp002, πE = πg.ResolveGlobal(πF, ß__import__); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 1176: mod = sys.modules[module]
							πF.SetLineno(1176)
							if πE = πg.CheckLocal(πF, µmodule, "module"); πE != nil {
								continue
							}
							πTemp002 = µmodule
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
							// line 1177: klass = getattr(mod, name)
							πF.SetLineno(1177)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µmod, "mod"); πE != nil {
								continue
							}
							πTemp001[0] = µmod
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp001[1] = µname
							if πTemp002, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µklass = πTemp003
							// line 1178: return klass
							πF.SetLineno(1178)
							if πE = πg.CheckLocal(πF, µklass, "klass"); πE != nil {
								continue
							}
							πR = µklass
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßfind_class.ToObject(), πTemp045); πE != nil {
						continue
					}
					// line 1180: def load_reduce(self):
					πF.SetLineno(1180)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp046 = πg.NewFunction(πg.NewCode("load_reduce", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstack *πg.Object = πg.UnboundLocal
						_ = µstack
						var µargs *πg.Object = πg.UnboundLocal
						_ = µargs
						var µfunc *πg.Object = πg.UnboundLocal
						_ = µfunc
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
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
							// line 1181: stack = self.stack
							πF.SetLineno(1181)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							µstack = πTemp001
							// line 1182: args = stack.pop()
							πF.SetLineno(1182)
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstack, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µargs = πTemp002
							// line 1183: func = stack[-1]
							πF.SetLineno(1183)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstack, πTemp001); πE != nil {
								continue
							}
							µfunc = πTemp002
							// line 1184: value = func(*args)
							πF.SetLineno(1184)
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfunc, "func"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Invoke(πF, µfunc, nil, µargs, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp001
							// line 1185: stack[-1] = value
							πF.SetLineno(1185)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.SetItem(πF, µstack, πTemp002, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_reduce.ToObject(), πTemp046); πE != nil {
						continue
					}
					// line 1186: dispatch[REDUCE] = load_reduce
					πF.SetLineno(1186)
					if πTemp047, πE = πg.ResolveClass(πF, πClass, nil, ßload_reduce); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp048}, πTemp047); πE != nil {
						continue
					}
					if πTemp049, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp051, πE = πg.ResolveClass(πF, πClass, nil, ßREDUCE); πE != nil {
						continue
					}
					πTemp050 = πTemp051
					if πE = πg.SetItem(πF, πTemp049, πTemp050, πTemp048); πE != nil {
						continue
					}
					// line 1188: def load_pop(self):
					πF.SetLineno(1188)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp047 = πg.NewFunction(πg.NewCode("load_pop", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 1189: del self.stack[-1]
							πF.SetLineno(1189)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_pop.ToObject(), πTemp047); πE != nil {
						continue
					}
					// line 1190: dispatch[POP] = load_pop
					πF.SetLineno(1190)
					if πTemp048, πE = πg.ResolveClass(πF, πClass, nil, ßload_pop); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp049}, πTemp048); πE != nil {
						continue
					}
					if πTemp050, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp052, πE = πg.ResolveClass(πF, πClass, nil, ßPOP); πE != nil {
						continue
					}
					πTemp051 = πTemp052
					if πE = πg.SetItem(πF, πTemp050, πTemp051, πTemp049); πE != nil {
						continue
					}
					// line 1192: def load_pop_mark(self):
					πF.SetLineno(1192)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp048 = πg.NewFunction(πg.NewCode("load_pop_mark", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µk *πg.Object = πg.UnboundLocal
						_ = µk
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
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
							// line 1193: k = self.marker()
							πF.SetLineno(1193)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmarker, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µk = πTemp002
							// line 1194: del self.stack[k:]
							πF.SetLineno(1194)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µk, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_pop_mark.ToObject(), πTemp048); πE != nil {
						continue
					}
					// line 1195: dispatch[POP_MARK] = load_pop_mark
					πF.SetLineno(1195)
					if πTemp049, πE = πg.ResolveClass(πF, πClass, nil, ßload_pop_mark); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp050}, πTemp049); πE != nil {
						continue
					}
					if πTemp051, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp053, πE = πg.ResolveClass(πF, πClass, nil, ßPOP_MARK); πE != nil {
						continue
					}
					πTemp052 = πTemp053
					if πE = πg.SetItem(πF, πTemp051, πTemp052, πTemp050); πE != nil {
						continue
					}
					// line 1197: def load_dup(self):
					πF.SetLineno(1197)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp049 = πg.NewFunction(πg.NewCode("load_dup", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 1198: self.append(self.stack[-1])
							πF.SetLineno(1198)
							πTemp001 = πF.MakeArgs(1)
							if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_dup.ToObject(), πTemp049); πE != nil {
						continue
					}
					// line 1199: dispatch[DUP] = load_dup
					πF.SetLineno(1199)
					if πTemp050, πE = πg.ResolveClass(πF, πClass, nil, ßload_dup); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp051}, πTemp050); πE != nil {
						continue
					}
					if πTemp052, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp054, πE = πg.ResolveClass(πF, πClass, nil, ßDUP); πE != nil {
						continue
					}
					πTemp053 = πTemp054
					if πE = πg.SetItem(πF, πTemp052, πTemp053, πTemp051); πE != nil {
						continue
					}
					// line 1201: def load_get(self):
					πF.SetLineno(1201)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp050 = πg.NewFunction(πg.NewCode("load_get", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
							// line 1202: self.append(self.memo[self.readline()[:-1]])
							πF.SetLineno(1202)
							πTemp001 = πF.MakeArgs(1)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp004, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp006, πTemp003); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_get.ToObject(), πTemp050); πE != nil {
						continue
					}
					// line 1203: dispatch[GET] = load_get
					πF.SetLineno(1203)
					if πTemp051, πE = πg.ResolveClass(πF, πClass, nil, ßload_get); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp052}, πTemp051); πE != nil {
						continue
					}
					if πTemp053, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp055, πE = πg.ResolveClass(πF, πClass, nil, ßGET); πE != nil {
						continue
					}
					πTemp054 = πTemp055
					if πE = πg.SetItem(πF, πTemp053, πTemp054, πTemp052); πE != nil {
						continue
					}
					// line 1205: def load_binget(self):
					πF.SetLineno(1205)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp051 = πg.NewFunction(πg.NewCode("load_binget", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
							// line 1206: i = ord(self.read(1))
							πF.SetLineno(1206)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
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
							µi = πTemp004
							// line 1207: self.append(self.memo[repr(i)])
							πF.SetLineno(1207)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp002[0] = µi
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp003 = πTemp005
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_binget.ToObject(), πTemp051); πE != nil {
						continue
					}
					// line 1208: dispatch[BINGET] = load_binget
					πF.SetLineno(1208)
					if πTemp052, πE = πg.ResolveClass(πF, πClass, nil, ßload_binget); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp053}, πTemp052); πE != nil {
						continue
					}
					if πTemp054, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp056, πE = πg.ResolveClass(πF, πClass, nil, ßBINGET); πE != nil {
						continue
					}
					πTemp055 = πTemp056
					if πE = πg.SetItem(πF, πTemp054, πTemp055, πTemp053); πE != nil {
						continue
					}
					// line 1210: def load_long_binget(self):
					πF.SetLineno(1210)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp052 = πg.NewFunction(πg.NewCode("load_long_binget", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
							// line 1211: i = mloads('i' + self.read(4))
							πF.SetLineno(1211)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, ßi.ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmloads); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp004
							// line 1212: self.append(self.memo[repr(i)])
							πF.SetLineno(1212)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp003[0] = µi
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
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
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_long_binget.ToObject(), πTemp052); πE != nil {
						continue
					}
					// line 1213: dispatch[LONG_BINGET] = load_long_binget
					πF.SetLineno(1213)
					if πTemp053, πE = πg.ResolveClass(πF, πClass, nil, ßload_long_binget); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp054}, πTemp053); πE != nil {
						continue
					}
					if πTemp055, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp057, πE = πg.ResolveClass(πF, πClass, nil, ßLONG_BINGET); πE != nil {
						continue
					}
					πTemp056 = πTemp057
					if πE = πg.SetItem(πF, πTemp055, πTemp056, πTemp054); πE != nil {
						continue
					}
					// line 1215: def load_put(self):
					πF.SetLineno(1215)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp053 = πg.NewFunction(πg.NewCode("load_put", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
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
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 *πg.Object
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
							// line 1216: self.memo[self.readline()[:-1]] = self.stack[-1]
							πF.SetLineno(1216)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp006, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßreadline, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
								continue
							}
							πTemp004 = πTemp006
							if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_put.ToObject(), πTemp053); πE != nil {
						continue
					}
					// line 1217: dispatch[PUT] = load_put
					πF.SetLineno(1217)
					if πTemp054, πE = πg.ResolveClass(πF, πClass, nil, ßload_put); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp055}, πTemp054); πE != nil {
						continue
					}
					if πTemp056, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp058, πE = πg.ResolveClass(πF, πClass, nil, ßPUT); πE != nil {
						continue
					}
					πTemp057 = πTemp058
					if πE = πg.SetItem(πF, πTemp056, πTemp057, πTemp055); πE != nil {
						continue
					}
					// line 1219: def load_binput(self):
					πF.SetLineno(1219)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp054 = πg.NewFunction(πg.NewCode("load_binput", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
						var πTemp008 *πg.Object
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
							// line 1220: i = ord(self.read(1))
							πF.SetLineno(1220)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
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
							µi = πTemp004
							// line 1221: self.memo[repr(i)] = self.stack[-1]
							πF.SetLineno(1221)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp006 = πTemp008
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_binput.ToObject(), πTemp054); πE != nil {
						continue
					}
					// line 1222: dispatch[BINPUT] = load_binput
					πF.SetLineno(1222)
					if πTemp055, πE = πg.ResolveClass(πF, πClass, nil, ßload_binput); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp056}, πTemp055); πE != nil {
						continue
					}
					if πTemp057, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp059, πE = πg.ResolveClass(πF, πClass, nil, ßBINPUT); πE != nil {
						continue
					}
					πTemp058 = πTemp059
					if πE = πg.SetItem(πF, πTemp057, πTemp058, πTemp056); πE != nil {
						continue
					}
					// line 1224: def load_long_binput(self):
					πF.SetLineno(1224)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp055 = πg.NewFunction(πg.NewCode("load_long_binput", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
							// line 1225: i = mloads('i' + self.read(4))
							πF.SetLineno(1225)
							πTemp001 = πF.MakeArgs(1)
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewInt(4).ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßread, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, ßi.ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmloads); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µi = πTemp004
							// line 1226: self.memo[repr(i)] = self.stack[-1]
							πF.SetLineno(1226)
							if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp004
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ßmemo, nil); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp001[0] = µi
							if πTemp007, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp006 = πTemp008
							if πE = πg.SetItem(πF, πTemp005, πTemp006, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_long_binput.ToObject(), πTemp055); πE != nil {
						continue
					}
					// line 1227: dispatch[LONG_BINPUT] = load_long_binput
					πF.SetLineno(1227)
					if πTemp056, πE = πg.ResolveClass(πF, πClass, nil, ßload_long_binput); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp057}, πTemp056); πE != nil {
						continue
					}
					if πTemp058, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp060, πE = πg.ResolveClass(πF, πClass, nil, ßLONG_BINPUT); πE != nil {
						continue
					}
					πTemp059 = πTemp060
					if πE = πg.SetItem(πF, πTemp058, πTemp059, πTemp057); πE != nil {
						continue
					}
					// line 1229: def load_append(self):
					πF.SetLineno(1229)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp056 = πg.NewFunction(πg.NewCode("load_append", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstack *πg.Object = πg.UnboundLocal
						_ = µstack
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µlist *πg.Object = πg.UnboundLocal
						_ = µlist
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
							// line 1230: stack = self.stack
							πF.SetLineno(1230)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							µstack = πTemp001
							// line 1231: value = stack.pop()
							πF.SetLineno(1231)
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstack, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp002
							// line 1232: list = stack[-1]
							πF.SetLineno(1232)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstack, πTemp001); πE != nil {
								continue
							}
							µlist = πTemp002
							// line 1233: list.append(value)
							πF.SetLineno(1233)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[0] = µvalue
							if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlist, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_append.ToObject(), πTemp056); πE != nil {
						continue
					}
					// line 1234: dispatch[APPEND] = load_append
					πF.SetLineno(1234)
					if πTemp057, πE = πg.ResolveClass(πF, πClass, nil, ßload_append); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp058}, πTemp057); πE != nil {
						continue
					}
					if πTemp059, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp061, πE = πg.ResolveClass(πF, πClass, nil, ßAPPEND); πE != nil {
						continue
					}
					πTemp060 = πTemp061
					if πE = πg.SetItem(πF, πTemp059, πTemp060, πTemp058); πE != nil {
						continue
					}
					// line 1236: def load_appends(self):
					πF.SetLineno(1236)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp057 = πg.NewFunction(πg.NewCode("load_appends", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstack *πg.Object = πg.UnboundLocal
						_ = µstack
						var µmark *πg.Object = πg.UnboundLocal
						_ = µmark
						var µlist *πg.Object = πg.UnboundLocal
						_ = µlist
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
							// line 1237: stack = self.stack
							πF.SetLineno(1237)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							µstack = πTemp001
							// line 1238: mark = self.marker()
							πF.SetLineno(1238)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmarker, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µmark = πTemp002
							// line 1239: list = stack[mark - 1]
							πF.SetLineno(1239)
							if πE = πg.CheckLocal(πF, µmark, "mark"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µmark, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstack, πTemp001); πE != nil {
								continue
							}
							µlist = πTemp002
							// line 1240: list.extend(stack[mark + 1:])
							πF.SetLineno(1240)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µmark, "mark"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µmark, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstack, πTemp001); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							if πE = πg.CheckLocal(πF, µlist, "list"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µlist, ßextend, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1241: del stack[mark:]
							πF.SetLineno(1241)
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmark, "mark"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µmark, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µstack, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_appends.ToObject(), πTemp057); πE != nil {
						continue
					}
					// line 1242: dispatch[APPENDS] = load_appends
					πF.SetLineno(1242)
					if πTemp058, πE = πg.ResolveClass(πF, πClass, nil, ßload_appends); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp059}, πTemp058); πE != nil {
						continue
					}
					if πTemp060, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp062, πE = πg.ResolveClass(πF, πClass, nil, ßAPPENDS); πE != nil {
						continue
					}
					πTemp061 = πTemp062
					if πE = πg.SetItem(πF, πTemp060, πTemp061, πTemp059); πE != nil {
						continue
					}
					// line 1244: def load_setitem(self):
					πF.SetLineno(1244)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp058 = πg.NewFunction(πg.NewCode("load_setitem", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstack *πg.Object = πg.UnboundLocal
						_ = µstack
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
						var µkey *πg.Object = πg.UnboundLocal
						_ = µkey
						var µdict *πg.Object = πg.UnboundLocal
						_ = µdict
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
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
							// line 1245: stack = self.stack
							πF.SetLineno(1245)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							µstack = πTemp001
							// line 1246: value = stack.pop()
							πF.SetLineno(1246)
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstack, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp002
							// line 1247: key = stack.pop()
							πF.SetLineno(1247)
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstack, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µkey = πTemp002
							// line 1248: dict = stack[-1]
							πF.SetLineno(1248)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstack, πTemp001); πE != nil {
								continue
							}
							µdict = πTemp002
							// line 1249: dict[key] = value
							πF.SetLineno(1249)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp002 = µkey
							if πE = πg.SetItem(πF, µdict, πTemp002, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_setitem.ToObject(), πTemp058); πE != nil {
						continue
					}
					// line 1250: dispatch[SETITEM] = load_setitem
					πF.SetLineno(1250)
					if πTemp059, πE = πg.ResolveClass(πF, πClass, nil, ßload_setitem); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp060}, πTemp059); πE != nil {
						continue
					}
					if πTemp061, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp063, πE = πg.ResolveClass(πF, πClass, nil, ßSETITEM); πE != nil {
						continue
					}
					πTemp062 = πTemp063
					if πE = πg.SetItem(πF, πTemp061, πTemp062, πTemp060); πE != nil {
						continue
					}
					// line 1252: def load_setitems(self):
					πF.SetLineno(1252)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp059 = πg.NewFunction(πg.NewCode("load_setitems", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstack *πg.Object = πg.UnboundLocal
						_ = µstack
						var µmark *πg.Object = πg.UnboundLocal
						_ = µmark
						var µdict *πg.Object = πg.UnboundLocal
						_ = µdict
						var µi *πg.Object = πg.UnboundLocal
						_ = µi
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
							// line 1253: stack = self.stack
							πF.SetLineno(1253)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							µstack = πTemp001
							// line 1254: mark = self.marker()
							πF.SetLineno(1254)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßmarker, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µmark = πTemp002
							// line 1255: dict = stack[mark - 1]
							πF.SetLineno(1255)
							if πE = πg.CheckLocal(πF, µmark, "mark"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Sub(πF, µmark, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstack, πTemp001); πE != nil {
								continue
							}
							µdict = πTemp002
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µmark, "mark"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, µmark, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp003[0] = πTemp002
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							πTemp004[0] = µstack
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp003[1] = πTemp005
							πTemp003[2] = πg.NewInt(2).ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
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
								µi = πTemp002
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(1)
							// line 1257: dict[stack[i]] = stack[i + 1]
							πF.SetLineno(1257)
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp005
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetItem(πF, µstack, πTemp002); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp005); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.GetItem(πF, µstack, πTemp009); πE != nil {
								continue
							}
							πTemp008 = πTemp010
							if πE = πg.SetItem(πF, µdict, πTemp008, πTemp002); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 1259: del stack[mark:]
							πF.SetLineno(1259)
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmark, "mark"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µmark, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.DelItem(πF, µstack, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_setitems.ToObject(), πTemp059); πE != nil {
						continue
					}
					// line 1260: dispatch[SETITEMS] = load_setitems
					πF.SetLineno(1260)
					if πTemp060, πE = πg.ResolveClass(πF, πClass, nil, ßload_setitems); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp061}, πTemp060); πE != nil {
						continue
					}
					if πTemp062, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp064, πE = πg.ResolveClass(πF, πClass, nil, ßSETITEMS); πE != nil {
						continue
					}
					πTemp063 = πTemp064
					if πE = πg.SetItem(πF, πTemp062, πTemp063, πTemp061); πE != nil {
						continue
					}
					// line 1262: def load_build(self):
					πF.SetLineno(1262)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp060 = πg.NewFunction(πg.NewCode("load_build", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µstack *πg.Object = πg.UnboundLocal
						_ = µstack
						var µstate *πg.Object = πg.UnboundLocal
						_ = µstate
						var µinst *πg.Object = πg.UnboundLocal
						_ = µinst
						var µsetstate *πg.Object = πg.UnboundLocal
						_ = µsetstate
						var µslotstate *πg.Object = πg.UnboundLocal
						_ = µslotstate
						var µd *πg.Object = πg.UnboundLocal
						_ = µd
						var µk *πg.Object = πg.UnboundLocal
						_ = µk
						var µv *πg.Object = πg.UnboundLocal
						_ = µv
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
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 *πg.BaseException
						_ = πTemp009
						var πTemp010 *πg.Traceback
						_ = πTemp010
						var πR *πg.Object
						_ = πR
						var πE *πg.BaseException
						_ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 9:
								goto Label9
							case 11:
								goto Label11
							case 12:
								goto Label12
							case 13:
								goto Label13
							case 17:
								goto Label17
							case 18:
								goto Label18
							case 22:
								goto Label22
							case 23:
								goto Label23
							default:
								panic("unexpected function state")
							}
							// line 1263: stack = self.stack
							πF.SetLineno(1263)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							µstack = πTemp001
							// line 1264: state = stack.pop()
							πF.SetLineno(1264)
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µstack, ßpop, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							µstate = πTemp002
							// line 1265: inst = stack[-1]
							πF.SetLineno(1265)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πE = πg.CheckLocal(πF, µstack, "stack"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µstack, πTemp001); πE != nil {
								continue
							}
							µinst = πTemp002
							// line 1266: setstate = getattr(inst, "__setstate__", None)
							πF.SetLineno(1266)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µinst, "inst"); πE != nil {
								continue
							}
							πTemp003[0] = µinst
							πTemp003[1] = ß__setstate__.ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003[2] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µsetstate = πTemp002
							if πE = πg.CheckLocal(πF, µsetstate, "setstate"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µsetstate); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 1267: if setstate:
							πF.SetLineno(1267)
						Label1:
							// line 1268: setstate(state)
							πF.SetLineno(1268)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp003[0] = µstate
							if πE = πg.CheckLocal(πF, µsetstate, "setstate"); πE != nil {
								continue
							}
							if πTemp001, πE = µsetstate.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1269: return
							πF.SetLineno(1269)
							πR = πg.None
							continue
							goto Label2
						Label2:
							// line 1270: slotstate = None
							πF.SetLineno(1270)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							µslotstate = πTemp001
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp003[0] = µstate
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							πTemp003[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001 = πTemp005
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label3
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp003[0] = µstate
							if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewInt(2).ToObject()); πE != nil {
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
							// line 1271: if isinstance(state, tuple) and len(state) == 2:
							πF.SetLineno(1271)
						Label4:
							// line 1272: state, slotstate = state
							πF.SetLineno(1272)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µstate); πE != nil {
								continue
							}
							µstate = πTemp001
							µslotstate = πTemp002
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µstate); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							goto Label7
							// line 1273: if state:
							πF.SetLineno(1273)
						Label6:
							// line 1274: try:
							πF.SetLineno(1274)
							πF.PushCheckpoint(9)
							// line 1275: d = inst.__dict__
							πF.SetLineno(1275)
							if πE = πg.CheckLocal(πF, µinst, "inst"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µinst, ß__dict__, nil); πE != nil {
								continue
							}
							µd = πTemp001
							// line 1276: try:
							πF.SetLineno(1276)
							πF.PushCheckpoint(11)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßiteritems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µk = πTemp005
								µv = πTemp006
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(12)
							// line 1278: d[intern(k)] = v
							πF.SetLineno(1278)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µv); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[0] = µk
							if πTemp006, πE = πg.ResolveGlobal(πF, ßintern); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp005 = πTemp008
							if πE = πg.SetItem(πF, µd, πTemp005, πTemp002); πE != nil {
								continue
							}
							continue
						Label13:
							if πE != nil || πR != nil {
								continue
							}
						Label14:
							πF.PopCheckpoint()
							goto Label10
						Label11:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label15
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 1281: except TypeError:
							πF.SetLineno(1281)
						Label15:
							// line 1282: d.update(state)
							πF.SetLineno(1282)
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							πTemp003[0] = µstate
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µd, ßupdate, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πF.RestoreExc(nil, nil)
							goto Label10
						Label10:
							πF.PopCheckpoint()
							goto Label8
						Label9:
							if πE == nil {
								continue
							}
							πE = nil
							πTemp009, πTemp010 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label16
							}
							πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
							continue
							// line 1284: except RuntimeError:
							πF.SetLineno(1284)
						Label16:
							if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µstate, ßitems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µk = πTemp005
								µv = πTemp006
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(17)
							// line 1295: setattr(inst, k, v)
							πF.SetLineno(1295)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µinst, "inst"); πE != nil {
								continue
							}
							πTemp003[0] = µinst
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[1] = µk
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp003[2] = µv
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label18:
							if πE != nil || πR != nil {
								continue
							}
						Label19:
							πF.RestoreExc(nil, nil)
							goto Label8
						Label8:
							goto Label7
						Label7:
							if πE = πg.CheckLocal(πF, µslotstate, "slotstate"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µslotstate); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label20
							}
							goto Label21
							// line 1296: if slotstate:
							πF.SetLineno(1296)
						Label20:
							if πE = πg.CheckLocal(πF, µslotstate, "slotstate"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µslotstate, ßitems, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp005); πE != nil {
								continue
							}
							πF.PushCheckpoint(23)
							πTemp004 = false
						Label22:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp004 {
								πF.PopCheckpoint()
								goto Label24
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp002); πE != nil {
									continue
								}
								µk = πTemp005
								µv = πTemp006
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(22)
							// line 1298: setattr(inst, k, v)
							πF.SetLineno(1298)
							πTemp003 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µinst, "inst"); πE != nil {
								continue
							}
							πTemp003[0] = µinst
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp003[1] = µk
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							πTemp003[2] = µv
							if πTemp002, πE = πg.ResolveGlobal(πF, ßsetattr); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							continue
						Label23:
							if πE != nil || πR != nil {
								continue
							}
						Label24:
							goto Label21
						Label21:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßload_build.ToObject(), πTemp060); πE != nil {
						continue
					}
					// line 1299: dispatch[BUILD] = load_build
					πF.SetLineno(1299)
					if πTemp061, πE = πg.ResolveClass(πF, πClass, nil, ßload_build); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp062}, πTemp061); πE != nil {
						continue
					}
					if πTemp063, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp065, πE = πg.ResolveClass(πF, πClass, nil, ßBUILD); πE != nil {
						continue
					}
					πTemp064 = πTemp065
					if πE = πg.SetItem(πF, πTemp063, πTemp064, πTemp062); πE != nil {
						continue
					}
					// line 1301: def load_mark(self):
					πF.SetLineno(1301)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp061 = πg.NewFunction(πg.NewCode("load_mark", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 1302: self.append(self.mark)
							πF.SetLineno(1302)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßmark, nil); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßappend, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßload_mark.ToObject(), πTemp061); πE != nil {
						continue
					}
					// line 1303: dispatch[MARK] = load_mark
					πF.SetLineno(1303)
					if πTemp062, πE = πg.ResolveClass(πF, πClass, nil, ßload_mark); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp063}, πTemp062); πE != nil {
						continue
					}
					if πTemp064, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp066, πE = πg.ResolveClass(πF, πClass, nil, ßMARK); πE != nil {
						continue
					}
					πTemp065 = πTemp066
					if πE = πg.SetItem(πF, πTemp064, πTemp065, πTemp063); πE != nil {
						continue
					}
					// line 1305: def load_stop(self):
					πF.SetLineno(1305)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp062 = πg.NewFunction(πg.NewCode("load_stop", "/usr/lib/python2.7/pickle.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]
						_ = µself
						var µvalue *πg.Object = πg.UnboundLocal
						_ = µvalue
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
							// line 1306: value = self.stack.pop()
							πF.SetLineno(1306)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßstack, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßpop, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µvalue = πTemp001
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp003[0] = µvalue
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_Stop); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 1307: raise _Stop(value)
							πF.SetLineno(1307)
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
					if πE = πClass.SetItem(πF, ßload_stop.ToObject(), πTemp062); πE != nil {
						continue
					}
					// line 1308: dispatch[STOP] = load_stop
					πF.SetLineno(1308)
					if πTemp063, πE = πg.ResolveClass(πF, πClass, nil, ßload_stop); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp064}, πTemp063); πE != nil {
						continue
					}
					if πTemp065, πE = πg.ResolveClass(πF, πClass, nil, ßdispatch); πE != nil {
						continue
					}
					if πTemp067, πE = πg.ResolveClass(πF, πClass, nil, ßSTOP); πE != nil {
						continue
					}
					πTemp066 = πTemp067
					if πE = πg.SetItem(πF, πTemp065, πTemp066, πTemp064); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp011, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp011 == nil {
				πTemp011 = πg.TypeType.ToObject()
			}
			if πTemp012, πE = πTemp011.Call(πF, []*πg.Object{πg.NewStr("Unpickler").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßUnpickler.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 1312: class _EmptyClass:
			πF.SetLineno(1312)
			πTemp002 = make([]*πg.Object, 0)
			πTemp004 = πg.NewDict()
			if πTemp010, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp010); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_EmptyClass", "/usr/lib/python2.7/pickle.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default:
						panic("unexpected function state")
					}
					// line 1313: pass
					πF.SetLineno(1313)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp011, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp011 == nil {
				πTemp011 = πg.TypeType.ToObject()
			}
			if πTemp012, πE = πTemp011.Call(πF, []*πg.Object{πg.NewStr("_EmptyClass").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_EmptyClass.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 1317: import binascii as _binascii
			πF.SetLineno(1317)
			if πTemp002, πE = πg.ImportModule(πF, "binascii"); πE != nil {
				continue
			}
			πTemp010 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_binascii.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 1319: def encode_long(x):
			πF.SetLineno(1319)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "x", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("encode_long", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]
				_ = µx
				var µashex *πg.Object = πg.UnboundLocal
				_ = µashex
				var µnjunkchars *πg.Object = πg.UnboundLocal
				_ = µnjunkchars
				var µnibbles *πg.Object = πg.UnboundLocal
				_ = µnibbles
				var µnbits *πg.Object = πg.UnboundLocal
				_ = µnbits
				var µnewnibbles *πg.Object = πg.UnboundLocal
				_ = µnewnibbles
				var µbinary *πg.Object = πg.UnboundLocal
				_ = µbinary
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
				var πTemp006 *πg.Object
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
					default:
						panic("unexpected function state")
					}
					// line 1320: r"""Encode a long to a two's complement little-endian binary string.
					πF.SetLineno(1320)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µx, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 1341: if x == 0:
					πF.SetLineno(1341)
				Label1:
					// line 1342: return ''
					πF.SetLineno(1342)
					πR = ß.ToObject()
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µx, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 1343: if x > 0:
					πF.SetLineno(1343)
				Label3:
					// line 1344: ashex = hex(x)
					πF.SetLineno(1344)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhex); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µashex = πTemp004
					// line 1345: assert ashex.startswith("0x")
					πF.SetLineno(1345)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = ß0x.ToObject()
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µashex, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
						continue
					}
					// line 1346: njunkchars = 2 + ashex.endswith('L')
					πF.SetLineno(1346)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = ßL.ToObject()
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µashex, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Add(πF, πg.NewInt(2).ToObject(), πTemp005); πE != nil {
						continue
					}
					µnjunkchars = πTemp001
					// line 1347: nibbles = len(ashex) - njunkchars
					πF.SetLineno(1347)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					πTemp003[0] = µashex
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µnjunkchars, "njunkchars"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp005, µnjunkchars); πE != nil {
						continue
					}
					µnibbles = πTemp001
					if πE = πg.CheckLocal(πF, µnibbles, "nibbles"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, µnibbles, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					πTemp003 = πF.MakeArgs(2)
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µashex, πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					πTemp003[1] = πg.NewInt(16).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.GE(πF, πTemp005, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label7
					}
					goto Label8
					// line 1348: if nibbles & 1:
					πF.SetLineno(1348)
				Label6:
					// line 1350: ashex = "0x0" + ashex[2:]
					πF.SetLineno(1350)
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µashex, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, ß0x0.ToObject(), πTemp005); πE != nil {
						continue
					}
					µashex = πTemp001
					goto Label8
					// line 1351: elif int(ashex[2], 16) >= 8:
					πF.SetLineno(1351)
				Label7:
					// line 1353: ashex = "0x00" + ashex[2:]
					πF.SetLineno(1353)
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µashex, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, ß0x00.ToObject(), πTemp005); πE != nil {
						continue
					}
					µashex = πTemp001
					goto Label8
				Label8:
					goto Label5
				Label4:
					// line 1358: ashex = hex(-x)
					πF.SetLineno(1358)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Neg(πF, µx); πE != nil {
						continue
					}
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhex); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µashex = πTemp004
					// line 1359: assert ashex.startswith("0x")
					πF.SetLineno(1359)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = ß0x.ToObject()
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µashex, ßstartswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Assert(πF, πTemp004, nil); πE != nil {
						continue
					}
					// line 1360: njunkchars = 2 + ashex.endswith('L')
					πF.SetLineno(1360)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = ßL.ToObject()
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µashex, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Add(πF, πg.NewInt(2).ToObject(), πTemp005); πE != nil {
						continue
					}
					µnjunkchars = πTemp001
					// line 1361: nibbles = len(ashex) - njunkchars
					πF.SetLineno(1361)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					πTemp003[0] = µashex
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µnjunkchars, "njunkchars"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp005, µnjunkchars); πE != nil {
						continue
					}
					µnibbles = πTemp001
					if πE = πg.CheckLocal(πF, µnibbles, "nibbles"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.And(πF, µnibbles, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label9
					}
					goto Label10
					// line 1362: if nibbles & 1:
					πF.SetLineno(1362)
				Label9:
					// line 1364: nibbles += 1
					πF.SetLineno(1364)
					if πE = πg.CheckLocal(πF, µnibbles, "nibbles"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µnibbles, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µnibbles = πTemp001
					goto Label10
				Label10:
					// line 1365: nbits = nibbles * 4
					πF.SetLineno(1365)
					if πE = πg.CheckLocal(πF, µnibbles, "nibbles"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Mul(πF, µnibbles, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					µnbits = πTemp001
					// line 1366: x += 1L << nbits
					πF.SetLineno(1366)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnbits, "nbits"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LShift(πF, πg.NewLongFromBytes([]byte{0x1}).ToObject(), µnbits); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IAdd(πF, µx, πTemp001); πE != nil {
						continue
					}
					µx = πTemp004
					// line 1367: assert x > 0
					πF.SetLineno(1367)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GT(πF, µx, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp001, nil); πE != nil {
						continue
					}
					// line 1368: ashex = hex(x)
					πF.SetLineno(1368)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp003[0] = µx
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhex); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µashex = πTemp004
					// line 1369: njunkchars = 2 + ashex.endswith('L')
					πF.SetLineno(1369)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = ßL.ToObject()
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µashex, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Add(πF, πg.NewInt(2).ToObject(), πTemp005); πE != nil {
						continue
					}
					µnjunkchars = πTemp001
					// line 1370: newnibbles = len(ashex) - njunkchars
					πF.SetLineno(1370)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					πTemp003[0] = µashex
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.CheckLocal(πF, µnjunkchars, "njunkchars"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, πTemp005, µnjunkchars); πE != nil {
						continue
					}
					µnewnibbles = πTemp001
					if πE = πg.CheckLocal(πF, µnewnibbles, "newnibbles"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnibbles, "nibbles"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µnewnibbles, µnibbles); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label11
					}
					goto Label12
					// line 1371: if newnibbles < nibbles:
					πF.SetLineno(1371)
				Label11:
					// line 1372: ashex = "0x" + "0" * (nibbles - newnibbles) + ashex[2:]
					πF.SetLineno(1372)
					if πE = πg.CheckLocal(πF, µnibbles, "nibbles"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnewnibbles, "newnibbles"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, µnibbles, µnewnibbles); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Mul(πF, ß0.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, ß0x.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µashex, πTemp005); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp004, πTemp006); πE != nil {
						continue
					}
					µashex = πTemp001
					goto Label12
				Label12:
					πTemp003 = πF.MakeArgs(2)
					πTemp004 = πg.NewInt(2).ToObject()
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µashex, πTemp004); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					πTemp003[1] = πg.NewInt(16).ToObject()
					if πTemp004, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.LT(πF, πTemp005, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label13
					}
					goto Label14
					// line 1373: if int(ashex[2], 16) < 8:
					πF.SetLineno(1373)
				Label13:
					// line 1375: ashex = "0xff" + ashex[2:]
					πF.SetLineno(1375)
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µashex, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, ß0xff.ToObject(), πTemp005); πE != nil {
						continue
					}
					µashex = πTemp001
					goto Label14
				Label14:
					goto Label5
				Label5:
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = ßL.ToObject()
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µashex, ßendswith, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label15
					}
					goto Label16
					// line 1377: if ashex.endswith('L'):
					πF.SetLineno(1377)
				Label15:
					// line 1378: ashex = ashex[2:-1]
					πF.SetLineno(1378)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πTemp004, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µashex, πTemp001); πE != nil {
						continue
					}
					µashex = πTemp004
					goto Label17
				Label16:
					// line 1380: ashex = ashex[2:]
					πF.SetLineno(1380)
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µashex, πTemp001); πE != nil {
						continue
					}
					µashex = πTemp004
					goto Label17
				Label17:
					// line 1381: assert len(ashex) & 1 == 0, (x, ashex)
					πF.SetLineno(1381)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µx, µashex).ToObject()
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					πTemp003[0] = µashex
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.And(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp005, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πE = πg.Assert(πF, πTemp004, πTemp001); πE != nil {
						continue
					}
					// line 1382: binary = _binascii.unhexlify(ashex)
					πF.SetLineno(1382)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					πTemp003[0] = µashex
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_binascii); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp001, ßunhexlify, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µbinary = πTemp001
					// line 1383: return binary[::-1]
					πF.SetLineno(1383)
					if πTemp004, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πTemp004}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbinary, "binary"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µbinary, πTemp001); πE != nil {
						continue
					}
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
			if πE = πF.Globals().SetItem(πF, ßencode_long.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 1320: r"""Encode a long to a two's complement little-endian binary string.
			πF.SetLineno(1320)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Encode a long to a two's complement little-endian binary string.\n    Note that 0L is a special case, returning an empty string, to save a\n    byte in the LONG1 pickling context.\n\n    >>> encode_long(0L)\n    ''\n    >>> encode_long(255L)\n    '\\xff\\x00'\n    >>> encode_long(32767L)\n    '\\xff\\x7f'\n    >>> encode_long(-256L)\n    '\\x00\\xff'\n    >>> encode_long(-32768L)\n    '\\x00\\x80'\n    >>> encode_long(-128L)\n    '\\x80'\n    >>> encode_long(127L)\n    '\\x7f'\n    >>>\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßencode_long); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 1385: def decode_long(data):
			πF.SetLineno(1385)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "data", Def: nil}
			πTemp011 = πg.NewFunction(πg.NewCode("decode_long", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdata *πg.Object = πArgs[0]
				_ = µdata
				var µnbytes *πg.Object = πg.UnboundLocal
				_ = µnbytes
				var µashex *πg.Object = πg.UnboundLocal
				_ = µashex
				var µn *πg.Object = πg.UnboundLocal
				_ = µn
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
					// line 1386: r"""Decode a long from a two's complement little-endian binary string.
					πF.SetLineno(1386)
					// line 1405: nbytes = len(data)
					πF.SetLineno(1405)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					πTemp001[0] = µdata
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnbytes = πTemp003
					if πE = πg.CheckLocal(πF, µnbytes, "nbytes"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, µnbytes, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 1406: if nbytes == 0:
					πF.SetLineno(1406)
				Label1:
					// line 1407: return 0L
					πF.SetLineno(1407)
					πR = πg.NewLongFromBytes([]byte{}).ToObject()
					continue
					goto Label2
				Label2:
					// line 1408: ashex = _binascii.hexlify(data[::-1])
					πF.SetLineno(1408)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.None, πTemp003}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µdata, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_binascii); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßhexlify, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µashex = πTemp002
					// line 1409: n = long(ashex, 16) # quadratic time before Python 2.3; linear now
					πF.SetLineno(1409)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µashex, "ashex"); πE != nil {
						continue
					}
					πTemp001[0] = µashex
					πTemp001[1] = πg.NewInt(16).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µn = πTemp003
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp005
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µdata, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, πTemp005, πg.NewStr("\x80").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 1410: if data[-1] >= '\x80':
					πF.SetLineno(1410)
				Label3:
					// line 1411: n -= 1L << (nbytes * 8)
					πF.SetLineno(1411)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnbytes, "nbytes"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Mul(πF, µnbytes, πg.NewInt(8).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LShift(πF, πg.NewLongFromBytes([]byte{0x1}).ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ISub(πF, µn, πTemp002); πE != nil {
						continue
					}
					µn = πTemp003
					goto Label4
				Label4:
					// line 1412: return n
					πF.SetLineno(1412)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πR = µn
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdecode_long.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 1386: r"""Decode a long from a two's complement little-endian binary string.
			πF.SetLineno(1386)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Decode a long from a two's complement little-endian binary string.\n    This is overriden on PyPy by a RPython version that has linear complexity.\n\n    >>> decode_long('')\n    0L\n    >>> decode_long(\"\\xff\\x00\")\n    255L\n    >>> decode_long(\"\\xff\\x7f\")\n    32767L\n    >>> decode_long(\"\\x00\\xff\")\n    -256L\n    >>> decode_long(\"\\x00\\x80\")\n    -32768L\n    >>> decode_long(\"\\x80\")\n    -128L\n    >>> decode_long(\"\\x7f\")\n    127L\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßdecode_long); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 1414: try:
			πF.SetLineno(1414)
			πF.PushCheckpoint(8)
			// line 1415: from __pypy__ import decode_long
			πF.SetLineno(1415)
			if πTemp002, πE = πg.ImportModule(πF, "__pypy__"); πE != nil {
				continue
			}
			πTemp012 = πTemp002[0]
			if πTemp013, πE = πg.GetAttrImport(πF, πTemp012, ßdecode_long); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßdecode_long.ToObject(), πTemp013); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label7
		Label8:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp006, πTemp007 = πF.ExcInfo()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp012); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label9
			}
			πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
			continue
			// line 1416: except ImportError:
			πF.SetLineno(1416)
		Label9:
			// line 1417: pass
			πF.SetLineno(1417)
			πF.RestoreExc(nil, nil)
			goto Label7
		Label7:
			// line 1421: try:
			πF.SetLineno(1421)
			πF.PushCheckpoint(11)
			// line 1422: from cStringIO import StringIO
			πF.SetLineno(1422)
			if πTemp002, πE = πg.ImportModule(πF, "cStringIO"); πE != nil {
				continue
			}
			πTemp012 = πTemp002[0]
			if πTemp013, πE = πg.GetAttrImport(πF, πTemp012, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp013); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			goto Label10
		Label11:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp006, πTemp007 = πF.ExcInfo()
			if πTemp012, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp012); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label12
			}
			πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
			continue
			// line 1423: except ImportError:
			πF.SetLineno(1423)
		Label12:
			// line 1424: from StringIO import StringIO
			πF.SetLineno(1424)
			if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp012 = πTemp002[0]
			if πTemp013, πE = πg.GetAttrImport(πF, πTemp012, ßStringIO); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp013); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label10
		Label10:
			// line 1426: def dump(obj, file, protocol=None):
			πF.SetLineno(1426)
			πTemp009 = make([]πg.Param, 3)
			πTemp009[0] = πg.Param{Name: "obj", Def: nil}
			πTemp009[1] = πg.Param{Name: "file", Def: nil}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp009[2] = πg.Param{Name: "protocol", Def: πTemp013}
			πTemp012 = πg.NewFunction(πg.NewCode("dump", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]
				_ = µobj
				var µfile *πg.Object = πArgs[1]
				_ = µfile
				var µprotocol *πg.Object = πArgs[2]
				_ = µprotocol
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
					// line 1427: Pickler(file, protocol).dump(obj)
					πF.SetLineno(1427)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp001[0] = µobj
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp002[0] = µfile
					if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
						continue
					}
					πTemp002[1] = µprotocol
					if πTemp003, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßdump, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdump.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 1429: def dumps(obj, protocol=None):
			πF.SetLineno(1429)
			πTemp009 = make([]πg.Param, 2)
			πTemp009[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp009[1] = πg.Param{Name: "protocol", Def: πTemp014}
			πTemp013 = πg.NewFunction(πg.NewCode("dumps", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]
				_ = µobj
				var µprotocol *πg.Object = πArgs[1]
				_ = µprotocol
				var µfile *πg.Object = πg.UnboundLocal
				_ = µfile
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
					// line 1430: file = StringIO()
					πF.SetLineno(1430)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µfile = πTemp002
					// line 1431: Pickler(file, protocol).dump(obj)
					πF.SetLineno(1431)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp003[0] = µobj
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp004[0] = µfile
					if πE = πg.CheckLocal(πF, µprotocol, "protocol"); πE != nil {
						continue
					}
					πTemp004[1] = µprotocol
					if πTemp001, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdump, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 1432: return file.getvalue()
					πF.SetLineno(1432)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µfile, ßgetvalue, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßdumps.ToObject(), πTemp013); πE != nil {
				continue
			}
			// line 1434: def load(file):
			πF.SetLineno(1434)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "file", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("load", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfile *πg.Object = πArgs[0]
				_ = µfile
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
					// line 1435: return Unpickler(file).load()
					πF.SetLineno(1435)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp001[0] = µfile
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUnpickler); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßload, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßload.ToObject(), πTemp014); πE != nil {
				continue
			}
			// line 1437: def loads(str):
			πF.SetLineno(1437)
			πTemp009 = make([]πg.Param, 1)
			πTemp009[0] = πg.Param{Name: "str", Def: nil}
			πTemp015 = πg.NewFunction(πg.NewCode("loads", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µstr *πg.Object = πArgs[0]
				_ = µstr
				var µfile *πg.Object = πg.UnboundLocal
				_ = µfile
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
					// line 1438: file = StringIO(str)
					πF.SetLineno(1438)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µstr, "str"); πE != nil {
						continue
					}
					πTemp001[0] = µstr
					if πTemp002, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µfile = πTemp003
					// line 1439: return Unpickler(file).load()
					πF.SetLineno(1439)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
						continue
					}
					πTemp001[0] = µfile
					if πTemp002, πE = πg.ResolveGlobal(πF, ßUnpickler); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßload, nil); πE != nil {
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
			// line 1443: def _test():
			πF.SetLineno(1443)
			πTemp009 = make([]πg.Param, 0)
			πTemp016 = πg.NewFunction(πg.NewCode("_test", "/usr/lib/python2.7/pickle.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdoctest *πg.Object = πg.UnboundLocal
				_ = µdoctest
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
					// line 1444: import doctest
					πF.SetLineno(1444)
					if πTemp002, πE = πg.ImportModule(πF, "doctest"); πE != nil {
						continue
					}
					πTemp001 = πTemp002[0]
					µdoctest = πTemp001
					// line 1445: return doctest.testmod()
					πF.SetLineno(1445)
					if πE = πg.CheckLocal(πF, µdoctest, "doctest"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µdoctest, ßtestmod, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_test.ToObject(), πTemp016); πE != nil {
				continue
			}
			// line 1448: try:
			πF.SetLineno(1448)
			πF.PushCheckpoint(14)
			// line 1449: import cpyext
			πF.SetLineno(1449)
			if πTemp002, πE = πg.ImportModule(πF, "cpyext"); πE != nil {
				continue
			}
			πTemp017 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßcpyext.ToObject(), πTemp017); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			// line 1453: Pickler.dispatch[cpyext.FunctionType] = Pickler.save_global
			πF.SetLineno(1453)
			if πTemp017, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
				continue
			}
			if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßsave_global, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp017}, πTemp018); πE != nil {
				continue
			}
			if πTemp019, πE = πg.ResolveGlobal(πF, ßPickler); πE != nil {
				continue
			}
			if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßdispatch, nil); πE != nil {
				continue
			}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßcpyext); πE != nil {
				continue
			}
			if πTemp022, πE = πg.GetAttr(πF, πTemp021, ßFunctionType, nil); πE != nil {
				continue
			}
			πTemp019 = πTemp022
			if πE = πg.SetItem(πF, πTemp020, πTemp019, πTemp017); πE != nil {
				continue
			}
			goto Label13
		Label14:
			if πE == nil {
				continue
			}
			πE = nil
			πTemp006, πTemp007 = πF.ExcInfo()
			if πTemp017, πE = πg.ResolveGlobal(πF, ßImportError); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp017); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label15
			}
			πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
			continue
			// line 1450: except ImportError:
			πF.SetLineno(1450)
		Label15:
			// line 1451: pass
			πF.SetLineno(1451)
			πF.RestoreExc(nil, nil)
			goto Label13
		Label13:
			if πTemp018, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp017, πE = πg.Eq(πF, πTemp018, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp008, πE = πg.IsTrue(πF, πTemp017); πE != nil {
				continue
			}
			if πTemp008 {
				goto Label16
			}
			goto Label17
			// line 1456: if __name__ == "__main__":
			πF.SetLineno(1456)
		Label16:
			// line 1457: _test()
			πF.SetLineno(1457)
			if πTemp017, πE = πg.ResolveGlobal(πF, ß_test); πE != nil {
				continue
			}
			if πTemp018, πE = πTemp017.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label17
		Label17:
		}
		return nil, πE
	})
	πg.RegisterModule("pickle", Code)
}
