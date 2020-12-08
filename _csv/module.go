package _csv
import (
	πg "github.com/pygolin/runtime"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/_csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßDialect := πg.InternStr("Dialect")
		ßEAT_CRNL := πg.InternStr("EAT_CRNL")
		ßESCAPED_CHAR := πg.InternStr("ESCAPED_CHAR")
		ßESCAPE_IN_QUOTED_FIELD := πg.InternStr("ESCAPE_IN_QUOTED_FIELD")
		ßError := πg.InternStr("Error")
		ßException := πg.InternStr("Exception")
		ßFalse := πg.InternStr("False")
		ßIN_FIELD := πg.InternStr("IN_FIELD")
		ßIN_QUOTED_FIELD := πg.InternStr("IN_QUOTED_FIELD")
		ßKeyError := πg.InternStr("KeyError")
		ßNone := πg.InternStr("None")
		ßQUOTE_ALL := πg.InternStr("QUOTE_ALL")
		ßQUOTE_IN_QUOTED_FIELD := πg.InternStr("QUOTE_IN_QUOTED_FIELD")
		ßQUOTE_MINIMAL := πg.InternStr("QUOTE_MINIMAL")
		ßQUOTE_NONE := πg.InternStr("QUOTE_NONE")
		ßQUOTE_NONNUMERIC := πg.InternStr("QUOTE_NONNUMERIC")
		ßReader := πg.InternStr("Reader")
		ßRuntimeError := πg.InternStr("RuntimeError")
		ßSTART_FIELD := πg.InternStr("START_FIELD")
		ßSTART_RECORD := πg.InternStr("START_RECORD")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßWriter := πg.InternStr("Writer")
		ß_ := πg.InternStr("_")
		ß__all__ := πg.InternStr("__all__")
		ß__class__ := πg.InternStr("__class__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__slots__ := πg.InternStr("__slots__")
		ß__version__ := πg.InternStr("__version__")
		ß_call_dialect := πg.InternStr("_call_dialect")
		ß_delimiter := πg.InternStr("_delimiter")
		ß_dialects := πg.InternStr("_dialects")
		ß_doublequote := πg.InternStr("_doublequote")
		ß_escapechar := πg.InternStr("_escapechar")
		ß_field_limit := πg.InternStr("_field_limit")
		ß_join_append := πg.InternStr("_join_append")
		ß_join_reset := πg.InternStr("_join_reset")
		ß_lineterminator := πg.InternStr("_lineterminator")
		ß_parse_add_char := πg.InternStr("_parse_add_char")
		ß_parse_eol := πg.InternStr("_parse_eol")
		ß_parse_process_char := πg.InternStr("_parse_process_char")
		ß_parse_reset := πg.InternStr("_parse_reset")
		ß_parse_save_field := πg.InternStr("_parse_save_field")
		ß_quotechar := πg.InternStr("_quotechar")
		ß_quoting := πg.InternStr("_quoting")
		ß_skipinitialspace := πg.InternStr("_skipinitialspace")
		ß_strict := πg.InternStr("_strict")
		ßall := πg.InternStr("all")
		ßappend := πg.InternStr("append")
		ßbasestring := πg.InternStr("basestring")
		ßbool := πg.InternStr("bool")
		ßcallable := πg.InternStr("callable")
		ßdelimiter := πg.InternStr("delimiter")
		ßdialect := πg.InternStr("dialect")
		ßdoublequote := πg.InternStr("doublequote")
		ßescapechar := πg.InternStr("escapechar")
		ßfield := πg.InternStr("field")
		ßfield_size_limit := πg.InternStr("field_size_limit")
		ßfields := πg.InternStr("fields")
		ßfloat := πg.InternStr("float")
		ßget_dialect := πg.InternStr("get_dialect")
		ßgetattr := πg.InternStr("getattr")
		ßhasattr := πg.InternStr("hasattr")
		ßinput_iter := πg.InternStr("input_iter")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßiter := πg.InternStr("iter")
		ßitervalues := πg.InternStr("itervalues")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßline_num := πg.InternStr("line_num")
		ßlineterminator := πg.InternStr("lineterminator")
		ßlist := πg.InternStr("list")
		ßlist_dialects := πg.InternStr("list_dialects")
		ßlong := πg.InternStr("long")
		ßnext := πg.InternStr("next")
		ßnum_fields := πg.InternStr("num_fields")
		ßnumeric_field := πg.InternStr("numeric_field")
		ßobject := πg.InternStr("object")
		ßproperty := πg.InternStr("property")
		ßquotechar := πg.InternStr("quotechar")
		ßquoting := πg.InternStr("quoting")
		ßrange := πg.InternStr("range")
		ßreader := πg.InternStr("reader")
		ßrec := πg.InternStr("rec")
		ßregister_dialect := πg.InternStr("register_dialect")
		ßreplace := πg.InternStr("replace")
		ßrepr := πg.InternStr("repr")
		ßskipinitialspace := πg.InternStr("skipinitialspace")
		ßstate := πg.InternStr("state")
		ßstr := πg.InternStr("str")
		ßstrict := πg.InternStr("strict")
		ßtuple := πg.InternStr("tuple")
		ßundefined := πg.InternStr("undefined")
		ßunregister_dialect := πg.InternStr("unregister_dialect")
		ßwrite := πg.InternStr("write")
		ßwriteline := πg.InternStr("writeline")
		ßwriter := πg.InternStr("writer")
		ßwriterow := πg.InternStr("writerow")
		ßwriterows := πg.InternStr("writerows")
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
		var πTemp007 *πg.Dict
		_ = πTemp007
		var πTemp008 []πg.Param
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
			// line 1: __doc__ = """CSV parsing and writing.
			πF.SetLineno(1)
			// line 1: __doc__ = """CSV parsing and writing.
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("CSV parsing and writing.\n\nThis module provides classes that assist in the reading and writing\nof Comma Separated Value (CSV) files, and implements the interface\ndescribed by PEP 305.  Although many CSV files are simple to parse,\nthe format is not formally defined by a stable specification and\nis subtle enough that parsing lines of a CSV file with something\nlike line.split(\",\") is bound to fail.  The module supports three\nbasic APIs: reading, writing, and registration of dialects.\n\n\nDIALECT REGISTRATION:\n\nReaders and writers support a dialect argument, which is a convenient\nhandle on a group of settings.  When the dialect argument is a string,\nit identifies one of the dialects previously registered with the module.\nIf it is a class or instance, the attributes of the argument are used as\nthe settings for the reader or writer:\n\n    class excel:\n        delimiter = ','\n        quotechar = '\"'\n        escapechar = None\n        doublequote = True\n        skipinitialspace = False\n        lineterminator = '\\r\\n'\n        quoting = QUOTE_MINIMAL\n\nSETTINGS:\n\n    * quotechar - specifies a one-character string to use as the\n        quoting character.  It defaults to '\"'.\n    * delimiter - specifies a one-character string to use as the\n        field separator.  It defaults to ','.\n    * skipinitialspace - specifies how to interpret whitespace which\n        immediately follows a delimiter.  It defaults to False, which\n        means that whitespace immediately following a delimiter is part\n        of the following field.\n    * lineterminator -  specifies the character sequence which should\n        terminate rows.\n    * quoting - controls when quotes should be generated by the writer.\n        It can take on any of the following module constants:\n\n        csv.QUOTE_MINIMAL means only when required, for example, when a\n            field contains either the quotechar or the delimiter\n        csv.QUOTE_ALL means that quotes are always placed around fields.\n        csv.QUOTE_NONNUMERIC means that quotes are always placed around\n            fields which do not parse as integers or floating point\n            numbers.\n        csv.QUOTE_NONE means that quotes are never placed around fields.\n    * escapechar - specifies a one-character string used to escape\n        the delimiter when quoting is set to QUOTE_NONE.\n    * doublequote - controls the handling of quotes inside fields.  When\n        True, two consecutive quotes are interpreted as one during read,\n        and when writing, each quote character embedded in the data is\n        written as two quotes.\n").ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("CSV parsing and writing.\n\nThis module provides classes that assist in the reading and writing\nof Comma Separated Value (CSV) files, and implements the interface\ndescribed by PEP 305.  Although many CSV files are simple to parse,\nthe format is not formally defined by a stable specification and\nis subtle enough that parsing lines of a CSV file with something\nlike line.split(\",\") is bound to fail.  The module supports three\nbasic APIs: reading, writing, and registration of dialects.\n\n\nDIALECT REGISTRATION:\n\nReaders and writers support a dialect argument, which is a convenient\nhandle on a group of settings.  When the dialect argument is a string,\nit identifies one of the dialects previously registered with the module.\nIf it is a class or instance, the attributes of the argument are used as\nthe settings for the reader or writer:\n\n    class excel:\n        delimiter = ','\n        quotechar = '\"'\n        escapechar = None\n        doublequote = True\n        skipinitialspace = False\n        lineterminator = '\\r\\n'\n        quoting = QUOTE_MINIMAL\n\nSETTINGS:\n\n    * quotechar - specifies a one-character string to use as the\n        quoting character.  It defaults to '\"'.\n    * delimiter - specifies a one-character string to use as the\n        field separator.  It defaults to ','.\n    * skipinitialspace - specifies how to interpret whitespace which\n        immediately follows a delimiter.  It defaults to False, which\n        means that whitespace immediately following a delimiter is part\n        of the following field.\n    * lineterminator -  specifies the character sequence which should\n        terminate rows.\n    * quoting - controls when quotes should be generated by the writer.\n        It can take on any of the following module constants:\n\n        csv.QUOTE_MINIMAL means only when required, for example, when a\n            field contains either the quotechar or the delimiter\n        csv.QUOTE_ALL means that quotes are always placed around fields.\n        csv.QUOTE_NONNUMERIC means that quotes are always placed around\n            fields which do not parse as integers or floating point\n            numbers.\n        csv.QUOTE_NONE means that quotes are never placed around fields.\n    * escapechar - specifies a one-character string used to escape\n        the delimiter when quoting is set to QUOTE_NONE.\n    * doublequote - controls the handling of quotes inside fields.  When\n        True, two consecutive quotes are interpreted as one during read,\n        and when writing, each quote character embedded in the data is\n        written as two quotes.\n").ToObject()); πE != nil {
				continue
			}
			// line 59: __version__ = "1.0"
			πF.SetLineno(59)
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πg.NewStr("1.0").ToObject()); πE != nil {
				continue
			}
			// line 61: __all__ = [
			πF.SetLineno(61)
			πTemp001 = make([]*πg.Object, 21)
			πTemp001[0] = ßDialect.ToObject()
			πTemp001[1] = ßError.ToObject()
			πTemp001[2] = ßQUOTE_ALL.ToObject()
			πTemp001[3] = ßQUOTE_MINIMAL.ToObject()
			πTemp001[4] = ßQUOTE_NONE.ToObject()
			πTemp001[5] = ßQUOTE_NONNUMERIC.ToObject()
			πTemp001[6] = ßReader.ToObject()
			πTemp001[7] = ßWriter.ToObject()
			πTemp001[8] = ß__doc__.ToObject()
			πTemp001[9] = ß__version__.ToObject()
			πTemp001[10] = ß_call_dialect.ToObject()
			πTemp001[11] = ß_dialects.ToObject()
			πTemp001[12] = ß_field_limit.ToObject()
			πTemp001[13] = ßfield_size_limit.ToObject()
			πTemp001[14] = ßget_dialect.ToObject()
			πTemp001[15] = ßlist_dialects.ToObject()
			πTemp001[16] = ßreader.ToObject()
			πTemp001[17] = ßregister_dialect.ToObject()
			πTemp001[18] = ßundefined.ToObject()
			πTemp001[19] = ßunregister_dialect.ToObject()
			πTemp001[20] = ßwriter.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 69: QUOTE_MINIMAL, QUOTE_ALL, QUOTE_NONNUMERIC, QUOTE_NONE = range(4)
			πF.SetLineno(69)
			πTemp001 = πF.MakeArgs(1)
			πTemp001[0] = πg.NewInt(4).ToObject()
			if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp001)
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßQUOTE_MINIMAL.ToObject(), πTemp002); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßQUOTE_ALL.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßQUOTE_NONNUMERIC.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßQUOTE_NONE.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 70: _dialects = {}
			πF.SetLineno(70)
			πTemp007 = πg.NewDict()
			πTemp002 = πTemp007.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_dialects.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 71: _field_limit = 128 * 1024 # max parsed field size
			πF.SetLineno(71)
			if πTemp002, πE = πg.Mul(πF, πg.NewInt(128).ToObject(), πg.NewInt(1024).ToObject()); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_field_limit.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 73: class Error(Exception):
			πF.SetLineno(73)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
				continue
			}
			πTemp001[0] = πTemp004
			πTemp007 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Error", "/usr/lib/python2.7/_csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 74: pass
					πF.SetLineno(74)
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Error").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßError.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 76: class Dialect(object):
			πF.SetLineno(76)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp004
			πTemp007 = πg.NewDict()
			if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Dialect", "/usr/lib/python2.7/_csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 77: """CSV dialect
					πF.SetLineno(77)
					// line 77: """CSV dialect
					πF.SetLineno(77)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("CSV dialect\n\n    The Dialect type records CSV parsing and generation options.").ToObject()); πE != nil {
						continue
					}
					// line 81: __slots__ = ["_delimiter", "_doublequote", "_escapechar",
					πF.SetLineno(81)
					πTemp001 = make([]*πg.Object, 8)
					πTemp001[0] = ß_delimiter.ToObject()
					πTemp001[1] = ß_doublequote.ToObject()
					πTemp001[2] = ß_escapechar.ToObject()
					πTemp001[3] = ß_lineterminator.ToObject()
					πTemp001[4] = ß_quotechar.ToObject()
					πTemp001[5] = ß_quoting.ToObject()
					πTemp001[6] = ß_skipinitialspace.ToObject()
					πTemp001[7] = ß_strict.ToObject()
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 85: def __new__(cls, dialect, **kwargs):
					πF.SetLineno(85)
					πTemp003 = make([]πg.Param, 2)
					πTemp003[0] = πg.Param{Name: "cls", Def: nil}
					πTemp003[1] = πg.Param{Name: "dialect", Def: nil}
					πTemp002 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/_csv.py", πTemp003, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µdialect *πg.Object = πArgs[1]; _ = µdialect
						var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
						var µname *πg.Object = πg.UnboundLocal; _ = µname
						var µself *πg.Object = πg.UnboundLocal; _ = µself
						var µset_char *πg.Object = πg.UnboundLocal; _ = µset_char
						var µset_str *πg.Object = πg.UnboundLocal; _ = µset_str
						var µset_quoting *πg.Object = πg.UnboundLocal; _ = µset_quoting
						var µattributes *πg.Object = πg.UnboundLocal; _ = µattributes
						var µnotset *πg.Object = πg.UnboundLocal; _ = µnotset
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var µconverter *πg.Object = πg.UnboundLocal; _ = µconverter
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Object
						_ = πTemp007
						var πTemp008 []*πg.Object
						_ = πTemp008
						var πTemp009 []πg.Param
						_ = πTemp009
						var πTemp010 *πg.Dict
						_ = πTemp010
						var πTemp011 *πg.Object
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.Object
						_ = πTemp013
						var πTemp014 *πg.Object
						_ = πTemp014
						var πTemp015 bool
						_ = πTemp015
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 13: goto Label13
							case 14: goto Label14
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µkwargs); πE != nil {
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
								µname = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Add(πF, ß_.ToObject(), µname); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßDialect); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ß__slots__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, πTemp007, πTemp005); πE != nil {
								continue
							}
							πTemp004 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label4
							}
							goto Label5
							// line 88: if '_' + name not in Dialect.__slots__:
							πF.SetLineno(88)
						Label4:
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple1(µname).ToObject()
							if πTemp004, πE = πg.Mod(πF, πg.NewStr("unexpected keyword argument '%s'").ToObject(), πTemp005); πE != nil {
								continue
							}
							πTemp008[0] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 89: raise TypeError("unexpected keyword argument '%s'" %
							πF.SetLineno(89)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(µdialect != πTemp004).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label6
							}
							goto Label7
							// line 92: if dialect is not None:
							πF.SetLineno(92)
						Label6:
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πTemp008[0] = µdialect
							if πTemp001, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
								continue
							}
							πTemp008[1] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							if πTemp002, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label8
							}
							goto Label9
							// line 93: if isinstance(dialect, basestring):
							πF.SetLineno(93)
						Label8:
							// line 94: dialect = get_dialect(dialect)
							πF.SetLineno(94)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πTemp008[0] = µdialect
							if πTemp001, πE = πg.ResolveGlobal(πF, ßget_dialect); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µdialect = πTemp004
							goto Label9
						Label9:
							πTemp008 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πTemp008[0] = µdialect
							if πTemp004, πE = πg.ResolveGlobal(πF, ßDialect); πE != nil {
								continue
							}
							πTemp008[1] = πTemp004
							if πTemp004, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001 = πTemp005
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label10
							}
							πTemp008 = πF.MakeArgs(1)
							πTemp009 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/_csv.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µkwargs, ßitervalues, nil); πE != nil {
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
											µvalue = πTemp002
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 98: and all(value is None for value in kwargs.itervalues())):
										πF.SetLineno(98)
										if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
											continue
										}
										if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
											continue
										}
										πTemp002 = πg.GetBool(µvalue == πTemp003).ToObject()
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
							if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp008[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßall); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp001 = πTemp006
						Label10:
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label11
							}
							goto Label12
							// line 97: if (isinstance(dialect, Dialect)
							πF.SetLineno(97)
						Label11:
							// line 99: return dialect
							πF.SetLineno(99)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πR = µdialect
							continue
							goto Label12
						Label12:
							goto Label7
						Label7:
							// line 101: self = object.__new__(cls)
							πF.SetLineno(101)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp008[0] = µcls
							if πTemp001, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp001, ß__new__, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp005.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µself = πTemp001
							// line 104: def set_char(x):
							πF.SetLineno(104)
							πTemp009 = make([]πg.Param, 1)
							πTemp009[0] = πg.Param{Name: "x", Def: nil}
							πTemp001 = πg.NewFunction(πg.NewCode("set_char", "/usr/lib/python2.7/_csv.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
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
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πTemp001 = πg.GetBool(µx == πTemp002).ToObject()
									if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label1
									}
									goto Label2
									// line 105: if x is None:
									πF.SetLineno(105)
								Label1:
									// line 106: return None
									πF.SetLineno(106)
									if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
										continue
									}
									πR = πTemp001
									continue
									goto Label2
								Label2:
									πTemp004 = πF.MakeArgs(2)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp004[0] = µx
									if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
										continue
									}
									πTemp004[1] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
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
									if !πTemp003 {
										goto Label3
									}
									πTemp004 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp004[0] = µx
									if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
										continue
									}
									if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									if πTemp002, πE = πg.LE(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
										continue
									}
									πTemp001 = πTemp002
								Label3:
									if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp003 {
										goto Label4
									}
									goto Label5
									// line 107: if isinstance(x, str) and len(x) <= 1:
									πF.SetLineno(107)
								Label4:
									// line 108: return x
									πF.SetLineno(108)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πR = µx
									continue
									goto Label5
								Label5:
									πTemp004 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
										continue
									}
									πTemp002 = πg.NewTuple1(µname).ToObject()
									if πTemp001, πE = πg.Mod(πF, πg.NewStr("%r must be a 1-character string").ToObject(), πTemp002); πE != nil {
										continue
									}
									πTemp004[0] = πTemp001
									if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
										continue
									}
									if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp004)
									// line 109: raise TypeError("%r must be a 1-character string" % (name,))
									πF.SetLineno(109)
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
							µset_char = πTemp001
							// line 110: def set_str(x):
							πF.SetLineno(110)
							πTemp009 = make([]πg.Param, 1)
							πTemp009[0] = πg.Param{Name: "x", Def: nil}
							πTemp005 = πg.NewFunction(πg.NewCode("set_str", "/usr/lib/python2.7/_csv.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
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
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp001[0] = µx
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
									// line 111: if isinstance(x, str):
									πF.SetLineno(111)
								Label1:
									// line 112: return x
									πF.SetLineno(112)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πR = µx
									continue
									goto Label2
								Label2:
									πTemp001 = πF.MakeArgs(1)
									if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
										continue
									}
									πTemp003 = πg.NewTuple1(µname).ToObject()
									if πTemp002, πE = πg.Mod(πF, πg.NewStr("%r must be a string").ToObject(), πTemp003); πE != nil {
										continue
									}
									πTemp001[0] = πTemp002
									if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp001)
									// line 113: raise TypeError("%r must be a string" % (name,))
									πF.SetLineno(113)
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
							µset_str = πTemp005
							// line 114: def set_quoting(x):
							πF.SetLineno(114)
							πTemp009 = make([]πg.Param, 1)
							πTemp009[0] = πg.Param{Name: "x", Def: nil}
							πTemp006 = πg.NewFunction(πg.NewCode("set_quoting", "/usr/lib/python2.7/_csv.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µx *πg.Object = πArgs[0]; _ = µx
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
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πTemp002 = πF.MakeArgs(1)
									πTemp002[0] = πg.NewInt(4).ToObject()
									if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
										continue
									}
									if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									if πTemp005, πE = πg.Contains(πF, πTemp004, µx); πE != nil {
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
									// line 115: if x in range(4):
									πF.SetLineno(115)
								Label1:
									// line 116: return x
									πF.SetLineno(116)
									if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
										continue
									}
									πR = µx
									continue
									goto Label2
								Label2:
									πTemp002 = πF.MakeArgs(1)
									πTemp002[0] = πg.NewStr("bad 'quoting' value").ToObject()
									if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
										continue
									}
									if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
										continue
									}
									πF.FreeArgs(πTemp002)
									// line 117: raise TypeError("bad 'quoting' value")
									πF.SetLineno(117)
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
							µset_quoting = πTemp006
							// line 119: attributes = {"delimiter": (',', set_char),
							πF.SetLineno(119)
							πTemp010 = πg.NewDict()
							if πE = πg.CheckLocal(πF, µset_char, "set_char"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(πg.NewStr(",").ToObject(), µset_char).ToObject()
							if πE = πTemp010.SetItem(πF, ßdelimiter.ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp012, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(πTemp011, πTemp012).ToObject()
							if πE = πTemp010.SetItem(πF, ßdoublequote.ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µset_char, "set_char"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(πTemp011, µset_char).ToObject()
							if πE = πTemp010.SetItem(πF, ßescapechar.ToObject(), πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µset_str, "set_str"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(πg.NewStr("\r\n").ToObject(), µset_str).ToObject()
							if πE = πTemp010.SetItem(πF, ßlineterminator.ToObject(), πTemp007); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µset_char, "set_char"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(πg.NewStr("\"").ToObject(), µset_char).ToObject()
							if πE = πTemp010.SetItem(πF, ßquotechar.ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßQUOTE_MINIMAL); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µset_quoting, "set_quoting"); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(πTemp011, µset_quoting).ToObject()
							if πE = πTemp010.SetItem(πF, ßquoting.ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πTemp012, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(πTemp011, πTemp012).ToObject()
							if πE = πTemp010.SetItem(πF, ßskipinitialspace.ToObject(), πTemp007); πE != nil {
								continue
							}
							if πTemp011, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πTemp012, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
								continue
							}
							πTemp007 = πg.NewTuple2(πTemp011, πTemp012).ToObject()
							if πE = πTemp010.SetItem(πF, ßstrict.ToObject(), πTemp007); πE != nil {
								continue
							}
							πTemp007 = πTemp010.ToObject()
							µattributes = πTemp007
							// line 130: notset = object()
							πF.SetLineno(130)
							if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, nil, nil); πE != nil {
								continue
							}
							µnotset = πTemp011
							if πTemp011, πE = πg.ResolveGlobal(πF, ßDialect); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, πTemp011, ß__slots__, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Iter(πF, πTemp012); πE != nil {
								continue
							}
							πF.PushCheckpoint(14)
							πTemp002 = false
						Label13:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp002 {
								πF.PopCheckpoint()
								goto Label15
							}
							if πTemp011, πE = πg.Next(πF, πTemp007); πE != nil {
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
								µname = πTemp011
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(13)            
							// line 132: name = name[1:]
							πF.SetLineno(132)
							if πTemp011, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µname, πTemp011); πE != nil {
								continue
							}
							µname = πTemp012
							// line 133: value = notset
							πF.SetLineno(133)
							if πE = πg.CheckLocal(πF, µnotset, "notset"); πE != nil {
								continue
							}
							µvalue = µnotset
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µkwargs, µname); πE != nil {
								continue
							}
							πTemp011 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp011); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label16
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp011 = πg.GetBool(µdialect != πTemp012).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp011); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label17
							}
							goto Label18
							// line 134: if name in kwargs:
							πF.SetLineno(134)
						Label16:
							// line 135: value = kwargs[name]
							πF.SetLineno(135)
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp011 = µname
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µkwargs, πTemp011); πE != nil {
								continue
							}
							µvalue = πTemp012
							goto Label18
							// line 136: elif dialect is not None:
							πF.SetLineno(136)
						Label17:
							// line 137: value = getattr(dialect, name, notset)
							πF.SetLineno(137)
							πTemp008 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πTemp008[0] = µdialect
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp008[1] = µname
							if πE = πg.CheckLocal(πF, µnotset, "notset"); πE != nil {
								continue
							}
							πTemp008[2] = µnotset
							if πTemp011, πE = πg.ResolveGlobal(πF, ßgetattr); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp011.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µvalue = πTemp012
							goto Label18
						Label18:
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnotset, "notset"); πE != nil {
								continue
							}
							πTemp011 = πg.GetBool(µvalue == µnotset).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp011); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label19
							}
							goto Label20
							// line 140: if value is notset:
							πF.SetLineno(140)
						Label19:
							// line 141: value = attributes[name][0]
							πF.SetLineno(141)
							πTemp011 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp013 = µname
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µattributes, πTemp013); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, πTemp014, πTemp011); πE != nil {
								continue
							}
							µvalue = πTemp012
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.Eq(πF, µname, ßquoting.ToObject()); πE != nil {
								continue
							}
							πTemp011 = πTemp012
							if πTemp003, πE = πg.IsTrue(πF, πTemp011); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label22
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µself, ßquotechar, nil); πE != nil {
								continue
							}
							if πTemp015, πE = πg.IsTrue(πF, πTemp013); πE != nil {
								continue
							}
							πTemp012 = πg.GetBool(!πTemp015).ToObject()
							πTemp011 = πTemp012
						Label22:
							if πTemp003, πE = πg.IsTrue(πF, πTemp011); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label23
							}
							goto Label24
							// line 142: if name == 'quoting' and not self.quotechar:
							πF.SetLineno(142)
						Label23:
							// line 143: value = QUOTE_NONE
							πF.SetLineno(143)
							if πTemp011, πE = πg.ResolveGlobal(πF, ßQUOTE_NONE); πE != nil {
								continue
							}
							µvalue = πTemp011
							goto Label24
						Label24:
							goto Label21
						Label20:
							// line 145: converter = attributes[name][1]
							πF.SetLineno(145)
							πTemp011 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							πTemp013 = µname
							if πE = πg.CheckLocal(πF, µattributes, "attributes"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µattributes, πTemp013); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, πTemp014, πTemp011); πE != nil {
								continue
							}
							µconverter = πTemp012
							if πE = πg.CheckLocal(πF, µconverter, "converter"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µconverter); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label25
							}
							goto Label26
							// line 146: if converter:
							πF.SetLineno(146)
						Label25:
							// line 147: value = converter(value)
							πF.SetLineno(147)
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp008[0] = µvalue
							if πE = πg.CheckLocal(πF, µconverter, "converter"); πE != nil {
								continue
							}
							if πTemp011, πE = µconverter.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							µvalue = πTemp011
							goto Label26
						Label26:
							goto Label21
						Label21:
							// line 150: self.__dict__['_' + name] = value
							πF.SetLineno(150)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ß__dict__, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.Add(πF, ß_.ToObject(), µname); πE != nil {
								continue
							}
							πTemp013 = πTemp014
							if πE = πg.SetItem(πF, πTemp012, πTemp013, πTemp011); πE != nil {
								continue
							}
							continue
						Label14:
							if πE != nil || πR != nil {
								continue
							}
						Label15:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßdelimiter, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp011); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label27
							}
							goto Label28
							// line 152: if not self.delimiter:
							πF.SetLineno(152)
						Label27:
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("delimiter must be set").ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 153: raise TypeError("delimiter must be set")
							πF.SetLineno(153)
							πE = πF.Raise(πTemp011, nil, nil)
							continue
							goto Label28
						Label28:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßquoting, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πg.ResolveGlobal(πF, ßQUOTE_NONE); πE != nil {
								continue
							}
							if πTemp011, πE = πg.NE(πF, πTemp012, πTemp013); πE != nil {
								continue
							}
							πTemp007 = πTemp011
							if πTemp002, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if !πTemp002 {
								goto Label29
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetAttr(πF, µself, ßquotechar, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp012); πE != nil {
								continue
							}
							πTemp011 = πg.GetBool(!πTemp003).ToObject()
							πTemp007 = πTemp011
						Label29:
							if πTemp002, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label30
							}
							goto Label31
							// line 155: if self.quoting != QUOTE_NONE and not self.quotechar:
							πF.SetLineno(155)
						Label30:
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("quotechar must be set if quoting enabled").ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 156: raise TypeError("quotechar must be set if quoting enabled")
							πF.SetLineno(156)
							πE = πF.Raise(πTemp011, nil, nil)
							continue
							goto Label31
						Label31:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp011, πE = πg.GetAttr(πF, µself, ßlineterminator, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IsTrue(πF, πTemp011); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(!πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp007); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label32
							}
							goto Label33
							// line 158: if not self.lineterminator:
							πF.SetLineno(158)
						Label32:
							πTemp008 = πF.MakeArgs(1)
							πTemp008[0] = πg.NewStr("lineterminator must be set").ToObject()
							if πTemp007, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp011, πE = πTemp007.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							// line 159: raise TypeError("lineterminator must be set")
							πF.SetLineno(159)
							πE = πF.Raise(πTemp011, nil, nil)
							continue
							goto Label33
						Label33:
							// line 161: return self
							πF.SetLineno(161)
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp002); πE != nil {
						continue
					}
					// line 163: delimiter        = property(lambda self: self._delimiter)
					πF.SetLineno(163)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_csv.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 163: delimiter        = property(lambda self: self._delimiter)
							πF.SetLineno(163)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_delimiter, nil); πE != nil {
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
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßdelimiter.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 164: doublequote      = property(lambda self: self._doublequote)
					πF.SetLineno(164)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_csv.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 164: doublequote      = property(lambda self: self._doublequote)
							πF.SetLineno(164)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_doublequote, nil); πE != nil {
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
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßdoublequote.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 165: escapechar       = property(lambda self: self._escapechar)
					πF.SetLineno(165)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_csv.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 165: escapechar       = property(lambda self: self._escapechar)
							πF.SetLineno(165)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_escapechar, nil); πE != nil {
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
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßescapechar.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 166: lineterminator   = property(lambda self: self._lineterminator)
					πF.SetLineno(166)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_csv.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 166: lineterminator   = property(lambda self: self._lineterminator)
							πF.SetLineno(166)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_lineterminator, nil); πE != nil {
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
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßlineterminator.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 167: quotechar        = property(lambda self: self._quotechar)
					πF.SetLineno(167)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_csv.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 167: quotechar        = property(lambda self: self._quotechar)
							πF.SetLineno(167)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_quotechar, nil); πE != nil {
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
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßquotechar.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 168: quoting          = property(lambda self: self._quoting)
					πF.SetLineno(168)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_csv.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 168: quoting          = property(lambda self: self._quoting)
							πF.SetLineno(168)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_quoting, nil); πE != nil {
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
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßquoting.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 169: skipinitialspace = property(lambda self: self._skipinitialspace)
					πF.SetLineno(169)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_csv.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 169: skipinitialspace = property(lambda self: self._skipinitialspace)
							πF.SetLineno(169)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_skipinitialspace, nil); πE != nil {
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
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßskipinitialspace.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 170: strict           = property(lambda self: self._strict)
					πF.SetLineno(170)
					πTemp001 = πF.MakeArgs(1)
					πTemp003 = make([]πg.Param, 1)
					πTemp003[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/_csv.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 170: strict           = property(lambda self: self._strict)
							πF.SetLineno(170)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_strict, nil); πE != nil {
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
					πTemp001[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πClass.SetItem(πF, ßstrict.ToObject(), πTemp005); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp003, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp003 == nil {
				πTemp003 = πg.TypeType.ToObject()
			}
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Dialect").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDialect.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 173: def _call_dialect(dialect_inst, kwargs):
			πF.SetLineno(173)
			πTemp008 = make([]πg.Param, 2)
			πTemp008[0] = πg.Param{Name: "dialect_inst", Def: nil}
			πTemp008[1] = πg.Param{Name: "kwargs", Def: nil}
			πTemp002 = πg.NewFunction(πg.NewCode("_call_dialect", "/usr/lib/python2.7/_csv.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdialect_inst *πg.Object = πArgs[0]; _ = µdialect_inst
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
					// line 174: return Dialect(dialect_inst, **kwargs)
					πF.SetLineno(174)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µdialect_inst, "dialect_inst"); πE != nil {
						continue
					}
					πTemp001[0] = µdialect_inst
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßDialect); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, nil, µkwargs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_call_dialect.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 176: def register_dialect(name, dialect=None, **kwargs):
			πF.SetLineno(176)
			πTemp008 = make([]πg.Param, 2)
			πTemp008[0] = πg.Param{Name: "name", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp008[1] = πg.Param{Name: "dialect", Def: πTemp004}
			πTemp003 = πg.NewFunction(πg.NewCode("register_dialect", "/usr/lib/python2.7/_csv.py", πTemp008, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var µdialect *πg.Object = πArgs[1]; _ = µdialect
				var µkwargs *πg.Object = πArgs[2]; _ = µkwargs
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
					// line 177: """Create a mapping from a string name to a dialect class.
					πF.SetLineno(177)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002[0] = µname
					if πTemp003, πE = πg.ResolveGlobal(πF, ßbasestring); πE != nil {
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
					// line 179: if not isinstance(name, basestring):
					πF.SetLineno(179)
				Label1:
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = πg.NewStr("dialect name must be a string or unicode").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					// line 180: raise TypeError("dialect name must be a string or unicode")
					πF.SetLineno(180)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label2
				Label2:
					// line 182: dialect = _call_dialect(dialect, kwargs)
					πF.SetLineno(182)
					πTemp002 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
						continue
					}
					πTemp002[0] = µdialect
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					πTemp002[1] = µkwargs
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_call_dialect); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					µdialect = πTemp003
					// line 183: _dialects[name] = dialect
					πF.SetLineno(183)
					if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdialect); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_dialects); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004 = µname
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
			if πE = πF.Globals().SetItem(πF, ßregister_dialect.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 177: """Create a mapping from a string name to a dialect class.
			πF.SetLineno(177)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Create a mapping from a string name to a dialect class.\n    dialect = csv.register_dialect(name, dialect)").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßregister_dialect); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 185: def unregister_dialect(name):
			πF.SetLineno(185)
			πTemp008 = make([]πg.Param, 1)
			πTemp008[0] = πg.Param{Name: "name", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("unregister_dialect", "/usr/lib/python2.7/_csv.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 *πg.BaseException
				_ = πTemp003
				var πTemp004 *πg.Traceback
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
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 186: """Delete the name/dialect mapping associated with a string name.\n
					πF.SetLineno(186)
					// line 188: try:
					πF.SetLineno(188)
					πF.PushCheckpoint(2)
					// line 189: del _dialects[name]
					πF.SetLineno(189)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_dialects); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp002 = µname
					if πE = πg.DelItem(πF, πTemp001, πTemp002); πE != nil {
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
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
					// line 190: except KeyError:
					πF.SetLineno(190)
				Label3:
					πTemp006 = πF.MakeArgs(1)
					πTemp006[0] = πg.NewStr("unknown dialect").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					// line 191: raise Error("unknown dialect")
					πF.SetLineno(191)
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
			if πE = πF.Globals().SetItem(πF, ßunregister_dialect.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 186: """Delete the name/dialect mapping associated with a string name.\n
			πF.SetLineno(186)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Delete the name/dialect mapping associated with a string name.\n\n    csv.unregister_dialect(name)").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßunregister_dialect); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 193: def get_dialect(name):
			πF.SetLineno(193)
			πTemp008 = make([]πg.Param, 1)
			πTemp008[0] = πg.Param{Name: "name", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("get_dialect", "/usr/lib/python2.7/_csv.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µname *πg.Object = πArgs[0]; _ = µname
				var πTemp001 *πg.Object
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 194: """Return the dialect instance associated with name.
					πF.SetLineno(194)
					// line 196: try:
					πF.SetLineno(196)
					πF.PushCheckpoint(2)
					// line 197: return _dialects[name]
					πF.SetLineno(197)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp001 = µname
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_dialects); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label1
				Label2:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp004, πTemp005 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
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
					// line 198: except KeyError:
					πF.SetLineno(198)
				Label3:
					πTemp007 = πF.MakeArgs(1)
					πTemp007[0] = πg.NewStr("unknown dialect").ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp007)
					// line 199: raise Error("unknown dialect")
					πF.SetLineno(199)
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
			if πE = πF.Globals().SetItem(πF, ßget_dialect.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 194: """Return the dialect instance associated with name.
			πF.SetLineno(194)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Return the dialect instance associated with name.\n    dialect = csv.get_dialect(name)").ToObject()); πE != nil {
				continue
			}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßget_dialect); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 201: def list_dialects():
			πF.SetLineno(201)
			πTemp008 = make([]πg.Param, 0)
			πTemp006 = πg.NewFunction(πg.NewCode("list_dialects", "/usr/lib/python2.7/_csv.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 202: """Return a list of all know dialect names
					πF.SetLineno(202)
					// line 204: return list(_dialects)
					πF.SetLineno(204)
					πTemp001 = πF.MakeArgs(1)
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_dialects); πE != nil {
						continue
					}
					πTemp001[0] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßlist_dialects.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 202: """Return a list of all know dialect names
			πF.SetLineno(202)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Return a list of all know dialect names\n    names = csv.list_dialects()").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßlist_dialects); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 206: class Reader(object):
			πF.SetLineno(206)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp007 = πg.NewDict()
			if πTemp009, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp009); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Reader", "/usr/lib/python2.7/_csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp007
				_ = πClass
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
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 []πg.Param
				_ = πTemp011
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 207: """CSV reader
					πF.SetLineno(207)
					// line 207: """CSV reader
					πF.SetLineno(207)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("CSV reader\n\n    Reader objects are responsible for reading and parsing tabular data\n    in CSV format.").ToObject()); πE != nil {
						continue
					}
					// line 213: (START_RECORD, START_FIELD, ESCAPED_CHAR, IN_FIELD,
					πF.SetLineno(213)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewInt(8).ToObject()
					if πTemp002, πE = πg.ResolveClass(πF, πClass, nil, ßrange); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}, πg.TieTarget{Target: &πTemp010}}}, πTemp003); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßSTART_RECORD.ToObject(), πTemp002); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßSTART_FIELD.ToObject(), πTemp004); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßESCAPED_CHAR.ToObject(), πTemp005); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßIN_FIELD.ToObject(), πTemp006); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßIN_QUOTED_FIELD.ToObject(), πTemp007); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßESCAPE_IN_QUOTED_FIELD.ToObject(), πTemp008); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßQUOTE_IN_QUOTED_FIELD.ToObject(), πTemp009); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßEAT_CRNL.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 217: def __init__(self, iterator, dialect=None, **kwargs):
					πF.SetLineno(217)
					πTemp011 = make([]πg.Param, 3)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "iterator", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp011[2] = πg.Param{Name: "dialect", Def: πTemp003}
					πTemp002 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_csv.py", πTemp011, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µiterator *πg.Object = πArgs[1]; _ = µiterator
						var µdialect *πg.Object = πArgs[2]; _ = µdialect
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
							// line 218: self.dialect = _call_dialect(dialect, kwargs)
							πF.SetLineno(218)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πTemp001[0] = µdialect
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp001[1] = µkwargs
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_call_dialect); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdialect, πTemp002); πE != nil {
								continue
							}
							// line 219: self.input_iter = iter(iterator)
							πF.SetLineno(219)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µiterator, "iterator"); πE != nil {
								continue
							}
							πTemp001[0] = µiterator
							if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßinput_iter, πTemp002); πE != nil {
								continue
							}
							// line 220: self.line_num = 0
							πF.SetLineno(220)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_num, πTemp002); πE != nil {
								continue
							}
							// line 222: self._parse_reset()
							πF.SetLineno(222)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_parse_reset, nil); πE != nil {
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
					// line 224: def _parse_reset(self):
					πF.SetLineno(224)
					πTemp011 = make([]πg.Param, 1)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_parse_reset", "/usr/lib/python2.7/_csv.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 225: self.field = ''
							πF.SetLineno(225)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ß.ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfield, πTemp001); πE != nil {
								continue
							}
							// line 226: self.fields = []
							πF.SetLineno(226)
							πTemp002 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp002...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfields, πTemp003); πE != nil {
								continue
							}
							// line 227: self.state = self.START_RECORD
							πF.SetLineno(227)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßSTART_RECORD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp003); πE != nil {
								continue
							}
							// line 228: self.numeric_field = False
							πF.SetLineno(228)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnumeric_field, πTemp003); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_parse_reset.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 230: def __iter__(self):
					πF.SetLineno(230)
					πTemp011 = make([]πg.Param, 1)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/_csv.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 231: return self
							πF.SetLineno(231)
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 233: def next(self):
					πF.SetLineno(233)
					πTemp011 = make([]πg.Param, 1)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/_csv.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var µpos *πg.Object = πg.UnboundLocal; _ = µpos
						var µfields *πg.Object = πg.UnboundLocal; _ = µfields
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
						var πTemp006 *πg.BaseException
						_ = πTemp006
						var πTemp007 *πg.Traceback
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
							case 1: goto Label1
							case 2: goto Label2
							case 11: goto Label11
							case 12: goto Label12
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 234: self._parse_reset()
							πF.SetLineno(234)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_reset, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 235: while True:
							πF.SetLineno(235)
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 236: try:
							πF.SetLineno(236)
							πF.PushCheckpoint(5)
							// line 237: line = next(self.input_iter)
							πF.SetLineno(237)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßinput_iter, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßnext); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µline = πTemp002
							πF.PopCheckpoint()
							goto Label4
						Label5:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp006, πTemp007 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
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
							// line 238: except StopIteration:
							πF.SetLineno(238)
						Label6:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.GT(πF, πTemp008, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							goto Label8
							// line 240: if len(self.field) > 0:
							πF.SetLineno(240)
						Label7:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("newline inside string").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 241: raise Error("newline inside string")
							πF.SetLineno(241)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label8
						Label8:
							// line 242: raise
							πF.SetLineno(242)
							πE = πF.Raise(nil, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label4
						Label4:
							// line 244: self.line_num += 1
							πF.SetLineno(244)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßline_num, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_num, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, µline, πg.NewStr("\x00").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label9
							}
							goto Label10
							// line 246: if '\0' in line:
							πF.SetLineno(246)
						Label9:
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("line contains NULL byte").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 247: raise Error("line contains NULL byte")
							πF.SetLineno(247)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label10
						Label10:
							// line 248: pos = 0
							πF.SetLineno(248)
							µpos = πg.NewInt(0).ToObject()
							// line 249: while pos < len(line):
							πF.SetLineno(249)
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
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp005[0] = µline
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp001, πE = πg.LT(πF, µpos, πTemp008); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(11)            
							// line 250: pos = self._parse_process_char(line, pos)
							πF.SetLineno(250)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp005[0] = µline
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							πTemp005[1] = µpos
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_process_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µpos = πTemp002
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							// line 251: self._parse_eol()
							πF.SetLineno(251)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_eol, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßSTART_RECORD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp008); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label14
							}
							goto Label15
							// line 253: if self.state == self.START_RECORD:
							πF.SetLineno(253)
						Label14:
							// line 254: break
							πF.SetLineno(254)
							πTemp003 = true
							continue
							goto Label15
						Label15:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 256: fields = self.fields
							πF.SetLineno(256)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfields, nil); πE != nil {
								continue
							}
							µfields = πTemp001
							// line 257: self.fields = []
							πF.SetLineno(257)
							πTemp005 = make([]*πg.Object, 0)
							πTemp001 = πg.NewList(πTemp005...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfields, πTemp002); πE != nil {
								continue
							}
							// line 258: return fields
							πF.SetLineno(258)
							if πE = πg.CheckLocal(πF, µfields, "fields"); πE != nil {
								continue
							}
							πR = µfields
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 260: def _parse_process_char(self, line, pos):
					πF.SetLineno(260)
					πTemp011 = make([]πg.Param, 3)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "line", Def: nil}
					πTemp011[2] = πg.Param{Name: "pos", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("_parse_process_char", "/usr/lib/python2.7/_csv.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µline *πg.Object = πArgs[1]; _ = µline
						var µpos *πg.Object = πArgs[2]; _ = µpos
						var µc *πg.Object = πg.UnboundLocal; _ = µc
						var µpos2 *πg.Object = πg.UnboundLocal; _ = µpos2
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
						var πTemp008 *πg.Object
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 11: goto Label11
							case 12: goto Label12
							default: panic("unexpected function state")
							}
							// line 261: c = line[pos]
							πF.SetLineno(261)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							πTemp001 = µpos
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µline, πTemp001); πE != nil {
								continue
							}
							µc = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßIN_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßSTART_RECORD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßSTART_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßESCAPED_CHAR, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßIN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßESCAPE_IN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßQUOTE_IN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßEAT_CRNL, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 262: if self.state == self.IN_FIELD:
							πF.SetLineno(262)
						Label1:
							// line 264: pos2 = pos
							πF.SetLineno(264)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							µpos2 = µpos
							// line 265: while True:
							πF.SetLineno(265)
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
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(11)            
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.Contains(πF, πg.NewStr("\n\r").ToObject(), µc); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp005).ToObject()
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label14
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßescapechar, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µc, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label15
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdelimiter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µc, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label16
							}
							goto Label17
							// line 266: if c in '\n\r':
							πF.SetLineno(266)
						Label14:
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µpos2, µpos); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label19
							}
							goto Label20
							// line 268: if pos2 > pos:
							πF.SetLineno(268)
						Label19:
							// line 269: self._parse_add_char(line[pos:pos2])
							πF.SetLineno(269)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µpos, µpos2, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µline, πTemp001); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 270: pos = pos2
							πF.SetLineno(270)
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							µpos = µpos2
							goto Label20
						Label20:
							// line 271: self._parse_save_field()
							πF.SetLineno(271)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_save_field, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 272: self.state = self.EAT_CRNL
							πF.SetLineno(272)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßEAT_CRNL, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label18
							// line 273: elif c == self.dialect.escapechar:
							πF.SetLineno(273)
						Label15:
							// line 275: pos2 -= 1
							πF.SetLineno(275)
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ISub(πF, µpos2, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µpos2 = πTemp001
							// line 276: self.state = self.ESCAPED_CHAR
							πF.SetLineno(276)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßESCAPED_CHAR, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label18
							// line 277: elif c == self.dialect.delimiter:
							πF.SetLineno(277)
						Label16:
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µpos2, µpos); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label21
							}
							goto Label22
							// line 279: if pos2 > pos:
							πF.SetLineno(279)
						Label21:
							// line 280: self._parse_add_char(line[pos:pos2])
							πF.SetLineno(280)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µpos, µpos2, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µline, πTemp001); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 281: pos = pos2
							πF.SetLineno(281)
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							µpos = µpos2
							goto Label22
						Label22:
							// line 282: self._parse_save_field()
							πF.SetLineno(282)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_save_field, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 283: self.state = self.START_FIELD
							πF.SetLineno(283)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßSTART_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label18
						Label17:
							// line 286: pos2 += 1
							πF.SetLineno(286)
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.IAdd(πF, µpos2, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µpos2 = πTemp001
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp006[0] = µline
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp001, πE = πg.LT(πF, µpos2, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label23
							}
							goto Label24
							// line 287: if pos2 < len(line):
							πF.SetLineno(287)
						Label23:
							// line 288: c = line[pos2]
							πF.SetLineno(288)
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							πTemp001 = µpos2
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µline, πTemp001); πE != nil {
								continue
							}
							µc = πTemp002
							// line 289: continue
							πF.SetLineno(289)
							continue
							goto Label24
						Label24:
							goto Label18
						Label18:
							// line 290: break
							πF.SetLineno(290)
							πTemp004 = true
							continue
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µpos2, µpos); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label25
							}
							goto Label26
							// line 291: if pos2 > pos:
							πF.SetLineno(291)
						Label25:
							// line 292: self._parse_add_char(line[pos:pos2])
							πF.SetLineno(292)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µpos, µpos2, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µline, πTemp001); πE != nil {
								continue
							}
							πTemp006[0] = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 293: pos = pos2 - 1
							πF.SetLineno(293)
							if πE = πg.CheckLocal(πF, µpos2, "pos2"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Sub(πF, µpos2, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µpos = πTemp001
							goto Label26
						Label26:
							goto Label10
							// line 295: elif self.state == self.START_RECORD:
							πF.SetLineno(295)
						Label2:
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πg.NewStr("\n\r").ToObject(), µc); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label27
							}
							goto Label28
							// line 296: if c in '\n\r':
							πF.SetLineno(296)
						Label27:
							// line 297: self.state = self.EAT_CRNL
							πF.SetLineno(297)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßEAT_CRNL, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label29
						Label28:
							// line 299: self.state = self.START_FIELD
							πF.SetLineno(299)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßSTART_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							// line 301: self._parse_process_char(line, pos)
							πF.SetLineno(301)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							πTemp006[0] = µline
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							πTemp006[1] = µpos
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_process_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label29
						Label29:
							goto Label10
							// line 303: elif self.state == self.START_FIELD:
							πF.SetLineno(303)
						Label3:
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πg.NewStr("\n\r").ToObject(), µc); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label30
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßquotechar, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µc, πTemp007); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label31
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßquoting, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßQUOTE_NONE); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label31:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label32
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßescapechar, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µc, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label33
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µc, πg.NewStr(" ").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label34
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßskipinitialspace, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label34:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label35
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdelimiter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µc, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label36
							}
							goto Label37
							// line 304: if c in '\n\r':
							πF.SetLineno(304)
						Label30:
							// line 306: self._parse_save_field()
							πF.SetLineno(306)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_save_field, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 307: self.state = self.EAT_CRNL
							πF.SetLineno(307)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßEAT_CRNL, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label38
							// line 308: elif (c == self.dialect.quotechar
							πF.SetLineno(308)
						Label32:
							// line 311: self.state = self.IN_QUOTED_FIELD
							πF.SetLineno(311)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßIN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label38
							// line 312: elif c == self.dialect.escapechar:
							πF.SetLineno(312)
						Label33:
							// line 314: self.state = self.ESCAPED_CHAR
							πF.SetLineno(314)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßESCAPED_CHAR, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label38
							// line 315: elif c == ' ' and self.dialect.skipinitialspace:
							πF.SetLineno(315)
						Label35:
							// line 317: pass
							πF.SetLineno(317)
							goto Label38
							// line 318: elif c == self.dialect.delimiter:
							πF.SetLineno(318)
						Label36:
							// line 320: self._parse_save_field()
							πF.SetLineno(320)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_save_field, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							goto Label38
						Label37:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßquoting, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ResolveGlobal(πF, ßQUOTE_NONNUMERIC); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp003, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label39
							}
							goto Label40
							// line 323: if self.dialect.quoting == QUOTE_NONNUMERIC:
							πF.SetLineno(323)
						Label39:
							// line 324: self.numeric_field = True
							πF.SetLineno(324)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnumeric_field, πTemp002); πE != nil {
								continue
							}
							goto Label40
						Label40:
							// line 325: self._parse_add_char(c)
							πF.SetLineno(325)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 326: self.state = self.IN_FIELD
							πF.SetLineno(326)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßIN_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label38
						Label38:
							goto Label10
							// line 328: elif self.state == self.ESCAPED_CHAR:
							πF.SetLineno(328)
						Label4:
							// line 329: self._parse_add_char(c)
							πF.SetLineno(329)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 330: self.state = self.IN_FIELD
							πF.SetLineno(330)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßIN_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label10
							// line 332: elif self.state == self.IN_QUOTED_FIELD:
							πF.SetLineno(332)
						Label5:
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßescapechar, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µc, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label41
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßquotechar, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µc, πTemp007); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label42
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßquoting, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßQUOTE_NONE); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label42:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label43
							}
							goto Label44
							// line 333: if c == self.dialect.escapechar:
							πF.SetLineno(333)
						Label41:
							// line 335: self.state = self.ESCAPE_IN_QUOTED_FIELD
							πF.SetLineno(335)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßESCAPE_IN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label45
							// line 336: elif (c == self.dialect.quotechar
							πF.SetLineno(336)
						Label43:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdoublequote, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label46
							}
							goto Label47
							// line 338: if self.dialect.doublequote:
							πF.SetLineno(338)
						Label46:
							// line 340: self.state = self.QUOTE_IN_QUOTED_FIELD
							πF.SetLineno(340)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßQUOTE_IN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label48
						Label47:
							// line 343: self.state = self.IN_FIELD
							πF.SetLineno(343)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßIN_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label48
						Label48:
							goto Label45
						Label44:
							// line 346: self._parse_add_char(c)
							πF.SetLineno(346)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							goto Label45
						Label45:
							goto Label10
							// line 348: elif self.state == self.ESCAPE_IN_QUOTED_FIELD:
							πF.SetLineno(348)
						Label6:
							// line 349: self._parse_add_char(c)
							πF.SetLineno(349)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 350: self.state = self.IN_QUOTED_FIELD
							πF.SetLineno(350)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßIN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label10
							// line 352: elif self.state == self.QUOTE_IN_QUOTED_FIELD:
							πF.SetLineno(352)
						Label7:
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßquotechar, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µc, πTemp007); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp004 {
								goto Label49
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßquoting, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßQUOTE_NONE); πE != nil {
								continue
							}
							if πTemp002, πE = πg.NE(πF, πTemp007, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πTemp002
						Label49:
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label50
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdelimiter, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µc, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label51
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πg.NewStr("\r\n").ToObject(), µc); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label52
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstrict, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label53
							}
							goto Label54
							// line 354: if (c == self.dialect.quotechar
							πF.SetLineno(354)
						Label50:
							// line 357: self._parse_add_char(c)
							πF.SetLineno(357)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 358: self.state = self.IN_QUOTED_FIELD
							πF.SetLineno(358)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßIN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label55
							// line 359: elif c == self.dialect.delimiter:
							πF.SetLineno(359)
						Label51:
							// line 361: self._parse_save_field()
							πF.SetLineno(361)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_save_field, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 362: self.state = self.START_FIELD
							πF.SetLineno(362)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßSTART_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label55
							// line 363: elif c in '\r\n':
							πF.SetLineno(363)
						Label52:
							// line 365: self._parse_save_field()
							πF.SetLineno(365)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_save_field, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 366: self.state = self.EAT_CRNL
							πF.SetLineno(366)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßEAT_CRNL, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label55
							// line 367: elif not self.dialect.strict:
							πF.SetLineno(367)
						Label53:
							// line 368: self._parse_add_char(c)
							πF.SetLineno(368)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp006[0] = µc
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 369: self.state = self.IN_FIELD
							πF.SetLineno(369)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßIN_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label55
						Label54:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp003, ßdelimiter, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp003, ßquotechar, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("'%c' expected after '%c'").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 371: raise Error("'%c' expected after '%c'" %
							πF.SetLineno(371)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label55
						Label55:
							goto Label10
							// line 374: elif self.state == self.EAT_CRNL:
							πF.SetLineno(374)
						Label8:
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Contains(πF, πg.NewStr("\r\n").ToObject(), µc); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label56
							}
							goto Label57
							// line 375: if c not in '\r\n':
							πF.SetLineno(375)
						Label56:
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("new-line character seen in unquoted field - do you need to open the file in universal-newline mode?").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 376: raise Error("new-line character seen in unquoted field - "
							πF.SetLineno(376)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label57
						Label57:
							goto Label10
						Label9:
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple1(πTemp003).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("unknown state: %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 381: raise RuntimeError("unknown state: %r" % (self.state,))
							πF.SetLineno(381)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label10
						Label10:
							// line 383: return pos + 1
							πF.SetLineno(383)
							if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, µpos, πg.NewInt(1).ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_parse_process_char.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 385: def _parse_eol(self):
					πF.SetLineno(385)
					πTemp011 = make([]πg.Param, 1)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_parse_eol", "/usr/lib/python2.7/_csv.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßEAT_CRNL, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßSTART_RECORD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßIN_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßSTART_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label4
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßESCAPED_CHAR, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßIN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßESCAPE_IN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßQUOTE_IN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label8
							}
							goto Label9
							// line 386: if self.state == self.EAT_CRNL:
							πF.SetLineno(386)
						Label1:
							// line 387: self.state = self.START_RECORD
							πF.SetLineno(387)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßSTART_RECORD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label10
							// line 388: elif self.state == self.START_RECORD:
							πF.SetLineno(388)
						Label2:
							// line 390: pass
							πF.SetLineno(390)
							goto Label10
							// line 391: elif self.state == self.IN_FIELD:
							πF.SetLineno(391)
						Label3:
							// line 394: self._parse_save_field()
							πF.SetLineno(394)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_save_field, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 395: self.state = self.START_RECORD
							πF.SetLineno(395)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßSTART_RECORD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label10
							// line 396: elif self.state == self.START_FIELD:
							πF.SetLineno(396)
						Label4:
							// line 398: self._parse_save_field()
							πF.SetLineno(398)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_save_field, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 399: self.state = self.START_RECORD
							πF.SetLineno(399)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßSTART_RECORD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label10
							// line 400: elif self.state == self.ESCAPED_CHAR:
							πF.SetLineno(400)
						Label5:
							// line 401: self._parse_add_char('\n')
							πF.SetLineno(401)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 402: self.state = self.IN_FIELD
							πF.SetLineno(402)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßIN_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label10
							// line 403: elif self.state == self.IN_QUOTED_FIELD:
							πF.SetLineno(403)
						Label6:
							// line 404: pass
							πF.SetLineno(404)
							goto Label10
							// line 405: elif self.state == self.ESCAPE_IN_QUOTED_FIELD:
							πF.SetLineno(405)
						Label7:
							// line 406: self._parse_add_char('\n')
							πF.SetLineno(406)
							πTemp005 = πF.MakeArgs(1)
							πTemp005[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_add_char, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 407: self.state = self.IN_QUOTED_FIELD
							πF.SetLineno(407)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßIN_QUOTED_FIELD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label10
							// line 408: elif self.state == self.QUOTE_IN_QUOTED_FIELD:
							πF.SetLineno(408)
						Label8:
							// line 410: self._parse_save_field()
							πF.SetLineno(410)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_parse_save_field, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 411: self.state = self.START_RECORD
							πF.SetLineno(411)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßSTART_RECORD, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstate, πTemp002); πE != nil {
								continue
							}
							goto Label10
						Label9:
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßstate, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple1(πTemp003).ToObject()
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("unknown state: %r").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßRuntimeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							// line 413: raise RuntimeError("unknown state: %r" % (self.state,))
							πF.SetLineno(413)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label10
						Label10:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß_parse_eol.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 415: def _parse_save_field(self):
					πF.SetLineno(415)
					πTemp011 = make([]πg.Param, 1)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("_parse_save_field", "/usr/lib/python2.7/_csv.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfield *πg.Object = πg.UnboundLocal; _ = µfield
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
							// line 416: field, self.field = self.field, ''
							πF.SetLineno(416)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple2(πTemp002, ß.ToObject()).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}}}, πTemp001); πE != nil {
								continue
							}
							µfield = πTemp002
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfield, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnumeric_field, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 417: if self.numeric_field:
							πF.SetLineno(417)
						Label1:
							// line 418: self.numeric_field = False
							πF.SetLineno(418)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnumeric_field, πTemp002); πE != nil {
								continue
							}
							// line 419: field = float(field)
							πF.SetLineno(419)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp005[0] = µfield
							if πTemp001, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µfield = πTemp002
							goto Label2
						Label2:
							// line 420: self.fields.append(field)
							πF.SetLineno(420)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp005[0] = µfield
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfields, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_parse_save_field.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 422: def _parse_add_char(self, c):
					πF.SetLineno(422)
					πTemp011 = make([]πg.Param, 2)
					πTemp011[0] = πg.Param{Name: "self", Def: nil}
					πTemp011[1] = πg.Param{Name: "c", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("_parse_add_char", "/usr/lib/python2.7/_csv.py", πTemp011, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µc *πg.Object = πArgs[1]; _ = µc
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
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
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
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp003[0] = µc
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ResolveGlobal(πF, ß_field_limit); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp002, πTemp004); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 423: if len(self.field) + len(c) > _field_limit:
							πF.SetLineno(423)
						Label1:
							πTemp003 = πF.MakeArgs(1)
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_field_limit); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("field larger than field limit (%d)").ToObject(), πTemp002); πE != nil {
								continue
							}
							πTemp003[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							// line 424: raise Error("field larger than field limit (%d)" % (_field_limit))
							πF.SetLineno(424)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label2
						Label2:
							// line 425: self.field += c
							πF.SetLineno(425)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfield, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, µc); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfield, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_parse_add_char.ToObject(), πTemp009); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp010, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp010 == nil {
				πTemp010 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp010.Call(πF, []*πg.Object{πg.NewStr("Reader").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßReader.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 428: class Writer(object):
			πF.SetLineno(428)
			πTemp001 = make([]*πg.Object, 1)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp001[0] = πTemp011
			πTemp007 = πg.NewDict()
			if πTemp009, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp007.SetItem(πF, ß__module__.ToObject(), πTemp009); πE != nil {
				continue
			}
			_, πE = πg.NewCode("Writer", "/usr/lib/python2.7/_csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 429: """CSV writer
					πF.SetLineno(429)
					// line 429: """CSV writer
					πF.SetLineno(429)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("CSV writer\n\n    Writer objects are responsible for generating tabular data\n    in CSV format from sequence input.").ToObject()); πE != nil {
						continue
					}
					// line 434: def __init__(self, file, dialect=None, **kwargs):
					πF.SetLineno(434)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "file", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "dialect", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/_csv.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfile *πg.Object = πArgs[1]; _ = µfile
						var µdialect *πg.Object = πArgs[2]; _ = µdialect
						var µkwargs *πg.Object = πArgs[3]; _ = µkwargs
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
								continue
							}
							πTemp004[0] = µfile
							πTemp004[1] = ßwrite.ToObject()
							if πTemp005, πE = πg.ResolveGlobal(πF, ßhasattr); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πTemp006
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label1
							}
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfile, "file"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µfile, ßwrite, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßcallable); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							πTemp002 = πTemp006
						Label1:
							if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label2
							}
							goto Label3
							// line 435: if not (hasattr(file, 'write') and callable(file.write)):
							πF.SetLineno(435)
						Label2:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("argument 1 must have a 'write' method").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 436: raise TypeError("argument 1 must have a 'write' method")
							πF.SetLineno(436)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label3
						Label3:
							// line 437: self.writeline = file.write
							πF.SetLineno(437)
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
							if πE = πg.SetAttr(πF, µself, ßwriteline, πTemp002); πE != nil {
								continue
							}
							// line 438: self.dialect = _call_dialect(dialect, kwargs)
							πF.SetLineno(438)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πTemp004[0] = µdialect
							if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
								continue
							}
							πTemp004[1] = µkwargs
							if πTemp001, πE = πg.ResolveGlobal(πF, ß_call_dialect); πE != nil {
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
							if πE = πg.SetAttr(πF, µself, ßdialect, πTemp001); πE != nil {
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
					// line 440: def _join_reset(self):
					πF.SetLineno(440)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_join_reset", "/usr/lib/python2.7/_csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 441: self.rec = []
							πF.SetLineno(441)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrec, πTemp003); πE != nil {
								continue
							}
							// line 442: self.num_fields = 0
							πF.SetLineno(442)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnum_fields, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_join_reset.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 444: def _join_append(self, field, quoted, quote_empty):
					πF.SetLineno(444)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "field", Def: nil}
					πTemp002[2] = πg.Param{Name: "quoted", Def: nil}
					πTemp002[3] = πg.Param{Name: "quote_empty", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_join_append", "/usr/lib/python2.7/_csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µfield *πg.Object = πArgs[1]; _ = µfield
						var µquoted *πg.Object = πArgs[2]; _ = µquoted
						var µquote_empty *πg.Object = πArgs[3]; _ = µquote_empty
						var µdialect *πg.Object = πg.UnboundLocal; _ = µdialect
						var µneed_escape *πg.Object = πg.UnboundLocal; _ = µneed_escape
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 17: goto Label17
							case 18: goto Label18
							case 6: goto Label6
							case 7: goto Label7
							default: panic("unexpected function state")
							}
							// line 445: dialect = self.dialect
							πF.SetLineno(445)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							µdialect = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßnum_fields, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 447: if self.num_fields > 0:
							πF.SetLineno(447)
						Label1:
							// line 448: self.rec.append(dialect.delimiter)
							πF.SetLineno(448)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdialect, ßdelimiter, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrec, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdialect, ßquoting, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßQUOTE_NONE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 450: if dialect.quoting == QUOTE_NONE:
							πF.SetLineno(450)
						Label3:
							// line 451: need_escape = tuple(dialect.lineterminator) + (
							πF.SetLineno(451)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdialect, ßlineterminator, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µdialect, ßescapechar, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µdialect, ßdelimiter, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µdialect, ßquotechar, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple3(πTemp006, πTemp007, πTemp008).ToObject()
							if πTemp001, πE = πg.Add(πF, πTemp005, πTemp002); πE != nil {
								continue
							}
							µneed_escape = πTemp001
							goto Label5
						Label4:
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µdialect, ßlineterminator, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp005
							if πTemp005, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µdialect, ßdelimiter, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µdialect, ßescapechar, nil); πE != nil {
								continue
							}
							πTemp005 = πg.NewTuple2(πTemp007, πTemp008).ToObject()
							if πTemp002, πE = πg.Add(πF, πTemp006, πTemp005); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
								continue
							}
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
							if πTemp002, πE = πg.Next(πF, πTemp001); πE != nil {
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
								µc = πTemp002
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(6)            
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp002 = µc
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp009 {
								goto Label9
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Contains(πF, µfield, µc); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp010).ToObject()
							πTemp002 = πTemp005
						Label9:
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label10
							}
							goto Label11
							// line 458: if c and c in field:
							πF.SetLineno(458)
						Label10:
							// line 459: quoted = True
							πF.SetLineno(459)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µquoted = πTemp002
							goto Label11
						Label11:
							continue
						Label7:
							if πE != nil || πR != nil {
								continue
							}
						Label8:
							// line 461: need_escape = ()
							πF.SetLineno(461)
							πTemp001 = πg.NewTuple0().ToObject()
							µneed_escape = πTemp001
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdialect, ßquotechar, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Contains(πF, µfield, πTemp002); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp003).ToObject()
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label12
							}
							goto Label13
							// line 462: if dialect.quotechar in field:
							πF.SetLineno(462)
						Label12:
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdialect, ßdoublequote, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label14
							}
							goto Label15
							// line 463: if dialect.doublequote:
							πF.SetLineno(463)
						Label14:
							// line 464: field = field.replace(dialect.quotechar,
							πF.SetLineno(464)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdialect, ßquotechar, nil); πE != nil {
								continue
							}
							πTemp004[0] = πTemp001
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdialect, ßquotechar, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mul(πF, πTemp002, πg.NewInt(2).ToObject()); πE != nil {
								continue
							}
							πTemp004[1] = πTemp001
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µfield, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µfield = πTemp002
							// line 466: quoted = True
							πF.SetLineno(466)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µquoted = πTemp001
							goto Label16
						Label15:
							// line 468: need_escape = (dialect.quotechar,)
							πF.SetLineno(468)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdialect, ßquotechar, nil); πE != nil {
								continue
							}
							πTemp001 = πg.NewTuple1(πTemp002).ToObject()
							µneed_escape = πTemp001
							goto Label16
						Label16:
							goto Label13
						Label13:
							goto Label5
						Label5:
							if πE = πg.CheckLocal(πF, µneed_escape, "need_escape"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µneed_escape); πE != nil {
								continue
							}
							πF.PushCheckpoint(18)
							πTemp003 = false
						Label17:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
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
								πTemp009 = !isStop
							} else {
								πTemp009 = true
								µc = πTemp002
							}
							if πE != nil || !πTemp009 {
								continue
							}
							πF.PushCheckpoint(17)            
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp002 = µc
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp009 {
								goto Label20
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Contains(πF, µfield, µc); πE != nil {
								continue
							}
							πTemp005 = πg.GetBool(πTemp010).ToObject()
							πTemp002 = πTemp005
						Label20:
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label21
							}
							goto Label22
							// line 472: if c and c in field:
							πF.SetLineno(472)
						Label21:
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µdialect, ßescapechar, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label23
							}
							goto Label24
							// line 473: if not dialect.escapechar:
							πF.SetLineno(473)
						Label23:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("need to escape, but no escapechar set").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 474: raise Error("need to escape, but no escapechar set")
							πF.SetLineno(474)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label24
						Label24:
							// line 475: field = field.replace(c, dialect.escapechar + c)
							πF.SetLineno(475)
							πTemp004 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							πTemp004[0] = µc
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µdialect, ßescapechar, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp005, µc); πE != nil {
								continue
							}
							πTemp004[1] = πTemp002
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µfield, ßreplace, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							µfield = πTemp005
							goto Label22
						Label22:
							continue
						Label18:
							if πE != nil || πR != nil {
								continue
							}
						Label19:
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µfield, ß.ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp003 {
								goto Label25
							}
							if πE = πg.CheckLocal(πF, µquote_empty, "quote_empty"); πE != nil {
								continue
							}
							πTemp001 = µquote_empty
						Label25:
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label26
							}
							goto Label27
							// line 478: if field == '' and quote_empty:
							πF.SetLineno(478)
						Label26:
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdialect, ßquoting, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ResolveGlobal(πF, ßQUOTE_NONE); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label28
							}
							goto Label29
							// line 479: if dialect.quoting == QUOTE_NONE:
							πF.SetLineno(479)
						Label28:
							πTemp004 = πF.MakeArgs(1)
							πTemp004[0] = πg.NewStr("single empty field record must be quoted").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 480: raise Error("single empty field record must be quoted")
							πF.SetLineno(480)
							πE = πF.Raise(πTemp002, nil, nil)
							continue
							goto Label29
						Label29:
							// line 481: quoted = 1
							πF.SetLineno(481)
							µquoted = πg.NewInt(1).ToObject()
							goto Label27
						Label27:
							if πE = πg.CheckLocal(πF, µquoted, "quoted"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µquoted); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label30
							}
							goto Label31
							// line 483: if quoted:
							πF.SetLineno(483)
						Label30:
							// line 484: field = dialect.quotechar + field + dialect.quotechar
							πF.SetLineno(484)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µdialect, ßquotechar, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Add(πF, πTemp005, µfield); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µdialect, ßquotechar, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Add(πF, πTemp002, πTemp005); πE != nil {
								continue
							}
							µfield = πTemp001
							goto Label31
						Label31:
							// line 486: self.rec.append(field)
							πF.SetLineno(486)
							πTemp004 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp004[0] = µfield
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrec, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							// line 487: self.num_fields += 1
							πF.SetLineno(487)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnum_fields, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, πTemp001, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßnum_fields, πTemp002); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_join_append.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 491: def writerow(self, row):
					πF.SetLineno(491)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "row", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("writerow", "/usr/lib/python2.7/_csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrow *πg.Object = πArgs[1]; _ = µrow
						var µdialect *πg.Object = πg.UnboundLocal; _ = µdialect
						var µrowlen *πg.Object = πg.UnboundLocal; _ = µrowlen
						var µfield *πg.Object = πg.UnboundLocal; _ = µfield
						var µquoted *πg.Object = πg.UnboundLocal; _ = µquoted
						var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
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
							case 2: goto Label2
							case 11: goto Label11
							case 4: goto Label4
							case 5: goto Label5
							default: panic("unexpected function state")
							}
							// line 492: dialect = self.dialect
							πF.SetLineno(492)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßdialect, nil); πE != nil {
								continue
							}
							µdialect = πTemp001
							// line 493: try:
							πF.SetLineno(493)
							πF.PushCheckpoint(2)
							// line 494: rowlen = len(row)
							πF.SetLineno(494)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp002[0] = µrow
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µrowlen = πTemp003
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
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
							// line 495: except TypeError:
							πF.SetLineno(495)
						Label3:
							πTemp002 = πF.MakeArgs(1)
							πTemp002[0] = πg.NewStr("sequence expected").ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 496: raise Error("sequence expected")
							πF.SetLineno(496)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 499: self._join_reset()
							πF.SetLineno(499)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_join_reset, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µrow); πE != nil {
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
								µfield = πTemp003
							}
							if πE != nil || !πTemp007 {
								continue
							}
							πF.PushCheckpoint(4)            
							// line 502: quoted = False
							πF.SetLineno(502)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µquoted = πTemp003
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µdialect, ßquoting, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßQUOTE_NONNUMERIC); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label7
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µdialect, ßquoting, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßQUOTE_ALL); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, πTemp008, πTemp009); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label8
							}
							goto Label9
							// line 503: if dialect.quoting == QUOTE_NONNUMERIC:
							πF.SetLineno(503)
						Label7:
							// line 504: try:
							πF.SetLineno(504)
							πF.PushCheckpoint(11)
							// line 505: float(field)
							πF.SetLineno(505)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp002[0] = µfield
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πF.PopCheckpoint()
							goto Label10
						Label11:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp004, πTemp005 = πF.ExcInfo()
							goto Label12
							// line 506: except:
							πF.SetLineno(506)
						Label12:
							// line 507: quoted = True
							πF.SetLineno(507)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µquoted = πTemp003
							πF.RestoreExc(nil, nil)
							goto Label10
						Label10:
							goto Label9
							// line 510: elif dialect.quoting == QUOTE_ALL:
							πF.SetLineno(510)
						Label8:
							// line 511: quoted = True
							πF.SetLineno(511)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µquoted = πTemp003
							goto Label9
						Label9:
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp003 = πg.GetBool(µfield == πTemp008).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label13
							}
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp002[0] = µfield
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp007, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label14
							}
							goto Label15
							// line 513: if field is None:
							πF.SetLineno(513)
						Label13:
							// line 514: value = ""
							πF.SetLineno(514)
							µvalue = ß.ToObject()
							goto Label16
							// line 515: elif isinstance(field, float):
							πF.SetLineno(515)
						Label14:
							// line 516: value = repr(field)
							πF.SetLineno(516)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp002[0] = µfield
							if πTemp003, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µvalue = πTemp008
							goto Label16
						Label15:
							// line 518: value = str(field)
							πF.SetLineno(518)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µfield, "field"); πE != nil {
								continue
							}
							πTemp002[0] = µfield
							if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							µvalue = πTemp008
							goto Label16
						Label16:
							// line 519: self._join_append(value, quoted, rowlen == 1)
							πF.SetLineno(519)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							πTemp002[0] = µvalue
							if πE = πg.CheckLocal(πF, µquoted, "quoted"); πE != nil {
								continue
							}
							πTemp002[1] = µquoted
							if πE = πg.CheckLocal(πF, µrowlen, "rowlen"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Eq(πF, µrowlen, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp002[2] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_join_append, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 522: self.rec.append(dialect.lineterminator)
							πF.SetLineno(522)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µdialect, ßlineterminator, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp001
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrec, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßappend, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							// line 524: self.writeline(''.join(self.rec))
							πF.SetLineno(524)
							πTemp002 = πF.MakeArgs(1)
							πTemp010 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßrec, nil); πE != nil {
								continue
							}
							πTemp010[0] = πTemp001
							if πTemp001, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp010)
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßwriteline, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwriterow.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 526: def writerows(self, rows):
					πF.SetLineno(526)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "rows", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("writerows", "/usr/lib/python2.7/_csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrows *πg.Object = πArgs[1]; _ = µrows
						var µrow *πg.Object = πg.UnboundLocal; _ = µrow
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
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, µrows); πE != nil {
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
								µrow = πTemp004
							}
							if πE != nil || !πTemp003 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 528: self.writerow(row)
							πF.SetLineno(528)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp005[0] = µrow
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßwriterow, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwriterows.ToObject(), πTemp006); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp010, πE = πTemp007.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp010 == nil {
				πTemp010 = πg.TypeType.ToObject()
			}
			if πTemp011, πE = πTemp010.Call(πF, []*πg.Object{πg.NewStr("Writer").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp007.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßWriter.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 530: def reader(*args, **kwargs):
			πF.SetLineno(530)
			πTemp008 = make([]πg.Param, 0)
			πTemp009 = πg.NewFunction(πg.NewCode("reader", "/usr/lib/python2.7/_csv.py", πTemp008, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
					// line 531: """
					πF.SetLineno(531)
					// line 546: return Reader(*args, **kwargs)
					πF.SetLineno(546)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßReader); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwargs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßreader.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 531: """
			πF.SetLineno(531)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("\n    csv_reader = reader(iterable [, dialect='excel']\n                       [optional keyword args])\n    for row in csv_reader:\n        process(row)\n\n    The \"iterable\" argument can be any object that returns a line\n    of input for each iteration, such as a file object or a list.  The\n    optional \"dialect\" parameter is discussed below.  The function\n    also accepts optional keyword arguments which override settings\n    provided by the dialect.\n\n    The returned object is an iterator.  Each iteration returns a row\n    of the CSV file (which can span multiple input lines)").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßreader); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 548: def writer(*args, **kwargs):
			πF.SetLineno(548)
			πTemp008 = make([]πg.Param, 0)
			πTemp010 = πg.NewFunction(πg.NewCode("writer", "/usr/lib/python2.7/_csv.py", πTemp008, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µargs *πg.Object = πArgs[0]; _ = µargs
				var µkwargs *πg.Object = πArgs[1]; _ = µkwargs
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
					// line 549: """
					πF.SetLineno(549)
					// line 562: return Writer(*args, **kwargs)
					πF.SetLineno(562)
					if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkwargs, "kwargs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ResolveGlobal(πF, ßWriter); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Invoke(πF, πTemp001, nil, µargs, nil, µkwargs); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßwriter.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 549: """
			πF.SetLineno(549)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("\n    csv_writer = csv.writer(fileobj [, dialect='excel']\n                            [optional keyword args])\n    for row in sequence:\n        csv_writer.writerow(row)\n\n    [or]\n\n    csv_writer = csv.writer(fileobj [, dialect='excel']\n                            [optional keyword args])\n    csv_writer.writerows(rows)\n\n    The \"fileobj\" argument can be any object that supports the file API.").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßwriter); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 565: undefined = object()
			πF.SetLineno(565)
			if πTemp011, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			if πTemp012, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßundefined.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 566: def field_size_limit(limit=undefined):
			πF.SetLineno(566)
			πTemp008 = make([]πg.Param, 1)
			if πTemp012, πE = πg.ResolveGlobal(πF, ßundefined); πE != nil {
				continue
			}
			πTemp008[0] = πg.Param{Name: "limit", Def: πTemp012}
			πTemp011 = πg.NewFunction(πg.NewCode("field_size_limit", "/usr/lib/python2.7/_csv.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µlimit *πg.Object = πArgs[0]; _ = µlimit
				var µold_limit *πg.Object = πg.UnboundLocal; _ = µold_limit
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 567: """Sets an upper limit on parsed fields.
					πF.SetLineno(567)
					// line 573: global _field_limit
					πF.SetLineno(573)
					// line 574: old_limit = _field_limit
					πF.SetLineno(574)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_field_limit); πE != nil {
						continue
					}
					µold_limit = πTemp001
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßundefined); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µlimit != πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 576: if limit is not undefined:
					πF.SetLineno(576)
				Label1:
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					πTemp004[0] = µlimit
					if πTemp005, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp005, πTemp006).ToObject()
					πTemp004[1] = πTemp002
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
						goto Label3
					}
					goto Label4
					// line 577: if not isinstance(limit, (int, long)):
					πF.SetLineno(577)
				Label3:
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, µlimit, ß__class__, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, πTemp005, ß__name__, nil); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(πTemp006).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("int expected, got %s").ToObject(), πTemp002); πE != nil {
						continue
					}
					πTemp004[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 578: raise TypeError("int expected, got %s" %
					πF.SetLineno(578)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label4
				Label4:
					// line 580: _field_limit = limit
					πF.SetLineno(580)
					if πE = πg.CheckLocal(πF, µlimit, "limit"); πE != nil {
						continue
					}
					if πE = πF.Globals().SetItem(πF, ß_field_limit.ToObject(), µlimit); πE != nil {
						continue
					}
					goto Label2
				Label2:
					// line 582: return old_limit
					πF.SetLineno(582)
					if πE = πg.CheckLocal(πF, µold_limit, "old_limit"); πE != nil {
						continue
					}
					πR = µold_limit
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfield_size_limit.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 567: """Sets an upper limit on parsed fields.
			πF.SetLineno(567)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Sets an upper limit on parsed fields.\n    csv.field_size_limit([limit])\n\n    Returns old limit. If limit is not given, no new limit is set and\n    the old limit is returned").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßfield_size_limit); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("_csv", Code)
}

