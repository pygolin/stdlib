package urlparse
import (
	πg "github.com/pygolin/runtime"
	// _ "github.com/pygolin/stdlib/operator"
	// _ "github.com/pygolin/stdlib/re"
)
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "/usr/lib/python2.7/urlparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ß := πg.InternStr("")
		ß0123456789 := πg.InternStr("0123456789")
		ß0123456789ABCDEFabcdef := πg.InternStr("0123456789ABCDEFabcdef")
		ßKeyError := πg.InternStr("KeyError")
		ßMAX_CACHE_SIZE := πg.InternStr("MAX_CACHE_SIZE")
		ßNameError := πg.InternStr("NameError")
		ßNone := πg.InternStr("None")
		ßOrderedDict := πg.InternStr("OrderedDict")
		ßParseResult := πg.InternStr("ParseResult")
		ßResultMixin := πg.InternStr("ResultMixin")
		ßSplitResult := πg.InternStr("SplitResult")
		ßTrue := πg.InternStr("True")
		ßTypeError := πg.InternStr("TypeError")
		ßValueError := πg.InternStr("ValueError")
		ß_ParseResult := πg.InternStr("_ParseResult")
		ß_SplitResult := πg.InternStr("_SplitResult")
		ß__all__ := πg.InternStr("__all__")
		ß__dict__ := πg.InternStr("__dict__")
		ß__doc__ := πg.InternStr("__doc__")
		ß__getnewargs__ := πg.InternStr("__getnewargs__")
		ß__getstate__ := πg.InternStr("__getstate__")
		ß__metaclass__ := πg.InternStr("__metaclass__")
		ß__module__ := πg.InternStr("__module__")
		ß__name__ := πg.InternStr("__name__")
		ß__new__ := πg.InternStr("__new__")
		ß__repr__ := πg.InternStr("__repr__")
		ß__slots__ := πg.InternStr("__slots__")
		ß_asciire := πg.InternStr("_asciire")
		ß_asdict := πg.InternStr("_asdict")
		ß_fields := πg.InternStr("_fields")
		ß_hexdig := πg.InternStr("_hexdig")
		ß_hextochr := πg.InternStr("_hextochr")
		ß_is_unicode := πg.InternStr("_is_unicode")
		ß_itemgetter := πg.InternStr("_itemgetter")
		ß_make := πg.InternStr("_make")
		ß_parse_cache := πg.InternStr("_parse_cache")
		ß_property := πg.InternStr("_property")
		ß_replace := πg.InternStr("_replace")
		ß_splitnetloc := πg.InternStr("_splitnetloc")
		ß_splitparams := πg.InternStr("_splitparams")
		ß_tuple := πg.InternStr("_tuple")
		ßany := πg.InternStr("any")
		ßappend := πg.InternStr("append")
		ßbool := πg.InternStr("bool")
		ßchr := πg.InternStr("chr")
		ßclassmethod := πg.InternStr("classmethod")
		ßclear := πg.InternStr("clear")
		ßclear_cache := πg.InternStr("clear_cache")
		ßcompile := πg.InternStr("compile")
		ßdecode := πg.InternStr("decode")
		ßdict := πg.InternStr("dict")
		ßfile := πg.InternStr("file")
		ßfind := πg.InternStr("find")
		ßfragment := πg.InternStr("fragment")
		ßftp := πg.InternStr("ftp")
		ßget := πg.InternStr("get")
		ßgeturl := πg.InternStr("geturl")
		ßgit := πg.InternStr("git")
		ßgopher := πg.InternStr("gopher")
		ßhdl := πg.InternStr("hdl")
		ßhostname := πg.InternStr("hostname")
		ßhttp := πg.InternStr("http")
		ßhttps := πg.InternStr("https")
		ßimap := πg.InternStr("imap")
		ßint := πg.InternStr("int")
		ßisinstance := πg.InternStr("isinstance")
		ßitemgetter := πg.InternStr("itemgetter")
		ßjoin := πg.InternStr("join")
		ßkeys := πg.InternStr("keys")
		ßlatin1 := πg.InternStr("latin1")
		ßlen := πg.InternStr("len")
		ßlower := πg.InternStr("lower")
		ßmailto := πg.InternStr("mailto")
		ßmap := πg.InternStr("map")
		ßmin := πg.InternStr("min")
		ßmms := πg.InternStr("mms")
		ßnetloc := πg.InternStr("netloc")
		ßnews := πg.InternStr("news")
		ßnfs := πg.InternStr("nfs")
		ßnntp := πg.InternStr("nntp")
		ßnon_hierarchical := πg.InternStr("non_hierarchical")
		ßobject := πg.InternStr("object")
		ßoperator := πg.InternStr("operator")
		ßparams := πg.InternStr("params")
		ßparse_qs := πg.InternStr("parse_qs")
		ßparse_qsl := πg.InternStr("parse_qsl")
		ßpassword := πg.InternStr("password")
		ßpath := πg.InternStr("path")
		ßpop := πg.InternStr("pop")
		ßport := πg.InternStr("port")
		ßproperty := πg.InternStr("property")
		ßprospero := πg.InternStr("prospero")
		ßquery := πg.InternStr("query")
		ßrange := πg.InternStr("range")
		ßre := πg.InternStr("re")
		ßremove := πg.InternStr("remove")
		ßreplace := πg.InternStr("replace")
		ßrfind := πg.InternStr("rfind")
		ßrsplit := πg.InternStr("rsplit")
		ßrsync := πg.InternStr("rsync")
		ßrtsp := πg.InternStr("rtsp")
		ßrtspu := πg.InternStr("rtspu")
		ßscheme := πg.InternStr("scheme")
		ßscheme_chars := πg.InternStr("scheme_chars")
		ßsftp := πg.InternStr("sftp")
		ßshttp := πg.InternStr("shttp")
		ßsip := πg.InternStr("sip")
		ßsips := πg.InternStr("sips")
		ßsnews := πg.InternStr("snews")
		ßsplit := πg.InternStr("split")
		ßstr := πg.InternStr("str")
		ßsvn := πg.InternStr("svn")
		ßtel := πg.InternStr("tel")
		ßtelnet := πg.InternStr("telnet")
		ßtuple := πg.InternStr("tuple")
		ßtype := πg.InternStr("type")
		ßunicode := πg.InternStr("unicode")
		ßunquote := πg.InternStr("unquote")
		ßurldefrag := πg.InternStr("urldefrag")
		ßurljoin := πg.InternStr("urljoin")
		ßurlparse := πg.InternStr("urlparse")
		ßurlsplit := πg.InternStr("urlsplit")
		ßurlunparse := πg.InternStr("urlunparse")
		ßurlunsplit := πg.InternStr("urlunsplit")
		ßusername := πg.InternStr("username")
		ßuses_fragment := πg.InternStr("uses_fragment")
		ßuses_netloc := πg.InternStr("uses_netloc")
		ßuses_params := πg.InternStr("uses_params")
		ßuses_query := πg.InternStr("uses_query")
		ßuses_relative := πg.InternStr("uses_relative")
		ßwais := πg.InternStr("wais")
		ßzip := πg.InternStr("zip")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Dict
		_ = πTemp004
		var πTemp005 []πg.Param
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
		var πTemp020 *πg.Object
		_ = πTemp020
		var πTemp021 *πg.Object
		_ = πTemp021
		var πTemp022 *πg.Object
		_ = πTemp022
		var πTemp023 *πg.Object
		_ = πTemp023
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			case 2: goto Label2
			default: panic("unexpected function state")
			}
			// line 1: """Parse (absolute and relative) URLs.
			πF.SetLineno(1)
			// line 1: """Parse (absolute and relative) URLs.
			πF.SetLineno(1)
			if πE = πF.Globals().SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Parse (absolute and relative) URLs.\n\nurlparse module is based upon the following RFC specifications.\n\nRFC 3986 (STD66): \"Uniform Resource Identifiers\" by T. Berners-Lee, R. Fielding\nand L.  Masinter, January 2005.\n\nRFC 2732 : \"Format for Literal IPv6 Addresses in URL's by R.Hinden, B.Carpenter\nand L.Masinter, December 1999.\n\nRFC 2396:  \"Uniform Resource Identifiers (URI)\": Generic Syntax by T.\nBerners-Lee, R. Fielding, and L. Masinter, August 1998.\n\nRFC 2368: \"The mailto URL scheme\", by P.Hoffman , L Masinter, J. Zwinski, July 1998.\n\nRFC 1808: \"Relative Uniform Resource Locators\", by R. Fielding, UC Irvine, June\n1995.\n\nRFC 1738: \"Uniform Resource Locators (URL)\" by T. Berners-Lee, L. Masinter, M.\nMcCahill, December 1994\n\nRFC 3986 is considered the current standard and any future changes to\nurlparse module should conform with it.  The urlparse module is\ncurrently not entirely compliant with this RFC due to defacto\nscenarios for parsing, and for backward compatibility purposes, some\nparsing quirks from older RFCs are retained. The testcases in\ntest_urlparse.py provides a good indicator of parsing behavior.\n\n").ToObject()); πE != nil {
				continue
			}
			// line 31: import re
			πF.SetLineno(31)
			if πTemp002, πE = πg.ImportModule(πF, "re"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßre.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 33: import operator
			πF.SetLineno(33)
			if πTemp002, πE = πg.ImportModule(πF, "operator"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßoperator.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 34: _itemgetter = operator.itemgetter
			πF.SetLineno(34)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßoperator); πE != nil {
				continue
			}
			if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßitemgetter, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_itemgetter.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 35: _property = property
			πF.SetLineno(35)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßproperty); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_property.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 36: _tuple = tuple
			πF.SetLineno(36)
			if πTemp001, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_tuple.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 38: __all__ = ["urlparse", "urlunparse", "urljoin", "urldefrag",
			πF.SetLineno(38)
			πTemp002 = make([]*πg.Object, 8)
			πTemp002[0] = ßurlparse.ToObject()
			πTemp002[1] = ßurlunparse.ToObject()
			πTemp002[2] = ßurljoin.ToObject()
			πTemp002[3] = ßurldefrag.ToObject()
			πTemp002[4] = ßurlsplit.ToObject()
			πTemp002[5] = ßurlunsplit.ToObject()
			πTemp002[6] = ßparse_qs.ToObject()
			πTemp002[7] = ßparse_qsl.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ß__all__.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 42: uses_relative = ['ftp', 'http', 'gopher', 'nntp', 'imap',
			πF.SetLineno(42)
			πTemp002 = make([]*πg.Object, 17)
			πTemp002[0] = ßftp.ToObject()
			πTemp002[1] = ßhttp.ToObject()
			πTemp002[2] = ßgopher.ToObject()
			πTemp002[3] = ßnntp.ToObject()
			πTemp002[4] = ßimap.ToObject()
			πTemp002[5] = ßwais.ToObject()
			πTemp002[6] = ßfile.ToObject()
			πTemp002[7] = ßhttps.ToObject()
			πTemp002[8] = ßshttp.ToObject()
			πTemp002[9] = ßmms.ToObject()
			πTemp002[10] = ßprospero.ToObject()
			πTemp002[11] = ßrtsp.ToObject()
			πTemp002[12] = ßrtspu.ToObject()
			πTemp002[13] = ß.ToObject()
			πTemp002[14] = ßsftp.ToObject()
			πTemp002[15] = ßsvn.ToObject()
			πTemp002[16] = πg.NewStr("svn+ssh").ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßuses_relative.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 46: uses_netloc = ['ftp', 'http', 'gopher', 'nntp', 'telnet',
			πF.SetLineno(46)
			πTemp002 = make([]*πg.Object, 23)
			πTemp002[0] = ßftp.ToObject()
			πTemp002[1] = ßhttp.ToObject()
			πTemp002[2] = ßgopher.ToObject()
			πTemp002[3] = ßnntp.ToObject()
			πTemp002[4] = ßtelnet.ToObject()
			πTemp002[5] = ßimap.ToObject()
			πTemp002[6] = ßwais.ToObject()
			πTemp002[7] = ßfile.ToObject()
			πTemp002[8] = ßmms.ToObject()
			πTemp002[9] = ßhttps.ToObject()
			πTemp002[10] = ßshttp.ToObject()
			πTemp002[11] = ßsnews.ToObject()
			πTemp002[12] = ßprospero.ToObject()
			πTemp002[13] = ßrtsp.ToObject()
			πTemp002[14] = ßrtspu.ToObject()
			πTemp002[15] = ßrsync.ToObject()
			πTemp002[16] = ß.ToObject()
			πTemp002[17] = ßsvn.ToObject()
			πTemp002[18] = πg.NewStr("svn+ssh").ToObject()
			πTemp002[19] = ßsftp.ToObject()
			πTemp002[20] = ßnfs.ToObject()
			πTemp002[21] = ßgit.ToObject()
			πTemp002[22] = πg.NewStr("git+ssh").ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßuses_netloc.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 50: uses_params = ['ftp', 'hdl', 'prospero', 'http', 'imap',
			πF.SetLineno(50)
			πTemp002 = make([]*πg.Object, 15)
			πTemp002[0] = ßftp.ToObject()
			πTemp002[1] = ßhdl.ToObject()
			πTemp002[2] = ßprospero.ToObject()
			πTemp002[3] = ßhttp.ToObject()
			πTemp002[4] = ßimap.ToObject()
			πTemp002[5] = ßhttps.ToObject()
			πTemp002[6] = ßshttp.ToObject()
			πTemp002[7] = ßrtsp.ToObject()
			πTemp002[8] = ßrtspu.ToObject()
			πTemp002[9] = ßsip.ToObject()
			πTemp002[10] = ßsips.ToObject()
			πTemp002[11] = ßmms.ToObject()
			πTemp002[12] = ß.ToObject()
			πTemp002[13] = ßsftp.ToObject()
			πTemp002[14] = ßtel.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßuses_params.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 56: non_hierarchical = ['gopher', 'hdl', 'mailto', 'news',
			πF.SetLineno(56)
			πTemp002 = make([]*πg.Object, 10)
			πTemp002[0] = ßgopher.ToObject()
			πTemp002[1] = ßhdl.ToObject()
			πTemp002[2] = ßmailto.ToObject()
			πTemp002[3] = ßnews.ToObject()
			πTemp002[4] = ßtelnet.ToObject()
			πTemp002[5] = ßwais.ToObject()
			πTemp002[6] = ßimap.ToObject()
			πTemp002[7] = ßsnews.ToObject()
			πTemp002[8] = ßsip.ToObject()
			πTemp002[9] = ßsips.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßnon_hierarchical.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 58: uses_query = ['http', 'wais', 'imap', 'https', 'shttp', 'mms',
			πF.SetLineno(58)
			πTemp002 = make([]*πg.Object, 12)
			πTemp002[0] = ßhttp.ToObject()
			πTemp002[1] = ßwais.ToObject()
			πTemp002[2] = ßimap.ToObject()
			πTemp002[3] = ßhttps.ToObject()
			πTemp002[4] = ßshttp.ToObject()
			πTemp002[5] = ßmms.ToObject()
			πTemp002[6] = ßgopher.ToObject()
			πTemp002[7] = ßrtsp.ToObject()
			πTemp002[8] = ßrtspu.ToObject()
			πTemp002[9] = ßsip.ToObject()
			πTemp002[10] = ßsips.ToObject()
			πTemp002[11] = ß.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßuses_query.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 60: uses_fragment = ['ftp', 'hdl', 'http', 'gopher', 'news',
			πF.SetLineno(60)
			πTemp002 = make([]*πg.Object, 13)
			πTemp002[0] = ßftp.ToObject()
			πTemp002[1] = ßhdl.ToObject()
			πTemp002[2] = ßhttp.ToObject()
			πTemp002[3] = ßgopher.ToObject()
			πTemp002[4] = ßnews.ToObject()
			πTemp002[5] = ßnntp.ToObject()
			πTemp002[6] = ßwais.ToObject()
			πTemp002[7] = ßhttps.ToObject()
			πTemp002[8] = ßshttp.ToObject()
			πTemp002[9] = ßsnews.ToObject()
			πTemp002[10] = ßfile.ToObject()
			πTemp002[11] = ßprospero.ToObject()
			πTemp002[12] = ß.ToObject()
			πTemp001 = πg.NewList(πTemp002...).ToObject()
			if πE = πF.Globals().SetItem(πF, ßuses_fragment.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 65: scheme_chars = ('abcdefghijklmnopqrstuvwxyz'
			πF.SetLineno(65)
			if πE = πF.Globals().SetItem(πF, ßscheme_chars.ToObject(), πg.NewStr("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-.").ToObject()); πE != nil {
				continue
			}
			// line 70: MAX_CACHE_SIZE = 20
			πF.SetLineno(70)
			if πE = πF.Globals().SetItem(πF, ßMAX_CACHE_SIZE.ToObject(), πg.NewInt(20).ToObject()); πE != nil {
				continue
			}
			// line 71: _parse_cache = {}
			πF.SetLineno(71)
			πTemp004 = πg.NewDict()
			πTemp001 = πTemp004.ToObject()
			if πE = πF.Globals().SetItem(πF, ß_parse_cache.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 74: def clear_cache():
			πF.SetLineno(74)
			πTemp005 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("clear_cache", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 75: """Clear the parse cache."""
					πF.SetLineno(75)
					// line 76: _parse_cache.clear()
					πF.SetLineno(76)
					if πTemp001, πE = πg.ResolveGlobal(πF, ß_parse_cache); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßclear_cache.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 75: """Clear the parse cache."""
			πF.SetLineno(75)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Clear the parse cache.").ToObject()); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßclear_cache); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp006, ß__doc__, πTemp003); πE != nil {
				continue
			}
			// line 79: class ResultMixin(object):
			πF.SetLineno(79)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ResultMixin", "/usr/lib/python2.7/urlparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
				_ = πClass
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []πg.Param
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
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 80: """Shared methods for the parsed result objects."""
					πF.SetLineno(80)
					// line 80: """Shared methods for the parsed result objects."""
					πF.SetLineno(80)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("Shared methods for the parsed result objects.").ToObject()); πE != nil {
						continue
					}
					// line 83: def username(self):
					πF.SetLineno(83)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("username", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnetloc *πg.Object = πg.UnboundLocal; _ = µnetloc
						var µuserinfo *πg.Object = πg.UnboundLocal; _ = µuserinfo
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 84: netloc = self.netloc
							πF.SetLineno(84)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnetloc, nil); πE != nil {
								continue
							}
							µnetloc = πTemp001
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µnetloc, πg.NewStr("@").ToObject()); πE != nil {
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
							// line 85: if "@" in netloc:
							πF.SetLineno(85)
						Label1:
							// line 86: userinfo = netloc.rsplit("@", 1)[0]
							πF.SetLineno(86)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("@").ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnetloc, ßrsplit, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µuserinfo = πTemp003
							if πE = πg.CheckLocal(πF, µuserinfo, "userinfo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µuserinfo, πg.NewStr(":").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 87: if ":" in userinfo:
							πF.SetLineno(87)
						Label3:
							// line 88: userinfo = userinfo.split(":", 1)[0]
							πF.SetLineno(88)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr(":").ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µuserinfo, "userinfo"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µuserinfo, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µuserinfo = πTemp003
							goto Label4
						Label4:
							// line 89: return userinfo
							πF.SetLineno(89)
							if πE = πg.CheckLocal(πF, µuserinfo, "userinfo"); πE != nil {
								continue
							}
							πR = µuserinfo
							continue
							goto Label2
						Label2:
							// line 90: return None
							πF.SetLineno(90)
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
					if πE = πClass.SetItem(πF, ßusername.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 91: username = property(username)
					πF.SetLineno(91)
					πTemp003 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßusername); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßusername.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 94: def password(self):
					πF.SetLineno(94)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("password", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnetloc *πg.Object = πg.UnboundLocal; _ = µnetloc
						var µuserinfo *πg.Object = πg.UnboundLocal; _ = µuserinfo
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 95: netloc = self.netloc
							πF.SetLineno(95)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µself, ßnetloc, nil); πE != nil {
								continue
							}
							µnetloc = πTemp001
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µnetloc, πg.NewStr("@").ToObject()); πE != nil {
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
							// line 96: if "@" in netloc:
							πF.SetLineno(96)
						Label1:
							// line 97: userinfo = netloc.rsplit("@", 1)[0]
							πF.SetLineno(97)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr("@").ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µnetloc, ßrsplit, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							µuserinfo = πTemp003
							if πE = πg.CheckLocal(πF, µuserinfo, "userinfo"); πE != nil {
								continue
							}
							if πTemp002, πE = πg.Contains(πF, µuserinfo, πg.NewStr(":").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp002).ToObject()
							if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp002 {
								goto Label3
							}
							goto Label4
							// line 98: if ":" in userinfo:
							πF.SetLineno(98)
						Label3:
							// line 99: return userinfo.split(":", 1)[1]
							πF.SetLineno(99)
							πTemp001 = πg.NewInt(1).ToObject()
							πTemp004 = πF.MakeArgs(2)
							πTemp004[0] = πg.NewStr(":").ToObject()
							πTemp004[1] = πg.NewInt(1).ToObject()
							if πE = πg.CheckLocal(πF, µuserinfo, "userinfo"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, µuserinfo, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp004)
							if πTemp003, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
								continue
							}
							πR = πTemp003
							continue
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 100: return None
							πF.SetLineno(100)
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
					if πE = πClass.SetItem(πF, ßpassword.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 101: password = property(password)
					πF.SetLineno(101)
					πTemp003 = πF.MakeArgs(1)
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßpassword); πE != nil {
						continue
					}
					πTemp003[0] = πTemp005
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßpassword.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 104: def hostname(self):
					πF.SetLineno(104)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("hostname", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnetloc *πg.Object = πg.UnboundLocal; _ = µnetloc
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
						var πTemp007 bool
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
							// line 105: netloc = self.netloc.split('@')[-1]
							πF.SetLineno(105)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("@").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µself, ßnetloc, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
								continue
							}
							µnetloc = πTemp002
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µnetloc, πg.NewStr("[").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							πTemp001 = πTemp002
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp006 {
								goto Label1
							}
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.Contains(πF, µnetloc, πg.NewStr("]").ToObject()); πE != nil {
								continue
							}
							πTemp002 = πg.GetBool(πTemp007).ToObject()
							πTemp001 = πTemp002
						Label1:
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label2
							}
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp006, πE = πg.Contains(πF, µnetloc, πg.NewStr(":").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp006).ToObject()
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label3
							}
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Eq(πF, µnetloc, ß.ToObject()); πE != nil {
								continue
							}
							if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp006 {
								goto Label4
							}
							goto Label5
							// line 106: if '[' in netloc and ']' in netloc:
							πF.SetLineno(106)
						Label2:
							// line 107: return netloc.split(']')[0][1:].lower()
							πF.SetLineno(107)
							if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
								continue
							}
							πTemp004 = πg.NewInt(0).ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("]").ToObject()
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, µnetloc, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp009, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp005, πE = πg.GetItem(πF, πTemp009, πTemp004); πE != nil {
								continue
							}
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßlower, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label6
							// line 108: elif ':' in netloc:
							πF.SetLineno(108)
						Label3:
							// line 109: return netloc.split(':')[0].lower()
							πF.SetLineno(109)
							πTemp001 = πg.NewInt(0).ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr(":").ToObject()
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnetloc, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßlower, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label6
							// line 110: elif netloc == '':
							πF.SetLineno(110)
						Label4:
							// line 111: return None
							πF.SetLineno(111)
							if πTemp001, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
								continue
							}
							πR = πTemp001
							continue
							goto Label6
						Label5:
							// line 113: return netloc.lower()
							πF.SetLineno(113)
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.GetAttr(πF, µnetloc, ßlower, nil); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
								continue
							}
							πR = πTemp002
							continue
							goto Label6
						Label6:
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ßhostname.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 114: hostname = property(hostname)
					πF.SetLineno(114)
					πTemp003 = πF.MakeArgs(1)
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßhostname); πE != nil {
						continue
					}
					πTemp003[0] = πTemp006
					if πTemp006, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßhostname.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 117: def port(self):
					πF.SetLineno(117)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp006 = πg.NewFunction(πg.NewCode("port", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var µnetloc *πg.Object = πg.UnboundLocal; _ = µnetloc
						var µport *πg.Object = πg.UnboundLocal; _ = µport
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
						var πTemp006 []*πg.Object
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
							default: panic("unexpected function state")
							}
							// line 118: netloc = self.netloc.split('@')[-1].split(']')[-1]
							πF.SetLineno(118)
							if πTemp002, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp001 = πTemp002
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr("]").ToObject()
							if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
								continue
							}
							πTemp004 = πTemp005
							πTemp006 = πF.MakeArgs(1)
							πTemp006[0] = πg.NewStr("@").ToObject()
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp007, πE = πg.GetAttr(πF, µself, ßnetloc, nil); πE != nil {
								continue
							}
							if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp007, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp006)
							if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							µnetloc = πTemp002
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.Contains(πF, µnetloc, πg.NewStr(":").ToObject()); πE != nil {
								continue
							}
							πTemp001 = πg.GetBool(πTemp009).ToObject()
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label1
							}
							goto Label2
							// line 119: if ':' in netloc:
							πF.SetLineno(119)
						Label1:
							// line 120: port = netloc.split(':')[1]
							πF.SetLineno(120)
							πTemp001 = πg.NewInt(1).ToObject()
							πTemp003 = πF.MakeArgs(1)
							πTemp003[0] = πg.NewStr(":").ToObject()
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µnetloc, ßsplit, nil); πE != nil {
								continue
							}
							if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							if πTemp002, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
								continue
							}
							µport = πTemp002
							if πE = πg.CheckLocal(πF, µport, "port"); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, µport); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label3
							}
							goto Label4
							// line 121: if port:
							πF.SetLineno(121)
						Label3:
							// line 122: port = int(port, 10)
							πF.SetLineno(122)
							πTemp003 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µport, "port"); πE != nil {
								continue
							}
							πTemp003[0] = µport
							πTemp003[1] = πg.NewInt(10).ToObject()
							if πTemp001, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
								continue
							}
							if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp003)
							µport = πTemp002
							if πE = πg.CheckLocal(πF, µport, "port"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.LE(πF, πg.NewInt(0).ToObject(), µport); πE != nil {
								continue
							}
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if !πTemp009 {
								goto Label5
							}
							if πTemp001, πE = πg.LE(πF, µport, πg.NewInt(65535).ToObject()); πE != nil {
								continue
							}
						Label5:
							if πTemp009, πE = πg.IsTrue(πF, πTemp001); πE != nil {
								continue
							}
							if πTemp009 {
								goto Label6
							}
							goto Label7
							// line 124: if (0 <= port <= 65535):
							πF.SetLineno(124)
						Label6:
							// line 125: return port
							πF.SetLineno(125)
							if πE = πg.CheckLocal(πF, µport, "port"); πE != nil {
								continue
							}
							πR = µport
							continue
							goto Label7
						Label7:
							goto Label4
						Label4:
							goto Label2
						Label2:
							// line 126: return None
							πF.SetLineno(126)
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
					if πE = πClass.SetItem(πF, ßport.ToObject(), πTemp006); πE != nil {
						continue
					}
					// line 127: port = property(port)
					πF.SetLineno(127)
					πTemp003 = πF.MakeArgs(1)
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßport); πE != nil {
						continue
					}
					πTemp003[0] = πTemp007
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ßproperty); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πClass.SetItem(πF, ßport.ToObject(), πTemp008); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("ResultMixin").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßResultMixin.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 130: class _SplitResult(tuple):
			πF.SetLineno(130)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_SplitResult", "/usr/lib/python2.7/urlparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 []*πg.Object
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
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πTemp013 πg.KWArgs
				_ = πTemp013
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 131: 'SplitResult(scheme, netloc, path, query, fragment)'
					πF.SetLineno(131)
					// line 131: 'SplitResult(scheme, netloc, path, query, fragment)'
					πF.SetLineno(131)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("SplitResult(scheme, netloc, path, query, fragment)").ToObject()); πE != nil {
						continue
					}
					// line 133: __slots__ = ()
					πF.SetLineno(133)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 135: _fields = ('scheme', 'netloc', 'path', 'query', 'fragment')
					πF.SetLineno(135)
					πTemp001 = πg.NewTuple5(ßscheme.ToObject(), ßnetloc.ToObject(), ßpath.ToObject(), ßquery.ToObject(), ßfragment.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß_fields.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 137: def __new__(_cls, scheme, netloc, path, query, fragment):
					πF.SetLineno(137)
					πTemp002 = make([]πg.Param, 6)
					πTemp002[0] = πg.Param{Name: "_cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "scheme", Def: nil}
					πTemp002[2] = πg.Param{Name: "netloc", Def: nil}
					πTemp002[3] = πg.Param{Name: "path", Def: nil}
					πTemp002[4] = πg.Param{Name: "query", Def: nil}
					πTemp002[5] = πg.Param{Name: "fragment", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µ_cls *πg.Object = πArgs[0]; _ = µ_cls
						var µscheme *πg.Object = πArgs[1]; _ = µscheme
						var µnetloc *πg.Object = πArgs[2]; _ = µnetloc
						var µpath *πg.Object = πArgs[3]; _ = µpath
						var µquery *πg.Object = πArgs[4]; _ = µquery
						var µfragment *πg.Object = πArgs[5]; _ = µfragment
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
							// line 138: 'Create new instance of SplitResult(scheme, netloc, path, query, fragment)'
							πF.SetLineno(138)
							// line 139: return _tuple.__new__(_cls, (scheme, netloc, path, query, fragment))
							πF.SetLineno(139)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µ_cls, "_cls"); πE != nil {
								continue
							}
							πTemp001[0] = µ_cls
							if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple5(µscheme, µnetloc, µpath, µquery, µfragment).ToObject()
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_tuple); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 138: 'Create new instance of SplitResult(scheme, netloc, path, query, fragment)'
					πF.SetLineno(138)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Create new instance of SplitResult(scheme, netloc, path, query, fragment)").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__new__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 142: def _make(cls, iterable, new=tuple.__new__, len=len):
					πF.SetLineno(142)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__new__, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "new", Def: πTemp005}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlen); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "len", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("_make", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µnew *πg.Object = πArgs[2]; _ = µnew
						var µlen *πg.Object = πArgs[3]; _ = µlen
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 143: 'Make a new SplitResult object from a sequence or iterable'
							πF.SetLineno(143)
							// line 144: result = new(cls, iterable)
							πF.SetLineno(144)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp001[1] = µiterable
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πTemp002, πE = µnew.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							if πTemp003, πE = µlen.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp003, πg.NewInt(5).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 145: if len(result) != 5:
							πF.SetLineno(145)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp005[0] = µresult
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							if πTemp003, πE = µlen.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Expected 5 arguments, got %d").ToObject(), πTemp003); πE != nil {
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
							// line 146: raise TypeError('Expected 5 arguments, got %d' % len(result))
							πF.SetLineno(146)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 147: return result
							πF.SetLineno(147)
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
					if πE = πClass.SetItem(πF, ß_make.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 143: 'Make a new SplitResult object from a sequence or iterable'
					πF.SetLineno(143)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Make a new SplitResult object from a sequence or iterable").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß_make); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 148: _make = classmethod(_make)
					πF.SetLineno(148)
					πTemp006 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_make); πE != nil {
						continue
					}
					πTemp006[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ß_make.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 150: def __repr__(self):
					πF.SetLineno(150)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 151: 'Return a nicely formatted representation string'
							πF.SetLineno(151)
							// line 152: return 'SplitResult(scheme=%r, netloc=%r, path=%r, query=%r, fragment=%r)' % self
							πF.SetLineno(152)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("SplitResult(scheme=%r, netloc=%r, path=%r, query=%r, fragment=%r)").ToObject(), µself); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 151: 'Return a nicely formatted representation string'
					πF.SetLineno(151)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return a nicely formatted representation string").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ß__repr__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 154: def _asdict(self):
					πF.SetLineno(154)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_asdict", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 155: 'Return a new OrderedDict which maps field names to their values'
							πF.SetLineno(155)
							// line 156: return OrderedDict(zip(self._fields, self))
							πF.SetLineno(156)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_fields, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßOrderedDict); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_asdict.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 155: 'Return a new OrderedDict which maps field names to their values'
					πF.SetLineno(155)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Return a new OrderedDict which maps field names to their values").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ß_asdict); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 158: def _replace(_self, **kwds):
					πF.SetLineno(158)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "_self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_replace", "/usr/lib/python2.7/urlparse.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µ_self *πg.Object = πArgs[0]; _ = µ_self
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 159: 'Return a new SplitResult object replacing specified fields with new values'
							πF.SetLineno(159)
							// line 160: result = _self._make(map(kwds.pop, ('scheme', 'netloc', 'path', 'query', 'fragment'), _self))
							πF.SetLineno(160)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewTuple5(ßscheme.ToObject(), ßnetloc.ToObject(), ßpath.ToObject(), ßquery.ToObject(), ßfragment.ToObject()).ToObject()
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µ_self, "_self"); πE != nil {
								continue
							}
							πTemp002[2] = µ_self
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µ_self, "_self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µ_self, ß_make, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp004
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µkwds); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 161: if kwds:
							πF.SetLineno(161)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µkwds, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Got unexpected field names: %r").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 162: raise ValueError('Got unexpected field names: %r' % kwds.keys())
							πF.SetLineno(162)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 163: return result
							πF.SetLineno(163)
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
					if πE = πClass.SetItem(πF, ß_replace.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 159: 'Return a new SplitResult object replacing specified fields with new values'
					πF.SetLineno(159)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Return a new SplitResult object replacing specified fields with new values").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_replace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 165: def __getnewargs__(self):
					πF.SetLineno(165)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__getnewargs__", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 166: 'Return self as a plain tuple.  Used by copy and pickle.'
							πF.SetLineno(166)
							// line 167: return tuple(self)
							πF.SetLineno(167)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
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
					if πE = πClass.SetItem(πF, ß__getnewargs__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 166: 'Return self as a plain tuple.  Used by copy and pickle.'
					πF.SetLineno(166)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Return self as a plain tuple.  Used by copy and pickle.").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß__getnewargs__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 169: __dict__ = _property(_asdict)
					πF.SetLineno(169)
					πTemp006 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_asdict); πE != nil {
						continue
					}
					πTemp006[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ß__dict__.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 171: def __getstate__(self):
					πF.SetLineno(171)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__getstate__", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 172: 'Exclude the OrderedDict from pickling'
							πF.SetLineno(172)
							// line 173: pass
							πF.SetLineno(173)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__getstate__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 172: 'Exclude the OrderedDict from pickling'
					πF.SetLineno(172)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Exclude the OrderedDict from pickling").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ß__getstate__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 175: scheme = _property(_itemgetter(0), doc='Alias for field number 0')
					πF.SetLineno(175)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(0).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 0").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßscheme.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 177: netloc = _property(_itemgetter(1), doc='Alias for field number 1')
					πF.SetLineno(177)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(1).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 1").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßnetloc.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 179: path = _property(_itemgetter(2), doc='Alias for field number 2')
					πF.SetLineno(179)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(2).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 2").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßpath.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 181: query = _property(_itemgetter(3), doc='Alias for field number 3')
					πF.SetLineno(181)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(3).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 3").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßquery.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 183: fragment = _property(_itemgetter(4), doc='Alias for field number 4')
					πF.SetLineno(183)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(4).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 4").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßfragment.ToObject(), πTemp011); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("_SplitResult").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_SplitResult.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 186: class SplitResult(_SplitResult, ResultMixin):
			πF.SetLineno(186)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_SplitResult); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßResultMixin); πE != nil {
				continue
			}
			πTemp002[1] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("SplitResult", "/usr/lib/python2.7/urlparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 188: __slots__ = ()
					πF.SetLineno(188)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 190: def geturl(self):
					πF.SetLineno(190)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("geturl", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 191: return urlunsplit(self)
							πF.SetLineno(191)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßurlunsplit); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgeturl.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("SplitResult").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßSplitResult.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 193: class _ParseResult(tuple):
			πF.SetLineno(193)
			πTemp002 = make([]*πg.Object, 1)
			if πTemp007, πE = πg.ResolveGlobal(πF, ßtuple); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("_ParseResult", "/usr/lib/python2.7/urlparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
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
				var πTemp006 []*πg.Object
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
				var πTemp012 []*πg.Object
				_ = πTemp012
				var πTemp013 πg.KWArgs
				_ = πTemp013
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 194: 'ParseResult(scheme, netloc, path, params, query, fragment)'
					πF.SetLineno(194)
					// line 194: 'ParseResult(scheme, netloc, path, params, query, fragment)'
					πF.SetLineno(194)
					if πE = πClass.SetItem(πF, ß__doc__.ToObject(), πg.NewStr("ParseResult(scheme, netloc, path, params, query, fragment)").ToObject()); πE != nil {
						continue
					}
					// line 196: __slots__ = ()
					πF.SetLineno(196)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 198: _fields = ('scheme', 'netloc', 'path', 'params', 'query', 'fragment')
					πF.SetLineno(198)
					πTemp001 = πg.NewTuple6(ßscheme.ToObject(), ßnetloc.ToObject(), ßpath.ToObject(), ßparams.ToObject(), ßquery.ToObject(), ßfragment.ToObject()).ToObject()
					if πE = πClass.SetItem(πF, ß_fields.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 200: def __new__(_cls, scheme, netloc, path, params, query, fragment):
					πF.SetLineno(200)
					πTemp002 = make([]πg.Param, 7)
					πTemp002[0] = πg.Param{Name: "_cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "scheme", Def: nil}
					πTemp002[2] = πg.Param{Name: "netloc", Def: nil}
					πTemp002[3] = πg.Param{Name: "path", Def: nil}
					πTemp002[4] = πg.Param{Name: "params", Def: nil}
					πTemp002[5] = πg.Param{Name: "query", Def: nil}
					πTemp002[6] = πg.Param{Name: "fragment", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("__new__", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µ_cls *πg.Object = πArgs[0]; _ = µ_cls
						var µscheme *πg.Object = πArgs[1]; _ = µscheme
						var µnetloc *πg.Object = πArgs[2]; _ = µnetloc
						var µpath *πg.Object = πArgs[3]; _ = µpath
						var µparams *πg.Object = πArgs[4]; _ = µparams
						var µquery *πg.Object = πArgs[5]; _ = µquery
						var µfragment *πg.Object = πArgs[6]; _ = µfragment
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
							// line 201: 'Create new instance of ParseResult(scheme, netloc, path, params, query, fragment)'
							πF.SetLineno(201)
							// line 202: return _tuple.__new__(_cls, (scheme, netloc, path, params, query, fragment))
							πF.SetLineno(202)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µ_cls, "_cls"); πE != nil {
								continue
							}
							πTemp001[0] = µ_cls
							if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µparams, "params"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
								continue
							}
							if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
								continue
							}
							πTemp002 = πg.NewTuple6(µscheme, µnetloc, µpath, µparams, µquery, µfragment).ToObject()
							πTemp001[1] = πTemp002
							if πTemp002, πE = πg.ResolveGlobal(πF, ß_tuple); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, πTemp002, ß__new__, nil); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__new__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 201: 'Create new instance of ParseResult(scheme, netloc, path, params, query, fragment)'
					πF.SetLineno(201)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πg.NewStr("Create new instance of ParseResult(scheme, netloc, path, params, query, fragment)").ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß__new__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp004, ß__doc__, πTemp003); πE != nil {
						continue
					}
					// line 205: def _make(cls, iterable, new=tuple.__new__, len=len):
					πF.SetLineno(205)
					πTemp002 = make([]πg.Param, 4)
					πTemp002[0] = πg.Param{Name: "cls", Def: nil}
					πTemp002[1] = πg.Param{Name: "iterable", Def: nil}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßtuple); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp004, ß__new__, nil); πE != nil {
						continue
					}
					πTemp002[2] = πg.Param{Name: "new", Def: πTemp005}
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßlen); πE != nil {
						continue
					}
					πTemp002[3] = πg.Param{Name: "len", Def: πTemp004}
					πTemp003 = πg.NewFunction(πg.NewCode("_make", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µcls *πg.Object = πArgs[0]; _ = µcls
						var µiterable *πg.Object = πArgs[1]; _ = µiterable
						var µnew *πg.Object = πArgs[2]; _ = µnew
						var µlen *πg.Object = πArgs[3]; _ = µlen
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 206: 'Make a new ParseResult object from a sequence or iterable'
							πF.SetLineno(206)
							// line 207: result = new(cls, iterable)
							πF.SetLineno(207)
							πTemp001 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µcls, "cls"); πE != nil {
								continue
							}
							πTemp001[0] = µcls
							if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
								continue
							}
							πTemp001[1] = µiterable
							if πE = πg.CheckLocal(πF, µnew, "new"); πE != nil {
								continue
							}
							if πTemp002, πE = µnew.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp002
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp001[0] = µresult
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							if πTemp003, πE = µlen.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							if πTemp002, πE = πg.NE(πF, πTemp003, πg.NewInt(6).ToObject()); πE != nil {
								continue
							}
							if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
								continue
							}
							if πTemp004 {
								goto Label1
							}
							goto Label2
							// line 208: if len(result) != 6:
							πF.SetLineno(208)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							πTemp005 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µresult, "result"); πE != nil {
								continue
							}
							πTemp005[0] = µresult
							if πE = πg.CheckLocal(πF, µlen, "len"); πE != nil {
								continue
							}
							if πTemp003, πE = µlen.Call(πF, πTemp005, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp005)
							if πTemp002, πE = πg.Mod(πF, πg.NewStr("Expected 6 arguments, got %d").ToObject(), πTemp003); πE != nil {
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
							// line 209: raise TypeError('Expected 6 arguments, got %d' % len(result))
							πF.SetLineno(209)
							πE = πF.Raise(πTemp003, nil, nil)
							continue
							goto Label2
						Label2:
							// line 210: return result
							πF.SetLineno(210)
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
					if πE = πClass.SetItem(πF, ß_make.ToObject(), πTemp003); πE != nil {
						continue
					}
					// line 206: 'Make a new ParseResult object from a sequence or iterable'
					πF.SetLineno(206)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewStr("Make a new ParseResult object from a sequence or iterable").ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveClass(πF, πClass, nil, ß_make); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp005, ß__doc__, πTemp004); πE != nil {
						continue
					}
					// line 211: _make = classmethod(_make)
					πF.SetLineno(211)
					πTemp006 = πF.MakeArgs(1)
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ß_make); πE != nil {
						continue
					}
					πTemp006[0] = πTemp004
					if πTemp004, πE = πg.ResolveClass(πF, πClass, nil, ßclassmethod); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ß_make.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 213: def __repr__(self):
					πF.SetLineno(213)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp004 = πg.NewFunction(πg.NewCode("__repr__", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 214: 'Return a nicely formatted representation string'
							πF.SetLineno(214)
							// line 215: return 'ParseResult(scheme=%r, netloc=%r, path=%r, params=%r, query=%r, fragment=%r)' % self
							πF.SetLineno(215)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp001, πE = πg.Mod(πF, πg.NewStr("ParseResult(scheme=%r, netloc=%r, path=%r, params=%r, query=%r, fragment=%r)").ToObject(), µself); πE != nil {
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
					if πE = πClass.SetItem(πF, ß__repr__.ToObject(), πTemp004); πE != nil {
						continue
					}
					// line 214: 'Return a nicely formatted representation string'
					πF.SetLineno(214)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πg.NewStr("Return a nicely formatted representation string").ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.ResolveClass(πF, πClass, nil, ß__repr__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp005); πE != nil {
						continue
					}
					// line 217: def _asdict(self):
					πF.SetLineno(217)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp005 = πg.NewFunction(πg.NewCode("_asdict", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
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
							// line 218: 'Return a new OrderedDict which maps field names to their values'
							πF.SetLineno(218)
							// line 219: return OrderedDict(zip(self._fields, self))
							πF.SetLineno(219)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(2)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µself, ß_fields, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp002[1] = µself
							if πTemp003, πE = πg.ResolveGlobal(πF, ßzip); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πTemp003, πE = πg.ResolveGlobal(πF, ßOrderedDict); πE != nil {
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
					if πE = πClass.SetItem(πF, ß_asdict.ToObject(), πTemp005); πE != nil {
						continue
					}
					// line 218: 'Return a new OrderedDict which maps field names to their values'
					πF.SetLineno(218)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp007}, πg.NewStr("Return a new OrderedDict which maps field names to their values").ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.ResolveClass(πF, πClass, nil, ß_asdict); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp008, ß__doc__, πTemp007); πE != nil {
						continue
					}
					// line 221: def _replace(_self, **kwds):
					πF.SetLineno(221)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "_self", Def: nil}
					πTemp007 = πg.NewFunction(πg.NewCode("_replace", "/usr/lib/python2.7/urlparse.py", πTemp002, πg.CodeFlagKWArg, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µ_self *πg.Object = πArgs[0]; _ = µ_self
						var µkwds *πg.Object = πArgs[1]; _ = µkwds
						var µresult *πg.Object = πg.UnboundLocal; _ = µresult
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
						var πTemp006 *πg.Object
						_ = πTemp006
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 222: 'Return a new ParseResult object replacing specified fields with new values'
							πF.SetLineno(222)
							// line 223: result = _self._make(map(kwds.pop, ('scheme', 'netloc', 'path', 'params', 'query', 'fragment'), _self))
							πF.SetLineno(223)
							πTemp001 = πF.MakeArgs(1)
							πTemp002 = πF.MakeArgs(3)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µkwds, ßpop, nil); πE != nil {
								continue
							}
							πTemp002[0] = πTemp003
							πTemp003 = πg.NewTuple6(ßscheme.ToObject(), ßnetloc.ToObject(), ßpath.ToObject(), ßparams.ToObject(), ßquery.ToObject(), ßfragment.ToObject()).ToObject()
							πTemp002[1] = πTemp003
							if πE = πg.CheckLocal(πF, µ_self, "_self"); πE != nil {
								continue
							}
							πTemp002[2] = µ_self
							if πTemp003, πE = πg.ResolveGlobal(πF, ßmap); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp002)
							πTemp001[0] = πTemp004
							if πE = πg.CheckLocal(πF, µ_self, "_self"); πE != nil {
								continue
							}
							if πTemp003, πE = πg.GetAttr(πF, µ_self, ß_make, nil); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							µresult = πTemp004
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp005, πE = πg.IsTrue(πF, µkwds); πE != nil {
								continue
							}
							if πTemp005 {
								goto Label1
							}
							goto Label2
							// line 224: if kwds:
							πF.SetLineno(224)
						Label1:
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µkwds, "kwds"); πE != nil {
								continue
							}
							if πTemp004, πE = πg.GetAttr(πF, µkwds, ßkeys, nil); πE != nil {
								continue
							}
							if πTemp006, πE = πTemp004.Call(πF, nil, nil); πE != nil {
								continue
							}
							if πTemp003, πE = πg.Mod(πF, πg.NewStr("Got unexpected field names: %r").ToObject(), πTemp006); πE != nil {
								continue
							}
							πTemp001[0] = πTemp003
							if πTemp003, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
								continue
							}
							if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
								continue
							}
							πF.FreeArgs(πTemp001)
							// line 225: raise ValueError('Got unexpected field names: %r' % kwds.keys())
							πF.SetLineno(225)
							πE = πF.Raise(πTemp004, nil, nil)
							continue
							goto Label2
						Label2:
							// line 226: return result
							πF.SetLineno(226)
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
					if πE = πClass.SetItem(πF, ß_replace.ToObject(), πTemp007); πE != nil {
						continue
					}
					// line 222: 'Return a new ParseResult object replacing specified fields with new values'
					πF.SetLineno(222)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp008}, πg.NewStr("Return a new ParseResult object replacing specified fields with new values").ToObject()); πE != nil {
						continue
					}
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_replace); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp009, ß__doc__, πTemp008); πE != nil {
						continue
					}
					// line 228: def __getnewargs__(self):
					πF.SetLineno(228)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp008 = πg.NewFunction(πg.NewCode("__getnewargs__", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 229: 'Return self as a plain tuple.  Used by copy and pickle.'
							πF.SetLineno(229)
							// line 230: return tuple(self)
							πF.SetLineno(230)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
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
					if πE = πClass.SetItem(πF, ß__getnewargs__.ToObject(), πTemp008); πE != nil {
						continue
					}
					// line 229: 'Return self as a plain tuple.  Used by copy and pickle.'
					πF.SetLineno(229)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Return self as a plain tuple.  Used by copy and pickle.").ToObject()); πE != nil {
						continue
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß__getnewargs__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
						continue
					}
					// line 232: __dict__ = _property(_asdict)
					πF.SetLineno(232)
					πTemp006 = πF.MakeArgs(1)
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_asdict); πE != nil {
						continue
					}
					πTemp006[0] = πTemp009
					if πTemp009, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp009.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ß__dict__.ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 234: def __getstate__(self):
					πF.SetLineno(234)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp009 = πg.NewFunction(πg.NewCode("__getstate__", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µself *πg.Object = πArgs[0]; _ = µself
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						for ; πF.State() >= 0; πF.PopCheckpoint() {
							switch πF.State() {
							case 0:
							default: panic("unexpected function state")
							}
							// line 235: 'Exclude the OrderedDict from pickling'
							πF.SetLineno(235)
							// line 236: pass
							πF.SetLineno(236)
						}
						if πE != nil {
							πR = nil
						} else if πR == nil {
							πR = πg.None
						}
						return πR, πE
					}), πF.Globals()).ToObject()
					if πE = πClass.SetItem(πF, ß__getstate__.ToObject(), πTemp009); πE != nil {
						continue
					}
					// line 235: 'Exclude the OrderedDict from pickling'
					πF.SetLineno(235)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Exclude the OrderedDict from pickling").ToObject()); πE != nil {
						continue
					}
					if πTemp011, πE = πg.ResolveClass(πF, πClass, nil, ß__getstate__); πE != nil {
						continue
					}
					if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
						continue
					}
					// line 238: scheme = _property(_itemgetter(0), doc='Alias for field number 0')
					πF.SetLineno(238)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(0).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 0").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßscheme.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 240: netloc = _property(_itemgetter(1), doc='Alias for field number 1')
					πF.SetLineno(240)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(1).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 1").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßnetloc.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 242: path = _property(_itemgetter(2), doc='Alias for field number 2')
					πF.SetLineno(242)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(2).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 2").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßpath.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 244: params = _property(_itemgetter(3), doc='Alias for field number 3')
					πF.SetLineno(244)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(3).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 3").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßparams.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 246: query = _property(_itemgetter(4), doc='Alias for field number 4')
					πF.SetLineno(246)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(4).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 4").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßquery.ToObject(), πTemp011); πE != nil {
						continue
					}
					// line 248: fragment = _property(_itemgetter(5), doc='Alias for field number 5')
					πF.SetLineno(248)
					πTemp006 = πF.MakeArgs(1)
					πTemp012 = πF.MakeArgs(1)
					πTemp012[0] = πg.NewInt(5).ToObject()
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_itemgetter); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp012, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp012)
					πTemp006[0] = πTemp011
					πTemp013 = πg.KWArgs{
						{"doc", πg.NewStr("Alias for field number 5").ToObject()},
					}
					if πTemp010, πE = πg.ResolveClass(πF, πClass, nil, ß_property); πE != nil {
						continue
					}
					if πTemp011, πE = πTemp010.Call(πF, πTemp006, πTemp013); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πE = πClass.SetItem(πF, ßfragment.ToObject(), πTemp011); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("_ParseResult").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ß_ParseResult.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 251: class ParseResult(_ParseResult, ResultMixin):
			πF.SetLineno(251)
			πTemp002 = make([]*πg.Object, 2)
			if πTemp007, πE = πg.ResolveGlobal(πF, ß_ParseResult); πE != nil {
				continue
			}
			πTemp002[0] = πTemp007
			if πTemp007, πE = πg.ResolveGlobal(πF, ßResultMixin); πE != nil {
				continue
			}
			πTemp002[1] = πTemp007
			πTemp004 = πg.NewDict()
			if πTemp003, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
				continue
			}
			if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp003); πE != nil {
				continue
			}
			_, πE = πg.NewCode("ParseResult", "/usr/lib/python2.7/urlparse.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
				πClass := πTemp004
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
					// line 253: __slots__ = ()
					πF.SetLineno(253)
					πTemp001 = πg.NewTuple0().ToObject()
					if πE = πClass.SetItem(πF, ß__slots__.ToObject(), πTemp001); πE != nil {
						continue
					}
					// line 255: def geturl(self):
					πF.SetLineno(255)
					πTemp002 = make([]πg.Param, 1)
					πTemp002[0] = πg.Param{Name: "self", Def: nil}
					πTemp001 = πg.NewFunction(πg.NewCode("geturl", "/usr/lib/python2.7/urlparse.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
							// line 256: return urlunparse(self)
							πF.SetLineno(256)
							πTemp001 = πF.MakeArgs(1)
							if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
								continue
							}
							πTemp001[0] = µself
							if πTemp002, πE = πg.ResolveGlobal(πF, ßurlunparse); πE != nil {
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
					if πE = πClass.SetItem(πF, ßgeturl.ToObject(), πTemp001); πE != nil {
						continue
					}
				}
				return nil, nil
			}).Eval(πF, πF.Globals(), nil, nil)
			if πE != nil {
				continue
			}
			if πTemp006, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
				continue
			}
			if πTemp006 == nil {
				πTemp006 = πg.TypeType.ToObject()
			}
			if πTemp007, πE = πTemp006.Call(πF, []*πg.Object{πg.NewStr("ParseResult").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
				continue
			}
			if πE = πF.Globals().SetItem(πF, ßParseResult.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 259: def urlparse(url, scheme='', allow_fragments=True):
			πF.SetLineno(259)
			πTemp005 = make([]πg.Param, 3)
			πTemp005[0] = πg.Param{Name: "url", Def: nil}
			πTemp005[1] = πg.Param{Name: "scheme", Def: ß.ToObject()}
			if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp005[2] = πg.Param{Name: "allow_fragments", Def: πTemp006}
			πTemp003 = πg.NewFunction(πg.NewCode("urlparse", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µurl *πg.Object = πArgs[0]; _ = µurl
				var µscheme *πg.Object = πArgs[1]; _ = µscheme
				var µallow_fragments *πg.Object = πArgs[2]; _ = µallow_fragments
				var µtuple *πg.Object = πg.UnboundLocal; _ = µtuple
				var µnetloc *πg.Object = πg.UnboundLocal; _ = µnetloc
				var µquery *πg.Object = πg.UnboundLocal; _ = µquery
				var µfragment *πg.Object = πg.UnboundLocal; _ = µfragment
				var µparams *πg.Object = πg.UnboundLocal; _ = µparams
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
				var πTemp008 bool
				_ = πTemp008
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 260: """Parse a URL into 6 components:
					πF.SetLineno(260)
					// line 265: tuple = urlsplit(url, scheme, allow_fragments)
					πF.SetLineno(265)
					πTemp001 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001[0] = µurl
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					πTemp001[1] = µscheme
					if πE = πg.CheckLocal(πF, µallow_fragments, "allow_fragments"); πE != nil {
						continue
					}
					πTemp001[2] = µallow_fragments
					if πTemp002, πE = πg.ResolveGlobal(πF, ßurlsplit); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µtuple = πTemp003
					// line 266: scheme, netloc, url, query, fragment = tuple
					πF.SetLineno(266)
					if πE = πg.CheckLocal(πF, µtuple, "tuple"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, µtuple); πE != nil {
						continue
					}
					µscheme = πTemp002
					µnetloc = πTemp003
					µurl = πTemp004
					µquery = πTemp005
					µfragment = πTemp006
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßuses_params); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πTemp004, µscheme); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008).ToObject()
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µurl, πg.NewStr(";").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008).ToObject()
					πTemp002 = πTemp003
				Label1:
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label2
					}
					goto Label3
					// line 267: if scheme in uses_params and ';' in url:
					πF.SetLineno(267)
				Label2:
					// line 268: url, params = _splitparams(url)
					πF.SetLineno(268)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001[0] = µurl
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_splitparams); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µurl = πTemp002
					µparams = πTemp004
					goto Label4
				Label3:
					// line 270: params = ''
					πF.SetLineno(270)
					µparams = ß.ToObject()
					goto Label4
				Label4:
					// line 271: return ParseResult(scheme, netloc, url, params, query, fragment)
					πF.SetLineno(271)
					πTemp001 = πF.MakeArgs(6)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					πTemp001[0] = µscheme
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					πTemp001[1] = µnetloc
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001[2] = µurl
					if πE = πg.CheckLocal(πF, µparams, "params"); πE != nil {
						continue
					}
					πTemp001[3] = µparams
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					πTemp001[4] = µquery
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					πTemp001[5] = µfragment
					if πTemp002, πE = πg.ResolveGlobal(πF, ßParseResult); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßurlparse.ToObject(), πTemp003); πE != nil {
				continue
			}
			// line 260: """Parse a URL into 6 components:
			πF.SetLineno(260)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πg.NewStr("Parse a URL into 6 components:\n    <scheme>://<netloc>/<path>;<params>?<query>#<fragment>\n    Return a 6-tuple: (scheme, netloc, path, params, query, fragment).\n    Note that we don't break the components up in smaller bits\n    (e.g. netloc is a single string) and we don't expand % escapes.").ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.ResolveGlobal(πF, ßurlparse); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp007, ß__doc__, πTemp006); πE != nil {
				continue
			}
			// line 273: def _splitparams(url):
			πF.SetLineno(273)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "url", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("_splitparams", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µurl *πg.Object = πArgs[0]; _ = µurl
				var µi *πg.Object = πg.UnboundLocal; _ = µi
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
				var πTemp007 *πg.Object
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µurl, πg.NewStr("/").ToObject()); πE != nil {
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
					// line 274: if '/'  in url:
					πF.SetLineno(274)
				Label1:
					// line 275: i = url.find(';', url.rfind('/'))
					πF.SetLineno(275)
					πTemp003 = πF.MakeArgs(2)
					πTemp003[0] = πg.NewStr(";").ToObject()
					πTemp004 = πF.MakeArgs(1)
					πTemp004[0] = πg.NewStr("/").ToObject()
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µurl, ßrfind, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp004)
					πTemp003[1] = πTemp005
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µurl, ßfind, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µi = πTemp005
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label4
					}
					goto Label5
					// line 276: if i < 0:
					πF.SetLineno(276)
				Label4:
					// line 277: return url, ''
					πF.SetLineno(277)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µurl, ß.ToObject()).ToObject()
					πR = πTemp001
					continue
					goto Label5
				Label5:
					goto Label3
				Label2:
					// line 279: i = url.find(';')
					πF.SetLineno(279)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr(";").ToObject()
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µurl, ßfind, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µi = πTemp005
					goto Label3
				Label3:
					// line 280: return url[:i], url[i+1:]
					πF.SetLineno(280)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µurl, πTemp005); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πTemp007, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µurl, πTemp005); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_splitparams.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 282: def _splitnetloc(url, start=0):
			πF.SetLineno(282)
			πTemp005 = make([]πg.Param, 2)
			πTemp005[0] = πg.Param{Name: "url", Def: nil}
			πTemp005[1] = πg.Param{Name: "start", Def: πg.NewInt(0).ToObject()}
			πTemp007 = πg.NewFunction(πg.NewCode("_splitnetloc", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µurl *πg.Object = πArgs[0]; _ = µurl
				var µstart *πg.Object = πArgs[1]; _ = µstart
				var µdelim *πg.Object = πg.UnboundLocal; _ = µdelim
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µwdelim *πg.Object = πg.UnboundLocal; _ = µwdelim
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 283: delim = len(url)   # position of end of domain part of url, default is end
					πF.SetLineno(283)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001[0] = µurl
					if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdelim = πTemp003
					if πTemp002, πE = πg.Iter(πF, πg.NewStr("/?#").ToObject()); πE != nil {
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
					// line 285: wdelim = url.find(c, start)        # find first of this delim
					πF.SetLineno(285)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					πTemp001[0] = µc
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					πTemp001[1] = µstart
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µurl, ßfind, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µwdelim = πTemp006
					if πE = πg.CheckLocal(πF, µwdelim, "wdelim"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GE(πF, µwdelim, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp005 {
						goto Label4
					}
					goto Label5
					// line 286: if wdelim >= 0:                    # if found
					πF.SetLineno(286)
				Label4:
					// line 287: delim = min(delim, wdelim)     # use earliest delim position
					πF.SetLineno(287)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
						continue
					}
					πTemp001[0] = µdelim
					if πE = πg.CheckLocal(πF, µwdelim, "wdelim"); πE != nil {
						continue
					}
					πTemp001[1] = µwdelim
					if πTemp003, πE = πg.ResolveGlobal(πF, ßmin); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µdelim = πTemp006
					goto Label5
				Label5:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 288: return url[start:delim], url[delim:]   # return (domain, rest)
					πF.SetLineno(288)
					if πE = πg.CheckLocal(πF, µstart, "start"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µstart, µdelim, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µurl, πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdelim, "delim"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{µdelim, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µurl, πTemp003); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(πTemp006, πTemp007).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ß_splitnetloc.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 290: def urlsplit(url, scheme='', allow_fragments=True):
			πF.SetLineno(290)
			πTemp005 = make([]πg.Param, 3)
			πTemp005[0] = πg.Param{Name: "url", Def: nil}
			πTemp005[1] = πg.Param{Name: "scheme", Def: ß.ToObject()}
			if πTemp009, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp005[2] = πg.Param{Name: "allow_fragments", Def: πTemp009}
			πTemp008 = πg.NewFunction(πg.NewCode("urlsplit", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µurl *πg.Object = πArgs[0]; _ = µurl
				var µscheme *πg.Object = πArgs[1]; _ = µscheme
				var µallow_fragments *πg.Object = πArgs[2]; _ = µallow_fragments
				var µkey *πg.Object = πg.UnboundLocal; _ = µkey
				var µcached *πg.Object = πg.UnboundLocal; _ = µcached
				var µnetloc *πg.Object = πg.UnboundLocal; _ = µnetloc
				var µquery *πg.Object = πg.UnboundLocal; _ = µquery
				var µfragment *πg.Object = πg.UnboundLocal; _ = µfragment
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µv *πg.Object = πg.UnboundLocal; _ = µv
				var µc *πg.Object = πg.UnboundLocal; _ = µc
				var µrest *πg.Object = πg.UnboundLocal; _ = µrest
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
				var πTemp008 bool
				_ = πTemp008
				var πTemp009 []πg.Param
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 21: goto Label21
					case 22: goto Label22
					default: panic("unexpected function state")
					}
					// line 291: """Parse a URL into 5 components:
					πF.SetLineno(291)
					// line 296: allow_fragments = bool(allow_fragments)
					πF.SetLineno(296)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µallow_fragments, "allow_fragments"); πE != nil {
						continue
					}
					πTemp001[0] = µallow_fragments
					if πTemp002, πE = πg.ResolveGlobal(πF, ßbool); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µallow_fragments = πTemp003
					// line 297: key = url, scheme, allow_fragments, type(url), type(scheme)
					πF.SetLineno(297)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µallow_fragments, "allow_fragments"); πE != nil {
						continue
					}
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001[0] = µurl
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					πTemp001[0] = µscheme
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtype); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp002 = πg.NewTuple5(µurl, µscheme, µallow_fragments, πTemp004, πTemp005).ToObject()
					µkey = πTemp002
					// line 298: cached = _parse_cache.get(key, None)
					πF.SetLineno(298)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp001[0] = µkey
					if πTemp002, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
						continue
					}
					πTemp001[1] = πTemp002
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_parse_cache); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µcached = πTemp002
					if πE = πg.CheckLocal(πF, µcached, "cached"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µcached); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					goto Label2
					// line 299: if cached:
					πF.SetLineno(299)
				Label1:
					// line 300: return cached
					πF.SetLineno(300)
					if πE = πg.CheckLocal(πF, µcached, "cached"); πE != nil {
						continue
					}
					πR = µcached
					continue
					goto Label2
				Label2:
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_parse_cache); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßMAX_CACHE_SIZE); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GE(πF, πTemp004, πTemp003); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 301: if len(_parse_cache) >= MAX_CACHE_SIZE: # avoid runaway growth
					πF.SetLineno(301)
				Label3:
					// line 302: clear_cache()
					πF.SetLineno(302)
					if πTemp002, πE = πg.ResolveGlobal(πF, ßclear_cache); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					goto Label4
				Label4:
					// line 303: netloc = query = fragment = ''
					πF.SetLineno(303)
					µnetloc = ß.ToObject()
					µquery = ß.ToObject()
					µfragment = ß.ToObject()
					// line 304: i = url.find(':')
					πF.SetLineno(304)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr(":").ToObject()
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µurl, ßfind, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µi = πTemp003
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GT(πF, µi, πg.NewInt(0).ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label5
					}
					goto Label6
					// line 305: if i > 0:
					πF.SetLineno(305)
				Label5:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µurl, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp004, ßhttp.ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label7
					}
					goto Label8
					// line 306: if url[:i] == 'http': # optimize the common case
					πF.SetLineno(306)
				Label7:
					// line 307: scheme = url[:i].lower()
					πF.SetLineno(307)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µurl, πTemp002); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßlower, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					µscheme = πTemp003
					// line 308: url = url[i+1:]
					πF.SetLineno(308)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{πTemp003, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µurl, πTemp002); πE != nil {
						continue
					}
					µurl = πTemp003
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µurl, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewStr("//").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 309: if url[:2] == '//':
					πF.SetLineno(309)
				Label9:
					// line 310: netloc, url = _splitnetloc(url, 2)
					πF.SetLineno(310)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001[0] = µurl
					πTemp001[1] = πg.NewInt(2).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_splitnetloc); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µnetloc = πTemp002
					µurl = πTemp004
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µnetloc, πg.NewStr("[").ToObject()); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp008).ToObject()
					πTemp003 = πTemp004
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label12
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µnetloc, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp008).ToObject()
					πTemp003 = πTemp004
				Label12:
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µnetloc, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp008).ToObject()
					πTemp003 = πTemp004
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label13
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µnetloc, πg.NewStr("[").ToObject()); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp008).ToObject()
					πTemp003 = πTemp004
				Label13:
					πTemp002 = πTemp003
				Label11:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label14
					}
					goto Label15
					// line 311: if (('[' in netloc and ']' not in netloc) or
					πF.SetLineno(311)
				Label14:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("Invalid IPv6 URL").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 313: raise ValueError("Invalid IPv6 URL")
					πF.SetLineno(313)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label15
				Label15:
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µallow_fragments, "allow_fragments"); πE != nil {
						continue
					}
					πTemp002 = µallow_fragments
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label16
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µurl, πg.NewStr("#").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp007).ToObject()
					πTemp002 = πTemp003
				Label16:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label17
					}
					goto Label18
					// line 314: if allow_fragments and '#' in url:
					πF.SetLineno(314)
				Label17:
					// line 315: url, fragment = url.split('#', 1)
					πF.SetLineno(315)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("#").ToObject()
					πTemp001[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µurl, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µurl = πTemp002
					µfragment = πTemp004
					goto Label18
				Label18:
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µurl, πg.NewStr("?").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label19
					}
					goto Label20
					// line 316: if '?' in url:
					πF.SetLineno(316)
				Label19:
					// line 317: url, query = url.split('?', 1)
					πF.SetLineno(317)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("?").ToObject()
					πTemp001[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µurl, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp004}}}, πTemp003); πE != nil {
						continue
					}
					µurl = πTemp002
					µquery = πTemp004
					goto Label20
				Label20:
					// line 318: v = SplitResult(scheme, netloc, url, query, fragment)
					πF.SetLineno(318)
					πTemp001 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					πTemp001[0] = µscheme
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					πTemp001[1] = µnetloc
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001[2] = µurl
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					πTemp001[3] = µquery
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					πTemp001[4] = µfragment
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSplitResult); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µv = πTemp003
					// line 319: _parse_cache[key] = v
					πF.SetLineno(319)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µv); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_parse_cache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp004 = µkey
					if πE = πg.SetItem(πF, πTemp003, πTemp004, πTemp002); πE != nil {
						continue
					}
					// line 320: return v
					πF.SetLineno(320)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					πR = µv
					continue
					goto Label8
				Label8:
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µurl, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
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
						µc = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(21)            
					if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßscheme_chars); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, πTemp004, µc); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(!πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label24
					}
					goto Label25
					// line 322: if c not in scheme_chars:
					πF.SetLineno(322)
				Label24:
					// line 323: break
					πF.SetLineno(323)
					πTemp006 = true
					continue
					goto Label25
				Label25:
					continue
				Label22:
					if πE != nil || πR != nil {
						continue
					}
					// line 327: rest = url[i+1:]
					πF.SetLineno(327)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µurl, πTemp003); πE != nil {
						continue
					}
					µrest = πTemp004
					if πE = πg.CheckLocal(πF, µrest, "rest"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, µrest); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp008).ToObject()
					πTemp003 = πTemp004
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label26
					}
					πTemp001 = πF.MakeArgs(1)
					πTemp009 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/urlparse.py", πTemp009, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µc *πg.Object = πg.UnboundLocal; _ = µc
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
								if πE = πg.CheckLocal(πF, µrest, "rest"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µrest); πE != nil {
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
									µc = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 328: if not rest or any(c not in '0123456789' for c in rest):
								πF.SetLineno(328)
								if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.Contains(πF, ß0123456789.ToObject(), µc); πE != nil {
									continue
								}
								πTemp004 = πg.GetBool(!πTemp003).ToObject()
								πF.PushCheckpoint(4)
								return πTemp004, nil
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
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					πTemp001[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßany); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πTemp003 = πTemp010
				Label26:
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label27
					}
					goto Label28
					// line 328: if not rest or any(c not in '0123456789' for c in rest):
					πF.SetLineno(328)
				Label27:
					// line 330: scheme, url = url[:i].lower(), rest
					πF.SetLineno(330)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πg.None, µi, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µurl, πTemp005); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetAttr(πF, πTemp010, ßlower, nil); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp005.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µrest, "rest"); πE != nil {
						continue
					}
					πTemp003 = πg.NewTuple2(πTemp010, µrest).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp010}}}, πTemp003); πE != nil {
						continue
					}
					µscheme = πTemp005
					µurl = πTemp010
					goto Label28
				Label28:
				Label23:
					goto Label6
				Label6:
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µurl, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Eq(πF, πTemp005, πg.NewStr("//").ToObject()); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label29
					}
					goto Label30
					// line 332: if url[:2] == '//':
					πF.SetLineno(332)
				Label29:
					// line 333: netloc, url = _splitnetloc(url, 2)
					πF.SetLineno(333)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001[0] = µurl
					πTemp001[1] = πg.NewInt(2).ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_splitnetloc); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µnetloc = πTemp002
					µurl = πTemp005
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µnetloc, πg.NewStr("[").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp008).ToObject()
					πTemp003 = πTemp005
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label32
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µnetloc, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp008).ToObject()
					πTemp003 = πTemp005
				Label32:
					πTemp002 = πTemp003
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label31
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µnetloc, πg.NewStr("]").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(πTemp008).ToObject()
					πTemp003 = πTemp005
					if πTemp007, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label33
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, µnetloc, πg.NewStr("[").ToObject()); πE != nil {
						continue
					}
					πTemp005 = πg.GetBool(!πTemp008).ToObject()
					πTemp003 = πTemp005
				Label33:
					πTemp002 = πTemp003
				Label31:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label34
					}
					goto Label35
					// line 334: if (('[' in netloc and ']' not in netloc) or
					πF.SetLineno(334)
				Label34:
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("Invalid IPv6 URL").ToObject()
					if πTemp002, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 336: raise ValueError("Invalid IPv6 URL")
					πF.SetLineno(336)
					πE = πF.Raise(πTemp003, nil, nil)
					continue
					goto Label35
				Label35:
					goto Label30
				Label30:
					if πE = πg.CheckLocal(πF, µallow_fragments, "allow_fragments"); πE != nil {
						continue
					}
					πTemp002 = µallow_fragments
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label36
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µurl, πg.NewStr("#").ToObject()); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp007).ToObject()
					πTemp002 = πTemp003
				Label36:
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label37
					}
					goto Label38
					// line 337: if allow_fragments and '#' in url:
					πF.SetLineno(337)
				Label37:
					// line 338: url, fragment = url.split('#', 1)
					πF.SetLineno(338)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("#").ToObject()
					πTemp001[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µurl, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µurl = πTemp002
					µfragment = πTemp005
					goto Label38
				Label38:
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Contains(πF, µurl, πg.NewStr("?").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πg.GetBool(πTemp006).ToObject()
					if πTemp006, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label39
					}
					goto Label40
					// line 339: if '?' in url:
					πF.SetLineno(339)
				Label39:
					// line 340: url, query = url.split('?', 1)
					πF.SetLineno(340)
					πTemp001 = πF.MakeArgs(2)
					πTemp001[0] = πg.NewStr("?").ToObject()
					πTemp001[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µurl, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp005}}}, πTemp003); πE != nil {
						continue
					}
					µurl = πTemp002
					µquery = πTemp005
					goto Label40
				Label40:
					// line 341: v = SplitResult(scheme, netloc, url, query, fragment)
					πF.SetLineno(341)
					πTemp001 = πF.MakeArgs(5)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					πTemp001[0] = µscheme
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					πTemp001[1] = µnetloc
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001[2] = µurl
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					πTemp001[3] = µquery
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					πTemp001[4] = µfragment
					if πTemp002, πE = πg.ResolveGlobal(πF, ßSplitResult); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µv = πTemp003
					// line 342: _parse_cache[key] = v
					πF.SetLineno(342)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, µv); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ß_parse_cache); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µkey, "key"); πE != nil {
						continue
					}
					πTemp005 = µkey
					if πE = πg.SetItem(πF, πTemp003, πTemp005, πTemp002); πE != nil {
						continue
					}
					// line 343: return v
					πF.SetLineno(343)
					if πE = πg.CheckLocal(πF, µv, "v"); πE != nil {
						continue
					}
					πR = µv
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßurlsplit.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 291: """Parse a URL into 5 components:
			πF.SetLineno(291)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp009}, πg.NewStr("Parse a URL into 5 components:\n    <scheme>://<netloc>/<path>?<query>#<fragment>\n    Return a 5-tuple: (scheme, netloc, path, query, fragment).\n    Note that we don't break the components up in smaller bits\n    (e.g. netloc is a single string) and we don't expand % escapes.").ToObject()); πE != nil {
				continue
			}
			if πTemp010, πE = πg.ResolveGlobal(πF, ßurlsplit); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp010, ß__doc__, πTemp009); πE != nil {
				continue
			}
			// line 345: def urlunparse(data):
			πF.SetLineno(345)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "data", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("urlunparse", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdata *πg.Object = πArgs[0]; _ = µdata
				var µscheme *πg.Object = πg.UnboundLocal; _ = µscheme
				var µnetloc *πg.Object = πg.UnboundLocal; _ = µnetloc
				var µurl *πg.Object = πg.UnboundLocal; _ = µurl
				var µparams *πg.Object = πg.UnboundLocal; _ = µparams
				var µquery *πg.Object = πg.UnboundLocal; _ = µquery
				var µfragment *πg.Object = πg.UnboundLocal; _ = µfragment
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
					// line 346: """Put a parsed URL back together again.  This may result in a
					πF.SetLineno(346)
					// line 350: scheme, netloc, url, params, query, fragment = data
					πF.SetLineno(350)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}}}, µdata); πE != nil {
						continue
					}
					µscheme = πTemp001
					µnetloc = πTemp002
					µurl = πTemp003
					µparams = πTemp004
					µquery = πTemp005
					µfragment = πTemp006
					if πE = πg.CheckLocal(πF, µparams, "params"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µparams); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label1
					}
					goto Label2
					// line 351: if params:
					πF.SetLineno(351)
				Label1:
					// line 352: url = "%s;%s" % (url, params)
					πF.SetLineno(352)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparams, "params"); πE != nil {
						continue
					}
					πTemp002 = πg.NewTuple2(µurl, µparams).ToObject()
					if πTemp001, πE = πg.Mod(πF, πg.NewStr("%s;%s").ToObject(), πTemp002); πE != nil {
						continue
					}
					µurl = πTemp001
					goto Label2
				Label2:
					// line 353: return urlunsplit((scheme, netloc, url, query, fragment))
					πF.SetLineno(353)
					πTemp008 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple5(µscheme, µnetloc, µurl, µquery, µfragment).ToObject()
					πTemp008[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßurlunsplit); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp001.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
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
			if πE = πF.Globals().SetItem(πF, ßurlunparse.ToObject(), πTemp009); πE != nil {
				continue
			}
			// line 346: """Put a parsed URL back together again.  This may result in a
			πF.SetLineno(346)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp010}, πg.NewStr("Put a parsed URL back together again.  This may result in a\n    slightly different, but equivalent URL, if the URL that was parsed\n    originally had redundant delimiters, e.g. a ? with an empty query\n    (the draft states that these are equivalent).").ToObject()); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ßurlunparse); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp011, ß__doc__, πTemp010); πE != nil {
				continue
			}
			// line 355: def urlunsplit(data):
			πF.SetLineno(355)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "data", Def: nil}
			πTemp010 = πg.NewFunction(πg.NewCode("urlunsplit", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µdata *πg.Object = πArgs[0]; _ = µdata
				var µscheme *πg.Object = πg.UnboundLocal; _ = µscheme
				var µnetloc *πg.Object = πg.UnboundLocal; _ = µnetloc
				var µurl *πg.Object = πg.UnboundLocal; _ = µurl
				var µquery *πg.Object = πg.UnboundLocal; _ = µquery
				var µfragment *πg.Object = πg.UnboundLocal; _ = µfragment
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
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 356: """Combine the elements of a tuple as returned by urlsplit() into a
					πF.SetLineno(356)
					// line 361: scheme, netloc, url, query, fragment = data
					πF.SetLineno(361)
					if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
						continue
					}
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp002}, πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp005}}}, µdata); πE != nil {
						continue
					}
					µscheme = πTemp001
					µnetloc = πTemp002
					µurl = πTemp003
					µquery = πTemp004
					µfragment = πTemp005
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					πTemp001 = µnetloc
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label1
					}
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					πTemp002 = µscheme
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label2
					}
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßuses_netloc); πE != nil {
						continue
					}
					if πTemp008, πE = πg.Contains(πF, πTemp004, µscheme); πE != nil {
						continue
					}
					πTemp003 = πg.GetBool(πTemp008).ToObject()
					πTemp002 = πTemp003
					if πTemp007, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label2
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µurl, πTemp004); πE != nil {
						continue
					}
					if πTemp003, πE = πg.NE(πF, πTemp005, πg.NewStr("//").ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp003
				Label2:
					πTemp001 = πTemp002
				Label1:
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label3
					}
					goto Label4
					// line 362: if netloc or (scheme and scheme in uses_netloc and url[:2] != '//'):
					πF.SetLineno(362)
				Label3:
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001 = µurl
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp006 {
						goto Label5
					}
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetItem(πF, µurl, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.NE(πF, πTemp004, πg.NewStr("/").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp002
				Label5:
					if πTemp006, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label6
					}
					goto Label7
					// line 363: if url and url[:1] != '/': url = '/' + url
					πF.SetLineno(363)
				Label6:
					// line 363: if url and url[:1] != '/': url = '/' + url
					πF.SetLineno(363)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πg.NewStr("/").ToObject(), µurl); πE != nil {
						continue
					}
					µurl = πTemp001
					goto Label7
				Label7:
					// line 364: url = '//' + (netloc or '') + url
					πF.SetLineno(364)
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					πTemp003 = µnetloc
					if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label8
					}
					πTemp003 = ß.ToObject()
				Label8:
					if πTemp002, πE = πg.Add(πF, πg.NewStr("//").ToObject(), πTemp003); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, µurl); πE != nil {
						continue
					}
					µurl = πTemp001
					goto Label4
				Label4:
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µscheme); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label9
					}
					goto Label10
					// line 365: if scheme:
					πF.SetLineno(365)
				Label9:
					// line 366: url = scheme + ':' + url
					πF.SetLineno(366)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µscheme, πg.NewStr(":").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, µurl); πE != nil {
						continue
					}
					µurl = πTemp001
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µquery); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label11
					}
					goto Label12
					// line 367: if query:
					πF.SetLineno(367)
				Label11:
					// line 368: url = url + '?' + query
					πF.SetLineno(368)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µurl, πg.NewStr("?").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, µquery); πE != nil {
						continue
					}
					µurl = πTemp001
					goto Label12
				Label12:
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.IsTrue(πF, µfragment); πE != nil {
						continue
					}
					if πTemp006 {
						goto Label13
					}
					goto Label14
					// line 369: if fragment:
					πF.SetLineno(369)
				Label13:
					// line 370: url = url + '#' + fragment
					πF.SetLineno(370)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Add(πF, µurl, πg.NewStr("#").ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, πTemp002, µfragment); πE != nil {
						continue
					}
					µurl = πTemp001
					goto Label14
				Label14:
					// line 371: return url
					πF.SetLineno(371)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πR = µurl
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßurlunsplit.ToObject(), πTemp010); πE != nil {
				continue
			}
			// line 356: """Combine the elements of a tuple as returned by urlsplit() into a
			πF.SetLineno(356)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp011}, πg.NewStr("Combine the elements of a tuple as returned by urlsplit() into a\n    complete URL as a string. The data argument can be any five-item iterable.\n    This may result in a slightly different, but equivalent URL, if the URL that\n    was parsed originally had unnecessary delimiters (for example, a ? with an\n    empty query; the RFC states that these are equivalent).").ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßurlunsplit); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp012, ß__doc__, πTemp011); πE != nil {
				continue
			}
			// line 373: def urljoin(base, url, allow_fragments=True):
			πF.SetLineno(373)
			πTemp005 = make([]πg.Param, 3)
			πTemp005[0] = πg.Param{Name: "base", Def: nil}
			πTemp005[1] = πg.Param{Name: "url", Def: nil}
			if πTemp012, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
				continue
			}
			πTemp005[2] = πg.Param{Name: "allow_fragments", Def: πTemp012}
			πTemp011 = πg.NewFunction(πg.NewCode("urljoin", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µbase *πg.Object = πArgs[0]; _ = µbase
				var µurl *πg.Object = πArgs[1]; _ = µurl
				var µallow_fragments *πg.Object = πArgs[2]; _ = µallow_fragments
				var µbscheme *πg.Object = πg.UnboundLocal; _ = µbscheme
				var µbnetloc *πg.Object = πg.UnboundLocal; _ = µbnetloc
				var µbpath *πg.Object = πg.UnboundLocal; _ = µbpath
				var µbparams *πg.Object = πg.UnboundLocal; _ = µbparams
				var µbquery *πg.Object = πg.UnboundLocal; _ = µbquery
				var µbfragment *πg.Object = πg.UnboundLocal; _ = µbfragment
				var µscheme *πg.Object = πg.UnboundLocal; _ = µscheme
				var µnetloc *πg.Object = πg.UnboundLocal; _ = µnetloc
				var µpath *πg.Object = πg.UnboundLocal; _ = µpath
				var µparams *πg.Object = πg.UnboundLocal; _ = µparams
				var µquery *πg.Object = πg.UnboundLocal; _ = µquery
				var µfragment *πg.Object = πg.UnboundLocal; _ = µfragment
				var µsegments *πg.Object = πg.UnboundLocal; _ = µsegments
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µn *πg.Object = πg.UnboundLocal; _ = µn
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
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 bool
				_ = πTemp010
				var πTemp011 bool
				_ = πTemp011
				var πTemp012 bool
				_ = πTemp012
				var πTemp013 []*πg.Object
				_ = πTemp013
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 21: goto Label21
					case 22: goto Label22
					case 24: goto Label24
					case 25: goto Label25
					case 27: goto Label27
					case 28: goto Label28
					default: panic("unexpected function state")
					}
					// line 374: """Join a base URL and a possibly relative URL to form an absolute
					πF.SetLineno(374)
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µbase); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 376: if not base:
					πF.SetLineno(376)
				Label1:
					// line 377: return url
					πF.SetLineno(377)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πR = µurl
					continue
					goto Label2
				Label2:
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µurl); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label3
					}
					goto Label4
					// line 378: if not url:
					πF.SetLineno(378)
				Label3:
					// line 379: return base
					πF.SetLineno(379)
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					πR = µbase
					continue
					goto Label4
				Label4:
					// line 380: bscheme, bnetloc, bpath, bparams, bquery, bfragment = \
					πF.SetLineno(380)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µbase, "base"); πE != nil {
						continue
					}
					πTemp003[0] = µbase
					πTemp003[1] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µallow_fragments, "allow_fragments"); πE != nil {
						continue
					}
					πTemp003[2] = µallow_fragments
					if πTemp001, πE = πg.ResolveGlobal(πF, ßurlparse); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
						continue
					}
					µbscheme = πTemp001
					µbnetloc = πTemp005
					µbpath = πTemp006
					µbparams = πTemp007
					µbquery = πTemp008
					µbfragment = πTemp009
					// line 382: scheme, netloc, path, params, query, fragment = \
					πF.SetLineno(382)
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp003[0] = µurl
					if πE = πg.CheckLocal(πF, µbscheme, "bscheme"); πE != nil {
						continue
					}
					πTemp003[1] = µbscheme
					if πE = πg.CheckLocal(πF, µallow_fragments, "allow_fragments"); πE != nil {
						continue
					}
					πTemp003[2] = µallow_fragments
					if πTemp001, πE = πg.ResolveGlobal(πF, ßurlparse); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
						continue
					}
					µscheme = πTemp001
					µnetloc = πTemp005
					µpath = πTemp006
					µparams = πTemp007
					µquery = πTemp008
					µfragment = πTemp009
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbscheme, "bscheme"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.NE(πF, µscheme, µbscheme); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label5
					}
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ResolveGlobal(πF, ßuses_relative); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Contains(πF, πTemp005, µscheme); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp010).ToObject()
					πTemp001 = πTemp004
				Label5:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label6
					}
					goto Label7
					// line 384: if scheme != bscheme or scheme not in uses_relative:
					πF.SetLineno(384)
				Label6:
					// line 385: return url
					πF.SetLineno(385)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πR = µurl
					continue
					goto Label7
				Label7:
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.ResolveGlobal(πF, ßuses_netloc); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, πTemp004, µscheme); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label8
					}
					goto Label9
					// line 386: if scheme in uses_netloc:
					πF.SetLineno(386)
				Label8:
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µnetloc); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label10
					}
					goto Label11
					// line 387: if netloc:
					πF.SetLineno(387)
				Label10:
					// line 388: return urlunparse((scheme, netloc, path,
					πF.SetLineno(388)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparams, "params"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple6(µscheme, µnetloc, µpath, µparams, µquery, µfragment).ToObject()
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßurlunparse); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label11
				Label11:
					// line 390: netloc = bnetloc
					πF.SetLineno(390)
					if πE = πg.CheckLocal(πF, µbnetloc, "bnetloc"); πE != nil {
						continue
					}
					µnetloc = µbnetloc
					goto Label9
				Label9:
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(1).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µpath, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewStr("/").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label12
					}
					goto Label13
					// line 391: if path[:1] == '/':
					πF.SetLineno(391)
				Label12:
					// line 392: return urlunparse((scheme, netloc, path,
					πF.SetLineno(392)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparams, "params"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple6(µscheme, µnetloc, µpath, µparams, µquery, µfragment).ToObject()
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßurlunparse); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label13
				Label13:
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, µpath); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp010).ToObject()
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µparams, "params"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.IsTrue(πF, µparams); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp010).ToObject()
					πTemp001 = πTemp004
				Label14:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label15
					}
					goto Label16
					// line 394: if not path and not params:
					πF.SetLineno(394)
				Label15:
					// line 395: path = bpath
					πF.SetLineno(395)
					if πE = πg.CheckLocal(πF, µbpath, "bpath"); πE != nil {
						continue
					}
					µpath = µbpath
					// line 396: params = bparams
					πF.SetLineno(396)
					if πE = πg.CheckLocal(πF, µbparams, "bparams"); πE != nil {
						continue
					}
					µparams = µbparams
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, µquery); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(!πTemp002).ToObject()
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label17
					}
					goto Label18
					// line 397: if not query:
					πF.SetLineno(397)
				Label17:
					// line 398: query = bquery
					πF.SetLineno(398)
					if πE = πg.CheckLocal(πF, µbquery, "bquery"); πE != nil {
						continue
					}
					µquery = µbquery
					goto Label18
				Label18:
					// line 399: return urlunparse((scheme, netloc, path,
					πF.SetLineno(399)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µparams, "params"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple6(µscheme, µnetloc, µpath, µparams, µquery, µfragment).ToObject()
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßurlunparse); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πR = πTemp004
					continue
					goto Label16
				Label16:
					// line 401: segments = bpath.split('/')[:-1] + path.split('/')
					πF.SetLineno(401)
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("/").ToObject()
					if πE = πg.CheckLocal(πF, µbpath, "bpath"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetAttr(πF, µbpath, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp006.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
						continue
					}
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr("/").ToObject()
					if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µpath, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					µsegments = πTemp001
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µsegments, πTemp004); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Eq(πF, πTemp005, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label19
					}
					goto Label20
					// line 403: if segments[-1] == '.':
					πF.SetLineno(403)
				Label19:
					// line 404: segments[-1] = ''
					πF.SetLineno(404)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ß.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πE = πg.SetItem(πF, µsegments, πTemp004, πTemp001); πE != nil {
						continue
					}
					goto Label20
				Label20:
					// line 405: while '.' in segments:
					πF.SetLineno(405)
					πF.PushCheckpoint(22)
					πTemp002 = false
				Label21:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label23
					}
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Contains(πF, µsegments, πg.NewStr(".").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πg.GetBool(πTemp011).ToObject()
					if πTemp010, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(21)            
					// line 406: segments.remove('.')
					πF.SetLineno(406)
					πTemp003 = πF.MakeArgs(1)
					πTemp003[0] = πg.NewStr(".").ToObject()
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.GetAttr(πF, µsegments, ßremove, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					continue
				Label22:
					if πE != nil || πR != nil {
						continue
					}
				Label23:
					// line 407: while 1:
					πF.SetLineno(407)
					πF.PushCheckpoint(25)
					πTemp002 = false
				Label24:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp002 {
						πF.PopCheckpoint()
						goto Label26
					}
					if πTemp010, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE != nil || !πTemp010 {
						continue
					}
					πF.PushCheckpoint(24)            
					// line 408: i = 1
					πF.SetLineno(408)
					µi = πg.NewInt(1).ToObject()
					// line 409: n = len(segments) - 1
					πF.SetLineno(409)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					πTemp003[0] = µsegments
					if πTemp004, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Sub(πF, πTemp005, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µn = πTemp001
					// line 410: while i < n:
					πF.SetLineno(410)
					πF.PushCheckpoint(28)
					πTemp010 = false
				Label27:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp010 {
						πF.PopCheckpoint()
						goto Label29
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LT(πF, µi, µn); πE != nil {
						continue
					}
					if πTemp011, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πE != nil || !πTemp011 {
						continue
					}
					πF.PushCheckpoint(27)            
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp005 = µi
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µsegments, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp006, πg.NewStr("..").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πTemp011, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp011 {
						goto Label30
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Sub(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µsegments, πTemp005); πE != nil {
						continue
					}
					πTemp005 = πg.NewTuple2(ß.ToObject(), πg.NewStr("..").ToObject()).ToObject()
					if πTemp012, πE = πg.Contains(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(!πTemp012).ToObject()
					πTemp001 = πTemp004
				Label30:
					if πTemp011, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp011 {
						goto Label31
					}
					goto Label32
					// line 411: if (segments[i] == '..'
					πF.SetLineno(411)
				Label31:
					// line 413: del segments[i-1:i+1]
					πF.SetLineno(413)
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πTemp004, πTemp005, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.DelItem(πF, µsegments, πTemp001); πE != nil {
						continue
					}
					// line 414: break
					πF.SetLineno(414)
					πTemp010 = true
					continue
					goto Label32
				Label32:
					// line 415: i = i+1
					πF.SetLineno(415)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					µi = πTemp001
					continue
				Label28:
					if πE != nil || πR != nil {
						continue
					}
					// line 417: break
					πF.SetLineno(417)
					πTemp002 = true
					continue
				Label29:
					continue
				Label25:
					if πE != nil || πR != nil {
						continue
					}
				Label26:
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					πTemp003 = make([]*πg.Object, 2)
					πTemp003[0] = ß.ToObject()
					πTemp003[1] = πg.NewStr("..").ToObject()
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πTemp001, πE = πg.Eq(πF, µsegments, πTemp004); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label33
					}
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					πTemp003[0] = µsegments
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp004, πE = πg.GE(πF, πTemp006, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp004
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if !πTemp002 {
						goto Label34
					}
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp005 = πTemp006
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µsegments, πTemp005); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Eq(πF, πTemp006, πg.NewStr("..").ToObject()); πE != nil {
						continue
					}
					πTemp001 = πTemp004
				Label34:
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label35
					}
					goto Label36
					// line 418: if segments == ['', '..']:
					πF.SetLineno(418)
				Label33:
					// line 419: segments[-1] = ''
					πF.SetLineno(419)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ß.ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.Neg(πF, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp004 = πTemp005
					if πE = πg.SetItem(πF, µsegments, πTemp004, πTemp001); πE != nil {
						continue
					}
					goto Label36
					// line 420: elif len(segments) >= 2 and segments[-1] == '..':
					πF.SetLineno(420)
				Label35:
					// line 421: segments[-2:] = ['']
					πF.SetLineno(421)
					πTemp003 = make([]*πg.Object, 1)
					πTemp003[0] = ß.ToObject()
					πTemp001 = πg.NewList(πTemp003...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πTemp001); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Neg(πF, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp005, πE = πg.SliceType.Call(πF, πg.Args{πTemp006, πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.SetItem(πF, µsegments, πTemp005, πTemp004); πE != nil {
						continue
					}
					goto Label36
				Label36:
					// line 422: return urlunparse((scheme, netloc, '/'.join(segments),
					πF.SetLineno(422)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µscheme, "scheme"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µnetloc, "netloc"); πE != nil {
						continue
					}
					πTemp013 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µsegments, "segments"); πE != nil {
						continue
					}
					πTemp013[0] = µsegments
					if πTemp004, πE = πg.GetAttr(πF, πg.NewStr("/").ToObject(), ßjoin, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp013, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp013)
					if πE = πg.CheckLocal(πF, µparams, "params"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µquery, "query"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfragment, "fragment"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple6(µscheme, µnetloc, πTemp005, µparams, µquery, µfragment).ToObject()
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßurlunparse); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
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
			if πE = πF.Globals().SetItem(πF, ßurljoin.ToObject(), πTemp011); πE != nil {
				continue
			}
			// line 374: """Join a base URL and a possibly relative URL to form an absolute
			πF.SetLineno(374)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp012}, πg.NewStr("Join a base URL and a possibly relative URL to form an absolute\n    interpretation of the latter.").ToObject()); πE != nil {
				continue
			}
			if πTemp013, πE = πg.ResolveGlobal(πF, ßurljoin); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp013, ß__doc__, πTemp012); πE != nil {
				continue
			}
			// line 425: def urldefrag(url):
			πF.SetLineno(425)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "url", Def: nil}
			πTemp012 = πg.NewFunction(πg.NewCode("urldefrag", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µurl *πg.Object = πArgs[0]; _ = µurl
				var µs *πg.Object = πg.UnboundLocal; _ = µs
				var µn *πg.Object = πg.UnboundLocal; _ = µn
				var µp *πg.Object = πg.UnboundLocal; _ = µp
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µq *πg.Object = πg.UnboundLocal; _ = µq
				var µfrag *πg.Object = πg.UnboundLocal; _ = µfrag
				var µdefrag *πg.Object = πg.UnboundLocal; _ = µdefrag
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
					// line 426: """Removes any existing fragment from URL.
					πF.SetLineno(426)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Contains(πF, µurl, πg.NewStr("#").ToObject()); πE != nil {
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
					// line 432: if '#' in url:
					πF.SetLineno(432)
				Label1:
					// line 433: s, n, p, a, q, frag = urlparse(url)
					πF.SetLineno(433)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp003[0] = µurl
					if πTemp001, πE = πg.ResolveGlobal(πF, ßurlparse); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp001}, πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}, πg.TieTarget{Target: &πTemp008}, πg.TieTarget{Target: &πTemp009}}}, πTemp004); πE != nil {
						continue
					}
					µs = πTemp001
					µn = πTemp005
					µp = πTemp006
					µa = πTemp007
					µq = πTemp008
					µfrag = πTemp009
					// line 434: defrag = urlunparse((s, n, p, a, q, ''))
					πF.SetLineno(434)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µp, "p"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µq, "q"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple6(µs, µn, µp, µa, µq, ß.ToObject()).ToObject()
					πTemp003[0] = πTemp001
					if πTemp001, πE = πg.ResolveGlobal(πF, ßurlunparse); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					µdefrag = πTemp004
					// line 435: return defrag, frag
					πF.SetLineno(435)
					if πE = πg.CheckLocal(πF, µdefrag, "defrag"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µfrag, "frag"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µdefrag, µfrag).ToObject()
					πR = πTemp001
					continue
					goto Label3
				Label2:
					// line 437: return url, ''
					πF.SetLineno(437)
					if πE = πg.CheckLocal(πF, µurl, "url"); πE != nil {
						continue
					}
					πTemp001 = πg.NewTuple2(µurl, ß.ToObject()).ToObject()
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
			if πE = πF.Globals().SetItem(πF, ßurldefrag.ToObject(), πTemp012); πE != nil {
				continue
			}
			// line 426: """Removes any existing fragment from URL.
			πF.SetLineno(426)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp013}, πg.NewStr("Removes any existing fragment from URL.\n\n    Returns a tuple of the defragmented URL and the fragment.  If\n    the URL contained no fragments, the second element is the\n    empty string.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp014, πE = πg.ResolveGlobal(πF, ßurldefrag); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp014, ß__doc__, πTemp013); πE != nil {
				continue
			}
			// line 439: try:
			πF.SetLineno(439)
			πF.PushCheckpoint(2)
			// line 440: unicode
			πF.SetLineno(440)
			if πTemp013, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
				continue
			}
			πF.PopCheckpoint()
			// line 445: def _is_unicode(x):
			πF.SetLineno(445)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "x", Def: nil}
			πTemp013 = πg.NewFunction(πg.NewCode("_is_unicode", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
					// line 446: return isinstance(x, unicode)
					πF.SetLineno(446)
					πTemp001 = πF.MakeArgs(2)
					if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
						continue
					}
					πTemp001[0] = µx
					if πTemp002, πE = πg.ResolveGlobal(πF, ßunicode); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ß_is_unicode.ToObject(), πTemp013); πE != nil {
				continue
			}
			goto Label1
		Label2:
			if πE == nil {
			  continue
			}
			πE = nil
			πTemp015, πTemp016 = πF.ExcInfo()
			if πTemp014, πE = πg.ResolveGlobal(πF, ßNameError); πE != nil {
				continue
			}
			if πTemp017, πE = πg.IsInstance(πF, πTemp015.ToObject(), πTemp014); πE != nil {
				continue
			}
			if πTemp017 {
				goto Label3
			}
			πE = πF.Raise(πTemp015.ToObject(), nil, πTemp016.ToObject())
			continue
			// line 441: except NameError:
			πF.SetLineno(441)
		Label3:
			// line 442: def _is_unicode(x):
			πF.SetLineno(442)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "x", Def: nil}
			πTemp014 = πg.NewFunction(πg.NewCode("_is_unicode", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µx *πg.Object = πArgs[0]; _ = µx
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 443: return 0
					πF.SetLineno(443)
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
			if πE = πF.Globals().SetItem(πF, ß_is_unicode.ToObject(), πTemp014); πE != nil {
				continue
			}
			πF.RestoreExc(nil, nil)
			goto Label1
		Label1:
			// line 453: _hexdig = '0123456789ABCDEFabcdef'
			πF.SetLineno(453)
			if πE = πF.Globals().SetItem(πF, ß_hexdig.ToObject(), ß0123456789ABCDEFabcdef.ToObject()); πE != nil {
				continue
			}
			// line 454: _hextochr = dict((a+b, chr(int(a+b,16)))
			πF.SetLineno(454)
			πTemp002 = πF.MakeArgs(1)
			πTemp005 = make([]πg.Param, 0)
			πTemp018 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µa *πg.Object = πg.UnboundLocal; _ = µa
				var µb *πg.Object = πg.UnboundLocal; _ = µb
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
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Object
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
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
						if πTemp002, πE = πg.ResolveGlobal(πF, ß_hexdig); πE != nil {
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
							µa = πTemp002
						}
						if πE != nil || !πTemp004 {
							continue
						}
						πF.PushCheckpoint(1)            
						if πTemp005, πE = πg.ResolveGlobal(πF, ß_hexdig); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
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
						if πTemp005, πE = πg.Next(πF, πTemp002); πE != nil {
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
							µb = πTemp005
						}
						if πE != nil || !πTemp006 {
							continue
						}
						πF.PushCheckpoint(4)            
						// line 454: _hextochr = dict((a+b, chr(int(a+b,16)))
						πF.SetLineno(454)
						if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.Add(πF, µa, µb); πE != nil {
							continue
						}
						πTemp008 = πF.MakeArgs(1)
						πTemp009 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Add(πF, µa, µb); πE != nil {
							continue
						}
						πTemp009[0] = πTemp010
						πTemp009[1] = πg.NewInt(16).ToObject()
						if πTemp010, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
							continue
						}
						if πTemp011, πE = πTemp010.Call(πF, πTemp009, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp009)
						πTemp008[0] = πTemp011
						if πTemp010, πE = πg.ResolveGlobal(πF, ßchr); πE != nil {
							continue
						}
						if πTemp011, πE = πTemp010.Call(πF, πTemp008, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp008)
						πTemp005 = πg.NewTuple2(πTemp007, πTemp011).ToObject()
						πF.PushCheckpoint(7)
						return πTemp005, nil
					Label7:
						πTemp007 = πSent
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
			if πTemp019, πE = πTemp018.Call(πF, nil, nil); πE != nil {
				continue
			}
			πTemp002[0] = πTemp019
			if πTemp019, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
				continue
			}
			if πTemp020, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_hextochr.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 456: _asciire = re.compile('([\x00-\x7f]+)')
			πF.SetLineno(456)
			πTemp002 = πF.MakeArgs(1)
			πTemp002[0] = πg.NewStr("([\x00-\x7f]+)").ToObject()
			if πTemp019, πE = πg.ResolveGlobal(πF, ßre); πE != nil {
				continue
			}
			if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßcompile, nil); πE != nil {
				continue
			}
			if πTemp019, πE = πTemp020.Call(πF, πTemp002, nil); πE != nil {
				continue
			}
			πF.FreeArgs(πTemp002)
			if πE = πF.Globals().SetItem(πF, ß_asciire.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 458: def unquote(s):
			πF.SetLineno(458)
			πTemp005 = make([]πg.Param, 1)
			πTemp005[0] = πg.Param{Name: "s", Def: nil}
			πTemp019 = πg.NewFunction(πg.NewCode("unquote", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µs *πg.Object = πArgs[0]; _ = µs
				var µbits *πg.Object = πg.UnboundLocal; _ = µbits
				var µres *πg.Object = πg.UnboundLocal; _ = µres
				var µappend *πg.Object = πg.UnboundLocal; _ = µappend
				var µi *πg.Object = πg.UnboundLocal; _ = µi
				var µitem *πg.Object = πg.UnboundLocal; _ = µitem
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
				var πTemp006 *πg.Object
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πTemp008 []*πg.Object
				_ = πTemp008
				var πTemp009 []*πg.Object
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
					case 14: goto Label14
					case 10: goto Label10
					case 11: goto Label11
					case 5: goto Label5
					case 6: goto Label6
					default: panic("unexpected function state")
					}
					// line 459: """unquote('abc%20def') -> 'abc def'."""
					πF.SetLineno(459)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_is_unicode); πE != nil {
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
					// line 460: if _is_unicode(s):
					πF.SetLineno(460)
				Label1:
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Contains(πF, µs, πg.NewStr("%").ToObject()); πE != nil {
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
					// line 461: if '%' not in s:
					πF.SetLineno(461)
				Label3:
					// line 462: return s
					πF.SetLineno(462)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πR = µs
					continue
					goto Label4
				Label4:
					// line 463: bits = _asciire.split(s)
					πF.SetLineno(463)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πTemp001[0] = µs
					if πTemp002, πE = πg.ResolveGlobal(πF, ß_asciire); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µbits = πTemp002
					// line 464: res = [bits[0]]
					πF.SetLineno(464)
					πTemp001 = make([]*πg.Object, 1)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µbits, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µres = πTemp002
					// line 465: append = res.append
					πF.SetLineno(465)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µres, ßappend, nil); πE != nil {
						continue
					}
					µappend = πTemp002
					πTemp001 = πF.MakeArgs(3)
					πTemp001[0] = πg.NewInt(1).ToObject()
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					πTemp005[0] = µbits
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[1] = πTemp006
					πTemp001[2] = πg.NewInt(2).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
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
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µi = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(5)            
					// line 467: append(unquote(str(bits[i])).decode('latin1'))
					πF.SetLineno(467)
					πTemp001 = πF.MakeArgs(1)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ßlatin1.ToObject()
					πTemp008 = πF.MakeArgs(1)
					πTemp009 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					πTemp003 = µi
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µbits, πTemp003); πE != nil {
						continue
					}
					πTemp009[0] = πTemp006
					if πTemp003, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp009, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp009)
					πTemp008[0] = πTemp006
					if πTemp003, πE = πg.ResolveGlobal(πF, ßunquote); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp008, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp008)
					if πTemp003, πE = πg.GetAttr(πF, πTemp006, ßdecode, nil); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp003, πE = µappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 468: append(bits[i + 1])
					πF.SetLineno(468)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.Add(πF, µi, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp006
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µbits, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp003, πE = µappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					continue
				Label6:
					if πE != nil || πR != nil {
						continue
					}
				Label7:
					// line 469: return ''.join(res)
					πF.SetLineno(469)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					πTemp001[0] = µres
					if πTemp002, πE = πg.GetAttr(πF, ß.ToObject(), ßjoin, nil); πE != nil {
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
					// line 471: bits = s.split('%')
					πF.SetLineno(471)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("%").ToObject()
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µs, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					µbits = πTemp003
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					πTemp001[0] = µbits
					if πTemp003, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					if πTemp002, πE = πg.Eq(πF, πTemp006, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
						continue
					}
					if πTemp004 {
						goto Label8
					}
					goto Label9
					// line 473: if len(bits) == 1:
					πF.SetLineno(473)
				Label8:
					// line 474: return s
					πF.SetLineno(474)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					πR = µs
					continue
					goto Label9
				Label9:
					// line 475: res = [bits[0]]
					πF.SetLineno(475)
					πTemp001 = make([]*πg.Object, 1)
					πTemp002 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetItem(πF, µbits, πTemp002); πE != nil {
						continue
					}
					πTemp001[0] = πTemp003
					πTemp002 = πg.NewList(πTemp001...).ToObject()
					µres = πTemp002
					// line 476: append = res.append
					πF.SetLineno(476)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, µres, ßappend, nil); πE != nil {
						continue
					}
					µappend = πTemp002
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µbits, "bits"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µbits, πTemp003); πE != nil {
						continue
					}
					if πTemp002, πE = πg.Iter(πF, πTemp006); πE != nil {
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
						µitem = πTemp003
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(10)            
					// line 478: try:
					πF.SetLineno(478)
					πF.PushCheckpoint(14)
					// line 479: append(_hextochr[item[:2]])
					πF.SetLineno(479)
					πTemp001 = πF.MakeArgs(1)
					if πTemp006, πE = πg.SliceType.Call(πF, πg.Args{πg.None, πg.NewInt(2).ToObject(), πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µitem, πTemp006); πE != nil {
						continue
					}
					πTemp003 = πTemp010
					if πTemp010, πE = πg.ResolveGlobal(πF, ß_hextochr); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, πTemp010, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp003, πE = µappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 480: append(item[2:])
					πF.SetLineno(480)
					πTemp001 = πF.MakeArgs(1)
					if πTemp003, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(2).ToObject(), πg.None, πg.None}, nil); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					if πTemp006, πE = πg.GetItem(πF, µitem, πTemp003); πE != nil {
						continue
					}
					πTemp001[0] = πTemp006
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp003, πE = µappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.PopCheckpoint()
					goto Label13
				Label14:
					if πE == nil {
					  continue
					}
					πE = nil
					πTemp011, πTemp012 = πF.ExcInfo()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßKeyError); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsInstance(πF, πTemp011.ToObject(), πTemp003); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label15
					}
					πE = πF.Raise(πTemp011.ToObject(), nil, πTemp012.ToObject())
					continue
					// line 481: except KeyError:
					πF.SetLineno(481)
				Label15:
					// line 482: append('%')
					πF.SetLineno(482)
					πTemp001 = πF.MakeArgs(1)
					πTemp001[0] = πg.NewStr("%").ToObject()
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp003, πE = µappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					// line 483: append(item)
					πF.SetLineno(483)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µitem, "item"); πE != nil {
						continue
					}
					πTemp001[0] = µitem
					if πE = πg.CheckLocal(πF, µappend, "append"); πE != nil {
						continue
					}
					if πTemp003, πE = µappend.Call(πF, πTemp001, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp001)
					πF.RestoreExc(nil, nil)
					goto Label13
				Label13:
					continue
				Label11:
					if πE != nil || πR != nil {
						continue
					}
				Label12:
					// line 484: return ''.join(res)
					πF.SetLineno(484)
					πTemp001 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
						continue
					}
					πTemp001[0] = µres
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
			if πE = πF.Globals().SetItem(πF, ßunquote.ToObject(), πTemp019); πE != nil {
				continue
			}
			// line 459: """unquote('abc%20def') -> 'abc def'."""
			πF.SetLineno(459)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πg.NewStr("unquote('abc%20def') -> 'abc def'.").ToObject()); πE != nil {
				continue
			}
			if πTemp021, πE = πg.ResolveGlobal(πF, ßunquote); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp021, ß__doc__, πTemp020); πE != nil {
				continue
			}
			// line 486: def parse_qs(qs, keep_blank_values=0, strict_parsing=0):
			πF.SetLineno(486)
			πTemp005 = make([]πg.Param, 3)
			πTemp005[0] = πg.Param{Name: "qs", Def: nil}
			πTemp005[1] = πg.Param{Name: "keep_blank_values", Def: πg.NewInt(0).ToObject()}
			πTemp005[2] = πg.Param{Name: "strict_parsing", Def: πg.NewInt(0).ToObject()}
			πTemp020 = πg.NewFunction(πg.NewCode("parse_qs", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µqs *πg.Object = πArgs[0]; _ = µqs
				var µkeep_blank_values *πg.Object = πArgs[1]; _ = µkeep_blank_values
				var µstrict_parsing *πg.Object = πArgs[2]; _ = µstrict_parsing
				var µdict *πg.Object = πg.UnboundLocal; _ = µdict
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
				var πTemp001 *πg.Dict
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
					// line 487: """Parse a query given as a string argument.
					πF.SetLineno(487)
					// line 504: dict = {}
					πF.SetLineno(504)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					µdict = πTemp002
					πTemp003 = πF.MakeArgs(3)
					if πE = πg.CheckLocal(πF, µqs, "qs"); πE != nil {
						continue
					}
					πTemp003[0] = µqs
					if πE = πg.CheckLocal(πF, µkeep_blank_values, "keep_blank_values"); πE != nil {
						continue
					}
					πTemp003[1] = µkeep_blank_values
					if πE = πg.CheckLocal(πF, µstrict_parsing, "strict_parsing"); πE != nil {
						continue
					}
					πTemp003[2] = µstrict_parsing
					if πTemp004, πE = πg.ResolveGlobal(πF, ßparse_qsl); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
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
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
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
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp005}, πg.TieTarget{Target: &πTemp008}}}, πTemp004); πE != nil {
							continue
						}
						µname = πTemp005
						µvalue = πTemp008
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.Contains(πF, µdict, µname); πE != nil {
						continue
					}
					πTemp004 = πg.GetBool(πTemp007).ToObject()
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label4
					}
					goto Label5
					// line 506: if name in dict:
					πF.SetLineno(506)
				Label4:
					// line 507: dict[name].append(value)
					πF.SetLineno(507)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp003[0] = µvalue
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp004 = µname
					if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µdict, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßappend, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					goto Label6
				Label5:
					// line 509: dict[name] = [value]
					πF.SetLineno(509)
					πTemp003 = make([]*πg.Object, 1)
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp003[0] = µvalue
					πTemp004 = πg.NewList(πTemp003...).ToObject()
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp005}, πTemp004); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					πTemp008 = µname
					if πE = πg.SetItem(πF, µdict, πTemp008, πTemp005); πE != nil {
						continue
					}
					goto Label6
				Label6:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 510: return dict
					πF.SetLineno(510)
					if πE = πg.CheckLocal(πF, µdict, "dict"); πE != nil {
						continue
					}
					πR = µdict
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßparse_qs.ToObject(), πTemp020); πE != nil {
				continue
			}
			// line 487: """Parse a query given as a string argument.
			πF.SetLineno(487)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp021}, πg.NewStr("Parse a query given as a string argument.\n\n        Arguments:\n\n        qs: percent-encoded query string to be parsed\n\n        keep_blank_values: flag indicating whether blank values in\n            percent-encoded queries should be treated as blank strings.\n            A true value indicates that blanks should be retained as\n            blank strings.  The default false value indicates that\n            blank values are to be ignored and treated as if they were\n            not included.\n\n        strict_parsing: flag indicating what to do with parsing errors.\n            If false (the default), errors are silently ignored.\n            If true, errors raise a ValueError exception.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp022, πE = πg.ResolveGlobal(πF, ßparse_qs); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp022, ß__doc__, πTemp021); πE != nil {
				continue
			}
			// line 512: def parse_qsl(qs, keep_blank_values=0, strict_parsing=0):
			πF.SetLineno(512)
			πTemp005 = make([]πg.Param, 3)
			πTemp005[0] = πg.Param{Name: "qs", Def: nil}
			πTemp005[1] = πg.Param{Name: "keep_blank_values", Def: πg.NewInt(0).ToObject()}
			πTemp005[2] = πg.Param{Name: "strict_parsing", Def: πg.NewInt(0).ToObject()}
			πTemp021 = πg.NewFunction(πg.NewCode("parse_qsl", "/usr/lib/python2.7/urlparse.py", πTemp005, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µqs *πg.Object = πArgs[0]; _ = µqs
				var µkeep_blank_values *πg.Object = πArgs[1]; _ = µkeep_blank_values
				var µstrict_parsing *πg.Object = πArgs[2]; _ = µstrict_parsing
				var µpairs *πg.Object = πg.UnboundLocal; _ = µpairs
				var µr *πg.Object = πg.UnboundLocal; _ = µr
				var µname_value *πg.Object = πg.UnboundLocal; _ = µname_value
				var µnv *πg.Object = πg.UnboundLocal; _ = µnv
				var µname *πg.Object = πg.UnboundLocal; _ = µname
				var µvalue *πg.Object = πg.UnboundLocal; _ = µvalue
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
				var πTemp011 []*πg.Object
				_ = πTemp011
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 513: """Parse a query given as a string argument.
					πF.SetLineno(513)
					// line 531: pairs = [s2 for s1 in qs.split('&') for s2 in s1.split(';')]
					πF.SetLineno(531)
					πTemp003 = make([]πg.Param, 0)
					πTemp002 = πg.NewFunction(πg.NewCode("<generator>", "/usr/lib/python2.7/urlparse.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µs1 *πg.Object = πg.UnboundLocal; _ = µs1
						var µs2 *πg.Object = πg.UnboundLocal; _ = µs2
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
						var πTemp008 bool
						_ = πTemp008
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
								πTemp002 = πF.MakeArgs(1)
								πTemp002[0] = πg.NewStr("&").ToObject()
								if πE = πg.CheckLocal(πF, µqs, "qs"); πE != nil {
									continue
								}
								if πTemp003, πE = πg.GetAttr(πF, µqs, ßsplit, nil); πE != nil {
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
									µs1 = πTemp003
								}
								if πE != nil || !πTemp006 {
									continue
								}
								πF.PushCheckpoint(1)            
								πTemp002 = πF.MakeArgs(1)
								πTemp002[0] = πg.NewStr(";").ToObject()
								if πE = πg.CheckLocal(πF, µs1, "s1"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µs1, ßsplit, nil); πE != nil {
									continue
								}
								if πTemp007, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp002)
								if πTemp003, πE = πg.Iter(πF, πTemp007); πE != nil {
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
									µs2 = πTemp004
								}
								if πE != nil || !πTemp008 {
									continue
								}
								πF.PushCheckpoint(4)            
								// line 531: pairs = [s2 for s1 in qs.split('&') for s2 in s1.split(';')]
								πF.SetLineno(531)
								if πE = πg.CheckLocal(πF, µs2, "s2"); πE != nil {
									continue
								}
								πF.PushCheckpoint(7)
								return µs2, nil
							Label7:
								πTemp004 = πSent
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
					if πTemp004, πE = πTemp002.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.ListType.Call(πF, πg.Args{πTemp004}, nil); πE != nil {
						continue
					}
					µpairs = πTemp001
					// line 532: r = []
					πF.SetLineno(532)
					πTemp005 = make([]*πg.Object, 0)
					πTemp001 = πg.NewList(πTemp005...).ToObject()
					µr = πTemp001
					if πE = πg.CheckLocal(πF, µpairs, "pairs"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, µpairs); πE != nil {
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
						µname_value = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					if πE = πg.CheckLocal(πF, µname_value, "name_value"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µname_value); πE != nil {
						continue
					}
					πTemp008 = πg.GetBool(!πTemp009).ToObject()
					πTemp004 = πTemp008
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if !πTemp007 {
						goto Label4
					}
					if πE = πg.CheckLocal(πF, µstrict_parsing, "strict_parsing"); πE != nil {
						continue
					}
					if πTemp009, πE = πg.IsTrue(πF, µstrict_parsing); πE != nil {
						continue
					}
					πTemp008 = πg.GetBool(!πTemp009).ToObject()
					πTemp004 = πTemp008
				Label4:
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label5
					}
					goto Label6
					// line 534: if not name_value and not strict_parsing:
					πF.SetLineno(534)
				Label5:
					// line 535: continue
					πF.SetLineno(535)
					continue
					goto Label6
				Label6:
					// line 536: nv = name_value.split('=', 1)
					πF.SetLineno(536)
					πTemp005 = πF.MakeArgs(2)
					πTemp005[0] = πg.NewStr("=").ToObject()
					πTemp005[1] = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µname_value, "name_value"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µname_value, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µnv = πTemp008
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µnv, "nv"); πE != nil {
						continue
					}
					πTemp005[0] = µnv
					if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					if πTemp004, πE = πg.NE(πF, πTemp010, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label7
					}
					goto Label8
					// line 537: if len(nv) != 2:
					πF.SetLineno(537)
				Label7:
					if πE = πg.CheckLocal(πF, µstrict_parsing, "strict_parsing"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µstrict_parsing); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label9
					}
					goto Label10
					// line 538: if strict_parsing:
					πF.SetLineno(538)
				Label9:
					if πTemp004, πE = πg.ResolveGlobal(πF, ßValueError); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µname_value, "name_value"); πE != nil {
						continue
					}
					πTemp010 = πg.NewTuple1(µname_value).ToObject()
					if πTemp008, πE = πg.Mod(πF, πg.NewStr("bad query field: %r").ToObject(), πTemp010); πE != nil {
						continue
					}
					// line 539: raise ValueError, "bad query field: %r" % (name_value,)
					πF.SetLineno(539)
					πE = πF.Raise(πTemp004, πTemp008, nil)
					continue
					goto Label10
				Label10:
					if πE = πg.CheckLocal(πF, µkeep_blank_values, "keep_blank_values"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.IsTrue(πF, µkeep_blank_values); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label11
					}
					goto Label12
					// line 541: if keep_blank_values:
					πF.SetLineno(541)
				Label11:
					// line 542: nv.append('')
					πF.SetLineno(542)
					πTemp005 = πF.MakeArgs(1)
					πTemp005[0] = ß.ToObject()
					if πE = πg.CheckLocal(πF, µnv, "nv"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µnv, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label13
				Label12:
					// line 544: continue
					πF.SetLineno(544)
					continue
					goto Label13
				Label13:
					goto Label8
				Label8:
					πTemp005 = πF.MakeArgs(1)
					πTemp008 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µnv, "nv"); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetItem(πF, µnv, πTemp008); πE != nil {
						continue
					}
					πTemp005[0] = πTemp010
					if πTemp008, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp010, πE = πTemp008.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					πTemp004 = πTemp010
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label14
					}
					if πE = πg.CheckLocal(πF, µkeep_blank_values, "keep_blank_values"); πE != nil {
						continue
					}
					πTemp004 = µkeep_blank_values
				Label14:
					if πTemp007, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp007 {
						goto Label15
					}
					goto Label16
					// line 545: if len(nv[1]) or keep_blank_values:
					πF.SetLineno(545)
				Label15:
					// line 546: name = unquote(nv[0].replace('+', ' '))
					πF.SetLineno(546)
					πTemp005 = πF.MakeArgs(1)
					πTemp011 = πF.MakeArgs(2)
					πTemp011[0] = πg.NewStr("+").ToObject()
					πTemp011[1] = πg.NewStr(" ").ToObject()
					πTemp004 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µnv, "nv"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µnv, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp008, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp005[0] = πTemp008
					if πTemp004, πE = πg.ResolveGlobal(πF, ßunquote); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µname = πTemp008
					// line 547: value = unquote(nv[1].replace('+', ' '))
					πF.SetLineno(547)
					πTemp005 = πF.MakeArgs(1)
					πTemp011 = πF.MakeArgs(2)
					πTemp011[0] = πg.NewStr("+").ToObject()
					πTemp011[1] = πg.NewStr(" ").ToObject()
					πTemp004 = πg.NewInt(1).ToObject()
					if πE = πg.CheckLocal(πF, µnv, "nv"); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetItem(πF, µnv, πTemp004); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp008, ßreplace, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp011, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp011)
					πTemp005[0] = πTemp008
					if πTemp004, πE = πg.ResolveGlobal(πF, ßunquote); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					µvalue = πTemp008
					// line 548: r.append((name, value))
					πF.SetLineno(548)
					πTemp005 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µvalue, "value"); πE != nil {
						continue
					}
					πTemp004 = πg.NewTuple2(µname, µvalue).ToObject()
					πTemp005[0] = πTemp004
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µr, ßappend, nil); πE != nil {
						continue
					}
					if πTemp008, πE = πTemp004.Call(πF, πTemp005, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp005)
					goto Label16
				Label16:
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
					// line 550: return r
					πF.SetLineno(550)
					if πE = πg.CheckLocal(πF, µr, "r"); πE != nil {
						continue
					}
					πR = µr
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßparse_qsl.ToObject(), πTemp021); πE != nil {
				continue
			}
			// line 513: """Parse a query given as a string argument.
			πF.SetLineno(513)
			if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp022}, πg.NewStr("Parse a query given as a string argument.\n\n    Arguments:\n\n    qs: percent-encoded query string to be parsed\n\n    keep_blank_values: flag indicating whether blank values in\n        percent-encoded queries should be treated as blank strings.  A\n        true value indicates that blanks should be retained as blank\n        strings.  The default false value indicates that blank values\n        are to be ignored and treated as if they were  not included.\n\n    strict_parsing: flag indicating what to do with parsing errors. If\n        false (the default), errors are silently ignored. If true,\n        errors raise a ValueError exception.\n\n    Returns a list, as G-d intended.\n    ").ToObject()); πE != nil {
				continue
			}
			if πTemp023, πE = πg.ResolveGlobal(πF, ßparse_qsl); πE != nil {
				continue
			}
			if πE = πg.SetAttr(πF, πTemp023, ß__doc__, πTemp022); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("urlparse", Code)
}

