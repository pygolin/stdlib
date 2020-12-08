package json
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/json/decoder"
	// _ "github.com/pygolin/stdlib/json/encoder"
	// _ "github.com/pygolin/stdlib/json_scanner"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/json/__init__.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßFalse := πg.InternStr("False")
		ßJSONDecoder := πg.InternStr("JSONDecoder")
		ßJSONEncoder := πg.InternStr("JSONEncoder")
		ßNone := πg.InternStr("None")
		ßTrue := πg.InternStr("True")
		ß__all__ := πg.InternStr("__all__")
		ß__author__ := πg.InternStr("__author__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__version__ := πg.InternStr("__version__")
		ß_default_decoder := πg.InternStr("_default_decoder")
		ß_default_encoder := πg.InternStr("_default_encoder")
		ßdecode := πg.InternStr("decode")
		ßdecoder := πg.InternStr("decoder")
		ßdump := πg.InternStr("dump")
		ßdumps := πg.InternStr("dumps")
		ßencode := πg.InternStr("encode")
		ßencoder := πg.InternStr("encoder")
		ßiterencode := πg.InternStr("iterencode")
		ßjson := πg.InternStr("json")
		ßjson_scanner := πg.InternStr("json_scanner")
		ßload := πg.InternStr("load")
		ßloads := πg.InternStr("loads")
		ßobject_hook := πg.InternStr("object_hook")
		ßobject_pairs_hook := πg.InternStr("object_pairs_hook")
		ßparse_constant := πg.InternStr("parse_constant")
		ßparse_float := πg.InternStr("parse_float")
		ßparse_int := πg.InternStr("parse_int")
		ßread := πg.InternStr("read")
		ßscanner := πg.InternStr("scanner")
		ßwrite := πg.InternStr("write")
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
		var πTemp009 πg.KWArgs
		_ = πTemp009
		var πTemp010 []πg.Param
		_ = πTemp010
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: r"""JSON (JavaScript Object Notation) <http://json.org> is a subset of
			πF.SetLineno(1)
			// line 1: r"""JSON (JavaScript Object Notation) <http://json.org> is a subset of
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("JSON (JavaScript Object Notation) <http://json.org> is a subset of\nJavaScript syntax (ECMA-262 3rd edition) used as a lightweight data\ninterchange format.\n\n:mod:`json` exposes an API familiar to users of the standard library\n:mod:`marshal` and :mod:`pickle` modules. It is the externally maintained\nversion of the :mod:`json` library contained in Python 2.6, but maintains\ncompatibility with Python 2.4 and Python 2.5 and (currently) has\nsignificant performance advantages, even without using the optional C\nextension for speedups.\n\nEncoding basic Python object hierarchies::\n\n    >>> import json\n    >>> json.dumps(['foo', {'bar': ('baz', None, 1.0, 2)}])\n    '[\"foo\", {\"bar\": [\"baz\", null, 1.0, 2]}]'\n    >>> print json.dumps(\"\\\"foo\\bar\")\n    \"\\\"foo\\bar\"\n    >>> print json.dumps(u'\\u1234')\n    \"\\u1234\"\n    >>> print json.dumps('\\\\')\n    \"\\\\\"\n    >>> print json.dumps({\"c\": 0, \"b\": 0, \"a\": 0}, sort_keys=True)\n    {\"a\": 0, \"b\": 0, \"c\": 0}\n    >>> from StringIO import StringIO\n    >>> io = StringIO()\n    >>> json.dump(['streaming API'], io)\n    >>> io.getvalue()\n    '[\"streaming API\"]'\n\nCompact encoding::\n\n    >>> import json\n    >>> json.dumps([1,2,3,{'4': 5, '6': 7}], sort_keys=True, separators=(',',':'))\n    '[1,2,3,{\"4\":5,\"6\":7}]'\n\nPretty printing::\n\n    >>> import json\n    >>> print json.dumps({'4': 5, '6': 7}, sort_keys=True,\n    ...                  indent=4, separators=(',', ': '))\n    {\n        \"4\": 5,\n        \"6\": 7\n    }\n\nDecoding JSON::\n\n    >>> import json\n    >>> obj = [u'foo', {u'bar': [u'baz', None, 1.0, 2]}]\n    >>> json.loads('[\"foo\", {\"bar\":[\"baz\", null, 1.0, 2]}]') == obj\n    True\n    >>> json.loads('\"\\\\\"foo\\\\bar\"') == u'\"foo\\x08ar'\n    True\n    >>> from StringIO import StringIO\n    >>> io = StringIO('[\"streaming API\"]')\n    >>> json.load(io)[0] == 'streaming API'\n    True\n\nSpecializing JSON object decoding::\n\n    >>> import json\n    >>> def as_complex(dct):\n    ...     if '__complex__' in dct:\n    ...         return complex(dct['real'], dct['imag'])\n    ...     return dct\n    ...\n    >>> json.loads('{\"__complex__\": true, \"real\": 1, \"imag\": 2}',\n    ...     object_hook=as_complex)\n    (1+2j)\n    >>> from decimal import Decimal\n    >>> json.loads('1.1', parse_float=Decimal) == Decimal('1.1')\n    True\n\nSpecializing JSON object encoding::\n\n    >>> import json\n    >>> def encode_complex(obj):\n    ...     if isinstance(obj, complex):\n    ...         return [obj.real, obj.imag]\n    ...     raise TypeError(repr(o) + \" is not JSON serializable\")\n    ...\n    >>> json.dumps(2 + 1j, default=encode_complex)\n    '[2.0, 1.0]'\n    >>> json.JSONEncoder(default=encode_complex).encode(2 + 1j)\n    '[2.0, 1.0]'\n    >>> ''.join(json.JSONEncoder(default=encode_complex).iterencode(2 + 1j))\n    '[2.0, 1.0]'\n\n\nUsing json.tool from the shell to validate and pretty-print::\n\n    $ echo '{\"json\":\"obj\"}' | python -m json.tool\n    {\n        \"json\": \"obj\"\n    }\n    $ echo '{ 1.2:3.4}' | python -m json.tool\n    Expecting property name enclosed in double quotes: line 1 column 3 (char 2)\n").ToObject()); πE != nil {
				continue
			}
			// line 100: __version__ = '2.0.9'
			πF.SetLineno(100)
			if πE = πF.Globals().SetItem(πF, ß__version__.ToObject(), πg.NewStr("2.0.9").ToObject()); πE != nil {
				continue
			}
			// line 101: __all__ = [
			πF.SetLineno(101)
			πTemp001 = make([]*πg.Object, 6)
			πTemp001[0] = ßdump.ToObject()
			πTemp001[1] = ßdumps.ToObject()
			πTemp001[2] = ßload.ToObject()
			πTemp001[3] = ßloads.ToObject()
			πTemp001[4] = ßJSONDecoder.ToObject()
			πTemp001[5] = ßJSONEncoder.ToObject()
			πTemp002 = πg.NewList(πTemp001...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 106: __author__ = 'Bob Ippolito <bob@redivi.com>'
			πF.SetLineno(106)
			if πE = πF.Globals().SetItem(πF, ß__author__.ToObject(), πg.NewStr("Bob Ippolito <bob@redivi.com>").ToObject()); πE != nil {
				continue
			}
			// line 110: import json.decoder
			πF.SetLineno(110)
			if πTemp001, πE = πg.ImportModule(πF, "json.decoder"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßjson.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 111: import json.encoder
			πF.SetLineno(111)
			if πTemp001, πE = πg.ImportModule(πF, "json.encoder"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßjson.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 112: import json_scanner
			πF.SetLineno(112)
			if πTemp001, πE = πg.ImportModule(πF, "json_scanner"); πE != nil {
				continue
			}
			πTemp002 = πTemp001[0]
			if πE = πF.Globals().SetItem(πF, ßjson_scanner.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 113: JSONDecoder = json.decoder.JSONDecoder
			πF.SetLineno(113)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßjson); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdecoder, nil); πE != nil {
				continue
			}
			if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßJSONDecoder, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßJSONDecoder.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 114: JSONEncoder = json.encoder.JSONEncoder
			πF.SetLineno(114)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßjson); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßencoder, nil); πE != nil {
				continue
			}
			if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßJSONEncoder, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßJSONEncoder.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 115: scanner = json_scanner
			πF.SetLineno(115)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßjson_scanner); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßscanner.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 117: _default_encoder = JSONEncoder(
			πF.SetLineno(117)
			if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πTemp008, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp009 = πg.KWArgs{
				{"skipkeys", πTemp002},
				{"ensure_ascii", πTemp003},
				{"check_circular", πTemp004},
				{"allow_nan", πTemp005},
				{"indent", πTemp006},
				{"separators", πTemp007},
				{"encoding", πg.NewStr("utf-8").ToObject()},
				{"default", πTemp008},
			}
			if πTemp002, πE = πg.ResolveGlobal(πF, ßJSONEncoder); πE != nil {
				continue
			}
			if πTemp003, πE = πTemp002.Call(πF, nil, πTemp009); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_default_encoder.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 128: def dump(obj, fp, skipkeys=False, ensure_ascii=True, check_circular=True,
			πF.SetLineno(128)
			πTemp010 = make([]πg.Param, 12)
			πTemp010[0] = πg.Param{Name: "obj", Def: nil}
			πTemp010[1] = πg.Param{Name: "fp", Def: nil}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp010[2] = πg.Param{Name: "skipkeys", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp010[3] = πg.Param{Name: "ensure_ascii", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp010[4] = πg.Param{Name: "check_circular", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp010[5] = πg.Param{Name: "allow_nan", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[6] = πg.Param{Name: "cls", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[7] = πg.Param{Name: "indent", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[8] = πg.Param{Name: "separators", Def: πTemp003}
			πTemp010[9] = πg.Param{Name: "encoding", Def: πg.NewStr("utf-8").ToObject()}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[10] = πg.Param{Name: "default", Def: πTemp003}
			if πTemp003, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp010[11] = πg.Param{Name: "sort_keys", Def: πTemp003}
			πTemp002 = πg.NewFunction(πg.NewCode("dump", "/usr/lib/python2.7/json/__init__.py", πTemp010, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µfp *πg.Object = πArgs[1]; _ = µfp
				var µskipkeys *πg.Object = πArgs[2]; _ = µskipkeys
				var µensure_ascii *πg.Object = πArgs[3]; _ = µensure_ascii
				var µcheck_circular *πg.Object = πArgs[4]; _ = µcheck_circular
				var µallow_nan *πg.Object = πArgs[5]; _ = µallow_nan
				var µcls *πg.Object = πArgs[6]; _ = µcls
				var µindent *πg.Object = πArgs[7]; _ = µindent
				var µseparators *πg.Object = πArgs[8]; _ = µseparators
				var µencoding *πg.Object = πArgs[9]; _ = µencoding
				var µdefault *πg.Object = πArgs[10]; _ = µdefault
				var µsort_keys *πg.Object = πArgs[11]; _ = µsort_keys
				var µkw *πg.Object = πArgs[12]; _ = µkw
				var µiterable *πg.Object = πg.UnboundLocal; _ = µiterable
				var µchunk *πg.Object = πg.UnboundLocal; _ = µchunk
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 8: goto Label8
					case 7: goto Label7
					default: panic("unexpected function state")
					}
					// line 131: """Serialize ``obj`` as a JSON formatted stream to ``fp`` (a
					πF.SetLineno(131)
					if πE = πg.CheckLocal(πF, µskipkeys, "skipkeys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µskipkeys); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µensure_ascii, "ensure_ascii"); πE != nil {
						continue
					}
					πTemp001 = µensure_ascii
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µcheck_circular, "check_circular"); πE != nil {
						continue
					}
					πTemp001 = µcheck_circular
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µallow_nan, "allow_nan"); πE != nil {
						continue
					}
					πTemp001 = µallow_nan
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µcls == πTemp005).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µindent == πTemp005).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µseparators, "separators"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µseparators == πTemp005).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µencoding, πg.NewStr("utf-8").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µdefault == πTemp005).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µsort_keys, "sort_keys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µsort_keys); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µkw); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp004).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 181: if (not skipkeys and ensure_ascii and
					πF.SetLineno(181)
				Label2:
					// line 185: iterable = _default_encoder.iterencode(obj)
					πF.SetLineno(185)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_default_encoder); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßiterencode, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					µiterable = πTemp001
					goto Label4
				Label3:
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µcls == πTemp003).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					goto Label6
					// line 187: if cls is None:
					πF.SetLineno(187)
				Label5:
					// line 188: cls = JSONEncoder
					πF.SetLineno(188)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßJSONEncoder); πE != nil {
						continue
					}
					µcls = πTemp001
					goto Label6
				Label6:
					// line 189: iterable = cls(skipkeys=skipkeys, ensure_ascii=ensure_ascii,
					πF.SetLineno(189)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					if πE = πg.CheckLocal(πF, µskipkeys, "skipkeys"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µensure_ascii, "ensure_ascii"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcheck_circular, "check_circular"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µallow_nan, "allow_nan"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µseparators, "separators"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsort_keys, "sort_keys"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"skipkeys", µskipkeys},
						{"ensure_ascii", µensure_ascii},
						{"check_circular", µcheck_circular},
						{"allow_nan", µallow_nan},
						{"indent", µindent},
						{"separators", µseparators},
						{"encoding", µencoding},
						{"default", µdefault},
						{"sort_keys", µsort_keys},
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µcls, nil, nil, πTemp007, µkw); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßiterencode, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					µiterable = πTemp001
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µiterable); πE != nil {
						continue
					}
					πF.PushCheckpoint(8)
					πTemp002 = false
				Label7:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
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
						πTemp004 = !isStop
					} else {
						πTemp004 = true
						µchunk = πTemp003
					}
					if πE != nil || !πTemp004 {
						continue
					}
					πF.PushCheckpoint(7)            
					// line 196: fp.write(chunk)
					πF.SetLineno(196)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µchunk, "chunk"); πE != nil {
						continue
					}
					πTemp006[0] = µchunk
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µfp, ßwrite, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					continue
				Label8:
					if πE != nil || πR != nil {
						continue
					}
				Label9:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßdump.ToObject(), πTemp002); πE != nil {
				continue
			}
			// line 131: """Serialize ``obj`` as a JSON formatted stream to ``fp`` (a
			πF.SetLineno(131)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Serialize ``obj`` as a JSON formatted stream to ``fp`` (a\n    ``.write()``-supporting file-like object).\n\n    If ``skipkeys`` is true then ``dict`` keys that are not basic types\n    (``str``, ``unicode``, ``int``, ``long``, ``float``, ``bool``, ``None``)\n    will be skipped instead of raising a ``TypeError``.\n\n    If ``ensure_ascii`` is true (the default), all non-ASCII characters in the\n    output are escaped with ``\\uXXXX`` sequences, and the result is a ``str``\n    instance consisting of ASCII characters only.  If ``ensure_ascii`` is\n    ``False``, some chunks written to ``fp`` may be ``unicode`` instances.\n    This usually happens because the input contains unicode strings or the\n    ``encoding`` parameter is used. Unless ``fp.write()`` explicitly\n    understands ``unicode`` (as in ``codecs.getwriter``) this is likely to\n    cause an error.\n\n    If ``check_circular`` is false, then the circular reference check\n    for container types will be skipped and a circular reference will\n    result in an ``OverflowError`` (or worse).\n\n    If ``allow_nan`` is false, then it will be a ``ValueError`` to\n    serialize out of range ``float`` values (``nan``, ``inf``, ``-inf``)\n    in strict compliance of the JSON specification, instead of using the\n    JavaScript equivalents (``NaN``, ``Infinity``, ``-Infinity``).\n\n    If ``indent`` is a non-negative integer, then JSON array elements and\n    object members will be pretty-printed with that indent level. An indent\n    level of 0 will only insert newlines. ``None`` is the most compact\n    representation.  Since the default item separator is ``', '``,  the\n    output might include trailing whitespace when ``indent`` is specified.\n    You can use ``separators=(',', ': ')`` to avoid this.\n\n    If ``separators`` is an ``(item_separator, dict_separator)`` tuple\n    then it will be used instead of the default ``(', ', ': ')`` separators.\n    ``(',', ':')`` is the most compact JSON representation.\n\n    ``encoding`` is the character encoding for str instances, default is UTF-8.\n\n    ``default(obj)`` is a function that should return a serializable version\n    of obj or raise TypeError. The default simply raises TypeError.\n\n    If *sort_keys* is ``True`` (default: ``False``), then the output of\n    dictionaries will be sorted by key.\n\n    To use a custom ``JSONEncoder`` subclass (e.g. one that overrides the\n    ``.default()`` method to serialize additional types), specify it with\n    the ``cls`` kwarg; otherwise ``JSONEncoder`` is used.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßdump); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 199: def dumps(obj, skipkeys=False, ensure_ascii=True, check_circular=True,
			πF.SetLineno(199)
			πTemp010 = make([]πg.Param, 11)
			πTemp010[0] = πg.Param{Name: "obj", Def: nil}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp010[1] = πg.Param{Name: "skipkeys", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp010[2] = πg.Param{Name: "ensure_ascii", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp010[3] = πg.Param{Name: "check_circular", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp010[4] = πg.Param{Name: "allow_nan", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[5] = πg.Param{Name: "cls", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[6] = πg.Param{Name: "indent", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[7] = πg.Param{Name: "separators", Def: πTemp004}
			πTemp010[8] = πg.Param{Name: "encoding", Def: πg.NewStr("utf-8").ToObject()}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[9] = πg.Param{Name: "default", Def: πTemp004}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
				continue
			}
			πTemp010[10] = πg.Param{Name: "sort_keys", Def: πTemp004}
			πTemp003 = πg.NewFunction(πg.NewCode("dumps", "/usr/lib/python2.7/json/__init__.py", πTemp010, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µobj *πg.Object = πArgs[0]; _ = µobj
				var µskipkeys *πg.Object = πArgs[1]; _ = µskipkeys
				var µensure_ascii *πg.Object = πArgs[2]; _ = µensure_ascii
				var µcheck_circular *πg.Object = πArgs[3]; _ = µcheck_circular
				var µallow_nan *πg.Object = πArgs[4]; _ = µallow_nan
				var µcls *πg.Object = πArgs[5]; _ = µcls
				var µindent *πg.Object = πArgs[6]; _ = µindent
				var µseparators *πg.Object = πArgs[7]; _ = µseparators
				var µencoding *πg.Object = πArgs[8]; _ = µencoding
				var µdefault *πg.Object = πArgs[9]; _ = µdefault
				var µsort_keys *πg.Object = πArgs[10]; _ = µsort_keys
				var µkw *πg.Object = πArgs[11]; _ = µkw
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 202: """Serialize ``obj`` to a JSON formatted ``str``.
					πF.SetLineno(202)
					if πE = πg.CheckLocal(πF, µskipkeys, "skipkeys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µskipkeys); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µensure_ascii, "ensure_ascii"); πE != nil {
						continue
					}
					πTemp001 = µensure_ascii
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µcheck_circular, "check_circular"); πE != nil {
						continue
					}
					πTemp001 = µcheck_circular
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µallow_nan, "allow_nan"); πE != nil {
						continue
					}
					πTemp001 = µallow_nan
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µcls == πTemp005).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µindent == πTemp005).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µseparators, "separators"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µseparators == πTemp005).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Eq(πF, µencoding, πg.NewStr("utf-8").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µdefault == πTemp005).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µsort_keys, "sort_keys"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µsort_keys); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, µkw); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp004).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 246: if (not skipkeys and ensure_ascii and
					πF.SetLineno(246)
				Label2:
					// line 250: return _default_encoder.encode(obj)
					πF.SetLineno(250)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_default_encoder); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßencode, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πR = πTemp001
					continue
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µcls == πTemp003).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 251: if cls is None:
					πF.SetLineno(251)
				Label4:
					// line 252: cls = JSONEncoder
					πF.SetLineno(252)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßJSONEncoder); πE != nil {
						continue
					}
					µcls = πTemp001
					goto Label5
				Label5:
					// line 253: return cls(
					πF.SetLineno(253)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
						continue
					}
					πTemp006[0] = µobj
					if πE = πg.CheckLocal(πF, µskipkeys, "skipkeys"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µensure_ascii, "ensure_ascii"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcheck_circular, "check_circular"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µallow_nan, "allow_nan"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µindent, "indent"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µseparators, "separators"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdefault, "default"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsort_keys, "sort_keys"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"skipkeys", µskipkeys},
						{"ensure_ascii", µensure_ascii},
						{"check_circular", µcheck_circular},
						{"allow_nan", µallow_nan},
						{"indent", µindent},
						{"separators", µseparators},
						{"encoding", µencoding},
						{"default", µdefault},
						{"sort_keys", µsort_keys},
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µcls, nil, nil, πTemp007, µkw); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßencode, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
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
			if πE = πF.Globals().SetItem(πF, ßdumps.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 202: """Serialize ``obj`` to a JSON formatted ``str``.
			πF.SetLineno(202)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Serialize ``obj`` to a JSON formatted ``str``.\n\n    If ``skipkeys`` is true then ``dict`` keys that are not basic types\n    (``str``, ``unicode``, ``int``, ``long``, ``float``, ``bool``, ``None``)\n    will be skipped instead of raising a ``TypeError``.\n\n\n    If ``ensure_ascii`` is false, all non-ASCII characters are not escaped, and\n    the return value may be a ``unicode`` instance. See ``dump`` for details.\n\n    If ``check_circular`` is false, then the circular reference check\n    for container types will be skipped and a circular reference will\n    result in an ``OverflowError`` (or worse).\n\n    If ``allow_nan`` is false, then it will be a ``ValueError`` to\n    serialize out of range ``float`` values (``nan``, ``inf``, ``-inf``) in\n    strict compliance of the JSON specification, instead of using the\n    JavaScript equivalents (``NaN``, ``Infinity``, ``-Infinity``).\n\n    If ``indent`` is a non-negative integer, then JSON array elements and\n    object members will be pretty-printed with that indent level. An indent\n    level of 0 will only insert newlines. ``None`` is the most compact\n    representation.  Since the default item separator is ``', '``,  the\n    output might include trailing whitespace when ``indent`` is specified.\n    You can use ``separators=(',', ': ')`` to avoid this.\n\n    If ``separators`` is an ``(item_separator, dict_separator)`` tuple\n    then it will be used instead of the default ``(', ', ': ')`` separators.\n    ``(',', ':')`` is the most compact JSON representation.\n\n    ``encoding`` is the character encoding for str instances, default is UTF-8.\n\n    ``default(obj)`` is a function that should return a serializable version\n    of obj or raise TypeError. The default simply raises TypeError.\n\n    If *sort_keys* is ``True`` (default: ``False``), then the output of\n    dictionaries will be sorted by key.\n\n    To use a custom ``JSONEncoder`` subclass (e.g. one that overrides the\n    ``.default()`` method to serialize additional types), specify it with\n    the ``cls`` kwarg; otherwise ``JSONEncoder`` is used.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßdumps); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
				continue
			}
			// line 260: _default_decoder = JSONDecoder(encoding=None, object_hook=None,
			πF.SetLineno(260)
			if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp009 = πg.KWArgs{
				{"encoding", πTemp004},
				{"object_hook", πTemp005},
				{"object_pairs_hook", πTemp006},
			}
			if πTemp004, πE = πg.ResolveGlobal(πF, ßJSONDecoder); πE != nil {
				continue
			}
			if πTemp005, πE = πTemp004.Call(πF, nil, πTemp009); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_default_decoder.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 264: def load(fp, encoding=None, cls=None, object_hook=None, parse_float=None,
			πF.SetLineno(264)
			πTemp010 = make([]πg.Param, 8)
			πTemp010[0] = πg.Param{Name: "fp", Def: nil}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[1] = πg.Param{Name: "encoding", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[2] = πg.Param{Name: "cls", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[3] = πg.Param{Name: "object_hook", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[4] = πg.Param{Name: "parse_float", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[5] = πg.Param{Name: "parse_int", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[6] = πg.Param{Name: "parse_constant", Def: πTemp005}
			if πTemp005, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[7] = πg.Param{Name: "object_pairs_hook", Def: πTemp005}
			πTemp004 = πg.NewFunction(πg.NewCode("load", "/usr/lib/python2.7/json/__init__.py", πTemp010, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfp *πg.Object = πArgs[0]; _ = µfp
				var µencoding *πg.Object = πArgs[1]; _ = µencoding
				var µcls *πg.Object = πArgs[2]; _ = µcls
				var µobject_hook *πg.Object = πArgs[3]; _ = µobject_hook
				var µparse_float *πg.Object = πArgs[4]; _ = µparse_float
				var µparse_int *πg.Object = πArgs[5]; _ = µparse_int
				var µparse_constant *πg.Object = πArgs[6]; _ = µparse_constant
				var µobject_pairs_hook *πg.Object = πArgs[7]; _ = µobject_pairs_hook
				var µkw *πg.Object = πArgs[8]; _ = µkw
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
					// line 266: """Deserialize ``fp`` (a ``.read()``-supporting file-like object containing
					πF.SetLineno(266)
					// line 293: return loads(fp.read(),
					πF.SetLineno(293)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfp, "fp"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µfp, ßread, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparse_float, "parse_float"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparse_int, "parse_int"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparse_constant, "parse_constant"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
						continue
					}
					πTemp004 = πg.KWArgs{
						{"encoding", µencoding},
						{"cls", µcls},
						{"object_hook", µobject_hook},
						{"parse_float", µparse_float},
						{"parse_int", µparse_int},
						{"parse_constant", µparse_constant},
						{"object_pairs_hook", µobject_pairs_hook},
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.ResolveGlobal(πF, ßloads); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Invoke(πF, πTemp002, πTemp001, nil, πTemp004, µkw); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßload.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 266: """Deserialize ``fp`` (a ``.read()``-supporting file-like object containing
			πF.SetLineno(266)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Deserialize ``fp`` (a ``.read()``-supporting file-like object containing\n    a JSON document) to a Python object.\n\n    If the contents of ``fp`` is encoded with an ASCII based encoding other\n    than utf-8 (e.g. latin-1), then an appropriate ``encoding`` name must\n    be specified. Encodings that are not ASCII based (such as UCS-2) are\n    not allowed, and should be wrapped with\n    ``codecs.getreader(fp)(encoding)``, or simply decoded to a ``unicode``\n    object and passed to ``loads()``\n\n    ``object_hook`` is an optional function that will be called with the\n    result of any object literal decode (a ``dict``). The return value of\n    ``object_hook`` will be used instead of the ``dict``. This feature\n    can be used to implement custom decoders (e.g. JSON-RPC class hinting).\n\n    ``object_pairs_hook`` is an optional function that will be called with the\n    result of any object literal decoded with an ordered list of pairs.  The\n    return value of ``object_pairs_hook`` will be used instead of the ``dict``.\n    This feature can be used to implement custom decoders that rely on the\n    order that the key and value pairs are decoded (for example,\n    collections.OrderedDict will remember the order of insertion). If\n    ``object_hook`` is also defined, the ``object_pairs_hook`` takes priority.\n\n    To use a custom ``JSONDecoder`` subclass, specify it with the ``cls``\n    kwarg; otherwise ``JSONDecoder`` is used.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßload); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp005); πE != nil {
				continue
			}
			// line 300: def loads(s, encoding=None, cls=None, object_hook=None, parse_float=None,
			πF.SetLineno(300)
			πTemp010 = make([]πg.Param, 8)
			πTemp010[0] = πg.Param{Name: "s", Def: nil}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[1] = πg.Param{Name: "encoding", Def: πTemp006}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[2] = πg.Param{Name: "cls", Def: πTemp006}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[3] = πg.Param{Name: "object_hook", Def: πTemp006}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[4] = πg.Param{Name: "parse_float", Def: πTemp006}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[5] = πg.Param{Name: "parse_int", Def: πTemp006}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[6] = πg.Param{Name: "parse_constant", Def: πTemp006}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
				continue
			}
			πTemp010[7] = πg.Param{Name: "object_pairs_hook", Def: πTemp006}
			πTemp005 = πg.NewFunction(πg.NewCode("loads", "/usr/lib/python2.7/json/__init__.py", πTemp010, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µencoding *πg.Object = πArgs[1]; _ = µencoding
				var µcls *πg.Object = πArgs[2]; _ = µcls
				var µobject_hook *πg.Object = πArgs[3]; _ = µobject_hook
				var µparse_float *πg.Object = πArgs[4]; _ = µparse_float
				var µparse_int *πg.Object = πArgs[5]; _ = µparse_int
				var µparse_constant *πg.Object = πArgs[6]; _ = µparse_constant
				var µobject_pairs_hook *πg.Object = πArgs[7]; _ = µobject_pairs_hook
				var µkw *πg.Object = πArgs[8]; _ = µkw
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
				var πTemp006 []*πg.Object
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 302: """Deserialize ``s`` (a ``str`` or ``unicode`` instance containing a JSON
					πF.SetLineno(302)
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µcls == πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µencoding == πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µobject_hook == πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µparse_int, "parse_int"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µparse_int == πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µparse_float, "parse_float"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µparse_float == πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µparse_constant, "parse_constant"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µparse_constant == πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(µobject_pairs_hook == πTemp004).ToObject()
					πTemp001 = πTemp003
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, µkw); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp005).ToObject()
					πTemp001 = πTemp003
				Label1:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label2
					}
					goto Label3
					// line 342: if (cls is None and encoding is None and object_hook is None and
					πF.SetLineno(342)
				Label2:
					// line 345: return _default_decoder.decode(s)
					πF.SetLineno(345)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[0] = µs
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_default_decoder); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πR = πTemp001
					continue
					goto Label3
				Label3:
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µcls == πTemp003).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 346: if cls is None:
					πF.SetLineno(346)
				Label4:
					// line 347: cls = JSONDecoder
					πF.SetLineno(347)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßJSONDecoder); πE != nil {
						continue
					}
					µcls = πTemp001
					goto Label5
				Label5:
					if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µobject_hook != πTemp003).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					goto Label7
					// line 348: if object_hook is not None:
					πF.SetLineno(348)
				Label6:
					// line 349: kw['object_hook'] = object_hook
					πF.SetLineno(349)
					if πE = πg.CheckLocal(πF, µobject_hook, "object_hook"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µobject_hook); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					πTemp003 = ßobject_hook.ToObject()
					if πE = πg.SetItem(πF, µkw, πTemp003, πTemp001); πE != nil {
						continue
					}
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µobject_pairs_hook != πTemp003).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					goto Label9
					// line 350: if object_pairs_hook is not None:
					πF.SetLineno(350)
				Label8:
					// line 351: kw['object_pairs_hook'] = object_pairs_hook
					πF.SetLineno(351)
					if πE = πg.CheckLocal(πF, µobject_pairs_hook, "object_pairs_hook"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µobject_pairs_hook); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					πTemp003 = ßobject_pairs_hook.ToObject()
					if πE = πg.SetItem(πF, µkw, πTemp003, πTemp001); πE != nil {
						continue
					}
					goto Label9
				Label9:
					if πE = πg.CheckLocal(πF, µparse_float, "parse_float"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µparse_float != πTemp003).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					goto Label11
					// line 352: if parse_float is not None:
					πF.SetLineno(352)
				Label10:
					// line 353: kw['parse_float'] = parse_float
					πF.SetLineno(353)
					if πE = πg.CheckLocal(πF, µparse_float, "parse_float"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µparse_float); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					πTemp003 = ßparse_float.ToObject()
					if πE = πg.SetItem(πF, µkw, πTemp003, πTemp001); πE != nil {
						continue
					}
					goto Label11
				Label11:
					if πE = πg.CheckLocal(πF, µparse_int, "parse_int"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µparse_int != πTemp003).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					goto Label13
					// line 354: if parse_int is not None:
					πF.SetLineno(354)
				Label12:
					// line 355: kw['parse_int'] = parse_int
					πF.SetLineno(355)
					if πE = πg.CheckLocal(πF, µparse_int, "parse_int"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µparse_int); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					πTemp003 = ßparse_int.ToObject()
					if πE = πg.SetItem(πF, µkw, πTemp003, πTemp001); πE != nil {
						continue
					}
					goto Label13
				Label13:
					if πE = πg.CheckLocal(πF, µparse_constant, "parse_constant"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(µparse_constant != πTemp003).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label14
					}
					goto Label15
					// line 356: if parse_constant is not None:
					πF.SetLineno(356)
				Label14:
					// line 357: kw['parse_constant'] = parse_constant
					πF.SetLineno(357)
					if πE = πg.CheckLocal(πF, µparse_constant, "parse_constant"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µparse_constant); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					πTemp003 = ßparse_constant.ToObject()
					if πE = πg.SetItem(πF, µkw, πTemp003, πTemp001); πE != nil {
						continue
					}
					goto Label15
				Label15:
					// line 358: return cls(encoding=encoding, **kw).decode(s)
					πF.SetLineno(358)
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp006[0] = µs
					if πE = πg.CheckLocal(πF, µencoding, "encoding"); πE != nil {
						continue
					}
					πTemp007 = πg.KWArgs{
						{"encoding", µencoding},
					}
					if πE = πg.CheckLocal(πF, µkw, "kw"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Invoke(πF, µcls, nil, nil, πTemp007, µkw); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πTemp003.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
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
			if πE = πF.Globals().SetItem(πF, ßloads.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 302: """Deserialize ``s`` (a ``str`` or ``unicode`` instance containing a JSON
			πF.SetLineno(302)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Deserialize ``s`` (a ``str`` or ``unicode`` instance containing a JSON\n    document) to a Python object.\n\n    If ``s`` is a ``str`` instance and is encoded with an ASCII based encoding\n    other than utf-8 (e.g. latin-1) then an appropriate ``encoding`` name\n    must be specified. Encodings that are not ASCII based (such as UCS-2)\n    are not allowed and should be decoded to ``unicode`` first.\n\n    ``object_hook`` is an optional function that will be called with the\n    result of any object literal decode (a ``dict``). The return value of\n    ``object_hook`` will be used instead of the ``dict``. This feature\n    can be used to implement custom decoders (e.g. JSON-RPC class hinting).\n\n    ``object_pairs_hook`` is an optional function that will be called with the\n    result of any object literal decoded with an ordered list of pairs.  The\n    return value of ``object_pairs_hook`` will be used instead of the ``dict``.\n    This feature can be used to implement custom decoders that rely on the\n    order that the key and value pairs are decoded (for example,\n    collections.OrderedDict will remember the order of insertion). If\n    ``object_hook`` is also defined, the ``object_pairs_hook`` takes priority.\n\n    ``parse_float``, if specified, will be called with the string\n    of every JSON float to be decoded. By default this is equivalent to\n    float(num_str). This can be used to use another datatype or parser\n    for JSON floats (e.g. decimal.Decimal).\n\n    ``parse_int``, if specified, will be called with the string\n    of every JSON int to be decoded. By default this is equivalent to\n    int(num_str). This can be used to use another datatype or parser\n    for JSON integers (e.g. float).\n\n    ``parse_constant``, if specified, will be called with one of the\n    following strings: -Infinity, Infinity, NaN, null, true, false.\n    This can be used to raise an exception if invalid JSON numbers\n    are encountered.\n\n    To use a custom ``JSONDecoder`` subclass, specify it with the ``cls``\n    kwarg; otherwise ``JSONDecoder`` is used.\n\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßloads); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("json", Code)
}

