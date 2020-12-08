package decoder
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/sys"
	// _ "github.com/pygolin/stdlib/re"
	// _ "github.com/pygolin/stdlib/json_scanner"
	// _ "github.com/pygolin/stdlib/_struct"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/json/decoder.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ßBACKSLASH := πg.InternStr("BACKSLASH")
		ßDEFAULT_ENCODING := πg.InternStr("DEFAULT_ENCODING")
		ßDOTALL := πg.InternStr("DOTALL")
		ßFLAGS := πg.InternStr("FLAGS")
		ßIndexError := πg.InternStr("IndexError")
		ßInfinity := πg.InternStr("Infinity")
		ßJSONArray := πg.InternStr("JSONArray")
		ßJSONDecoder := πg.InternStr("JSONDecoder")
		ßJSONObject := πg.InternStr("JSONObject")
		ßKeyError := πg.InternStr("KeyError")
		ßMULTILINE := πg.InternStr("MULTILINE")
		ßNaN := πg.InternStr("NaN")
		ßNegInf := πg.InternStr("NegInf")
		ßNone := πg.InternStr("None")
		ßPosInf := πg.InternStr("PosInf")
		ßSTRINGCHUNK := πg.InternStr("STRINGCHUNK")
		ßStopIteration := πg.InternStr("StopIteration")
		ßTrue := πg.InternStr("True")
		ßVERBOSE := πg.InternStr("VERBOSE")
		ßValueError := πg.InternStr("ValueError")
		ßWHITESPACE := πg.InternStr("WHITESPACE")
		ßWHITESPACE_STR := πg.InternStr("WHITESPACE_STR")
		ß_CONSTANTS := πg.InternStr("_CONSTANTS")
		ß__all__ := πg.InternStr("__all__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__getitem__ := πg.InternStr("__getitem__")
		ß__init__ := πg.InternStr("__init__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß_decode_uXXXX := πg.InternStr("_decode_uXXXX")
		ß_floatconstants := πg.InternStr("_floatconstants")
		ßappend := πg.InternStr("append")
		ßb := πg.InternStr("b")
		ßc_scanstring := πg.InternStr("c_scanstring")
		ßcompile := πg.InternStr("compile")
		ßcount := πg.InternStr("count")
		ßdecode := πg.InternStr("decode")
		ßdict := πg.InternStr("dict")
		ßencoding := πg.InternStr("encoding")
		ßend := πg.InternStr("end")
		ßerrmsg := πg.InternStr("errmsg")
		ßf := πg.InternStr("f")
		ßfloat := πg.InternStr("float")
		ßgroups := πg.InternStr("groups")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßjoin := πg.InternStr("join")
		ßlen := πg.InternStr("len")
		ßlinecol := πg.InternStr("linecol")
		ßmake_scanner := πg.InternStr("make_scanner")
		ßmatch := πg.InternStr("match")
		ßmaxunicode := πg.InternStr("maxunicode")
		ßn := πg.InternStr("n")
		ßobject := πg.InternStr("object")
		ßobject_hook := πg.InternStr("object_hook")
		ßobject_pairs_hook := πg.InternStr("object_pairs_hook")
		ßparse_array := πg.InternStr("parse_array")
		ßparse_constant := πg.InternStr("parse_constant")
		ßparse_float := πg.InternStr("parse_float")
		ßparse_int := πg.InternStr("parse_int")
		ßparse_object := πg.InternStr("parse_object")
		ßparse_string := πg.InternStr("parse_string")
		ßpy_scanstring := πg.InternStr("py_scanstring")
		ßr := πg.InternStr("r")
		ßraw_decode := πg.InternStr("raw_decode")
		ßre := πg.InternStr("re")
		ßrepr := πg.InternStr("repr")
		ßrindex := πg.InternStr("rindex")
		ßscan_once := πg.InternStr("scan_once")
		ßscanner := πg.InternStr("scanner")
		ßscanstring := πg.InternStr("scanstring")
		ßstrict := πg.InternStr("strict")
		ßstruct := πg.InternStr("struct")
		ßsys := πg.InternStr("sys")
		ßt := πg.InternStr("t")
		ßu := πg.InternStr("u")
		ßunichr := πg.InternStr("unichr")
		ßunicode := πg.InternStr("unicode")
		ßunpack := πg.InternStr("unpack")
		ßxX := πg.InternStr("xX")
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
		var πTemp007 []πg.Param
		_ = πTemp007
		var πTemp008 *πg.Dict
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 bool
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
			// line 1: """Implementation of JSONDecoder
			πF.SetLineno(1)
			// line 1: """Implementation of JSONDecoder
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Implementation of JSONDecoder\n").ToObject()); πE != nil {
				continue
			}
			// line 3: import re
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
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
			// line 5: import _struct as struct
			πF.SetLineno(5)
			if πTemp002, πE = πg.ImportModule(πF, "_struct"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßstruct.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 8: import json_scanner as scanner
			πF.SetLineno(8)
			if πTemp002, πE = πg.ImportModule(πF, "json_scanner"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßscanner.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 14: c_scanstring = None
			πF.SetLineno(14)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßc_scanstring.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 16: __all__ = ['JSONDecoder']
			πF.SetLineno(16)
			πTemp002 = make([]*πg.Object, 1)
			πTemp002[0] = ßJSONDecoder.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 18: FLAGS = re.VERBOSE | re.MULTILINE | re.DOTALL
			πF.SetLineno(18)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßVERBOSE, nil); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp004, ßMULTILINE, nil); πE != nil {
				continue
			}
			if πTemp003, πE = πg.Or(πF, πTemp005, πTemp006); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßDOTALL, nil); πE != nil {
				continue
			}
			if πTemp001, πE = πg.Or(πF, πTemp003, πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßFLAGS.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 20: def _floatconstants():
			πF.SetLineno(20)
			πTemp007 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("_floatconstants", "/usr/lib/python2.7/json/decoder.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µnan *πg.Object = πg.UnboundLocal; _ = µnan
				var µinf *πg.Object = πg.UnboundLocal; _ = µinf
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
					// line 21: nan = struct.unpack('>d', b'\x7f\xf8\x00\x00\x00\x00\x00\x00')
					πF.SetLineno(21)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr(">d").ToObject()
					πTemp001[1] = πg.NewStr("\x7f\xf8\x00\x00\x00\x00\x00\x00").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstruct); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunpack, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µnan = πTemp002
					// line 22: inf = struct.unpack('>d', b'\x7f\xf0\x00\x00\x00\x00\x00\x00')
					πF.SetLineno(22)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr(">d").ToObject()
					πTemp001[1] = πg.NewStr("\x7f\xf0\x00\x00\x00\x00\x00\x00").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßstruct); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßunpack, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µinf = πTemp002
					// line 23: nan = nan[0]
					πF.SetLineno(23)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µnan, "nan"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µnan, πTemp002); πE != nil {
						continue
					}
					µnan = πTemp003
					// line 24: inf = inf[0]
					πF.SetLineno(24)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µinf, "inf"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µinf, πTemp002); πE != nil {
						continue
					}
					µinf = πTemp003
					// line 25: return nan, inf, -inf
					πF.SetLineno(25)
					if πE = πg.CheckLocal(πF, µnan, "nan"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinf, "inf"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µinf, "inf"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Neg(πF, µinf); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple3(µnan, µinf, πTemp003).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_floatconstants.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: NaN, PosInf, NegInf = _floatconstants()
			πF.SetLineno(27)
			if πTemp003, πE = πg.ResolveGlobal(πF, ß_floatconstants); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
			if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNaN.ToObject(), πTemp003); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßPosInf.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßNegInf.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 30: def linecol(doc, pos):
			πF.SetLineno(30)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "doc", Def: nil}
			πTemp007[1] = πg.Param{Name: "pos", Def: nil}
			πTemp003 = πg.NewFunction(πg.NewCode("linecol", "/usr/lib/python2.7/json/decoder.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdoc *πg.Object = πArgs[0]; _ = µdoc
				var µpos *πg.Object = πArgs[1]; _ = µpos
				var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
				var µcolno *πg.Object = πg.UnboundLocal; _ = µcolno
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
					// line 31: lineno = doc.count('\n', 0, pos) + 1
					πF.SetLineno(31)
					πTemp002 = πF.MakeArgs(3)
					πTemp002[0] = πg.NewStr("\n").ToObject()
					πTemp002[1] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp002[2] = µpos
					if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µdoc, ßcount, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Add(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µlineno = πTemp001
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µlineno, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 32: if lineno == 1:
					πF.SetLineno(32)
				Label1:
					// line 33: colno = pos + 1
					πF.SetLineno(33)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µpos, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µcolno = πTemp001
					goto Label3
				Label2:
					// line 35: colno = pos - doc.rindex('\n', 0, pos)
					πF.SetLineno(35)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp002 = πF.MakeArgs(3)
					πTemp002[0] = πg.NewStr("\n").ToObject()
					πTemp002[1] = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp002[2] = µpos
					if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µdoc, ßrindex, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Sub(πF, µpos, πTemp004); πE != nil {
						continue
					}
					µcolno = πTemp001
					goto Label3
				Label3:
					// line 36: return lineno, colno
					πF.SetLineno(36)
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcolno, "colno"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µlineno, µcolno).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßlinecol.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 39: def errmsg(msg, doc, pos, end=None):
			πF.SetLineno(39)
			πTemp007 = make([]πg.Param, 4)
			πTemp007[0] = πg.Param{Name: "msg", Def: nil}
			πTemp007[1] = πg.Param{Name: "doc", Def: nil}
			πTemp007[2] = πg.Param{Name: "pos", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007[3] = πg.Param{Name: "end", Def: πTemp005}
			πTemp004 = πg.NewFunction(πg.NewCode("errmsg", "/usr/lib/python2.7/json/decoder.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µmsg *πg.Object = πArgs[0]; _ = µmsg
				var µdoc *πg.Object = πArgs[1]; _ = µdoc
				var µpos *πg.Object = πArgs[2]; _ = µpos
				var µend *πg.Object = πArgs[3]; _ = µend
				var µlineno *πg.Object = πg.UnboundLocal; _ = µlineno
				var µcolno *πg.Object = πg.UnboundLocal; _ = µcolno
				var µfmt *πg.Object = πg.UnboundLocal; _ = µfmt
				var µendlineno *πg.Object = πg.UnboundLocal; _ = µendlineno
				var µendcolno *πg.Object = πg.UnboundLocal; _ = µendcolno
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 41: lineno, colno = linecol(doc, pos)
					πF.SetLineno(41)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
						continue
					}
					πTemp001[0] = µdoc
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp001[1] = µpos
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlinecol); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µlineno = πTemp002
					µcolno = πTemp004
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(µend == πTemp003).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label1
					}
					goto Label2
					// line 42: if end is None:
					πF.SetLineno(42)
				Label1:
					// line 45: fmt = '%s: line %d column %d (char %d)'
					πF.SetLineno(45)
					µfmt = πg.NewStr("%s: line %d column %d (char %d)").ToObject()
					// line 46: return fmt % (msg, lineno, colno, pos)
					πF.SetLineno(46)
					if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcolno, "colno"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple4(µmsg, µlineno, µcolno, µpos).ToObject()
					if πTemp002, πE = πg.Mod(πF, µfmt, πTemp003); πE != nil {
						continue
					}
					πR = πTemp002
					continue
					goto Label2
				Label2:
					// line 47: endlineno, endcolno = linecol(doc, end)
					πF.SetLineno(47)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdoc, "doc"); πE != nil {
						continue
					}
					πTemp001[0] = µdoc
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001[1] = µend
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlinecol); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µendlineno = πTemp002
					µendcolno = πTemp004
					// line 50: fmt = '%s: line %d column %d - line %d column %d (char %d - %d)'
					πF.SetLineno(50)
					µfmt = πg.NewStr("%s: line %d column %d - line %d column %d (char %d - %d)").ToObject()
					// line 51: return fmt % (msg, lineno, colno, endlineno, endcolno, pos, end)
					πF.SetLineno(51)
					if πE = πg.CheckLocal(πF, µfmt, "fmt"); πE != nil {
						continue
					}
					πTemp001 = make([]*πg.Object, 7)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp001[0] = µmsg
					if πE = πg.CheckLocal(πF, µlineno, "lineno"); πE != nil {
						continue
					}
					πTemp001[1] = µlineno
					if πE = πg.CheckLocal(πF, µcolno, "colno"); πE != nil {
						continue
					}
					πTemp001[2] = µcolno
					if πE = πg.CheckLocal(πF, µendlineno, "endlineno"); πE != nil {
						continue
					}
					πTemp001[3] = µendlineno
					if πE = πg.CheckLocal(πF, µendcolno, "endcolno"); πE != nil {
						continue
					}
					πTemp001[4] = µendcolno
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp001[5] = µpos
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001[6] = µend
					πTemp003 = πg.NewTuple(πTemp001...).ToObject()
					if πTemp002, πE = πg.Mod(πF, µfmt, πTemp003); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßerrmsg.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 54: _CONSTANTS = {
			πF.SetLineno(54)
			πTemp008 = πg.NewDict()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNegInf); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πg.NewStr("-Infinity").ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßPosInf); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ßInfinity.ToObject(), πTemp005); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNaN); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ßNaN.ToObject(), πTemp005); πE != nil {
				continue
			}
			πTemp005 = πTemp008.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_CONSTANTS.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 60: STRINGCHUNK = re.compile(r'(.*?)(["\\\x00-\x1f])', FLAGS)
			πF.SetLineno(60)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("(.*?)([\"\\\\\\x00-\\x1f])").ToObject()
			if πTemp005, πE = πg.ResolveGlobal(πF, ßFLAGS); πE != nil {
				continue
			}
			πTemp002[1] = πTemp005
			if πTemp005, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßSTRINGCHUNK.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 61: BACKSLASH = {
			πF.SetLineno(61)
			πTemp008 = πg.NewDict()
			if πE = πTemp008.SetItem(πF, πg.NewStr("\"").ToObject(), πg.NewUnicode("\"").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πg.NewStr("\\").ToObject(), πg.NewUnicode("\\").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, πg.NewStr("/").ToObject(), πg.NewUnicode("/").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ßb.ToObject(), πg.NewUnicode("\x08").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ßf.ToObject(), πg.NewUnicode("\x0c").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ßn.ToObject(), πg.NewUnicode("\n").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ßr.ToObject(), πg.NewUnicode("\r").ToObject()); πE != nil {
				continue
			}
			if πE = πTemp008.SetItem(πF, ßt.ToObject(), πg.NewUnicode("\t").ToObject()); πE != nil {
				continue
			}
			πTemp005 = πTemp008.ToObject()
			if πE = πF.Globals().SetItem(πF, ßBACKSLASH.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 66: DEFAULT_ENCODING = "utf-8"
			πF.SetLineno(66)
			if πE = πF.Globals().SetItem(πF, ßDEFAULT_ENCODING.ToObject(), πg.NewStr("utf-8").ToObject()); πE != nil {
				continue
			}
			// line 68: def _decode_uXXXX(s, pos):
			πF.SetLineno(68)
			πTemp007 = make([]πg.Param, 2)
			πTemp007[0] = πg.Param{Name: "s", Def: nil}
			πTemp007[1] = πg.Param{Name: "pos", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("_decode_uXXXX", "/usr/lib/python2.7/json/decoder.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µpos *πg.Object = πArgs[1]; _ = µpos
				var µesc *πg.Object = πg.UnboundLocal; _ = µesc
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
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
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 []*πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 69: esc = s[pos + 1:pos + 5]
					πF.SetLineno(69)
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µpos, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µpos, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp002, πTemp003, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µesc = πTemp002
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µesc, "esc"); πE != nil {
						continue
					}
					πTemp005[0] = µesc
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewInt(4).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp004 {
						goto Label1
					}
					πTemp003 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µesc, "esc"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µesc, πTemp003); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, ßxX.ToObject(), πTemp006); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(!πTemp007).ToObject()
					πTemp001 = πTemp002
				Label1:
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label2
					}
					goto Label3
					// line 70: if len(esc) == 4 and esc[1] not in 'xX':
					πF.SetLineno(70)
				Label2:
					// line 71: try:
					πF.SetLineno(71)
					πF.PushCheckpoint(5)
					// line 72: return int(esc, 16)
					πF.SetLineno(72)
					πTemp005 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µesc, "esc"); πE != nil {
						continue
					}
					πTemp005[0] = µesc
					πTemp005[1] = πg.NewInt(16).ToObject()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πR = πTemp002
					continue
					πF.PopCheckpoint()
					goto Label4
				Label5:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 73: except ValueError:
					πF.SetLineno(73)
				Label6:
					// line 74: pass
					πF.SetLineno(74)
					πF.RestoreExc(nil, nil)
					goto Label4
				Label4:
					goto Label3
				Label3:
					// line 75: msg = "Invalid \\uXXXX escape"
					πF.SetLineno(75)
					µmsg = πg.NewStr("Invalid \\uXXXX escape").ToObject()
					πTemp005 = πF.MakeArgs(1)
					πTemp010 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp010[0] = µmsg
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp010[1] = µs
					if πE = πg.CheckLocal(πF, µpos, "pos"); πE != nil {
						continue
					}
					πTemp010[2] = µpos
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp010, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp010)
					πTemp005[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					// line 76: raise ValueError(errmsg(msg, s, pos))
					πF.SetLineno(76)
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
			if πE = πF.Globals().SetItem(πF, ß_decode_uXXXX.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 78: def py_scanstring(s, end, encoding=None, strict=True,
			πF.SetLineno(78)
			πTemp007 = make([]πg.Param, 6)
			πTemp007[0] = πg.Param{Name: "s", Def: nil}
			πTemp007[1] = πg.Param{Name: "end", Def: nil}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp007[2] = πg.Param{Name: "encoding", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp007[3] = πg.Param{Name: "strict", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßBACKSLASH); πE != nil {
				continue
			}
			πTemp007[4] = πg.Param{Name: "_b", Def: πTemp009}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßSTRINGCHUNK); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßmatch, nil); πE != nil {
				continue
			}
			πTemp007[5] = πg.Param{Name: "_m", Def: πTemp010}
			πTemp006 = πg.NewFunction(πg.NewCode("py_scanstring", "/usr/lib/python2.7/json/decoder.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µend *πg.Object = πArgs[1]; _ = µend
				var µencoding *πg.Object = πArgs[2]; _ = µencoding
				var µstrict *πg.Object = πArgs[3]; _ = µstrict
				var µ_b *πg.Object = πArgs[4]; _ = µ_b
				var µ_m *πg.Object = πArgs[5]; _ = µ_m
				var µchunks *πg.Object = πg.UnboundLocal; _ = µchunks
				var µ_append *πg.Object = πg.UnboundLocal; _ = µ_append
				var µbegin *πg.Object = πg.UnboundLocal; _ = µbegin
				var µchunk *πg.Object = πg.UnboundLocal; _ = µchunk
				var µcontent *πg.Object = πg.UnboundLocal; _ = µcontent
				var µterminator *πg.Object = πg.UnboundLocal; _ = µterminator
				var µmsg *πg.Object = πg.UnboundLocal; _ = µmsg
				var µesc *πg.Object = πg.UnboundLocal; _ = µesc
				var µchar *πg.Object = πg.UnboundLocal; _ = µchar
				var µuni *πg.Object = πg.UnboundLocal; _ = µuni
				var µuni2 *πg.Object = πg.UnboundLocal; _ = µuni2
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 *πg.BaseException
				_ = πTemp008
				var πTemp009 *πg.Traceback
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 19: goto Label19
					case 25: goto Label25
					case 3: goto Label3
					case 4: goto Label4
					default: panic("unexpected function state")
					}
					// line 80: """Scan the string s for a JSON string. End is the index of the
					πF.SetLineno(80)
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µencoding == πTemp002).ToObject()
					if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp003 {
						goto Label1
					}
					goto Label2
					// line 88: if encoding is None:
					πF.SetLineno(88)
				Label1:
					// line 89: encoding = DEFAULT_ENCODING
					πF.SetLineno(89)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßDEFAULT_ENCODING); πE != nil {
						continue
					}
					µencoding = πTemp001
					goto Label2
				Label2:
					// line 90: chunks = []
					πF.SetLineno(90)
					πTemp004 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp004...).ToObject()
					µchunks = πTemp001
					// line 91: _append = chunks.append
					πF.SetLineno(91)
					if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µchunks, ßappend, nil); πE != nil {
						continue
					}
					µ_append = πTemp001
					// line 92: begin = end - 1
					πF.SetLineno(92)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µbegin = πTemp001
					// line 93: while 1:
					πF.SetLineno(93)
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
					if πTemp005, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(3)            
					// line 94: chunk = _m(s, end)
					πF.SetLineno(94)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp004[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp004[1] = µend
					if πE = πg.CheckLocal(πF, µ_m, "_m"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_m.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µchunk = πTemp001
					if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µchunk == πTemp002).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label6
					}
					goto Label7
					// line 95: if chunk is None:
					πF.SetLineno(95)
				Label6:
					πTemp004 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					πTemp006[0] = πg.NewStr("Unterminated string starting at").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[1] = µs
					if πE = πg.CheckLocal(πF, µbegin, "begin"); πE != nil {
						continue
					}
					πTemp006[2] = µbegin
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 96: raise ValueError(
					πF.SetLineno(96)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label7
				Label7:
					// line 98: end = chunk.end()
					πF.SetLineno(98)
					if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µchunk, ßend, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend = πTemp002
					// line 99: content, terminator = chunk.groups()
					πF.SetLineno(99)
					if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µchunk, ßgroups, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp007}}}, πTemp002); πE != nil {
						continue
					}
					µcontent = πTemp001
					µterminator = πTemp007
					if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µcontent); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label8
					}
					goto Label9
					// line 101: if content:
					πF.SetLineno(101)
				Label8:
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
						continue
					}
					πTemp004[0] = µcontent
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					πTemp004[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp005, πE = πg.IsTrue(πF, πTemp007); πE != nil {
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
					// line 102: if not isinstance(content, unicode):
					πF.SetLineno(102)
				Label10:
					// line 103: content = unicode(content, encoding)
					πF.SetLineno(103)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
						continue
					}
					πTemp004[0] = µcontent
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp004[1] = µencoding
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µcontent = πTemp002
					goto Label11
				Label11:
					// line 104: _append(content)
					πF.SetLineno(104)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µcontent, "content"); πE != nil {
						continue
					}
					πTemp004[0] = µcontent
					if πE = πg.CheckLocal(πF, µ_append, "_append"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_append.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µterminator, "terminator"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µterminator, πg.NewStr("\"").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µterminator, "terminator"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µterminator, πg.NewStr("\\").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					goto Label14
					// line 107: if terminator == '"':
					πF.SetLineno(107)
				Label12:
					// line 108: break
					πF.SetLineno(108)
					πTemp003 = true
					continue
					goto Label14
					// line 109: elif terminator != '\\':
					πF.SetLineno(109)
				Label13:
					if πE = πg.CheckLocal(πF, µstrict, "strict"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µstrict); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label15
					}
					goto Label16
					// line 110: if strict:
					πF.SetLineno(110)
				Label15:
					// line 111: msg = "Invalid control character %r at" % (terminator,)
					πF.SetLineno(111)
					if πE = πg.CheckLocal(πF, µterminator, "terminator"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple1(µterminator).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("Invalid control character %r at").ToObject(), πTemp002); πE != nil {
						continue
					}
					µmsg = πTemp001
					πTemp004 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp006[0] = µmsg
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[1] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp006[2] = µend
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 113: raise ValueError(errmsg(msg, s, end))
					πF.SetLineno(113)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label17
				Label16:
					// line 115: _append(terminator)
					πF.SetLineno(115)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µterminator, "terminator"); πE != nil {
						continue
					}
					πTemp004[0] = µterminator
					if πE = πg.CheckLocal(πF, µ_append, "_append"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_append.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 116: continue
					πF.SetLineno(116)
					continue
					goto Label17
				Label17:
					goto Label14
				Label14:
					// line 117: try:
					πF.SetLineno(117)
					πF.PushCheckpoint(19)
					// line 118: esc = s[end]
					πF.SetLineno(118)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µesc = πTemp002
					πF.PopCheckpoint()
					goto Label18
				Label19:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label20
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 119: except IndexError:
					πF.SetLineno(119)
				Label20:
					πTemp004 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					πTemp006[0] = πg.NewStr("Unterminated string starting at").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[1] = µs
					if πE = πg.CheckLocal(πF, µbegin, "begin"); πE != nil {
						continue
					}
					πTemp006[2] = µbegin
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 120: raise ValueError(
					πF.SetLineno(120)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label18
				Label18:
					if πE = πg.CheckLocal(πF, µesc, "esc"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µesc, ßu.ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label21
					}
					goto Label22
					// line 123: if esc != 'u':
					πF.SetLineno(123)
				Label21:
					// line 124: try:
					πF.SetLineno(124)
					πF.PushCheckpoint(25)
					// line 125: char = _b[esc]
					πF.SetLineno(125)
					if πE = πg.CheckLocal(πF, µesc, "esc"); πE != nil {
						continue
					}
					πTemp001 = µesc
					if πE = πg.CheckLocal(πF, µ_b, "_b"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µ_b, πTemp001); πE != nil {
						continue
					}
					µchar = πTemp002
					πF.PopCheckpoint()
					goto Label24
				Label25:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp008, πTemp009 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label26
					}
					πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
					continue
					// line 126: except KeyError:
					πF.SetLineno(126)
				Label26:
					// line 127: msg = "Invalid \\escape: " + repr(esc)
					πF.SetLineno(127)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µesc, "esc"); πE != nil {
						continue
					}
					πTemp004[0] = µesc
					if πTemp002, πE = πg.ResolveGlobal(πF, ßrepr); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πTemp001, πE = πg.Add(πF, πg.NewStr("Invalid \\escape: ").ToObject(), πTemp007); πE != nil {
						continue
					}
					µmsg = πTemp001
					πTemp004 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µmsg, "msg"); πE != nil {
						continue
					}
					πTemp006[0] = µmsg
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[1] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp006[2] = µend
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					// line 128: raise ValueError(errmsg(msg, s, end))
					πF.SetLineno(128)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label24
				Label24:
					// line 129: end += 1
					πF.SetLineno(129)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					goto Label23
				Label22:
					// line 132: uni = _decode_uXXXX(s, end)
					πF.SetLineno(132)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp004[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp004[1] = µend
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_decode_uXXXX); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µuni = πTemp002
					// line 133: end += 5
					πF.SetLineno(133)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(5).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					if πTemp007, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp007, ßmaxunicode, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, πTemp010, πg.NewInt(65535).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label27
					}
					if πE = πg.CheckLocal(πF, µuni, "uni"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.LE(πF, πg.NewInt(55296).ToObject(), µuni); πE != nil {
						continue
					}
					if πTemp011, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp011 {
						goto Label28
					}
					if πTemp002, πE = πg.LE(πF, µuni, πg.NewInt(56319).ToObject()); πE != nil {
						continue
					}
				Label28:
					πTemp001 = πTemp002
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label27
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Add(πF, µend, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.SliceType.Call(πF, πg.Args{µend, πTemp010, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µs, πTemp007); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp010, πg.NewStr("\\u").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label27:
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label29
					}
					goto Label30
					// line 135: if sys.maxunicode > 65535 and \
					πF.SetLineno(135)
				Label29:
					// line 137: uni2 = _decode_uXXXX(s, end + 1)
					πF.SetLineno(137)
					πTemp004 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp004[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004[1] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_decode_uXXXX); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µuni2 = πTemp002
					if πE = πg.CheckLocal(πF, µuni2, "uni2"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, πg.NewInt(56320).ToObject(), µuni2); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp005 {
						goto Label31
					}
					if πTemp001, πE = πg.LE(πF, µuni2, πg.NewInt(57343).ToObject()); πE != nil {
						continue
					}
				Label31:
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label32
					}
					goto Label33
					// line 138: if 0xdc00 <= uni2 <= 0xdfff:
					πF.SetLineno(138)
				Label32:
					// line 139: uni = 0x10000 + (((uni - 0xd800) << 10) | (uni2 - 0xdc00))
					πF.SetLineno(139)
					if πE = πg.CheckLocal(πF, µuni, "uni"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Sub(πF, µuni, πg.NewInt(55296).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.LShift(πF, πTemp010, πg.NewInt(10).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µuni2, "uni2"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Sub(πF, µuni2, πg.NewInt(56320).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Or(πF, πTemp007, πTemp010); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewInt(65536).ToObject(), πTemp002); πE != nil {
						continue
					}
					µuni = πTemp001
					// line 140: end += 6
					πF.SetLineno(140)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(6).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					goto Label33
				Label33:
					goto Label30
				Label30:
					// line 141: char = unichr(uni)
					πF.SetLineno(141)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µuni, "uni"); πE != nil {
						continue
					}
					πTemp004[0] = µuni
					if πTemp001, πE = πg.ResolveGlobal(πF, ßunichr); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					µchar = πTemp002
					goto Label23
				Label23:
					// line 143: _append(char)
					πF.SetLineno(143)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchar, "char"); πE != nil {
						continue
					}
					πTemp004[0] = µchar
					if πE = πg.CheckLocal(πF, µ_append, "_append"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_append.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					continue
				Label4:
					if πE != nil || πR != nil {
						continue
					}
				Label5:
					// line 144: return u''.join(chunks), end
					πF.SetLineno(144)
					πTemp004 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchunks, "chunks"); πE != nil {
						continue
					}
					πTemp004[0] = µchunks
					if πTemp002, πE = πg.GetAttr(πF, πg.NewUnicode("").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp007, µend).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßpy_scanstring.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 80: """Scan the string s for a JSON string. End is the index of the
			πF.SetLineno(80)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Scan the string s for a JSON string. End is the index of the\n    character in s after the quote that started the JSON string.\n    Unescapes all valid JSON string escape sequences and raises ValueError\n    on attempt to decode an invalid string. If strict is False then literal\n    control characters are allowed in the string.\n\n    Returns a tuple of the decoded string and the index of the character in s\n    after the end quote.").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßpy_scanstring); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 148: scanstring = c_scanstring or py_scanstring
			πF.SetLineno(148)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßc_scanstring); πE != nil {
				continue
			}
			πTemp009 = πTemp010
			if πTemp011, πE = πg.IsTrue(πF, πTemp009); πE != nil {
				continue
			}
			if πTemp011 {
				goto Label1
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßpy_scanstring); πE != nil {
				continue
			}
			πTemp009 = πTemp010
		Label1:
			if πE = πF.Globals().SetItem(πF, ßscanstring.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 150: WHITESPACE = re.compile(r'[ \t\n\r]*', FLAGS)
			πF.SetLineno(150)
			πTemp002 = πF.MakeArgs(2)
			πTemp002[0] = πg.NewStr("[ \\t\\n\\r]*").ToObject()
			if πTemp009, πE = πg.ResolveGlobal(πF, ßFLAGS); πE != nil {
				continue
			}
			πTemp002[1] = πTemp009
			if πTemp009, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp009, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ßWHITESPACE.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 151: WHITESPACE_STR = ' \t\n\r'
			πF.SetLineno(151)
			if πE = πF.Globals().SetItem(πF, ßWHITESPACE_STR.ToObject(), πg.NewStr(" \t\n\r").ToObject()); πE != nil {
				continue
			}
			// line 153: def JSONObject(s_and_end, encoding, strict, scan_once, object_hook,
			πF.SetLineno(153)
			πTemp007 = make([]πg.Param, 8)
			πTemp007[0] = πg.Param{Name: "s_and_end", Def: nil}
			πTemp007[1] = πg.Param{Name: "encoding", Def: nil}
			πTemp007[2] = πg.Param{Name: "strict", Def: nil}
			πTemp007[3] = πg.Param{Name: "scan_once", Def: nil}
			πTemp007[4] = πg.Param{Name: "object_hook", Def: nil}
			πTemp007[5] = πg.Param{Name: "object_pairs_hook", Def: nil}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßWHITESPACE); πE != nil {
				continue
			}
			if πTemp012, πE = πg.GetAttr(πF, πTemp010, ßmatch, nil); πE != nil {
				continue
			}
			πTemp007[6] = πg.Param{Name: "_w", Def: πTemp012}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßWHITESPACE_STR); πE != nil {
				continue
			}
			πTemp007[7] = πg.Param{Name: "_ws", Def: πTemp010}
			πTemp009 = πg.NewFunction(πg.NewCode("JSONObject", "/usr/lib/python2.7/json/decoder.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs_and_end *πg.Object = πArgs[0]; _ = µs_and_end
				var µencoding *πg.Object = πArgs[1]; _ = µencoding
				var µstrict *πg.Object = πArgs[2]; _ = µstrict
				var µscan_once *πg.Object = πArgs[3]; _ = µscan_once
				var µobject_hook *πg.Object = πArgs[4]; _ = µobject_hook
				var µobject_pairs_hook *πg.Object = πArgs[5]; _ = µobject_pairs_hook
				var µ_w *πg.Object = πArgs[6]; _ = µ_w
				var µ_ws *πg.Object = πArgs[7]; _ = µ_ws
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µend *πg.Object = πg.UnboundLocal; _ = µend
				var µpairs *πg.Object = πg.UnboundLocal; _ = µpairs
				var µpairs_append *πg.Object = πg.UnboundLocal; _ = µpairs_append
				var µnextchar *πg.Object = πg.UnboundLocal; _ = µnextchar
				var µresult *πg.Object = πg.UnboundLocal; _ = µresult
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 bool
				_ = πTemp004
				var πTemp005 *πg.Dict
				_ = πTemp005
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.BaseException
				_ = πTemp009
				var πTemp010 *πg.Traceback
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 38: goto Label38
					case 12: goto Label12
					case 13: goto Label13
					case 20: goto Label20
					case 27: goto Label27
					case 30: goto Label30
					default: panic("unexpected function state")
					}
					// line 155: s, end = s_and_end
					πF.SetLineno(155)
					if πE = πg.CheckLocal(πF, µs_and_end, "s_and_end"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µs_and_end); πE != nil {
						continue
					}
					µs = πTemp001
					µend = πTemp002
					// line 156: pairs = []
					πF.SetLineno(156)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µpairs = πTemp001
					// line 157: pairs_append = pairs.append
					πF.SetLineno(157)
					if πE = πg.CheckLocal(πF, µpairs, "pairs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µpairs, ßappend, nil); πE != nil {
						continue
					}
					µpairs_append = πTemp001
					// line 160: nextchar = s[end:end + 1]
					πF.SetLineno(160)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µend, πTemp002, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µnextchar, πg.NewStr("\"").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 162: if nextchar != '"':
					πF.SetLineno(162)
				Label1:
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µ_ws, µnextchar); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 163: if nextchar in _ws:
					πF.SetLineno(163)
				Label3:
					// line 164: end = _w(s, end).end()
					πF.SetLineno(164)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp003[1] = µend
					if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_w.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend = πTemp001
					// line 165: nextchar = s[end:end + 1]
					πF.SetLineno(165)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µend, πTemp002, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µnextchar, πg.NewStr("}").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µnextchar, πg.NewStr("\"").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label6
					}
					goto Label7
					// line 167: if nextchar == '}':
					πF.SetLineno(167)
				Label5:
					if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µobject_pairs_hook != πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					goto Label9
					// line 168: if object_pairs_hook is not None:
					πF.SetLineno(168)
				Label8:
					// line 169: result = object_pairs_hook(pairs)
					πF.SetLineno(169)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpairs, "pairs"); πE != nil {
						continue
					}
					πTemp003[0] = µpairs
					if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
						continue
					}
					if πTemp001, πE = µobject_pairs_hook.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp001
					// line 170: return result, end + 1
					πF.SetLineno(170)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µresult, πTemp002).ToObject()
					πR = πTemp001
					continue
					goto Label9
				Label9:
					// line 171: pairs = {}
					πF.SetLineno(171)
					πTemp005 = πg.NewDict()
					πTemp001 = πTemp005.ToObject()
					µpairs = πTemp001
					if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µobject_hook != πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label10
					}
					goto Label11
					// line 172: if object_hook is not None:
					πF.SetLineno(172)
				Label10:
					// line 173: pairs = object_hook(pairs)
					πF.SetLineno(173)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpairs, "pairs"); πE != nil {
						continue
					}
					πTemp003[0] = µpairs
					if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
						continue
					}
					if πTemp001, πE = µobject_hook.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µpairs = πTemp001
					goto Label11
				Label11:
					// line 174: return pairs, end + 1
					πF.SetLineno(174)
					if πE = πg.CheckLocal(πF, µpairs, "pairs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µpairs, πTemp002).ToObject()
					πR = πTemp001
					continue
					goto Label7
					// line 175: elif nextchar != '"':
					πF.SetLineno(175)
				Label6:
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					πTemp006[0] = πg.NewStr("Expecting property name enclosed in double quotes").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[1] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp006[2] = µend
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 176: raise ValueError(errmsg(
					πF.SetLineno(176)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label7
				Label7:
					goto Label2
				Label2:
					// line 178: end += 1
					πF.SetLineno(178)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					// line 179: while True:
					πF.SetLineno(179)
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
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(12)            
					// line 180: key, end = scanstring(s, end, encoding, strict)
					πF.SetLineno(180)
					πTemp003 = πF.MakeArgs(4)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp003[1] = µend
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp003[2] = µencoding
					if πE = πg.CheckLocal(πF, µstrict, "strict"); πE != nil {
						continue
					}
					πTemp003[3] = µstrict
					if πTemp001, πE = πg.ResolveGlobal(πF, ßscanstring); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp008}}}, πTemp002); πE != nil {
						continue
					}
					µkey = πTemp001
					µend = πTemp008
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µend, πTemp008, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp008, πg.NewStr(":").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label15
					}
					goto Label16
					// line 184: if s[end:end + 1] != ':':
					πF.SetLineno(184)
				Label15:
					// line 185: end = _w(s, end).end()
					πF.SetLineno(185)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp003[1] = µend
					if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_w.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend = πTemp001
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µend, πTemp008, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, πTemp008, πg.NewStr(":").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label17
					}
					goto Label18
					// line 186: if s[end:end + 1] != ':':
					πF.SetLineno(186)
				Label17:
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					πTemp006[0] = πg.NewStr("Expecting ':' delimiter").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[1] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp006[2] = µend
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 187: raise ValueError(errmsg("Expecting ':' delimiter", s, end))
					πF.SetLineno(187)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label18
				Label18:
					goto Label16
				Label16:
					// line 188: end += 1
					πF.SetLineno(188)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					// line 190: try:
					πF.SetLineno(190)
					πF.PushCheckpoint(20)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp002 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µ_ws, πTemp008); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label21
					}
					goto Label22
					// line 191: if s[end] in _ws:
					πF.SetLineno(191)
				Label21:
					// line 192: end += 1
					πF.SetLineno(192)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp002 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µ_ws, πTemp008); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label23
					}
					goto Label24
					// line 193: if s[end] in _ws:
					πF.SetLineno(193)
				Label23:
					// line 194: end = _w(s, end + 1).end()
					πF.SetLineno(194)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_w.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend = πTemp001
					goto Label24
				Label24:
					goto Label22
				Label22:
					πF.PopCheckpoint()
					goto Label19
				Label20:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label25
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 195: except IndexError:
					πF.SetLineno(195)
				Label25:
					// line 196: pass
					πF.SetLineno(196)
					πF.RestoreExc(nil, nil)
					goto Label19
				Label19:
					// line 198: try:
					πF.SetLineno(198)
					πF.PushCheckpoint(27)
					// line 199: value, end = scan_once(s, end)
					πF.SetLineno(199)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp003[1] = µend
					if πE = πg.CheckLocal(πF, µscan_once, "scan_once"); πE != nil {
						continue
					}
					if πTemp001, πE = µscan_once.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp008}}}, πTemp001); πE != nil {
						continue
					}
					µvalue = πTemp002
					µend = πTemp008
					πF.PopCheckpoint()
					goto Label26
				Label27:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label28
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 200: except StopIteration:
					πF.SetLineno(200)
				Label28:
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					πTemp006[0] = πg.NewStr("Expecting object").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[1] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp006[2] = µend
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 201: raise ValueError(errmsg("Expecting object", s, end))
					πF.SetLineno(201)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label26
				Label26:
					// line 202: pairs_append((key, value))
					πF.SetLineno(202)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µkey, µvalue).ToObject()
					πTemp003[0] = πTemp001
					if πE = πg.CheckLocal(πF, µpairs_append, "pairs_append"); πE != nil {
						continue
					}
					if πTemp001, πE = µpairs_append.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 204: try:
					πF.SetLineno(204)
					πF.PushCheckpoint(30)
					// line 205: nextchar = s[end]
					πF.SetLineno(205)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µ_ws, µnextchar); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label31
					}
					goto Label32
					// line 206: if nextchar in _ws:
					πF.SetLineno(206)
				Label31:
					// line 207: end = _w(s, end + 1).end()
					πF.SetLineno(207)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_w.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend = πTemp001
					// line 208: nextchar = s[end]
					πF.SetLineno(208)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					goto Label32
				Label32:
					πF.PopCheckpoint()
					goto Label29
				Label30:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label33
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 209: except IndexError:
					πF.SetLineno(209)
				Label33:
					// line 210: nextchar = ''
					πF.SetLineno(210)
					µnextchar = ß.ToObject()
					πF.RestoreExc(nil, nil)
					goto Label29
				Label29:
					// line 211: end += 1
					πF.SetLineno(211)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µnextchar, πg.NewStr("}").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label34
					}
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µnextchar, πg.NewStr(",").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label35
					}
					goto Label36
					// line 213: if nextchar == '}':
					πF.SetLineno(213)
				Label34:
					// line 214: break
					πF.SetLineno(214)
					πTemp004 = true
					continue
					goto Label36
					// line 215: elif nextchar != ',':
					πF.SetLineno(215)
				Label35:
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					πTemp006[0] = πg.NewStr("Expecting ',' delimiter").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[1] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 216: raise ValueError(errmsg("Expecting ',' delimiter", s, end - 1))
					πF.SetLineno(216)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label36
				Label36:
					// line 218: try:
					πF.SetLineno(218)
					πF.PushCheckpoint(38)
					// line 219: nextchar = s[end]
					πF.SetLineno(219)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µ_ws, µnextchar); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label39
					}
					goto Label40
					// line 220: if nextchar in _ws:
					πF.SetLineno(220)
				Label39:
					// line 221: end += 1
					πF.SetLineno(221)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					// line 222: nextchar = s[end]
					πF.SetLineno(222)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µ_ws, µnextchar); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label41
					}
					goto Label42
					// line 223: if nextchar in _ws:
					πF.SetLineno(223)
				Label41:
					// line 224: end = _w(s, end + 1).end()
					πF.SetLineno(224)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_w.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend = πTemp001
					// line 225: nextchar = s[end]
					πF.SetLineno(225)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					goto Label42
				Label42:
					goto Label40
				Label40:
					πF.PopCheckpoint()
					goto Label37
				Label38:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp009, πTemp010 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp009.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label43
					}
					πE = πF.Raise(πTemp009.ToObject(), nil, πTemp010.ToObject())
					continue
					// line 226: except IndexError:
					πF.SetLineno(226)
				Label43:
					// line 227: nextchar = ''
					πF.SetLineno(227)
					µnextchar = ß.ToObject()
					πF.RestoreExc(nil, nil)
					goto Label37
				Label37:
					// line 229: end += 1
					πF.SetLineno(229)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µnextchar, πg.NewStr("\"").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label44
					}
					goto Label45
					// line 230: if nextchar != '"':
					πF.SetLineno(230)
				Label44:
					πTemp003 = πF.MakeArgs(1)
					πTemp006 = πF.MakeArgs(3)
					πTemp006[0] = πg.NewStr("Expecting property name enclosed in double quotes").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[1] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Sub(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp006[2] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp003[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 231: raise ValueError(errmsg(
					πF.SetLineno(231)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label45
				Label45:
					continue
				Label13:
					if πE != nil || πR != nil {
						continue
					}
				Label14:
					if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µobject_pairs_hook != πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label46
					}
					goto Label47
					// line 233: if object_pairs_hook is not None:
					πF.SetLineno(233)
				Label46:
					// line 234: result = object_pairs_hook(pairs)
					πF.SetLineno(234)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpairs, "pairs"); πE != nil {
						continue
					}
					πTemp003[0] = µpairs
					if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
						continue
					}
					if πTemp001, πE = µobject_pairs_hook.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µresult = πTemp001
					// line 235: return result, end
					πF.SetLineno(235)
					if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µresult, µend).ToObject()
					πR = πTemp001
					continue
					goto Label47
				Label47:
					// line 236: pairs = dict(pairs)
					πF.SetLineno(236)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpairs, "pairs"); πE != nil {
						continue
					}
					πTemp003[0] = µpairs
					if πTemp001, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µpairs = πTemp002
					if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µobject_hook != πTemp002).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label48
					}
					goto Label49
					// line 237: if object_hook is not None:
					πF.SetLineno(237)
				Label48:
					// line 238: pairs = object_hook(pairs)
					πF.SetLineno(238)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µpairs, "pairs"); πE != nil {
						continue
					}
					πTemp003[0] = µpairs
					if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
						continue
					}
					if πTemp001, πE = µobject_hook.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µpairs = πTemp001
					goto Label49
				Label49:
					// line 239: return pairs, end
					πF.SetLineno(239)
					if πE = πg.CheckLocal(πF, µpairs, "pairs"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µpairs, µend).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßJSONObject.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 241: def JSONArray(s_and_end, scan_once, _w=WHITESPACE.match, _ws=WHITESPACE_STR):
			πF.SetLineno(241)
			πTemp007 = make([]πg.Param, 4)
			πTemp007[0] = πg.Param{Name: "s_and_end", Def: nil}
			πTemp007[1] = πg.Param{Name: "scan_once", Def: nil}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßWHITESPACE); πE != nil {
				continue
			}
			if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßmatch, nil); πE != nil {
				continue
			}
			πTemp007[2] = πg.Param{Name: "_w", Def: πTemp013}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßWHITESPACE_STR); πE != nil {
				continue
			}
			πTemp007[3] = πg.Param{Name: "_ws", Def: πTemp012}
			πTemp010 = πg.NewFunction(πg.NewCode("JSONArray", "/usr/lib/python2.7/json/decoder.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs_and_end *πg.Object = πArgs[0]; _ = µs_and_end
				var µscan_once *πg.Object = πArgs[1]; _ = µscan_once
				var µ_w *πg.Object = πArgs[2]; _ = µ_w
				var µ_ws *πg.Object = πArgs[3]; _ = µ_ws
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µend *πg.Object = πg.UnboundLocal; _ = µend
				var µvalues *πg.Object = πg.UnboundLocal; _ = µvalues
				var µnextchar *πg.Object = πg.UnboundLocal; _ = µnextchar
				var µ_append *πg.Object = πg.UnboundLocal; _ = µ_append
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
				var πTemp007 *πg.BaseException
				_ = πTemp007
				var πTemp008 *πg.Traceback
				_ = πTemp008
				var πTemp009 []*πg.Object
				_ = πTemp009
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 9: goto Label9
					case 5: goto Label5
					case 6: goto Label6
					case 17: goto Label17
					default: panic("unexpected function state")
					}
					// line 242: s, end = s_and_end
					πF.SetLineno(242)
					if πE = πg.CheckLocal(πF, µs_and_end, "s_and_end"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}}}, µs_and_end); πE != nil {
						continue
					}
					µs = πTemp001
					µend = πTemp002
					// line 243: values = []
					πF.SetLineno(243)
					πTemp003 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					µvalues = πTemp001
					// line 244: nextchar = s[end:end + 1]
					πF.SetLineno(244)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µend, πTemp002, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µ_ws, µnextchar); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp004).ToObject()
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label1
					}
					goto Label2
					// line 245: if nextchar in _ws:
					πF.SetLineno(245)
				Label1:
					// line 246: end = _w(s, end + 1).end()
					πF.SetLineno(246)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_w.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend = πTemp001
					// line 247: nextchar = s[end:end + 1]
					πF.SetLineno(247)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µend, πTemp002, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µnextchar, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label3
					}
					goto Label4
					// line 249: if nextchar == ']':
					πF.SetLineno(249)
				Label3:
					// line 250: return values, end + 1
					πF.SetLineno(250)
					if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µvalues, πTemp002).ToObject()
					πR = πTemp001
					continue
					goto Label4
				Label4:
					// line 251: _append = values.append
					πF.SetLineno(251)
					if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µvalues, ßappend, nil); πE != nil {
						continue
					}
					µ_append = πTemp001
					// line 252: while True:
					πF.SetLineno(252)
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
					if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp005 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 253: try:
					πF.SetLineno(253)
					πF.PushCheckpoint(9)
					// line 254: value, end = scan_once(s, end)
					πF.SetLineno(254)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp003[1] = µend
					if πE = πg.CheckLocal(πF, µscan_once, "scan_once"); πE != nil {
						continue
					}
					if πTemp001, πE = µscan_once.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp006}}}, πTemp001); πE != nil {
						continue
					}
					µvalue = πTemp002
					µend = πTemp006
					πF.PopCheckpoint()
					goto Label8
				Label9:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label10
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 255: except StopIteration:
					πF.SetLineno(255)
				Label10:
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(3)
					πTemp009[0] = πg.NewStr("Expecting object").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp009[1] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp009[2] = µend
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 256: raise ValueError(errmsg("Expecting object", s, end))
					πF.SetLineno(256)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					πF.RestoreExc(nil, nil)
					goto Label8
				Label8:
					// line 257: _append(value)
					πF.SetLineno(257)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp003[0] = µvalue
					if πE = πg.CheckLocal(πF, µ_append, "_append"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_append.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 258: nextchar = s[end:end + 1]
					πF.SetLineno(258)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µend, πTemp002, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µ_ws, µnextchar); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label11
					}
					goto Label12
					// line 259: if nextchar in _ws:
					πF.SetLineno(259)
				Label11:
					// line 260: end = _w(s, end + 1).end()
					πF.SetLineno(260)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_w.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend = πTemp001
					// line 261: nextchar = s[end:end + 1]
					πF.SetLineno(261)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{µend, πTemp002, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetItem(πF, µs, πTemp001); πE != nil {
						continue
					}
					µnextchar = πTemp002
					goto Label12
				Label12:
					// line 262: end += 1
					πF.SetLineno(262)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, µnextchar, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µnextchar, "nextchar"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.NE(πF, µnextchar, πg.NewStr(",").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label14
					}
					goto Label15
					// line 263: if nextchar == ']':
					πF.SetLineno(263)
				Label13:
					// line 264: break
					πF.SetLineno(264)
					πTemp004 = true
					continue
					goto Label15
					// line 265: elif nextchar != ',':
					πF.SetLineno(265)
				Label14:
					πTemp003 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(3)
					πTemp009[0] = πg.NewStr("Expecting ',' delimiter").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp009[1] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp009[2] = µend
					if πTemp001, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp003[0] = πTemp002
					if πTemp001, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					// line 266: raise ValueError(errmsg("Expecting ',' delimiter", s, end))
					πF.SetLineno(266)
					πE = πF.Raise(πTemp002, nil, nil)
					continue
					goto Label15
				Label15:
					// line 267: try:
					πF.SetLineno(267)
					πF.PushCheckpoint(17)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp002 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µ_ws, πTemp006); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label18
					}
					goto Label19
					// line 268: if s[end] in _ws:
					πF.SetLineno(268)
				Label18:
					// line 269: end += 1
					πF.SetLineno(269)
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.IAdd(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µend = πTemp001
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp002 = µend
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µs, πTemp002); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µ_ws, "_ws"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Contains(πF, µ_ws, πTemp006); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp005).ToObject()
					if πTemp005, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label20
					}
					goto Label21
					// line 270: if s[end] in _ws:
					πF.SetLineno(270)
				Label20:
					// line 271: end = _w(s, end + 1).end()
					πF.SetLineno(271)
					πTemp003 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp003[0] = µs
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µend, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[1] = πTemp001
					if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
						continue
					}
					if πTemp001, πE = µ_w.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßend, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend = πTemp001
					goto Label21
				Label21:
					goto Label19
				Label19:
					πF.PopCheckpoint()
					goto Label16
				Label17:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp007, πTemp008 = πF.ExcInfo()
					if πTemp001, πE = πg.ResolveGlobal(πF, ßIndexError); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsInstance(πF, πTemp007.ToObject(), πTemp001); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label22
					}
					πE = πF.Raise(πTemp007.ToObject(), nil, πTemp008.ToObject())
					continue
					// line 272: except IndexError:
					πF.SetLineno(272)
				Label22:
					// line 273: pass
					πF.SetLineno(273)
					πF.RestoreExc(nil, nil)
					goto Label16
				Label16:
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 275: return values, end
					πF.SetLineno(275)
					if πE = πg.CheckLocal(πF, µvalues, "values"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µvalues, µend).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßJSONArray.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 277: class JSONDecoder(object):
			πF.SetLineno(277)
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
			_, πE = πg.NewCode("JSONDecoder", "/usr/lib/python2.7/json/decoder.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp008
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
					// line 278: """Simple JSON <http://json.org> decoder
					πF.SetLineno(278)
					// line 278: """Simple JSON <http://json.org> decoder
					πF.SetLineno(278)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Simple JSON <http://json.org> decoder\n\n    Performs the following translations in decoding by default:\n\n    +---------------+-------------------+\n    | JSON          | Python            |\n    +===============+===================+\n    | object        | dict              |\n    +---------------+-------------------+\n    | array         | list              |\n    +---------------+-------------------+\n    | string        | unicode           |\n    +---------------+-------------------+\n    | number (int)  | int, long         |\n    +---------------+-------------------+\n    | number (real) | float             |\n    +---------------+-------------------+\n    | true          | True              |\n    +---------------+-------------------+\n    | false         | False             |\n    +---------------+-------------------+\n    | null          | None              |\n    +---------------+-------------------+\n\n    It also understands ``NaN``, ``Infinity``, and ``-Infinity`` as\n    their corresponding ``float`` values, which is outside the JSON spec.\n\n    ").ToObject()); πE != nil {
						continue
					}
					// line 307: def __init__(self, encoding=None, object_hook=None, parse_float=None,
					πF.SetLineno(307)
					πTemp002 = make([]πg.Param, 8)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[1] = πg.Param{Name: "encoding", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "object_hook", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "parse_float", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[4] = πg.Param{Name: "parse_int", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[5] = πg.Param{Name: "parse_constant", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
						continue
					}
					πTemp002[6] = πg.Param{Name: "strict", Def: πTemp003}
					if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßNone); πE != nil {
						continue
					}
					πTemp002[7] = πg.Param{Name: "object_pairs_hook", Def: πTemp003}
					πTemp001 = πg.NewFunction(πg.NewCode("__init__", "/usr/lib/python2.7/json/decoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µencoding *πg.Object = πArgs[1]; _ = µencoding
						var µobject_hook *πg.Object = πArgs[2]; _ = µobject_hook
						var µparse_float *πg.Object = πArgs[3]; _ = µparse_float
						var µparse_int *πg.Object = πArgs[4]; _ = µparse_int
						var µparse_constant *πg.Object = πArgs[5]; _ = µparse_constant
						var µstrict *πg.Object = πArgs[6]; _ = µstrict
						var µobject_pairs_hook *πg.Object = πArgs[7]; _ = µobject_pairs_hook
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
							// line 310: """``encoding`` determines the encoding used to interpret any ``str``
							πF.SetLineno(310)
							// line 352: self.encoding = encoding
							πF.SetLineno(352)
							if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µencoding); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßencoding, πTemp001); πE != nil {
								continue
							}
							// line 353: self.object_hook = object_hook
							πF.SetLineno(353)
							if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µobject_hook); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßobject_hook, πTemp001); πE != nil {
								continue
							}
							// line 354: self.object_pairs_hook = object_pairs_hook
							πF.SetLineno(354)
							if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µobject_pairs_hook); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßobject_pairs_hook, πTemp001); πE != nil {
								continue
							}
							// line 355: self.parse_float = parse_float or float
							πF.SetLineno(355)
							if πE = πg.CheckLocal(πF, µparse_float, "parse_float"); πE != nil {
								continue
							}
							πTemp001 = µparse_float
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label1
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label1:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparse_float, πTemp003); πE != nil {
								continue
							}
							// line 356: self.parse_int = parse_int or int
							πF.SetLineno(356)
							if πE = πg.CheckLocal(πF, µparse_int, "parse_int"); πE != nil {
								continue
							}
							πTemp001 = µparse_int
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label2
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							πTemp001 = πTemp003
						Label2:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparse_int, πTemp003); πE != nil {
								continue
							}
							// line 357: self.parse_constant = parse_constant or _CONSTANTS.__getitem__
							πF.SetLineno(357)
							if πE = πg.CheckLocal(πF, µparse_constant, "parse_constant"); πE != nil {
								continue
							}
							πTemp001 = µparse_constant
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							if πTemp003, πE = πg.ResolveGlobal(πF, ß_CONSTANTS); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ß__getitem__, nil); πE != nil {
								continue
							}
							πTemp001 = πTemp004
						Label3:
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparse_constant, πTemp003); πE != nil {
								continue
							}
							// line 358: self.strict = strict
							πF.SetLineno(358)
							if πE = πg.CheckLocal(πF, µstrict, "strict"); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µstrict); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßstrict, πTemp001); πE != nil {
								continue
							}
							// line 359: self.parse_object = JSONObject
							πF.SetLineno(359)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßJSONObject); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparse_object, πTemp003); πE != nil {
								continue
							}
							// line 360: self.parse_array = JSONArray
							πF.SetLineno(360)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßJSONArray); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparse_array, πTemp003); πE != nil {
								continue
							}
							// line 361: self.parse_string = scanstring
							πF.SetLineno(361)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßscanstring); πE != nil {
								continue
							}
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßparse_string, πTemp003); πE != nil {
								continue
							}
							// line 362: self.scan_once = scanner.make_scanner(self)
							πF.SetLineno(362)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp005[0] = µself
							if πTemp001, πE = πg.ResolveGlobal(πF, ßscanner); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmake_scanner, nil); πE != nil {
								continue
							}
							if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πE = πg.SetAttr(πF, µself, ßscan_once, πTemp003); πE != nil {
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
					// line 310: """``encoding`` determines the encoding used to interpret any ``str``
					πF.SetLineno(310)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("``encoding`` determines the encoding used to interpret any ``str``\n        objects decoded by this instance (utf-8 by default).  It has no\n        effect when decoding ``unicode`` objects.\n\n        Note that currently only encodings that are a superset of ASCII work,\n        strings of other encodings should be passed in as ``unicode``.\n\n        ``object_hook``, if specified, will be called with the result\n        of every JSON object decoded and its return value will be used in\n        place of the given ``dict``.  This can be used to provide custom\n        deserializations (e.g. to support JSON-RPC class hinting).\n\n        ``object_pairs_hook``, if specified will be called with the result of\n        every JSON object decoded with an ordered list of pairs.  The return\n        value of ``object_pairs_hook`` will be used instead of the ``dict``.\n        This feature can be used to implement custom decoders that rely on the\n        order that the key and value pairs are decoded (for example,\n        collections.OrderedDict will remember the order of insertion). If\n        ``object_hook`` is also defined, the ``object_pairs_hook`` takes\n        priority.\n\n        ``parse_float``, if specified, will be called with the string\n        of every JSON float to be decoded. By default this is equivalent to\n        float(num_str). This can be used to use another datatype or parser\n        for JSON floats (e.g. decimal.Decimal).\n\n        ``parse_int``, if specified, will be called with the string\n        of every JSON int to be decoded. By default this is equivalent to\n        int(num_str). This can be used to use another datatype or parser\n        for JSON integers (e.g. float).\n\n        ``parse_constant``, if specified, will be called with one of the\n        following strings: -Infinity, Infinity, NaN.\n        This can be used to raise an exception if invalid JSON numbers\n        are encountered.\n\n        If ``strict`` is false (true is the default), then control\n        characters will be allowed inside strings.  Control characters in\n        this context are those with character codes in the 0-31 range,\n        including ``'\\t'`` (tab), ``'\\n'``, ``'\\r'`` and ``'\\0'``.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__init__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 364: def decode(self, s, _w=WHITESPACE.match):
					πF.SetLineno(364)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "s", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßWHITESPACE); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßmatch, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "_w", Def: πTemp005}
					πTemp003 = πg.NewFunction(πg.NewCode("decode", "/usr/lib/python2.7/json/decoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πArgs[1]; _ = µs
						var µ_w *πg.Object = πArgs[2]; _ = µ_w
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var µend *πg.Object = πg.UnboundLocal; _ = µend
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
						var πTemp006 *πg.Object
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
							default: panic("unexpected function state")
							}
							// line 365: """Return the Python representation of ``s`` (a ``str`` or ``unicode``
							πF.SetLineno(365)
							// line 369: obj, end = self.raw_decode(s, idx=_w(s, 0).end())
							πF.SetLineno(369)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[0] = µs
							πTemp002[1] = πg.NewInt(0).ToObject()
							if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
								continue
							}
							if πTemp003, πE = µ_w.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							πTemp005 = πg.KWArgs{
								{"idx", πTemp003},
							}
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ßraw_decode, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, πTemp005); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp006}}}, πTemp004); πE != nil {
								continue
							}
							µobj = πTemp003
							µend = πTemp006
							// line 370: end = _w(s, end).end()
							πF.SetLineno(370)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp001[1] = µend
							if πE = πg.CheckLocal(πF, µ_w, "_w"); πE != nil {
								continue
							}
							if πTemp003, πE = µ_w.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßend, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							µend = πTemp003
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp003, πE = πg.NE(πF, µend, πTemp006); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label1
							}
							goto Label2
							// line 371: if end != len(s):
							πF.SetLineno(371)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(4)
							πTemp002[0] = πg.NewStr("Extra data").ToObject()
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp002[1] = µs
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp002[2] = µend
							πTemp008 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp008[0] = µs
							if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp008)
							πTemp002[3] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßerrmsg); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 372: raise ValueError(errmsg("Extra data", s, end, len(s)))
							πF.SetLineno(372)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 373: return obj
							πF.SetLineno(373)
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
					if πE = πClass.SetItem(πF, ßdecode.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 365: """Return the Python representation of ``s`` (a ``str`` or ``unicode``
					πF.SetLineno(365)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Return the Python representation of ``s`` (a ``str`` or ``unicode``\n        instance containing a JSON document)\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßdecode); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 375: def raw_decode(self, s, idx=0):
					πF.SetLineno(375)
					πTemp002 = make([]πg.Param, 3)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp002[1] = πg.Param{Name: "s", Def: nil}
					πTemp002[2] = πg.Param{Name: "idx", Def: πg.NewInt(0).ToObject()}
					πTemp004 = πg.NewFunction(πg.NewCode("raw_decode", "/usr/lib/python2.7/json/decoder.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µs *πg.Object = πArgs[1]; _ = µs
						var µidx *πg.Object = πArgs[2]; _ = µidx
						var µobj *πg.Object = πg.UnboundLocal; _ = µobj
						var µend *πg.Object = πg.UnboundLocal; _ = µend
						var πTemp001 []*πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.BaseException
						_ = πTemp005
						var πTemp006 *πg.Traceback
						_ = πTemp006
						var πTemp007 bool
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							case 2: goto Label2
							default: panic("unexpected function state")
							}
							// line 376: """Decode a JSON document from ``s`` (a ``str`` or ``unicode``
							πF.SetLineno(376)
							// line 384: try:
							πF.SetLineno(384)
							πF.PushCheckpoint(2)
							// line 385: obj, end = self.scan_once(s, idx)
							πF.SetLineno(385)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
								continue
							}
							πTemp001[0] = µs
							if πE = πg.CheckLocal(πF, µidx, "idx"); πE != nil {
								continue
							}
							πTemp001[1] = µidx
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetAttr(πF, µself, ßscan_once, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
								continue
							}
							µobj = πTemp002
							µend = πTemp004
							πF.PopCheckpoint()
							goto Label1
						Label2:
							if πE == nil {
							  continue
							}
							πE = nil
							πTemp005, πTemp006 = πF.ExcInfo()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßStopIteration); πE != nil {
								continue
							}
							if πTemp007, πE = πg.IsInstance(πF, πTemp005.ToObject(), πTemp002); πE != nil {
								continue
							}
							if πTemp007 {
								goto Label3
							}
							πE = πF.Raise(πTemp005.ToObject(), nil, πTemp006.ToObject())
							continue
							// line 386: except StopIteration:
							πF.SetLineno(386)
						Label3:
							πTemp001 = πF.MakeArgs(1)
							πTemp001[0] = πg.NewStr("No JSON object could be decoded").ToObject()
							if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 387: raise ValueError("No JSON object could be decoded")
							πF.SetLineno(387)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							πF.RestoreExc(nil, nil)
							goto Label1
						Label1:
							// line 388: return obj, end
							πF.SetLineno(388)
							if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µend, "end"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple2(µobj, µend).ToObject()
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
					if πE = πClass.SetItem(πF, ßraw_decode.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 376: """Decode a JSON document from ``s`` (a ``str`` or ``unicode``
					πF.SetLineno(376)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Decode a JSON document from ``s`` (a ``str`` or ``unicode``\n        beginning with a JSON document) and return a 2-tuple of the Python\n        representation and the index in ``s`` where the document ended.\n\n        This can be used to decode a JSON document from a string that may\n        have extraneous data at the end.\n\n        ").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßraw_decode); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
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
			if πTemp014, πE = πTemp013.Call(πF, []*πg.Object{πg.NewStr("JSONDecoder").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp008.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßJSONDecoder.ToObject(), πTemp014); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("json.decoder", Code)
}

