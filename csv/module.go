package csv
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/_csv"
	// _ "github.com/pygolin/stdlib/functools"
	// _ "github.com/pygolin/stdlib/StringIO"
	// _ "github.com/pygolin/stdlib/re"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßDOTALL := πg.InternStr("DOTALL")
		ßDialect := πg.InternStr("Dialect")
		ßDictReader := πg.InternStr("DictReader")
		ßDictWriter := πg.InternStr("DictWriter")
		ßError := πg.InternStr("Error")
		ßFalse := πg.InternStr("False")
		ßKeyError := πg.InternStr("KeyError")
		ßMULTILINE := πg.InternStr("MULTILINE")
		ßNone := πg.InternStr("None")
		ßOverflowError := πg.InternStr("OverflowError")
		ßQUOTE_ALL := πg.InternStr("QUOTE_ALL")
		ßQUOTE_MINIMAL := πg.InternStr("QUOTE_MINIMAL")
		ßQUOTE_NONE := πg.InternStr("QUOTE_NONE")
		ßQUOTE_NONNUMERIC := πg.InternStr("QUOTE_NONNUMERIC")
		ßSniffer := πg.InternStr("Sniffer")
		ßStopIteration := πg.InternStr("StopIteration")
		ßStringIO := πg.InternStr("StringIO")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß_Dialect := πg.InternStr("_Dialect")
		ß_StringIO := πg.InternStr("_StringIO")
		ß__all__ := πg.InternStr("__all__")
		ß__class__ := πg.InternStr("__class__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__init__ := πg.InternStr("__init__")
		ß__iter__ := πg.InternStr("__iter__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__version__ := πg.InternStr("__version__")
		ß_csv := πg.InternStr("_csv")
		ß_dict_to_list := πg.InternStr("_dict_to_list")
		ß_fieldnames := πg.InternStr("_fieldnames")
		ß_guess_delimiter := πg.InternStr("_guess_delimiter")
		ß_guess_quote_and_delimiter := πg.InternStr("_guess_quote_and_delimiter")
		ß_name := πg.InternStr("_name")
		ß_valid := πg.InternStr("_valid")
		ß_validate := πg.InternStr("_validate")
		ßappend := πg.InternStr("append")
		ßchr := πg.InternStr("chr")
		ßcompile := πg.InternStr("compile")
		ßcomplex := πg.InternStr("complex")
		ßcount := πg.InternStr("count")
		ßdelim := πg.InternStr("delim")
		ßdelimiter := πg.InternStr("delimiter")
		ßdialect := πg.InternStr("dialect")
		ßdict := πg.InternStr("dict")
		ßdoublequote := πg.InternStr("doublequote")
		ßescape := πg.InternStr("escape")
		ßescapechar := πg.InternStr("escapechar")
		ßexcel := πg.InternStr("excel")
		ßexcel_tab := πg.InternStr("excel_tab")
		ßextrasaction := πg.InternStr("extrasaction")
		ßfield_size_limit := πg.InternStr("field_size_limit")
		ßfieldnames := πg.InternStr("fieldnames")
		ßfilter := πg.InternStr("filter")
		ßfindall := πg.InternStr("findall")
		ßfloat := πg.InternStr("float")
		ßfunctools := πg.InternStr("functools")
		ßget := πg.InternStr("get")
		ßget_dialect := πg.InternStr("get_dialect")
		ßgetattr := πg.InternStr("getattr")
		ßglobals := πg.InternStr("globals")
		ßgroupindex := πg.InternStr("groupindex")
		ßhas_header := πg.InternStr("has_header")
		ßignore := πg.InternStr("ignore")
		ßint := πg.InternStr("int")
		ßitems := πg.InternStr("items")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlen := πg.InternStr("len")
		ßline_num := πg.InternStr("line_num")
		ßlineterminator := πg.InternStr("lineterminator")
		ßlist_dialects := πg.InternStr("list_dialects")
		ßlong := πg.InternStr("long")
		ßlower := πg.InternStr("lower")
		ßmin := πg.InternStr("min")
		ßname := πg.InternStr("name")
		ßnext := πg.InternStr("next")
		ßobject := πg.InternStr("object")
		ßpreferred := πg.InternStr("preferred")
		ßproperty := πg.InternStr("property")
		ßquote := πg.InternStr("quote")
		ßquotechar := πg.InternStr("quotechar")
		ßquoting := πg.InternStr("quoting")
		ßraise := πg.InternStr("raise")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßreader := πg.InternStr("reader")
		ßreduce := πg.InternStr("reduce")
		ßregister_dialect := πg.InternStr("register_dialect")
		ßremove := πg.InternStr("remove")
		ßrepr := πg.InternStr("repr")
		ßrestkey := πg.InternStr("restkey")
		ßrestval := πg.InternStr("restval")
		ßsearch := πg.InternStr("search")
		ßsetter := πg.InternStr("setter")
		ßskipinitialspace := πg.InternStr("skipinitialspace")
		ßsniff := πg.InternStr("sniff")
		ßsniffed := πg.InternStr("sniffed")
		ßsort := πg.InternStr("sort")
		ßspace := πg.InternStr("space")
		ßsplit := πg.InternStr("split")
		ßstr := πg.InternStr("str")
		ßtype := πg.InternStr("type")
		ßunregister_dialect := πg.InternStr("unregister_dialect")
		ßwriteheader := πg.InternStr("writeheader")
		ßwriter := πg.InternStr("writer")
		ßwriterow := πg.InternStr("writerow")
		ßwriterows := πg.InternStr("writerows")
		ßzip := πg.InternStr("zip")
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
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\ncsv.py - read/write/investigate CSV files\n").ToObject()); πE != nil {
				continue
			}
			// line 6: import re
			πF.SetLineno(6)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 7: import functools
			πF.SetLineno(7)
			if πTemp002, πE = πg.ImportModule(πF, "functools"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßfunctools.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: reduce = functools.reduce
			πF.SetLineno(8)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßfunctools); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßreduce, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßreduce.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 12: import _csv
			πF.SetLineno(12)
			if πTemp002, πE = πg.ImportModule(πF, "_csv"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_csv.ToObject(), πTemp001); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_csv); πE != nil {
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
			// line 14: globals()[name] = getattr(_csv, name)
			πF.SetLineno(14)
			πTemp002 = πF.MakeArgs(2)
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_csv); πE != nil {
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
			// line 22: _Dialect = _csv.Dialect
			πF.SetLineno(22)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_csv); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßDialect, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_Dialect.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 24: import StringIO as _StringIO
			πF.SetLineno(24)
			if πTemp002, πE = πg.ImportModule(πF, "StringIO"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ß_StringIO.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 25: StringIO = _StringIO.StringIO
			πF.SetLineno(25)
			if πTemp001, πE = πg.ResolveGlobal(πF, ß_StringIO); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßStringIO, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßStringIO.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 31: __all__ = [ "QUOTE_MINIMAL", "QUOTE_ALL", "QUOTE_NONNUMERIC", "QUOTE_NONE",
			πF.SetLineno(31)
			πTemp002 = make([]*πg.Object, 20)
			πTemp002[0] = ßQUOTE_MINIMAL.ToObject()
			πTemp002[1] = ßQUOTE_ALL.ToObject()
			πTemp002[2] = ßQUOTE_NONNUMERIC.ToObject()
			πTemp002[3] = ßQUOTE_NONE.ToObject()
			πTemp002[4] = ßError.ToObject()
			πTemp002[5] = ßDialect.ToObject()
			πTemp002[6] = ß__doc__.ToObject()
			πTemp002[7] = ßexcel.ToObject()
			πTemp002[8] = ßexcel_tab.ToObject()
			πTemp002[9] = ßfield_size_limit.ToObject()
			πTemp002[10] = ßreader.ToObject()
			πTemp002[11] = ßwriter.ToObject()
			πTemp002[12] = ßregister_dialect.ToObject()
			πTemp002[13] = ßget_dialect.ToObject()
			πTemp002[14] = ßlist_dialects.ToObject()
			πTemp002[15] = ßSniffer.ToObject()
			πTemp002[16] = ßunregister_dialect.ToObject()
			πTemp002[17] = ß__version__.ToObject()
			πTemp002[18] = ßDictReader.ToObject()
			πTemp002[19] = ßDictWriter.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 37: class Dialect(object):
			πF.SetLineno(37)
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
			_, πE = πg.NewCode("Dialect", "/usr/lib/python2.7/csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
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
					// line 38: """Describe an Excel dialect.
					πF.SetLineno(38)
					// line 38: """Describe an Excel dialect.
					πF.SetLineno(38)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Describe an Excel dialect.\n\n    This must be subclassed (see csv.excel).  Valid attributes are:\n    delimiter, quotechar, escapechar, doublequote, skipinitialspace,\n    lineterminator, quoting.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 45: _name = ""
					πF.SetLineno(45)
					if πE = πClass.SetItem(πF, ß_name.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					// line 46: _valid = False
					πF.SetLineno(46)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ß_valid.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 48: delimiter = None
					πF.SetLineno(48)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdelimiter.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 49: quotechar = None
					πF.SetLineno(49)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßquotechar.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 50: escapechar = None
					πF.SetLineno(50)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßescapechar.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 51: doublequote = None
					πF.SetLineno(51)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdoublequote.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 52: skipinitialspace = None
					πF.SetLineno(52)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßskipinitialspace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 53: lineterminator = None
					πF.SetLineno(53)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßlineterminator.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 54: quoting = None
					πF.SetLineno(54)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßquoting.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 56: def __init__(self):
					πF.SetLineno(56)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ß__class__, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßDialect); πE != nil {
								continue
							}
							if πTemp001, πE = πg.NE(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 57: if self.__class__ != Dialect:
							πF.SetLineno(57)
						Label1:
							// line 58: self._valid = True
							πF.SetLineno(58)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_valid, πTemp002); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 59: self._validate()
							πF.SetLineno(59)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_validate, nil); πE != nil {
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
					// line 61: def _validate(self):
					πF.SetLineno(61)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("_validate", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µe *πg.Object = πg.UnboundLocal; _ = µe
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 62: try:
							πF.SetLineno(62)
							πF.PushCheckpoint(2)
							// line 63: _Dialect(self)
							πF.SetLineno(63)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_Dialect); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
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
							πE = πF.Raise(πTemp004.ToObject(), nil, πTemp005.ToObject())
							continue
							// line 64: except TypeError, e:
							πF.SetLineno(64)
						Label3:
							µe = πTemp004.ToObject()
							πTemp001 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µe, "e"); πE != nil {
								continue
							}
							πTemp007[0] = µe
							if πTemp002, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							πTemp001[0] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 66: raise Error(str(e))
							πF.SetLineno(66)
							πE = πF.Raise(πTemp003, nil, nil)
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
					if πE = πClass.SetItem(πF, ß_validate.ToObject(), πTemp003); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Dialect").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDialect.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 68: class excel(Dialect):
			πF.SetLineno(68)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßDialect); πE != nil {
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
			_, πE = πg.NewCode("excel", "/usr/lib/python2.7/csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 69: """Describe the usual properties of Excel-generated CSV files."""
					πF.SetLineno(69)
					// line 69: """Describe the usual properties of Excel-generated CSV files."""
					πF.SetLineno(69)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Describe the usual properties of Excel-generated CSV files.").ToObject()); πE != nil {
						continue
					}
					// line 70: delimiter = ','
					πF.SetLineno(70)
					if πE = πClass.SetItem(πF, ßdelimiter.ToObject(), πg.NewStr(",").ToObject()); πE != nil {
						continue
					}
					// line 71: quotechar = '"'
					πF.SetLineno(71)
					if πE = πClass.SetItem(πF, ßquotechar.ToObject(), πg.NewStr("\"").ToObject()); πE != nil {
						continue
					}
					// line 72: doublequote = True
					πF.SetLineno(72)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßdoublequote.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 73: skipinitialspace = False
					πF.SetLineno(73)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßskipinitialspace.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 74: lineterminator = '\r\n'
					πF.SetLineno(74)
					if πE = πClass.SetItem(πF, ßlineterminator.ToObject(), πg.NewStr("\r\n").ToObject()); πE != nil {
						continue
					}
					// line 75: quoting = QUOTE_MINIMAL
					πF.SetLineno(75)
					if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßQUOTE_MINIMAL); πE != nil {
						continue
					}
					if πE = πClass.SetItem(πF, ßquoting.ToObject(), πTemp001); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("excel").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßexcel.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 76: register_dialect("excel", excel)
			πF.SetLineno(76)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = ßexcel.ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßexcel); πE != nil {
				continue
			}
			πTemp002[1] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßregister_dialect); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 78: class excel_tab(excel):
			πF.SetLineno(78)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßexcel); πE != nil {
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
			_, πE = πg.NewCode("excel_tab", "/usr/lib/python2.7/csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp010
				_ = πClass
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 79: """Describe the usual properties of Excel-generated TAB-delimited files."""
					πF.SetLineno(79)
					// line 79: """Describe the usual properties of Excel-generated TAB-delimited files."""
					πF.SetLineno(79)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Describe the usual properties of Excel-generated TAB-delimited files.").ToObject()); πE != nil {
						continue
					}
					// line 80: delimiter = '\t'
					πF.SetLineno(80)
					if πE = πClass.SetItem(πF, ßdelimiter.ToObject(), πg.NewStr("\t").ToObject()); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("excel_tab").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßexcel_tab.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 81: register_dialect("excel-tab", excel_tab)
			πF.SetLineno(81)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("excel-tab").ToObject()
			if πTemp001, πE = πg.ResolveGlobal(πF, ßexcel_tab); πE != nil {
				continue
			}
			πTemp002[1] = πTemp001
			if πTemp001, πE = πg.ResolveGlobal(πF, ßregister_dialect); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			// line 84: class DictReader(object):
			πF.SetLineno(84)
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
			_, πE = πg.NewCode("DictReader", "/usr/lib/python2.7/csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp005 []*πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 85: def __init__(self, f, fieldnames=None, restkey=None, restval=None,
					πF.SetLineno(85)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "f", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "fieldnames", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "restkey", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "restval", Def: πTemp003}
					// line 86: dialect="excel", *args, **kwds):
					πF.SetLineno(86)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), ßexcel.ToObject()); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "dialect", Def: ßexcel.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/csv.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πArgs[1]; _ = µf
						var µfieldnames *πg.Object = πArgs[2]; _ = µfieldnames
						var µrestkey *πg.Object = πArgs[3]; _ = µrestkey
						var µrestval *πg.Object = πArgs[4]; _ = µrestval
						var µdialect *πg.Object = πArgs[5]; _ = µdialect
						var µargs *πg.Object = πArgs[6]; _ = µargs
						var µkwds *πg.Object = πArgs[7]; _ = µkwds
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
							// line 87: self._fieldnames = fieldnames   # list of keys for the dict
							πF.SetLineno(87)
							if πE = πg.CheckLocal(πF, µfieldnames, "fieldnames"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfieldnames); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_fieldnames, πTemp001); πE != nil {
								continue
							}
							// line 88: self.restkey = restkey          # key to catch long rows
							πF.SetLineno(88)
							if πE = πg.CheckLocal(πF, µrestkey, "restkey"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µrestkey); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrestkey, πTemp001); πE != nil {
								continue
							}
							// line 89: self.restval = restval          # default value for short rows
							πF.SetLineno(89)
							if πE = πg.CheckLocal(πF, µrestval, "restval"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µrestval); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrestval, πTemp001); πE != nil {
								continue
							}
							// line 90: self.reader = reader(f, dialect, *args, **kwds)
							πF.SetLineno(90)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp002[0] = µf
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πTemp002[1] = µdialect
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßreader); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Invoke(πF, πTemp001, πTemp002, µargs, nil, µkwds); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßreader, πTemp001); πE != nil {
								continue
							}
							// line 91: self.dialect = dialect
							πF.SetLineno(91)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µdialect); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßdialect, πTemp001); πE != nil {
								continue
							}
							// line 92: self.line_num = 0
							πF.SetLineno(92)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_num, πTemp001); πE != nil {
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
					// line 94: def __iter__(self):
					πF.SetLineno(94)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("__iter__", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 95: return self
							πF.SetLineno(95)
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
					if πE = πClass.SetItem(πF, ß__iter__.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 98: def fieldnames(self):
					πF.SetLineno(98)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("fieldnames", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πTemp001 *πg.Object
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_fieldnames, nil); πE != nil {
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
							// line 99: if self._fieldnames is None:
							πF.SetLineno(99)
						Label1:
							// line 100: try:
							πF.SetLineno(100)
							πF.PushCheckpoint(4)
							// line 101: self._fieldnames = self.reader.next()
							πF.SetLineno(101)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnext, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_fieldnames, πTemp002); πE != nil {
								continue
							}
							πF.PopCheckpoint()
							goto Label3
						Label4:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label5
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 102: except StopIteration:
							πF.SetLineno(102)
						Label5:
							// line 103: pass
							πF.SetLineno(103)
							πF.RestoreExc(nil, nil)
							goto Label3
						Label3:
							goto Label2
						Label2:
							// line 104: self.line_num = self.reader.line_num
							πF.SetLineno(104)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßline_num, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_num, πTemp001); πE != nil {
								continue
							}
							// line 105: return self._fieldnames
							πF.SetLineno(105)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ß_fieldnames, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfieldnames.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 106: fieldnames = property(fieldnames)
					πF.SetLineno(106)
					πTemp005 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßfieldnames); πE != nil {
						continue
					}
					πTemp005[0] = πTemp006
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßfieldnames.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 113: def fieldnames(self, value):
					πF.SetLineno(113)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "value", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("fieldnames", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 114: self._fieldnames = value
							πF.SetLineno(114)
							if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µvalue); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ß_fieldnames, πTemp001); πE != nil {
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
					if πE = πClass.SetItem(πF, ßfieldnames.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 112: @fieldnames.setter
					πF.SetLineno(112)
					πTemp005 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßfieldnames); πE != nil {
						continue
					}
					πTemp005[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßfieldnames); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßsetter, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πE = πClass.SetItem(πF, ßfieldnames.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 116: def next(self):
					πF.SetLineno(116)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("next", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrow *πg.Object = πg.UnboundLocal; _ = µrow
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var µlf *πg.Object = πg.UnboundLocal; _ = µlf
						var µlr *πg.Object = πg.UnboundLocal; _ = µlr
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
							case 9: goto Label9
							case 10: goto Label10
							case 3: goto Label3
							case 4: goto Label4
							default: panic("unexpected function state")
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßline_num, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 117: if self.line_num == 0:
							πF.SetLineno(117)
						Label1:
							// line 119: self.fieldnames
							πF.SetLineno(119)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfieldnames, nil); πE != nil {
								continue
							}
							goto Label2
						Label2:
							// line 120: row = self.reader.next()
							πF.SetLineno(120)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnext, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrow = πTemp001
							// line 121: self.line_num = self.reader.line_num
							πF.SetLineno(121)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßline_num, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßline_num, πTemp001); πE != nil {
								continue
							}
							// line 126: while row == []:
							πF.SetLineno(126)
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
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp005 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp005...).ToObject()
							if πTemp001, πE = πg.Eq(πF, µrow, πTemp002); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(3)            
							// line 127: row = self.reader.next()
							πF.SetLineno(127)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßreader, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnext, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µrow = πTemp001
							continue
						Label4:
							if πE != nil || πR != nil {
								continue
							}
						Label5:
							// line 128: d = dict(zip(self.fieldnames, row))
							πF.SetLineno(128)
							πTemp005 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfieldnames, nil); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp006[1] = µrow
							if πTemp001, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp005[0] = πTemp002
							if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µd = πTemp002
							// line 129: lf = len(self.fieldnames)
							πF.SetLineno(129)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßfieldnames, nil); πE != nil {
								continue
							}
							πTemp005[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µlf = πTemp002
							// line 130: lr = len(row)
							πF.SetLineno(130)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp005[0] = µrow
							if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							µlr = πTemp002
							if πE = πg.CheckLocal(πF, µlf, "lf"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlr, "lr"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LT(πF, µlf, µlr); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label6
							}
							if πE = πg.CheckLocal(πF, µlf, "lf"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µlr, "lr"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GT(πF, µlf, µlr); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label7
							}
							goto Label8
							// line 131: if lf < lr:
							πF.SetLineno(131)
						Label6:
							// line 132: d[self.restkey] = row[lf:]
							πF.SetLineno(132)
							if πE = πg.CheckLocal(πF, µlf, "lf"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µlf, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, µrow, πTemp001); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßrestkey, nil); πE != nil {
								continue
							}
							πTemp007 = πTemp008
							if πE = πg.SetItem(πF, µd, πTemp007, πTemp001); πE != nil {
								continue
							}
							goto Label8
							// line 133: elif lf > lr:
							πF.SetLineno(133)
						Label7:
							if πE = πg.CheckLocal(πF, µlr, "lr"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µlr, πg.None, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µself, ßfieldnames, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp002); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Iter(πF, πTemp007); πE != nil {
								continue
							}
							πF.PushCheckpoint(10)
							πTemp003 = false
						Label9:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp003 {
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
								πTemp004 = !isStop
							} else {
								πTemp004 = true
								µkey = πTemp002
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(9)            
							// line 135: d[key] = self.restval
							πF.SetLineno(135)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßrestval, nil); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp008 = µkey
							if πE = πg.SetItem(πF, µd, πTemp008, πTemp007); πE != nil {
								continue
							}
							continue
						Label10:
							if πE != nil || πR != nil {
								continue
							}
						Label11:
							goto Label8
						Label8:
							// line 136: return d
							πF.SetLineno(136)
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
					if πE = πClass.SetItem(πF, ßnext.ToObject(), πTemp007); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DictReader").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDictReader.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 139: class DictWriter(object):
			πF.SetLineno(139)
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
			_, πE = πg.NewCode("DictWriter", "/usr/lib/python2.7/csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 140: def __init__(self, f, fieldnames, restval="", extrasaction="raise",
					πF.SetLineno(140)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "f", Def: nil}
					πTemp002[2] = πg.Param{Name: "fieldnames", Def: nil}
					// line 140: def __init__(self, f, fieldnames, restval="", extrasaction="raise",
					πF.SetLineno(140)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), ß.ToObject()); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "restval", Def: ß.ToObject()}
					πTemp002[4] = πg.Param{Name: "extrasaction", Def: ßraise.ToObject()}
					πTemp002[5] = πg.Param{Name: "dialect", Def: ßexcel.ToObject()}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/csv.py", πTemp002, πg.CodeFlagVarArg | πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µf *πg.Object = πArgs[1]; _ = µf
						var µfieldnames *πg.Object = πArgs[2]; _ = µfieldnames
						var µrestval *πg.Object = πArgs[3]; _ = µrestval
						var µextrasaction *πg.Object = πArgs[4]; _ = µextrasaction
						var µdialect *πg.Object = πArgs[5]; _ = µdialect
						var µargs *πg.Object = πArgs[6]; _ = µargs
						var µkwds *πg.Object = πArgs[7]; _ = µkwds
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
							// line 142: self.fieldnames = fieldnames    # list of keys for the dict
							πF.SetLineno(142)
							if πE = πg.CheckLocal(πF, µfieldnames, "fieldnames"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µfieldnames); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßfieldnames, πTemp001); πE != nil {
								continue
							}
							// line 143: self.restval = restval          # for writing short dicts
							πF.SetLineno(143)
							if πE = πg.CheckLocal(πF, µrestval, "restval"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µrestval); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßrestval, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µextrasaction, "extrasaction"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µextrasaction, ßlower, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(ßraise.ToObject(), ßignore.ToObject()).ToObject()
							if πTemp004, πE = πg.Contains(πF, πTemp002, πTemp003); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(!πTemp004).ToObject()
							if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 144: if extrasaction.lower() not in ("raise", "ignore"):
							πF.SetLineno(144)
						Label1:
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µextrasaction, "extrasaction"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("extrasaction (%s) must be 'raise' or 'ignore'").ToObject(), µextrasaction); πE != nil {
								continue
							}
							// line 145: raise ValueError, \
							πF.SetLineno(145)
							πE = πF.Raise(πTemp001, πTemp002, nil)
							continue
							goto Label2
						Label2:
							// line 148: self.extrasaction = extrasaction
							πF.SetLineno(148)
							if πE = πg.CheckLocal(πF, µextrasaction, "extrasaction"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µextrasaction); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßextrasaction, πTemp001); πE != nil {
								continue
							}
							// line 149: self.writer = writer(f, dialect, *args, **kwds)
							πF.SetLineno(149)
							πTemp005 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µf, "f"); πE != nil {
								continue
							}
							πTemp005[0] = µf
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πTemp005[1] = µdialect
							if πE = πg.CheckLocal(πF, µargs, "args"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ResolveGlobal(πF, ßwriter); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Invoke(πF, πTemp001, πTemp005, µargs, nil, µkwds); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßwriter, πTemp001); πE != nil {
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
					// line 151: def writeheader(self):
					πF.SetLineno(151)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp003 = πg.NewFunction(πg.NewCode("writeheader", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µheader *πg.Object = πg.UnboundLocal; _ = µheader
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
							// line 152: header = dict(zip(self.fieldnames, self.fieldnames))
							πF.SetLineno(152)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfieldnames, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßfieldnames, nil); πE != nil {
								continue
							}
							πTemp002[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µheader = πTemp004
							// line 153: self.writerow(header)
							πF.SetLineno(153)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							πTemp001[0] = µheader
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwriterow, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwriteheader.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 155: def _dict_to_list(self, rowdict):
					πF.SetLineno(155)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "rowdict", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_dict_to_list", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrowdict *πg.Object = πArgs[1]; _ = µrowdict
						var µwrong_fields *πg.Object = πg.UnboundLocal; _ = µwrong_fields
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 []πg.Param
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 []*πg.Object
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
							if πTemp002, πE = πg.GetAttr(πF, µself, ßextrasaction, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, πTemp002, ßraise.ToObject()); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label1
							}
							goto Label2
							// line 156: if self.extrasaction == "raise":
							πF.SetLineno(156)
						Label1:
							// line 157: wrong_fields = [k for k in rowdict if k not in self.fieldnames]
							πF.SetLineno(157)
							πTemp004 = make([]πg.Param, 0)
							πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/csv.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µk *πg.Object = πg.UnboundLocal; _ = µk
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
										if πE = πg.CheckLocal(πF, µrowdict, "rowdict"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µrowdict); πE != nil {
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
											µk = πTemp004
										}
										if πE != nil || !πTemp003 {
											continue
										}
										πF.PushCheckpoint(1)            
										if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp005, πE = πg.GetAttr(πF, µself, ßfieldnames, nil); πE != nil {
											continue
										}
										if πTemp003, πE = πg.Contains(πF, πTemp005, µk); πE != nil {
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
										// line 157: wrong_fields = [k for k in rowdict if k not in self.fieldnames]
										πF.SetLineno(157)
									Label4:
										// line 157: wrong_fields = [k for k in rowdict if k not in self.fieldnames]
										πF.SetLineno(157)
										if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
											continue
										}
										πF.PushCheckpoint(6)
										return µk, nil
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
							if πTemp005, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp005}, nil); πE != nil {
								continue
							}
							µwrong_fields = πTemp001
							if πE = πg.CheckLocal(πF, µwrong_fields, "wrong_fields"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IsTrue(πF, µwrong_fields); πE != nil {
								continue
							}
							if πTemp003 {
								goto Label3
							}
							goto Label4
							// line 158: if wrong_fields:
							πF.SetLineno(158)
						Label3:
							πTemp006 = πF.MakeArgs(1)
							πTemp007 = πF.MakeArgs(1)
							πTemp004 = make([]πg.Param, 0)
							πTemp008 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/csv.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
										if πE = πg.CheckLocal(πF, µwrong_fields, "wrong_fields"); πE != nil {
											continue
										}
										if πTemp001, πE = πg.Iter(πF, µwrong_fields); πE != nil {
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
										// line 160: + ", ".join([repr(x) for x in wrong_fields]))
										πF.SetLineno(160)
										πTemp005 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
											continue
										}
										πTemp005[0] = µx
										if πTemp004, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
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
							if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
								continue
							}
							πTemp007[0] = πTemp005
							if πTemp005, πE = πg.GetAttr(πF, πg.NewStr(", ").ToObject(), ßjoin, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp005.Call(πF, πTemp007, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp007)
							if πTemp001, πE = πg.Add(πF, πg.NewStr("dict contains fields not in fieldnames: ").ToObject(), πTemp009); πE != nil {
								continue
							}
							πTemp006[0] = πTemp001
							if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							// line 159: raise ValueError("dict contains fields not in fieldnames: "
							πF.SetLineno(159)
							πE = πF.Raise(πTemp005, nil, nil)
							continue
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 161: return [rowdict.get(key, self.restval) for key in self.fieldnames]
							πF.SetLineno(161)
							πTemp004 = make([]πg.Param, 0)
							πTemp005 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/csv.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µkey *πg.Object = πg.UnboundLocal; _ = µkey
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
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										default: panic("unexpected function state")
										}
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßfieldnames, nil); πE != nil {
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
											µkey = πTemp002
										}
										if πE != nil || !πTemp004 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 161: return [rowdict.get(key, self.restval) for key in self.fieldnames]
										πF.SetLineno(161)
										πTemp005 = πF.MakeArgs(2)
										if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
											continue
										}
										πTemp005[0] = µkey
										if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µself, ßrestval, nil); πE != nil {
											continue
										}
										πTemp005[1] = πTemp002
										if πE = πg.CheckLocal(πF, µrowdict, "rowdict"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µrowdict, ßget, nil); πE != nil {
											continue
										}
										if πTemp006, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp005)
										πF.PushCheckpoint(4)
										return πTemp006, nil
									Label4:
										πTemp002 = πSent
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
							if πTemp009, πE = πTemp005.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_dict_to_list.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 163: def writerow(self, rowdict):
					πF.SetLineno(163)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "rowdict", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("writerow", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrowdict *πg.Object = πArgs[1]; _ = µrowdict
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
							// line 164: return self.writer.writerow(self._dict_to_list(rowdict))
							πF.SetLineno(164)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrowdict, "rowdict"); πE != nil {
								continue
							}
							πTemp002[0] = µrowdict
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_dict_to_list, nil); πE != nil {
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
							if πTemp003, πE = πg.GetAttr(πF, µself, ßwriter, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwriterow, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwriterow.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 166: def writerows(self, rowdicts):
					πF.SetLineno(166)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "rowdicts", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("writerows", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µrowdicts *πg.Object = πArgs[1]; _ = µrowdicts
						var µrows *πg.Object = πg.UnboundLocal; _ = µrows
						var µrowdict *πg.Object = πg.UnboundLocal; _ = µrowdict
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
						var πTemp006 []*πg.Object
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
							// line 167: rows = []
							πF.SetLineno(167)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µrows = πTemp002
							if πE = πg.CheckLocal(πF, µrowdicts, "rowdicts"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µrowdicts); πE != nil {
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
								µrowdict = πTemp005
							}
							if πE != nil || !πTemp004 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 169: rows.append(self._dict_to_list(rowdict))
							πF.SetLineno(169)
							πTemp001 = πF.MakeArgs(1)
							πTemp006 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrowdict, "rowdict"); πE != nil {
								continue
							}
							πTemp006[0] = µrowdict
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µself, ß_dict_to_list, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							πTemp001[0] = πTemp007
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µrows, ßappend, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 170: return self.writer.writerows(rows)
							πF.SetLineno(170)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrows, "rows"); πE != nil {
								continue
							}
							πTemp001[0] = µrows
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßwriter, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp002, ßwriterows, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ßwriterows.ToObject(), πTemp006); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("DictWriter").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßDictWriter.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 177: complex = float
			πF.SetLineno(177)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßcomplex.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 179: class Sniffer(object):
			πF.SetLineno(179)
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
			_, πE = πg.NewCode("Sniffer", "/usr/lib/python2.7/csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 180: '''
					πF.SetLineno(180)
					// line 180: '''
					πF.SetLineno(180)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("\n    \"Sniffs\" the format of a CSV file (i.e. delimiter, quotechar)\n    Returns a Dialect object.\n    ").ToObject()); πE != nil {
						continue
					}
					// line 184: def __init__(self):
					πF.SetLineno(184)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 186: self.preferred = [',', '\t', ';', ' ', ':']
							πF.SetLineno(186)
							πTemp001 = make([]*πg.Object, 5)
							πTemp001[0] = πg.NewStr(",").ToObject()
							πTemp001[1] = πg.NewStr("\t").ToObject()
							πTemp001[2] = πg.NewStr(";").ToObject()
							πTemp001[3] = πg.NewStr(" ").ToObject()
							πTemp001[4] = πg.NewStr(":").ToObject()
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßpreferred, πTemp003); πE != nil {
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
					// line 189: def sniff(self, sample, delimiters=None):
					πF.SetLineno(189)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "sample", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "delimiters", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("sniff", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsample *πg.Object = πArgs[1]; _ = µsample
						var µdelimiters *πg.Object = πArgs[2]; _ = µdelimiters
						var µquotechar *πg.Object = πg.UnboundLocal; _ = µquotechar
						var µdoublequote *πg.Object = πg.UnboundLocal; _ = µdoublequote
						var µdelimiter *πg.Object = πg.UnboundLocal; _ = µdelimiter
						var µskipinitialspace *πg.Object = πg.UnboundLocal; _ = µskipinitialspace
						var µdialect *πg.Object = πg.UnboundLocal; _ = µdialect
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
						var πTemp008 *πg.Dict
						_ = πTemp008
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 190: """
							πF.SetLineno(190)
							// line 194: quotechar, doublequote, delimiter, skipinitialspace = \
							πF.SetLineno(194)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsample, "sample"); πE != nil {
								continue
							}
							πTemp001[0] = µsample
							if πE = πg.CheckLocal(πF, µdelimiters, "delimiters"); πE != nil {
								continue
							}
							πTemp001[1] = µdelimiters
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_guess_quote_and_delimiter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp003); πE != nil {
								continue
							}
							µquotechar = πTemp002
							µdoublequote = πTemp004
							µdelimiter = πTemp005
							µskipinitialspace = πTemp006
							if πE = πg.CheckLocal(πF, µdelimiter, "delimiter"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µdelimiter); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 196: if not delimiter:
							πF.SetLineno(196)
						Label1:
							// line 197: delimiter, skipinitialspace = self._guess_delimiter(sample,
							πF.SetLineno(197)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µsample, "sample"); πE != nil {
								continue
							}
							πTemp001[0] = µsample
							if πE = πg.CheckLocal(πF, µdelimiters, "delimiters"); πE != nil {
								continue
							}
							πTemp001[1] = µdelimiters
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ß_guess_delimiter, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µdelimiter = πTemp002
							µskipinitialspace = πTemp004
							goto Label2
						Label2:
							if πE = πg.CheckLocal(πF, µdelimiter, "delimiter"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, µdelimiter); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp007).ToObject()
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							goto Label4
							// line 200: if not delimiter:
							πF.SetLineno(200)
						Label3:
							if πTemp002, πE = πg.ResolveGlobal(πF, ßError); πE != nil {
								continue
							}
							// line 201: raise Error, "Could not determine delimiter"
							πF.SetLineno(201)
							πE = πF.Raise(πTemp002, πg.NewStr("Could not determine delimiter").ToObject(), nil)
							continue
							goto Label4
						Label4:
							// line 203: class dialect(Dialect):
							πF.SetLineno(203)
							πTemp001 = make([]*πg.Object, 1)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßDialect); πE != nil {
								continue
							}
							πTemp001[0] = πTemp004
							πTemp008 = πg.NewDict()
							if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
								continue
							}
							if πE = πTemp008.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
								continue
							}
							_, πE = πg.NewCode("dialect", "/usr/lib/python2.7/csv.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
								πClass := πTemp008
								_ = πClass
								var πTemp001 *πg.Object
								_ = πTemp001
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 204: _name = "sniffed"
									πF.SetLineno(204)
									// line 204: _name = "sniffed"
									πF.SetLineno(204)
									if πE = πClass.SetItem(πF, ß__doc__.ToObject(), ßsniffed.ToObject()); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ß_name.ToObject(), ßsniffed.ToObject()); πE != nil {
										continue
									}
									// line 205: lineterminator = '\r\n'
									πF.SetLineno(205)
									if πE = πClass.SetItem(πF, ßlineterminator.ToObject(), πg.NewStr("\r\n").ToObject()); πE != nil {
										continue
									}
									// line 206: quoting = QUOTE_MINIMAL
									πF.SetLineno(206)
									if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßQUOTE_MINIMAL); πE != nil {
										continue
									}
									if πE = πClass.SetItem(πF, ßquoting.ToObject(), πTemp001); πE != nil {
										continue
									}
								}
								return nil, nil
							}).Eval(πF, πF.Globals(), nil, nil)
							if πE != nil {
								continue
							}
							if πTemp003, πE = πTemp008.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
								continue
							}
							if πTemp003 == nil {
								πTemp003 = πg.TypeType.ToObject()
							}
							if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("dialect").ToObject(), πg.NewTuple(πTemp001...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
								continue
							}
							µdialect = πTemp004
							// line 209: dialect.doublequote = doublequote
							πF.SetLineno(209)
							if πE = πg.CheckLocal(πF, µdoublequote, "doublequote"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdoublequote); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µdialect, ßdoublequote, πTemp002); πE != nil {
								continue
							}
							// line 210: dialect.delimiter = delimiter
							πF.SetLineno(210)
							if πE = πg.CheckLocal(πF, µdelimiter, "delimiter"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µdelimiter); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µdialect, ßdelimiter, πTemp002); πE != nil {
								continue
							}
							// line 212: dialect.quotechar = quotechar or '"'
							πF.SetLineno(212)
							if πE = πg.CheckLocal(πF, µquotechar, "quotechar"); πE != nil {
								continue
							}
							πTemp002 = µquotechar
							if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label5
							}
							πTemp002 = πg.NewStr("\"").ToObject()
						Label5:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µdialect, ßquotechar, πTemp003); πE != nil {
								continue
							}
							// line 213: dialect.skipinitialspace = skipinitialspace
							πF.SetLineno(213)
							if πE = πg.CheckLocal(πF, µskipinitialspace, "skipinitialspace"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µskipinitialspace); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µdialect, ßskipinitialspace, πTemp002); πE != nil {
								continue
							}
							// line 215: return dialect
							πF.SetLineno(215)
							if πE = πg.CheckLocal(πF, µdialect, "dialect"); πE != nil {
								continue
							}
							πR = µdialect
							continue
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßsniff.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 190: """
					πF.SetLineno(190)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("\n        Returns a dialect (or None) corresponding to the sample\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßsniff); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 218: def _guess_quote_and_delimiter(self, data, delimiters):
					πF.SetLineno(218)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp002[2] = πg.Param{Name: "delimiters", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("_guess_quote_and_delimiter", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdata *πg.Object = πArgs[1]; _ = µdata
						var µdelimiters *πg.Object = πArgs[2]; _ = µdelimiters
						var µmatches *πg.Object = πg.UnboundLocal; _ = µmatches
						var µrestr *πg.Object = πg.UnboundLocal; _ = µrestr
						var µregexp *πg.Object = πg.UnboundLocal; _ = µregexp
						var µquotes *πg.Object = πg.UnboundLocal; _ = µquotes
						var µdelims *πg.Object = πg.UnboundLocal; _ = µdelims
						var µspaces *πg.Object = πg.UnboundLocal; _ = µspaces
						var µm *πg.Object = πg.UnboundLocal; _ = µm
						var µn *πg.Object = πg.UnboundLocal; _ = µn
						var µkey *πg.Object = πg.UnboundLocal; _ = µkey
						var µquotechar *πg.Object = πg.UnboundLocal; _ = µquotechar
						var µdelim *πg.Object = πg.UnboundLocal; _ = µdelim
						var µskipinitialspace *πg.Object = πg.UnboundLocal; _ = µskipinitialspace
						var µdq_regexp *πg.Object = πg.UnboundLocal; _ = µdq_regexp
						var µdoublequote *πg.Object = πg.UnboundLocal; _ = µdoublequote
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
						var πTemp009 *πg.Dict
						_ = πTemp009
						var πTemp010 *πg.BaseException
						_ = πTemp010
						var πTemp011 *πg.Traceback
						_ = πTemp011
						var πTemp012 bool
						_ = πTemp012
						var πTemp013 bool
						_ = πTemp013
						var πTemp014 []πg.Param
						_ = πTemp014
						var πTemp015 []*πg.Object
						_ = πTemp015
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 1: goto Label1
							case 2: goto Label2
							case 8: goto Label8
							case 9: goto Label9
							case 14: goto Label14
							case 21: goto Label21
							default: panic("unexpected function state")
							}
							// line 219: """
							πF.SetLineno(219)
							// line 230: matches = []
							πF.SetLineno(230)
							πTemp001 = make([]*πg.Object, 0)
							πTemp002 = πg.NewList(πTemp001...).ToObject()
							µmatches = πTemp002
							πTemp003 = πg.NewTuple4(πg.NewStr("(?P<delim>[^\\w\n\"'])(?P<space> ?)(?P<quote>[\"']).*?(?P=quote)(?P=delim)").ToObject(), πg.NewStr("(?:^|\n)(?P<quote>[\"']).*?(?P=quote)(?P<delim>[^\\w\n\"'])(?P<space> ?)").ToObject(), πg.NewStr("(?P<delim>>[^\\w\n\"'])(?P<space> ?)(?P<quote>[\"']).*?(?P=quote)(?:$|\n)").ToObject(), πg.NewStr("(?:^|\n)(?P<quote>[\"']).*?(?P=quote)(?:$|\n)").ToObject()).ToObject()
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
								µrestr = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 235: regexp = re.compile(restr, re.DOTALL | re.MULTILINE)
							πF.SetLineno(235)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µrestr, "restr"); πE != nil {
								continue
							}
							πTemp001[0] = µrestr
							if πTemp006, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßDOTALL, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßMULTILINE, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Or(πF, πTemp007, πTemp008); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßcompile, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µregexp = πTemp003
							// line 236: matches = regexp.findall(data)
							πF.SetLineno(236)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πE = πg.CheckLocal(πF, µregexp, "regexp"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µregexp, ßfindall, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmatches = πTemp006
							if πE = πg.CheckLocal(πF, µmatches, "matches"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µmatches); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label4
							}
							goto Label5
							// line 237: if matches:
							πF.SetLineno(237)
						Label4:
							// line 238: break
							πF.SetLineno(238)
							πTemp004 = true
							continue
							goto Label5
						Label5:
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µmatches, "matches"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µmatches); πE != nil {
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
							// line 240: if not matches:
							πF.SetLineno(240)
						Label6:
							// line 242: return ('', False, None, 0)
							πF.SetLineno(242)
							if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(ß.ToObject(), πTemp003, πTemp006, πg.NewInt(0).ToObject()).ToObject()
							πR = πTemp002
							continue
							goto Label7
						Label7:
							// line 243: quotes = {}
							πF.SetLineno(243)
							πTemp009 = πg.NewDict()
							πTemp002 = πTemp009.ToObject()
							µquotes = πTemp002
							// line 244: delims = {}
							πF.SetLineno(244)
							πTemp009 = πg.NewDict()
							πTemp002 = πTemp009.ToObject()
							µdelims = πTemp002
							// line 245: spaces = 0
							πF.SetLineno(245)
							µspaces = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µmatches, "matches"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µmatches); πE != nil {
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
								µm = πTemp003
							}
							if πE != nil || !πTemp005 {
								continue
							}
							πF.PushCheckpoint(8)            
							// line 247: n = regexp.groupindex['quote'] - 1
							πF.SetLineno(247)
							πTemp006 = ßquote.ToObject()
							if πE = πg.CheckLocal(πF, µregexp, "regexp"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µregexp, ßgroupindex, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µn = πTemp003
							// line 248: key = m[n]
							πF.SetLineno(248)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp003 = µn
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µm, πTemp003); πE != nil {
								continue
							}
							µkey = πTemp006
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µkey); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label11
							}
							goto Label12
							// line 249: if key:
							πF.SetLineno(249)
						Label11:
							// line 250: quotes[key] = quotes.get(key, 0) + 1
							πF.SetLineno(250)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µquotes, ßget, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Add(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007 = µkey
							if πE = πg.SetItem(πF, µquotes, πTemp007, πTemp006); πE != nil {
								continue
							}
							goto Label12
						Label12:
							// line 251: try:
							πF.SetLineno(251)
							πF.PushCheckpoint(14)
							// line 252: n = regexp.groupindex['delim'] - 1
							πF.SetLineno(252)
							πTemp006 = ßdelim.ToObject()
							if πE = πg.CheckLocal(πF, µregexp, "regexp"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µregexp, ßgroupindex, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µn = πTemp003
							// line 253: key = m[n]
							πF.SetLineno(253)
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp003 = µn
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µm, πTemp003); πE != nil {
								continue
							}
							µkey = πTemp006
							πF.PopCheckpoint()
							goto Label13
						Label14:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label15
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 254: except KeyError:
							πF.SetLineno(254)
						Label15:
							// line 255: continue
							πF.SetLineno(255)
							continue
							πF.RestoreExc(nil, nil)
							goto Label13
						Label13:
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp003 = µkey
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if !πTemp005 {
								goto Label16
							}
							if πE = πg.CheckLocal(πF, µdelimiters, "delimiters"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(µdelimiters == πTemp008).ToObject()
							πTemp006 = πTemp007
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label17
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdelimiters, "delimiters"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.Contains(πF, µdelimiters, µkey); πE != nil {
								continue
							}
							πTemp007 = πg.GetBool(πTemp013).ToObject()
							πTemp006 = πTemp007
						Label17:
							πTemp003 = πTemp006
						Label16:
							if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label18
							}
							goto Label19
							// line 256: if key and (delimiters is None or key in delimiters):
							πF.SetLineno(256)
						Label18:
							// line 257: delims[key] = delims.get(key, 0) + 1
							πF.SetLineno(257)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp001[0] = µkey
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µdelims, ßget, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Add(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
								continue
							}
							πTemp007 = µkey
							if πE = πg.SetItem(πF, µdelims, πTemp007, πTemp006); πE != nil {
								continue
							}
							goto Label19
						Label19:
							// line 258: try:
							πF.SetLineno(258)
							πF.PushCheckpoint(21)
							// line 259: n = regexp.groupindex['space'] - 1
							πF.SetLineno(259)
							πTemp006 = ßspace.ToObject()
							if πE = πg.CheckLocal(πF, µregexp, "regexp"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µregexp, ßgroupindex, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetItem(πF, πTemp008, πTemp006); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Sub(πF, πTemp007, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µn = πTemp003
							πF.PopCheckpoint()
							goto Label20
						Label21:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp010, πTemp011 = πF.ExcInfo()
							if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsInstance(πF, πTemp010.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label22
							}
							πE = πF.Raise(πTemp010.ToObject(), nil, πTemp011.ToObject())
							continue
							// line 260: except KeyError:
							πF.SetLineno(260)
						Label22:
							// line 261: continue
							πF.SetLineno(261)
							continue
							πF.RestoreExc(nil, nil)
							goto Label20
						Label20:
							if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
								continue
							}
							πTemp003 = µn
							if πE = πg.CheckLocal(πF, µm, "m"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µm, πTemp003); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label23
							}
							goto Label24
							// line 262: if m[n]:
							πF.SetLineno(262)
						Label23:
							// line 263: spaces += 1
							πF.SetLineno(263)
							if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.IAdd(πF, µspaces, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µspaces = πTemp003
							goto Label24
						Label24:
							continue
						Label9:
							if πE != nil || πR != nil {
								continue
							}
						Label10:
							// line 265: quotechar = reduce(lambda a, b, quotes = quotes:
							πF.SetLineno(265)
							πTemp001 = πF.MakeArgs(2)
							πTemp014 = make([]πg.Param, 3)
							πTemp014[0] = πg.Param{Name: "a", Def: nil}
							πTemp014[1] = πg.Param{Name: "b", Def: nil}
							if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
								continue
							}
							πTemp014[2] = πg.Param{Name: "quotes", Def: µquotes}
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/csv.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µa *πg.Object = πArgs[0]; _ = µa
								var µb *πg.Object = πArgs[1]; _ = µb
								var µquotes *πg.Object = πArgs[2]; _ = µquotes
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
								var πTemp008 *πg.Object
								_ = πTemp008
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 265: quotechar = reduce(lambda a, b, quotes = quotes:
									πF.SetLineno(265)
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									πTemp006 = µa
									if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
										continue
									}
									if πTemp007, πE = πg.GetItem(πF, µquotes, πTemp006); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									πTemp006 = µb
									if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
										continue
									}
									if πTemp008, πE = πg.GetItem(πF, µquotes, πTemp006); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GT(πF, πTemp007, πTemp008); πE != nil {
										continue
									}
									πTemp003 = πTemp005
									if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									if !πTemp004 {
										goto Label2
									}
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									πTemp003 = µa
								Label2:
									πTemp001 = πTemp003
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label1
									}
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									πTemp001 = µb
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
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µquotes, "quotes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µquotes, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßreduce); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µquotechar = πTemp003
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, µdelims); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label25
							}
							goto Label26
							// line 268: if delims:
							πF.SetLineno(268)
						Label25:
							// line 269: delim = reduce(lambda a, b, delims = delims:
							πF.SetLineno(269)
							πTemp001 = πF.MakeArgs(2)
							πTemp014 = make([]πg.Param, 3)
							πTemp014[0] = πg.Param{Name: "a", Def: nil}
							πTemp014[1] = πg.Param{Name: "b", Def: nil}
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							πTemp014[2] = πg.Param{Name: "delims", Def: µdelims}
							πTemp002 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/csv.py", πTemp014, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µa *πg.Object = πArgs[0]; _ = µa
								var µb *πg.Object = πArgs[1]; _ = µb
								var µdelims *πg.Object = πArgs[2]; _ = µdelims
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
								var πTemp008 *πg.Object
								_ = πTemp008
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 269: delim = reduce(lambda a, b, delims = delims:
									πF.SetLineno(269)
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									πTemp006 = µa
									if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
										continue
									}
									if πTemp007, πE = πg.GetItem(πF, µdelims, πTemp006); πE != nil {
										continue
									}
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									πTemp006 = µb
									if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
										continue
									}
									if πTemp008, πE = πg.GetItem(πF, µdelims, πTemp006); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GT(πF, πTemp007, πTemp008); πE != nil {
										continue
									}
									πTemp003 = πTemp005
									if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									if !πTemp004 {
										goto Label2
									}
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									πTemp003 = µa
								Label2:
									πTemp001 = πTemp003
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label1
									}
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									πTemp001 = µb
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
							πTemp001[0] = πTemp002
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdelims, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
							if πTemp002, πE = πg.ResolveGlobal(πF, ßreduce); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdelim = πTemp003
							// line 271: skipinitialspace = delims[delim] == spaces
							πF.SetLineno(271)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							πTemp003 = µdelim
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µdelims, πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µspaces, "spaces"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, πTemp006, µspaces); πE != nil {
								continue
							}
							µskipinitialspace = πTemp002
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Eq(πF, µdelim, πg.NewStr("\n").ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label28
							}
							goto Label29
							// line 272: if delim == '\n': # most likely a file with a single column
							πF.SetLineno(272)
						Label28:
							// line 273: delim = ''
							πF.SetLineno(273)
							µdelim = ß.ToObject()
							goto Label29
						Label29:
							goto Label27
						Label26:
							// line 276: delim = ''
							πF.SetLineno(276)
							µdelim = ß.ToObject()
							// line 277: skipinitialspace = 0
							πF.SetLineno(277)
							µskipinitialspace = πg.NewInt(0).ToObject()
							goto Label27
						Label27:
							// line 281: dq_regexp = re.compile(
							πF.SetLineno(281)
							πTemp001 = πF.MakeArgs(2)
							πTemp009 = πg.NewDict()
							πTemp015 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							πTemp015[0] = µdelim
							if πTemp003, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp003, ßescape, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp006.Call(πF, πTemp015, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp015)
							if πE = πTemp009.SetItem(πF, ßdelim.ToObject(), πTemp003); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µquotechar, "quotechar"); πE != nil {
								continue
							}
							if πE = πTemp009.SetItem(πF, ßquote.ToObject(), µquotechar); πE != nil {
								continue
							}
							πTemp003 = πTemp009.ToObject()
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("((%(delim)s)|^)\\W*%(quote)s[^%(delim)s\\n]*%(quote)s[^%(delim)s\\n]*%(quote)s\\W*((%(delim)s)|$)").ToObject(), πTemp003); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßMULTILINE, nil); πE != nil {
								continue
							}
							πTemp001[1] = πTemp003
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
							µdq_regexp = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πE = πg.CheckLocal(πF, µdq_regexp, "dq_regexp"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdq_regexp, ßsearch, nil); πE != nil {
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
								goto Label30
							}
							goto Label31
							// line 287: if dq_regexp.search(data):
							πF.SetLineno(287)
						Label30:
							// line 288: doublequote = True
							πF.SetLineno(288)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
								continue
							}
							µdoublequote = πTemp002
							goto Label32
						Label31:
							// line 290: doublequote = False
							πF.SetLineno(290)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
								continue
							}
							µdoublequote = πTemp002
							goto Label32
						Label32:
							// line 292: return (quotechar, doublequote, delim, skipinitialspace)
							πF.SetLineno(292)
							if πE = πg.CheckLocal(πF, µquotechar, "quotechar"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdoublequote, "doublequote"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µskipinitialspace, "skipinitialspace"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple4(µquotechar, µdoublequote, µdelim, µskipinitialspace).ToObject()
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
					if πE = πClass.SetItem(πF, ß_guess_quote_and_delimiter.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 219: """
					πF.SetLineno(219)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("\n        Looks for text enclosed between two identical quotes\n        (the probable quotechar) which are preceded and followed\n        by the same character (the probable delimiter).\n        For example:\n                         ,'some text',\n        The quote with the most wins, same with the delimiter.\n        If there is no quotechar the delimiter can't be determined\n        this way.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ß_guess_quote_and_delimiter); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 295: def _guess_delimiter(self, data, delimiters):
					πF.SetLineno(295)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "data", Def: nil}
					πTemp002[2] = πg.Param{Name: "delimiters", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_guess_delimiter", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µdata *πg.Object = πArgs[1]; _ = µdata
						var µdelimiters *πg.Object = πArgs[2]; _ = µdelimiters
						var µascii *πg.Object = πg.UnboundLocal; _ = µascii
						var µchunkLength *πg.Object = πg.UnboundLocal; _ = µchunkLength
						var µiteration *πg.Object = πg.UnboundLocal; _ = µiteration
						var µcharFrequency *πg.Object = πg.UnboundLocal; _ = µcharFrequency
						var µmodes *πg.Object = πg.UnboundLocal; _ = µmodes
						var µdelims *πg.Object = πg.UnboundLocal; _ = µdelims
						var µstart *πg.Object = πg.UnboundLocal; _ = µstart
						var µend *πg.Object = πg.UnboundLocal; _ = µend
						var µline *πg.Object = πg.UnboundLocal; _ = µline
						var µchar *πg.Object = πg.UnboundLocal; _ = µchar
						var µmetaFrequency *πg.Object = πg.UnboundLocal; _ = µmetaFrequency
						var µfreq *πg.Object = πg.UnboundLocal; _ = µfreq
						var µitems *πg.Object = πg.UnboundLocal; _ = µitems
						var µmodeList *πg.Object = πg.UnboundLocal; _ = µmodeList
						var µtotal *πg.Object = πg.UnboundLocal; _ = µtotal
						var µconsistency *πg.Object = πg.UnboundLocal; _ = µconsistency
						var µthreshold *πg.Object = πg.UnboundLocal; _ = µthreshold
						var µk *πg.Object = πg.UnboundLocal; _ = µk
						var µv *πg.Object = πg.UnboundLocal; _ = µv
						var µdelim *πg.Object = πg.UnboundLocal; _ = µdelim
						var µskipinitialspace *πg.Object = πg.UnboundLocal; _ = µskipinitialspace
						var µd *πg.Object = πg.UnboundLocal; _ = µd
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 []*πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 []πg.Param
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 *πg.Dict
						_ = πTemp007
						var πTemp008 *πg.Object
						_ = πTemp008
						var πTemp009 bool
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 bool
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
						var πTemp019 bool
						_ = πTemp019
						var πTemp020 bool
						_ = πTemp020
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
							case 10: goto Label10
							case 11: goto Label11
							case 40: goto Label40
							case 19: goto Label19
							case 20: goto Label20
							case 23: goto Label23
							case 24: goto Label24
							case 39: goto Label39
							default: panic("unexpected function state")
							}
							// line 296: """
							πF.SetLineno(296)
							// line 314: data = filter(None, data.split('\n'))
							πF.SetLineno(314)
							πTemp001 = πF.MakeArgs(2)
							if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("\n").ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µdata, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[1] = πTemp004
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfilter); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µdata = πTemp004
							// line 316: ascii = [chr(c) for c in range(127)] # 7-bit ASCII
							πF.SetLineno(316)
							πTemp005 = make([]πg.Param, 0)
							πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/csv.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µc *πg.Object = πg.UnboundLocal; _ = µc
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
								return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
									for ; πF.State() >= 0; πF.PopCheckpoint() {
										switch πF.State() {
										case 0:
										case 1: goto Label1
										case 2: goto Label2
										case 4: goto Label4
										default: panic("unexpected function state")
										}
										πTemp002 = πF.MakeArgs(1)
										πTemp002[0] = πg.NewInt(127).ToObject()
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
											µc = πTemp003
										}
										if πE != nil || !πTemp006 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 316: ascii = [chr(c) for c in range(127)] # 7-bit ASCII
										πF.SetLineno(316)
										πTemp002 = πF.MakeArgs(1)
										if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
											continue
										}
										πTemp002[0] = µc
										if πTemp003, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
											continue
										}
										if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
											continue
										}
										πF.FreeArgs(πTemp002)
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
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp006}, nil); πE != nil {
								continue
							}
							µascii = πTemp002
							// line 319: chunkLength = min(10, len(data))
							πF.SetLineno(319)
							πTemp001 = πF.MakeArgs(2)
							πTemp001[0] = πg.NewInt(10).ToObject()
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003[0] = µdata
							if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[1] = πTemp006
							if πTemp002, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µchunkLength = πTemp006
							// line 320: iteration = 0
							πF.SetLineno(320)
							µiteration = πg.NewInt(0).ToObject()
							// line 321: charFrequency = {}
							πF.SetLineno(321)
							πTemp007 = πg.NewDict()
							πTemp002 = πTemp007.ToObject()
							µcharFrequency = πTemp002
							// line 322: modes = {}
							πF.SetLineno(322)
							πTemp007 = πg.NewDict()
							πTemp002 = πTemp007.ToObject()
							µmodes = πTemp002
							// line 323: delims = {}
							πF.SetLineno(323)
							πTemp007 = πg.NewDict()
							πTemp002 = πTemp007.ToObject()
							µdelims = πTemp002
							// line 324: start, end = 0, min(chunkLength, len(data))
							πF.SetLineno(324)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchunkLength, "chunkLength"); πE != nil {
								continue
							}
							πTemp001[0] = µchunkLength
							πTemp003 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp003[0] = µdata
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							πTemp001[1] = πTemp008
							if πTemp006, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp002 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp008).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
								continue
							}
							µstart = πTemp006
							µend = πTemp008
							// line 325: while start < len(data):
							πF.SetLineno(325)
							πF.PushCheckpoint(2)
							πTemp009 = false
						Label1:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							πTemp001[0] = µdata
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.LT(πF, µstart, πTemp008); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 326: iteration += 1
							πF.SetLineno(326)
							if πE = πg.CheckLocal(πF, µiteration, "iteration"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µiteration, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µiteration = πTemp002
							if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{µstart, µend, πg.None}, nil); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µdata, πTemp006); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(5)
							πTemp010 = false
						Label4:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp010 {
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
								πTemp011 = !isStop
							} else {
								πTemp011 = true
								µline = πTemp006
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(4)            
							if πE = πg.CheckLocal(πF, µascii, "ascii"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Iter(πF, µascii); πE != nil {
								continue
							}
							πF.PushCheckpoint(8)
							πTemp011 = false
						Label7:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp011 {
								πF.PopCheckpoint()
								goto Label9
							}
							if πTemp008, πE = πg.Next(πF, πTemp006); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp012 = !isStop
							} else {
								πTemp012 = true
								µchar = πTemp008
							}
							if πE != nil || !πTemp012 {
								continue
							}
							πF.PushCheckpoint(7)            
							// line 329: metaFrequency = charFrequency.get(char, {})
							πF.SetLineno(329)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp001[0] = µchar
							πTemp007 = πg.NewDict()
							πTemp008 = πTemp007.ToObject()
							πTemp001[1] = πTemp008
							if πE = πg.CheckLocal(πF, µcharFrequency, "charFrequency"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µcharFrequency, ßget, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µmetaFrequency = πTemp013
							// line 331: freq = line.count(char)
							πF.SetLineno(331)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp001[0] = µchar
							if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µline, ßcount, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µfreq = πTemp013
							// line 333: metaFrequency[freq] = metaFrequency.get(freq, 0) + 1
							πF.SetLineno(333)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µfreq, "freq"); πE != nil {
								continue
							}
							πTemp001[0] = µfreq
							πTemp001[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µmetaFrequency, "metaFrequency"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetAttr(πF, µmetaFrequency, ßget, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp013.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp008, πE = πg.Add(πF, πTemp014, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmetaFrequency, "metaFrequency"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfreq, "freq"); πE != nil {
								continue
							}
							πTemp014 = µfreq
							if πE = πg.SetItem(πF, µmetaFrequency, πTemp014, πTemp013); πE != nil {
								continue
							}
							// line 334: charFrequency[char] = metaFrequency
							πF.SetLineno(334)
							if πE = πg.CheckLocal(πF, µmetaFrequency, "metaFrequency"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, µmetaFrequency); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcharFrequency, "charFrequency"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp013 = µchar
							if πE = πg.SetItem(πF, µcharFrequency, πTemp013, πTemp008); πE != nil {
								continue
							}
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
							if πE = πg.CheckLocal(πF, µcharFrequency, "charFrequency"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µcharFrequency, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp008); πE != nil {
								continue
							}
							πF.PushCheckpoint(11)
							πTemp010 = false
						Label10:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp010 {
								πF.PopCheckpoint()
								goto Label12
							}
							if πTemp006, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µchar = πTemp006
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(10)            
							// line 337: items = charFrequency[char].items()
							πF.SetLineno(337)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp006 = µchar
							if πE = πg.CheckLocal(πF, µcharFrequency, "charFrequency"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µcharFrequency, πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp008, ßitems, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							µitems = πTemp008
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp001[0] = µitems
							if πTemp013, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp013.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp008, πE = πg.Eq(πF, πTemp014, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp008
							if πTemp011, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp011 {
								goto Label13
							}
							πTemp013 = πg.NewInt(0).ToObject()
							πTemp015 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp016, πE = πg.GetItem(πF, µitems, πTemp015); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, πTemp016, πTemp013); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Eq(πF, πTemp014, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp008
						Label13:
							if πTemp011, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label14
							}
							goto Label15
							// line 338: if len(items) == 1 and items[0][0] == 0:
							πF.SetLineno(338)
						Label14:
							// line 339: continue
							πF.SetLineno(339)
							continue
							goto Label15
						Label15:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp001[0] = µitems
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.GT(πF, πTemp013, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp011, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label16
							}
							goto Label17
							// line 341: if len(items) > 1:
							πF.SetLineno(341)
						Label16:
							// line 342: modes[char] = reduce(lambda a, b: a[1] > b[1] and a or b,
							πF.SetLineno(342)
							πTemp001 = πF.MakeArgs(2)
							πTemp005 = make([]πg.Param, 2)
							πTemp005[0] = πg.Param{Name: "a", Def: nil}
							πTemp005[1] = πg.Param{Name: "b", Def: nil}
							πTemp006 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/csv.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µa *πg.Object = πArgs[0]; _ = µa
								var µb *πg.Object = πArgs[1]; _ = µb
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
								var πTemp008 *πg.Object
								_ = πTemp008
								var πR *πg.Object; _ = πR
								var πE *πg.BaseException; _ = πE
								for ; πF.State() >= 0; πF.PopCheckpoint() {
									switch πF.State() {
									case 0:
									default: panic("unexpected function state")
									}
									// line 342: modes[char] = reduce(lambda a, b: a[1] > b[1] and a or b,
									πF.SetLineno(342)
									πTemp006 = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									if πTemp007, πE = πg.GetItem(πF, µa, πTemp006); πE != nil {
										continue
									}
									πTemp006 = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									if πTemp008, πE = πg.GetItem(πF, µb, πTemp006); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GT(πF, πTemp007, πTemp008); πE != nil {
										continue
									}
									πTemp003 = πTemp005
									if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
										continue
									}
									if !πTemp004 {
										goto Label2
									}
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									πTemp003 = µa
								Label2:
									πTemp001 = πTemp003
									if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
										continue
									}
									if πTemp002 {
										goto Label1
									}
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									πTemp001 = µb
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
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp001[1] = µitems
							if πTemp006, πE = πg.ResolveGlobal(πF, ßreduce); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodes, "modes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp013 = µchar
							if πE = πg.SetItem(πF, µmodes, πTemp013, πTemp006); πE != nil {
								continue
							}
							// line 346: items.remove(modes[char])
							πF.SetLineno(346)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp006 = µchar
							if πE = πg.CheckLocal(πF, µmodes, "modes"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µmodes, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µitems, ßremove, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 347: modes[char] = (modes[char][0], modes[char][1]
							πF.SetLineno(347)
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp014 = µchar
							if πE = πg.CheckLocal(πF, µmodes, "modes"); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetItem(πF, µmodes, πTemp014); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, πTemp015, πTemp008); πE != nil {
								continue
							}
							πTemp014 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp016 = µchar
							if πE = πg.CheckLocal(πF, µmodes, "modes"); πE != nil {
								continue
							}
							if πTemp017, πE = πg.GetItem(πF, µmodes, πTemp016); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetItem(πF, πTemp017, πTemp014); πE != nil {
								continue
							}
							πTemp014 = πg.NewInt(1).ToObject()
							πTemp001 = πF.MakeArgs(2)
							πTemp005 = make([]πg.Param, 2)
							πTemp005[0] = πg.Param{Name: "a", Def: nil}
							πTemp005[1] = πg.Param{Name: "b", Def: nil}
							πTemp017 = πg.NewFunction(πg.NewCode("<lambda>", "/usr/lib/python2.7/csv.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µa *πg.Object = πArgs[0]; _ = µa
								var µb *πg.Object = πArgs[1]; _ = µb
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
									// line 348: - reduce(lambda a, b: (0, a[1] + b[1]),
									πF.SetLineno(348)
									πTemp003 = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
										continue
									}
									if πTemp004, πE = πg.GetItem(πF, µa, πTemp003); πE != nil {
										continue
									}
									πTemp003 = πg.NewInt(1).ToObject()
									if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
										continue
									}
									if πTemp005, πE = πg.GetItem(πF, µb, πTemp003); πE != nil {
										continue
									}
									if πTemp002, πE = πg.Add(πF, πTemp004, πTemp005); πE != nil {
										continue
									}
									πTemp001 = πg.NewTuple2(πg.NewInt(0).ToObject(), πTemp002).ToObject()
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
							πTemp001[0] = πTemp017
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							πTemp001[1] = µitems
							if πTemp017, πE = πg.ResolveGlobal(πF, ßreduce); πE != nil {
								continue
							}
							if πTemp018, πE = πTemp017.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp016, πE = πg.GetItem(πF, πTemp018, πTemp014); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Sub(πF, πTemp015, πTemp016); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(πTemp013, πTemp008).ToObject()
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πTemp006); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodes, "modes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp013 = µchar
							if πE = πg.SetItem(πF, µmodes, πTemp013, πTemp008); πE != nil {
								continue
							}
							goto Label18
						Label17:
							// line 351: modes[char] = items[0]
							πF.SetLineno(351)
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µitems, πTemp006); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp008); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µmodes, "modes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
								continue
							}
							πTemp013 = µchar
							if πE = πg.SetItem(πF, µmodes, πTemp013, πTemp006); πE != nil {
								continue
							}
							goto Label18
						Label18:
							continue
						Label11:
							if πE != nil || πR != nil {
								continue
							}
						Label12:
							// line 354: modeList = modes.items()
							πF.SetLineno(354)
							if πE = πg.CheckLocal(πF, µmodes, "modes"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µmodes, ßitems, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							µmodeList = πTemp006
							// line 355: total = float(chunkLength * iteration)
							πF.SetLineno(355)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µchunkLength, "chunkLength"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µiteration, "iteration"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Mul(πF, µchunkLength, µiteration); πE != nil {
								continue
							}
							πTemp001[0] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µtotal = πTemp006
							// line 357: consistency = 1.0
							πF.SetLineno(357)
							µconsistency = πg.NewFloat(1.0).ToObject()
							// line 359: threshold = 0.9
							πF.SetLineno(359)
							µthreshold = πg.NewFloat(0.9).ToObject()
							// line 360: while len(delims) == 0 and consistency >= threshold:
							πF.SetLineno(360)
							πF.PushCheckpoint(20)
							πTemp010 = false
						Label19:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp010 {
								πF.PopCheckpoint()
								goto Label21
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							πTemp001[0] = µdelims
							if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.Eq(πF, πTemp013, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp002 = πTemp006
							if πTemp012, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label22
							}
							if πE = πg.CheckLocal(πF, µconsistency, "consistency"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µthreshold, "threshold"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GE(πF, µconsistency, µthreshold); πE != nil {
								continue
							}
							πTemp002 = πTemp006
						Label22:
							if πTemp011, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(19)            
							if πE = πg.CheckLocal(πF, µmodeList, "modeList"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, µmodeList); πE != nil {
								continue
							}
							πF.PushCheckpoint(24)
							πTemp011 = false
						Label23:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp011 {
								πF.PopCheckpoint()
								goto Label25
							}
							if πTemp006, πE = πg.Next(πF, πTemp002); πE != nil {
								isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
								if exc != nil {
									πE = exc
								} else if isStop {
									πE = nil
									πF.RestoreExc(nil, nil)
								}
								πTemp012 = !isStop
							} else {
								πTemp012 = true
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp013}}}, πTemp006); πE != nil {
									continue
								}
								µk = πTemp008
								µv = πTemp013
							}
							if πE != nil || !πTemp012 {
								continue
							}
							πF.PushCheckpoint(23)            
							πTemp013 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µv, πTemp013); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GT(πF, πTemp014, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp008
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label26
							}
							πTemp013 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µv, πTemp013); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GT(πF, πTemp014, πg.NewInt(0).ToObject()); πE != nil {
								continue
							}
							πTemp006 = πTemp008
						Label26:
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label27
							}
							goto Label28
							// line 362: if v[0] > 0 and v[1] > 0:
							πF.SetLineno(362)
						Label27:
							πTemp014 = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πTemp015, πE = πg.GetItem(πF, µv, πTemp014); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µtotal, "total"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.Div(πF, πTemp015, µtotal); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µconsistency, "consistency"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GE(πF, πTemp013, µconsistency); πE != nil {
								continue
							}
							πTemp006 = πTemp008
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if !πTemp012 {
								goto Label29
							}
							if πE = πg.CheckLocal(πF, µdelimiters, "delimiters"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp013 = πg.GetBool(µdelimiters == πTemp014).ToObject()
							πTemp008 = πTemp013
							if πTemp019, πE = πg.IsTrue(πF, πTemp008); πE != nil {
								continue
							}
							if πTemp019 {
								goto Label30
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdelimiters, "delimiters"); πE != nil {
								continue
							}
							if πTemp020, πE = πg.Contains(πF, µdelimiters, µk); πE != nil {
								continue
							}
							πTemp013 = πg.GetBool(πTemp020).ToObject()
							πTemp008 = πTemp013
						Label30:
							πTemp006 = πTemp008
						Label29:
							if πTemp012, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp012 {
								goto Label31
							}
							goto Label32
							// line 363: if ((v[1]/total) >= consistency and
							πF.SetLineno(363)
						Label31:
							// line 365: delims[k] = v
							πF.SetLineno(365)
							if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µv); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
								continue
							}
							πTemp008 = µk
							if πE = πg.SetItem(πF, µdelims, πTemp008, πTemp006); πE != nil {
								continue
							}
							goto Label32
						Label32:
							goto Label28
						Label28:
							continue
						Label24:
							if πE != nil || πR != nil {
								continue
							}
						Label25:
							// line 366: consistency -= 0.01
							πF.SetLineno(366)
							if πE = πg.CheckLocal(πF, µconsistency, "consistency"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ISub(πF, µconsistency, πg.NewFloat(0.01).ToObject()); πE != nil {
								continue
							}
							µconsistency = πTemp002
							continue
						Label20:
							if πE != nil || πR != nil {
								continue
							}
						Label21:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							πTemp001[0] = µdelims
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label33
							}
							goto Label34
							// line 368: if len(delims) == 1:
							πF.SetLineno(368)
						Label33:
							// line 369: delim = delims.keys()[0]
							πF.SetLineno(369)
							πTemp002 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µdelims, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, πTemp013, πTemp002); πE != nil {
								continue
							}
							µdelim = πTemp006
							// line 370: skipinitialspace = (data[0].count(delim) ==
							πF.SetLineno(370)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							πTemp001[0] = µdelim
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, µdata, πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp008, ßcount, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Mod(πF, πg.NewStr("%c ").ToObject(), µdelim); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							πTemp006 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µdata, πTemp006); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, πTemp013, ßcount, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp008, πTemp013); πE != nil {
								continue
							}
							µskipinitialspace = πTemp002
							// line 372: return (delim, skipinitialspace)
							πF.SetLineno(372)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µskipinitialspace, "skipinitialspace"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µdelim, µskipinitialspace).ToObject()
							πR = πTemp002
							continue
							goto Label34
						Label34:
							// line 375: start = end
							πF.SetLineno(375)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							µstart = µend
							// line 376: end += chunkLength
							πF.SetLineno(376)
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µchunkLength, "chunkLength"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.IAdd(πF, µend, µchunkLength); πE != nil {
								continue
							}
							µend = πTemp002
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µdelims); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(!πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label35
							}
							goto Label36
							// line 378: if not delims:
							πF.SetLineno(378)
						Label35:
							// line 379: return ('', 0)
							πF.SetLineno(379)
							πTemp002 = πg.NewTuple2(ß.ToObject(), πg.NewInt(0).ToObject()).ToObject()
							πR = πTemp002
							continue
							goto Label36
						Label36:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							πTemp001[0] = µdelims
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.GT(πF, πTemp008, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label37
							}
							goto Label38
							// line 382: if len(delims) > 1:
							πF.SetLineno(382)
						Label37:
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µself, ßpreferred, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(40)
							πTemp009 = false
						Label39:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp009 {
								πF.PopCheckpoint()
								goto Label41
							}
							if πTemp006, πE = πg.Next(πF, πTemp002); πE != nil {
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
								µd = πTemp006
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(39)            
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µdelims, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp008.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp010, πE = πg.Contains(πF, πTemp013, µd); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(πTemp010).ToObject()
							if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label42
							}
							goto Label43
							// line 384: if d in delims.keys():
							πF.SetLineno(384)
						Label42:
							// line 385: skipinitialspace = (data[0].count(d) ==
							πF.SetLineno(385)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							πTemp001[0] = µd
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µdata, πTemp008); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp013, ßcount, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("%c ").ToObject(), µd); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µdata, πTemp008); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp014, ßcount, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp006, πE = πg.Eq(πF, πTemp013, πTemp014); πE != nil {
								continue
							}
							µskipinitialspace = πTemp006
							// line 387: return (d, skipinitialspace)
							πF.SetLineno(387)
							if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µskipinitialspace, "skipinitialspace"); πE != nil {
								continue
							}
							πTemp006 = πg.NewTuple2(µd, µskipinitialspace).ToObject()
							πR = πTemp006
							continue
							goto Label43
						Label43:
							continue
						Label40:
							if πE != nil || πR != nil {
								continue
							}
						Label41:
							goto Label38
						Label38:
							// line 391: items = [(v,k) for (k,v) in delims.items()]
							πF.SetLineno(391)
							πTemp005 = make([]πg.Param, 0)
							πTemp006 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/csv.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
								var µk *πg.Object = πg.UnboundLocal; _ = µk
								var µv *πg.Object = πg.UnboundLocal; _ = µv
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
										if πE = πg.CheckLocal(πF, µdelims, "delims"); πE != nil {
											continue
										}
										if πTemp002, πE = πg.GetAttr(πF, µdelims, ßitems, nil); πE != nil {
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
											µk = πTemp003
											µv = πTemp006
										}
										if πE != nil || !πTemp005 {
											continue
										}
										πF.PushCheckpoint(1)            
										// line 391: items = [(v,k) for (k,v) in delims.items()]
										πF.SetLineno(391)
										if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
											continue
										}
										if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
											continue
										}
										πTemp002 = πg.NewTuple2(µv, µk).ToObject()
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
							if πTemp008, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
								continue
							}
							µitems = πTemp002
							// line 392: items.sort()
							πF.SetLineno(392)
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µitems, ßsort, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πTemp002.Call(πF, nil, nil); πE != nil {
								continue
							}
							// line 393: delim = items[-1][1]
							πF.SetLineno(393)
							πTemp002 = πg.NewInt(1).ToObject()
							if πTemp014, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp013 = πTemp014
							if πE = πg.CheckLocal(πF, µitems, "items"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µitems, πTemp013); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetItem(πF, πTemp014, πTemp002); πE != nil {
								continue
							}
							µdelim = πTemp008
							// line 395: skipinitialspace = (data[0].count(delim) ==
							πF.SetLineno(395)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							πTemp001[0] = µdelim
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp013, πE = πg.GetItem(πF, µdata, πTemp008); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp013, ßcount, nil); πE != nil {
								continue
							}
							if πTemp013, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.Mod(πF, πg.NewStr("%c ").ToObject(), µdelim); πE != nil {
								continue
							}
							πTemp001[0] = πTemp008
							πTemp008 = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
								continue
							}
							if πTemp014, πE = πg.GetItem(πF, µdata, πTemp008); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp014, ßcount, nil); πE != nil {
								continue
							}
							if πTemp014, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.Eq(πF, πTemp013, πTemp014); πE != nil {
								continue
							}
							µskipinitialspace = πTemp002
							// line 397: return (delim, skipinitialspace)
							πF.SetLineno(397)
							if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µskipinitialspace, "skipinitialspace"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µdelim, µskipinitialspace).ToObject()
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
					if πE = πClass.SetItem(πF, ß_guess_delimiter.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 296: """
					πF.SetLineno(296)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("\n        The delimiter /should/ occur the same number of times on\n        each row. However, due to malformed data, it may not. We don't want\n        an all or nothing approach, so we allow for small variations in this\n        number.\n          1) build a table of the frequency of each character on every line.\n          2) build a table of frequencies of this frequency (meta-frequency?),\n             e.g.  'x occurred 5 times in 10 rows, 6 times in 1000 rows,\n             7 times in 2 rows'\n          3) use the mode of the meta-frequency to determine the /expected/\n             frequency for that character\n          4) find out how often the character actually meets that goal\n          5) the character that best meets its goal is the delimiter\n        For performance reasons, the data is evaluated in chunks, so it can\n        try and evaluate the smallest portion of the data possible, evaluating\n        additional chunks as necessary.\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ß_guess_delimiter); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
						continue
					}
					// line 400: def has_header(self, sample):
					πF.SetLineno(400)
					πTemp002 = make([]πg.Param, 2)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "sample", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("has_header", "/usr/lib/python2.7/csv.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µsample *πg.Object = πArgs[1]; _ = µsample
						var µrdr *πg.Object = πg.UnboundLocal; _ = µrdr
						var µheader *πg.Object = πg.UnboundLocal; _ = µheader
						var µcolumns *πg.Object = πg.UnboundLocal; _ = µcolumns
						var µcolumnTypes *πg.Object = πg.UnboundLocal; _ = µcolumnTypes
						var µi *πg.Object = πg.UnboundLocal; _ = µi
						var µchecked *πg.Object = πg.UnboundLocal; _ = µchecked
						var µrow *πg.Object = πg.UnboundLocal; _ = µrow
						var µcol *πg.Object = πg.UnboundLocal; _ = µcol
						var µthisType *πg.Object = πg.UnboundLocal; _ = µthisType
						var µhasHeader *πg.Object = πg.UnboundLocal; _ = µhasHeader
						var µcolType *πg.Object = πg.UnboundLocal; _ = µcolType
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 []*πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Dict
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πTemp008 bool
						_ = πTemp008
						var πTemp009 *πg.Object
						_ = πTemp009
						var πTemp010 bool
						_ = πTemp010
						var πTemp011 bool
						_ = πTemp011
						var πTemp012 *πg.Object
						_ = πTemp012
						var πTemp013 *πg.BaseException
						_ = πTemp013
						var πTemp014 *πg.Traceback
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
							case 11: goto Label11
							case 12: goto Label12
							case 14: goto Label14
							case 15: goto Label15
							case 18: goto Label18
							case 27: goto Label27
							case 28: goto Label28
							case 37: goto Label37
							default: panic("unexpected function state")
							}
							// line 410: rdr = reader(StringIO(sample), self.sniff(sample))
							πF.SetLineno(410)
							πTemp001 = πF.MakeArgs(2)
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsample, "sample"); πE != nil {
								continue
							}
							πTemp002[0] = µsample
							if πTemp003, πE = πg.ResolveGlobal(πF, ßStringIO); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							πTemp002 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µsample, "sample"); πE != nil {
								continue
							}
							πTemp002[0] = µsample
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßsniff, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[1] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßreader); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µrdr = πTemp004
							// line 412: header = rdr.next() # assume first row is header
							πF.SetLineno(412)
							if πE = πg.CheckLocal(πF, µrdr, "rdr"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µrdr, ßnext, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
								continue
							}
							µheader = πTemp004
							// line 414: columns = len(header)
							πF.SetLineno(414)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							πTemp001[0] = µheader
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µcolumns = πTemp004
							// line 415: columnTypes = {}
							πF.SetLineno(415)
							πTemp005 = πg.NewDict()
							πTemp003 = πTemp005.ToObject()
							µcolumnTypes = πTemp003
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							πTemp001[0] = µcolumns
							if πTemp004, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.Iter(πF, πTemp006); πE != nil {
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
								µi = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(1)            
							// line 416: for i in range(columns): columnTypes[i] = None
							πF.SetLineno(416)
							if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp004); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolumnTypes, "columnTypes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
								continue
							}
							πTemp009 = µi
							if πE = πg.SetItem(πF, µcolumnTypes, πTemp009, πTemp006); πE != nil {
								continue
							}
							continue
						Label2:
							if πE != nil || πR != nil {
								continue
							}
						Label3:
							// line 418: checked = 0
							πF.SetLineno(418)
							µchecked = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µrdr, "rdr"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, µrdr); πE != nil {
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
								µrow = πTemp004
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(4)            
							if πE = πg.CheckLocal(πF, µchecked, "checked"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GT(πF, µchecked, πg.NewInt(20).ToObject()); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label7
							}
							goto Label8
							// line 421: if checked > 20:
							πF.SetLineno(421)
						Label7:
							// line 422: break
							πF.SetLineno(422)
							πTemp007 = true
							continue
							goto Label8
						Label8:
							// line 423: checked += 1
							πF.SetLineno(423)
							if πE = πg.CheckLocal(πF, µchecked, "checked"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IAdd(πF, µchecked, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µchecked = πTemp004
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							πTemp001[0] = µrow
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µcolumns, "columns"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.NE(πF, πTemp009, µcolumns); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label9
							}
							goto Label10
							// line 425: if len(row) != columns:
							πF.SetLineno(425)
						Label9:
							// line 426: continue # skip rows that have irregular number of columns
							πF.SetLineno(426)
							continue
							goto Label10
						Label10:
							if πE = πg.CheckLocal(πF, µcolumnTypes, "columnTypes"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetAttr(πF, µcolumnTypes, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp006.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πg.Iter(πF, πTemp009); πE != nil {
								continue
							}
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
							if πTemp006, πE = πg.Next(πF, πTemp004); πE != nil {
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
								µcol = πTemp006
							}
							if πE != nil || !πTemp010 {
								continue
							}
							πF.PushCheckpoint(11)            
							πTemp001 = make([]*πg.Object, 4)
							if πTemp009, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
								continue
							}
							πTemp001[1] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							πTemp001[2] = πTemp009
							if πTemp009, πE = πg.ResolveGlobal(πF, ßcomplex); πE != nil {
								continue
							}
							πTemp001[3] = πTemp009
							πTemp009 = πg.NewList(πTemp001...).ToObject()
							if πTemp006, πE = πg.Iter(πF, πTemp009); πE != nil {
								continue
							}
							πF.PushCheckpoint(15)
							πTemp010 = false
						Label14:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp010 {
								πF.PopCheckpoint()
								goto Label16
							}
							if πTemp009, πE = πg.Next(πF, πTemp006); πE != nil {
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
								µthisType = πTemp009
							}
							if πE != nil || !πTemp011 {
								continue
							}
							πF.PushCheckpoint(14)            
							// line 431: try:
							πF.SetLineno(431)
							πF.PushCheckpoint(18)
							// line 432: thisType(row[col])
							πF.SetLineno(432)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp009 = µcol
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µrow, πTemp009); πE != nil {
								continue
							}
							πTemp001[0] = πTemp012
							if πE = πg.CheckLocal(πF, µthisType, "thisType"); πE != nil {
								continue
							}
							if πTemp009, πE = µthisType.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 433: break
							πF.SetLineno(433)
							πTemp010 = true
							continue
							πF.PopCheckpoint()
							goto Label17
						Label18:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp013, πTemp014 = πF.ExcInfo()
							if πTemp012, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp015, πE = πg.ResolveGlobal(πF, ßOverflowError); πE != nil {
								continue
							}
							πTemp009 = πg.NewTuple2(πTemp012, πTemp015).ToObject()
							if πTemp011, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp009); πE != nil {
								continue
							}
							if πTemp011 {
								goto Label19
							}
							πE = πF.Raise(πTemp013.ToObject(), nil, πTemp014.ToObject())
							continue
							// line 434: except (ValueError, OverflowError):
							πF.SetLineno(434)
						Label19:
							// line 435: pass
							πF.SetLineno(435)
							πF.RestoreExc(nil, nil)
							goto Label17
						Label17:
							continue
						Label15:
							if πE != nil || πR != nil {
								continue
							}
							// line 438: thisType = len(row[col])
							πF.SetLineno(438)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp009 = µcol
							if πE = πg.CheckLocal(πF, µrow, "row"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µrow, πTemp009); πE != nil {
								continue
							}
							πTemp001[0] = πTemp012
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µthisType = πTemp012
						Label16:
							if πE = πg.CheckLocal(πF, µthisType, "thisType"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßlong); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Eq(πF, µthisType, πTemp009); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label20
							}
							goto Label21
							// line 441: if thisType == long:
							πF.SetLineno(441)
						Label20:
							// line 442: thisType = int
							πF.SetLineno(442)
							if πTemp006, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							µthisType = πTemp006
							goto Label21
						Label21:
							if πE = πg.CheckLocal(πF, µthisType, "thisType"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp009 = µcol
							if πE = πg.CheckLocal(πF, µcolumnTypes, "columnTypes"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µcolumnTypes, πTemp009); πE != nil {
								continue
							}
							if πTemp006, πE = πg.NE(πF, µthisType, πTemp012); πE != nil {
								continue
							}
							if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label22
							}
							goto Label23
							// line 444: if thisType != columnTypes[col]:
							πF.SetLineno(444)
						Label22:
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp009 = µcol
							if πE = πg.CheckLocal(πF, µcolumnTypes, "columnTypes"); πE != nil {
								continue
							}
							if πTemp012, πE = πg.GetItem(πF, µcolumnTypes, πTemp009); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πTemp006 = πg.GetBool(πTemp012 == πTemp009).ToObject()
							if πTemp010, πE = πg.IsTrue(πF, πTemp006); πE != nil {
								continue
							}
							if πTemp010 {
								goto Label24
							}
							goto Label25
							// line 445: if columnTypes[col] is None: # add new column type
							πF.SetLineno(445)
						Label24:
							// line 446: columnTypes[col] = thisType
							πF.SetLineno(446)
							if πE = πg.CheckLocal(πF, µthisType, "thisType"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, µthisType); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcolumnTypes, "columnTypes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp009 = µcol
							if πE = πg.SetItem(πF, µcolumnTypes, πTemp009, πTemp006); πE != nil {
								continue
							}
							goto Label26
						Label25:
							// line 450: del columnTypes[col]
							πF.SetLineno(450)
							if πE = πg.CheckLocal(πF, µcolumnTypes, "columnTypes"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp006 = µcol
							if πE = πg.DelItem(πF, µcolumnTypes, πTemp006); πE != nil {
								continue
							}
							goto Label26
						Label26:
							goto Label23
						Label23:
							continue
						Label12:
							if πE != nil || πR != nil {
								continue
							}
						Label13:
							continue
						Label5:
							if πE != nil || πR != nil {
								continue
							}
						Label6:
							// line 454: hasHeader = 0
							πF.SetLineno(454)
							µhasHeader = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µcolumnTypes, "columnTypes"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µcolumnTypes, ßitems, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Iter(πF, πTemp006); πE != nil {
								continue
							}
							πF.PushCheckpoint(28)
							πTemp007 = false
						Label27:
							if πE != nil || πR != nil {
								continue
							}
							if πTemp007 {
								πF.PopCheckpoint()
								goto Label29
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
								if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
									continue
								}
								µcol = πTemp006
								µcolType = πTemp009
							}
							if πE != nil || !πTemp008 {
								continue
							}
							πF.PushCheckpoint(27)            
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcolType, "colType"); πE != nil {
								continue
							}
							πTemp001[0] = µcolType
							if πTemp006, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewInt(0).ToObject()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
								continue
							}
							if πTemp012, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.Eq(πF, πTemp009, πTemp012); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label30
							}
							goto Label31
							// line 456: if type(colType) == type(0): # it's a length
							πF.SetLineno(456)
						Label30:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp006 = µcol
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.GetItem(πF, µheader, πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp009
							if πTemp006, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.CheckLocal(πF, µcolType, "colType"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.NE(πF, πTemp009, µcolType); πE != nil {
								continue
							}
							if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label33
							}
							goto Label34
							// line 457: if len(header[col]) != colType:
							πF.SetLineno(457)
						Label33:
							// line 458: hasHeader += 1
							πF.SetLineno(458)
							if πE = πg.CheckLocal(πF, µhasHeader, "hasHeader"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IAdd(πF, µhasHeader, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µhasHeader = πTemp004
							goto Label35
						Label34:
							// line 460: hasHeader -= 1
							πF.SetLineno(460)
							if πE = πg.CheckLocal(πF, µhasHeader, "hasHeader"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ISub(πF, µhasHeader, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µhasHeader = πTemp004
							goto Label35
						Label35:
							goto Label32
						Label31:
							// line 462: try:
							πF.SetLineno(462)
							πF.PushCheckpoint(37)
							// line 463: colType(header[col])
							πF.SetLineno(463)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µcol, "col"); πE != nil {
								continue
							}
							πTemp004 = µcol
							if πE = πg.CheckLocal(πF, µheader, "header"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.GetItem(πF, µheader, πTemp004); πE != nil {
								continue
							}
							πTemp001[0] = πTemp006
							if πE = πg.CheckLocal(πF, µcolType, "colType"); πE != nil {
								continue
							}
							if πTemp004, πE = µcolType.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							πF.PopCheckpoint()
							// line 467: hasHeader -= 1
							πF.SetLineno(467)
							if πE = πg.CheckLocal(πF, µhasHeader, "hasHeader"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.ISub(πF, µhasHeader, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µhasHeader = πTemp004
							goto Label36
						Label37:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp013, πTemp014 = πF.ExcInfo()
							if πTemp006, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp009, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
								continue
							}
							πTemp004 = πg.NewTuple2(πTemp006, πTemp009).ToObject()
							if πTemp008, πE = πg.IsInstance(πF, πTemp013.ToObject(), πTemp004); πE != nil {
								continue
							}
							if πTemp008 {
								goto Label38
							}
							πE = πF.Raise(πTemp013.ToObject(), nil, πTemp014.ToObject())
							continue
							// line 464: except (ValueError, TypeError):
							πF.SetLineno(464)
						Label38:
							// line 465: hasHeader += 1
							πF.SetLineno(465)
							if πE = πg.CheckLocal(πF, µhasHeader, "hasHeader"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IAdd(πF, µhasHeader, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							µhasHeader = πTemp004
							πF.RestoreExc(nil, nil)
							goto Label36
						Label36:
							goto Label32
						Label32:
							continue
						Label28:
							if πE != nil || πR != nil {
								continue
							}
						Label29:
							// line 469: return hasHeader > 0
							πF.SetLineno(469)
							if πE = πg.CheckLocal(πF, µhasHeader, "hasHeader"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GT(πF, µhasHeader, πg.NewInt(0).ToObject()); πE != nil {
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
					if πE = πClass.SetItem(πF, ßhas_header.ToObject(), πTemp006); πE != nil {
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
			if πTemp004, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Sniffer").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp010.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSniffer.ToObject(), πTemp004); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("csv", Code)
}

